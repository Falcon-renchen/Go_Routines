package main

import (
	"fmt"
	"time"
)

//定义一个任务类型
type Task struct {
	f func() error	//Task里面有一个具体业务
}
//创建一个Task任务
func NewTask(arg_f func() error) *Task {
	t := Task{f:arg_f}
	return &t
}
//Task需要一个执行业务的方法
func (t *Task) Execute() {
	t.f()	//调用业务中已经绑好的业务方法
}

//定义一个协程池的类型
type Pool struct {
	//对外的Task入口
	EntryChannel chan *Task

	//对内的Task队列
	JobChannel chan *Task

	//协程池中最大的worker数量
	worker_num int
}
//创建Pool函数
func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		JobChannel:   make(chan *Task),
		worker_num:   cap,
	}
	return &p
}
//协程池创建一个Worker，并且让这个Worker工作
func (p *Pool) worker(worker_ID int) {
	//一个worker的具体工作

	//永久的从JobChannel去取任务
	for task := range p.JobChannel {
		//一旦拿到任务，执行这个任务
		task.Execute()
		fmt.Println("worker_ID",worker_ID,"执行完了任务")
	}
}

//执行协程池
func (p *Pool) run() {
	//根据worker_num去工作
	for i:=0; i<p.worker_num; i++ {
		go p.worker(i)
	}
	//从EntryChannel中取数据，然后发送给JobChannel
	for task := range p.EntryChannel {
		//一旦有task读到
		p.JobChannel <- task
	}
}

//主函数，测试协程池的运行
func main() {
	//创建一些任务
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})
	//创建一个协程池,这个协程池能容纳最大的worker数量是4
	p := NewPool(4)
	task_num := 0
	//将这些任务交给协程池Pool
	go func() {
		for {
			p.EntryChannel <- t
			task_num += 1
			fmt.Println("当前一共执行了",task_num,"个任务")
		}
	}()
	//启动pool,让pool开始工作，此时pool回创建worker，让worker工作
	p.run()
}
