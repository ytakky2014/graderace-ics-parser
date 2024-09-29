# 内容
- graderace-ics-parser はJRAが提供している重賞カレンダーをパースして、重賞の日程と重賞名を返すツールです
- JRA公式データのURLも生成し表示します
- 現在は2024年の重賞のみ対応しています。
- 3連休での開催を考慮して月曜日の重賞も表示されます
- 年末等で土日月以外に重賞があった場合の考慮はされていません。

# レーシングカレンダーのダウンロード
- [JRAの公式サイトのトップからホーム>競馬メニュー>レーシングカレンダー](https://www.jra.go.jp/keiba/calendar/)>他のカレンダーと連携すると遷移しファイルをダウンロードして、ディレクトリ内に配置してください

# 実行
- go run main.go

# 実行例
```
~/graderace-ics-parser$ go run main.go
20241005はサウジアラビアロイヤルカップ(GIII)があります。
20241006は京都大賞典(GII)があります。
20241006は毎日王冠(GII)があります。
20241007は重賞非開催日です
https://www.jra.go.jp/keiba/thisweek/2024/1005_1/
https://www.jra.go.jp/keiba/thisweek/2024/1006_1/
https://www.jra.go.jp/keiba/thisweek/2024/1006_2/
```
