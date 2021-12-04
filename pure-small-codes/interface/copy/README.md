
# how to learn

```
GOOS=darwin GOARCH=amd64 go build -o copy -gcflags '-N -l' # コンパイル最適化しないようにする
go tool objdump copy # コンパイルされたバイナリをアセンブリへ解析する (逆アセンブリ)
```

# output

f()

```
  main.go:8             0x108b681               48c744242001000000      MOVQ $0x1, 0x20(SP)
  main.go:9             0x108b68a               440f117c2430            MOVUPS X15, 0x30(SP)
  main.go:10            0x108b690               488d542420              LEAQ 0x20(SP), DX
  main.go:10            0x108b695               4889542428              MOVQ DX, 0x28(SP)
  main.go:10            0x108b69a               488d355f490000          LEAQ runtime.types+17824(SB), SI
  main.go:10            0x108b6a1               4889742430              MOVQ SI, 0x30(SP)
  main.go:10            0x108b6a6               4889542438              MOVQ DX, 0x38(SP)
```

g()

```
  main.go:16            0x108b721               48c744242001000000      MOVQ $0x1, 0x20(SP)
  main.go:17            0x108b72a               440f117c2438            MOVUPS X15, 0x38(SP)
  main.go:18            0x108b730               488b542420              MOVQ 0x20(SP), DX
  main.go:18            0x108b735               4889542428              MOVQ DX, 0x28(SP)
  main.go:18            0x108b73a               488d153f740000          LEAQ runtime.types+28960(SB), DX
  main.go:18            0x108b741               4889542438              MOVQ DX, 0x38(SP)
  main.go:18            0x108b746               488d542428              LEAQ 0x28(SP), DX
  main.go:18            0x108b74b               4889542440              MOVQ DX, 0x40(SP)
```

# learning point

* f()の方ではinterface{}型としてのaの値ポインタはiと同じになる
* g()の方ではaの値ポインタは新しく作られる

# わからない点

f()とg()での逆アセンブリの結果を見てもコピーの様子が違う意味が読み取れない

→ アセンブリの読み方をもっと学ぶ必要あり (2021年12月4日時点)