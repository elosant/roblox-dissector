// +build !nfqueue

package main

import (
	"context"
	"errors"
)

func CaptureFromNFQueue(_ context.Context, session *CaptureSession) error {
	session.ReportDone()
	return errors.New("nfqueue not built")
}



