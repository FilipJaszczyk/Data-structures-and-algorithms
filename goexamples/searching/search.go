package searching

func BinarySearch(list []int, number int) bool {
	head, tail := 0, len(list)-1
	for head <= tail {
		mid := (head + tail) / 2
		if list[mid] == number {
			return true
		} else if list[mid] > number {
			tail = mid - 1
		} else if list[mid] < number {
			head = mid + 1
		}
	}
	return false
}

func RecursiveBinarySearch(list []int, number int) bool {
	head, tail := 0, len(list)-1
	if head <= tail {
		mid := (head + tail) / 2
		if list[mid] == number {
			return true
		} else if list[mid] > number {
			return RecursiveBinarySearch(list[:mid], number)
		} else if list[mid] < number {
			return RecursiveBinarySearch(list[mid+1:], number)
		}
	}
	return false
}
