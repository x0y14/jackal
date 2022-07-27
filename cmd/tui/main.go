package main

import (
	"context"
	"flag"
	"github.com/bufbuild/connect-go"
	tea "github.com/charmbracelet/bubbletea"
	chatv1 "github.com/x0y14/jackal/gen/chat/v1"
	"github.com/x0y14/jackal/gen/chat/v1/chatv1connect"
	notifyv1 "github.com/x0y14/jackal/gen/notify/v1"
	"github.com/x0y14/jackal/gen/notify/v1/notifyv1connect"
	typesv1 "github.com/x0y14/jackal/gen/types/v1"
	"github.com/x0y14/jackal/tui"
	"log"
	"net/http"
)

func main() {
	var userId string
	var displayName string
	var receiver string
	flag.StringVar(&userId, "userid", "", "")
	flag.StringVar(&displayName, "name", "", "")
	flag.StringVar(&receiver, "receiver", "", "")
	flag.Parse()

	chatClient := chatv1connect.NewChatServiceClient(
		http.DefaultClient,
		"http://localhost:8081")

	notifyClient := notifyv1connect.NewNotifyServiceClient(
		http.DefaultClient,
		"http://localhost:8082")

	_, err := chatClient.CreateUser(context.Background(), connect.NewRequest(&chatv1.CreateUserRequest{
		User: &typesv1.User{
			UserId:      userId,
			DisplayName: displayName,
		},
	}))
	if err != nil {
		log.Printf("failed to create user: %v", err)
	}

	p := tea.NewProgram(tui.InitialModel(chatClient, userId, receiver))

	go func() {
		req := connect.NewRequest(&notifyv1.FetchMessageRequest{LastMessageId: 0})
		req.Header().Set("X-User-ID", userId)
		stream, err := notifyClient.FetchMessage(
			context.Background(),
			req)
		if err != nil {
			log.Fatalf("failed to fetch message: %v", err)
		}
		for stream.Receive() {
			p.Send(tui.ReceiveRespMsg{Resp: stream.Msg()})
		}
	}()

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
