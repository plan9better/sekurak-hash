package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const THREADS = 6

 //aaaa
 // const myhash = "61be55a8e2f6b4e172338bddf184d6dbee29c98853e0a0485ecee7f27b9af0b4"

// 300th line
const myhash = "257d84791c8008f78b93f0170a4ae15222fb28fd9ae60fc9d8ec0acddfa356a0"

// aaabażur
//const myhash = "83a82bdd5761c7e6f963c53293770928913539cac3f204af25585e491f9ff1a2"

//const myhash = "cca34228ccaf5c4778f4d3c10f1a7af74f09897af5271b0844b18ba238328139"

// const myhash = "5fdda53345ee02ae5cde71a595cb31ecef263ec908f13a0017e15231cd282ca1"
var wg sync.WaitGroup
var shutdown bool

func calculate(lines []string, stat *int) {
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
			// hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

			//if concat == "aalborscyaalborscy" {
			//	println(concat, hash)
			//}

			if hash == myhash {
				println("Password: ", concat)
				shutdown = true
				return
			}

		}
	}
}

func main() {

	bytes, _ := os.ReadFile("slowa.txt")
	str := string(bytes)
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		lines[i] = strings.Trim(line, "\r")
	}

	// Split the slice into however many threads you have to speed up execution
	// by claculating the hashes concurrently
	n_lines := len(lines)
	// n_lines := 100
	lines_per_thread := n_lines / THREADS

	thread_slices := make([][]string, THREADS)
	for i := range thread_slices {
		thread_slices[i] = make([]string, lines_per_thread)
	}

	thread_slices[0] = lines[0:lines_per_thread]
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

	// for idx, slc := range thread_slices {
	// 	fmt.Printf("%d: %v\n", idx+1, slc[idx])
	// }


	start := time.Now()
	shutdown = false
  status := make([]int, THREADS);

	for idx := range thread_slices {
		wg.Add(1)
		go calculate(thread_slices[idx], &status[idx]);
	}

  for !shutdown{
    for idx := range thread_slices{
      fmt.Printf("------------------\n");
      fmt.Printf("%d: %d/%d\n", idx + 1, status[idx], lines_per_thread);
    }
    time.Sleep(time.Second)
  }

	wg.Wait()
	elapsed := time.Since(start)
	println("Time: ", elapsed.String())

}
