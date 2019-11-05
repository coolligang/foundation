package expect

import (
	"time"
)

type Reporter interface {
	TestCase()
	Result(time.Duration)
}
