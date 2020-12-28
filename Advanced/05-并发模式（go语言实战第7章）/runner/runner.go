package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

/* 使用通道来监视程序的执行时间

 */

type Runner struct {
	// 收发os.Signal接口类型的值，用来从主机操作系统接收中断事件
	interrupt chan os.Signal

	// 报告处理任务已经完成
	complete chan error

	// 报告处理任务已经超时
	timeout <-chan time.Time

	// 一组以索引顺序依次执行的函数
	tasks []func(int)
}

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

func NewRunner(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// 添加一个tasks任务附加到Runner上。
// 这个tasks是一个接收一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// 执行全部任务，并监视通道事件
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时发出的信号
	case err := <-r.complete:
		return err
	// 当任务处理程序运行超时时发出的信号
	case <-r.timeout:
		return ErrTimeout
	}
}

// run执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测到操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

// gotInterrupt 验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件被触发时发出的信号
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
