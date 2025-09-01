package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MohummedSoliman/booking/pkg/models"
)

type postData struct {
	key   string
	value string
}

var tests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"generals", "/generals", "GET", http.StatusOK},
	{"majors", "/majors", "GET", http.StatusOK},
	{"contact", "/contact", "GET", http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range tests {
		res, err := ts.Client().Get(ts.URL + test.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if res.StatusCode != test.expectedStatusCode {
			t.Errorf("for %s, expected %d but got %d", test.name, test.expectedStatusCode, res.StatusCode)
		}
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	req := httptest.NewRequest("GET", "/make-reservation", nil)
	ctx := getContext(req)
	req = req.WithContext(ctx)

	resRecorder := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.MakeReservation)
	handler.ServeHTTP(resRecorder, req)

	if resRecorder.Code != http.StatusOK {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", resRecorder.Code, http.StatusOK)
	}

	// Test case when reservation is not in the session
	req = httptest.NewRequest("GET", "/make-reservation", nil)
	ctx = getContext(req)
	req = req.WithContext(ctx)

	resRecorder = httptest.NewRecorder()
	handler.ServeHTTP(resRecorder, req)

	if resRecorder.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", resRecorder.Code, http.StatusTemporaryRedirect)
	}

	// Test with non-existent room
	req = httptest.NewRequest("GET", "/make-reservation", nil)
	ctx = getContext(req)
	req = req.WithContext(ctx)

	resRecorder = httptest.NewRecorder()
	reservation.RoomID = 3
	session.Put(ctx, "reservation", reservation)
	handler.ServeHTTP(resRecorder, req)

	if resRecorder.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", resRecorder.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepository_PostReservation(t *testing.T) {
	reqBody := "start_date=2030-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2030-01-03")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=mohamed")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=ibrahim")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=mo@mail.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=01010101223")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	req := httptest.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx := getContext(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resRecoder := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(resRecoder, req)

	if resRecoder.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code: got %d, wanted %d", resRecoder.Code, http.StatusSeeOther)
	}

	// Test for messing form body
	req = httptest.NewRequest("POST", "/make-reservation", nil)
	ctx = getContext(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "x-www-form-urlencoded")

	resRecoder = httptest.NewRecorder()
	handler.ServeHTTP(resRecoder, req)

	if resRecoder.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for messing reqeust body: got %d, wanted %d", resRecoder.Code, http.StatusTemporaryRedirect)
	}

	// Test for invalid start date
	reqBody = "start_date=2030-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2030-01-03")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=mohamed")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=ibrahim")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=mo@mail.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=01010101223")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	req = httptest.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getContext(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "x-www-form-urlencoded")

	resRecoder = httptest.NewRecorder()
	handler.ServeHTTP(resRecoder, req)
	if resRecoder.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid start date: got %d, wanted %d", resRecoder.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepository_AvailabilityJSON(t *testing.T) {
	reqBody := "start=2030-01-01&end=2030-01-02&room_id=1"
	req := httptest.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))
	ctx := getContext(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Repo.AvailabilityJSON)
	handler.ServeHTTP(resRecorder, req)

	var js jsonResponse
	err := json.Unmarshal(resRecorder.Body.Bytes(), &js)
	if err != nil {
		t.Error("failed to parse response json")
	}
}

func getContext(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
