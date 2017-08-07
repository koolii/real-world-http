// golangで任意の文字列をPOST送信
package main

import (
  "log"
  "net/http"
  "strings"
)

func main() {
  reader := strings.NewReader("テキスト")
  resp, err := http.Post("http://localhost:18888", "text/plain", reader)
  if err != nil {
    panic(err)
  }
  log.Println("status:", resp.Status)
}
