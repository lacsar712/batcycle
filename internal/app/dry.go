package app

import (
	"context"
	"time"

	"github.com/lacsar712/batcycle/internal/clock"
)

type DryRamp struct {
	clk   clock.Clock
	tick  time.Duration
	steps int
}

func NewDryRamp(clk clock.Clock, tick time.Duration, steps int) *DryRamp {
	if steps <= 0 {
		steps = 40
	}
	return &DryRamp{clk: clk, tick: tick, steps: steps}
}

func (r *DryRamp) Ramp(ctx context.Context, target float64, apply func(float64)) error {
	step := target / float64(r.steps)
	if step <= 0 {
		step = 0.5
	}
	cur := 0.0
	for cur < target {
		// Honor batch abort: a canceled context means the charge ramp must
		// stop chasing target immediately, otherwise the tail of the ramp
		// keeps creeping past the abort point on the sampling curve.
		if err := ctx.Err(); err != nil {
			return err
		}
		cur += step
		if cur > target {
			cur = target
		}
		apply(cur)
		if pc, ok := r.clk.(*clock.ProcessClock); ok {
			pc.Step()
		}
		time.Sleep(2 * time.Millisecond)
	}
	return nil
}

func (a *App) RunDryRamp(ctx context.Context, target float64) error {
	return a.dryRamp.Ramp(ctx, target, func(v float64) { _ = v })
}
