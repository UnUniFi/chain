# security design

利用状況整理  
 ユーザーが、パブリック UI を通してシステムを利用し  
 トークンや NFT のような資産がブロックチェーンで保管し  
 機密性は担保しない  
 ブロックチェーンでデータそのものを保護する

現状

- 守るべきもの
  - システムのコンセンサス
  - 利用者の NFT
  - 利用者の残高
  - トークン価値

基本機能

- アカウント管理
  - 特定  
    できない
  - 防御  
    できない
  - 検知  
    多分できない
  - 対応  
    ブラックリストに入れて、Send できないようにする
  - 復旧  
    新しいアカウントに送付
- 認証  
  Kepler にお任せ
- アクセス制御
  - 特定  
     役割の定義
  - 防御  
     役割によるアクセス制御
  - 検知  
     アクセス制御の監視
  - 対応  
     アクセス権限の無効化
  - 復旧  
     アクセス権限の有効化
- データ保護  
  ブロックチェーンだから省略
- バックアップ  
  VPS のバックアップ  
  チェーンのバックアップ
- 出品
- 入札
- 落札
- 共通機能
- ログ管理
- ログ設計

直接的

- 誤認させる  
  fake NFT→ 認証プロセスを設ける
- ~~51% attack~~
- Goldfinger attack  
  NFT 認証を壊す →tendermint に任せつつ VPS として DDos 攻撃に耐える設計にしておく
- ~~Hard Fork~~
- marketplace manipulation  
  Mischievously listing items at auction→delay process を設ける  
  fake nft などを大量出品して検索させなくする →GUU ステーキングを強制させつつ X 品目から手数料が発生するようにする
- Sybil Attack  
  misleading bidding→delay process で対応  
  Price Manipulation by Bidding→ 自動刻み入札を採用(like ヤフオク)して、入札ハードルを下げて 1 位と 2 位の入札額の差をうめる
- DNS Hijack  
  ~~BGP Hijack~~
- Eclipse Attack  
  攻撃方法を考える
- Wallet Attack
- DDoS  
  DDoS msg→ システムアーキテクトとして対応する
- ~~Dusting Attack~~
- attack dapp  
  間接的
- Pricefeed の元を乗っ取る →IBC 接続に変更する
