package utils_test

import (
	"testing"

	"github.com/aejoy/vk-yourstickers/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetScreenName(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		userID int
		typ    string
	}{
		{name: "divided", input: "[id375789362|@den.boytsov]", userID: 375789362, typ: "user"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userID, typ, err := utils.GetScreenName(test.input)
			assert.NoError(t, err, err)
			assert.Equal(t, test.userID, userID)
			assert.Equal(t, test.typ, typ)
		})
	}
}
