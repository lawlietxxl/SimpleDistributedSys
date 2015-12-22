package main
import (
	"runner"
	"client"
	"args_struct"
	"task"
	"fmt"
	"time"
)

func main() {

	m := runner.NewMaster("", 2222)
	m.StartServer()


	w1 := runner.NewWorker("", 2233)
	w2 := runner.NewWorker("", 2244)
	w3 := runner.NewWorker("", 2255)
	w1.StartServer()
	w2.StartServer()
	w3.StartServer()

	w1.Register(m.ToIpString())
	w2.Register(m.ToIpString())
	w3.Register(m.ToIpString())

	m.LogWorkers()

	c1 := new(client.Client)
	c2 := new(client.Client)
	c3 := new(client.Client)

	flagChans := make(chan bool, 3)
	flagChans <- true
	flagChans <- true
	flagChans <- true



	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
			fmt.Print("\n")
		}
	}()


	go runningTest(c1, m, task.TASK.MINUS, flagChans)
	go runningTest(c2, m, task.TASK.MINUS, flagChans)
	go runningTest(c3, m, task.TASK.MINUS, flagChans)




	time.Sleep(time.Second)
	if <- flagChans{
		if <- flagChans{
			if <- flagChans{

			}
		}
	}
}

func runningTest(c *client.Client, m *runner.Master, task string, flagsChans chan bool){
	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
			fmt.Print("\n")
		}
	}()

	<- flagsChans
	for i:=1; i <1000000;i++{
		a := args_struct.ArithArgs{float64(i), float64(i+1)}
		ipString := c.GetWorkerIpString(m.ToIpString())
		if(i%1000 == 0) {
			fmt.Printf("%d iter: %f\n", i, c.WorkArith(ipString, &a, task))
		}
	}
	flagsChans <- true
}