package methods

import "strconv"

func TimeArray() []string {
	var result []string
	for i := 0; i < 24; i++ {
		for j := 0; j <= 30; j += 30 {
			var h string
			var m string
			if i < 10 {
				h = "0" + strconv.Itoa(i)
			} else {
				h = strconv.Itoa(i)
			}

			if j == 0 {
				m = "00"
			} else {
				m = "30"
			}

			time := h + ":" + m
			result = append(result, time)
		}
	}
	return result
}