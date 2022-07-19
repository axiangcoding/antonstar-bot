package app

import "github.com/axiangcoding/ax-web/logging"

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
