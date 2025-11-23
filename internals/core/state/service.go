package app_state

import (
	"context"
	"fmt"
	"net"
	"time"
)

const TIMEOUT = 2

type AppStateSerivce struct {
	ConnectivityState *ConnectivityBus
	probeTimeout      int
	currentTime       int64
	nextProbeTime     int64
}

func NewAppStateService() *AppStateSerivce {
	bus := NewConnectityBus()
	return &AppStateSerivce{
		ConnectivityState: bus,
		probeTimeout:      TIMEOUT,
	}
}

func (ass *AppStateSerivce) probe(ctx context.Context) AppState {
	d := net.Dialer{Timeout: time.Second * 2}

	conn, err := d.DialContext(ctx, "tcp", "8.8.8.8:53")

	fmt.Printf("Probing %v", err)

	if err != nil {
		return AppOfflineState
	}
	_ = conn.Close()
	return AppOnlineState
}

func (ass *AppStateSerivce) Check(ctx context.Context, isForce bool) AppState {
	prevStatus := ass.ConnectivityState.Status()
	status := ass.probe(ctx)

	if status != prevStatus || isForce {
		ass.ConnectivityState.Publish(status)
		ass.probeTimeout = TIMEOUT
	} else if prevStatus == AppOfflineState {
		if ass.probeTimeout < 320 {
			ass.probeTimeout *= 2
		}
	}

	return status
}

func (ass *AppStateSerivce) Heartbeat(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	ass.currentTime = time.Now().Unix()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ass.currentTime = time.Now().Unix()
			if ass.currentTime >= ass.nextProbeTime {
				ass.Check(ctx, false)
				ass.nextProbeTime = ass.currentTime + int64(ass.probeTimeout)
			}
		}
	}
}
