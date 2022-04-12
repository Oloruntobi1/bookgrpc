package main

import (
	"context"
	"fmt"
	"log"

	users "github.com/Oloruntobi1/bookgrpc/unaryexample/protos"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client..")

	opts := grpc.WithInsecure()
	con, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatalf("Error connecting: %v \n", err)
	}

	defer con.Close()
	c := users.NewUsersClient(con)
	getUsers(c)
}

// getUsers function
func getUsers(c users.UsersClient) {
	req := &users.GetUsersReq{
		Status: users.UserStatus_USER_STATUS_UNKNOWN,
	}

	res, err := c.GetUsers(context.Background(), req)
	if err != nil {
		log.Fatalf("Error on GetUsers rpc call: %v \n", err)
	}
	fmt.Printf("Response: %+v\n", res)
}