package sugar

func Contains[K comparable](elements []K, value K) bool {
	for _, ele := range elements {
		if ele == value {
			return true
		}
	}
	return false
}

func Delete[K comparable](elements []K, value K) []K {
	var res []K
	for _, ele := range elements {
		if value != ele {
			res = append(res, ele)
		}
	}
	return res
}

func Diff[K comparable](a, b []K) []K {
	var diff []K
	temp := map[K]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}
	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diff = append(diff, val)
		}
	}

	return diff
}
