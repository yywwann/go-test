package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	genUTCTime()
	genUTCTime1()
	parseTime()
	parseDuration()

	now := time.Now()
	formatTime := timeStampFormat(now.UnixNano())
	fmt.Println(formatTime)

	fmt.Println(now)
	fmt.Println(now.UnixNano()) // 生成时间戳19位
	fmt.Println(now.UnixNano() / 1e6)
	fmt.Println(now.Unix()) // 时间戳10位

	// 生成两个时间的差值 ：时间段
	lastTime := time.Unix(0, 1550714939515549582)
	dur := now.UTC().Sub(lastTime.UTC())
	fmt.Printf("duration is: %v\n", dur)
	fmt.Printf("duration to string is: %s\n", dur.String())
	fmt.Printf("duration to Nanoseconds is: %v\n", dur.Nanoseconds())
	fmt.Printf("duration to Seconds is: %v\n", dur.Seconds())
	fmt.Printf("duration to Minutes is: %v\n", dur.Minutes())
	fmt.Printf("duration to Hour is: %v\n", dur.Hours())

	secondTime := time.Unix(0, 1550721050238199291)
	dur1 := now.UTC().Sub(secondTime.UTC())
	fmt.Printf("Second Duration is: %v\n", dur1)
	fmt.Printf("duration to Truncate is: %v\n", dur.Truncate(dur1))
	fmt.Printf("duration to Round is: %v\n", dur.Round(dur1))

	// 给now时间加上一个时间段，得到一个新的时间
	fur := now.Add(dur)
	fmt.Println(fur)

	// 比较两个time.Time格式的时间变量
	boolEq := now.Equal(fur)
	boolBe := now.Before(fur) // whether now is before fur
	boolAf := now.After(fur)  // whether now is after fur
	fmt.Println(boolEq)
	fmt.Println(boolBe)
	fmt.Println(boolAf)

	// 当前时间减去实参里的时间，得到的时间段
	durSince := time.Since(secondTime) // 正数
	durUntil := time.Until(secondTime) // 负数
	fmt.Printf("Since the now is: %v, Until the now is: %v\n", durSince, durUntil)

	// 给指定时间增加年、月、日
	AddTime := now.AddDate(1, 2, 3)
	fmt.Printf("AddDate to now is: %v\n", AddTime)

	// 得到指定时间的年、月、日， 均为int类型
	year, month, day := now.Date()
	fmt.Printf("year: %d, month: %d, day: %d\n", year, month, day)
	yearOnly := now.Year()
	fmt.Printf("only year(int) is: %d\n", yearOnly)
	monthOnly := now.Month().String()
	fmt.Printf("only month(Month.String()) is: %s\n", monthOnly)
	weekdayOnly := now.Weekday().String()
	fmt.Printf("only weekday(Weekday.String()) is: %s\n", weekdayOnly)
	dayOnly := now.Day()
	fmt.Printf("only day(int) is: %d\n", dayOnly)

	// 得到时、分、秒
	hour, min, sec := now.Clock()
	fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)
	Hour := now.Hour()
	fmt.Printf("Hour(int) is: %d\n", Hour)
	Minute := now.Minute()
	fmt.Printf("Minute(int) is: %d\n", Minute)
	Second := now.Second()
	fmt.Printf("Second(int) is: %d\n", Second)
	Nanosecond := now.Nanosecond()
	fmt.Printf("Nanosecond(int) is: %d\n", Nanosecond)
	YearDay := now.YearDay()
	fmt.Printf("YearDay(int) is: %d\n", YearDay)

	// 序列化与反序列化
	byteTime, err := secondTime.MarshalBinary()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("MarshalBinary the secondTime is: %v\n", byteTime)
	tmpByteTime := time.Time{}
	err = tmpByteTime.UnmarshalBinary(byteTime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("UnmarshalBinary the secondTime is: %v\n", tmpByteTime)

	jsonTime, err := secondTime.MarshalJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("MarshalJSON the secondTime is: %v\n", jsonTime)
	tmpJsonTime := time.Time{}
	err = tmpJsonTime.UnmarshalJSON(jsonTime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("UnmarshalJSON the secondTime is: %v\n", tmpJsonTime)

	textTime, err := secondTime.MarshalText()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("MarshalText the secondTime is: %v\n", textTime)
	tmpTextTime := time.Time{}
	err = tmpJsonTime.UnmarshalText(textTime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("UnmarshalText the secondTime is: %v\n", tmpTextTime)

	mockTime := time.Date(1, 2, 3, 4, 5, 6, 7, now.Location())
	fmt.Printf("mock a time follow local location is: %v\n", mockTime)

	// Location from zoneinfo.go
	location := now.Location()
	fmt.Printf("Current location is: %v\n", location)
	fmt.Printf("Current location(string) is: %s\n", location.String())

	fixLocation := time.FixedZone("local", -8)
	fmt.Printf("fixed the local zone is: %v\n", fixLocation)

}

func genUTCTime() {
	now := time.Now().UTC()
	UTCTimeFormat := now.Format("2006-01-02T15:04:05.000Z")
	fmt.Println(UTCTimeFormat)
}

func genUTCTime1() {
	now := time.Now()
	// LoadLocation()，可以在"$GOROOT/lib/time/zoneinfo.zip"的解压后的文件夹中找到需要设置的时区
	local, _ := time.LoadLocation("UTC") // "" "UTC" "Local" "America/Los_Angeles"
	UTCTimeFormat := now.In(local).Format("2006-01-02T15:04:05.000Z")
	fmt.Println(UTCTimeFormat)
}

// 格式化当前时间，并进行解析
func parseTime() {
	now := time.Now().UTC()
	UTCTimeFormat := now.Format("2006-01-02T15:04:05.000Z")
	// 较之ParseInLocation()，没有指定时区，默认时区为"UTC"
	parseResult, err := time.Parse("2006-01-02T15:04:05.000Z", UTCTimeFormat)
	if err != nil {
		fmt.Printf("parse the time has error is: %v\n", err)
	}
	fmt.Printf("the result of parse(UTC) is: %v\n", parseResult)

	// 增加一个参数，时区
	parseResult1, err := time.ParseInLocation("2006-01-02T15:04:05.000Z", UTCTimeFormat, time.Local)
	if err != nil {
		fmt.Printf("parse the time has error is: %v\n", err)
	}
	fmt.Printf("the result of parse(local) is: %v\n", parseResult1)

	dateFormat := parseResult.Format("2006-01-02")
	fmt.Printf("format the parsed result is: %s\n", dateFormat)
}

// 解析时间段
func parseDuration() {
	strDuration := "2h49m39s"
	duration, err := time.ParseDuration(strDuration)
	if err != nil {
		fmt.Printf("Parse the duration has error is: %v\n", err)
	}
	fmt.Printf("Parse the duration is: %v, its type is: %v\n", duration, reflect.TypeOf(duration))

}

func timeStampFormat(timestamp int64) string {
	tm := time.Unix(0, timestamp).UTC()
	UTCTimeFormat := tm.Format("2006-01-02T15:04:05.000Z")
	return UTCTimeFormat
}
