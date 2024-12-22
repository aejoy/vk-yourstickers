package stickers

import (
	"github.com/aejoy/vk-yourstickers/internal/domain"
	"github.com/aejoy/vk-yourstickers/internal/interfaces"
)

type Service struct {
	repo  interfaces.Repository
	cache interfaces.Cache
}

func NewStickersService(repo interfaces.Repository, cache interfaces.Cache) *Service {
	return &Service{repo, cache}
}

func (s *Service) GetStickerPacks(stickerPackIDs []int) (map[int]*domain.StickerPack, error) {
	stickerPacks, err := s.cache.GetStickerPacks(stickerPackIDs)
	if err != nil {
		return nil, err
	}

	undefinedStickerPackIDs := make([]int, 0)

	for _, stickerPackID := range stickerPackIDs {
		if _, ok := stickerPacks[stickerPackID]; !ok {
			undefinedStickerPackIDs = append(undefinedStickerPackIDs, stickerPackID)
		}
	}

	if len(undefinedStickerPackIDs) != 0 {
		savedStickerPacks, err := s.repo.GetStickerPacks(undefinedStickerPackIDs)
		if err != nil {
			return nil, err
		}

		for stickerPackID, stickerPack := range savedStickerPacks {
			stickerPacks[stickerPackID] = stickerPack

			if err := s.cache.CreateStickerPack(stickerPackID, stickerPack.Title, stickerPack.PreviewURL); err != nil {
				return nil, err
			}
		}
	}

	return stickerPacks, nil
}

func (s *Service) CreateStickerPack(id int, title, previewURL string) error {
	if err := s.cache.CreateStickerPack(id, title, previewURL); err != nil {
		return err
	}

	return s.repo.CreateStickerPack(id, title, previewURL)
}
