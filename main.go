package main

import (
	"fmt"
	"log"
	"os"
	"time"

	ics "github.com/arran4/golang-ical"
)

type gradeRace struct {
	date string
	name string
}

func main() {
	gradeRaces, err := icsData()
	if err != nil {
		log.Fatal(err)
	}

	weekendDays := newxtWeekendDays()
	var urls []string
	// 3連休等で月曜日開催もあるので月曜日もチェックする
	// 年末等で土日月以外の可能性もあるがその点に関しては考慮しない
	// 開催があったとしても年一あるかのレアケースのため
	for _, w := range weekendDays {
		urls = append(urls, gradeRaceHeld(gradeRaces, w)...)
	}
	for _, url := range urls {
		fmt.Printf("%s", url)
	}
}

// gradeRaceHeld はchekDayに重賞が開催されるかを確認する
// 重賞が開催される場合JRAの公式データページのURLを返す
// 同日で複数重賞が開催される可能性があるので、返却値はslice
func gradeRaceHeld(graces []gradeRace, checkDay time.Time) []string {
	num := 0
	urls := []string{}
	for _, data := range graces {
		if data.date == checkDay.Format("20060102") {
			num++
			fmt.Printf("%sは%sがあります。\n", data.date, data.name)
			urls = append(urls, fmt.Sprintf("https://www.jra.go.jp/keiba/thisweek/%s/%s_%d/\n", checkDay.Format("2006"), checkDay.Format("0102"), num))
		}
	}
	if num == 0 {
		fmt.Printf("%sは重賞非開催日です\n", checkDay.Format("20060102"))
	}
	return urls
}

// nextWeekendDays は実行日を起点として次の週末を返す
// 3連休も考慮して土日月とする
func newxtWeekendDays() []time.Time {
	const weekDay = 6
	now := time.Now()
	nsd := weekDay - now.Weekday()
	if nsd <= 0 {
		nsd = nsd + 7
	}

	nextWeekendStaurday := now.AddDate(0, 0, int(nsd))
	nextWeekendSunday := nextWeekendStaurday.AddDate(0, 0, 1)
	nextWeekendMonday := nextWeekendSunday.AddDate(0, 0, 1)

	return []time.Time{nextWeekendStaurday, nextWeekendSunday, nextWeekendMonday}
}

// icsDataは jraのicsデータを解析して、重賞開催日と重賞名の構造体のスライスを返す
// pathは今後拡張予定
func icsData() ([]gradeRace, error) {
	path := "jrarace2024.ics"
	f, err := os.Open(path)
	if err != nil {
		return []gradeRace{}, err
	}
	defer f.Close()

	cal, err := ics.ParseCalendar(f)
	if err != nil {
		return []gradeRace{}, err
	}
	events := cal.Events()
	gradeRaces := []gradeRace{}

	for _, e := range events {
		date := e.GetProperty(ics.ComponentPropertyDtStart).Value
		name := e.GetProperty(ics.ComponentPropertySummary).Value
		gradeRaces = append(gradeRaces, gradeRace{date: date, name: name})
	}
	return gradeRaces, nil
}
