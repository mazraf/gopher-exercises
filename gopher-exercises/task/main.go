package main

import (
	"log"
	"path/filepath"

	"github.com/farzamalam/gopher-exercises/task/cmd"
	"github.com/farzamalam/gopher-exercises/task/db"
	"github.com/mitchellh/go-homedir"
)

// main is used to get the home dir path and initialize task.db their
// by calling db.Init()
// then it executes the RootCmd.
func main() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("Error while getting the home dir.")
	}
	dbPath := filepath.Join(home, "tasks.db")
	err = db.Init(dbPath)
	if err != nil {
		log.Fatal("Error while Init() db")
	}
	// Added a db.Close()
	defer db.GetDB().Close()
	cmd.RootCmd.Execute()
}
