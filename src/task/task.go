package task

const (
	plus = "Worker.Plus"
	minus = "Worker.Minus"
	divide = "Worker.Divide"
	multiple = "Worker.Multiple"

	register_worker = "Master.RegisterWorker"
	find_worker = "Master.FindWorker"
)


type taskCluster struct  {
	PLUS string
	MINUS string
	DIVIDE string
	MULTIPLE string
}

var TASK = taskCluster{
	plus,
	minus,
	divide,
	multiple,
}

type masterTask struct {
	REGISTER_WORKER string
	FIND_WORKER string
}

var MASTER_TASK = masterTask{
	register_worker,
	find_worker,
}