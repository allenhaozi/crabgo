package app

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/allenhaozi/crabgo/cmd/app/options"
	"github.com/allenhaozi/crabgo/pkg/register"
	"github.com/allenhaozi/crabgo/pkg/restapi"
	"github.com/allenhaozi/crabgo/pkg/server"
	"github.com/allenhaozi/crabgo/pkg/utils"
)

func NewIAMCommand() *cobra.Command {

	opts, err := options.NewOptions()
	//init log
	if err != nil {
		fmt.Printf("parse option occur error, error:%s", err.Error())
		os.Exit(1)
	}
	cmd := &cobra.Command{
		Use:  "crabgo",
		Long: "crabgo project",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCommand(cmd, args, opts); err != nil {
				fmt.Printf("runCommand occur err, msg:%s", err.Error())
				os.Exit(1)
			}
		},
	}
	return cmd
}

func runCommand(cmd *cobra.Command, args []string, opts *options.Options) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return Run(ctx, opts.Config)
}

func Run(ctx context.Context, cfg *register.Config) error {
	// initial log
	InitLog(cfg)

	utils.InstallgoroutineDumpGenerator()

	if err := buildServiceChain(ctx, cfg); err != nil {
		return err
	}
	return StartWebServer(ctx, cfg)
}

func buildServiceChain(ctx context.Context, cfg *register.Config) error {
	// register image
	var err error
	return err
}

func StartWebServer(ctx context.Context, cfg *register.Config) error {
	var err error
	//start restful api web server
	webServer := server.NewWebServer()

	restfulApiList := restapi.Setup(cfg)

	log.Info("starts registering restful apis")

	webServer.RegisterHttpHandler(restfulApiList)

	log.Info("restful apis are registered")

	err = webServer.StartHttpServer(cfg)
	return err
}

//func InitLog(cfg *register.Config) {
func InitLog(cfg *register.Config) {

	log.SetLevel(cfg.GeneralConfig.CrabConfig.LogLevel)

	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
