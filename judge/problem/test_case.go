package problem

type TestCase struct {
	Input          string
	ExpectedOutput []string // at least one of these has to match
}

type TestCaseResult struct {
	Input          string // their input
	Stdout         string // their stdout
	ExpectedOutput string // the test's stdout
	Passed         bool
	Elapsed        int64 // time it took to run test case in ms
}
