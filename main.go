package main

import (
	"ResourceServer/server"
	"ResourceServer/service"
	"flag"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"

	_ "ResourceServer/docs"
)

// @title Resource Server
// @version 0.2
// @description Server for downloading images from a URL in a local folder. With the further ability to receive these images by a special assigned name.

// @contact.name Melenium2
// @contact.email osumelenium@gmail.com

// @BasePath /
func main() {
	var (
		resourcePath      string
		parallelUploading int
		port              string
	)

	flag.StringVar(&resourcePath, "path", server.DefaultSaveFolder, "path to resource folder in filesystem")
	flag.IntVar(&parallelUploading, "parallel", server.DefaultWorkers, "number of workers for concurrency uploading")
	flag.StringVar(&port, "port", server.DefaultPort, "listening port")
	flag.Parse()

	app := fiber.New()
	config := server.NewConfig()
	{
		config.ServeFolder = resourcePath
		config.ParallelWorkers = parallelUploading
		config.Port = port
	}
	workService := service.New(config.ServeFolder, parallelUploading)
	serv := server.New(app, workService, config)
	_ = serv.InitRoutes()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		_ = <-c
		_ = serv.Shutdown()
	}()

	if err := serv.Listen(); err != nil {
		log.Fatal(err)
	}

	log.Println("Shutdown")
}
