package main
import (
	"fmt"
	"strconv"
)

// struct
type Vector3 struct {
	x int
	y float32
	z int
}

// acts as a contructor
func (v *Vector3) Init(x int, y float32, z int) {
	v.x = x
	v.y = y
	v.z = z	
}

// main function
func main() {

	vec := new(Vector3)
	vec.Init(1,0,1)

	vy := fmt.Sprintf("%f", vec.y)

	fmt.Printf(strconv.Itoa(vec.x) + " X \n")
	fmt.Printf(vy + " Y \n")
	fmt.Printf(strconv.Itoa(vec.z) + " Z \n")
}
