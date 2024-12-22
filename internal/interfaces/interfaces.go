package interfaces

import (
	"github.com/aejoy/vk-yourstickers/internal/domain"
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/update"
)

type Repository interface {
	GetStickerPacks(stickerPackIDs []int) (map[int]*domain.StickerPack, error)
	CreateStickerPack(id int, title, previewURL string) error
}

type Cache interface {
	GetStickerPacks(stickerPacksIDs []int) (map[int]*domain.StickerPack, error)
	CreateStickerPack(id int, title, previewURL string) error
}

type Service interface {
	GetStickerPacks(stickerPackIDs []int) (map[int]*domain.StickerPack, error)
	CreateStickerPack(id int, title, previewURL string) error
}

type Command interface {
	Execute(bot *api.API, message update.Message, args []string) error
}
