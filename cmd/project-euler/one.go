package euler

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	oneCmd = &cobra.Command{
		Use:     "one",
		Aliases: []string{"1"},
		Short:   "Multiples of 3 and 5",
		Long: `
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
`,
		Run: runOneCmd,
	}
)

func init() {
}

func setOneCmdFlags() {
	oneCmd.Flags().IntP("below", "b", 1000, "Number to count up to but not include")
	oneCmd.Flags().IntP("num1", "1", 3, "First num")
	oneCmd.Flags().IntP("num2", "2", 5, "Second num")
}

func runOneCmd(cmd *cobra.Command, args []string) {
	below := viper.GetInt("below")
	num1 := viper.GetInt("num1")
	num2 := viper.GetInt("num2")

	fmt.Println(solve1(num1, num2, below))
}

func solve1(num1, num2, below int) int {
	log.WithFields(log.Fields{
		"num1": num1,
		"num2": num2,
		"below": below,
	}).Debug("Starting solution...")

	sum1 := sumSequence(num1, below)
	sum2 := sumSequence(num2, below)
	num3 := num1 * num2
	sum3 := sumSequence(num3, below)
	
	return (sum1 + sum2 - sum3)
}

func sumSequence(increment, below int) int {
	n := float64((below-1) / increment)
	sum := (float64(increment) + n*float64(increment)) * n/2
	log.Debugf("%d: n: %f %f, Sum: %f", increment, n, n/2, sum)
	return int(sum)
}
