package store

import (
	"time"

	"github.com/lacsar712/batcycle/internal/model"
)

type ScheduleStore struct {
	mem *Memory
}

func NewScheduleStore(mem *Memory) *ScheduleStore {
	return &ScheduleStore{mem: mem}
}

func (s *ScheduleStore) Save(sched model.DryingSchedule) {
	s.mem.PutSchedule(sched)
}

func (s *ScheduleStore) Load(id model.ScheduleID) (model.DryingSchedule, error) {
	sched, ok := s.mem.GetSchedule(id)
	if !ok {
		return model.DryingSchedule{}, model.Wrap("schedule", "load", model.ErrNotFound)
	}
	return sched, nil
}

func (s *ScheduleStore) SnapshotClone(id model.ScheduleID) (model.DryingSchedule, error) {
	sched, err := s.Load(id)
	if err != nil {
		return model.DryingSchedule{}, err
	}
	return sched.Clone(), nil
}

func (s *ScheduleStore) ActiveEntry(sched model.DryingSchedule, now time.Time) (model.DryingScheduleEntry, bool) {
	for _, e := range sched.Entries {
		if !now.Before(e.Start) && now.Before(e.End) {
			return e, true
		}
	}
	return model.DryingScheduleEntry{}, false
}

// PendingEntry reports the next schedule entry whose soak window has not yet
// opened (now < Start), if any. A pending entry means the schedule is not
// empty: an equalization soak is merely not yet due, which must not be
// confused with a schedule that has been emptied.
func (s *ScheduleStore) PendingEntry(sched model.DryingSchedule, now time.Time) (model.DryingScheduleEntry, bool) {
	var pending model.DryingScheduleEntry
	found := false
	for _, e := range sched.Entries {
		if now.Before(e.Start) {
			if !found || e.Start.Before(pending.Start) {
				pending = e
				found = true
			}
		}
	}
	return pending, found
}

func (s *ScheduleStore) EntriesOverlapping(sched model.DryingSchedule, start, end time.Time) []model.DryingScheduleEntry {
	var out []model.DryingScheduleEntry
	for _, e := range sched.Entries {
		if e.End.After(start) && e.Start.Before(end) {
			out = append(out, e)
		}
	}
	return out
}
