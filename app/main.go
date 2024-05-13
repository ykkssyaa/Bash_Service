package main

import (
	"context"
	"errors"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	config "github.com/ykkssyaa/Bash_Service/internal/configs"
	"github.com/ykkssyaa/Bash_Service/internal/consts"
	"github.com/ykkssyaa/Bash_Service/internal/gateway"
	"github.com/ykkssyaa/Bash_Service/internal/server"
	"github.com/ykkssyaa/Bash_Service/internal/service"
	lg "github.com/ykkssyaa/Bash_Service/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := lg.InitLogger()

	logger.Info.Print("Executing InitConfig.")
	if err := config.InitConfig(); err != nil {
		logger.Err.Fatalf(err.Error())
	}

	logger.Info.Print("Connecting to Postgres.")
	db, err := gateway.NewPostgresDB(viper.GetString("POSTGRES_STRING"))

	if err != nil {
		logger.Err.Fatalf(err.Error())
	}

	storageSize := consts.CmdPullSize

	logger.Info.Print("Creating Gateways.")
	gateways := gateway.NewGateway(db, storageSize)

	logger.Info.Print("Creating Services.")
	services := service.NewService(gateways, logger)

	logger.Info.Print("Creating server.")

	port := viper.GetString("PORT")
	srv := server.NewHttpServer(":"+port, services)

	logger.Info.Print("Starting the server on port: " + port + "\n\n")

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Err.Fatalf("error occured while running http server: \"%s\" \n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info.Println("Server Shutting Down.")

	if err = srv.Shutdown(context.Background()); err != nil {
		logger.Err.Fatalf("error occured while server shutting down: \"%s\" \n", err.Error())
	}

	logger.Info.Println("DB connection closing.")
	if err := db.Close(); err != nil {
		logger.Err.Fatalf("error occured on db connection close: \"%s\" \n", err.Error())
	}
}
