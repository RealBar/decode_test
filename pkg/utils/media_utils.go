package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"io"
	"io/ioutil"
)

func ShrinkImage(src image.Image, width int, height int) image.Image {
	return imaging.Resize(src, width, height, imaging.NearestNeighbor)
}

func GenerateMD5(data io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(bytes)), nil
}
