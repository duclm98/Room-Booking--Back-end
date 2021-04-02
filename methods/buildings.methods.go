package methods

import (
	form "echo-demo/forms"
)

func Contains(arr []form.Booking, id uint) bool {
	for i := range arr {
		if arr[i].RoomID == id {
			return true
		}
	}
	return false
}