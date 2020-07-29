// +build !bettercap

package main

import (
	"context"
	"errors"
)

func CaptureFromBettercap(_ context.Context, session *CaptureSession) error {
	session.ReportDone()
	return errors.New("bettercap not built")
}



