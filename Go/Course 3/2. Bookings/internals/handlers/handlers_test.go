package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type param struct {
	key   string
	value string
}

type routeTest struct {
	name               string
	path               string
	method             string
	params             []param
	expectedStatusCode int
}

var routeTests = []routeTest{
	{
		name:               "home",
		path:               "/",
		method:             http.MethodGet,
		params:             []param{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "about",
		path:               "/about",
		method:             http.MethodGet,
		params:             []param{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "contact",
		path:               "/contact",
		method:             http.MethodGet,
		params:             []param{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "search-availability",
		path:               "/search-availability",
		method:             http.MethodGet,
		params:             []param{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:   "search-availability",
		path:   "/search-availability",
		method: http.MethodPost,
		params: []param{
			{key: "start", value: "2020-01-01"},
			{key: "end", value: "2020-01-05"},
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:   "search-availability-json",
		path:   "/search-availability-json",
		method: http.MethodPost,
		params: []param{
			{key: "start", value: "2020-01-01"},
			{key: "end", value: "2020-01-05"},
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "make-reservation",
		path:               "/make-reservation",
		method:             http.MethodGet,
		params:             []param{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:   "make-reservation",
		path:   "/make-reservation",
		method: http.MethodPost,
		params: []param{
			{key: "first_name", value: "Vinícius"},
			{key: "last_name", value: "Souza"},
			{key: "email", value: "viniciusps01@gmail.com"},
			{key: "phone", value: "63992710348"},
		},
		expectedStatusCode: http.StatusOK,
	},

	{
		name:               "reservation-summary",
		path:               "/reservation-summary",
		method:             http.MethodGet,
		params:             []param{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "majors-suite",
		path:               "/majors-suite",
		method:             http.MethodGet,
		params:             []param{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "generals-quarters",
		path:               "/generals-quarters",
		method:             http.MethodGet,
		params:             []param{},
		expectedStatusCode: http.StatusOK,
	},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewServer(routes)

	defer ts.Close()

	for _, test := range routeTests {
		var err error
		var res *http.Response

		switch test.method {
		case http.MethodGet:
			res, err = ts.Client().Get(ts.URL + test.path)
		case http.MethodPost:
			values := url.Values{}

			for _, item := range test.params {
				values.Add(item.key, item.value)
			}

			url := ts.URL + test.path

			res, err = ts.Client().PostForm(url, values)

		default:
			t.Error("No test written for testing requests of type", test.method)
			return
		}

		if err != nil {
			t.Log(err)
			log.Fatal(err)
		}

		if test.expectedStatusCode != res.StatusCode {
			t.Errorf(
				"%v:%v expects %d, but got %d",
				test.method,
				test.name,
				test.expectedStatusCode,
				res.StatusCode,
			)
		}
	}
}
