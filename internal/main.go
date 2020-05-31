package main

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/thunpin/go-study-swagger/pkg/swagger/server/restapi"
	"github.com/thunpin/go-study-swagger/pkg/swagger/server/restapi/operations"
)

func main() {
	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewHelloAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	server.Port = 8080
	log.Printf("Listening on port %d", server.Port)

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)
	api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(GetHelloUser)
	api.GetOopsUserHandler = operations.GetOopsUserHandlerFunc(GetOopsUser)

	// Start listening using having the handlers and port
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

//Health route returns OK
func Health(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")
}

//GetHelloUser returns Hello + your name
func GetHelloUser(params operations.GetHelloUserParams) middleware.Responder {
	return operations.NewGetHelloUserOK().WithPayload("Hello " + params.User + "!")
}

//GetOopsUser returns Hello + your name
func GetOopsUser(params operations.GetOopsUserParams) middleware.Responder {
	return operations.NewGetOopsUserOK().WithPayload("Oops " + params.User + "!")
}
