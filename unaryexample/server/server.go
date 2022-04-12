package main

import (
	"context"
	"net"
	usersv1 "github.com/Oloruntobi1/bookgrpc/unaryexample/protos"

	"fmt"
	"log"
	"google.golang.org/grpc"
)

type server struct {
	usersv1.UnimplementedUsersServer
}

func main() {
	fmt.Println("Starting server..")

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}

	s := grpc.NewServer()
	usersv1.RegisterUsersServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fails to serve: %v", err)
	}
}

// GetUsers function
func (*server) GetUsers(ctx context.Context, req *usersv1.GetUsersReq) (*usersv1.GetUsersRes, error) {
	status := req.GetStatus()
	userList := getUserList()
	usersFiltered := []*usersv1.User{}
	switch status {
	case usersv1.UserStatus_USER_STATUS_ACTIVE:
		usersFiltered = filterBy("active", userList)
	case usersv1.UserStatus_USER_STATUS_BLOCKED:
		usersFiltered = filterBy("blocked", userList)
	case usersv1.UserStatus_USER_STATUS_SUSPENDED:
		usersFiltered = filterBy("suspended", userList)
	default:
		usersFiltered = userList
	}

	res := usersv1.GetUsersRes{
		Users: usersFiltered,
	}
	return &res, nil
}

// getUserList function
func getUserList() []*usersv1.User {
	userObj := []*usersv1.User{}
	userObj = append(userObj, &usersv1.User{Name: "John", LastName: "Phill", Age: 34, Email: "john@gmail.com", Status: "active"})
	userObj = append(userObj, &usersv1.User{Name: "Carl", LastName: "Meertz", Age: 23, Email: "carl@gmail.com", Status: "active"})
	userObj = append(userObj, &usersv1.User{Name: "Susan", LastName: "Zeanz", Age: 30, Email: "susan@gmail.com", Status: "blocked"})
	userObj = append(userObj, &usersv1.User{Name: "Marylen", LastName: "Inc", Age: 29, Email: "marylen@gmail.com", Status: "blocked"})
	userObj = append(userObj, &usersv1.User{Name: "Peet", LastName: "Green", Age: 25, Email: "peet@gmail.com", Status: "ignored"})
	userObj = append(userObj, &usersv1.User{Name: "Maty", LastName: "Jackson", Age: 28, Email: "maty@gmail.com", Status: "suspended"})
	return userObj
}

// filterBy function
func filterBy(status string, userList []*usersv1.User) []*usersv1.User {
	usersFiltered := []*usersv1.User{}
	for _, v := range userList {
		if (v.Status == "blocked" || v.Status == "ignored") && status == "blocked" {
			usersFiltered = append(usersFiltered, v)
		} else if v.Status == status {
			usersFiltered = append(usersFiltered, v)
		}
	}
	return usersFiltered
}