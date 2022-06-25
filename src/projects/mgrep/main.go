package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const Debug = true

// Job a file path to search for instances of the search term
type Job struct {
	Path string
}

// Result details of a line in the file where the search term was found
type Result struct {
	Line    string
	LineNum int
	Path    string
}

// produce recursively reads files in the path and sends a job to the
// jobs channel for each file in the path
func produce(jobs chan<- Job, path string) error {
	if Debug {
		fmt.Println("DEBUG", "produce", path)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("readdir: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			nextPath := filepath.Join(path, entry.Name())
			produce(jobs, nextPath)
		} else {
			jobs <- Job{Path: filepath.Join(path, entry.Name())}
		}
	}

	return nil
}

// consume takes a job and prints lines that contain the search term
// to standard output
func consume(job Job, searchTerm string) {
	if Debug {
		fmt.Println("DEBUG", "consume", job.Path)
	}

	results, err := findInFile(job.Path, searchTerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, r := range results {
		fmt.Printf("%v[%v]:%v\n", r.Path, r.LineNum, r.Line)
	}
}

// findInFile read a file indicated by the path and return lines
// that contain the search term
func findInFile(path string, searchTerm string) ([]Result, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("findInFile: %v", err)
	}

	results := make([]Result, 0)

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchTerm) {
			r := Result{scanner.Text(), lineNum, path}
			results = append(results, r)
		}
		lineNum += 1
	}

	return results, nil
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("usage: mgrep searchterm searchpath")
	}

	searchTerm, searchPath := os.Args[1], os.Args[2]
	var wg sync.WaitGroup
	jobs := make(chan Job, 100)
	numWorkers := 10

	go func() {
		defer func() {
			if Debug {
				fmt.Println("DEBUG", "close(jobs)")
			}
			close(jobs)
		}()

		err := produce(jobs, searchPath)
		if err != nil {
			log.Fatal(err)
		}
	}()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for job := range jobs {
				consume(job, searchTerm)
			}
		}()
	}

	wg.Wait()
}
