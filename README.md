# go-utils

A small, tested Go utility library. Each utility lives in its own source file and has automated tests.

## Requirements

Go 1.24 or newer.

## Build and test

```sh
go build ./...
go test -count=1 ./...
```

## Usage

```go
import "github.com/hjunior29/go-utils/pkg/utils"

reversed := utils.ReverseString("hello") // "olleh"
```

## Initial utilities

- `ReverseString` reverses Unicode code points.
- `WordCount` counts whitespace-separated words.
- `ClampInt` clamps an integer to inclusive bounds and rejects reversed bounds.

## Adding utilities

Use English for code, comments, documentation, and tests. Add one utility per file with meaningful tests covering normal, empty, boundary, and invalid input where applicable. Do not add dependencies without review.

Run the complete build and test suite before submitting changes. Keep public exports synchronized when adding modules.

## License

MIT
