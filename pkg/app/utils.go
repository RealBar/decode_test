package app

import "image"

func GetMediaFormat(formatStr string) MediaFormat {
	var res = Unknown
	switch formatStr {
	case "png":
		res = Png
	case "jpeg":
		res = Jpeg
	case "gif":
		res = Gif
	}
	return res
}

func ThumbnailSize(bound image.Rectangle) (int, int, bool) {
	srcW := bound.Dx()
	srcH := bound.Dy()
	var dstW, dstH int
	var needShrink bool
	if srcH > ThumbnailThreshold {
		needShrink = true
		if srcW == srcH {
			dstH = ThumbnailThreshold
			dstW = ThumbnailThreshold
		} else if srcH > srcW {
			dstH = ThumbnailThreshold
			dstW = (srcH - ThumbnailThreshold) * srcW / srcH
		} else {
			dstW = ThumbnailThreshold
			dstH = (srcW - ThumbnailThreshold) * srcH / srcW
		}
	} else if srcW > ThumbnailThreshold {
		needShrink = true
		dstW = ThumbnailThreshold
		dstH = (srcW - ThumbnailThreshold) * srcH / srcW
	}
	return dstW, dstH, needShrink
}
