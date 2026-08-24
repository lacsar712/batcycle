package app

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lacsar712/batcycle/internal/config"
	"github.com/lacsar712/batcycle/internal/model"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	anchor := a.clk.Now().Add(-3 * time.Minute)
	err = a.ConfirmGradientHold(context.Background(), anchor)
	if err == nil {
		t.Fatal("expected gradient hold error")
	}
	if !errors.Is(err, model.ErrGradientHold) {
		t.Fatalf("expected ErrGradientHold, got %v", err)
	}
	_ = time.Second
}
