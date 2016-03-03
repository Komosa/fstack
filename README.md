## fstack - stack in file
This lib provides file based stack of strings.

Status: only one unit test at the moment, but API will remain stable.

List of operations:
- `New(filename string) (*Stack, error)`, open given file and read stack from it;
- methods with `*Stack` receiver:
    - check if stack is `Empty() bool`;
    - `Push(s string)` appends value to the stack (at end of file, if you want to edit it also manually);
    - `Top() string` read value from top of stack, `""` if `Empty()`;
    - `Pop()` pop value from the stack, does nothing if already `Empty()`;
    - `Clear()` pop all values;
    - `Size() int`;
    - `Sync(perm os.FileMode) error` write stack content back to underlying file; returned value comes from underlying `file.Close()`; use _perm_ for file creation if file doesn't exist before.

Note: This is not mechanism for synchronization.
Eg. `Sync()` will always overwrite file.

### When to use?
Use it if you are writing app which needs *small and simple* state.
Also when state sync speed is not bottleneck.

My example is tool for reporting my activity at job to jira.
Stack is for putting actual task and stack, starting work on another task (eg, incoming bug), and going back to the task from stack when finished.
I don't want to run program all the time (so, I don't use in memory storage).
I also don't want to use anything _heavy_, as database.
I need sometimes grab what I'm doing and add some persistence to data, and obviously using file comes to my mind.
