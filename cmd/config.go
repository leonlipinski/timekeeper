package cmd

import (
	"fmt"
	"timekeeper/models"
	"timekeeper/storage"
	"timekeeper/utils"

	"github.com/spf13/cobra"
)

var customerName string
var processName string
var moveFiles bool
var version bool
var versionString = "0.1.0"

// configCmd represents the config command
// This command is used to configure the timekeeper application
// It allows the user to add a customer to the allowed list
// and to move the entries.csv file to a new location
// The command also renames the entries.csv file to entries-YYYY-MM-DD-to-YYYY-MM-DD.csv
// based on the dates found in the file
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the timekeeper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Configure the timekeeper")
		if version {
			fmt.Println("Version:", versionString)
			return
		}
		if customerName != "" {
			fmt.Println("Adding customer:", customerName)
			customer := models.Customer{Name: customerName}
			storage.SaveCustomerToCSV(customer)
		}
		if (processName != "" && customerName != "") {
			fmt.Println("Adding process:", processName)
			process := models.Process{
				Customer: customerName,
				Process:  processName,
			}
			storage.SaveProcessToCSV(process)
		} else {
			fmt.Println("No customer and process name provided")
			return
		}
		if moveFiles {

			oldFilePath, _ := storage.GetConfigPath("entries.csv")
			FileExists := storage.FileExists(oldFilePath)
			if !FileExists {
				fmt.Println("File does not exist:", oldFilePath)
				return
			}
			fmt.Println("Moving files to new location")
			dates, _ := storage.ReadDatesFromEntriesCSV()
			minDate, maxDate, _ := utils.FindMinMaxDates(dates)
			if minDate == "" || maxDate == "" {
				fmt.Println("No valid dates found in entries.csv")
				return
			}
			fmt.Println("Min Date:", minDate)
			fmt.Println("Max Date:", maxDate)
			fileName := fmt.Sprintf("entries-%s-to-%s.csv", minDate, maxDate)


			newFilePath, _ := storage.GetConfigPath(fileName)

			fmt.Printf("Moving %s to %s\n", oldFilePath, newFilePath)

			err := storage.RenameFile(
				oldFilePath,
				newFilePath,
			)
			if err != nil {
				fmt.Println("Error moving files:", err)
				return
			}
			fmt.Println("Files moved successfully")
		}
	},
}

// init function initializes the config command and its flags
func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVarP(&customerName, "customer", "c", "", "Add a customer to the allowed list")
	configCmd.Flags().StringVarP(&processName, "process", "p", "", "Add a process to the allowed list")
	configCmd.Flags().BoolVarP(&moveFiles, "archive", "a", false, "Renames entries.csv to entries-YYYY-MM-DD-to-YYYY-MM-DD.csv")
	configCmd.Flags().BoolVarP(&version, "version", "v", false, "Show version information")

}
