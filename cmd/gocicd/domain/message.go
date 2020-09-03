package domain

import (
	"fmt"

	"github.com/davidenq/go-cicd/cmd/gocicd/types"
)

//GenerateMessage .
func GenerateMessage(bodyMessage *types.BodyMessage) string {
	message := fmt.Sprintf("Hello %s your message will be send", bodyMessage.To)
	return message
}
