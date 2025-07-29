package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// How to make self-signed Certificate (but browser will warn client with untrusted connection!!!)

// install openssl

// openssl req -x509 -nodes -days 365 \
//  -newkey rsa:2048 \
//  -keyout localhost.key \
//  -out localhost.crt \
//  -subj "/CN=localhost"

// this will create a self signed certs localhost.key and  localhost.crt
// if needed you need to move these files to project folder

// then use them with srv.ListenAndServerTLS("localhost.crt ", "localhost.key")

// NExt this is implementation of local trusted Cert with mkcert
func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write([]byte("Hello from server!!! \n You can trust me !!! OR NO XEXE"))
	})

	srv := http.Server{
		Addr:         ":8082",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  130 * time.Second,
	}

	/*
		cert, key := "localhost.crt", "localhost.key"
	*/

	// install choco
	// with choco install mkcert that is Locally CA (like GoDaddy)
	// with mkcert -install

	// move to project folder
	//with mkcert localhost 127.0.0.1 ::1 added certs to folder your are currently on
	// and use this created "localhost+2.pem", "localhost+2-key.pem" certs in server

	quit, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGALRM)
	defer stop()
	go func() {
		if err := srv.ListenAndServeTLS("localhost+2.pem", "localhost+2-key.pem"); err != nil && err != http.ErrServerClosed {
			fmt.Println("Can not run due to ", err.Error())
		}
	}()

	log.Println("Server is running on :8082")

	<-quit.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server is shutdown with err :=", err.Error())
	}

	fmt.Println("Server stops running")
}
