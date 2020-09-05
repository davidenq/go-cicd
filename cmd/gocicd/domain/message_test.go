package domain

import (
	"testing"

	"github.com/davidenq/go-cicd/cmd/gocicd/types"
)

//TestGenerateMessage .
func TestGenerateMessage(t *testing.T) {

	t.Run("should return: Hello Juan Perez your message will be send", func(t *testing.T) {
		want := "Hello Juan Perez your message will be send"
		dataMessage := &types.BodyMessage{
			Message:       "This is a test",
			To:            "Juan Perez",
			From:          "Rita Asturia",
			TimeToLifeSec: 45,
		}
		got := GenerateMessage(dataMessage)
		if got != want {
			t.Errorf("got '%s', want '%s", got, want)
		}
	})
}
