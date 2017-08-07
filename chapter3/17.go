// $ curl file://main.go
// golangでローカルファイルにアクセスするfileスキーマを有効化する
package main

import (
  "log"
  "net/http"
  "net/http/httputil"
)

func main() {
  transport := &http.Transport{}
  transport.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))
  client := http.Client{
    Transport: transport,
  }
  resp, err := client.Get("file://./main.go")
  if err != nil {
    panic(err)
  }
  dump, err := httputil.DumpResponse(resp, true)
  if err != nil {
    panic(err)
  }
  log.Println(string(dump))
}

// 通信バックエンドのhttp.Transportには、これ以外のスキーマ用のトランスポートを追加する
// RegisterProtocolメソッドがある。このメソッドに登録できる、ファイルアクセス用バックエンド
// http.NewFileTransport()もあります。これらを使うことで、ローカルファイルアクセスが出来るようになる。
// ローカルファイルの内容がレスポンスボディに格納されて帰ってきます。
