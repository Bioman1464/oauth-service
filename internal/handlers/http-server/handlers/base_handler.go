package handlers

import log "github.com/sirupsen/logrus"

type BaseHandler struct {
	Log *log.Logger
}

func NewBaseHandler(log *log.Logger) BaseHandler {
	return BaseHandler{Log: log}
}
