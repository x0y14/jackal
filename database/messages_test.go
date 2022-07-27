package database

import (
	"fmt"
	typesv1 "github.com/x0y14/jackal/gen/types/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"testing"
)

func init() {
	err := Init("../data/sqlite/jackal.sqlite")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateMessage(t *testing.T) {
	defer Close()
	messageId, err := CreateMessage(&typesv1.Message{
		//MessageId: 0,
		From:      "a",
		To:        "b",
		Text:      "hello",
		Metadata:  "{}",
		Kind:      0,
		CreatedAt: timestamppb.Now(),
	})
	if err != nil {
		t.Fatalf("failed to create message: %v", err)
	}
	fmt.Println(messageId)
}

func TestGetMessage(t *testing.T) {
	defer Close()
	message, err := GetMessage(1)
	if err != nil {
		t.Fatalf("failed to get message: %v", err)
	}

	fmt.Println(message)
}
