package util

import (
	"sort"
	"time"
)

const (
	NormalActionDuration time.Duration = time.Minute * 15   // 15'
	AccessDuration       time.Duration = time.Hour * 24     // 1 ngày
	RefreshDuration      time.Duration = AccessDuration * 7 // 1 tuần
)

func GetPrimitiveTime() time.Time {
	// 1/1/1900 - 00:00:00
	return time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
}

func IsActionExpired(exp time.Time) bool {
	return time.Now().After(exp)
}

// Generic function to sort any slice of structs based on a time.Time field
func SortByTime[T any](items []T, getTime func(T) time.Time, ascending bool) {
	sort.Slice(items, func(i, j int) bool {
		if ascending {
			return getTime(items[i]).Before(getTime(items[j]))
		}

		return getTime(items[i]).After(getTime(items[j]))
	})
}
