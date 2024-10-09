# 簡易Slackアプリ

## テーブル設計

### ・userテーブル
```
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | int          | NO   | PRI | NULL    | auto_increment |
| name       | varchar(255) | NO   |     | NULL    |                |
| age        | int          | NO   |     | NULL    |                |
| email      | varchar(255) | NO   |     | NULL    |                |
| created_at | datetime     | NO   |     | NULL    |                |
| updated_at | datetime     | YES  |     | NULL    |                |
+------------+--------------+------+-----+---------+----------------+
```
### ・channelテーブル
```
+----------------+--------------+------+-----+---------+----------------+
| Field          | Type         | Null | Key | Default | Extra          |
+----------------+--------------+------+-----+---------+----------------+
| id             | int          | NO   | PRI | NULL    | auto_increment |
| channel_name   | varchar(255) | NO   |     | NULL    |                |
| create_user_id | int          | NO   |     | NULL    |                |
| created_at     | datetime     | NO   |     | NULL    |                |
| updated_at     | datetime     | YES  |     | NULL    |                |
+----------------+--------------+------+-----+---------+----------------+
```
### ・messageテーブル (channelテーブルに紐づく)
```
+------------+----------+------+-----+---------+----------------+
| Field      | Type     | Null | Key | Default | Extra          |
+------------+----------+------+-----+---------+----------------+
| id         | int      | NO   | PRI | NULL    | auto_increment |
| channel_id | int      | NO   |     | NULL    |                |
| user_id    | int      | NO   |     | NULL    |                |
| message    | text     | NO   |     | NULL    |                |
| created_at | datetime | NO   |     | NULL    |                |
| updated_at | datetime | YES  |     | NULL    |                |
+------------+----------+------+-----+---------+----------------+
```

## API設計
・ユーザー作成 POST localhost:8080/user
```
root@DESKTOP-RTPGFSE:~# curl -X POST -H "Content-Type: application/json" -d '{"name":"Taro", "age":20, "email":"hoge@gmail.com"}' localhost:8080/user
```
・ユーザー取得 GET localhost:8080/user/:user_id
```
root@DESKTOP-RTPGFSE:~# curl -X GET localhost:8080/user/1
{"name":"Taro","age":20,"email":"hoge@gmail.com","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}
```
・ユーザー情報更新 PUT localhost:8080/user/:user_id
```
root@DESKTOP-RTPGFSE:~# curl -X PUT -H "Content-Type: application/json" -d '{"name":"Ken", "age":30, "email":"hogehoge@gmail.com"}' localhost:8080/user/1
```
・ユーザー削除 DELETE localhost:8080/user/:user_id
```
root@DESKTOP-RTPGFSE:~# curl -X DELETE localhost:8080/user/1
```
・チャンネル作成 POST localhost:8080/channel
```
root@DESKTOP-RTPGFSE:~# curl -X POST -H "Content-Type: application/json" -d '{"channel_name":"Today Issue", "create_user_id":1}' localhost:8080/channel
```
・チャンネル一覧取得 GET localhost:8080/channel (チャンネルのみ取得)
```
root@DESKTOP-RTPGFSE:~# curl -X GET localhost:8080/channel
[{"channel_name":"Pressing matter","create_user_id":1,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},{"channel_name":"Today Issue","create_user_id":1,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]
```
・チャンネル詳細取得 GET localhost:8080/channel/:channel_id (チャンネルに紐づくメッセージも全て取得)
```
未実装
```
・チャンネル情報更新 PUT localhost:8080/channel/:channel_id
```
root@DESKTOP-RTPGFSE:~# curl -X PUT -H "Content-Type: application/json" -d '{"channel_name":"2024_10_Issue"}' localhost:8080/channel/1
```
・チャンネル削除 DELETE localhost:8080/channel/:channel_id (チャンネルに紐づくメッセージも全て削除)
```
root@DESKTOP-RTPGFSE:~# curl -X DELETE localhost:8080/channel/1
```
・メッセージ投稿 POST localhost:8080/message
```
root@DESKTOP-RTPGFSE:~# curl -X POST -H "Content-Type: application/json" -d '{"channel_id":1, "user_id":1, "message":"test posting"}' localhost:8080/message
```
・メッセージ更新 PUT localhost:8080/message/:message_id
```
未実装
```
・メッセージ削除 DELETE localhost:8080/message/:message_id
```
未実装
```

