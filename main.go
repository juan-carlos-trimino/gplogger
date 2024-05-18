package logger

import (
  "encoding/json"
  "fmt"
  "time"
)

type logEntry struct {
  DateTime string "json:\"date_time\""
  Level string `json:"level"`
  CorrelationId string `json:"correlation_id"`
  Message string `json:"msg"`
}

func LogInfo(msg, cid string) {
  printMessage(msg, cid, "INFO")
}

func LogWarning(msg, cid string) {
  printMessage(msg, cid, "WARN")
}

func LogError(msg, cid string) {
  printMessage(msg, cid, "ERROR")
}

func printMessage(msg, cid, level string) {
  le := logEntry{
    //time.Now() returns the current local time; using the current time in UTC.
    DateTime: time.Now().UTC().Format(time.RFC3339Nano),
    CorrelationId: cid,
    Level: level,
    Message: msg,
  }
  json, err := json.Marshal(le)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(json))
}
