package utils

import (
	"fmt"
	"strings"

	"github.com/aejoy/vk-yourstickers/internal/domain"
	"github.com/aejoy/vk-yourstickers/pkg/responses"
)

func FilterOnlyAvailableStickerPacks(stickerPack responses.StickerPack) *domain.StickerPack {
	if !stickerPack.HasAnimation && !stickerPack.IsVmoji {
		previewURL := strings.ReplaceAll(stickerPack.Previews[len(stickerPack.Previews)-1].URL, "thumb-102-", "preview-592-1")

		if previewURL != "" {
			return &domain.StickerPack{
				ID:         stickerPack.ID,
				Title:      stickerPack.Title,
				PreviewURL: previewURL,
			}
		}
	}

	return nil
}

func GetUndefinedStickerIDs(existsStickerIDs map[int]*domain.StickerPack, stickerIDs []int) []string {
	undefinedStickerPackIDs := make([]string, 0, len(existsStickerIDs))

	for _, stickerID := range stickerIDs {
		if _, ok := existsStickerIDs[stickerID]; !ok {
			undefinedStickerPackIDs = append(undefinedStickerPackIDs, fmt.Sprint(stickerID))
		}
	}

	return undefinedStickerPackIDs
}
