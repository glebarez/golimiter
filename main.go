package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/mirecl/golimiter/internal/analysis"
	"github.com/mirecl/golimiter/internal/linters"
)

// Version golimiter linter.
const Version string = "0.3.4"

func main() {
	jsonFlag := flag.Bool("json", false, "format report")
	versionFlag := flag.Bool("version", false, "version golimiter")

	flag.Parse()

	if *versionFlag {
		fmt.Printf("golimiter %s\n", Version)
		return
	}

	cfg, err := analysis.ReadConfig()
	if err != nil {
		panic(err)
	}

	allIssues := analysis.Run(cfg,
		linters.NewNoGeneric(),
		linters.NewNoInit(),
		linters.NewNoGoroutine(),
		linters.NewNoNoLint(),
		linters.NewNoDefer(),
		linters.NewNoLength(),
		linters.NewNoPrefix(),
	)

	if *jsonFlag {
		if allIssuesBytes, err := json.Marshal(allIssues); err == nil {
			fmt.Println(string(allIssuesBytes))
		}
		return
	}

	for linter, issues := range allIssues {
		for _, issue := range issues {
			position := fmt.Sprintf("%s:%v", analysis.GetPathRelative(issue.Filename), issue.Line)
			fmt.Printf("%s \033[31m%s: %s. \033[0m\033[30m(%s)\033[0m\n", position, linter, issue.Message, issue.Hash)
		}
	}
}
