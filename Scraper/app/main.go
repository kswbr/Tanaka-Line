package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
)

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  url := "https://kushi-tanaka.com/news/"  // アクセスするURL

  resp, err := http.Get(url)      // GETリクエストでアクセスする
  if err != nil {                 // err ってのはエラーの時にエラーの内容が入ってくる
      panic(err)                  // panicは処理を中断してエラーの中身を出力する
  }
  defer resp.Body.Close()         // 関数が終了するとクローズする

  byteArray, err := ioutil.ReadAll(resp.Body) // 帰ってきたレスポンスの中身を取り出す
  if err != nil {
      panic(err)
  }
  fmt.Fprintf(w, string(byteArray))
}
