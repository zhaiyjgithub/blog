package main

import (
	"testing"
)

func TestNewSqsService(t *testing.T) {
	n := "SenderWorker"
	NewSqsService(n)
}
