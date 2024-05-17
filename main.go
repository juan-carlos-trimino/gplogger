package logger

import (
  "fmt"
  "time"
)

func LogInfo(message string) {
  /***
  time.Now() returns the current local time; using the current time in UTC.
  ***/
  fmt.Printf("%s INFO - %v", time.Now().UTC().Format(time.RFC3339Nano), message)
}

func LogWarning(message string) {
  fmt.Printf("%s WARN - %v", time.Now().UTC().Format(time.RFC3339Nano), message)
}

func LogError(message string) {
  fmt.Printf("%s ERROR - %v", time.Now().UTC().Format(time.RFC3339Nano), message)
}
