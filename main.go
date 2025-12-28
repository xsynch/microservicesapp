package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"
	"net/http"
	
	"github.com/go-openapi/runtime/middleware"
	
	"github.com/gorilla/mux"
	"github.com/xsynch/microserviceapp/handlers"
)

func main(){
	fmt.Println("This is a new start")
	l := log.New(os.Stdout,"product-api: ",log.LstdFlags)
	
	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)
	
	
	// sm := http.NewServeMux()
	sm := mux.NewRouter()
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/",ph.GetProducts)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}",ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)
	
	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/",ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/",ph.DeleteProduct)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts,nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	
	// sm.Handle("/",ph)
	// sm.Handle("/goodbye",gh)

	


	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,

	}
	go func(){
		err := s.ListenAndServe()	
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal,1)
	signal.Notify(sigChan,syscall.SIGINT)
	signal.Notify(sigChan,syscall.SIGTERM)

	sig := <- sigChan
	l.Println("Received Terminate, shutting down gradefully",sig)
	
	
	tc,_ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	
}