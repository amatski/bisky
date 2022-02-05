package problem

// ProblemDataFetcher is responsible for fetching data for the problem/validator
type ProblemDataFetcher interface {
	CodeTemplate(problemId string, ext string) (string, error)
	HookTemplate(problemId string, ext string) (string, error)
	Solution(problemId string, ext string) (string, error)
	TestCase(problemId string, testCaseId int) (*TestCase, error)
}
