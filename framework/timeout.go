package framework

import (
	"context"
	"fmt"
	"log"
	"time"
)

func TimeoutHandler(fun Controller, d time.Duration) Controller {

	return func(c *Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationContext, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		c.request.WithContext(durationContext)

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			fun(c)
			finish <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			log.Println(p)
			c.responseWriter.WriteHeader(500)
		case <-finish:
			fmt.Println("finish")
		case <-durationContext.Done():
			c.SetHasTimeout()
			c.responseWriter.Write([]byte("timeout"))
		}
		return nil
	}
}
