package utils

import "github.com/aejoy/vk-yourstickers/pkg/consts"

func GetPaginationBounds(countPerPage, itemsCount, page int) (int, int, int, error) {
	pages := itemsCount / countPerPage

	if page > pages {
		return pages, 0, 0, consts.ErrPageNotFound
	}

	offset := (page - 1) * countPerPage
	end := offset + countPerPage

	return pages, offset, end, nil
}
