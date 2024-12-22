package commands

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/aejoy/vkgo/update"

	"github.com/aejoy/vkgo/api"
)

type PingCmd struct {
}

func NewPingCmd() PingCmd {
	return PingCmd{}
}

func (c PingCmd) Execute(bot *api.API, _ []string, message update.Message) error {
	now, text := time.Now(), "✦ Понг!"

	sent, err := bot.SendMessage(message.ChatID, text)
	if err != nil {
		return errors.Wrap(err, "SendMessage")
	}

	_, err = bot.EditMessage(sent.ChatID, sent.ChatMessageID,
		fmt.Sprintf("%s (%v)", text, time.Since(now)))

	return err
}
