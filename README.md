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
    - [gazelle](https://github.com/bazelbuild/bazel-gazelle#running-gazelle-with-go)
```
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6b65cb7917b4d1709f9410ffe00ecf3e160edf674b78c54a894471320862184f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.39.0/rules_go-v0.39.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.39.0/rules_go-v0.39.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098274d14a1d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
    ],
)


load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

go_rules_dependencies()

go_register_toolchains(version = "1.19.5")

gazelle_dependencies()
```
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

## test
- build
    - `bazel run //:gazelle`
    - `bazel run //:gazelle -- update-repos -from_file=go.mod`
- test
    - `bazel test //...`
```
bazel?[main]: % bazel test //...
INFO: Analyzed 3 targets (0 packages loaded, 0 targets configured).
INFO: Found 2 targets and 1 test target...
INFO: Elapsed time: 0.190s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
//cmd:cmd_test                                                  (cached) PASSED in 0.4s

Executed 0 out of 1 test: 1 test passes.
There were tests whose specified size is too big. Use the --test_verbose_timeout_warnings command line option to see which ones these are.
```
