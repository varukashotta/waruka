package rectangle

import "math"
import "fmt"

/*
 * init function added
 */
func init() {
	fmt.Println("rectangle package initialized")
}

/*
 * init function added
 */
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

/*
 * Diagonal function added
 */
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
