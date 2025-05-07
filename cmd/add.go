package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"timekeeper/models"
	"timekeeper/storage"
	"timekeeper/utils"

	"github.com/spf13/cobra"
)

var customer string
var task string
var timeStr string
var addDate string
var process string
var weekdays = []string{
	"monday",
	"tuesday",
	"wednesday",
	"thursday",
	"friday",
	"saturday",
	"sunday",
}

func contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

func lastWeekday(target time.Weekday, from time.Time) time.Time {
    daysAgo := (int(from.Weekday()) - int(target) + 7) % 7
    return from.AddDate(0, 0, -daysAgo)
}


// addCmd represents the add command
// This command allows the user to add a new time entry with customer, task, time spent, and date.
// It validates the customer against a list of allowed customers and formats the time input.
// The time can be specified in minutes or hours, and the date can be specified in YYYY-MM-DD format.
// If the date is not provided, it defaults to the current date.
// If the date is "yesterday", it will be set to the previous day.
// The command also provides tab completion for the customer and task flags.
// It uses the storage package to save the entry to a CSV file.
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new time entry",
	Run: func(cmd *cobra.Command, args []string) {
		customerMatch := false
		processMatch := false
		for _, c := range allowedCustomers {
			if c.Name == customer {
				customerMatch = true
				break
			}
		}
		for _, p := range allowedProcesses {
			if p.Process == process {
				processMatch = true
				break
			}
		}
		if !customerMatch {
			fmt.Printf("Customer '%s' not found in allowed list\n", customer)
			fmt.Printf("Please run 'timekeeper config -c \"%s\"' to add the customer\n", customer)
			os.Exit(1)
		}
		if !processMatch {
			fmt.Printf("Process '%s' not found in allowed list\n", process)
			fmt.Printf("Please run 'timekeeper config -c \"%s\" -p \"%s\"' to add the process\n", customer, process)
			os.Exit(1)
		}

		var minutesInt int = 0
		normalized := strings.ReplaceAll(timeStr, ",", ".") // handle "1,5h" -> "1.5h"

		if strings.HasSuffix(normalized, "m") {
			normalized = strings.TrimSuffix(normalized, "m")
			minutes, err := strconv.ParseFloat(normalized, 64)
			if err == nil {
				minutesInt = int(minutes)
			}
		} else if strings.HasSuffix(normalized, "h") {
			normalized = strings.TrimSuffix(normalized, "h")
			hours, err := strconv.ParseFloat(normalized, 64)
			if err == nil {
				minutesInt = int(hours * 60)
			}
		}
		minutesInt = utils.RoundToNearest15(minutesInt)
		if customer == "" || process == "" || task == "" || minutesInt <= 0 {
			fmt.Println("Usage: timekeeper add -c <customer> -p <process> -t <task> -m <minutes> [-d YYYY-MM-DD]")
			return
		}

		lowerCaseAddDate := strings.ToLower(addDate)
		var days = map[string]time.Weekday{
			"monday":    time.Monday,
			"tuesday":   time.Tuesday,
			"wednesday": time.Wednesday,
			"thursday":  time.Thursday,
			"friday":    time.Friday,
			"saturday":  time.Saturday,
			"sunday":    time.Sunday,
		}
		if addDate == "" {
			addDate = time.Now().Format("2006-01-02")
		} else if addDate == "yesterday" {
			addDate = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		} else if addDate == "today" {
			addDate = time.Now().Format("2006-01-02")
		} else if contains(weekdays, lowerCaseAddDate ) {
			today := time.Date(2025, 4, 14, 0, 0, 0, 0, time.UTC)
			lastDay := lastWeekday(days[lowerCaseAddDate], today)
			addDate = lastDay.Format("2006-01-02")

			fmt.Println("Today is:", today.Weekday(), today.Format("2006-01-02"))
			fmt.Println("Last Day:", addDate)
		}

		entry := models.Entry{
			Customer: customer,
			Process:  process,
			Task:     task,
			Minutes:  minutesInt,
			Date:     addDate,
		}

		err := storage.SaveEntryToCSV(entry)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("✔ Entry saved successfully.")
		}
	},
}

// allowedCustomers is a slice of Customer structs that holds the list of allowed customers.
var allowedCustomers []models.Customer
var allowedProcesses []models.Process

// init function initializes the add command and its flags.
func init() {
	var err error
	allowedCustomers, err = storage.CustomerFromCSV()
	if err != nil {
		fmt.Println("Could not load customers from CSV:", err)
	}

	allowedProcesses, err = storage.ProcessFromCSV()
	if err != nil {
		fmt.Println("Could not load customers from CSV:", err)
	}


    rootCmd.AddCommand(addCmd)

    addCmd.Flags().StringVarP(&customer, "customer", "c", "", "Customer name")
	addCmd.Flags().StringVarP(&process, "process", "p", "", "Process name")
    addCmd.Flags().StringVarP(&task, "task", "t", "", "Task description")
    addCmd.Flags().StringVarP(&timeStr, "time", "m", "0", "Time spent in minutes")
	addCmd.Flags().StringVarP(&addDate, "date", "d", "", "Date of entry (format: YYYY-MM-DD)")

	addCmd.RegisterFlagCompletionFunc("customer", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		customers, err := storage.CustomerFromCSV()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		var suggestions []string
		for _, c := range customers {
			if strings.HasPrefix(strings.ToLower(c.Name), strings.ToLower(toComplete)) {
				suggestions = append(suggestions, c.Name)
			}
		}
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	})
	addCmd.RegisterFlagCompletionFunc("process", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		processes, err := storage.ProcessFromCSV()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		var suggestions []string
		for _, p := range processes {
			if strings.HasPrefix(strings.ToLower(p.Process), strings.ToLower(toComplete)) {
				suggestions = append(suggestions, p.Process)
			}
		}
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	})
	addCmd.RegisterFlagCompletionFunc("task", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		tasks, err := storage.ReadEntriesFromCSV()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		var suggestions []string
		for _, t := range tasks {
			if strings.HasPrefix(strings.ToLower(t.Task), strings.ToLower(toComplete)) {
				suggestions = append(suggestions, t.Task)
			}
		}
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	})
	addCmd.RegisterFlagCompletionFunc("date", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		dates := []string{
			time.Now().Format("2006-01-02"),
			"today",
			"yesterday",
			"monday",
			"tuesday",
			"wednesday",
			"thursday",
			"friday",
			"saturday",
			"sunday",
		}

		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		var suggestions []string
		for _, d := range dates {
			if strings.HasPrefix(strings.ToLower(d), strings.ToLower(toComplete)) {
				suggestions = append(suggestions, d)
			}
		}
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	})
}
