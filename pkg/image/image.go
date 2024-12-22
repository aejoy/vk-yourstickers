package image

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"net/http"

	"github.com/pkg/errors"

	"github.com/aejoy/vk-yourstickers/pkg/consts"
	"github.com/nfnt/resize"
)

func FetchImage(url string) (image.Image, error) {
	timeout, cancel := context.WithTimeout(context.Background(), consts.DefaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(timeout, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "FetchImage")
	}

	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "DecodeImage")
	}

	return img, nil
}

func ResizeImage(img image.Image, newHeight, newWidth uint) image.Image {
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	return resizedImg
}

func ToJPEG(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: consts.CarouselPhotoQuality}); err != nil {
		return nil, errors.Wrap(err, "EncodeImage")
	}

	return buf.Bytes(), nil
}

func Fetch(url string) ([]byte, error) {
	img, err := FetchImage(url)
	if err != nil {
		return nil, err
	}

	jpegImg, err := ToJPEG(ResizeImage(img, consts.CarouselPhotoHeight, consts.CarouselPhotoWidth))
	if err != nil {
		return nil, err
	}

	return jpegImg, nil
}
