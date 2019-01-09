Test package Runner

Run console command with arguments and environment variables.

## How to use

```go
package main

import (
	"github.com/vovanada/runner"
	"log"
)

func main() {
	r, err := runner.RunProcess("ping", []string{"127.0.0.1", "-c", "4000"}, []string{})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("PID - %v", r.PID())
}

```