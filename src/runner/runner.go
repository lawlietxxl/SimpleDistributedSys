package runner

import (
	"args_struct"
	"net"
	"net/rpc"
	"strconv"
	"net/http"
	"fmt"
	"os"
	"task"
)

type basicRunner struct {
	ipString string
	port     int
	listener *net.Listener
}

func newBasicRunner(ipString string, port int) *basicRunner {
	br := new(basicRunner)
	br.ipString = ipString
	br.port = port
	return br
}




func (b *basicRunner) ToIpString() string {
	return b.ipString+":"+strconv.Itoa(b.port)
}

type Master struct {
	*basicRunner
	workers []*Worker
	workingLocation int
}

func NewMaster(ipString string, port int) (*Master){
	m := new(Master)
	m.basicRunner = newBasicRunner(ipString, port)
	m.workingLocation = 0
	return m
}
func (m *Master) StartServer(){
	rpc.Register(m)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp" , ":"+strconv.Itoa(m.port))
	checkError(err)
	go http.Serve(l, nil)
}
func (m *Master) LogWorkers(){
	for _, w:= range m.workers{
		fmt.Printf("%s:%d\n", w.ipString, w.port)
	}
}
func (m *Master) nextWorker()(*Worker){
	m.workingLocation = (m.workingLocation+1)%len(m.workers)
	return m.workers[m.workingLocation]
}



func (m *Master) RegisterWorker(w *args_struct.RegisterWorkerArgs, reply *int) error {
	nw := NewWorker(w.IpString, w.Port)
	m.workers = append(m.workers, nw)
	return nil
}

func (m *Master) FindWorker(a *args_struct.FindWorkerArgs, reply *string) error {
	w := m.nextWorker()
	*reply = w.ToIpString()
	return nil
}






type Worker struct {
	*basicRunner
}

func NewWorker(ipString string, port int)(*Worker) {
	w := new(Worker)
	w.basicRunner = newBasicRunner(ipString, port)
	return  w
}

func (w *Worker)StartServer(){
	rpc.Register(w)
	//rpc.HandleHTTP()
	l, err := net.Listen("tcp" , ":"+strconv.Itoa(w.port))
	checkError(err)
	go http.Serve(l, nil)
}
func (w *Worker)Register(ipString string){
	client, _ := rpc.DialHTTP("tcp", ipString)
	a := args_struct.RegisterWorkerArgs{w.ipString, w.port}
	var reply int
	client.Call(task.MASTER_TASK.REGISTER_WORKER, &a, &reply)
}


func (*Worker) Plus(a *args_struct.ArithArgs, rep *float64) error {
	*rep = a.Arg_l + a.Arg_r
	return nil
}

func (*Worker) Minus(a *args_struct.ArithArgs, rep *float64) error{
	*rep = a.Arg_l - a.Arg_r
	return nil
}

func (*Worker) Divide(a *args_struct.ArithArgs, rep *float64) error{
	*rep = a.Arg_l / a.Arg_r
	return nil
}

func (*Worker) Multiple(a *args_struct.ArithArgs, rep *float64) error {
	*rep = a.Arg_l * a.Arg_r
	return nil
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}