package mr

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

// 1管理任务分配：将 map 和 reduce 任务分配给 worker。
// 2跟踪任务状态：记录每个任务（map/reduce）的状态（未分配、进行中、已完成）。
// 3处理 worker 的 RPC 请求：实现 RPC handler，worker 通过 RPC 请求任务、汇报任务完成等。
// 4判断作业是否完成：实现 Done() 方法，判断所有任务是否都已完成。
// 5初始化协调者：在 MakeCoordinator 函数中初始化任务、状态等。

// 定义
type Coordinator struct {
	// Your definitions here.

}

// Your code here -- RPC handlers for the worker to call.

// 3 处理rpc请求：worker 通过 RPC 请求任务、汇报任务完成等。
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

// 2 跟踪任务状态：
// start a thread that listens for RPCs from worker.go
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// 4 done returns true if the entire job has finished.
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
func (c *Coordinator) Done() bool {
	ret := false

	// Your code here.
	//检查是否所有的map和reduce任务都已完成
	// 这里可以检查任务状态，判断是否所有任务都已完成

	return ret
}

// manager角色：初始化任务和状态
// 这里实现任务分配吗？
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}

	// Your code here.

	c.server()
	return &c
}
