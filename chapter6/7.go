// OSに依存せずに証明書を読み込んでHTTPS通信を行うクライアントコード

package main

import (
  "crypto/tls"
  "crypto/x509"
  "io/ioutil"
  "log"
  "net/http"
  "net/http/httputil"
)

func main() {
  // 証明書をよみこむ
  cert, err := ioutio.ReadFile("ca.crt")
  if err != nil {
    panic(err)
  }
  certPool := x509.NewCertPool()
  certPool.AppendCertsFromPEM(cert)
  tlsConfig := &tls.Config{
    RootCAs: certPool,
  }
  tlsConfig.BuildNameToCertificate()

  // クライアントを作成
  client := &http.Client{
    Transport: &http.Transport{
      TLSClientConfig: tlsConfig,
    },
  }

  // 通信を行う
  resp, err := client.Get("https://localhost:18443")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  dump, err := httputil.DumpResponse(resp, true)
  if err != nil {
    panic(err)
  }
  log.Println(string(dump))
}
