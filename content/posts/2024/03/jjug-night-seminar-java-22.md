---
title: "JJUG ナイトセミナー 「Java 22 リリース記念イベント」 参加メモ #jjug"
date: 2024-03-31T12:30:00+0900
draft: false
description: JJUGナイトセミナー「Java22リリース記念イベント」参加時のメモ
categories:
    - 勉強会
tags:
    - JJUG
---

2024-03-27に開催されたJJUGナイトセミナー「[Java 22リリース記念イベント](vhttps://jjug.doorkeeper.jp/events/169972)」に参加したメモ。

{{< linkCard "https://jjug.doorkeeper.jp/events/169972/" "【オフライン】JJUGナイトセミナー「Java 22 リリース記念イベント」3/27(水) 開催" >}}

多分資料と合わせて見るかJEPと付き合わせないと自分でも思い出せなくなる気がするけど、まだ資料は公開されていないかな? 時期に公開されると思うのでそのときにリンク追記します。

セッション登壇された久保田([@sugarlife](https://twitter.com/sugarlife/with_replies))さん、運営スタッフの皆さん、ありがとうございました。

## JEPs in JDK 22
- JEP 423: REgion Pinning for G1  
JNIを使っているとGCが無効化される(GC Locker)  
→critical objectがいるRegionをPinningする(Evacuationさせない)ことでGCを待たせることなく処理できるように  
※Pinning自体は既存の仕組み  
OOME発生可能性の緩和とパフォーマンス改善(全体のstall回避)の効果  
Shenandoahは(おそらく当初から)対応済み、ZGCはこれから。

- JEP 447: Statements before super(...) (preview)  
親クラスのコンストラクタ(`super()`)呼び出し前に処理がかける  
例: コンストラクタ引数の値を検証してから呼びたい (従来: コンストラクタに渡してから検証するしか無かった。あるいは、検証メソッドをstaticで書く必要があった)  
例: 値の準備(コンストラクタ引数で受け取った値を加工して親コンストラクタ・別のコンストラクタに渡す)  
例: 親コンストラクタの引数を共有したい(ローカル変数を経由できる)  
制約: コンストラクト前のインスタンスはアクセス不可(コンストラクタを呼ぶ前にフィールド・メソッドを触る、等はエラー)  
エラーははっきりとした表現  
例: `cannot reference this before supertype constructor has been called`

- JEP 456: Unnamed Variables & Patterns  
不要な変数やパターンを`_`(Java9で禁止された)でかけるようにし、不要なことを明確にする等  
JDK 21から内容の変更なし

- JEP 4463: Implicitly Declared Classes and Instance Main Methods (Second Preview)  
"おまじない"の削減  
`void main()` or `void main(String[] args)`であれば、アクセス修飾子等は何でも良い  
クラス宣言も不要

- JEP 454: Foreign Function & Memory API  
Java -> Native: downcall  
Native -> Java: upcall

- JEP 460: Vector API (Seventh Incubator)  
SIMD演算(ベクトル演算)のサポート

- JEP 457: Class-File API (Preview)
ASMをJDK内部でも使っており、クラスファイルフォーマットが変わるとASMの対応に時間がかかる  
→クラスファイルAPIを提供することで、JDK側で互換性を担保したい  
フィールド、メソッド、属性、バイトコード命令などのクラスファイルの実体を不変のオブジェクトとして扱うAPI(`java.lang.classfile`)を提供  
(感想)ASMと同じ書き方のはずだが、こちらのほうが読みやすく見える気がする  
(感想 with lambda)良き

- JEP 459: Launch Multi-File Source-Code Programs  
ソースコードから直接実行する Source-file mode が複数ファイルに対応

- JEP 459: String Templates (Second Preview)  
Advanced usage: https://github.com/bitterfox/json-string-template

- JEP 462: Structured Concurrency (Second Preview)  
内容としてはJDK21から変更なし

- JEP 464: Scoped Values (Second Preview)  
JDK 21から変更なし  
ThreadLocalよりも扱いやすく  
Structured Concurrencyとの組み合わせでも有用(ThreadLocalだと子スレッドにコピーされてしみ合う)

- JEP 461: Stream Gatherers (Preview)  
中間操作のユーザ定義をサポートする  
`stream.gather(..).gather(..).gather(..).collect(..)`  
`java.util.stream.Collector`に対する`java.util.stream.Gatherer`  
built-in gatherersもある  
fold: Stream#reduceと違って順序保証があり、BiFunctionを受け取るので入力と出力の方が違っても良い  
scan: インクリメンタルな累積  
mapConcurrent: mapをVirtual Threadを使ってconcurrentに  
windowFixed: 順番に詰める  
windowSliding: 一つずつずらしながら末尾まで詰める  
Gathererで自前実装する: Initializer, Integrator, Combiner(sequencialでない場合のみ), Finisher  
    Integrator: downstream: 出力のイメージでいいかな  
    FInisher: 最後に余ったものをdownstreamに突っ込んで終わる

## JDK 23
- JEP 455: Primitive Types in Patterns, instanceof, and switch (Preview)
