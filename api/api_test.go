package api_test

import (
	"fmt"
	"github.com/alexanderpyoung/bbgUser/api"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader
	usersUrl string
)

func init() {
	server = httptest.NewServer(api.Handlers())
	usersUrl = fmt.Sprintf("%s/users", server.URL)
}

func TestCreateUserIncomplete(t *testing.T) {
	userJson := `{"username": "newuser", "password": "password", "email": "xandy@email.com"}`

	reader = strings.NewReader(userJson)

	request, err := http.NewRequest("POST", usersUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 400 {
		t.Errorf("Bad Request expected: %d", res.StatusCode)
	}
}

func TestCreateUserComplete(t *testing.T) {
	userJson := `{"username": "newuser", "password": "password", "email": "xandy@email.com", "friend": "Xandy"}`

	reader = strings.NewReader(userJson)

	request, err := http.NewRequest("POST", usersUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
