package cache

import (
	"context"
	"fmt"

	"github.com/aejoy/vk-yourstickers/internal/domain"
	"github.com/aejoy/vk-yourstickers/pkg/consts"
	"github.com/redis/rueidis"
)

type Cache struct {
	rueidis.Client
}

func NewCache(address string) (*Cache, error) {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{address}})
	if err != nil {
		return nil, err
	}

	return &Cache{client}, nil
}

func (c *Cache) GetStickerPacks(stickerPacksIDs []int) (map[int]*domain.StickerPack, error) {
	stickerPacks := make(map[int]*domain.StickerPack)

	for _, stickerPackID := range stickerPacksIDs {
		timeout, cancel := context.WithTimeout(context.Background(), consts.DefaultTimeout)
		defer cancel()

		key := fmt.Sprintf("sticker:%d", stickerPackID)
		res := c.Client.Do(timeout, c.Client.B().Hgetall().Key(key).Build())

		hashMap, err := res.AsMap()
		if err != nil {
			return nil, err
		}

		if len(hashMap) == 0 {
			continue
		}

		stickerPack := domain.StickerPack{}

		idKey := hashMap["id"]
		titleKey := hashMap["title"]
		previewKey := hashMap["preview_url"]

		id, err := idKey.AsInt64()
		if err != nil {
			return nil, err
		}

		title, err := titleKey.ToString()
		if err != nil {
			return nil, err
		}

		previewURL, err := previewKey.ToString()
		if err != nil {
			return nil, err
		}

		stickerPack.ID = int(id)
		stickerPack.Title = title
		stickerPack.PreviewURL = previewURL

		stickerPacks[stickerPack.ID] = &stickerPack
	}

	return stickerPacks, nil
}

func (c *Cache) CreateStickerPack(id int, title, previewURL string) error {
	stickerPackID := fmt.Sprint(id)
	key := "sticker:" + stickerPackID

	timeout, cancel := context.WithTimeout(context.Background(), consts.DefaultTimeout)
	defer cancel()

	if err := c.Client.Do(timeout,
		c.Client.B().Hset().Key(key).FieldValue().
			FieldValue("id", stickerPackID).
			FieldValue("title", title).
			FieldValue("preview_url", previewURL).
			Build()).Error(); err != nil {
		return err
	}

	expTimeout, expCancel := context.WithTimeout(context.Background(), consts.DefaultTimeout)
	defer expCancel()

	return c.Client.Do(expTimeout, c.Client.B().Expire().Key(key).Seconds(consts.Day).Build()).Error()
}

func (c *Cache) Load(stickerPacksIDs []domain.StickerPack) error {
	for _, stickerPackID := range stickerPacksIDs {
		if err := c.CreateStickerPack(stickerPackID.ID, stickerPackID.Title, stickerPackID.PreviewURL); err != nil {
			return err
		}
	}

	return nil
}
