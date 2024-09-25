package interval

import (
	"cmp"
	"slices"
)

type Interval[T cmp.Ordered] struct {
	start T
	end   T
}

func New[T cmp.Ordered](start, end T) Interval[T] {
	if end < start {
		panic("interval start values must be less than the end value")
	}

	return Interval[T]{
		start: start,
		end:   end,
	}
}

func (this *Interval[T]) Start() T {
	return this.start
}

func (this *Interval[T]) End() T {
	return this.start
}

// Overlap returns true if this Interval overlaps with the other interval;
// otherwise, false.
func (this *Interval[T]) Overlap(other Interval[T]) bool {
	return this.end >= other.start && other.end >= this.start
}

// Merge merges this interval with the other interval and returns the new merged interval.
// It also returns true if the intervals overlapped and were merged; otherwise, false.
func (this *Interval[T]) Merge(other Interval[T]) (Interval[T], bool) {
	if !this.Overlap(other) {
		bad := new(T)
		return New(*bad, *bad), false
	}

	return New(min(this.start, other.start), max(this.end, other.end)), true
}

// StartsBefore returns true if this interval starts before the other interval;
// otherwise, false.
func (this *Interval[T]) StartsBefore(other Interval[T]) bool {
	return this.start < other.start
}

// StartsAfter returns true if this interval starts after the other interval;
// otherwise, false.
func (this *Interval[T]) StartsAfter(other Interval[T]) bool {
	return this.start > other.start
}

// EndsBefore returns true if this interval ends before the other interval;
// otherwise, false.
func (this *Interval[T]) EndsBefore(other Interval[T]) bool {
	return this.end < other.end
}

// EndsAfter returns true if this interval ends after the other interval;
// otherwise, false.
func (this *Interval[T]) EndsAfter(other Interval[T]) bool {
	return this.end > other.end
}

// MergeAll merges all overlapping intervals in the given collection of intervals.
func MergeAll[T cmp.Ordered](intervals ...Interval[T]) []Interval[T] {
	if len(intervals) <= 1 {
		return intervals
	}

	merged := []Interval[T]{}

	// So, if we don't sort the intervals, then each interval must be
	// visited N^2 times. If we sort by start time NlogN, this'll save us work.
	compare := func(a, b Interval[T]) int {
		if a.StartsBefore(b) {
			return -1
		} else if a.StartsAfter(b) {
			return 1
		}

		return 0
	}

	slices.SortFunc(intervals, compare)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		curr := merged[len(merged)-1]

		if new, ok := curr.Merge(intervals[i]); ok {
			// These intervals overlapped and created a new one;
			// assign the new merged interval to the last of merged.
			merged[len(merged)-1] = new
		} else {
			// The intervals did not overlap
			// Because we're sorted by start time, add a new
			// interval to our merged collection
			merged = append(merged, intervals[i])
		}
	}

	return merged
}
