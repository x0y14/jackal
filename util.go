package jackal

import (
	"fmt"
	"sort"
	"strings"
)

func CreateChatId(userId1, userId2 string) (string, error) {
	if userId1 == userId2 {
		return "", fmt.Errorf("same userId")
	}
	ids := []string{userId1, userId2}
	sort.Strings(ids)
	return fmt.Sprintf("di|%s", strings.Join(ids, ".")), nil
}
