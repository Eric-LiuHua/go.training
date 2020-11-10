package sorts

// 冒泡排序
func Sort_Bubble(arr []int) []int {
	arr_len := len(arr)
	if arr_len > 1 {
		for i := arr_len - 1; i > 0; i-- {
			for j := 1; j <= i; j++ {
				if arr[j-1] > arr[j] {
					arr[j-1], arr[j] = arr[j], arr[j-1]
				}
			}
		}
	}
	return arr
}

//计数排序
func Sort_Counter(arr []int) []int {
	var arr_len int = len(arr)
	if arr_len > 1 {
		var min int = arr[0]
		var max int = arr[0]

		for _, a := range arr {
			if a > max {
				max = a
			}
			if a < min {
				min = a
			}
		}
		//Go用切片slice来构造动态数组
		var counter []int = make([]int, max-min+1)

		for _, a := range arr {
			counter[a-min]++
		}
		var cur int = 0
		for i := 0; i < max-min+1; i++ {
			if counter[i] > 0 {
				for {
					if counter[i] <= 0 {
						break
					}
					arr[cur] = i + min
					counter[i] -= 1
					cur += 1

				}
			}
		}

	}

	return arr
}

//归并排序
func Sort_Merge(arr []int) []int {
	arr_len := len(arr)
	if arr_len > 1 {
		var tmp []int = make([]int, arr_len)
		sort_Merge_Part(arr, tmp, 0, arr_len-1)
	}
	return arr
}
func sort_Merge_Part(arr []int, tmp []int, left int, right int) {
	if left < right {
		mid := left + ((right - left) >> 1)
		var i int = left
		var j int = mid + 1

		sort_Merge_Part(arr, tmp, left, mid)
		sort_Merge_Part(arr, tmp, j, right)

		var cur int = 0

		for {
			if i <= mid && j <= right {
				if arr[i] > arr[j] {
					tmp[cur] = arr[j]
					cur += 1
					j += 1
				} else {
					tmp[cur] = arr[i]
					cur += 1
					i += 1
				}
			} else if i <= mid {
				tmp[cur] = arr[i]
				cur += 1
				i += 1
			} else if j <= right {
				tmp[cur] = arr[j]
				cur += 1
				j += 1
			} else {
				break
			}
		}

		//把记录到到数据，从新放入队列
		for k := 0; k < cur; k++ {
			arr[left+k] = tmp[k]
		}
	}
}
