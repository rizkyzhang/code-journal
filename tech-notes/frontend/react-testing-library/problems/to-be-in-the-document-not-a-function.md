## Problem

toBeInTheDocument() not a function

## Solution

toBeInTheDocument is not part of RTL. You need to install jest-dom to enable it.

And then import it in your test files by: `import '@testing-library/jest-dom'`

## Reference

https://stackoverflow.com/questions/56547215/react-testing-library-why-is-tobeinthedocument-not-a-function
