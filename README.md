# Connect

> **Note**  
> Segment has paused maintenance on this project, but may return it to an active status in the future. Issues and pull requests from external contributors are not being considered, although internal contributions may appear from time to time. The project remains available under its open source license for anyone to use.

Segment Connect is a framework to develop custom integrations that listen on Webhooks.


Creating a new integration is as simple as having a struct that satisfies this interface:
```go
type Integration interface {
	Init() error
	Process(r io.ReadCloser) error
}
```

Each request is routed through the `Process(...)` function.

## Example Integration

```go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type LogIntegration struct{}

func (l *LogIntegration) Init() error { return nil }

func (l *LogIntegration) Process(r io.ReadCloser) error {
	defer r.Close() // Closes the reader when process returns

	// Read the entire body in memory
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Printf("-> Received body: %v\n", string(b))
}

func main() {
	connect.Run(new(LogIntegration))
}
```
