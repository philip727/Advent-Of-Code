package main


func getSideArea(a uint8, b uint8) uint8 {
	return a * b
}

func GetSurfaceAreaOfBox(dim [3]uint8) uint8 {
	a := getSideArea(dim[0], dim[1])
	b := getSideArea(dim[1], dim[2])
	c := getSideArea(dim[2], dim[0])
	smallest := min(a, b, c)

	area := 2*a + 2*b + 2*c + smallest

	return area
}

func main() {
}
