# メモ
## 標準入力の受け取り
- `fmt.scan`はバッファリングしないので、大量に入力するときは遅い。TLEになる場合は`bufio.scanner`を使ったライブラリを自作して使うのがいい。
    - https://zenn.dev/hsaki/books/golang-io-package/viewer/intro