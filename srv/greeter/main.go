package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/service/grpc"

	log "github.com/sirupsen/logrus"

	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	logger "github.com/xmlking/micro-starter-kit/shared/log"
	logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
	transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
	"github.com/xmlking/micro-starter-kit/srv/greeter/handler"
	greeterPB "github.com/xmlking/micro-starter-kit/srv/greeter/proto/greeter"
)

const (
	serviceName = "greetersrv"
)

var (
	configDir  string
	configFile string
	cfg        myConfig.ServiceConfiguration
)

func main() {
	// New Service
	service := grpc.NewService(
		// optional cli flag to override config.
		// comment out if you don't need to override any base config via CLI
		micro.Flags(
			cli.StringFlag{
				Name:        "configDir, d",
				Value:       "/config",
				Usage:       "Path to the config directory. Defaults to 'config'",
				EnvVar:      "CONFIG_DIR",
				Destination: &configDir,
			},
			cli.StringFlag{
				Name:        "configFile, f",
				Value:       "config.yaml",
				Usage:       "Config file in configDir. Defaults to 'config.yaml'",
				EnvVar:      "CONFIG_FILE",
				Destination: &configFile,
			}),
		micro.Name(serviceName),
		micro.Version(myConfig.Version),
	)

	// Initialize service
	service.Init(
		micro.Action(func(c *cli.Context) {
			// load config
			myConfig.InitConfig(configDir, configFile)
			_ = config.Scan(&cfg)
			logger.InitLogger(cfg.Log)
		}),
	)

	// Initialize Features
	var options []micro.Option
	// Wrappers are invoked in the order as they added
	if cfg.Features["reqlogs"].Enabled {
		options = append(options, micro.WrapHandler(logWrapper.NewHandlerWrapper()))
	}
	if cfg.Features["translogs"].Enabled {
		topic := config.Get("features", "translogs", "topic").String("recordersrv")
		publisher := micro.NewPublisher(topic, service.Client())
		options = append(options, micro.WrapHandler(transWrapper.NewHandlerWrapper(publisher)))
	}

	// Initialize Features
	service.Init(
		options...,
	)

	// Register Handler
	_ = greeterPB.RegisterGreeterHandler(service.Server(), handler.NewGreeterHandler())
	myConfig.PrintBuildInfo()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
