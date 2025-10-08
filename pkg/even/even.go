package even

func IsEven(number int) (bool, error) {
	// TODO: Implement error handling for negative numbers
	if number%2 == 0 {
		return true, nil
	}
	return false, nil
}
