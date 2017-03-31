package expand

import "testing"

func TestExpand(t *testing.T) {
	currnetURL := "http://baidu.com/a/b/c.html"
	pendingTest := []string{
		"http://baidu.com/result",
		"//baidu.com/result",
		"/result/ab/bb/cb",
		"result.html",
	}

	for _, item := range pendingTest {
		t.Log(Expand(currnetURL, item))
	}

	t.Log("\n\n----------------------\n\n")
	currnetURL = "http://baidu.com/a/b/c/"
	pendingTest = []string{
		"http://baidu.com/result",
		"//baidu.com/result",
		"/result/ab/bb/cb",
		"result.html",
	}

	for _, item := range pendingTest {
		t.Log(Expand(currnetURL, item))
	}
}
