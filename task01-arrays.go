package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var length = len(input)
	var sum float32
	for i := 0; i < length; i++ {
		sum += input[i]
	}
	return sum / float32(length)
}
