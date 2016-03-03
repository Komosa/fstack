package fstack_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Komosa/fstack"
)

func TestNonMod(t *testing.T) {
	f, err := ioutil.TempFile("", "fstack_test_file")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	data := []byte(`bottom
middle
top
`)
	_, err = f.Write(data)
	f.Close()
	fatalMaybe(err, t)

	st, err := fstack.New(f.Name())
	fatalMaybe(err, t)

	if s := st.Top(); s != "top" {
		t.Errorf("top of stack should be `top`, got: %v", s)
	}
	if s := st.Size(); s != 3 {
		t.Errorf("stack should have exactly 3 elements, got: %v", s)
	}
	if st.Empty() {
		t.Error("stack should be not empty, got: empty=true")
	}

	err = st.Sync(0) // perms doesn't matter, file already exists
	fatalMaybe(err, t)

	got, err := ioutil.ReadFile(f.Name())
	fatalMaybe(err, t)

	if !bytes.Equal(data, got) {
		t.Errorf("file content differs, exp: %#+v, got: %#+v", data, got)
	}
}

func fatalMaybe(err error, tb testing.TB) {
	if err != nil {
		tb.Fatal(err)
	}
}
