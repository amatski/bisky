package judge

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/amatski/bisky/compiler"
	"github.com/amatski/bisky/judge/codegen"
	"github.com/amatski/bisky/judge/problem"
)

type RequestHandler struct {
	Fetcher problem.ProblemDataFetcher
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{Fetcher: &problem.DiskFetcher{ParentDir: "data"}}
}

type JudgeRequest struct {
	Language  string `json:"language"`
	Code      string `json:"code"` // their solution
	Problem   string `json:"problem"`
	TestCases []*problem.TestCase
}

type JudgeResponse struct {
	Results []problem.TestCaseResult
}

func parseAnswers(stdout string, secret string, numTests int) []string {
	s := strings.Split(stdout, "\n")
	answers := make([]string, 0, numTests)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == secret {
			answers = append(answers, s[i+1])
			i++
		}
	}

	return answers
}

// JudgeSolution judges the submitted solution
func (h *RequestHandler) JudgeSolution(req JudgeRequest) (*JudgeResponse, error) {
	hookTemplate, err := h.Fetcher.HookTemplate(req.Problem, codegen.LanguageToExt[req.Language])
	if err != nil {
		return nil, err
	}

	solution, err := codegen.NewCode(req.Language)
	if err != nil {
		return nil, err
	}
	secret := codegen.RandomSecret()

	solution.AddCode(hookTemplate)

	testCaseCode, err := codegen.TestCaseCalls(req.TestCases, req.Language, req.Problem, secret)
	if err != nil {
		return nil, err
	}

	solution.InjectValue("tests", testCaseCode)
	solution.InjectValue("solution", req.Code)

	out, err := compiler.NewRequestHandler().HandleCompileRequest(context.Background(), compiler.CompileEvent{
		Code:     solution.Code(),
		Language: req.Language,
	})

	if err != nil {
		return nil, err
	}

	// Parse the stdout to get their answers per test
	answers := parseAnswers(out, secret, len(req.TestCases))

	if len(answers) != len(req.TestCases) {
		return nil, errors.New("answers and test cases are not 1:1")
	}

	// convert answers to []problem.TestCaseResult
	// we would have to evaluate it against uhh expected
	results := []problem.TestCaseResult{}
	for idx, answerStdout := range answers {
		// should match testCases[idx].ExpectedOutput
		passed := false
		for _, expected := range req.TestCases[idx].ExpectedOutput {
			if expected == answerStdout {
				passed = true
				break
			}
		}

		res := problem.TestCaseResult{
			Passed:  passed,
			Stdout:  answerStdout,
			Elapsed: time.Second, // we don't know this yet
		}
		results = append(results, res)
	}

	return &JudgeResponse{
		Results: results,
	}, err
}
