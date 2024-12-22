package utils

import (
	"strconv"

	"github.com/aejoy/vk-yourstickers/pkg/consts"
)

func GetScreenName(in string) (userID int, typ string, err error) {
	if groups := consts.ScreenNameRegex.FindStringSubmatch(in); len(groups) > consts.MinScreenNameLength {
		userID, err = strconv.Atoi(groups[2])
		if err != nil {
			return
		}

		if userID < 0 {
			typ = "group"
		} else {
			typ = "user"
		}
	}

	return userID, typ, err
}
