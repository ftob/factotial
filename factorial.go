package factorial


func countingRecursive(x uint) uint {
	if x == 1 { return 1 }
	return x * countingRecursive(x - 1)
}