---
title: "[Eclipse][EGit] .projectがルートに配置されるようにGit管理する"
date: 2021-06-02T08:07:15+0900
draft: false
categories: 
tags:
    - eclipse
    - EGit
    - GIt
aliases:
    - /2021/06/02/264.html
---

<!-- wp:paragraph -->
<p>前提: Gitの<code>user.name</code>, <code>user.email</code>が適切に設定されていること。</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>この記事の手順で作成したリポジトリは <a href="https://github.com/maruTA-bis5/git-init-eclipse ">https://github.com/maruTA-bis5/git-init-eclipse </a>です。</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>1. EclipseのProject Explorerより、Git管理したいプロジェクトを右クリックし、Team &gt; Share Project...を選択する。</p>
<!-- /wp:paragraph -->

![share-project-menu](./share-project-menu.png)

<!-- wp:paragraph -->
<p>2. Share Projectダイアログの<code>Use or create repository in parent folder of project</code>を選択する。</p>
<!-- /wp:paragraph -->

![share-project-dialog](./share-project-dialog.png)

<!-- wp:paragraph -->
<p>3. 対象のプロジェクトを選択し、<code>Create Repository</code>ボタンをクリックする。</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>4. <code>Finish</code>をクリックする。</p>
<!-- /wp:paragraph -->

![share-project-dialog-git-init](./share-project-dialog-git-init.png)

<!-- wp:paragraph -->
<p>ここまででローカルリポジトリが作成されます。次は初回のコミットを行います。</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>5. <code>Git Staging</code>ビューを開く。</p>
<!-- /wp:paragraph -->

![git-staging-view](./git-staging-view.png)

<!-- wp:paragraph -->
<p>6. <code>Unstaged Changes</code>のファイルを<code>Staged Changes</code>に移動し、<code>Commit Message</code>を入力して<code>Commit</code>をクリックする。<br>このとき、<code>*.class</code>ファイルは表示されていませんが、<code>*.class</code>ファイルが出力されるディレクトリは<code>.gitignore</code>に記載されているため無視されます。他にも無視するファイルがあれば<code>.gitignore</code>に追記しておきます。<br>※<code>.gigignore</code>はProject Explorerには表示されないので、このタイミングで修正するか、Git Repositoriesビューから開く必要があります。</p>
<!-- /wp:paragraph -->

![git-staging-view-staged](./git-staging-view-staged.png)

<!-- wp:paragraph -->
<p>初回コミットが完了したら、リモートリポジトリにpushします。(リモートリポジトリでの管理が不要ならここまで)</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>7. Git Stagingビューの<code>Push HEAD</code>ボタンをクリックする。</p>
<!-- /wp:paragraph -->

![git-staging-view-push-head](./git-staging-view-push-head.png)

<!-- wp:paragraph -->
<p>8. <code>Location</code> &gt; <code>URI</code>にリモートリポジトリのURL、認証が必要なら<code>Authentication</code>を入力し<code>Preview &gt;</code>をクリックする。<br>画像ではSSHを使用する形で入力していますが、HTTP(S)でも問題ありません。<br>パスワード認証が必要な場合は、<code>Authentication</code>を入力しておきます。</p>
<!-- /wp:paragraph -->

![set-remote](./set-remote.png)

<!-- wp:paragraph -->
<p>9. ブランチの指定は基本的に変更しなくて良いはず。前のステップで入力した認証情報に誤りがあればこの段階でエラーが表示されるので、一度戻って修正する。問題なければ<code>Preview &gt;</code>をクリックする。</p>
<!-- /wp:paragraph -->

![select-branch](./select-branch.png)

<!-- wp:paragraph -->
<p>10. pushのプレビューを確認して<code>Push</code>をクリックする。</p>
<!-- /wp:paragraph -->

![push-preview](./push-preview.png)

<!-- wp:paragraph -->
<p>11. pushが完了するとpush結果が表示される。失敗した場合もこのダイアログが表示されるので、必ず<code>Message Details</code>の内容を確認する。</p>
<!-- /wp:paragraph -->

![push-result](./push-result.png)
