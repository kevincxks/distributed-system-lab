package mapreduce

import (
	"fmt"
	"sync"
)

// schedule starts and waits for all tasks in the given phase (Map or Reduce).
func (mr *Master) schedule(phase jobPhase) {
	var ntasks int
	var nios int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mr.files)
		nios = mr.nReduce
	case reducePhase:
		ntasks = mr.nReduce
		nios = len(mr.files)
	}

	fmt.Printf("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, nios)

	// All ntasks tasks have to be scheduled on workers, and only once all of
	// them have been completed successfully should the function return.
	// Remember that workers may fail, and that any given worker may finish
	// multiple tasks.
	//
	// TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
	//
	var wg sync.WaitGroup
	for i := 0; i < ntasks; i++ {
		wg.Add(1)
		taskArgs := DoTaskArgs{
			JobName:       mr.jobName,
			File:          mr.files[i],
			Phase:         phase,
			TaskNumber:    i,
		}
		if phase == "Map" {
			taskArgs.NumOtherPhase = mr.nReduce
		} else {
			taskArgs.NumOtherPhase = len(mr.files)
		}
		tmpWorker := <-mr.registerChannel
		go func(tmp string, args *DoTaskArgs, wg *sync.WaitGroup) {
			flag := call(tmp, "Worker.DoTask", args, new(struct{}))
			if !flag {
				tmp2 := <-mr.registerChannel
				call(tmp2, "Worker.DoTask", args, new(struct{}))
				mr.registerChannel <- tmp2
			}
			if flag {
				mr.registerChannel <- tmp
			}
			wg.Done()
		}(tmpWorker, &taskArgs, &wg)
	}
	wg.Wait()
	fmt.Printf("Schedule: %v phase done\n", phase)
}
