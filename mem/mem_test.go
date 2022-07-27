package mem

import (
	"fmt"
	typesv1 "github.com/x0y14/jackal/gen/types/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func init() {
	Init("localhost:6379")
}

func TestGenerateMessageId(t *testing.T) {
	msgId, err := GenerateMessageId()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msgId)
}

func TestStoreGetMessage(t *testing.T) {
	msgId, err := StoreMessage(
		&typesv1.Message{
			MessageId: 0,
			From:      "a",
			To:        "b",
			Text:      "hello",
			Metadata:  "{}",
			Kind:      1,
			CreatedAt: timestamppb.Now(),
		})
	if err != nil {
		t.Fatal(err)
	}

	msg, err := GetMessage(msgId)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(msg)
}
