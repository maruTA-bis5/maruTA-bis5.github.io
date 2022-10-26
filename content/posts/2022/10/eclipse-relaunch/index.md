---
title: "[Eclipse] `java.net.BindException: Address already in use`を回避する"
date: 2022-10-27T08:00:00+0900
draft: false
description: "Eclipseでサーバーアプリケーションを開発する際、2重に起動して`java.net.BindException: Address already in use`に遭遇する、という経験をした人も多いのではないでしょうか。
Eclipse 4.6(Neon)以降では、2重起動を回避する手段が用意されています。"
categories:
    - プログラミング
tags: 
    - Eclipse
---

Eclipseでサーバーアプリケーションを開発する際、2重に起動して`java.net.BindException: Address already in use`に遭遇する、という経験をした人も多いのではないでしょうか。  
Eclipse 4.6(Neon)以降では、2重起動を回避する手段が用意されています。

メニューからWindow > Preferencesを選択して設定ダイアログを開き、Run/Debug > Launchingの一番下にある`Terminate and relaunch while launching`をONに設定することで、すでに起動しているプロセスがあれば終了してから起動するようになります。  
![preference](./preference.png)

括弧書きにもあるとおり、Shiftキーを押しながら起動しても同じ挙動となりますが、毎度Shiftキーを押して起動するのは面倒なので、予め設定しておいたほうが便利でしょう。

#### 参考
{{< linkCard "https://www.eclipse.org/eclipse/news/4.6/platform.php#terminate-relaunch-history" "Eclipse Project Neon - New and Noteworthy" >}}
