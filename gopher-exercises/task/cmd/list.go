package cmd

import (
	"fmt"
	"os"

	"github.com/farzamalam/gopher-exercises/task/db"
	"github.com/spf13/cobra"
)

//init is used to add the listCmd var to RootCmd.
//it runs before main()
func init() {
	RootCmd.AddCommand(listCmd)
}

//listCmd variable contains the task of displaying all the task that are present
//in the bucket. It calls db.AllTasks() that returns all the tasks, it just prints on
//the console.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list is used to print all the tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Error while getting the tasks : ", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("No task in the bucket to complete!")
			return
		}
		fmt.Printf("To do list :\n")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}
