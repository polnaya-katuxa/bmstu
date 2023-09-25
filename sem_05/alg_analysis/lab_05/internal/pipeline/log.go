package pipeline

import (
	"fmt"
	"time"
)

func LogTasks(tasks []Task) {
	startSys := tasks[0].TimeFirst

	sumBefore1 := time.Duration(0).Microseconds()
	sumBefore2 := time.Duration(0).Microseconds()
	sumBefore3 := time.Duration(0).Microseconds()

	minBefore1 := tasks[0].TimeStart1.Sub(tasks[0].TimeFirst).Microseconds()
	minBefore2 := tasks[0].TimeStart2.Sub(tasks[0].TimeEnd1).Microseconds()
	minBefore3 := tasks[0].TimeStart3.Sub(tasks[0].TimeEnd2).Microseconds()

	maxBefore1 := tasks[0].TimeStart1.Sub(tasks[0].TimeFirst).Microseconds()
	maxBefore2 := tasks[0].TimeStart2.Sub(tasks[0].TimeEnd1).Microseconds()
	maxBefore3 := tasks[0].TimeStart3.Sub(tasks[0].TimeEnd2).Microseconds()

	totalSum := int64(0)
	totalMax := tasks[0].TimeEnd3.Sub(tasks[0].TimeFirst).Microseconds()
	totalMin := tasks[0].TimeEnd3.Sub(tasks[0].TimeFirst).Microseconds()

	for i, v := range tasks {
		t1 := tasks[i].TimeStart1.Sub(tasks[i].TimeFirst).Microseconds()
		t2 := tasks[i].TimeStart2.Sub(tasks[i].TimeEnd1).Microseconds()
		t3 := tasks[i].TimeStart3.Sub(tasks[i].TimeEnd2).Microseconds()

		sumBefore1 += t1
		sumBefore2 += t2
		sumBefore3 += t3

		if t1 < minBefore1 {
			minBefore1 = t1
		}
		if t2 < minBefore2 {
			minBefore2 = t2
		}
		if t3 < minBefore3 {
			minBefore3 = t3
		}

		if t1 > maxBefore1 {
			maxBefore1 = t1
		}
		if t2 > maxBefore2 {
			maxBefore2 = t2
		}
		if t3 > maxBefore3 {
			maxBefore3 = t3
		}

		fmt.Printf("TASK №%d: \n", i)
		fmt.Printf("Line 1:\n")
		fmt.Printf("\tstart %v\n", v.TimeStart1.Sub(startSys).Microseconds())
		fmt.Printf("\tend %v\n", v.TimeEnd1.Sub(startSys).Microseconds())
		fmt.Printf("Line 2:\n")
		fmt.Printf("\tstart %v\n", v.TimeStart2.Sub(startSys).Microseconds())
		fmt.Printf("\tend %v\n", v.TimeEnd2.Sub(startSys).Microseconds())
		fmt.Printf("Line 3:\n")
		fmt.Printf("\tstart %v\n", v.TimeStart3.Sub(startSys).Microseconds())
		fmt.Printf("\tend %v\n", v.TimeEnd3.Sub(startSys).Microseconds())

		total := v.TimeEnd3.Sub(v.TimeFirst).Microseconds()
		totalSum += total
		if total > totalMax {
			totalMax = total
		}
		if total < totalMin {
			totalMin = total
		}
	}

	avgBefore1 := sumBefore1 / int64(len(tasks))
	avgBefore2 := sumBefore2 / int64(len(tasks))
	avgBefore3 := sumBefore3 / int64(len(tasks))

	fmt.Println()
	fmt.Printf("Line 1:\n")
	fmt.Printf("\tsum %v\n", sumBefore1)
	fmt.Printf("\tavg %v\n", avgBefore1)
	fmt.Printf("\tmin %v\n", minBefore1)
	fmt.Printf("\tmax %v\n", maxBefore1)
	fmt.Printf("Line 2:\n")
	fmt.Printf("\tsum %v\n", sumBefore2)
	fmt.Printf("\tavg %v\n", avgBefore2)
	fmt.Printf("\tmin %v\n", minBefore2)
	fmt.Printf("\tmax %v\n", maxBefore2)
	fmt.Printf("Line 3:\n")
	fmt.Printf("\tsum %v\n", sumBefore3)
	fmt.Printf("\tavg %v\n", avgBefore3)
	fmt.Printf("\tmin %v\n", minBefore3)
	fmt.Printf("\tmax %v\n", maxBefore3)
	fmt.Println()
	t := tasks[len(tasks)-1].TimeEnd3.Sub(tasks[0].TimeFirst).Microseconds()
	fmt.Printf("Total system time: %v\n", t)
	fmt.Printf("\tavg %v\n", totalSum/int64(len(tasks)))
	fmt.Printf("\tmin %v\n", totalMin)
	fmt.Printf("\tmax %v\n", totalMax)
}
