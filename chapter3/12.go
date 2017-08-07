// http.Client構造体を使用して、クッキーの送受信を行う
package main

import (
  "log"
  "net/http"
  "net/http/cookiejar"
  "net/http/httputil"
)

func main() {
  // クッキーを保存するcookiejar(クッキーの瓶)のインスタンスを作成
  jar, err := cookiejar.New(nil)
  if err != nil {
    panic(err)
  }
  // クッキーを保存可能なhttp.Clientインスタンスを作成
  client := http.Client{
    Jar: jar,
  }
  // クッキーは初回アクセスでクッキーを受信し
  // 2回目以降のアクセスでクッキーをサーバに対して送信する仕組みなので2回アクセス
  for i := 0; i < 2; i++ {
    // `http.Get()` の代わりに、作成したクライアントのGet()をつかってアクセス
    resp, err := client.Get("http://localhost:18888/cookie")
    if err != nil {
      panic(err)
    }
    dump, err := httputil.DumpResponse(resp, true)
    if err != nil {
      panic(err)
    }
    log.Println(string(dump))
  }
}

// net/http/cookiejarは組み込みライブラリとして実装されているクッキー機能の一つだが
// オンメモリでしか管理されていないので永続化はされない
// 永続的に利用するにはサーバパーティを利用するか `https://github.com/juju/persistent-cookiejar` 
