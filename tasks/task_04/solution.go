package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	var stats Stats
	for i := 1; i < len(nums); i++ {
		stats.Count++
		now := nums[i] - nums[i-1]
		if now > stats.Max || i == 1 {
			stats.Max = now
		}
		if now < stats.Min || i == 1 {
			stats.Min = now
		}
		stats.Sum += now
	}
	return stats
}
