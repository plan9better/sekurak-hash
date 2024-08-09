package main

import (
	"crypto/sha256"
	"fmt"
	"sync"
	"time"
)

const THREADS = 6

var wg sync.WaitGroup
var shutdown bool

func calculate(lines []string, stat *int, myhash string, pass *string) {
	defer wg.Done()

	for _, line := range lines {
		*stat++
		for _, line2 := range lines {
			if shutdown {
				return
			}

			hasher := sha256.New()
			concat := fmt.Sprintf("%s%s", line, line2)
			hasher.Write([]byte(concat))

			hash := fmt.Sprintf("%x", hasher.Sum(nil))

			if concat == "abderytachabderytach" {
				println("What the fuck", hash)
			}

			if hash == myhash {
				*pass = concat
				shutdown = true
				return
			}

		}
	}
}

// Pass the slice of lines to this function
// dictionary []string
// hash string
// -> string
func GetPassword(lines []string, myhash string, breaker *int) string {

	// Split the slice into however many threads you have to speed up execution
	// by claculating the hashes concurrently
	n_lines := len(lines)
	// n_lines := 100
	lines_per_thread := n_lines / THREADS

	thread_slices := make([][]string, THREADS)
	for i := range thread_slices {
		thread_slices[i] = make([]string, lines_per_thread)
	}

	// thread_slices[0] = lines[0:lines_per_thread]
	thread_slices[THREADS-1] = lines[(THREADS-1)*lines_per_thread : n_lines]
	for i := 1; i < THREADS-1; i++ {
		thread_slices[i] = lines[i*lines_per_thread : (i+1)*lines_per_thread]
	}

	// The above for 6 threads would look like this:
	// t1 := lines[0:lines_per_thread]
	// t2 := lines[lines_per_thread : lines_per_thread*2]
	// t3 := lines[2*lines_per_thread : lines_per_thread*3]
	// t4 := lines[3*lines_per_thread : lines_per_thread*4]
	// t5 := lines[4*lines_per_thread : lines_per_thread*5]
	// t6 := lines[5*lines_per_thread : n_lines]

	start := time.Now()
	shutdown = false
	status := make([]int, THREADS)
	password := ""

	for idx := range thread_slices {
		wg.Add(1)
		go calculate(thread_slices[idx], &status[idx], myhash, &password)
	}

	for !shutdown {
		num := 0
		for idx := range thread_slices {
			num += status[idx]
		}
		*breaker = num
		time.Sleep(time.Second)
	}

	wg.Wait()
	elapsed := time.Since(start)
	println("Time: ", elapsed.String())

	return password
}
