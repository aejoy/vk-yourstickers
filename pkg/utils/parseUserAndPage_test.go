package utils_test

import (
	"testing"

	"github.com/aejoy/vk-yourstickers/pkg/utils"
	"github.com/aejoy/vkgo/update"
	"github.com/stretchr/testify/assert"
)

const Creator = 542439242

func TestParseUserAndPage(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		userID int
		page   int
	}{
		{name: "userID & page", args: []string{"[id690069912|@kirieshki120]", "2"}, userID: 690069912, page: 2},
		{name: "only userID", args: []string{"[id690069912|@kirieshki120]"}, userID: 690069912, page: 1},
		{name: "only page", args: []string{"1"}, userID: Creator, page: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userID, page, err := utils.ParseUserAndPage(test.args, update.Message{
				UserID: Creator,
			})

			assert.NoError(t, err, err)
			assert.Equal(t, test.userID, userID)
			assert.Equal(t, test.page, page)
		})
	}
}
