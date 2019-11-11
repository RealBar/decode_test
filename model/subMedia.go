package model

import (
	"context"
	"decode_test/pkg/app"
	"decode_test/pkg/e"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

//subMedia is media that generated from a media

type SubMediaRef struct {
	MediaID   int64                   `bson:"_id"`
	SubMedias map[string]SubMediaInfo `bson:"sub_media"`
}

type SubMediaInfo struct {
	Status     app.MediaStatus `bson:"status"`
	CreateTime int64           `bson:"ctime"`
	ModifyTime int64           `bson:"mtime"`
}

const SubMediaCollName = "sub_media_red_collection"

func (db *DBWrapper) CreateSubMediaRef(mediaID int64, subMediaID int64) error {
	now := time.Now().Unix()
	opts := options.FindOneAndUpdate().SetUpsert(true)
	info := SubMediaInfo{
		Status:     app.MediaStatusNormal,
		CreateTime: now,
		ModifyTime: now,
	}

	result := db.getColl(SubMediaCollName).
		FindOneAndUpdate(context.Background(),
			bson.D{
				{"_id", mediaID},
			},
			bson.D{
				{"$set",
					bson.D{
						{"sub_media." + strconv.FormatInt(subMediaID, 10), info},
					},
				},
			}, opts)
	err := result.Err()
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	return nil
}

func (db *DBWrapper) DeleteSubMediaRef(mediaID int64, subMediaID int64) error {
	now := time.Now().Unix()
	info := SubMediaInfo{
		Status:     app.MediaStatusDeleted,
		ModifyTime: now,
	}
	opts := options.FindOneAndUpdate().SetUpsert(false)
	result := db.getColl(SubMediaCollName).FindOneAndUpdate(context.Background(),
		bson.D{
			{"_id", mediaID},
			{"$set",
				bson.D{
					{"sub_media." + strconv.FormatInt(subMediaID, 10), info},
				},
			},
		}, opts)
	err := result.Err()
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	return nil
}

func (db *DBWrapper) QuerySubMediaIDs(mediaID int64) ([]int64, error) {
	var subMedia = new(SubMediaRef)
	var err error
	res := db.getColl(SubMediaCollName).FindOne(context.Background(), bson.D{
		{"_id", mediaID},
	})
	err = res.Decode(subMedia)
	if err != nil {
		return nil, err
	}
	subMediaIDs := make([]int64, len(subMedia.SubMedias))
	var i, errCount int
	var id int64
	for k, _ := range subMedia.SubMedias {
		id, err = strconv.ParseInt(k, 10, 64)
		if err != nil {
			logrus.WithError(err).Error("subMedia id parse error:" + k)
			errCount++
		} else {
			subMediaIDs[i] = id
		}
		i++
	}
	if i == errCount {
		return subMediaIDs, e.ErrAllMediaIDInvalid
	}
	return subMediaIDs, nil
}
