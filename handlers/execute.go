package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"straitjacket/engine"
)

type ExecutionResult struct {
	Stdout     string `json:"stdout"`
	Stderr     string `json:"stderr"`
	ExitStatus int    `json:"exit_status"`
	Time       string `json:"time"`
	Error      string `json:"error"`
}

func (ctx *Context) ExecuteHandler(res http.ResponseWriter, req *http.Request) {
	languageName := req.FormValue("language")
	source := req.FormValue("source")
	stdin := req.FormValue("stdin")
	timelimit := req.FormValue("timelimit")

	log.Println(languageName, source, stdin, timelimit)

	language, err := ctx.Engine.FindLanguage(languageName)
	if err != nil {
		panic(err)
	}

	runResult, err := language.Run(&engine.RunOptions{
		Source: source,
		Stdin:  stdin,
	})
	if err != nil {
		panic(err)
	}

	result := ExecutionResult{
		Stdout:     runResult.Stdout,
		Stderr:     runResult.Stderr,
		ExitStatus: runResult.ExitCode,
	}
	json, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	_, err = res.Write(json)
	if err != nil {
		panic(err)
	}
}
