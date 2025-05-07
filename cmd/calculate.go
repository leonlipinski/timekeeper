package cmd

import (
	"fmt"
	// "time"
	"strings"
	"strconv"
	"timekeeper/utils"
	"github.com/spf13/cobra"
)

var startTime string
var endTime string

// calculateCmd represents the calculate command
// This command calculates the total time spent between two timestamps
// It takes two flags: --start and --end, which represent the start and end times respectively
// The command then calculates the difference between the two times and formats it as a duration
// The command also rounds the total minutes to the nearest 15 minutes
// The command prints the total minutes spent and the formatted duration
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculate the total time spentbetween to timestamp",
	Run: func(cmd *cobra.Command, args []string) {
		startTimeParts := strings.Split(startTime, ":")
		endTimeParts := strings.Split(endTime, ":")
		endHour, _ := strconv.Atoi(endTimeParts[0])
		startHour, _ := strconv.Atoi(startTimeParts[0])
		calcHours := endHour-startHour

		endMinute, _ := strconv.Atoi(endTimeParts[1])
		startMinute, _ := strconv.Atoi(startTimeParts[1])
		calcMinutes := endMinute-startMinute

		totalMinutes := utils.FormatAsMinutes(calcHours, calcMinutes)
		calcTotalMinutes := utils.RoundToNearest15(totalMinutes)

		fmt.Printf("Without rounding: %d\n", totalMinutes)
		fmt.Printf("Total minutes spent: %d\n", calcTotalMinutes)
		fmt.Printf("Total time spent: %s\n", utils.FormatDurationAsTime(calcTotalMinutes))
	},
}

// init function initializes the calculate command and its flags
func init() {

	rootCmd.AddCommand(calculateCmd)

	calculateCmd.Flags().StringVarP(&startTime, "start", "s", "", "Start date in YYYY-MM-DD format")
	calculateCmd.Flags().StringVarP(&endTime, "end", "e", "", "End date in YYYY-MM-DD format")
	calculateCmd.MarkFlagRequired("start")
	calculateCmd.MarkFlagRequired("end")
}
