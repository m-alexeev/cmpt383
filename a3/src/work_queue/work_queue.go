package work_queue

// import("fmt")

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	//* initialize struct;
	q := new(WorkQueue)
	// TODO start nWorkers workers as goroutines
	q.Results = make (chan interface{}, maxJobs)
	q.Jobs = make (chan Worker, maxJobs) 
	for i := uint(0); i < nWorkers; i ++ {
		go q.worker()
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs (until shut down).
func (queue WorkQueue) worker() {
	// TODO: Listen on the .Jobs channel for incoming tasks. 
	// ! For each task...
		// * run tasks by calling .Run(),
		// * send the return value back on Results channel.
		// * Exit (return) when .Jobs is closed.
	for elem  := range queue.Jobs{
		task := elem
		queue.Results <- task.Run()
		// if len(queue.Jobs) == 0 {
		// 	return
		// }
	}
}

func (queue WorkQueue) Enqueue(work Worker) {
	// * put the work into the Jobs channel so a worker can find it and start the task.
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	// * close .Jobs and remove all remaining jobs from the channel.
	close(queue.Jobs)
	for range queue.Jobs{
		<-queue.Jobs
	}
}
