package service

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	connect_go "github.com/bufbuild/connect-go"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	"github.com/x0y14/jackal"
	"github.com/x0y14/jackal/database"
	v1 "github.com/x0y14/jackal/gen/chat/v1"
	"github.com/x0y14/jackal/mem"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChatService struct {
	Mq pulsar.Client
}

func (c *ChatService) CreateUser(
	_ context.Context,
	request *connect_go.Request[v1.CreateUserRequest]) (
	*connect_go.Response[v1.CreateUserResponse], error) {
	log.Info().Str("method", request.Spec().Procedure)

	err := database.CreateUser(request.Msg.User)
	if err != nil {
		log.Warn().Msg(err.Error())
		return nil, connect_go.NewError(connect_go.CodeInvalidArgument, err)
	}

	return connect_go.NewResponse(
		&v1.CreateUserResponse{
			User: request.Msg.User}), nil
}

func (c *ChatService) SendMessage(
	_ context.Context,
	request *connect_go.Request[v1.SendMessageRequest]) (
	*connect_go.Response[v1.SendMessageResponse], error) {
	log.Info().Str("method", request.Spec().Procedure)

	// ヘッダーからuserIdを取り出す
	userId := request.Header().Get("X-User-ID")
	if userId == "" {
		log.Warn().Msg("empty X-User-ID")
		return nil, connect_go.NewError(connect_go.CodeUnauthenticated, fmt.Errorf("pls set X-User-ID"))
	}

	// どちらのユーザーも存在することを確認
	_, err := database.GetUser(userId)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInvalidArgument, fmt.Errorf("sender, who are you"))
	}
	_, err = database.GetUser(request.Msg.Message.To)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInvalidArgument, fmt.Errorf("receiver, who are you"))
	}

	// 送信者、受信者のidをくっつけてチャットidを作成
	chatId, err := jackal.CreateChatId(userId, request.Msg.Message.To)
	if err != nil {
		log.Warn().Str("chatId", chatId).Msg("invalid chatId")
		return nil, connect_go.NewError(connect_go.CodeInvalidArgument, fmt.Errorf("invalid sender or receiver"))
	}

	// 一部情報を強制的に書き換え
	//request.Msg.Message.To = chatId
	request.Msg.Message.CreatedAt = timestamppb.Now()

	// sqlite
	//messageId, err := database.CreateMessage(request.Msg.Message)

	// redis
	messageId, err := mem.StoreMessage(request.Msg.Message)
	if err != nil {
		log.Error().Str("method", request.Spec().Procedure).Err(err)
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	// 受信したことをmqに教えてあげる
	producer, err := c.Mq.CreateProducer(pulsar.ProducerOptions{
		Topic: request.Msg.Message.To,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create pulsar.producer")
	}
	defer producer.Close()
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(fmt.Sprintf("%v", messageId)),
	})
	if err != nil {
		return nil, err
	}

	// idを書き換え
	request.Msg.Message.MessageId = messageId

	return connect_go.NewResponse(
		&v1.SendMessageResponse{
			Message: request.Msg.Message,
		}), nil
}
