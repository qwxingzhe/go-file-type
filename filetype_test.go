package go_file_type

import (
	"io/ioutil"
	"os"
	"testing"
)

type PathTest struct {
	path, result string
}

var loacltests = []PathTest{
	{"./testfile/0010.gif", "gif"},
	{"./testfile/0011.gif", "gif"},
	{"./testfile/002.png", "png"},
	{"./testfile/0030.jpg", "jpg"},
	//{"./testfile/0040.webp","webp"},
	//{"./testfile/0041.webp","webp"},
	//{"./testfile/0042.webp","webp"},
}

func TestUnit_GetFileTypeByByte(t *testing.T) {
	for _, test := range loacltests {
		f, _ := os.Open(test.path)
		fSrc, _ := ioutil.ReadAll(f)
		if s := GetFileTypeByByte(fSrc[:10]); s != test.result {
			t.Errorf("GetFileTypeByByte(%q) = %q, want %q", test.path, s, test.result)
		}
	}

}

func TestUnit_GetFileTypeByLocalPath(t *testing.T) {
	for _, test := range loacltests {
		if s := GetFileTypeByLocalPath(test.path); s != test.result {
			t.Errorf("GetFileTypeByPath(%q) = %q, want %q", test.path, s, test.result)
		}
	}
}

var urltests = []PathTest{
	{"http://img.gif.cn/temp_makegif/20210806/1628230128817044.gif", "gif"},
}

func TestUnit_GetFileTypeByUrl(t *testing.T) {
	for _, test := range urltests {
		if s := GetFileTypeByUrl(test.path); s != test.result {
			t.Errorf("GetFileTypeByUrl(%q) = %q, want %q", test.path, s, test.result)
		}
	}
}
