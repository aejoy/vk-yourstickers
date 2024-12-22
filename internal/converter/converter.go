package converter

import (
	"github.com/aejoy/vkgo/keyboard"
	"github.com/aejoy/vkgo/responses"
	"github.com/aejoy/vkgo/template"
)

func ToCarouselElement(title, previewURL string) template.Element {
	return template.Element{
		Title:       title,
		Description: "Набор в наличии",
		PhotoID:     previewURL,
		Action:      template.Action{Type: "open_photo"},
		Buttons:     &keyboard.New().Link("Магазин", "https://vk.com/stickers").Buttons[0],
	}
}

func ToStickerIDs(stickers []responses.Sticker) []int {
	stickerIDs := make([]int, 0, len(stickers))

	for _, sticker := range stickers {
		stickerIDs = append(stickerIDs, sticker.ID)
	}

	return stickerIDs
}
