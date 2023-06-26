
## Original from

https://go.dev/doc/tutorial/workspaces

## go1.18からのgo workspaceが解決する課題

go.mod に書き込んでいた `replace => ` の指定を毎回Git Commit時に削除するような手間を無くさせる。

→ 逆にいうと、本質的なことは go.mod の `replace => ` で同じことができるということ。

## やっていること

[go.work](./go.work) に記載先のモジュール(go.mod)名が第一優先の参照先になる。

## ここでの例

[./hello/go.mod](./hello/go.mod)が参照しているのは、`require golang.org/x/example`であり、通常はそのGoレポジトリ先のパッケージを見に行くが、go.workで `use ./example`とあり、[./example/go.mod](./example/go.mod) に`module golang.org/x/example`があるので、helloが参照するのは、./exampleにあるパッケージになる。