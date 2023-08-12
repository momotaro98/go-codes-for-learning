# underlying type と core type

core typeがunderlying typeより抽象的な概念。

## underlying type

Go公式Doc: https://go.dev/ref/spec#Underlying_types

<冒頭日本語訳>  
それぞれの型Tにはunderlying typeがある： Tが事前に宣言されたboolean型、numeric型、string型のいずれか、または型リテラルである場合、対応するunderlying typeはTそのものである。それ以外の場合、Tのunderlying typeは、Tが宣言で参照する型のunderlying typeです。  
</冒頭日本語訳>  

参考: https://zenn.dev/nobishii/articles/type_param_intro#underlying-type

> Go言語の全ての型は、それに対応する"underlying type"という型を持っています。
> 1つの型に対して、対応するunderlying typeは必ず1つだけ存在します。underlying typeを持たない型や、underlying typeを2つ以上持つ型は存在しません。

## core type

Go公式Doc: https://go.dev/ref/spec#Core_types

<冒頭日本語訳>  
Tがインタフェース型**では無い**場合、Tはそれぞれcore typeを持ち、これはTのunderlying typeと同じである。  
Tがインターフェース型の場合、がcore typeを持つのは、以下のいずれかの条件を満たす場合である。

* Tの型集合に含まれるすべての型のunderlying typeであるUが1つ存在する。
* Tの型集合は、identical element typeであるEを持つchannel型のみを含み、すべての方向性channelは同じ方向を持つ。

**上記2つ以外のインターフェース型はcore typeを持たない。**
</冒頭日本語訳>  

### core type がどういう意味を持つか

参考: https://zenn.dev/nobishii/articles/type_param_intro#core-type%E3%81%AE%E7%99%BB%E5%A0%B4%E5%A0%B4%E9%9D%A2

