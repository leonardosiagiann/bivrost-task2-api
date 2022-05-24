package main

import (
	"bivrost-task2/controllers"
	"bivrost-task2/database"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/koinworks/asgard-bivrost/libs"
	hmodels "github.com/koinworks/asgard-heimdal/models"
)

func init() {
	godotenv.Load()
	database.StartDB()
}

func main() {

	hostname, _ := os.Hostname()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	portNumber, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	serviceConfig := &hmodels.Service{
		Class:     "product-service",
		Key:       os.Getenv("APP_KEY"),
		Name:      os.Getenv("APP_NAME"),
		Version:   os.Getenv("APP_VERSION"),
		Host:      hostname,
		Port:      portNumber,
		Namespace: os.Getenv("K8S_NAMESPACE"),
		Metas:     make(hmodels.ServiceMetas),
	}

	registry, err := libs.InitRegistry(libs.RegistryConfig{
		Address:  os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Service:  serviceConfig,
	})

	if err != nil {
		log.Fatal(err)
	}

	server, err := libs.NewServer(registry)
	if err != nil {
		log.Fatal(err)
	}

	bivrostSvc := server.AsGatewayService(
		"/v1",
	)

	bivrostSvc.Get("/", controllers.PingHandler)
	bivrostSvc.Post("/createitem", controllers.CreateItem)
	bivrostSvc.Get("/list", controllers.GetItems)
	bivrostSvc.Post("/createorder", controllers.CreateOrder)
	bivrostSvc.Get("/orders", controllers.GetOrders)

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
