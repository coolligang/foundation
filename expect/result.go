package expect

import (
	"time"
)

type Result struct {
	config Config
	chain  Chain
	rtt    *time.Duration
}
