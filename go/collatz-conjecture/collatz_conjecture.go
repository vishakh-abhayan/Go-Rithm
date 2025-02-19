package collatzconjecture

func CollatzConjecture(n int) (int, error) {
	var value int
	for i:=0; value==1;i++{
		if n % 2 == 0{
			value = n/2
		}
		value = n * 3 + 1
	}
	return i, nil
}
