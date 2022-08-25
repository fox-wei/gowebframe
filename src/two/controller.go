package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fox-wei/gowebframe/framework"
)

func FooControllerHandler(c *framework.Context) error {

	finish := make(chan struct{}, 1)       //*负责结束通知
	panicChan := make(chan interface{}, 1) //*异常通知

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second)) //!创建超时Context
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		//*Do real action，模拟业务
		time.Sleep(10 * time.Millisecond)
		c.Json(200, "ok")

		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan: //*异常事件
		c.WriteMux().Lock()
		defer c.WriteMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish: //*结束事件
		fmt.Println("finish")
	case <-durationCtx.Done(): //*超时事件
		c.WriteMux().Lock()
		defer c.WriteMux().Unlock()
		c.Json(500, "time out") //*前端显示异常信息
		c.SetHasTimeout()
	}
	return nil
}
