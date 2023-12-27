---
tags:
  - golang-util
---
```go
package rupiah

import (
	"fmt"

	"github.com/bojanz/currency"
)

func FormatRupiah(value uint64) (string, error) {
	amount, err := currency.NewAmount(fmt.Sprint(value), "IDR")
	if err != nil {
		return "", err
	}
	locale := currency.NewLocale("id")
	formatter := currency.NewFormatter(locale)
	formatter.MaxDigits = 0

	return formatter.Format(amount), nil
}
```
