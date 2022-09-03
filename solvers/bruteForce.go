package solvers

import (
	t "main/solvers/types"
	u "main/utils"
	s "main/utils/structs"
	"sync"
)

func BruteSolve(limit int, initConf *s.Configuration, index int, solution *s.Configuration) {

	if index >= initConf.Len() {
		if initConf.Weight <= limit && initConf.Value > solution.Value {
			solution.Copy(initConf)
		}
		return
	}

	conf0 := initConf.Clone()
	conf1 := initConf.Clone()
	conf1.AddElement(index)

	BruteSolve(limit, conf0, index+1, solution)
	BruteSolve(limit, conf1, index+1, solution)
}

func GetBruteSolveJob(limit int, solution *s.Configuration, jobs s.Container[func()], lock *sync.RWMutex, cond *sync.Cond) t.Job {
	var job t.Job

	job = func(conf *s.Configuration, index int) {
		if index >= conf.Len() {
			u.DoubleCheck(func() bool { return conf.Weight <= limit && conf.Value > solution.Value },
				func() { solution.Copy(conf) }, func() {}, lock, false)
			return
		}

		conf0 := conf.Clone()
		conf1 := conf.Clone()
		conf1.AddElement(index)

		if conf.Len()-index > u.HeightLimit {
			cond.L.Lock()
			jobs.Push(func() { job(conf0, index+1) })
			cond.L.Unlock()
			cond.Signal()
		} else {
			job(conf0, index+1)
		}
		job(conf1, index+1)
	}

	return job
}
