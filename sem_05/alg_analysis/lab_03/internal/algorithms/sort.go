package algorithms

func Quicksort(arr []int) {
	MemoryInfo.Reset()

	quicksort(arr)
}

func quicksort(arr []int) {
	MemoryInfo.Add(memoryQuick(arr))
	defer MemoryInfo.Done(memoryQuick(arr))

	n := len(arr)

	less := make([]int, 0, n)
	greater := make([]int, 0, n)
	pivot := 0

	if n > 1 {
		pivot = arr[n/2]

		for i, v := range arr {
			if i != n/2 {
				if v < pivot {
					less = append(less, v)
				} else {
					greater = append(greater, v)
				}
			}
		}

		quicksort(less)
		quicksort(greater)

		l := len(less)
		g := len(greater)

		for i := 0; i < l; i++ {
			arr[i] = less[i]
		}
		arr[l] = pivot
		for i := l + 1; i < l+g+1; i++ {
			arr[i] = greater[i-l-1]
		}
	}
}

func flip(arr []int) {
	n := len(arr)

	for left := 0; left < n; left++ {
		arr[left], arr[n-1] = arr[n-1], arr[left]
		n--
	}
}

func getIndMax(arr []int) int {
	iMax := 0

	for i := range arr {
		if arr[i] > arr[iMax] {
			iMax = i
		}
	}

	return iMax
}

func Pancakesort(arr []int) {
	MemoryInfo.Reset()
	MemoryInfo.Add(memoryPancake(arr))
	defer MemoryInfo.Done(memoryPancake(arr))

	for n := len(arr); n > 1; n-- {
		iMax := getIndMax(arr[:n])
		if iMax != n-1 {
			flip(arr[:(iMax + 1)])
			flip(arr[:n])
		}
	}
}

func Beadsort(arr []int) {
	MemoryInfo.Reset()
	MemoryInfo.Add(memoryBead(arr))
	defer MemoryInfo.Done(memoryBead(arr))

	n := len(arr)
	max := arr[getIndMax(arr)]

	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, max)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < arr[i]; j++ {
			m[i][j]++
		}
	}

	for j := 0; j < max; j++ {
		beadsInColumn := 0

		for i := 0; i < n; i++ {
			if m[i][j] == 1 {
				beadsInColumn++
				m[i][j] = 0
			}
		}

		for i := n - beadsInColumn; i < n; i++ {
			m[i][j] = 1
		}
	}

	for i := 0; i < n; i++ {
		beadsInRow := 0

		for j := 0; j < max; j++ {
			beadsInRow += m[i][j]
		}

		arr[i] = beadsInRow
	}
}
