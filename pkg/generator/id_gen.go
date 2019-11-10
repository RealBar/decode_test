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

var idGenerator struct {
	ProcessID int64
	IdcID     int64
	RandGen   rand.Source
}

func Setup() {
	var idcID = app.DefaultIdcID
	var err error
	idcStr, b := os.LookupEnv(app.CONFIG_IDC_KEY)
	if !b {
		logrus.Error("idc id not found in environment, using default IdcID:", app.DefaultIdcID)
	} else {
		idcID, err = strconv.ParseInt(idcStr, 10, 64)
		if err != nil {
			logrus.Error("idc id invalid:", app.DefaultIdcID)
			panic(err)
		}
	}
	idGenerator.IdcID = idcID & 0x0f

	processID, err := cache.IncrBy(app.REDIS_PROCESS_KEY, 1)
	if err != nil {
		logrus.Error("get process id error")
		panic(err)
	}
	idGenerator.ProcessID = processID & 0xfff
	idGenerator.RandGen = rand.NewSource(time.Now().Unix())
}

func GenerateID(ownerID int64) int64 {
	nowInMillis := int64(time.Now().Nanosecond() * 1000_000 & 0xffffff)
	ownerID = ownerID & 0x7f
	randNum := idGenerator.RandGen.Int63() & 0xfff
	return ownerID<<52 | nowInMillis<<28 | idGenerator.IdcID<<24 | idGenerator.ProcessID<<12 | randNum
}
