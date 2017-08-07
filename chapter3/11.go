// $ curl -F "name=Micael Jackson" -F "thumbnail=@photo.png" http://localhost:18888
// golangでマルチパートフォームをPOST送信
package main

import (
  "bytes"
  "io"
  "log"
  "mime/multipart"
  "net/http"
  "os"
  "net/textproto"
)

func main() {
  // マルチパート部を組み立てた後のバイト列を格納するバッファを宣言
  var buffer bytes.Buffer
  // マルチパートを組み立てるライターを作る
  writer := multipart.NewWriter(&buffer)
  // ファイル以外のフィールドは `WiteField()` を使って登録
  writer.WriteField("name", "Michael Jackson")


  // 10で使用していたコード
  // ここから `io.Copy()` までがファイルを読み込む操作
  // まず個別のファイル書き込みの `io.Writer` を作る
  // fileWriter, err := writer.CreateFormFile("thumbnail", "phone.png")
  // if err != nil {
  //   panic(err)
  // }

  // 11で使用しているコード
  part := make(textproto.MIMEHeader)
  part.Set("Content-Type", "image/jpeg")
  part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.png"`)
  fileWriter, err := writer.CreatePart(part)
  if err != nil {
    panic(err)
  }
  // ファイルオープン
  readFile, err := os.Open("photo.png")
  if err != nil {
    panic(err)
  }
  defer readFile.Close()
  // `io.Copy()` を使って、ファイルの全コンテンツを、ファイル書き込み用の `io.Writer` にコピー
  io.Copy(fileWriter, readFile)

  // 最後にマルチパートの `io.Writer` をクロ＝逗子、バッファに全てを書き込む
  writer.Close()

  resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
  if err != nil {
    panic(err)
  }
  log.Println("status:", resp.Status)
}

// だがこのままだとMIMEタイプが事実上void型友言えるapplication/octet-stream型になってしまう
// 次でMIMEタイプを詰める
