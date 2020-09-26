package controller

type taskState int

// state of the task
const (
	CREATED taskState = iota
	PENDING
	RUNNING
	PAUSED
	SUCCESSFUL
	ERROR
)

type task struct {
	target module.Target
	tool   module.Tool
	state  taskState
}
