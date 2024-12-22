package db

import (
	"context"

	"github.com/aejoy/go-pkg/postgres"
	"github.com/aejoy/vk-yourstickers/internal/domain"
	"github.com/aejoy/vk-yourstickers/pkg/consts"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	conn *pgxpool.Pool
}

func NewDB(url string) (*DB, error) {
	db, err := postgres.NewPostgres([]string{url})
	if err != nil {
		return nil, err
	}

	return &DB{db[0]}, nil
}

func (db *DB) GetAllStickerPacks() ([]domain.StickerPack, error) {
	stickerPacks := make([]domain.StickerPack, 0)

	timeout, cancel := context.WithTimeout(context.Background(), consts.DefaultTimeout)
	defer cancel()

	rows, err := db.conn.Query(timeout, "SELECT id, title, preview_url FROM sticker_pack;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		stickerPack := domain.StickerPack{}
		if err := rows.Scan(&stickerPack.ID, &stickerPack.Title, &stickerPack.PreviewURL); err != nil {
			return nil, err
		}

		stickerPacks = append(stickerPacks, stickerPack)
	}

	return stickerPacks, nil
}

func (db *DB) GetStickerPacks(stickerPackIDs []int) (map[int]*domain.StickerPack, error) {
	stickerPacks := map[int]*domain.StickerPack{}

	timeout, cancel := context.WithTimeout(context.Background(), consts.DefaultTimeout)
	defer cancel()

	rows, err := db.conn.Query(timeout, "SELECT id, title, preview_url FROM sticker_pack WHERE id = ANY($1)", stickerPackIDs)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		stickerPack := new(domain.StickerPack)
		if err := rows.Scan(&stickerPack.ID, &stickerPack.Title, &stickerPack.PreviewURL); err != nil {
			return nil, err
		}

		stickerPacks[stickerPack.ID] = stickerPack
	}

	return stickerPacks, nil
}

func (db *DB) CreateStickerPack(id int, title, previewURL string) error {
	timeout, cancel := context.WithTimeout(context.Background(), consts.DefaultTimeout)
	defer cancel()

	_, err := db.conn.Exec(timeout, "INSERT INTO sticker_pack(id, title, preview_url) VALUES ($1, $2, $3);", id, title, previewURL)

	return err
}
