package cmd

import (
	// "container/list"
	"fmt"
	"timekeeper/storage"
	"timekeeper/utils"

	"github.com/spf13/cobra"
)

var listDate string
var byDate bool = false

// listCmd represents the list command
// This command lists all time entries from the entries.csv file
// It takes an optional --date flag to filter entries by date
// It also takes an optional --by-date flag to group entries by date
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all time entries",
	Run: func(cmd *cobra.Command, args []string) {
		entries, err := storage.ReadEntriesFromCSV()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if len(entries) == 0 {
			fmt.Println("No entries found.")
			return
		}

		// Use a map to group by Date + Customer + Task
		type key struct {
			Date     string
			Customer string
			Process  string
			Task     string
		}
		// Use a map to group by Date
		type date struct {
			Date     string
		}

		combined := make(map[key]int)
		minutesByDate := make(map[date]int)

		for _, entry := range entries {
			if listDate != "" && entry.Date != listDate {
				continue // Skip entries that don't match the date
			}
			k := key{Date: entry.Date, Customer: entry.Customer, Process: entry.Process, Task: entry.Task}
			combined[k] += entry.Minutes
			e := date{Date: entry.Date}
			minutesByDate[e] += entry.Minutes
		}

		// Print the combined results
		if byDate == false {
			fmt.Println("Combined time entries:")
			for k, totalMinutes := range combined {
				fmt.Printf("Date: %s, Customer: %s, Process: %s, Task: %s, Total Minutes: %s\n",
					k.Date, k.Customer, k.Process, k.Task, utils.FormatDurationAsTime(totalMinutes))
			}
		} else if byDate == true {
			fmt.Println("Combined time entries by date:")
			for e, totalMinutes := range minutesByDate {
				fmt.Printf("Date: %s, Total minutes worked: %s\n",
					e.Date, utils.FormatDurationAsTime(totalMinutes))
			}
		}
	},
}

// init function initializes the list command and its flags
func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&listDate, "date", "d", "", "Date to search for (YYYY-MM-DD)")
	listCmd.Flags().BoolVarP(&byDate, "by-date", "b", false, "Group by date")
}
