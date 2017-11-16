package downloader

import "testing"

func TestGet(t *testing.T) {
	if _, err := Get("baidu", "http://www.hao123.com"); err != nil {
		t.Fatal(err)
	}
}

func TestPostJson(t *testing.T) {

}
