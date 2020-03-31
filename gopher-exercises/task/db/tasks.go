package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

// taskBucket is the name of bucket of type []byte
var taskBucket = []byte("tasks")

// db is the *bolt.DB contains the instance of db.
var db *bolt.DB

// Init is used to create
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// GetDB is used to return the instance of *bolt.DB so that connection
// can be closed from the main.
func GetDB() *bolt.DB {
	return db
}

// CreateTask is used add a new task in the bucket.
// it uses bucket.NextSequence() to generate ids.
// it calls itob() to convert id of type int to []byte.
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		return b.Put(itob(id), []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// Task represents the task. It has Key as int and Value which is our task is of
// type string.
type Task struct {
	Key   int
	Value string
}

// AllTasks is used to return the slice of all the tasks that are present in the bucket.
// it uses cursor to iterate through all the key-value pair.
func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// DeleteTask is used to delete a task from the bucket.
func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

// itob is takes and int type and returns a []byte type using encoding/binary
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi is used to convert byte[] to int type using encoding/binary
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
