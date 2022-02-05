package problem

import (
	"fmt"
	"io/ioutil"
)

// DiskFetcher is a ProblemDataFetcher that uses the disk
type DiskFetcher struct {
	ParentDir string
}

func (d *DiskFetcher) CodeTemplate(problemId string, ext string) (string, error) {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/templates/%s%s", d.ParentDir, problemId, ext))
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func (d *DiskFetcher) HookTemplate(problemId string, ext string) (string, error) {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/hooks/%s%s", d.ParentDir, problemId, ext))
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func (d *DiskFetcher) Solution(problemId string, ext string) (string, error) {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/solutions/%s%s", d.ParentDir, problemId, ext))
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func (d *DiskFetcher) TestCase(problemId string, testCaseId int) (*TestCase, error) {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/test_cases/input/%s_%d", d.ParentDir, problemId, testCaseId))
	if err != nil {
		return nil, err
	}
	input := string(dat)

	dat, err = ioutil.ReadFile(fmt.Sprintf("%s/test_cases/output/%s_%d", d.ParentDir, problemId, testCaseId))
	if err != nil {
		return nil, err
	}

	output := string(dat)

	return &TestCase{
		Input:          input,
		ExpectedOutput: []string{output},
	}, nil
}
