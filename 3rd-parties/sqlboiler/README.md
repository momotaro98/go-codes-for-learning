# volatiletech/sqlboilerk

参考にした記事

* https://qiita.com/uhey22e/items/640a4ae861d123b15b53


# Step

0. スキーマ作成済みのDBを立ち上げる
0. `sqlboiler.toml` ファイルを修正して置く
0. Run `go get -u -t github.com/volatiletech/sqlboiler` and `go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql`
0. `sqlboiler --wipe mysql`

上記により自動生成のコード達が現れる。
