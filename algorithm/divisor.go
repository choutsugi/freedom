package algorithm

func MaxDivisor[T int | int8 | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](min, max T) T {

	var maxDivisor T
	complement := max % min
	if complement != 0 {
		maxDivisor = MaxDivisor(complement, min)
	} else {
		maxDivisor = min
	}
	return maxDivisor
}

func MaxDivisor4Slice[T int | int8 | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](arr []T) T {

	maxDivisor := arr[0]
	for i := 1; i < len(arr); i++ {

		maxDivisor = MaxDivisor(maxDivisor, arr[i])
	}

	return maxDivisor
}
