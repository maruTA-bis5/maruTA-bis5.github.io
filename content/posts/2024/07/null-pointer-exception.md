---
title: "NullPointerExceptionの原因調査方法 #Java"
date: 2024-07-06T16:40:00+0900
draft: false
description: NullPointerExceptionの原因となったnullの探し方を解説します
categories:
    - プログラミング
tags:
    - Java
    - 初心者向け
---

Javaを使って開発している人が良く遭遇する例外の一つである`java.lang.NullPointerException`(NPE)。経験が浅い人にとってはどう対処すれば良いか分からないという声をよく聞くので、原因となったnullの探し方を軽く紹介します。

## `NullPointerException`が発生する原因
そもそもの`NullPointerException`が発生する原因について軽く説明しておきます。

`NullPointerException`は、"値がnullであるオブジェクトに対するメソッド呼び出しやフィールドアクセス"を行おうとした際に発生する例外です。  
コード上は、_"nullに評価されるなにか" + ピリオド(`.`) + フィールド名 or メソッド呼び出し_ という形式であれば、`NullPointerException`の発生箇所だと判断することが可能です。

(他にも`throw null;`もNPEを発生させますが、あまり多く見かけないと思うのでこの記事では扱いません。)


## Java 17より前のバージョンを使っている場合
例えば、次のようなコードがあったとします。
```java
class Value {
    private String field;

    String getField() {
        return field;
    }
}

// どこかのメソッド内
Value v;

if (v.getField().isEmpty()) { // HERE
```
`HERE`と書いてあるif文でNullPointerExceptionが発生していれば、`v`もしくは`v.getField()`のどちらかが`null`と評価されていることが推測できます。

別の例を見てみましょう。
```java
void someMethod(int arg) {
    ...
}

Integer value = null;
someMethod(value);
```
このコードは`someMethod(value);`でNullPointerExceptionが発生します。確かに`value`は`null`に評価されるのですが、先程示した形式を満たしていません。

`Integer`という型は、プリミティブな`int`のwrapper classです。wrapper classの変数にはプリミティブ型と違って`null`を代入可能ですね。  
一方、`someMethod`の引数はプリミティブな`int`と定義されていますから、`Integer`を引数で渡すと、自動アンボクシング(Autounboxing)によって`Integer`から`int`への変換が行われます。  
`Integer`の自動アンボクシングでは[`Integer#intValue()`](https://docs.oracle.com/javase/jp/21/docs/api/java.base/java/lang/Integer.html#intValue())が使用される([JLS 5.1.8](https://docs.oracle.com/javase/specs/jls/se21/html/jls-5.html#jls-5.1.8))ため、  
"nullに評価されるオブジェクトに対するメソッド呼び出し"となることで`NullPointerException`が発生するのです。

## Java 17以降を使っている場合
[JEP 358: Helpful NullPointerExceptions](https://openjdk.org/jeps/358)が導入され、何がnullでどの操作に失敗したのかを例外のメッセージで説明してくれます。

```
jshell> void print(Object arg) {
   ...>     System.out.println(arg.toString());
   ...> }
|  次を作成しました: メソッド print(Object)

jshell> print(null)
|  例外java.lang.NullPointerException: Cannot invoke "Object.toString()" because "<parameter1>" is null
|        at print (#5:2)
|        at (#6:1)
```

※正確にはJava 14で導入された機能ですが、業務で非LTSかつメンテナンスが終わっているJava 14,15,16を使っている人はさすがにいないですよね?
