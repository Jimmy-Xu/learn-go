package main

import (
	"fmt"
	"time"
	"strings"
)

type Task struct {
	name string
}

type Result struct {
	result string
}

var (
	TIMEOUT = time.Duration(10*time.Second)
	chSend  = make(chan Task)
)

func main() {
	defer fmt.Println("bye")

	doneCh := Process(chSend,0)
	time.Sleep(time.Duration(1*time.Second))

	fmt.Println("start")

	//发送
	for i:=0; i<=10; i++ {
		task := Task{name: fmt.Sprintf("task-0%v", i)}
		select {
		case chSend <- task: //send task
			fmt.Printf("\nsend task:%v\n", task.name)
		case <-doneCh: //wait timeout
			fmt.Println("\nstop send task")
			return
		}
		time.Sleep(time.Duration(2*time.Second))
	}
}

func f(task Task) Result {
	fmt.Printf("f task:%v\n", task.name)
	time.Sleep(time.Duration(1*time.Second))
	return Result{
		result: strings.ToUpper(task.name),
	}
}

// 执行方
func doTask(doneCh <-chan struct{}, taskCh <-chan Task) (chan Result) {
	outCh := make(chan Result)
	go func() {
		fmt.Println("doTask: waiting")
		// close 是为了让调用方的range能够正常退出
		defer close(outCh)
		defer close(chSend)
		for {
			select {
			case task:= <-taskCh: //receive task
				fmt.Println("receive task")
				outCh <-f(task) // send result
			case <-doneCh: //wait timeout
				fmt.Println("doTask: doneCh")
				return
			}
		}
	}()

	return outCh
}

// 调用方
func Process(taskCh <-chan Task, num int)  (chan struct{}) {
	doneCh := make(chan struct{})

	outCh := doTask(doneCh, taskCh)

	// 超时检测
	go func() {
		<- time.After(TIMEOUT)
		fmt.Println("\ntime outCh, broadcast 'doneCh'")
		close(doneCh)
		time.Sleep(time.Duration(1*time.Second))
	}()

	// 接收结果
	go func () {
		// 因为goroutine执行完毕，或者超时，导致out被close
		for {
			res, ok := <-outCh //receive result
			if ok {
				fmt.Printf("reveive result:%v\n",res.result)
			} else {
				fmt.Printf("channel closed\n")
				break
			}
		}
	}()

	return doneCh
}
