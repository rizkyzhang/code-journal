- Writing a test is just like writing a function, with a few rules:
  - It needs to be in a file with a name like `xxx_test.go`
  - The test function must start with the word `Test`
  - The test function takes one argument only t `*testing.T`

---

- `t.Helper()` is needed to tell the test suite that a function is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper.

---

- TDD cycle:
  - Write a test
  - Make the compiler pass
  - Run the test, see that it fails and check the error message is meaningful
  - Write enough code to make the test pass
  - Refactor
