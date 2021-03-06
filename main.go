package main

import (
	"flag"
	"fmt"

	"github.com/bcskill/eth-main-side-swap/admin"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/bcskill/eth-main-side-swap/executor"
	"github.com/bcskill/eth-main-side-swap/model"
	"github.com/bcskill/eth-main-side-swap/observer"
	"github.com/bcskill/eth-main-side-swap/swap"
	"github.com/bcskill/eth-main-side-swap/util"
)

const (
	flagConfigType         = "config-type"
	flagConfigAwsRegion    = "aws-region"
	flagConfigAwsSecretKey = "aws-secret-key"
	flagConfigPath         = "config-path"
)

const (
	ConfigTypeLocal = "local"
	ConfigTypeAws   = "aws"
)

func initFlags() {
	flag.String(flagConfigPath, "", "config path")
	flag.String(flagConfigType, "", "config type, local or aws")
	flag.String(flagConfigAwsRegion, "", "aws s3 region")
	flag.String(flagConfigAwsSecretKey, "", "aws s3 secret key")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(fmt.Sprintf("bind flags error, err=%s", err))
	}
}

func printUsage() {
	fmt.Print("usage: ./swap --config-type [local or aws] --config-path config_file_path\n")
}

func main() {
	initFlags()

	configType := viper.GetString(flagConfigType)
	if configType == "" {
		printUsage()
		return
	}

	if configType != ConfigTypeAws && configType != ConfigTypeLocal {
		printUsage()
		return
	}

	var config *util.Config
	if configType == ConfigTypeAws {
		awsSecretKey := viper.GetString(flagConfigAwsSecretKey)
		if awsSecretKey == "" {
			printUsage()
			return
		}

		awsRegion := viper.GetString(flagConfigAwsRegion)
		if awsRegion == "" {
			printUsage()
			return
		}

		configContent, err := util.GetSecret(awsSecretKey, awsRegion)
		if err != nil {
			fmt.Printf("get aws config error, err=%s", err.Error())
			return
		}
		config = util.ParseConfigFromJson(configContent)
	} else {
		configFilePath := viper.GetString(flagConfigPath)
		if configFilePath == "" {
			printUsage()
			return
		}
		config = util.ParseConfigFromFile(configFilePath)
	}
	config.Validate()

	// init logger
	util.InitLogger(config.LogConfig)
	util.InitTgAlerter(config.AlertConfig)

	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.DBPath)
	if err != nil {
		panic(fmt.Sprintf("open db error, err=%s", err.Error()))
	}
	defer db.Close()
	model.InitTables(db)

	sideClient, err := ethclient.Dial(config.ChainConfig.SideProvider)
	if err != nil {
		panic("new side client error")
	}

	mainClient, err := ethclient.Dial(config.ChainConfig.MainProvider)
	if err != nil {
		panic("new main client error")
	}

	sideExecutor := executor.NewSideExecutor(sideClient, config.ChainConfig.SideSwapAgentAddr, config)
	sideObserver := observer.NewObserver(db, config.ChainConfig.SideStartHeight, config.ChainConfig.SideConfirmNum, config, sideExecutor)
	sideObserver.Start()

	mainExecutor := executor.NewMainExecutor(mainClient, config.ChainConfig.MainSwapAgentAddr, config)
	mainObserver := observer.NewObserver(db, config.ChainConfig.MainStartHeight, config.ChainConfig.MainConfirmNum, config, mainExecutor)
	mainObserver.Start()

	swapEngine, err := swap.NewSwapEngine(db, config, sideClient, mainClient)
	if err != nil {
		panic(fmt.Sprintf("create swap engine error, err=%s", err.Error()))
	}

	swapEngine.Start()

	swapPairEngine, err := swap.NewSwapPairEngine(db, config, sideClient, swapEngine)
	if err != nil {
		panic(fmt.Sprintf("create swap pair engine error, err=%s", err.Error()))
	}
	swapPairEngine.Start()

	signer, err := util.NewHmacSignerFromConfig(config)
	if err != nil {
		panic(fmt.Sprintf("new hmac singer error, err=%s", err.Error()))
	}
	admin := admin.NewAdmin(config, db, signer, swapEngine)
	go admin.Serve()

	select {}
}
