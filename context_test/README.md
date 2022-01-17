# timeout

## http 超时处理

```golang  
func (h *timeoutHandler) ServeHTTP(w ResponseWriter, r *Request) {
	ctx := h.testContext
	if ctx == nil {
		var cancelCtx context.CancelFunc
		ctx, cancelCtx = context.WithTimeout(r.Context(), h.dt)
		defer cancelCtx()
	}
	r = r.WithContext(ctx)
	done := make(chan struct{})
	tw := &timeoutWriter{
		w:   w,
		h:   make(Header),
		req: r,
	}
	panicChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		h.handler.ServeHTTP(tw, r)
		close(done)
	}()
	select {
	case p := <-panicChan:
		panic(p)
	case <-done:
		tw.mu.Lock()
		defer tw.mu.Unlock()
		dst := w.Header()
		for k, vv := range tw.h {
			dst[k] = vv
		}
		if !tw.wroteHeader {
			tw.code = StatusOK
		}
		w.WriteHeader(tw.code)
		w.Write(tw.wbuf.Bytes())
	case <-ctx.Done():
		tw.mu.Lock()
		defer tw.mu.Unlock()
		w.WriteHeader(StatusServiceUnavailable)
		io.WriteString(w, h.errorBody())
		tw.timedOut = true
	}
}
```

## rpc 超时处理

```go  
// UnaryTimeoutInterceptor returns a func that sets timeout to incoming unary requests.
func UnaryTimeoutInterceptor(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		var resp interface{}
		var err error
		var lock sync.Mutex
		done := make(chan struct{})
		// create channel with buffer size 1 to avoid goroutine leak
		panicChan := make(chan interface{}, 1)
		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			lock.Lock()
			defer lock.Unlock()
			resp, err = handler(ctx, req)
			close(done)
		}()

		select {
		case p := <-panicChan:
			panic(p)
		case <-done:
			lock.Lock()
			defer lock.Unlock()
			return resp, err
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
```