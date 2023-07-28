package promise

func NewPromise[R, E any](f func(chan<- R, chan<- E)) *Promise[R, E] {
	resolve, reject := make(chan R), make(chan E)
	go f(resolve, reject)
	return &Promise[R, E]{resolve, reject}
}

func (p *Promise[R, E]) Await() (R, E) {
	defer func() {
		close(p.resolve)
		close(p.reject)
	}()

	select {
	case r := <-p.resolve:
		var e E
		return r, e
	case e := <-p.reject:
		var r R
		return r, e
	}

}

type Promise[R, E any] struct {
	resolve chan R
	reject  chan E
}
