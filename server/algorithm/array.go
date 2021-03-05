package algorithm

// BubbleSort 冒泡排序：相邻的两两对换最好n，最差n^2
func BubbleSort(arr []int) {
	n := len(arr)
	disSort := false
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				disSort = true
			}
		}
		if !disSort {
			return
		}
	}
}

// SelectSort 选择排序：选择最小的排在最后面,n^2,优化后减少一半，但还是很慢
func SelectSort(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i-- {
		minIdx := i
		maxIdx := i
		for j := i + 1; j < n-i; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
				continue
			}

			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		if maxIdx == i && minIdx != n-i-1 {
			arr[maxIdx], arr[n-i-1] = arr[n-i-1], arr[maxIdx]
			arr[minIdx], arr[i] = arr[i], arr[minIdx]
		} else if maxIdx == i && minIdx == n-i-1 {
			arr[minIdx], arr[maxIdx] = arr[maxIdx], arr[maxIdx]
		} else {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
			arr[maxIdx], arr[n-i-1] = arr[n-i-1], arr[maxIdx]
		}
	}
}

// InsertSort 插入某一位，然后不断和前一位比较，复杂度n-1, n^2，数组规模 n 较小的大多数情况下，我们可以使用插入排序，它比冒泡排序，选择排序都快，甚至比任何的排序算法都快。大家都很少使用冒泡、直接选择，直接插入排序算法，因为在有大量元素的无序数列下，这些算法的效率都很低。
func InsertSort(arr []int) {
	n := len(arr)
	for i := 1; i <= n-1; i++ {
		deal := arr[i]
		j := i - 1
		if deal < arr[j] {
			for ; j >= 0 && deal < arr[j]; j-- {
				arr[j+1] = arr[j]
			}
			arr[j+1] = deal
		}
	}
}

// ShellSort 希尔排序，每次按长度一半进行插入排序 希尔排序的时间复杂度大约在这个范围：O(n^1.3)~O(n^2)
func ShellSort(arr []int) {
	n := len(arr)
	for step := n / 2; step >= 1; step /= 2 {
		for i := step; i < n; i += step {
			for j := i - step; j >= 0; j -= step {
				if arr[j+step] < arr[j] {
					arr[j], arr[j+step] = arr[j+step], arr[j]
					continue
				}
				break
			}
		}
	}
}
