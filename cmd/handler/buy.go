package handler

import (
	//"github.com/golang/glog"
	"github.com/wcake/go/jutils"
	"time"
)

type Buy struct {
	BuyID   string `json:"buyid"`
	BuyTime string `json:"buytime"`
}

func getBuyID() string {
	return jutils.GetUuid()
}

func getBuyTime() string {
	return time.Now().Format("20060102150405")
}
