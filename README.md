# bazel

## 参考
- https://christina04.hatenablog.com/entry/using-bazel-to-build-go

## インストール
- install bazel
    - `brew install bazel`
- install gazell
    - `go install github.com/bazelbuild/bazel-gazelle/cmd/gazelle@latest`

## 環境
あらかじめ以下のコードを用意しておく。
```
.
├── README.md
├── cmd
│   └── main.go
├── go.mod
└── go.sum
```

### cmd/main.go
実行するとxidを生成し標準出力する。

## BAZELの実行準備
- touch WORKSPACE
    - SET UPに書いてあるコードをpaste
    - ![gazelle](https://github.com/bazelbuild/bazel-gazelle#running-gazelle-with-go)
- touch BUILD.bazel
```
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/garireo2549/bazel
gazelle(name = "gazelle")
```

## ビルドファイルの生成
- `bazel run //:gazelle`
    - 各ディレクトリに`BUILD.bazel`が作成される
- `bazel run //:gazelle -- update-repos -from_file=go.mod`
    - 依存関係をWORKSPACEに自動生成

## 実行
- `bazel run //cmd`
```
bazel?[main]: % bazel run  //cmd
INFO: Analyzed target //cmd:cmd (25 packages loaded, 372 targets configured).
INFO: Found 1 target...
Target //cmd:cmd up-to-date:
  bazel-bin/cmd/cmd_/cmd
INFO: Elapsed time: 0.286s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Running command line: bazel-bin/cmd/cmd_/cmd
cgnojre49b3jidljdomg
```
