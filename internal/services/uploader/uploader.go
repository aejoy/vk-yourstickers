package uploader

import (
	"fmt"
	"strings"
	"sync"

	"github.com/aejoy/vk-yourstickers/internal/domain"
	"github.com/aejoy/vk-yourstickers/internal/interfaces"

	"github.com/aejoy/vk-yourstickers/pkg/image"
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/types"
)

// Pattern: Publish-Subscribe

type Service struct {
	botID           int
	albumID         int
	mutex           sync.Mutex
	pending         map[int][]chan string
	stickersService interfaces.Service
}

func NewUploaderService(botID, albumID int, stickersService interfaces.Service) *Service {
	return &Service{
		botID:           botID,
		albumID:         albumID,
		mutex:           sync.Mutex{},
		pending:         map[int][]chan string{},
		stickersService: stickersService,
	}
}

func (s *Service) UploadStickerPacks(userBot *api.API, infusedStickerPacks []*domain.StickerPack) []*domain.StickerPack {
	wg := &sync.WaitGroup{}

	for _, infusedStickerPack := range infusedStickerPacks {
		wg.Add(1)

		go s.UploadStickerPack(wg, userBot, infusedStickerPack)
	}

	wg.Wait()

	return infusedStickerPacks
}

func (s *Service) UploadStickerPack(wg *sync.WaitGroup, userBot *api.API, stickerPack *domain.StickerPack) {
	defer wg.Done()
	s.mutex.Lock()

	if chans, exist := s.pending[stickerPack.ID]; exist {
		ch := make(chan string)

		s.pending[stickerPack.ID] = append(chans, ch)
		s.mutex.Unlock()

		stickerPack.PreviewURL = <-ch

		return
	}

	s.pending[stickerPack.ID] = []chan string{}
	s.mutex.Unlock()

	defer delete(s.pending, stickerPack.ID)

	var img []byte

	for i := 1; i < 8; i++ {
		imgBytes, err := image.Fetch(strings.ReplaceAll(stickerPack.PreviewURL, "592-1", fmt.Sprintf("592-%d", i)))
		if err != nil {
			continue
		}

		img = imgBytes
	}

	uploadFile := types.UploadFile{FieldName: "file1", FileName: "sticker.png", Bytes: img}

	photos, err := userBot.UploadAlbumPhotos(s.albumID, []types.UploadFile{uploadFile}, s.botID)
	if err != nil {
		fmt.Println("uploadErr", err)

		return
	}

	photo := photos[0]
	stickerPack.PreviewURL = fmt.Sprintf("%d_%d", photo.OwnerID, photo.ID)

	if err := s.stickersService.CreateStickerPack(stickerPack.ID, stickerPack.Title, stickerPack.PreviewURL); err != nil {
		fmt.Println("repo.CreateStickerPack", err)
	}

	s.mutex.Lock()

	for _, ch := range s.pending[stickerPack.ID] {
		ch <- stickerPack.PreviewURL
		close(ch)
	}

	s.mutex.Unlock()
}
