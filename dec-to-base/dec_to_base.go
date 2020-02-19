package dec_to_base

import (
	"fmt"
	"math"
)

// DecToBase returns a string representation
// of the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity
//
func DecToBase(dec int, base int) string {
	var res string
	powArr := getBase(base, dec)
	for i:= len(powArr)-1 ; i<=0; i-- {
		// item from the array
		no := powArr[i]
		fmt.Println(no)
		// dividing
		divRes := dec / no
		if divRes == 0 {
			continue
		}
		fmt.Println(divRes)
		res += fmt.Sprint(divRes)
		dec = dec % no
	}


	return res
}

func getBase(base int, num int) []int {
	var output []int
	var pow float64
	for i := 0; pow <= float64(num); i++ {
		pow = math.Pow(float64(base), float64(i))
		output = append(output, int(pow))
	}

	return output
}


