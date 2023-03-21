package mem

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	typesv1 "github.com/x0y14/jackal/gen/types/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

var rdb *redis.Client

func Init(addr string) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	return err
}

func GenerateMessageId() (int64, error) {
	msgId, err := rdb.Incr(context.Background(), "message_id").Result()
	if err == redis.Nil {
		msgId = 0
		err = rdb.Set(context.Background(), "message_id", msgId, 0).Err()
		if err != nil {
			return 0, err
		}
	}
	return msgId, err
}

func StoreMessage(message *typesv1.Message) (int64, error) {
	msgId, err := GenerateMessageId()
	if err != nil {
		return 0, err
	}
	err = rdb.HSet(
		context.Background(),
		strconv.FormatInt(msgId, 10),
		map[string]string{
			"from":       message.From,
			"to":         message.To,
			"text":       message.Text,
			"metadata":   message.Metadata,
			"kind":       strconv.FormatInt(int64(message.Kind.Number()), 10),
			"created_at": strconv.FormatInt(message.CreatedAt.Seconds, 10),
		}).Err()
	if err != nil {
		return 0, err
	}
	return msgId, nil
}

type Message struct {
	MessageId string `redis:"message_id"`
	From      string `redis:"from"`
	To        string `redis:"to"`
	Text      string `redis:"text"`
	Metadata  string `redis:"metadata"`
	Kind      string `redis:"kind"`
	CreatedAt string `redis:"created_at"`
}

func GetMessage(messageId int64) (*typesv1.Message, error) {
	var msg Message
	err := rdb.HGetAll(context.Background(), strconv.FormatInt(messageId, 10)).Scan(&msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get hash")
	}

	kind, err := strconv.ParseInt(msg.Kind, 10, 32)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse kind")
	}
	createdAt, err := strconv.ParseInt(msg.CreatedAt, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse created at")
	}

	return &typesv1.Message{
		MessageId: messageId,
		From:      msg.From,
		To:        msg.To,
		Text:      msg.Text,
		Metadata:  msg.Metadata,
		Kind:      typesv1.MessageKind(kind),
		CreatedAt: &timestamppb.Timestamp{Seconds: createdAt},
	}, nil
}
