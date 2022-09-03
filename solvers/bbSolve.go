package solvers

import (
	t "main/solvers/types"
	u "main/utils"
	s "main/utils/structs"
	"sync"
)

func upperBound(current int, index int, values *[]int) int {
	res := current

	for i, v := range *values {
		if i >= index {
			res += v
		}
	}

	return res
}

func BBsolve(limit int, initConf *s.Configuration, index int, solution *s.Configuration) {

	if initConf.Weight > limit ||
		upperBound(initConf.Value, index, initConf.Values) < solution.Value {
		return
	}

	if index >= initConf.Len() {
		if initConf.Weight <= limit && initConf.Value > solution.Value {
			solution.Copy(initConf)
		}
		return
	}

	conf0 := initConf.Clone()
	conf1 := initConf.Clone()
	conf1.AddElement(index)

	BBsolve(limit, conf0, index+1, solution)
	BBsolve(limit, conf1, index+1, solution)
}

func GetBBSolveJob(limit int, solution *s.Configuration, jobs *s.Stack[func()], lock *sync.RWMutex, cond *sync.Cond) t.Job {

	var job t.Job

	job = func(conf *s.Configuration, index int) {
		if conf.Weight > limit {
			return
		}

		if index >= conf.Len() {
			u.DoubleCheck(func() bool { return conf.Value > solution.Value },
				func() { solution.Copy(conf) }, func() {}, lock, false)
			return
		}

		var possible bool
		ub := upperBound(conf.Value, index, conf.Values)

		u.DoubleCheck(func() bool { return ub < solution.Value },
			func() { possible = false }, func() { possible = true }, lock, true)

		if !possible {
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
