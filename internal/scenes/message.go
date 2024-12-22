package scenes

import (
	"fmt"
	"strings"

	"github.com/aejoy/vk-yourstickers/internal/interfaces"
	"github.com/aejoy/vk-yourstickers/internal/services/uploader"

	"github.com/aejoy/vk-yourstickers/internal/scenes/commands"
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/update"
)

type MessageScene struct {
	commands map[string]interfaces.Command
}

func NewMessageScene(userBot *api.API, uploaderService *uploader.Service, stickersService interfaces.Service) MessageScene {
	return MessageScene{
		commands: map[string]interfaces.Command{
			"пинг":    commands.NewPingCmd(),
			"помощь":  commands.NewHelpCmd(),
			"стикеры": commands.NewStickersCmd(userBot, uploaderService, stickersService),
		},
	}
}

func (s MessageScene) Message(bot *api.API, message update.Message) {
	if len(message.Text) == 0 {
		return
	}

	if prefix := string(message.Text[0]); prefix != ":" && prefix != "/" {
		return
	}

	message.Text = message.Text[1:]

	args := strings.Fields(message.Text)
	if len(args) == 0 {
		return
	}

	if cmd, ok := s.commands[args[0]]; ok {
		if err := cmd.Execute(bot, message, args[1:]); err != nil {
			if _, err := bot.SendMessage(message.ChatID, err.Error()); err != nil {
				fmt.Println("main.SendMessage error:", err)
			}
		}
	}
}
