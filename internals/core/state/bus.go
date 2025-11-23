package app_state

import (
	"sync"
)

type AppState int

const (
	AppOfflineState AppState = 0
	AppOnlineState  AppState = 1
)

type ConnectivityBus struct {
	mu     sync.RWMutex
	status AppState
	subs   map[int]chan AppState
	nextID int
}

func NewConnectityBus() *ConnectivityBus {
	return &ConnectivityBus{
		subs: make(map[int]chan AppState),
	}
}

func (bus *ConnectivityBus) Subscribe() (<-chan AppState, func()) {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	id := bus.nextID
	bus.nextID++

	ch := make(chan AppState, 16)
	bus.subs[id] = ch

	unsub := func() {
		bus.mu.Lock()
		defer bus.mu.Unlock()

		if c, ok := bus.subs[id]; ok {
			close(c)
			delete(bus.subs, id)
		}
	}

	return ch, unsub
}

func (bus *ConnectivityBus) Publish(state AppState) {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	bus.status = state

	for i, ch := range bus.subs {
		select {
		case ch <- state:
		default:
			close(ch)
			delete(bus.subs, i)
		}
	}
}

func (bus *ConnectivityBus) Status() AppState {
	return bus.status
}
