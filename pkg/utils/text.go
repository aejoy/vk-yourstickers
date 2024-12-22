package utils

import "fmt"

func GetObjectListsText(senderID, userID int, objectsName, yourObjects string) string {
	if senderID == userID {
		return yourObjects
	}

	return fmt.Sprintf("%s @id%d(Пользователя)", objectsName, userID)
}
