package main

import (
	"context"
	"time"
)

func locale(ctx context.Context) (string, error) {
	//if deadline, ok := ctx.Deadline(); ok {
	//	if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
	//		return "", context.DeadlineExceeded
	//	}
	//}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}
