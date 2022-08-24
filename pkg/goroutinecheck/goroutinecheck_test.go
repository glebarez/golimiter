package goroutinecheck_test

import (
	"testing"

	"github.com/mirecl/golimiter/pkg/goroutinecheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestGoroutine(t *testing.T) {
	testdata := analysistest.TestData()

	TestCases := []struct {
		name string
		pkg  []string
		cfg  *goroutinecheck.Config
	}{
		{
			name: "success analysis",
			pkg:  []string{"a"},
			cfg:  new(goroutinecheck.Config),
		},
		{
			name: "failed to 1 package",
			pkg:  []string{"b"},
			cfg: func() *goroutinecheck.Config {
				Limit := 2
				return &goroutinecheck.Config{Limit: &Limit}
			}(),
		},
		{
			name: "failed limit",
			pkg:  []string{"c"},
			cfg: func() *goroutinecheck.Config {
				Limit := 0
				return &goroutinecheck.Config{Limit: &Limit}
			}(),
		},
		{
			name: "failed limit in all package",
			pkg:  []string{"c", "d"},
			cfg: func() *goroutinecheck.Config {
				Limit := 0
				return &goroutinecheck.Config{Limit: &Limit}
			}(),
		},
		{
			name: "success analysis with test file",
			pkg:  []string{"e"},
			cfg:  new(goroutinecheck.Config),
		},
	}

	for _, testCase := range TestCases {
		t.Run(testCase.name, func(t *testing.T) {
			analyzer := goroutinecheck.New(testCase.cfg)
			analysistest.Run(t, testdata, analyzer, testCase.pkg...)
		})
	}
}
