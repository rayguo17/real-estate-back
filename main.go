package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goBack/models"
	"github.com/goBack/routers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	models.Init()
	gin.SetMode("test")
	api := routers.InitRouter()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: api,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	fmt.Println(http.StatusOK)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	s := <-quit

	fmt.Println("signal:", s)
	log.Println("shutdown server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	log.Println("Server exiting")

}
