package sync

import "testing"

type fsTester struct {
	copysrc string
	copydst string
}

func (fs *fsTester) Copy(src, dst string) error {
	fs.copysrc = src
	fs.copydst = dst
	return nil
}
func (fs *fsTester) GetNewerFile(f1, f2 string) (string, error) {
	switch {
	case f1 == "Test2" && f2 == "Test3":
		return "", ErrFilesHaveSameAge
	case f2 == "Test1":
		return f2, nil
	default:
		return f1, nil
	}
}

func (fs *fsTester) SkipFile(s string) bool {
	if s == "TestSkip" {
		return true
	}
	return false
}

var tests = []struct {
	file1     string
	file2     string
	expectsrc string
	expectdst string
	expecterr error
}{
	{"Test1", "Test2", "Test1", "Test2", nil},
	{"Test2", "Test1", "Test1", "Test2", nil},
	{"Test2", "Test3", "", "", ErrFilesHaveSameAge},
	{"TestSkip", "Test4", "", "", nil},
}

func TestSync(t *testing.T) {
	for _, test := range tests {
		fs := &fsTester{}
		err := Sync(fs, test.file1, test.file2)
		if err != test.expecterr {
			t.Fatalf("Wrong Error. Got:%v Expect: %v\n", err, test.expecterr)
		}
		if fs.copysrc != test.expectsrc && fs.copydst != test.expectdst {
			t.Fatalf("Copy is not called correct!\n Got: src:%v dst:%v\nWant: src:%v dst:%v",
				fs.copysrc,
				fs.copydst,
				test.expectsrc,
				test.expectdst)
		}

	}
}
