# Go Utils

A collection of Go utility functions for common operations.

## Installation

```bash
go get github.com/hjunior29/go-utils
```

## Usage

```go
import "github.com/hjunior29/go-utils/pkg/utils"

// String utilities
reversed := utils.Reverse("hello")     // "olleh"
capitalized := utils.Capitalize("go")  // "Go"

// Slice utilities
found := utils.Contains([]string{"a", "b"}, "a") // true

// Math utilities
max := utils.Max(5, 10) // 10
min := utils.Min(5, 10) // 5
```

## License

[MIT](LICENSE)
