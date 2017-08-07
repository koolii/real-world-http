// $ curl -X DELETE http://localhost:18888

package main

import (
  "log"
  "net/http"
  "net/http/httputil"
)

func main() {
  client := &http.Client{}
  request, err := http.NewRequest("DELETE", "http://localhost:18888", nil)
  if err != nil {
    panic(err)
  }
  resp, err := client.Do(request)
  if err != nil {
    panic(err)
  }
  dump, err := httputil.DumpResponse(resp, true)
  if err != nil {
    panic(err)
  }
  log.Println(string(dump))
}

// PostFormで送信する形式に変換するには、 `url.Values` を `io.Reader` インターフェース準拠のオブジェクトに変換する
// import (
//   "strings"
//   "net/url"
// )
//
// values := url.Values{"test":{"value"}}
// reader := strings.NewReader(values.Encode())

// ヘッダーの送信
// $ curl -H "Content-Type=image/jpeg" -d "@image.jpeg" http://localhost:18888
// request.Header.Add("Content-Type", "image/jpeg")

// BASIC認証
// request.SetBasicAuth("user-name", "pass")
// クッキーを手動でひとつ足す
// request.AddCookie(&http.Cookie({Name:"test",Value:"value"})
