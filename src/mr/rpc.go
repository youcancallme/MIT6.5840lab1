package mr

//
// RPC definitions.
//
// remember to capitalize all names.所有名称首字母都要大写
//

import (
	"os"
	"strconv"
)

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.实现RPC定义

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
// coordinator 会在 tmp 下生成一个带有当前用户 ID 的 socket 文件名，
// 这样每个用户都不会冲突，也能保证在 Athena 环境下正常运行。
func coordinatorSock() string {
	//文件名前缀
	s := "/var/tmp/5840-mr-"
	//获取当前用户的 UID，并拼接到文件名后面，使每个用户的 socket 文件名唯一。
	s += strconv.Itoa(os.Getuid())
	return s
}
