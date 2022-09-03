package types

import (
	s "main/utils/structs"
	"sync"
)

type Job func(conf *s.Configuration, index int)
type GetJob func(limit int, solution *s.Configuration, jobs *s.Stack[func()], lock *sync.RWMutex, cond *sync.Cond) Job
