package capacity

type TestBehaveInterface interface {
	TestMainLogic(index, coroutine int)
	IsTestPassed(index, concurrentUser int) (success bool)
	NextStep(prev int) (now int)
}
