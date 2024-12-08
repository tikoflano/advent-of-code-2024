package solution202406

import "sync"

func WaitAndClose[T interface{}](c chan T, wg *sync.WaitGroup) {
	wg.Wait()
	close(c)
}
