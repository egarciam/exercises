package leetcode

func isPossibleToSplit(nums []int) bool {
	if len(nums)%2 != 0 {
		return false
	}

	keys := make(map[int]int)

	for _, v := range nums {
		num, ok := keys[v]
		if ok { //Existe
			if num >= 2 { //Solo puede haber uno maximo
				return false
			}
			//ya esta en uno lo incluimos en el segundo
			keys[v] = keys[v] + 1
		} else {
			keys[v] = 1
		}
	}

	return true
}
