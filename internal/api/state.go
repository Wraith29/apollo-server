package api

import "sync"

type State struct {
	updateInProgress bool
	mutex            *sync.Mutex
}

func NewState() *State {
	return &State{
		updateInProgress: false,
		mutex:            &sync.Mutex{},
	}
}
