#280blockerget_filter_for_Golung
280blocker様の当月のブロックルールを取得するためのGo言語で書かれたコードです。
Chrome拡張機能:ublockOriginにインポートできるようAdblockPlus形式ファイル準拠で取得しています。

go run ./280getfilter.go
もしくは
exeを起動
どちらかしていただくと
ヘッドレスブラウザを起動し280blockerのフィルタをクリップボードに貼り付けます
後はublockOriginのカスタムフィルターに設定していただくか、ユーザご自身で取得内容を改変するなり操作してください

プログラミング・コードがわからないなど不慣れな方は、添付のexeアプリをご利用下さい。

このコードは改変・再配布等自由です。

※github.com/atotto/clipboard"及び"github.com/go-rod/rod"の利用規約についてはライブラリ作者様に従います。

※ブロックフィルタルールの利用規約については280blocker様に従います。
280blocker様のサイト
https://280blocker.net/

LICENSE → 【licence.txt】 in this repository.

ソースコード及びexe起動時にWindows のセキュリティ機能が leakless.exe というファイルをウィルスまたは望ましくないソフトウェアとして検出された場合
Windows のセキュリティ設定でこのファイルを許可する必要があります。具体的な手順は次のとおりです。
[ウィルスと脅威の防止] > [保護の履歴] で 該当exeに関する項目を [復元]
