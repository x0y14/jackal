package main

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	chatv1 "github.com/x0y14/jackal/gen/chat/v1"
	"github.com/x0y14/jackal/gen/chat/v1/chatv1connect"
	typesv1 "github.com/x0y14/jackal/gen/types/v1"
	"log"
	"net/http"
	"time"
)

const (
	SenderId   = "test-sender"
	ReceiverId = "test-receiver"
)

func main() {
	chatClient := chatv1connect.NewChatServiceClient(
		http.DefaultClient,
		"http://localhost:8081")

	_, err := chatClient.CreateUser(context.Background(), connect.NewRequest(&chatv1.CreateUserRequest{
		User: &typesv1.User{
			UserId:      SenderId,
			DisplayName: "JOHN",
		},
	}))
	if err != nil {
		log.Println(err)
	}

	count := 0

	for {
		time.Sleep(time.Second)

		msg := &chatv1.SendMessageRequest{
			Message: &typesv1.Message{
				//MessageId: 0,
				From:     SenderId,
				To:       ReceiverId,
				Text:     fmt.Sprintf("hello, %d", count),
				Metadata: "{}",
				Kind:     0,
				//CreatedAt: nil,
			}}
		req := connect.NewRequest(msg)
		req.Header().Set("X-User-ID", SenderId)

		_, err := chatClient.SendMessage(
			context.Background(),
			req,
		)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("SEND: %v", msg)
		count++
	}

}
