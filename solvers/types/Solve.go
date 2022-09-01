package types

import (
	s "main/utils/structs"
	"sync"
)

type Job func(conf *s.Configuration, index int)
type GoJob func(limit int, solution *s.Configuration, jobs *s.Stack, lock *sync.RWMutex, cond *sync.Cond) Job
