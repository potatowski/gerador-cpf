package src

func calculator(doc string, first int) int {
	var sum int
	if len(doc) > 9 {
		for _, value := range doc {
			if first == 1 {
				first = 9
			}

			var rs = int(value-'0') * first
			sum += rs
			first--
		}

		return sum
	}

	for _, value := range doc {
		var rs = int(value-'0') * first
		sum += rs
		first--
	}

	return sum
}

func getDigit(sum int) int {
	rest := sum % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}
