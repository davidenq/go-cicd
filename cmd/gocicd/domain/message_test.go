package domain

import (
	"testing"
)

//BodyMessage .
type BodyMessage struct {
	Message       string
	To            string
	From          string
	TimeToLifeSec int
}

//TestGenerateMessage .
func TestGenerateMessage(t *testing.T) {

	t.Run("should return: Hello Juan Perez your message will be send", func(t *testing.T) {
		want := "Hello Juan Perez your message will be send"
		dataMessage := &BodyMessage{
			Message:       "This is a test",
			To:            "Juan Perez",
			From:          "Rita Asturia",
			TimeToLifeSec: 45,
		}
		got := GenerateMessage(dataMessage)
		if got != want {
			t.Error("got '%s', want '%s", got, want)
		}
	})
}
