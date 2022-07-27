package database

import (
	"fmt"
	typesv1 "github.com/x0y14/jackal/gen/types/v1"
	"log"
	"testing"
)

func init() {
	err := Init("../data/sqlite/jackal.sqlite")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateUser(t *testing.T) {
	err := CreateUser(&typesv1.User{
		UserId:      "user_a",
		DisplayName: "A-san",
	})
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}
}

func TestGetUser(t *testing.T) {
	user, err := GetUser("user_a20")
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}
	fmt.Println(user)
}
