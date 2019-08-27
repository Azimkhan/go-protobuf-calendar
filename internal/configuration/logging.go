package configuration

import (
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
)

import "log"

func CreateLogger(config string) *zap.Logger {
	zapConf := zap.Config{}
	rawJson, e := ioutil.ReadFile(config)
	if e != nil {
		log.Fatal(e)
	}
	e = json.Unmarshal(rawJson, &zapConf)
	if e != nil {
		log.Fatal(e)
	}
	logger, e := zapConf.Build()
	if e != nil {
		log.Fatal(e)
	}
	return logger
}
