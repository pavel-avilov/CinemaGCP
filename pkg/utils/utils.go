package utils

func GenerateAvailableSeats(countSeats int) []int {
	var availableSeats []int
	if countSeats < 1 {
		return availableSeats
	}
	p := 1
	for {
		availableSeats = append(availableSeats, p)
		p += 1
		if p > countSeats {
			return availableSeats
		}
	}
}

func Difference(a, b []int) (diff []int) {
	m := make(map[int]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}
