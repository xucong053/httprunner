package examples

import (
	"testing"

	"github.com/httprunner/httpboomer"
)

func TestCaseBasicRequest(t *testing.T) {
	testcase := &httpboomer.TestCase{
		Config: httpboomer.TConfig{
			Name:    "request methods testcase in hardcode",
			BaseURL: "https://postman-echo.com",
			Verify:  false,
		},
		TestSteps: []httpboomer.IStep{
			httpboomer.Step("get with params").
				GET("/get").
				WithParams(map[string]interface{}{"foo1": "bar1", "foo2": "bar2"}).
				WithHeaders(map[string]string{"User-Agent": "HttpBoomer"}).
				Validate().
				AssertEqual("status_code", 200, "check status code").
				AssertEqual("headers.Connection", "keep-alive", "check header Connection").
				AssertEqual("headers.\"Content-Type\"", "application/json; charset=utf-8", "check header Content-Type").
				AssertEqual("body.args.foo1", "bar1", "check args foo1").
				AssertEqual("body.args.foo2", "bar2", "check args foo2"),
			httpboomer.Step("post raw text").
				POST("/post").
				WithHeaders(map[string]string{"User-Agent": "HttpBoomer", "Content-Type": "text/plain"}).
				WithData("This is expected to be sent back as part of response body.").
				Validate().
				AssertEqual("status_code", 200, "check status code").
				AssertEqual("body.data", "This is expected to be sent back as part of response body.", "check data"),
			httpboomer.Step("post form data").
				POST("/post").
				WithHeaders(map[string]string{"User-Agent": "HttpBoomer", "Content-Type": "application/x-www-form-urlencoded"}).
				WithParams(map[string]interface{}{"foo1": "bar1", "foo2": "bar2"}).
				Validate().
				AssertEqual("status_code", 200, "check status code").
				AssertEqual("body.form.foo1", "bar1", "check form foo1").
				AssertEqual("body.form.foo2", "bar2", "check form foo2"),
			httpboomer.Step("post json data").
				POST("/post").
				WithHeaders(map[string]string{"User-Agent": "HttpBoomer"}).
				WithJSON(map[string]interface{}{"foo1": "bar1", "foo2": "bar2"}).
				Validate().
				AssertEqual("status_code", 200, "check status code").
				AssertEqual("body.json.foo1", "bar1", "check json foo1").
				AssertEqual("body.json.foo2", "bar2", "check json foo2"),
			httpboomer.Step("put request").
				PUT("/put").
				WithHeaders(map[string]string{"User-Agent": "HttpBoomer", "Content-Type": "text/plain"}).
				WithData("This is expected to be sent back as part of response body.").
				Validate().
				AssertEqual("status_code", 200, "check status code").
				AssertEqual("body.data", "This is expected to be sent back as part of response body.", "check data"),
		},
	}

	err := httpboomer.Run(t, testcase)
	if err != nil {
		t.Fatalf("run testcase error: %v", err)
	}
}
