package resources

import (
	"log"
	"sort"
)

/*
There is no update activity, just use delete and add activity for now.
*/

// TimeFrame Defines a timeframe using Unix Epoch time start and end seconds.
type TimeFrame struct {
	UnixStartTime int64
	UnixEndTime   int64
}

// Activity Consolidates activities for resource.
type Activity struct {
	ActionID int64
	TimeFrame
}

// Resource Consolidates resource availability.
type Resource struct {
	overlapIndex int
	Activities   []Activity
	BusyInterval []TimeFrame
}

// NewResource Constructor for resources.
func NewResource() *Resource {
	return &Resource{
		Activities:   []Activity{},
		BusyInterval: []TimeFrame{},
	}
}

// AddActivity Method adds activity to a resource.
func (r *Resource) AddActivity(a ...Activity) {
	for _, v := range a {
		r.Activities = append(r.Activities, v)
	}
}

// DeleteActivity Method deletes activity from a resource.
func (r *Resource) DeleteActivity(a Activity) error {
	return nil
}

// updateBusyTime Method updates busy time once a change in activities occurs.
func (r *Resource) updateBusyTime() error {
	r.sortActivities()
	r.syncBusy()

	for r.hasOverlapping() {
		log.Println("overlapping index:", r.overlapIndex)

		newTimeFrame := mergeOverlapping(r.BusyInterval[r.overlapIndex], r.BusyInterval[r.overlapIndex+1])
		r.BusyInterval = r.removeTimeFrame(r.overlapIndex + 1)
		r.BusyInterval[r.overlapIndex] = newTimeFrame
	}

	log.Println("updated busy intervals:", r.BusyInterval)

	return nil
}

// sortActivities Method sorts activities once a change in activities occurs.
func (r *Resource) sortActivities() {
	sort.Slice(r.Activities, func(i, j int) bool {
		return r.Activities[i].UnixStartTime < r.Activities[j].UnixStartTime
	})
}

func (r *Resource) syncBusy() {
	r.BusyInterval = make([]TimeFrame, len(r.Activities)) // reset slice

	for ix, v := range r.Activities {
		r.BusyInterval[ix] = v.TimeFrame
	}
	log.Println("filled up busy intervals:", r.BusyInterval)
}

// hasOverlapping Helper given ordered by start time slice gives indication if it has overlapping time frames.
func (r *Resource) hasOverlapping() bool {
	for i := 0; i < len(r.BusyInterval)-1; i++ {
		if areOverlapping(r.BusyInterval[i], r.BusyInterval[i+1]) {
			r.overlapIndex = i
			return true
		}
	}
	return false
}

func (r *Resource) removeTimeFrame(index int) []TimeFrame {
	return append(r.BusyInterval[:index], r.BusyInterval[index+1:]...)
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

// areOverlapping Helper should provide if two ordered by start time time frames
// are overlapping.
func areOverlapping(t1, t2 TimeFrame) bool {
	log.Println("checking overlapping for:", t1.UnixEndTime, t2.UnixStartTime)
	if t1.UnixEndTime >= t2.UnixStartTime {
		return true
	}
	return false
}

func mergeOverlapping(t1, t2 TimeFrame) TimeFrame {
	return TimeFrame{
		UnixStartTime: t1.UnixStartTime,
		UnixEndTime:   max(t1.UnixEndTime, t2.UnixEndTime),
	}
}
