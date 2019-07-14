package euler

import (
	"fmt"
	"strconv"

  log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var (
	oneCmd = &cobra.Command{
		Use: "one",
		Aliases: []string{"1"},
		Short: "Multiples of 3 and 5",
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
	oneCmd.Flags().StringSliceP("nums", "n", []string{"3", "5"}, "Numbers to be multiple of")
}

func runOneCmd(cmd *cobra.Command, args []string) {
	below := viper.GetInt("below")
	nums_str := viper.GetStringSlice("nums")
	nums := make([]int, len(nums_str))
	for i, num_str := range nums_str {
		num, err := strconv.Atoi(num_str)
		if err != nil || num < 0 {
			log.Fatalf("'%s' is not a positive integer", num_str)
		}
		nums[i] = num
	}
	
	fmt.Println(solve(nums, below))
}

func solve(nums []int, below int) int {
	sum := 0
	for i := 1; i < below; i++ {
		divisible := false
		for _, num := range nums {
			if i % num == 0 {
				divisible = true
				break
			}
		}
		if divisible {
			sum += i
		}
	}
	return sum
}

