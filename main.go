package main

import (
	"net/http"
	_ "programmertio/golang-restful-api/app"
	_ "programmertio/golang-restful-api/controller"
	"programmertio/golang-restful-api/helper"
	"programmertio/golang-restful-api/middleware"
	_ "programmertio/golang-restful-api/repository"
	_ "programmertio/golang-restful-api/service"

	_ "github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
