package euler

import (
	"fmt"
	"math"
	"math/big"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	twentyFiveCmd = &cobra.Command{
		Use:     "twentyfive",
		Aliases: []string{"25"},
		Short:   "First Fibonacci number with n digits",
		Long: `
The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
`,
		Run: runTwentyFiveCmd,
	}
)

func init() {
}

func setTwentyFiveFlags() {
	twentyFiveCmd.Flags().IntP("digits", "d", 1000, "Number of digits")
}

func runTwentyFiveCmd(cmd *cobra.Command, args []string) {
	digits := viper.GetInt("digits")

	// fmt.Println(solve25(digits))
	fmt.Println(Fibonacci(new(big.Int).SetInt(7)))
}

// func solve25(digits int) int {
// 	log.WithFields(log.Fields{
// 		"digits": digits,
// 	}).Debug("Starting solution...")

// 	n := digits
// 	last_n := 0
// 	curDigits := 0
// 	inc := 0
// 	for {
// 		curDigits = CountDigits(Fibonacci(n))
// 		if curDigits == digits  && CountDigits(Fibonacci(n-1)) == digits-1 {
// 			break
// 		}

// 		if curDigits < digits {
// 			if last_n < n {
// 				inc = n/2
// 			} else {
// 				inc = (last_n - n) / 2 
// 			}
		
// 		} else {
// 			if last_n > n {
// 				inc = -n/2
// 			} else {
// 				inc = (last_n - n) / 2
// 			}
// 		}

// 		last_n = n
// 		n += inc
// 		if n < 0 {
// 			log.Fatal("n is less than 0")
// 		}
// 		fmt.Println(n, Fibonacci(n))
// 	}

// 	return n
// }

func Fibonacci(n *big.Int) *big.Int {
	nn := new(big.Int).SetInt(n)
	phi := big.NewFloat(math.Phi)
	n1 := new(big.Float).Exp(big.NewInt(), big.NewInt(), nil)
	fibn := new(big.Int).SetInt(

	)
	// fibn := int((math.Pow(math.Phi, new(big.Float).SetInt(n)) - math.Pow(-math.Phi, new(big.Float).SetInt(-n))) / (2*math.Phi-1))
	log.Debug(fmt.Sprintf("Fib(%d) = %d", n, fibn))
	return fibn
}

 func CountDigits(i int) (count int) {
 	for i != 0 {
 		i /= 10
 		count = count + 1
 	}
 	return count
 }