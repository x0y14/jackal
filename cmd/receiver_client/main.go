package main

import (
	"context"
	"github.com/bufbuild/connect-go"
	chatv1 "github.com/x0y14/jackal/gen/chat/v1"
	"github.com/x0y14/jackal/gen/chat/v1/chatv1connect"
	notifyv1 "github.com/x0y14/jackal/gen/notify/v1"
	"github.com/x0y14/jackal/gen/notify/v1/notifyv1connect"
	typesv1 "github.com/x0y14/jackal/gen/types/v1"
	"log"
	"net/http"
)

const (
	SenderId   = "test-sender"
	ReceiverId = "test-receiver"
)

func main() {
	chatClient := chatv1connect.NewChatServiceClient(
		http.DefaultClient,
		"http://localhost:8081")

	notifyClient := notifyv1connect.NewNotifyServiceClient(
		http.DefaultClient,
		"http://localhost:8082")

	_, err := chatClient.CreateUser(context.Background(), connect.NewRequest(&chatv1.CreateUserRequest{
		User: &typesv1.User{
			UserId:      ReceiverId,
			DisplayName: "TOM",
		},
	}))
	if err != nil {
		log.Printf("failed to create user: %v", err)
	}

	req := connect.NewRequest(&notifyv1.FetchMessageRequest{LastMessageId: 0})
	req.Header().Set("X-User-ID", ReceiverId)
	stream, err := notifyClient.FetchMessage(
		context.Background(),
		req,
	)
	if err != nil {
		log.Fatalf("failed to fetch msg: %v", err)
	}
	defer stream.Close()
	for stream.Receive() {
		msg := stream.Msg()
		log.Printf("RECEIVE: %v", msg.Message)
	}
	if stream.Err() != nil {
		log.Fatalf("stream err: %v", err)
	}
}
