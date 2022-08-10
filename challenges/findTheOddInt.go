package challenges

/*
*Given an array of integers, find the one that appears an odd number of times.
*There will always be only one integer that appears an odd number of times.

Examples
[7] should return 7, because it occurs 1 time (which is odd).
[0] should return 0, because it occurs 1 time (which is odd).
[1,1,2] should return 2, because it occurs 1 time (which is odd).
[0,1,0,1,0] should return 0, because it occurs 3 times (which is odd).
[1,2,2,3,3,3,4,3,3,3,2,2,1] should return 4, because it appears 1 time (which is odd).
*/

func FindOdd(seq []int) int {
	val := 0
	m := make(map[int]int)
	for _, n := range seq {
		m[n]++
	}
	for k, i := range m {
		if i%2 != 0 {
			val = k
		}
	}
	return val
}

func FindOddBestSolution(seq []int) int {
	res := 0
	for _, x := range seq {
		res ^= x
	}
	return res
}
