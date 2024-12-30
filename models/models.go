package models

import "github.com/vukovlevi/battleship/server/logger"

var Log *logger.Logger

func SetLogger(log *logger.Logger) {
    Log = log
}
