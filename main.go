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

var values []int
var weights []int

var limit int

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	solver := os.Args[1]
	id := os.Args[2]

	length, _ := strconv.Atoi(os.Args[3])
	vals := os.Args[5:]
	limit, _ = strconv.Atoi(os.Args[4])

	values = make([]int, length)
	weights = make([]int, length)

	for index, val := range vals {
		if index%2 == 0 {
			weights[int(index/2)], _ = strconv.Atoi(val)
		} else {
			values[int(index/2)], _ = strconv.Atoi(val)
		}
	}

	var solution *s.Configuration
	var start time.Time
	var elapsed time.Duration

	//---------------------------------------------

	switch solver {
	case "brute":
		solution = s.NewConf(length, &values, &weights)

		start = time.Now()

		u.StartRoutines(solvers.GetBruteSolveJob, length, limit, solution, runtime.NumCPU())
		elapsed = time.Since(start)
		break

	case "bb":
		solution = s.NewConf(length, &values, &weights)

		start = time.Now()

		u.StartRoutines(solvers.GetBBSolveJob, length, limit, solution, runtime.NumCPU())
		elapsed = time.Since(start)
		break

	default:
		fmt.Println("Unknown solver")
		break
	}

	if solution != nil {
		fmt.Print(id + " ")
		fmt.Print(solution.ToString() + " ")
		fmt.Println(elapsed)
	} else {
		fmt.Println("FAILED")
	}
}
