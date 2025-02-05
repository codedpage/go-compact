package main

import (
	"bufio"
	"fmt"
	"io"
	"sync"
)

// global variables for managing workers and tasks
var (
	wg    sync.WaitGroup
	tasks = make(chan Task)
)

// Task struct with a completion flag
type Task struct {
	Completed bool
}

// workerPool manages worker goroutines to process tasks
func workerPool(w int) {
	defer wg.Done()

	// Slice to keep track of the number of tasks each worker processed
	taskCounts := make([]int, w)

	var taskWg sync.WaitGroup

	// Launch workers
	for i := 0; i < w; i++ {
		taskWg.Add(1)
		go func(workerID int) {
			defer taskWg.Done()
			for task := range tasks {
				if processTask(&task) {
					taskCounts[workerID]++
				}
			}
		}(i)
	}

	taskWg.Wait()
	printSummary(taskCounts)
}

// printSummary displays the count of tasks processed by each worker
func printSummary(taskCounts []int) {
	for i, count := range taskCounts {
		fmt.Printf("Worker %d processed %d tasks\n", i+1, count)
	}
}

// publishTasks sends the specified number of tasks to the tasks channel
func publishTasks(t int) {
	defer wg.Done()
	for i := 0; i < t; i++ {
		tasks <- Task{}
	}
	close(tasks)
}

// processTask marks a task as completed and returns true if successful
func processTask(task *Task) bool {
	task.Completed = true
	return task.Completed
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	nTasks := 1000 //, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	t := int(nTasks)
	nWorkers := 5 //, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	w := int(nWorkers)

	wg.Add(1)
	go publishTasks(t)
	wg.Add(1)
	go workerPool(w)
	wg.Wait()
}

// Helper function to read a line from standard input
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return string(str)
}

// Helper function to check for errors
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
