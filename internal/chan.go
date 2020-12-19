package internal

type CloseChan struct {
	done chan struct{}
}

func NewCloseChan() *CloseChan {
	return &CloseChan{
		done: make(chan struct{}),
	}
}

func (r *CloseChan) Chan() chan struct{} {
	return r.done
}

func (r *CloseChan) IsClosed() bool {
	select {
	case <-r.done:
		return true
	default:
	}
	return false
}

func (r *CloseChan) Close() {
	if !r.IsClosed() {
		close(r.done)
	}
}
