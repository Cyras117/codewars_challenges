package challenges

import (
	"fmt"
	"math"
)

/*
*The rgb function is incomplete. Complete it so that passing in RGB decimal values will result in a hexadecimal representation being returned.
*Valid decimal values for RGB are 0 - 255. Any values that fall out of that range must be rounded to the closest valid value.
*Note: Your answer should always be 6 characters long, the shorthand with 3 will not work here.
*The following are examples of expected output values:

kata.rgb(255, 255, 255) -- returns FFFFFF
kata.rgb(255, 255, 300) -- returns FFFFFF
kata.rgb(0, 0, 0) -- returns 000000
kata.rgb(148, 0, 211) -- returns 9400D3

*/

func RGB(r, g, b int) string {
	const maxlim = 255
	const minlim = 0
	const base float64 = 16
	nres := make([][2]float64, 3)
	m := map[float64]string{
		10: "A", 11: "B", 12: "C", 13: "D", 14: "E", 15: "F",
	}
	var result string

	if r > maxlim {
		r = maxlim
	}
	if r < minlim {
		r = minlim
	}
	if g > maxlim {
		g = maxlim
	}
	if g < minlim {
		g = minlim
	}
	if b > maxlim {
		b = maxlim
	}
	if b < minlim {
		b = minlim
	}

	nres[0][0], nres[0][1] = math.Modf(float64(r) / base)
	nres[1][0], nres[1][1] = math.Modf(float64(g) / base)
	nres[2][0], nres[2][1] = math.Modf(float64(b) / base)

	for _, res := range nres {
		if res[0] < 10 {
			result += fmt.Sprint(res[0])
		}
		if res[0] > 9 {
			result += m[res[0]]
		}
		if res[1]*base < 10 {
			result += fmt.Sprint(res[1] * base)
		}
		if res[1]*base > 9 {
			result += m[res[1]*base]
		}
	}

	return result
}

////////Best solution///////////////
func getValid(x int) int {
	switch {
	case x < 0:
		return 0
	case x > 255:
		return 255
	default:
		return x
	}
}

func RGBbestSolution(r, g, b int) string {
	res := fmt.Sprintf("%02X%02X%02X", getValid(r), getValid(g), getValid(b))
	fmt.Println(res)
	return res
}
