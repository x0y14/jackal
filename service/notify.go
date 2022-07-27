package service

import (
	"context"
	"fmt"
	connect_go "github.com/bufbuild/connect-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	v1 "github.com/x0y14/jackal/gen/notify/v1"
	"github.com/x0y14/jackal/mem"
	"strconv"
)

type NotifyService struct {
	Rb *amqp.Connection
}

func (n *NotifyService) FetchMessage(
	ctx context.Context,
	request *connect_go.Request[v1.FetchMessageRequest],
	stream *connect_go.ServerStream[v1.FetchMessageResponse]) error {
	userId := request.Header().Get("X-User-ID")
	if userId == "" {
		return connect_go.NewError(connect_go.CodeUnauthenticated, fmt.Errorf("pls set X-User-ID"))
	}

	ch, err := n.Rb.Channel()
	if err != nil {
		log.Error().Err(err).Msg("failed to create rb.channel")
		return connect_go.NewError(connect_go.CodeInternal, err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		userId,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to declare rb.queue")
		return connect_go.NewError(connect_go.CodeInternal, err)
	}

	queue, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to consume")
		return connect_go.NewError(connect_go.CodeInternal, err)
	}

Loop:
	for {
		select {
		case <-ctx.Done():
			log.Info().Str("userId", userId).Msg("disconnected")
			break Loop
		case msg := <-queue:
			msgIdStr := string(msg.Body)
			msgId, err := strconv.ParseInt(string(msgIdStr), 10, 64)
			if err != nil {
				log.Error().Err(err).Msg("failed to parse msgId")
				return connect_go.NewError(connect_go.CodeInternal, err)
			}
			message, err := mem.GetMessage(msgId)
			if err != nil {
				log.Error().Err(err).Msg("failed to get message from redis")
				return connect_go.NewError(connect_go.CodeInternal, err)
			}
			err = stream.Send(&v1.FetchMessageResponse{Message: message})
			if err != nil {
				log.Error().Err(err).Msg("failed to send response to user via stream")
				return connect_go.NewError(connect_go.CodeInternal, err)
			}
		}
	}

	return nil
}
