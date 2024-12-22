package consts

import (
	"errors"
	"regexp"
	"time"
)

const (
	Day                  = 86400
	DefaultTimeout       = 5 * time.Second
	PacksPerPage         = 10
	CarouselPhotoHeight  = uint(544)
	CarouselPhotoWidth   = uint(884)
	CarouselPhotoQuality = 100
	SingleArgsCount      = 1
	TwoPairArgsCount     = 2
	MinScreenNameLength  = 3
	UnknownsNotEmpty     = 0
)

var (
	ScreenNameRegex          = regexp.MustCompile(`(?i)\[(id|club)([0-9]+)\|(.*)]`)
	ErrPageNotFound          = errors.New("page not found")
	ErrNotAvailableForGroups = errors.New("not available for groups")
)
