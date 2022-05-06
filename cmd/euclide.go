package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var euclideCmd = &cobra.Command{
	Use:   "euclide",
	Short: "Thuật toán Euclide tìm GCD",
	Long:  "Ex: ./tools euclide 1970 1066",
	Run: func(cmd *cobra.Command, args []string) {
		a, _ := strconv.Atoi(args[0])
		b, _ := strconv.Atoi(args[1])
		fmt.Println("a = ", a)
		fmt.Println("b = ", b)
		fmt.Println("GCD = ", gcd(a, b))
	},
}

func init() {
	rootCmd.AddCommand(euclideCmd)
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	//fmt.Println(a, b)
	return gcd(b, a%b)
}
