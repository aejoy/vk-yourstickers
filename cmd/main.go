package main

import (
	"log"
	"os"
	"strconv"

	"github.com/aejoy/vk-yourstickers/internal/repositories/cache"
	"github.com/aejoy/vk-yourstickers/internal/repositories/db"
	"github.com/aejoy/vk-yourstickers/internal/scenes"
	"github.com/aejoy/vk-yourstickers/internal/services/stickers"
	"github.com/aejoy/vk-yourstickers/internal/services/uploader"
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/longpoll"
	"github.com/aejoy/vkgo/scene"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Load()

	db, err := db.NewDB(os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	cache, err := cache.NewCache(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	allStickers, err := db.GetAllStickerPacks()
	if err != nil {
		panic(err)
	}

	if err := cache.Load(allStickers); err != nil {
		panic(err)
	}

	user, err := api.New(os.Getenv("USER_TOKEN"))
	if err != nil {
		panic(err)
	}

	bot, err := api.New(os.Getenv("TOKEN"), zap.NewNop())
	if err != nil {
		panic(err)
	}

	albumID, err := strconv.Atoi(os.Getenv("ALBUM_ID"))
	if err != nil {
		panic(err)
	}

	stickersService := stickers.NewStickersService(db, cache)
	scenes := scenes.NewScenes(user,
		uploader.NewUploaderService(bot.ID, albumID, stickersService),
		stickersService,
	)

	log.Fatalln(longpoll.Start(bot, scene.Message(scenes.Message)))
}
