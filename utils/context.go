package utils

import (
	"context"
	"time"

	"gogenggo/internals/types/constants"
)

func SetProccessTimeCtx(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, constants.StartTimeCtx, time.Now())
	return ctx
}

func GetStartTimeProcessCtx(ctx context.Context) time.Time {
	startTime := ctx.Value(constants.StartTimeCtx)
	if startTime != nil {
		return startTime.(time.Time)
	}

	return time.Now()
}
