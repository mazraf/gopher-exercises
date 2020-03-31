package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/farzamalam/gopher-exercises/task/db"

	"github.com/spf13/cobra"
)

// addCmd is used to add a new task to the bucket.
// it takes the task as cmd argument the calls db.CreateTask() with task argument to store the path.
// then prints a new task added msg.
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add is used to add new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong.", err.Error())
			os.Exit(1)
		}
		fmt.Printf("New task \"%s\" has been added.\n", task)
	},
}

// init is used to add the addCmd to RootCmd.
// it runs before main()
func init() {
	RootCmd.AddCommand(addCmd)
}
