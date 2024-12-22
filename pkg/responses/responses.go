package responses

import (
	"github.com/aejoy/vkgo/responses"
	"github.com/aejoy/vkgo/update"
)

type StickersReply struct {
	responses.ErrorInterface
	Response Stickers `json:"response"`
}

type Stickers struct {
	Items []StickerPack `json:"items"`
}

type StickerPack struct {
	ID           int                    `json:"id"`
	Type         string                 `json:"type"`
	IsNew        bool                   `json:"is_new"`
	Copyright    string                 `json:"copyright"`
	Purchased    int                    `json:"purchased"`
	Active       int                    `json:"active"`
	PurchaseDate int                    `json:"purchase_date"`
	Title        string                 `json:"title"`
	Stickers     []Sticker              `json:"stickers"`
	Icon         StickerPackIcon        `json:"icon"`
	Previews     []StickerPackPreview   `json:"previews"`
	VmojiAvatar  StickerPackVMojiAvatar `json:"vmoji_avatar"`
	IsVmoji      bool                   `json:"is_vmoji"`
	HasAnimation bool                   `json:"has_animation"`
}

type Sticker struct {
	InnerType            string             `json:"inner_type"`
	StickerID            int                `json:"sticker_id"`
	Images               []update.PhotoSize `json:"images"`
	ImagesWithBackground []update.PhotoSize `json:"images_with_background"`
	IsAllowed            bool               `json:"is_allowed"`
	Render               StickerRender      `json:"render"`
	VMoji                VMoji              `json:"vmoji"`
}

type StickerRender struct {
	ID          string               `json:"id"`
	Images      []StickerRenderImage `json:"images"`
	IsStub      bool                 `json:"is_stub"`
	IsRendering bool                 `json:"is_rendering"`
}

type StickerRenderImage struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Theme  string `json:"theme"`
}

type VMoji struct {
	CharacterID string `json:"character_id"`
}

type StickerPackIcon struct {
	BaseURL string `json:"base_url"`
}

type StickerPackPreview struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type StickerPackVMojiAvatar struct {
	ID          string `json:"id"`
	CharacterID string `json:"character_id"`
	Name        string `json:"name"`
	CanShare    bool   `json:"can_share"`
	IsActive    bool   `json:"is_active"`
}

type StickerCarouselElement struct {
	Name    string
	PhotoID string
}
