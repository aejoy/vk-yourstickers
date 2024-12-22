package utils

import (
	"strconv"

	"github.com/aejoy/vk-yourstickers/pkg/consts"
	"github.com/aejoy/vkgo/update"
)

func ParseUserAndPage(args []string, message update.Message) (userID, page int, err error) {
	if message.Reply != nil {
		userID = message.Reply.UserID
	}

	if len(message.Forwards) > 0 {
		userID = message.Forwards[0].UserID
	}

	for _, arg := range args {
		if userID == 0 {
			var typ string
			userID, typ, err = GetScreenName(arg)

			if typ == "group" {
				return userID, page, consts.ErrNotAvailableForGroups
			}

			continue
		}

		page, err = strconv.Atoi(arg)
		if err != nil {
			return userID, page, err
		}
	}

	if page <= 0 {
		page = 1
	}

	if userID == 0 {
		userID = message.UserID
	}

	return userID, page, err
}
