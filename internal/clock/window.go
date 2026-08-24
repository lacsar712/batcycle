package clock

import (
	"time"

	"github.com/lacsar712/batcycle/internal/model"
)

type AvgTempWindow struct {
	clk      Clock
	duration time.Duration
}

func NewAvgTempWindow(clk Clock, duration time.Duration) *AvgTempWindow {
	if duration <= 0 {
		duration = 2 * time.Minute
	}
	return &AvgTempWindow{clk: clk, duration: duration}
}

func (w *AvgTempWindow) Active(anchor time.Time) bool {
	// Measure elapsed against the injected process clock, not the wall clock:
	// the hold window must follow the (freezable) process clock, so that freezing
	// the process clock also freezes window progress.
	return w.clk.Now().Sub(anchor) < w.duration
}

func (w *AvgTempWindow) Require(anchor time.Time) error {
	if w.Active(anchor) {
		return nil
	}
	return model.ErrGradientHold
}
