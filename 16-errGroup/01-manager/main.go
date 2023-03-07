package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg := new(errgroup.Group)
	s := []string{"task1", "fake1", "task2", "fake2"}
	for _, v := range s {
		task := v
		eg.Go(func() error {
			return doTask(task)
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Printf("error :%v\n", err)
	}
	fmt.Println("finish")

}
func doTask(task string) error {
	if task == "fake1" || task == "fake2" {
		return fmt.Errorf("%v failed", task)
	}
	fmt.Printf("task %v completed\n", task)
	return nil
}
