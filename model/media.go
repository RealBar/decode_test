package model

import (
	"decode_test/pkg/generator"
	"time"
)

type MediaStatus uint8
type MediaType uint8
type MediaFormat uint8

const (
	_ = MediaStatus(iota)
	MediaStatusNormal
	MediaStatusDeleted
)
const (
	_ = MediaType(iota)
	ImageType
	VideoType
	AudioType
	TextType
)

const (
	_ = MediaFormat(iota)
	Jpeg
	Gif
	Png
	Mp4
	Mp3
	SimpleText
)

type Media struct {
	ID         int64       `gorm:"id"`
	OwnerID    int64       `gorm:"owner_id"`
	CreateTime int64       `gorm:"ctime"`
	ModifyTime int64       `gorm:"mtime"`
	Size       int64       `gorm:"size"`
	Name       string      `gorm:"name"`
	MD5        string      `gorm:"md5"`
	Type       MediaType   `gorm:"type"`
	Format     MediaFormat `gorm:"format"`
	Status     MediaStatus `gorm:"status"`
	ExtInfo    string      `gorm:"ext_info"`
}

func (db *DBWrapper) CreateMedia(name string, mType MediaType, mFormat MediaFormat, MD5 string, size int64, ownerID int64,
	extInfo string) error {
	id := generator.GenerateID(ownerID)
	now := time.Now().Unix()
	media := &Media{
		ID:         id,
		Name:       name,
		Size:       size,
		Type:       mType,
		Format:     mFormat,
		Status:     MediaStatusNormal,
		MD5:        MD5,
		CreateTime: now,
		ModifyTime: now,
		ExtInfo:    extInfo,
	}
	return db.mysqlDB.Create(media).Error
}
