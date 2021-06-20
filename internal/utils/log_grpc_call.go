package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LogGrpcCall logs gRPC calls
func LogGrpcCall(funcName string, req, res interface{}, err *error) {

	var logEvent *zerolog.Event

	if *err != nil {
		logEvent = log.Error()
	} else {
		logEvent = log.Debug()
	}

	logEvent.
		Err(*err).
		Interface("Request", req).
		Interface("Response", res).
		Msg(funcName)
}
