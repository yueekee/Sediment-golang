package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// StubPlayerScore 继承了 PlayerStore接口
type StubPlayerScore struct {
	scores map[string]int
}

func (s *StubPlayerScore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestServer(t *testing.T) {
	store := StubPlayerScore{
		map[string]int{
		"Pepper": 20,
		"Floyd": 10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T){
		// request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		request := newGetScoreRequest("Pepper")
		respone := httptest.NewRecorder()

		server.ServeHTTP(respone, request)

		got := respone.Body.String()
		want := "20"

		// if got != want {
		// 	t.Errorf("got '%s', want '%s'", got, want)
		// }
		assertResponseBody(t, got, want)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		// request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
	
		server.ServeHTTP(response, request)
	
		got := response.Body.String()
		want := "10"
	
		// if got != want {
		// 	t.Errorf("got '%s', want '%s'", got, want)
		// }
		assertResponseBody(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		//if got != want {
		//	t.Errorf("got status %d want %d", got, want)
		//}
		assertStatus(t, got, want)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerScore{
		map[string]int{},
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/" + name, nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
        t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
    }
}