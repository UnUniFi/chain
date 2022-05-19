# security design

利用状況整理
  ユーザーが、パブリックUIを通してシステムを利用し
  トークンやNFTのような資産がブロックチェーンで保管し
  機密性は担保しない
  ブロックチェーンでデータそのものを保護する

現状
  守るべきもの
    システムのコンセンサス
    利用者のNFT
    利用者の残高
    トークン価値
  利用者

基本機能
  アカウント管理
    特定
      できない
    防御
      できない
    検知
      多分できない
    対応
      ブラックリストに入れて、Sendできないようにする
    復旧
      新しいアカウントに送付
  認証
    Keplerにお任せ
  アクセス制御
    特定
      役割の定義
    防御
      役割によるアクセス制御
    検知
      アクセス制御の監視
    対応
      アクセス権限の無効化
    復旧
      アクセス権限の有効化
  データ保護
    ブロックチェーンだから省略
  バックアップ
    VPSのバックアップ
    チェーンのバックアップ
  出品
  入札
  落札
  共通機能
  ログ管理
  ログ設計

直接的
  誤認させる
    fake NFT→認証プロセスを設ける
  ~~51% attack~~
  Goldfinger attack
    NFT認証を壊す→tendermintに任せつつVPSとしてDDos攻撃に耐える設計にしておく
  ~~Hard Fork~~
  market manipulation
    Mischievously selling items at auction→delay processを設ける
    fake nftなどを大量出品して検索させなくする→GUUステーキングを強制させつつX品目から手数料が発生するようにする
  Sybil Attack
    misleading bidding→delay processで対応
    Price Manipulation by Bidding→自動刻み入札を採用(like ヤフオク)して、入札ハードルを下げて1位と2位の入札額の差をうめる
  DNS Hijack
  ~~BGP Hijack~~
  Eclipse Attack
    攻撃方法を考える
  Wallet Attack
  DDoS
    DDoS msg→システムアーキテクトとして対応する
  ~~Dusting Attack~~
  attack dapp
間接的
  Pricefeedの元を乗っ取る→IBC接続に変更する
