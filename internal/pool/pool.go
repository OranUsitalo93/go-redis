package pool

import (
	"context"
	"net"
	"sync/atomic"
)

// ... existing code ...

func (p *ConnPool) newConn(ctx context.Context) (*Conn, error) {
	cn, err := p.dial(ctx)
	if err != nil {
		return nil, err
	}

	if p.opt.OnConnect != nil {
		if err := p.opt.OnConnect(ctx, cn); err != nil {
			// Ensure connection is closed and pool counter is decremented on hook failure
			_ = cn.Close()
			atomic.AddInt32(&p.poolSize, -1)
			return nil, err
		}
	}

	return cn, nil
}

// ... existing code ...