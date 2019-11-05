package expect

import (
	"testing"
)

type Expect struct {
	config   Config
	builders []func(*TestCase)
	matchers []func(*Result)
}

func NewExpect(t *testing.T, baseURL string) *Expect {
	return WithConfig(Config{})
}

func WithConfig(config Config) *Expect {
	if config.Reporter == nil {
		panic("config.Reporter is nil")
	}
	return &Expect{
		config: config,
	}
}
