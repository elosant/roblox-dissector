// +build bettercap

package main

import (
	"context"
	"fmt"
)

func CaptureFromBettercap(_ context.Context, _ *CaptureSession) error {
	fmt.Println("Called bettercap callback")
	return nil
}
