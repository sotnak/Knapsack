package utils

import (
	t "main/solvers/types"
	s "main/utils/structs"
	"sync"
)

const HeightLimit = 5

func DoubleCheck(condition func() bool, than func(), els func(), lock *sync.RWMutex, readOnly bool) {
	if condition() {
		if readOnly {
			lock.RLock()
		} else {
			lock.Lock()
		}

		if condition() {
			than()
		} else {
			els()
		}

		if readOnly {
			lock.RUnlock()
		} else {
			lock.Unlock()
		}

	} else {
		els()
	}
}

func StartRoutines(getJob t.GetJob, newContainer func() s.Container[func()],
	length int, limit int, items *[]s.Item, numOfRoutines int) *s.Configuration {

	rwLock := &sync.RWMutex{}

	var myItems []s.Item = make([]s.Item, length)

	copy(myItems, *items)

	initConf := s.NewConf(length, &myItems)

	solution := s.NewConf(length, &myItems)

	waiting := 0

	cond := sync.NewCond(&sync.Mutex{})

	jobs := newContainer()

	solve := getJob(limit, solution, jobs, rwLock, cond)
	jobs.Push(func() { solve(initConf, 0) })

	var wg sync.WaitGroup

	for i := 0; i < numOfRoutines; i++ {
		wg.Add(1)

		func() {
			id := i
			go func() {
				defer wg.Done()
				routine(id, jobs, &waiting, numOfRoutines, cond)
			}()
		}()

	}

	wg.Wait()

	if !jobs.Empty() {
		panic("Routines terminated")
	}

	return solution
}

func routine(id int, jobs s.Container[func()], waiting *int, numOfRoutines int, cond *sync.Cond) {

	var job func()
	job = nil

	for {

		cond.L.Lock()

		(*waiting)++

		if (*waiting) == numOfRoutines && jobs.Empty() {
			//fmt.Println("closing")
			for i := 0; i < numOfRoutines; i++ {
				jobs.Push(nil)
			}
			cond.Broadcast()
		}

		for jobs.Empty() {
			cond.Wait()
		}

		job = jobs.Pop()
		//fmt.Println(job)

		(*waiting)--

		cond.L.Unlock()

		if job == nil {
			//fmt.Println(strconv.Itoa(id) + " leaving")
			break
		}

		//fmt.Println("working")
		job()
	}
}
