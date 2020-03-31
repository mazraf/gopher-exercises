package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/farzamalam/gopher-exercises/task/db"

	"github.com/spf13/cobra"
)

// doCmd is used to mark a task as complete.
// it deletes the entry from the bucket.
// it takes mulitple ids as slice from the cmd line.
// it calls the delete func with all those ids.
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do is used to mark task that are completed",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Invalid argument :", arg)
			} else {
				ids = append(ids, id)
			}
		}
		delete(ids)
	},
}

// delete is used to delete task/tasks from the bucket.
// it calls the db.AllTasks() and db.Deletes(task[id-1].Key)
// It prints the deleted task or error and does not return anything.
func delete(ids []int) {
	tasks, err := db.AllTasks()
	if err != nil {
		fmt.Println("Error while deleting the tasks : ", err)
		os.Exit(1)
	}
	for _, id := range ids {
		if id <= 0 || id > len(tasks) {
			fmt.Println("Invalid id : ", id)
			continue
		}
		task := tasks[id-1]
		err = db.DeleteTask(task.Key)
		if err != nil {
			fmt.Printf("Error while deleting the task with id : %d\n", id)
		} else {
			fmt.Printf("Marked %d as completed\n", id)
		}

	}
}

// init adds doCmd into RootCmd.
// it is called before main.
func init() {
	RootCmd.AddCommand(doCmd)
}
