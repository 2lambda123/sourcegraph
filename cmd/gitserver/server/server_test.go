package server

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"strings"
	"testing"
)

type Test struct {
	Name         string
	Request      *http.Request
	ExpectedCode int
	ExpectedBody string
}

func TestRequest(t *testing.T) {
	tests := []Test{
		{
			Name:         "WithOutput",
			Request:      httptest.NewRequest("POST", "/exec", strings.NewReader(`{"Repo": "github.com/gorilla/mux", "Args": ["testcommand"]}`)),
			ExpectedCode: http.StatusOK,
			ExpectedBody: "testoutput",
		},
		{
			Name:         "WithoutOutput",
			Request:      httptest.NewRequest("POST", "/exec", strings.NewReader(`{"Repo": "github.com/gorilla/mux", "Args": ["nooutput"]}`)),
			ExpectedCode: http.StatusOK,
			ExpectedBody: "",
		},
		{
			Name:         "NonexistingRepo",
			Request:      httptest.NewRequest("POST", "/exec", strings.NewReader(`{"Repo": "github.com/gorilla/doesnotexist", "Args": ["testcommand"]}`)),
			ExpectedCode: http.StatusConflict,
			ExpectedBody: `{"RepoNotFound":true,"CloneInProgress":false,"Error":"","ExitStatus":0,"Stderr":""}`,
		},
		{
			Name:         "Error1",
			Request:      httptest.NewRequest("POST", "/exec", strings.NewReader(`{"Repo": "github.com/gorilla/mux", "Args": ["testerror1"]}`)),
			ExpectedCode: http.StatusConflict,
			ExpectedBody: `{"RepoNotFound":false,"CloneInProgress":false,"Error":"testerror","ExitStatus":0,"Stderr":""}`,
		},
		{
			Name:         "Error2",
			Request:      httptest.NewRequest("POST", "/exec", strings.NewReader(`{"Repo": "github.com/gorilla/mux", "Args": ["testerror2"]}`)),
			ExpectedCode: http.StatusConflict,
			ExpectedBody: `{"RepoNotFound":false,"CloneInProgress":false,"Error":"","ExitStatus":1,"Stderr":"teststderr"}`,
		},
		{
			Name:         "EmptyBody",
			Request:      httptest.NewRequest("POST", "/exec", nil),
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: `EOF`,
		},
		{
			Name:         "EmptyInput",
			Request:      httptest.NewRequest("POST", "/exec", strings.NewReader("{}")),
			ExpectedCode: http.StatusConflict,
			ExpectedBody: `{"RepoNotFound":true,"CloneInProgress":false,"Error":"","ExitStatus":0,"Stderr":""}`,
		},
	}

	s := &Server{ReposDir: "/testroot"}
	h := s.Handler()

	repoExists = func(dir string) bool {
		return dir == "/testroot/github.com/gorilla/mux"
	}
	runCommand = func(cmd *exec.Cmd) (error, int) {
		switch cmd.Args[1] {
		case "testcommand":
			cmd.Stdout.Write([]byte("testoutput"))
		case "testerror1":
			return errors.New("testerror"), 0
		case "testerror2":
			cmd.Stderr.Write([]byte("teststderr"))
			return nil, 1
		}
		return nil, 0
	}
	noUpdates = true

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := httptest.ResponseRecorder{Body: new(bytes.Buffer)}
			h.ServeHTTP(&w, test.Request)

			if w.Code != test.ExpectedCode {
				t.Errorf("wrong status: expected %d, got %d", test.ExpectedCode, w.Code)
			}

			body := strings.TrimSpace(w.Body.String())
			if body != test.ExpectedBody {
				t.Errorf("wrong body: expected %q, got %q", test.ExpectedBody, body)
			}
		})
	}
}
