package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	solvers "main/solvers"
	u "main/utils"
	s "main/utils/structs"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	solver := os.Args[1]
	id := os.Args[2]

	length, _ := strconv.Atoi(os.Args[3])
	vals := os.Args[5:]
	limit, _ := strconv.Atoi(os.Args[4])

	items := make([]s.Item, length)

	for i := 0; i < length*2; i += 2 {
		var index int = i / 2

		weight, _ := strconv.Atoi(vals[i])
		value, _ := strconv.Atoi(vals[i+1])

		items[index] = s.Item{Value: value, Weight: weight}
	}

	var solution *s.Configuration
	var start time.Time
	var elapsed time.Duration

	//---------------------------------------------

	switch solver {
	case "brute":

		start = time.Now()

		solution = u.StartRoutines(solvers.GetBruteSolveJob, s.NewQueue[func()], length, limit, &items, runtime.NumCPU())
		elapsed = time.Since(start)
		break

	case "bb":

		start = time.Now()

		solution = u.StartRoutines(solvers.GetBBSolveJob, s.NewQueue[func()], length, limit, &items, runtime.NumCPU())
		elapsed = time.Since(start)
		break

	default:
		fmt.Println("Unknown solver")
		break
	}

	if solution != nil {
		fmt.Print(id + " ")
		fmt.Print(solution.ToString() + " ")
		fmt.Println(elapsed.Seconds())
	} else {
		fmt.Println("FAILED")
	}
}
