package database

import (
	typesv1 "github.com/x0y14/jackal/gen/types/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateMessage(message *typesv1.Message) (int64, error) {
	res, err := database.Exec(
		`insert into messages ("from", "to", text, metadata, kind, created_at) values (?, ?, ?, ?, ?, ?)`,
		message.From,
		message.To,
		message.Text,
		message.Metadata,
		message.Kind,
		message.CreatedAt.Seconds,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func GetMessage(messageId int64) (*typesv1.Message, error) {
	stmt, err := database.Prepare(`select "from", "to", text, metadata, kind, created_at from messages where message_id = ?`)
	if err != nil {
		return nil, err
	}

	var from string
	var to string
	var text string
	var metadata string
	var kind int
	var createdAt int
	err = stmt.QueryRow(messageId).Scan(&from, &to, &text, &metadata, &kind, &createdAt)
	if err != nil {
		return nil, err
	}

	return &typesv1.Message{
		MessageId: messageId,
		From:      from,
		To:        to,
		Text:      text,
		Metadata:  metadata,
		Kind:      typesv1.MessageKind(kind),
		CreatedAt: &timestamppb.Timestamp{Seconds: int64(createdAt)},
	}, nil
}
