package sum

func Ints(sliceInt ...int) int {
	return ints(sliceInt)
}

func ints(sliceInt []int) int {
	if len(sliceInt) == 0 {
		return 0
	}else{
		return ints(sliceInt[1:]) + sliceInt[0]
	}
}