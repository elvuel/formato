# formato

Simple format converter for json, yaml, toml.

Moved from `confandrea` utilities.

## Installation

```bash
$ go get github.com/elvuel/formato
```

## Usage

```go
//...
import (
    "github.com/elvuel/formato"
)

func main() {
    // Y2J converts YAML to JSON
    formato.Y2J(input)
    // J2Y converts JSON to YAML
    formato.J2Y(input)
    // T2J converts TOML to JSON
    formato.T2J(input)
    // J2T converts JSON to TOML
    formato.J2T(input)
}
//...

```

## License

MIT
