package model

import (
	"decode_test/pkg/app"
	"time"
)

type Media struct {
	ID         int64           `gorm:"id"`
	OwnerID    int64           `gorm:"owner_id"`
	CreateTime int64           `gorm:"ctime"`
	ModifyTime int64           `gorm:"mtime"`
	Size       int64           `gorm:"size"`
	Name       string          `gorm:"name"`
	MD5        string          `gorm:"md5"`
	Type       app.MediaType   `gorm:"type"`
	Format     app.MediaFormat `gorm:"format"`
	Status     app.MediaStatus `gorm:"status"`
	ExtInfo    string          `gorm:"ext_info"`
}

func (db *DBWrapper) CreateMedia(c *app.ApplicationContext, name string, mType app.MediaType, mFormat app.MediaFormat,
	MD5 string, size int64, ownerID int64, extInfo string) error {
	id := c.Gen().GenerateID(ownerID)
	now := time.Now().Unix()
	media := &Media{
		ID:         id,
		Name:       name,
		Size:       size,
		Type:       mType,
		Format:     mFormat,
		Status:     app.MediaStatusNormal,
		MD5:        MD5,
		CreateTime: now,
		ModifyTime: now,
		ExtInfo:    extInfo,
	}
	return db.mysqlDB.Create(media).Error
}
