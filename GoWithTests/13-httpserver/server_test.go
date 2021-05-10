package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// StubPlayerScore 继承了 PlayerStore接口
type StubPlayerScore struct {
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerScore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerScore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestServer(t *testing.T) {
	store := StubPlayerScore{
		map[string]int{
		"Pepper": 20,
		"Floyd": 10,
		},
		nil,
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
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
	
		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got '%s' want '%s'", store.winCalls[0], player)
		}
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

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/players/" + name, nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
        t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
    }
}

/* 
集成测试对较大型的测试很有用，但你必须牢记：
集成测试更难编写
测试失败时，可能很难知道原因（通常它是集成测试组件中的错误），因此可能更难修复
有时运行较慢（因为它们通常与“真实”组件一起使用，比如数据库）
*/

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
    store := InMemoryPlayerStore{}
    server := PlayerServer{&store}
    player := "Pepper"

    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

    response := httptest.NewRecorder()
    server.ServeHTTP(response, newGetScoreRequest(player))
    assertStatus(t, response.Code, http.StatusOK)

    assertResponseBody(t, response.Body.String(), "3")
}