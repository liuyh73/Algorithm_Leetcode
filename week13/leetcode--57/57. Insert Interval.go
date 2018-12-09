/**
 * Definition for an interval.
 * type Interval struct {
 *     Start int
 *     End   int
 * }
 */
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	intervals_merged := make([]Interval, 0)
	flag := false
	var i int
	var interval Interval
	for i, interval = range intervals {
		if interval.End < newInterval.Start {
			intervals_merged = append(intervals_merged, interval)
		} else if interval.Start > newInterval.End {
			flag = true
			intervals_merged = append(intervals_merged, newInterval)
			break
		} else {
			newInterval.Start = min(interval.Start, newInterval.Start)
			newInterval.End = max(interval.End, newInterval.End)
		}
	}
	if flag == true {
		intervals_merged = append(intervals_merged, intervals[i:]...)
	} else {
		intervals_merged = append(intervals_merged, newInterval)
	}
	return intervals_merged
}