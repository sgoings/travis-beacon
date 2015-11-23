package handlers

import (
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/sgoings/travis-beacon/storage"
)

var updateRouter *mux.Router

func init() {
	updateRouter = NewRouter(storage.NewInMemoryDB()).(*mux.Router)
}

func TestUpdateCharts(t *testing.T) {
	travisBody := `
{
  "payload": {
    "id": 1,
    "number": "1",
    "status": null,
    "started_at": null,
    "finished_at": null,
    "status_message": "Passed",
    "commit": "62aae5f70ceee39123ef",
    "branch": "master",
    "message": "the commit message",
    "compare_url": "https://github.com/svenfuchs/minimal/compare/master...develop",
    "committed_at": "2011-11-11T11: 11: 11Z",
    "committer_name": "Sven Fuchs",
    "committer_email": "svenfuchs@artweb-design.de",
    "author_name": "Sven Fuchs",
    "author_email": "svenfuchs@artweb-design.de",
    "type": "push",
    "build_url": "https://travis-ci.org/svenfuchs/minimal/builds/1",
    "repository": {
      "id": 1,
      "name": "minimal",
      "owner_name": "svenfuchs",
      "url": "http://github.com/svenfuchs/minimal"
     },
    "config": {
      "notifications": {
        "webhooks": ["http://evome.fr/notifications", "http://example.com/"]
      }
    },
    "matrix": [
      {
        "id": 2,
        "repository_id": 1,
        "number": "1.1",
        "state": "created",
        "started_at": null,
        "finished_at": null,
        "config": {
          "notifications": {
            "webhooks": ["http://evome.fr/notifications", "http://example.com/"]
          }
        },
        "status": null,
        "log": "Blahblah blah\\n[helm-lint] <start>{ \"charts\": { \"redis\": {\"name\": \"redis\", \"status\": 50, \"failed_criteria\": [\"lint-type-1\", \"lint-type-2\"]}, \"alpine\": {\"name\": \"alpine\", \"status\": 100} } }[helm-lint] <end>",
        "result": null,
        "parent_id": 1,
        "commit": "62aae5f70ceee39123ef",
        "branch": "master",
        "message": "the commit message",
        "committed_at": "2011-11-11T11: 11: 11Z",
        "committer_name": "Sven Fuchs",
        "committer_email": "svenfuchs@artweb-design.de",
        "author_name": "Sven Fuchs",
        "author_email": "svenfuchs@artweb-design.de",
        "compare_url": "https://github.com/svenfuchs/minimal/compare/master...develop"
      }
    ]
  }
}
`

	response := httptest.NewRecorder()
	request, err := http.NewRequest("POST", "/charts/", strings.NewReader(travisBody))
	if err != nil {
		t.Fatal("Creating 'POST /charts/' request failed!")
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "travis-token")

	updateRouter.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatal(response.Code)
	}
}
