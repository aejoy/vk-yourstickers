package commands

import (
	"fmt"
	"strings"

	"github.com/aejoy/vk-yourstickers/internal/converter"
	sharedResponses "github.com/aejoy/vk-yourstickers/pkg/responses"
	"github.com/aejoy/vk-yourstickers/pkg/utils"

	"github.com/aejoy/vk-yourstickers/internal/domain"
	"github.com/aejoy/vk-yourstickers/internal/interfaces"
	"github.com/aejoy/vk-yourstickers/internal/services/uploader"
	"github.com/aejoy/vk-yourstickers/pkg/consts"
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/template"
	"github.com/aejoy/vkgo/types"
	"github.com/aejoy/vkgo/update"
	"github.com/pkg/errors"
)

type StickersCmd struct {
	userBot         *api.API
	uploaderService *uploader.Service
	stickersService interfaces.Service
}

func NewStickersCmd(userBot *api.API, uploaderService *uploader.Service, stickersService interfaces.Service) StickersCmd {
	return StickersCmd{userBot, uploaderService, stickersService}
}

func (c StickersCmd) Execute(bot *api.API, message update.Message, args []string) error {
	fwdMsg, err := types.NewForward(message.ChatID, message.ChatMessageID, true)
	if err != nil {
		return errors.Wrap(err, "NewForward")
	}

	userID, page, err := utils.ParseUserAndPage(args, message)
	if err != nil {
		return err
	}

	stickers, err := c.userBot.GetStickers(userID)
	if err != nil {
		return errors.Wrap(err, "GetStickers")
	}

	stickersCount := len(stickers)

	pages, offset, end, err := utils.GetPaginationBounds(consts.PacksPerPage, stickersCount, page)
	if err != nil {
		return err
	}

	stickerIDs := converter.ToStickerIDs(stickers[offset:end])

	existsStickerPacks, err := c.stickersService.GetStickerPacks(stickerIDs)
	if err != nil {
		return errors.Wrap(err, "GetStickerPacks")
	}

	var unavailableStickerPacksCount int

	if unknownsStickerIDs := utils.GetUndefinedStickerIDs(existsStickerPacks, stickerIDs); len(unknownsStickerIDs) > consts.UnknownsNotEmpty {
		if err := c.UploadUnknownsStickerPacks(bot, message.ChatID, unknownsStickerIDs,
			&unavailableStickerPacksCount, existsStickerPacks); err != nil {
			return err
		}
	}

	text := fmt.Sprintf("💖 %s:\n\n🎾 Всего стикеров: %d\n💬 Страница: %d/%d", utils.GetObjectListsText(message.UserID, userID, "Стикеры", "Ваши стикеры"), stickersCount, page, pages)

	if unavailableStickerPacksCount > 0 {
		text += fmt.Sprintf("\n\n☢ Анимированных & VMoji: %d", unavailableStickerPacksCount)
	}

	if _, err := bot.SendMessage(types.SendMessage{
		ChatID:          message.ChatID,
		Text:            text,
		Template:        NewCarousel(existsStickerPacks),
		Forward:         fwdMsg,
		DisableMentions: true,
	}); err != nil {
		return errors.Wrap(err, "SendMessage")
	}

	return nil
}

func NewCarousel(existsStickerPacks map[int]*domain.StickerPack) string {
	carousel := template.NewCarousel()

	for _, existsStickerPack := range existsStickerPacks {
		carousel.Add(converter.ToCarouselElement(existsStickerPack.Title, existsStickerPack.PreviewURL))
	}

	return carousel.JSON()
}

func (c StickersCmd) UploadUnknownsStickerPacks(bot *api.API, chatID int,
	unknownsStickerIDs []string, unavailableStickerPacksCount *int,
	existsStickerPacks map[int]*domain.StickerPack,
) error {
	unknownsStickerPackIDsCount := len(unknownsStickerIDs)

	stickerPacksInfo := sharedResponses.StickersReply{}
	if err := c.userBot.Call("store.getProducts", fmt.Sprintf("type=stickers&product_ids=%v&extended=1", strings.Join(unknownsStickerIDs, ",")), &stickerPacksInfo); err != nil {
		fmt.Println("userBot.Call err", err)
		return nil
	}

	infusedStickerPacks := make([]*domain.StickerPack, 0, unknownsStickerPackIDsCount)

	for _, stickerPack := range stickerPacksInfo.Response.Items {
		if infusedStickerPack := utils.FilterOnlyAvailableStickerPacks(stickerPack); infusedStickerPack != nil {
			infusedStickerPacks = append(infusedStickerPacks, infusedStickerPack)
		} else {
			*unavailableStickerPacksCount++
		}
	}

	if len(infusedStickerPacks) > 0 {
		sent, err := bot.SendMessage(chatID, fmt.Sprintf("Выгружаю %d стикер-паков", unknownsStickerPackIDsCount))
		if err != nil {
			return errors.Wrap(err, "SendMessage")
		}

		for _, uploadedStickerPack := range c.uploaderService.UploadStickerPacks(c.userBot, infusedStickerPacks) {
			existsStickerPacks[uploadedStickerPack.ID] = uploadedStickerPack
		}

		if _, err := bot.DeleteMessages(chatID, sent.ChatMessageID, 1); err != nil {
			return err
		}
	}

	return nil
}
