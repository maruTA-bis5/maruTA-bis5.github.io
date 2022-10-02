---
title: 自作Mikutterプラグインまとめ（2014年3月版）
date: 2014-03-30T15:25:21+0900
draft: false
categories: 
    - プログラミング
tags:
    - mikutter
    - mikutter plugin
    - mikutter command
aliases:
    - /2014/03/30/9.html
---

![thumb.jpg](thumb.jpg "なんかTwitterの鳥っぽい感じがした by [Macomb Paynes](http://www.flickr.com/photos/24730945@N03/8691494438/) - [CC BY-NC-SA 2.0](http://creativecommons.org/licenses/by-nc-sa/2.0/)")

私はRubyで書かれたクロスプラットフォームのTwitterクライアント「mikutter」の愛用者です。  
mikutterの特徴の1つとして、簡単なRubyスクリプトでいくらでも機能を拡張できるプラグイン機構があります。  
私も幾つか自作プラグインを公開しているので、まとめておきます。  
mikutterプラグイン、mikutterコマンドなど興味がある方は、<a href="http://mikutter.hachune.net">http://mikutter.hachune.net</a>や<a href="http://yuzuki.hachune.net/wiki">http://yuzuki.hachune.net/wiki</a>をご覧ください。

### mikutter-focus-to-hometimeline
ホームタイムラインタブにフォーカスを移動するmikutterコマンドを提供するプラグインです。  
mikutterコアで投稿ボックスにフォーカスを移動するコマンドは提供されていますが、マウスを極力使いたくない人種（含自分）としてはホームTLにキーボードだけで移動したいのです。  
<a href="http://github.com/bis5/mikutter-focus-to-hometimeline">http://github.com/bis5/mikutter-focus-to-hometimeline</a>

### mikutter_growl_gntp
クロスプラットフォームの通知アプリケーションである「Growl」に対して新着通知を行うプラグインです。  
Linuxではlibnotifyやnotify-osdというプログラムで通知できる環境がほとんどですが、RubyからWindowsAPIを通じてバルーン通知やトースト通知を行うのは非常に面倒です。Growlを使えば、少ないコード量で簡単に通知機能を実装できます。  
Growlはリモートホストからの通知にも対応しており、このプラグインはリモートホスト上のGrowlに対しても通知を飛ばせるようになっています。mikutterプラグインでバッチ処理を行い、管理用ホストに完了通知を行うような、本来のプラグインホストとして乗りようにおいて有用な気がします（でっちあげ）。  
<a href="http://github.com/bis5/mikutter_growl_gntp">http://github.com/bis5/mikutter_growl_gntp</a>

### mikutter-googl
goo.glで短縮されたURLを展開して表示するプラグインです。  
展開するのにわざわざAPIキーを使っているので、今後利用者が増えてきたらAPI Limitに引っかかるかもしれません。そうなったらそうなったでそのときに考えたいとは思います。  
<a href="http://github.com/bis5/mikutter-googl">http://github.com/bis5/mikutter-googl</a>

### my_wishlist
Amazon.co.jpの欲しいものリストのURLを投稿ボックスに挿入します。  
設定画面で挿入する欲しいものリストのIDを設定する必要があります。  
欲しいものリストを公開してテロの被害者になりたい方におすすめのプラグインです。ちなみに私はまだ被害を受けたことはありません。<a href="http://github.com/bis5/my_wishlist">http://github.com/bis5/my_wishlist</a>

### mikutter_kokoro_no_koe
。○（こんな感じに心の声っぽくテキストを整形するプラグインです。）  
。○（それ以上でもそれ以下でもありません。）  
。○（形が気に入らない方はフォークして自分の心の声を表すプラグインを作ってください）  
<a href="http://github.com/bis5/mikutter_kokoro_no_koe"> http://github.com/bis5/mikutter_kokoro_no_koe</a>
