---
title: "JJUG CCC 2023 Spring 参加メモ #jjug_ccc"
date: 
draft: true
description: 
categories:
    - 勉強会
tags:
    - JJUG
    - JJUG CCC
---

2023/06/04 西新宿で開催された JJUG CCC 2023 Spring に現地参加したのでそのメモを放流しておきます。

## 全体を通して
前回2022 Fallでは試験的に一部オフライン開催だったのが、今回はオンライン・オフラインのハイブリッド開催だったということで、
朝から移動して参加するのってこんな感じだったな、と少し懐かしく思いながらの参加でした。

セッション予約はすべての時間を埋めてしまうクセがあり、あまり休まる時間がないので終わる頃にはだいぶヘトヘトなんですよね。
今回はパスしましたが、次回はスタッフとして参加するのもアリかもな、と思っています。
(会社が許せばスポンサー参加したいと思いつつ。これまで採用には全く関わっていませんけど、中途採用のきっかけとしてCCCのスポンサーになってみたらどうか、話してみようかな)

## 金融系子会社でレガシーシステムしか作ったことないけど、モダン開発に挑戦してみた
9:30 コンファレンスB

生命保険会社


生命保険システムの特徴: ライフサイクルが長い(保証期間)m高い品質(顧客のお金を扱う)、巨大

新技術を学ぶ際のハードル
技術の壁: 巨大なシステム&大人数で開発→品質を担保するために、技術的な部分が社内独自フレームワークに隠蔽されていて見えない
自動テストは共通機能のみ、リファクタも少ない
→作られたものをずっと同じ・決まったやり方で保守するため、改善されることもなく、モダンな技術との乖離が大きいい

文化の壁: コード変更に承認が必要になるので不要な変更を避ける
システムが巨大->影響範囲が大きく改善しじづらい
ライフサイクルが長い右安定志向が強い
→課題を感じても承認が必要で、新しいことに挑戦しづらい環境

モダン開発文化: 新しいものを学ぶ、常に変化する、新しい状況に対応する とは逆の環境

POCでの取り組み
フロントエンド: React + バックエンド SpringBoot + DDD, RESTful API

技術の壁への取り組み: Qiitaなどの情報だけでなく、原点となるドキュメントを当たる
DDD: エリックエヴァンス本から読んだら積んだ
→基礎知識の習得と合わせて実践(繰り返して経験を蓄積する) いきなり大規模なものを作るのではなく、シンプルな仕様から少しずつ追加していく
手続き型→RESTful & DDDは考え方の改革に近い→何回もAPI設計
モデルズ(クラス図、シーケンス図)を活用して全体の流れを抑えるのが良いと思う
テストコードを書くことになれるのも必要

文化の壁への取り組み: 新シ事を学ぶ文化の情勢 原点に当たる、もくもく会、資格取得)
新しいことにチャレンジ(繰り返し、専任で自発的に参加)

もとのスキルセットからは、3年位でやっと普通に開発できるようになった
マイクロサービスが銀の弾丸のように思えていたが、基幹システムでは適合しないケースもある

### QA
Q: 別の業界だがレガシーな組織にいて、モダンに作り変えるのは渦香椎と感じている。チームの年齢層?
A: 20代前半から後半。古い会社にありがちだが、上に話が伝わらないこともあり、あねージャー層を追加してもらって、いまでは新入社員〜3,40題がメイン

Q: 過去にレガシーな環境にいたことがある。新しいモダンな文化を取り入れるのは大変なこと。会社に掛け合うのに時間がかかったと思うが、会社として取り組む形にどうやって行ったのか
A: "2025年の可"がスタートで、会社として上が危機感を抱いていた。PoCが流行っているらしい、としてプロジェクトが作られたことが大きい。
並走していて忘れされれていきがちだったが、そうではないだろ、ということを言い続けていた


Q: 新しい技術を導入することを、上がGoをだしたしきい値(雰囲気)が知りたい
A: 意外と巻き込まれて消えていきそうな感じだったが、心の中に残っていたのか、継続的にアプローチし続けていたら倒れていた。アピールする人がいれば倒れてくれるのではないか、という印象

Q: 大きいシステム、ウォーターフォールのイメージだが、アジャイルは導入したのか
A: ハイブリッドな感じ。上の人がプログラムを作れない、新しいやり方での要件定義二日と来ていなかったので、要件定義はウォーターフォールで。設計ルールを決めないような要件定義をやってもらう感じに。いまは完全なアジャイルではないが、イテレーション(2weeks)でタスクを組んで消化する、という流れ


## サーバレスJavaの今 ~SNapStartとWeb Adapterを寄せて~
10:00 コンファレンスB
@yasai_ls
サーバーレスJavaを取り巻く最新トレンドの紹介
SpringBootとWuarkusを利用

### そもそもサーバレスってなんだっけ
4つの特徴: 
インフラのプロビジョニング・管理が不要: パッチ適用も含めてインスラ事業者が担当
自動スケーリング: 
利用した分だkの支払い: アイドル状態のリソースはなく、利用した時間分だけ支払う→コスト効率が良い
高い可用性と安全性: (AWS)複数のアベイラビリティゾーンに公開される

代表機能: AWS Lambda
Event DrivenなFunction as as Service
FIrecracker MicroVM上に実行環境、ランタイム、実行コードがデプロイされる(本セッションの対象は`実行コード`


### Lambda SnapStart
Re:Invent 2022で発表された機能

従来のLambdaのライフサイクル
initフェーズ: JVM実行環境の初期化・JARファイルのロード(コールドスタート)
→肥大化したJar(Fat,Uber)の良い込でコールドスタートにかかる時間が増える)

Lambda SnapStartによる変化
有効化してフェプロ医する→バージョン発行をトリガーにInitフェーズが実行されて、その後スナップショットが取得される
→スナップショットから起動することで、コールドスタートにかかる時間を軽減できる(コールドスタートフェーズが"スナップショットの復元フェーズ"に変わる)
→コールドスタートと比較して起動速度が早くなる

Lambda SnapStartの有効化: 
AWS SAM: AutoPublishAlias, SNapStartを指定
CDKなど、他のツールでもかんたんに指定できる

→SnapStartで起動時間が1/4まで短縮できた

SnapStartのトレードオフ
現時点で未対応: Provisioned COncurrency, ARM64, Amazon EFS, 512MBを超えるエフェメラルストレージ
トレードオフ: Lambdaが14日間呼び出されないと関数がInactiveに移行→起動しようとすると軌道に失敗→Snapshotが生成されて使えるようになる
→暖気生んても考慮したほうがいいかも

考慮点
一意性の確保: UUID生成など一位となる情報をスナップショットに含めると一意性を確保できない(SecureRandomを使うべき)→初期化フェーズで一位情報を生成ないこと(After Restoreで一位情報を生成するとお酔い))
ネットワーク確率: 不ナップショット復元後のNW確率の担保が必要(外部API，DB接続など)　SDKに関わる部分は自動的に再接続されるが、それ以外について(Runtime Hookを活用)

ランタイムフックの活用: スナップショットを作成する前、または復元したあとでフックを挟むことができる(`beforeCheckpoint()`, `afterRestore()`)

ベストプラクティス: すなっぷしょっとさくせいじいにおおくのくらすをしょきかする
→起動時のレイテンシーの現任となるクラスを事前ロードするとか、ダミー呼び出しを使用してクラスをロードしておいてスナップショットに含めておく
Quarkusでpreload classのメタデータ生成機能があり便利


### AWS Lambda Web Adapter
Spring BootをLambdaで起動する

AWS SAM, Serverless Framework: Deployment Framework .... かんたんにサーバレスアプリケーションをデプロイできるOSSという位置づけ

サーバレスジャヴァでフレームワークを使うメリット
使い慣れたFWを利用することで開発者体験の工場(特にローカル開発。Lambda関数だとLocalstack使ったりしないといけない)
移植性の向上(サーバーレス ,=> コンテナオーケストレーション 切り替えるときにコードを変えなくて良い)

Web Adapter: Webアプリケーションを(ほぼそのまま) AWS Lambdaで動かせる(Javaいがいでも)
HTTPリクエストをWeb Adapter経由でLambda関数に渡せる

SpringBoot appおWeb Adapterを使えるようにする
デプロイはAWS SAM CLIで3ステップでOK
DockerfileはAWS Labda Web Adapterの拡張機能をコピーする一文を加えれば後は普通に
Quakusでも同じ

コールドスタート時間: 6~7秒かかった
→Graal VM Native Image化→1秒未満で起動できた

Native Imageの考慮点: 
ピーク時のスループット:
JITではないので、ピーク時スループットはJITより遅い
→FaaSだとあまり考慮しなくても良いと思う
トレースエージェント: リフレクションの動的要素に対応できない
→トレース・エージェントツールを死闘して、到達可能性メタデータを作っておく必要がある

Web Adapterのトレードオフ:
API Gateway, Lambda Function URL, Application Load Balancerのみサポート
EventBridge, S3 event, SQSメッセージなどによるインテグレーションは、従来のEventRequestクラスを利用したハンドリングを推奨

Web Adapter以外の選択肢
フレームワーク共通furem-muwa-kukyo: Serverless Java Container
フレームワーク統合: Spring Cloud Function, Quarkus


### サーバレスJavaの技術スタック選定
ライブラリ・実装・ビルド・デプロイオプション・デプロイ、でそれぞれ選択肢がある
Maven/Gradle
Pure Java/Framework
Jar/Native
SnapStart/WebAdapter.none
Zip/Container

排他要素
NativeImage: SnapStart未対応。ライフらいくるフックが利用できない(使うメリットがない)
SnapStart: Container Imageに非対応

フレームワーク利用における統合の選択肢
- イベントハンドラークラスを利用(基本): イベントハンドラーと見つけ統合している→コンテナ移行のコストが高い
- フレームワークによる(HTTP)エンドポイントを利用→コンテナ移行が極めて用意


AWS Lambda Web Adapter Non Native Imageで高速起動
→Provisioned Concurrency + Application Auto Scalingを活用して暖機運転 + 自動スケール員・アウト

### なぜイベント駆動アーキテクチャなのか
コールドスタートで起動した関数の実行環境は再利用される(ウォームスタート)
→起動時間が早ければウォームスタートの回数も増える
なるべく早く起動し、なるべく早く終了すれば、スケールアウト街を減らしてウォームスタートで実行できる

### 適応度関数の計測
LambdaによるHTTP APIで何を死守したいのか: 2秒以内でレスポンスが帰ってくること
→レイテンシが2秒以内に収まっているか、適応度関数を用いて計測することが重要
例: CloudWatch Synthetics, K6

### べbりなツールセット
AWS Lambda Powertools for Java
X-Rayによるトレース、構造化ログ、カスタムメトリクスの非同期作成を用意にする
襟: コールドスタートをアノテーションとしてキャプチャ、Embedded Metric Format尼したがってメトリクスを標準出力するとカスタムメトリクスを非同期に作成

AWS Lambda Power Tuning
コストの最適化のため、最適なメモリサイズを可視化できる

### まとめ
サーバレスJavaの実装方法を理解し最適な選択を。


### WA
Q: Provisioned Concurrencyと(SnapStart)の違い
A: PC: コールドスタート後のウォームスタートの状態でいくつか確保しておいて、すぐに実行できるようにする機能。ランタイムがJavaでなくても利用可能。SnapStartはJava専用の機能。PCは常に暖機運転しているので、その分コストがかかる


## SpringBootでメッセージキュー＆非同期処理を使ったノウハウ紹介
11:00 コンファレンスC
リクルート

HRシステムのアーキテクト

アーキテクトとしての働き方
システムの安定稼働、開発を維持するために構造的な問題を解決する
技術の声優だけでなく、現場の課題認識と解決能力も求められる(シワゆるシステム設計だけにとどまらない)
→ブレの少ないアーキテクチャを示すことで生産性を上げたり、品質を向上させる)

西村さん: 思想を話す担当
山田さん: 技術を話す担当

過去のデブサミ,JJUGでの発表はレガシーシステムの発表が多かったが、今回はクラウド・新規サービスの話
(安定稼働。改善が重要視されていて、稼働から10年・20年立っていたシステム)

最近のHR事業の傾向
- クラウドで構築された新規システムが増えてきている(継続性・アジリティを考慮した経営判断があり、技術スタックが進化している)


### 非同期処理導入の背景
(前提) 今回話す「非同期処理」: Javaのスレッドを用いた非同期ではなく、メッセージキューを用いた非同期処理

メッセージキューがあると何が嬉しい? -> 非同期で(好きなタイミングで)処理に関われる
メッセージキューにためておけば、別々のタイミングでデプロイすることもできる

メッセージキューのSLAが高くないといけない
→マネージドなメッセージキューサービスがあるので、高いSLAのメッセージキューは比較的容易に用意できる

例: オンプレで、DBに登録されたデータを一括で外部連携したり、多くの処理を一度に行わないといけない(DB登録、メール送信、添付ファイル登録、外部API呼び出し、など)処理
→エラー処理を色々考えないといけないので、結果としてコードや仕様などの複雑度を上げていく傾向

クラウドをベースにすることで選択肢が広がった
オンプレと同じやり方で構築することもできるが、より最適な方法を模索する中で注目したのが非同期処理

メリット
- 時間のかかる処理を非同期化→Webの応答速度を一定にする、リトライも用意
- 定期バッチでまとめて処理していたもの(cron設定やエラー処理が必要)を随時個別に処理(件数が多ければ水平スケールできるのでパフォーマンスもよい)

適切にメッセージキューを使うことで柔軟かつ堅牢なシステム構築ができる
キューはマネージドなので、運用コストも上がらない

メリットだけを伝えても導入は難しい→いかにハードルを下げて実装してもらうか、が重要

Webサービスの代表的な技術スタック: ECS+Fargate, NextJS + SpringBoot 非同期処理はSpringBootで実装
※シンプルさを重視した技術選定を心がけている

- プロデューサーはSpringBootのみ: SQSへのメッセージ登録に成約を設ける
- コンシューマーもSpringBootで構築: Producer, Consumerを同じ技術スタックで実装できるようにした
- 設計指針もWebAPI(SpringBoot)に合わせた(Controller/Usecase/Service/Repository)
    - Message APIという名前にした(処理の起点がMessage Queueなので)→処理の起点が異なるだけの同じAPIとして認識できるように


### Spring Boot + SQSでの構築
SQS: Simple Queue Service

コンシューマーの話がメイン(プロデューサーはキューイングするだけ)

受信するコード
```java
var request = ReceiveMessageRequest.builder()
    .queueUrl(queueUrl)
    .maxNumberOfMessages(5)
    .build();
return sqsClient.receiveMessage(receiveMessageRequest).messages(); 
```
定期的に処理を実行して、キューにメッセージが届いているか確認する必要がある
→自動的にポーリングして、メッセージが届いていたら動か串済にできると良い

選択肢
1. AWS Lambda イベントソースマッピング: メッセージがキューに入っていたらLambda Functionを呼び出してくれる
2. Spring Cloud AWSのアノテーション駆動形リスナーエンドポイント: ライブラリがポーリングして、メッセージがあれば`@SqsListener`をつけたメソッドを呼んでくれる

→WebAPIと合わせてSpringBootを利用したいため、Spring Cloud AWSを利用

※Spring Cloud 2.x系を前提。3.xでは微妙に書き方が違う
```java
@Component
public class SampleMessageController {
    @SqsListener(value = "sampleQueue",
        deletionPolicy = SqsMessageDeletionPolicy.ON_SUCCESS)
    public void receiveMessage(String message) {
        System.out.println(message);
    }
}
```

DeletionPolicy: SQSはメッセージを受信しただけではキューからは削除されず、明示的に削除する必要がある。そのタイミングの設定
→正常に処理ができなかったときにリトライ可能にするため
ALWAYS
NEVER
NO_REDRIVE (デフォルト): リドライブポリシーが定義されていない場合、メッセージを削除する
ON_SUCCESS

リドライブポリシー: メッセージが一定回数移譲受信された際に、一時的に別のキュー(デッドレターキュー)に移動させる設定
→メッセージそのものに問題があって残り続けているメッセージを"壊れたメッセージ"として退避させる場合に用いる

→デッドレターキューを設定した上でON_SUCCESSを使用している

その他Tips
SpringのAspectを死闘すると、メッセージ処理の前後の処理を実装dキル
アプリケーションの停止時は20秒(デフォルト)まで全SqsListener処理完了まで待つ(時間がかかる処理があるなら調子しておくと良し)
起動時にはコアプールがリスナー数x2, 最大プールがリスナー数x(10+1)確保される

Web/Message APIを両立するためのプロジェクト構成
- web-api (controller/usecase)
- message-api (controller/usecase)
- service (業務ロジック)
- repository

ローカルのテスト用意性を上げる: LocalStackを使ってAWS上のサービスを模擬的に構築できる(SQSを使った動作確認がローカルで)

### 導入してみてのノウハウ共有

扱いやすさの代償
かんたんに誰でもかけるようにした結果、非同期であることを無視した実装が報じた
例: システム感のデータ更新→参照、非同期による更新処理が完了する前に、WebAPIで同期的にデータを取得することが起きた
連携の方針が決まっていなかった: 連携部分はMessageAPIでやる方針を立てて、粗結合を目指していた
"API"の思想: レネ慶する側の実装者からすると、両方とも単に"API"だという認知になってしまった(単にAPIコールすればよい、という認識になった))
→対策
運用で検知: 設計レビューをする
パターンの提示: 別システムからのデータ更新が本当に必要か、見直す
非同期処理は扱う上で難しいので、実装がかんたんであっても、設計からうまく付き合う必要がある

監視・スケール
SQSごとにSpringBoot側でControllerを用意すると、1プロセス中に複数のSQSを購読する形になる
MessageAPI側のスループットを超えるメッセージが登録されたり、何らかの理由でMessageAPIが遅延したりすると、キューにメッセージが貯まる
→一定時間、しきい値を超えるメッセージが溜まっていたりする場合にアラートを出すと良い
→処理遅延を解消するには水平スケール(プロセスを増やす)が楽

段階的リリース
ビッグバンリリースよりも、小さく段階的なリリースがこ泊まれるのが近年のトレンド。非同期処理でもこの考え方を採用できるケースがあるのでは。
例: あるでーたが更新されたのを起点にPDFを作成してS3に格納する(格納は非同期で良い)
→データ登録処理側にキューへの登録を追加することで、こうしんAPIはファイル生成側の処理時間やエラーに影響されなくなる

以下の順番での段階リリースを実施
1. SQSにメッセージを送信する機能
2. 更新されたデータXを取得する機能: キューが正しく処理されることを確認する
3. でーたXをPDFに変換する機能: 処理時間やインフラへの負荷を確認する
4. PDFをS3に格納する機能
→小さいリリースでリスクを最小化


### QA
Q: 段階的リリース、最初のリリースではConsumeするところはかんたんに作ったのか、設定でクリアされるようにしたのか
A: 後者。SQSは最大14日しか保持されないので、キューに入ったことだけ確認して、後は自動で消えるのみ

Q: SQSを使ったが、"処理結果をユーザが知りたい"というのに、RDB似ステータステーブルと作る苦しい実装になってしまった。そのへんの工夫あるか
A: 今まさに作っている。基本的には同じように、登録されたジョブを管理するテーブルは必要かと思う。シンプルなステータスだけを管理するテーブルとして、シンプルな形。



## Head toward Java 20 and 21
13:00 コンファレンスA

Java 20: 2023/03/21 7 JEPs
Java 21: 2023/09/19予定 2023.06.08から安定化フェーズ(Rampdown)突入 16 JEPs(確定していないものも含む)

### Project Amber
開発生産性を高めるためのJava言語改善
- Record Patterns
    - `instanceof`パターンマッチングのRecord版
        ```java
        record Point(int x, int y) {}
        if (obj instanceof Point(int x, int y)) {
            System.out.println(x+y);
        }
        ```
    - Generic Record Pattern`record Box<T>(T t){}` にも対応
    - 21でStandard       `
- Pattern Matching for switch
    ```java
    return switch (obj) {
        case Integer i -> String.format(...);
        case Double d -> ...
    }
    ```
    - switch labelにenumを利用している際、実行時に網羅されていない場合の例外がMatchException
        - sealed classはIncompatibleClassChangeError
    - Generics Record Patternの方推論にも対応
    - Qualified enum定数をcase定数として利用可能(後述))
    - 21でStandard
    - 正常処理とthrow exceptionをswitch文で書いたり、条件処理を書いたり
        ```java
        case String s when (s,length() >= 1) -> s,
        case String s -> "",
        default -> ""
        ```
    - Objectは原理的に網羅するのは不可能なので、defaultがないとコンパイル時エラーになる
    - Qualified enum label
        ```java
        enum Flag{A,B}
        switch (f) {
            // これができるようになった
            case Flag.A:
            case Flag.B:
        }
        ```
- String Template
    - JDK 21で初回のPreview
    - MessageFormatとかString.formatとか文字列連結とかStringBuilderでやっていたやつを書きやすく
    - `java.lang.StringTemplate.Processor`を利用する
        ```java
        var x = 10.0, y = 20.5;
        String s -> STR."\{ x } plus \{ y } equals\{ x + y }";
        ```
    - `java.util.FormatProcessor`(`FMT`)で書式を指定可能
    - Processorを自前で作ることも可能
- Unnamed
    - Unnamed Patterns and Variables
        - JDK21で初回Preview
        - 使わない変数を使わないと明示することができる
            ```java
            if (obj instanceof Point(_, int y)) { System.out.println(y);}
            if (obj instanceof Point(int _, int y)) { System.out.println(y);}
            ```
    - Unnamed Classes and Instance Main Methods (Preview)j\
        - JDK 21で導入開始
        - クラス、パッケージ、モジュールと板概念なしに、段階的にプログラムを学べるように
        ```java
        void main() {
            System.out.println("Hello World!");
        }
        ```
        - Source code launcher(`java`コマンドで直接)で実行すればコンパイル不要

### Project Loom
- Virtual Threads
    - 従来のスレッド = OSのシステムスレッドのラッパー
        - システムコールが必要なのでコストが高く、コンテキストスイッチもあるし、OSスレッド数上限に達する場合もある
    - OSではなくJVMが管理する軽量スレッドの導入
    - JDK 21でStandard
    - JDK 21でThreadLocal変数がサポートされた(既存ライブラリの移行コストを低減)
    - `Thread.ofVirtual()`で作る (プラットフォームスレッドは`THread.ofPlatform()` or `new Thread`)
    - JFR, JVM TI, Java Debug InterfaceなどはCirtual Threadsに対応済みだが、一部は互換性の都合でPlatform threadsに限定される
        - Thread.setPriority, Thread.Daemon
        - Thread.getAllStackTraces
        - JVM TIのGetAllThreadsやGetAlLStackTracesはVirtual Threadを返さない
        - etc...
- Scoped Values
    - スレッドないとこスレッドの療法で普遍データを共有するプログラミングモデルの提供
    - VirtualThreadでThreadLocalより使いやすいものを作りたい
        - ThreadLocalの問題
            - いつでも変更できるのでどこで変更されるか分かりづらい
            - removeが呼ばれるまで保持されるのでメモリリークする可能性
            - おやすれっどのThreadLocalが小スレッドに継承されるので、多くのスレッドを使う場合にオーバーヘッド(特にメモリ使用量)が増大する可能視
    - JDK 21でIncubatorからPreview(java.lang)
        ```java
        private static final ScopedValue<User> USER = ScopedVlue.newInstance();
        ScopedValue.runShere(USER, new User("dule"), () -> doSomething());
        void doSomething() {
            var usr = USER.get();
            // doSomething内で最バインドできる
        }
        ```
    - `StructuredTaskScope`によって生成された小スレッドには継承される
- Structured Concurrency
    - 並列処理を楽に (タスクを有るコードブロックでサブタスクに分けて同時進行できるようにし、最終的に同じコードブロックに戻る)
    - JDK 21でIncubatorからPreviewに
    - `StructuredTaskScope::fork`が`Future`ではなく`Subtask`を返すように
    - 1つでも成功すれば他のサブタスクはシャットダウンする場合は、`StructuredTaskScope.ShutdownOnSuccess`を使う
        - すべて失敗した場合は`ExecutionException`
    - いずれかが例外を出したらシャットダウンする場合は、`StructuredTaskScope.ShutdownOnFailure`を使う

### Project Valhalla
「クラスのように記述でき、intのように動作する」新しい型、ValueObjectの導入
JDK 20/21での変更はなし

### Project Panama
- Vector API
    - SIMD演算のサポート
        - ベクトル演算を実行するためのJava API
    - JDK 21でSixth Incubator
    - Valhallaが提供されてからPreview APIとしたい、とのこと
- Foreign Function & Memory API
    - ヒプ外のメモリを直接扱いたい
    JDK 21で7回目
        - 今はPreview API(`java.lang.foreign`)

### No Project
- JEP 404: Generational Shenandoah (Experimental)
    - Shenandoahを世代別GCにする
        - 世代を分ける(Old/Young)ことでメモリの仕様効率を高めたい
        - Javaはもともと世代別GCが主だったので、こちらのほうが効率が良かそうであった
    - Defenerated GC, Full GCの機会を減らす(STWの原因)
- JEP 439 Generational ZGC
    - 最終的に、非世代別ZGCは削除予定
    - `-XX:+ZGenerational`
- JEP 431: Sequenced Collection
    - 例: 逆順で反復処理する際の書き方が統一される
    - SequencedCollection
        - 要素が定義された順序を持つコレクション
        - Dequeに`reversed()`が追加された形
    - SequencedSet
        - SortedSetのように相対比較で絵要素を配置するコレクションはaddFirst/LastはできないのでUnsupportedOperationException
    - SequencedMap
        - NavigableMapにreversed, sequenced..., put...が追加された形
        - putFirst, putLastは、LinkedHashMapではさいはいちするが、SortedMap等、ソートされている状態が重要なものはUnsupportedOperationException
    - 既存コードとのコンフリクトに注意
- JEP 451 Prepare to Disallow Dynamic Loading of Agents
    - 動的にjava agent, jvm tiをあタッチすることをデフォルトで禁止する
    - `-XX:+EnableDynamicAgentLoading`で警告を抑制できる
    - `-Djdk.instrument.traceUsage`で該当箇所を調査
- JEP 449: Depracate the Windows 32-bit x86 Port for Removal
    - Virtual Thread対応が無理
- JEP 452: Key Encapsulation Mechanism API

### Other noteworthy changes
- Tool
    - `jfr view`
    - (JDK 22) `jfr query "select * from GrabageCollection" file.jfr`
    - (未定) Asynchronous Stack Trace VM API (JEP 435)
- GC
    - 色々入った
        - メモリフットプリントの作ge、いらない起動オプションの削減
- Security
    - 弱いセキュリティの無効化、ファイルがないときに例外、等

### QA
Q: テンプレート構文で`\`がつくのは?
A: (Twitterで確認して回答)


## 複雑性に立ち向かうためのサーバーサイドコード分割
14:00 コンファレンスC
サイボウズ(kintoneチーム)
前田さｎ

### コード分割とは
背景
リリースして10年経過し、src/main/java以下のファイル創業数 >= 35万

肥大化・複雑化したことで開発効率が低下
同リファクタリングすればいいかわからない
影響範囲の確認が難しい
覚えおることが多く、新メンバーがキャッチアップしにくい


→コード分割という虜海
- 機能ごとにパッケージを分割
- 機能感で関わらないならコード感も関わらない、依存しない
→独立して理解可能なコードの塊を増やすことで改善を目指す

機能の特徴を捉えて、昨日の協会とコードの協会を合わせて確認する
"機能が違っていても偶然同じ実装"を1つにまとめるのは良くない
→コードは機能を実現するためにあるので、昨日の特徴を捉えて、偶然を配した意思決定につなげたい


kintone
業務システム(アプリ)を作成することができるツール。業務の数だけアプリを作る
アプリはデータベースのようなもの
データを蓄積する以外に、ワークフロー、アクセス権、通知等の設定がある


アプリ機能: アプリに関する機能。データ(レコード)の表示・編集をする(表示は1件・複数・グラフ)
アプリ設定機能; データベースの設定。項目(フィールド)のデータ型・内容を設定。アクセス権、通知、連携用APIトークンの設定もここ

アプリ機能(com.kintone.app)とアプリ設定機能(com.kintone.appsettings)が完全に独立する形を目指す
相互に依存しないように、ArchUnitで検証


依存の管理: ArchUnit

なぜ機能単位でのコード分割?
kinounitaiousitapakke-jiwakehaaruteidosareteitanode,saraniosisusumeteizonnnokanriwofukumeteokonauyounisita 
新規昨日の開発と並行して進められる(コードの中身はほとんど変わらない)
PdMのメンタルモデルをそのまま適用できそうで、分割の基準を位置から考える必要がなかった

メンタルモデル: どのような部分から構成され、同関係しあっているか、という捉え方
PdMのメンタルモデルは、kintoneの機能提供を継続してきたことから、いろいろな立場の人に機能を提供し続けたことで、一貫したモデルが形成された
(もともと、アプリ vs アプリ設定をあまり区別せずに開発されてきたが、分けたほうが考えやすいと気づいた)

### トライアルプロジェクト
分割の仕組みは整えたものの、2~3人のメンバ画集に1時間ほど集まって行うのみだった(分割するのも小さな機能)
チームのメンバにヒアリングしたところ、やり方がイメージできていなく手を出しづらいことが判明

普段の開発とは別に、大きめの改善PoCを行う仕組み(プロジェクト)があり、目的、ゴール、メンバ、期間について承認を受ければ集中的に取り組むことが可能
→コード分割トライアルプロジェクト

分割についての学習を目的として、終了後にも参加メンバが分割を行えるように
アプリ設定機能を対象にした(コア機能である程度複雑だが、ほか昨日からの依存はなく、丁度いいサイズ)
呼びかけに応じてもらったメンバが参加し、202211~202301に、週4時間3ヶ月

コア機能の分割は初めてだったので、モブプログラミングで
一気に分割するのは難しいのでどこから初めてどこを目指すか考えたり、分割後の命名を考えながら
同じコードを皆がgら一緒に考えるので、認識がブレず、集中できる
100~200ふぁいるのdiffが出るので、PRをレビューするやり方だと大変。モブだと常にレビューしている形になるので、後でレビューする必要がない
治験の共有・学習にも効果的だった(知っている人は序盤だけリードして終盤は観察する、といった工夫で、つまづきやすいところを確認)

設定項目が20ある家の5つを分割した。
コードの可読性が向上していることは実感できた
参加メンバは各自の担当領域でコード分割を推進している右目的が達成された

### コード分割を通して見えてきた複雑性
例: APIトークン設定項目
ApiTOkenService: APIアプリを使う人がレコードを取得する際、アプリを作る人が設定を変更する佐野、両方で利用される getToken, list, update
ApiToken: APIトークンを表現するデータクラス canView, canAdd, ....

既存のApiTokenService, ApiTokenは設定関連が多いのでappsettingsで良さそうだが、appから利用ができない
→getTokenは設定を変更することとは関係がないと気づいた

→ApiTokenRightService, ApiTokenRightでトークンの取得、トークンに付与されている権限のデータを別のif, クラスに変更した
→ApiTokenRIghtServiceの実装はappsettings内に実装した。設定のテーブルを参照するため。DIでappsettingsへの依存は回避
(m: interfaceはappにあるのか、appserviceにあるのか・・・共通においてあるとしても、Serviceはapp向けだからappにないとおかしい気がする)

使う側・作る側で同じものを使っていたことが、コードを複雑にしていた
(m: これはkintoneみたいなアプリ固有の話だろうな。いや、普通の業務システムでもそのへん変えていいとは覆うけど)
→誰から呼ばれるかによって、必要とされる情報や処理が変わるので、分割することで可読性の工場が期待できる

### 今後の課題とまとめ

課題
コードを分割仕切る
共通部分の分析、整理もやりたい

さらなる分割も勧めたい(DB, サービス分割)


### QA
Q: 分割するに当たり、コードに手を加えるのでデグレードの担保も必要だが、ユニットテス尾はある程度揃っていた?
A: kintoneではE2E(selenium)や専任のQAがいるので、それで担保できたと思う

Q: コード分割が終わったところで、コード行数はどう変わったのか。また、創業数は規模を表すものとして適切な指標ではなく、循環的複雑度などの定量的な指標は分析しているのか
A: プロダクションコードが10安行で、テスト勧めて30万行。まだ終わっていないが、行数はあまり変わっていない。指標について、納得行く指標はまだない。メンバーに聞いて、メンテしやすくなっているかどうか、という印象(定性的な指標)のみ

Q: 分割でモブプロして人を育ててチームに配属された、という話。後で見回して、分割がおかしかった、というチェックなど
A: 今のメンバーが継続して行っているという状況。そこまで一気に進んでいるわけではないが、ヒアリングして絵状況を確認しながら勧めている。まだ深刻な問題は見つかってはいない


## Virtual Threads - 導入の背景と、効率的な使い方
15:00 コンファレンスA

### 3分でわかる Virtual Threads
- JVMが管理する軽量スレッド: ほか言語で言うところのFiber
- スループット工場が目的
    - (HTTPなどの)I/Oを多く含むシステムで効果的
    - 応答時間はちょっとだけ悪化
- 従来のスレッドと使い方は同じ
    - 使用上の注意点はある

### 導入の背景
- 1995 (Java 1,0)
    - 最初からスレッドに対応していた(他の言語では少ない)
    - SIngle Coreの時代
    - Thread = OSスレッドのラッパー
        - Threadオブジェクト生成: 思い & メモリ消費大
        - Threadの切り替え: OSスレッドのコンテキストスイッチ (not parallel. concurrent)
            - コンテキストスイッチ = OSスレッド、Threadおぶじぇくと、JVM Stackの退避・復帰が起きる
        - FYI: スタックトレース = 各スレッドのJVM Stackの情報を見ている
    - →コンテキストスイッチをなるべく発生させないスレッドスケジューリング
    - この頃の主戦場はアプレット(アニメーションのスレッドと、バックグラウンドでイメージをロードするスレッド、程度)
    - →サーバサイドで動かすようになりこの作りは厳しくなってきた
- 2004 (Java 5)
    - Multi Core黎明期
    - Concurrency Utilities
        - スレッドとタスクの分離
            - 普通の開発者はタスクを見る
        - スレッドプール導入
            - スレッド生成のコストは(少し)下がった
            - 応答時間重視: Executors.newFixedThreadPool
            - スループット重視: Executors.newCachedThreadPool
                - Therad per Request/Task
                    - リクエストが増えるとスループットが頭打ちになる(cached threadが足りなくなって新しいスレッドが必要になってしまう)
- 2011 (Java 7)
    - Big Dataの時代 (少し前にHadoop/MapReduceが流行った時代)
    - Fork/Join Framework
        - 分割統治法: タスクを細粒度に分割して処理する
        - Work-Stealingタスクスケジューラー
            - スレッドごとにタスクキューを作る
            - キューがからのときは他のスレッドからタスクを盗む(steal)
        - データ量・計算量が大きい処理を効率的に処理できるように
            - ↑の前提であるので、業務システムにはあまり効果的ではない
                - 業務タスクの特徴: 通信、DBアクセスなどI/O処理が多い
                    - 計算と比べると時間数の桁が3~5桁は違う
                    - I/O処理は待っているだけ→CPUが遊んでしまう
                        - →IO街の間に他の処理を行いたい
- 2014 (Java 8)
    - Project Lambda
    - ラムダ式の導入→処理を関数として記述できるように
    - CompletableFuture(非同期処理を関数で記述) やReactive Programming (Spring WebFlus, Oracle Helidon SE)で、I/O処理を非同期実行することでスループットを向上できる
        - とはいうものの、従来の地区字的な記述と考え方が異なるし、例外処理・デバッグが煩雑
            - 例外はStreamと同じだが、非同期だとスタックトレースが切れてしまうので、どこから呼ばれているのかわからなくなる
- 2023 (Java 21)
    - 従来の地区字的な記述で、例外処理を扱いやすく、デバッグもやりやすいママで、スループットを向上させたい → Virthal Threads

### 動作原理
Virtual Threads: Platform Theradとタスクの仲介
- タスクからはPlatform Threadと区別がつかない
- 実際にはPlatform Threadにマウントして実行される
    - マウントしたスレッドをCarrier THreadと呼ぶ
- 積極的なコンテキストスイッチ
    - OSスレッドのスイッチは重いので、限定継続を導入してコンテキストスイッチの高速化
        - 継続: ここまで処理をやったので、ここから再開するよ
    - Work-Stealingによるスレッドスケジューリング
- 一貫性のあるスタックトレース
    - Carrier THreadが変わってもスタックトレースは途切れない

Virtual ThreadをPlatform Threadにmountして処理を実行、I/O街が発生したらunmountして処理を止める。I/O割り込みがあったら、再びmountして処理を再開する
→Platform Threadが少なくても、I/O待ち時間の活用により大幅なスループット向上が見込まれる。頻繁なコンテキストスイッチによる応答時間の悪化はある
SocketChannelImpl#park -> VirtualThread#park -> yieldContinuation

### 効果的な使い方
- ExecutorServiceを介して利用する
-   - Executors.newVirtualThreadPerTaskExecutor()
    - or Executors.newTHreadPerTaskExecutor(Thread.ofVirtual().factory())
    - 上のはしたを呼んでいるだけ
    - 後はPlatform Threadと同様にtaskをsubmitするだけ
- Virtual THreadの直接生成も可能だが、基本はやらない
    - 独自のExecutorServiceを定義する場合だけ
- ExecutorServiceはJava 19からAutoClosableになった(try-with-resourcesで使えばshutdownを書かなくて良い)

VirtualThread使いたい
→FW使っている
  →WebFluxなどのリアクティブ or CompletableFutureゴリゴリかける →VT使う必要なし
  →Spring Boot, Jakarta EE等→FWの対応まち
→その他 or FW使っていない
  →ボトルベックは?
    →データ量/計算量→ParallelStream, Fork/Join FW, Vector APIのほうが適している
    →File読み込み→Memory Map(メモリの中にファイルを展開できる)
    →通信/DB→Virtual Thread

たいてい、フレームワークが対応するのを待てば十分

スレッドの注意点(特にVirtual Threads)
1. synchronizedをなるべく使わないようにする(ぶっちゃけやめましょう)
    - synchronized = モニタロック
        - リソースへのアクセスを1つのスレッドだけに限定し、ほかのスレッドをブロックするため、ボトルネックになりやすい
    - VIrtual ThraedではCarrier Threadをブロックしてしまう
    - 使わないようにするには:
        - イミュータブルクラスの活用
            - 状態が変化しなければ複数スレッドからもアクセス可能
            - record型の利用
        - 処理結果の受け渡しがある場合Callableを使う
        - どうしてもロックが必要であれば、synchronizedではなくReentrantLockクラスを使う(最適化もされて安全に使える)
            - 用途に合わせてjava.util.concurrent.lockを使う
2. ThreadLocalを使わない
    - ThreadLocalはミュータブル & ライフタイムが不明確
        -  Virtual Threadsはサポートしているがスケールしない
    - ScopedValueで代用する
        - ただしJava 21ではPreview
3. スレッドは使い回さない
    - Virtual Threadsは使い捨て
    - スレッドのプールはExecutorServiceに任せる
        - タスクを記述することに集中
4. ThreadくらすのAPIをもう一度チェック
    - `@Deprecated(forRemoval=true)`なメソッド
    - 直接触るのはバットプラクティス

## API化によるレガシーシステムの改善
16:15 コンファレンスC
パーソルキャリア
dodaのバックエンド開発

10年移譲運用されていて、クラウドリフトも経験している。2018年から開発チームの内製化を開始


### リビルドプロジェクト
フロントエンドの課題
- フレームワークを利用していないことによる限界
    - スタイルの変更が様々な画面に影響を及ぼす(例: 共通で配置されるボタンのスタイルが、画面のスタイルと衝突して崩れる)
    - 画面内の複雑な変化をjQueryで行うことが難しい
- CLS(画面の安定性)指標の悪化
    - クライアントサイドでレンダリングすることによる画面のガタツキ
        - Google検索順位の低下の可能性

サーバサイドの課題
- テストコードがかけていない・書きにくい
    - 親クラスに小クラス向けの分岐があったりstaticメソッドを呼んでいたり
    - 修正がどこに影響するのかわからない状態
- 可読性が低い
    - 性的コード解析でエラーが大量
    - フォーマットが統一されていない

インフラの課題
- 変更に時間がかかる
    - クラウドリフトしてEC2に載せただけなのでスケールしにくい

リビルドプロジェクトでの取り組み
- フロントエンド: React/Next,jsの導入
    - コンポーネント歌詞意図しないスタイルが当たることを防ぐ
    - Server Side Rendering
- サーバサイド
    - APIのために新しいシステムを作成
        - 画面と切り離すことで大規模な変更(ライブラリのバージョンアップなど)をしやすい状態にする
    - 負債を産まないための仕組みを整備→テストコードを書きやすく
- インフラ
    - コンテナ(ECS), IaC(CloudFormation)

### API開発の工夫
画面とは切り離されたが、フロントエンドとの連携をしやすいAPIを設計することが重要になった

#### API開発の流れ
1. サーバサイドとフロントエンドの認識合わせ
2. サーバサイド: IF定義作成(OpenAPI)
3. フロントエンド: モックを釐王して実装
4. サーバサイド: API実装
5. 統合してテスト

OpenAPIが共通の指針となるため、実装とかいりしないようにする必要がある
→不フロント・サーバどちらもライブラリで工夫している

サーバサイド: springdoc-openapを使う
- コントローラからOpenAPIを生成してswagger-uiで閲覧可能
- 必要な情報はアノテーションで記載
- コードから生成するため乖離は発生せず、なれたJavaのコードで型安全に生成できるため効率も良い
    - Controllerだけ作ってPR出す形

フロントエンド: aspidaを使ってOpenAPIから型情報を生成して使う
モックサーバはPrismで作成

#### API設計で気をつけていること
- 命名
    - xxxflagにしない
    - DBのカラム名をそのままレスポンスに利用しない (すでにある実装をベースに移行するため、意識しないと発生しやすい)
- 広報互換性を保つ(破壊的変更があるとフロントも修正画筆王になる)
- Web API: The Good Partsを酸法にしている

- 既存の仕様が良くない場合は、そのままAPIにするのではなく可能な範囲で再定義する
    - 10年運用されているので、途中で用件が変わって、その過程でおかしな仕様もある
    - せっかく作り直すので、変な仕様は再定義する
    - スケジュールもあるので可能な範囲を見極めて実施
- 画面に依存したAPIを作らない
    - 特定の画面のみで利用される文言などをAPIで返さない
        - 例: JSPにわたしていたDTPにセットしているタイトル部分をAPIに含めない
        - SEO対策のカラムの値を返さない(返すなら他の用途でも使える名前で返す)

### 技術負債を生まないための取り組み
大規模な修正はしやすくなったが、IFの状態が同じことの担保が必要。
内部の状態(コード品質やセキュリティ)についても一定の状態を保てないと開発効率が落ちる

- コンテナ(ECS)の利用
    - EC2だと状態の再現が難しい
    - オートスケーリング赤脳になった
        - これまではピーク時でもさばけるくらいのインスタンスを立てていた
        - アプリが落ちても自動復旧・高速な切り戻しが可能に
    - 脆弱性へ即時に対応
        - ECRのイメージスキャンやAmazon Inspector

ECSでJavaアプリケーションを動かす際に発生した課題
Fargate 1.4.0だと通信の影響で利用できないが、EFSが1.4.0でないと利用できないという問題。JFR/HeapDumpの出力をどうするか
→サイドカーをマウントしてinotifyで通知→s3に転送、とした

各種テスト
- 静的コード解析
    - CIで実行し、エラー発生時にはデプロイさせないようにした
    - SpotBugs, Spotless+google-java-formatを使用
    - ArchUnitによるアーキテクチャテストも使用
- 単体テスト
    - staticメソッド呼び出しをDIに置き換える
- インテグレーションテスト: API単位でのDBを含めたテスト
    - DB/外部API周りの整備が難しいところだと考えている
    - Dockerでテストの再現性をもたせるようにした
        - TestContainers + SprintBootTest
        - 外部APIはWireMockを1付きどうして、エンドポイントで振り分ける
        - DBはFlywayでセットアップ
    - `@Sql`でOracle固有のSQLを使用している箇所は、Beanをテスト用に差し替え

ライブラリの脆弱性の定期診断
- OWASP Dependency-Checkプラグインで日次でチェック
- 一定以上のCVSSスコアのものを検知した場合はエラーを発生させてメール通知
- テストを充実させたことで、バージョンアップを安心して実行できるように


### WA
Q: チーム体制、フロントエンドとバックエンドのエンジニアは別?
A: もともとは担当分けていなかったが、このプロジェクトを進める際は、別々のチーム。完成した後の画面は、また一緒の担当で各スクラムが触る。登壇者は移行完了後のチームには携わっていないが、Reactに慣れている人が多くない、というのが現状の課題。今後チームをwけるべきかは検討中

Q: API管理に関連して、フロントエンドがどのAPIを呼んでいるかは、同管理しているのか
A: まだ画面数がそこまで多くないのであまり課題にはなっていない。利用する画面が多くなってきたら、また検討したい

Q: API分割粒度・再利用のポリシー
A: 再利用性を考えて、個別に検討している。厳格なルールはない

Q: React導入に当たり、社内で勉強会はやった?
A: フロント担当の人が何度かやってる。チームによって個別にやっているかも

Q: 移行にかけた時間
A: dudaトップページから移行したが、3ヶ月~半年。その後は個別に順次勧めている

Q: チーム分けの基準
A: dudaはKPIに紐づくチーム分けになっていて、その中でフロント・サーバサイド療法を担当している。それ以外にプロジェクトを進めるチームがあり、今回の移行もプロジェクトの中でフロント/サーバに分けて勧めていた

Q: DBのカラム名とAPIの項目名が一致しないように、という話。どういう弊害があったか
A: カラム名・テーブル名、連番になっている(歴史的敬意 例: X001, X002)。その名前だと何を意図しているのかわかりにくいので、APIを使う人がわかるようなカラム名に変えて返す。

## Revisiting Design Patterns after 20
17:15 コンファレンスA

GoFのDesign Patternsは有用だったが20年が経過している
Eric EvansのDDD、Efective Java

Java 20: 2023/03/21
DP: general, reusable solution, 与えられた文脈の問題を解決するもの

Pattern Language
ベターなコードを書くことを可能にしてくれる

original DP: favor composition over inheritance
facor smaller interfaces (1メソッドだけとか)
favor existing interfaces (idをみえば何をやっているかわかる)) ^> Java8: Fnctional Interfaces 新しいインタフェースを作らなくても住むように

ソースコードはすべてGitHubに公開されているよ

### Command
```java
interface Command {
    void execute();
}
```
→java.lang.Runnable

### Observer
```java
interface Observer<T> {
    void observer(T t)
}
```
-> java.util.function.Consumer
Observer作るならCOんすめｒ、Observableは自分で作るべき。何が必要か考えて

### Strategy
interface作って、実装作って、状況によって作るインスタンスを変える
→Strategyインタフェースは作らずに、`Function`を使う
    新しい(自前の)インタフェースを学ばなくて良い
    ラムダ式やメソッド参照でかける
    this is the modern way!

### Template Method
in 2023, no template method

Template Methodは30年前のもの。今作るなら、Strategyで良い。

### Singleton
これはかんたん。
legacy singleton: static finalなINSTANCEとgetInstance()メソッド。
since java 8: use enum
```java
public enum LiveRevisitedSingleton {
    INSTANCE;
}
```
legacy: 長い
new: JVMが常に1インスタンスしか作られないことを保証してくれる


### Interpreter
legacy: enumで演算子をつくってどうのこうのする
    - トークンが演算子かどうかbullチェックしていた
new: Optionalを活用
(m: ちょいわからなかったから後でコード見る)


### Chain of Responsibility
3種類の通知を設定できるアプリを考える(google wallet, email, sns)
実際に通知するのは1つだけで、どれを使うかはユーザーが選択する
→チェーンを作る。通知を有効にしている方法なら通志する。そうでなければ次の通知方法に処理を進める
→コピペで作れる

new version
1. Predicateで条件を判断できる
2. 実際の通知内容はそれぞれ違う

→ConsumerとPredicateを実装したクラスを、それぞれの通知方法の実装にする
あるいは、PredicateとConsumerを持つrecordを作る

前者
NotifierのSreamを作ってfilterしてfindFirstしてifPresentで実行

old versionだと順番カエルのちょっと面倒だったよね

### Visitor
one of most complecated design pattern.
とてもエレガントなソリューションだけど、読むのが難しい
1つのオペレーションしか使わないならStrategyを使えばいいけど、他にも使うならVisitorがより良いソリューションになる
扱うチケットの種類が増えると、visitorインタフェース似メソッドを追加しないといけなくて、すべての実装クラスも修正が必要になる

modern version:
sealed classes + pattern matchingが使える

扱うチケットの種類が増えると、switch文(式)に新しいcaseラベルが必要になる(コンパイルエラーで検知できる)
型だけでなく、状態も使って分岐できるから、Visitorよりもフレキシブル


### Specification
DDDのデザインパターン
条件に名前をつけよう、って話。マジックナンバー、マジックコンディションは将来的に忘れる。
→コードがセルフドキュメントになる
`Specification`インタフェースでand/or/notも使って組み合わせられるようにする

now: Predicateで良い。and/or/notも提供される


github.com/yanaga/revisiting-design-pattern


### QA
Q: varを使っている意図
A: 位置を揃えたかったんだw
プロダクションコードだと、左右に同じ型が出てくるとき。


## JUnitテストをCI環境で並列で実行する方法とその速度、スケーラビリティ
18:15 コンファレンスB

今日: CI環境で、JUnitテストを複数のマシンで自動的に手分けして実行することで、CI街の時間を削減する、2023年5月現在のポピュラーな技術による、現実的かつシンプルな方法。

答え: jest --shard
※去年の5月くらいに出てきた
これをGradleでやりたい

1. すべてのテストクラスのリストをファイルに書き出しておく
2. 各マシンは自分の担当テストをリストから抽出して実行する

build.gradleで`includeTestMatching`の呼び出しが動的に変化するようにプログラミングすれば良い

→4並列までは時間短縮できたが、6並列移譲だと安定しなくなった
→8並列のCI一発で100円超えそうな課金が発生
→→カネがかかる割に対して早くなってない。並列度を上げても無駄そう

対して早くならなかった/速度が安定しなかった原因
1. 小規模すぎる(テストクラスが80個しかない)
2. Springの起動処理が重い
    - ※テストコンテキストは複数のテストクラス感で可能な限り使い回される
    - Springの起動時間は並列度を上げても変わらないので、テスト時間が短縮されても起動時間分の下駄が残っている
3. キャッシュ
    - `~/.gradle/*`. `node_modules/*`, docker cache

キャッシュが常にあるローカル環境 vs キャッシュが常に消えていて、どこかに対比させておいたキャッシュを毎回下記戻す環境は、勝負にならない
→`ubuntu-latest`で高速化しようとするのが無理筋
→→`self-hosted`



