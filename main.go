package main

import (
	"github.com/rodrigodmd/gogen/pkg/cli"
)

func main() {
	// ConfigRepo
	//promptRepo := prompt.NewRepository()
	//storageRepo := storage.NewConfigRepo()
	//statusRepo := storage.NewStatusRepo()
	//httpRepo := httpclient.New()

	// Services
	//confSrv := config.NewService(promptRepo, storageRepo)
	//statusSrv := status.New(statusRepo, promptRepo)

	//procSrv := processor.New(confSrv, statusSrv, promptRepo, httpRepo)
	//.NewService(promptRepo, storageRepo)

	command := cli.NewCommandHandler()
	command.AddRoutes()
	command.Run()
}
