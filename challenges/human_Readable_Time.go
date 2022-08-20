package challenges

import (
	"fmt"
	"math"
)

/*
*Write a function, which takes a non-negative integer (seconds) as input and returns the time in a human-readable format (HH:MM:SS)
*HH = hours, padded to 2 digits, range: 00 - 99
*MM = minutes, padded to 2 digits, range: 00 - 59
*SS = seconds, padded to 2 digits, range: 00 - 59
*The maximum time never exceeds 359999 (99:59:59)
*You can find some examples in the test fixtures.
 */

func HumanReadableTime(seconds int) string {
	number := float64(seconds)

	//hour
	h, haux := math.Modf(number / 3600)

	//min
	min, minaux := math.Modf(haux * 60)

	//sec
	sec, secaux := math.Modf(minaux * 60)
	sec = math.Round(sec + secaux)

	return fmt.Sprintf("%.2d:%.2d:%.2d", int(h), int(min), int(sec))
}
