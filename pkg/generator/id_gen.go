package generator

import (
	"decode_test/pkg/app"
	"decode_test/pkg/cache"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ID:
// OwnerID: 11 bits
// timestamp:24 bits
// IDC ID:4 bits
// process id:12 bits
// random num: 12 bits

type IDGenerator struct {
	ProcessID int64
	IdcID     int64
	RandGen   rand.Source
}

func Setup(logger *logrus.Logger, cache *cache.Cache) *IDGenerator {
	var idGenerator = new(IDGenerator)
	var idcID = app.DefaultIdcID
	var err error
	idcStr, b := os.LookupEnv(app.ConfigIdcKey)
	if !b {
		logger.Error("idc id not found in environment, using default IdcID:", app.DefaultIdcID)
	} else {
		idcID, err = strconv.ParseInt(idcStr, 10, 64)
		if err != nil {
			logger.Error("idc id invalid:", app.DefaultIdcID)
			panic(err)
		}
	}
	idGenerator.IdcID = idcID & 0x0f

	processID, err := cache.IncrBy(app.RedisProcessKey, 1)
	if err != nil {
		logger.Error("get process id error")
		panic(err)
	}
	idGenerator.ProcessID = processID & 0xfff
	idGenerator.RandGen = rand.NewSource(time.Now().Unix())
	return idGenerator
}

func (g *IDGenerator) GenerateID(ownerID int64) int64 {
	nowInMillis := int64(time.Now().Nanosecond() * 1000_000 & 0xffffff)
	ownerID = ownerID & 0x7f
	randNum := g.RandGen.Int63() & 0xfff
	return ownerID<<52 | nowInMillis<<28 | g.IdcID<<24 | g.ProcessID<<12 | randNum
}
