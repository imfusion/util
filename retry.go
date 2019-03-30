package util

import (
	"time"

	"github.com/stackcats/iris/utils"
)

var log = utils.NewLogger()

// Stop 类型的错误会中断重试
type Stop struct {
	error
}

// NoRetryError ...
func NoRetryError(err error) Stop {
	return Stop{err}
}

// Retry 调用 fn 最多重试 attempts 次，重试时间 2 倍递增
// attempts：最多重试次数；
// sleep：   调用失败后的等待时间；
// fn：      重试的函数。函数的类型为 func() error
func Retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(Stop); ok {
			return s.error
		}

		if attempts--; attempts > 0 {
			log.Warnf("retry func error: %s. attemps #%d after %s.", err.Error(), attempts, sleep)
			time.Sleep(sleep)
			return Retry(attempts, 2*sleep, fn)
		}
		return err
	}

	return nil
}
