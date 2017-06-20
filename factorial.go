package factorial


func countingRecursive(x uint) uint {
	if x == 1 { return 1 }
	return x * countingRecursive(x - 1)
}

func counting(x uint) uint {
	result := uint(1)
	for i := uint(2); i <= x; i++ {
		result *= i
	}
	return result
}