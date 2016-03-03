// file based stack of strings
package fstack

import (
	"io/ioutil"
	"os"
	"strings"
)

// Stack provides access to basic file-based stack.
type Stack struct {
	fname string
	lines []string
}

// Create new Stack from given file.
func New(filename string) (*Stack, error) {
	f, err := os.Open(filename)
	if os.IsNotExist(err) {
		return &Stack{fname: filename}, nil
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	st := &Stack{fname: filename, lines: strings.Split(string(data), "\n")}
	if len(st.Top()) == 0 {
		st.Pop()
	}
	return st, nil
}

// Check if stack is empty
func (st *Stack) Empty() bool {
	return len(st.lines) == 0
}

// Appends value to the stack.
// It will be written at end of file when `Sync()` will be called, in case if you want to edit it also manually.
func (st *Stack) Push(s string) {
	st.lines = append(st.lines, s)
}

// Read value from top of stack, `""` if `Empty()`
func (st *Stack) Top() string {
	if len(st.lines) == 0 {
		return ""
	}
	return st.lines[len(st.lines)-1]
}

// Pop value from the stack, does nothing if already `Empty()`
func (st *Stack) Pop() {
	if len(st.lines) <= 1 {
		st.lines = nil
		return
	}
	st.lines = st.lines[:len(st.lines)-1]
}

// Pop all values from the stack
func (st *Stack) Clear() {
	st.lines = nil
}

// Get size of the stack
func (st *Stack) Size() int {
	return len(st.lines)
}

// Write stack content back to original file.
// This method will replace original content.
// Returned value comes from underlying `file.Close()`.
// _perm_ is used for file creation if file doesn't exist before.
func (st *Stack) Sync(perm os.FileMode) error {
	return ioutil.WriteFile(st.fname, []byte(strings.Join(st.lines, "\n")+"\n"), perm)
}
