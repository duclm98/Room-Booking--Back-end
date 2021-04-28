package methods

import (
	dto "echo-demo/DTOs"
)

func Contains(arr []dto.Booking, id uint) bool {
	for i := range arr {
		if arr[i].RoomID == id {
			return true
		}
	}
	return false
}