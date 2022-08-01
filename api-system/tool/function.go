package tool

import "github.com/axiangcoding/ax-web/logging"

// GoWithRecover 在协程中如果出现fatal不会导致整个程序崩溃
func GoWithRecover(f func()) {
	go func(handler func()) {
		defer func() {
			if r := recover(); r != nil {
				logging.Errorf("recover from go func error. %s", r)
			}

		}()
		handler()
	}(f)
}
