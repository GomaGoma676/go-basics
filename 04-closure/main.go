package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func funcDefer() {
	defer fmt.Println("main func final-finish")
	defer fmt.Println("main func semi-finish")
	fmt.Println("hello world")
}
func trimExtension(files ...string) []string {
	out := make([]string, 0, len(files))
	for _, f := range files {
		out = append(out, strings.TrimSuffix(f, ".csv"))
	}
	return out
}
func fileChecker(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", errors.New("file not found")
	}
	defer f.Close()
	return name, nil
}
func addExt(f func(file string) string, name string) {
	fmt.Println(f(name))
}
func multiply() func(int) int {
	return func(n int) int {
		return n * 1000
	}
}
func countUp() func(int) int {
	count := 0
	return func(n int) int {
		count += n
		return count
	}
}

//var count int

func main() {
	funcDefer()
	files := []string{"file1.csv", "file2.csv", "file3.csv"}
	fmt.Println(trimExtension(files...))
	name, err := fileChecker("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

	i := 1
	func(i int) {
		fmt.Println(i)
	}(i)
	f1 := func(i int) int {
		return i + 1
	}
	fmt.Println(f1(i))

	f2 := func(file string) string {
		return file + ".csv"
	}
	addExt(f2, "file1")

	f3 := multiply()
	fmt.Println(f3(2))

	f4 := countUp()
	for i := 1; i <= 5; i++ {
		v := f4(2)
		fmt.Printf("%v\n", v)
	}
}
