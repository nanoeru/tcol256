// color project main.go
package tcol256

import (
	"fmt"
)

func Example() {
	//Print a test graphic showing all colors.

	fmt.Println("System colors:")
	for i := 0; i < 8; i++ {
		print_color("  ", NONE, i)
	}
	fmt.Println()
	for i := 8; i < 16; i++ {
		print_color("  ", NONE, i)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("RGB color cube, 6x6x6:")
	for g := 0; g < 6; g++ {
		for r := 0; r < 6; r++ {
			for b := 0; b < 6; b++ {
				print_color("  ", NONE, rgb(r, g, b))
			}
			fmt.Print("  ")
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Grayscale ramp, with RGB grays:")
	for v := 0; v < 24; v++ {
		print_color("  ", NONE, gray(v))
	}
	fmt.Println()

	print_color("  ", NONE, rgb(0, 0, 0))
	fmt.Print("              ")
	print_color("    ", NONE, rgb(1, 1, 1))
	fmt.Print("    ")
	print_color("    ", NONE, rgb(2, 2, 2))
	fmt.Print("    ")
	print_color("    ", NONE, rgb(3, 3, 3))
	fmt.Print("    ")
	print_color("    ", NONE, rgb(4, 4, 4))
	fmt.Print("    ")
	print_color("    ", NONE, rgb(4, 4, 4))
	fmt.Print("    ")
	print_color("    ", NONE, rgb(5, 5, 5))
	fmt.Println()
}
