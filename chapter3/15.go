// $ curl -x http://localhost:18888 http://github.com
// golangでプロキシを設定
package main

import (
  "log"
  "net/http"
  "net/http/httputil"
  "net/url"
)

func main() {
  proxyUrl, err := url.Parse("http://localhost:18888")
  if err != nil {
    panic(err)
  }
  client := http.Client{
    Transport: &http.Transport{
      Proxy: http.ProxyURL(proxyUrl),
    },
  }
  resp, err := client.Get("http://github.com")
  if err != nil {
    panic(err)
  }
  dump, err := httputil.DumpResponse(resp, true)
  if err != nil {
    panic(err)
  }
  log.Println(string(dump))
}

// client.Get()先は外部サイトになっていますが、プロキシの向き先はローカルのテストサーバです
// このコードを実行すると、外部向けには直接リクエストは飛ばずに、ローカルのサーバが一旦リクエストを受けます
// ですが、このローカルサーバが直接レスポンスを返しているので、このコードではgithub.comへのアクセスは発生しません

// クッキーと同様にコードを追加してhttpパッケージのDefaultTransportに設定すると、http.Get()等のグローバルな関数でもプロキシが使われます

// http.DefaultTransport = &http.Transport{
//   Proxy: http.ProxyURL(proxyUrl),
// }
