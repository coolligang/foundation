package capacity

type TestResult struct {
	Detail []TestResultElem
}

func newTestResult() (result *TestResult) {
	result = new(TestResult)
	result.Detail = make([]TestResultElem, 0)
	return result
}

func addTestResultElem(result *TestResult, resultElem TestResultElem) {
	result.Detail = append(result.Detail, resultElem)
}

type TestResultElem struct {
	ConcurrentUser int
	Samples        []int64
	Summary        []interface{}
	ErrorOccur     []bool
}

func newResultElem(concurrentUser int) (resultElem *TestResultElem) {
	resultElem = new(TestResultElem)
	resultElem.ConcurrentUser = concurrentUser
	resultElem.Samples = make([]int64, concurrentUser)
	resultElem.Summary = make([]interface{}, concurrentUser)
	resultElem.ErrorOccur = make([]bool, concurrentUser)
	return resultElem
}
