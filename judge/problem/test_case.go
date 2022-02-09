package problem

type TestCase struct {
	Input          string
	ExpectedOutput []string // at least one of these has to match
}

type TestCaseResult struct {
	Stdout         string // their original stdout
	ExpectedStdout string
	Passed         bool
	Elapsed        int64 // time it took to run test case in ms
}
