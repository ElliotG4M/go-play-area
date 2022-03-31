package loops

// Pretty standard for loop stuff
func conditionalLoop() {
	var i int
	for i < 5 {
		println(i)
		i++

		if i == 3 {
			continue
		}

		println("continuing...")

		if i == 4 {
			break
		}
	}
}

func usualForLoop() {
	for i := 0; i < 5; i++ {
		println(i)
	}
}

func infiniLoop() {
	i := 0
	for {
		if i > 5 {
			break
		}
		println(i)
		i++
	}
}

// Different syntax than typical foreach. Uses range keyword. Works on arrays, slices, maps, strings and channels
func forEach() {
	slice := []int{1, 2, 3}
	for index, value := range slice {
		println(index, "=>", value)
	}

	// We just want the indexes
	for index := range slice {
		println(index)
	}

	// We just want the values
	for _, value := range slice {
		println(value)
	}
}
