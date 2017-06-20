package factorial


func countingRecursive(x int) int {
	if x == 1 { return 1 }
	return x * countingRecursive(x - 1)
}