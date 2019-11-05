package capacity

import (
	"sync"
	"time"
)

type TestLogic struct {
	MaxConcurrentUser              int
	TestTimeForEachConCurrentUser  time.Duration
	WaitTimeForEachCoroutineSetup  time.Duration
	WaitTimeForCoroutineSingleStep time.Duration
	WaitGroup                      sync.WaitGroup
	IsStartFlag                    bool
	IsStopFlag                     bool
	Result                         *TestResult
}

func (capacity *TestLogic) LoadTest(instance TestBehaveInterface) (success bool, result TestResult) {
	capacity.Result = newTestResult()
	concurrentUser := 1
	for index := 1; concurrentUser < capacity.MaxConcurrentUser; index++ {
		success = capacity.loadTestForDefinedConCurrentUser(instance, index, concurrentUser)
		if !success {
			break
		}
		concurrentUser = instance.NextStep(concurrentUser)
	}
	return success, result
}

func (capacity *TestLogic) loadTestForDefinedConCurrentUser(instance TestBehaveInterface, index, concurrentUser int) (success bool) {

	resultElem := newResultElem(concurrentUser)
	addTestResultElem(capacity.Result, *resultElem)

	capacity.IsStartFlag = false
	capacity.IsStopFlag = false

	for coroutine := 1; coroutine <= concurrentUser; coroutine++ {
		capacity.WaitGroup.Add(1)
		go instance.TestMainLogic(index, coroutine)
	}
	time.Sleep(capacity.WaitTimeForEachCoroutineSetup * time.Duration(concurrentUser))
	capacity.WaitGroup.Add(1)
	go func() {
		defer capacity.WaitGroup.Done()
		capacity.IsStartFlag = true
		time.Sleep(capacity.TestTimeForEachConCurrentUser)
		capacity.IsStopFlag = true
	}()
	capacity.WaitGroup.Wait()
	return instance.IsTestPassed(index, concurrentUser)
}
