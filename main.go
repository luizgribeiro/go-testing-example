package main

import (
	"context"
	"encoding/json"
	"fmt"
	"luizgribeiro/testing/pkg/client"
	"luizgribeiro/testing/pkg/model"
	"luizgribeiro/testing/pkg/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {

	conn, err := pgx.Connect(context.Background(), "postgres://myuser:mypassword@127.0.0.1:5432/mydatabase")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	srv := &http.Server{
		Addr: ":8080",
	}

	accountClient := client.NewAccountClient(conn)

	userClient := client.NewUserClient(conn, accountClient)
	userService := service.NewUserService(conn, userClient)

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		switch r.Method {
		case http.MethodGet:
			fmt.Println("get enviado")
		case http.MethodPost:
			var userAccount model.UserAccount
			err := json.NewDecoder(r.Body).Decode(&userAccount)
			fmt.Printf("%+v\n", userAccount)
			if err != nil {
				fmt.Println("error")
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			userService.CreateUserWithAccount(ctx, &userAccount)
		}
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Starting HTTP Server")
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println("Error while starting server", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Failed to shutdown server")
		os.Exit(1)
	}
	fmt.Println("Server shutdown successfully")
}
