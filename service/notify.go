package service

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/rs/zerolog/log"
	v1 "github.com/x0y14/jackal/gen/notify/v1"
	"github.com/x0y14/jackal/mem"
	"strconv"
)

type NotifyService struct {
	Mq pulsar.Client
}

func (n *NotifyService) FetchMessage(
	ctx context.Context,
	request *connect_go.Request[v1.FetchMessageRequest],
	stream *connect_go.ServerStream[v1.FetchMessageResponse]) error {
	userId := request.Header().Get("X-User-ID")
	if userId == "" {
		return connect_go.NewError(connect_go.CodeUnauthenticated, fmt.Errorf("pls set X-User-ID"))
	}

	consumer, err := n.Mq.Subscribe(pulsar.ConsumerOptions{
		Topic:            userId,
		SubscriptionName: userId,
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create pulsar.consumer")
		return connect_go.NewError(connect_go.CodeInternal, err)
	}
	defer consumer.Close()

Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		default:
			msg, err := consumer.Receive(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("failed to receive message via pulsar.consumer")
				return connect_go.NewError(connect_go.CodeInternal, err)
			}
			consumer.Ack(msg)

			msgIdStr := msg.Payload()
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
