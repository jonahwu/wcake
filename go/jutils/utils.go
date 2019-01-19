package jutils

import (
	"github.com/satori/go.uuid"
	"os"
)

func GetHostName() string {
	hn, _ := os.Hostname()
	return hn

}

func GetUuid() string {
	return uuid.Must(uuid.NewV4()).String()
}
