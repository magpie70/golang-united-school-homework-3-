package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	var length = len(input)
	for i := 0; i < length; i++ {
		result = append(result, input[length-1-i])
	}
	return result
}
