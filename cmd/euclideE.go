package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var euclideECmd = &cobra.Command{
	Use:   "euclideE",
	Short: "Thuật toán Euclide mở rộng tìm GCD",
	Long:  "Ex: ./tools euclideE 1759 550",
	Run: func(cmd *cobra.Command, args []string) {
		a, _ := strconv.Atoi(args[0])
		b, _ := strconv.Atoi(args[1])
		fmt.Println("a = ", a)
		fmt.Println("b = ", b)
		d, x, y := gcdE(a, b)
		fmt.Printf("d = %d\nx = %d\ny = %d", d, x, y)
	},
}

func init() {
	rootCmd.AddCommand(euclideECmd)
}

func gcdE(a int, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	x2 := 1
	x1 := 0
	y2 := 0
	y1 := 1
	for b > 0 {
		q := int(a / b)
		r := a - q*b
		x := x2 - q*x1
		y := y2 - q*y1
		a = b
		b = r
		x2 = x1
		x1 = x
		y2 = y1
		y1 = y
		fmt.Print(q, r, x, y, a, b, x2, x1, y2, y1)
	}
	d := a
	x := x2
	y := y2
	return d, x, y
}
