package flow

func GetRemainder(x int, y int) (int, error) {
	if y != 0 {
		return x / y, nil
	}

	return 0, nil

}
