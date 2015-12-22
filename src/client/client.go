package client
import (
	"net/rpc"
	"args_struct"
	"task"
	"fmt"
)

type Client struct {
}

func (*Client)GetWorkerIpString(masterIpString string)string{
	client, err := rpc.DialHTTP("tcp", masterIpString)
	if err != nil {
		fmt.Println(err)
		//os.Exit(2)
	}
	defer client.Close()
	a := args_struct.FindWorkerArgs{}
	var reply string
	client.Call(task.MASTER_TASK.FIND_WORKER, &a, &reply)
	return reply
}

func (*Client)WorkerPlus(ipString string) float64{
	client, _ := rpc.DialHTTP("tcp", ipString)
	defer client.Close()
	a := args_struct.ArithArgs{1,1}
	var reply float64
	client.Call(task.TASK.PLUS, &a, &reply)
	return reply
}
func (*Client)WorkerMinus(ipString string) float64{
	client, _ := rpc.DialHTTP("tcp", ipString)
	defer client.Close()
	a := args_struct.ArithArgs{1,1}
	var reply float64
	client.Call(task.TASK.MINUS, &a, &reply)
	return reply
}
func (*Client)WorkerDivide(ipString string) float64{
	client, _ := rpc.DialHTTP("tcp", ipString)
	defer client.Close()
	a := args_struct.ArithArgs{1,1}
	var reply float64
	client.Call(task.TASK.DIVIDE, &a, &reply)
	return reply
}
func (*Client)WorkerMultiple(ipString string) float64{
	client, _ := rpc.DialHTTP("tcp", ipString)
	defer client.Close()
	a := args_struct.ArithArgs{1,1}
	var reply float64
	client.Call(task.TASK.MULTIPLE, &a, &reply)
	return reply
}

func (*Client)WorkArith(ipString string, arg *args_struct.ArithArgs, task string) float64{
	client, err := rpc.DialHTTP("tcp", ipString)
	if err != nil {
		fmt.Println(err)
		//os.Exit(2)
	}
	defer client.Close()
	var reply float64
	client.Call(task, arg, &reply)

	//client must be closed, or there will be no more memory
	return reply
}