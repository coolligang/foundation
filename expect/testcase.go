package expect

type TestCase struct {
	config   Config
	chain    Chain
	matchers []func(*Result)
}
