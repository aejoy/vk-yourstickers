package commands

import (
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/update"
)

type HelpCmd struct {
}

func NewHelpCmd() HelpCmd {
	return HelpCmd{}
}

func (c HelpCmd) Execute(bot *api.API, message update.Message, _ []string) error {
	_, err := bot.SendMessage(message.ChatID, []string{"article-227659026_316689_38f5db5004fadd7d7f"})
	return err
}
