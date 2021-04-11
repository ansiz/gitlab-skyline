package main

import (
	"encoding/json"
	"log"
	"time"
	"fmt"
)

type Origin struct {
	Date  string
	Day   int
	Total int
}

type Target struct {
	DayInYear int
	Day       int
	Total     int
}

const (
	layout = "2006-01-02"
)

func GetDayOfYear(date string) int {
	t, _ := time.Parse(layout, "2020-01-03")
	return t.YearDay()
}

func Convert2Target(o Origin) Target{
	return Target{
		DayInYear: GetDayOfYear(o.Date),
		Day: o.Day+1,
		Total: o.Total,
	}
}

func main() {
	originDatas := []Origin{}
	err := json.Unmarshal([]byte(data), &originDatas)
	if err != nil {
		log.Fatal(err)
	}
	martrix:=[][]int{}
	for _,d:= range originDatas {
		td:=Convert2Target(d)
		martrix = append(martrix, []int{td.DayInYear,td.Day,td.Total})
	}
	str,_:=json.Marshal(martrix)
	fmt.Printf("%s\n",str)
}

const data = `[
	{
	  "date": "2019-03-31",
	  "day": 0,
	  "month": 3,
	  "total": 0,
	  "week": 14,
	  "year": 2019
	},
	{
	  "date": "2019-04-01",
	  "day": 1,
	  "month": 4,
	  "total": 0,
	  "week": 14,
	  "year": 2019
	},
	{
	  "date": "2019-04-02",
	  "day": 2,
	  "month": 4,
	  "total": 0,
	  "week": 14,
	  "year": 2019
	},
	{
	  "date": "2019-04-03",
	  "day": 3,
	  "month": 4,
	  "total": 0,
	  "week": 14,
	  "year": 2019
	},
	{
	  "date": "2019-04-04",
	  "day": 4,
	  "month": 4,
	  "total": 0,
	  "week": 14,
	  "year": 2019
	},
	{
	  "date": "2019-04-05",
	  "day": 5,
	  "month": 4,
	  "total": 0,
	  "week": 14,
	  "year": 2019
	},
	{
	  "date": "2019-04-06",
	  "day": 6,
	  "month": 4,
	  "total": 0,
	  "week": 14,
	  "year": 2019
	},
	{
	  "date": "2019-04-07",
	  "day": 0,
	  "month": 4,
	  "total": 0,
	  "week": 15,
	  "year": 2019
	},
	{
	  "date": "2019-04-08",
	  "day": 1,
	  "month": 4,
	  "total": 0,
	  "week": 15,
	  "year": 2019
	},
	{
	  "date": "2019-04-09",
	  "day": 2,
	  "month": 4,
	  "total": 0,
	  "week": 15,
	  "year": 2019
	},
	{
	  "date": "2019-04-10",
	  "day": 3,
	  "month": 4,
	  "total": 0,
	  "week": 15,
	  "year": 2019
	},
	{
	  "date": "2019-04-11",
	  "day": 4,
	  "month": 4,
	  "total": 0,
	  "week": 15,
	  "year": 2019
	},
	{
	  "date": "2019-04-12",
	  "day": 5,
	  "month": 4,
	  "total": 0,
	  "week": 15,
	  "year": 2019
	},
	{
	  "date": "2019-04-13",
	  "day": 6,
	  "month": 4,
	  "total": 0,
	  "week": 15,
	  "year": 2019
	},
	{
	  "date": "2019-04-14",
	  "day": 0,
	  "month": 4,
	  "total": 0,
	  "week": 16,
	  "year": 2019
	},
	{
	  "date": "2019-04-15",
	  "day": 1,
	  "month": 4,
	  "total": 0,
	  "week": 16,
	  "year": 2019
	},
	{
	  "date": "2019-04-16",
	  "day": 2,
	  "month": 4,
	  "total": 0,
	  "week": 16,
	  "year": 2019
	},
	{
	  "date": "2019-04-17",
	  "day": 3,
	  "month": 4,
	  "total": 0,
	  "week": 16,
	  "year": 2019
	},
	{
	  "date": "2019-04-18",
	  "day": 4,
	  "month": 4,
	  "total": 0,
	  "week": 16,
	  "year": 2019
	},
	{
	  "date": "2019-04-19",
	  "day": 5,
	  "month": 4,
	  "total": 0,
	  "week": 16,
	  "year": 2019
	},
	{
	  "date": "2019-04-20",
	  "day": 6,
	  "month": 4,
	  "total": 0,
	  "week": 16,
	  "year": 2019
	},
	{
	  "date": "2019-04-21",
	  "day": 0,
	  "month": 4,
	  "total": 0,
	  "week": 17,
	  "year": 2019
	},
	{
	  "date": "2019-04-22",
	  "day": 1,
	  "month": 4,
	  "total": 0,
	  "week": 17,
	  "year": 2019
	},
	{
	  "date": "2019-04-23",
	  "day": 2,
	  "month": 4,
	  "total": 0,
	  "week": 17,
	  "year": 2019
	},
	{
	  "date": "2019-04-24",
	  "day": 3,
	  "month": 4,
	  "total": 0,
	  "week": 17,
	  "year": 2019
	},
	{
	  "date": "2019-04-25",
	  "day": 4,
	  "month": 4,
	  "total": 0,
	  "week": 17,
	  "year": 2019
	},
	{
	  "date": "2019-04-26",
	  "day": 5,
	  "month": 4,
	  "total": 0,
	  "week": 17,
	  "year": 2019
	},
	{
	  "date": "2019-04-27",
	  "day": 6,
	  "month": 4,
	  "total": 0,
	  "week": 17,
	  "year": 2019
	},
	{
	  "date": "2019-04-28",
	  "day": 0,
	  "month": 4,
	  "total": 0,
	  "week": 18,
	  "year": 2019
	},
	{
	  "date": "2019-04-29",
	  "day": 1,
	  "month": 4,
	  "total": 0,
	  "week": 18,
	  "year": 2019
	},
	{
	  "date": "2019-04-30",
	  "day": 2,
	  "month": 4,
	  "total": 0,
	  "week": 18,
	  "year": 2019
	},
	{
	  "date": "2019-05-01",
	  "day": 3,
	  "month": 5,
	  "total": 0,
	  "week": 18,
	  "year": 2019
	},
	{
	  "date": "2019-05-02",
	  "day": 4,
	  "month": 5,
	  "total": 0,
	  "week": 18,
	  "year": 2019
	},
	{
	  "date": "2019-05-03",
	  "day": 5,
	  "month": 5,
	  "total": 0,
	  "week": 18,
	  "year": 2019
	},
	{
	  "date": "2019-05-04",
	  "day": 6,
	  "month": 5,
	  "total": 0,
	  "week": 18,
	  "year": 2019
	},
	{
	  "date": "2019-05-05",
	  "day": 0,
	  "month": 5,
	  "total": 0,
	  "week": 19,
	  "year": 2019
	},
	{
	  "date": "2019-05-06",
	  "day": 1,
	  "month": 5,
	  "total": 0,
	  "week": 19,
	  "year": 2019
	},
	{
	  "date": "2019-05-07",
	  "day": 2,
	  "month": 5,
	  "total": 3,
	  "week": 19,
	  "year": 2019
	},
	{
	  "date": "2019-05-08",
	  "day": 3,
	  "month": 5,
	  "total": 3,
	  "week": 19,
	  "year": 2019
	},
	{
	  "date": "2019-05-09",
	  "day": 4,
	  "month": 5,
	  "total": 0,
	  "week": 19,
	  "year": 2019
	},
	{
	  "date": "2019-05-10",
	  "day": 5,
	  "month": 5,
	  "total": 0,
	  "week": 19,
	  "year": 2019
	},
	{
	  "date": "2019-05-11",
	  "day": 6,
	  "month": 5,
	  "total": 0,
	  "week": 19,
	  "year": 2019
	},
	{
	  "date": "2019-05-12",
	  "day": 0,
	  "month": 5,
	  "total": 0,
	  "week": 20,
	  "year": 2019
	},
	{
	  "date": "2019-05-13",
	  "day": 1,
	  "month": 5,
	  "total": 0,
	  "week": 20,
	  "year": 2019
	},
	{
	  "date": "2019-05-14",
	  "day": 2,
	  "month": 5,
	  "total": 0,
	  "week": 20,
	  "year": 2019
	},
	{
	  "date": "2019-05-15",
	  "day": 3,
	  "month": 5,
	  "total": 0,
	  "week": 20,
	  "year": 2019
	},
	{
	  "date": "2019-05-16",
	  "day": 4,
	  "month": 5,
	  "total": 0,
	  "week": 20,
	  "year": 2019
	},
	{
	  "date": "2019-05-17",
	  "day": 5,
	  "month": 5,
	  "total": 0,
	  "week": 20,
	  "year": 2019
	},
	{
	  "date": "2019-05-18",
	  "day": 6,
	  "month": 5,
	  "total": 0,
	  "week": 20,
	  "year": 2019
	},
	{
	  "date": "2019-05-19",
	  "day": 0,
	  "month": 5,
	  "total": 0,
	  "week": 21,
	  "year": 2019
	},
	{
	  "date": "2019-05-20",
	  "day": 1,
	  "month": 5,
	  "total": 0,
	  "week": 21,
	  "year": 2019
	},
	{
	  "date": "2019-05-21",
	  "day": 2,
	  "month": 5,
	  "total": 0,
	  "week": 21,
	  "year": 2019
	},
	{
	  "date": "2019-05-22",
	  "day": 3,
	  "month": 5,
	  "total": 0,
	  "week": 21,
	  "year": 2019
	},
	{
	  "date": "2019-05-23",
	  "day": 4,
	  "month": 5,
	  "total": 0,
	  "week": 21,
	  "year": 2019
	},
	{
	  "date": "2019-05-24",
	  "day": 5,
	  "month": 5,
	  "total": 0,
	  "week": 21,
	  "year": 2019
	},
	{
	  "date": "2019-05-25",
	  "day": 6,
	  "month": 5,
	  "total": 0,
	  "week": 21,
	  "year": 2019
	},
	{
	  "date": "2019-05-26",
	  "day": 0,
	  "month": 5,
	  "total": 0,
	  "week": 22,
	  "year": 2019
	},
	{
	  "date": "2019-05-27",
	  "day": 1,
	  "month": 5,
	  "total": 0,
	  "week": 22,
	  "year": 2019
	},
	{
	  "date": "2019-05-28",
	  "day": 2,
	  "month": 5,
	  "total": 0,
	  "week": 22,
	  "year": 2019
	},
	{
	  "date": "2019-05-29",
	  "day": 3,
	  "month": 5,
	  "total": 0,
	  "week": 22,
	  "year": 2019
	},
	{
	  "date": "2019-05-30",
	  "day": 4,
	  "month": 5,
	  "total": 0,
	  "week": 22,
	  "year": 2019
	},
	{
	  "date": "2019-05-31",
	  "day": 5,
	  "month": 5,
	  "total": 0,
	  "week": 22,
	  "year": 2019
	},
	{
	  "date": "2019-06-01",
	  "day": 6,
	  "month": 6,
	  "total": 0,
	  "week": 22,
	  "year": 2019
	},
	{
	  "date": "2019-06-02",
	  "day": 0,
	  "month": 6,
	  "total": 0,
	  "week": 23,
	  "year": 2019
	},
	{
	  "date": "2019-06-03",
	  "day": 1,
	  "month": 6,
	  "total": 0,
	  "week": 23,
	  "year": 2019
	},
	{
	  "date": "2019-06-04",
	  "day": 2,
	  "month": 6,
	  "total": 1,
	  "week": 23,
	  "year": 2019
	},
	{
	  "date": "2019-06-05",
	  "day": 3,
	  "month": 6,
	  "total": 1,
	  "week": 23,
	  "year": 2019
	},
	{
	  "date": "2019-06-06",
	  "day": 4,
	  "month": 6,
	  "total": 0,
	  "week": 23,
	  "year": 2019
	},
	{
	  "date": "2019-06-07",
	  "day": 5,
	  "month": 6,
	  "total": 0,
	  "week": 23,
	  "year": 2019
	},
	{
	  "date": "2019-06-08",
	  "day": 6,
	  "month": 6,
	  "total": 0,
	  "week": 23,
	  "year": 2019
	},
	{
	  "date": "2019-06-09",
	  "day": 0,
	  "month": 6,
	  "total": 0,
	  "week": 24,
	  "year": 2019
	},
	{
	  "date": "2019-06-10",
	  "day": 1,
	  "month": 6,
	  "total": 0,
	  "week": 24,
	  "year": 2019
	},
	{
	  "date": "2019-06-11",
	  "day": 2,
	  "month": 6,
	  "total": 0,
	  "week": 24,
	  "year": 2019
	},
	{
	  "date": "2019-06-12",
	  "day": 3,
	  "month": 6,
	  "total": 0,
	  "week": 24,
	  "year": 2019
	},
	{
	  "date": "2019-06-13",
	  "day": 4,
	  "month": 6,
	  "total": 3,
	  "week": 24,
	  "year": 2019
	},
	{
	  "date": "2019-06-14",
	  "day": 5,
	  "month": 6,
	  "total": 1,
	  "week": 24,
	  "year": 2019
	},
	{
	  "date": "2019-06-15",
	  "day": 6,
	  "month": 6,
	  "total": 0,
	  "week": 24,
	  "year": 2019
	},
	{
	  "date": "2019-06-16",
	  "day": 0,
	  "month": 6,
	  "total": 0,
	  "week": 25,
	  "year": 2019
	},
	{
	  "date": "2019-06-17",
	  "day": 1,
	  "month": 6,
	  "total": 1,
	  "week": 25,
	  "year": 2019
	},
	{
	  "date": "2019-06-18",
	  "day": 2,
	  "month": 6,
	  "total": 3,
	  "week": 25,
	  "year": 2019
	},
	{
	  "date": "2019-06-19",
	  "day": 3,
	  "month": 6,
	  "total": 4,
	  "week": 25,
	  "year": 2019
	},
	{
	  "date": "2019-06-20",
	  "day": 4,
	  "month": 6,
	  "total": 5,
	  "week": 25,
	  "year": 2019
	},
	{
	  "date": "2019-06-21",
	  "day": 5,
	  "month": 6,
	  "total": 8,
	  "week": 25,
	  "year": 2019
	},
	{
	  "date": "2019-06-22",
	  "day": 6,
	  "month": 6,
	  "total": 0,
	  "week": 25,
	  "year": 2019
	},
	{
	  "date": "2019-06-23",
	  "day": 0,
	  "month": 6,
	  "total": 0,
	  "week": 26,
	  "year": 2019
	},
	{
	  "date": "2019-06-24",
	  "day": 1,
	  "month": 6,
	  "total": 9,
	  "week": 26,
	  "year": 2019
	},
	{
	  "date": "2019-06-25",
	  "day": 2,
	  "month": 6,
	  "total": 7,
	  "week": 26,
	  "year": 2019
	},
	{
	  "date": "2019-06-26",
	  "day": 3,
	  "month": 6,
	  "total": 2,
	  "week": 26,
	  "year": 2019
	},
	{
	  "date": "2019-06-27",
	  "day": 4,
	  "month": 6,
	  "total": 5,
	  "week": 26,
	  "year": 2019
	},
	{
	  "date": "2019-06-28",
	  "day": 5,
	  "month": 6,
	  "total": 0,
	  "week": 26,
	  "year": 2019
	},
	{
	  "date": "2019-06-29",
	  "day": 6,
	  "month": 6,
	  "total": 0,
	  "week": 26,
	  "year": 2019
	},
	{
	  "date": "2019-06-30",
	  "day": 0,
	  "month": 6,
	  "total": 0,
	  "week": 27,
	  "year": 2019
	},
	{
	  "date": "2019-07-01",
	  "day": 1,
	  "month": 7,
	  "total": 29,
	  "week": 27,
	  "year": 2019
	},
	{
	  "date": "2019-07-02",
	  "day": 2,
	  "month": 7,
	  "total": 9,
	  "week": 27,
	  "year": 2019
	},
	{
	  "date": "2019-07-03",
	  "day": 3,
	  "month": 7,
	  "total": 11,
	  "week": 27,
	  "year": 2019
	},
	{
	  "date": "2019-07-04",
	  "day": 4,
	  "month": 7,
	  "total": 5,
	  "week": 27,
	  "year": 2019
	},
	{
	  "date": "2019-07-05",
	  "day": 5,
	  "month": 7,
	  "total": 3,
	  "week": 27,
	  "year": 2019
	},
	{
	  "date": "2019-07-06",
	  "day": 6,
	  "month": 7,
	  "total": 0,
	  "week": 27,
	  "year": 2019
	},
	{
	  "date": "2019-07-07",
	  "day": 0,
	  "month": 7,
	  "total": 0,
	  "week": 28,
	  "year": 2019
	},
	{
	  "date": "2019-07-08",
	  "day": 1,
	  "month": 7,
	  "total": 13,
	  "week": 28,
	  "year": 2019
	},
	{
	  "date": "2019-07-09",
	  "day": 2,
	  "month": 7,
	  "total": 9,
	  "week": 28,
	  "year": 2019
	},
	{
	  "date": "2019-07-10",
	  "day": 3,
	  "month": 7,
	  "total": 17,
	  "week": 28,
	  "year": 2019
	},
	{
	  "date": "2019-07-11",
	  "day": 4,
	  "month": 7,
	  "total": 12,
	  "week": 28,
	  "year": 2019
	},
	{
	  "date": "2019-07-12",
	  "day": 5,
	  "month": 7,
	  "total": 6,
	  "week": 28,
	  "year": 2019
	},
	{
	  "date": "2019-07-13",
	  "day": 6,
	  "month": 7,
	  "total": 0,
	  "week": 28,
	  "year": 2019
	},
	{
	  "date": "2019-07-14",
	  "day": 0,
	  "month": 7,
	  "total": 6,
	  "week": 29,
	  "year": 2019
	},
	{
	  "date": "2019-07-15",
	  "day": 1,
	  "month": 7,
	  "total": 16,
	  "week": 29,
	  "year": 2019
	},
	{
	  "date": "2019-07-16",
	  "day": 2,
	  "month": 7,
	  "total": 3,
	  "week": 29,
	  "year": 2019
	},
	{
	  "date": "2019-07-17",
	  "day": 3,
	  "month": 7,
	  "total": 13,
	  "week": 29,
	  "year": 2019
	},
	{
	  "date": "2019-07-18",
	  "day": 4,
	  "month": 7,
	  "total": 25,
	  "week": 29,
	  "year": 2019
	},
	{
	  "date": "2019-07-19",
	  "day": 5,
	  "month": 7,
	  "total": 12,
	  "week": 29,
	  "year": 2019
	},
	{
	  "date": "2019-07-20",
	  "day": 6,
	  "month": 7,
	  "total": 0,
	  "week": 29,
	  "year": 2019
	},
	{
	  "date": "2019-07-21",
	  "day": 0,
	  "month": 7,
	  "total": 0,
	  "week": 30,
	  "year": 2019
	},
	{
	  "date": "2019-07-22",
	  "day": 1,
	  "month": 7,
	  "total": 7,
	  "week": 30,
	  "year": 2019
	},
	{
	  "date": "2019-07-23",
	  "day": 2,
	  "month": 7,
	  "total": 26,
	  "week": 30,
	  "year": 2019
	},
	{
	  "date": "2019-07-24",
	  "day": 3,
	  "month": 7,
	  "total": 16,
	  "week": 30,
	  "year": 2019
	},
	{
	  "date": "2019-07-25",
	  "day": 4,
	  "month": 7,
	  "total": 24,
	  "week": 30,
	  "year": 2019
	},
	{
	  "date": "2019-07-26",
	  "day": 5,
	  "month": 7,
	  "total": 6,
	  "week": 30,
	  "year": 2019
	},
	{
	  "date": "2019-07-27",
	  "day": 6,
	  "month": 7,
	  "total": 0,
	  "week": 30,
	  "year": 2019
	},
	{
	  "date": "2019-07-28",
	  "day": 0,
	  "month": 7,
	  "total": 0,
	  "week": 31,
	  "year": 2019
	},
	{
	  "date": "2019-07-29",
	  "day": 1,
	  "month": 7,
	  "total": 0,
	  "week": 31,
	  "year": 2019
	},
	{
	  "date": "2019-07-30",
	  "day": 2,
	  "month": 7,
	  "total": 0,
	  "week": 31,
	  "year": 2019
	},
	{
	  "date": "2019-07-31",
	  "day": 3,
	  "month": 7,
	  "total": 0,
	  "week": 31,
	  "year": 2019
	},
	{
	  "date": "2019-08-01",
	  "day": 4,
	  "month": 8,
	  "total": 0,
	  "week": 31,
	  "year": 2019
	},
	{
	  "date": "2019-08-02",
	  "day": 5,
	  "month": 8,
	  "total": 0,
	  "week": 31,
	  "year": 2019
	},
	{
	  "date": "2019-08-03",
	  "day": 6,
	  "month": 8,
	  "total": 1,
	  "week": 31,
	  "year": 2019
	},
	{
	  "date": "2019-08-04",
	  "day": 0,
	  "month": 8,
	  "total": 0,
	  "week": 32,
	  "year": 2019
	},
	{
	  "date": "2019-08-05",
	  "day": 1,
	  "month": 8,
	  "total": 0,
	  "week": 32,
	  "year": 2019
	},
	{
	  "date": "2019-08-06",
	  "day": 2,
	  "month": 8,
	  "total": 0,
	  "week": 32,
	  "year": 2019
	},
	{
	  "date": "2019-08-07",
	  "day": 3,
	  "month": 8,
	  "total": 0,
	  "week": 32,
	  "year": 2019
	},
	{
	  "date": "2019-08-08",
	  "day": 4,
	  "month": 8,
	  "total": 9,
	  "week": 32,
	  "year": 2019
	},
	{
	  "date": "2019-08-09",
	  "day": 5,
	  "month": 8,
	  "total": 12,
	  "week": 32,
	  "year": 2019
	},
	{
	  "date": "2019-08-10",
	  "day": 6,
	  "month": 8,
	  "total": 0,
	  "week": 32,
	  "year": 2019
	},
	{
	  "date": "2019-08-11",
	  "day": 0,
	  "month": 8,
	  "total": 0,
	  "week": 33,
	  "year": 2019
	},
	{
	  "date": "2019-08-12",
	  "day": 1,
	  "month": 8,
	  "total": 10,
	  "week": 33,
	  "year": 2019
	},
	{
	  "date": "2019-08-13",
	  "day": 2,
	  "month": 8,
	  "total": 10,
	  "week": 33,
	  "year": 2019
	},
	{
	  "date": "2019-08-14",
	  "day": 3,
	  "month": 8,
	  "total": 3,
	  "week": 33,
	  "year": 2019
	},
	{
	  "date": "2019-08-15",
	  "day": 4,
	  "month": 8,
	  "total": 9,
	  "week": 33,
	  "year": 2019
	},
	{
	  "date": "2019-08-16",
	  "day": 5,
	  "month": 8,
	  "total": 9,
	  "week": 33,
	  "year": 2019
	},
	{
	  "date": "2019-08-17",
	  "day": 6,
	  "month": 8,
	  "total": 0,
	  "week": 33,
	  "year": 2019
	},
	{
	  "date": "2019-08-18",
	  "day": 0,
	  "month": 8,
	  "total": 0,
	  "week": 34,
	  "year": 2019
	},
	{
	  "date": "2019-08-19",
	  "day": 1,
	  "month": 8,
	  "total": 3,
	  "week": 34,
	  "year": 2019
	},
	{
	  "date": "2019-08-20",
	  "day": 2,
	  "month": 8,
	  "total": 4,
	  "week": 34,
	  "year": 2019
	},
	{
	  "date": "2019-08-21",
	  "day": 3,
	  "month": 8,
	  "total": 6,
	  "week": 34,
	  "year": 2019
	},
	{
	  "date": "2019-08-22",
	  "day": 4,
	  "month": 8,
	  "total": 20,
	  "week": 34,
	  "year": 2019
	},
	{
	  "date": "2019-08-23",
	  "day": 5,
	  "month": 8,
	  "total": 13,
	  "week": 34,
	  "year": 2019
	},
	{
	  "date": "2019-08-24",
	  "day": 6,
	  "month": 8,
	  "total": 0,
	  "week": 34,
	  "year": 2019
	},
	{
	  "date": "2019-08-25",
	  "day": 0,
	  "month": 8,
	  "total": 0,
	  "week": 35,
	  "year": 2019
	},
	{
	  "date": "2019-08-26",
	  "day": 1,
	  "month": 8,
	  "total": 10,
	  "week": 35,
	  "year": 2019
	},
	{
	  "date": "2019-08-27",
	  "day": 2,
	  "month": 8,
	  "total": 11,
	  "week": 35,
	  "year": 2019
	},
	{
	  "date": "2019-08-28",
	  "day": 3,
	  "month": 8,
	  "total": 13,
	  "week": 35,
	  "year": 2019
	},
	{
	  "date": "2019-08-29",
	  "day": 4,
	  "month": 8,
	  "total": 8,
	  "week": 35,
	  "year": 2019
	},
	{
	  "date": "2019-08-30",
	  "day": 5,
	  "month": 8,
	  "total": 10,
	  "week": 35,
	  "year": 2019
	},
	{
	  "date": "2019-08-31",
	  "day": 6,
	  "month": 8,
	  "total": 2,
	  "week": 35,
	  "year": 2019
	},
	{
	  "date": "2019-09-01",
	  "day": 0,
	  "month": 9,
	  "total": 0,
	  "week": 36,
	  "year": 2019
	},
	{
	  "date": "2019-09-02",
	  "day": 1,
	  "month": 9,
	  "total": 2,
	  "week": 36,
	  "year": 2019
	},
	{
	  "date": "2019-09-03",
	  "day": 2,
	  "month": 9,
	  "total": 0,
	  "week": 36,
	  "year": 2019
	},
	{
	  "date": "2019-09-04",
	  "day": 3,
	  "month": 9,
	  "total": 8,
	  "week": 36,
	  "year": 2019
	},
	{
	  "date": "2019-09-05",
	  "day": 4,
	  "month": 9,
	  "total": 0,
	  "week": 36,
	  "year": 2019
	},
	{
	  "date": "2019-09-06",
	  "day": 5,
	  "month": 9,
	  "total": 0,
	  "week": 36,
	  "year": 2019
	},
	{
	  "date": "2019-09-07",
	  "day": 6,
	  "month": 9,
	  "total": 0,
	  "week": 36,
	  "year": 2019
	},
	{
	  "date": "2019-09-08",
	  "day": 0,
	  "month": 9,
	  "total": 0,
	  "week": 37,
	  "year": 2019
	},
	{
	  "date": "2019-09-09",
	  "day": 1,
	  "month": 9,
	  "total": 8,
	  "week": 37,
	  "year": 2019
	},
	{
	  "date": "2019-09-10",
	  "day": 2,
	  "month": 9,
	  "total": 0,
	  "week": 37,
	  "year": 2019
	},
	{
	  "date": "2019-09-11",
	  "day": 3,
	  "month": 9,
	  "total": 18,
	  "week": 37,
	  "year": 2019
	},
	{
	  "date": "2019-09-12",
	  "day": 4,
	  "month": 9,
	  "total": 3,
	  "week": 37,
	  "year": 2019
	},
	{
	  "date": "2019-09-13",
	  "day": 5,
	  "month": 9,
	  "total": 0,
	  "week": 37,
	  "year": 2019
	},
	{
	  "date": "2019-09-14",
	  "day": 6,
	  "month": 9,
	  "total": 0,
	  "week": 37,
	  "year": 2019
	},
	{
	  "date": "2019-09-15",
	  "day": 0,
	  "month": 9,
	  "total": 0,
	  "week": 38,
	  "year": 2019
	},
	{
	  "date": "2019-09-16",
	  "day": 1,
	  "month": 9,
	  "total": 10,
	  "week": 38,
	  "year": 2019
	},
	{
	  "date": "2019-09-17",
	  "day": 2,
	  "month": 9,
	  "total": 36,
	  "week": 38,
	  "year": 2019
	},
	{
	  "date": "2019-09-18",
	  "day": 3,
	  "month": 9,
	  "total": 32,
	  "week": 38,
	  "year": 2019
	},
	{
	  "date": "2019-09-19",
	  "day": 4,
	  "month": 9,
	  "total": 2,
	  "week": 38,
	  "year": 2019
	},
	{
	  "date": "2019-09-20",
	  "day": 5,
	  "month": 9,
	  "total": 14,
	  "week": 38,
	  "year": 2019
	},
	{
	  "date": "2019-09-21",
	  "day": 6,
	  "month": 9,
	  "total": 0,
	  "week": 38,
	  "year": 2019
	},
	{
	  "date": "2019-09-22",
	  "day": 0,
	  "month": 9,
	  "total": 0,
	  "week": 39,
	  "year": 2019
	},
	{
	  "date": "2019-09-23",
	  "day": 1,
	  "month": 9,
	  "total": 10,
	  "week": 39,
	  "year": 2019
	},
	{
	  "date": "2019-09-24",
	  "day": 2,
	  "month": 9,
	  "total": 11,
	  "week": 39,
	  "year": 2019
	},
	{
	  "date": "2019-09-25",
	  "day": 3,
	  "month": 9,
	  "total": 11,
	  "week": 39,
	  "year": 2019
	},
	{
	  "date": "2019-09-26",
	  "day": 4,
	  "month": 9,
	  "total": 12,
	  "week": 39,
	  "year": 2019
	},
	{
	  "date": "2019-09-27",
	  "day": 5,
	  "month": 9,
	  "total": 3,
	  "week": 39,
	  "year": 2019
	},
	{
	  "date": "2019-09-28",
	  "day": 6,
	  "month": 9,
	  "total": 0,
	  "week": 39,
	  "year": 2019
	},
	{
	  "date": "2019-09-29",
	  "day": 0,
	  "month": 9,
	  "total": 16,
	  "week": 40,
	  "year": 2019
	},
	{
	  "date": "2019-09-30",
	  "day": 1,
	  "month": 9,
	  "total": 0,
	  "week": 40,
	  "year": 2019
	},
	{
	  "date": "2019-10-01",
	  "day": 2,
	  "month": 10,
	  "total": 0,
	  "week": 40,
	  "year": 2019
	},
	{
	  "date": "2019-10-02",
	  "day": 3,
	  "month": 10,
	  "total": 0,
	  "week": 40,
	  "year": 2019
	},
	{
	  "date": "2019-10-03",
	  "day": 4,
	  "month": 10,
	  "total": 0,
	  "week": 40,
	  "year": 2019
	},
	{
	  "date": "2019-10-04",
	  "day": 5,
	  "month": 10,
	  "total": 0,
	  "week": 40,
	  "year": 2019
	},
	{
	  "date": "2019-10-05",
	  "day": 6,
	  "month": 10,
	  "total": 0,
	  "week": 40,
	  "year": 2019
	},
	{
	  "date": "2019-10-06",
	  "day": 0,
	  "month": 10,
	  "total": 0,
	  "week": 41,
	  "year": 2019
	},
	{
	  "date": "2019-10-07",
	  "day": 1,
	  "month": 10,
	  "total": 0,
	  "week": 41,
	  "year": 2019
	},
	{
	  "date": "2019-10-08",
	  "day": 2,
	  "month": 10,
	  "total": 10,
	  "week": 41,
	  "year": 2019
	},
	{
	  "date": "2019-10-09",
	  "day": 3,
	  "month": 10,
	  "total": 20,
	  "week": 41,
	  "year": 2019
	},
	{
	  "date": "2019-10-10",
	  "day": 4,
	  "month": 10,
	  "total": 19,
	  "week": 41,
	  "year": 2019
	},
	{
	  "date": "2019-10-11",
	  "day": 5,
	  "month": 10,
	  "total": 32,
	  "week": 41,
	  "year": 2019
	},
	{
	  "date": "2019-10-12",
	  "day": 6,
	  "month": 10,
	  "total": 23,
	  "week": 41,
	  "year": 2019
	},
	{
	  "date": "2019-10-13",
	  "day": 0,
	  "month": 10,
	  "total": 0,
	  "week": 42,
	  "year": 2019
	},
	{
	  "date": "2019-10-14",
	  "day": 1,
	  "month": 10,
	  "total": 4,
	  "week": 42,
	  "year": 2019
	},
	{
	  "date": "2019-10-15",
	  "day": 2,
	  "month": 10,
	  "total": 1,
	  "week": 42,
	  "year": 2019
	},
	{
	  "date": "2019-10-16",
	  "day": 3,
	  "month": 10,
	  "total": 5,
	  "week": 42,
	  "year": 2019
	},
	{
	  "date": "2019-10-17",
	  "day": 4,
	  "month": 10,
	  "total": 23,
	  "week": 42,
	  "year": 2019
	},
	{
	  "date": "2019-10-18",
	  "day": 5,
	  "month": 10,
	  "total": 3,
	  "week": 42,
	  "year": 2019
	},
	{
	  "date": "2019-10-19",
	  "day": 6,
	  "month": 10,
	  "total": 0,
	  "week": 42,
	  "year": 2019
	},
	{
	  "date": "2019-10-20",
	  "day": 0,
	  "month": 10,
	  "total": 0,
	  "week": 43,
	  "year": 2019
	},
	{
	  "date": "2019-10-21",
	  "day": 1,
	  "month": 10,
	  "total": 8,
	  "week": 43,
	  "year": 2019
	},
	{
	  "date": "2019-10-22",
	  "day": 2,
	  "month": 10,
	  "total": 10,
	  "week": 43,
	  "year": 2019
	},
	{
	  "date": "2019-10-23",
	  "day": 3,
	  "month": 10,
	  "total": 11,
	  "week": 43,
	  "year": 2019
	},
	{
	  "date": "2019-10-24",
	  "day": 4,
	  "month": 10,
	  "total": 17,
	  "week": 43,
	  "year": 2019
	},
	{
	  "date": "2019-10-25",
	  "day": 5,
	  "month": 10,
	  "total": 17,
	  "week": 43,
	  "year": 2019
	},
	{
	  "date": "2019-10-26",
	  "day": 6,
	  "month": 10,
	  "total": 0,
	  "week": 43,
	  "year": 2019
	},
	{
	  "date": "2019-10-27",
	  "day": 0,
	  "month": 10,
	  "total": 0,
	  "week": 44,
	  "year": 2019
	},
	{
	  "date": "2019-10-28",
	  "day": 1,
	  "month": 10,
	  "total": 8,
	  "week": 44,
	  "year": 2019
	},
	{
	  "date": "2019-10-29",
	  "day": 2,
	  "month": 10,
	  "total": 10,
	  "week": 44,
	  "year": 2019
	},
	{
	  "date": "2019-10-30",
	  "day": 3,
	  "month": 10,
	  "total": 28,
	  "week": 44,
	  "year": 2019
	},
	{
	  "date": "2019-10-31",
	  "day": 4,
	  "month": 10,
	  "total": 10,
	  "week": 44,
	  "year": 2019
	},
	{
	  "date": "2019-11-01",
	  "day": 5,
	  "month": 11,
	  "total": 20,
	  "week": 44,
	  "year": 2019
	},
	{
	  "date": "2019-11-02",
	  "day": 6,
	  "month": 11,
	  "total": 0,
	  "week": 44,
	  "year": 2019
	},
	{
	  "date": "2019-11-03",
	  "day": 0,
	  "month": 11,
	  "total": 0,
	  "week": 45,
	  "year": 2019
	},
	{
	  "date": "2019-11-04",
	  "day": 1,
	  "month": 11,
	  "total": 8,
	  "week": 45,
	  "year": 2019
	},
	{
	  "date": "2019-11-05",
	  "day": 2,
	  "month": 11,
	  "total": 3,
	  "week": 45,
	  "year": 2019
	},
	{
	  "date": "2019-11-06",
	  "day": 3,
	  "month": 11,
	  "total": 0,
	  "week": 45,
	  "year": 2019
	},
	{
	  "date": "2019-11-07",
	  "day": 4,
	  "month": 11,
	  "total": 0,
	  "week": 45,
	  "year": 2019
	},
	{
	  "date": "2019-11-08",
	  "day": 5,
	  "month": 11,
	  "total": 7,
	  "week": 45,
	  "year": 2019
	},
	{
	  "date": "2019-11-09",
	  "day": 6,
	  "month": 11,
	  "total": 0,
	  "week": 45,
	  "year": 2019
	},
	{
	  "date": "2019-11-10",
	  "day": 0,
	  "month": 11,
	  "total": 0,
	  "week": 46,
	  "year": 2019
	},
	{
	  "date": "2019-11-11",
	  "day": 1,
	  "month": 11,
	  "total": 1,
	  "week": 46,
	  "year": 2019
	},
	{
	  "date": "2019-11-12",
	  "day": 2,
	  "month": 11,
	  "total": 38,
	  "week": 46,
	  "year": 2019
	},
	{
	  "date": "2019-11-13",
	  "day": 3,
	  "month": 11,
	  "total": 15,
	  "week": 46,
	  "year": 2019
	},
	{
	  "date": "2019-11-14",
	  "day": 4,
	  "month": 11,
	  "total": 4,
	  "week": 46,
	  "year": 2019
	},
	{
	  "date": "2019-11-15",
	  "day": 5,
	  "month": 11,
	  "total": 8,
	  "week": 46,
	  "year": 2019
	},
	{
	  "date": "2019-11-16",
	  "day": 6,
	  "month": 11,
	  "total": 0,
	  "week": 46,
	  "year": 2019
	},
	{
	  "date": "2019-11-17",
	  "day": 0,
	  "month": 11,
	  "total": 0,
	  "week": 47,
	  "year": 2019
	},
	{
	  "date": "2019-11-18",
	  "day": 1,
	  "month": 11,
	  "total": 19,
	  "week": 47,
	  "year": 2019
	},
	{
	  "date": "2019-11-19",
	  "day": 2,
	  "month": 11,
	  "total": 28,
	  "week": 47,
	  "year": 2019
	},
	{
	  "date": "2019-11-20",
	  "day": 3,
	  "month": 11,
	  "total": 9,
	  "week": 47,
	  "year": 2019
	},
	{
	  "date": "2019-11-21",
	  "day": 4,
	  "month": 11,
	  "total": 3,
	  "week": 47,
	  "year": 2019
	},
	{
	  "date": "2019-11-22",
	  "day": 5,
	  "month": 11,
	  "total": 4,
	  "week": 47,
	  "year": 2019
	},
	{
	  "date": "2019-11-23",
	  "day": 6,
	  "month": 11,
	  "total": 0,
	  "week": 47,
	  "year": 2019
	},
	{
	  "date": "2019-11-24",
	  "day": 0,
	  "month": 11,
	  "total": 9,
	  "week": 48,
	  "year": 2019
	},
	{
	  "date": "2019-11-25",
	  "day": 1,
	  "month": 11,
	  "total": 4,
	  "week": 48,
	  "year": 2019
	},
	{
	  "date": "2019-11-26",
	  "day": 2,
	  "month": 11,
	  "total": 0,
	  "week": 48,
	  "year": 2019
	},
	{
	  "date": "2019-11-27",
	  "day": 3,
	  "month": 11,
	  "total": 3,
	  "week": 48,
	  "year": 2019
	},
	{
	  "date": "2019-11-28",
	  "day": 4,
	  "month": 11,
	  "total": 2,
	  "week": 48,
	  "year": 2019
	},
	{
	  "date": "2019-11-29",
	  "day": 5,
	  "month": 11,
	  "total": 0,
	  "week": 48,
	  "year": 2019
	},
	{
	  "date": "2019-11-30",
	  "day": 6,
	  "month": 11,
	  "total": 0,
	  "week": 48,
	  "year": 2019
	},
	{
	  "date": "2019-12-01",
	  "day": 0,
	  "month": 12,
	  "total": 2,
	  "week": 49,
	  "year": 2019
	},
	{
	  "date": "2019-12-02",
	  "day": 1,
	  "month": 12,
	  "total": 3,
	  "week": 49,
	  "year": 2019
	},
	{
	  "date": "2019-12-03",
	  "day": 2,
	  "month": 12,
	  "total": 4,
	  "week": 49,
	  "year": 2019
	},
	{
	  "date": "2019-12-04",
	  "day": 3,
	  "month": 12,
	  "total": 0,
	  "week": 49,
	  "year": 2019
	},
	{
	  "date": "2019-12-05",
	  "day": 4,
	  "month": 12,
	  "total": 8,
	  "week": 49,
	  "year": 2019
	},
	{
	  "date": "2019-12-06",
	  "day": 5,
	  "month": 12,
	  "total": 8,
	  "week": 49,
	  "year": 2019
	},
	{
	  "date": "2019-12-07",
	  "day": 6,
	  "month": 12,
	  "total": 0,
	  "week": 49,
	  "year": 2019
	},
	{
	  "date": "2019-12-08",
	  "day": 0,
	  "month": 12,
	  "total": 0,
	  "week": 50,
	  "year": 2019
	},
	{
	  "date": "2019-12-09",
	  "day": 1,
	  "month": 12,
	  "total": 11,
	  "week": 50,
	  "year": 2019
	},
	{
	  "date": "2019-12-10",
	  "day": 2,
	  "month": 12,
	  "total": 5,
	  "week": 50,
	  "year": 2019
	},
	{
	  "date": "2019-12-11",
	  "day": 3,
	  "month": 12,
	  "total": 15,
	  "week": 50,
	  "year": 2019
	},
	{
	  "date": "2019-12-12",
	  "day": 4,
	  "month": 12,
	  "total": 9,
	  "week": 50,
	  "year": 2019
	},
	{
	  "date": "2019-12-13",
	  "day": 5,
	  "month": 12,
	  "total": 9,
	  "week": 50,
	  "year": 2019
	},
	{
	  "date": "2019-12-14",
	  "day": 6,
	  "month": 12,
	  "total": 2,
	  "week": 50,
	  "year": 2019
	},
	{
	  "date": "2019-12-15",
	  "day": 0,
	  "month": 12,
	  "total": 0,
	  "week": 51,
	  "year": 2019
	},
	{
	  "date": "2019-12-16",
	  "day": 1,
	  "month": 12,
	  "total": 8,
	  "week": 51,
	  "year": 2019
	},
	{
	  "date": "2019-12-17",
	  "day": 2,
	  "month": 12,
	  "total": 15,
	  "week": 51,
	  "year": 2019
	},
	{
	  "date": "2019-12-18",
	  "day": 3,
	  "month": 12,
	  "total": 3,
	  "week": 51,
	  "year": 2019
	},
	{
	  "date": "2019-12-19",
	  "day": 4,
	  "month": 12,
	  "total": 5,
	  "week": 51,
	  "year": 2019
	},
	{
	  "date": "2019-12-20",
	  "day": 5,
	  "month": 12,
	  "total": 28,
	  "week": 51,
	  "year": 2019
	},
	{
	  "date": "2019-12-21",
	  "day": 6,
	  "month": 12,
	  "total": 2,
	  "week": 51,
	  "year": 2019
	},
	{
	  "date": "2019-12-22",
	  "day": 0,
	  "month": 12,
	  "total": 0,
	  "week": 52,
	  "year": 2019
	},
	{
	  "date": "2019-12-23",
	  "day": 1,
	  "month": 12,
	  "total": 1,
	  "week": 52,
	  "year": 2019
	},
	{
	  "date": "2019-12-24",
	  "day": 2,
	  "month": 12,
	  "total": 8,
	  "week": 52,
	  "year": 2019
	},
	{
	  "date": "2019-12-25",
	  "day": 3,
	  "month": 12,
	  "total": 2,
	  "week": 52,
	  "year": 2019
	},
	{
	  "date": "2019-12-26",
	  "day": 4,
	  "month": 12,
	  "total": 16,
	  "week": 52,
	  "year": 2019
	},
	{
	  "date": "2019-12-27",
	  "day": 5,
	  "month": 12,
	  "total": 12,
	  "week": 52,
	  "year": 2019
	},
	{
	  "date": "2019-12-28",
	  "day": 6,
	  "month": 12,
	  "total": 0,
	  "week": 52,
	  "year": 2019
	},
	{
	  "date": "2019-12-29",
	  "day": 0,
	  "month": 12,
	  "total": 0,
	  "week": 53,
	  "year": 2019
	},
	{
	  "date": "2019-12-30",
	  "day": 1,
	  "month": 12,
	  "total": 7,
	  "week": 53,
	  "year": 2019
	},
	{
	  "date": "2019-12-31",
	  "day": 2,
	  "month": 12,
	  "total": 7,
	  "week": 53,
	  "year": 2019
	},
	{
	  "date": "2020-01-01",
	  "day": 3,
	  "month": 1,
	  "total": 7,
	  "week": 1,
	  "year": 2020
	},
	{
	  "date": "2020-01-02",
	  "day": 4,
	  "month": 1,
	  "total": 7,
	  "week": 1,
	  "year": 2020
	},
	{
	  "date": "2020-01-03",
	  "day": 5,
	  "month": 1,
	  "total": 16,
	  "week": 1,
	  "year": 2020
	},
	{
	  "date": "2020-01-04",
	  "day": 6,
	  "month": 1,
	  "total": 0,
	  "week": 1,
	  "year": 2020
	},
	{
	  "date": "2020-01-05",
	  "day": 0,
	  "month": 1,
	  "total": 0,
	  "week": 2,
	  "year": 2020
	},
	{
	  "date": "2020-01-06",
	  "day": 1,
	  "month": 1,
	  "total": 3,
	  "week": 2,
	  "year": 2020
	},
	{
	  "date": "2020-01-07",
	  "day": 2,
	  "month": 1,
	  "total": 25,
	  "week": 2,
	  "year": 2020
	},
	{
	  "date": "2020-01-08",
	  "day": 3,
	  "month": 1,
	  "total": 8,
	  "week": 2,
	  "year": 2020
	},
	{
	  "date": "2020-01-09",
	  "day": 4,
	  "month": 1,
	  "total": 12,
	  "week": 2,
	  "year": 2020
	},
	{
	  "date": "2020-01-10",
	  "day": 5,
	  "month": 1,
	  "total": 13,
	  "week": 2,
	  "year": 2020
	},
	{
	  "date": "2020-01-11",
	  "day": 6,
	  "month": 1,
	  "total": 0,
	  "week": 2,
	  "year": 2020
	},
	{
	  "date": "2020-01-12",
	  "day": 0,
	  "month": 1,
	  "total": 0,
	  "week": 3,
	  "year": 2020
	},
	{
	  "date": "2020-01-13",
	  "day": 1,
	  "month": 1,
	  "total": 0,
	  "week": 3,
	  "year": 2020
	},
	{
	  "date": "2020-01-14",
	  "day": 2,
	  "month": 1,
	  "total": 0,
	  "week": 3,
	  "year": 2020
	},
	{
	  "date": "2020-01-15",
	  "day": 3,
	  "month": 1,
	  "total": 6,
	  "week": 3,
	  "year": 2020
	},
	{
	  "date": "2020-01-16",
	  "day": 4,
	  "month": 1,
	  "total": 7,
	  "week": 3,
	  "year": 2020
	},
	{
	  "date": "2020-01-17",
	  "day": 5,
	  "month": 1,
	  "total": 1,
	  "week": 3,
	  "year": 2020
	},
	{
	  "date": "2020-01-18",
	  "day": 6,
	  "month": 1,
	  "total": 0,
	  "week": 3,
	  "year": 2020
	},
	{
	  "date": "2020-01-19",
	  "day": 0,
	  "month": 1,
	  "total": 0,
	  "week": 4,
	  "year": 2020
	},
	{
	  "date": "2020-01-20",
	  "day": 1,
	  "month": 1,
	  "total": 0,
	  "week": 4,
	  "year": 2020
	},
	{
	  "date": "2020-01-21",
	  "day": 2,
	  "month": 1,
	  "total": 0,
	  "week": 4,
	  "year": 2020
	},
	{
	  "date": "2020-01-22",
	  "day": 3,
	  "month": 1,
	  "total": 0,
	  "week": 4,
	  "year": 2020
	},
	{
	  "date": "2020-01-23",
	  "day": 4,
	  "month": 1,
	  "total": 0,
	  "week": 4,
	  "year": 2020
	},
	{
	  "date": "2020-01-24",
	  "day": 5,
	  "month": 1,
	  "total": 0,
	  "week": 4,
	  "year": 2020
	},
	{
	  "date": "2020-01-25",
	  "day": 6,
	  "month": 1,
	  "total": 0,
	  "week": 4,
	  "year": 2020
	},
	{
	  "date": "2020-01-26",
	  "day": 0,
	  "month": 1,
	  "total": 0,
	  "week": 5,
	  "year": 2020
	},
	{
	  "date": "2020-01-27",
	  "day": 1,
	  "month": 1,
	  "total": 0,
	  "week": 5,
	  "year": 2020
	},
	{
	  "date": "2020-01-28",
	  "day": 2,
	  "month": 1,
	  "total": 0,
	  "week": 5,
	  "year": 2020
	},
	{
	  "date": "2020-01-29",
	  "day": 3,
	  "month": 1,
	  "total": 0,
	  "week": 5,
	  "year": 2020
	},
	{
	  "date": "2020-01-30",
	  "day": 4,
	  "month": 1,
	  "total": 0,
	  "week": 5,
	  "year": 2020
	},
	{
	  "date": "2020-01-31",
	  "day": 5,
	  "month": 1,
	  "total": 0,
	  "week": 5,
	  "year": 2020
	},
	{
	  "date": "2020-02-01",
	  "day": 6,
	  "month": 2,
	  "total": 0,
	  "week": 5,
	  "year": 2020
	},
	{
	  "date": "2020-02-02",
	  "day": 0,
	  "month": 2,
	  "total": 3,
	  "week": 6,
	  "year": 2020
	},
	{
	  "date": "2020-02-03",
	  "day": 1,
	  "month": 2,
	  "total": 3,
	  "week": 6,
	  "year": 2020
	},
	{
	  "date": "2020-02-04",
	  "day": 2,
	  "month": 2,
	  "total": 8,
	  "week": 6,
	  "year": 2020
	},
	{
	  "date": "2020-02-05",
	  "day": 3,
	  "month": 2,
	  "total": 7,
	  "week": 6,
	  "year": 2020
	},
	{
	  "date": "2020-02-06",
	  "day": 4,
	  "month": 2,
	  "total": 7,
	  "week": 6,
	  "year": 2020
	},
	{
	  "date": "2020-02-07",
	  "day": 5,
	  "month": 2,
	  "total": 4,
	  "week": 6,
	  "year": 2020
	},
	{
	  "date": "2020-02-08",
	  "day": 6,
	  "month": 2,
	  "total": 0,
	  "week": 6,
	  "year": 2020
	},
	{
	  "date": "2020-02-09",
	  "day": 0,
	  "month": 2,
	  "total": 4,
	  "week": 7,
	  "year": 2020
	},
	{
	  "date": "2020-02-10",
	  "day": 1,
	  "month": 2,
	  "total": 16,
	  "week": 7,
	  "year": 2020
	},
	{
	  "date": "2020-02-11",
	  "day": 2,
	  "month": 2,
	  "total": 6,
	  "week": 7,
	  "year": 2020
	},
	{
	  "date": "2020-02-12",
	  "day": 3,
	  "month": 2,
	  "total": 17,
	  "week": 7,
	  "year": 2020
	},
	{
	  "date": "2020-02-13",
	  "day": 4,
	  "month": 2,
	  "total": 1,
	  "week": 7,
	  "year": 2020
	},
	{
	  "date": "2020-02-14",
	  "day": 5,
	  "month": 2,
	  "total": 4,
	  "week": 7,
	  "year": 2020
	},
	{
	  "date": "2020-02-15",
	  "day": 6,
	  "month": 2,
	  "total": 0,
	  "week": 7,
	  "year": 2020
	},
	{
	  "date": "2020-02-16",
	  "day": 0,
	  "month": 2,
	  "total": 0,
	  "week": 8,
	  "year": 2020
	},
	{
	  "date": "2020-02-17",
	  "day": 1,
	  "month": 2,
	  "total": 20,
	  "week": 8,
	  "year": 2020
	},
	{
	  "date": "2020-02-18",
	  "day": 2,
	  "month": 2,
	  "total": 8,
	  "week": 8,
	  "year": 2020
	},
	{
	  "date": "2020-02-19",
	  "day": 3,
	  "month": 2,
	  "total": 2,
	  "week": 8,
	  "year": 2020
	},
	{
	  "date": "2020-02-20",
	  "day": 4,
	  "month": 2,
	  "total": 0,
	  "week": 8,
	  "year": 2020
	},
	{
	  "date": "2020-02-21",
	  "day": 5,
	  "month": 2,
	  "total": 8,
	  "week": 8,
	  "year": 2020
	},
	{
	  "date": "2020-02-22",
	  "day": 6,
	  "month": 2,
	  "total": 0,
	  "week": 8,
	  "year": 2020
	},
	{
	  "date": "2020-02-23",
	  "day": 0,
	  "month": 2,
	  "total": 0,
	  "week": 9,
	  "year": 2020
	},
	{
	  "date": "2020-02-24",
	  "day": 1,
	  "month": 2,
	  "total": 6,
	  "week": 9,
	  "year": 2020
	},
	{
	  "date": "2020-02-25",
	  "day": 2,
	  "month": 2,
	  "total": 11,
	  "week": 9,
	  "year": 2020
	},
	{
	  "date": "2020-02-26",
	  "day": 3,
	  "month": 2,
	  "total": 13,
	  "week": 9,
	  "year": 2020
	},
	{
	  "date": "2020-02-27",
	  "day": 4,
	  "month": 2,
	  "total": 9,
	  "week": 9,
	  "year": 2020
	},
	{
	  "date": "2020-02-28",
	  "day": 5,
	  "month": 2,
	  "total": 28,
	  "week": 9,
	  "year": 2020
	},
	{
	  "date": "2020-02-29",
	  "day": 6,
	  "month": 2,
	  "total": 11,
	  "week": 9,
	  "year": 2020
	},
	{
	  "date": "2020-03-01",
	  "day": 0,
	  "month": 3,
	  "total": 16,
	  "week": 10,
	  "year": 2020
	},
	{
	  "date": "2020-03-02",
	  "day": 1,
	  "month": 3,
	  "total": 28,
	  "week": 10,
	  "year": 2020
	},
	{
	  "date": "2020-03-03",
	  "day": 2,
	  "month": 3,
	  "total": 17,
	  "week": 10,
	  "year": 2020
	},
	{
	  "date": "2020-03-04",
	  "day": 3,
	  "month": 3,
	  "total": 29,
	  "week": 10,
	  "year": 2020
	},
	{
	  "date": "2020-03-05",
	  "day": 4,
	  "month": 3,
	  "total": 11,
	  "week": 10,
	  "year": 2020
	},
	{
	  "date": "2020-03-06",
	  "day": 5,
	  "month": 3,
	  "total": 3,
	  "week": 10,
	  "year": 2020
	},
	{
	  "date": "2020-03-07",
	  "day": 6,
	  "month": 3,
	  "total": 0,
	  "week": 10,
	  "year": 2020
	},
	{
	  "date": "2020-03-08",
	  "day": 0,
	  "month": 3,
	  "total": 0,
	  "week": 11,
	  "year": 2020
	},
	{
	  "date": "2020-03-09",
	  "day": 1,
	  "month": 3,
	  "total": 15,
	  "week": 11,
	  "year": 2020
	},
	{
	  "date": "2020-03-10",
	  "day": 2,
	  "month": 3,
	  "total": 12,
	  "week": 11,
	  "year": 2020
	},
	{
	  "date": "2020-03-11",
	  "day": 3,
	  "month": 3,
	  "total": 21,
	  "week": 11,
	  "year": 2020
	},
	{
	  "date": "2020-03-12",
	  "day": 4,
	  "month": 3,
	  "total": 10,
	  "week": 11,
	  "year": 2020
	},
	{
	  "date": "2020-03-13",
	  "day": 5,
	  "month": 3,
	  "total": 5,
	  "week": 11,
	  "year": 2020
	},
	{
	  "date": "2020-03-14",
	  "day": 6,
	  "month": 3,
	  "total": 0,
	  "week": 11,
	  "year": 2020
	},
	{
	  "date": "2020-03-15",
	  "day": 0,
	  "month": 3,
	  "total": 0,
	  "week": 12,
	  "year": 2020
	},
	{
	  "date": "2020-03-16",
	  "day": 1,
	  "month": 3,
	  "total": 2,
	  "week": 12,
	  "year": 2020
	},
	{
	  "date": "2020-03-17",
	  "day": 2,
	  "month": 3,
	  "total": 26,
	  "week": 12,
	  "year": 2020
	},
	{
	  "date": "2020-03-18",
	  "day": 3,
	  "month": 3,
	  "total": 19,
	  "week": 12,
	  "year": 2020
	},
	{
	  "date": "2020-03-19",
	  "day": 4,
	  "month": 3,
	  "total": 23,
	  "week": 12,
	  "year": 2020
	},
	{
	  "date": "2020-03-20",
	  "day": 5,
	  "month": 3,
	  "total": 5,
	  "week": 12,
	  "year": 2020
	},
	{
	  "date": "2020-03-21",
	  "day": 6,
	  "month": 3,
	  "total": 0,
	  "week": 12,
	  "year": 2020
	},
	{
	  "date": "2020-03-22",
	  "day": 0,
	  "month": 3,
	  "total": 0,
	  "week": 13,
	  "year": 2020
	},
	{
	  "date": "2020-03-23",
	  "day": 1,
	  "month": 3,
	  "total": 0,
	  "week": 13,
	  "year": 2020
	},
	{
	  "date": "2020-03-24",
	  "day": 2,
	  "month": 3,
	  "total": 5,
	  "week": 13,
	  "year": 2020
	},
	{
	  "date": "2020-03-25",
	  "day": 3,
	  "month": 3,
	  "total": 11,
	  "week": 13,
	  "year": 2020
	},
	{
	  "date": "2020-03-26",
	  "day": 4,
	  "month": 3,
	  "total": 2,
	  "week": 13,
	  "year": 2020
	},
	{
	  "date": "2020-03-27",
	  "day": 5,
	  "month": 3,
	  "total": 0,
	  "week": 13,
	  "year": 2020
	},
	{
	  "date": "2020-03-28",
	  "day": 6,
	  "month": 3,
	  "total": 0,
	  "week": 13,
	  "year": 2020
	},
	{
	  "date": "2020-03-29",
	  "day": 0,
	  "month": 3,
	  "total": 1,
	  "week": 14,
	  "year": 2020
	},
	{
	  "date": "2020-03-30",
	  "day": 1,
	  "month": 3,
	  "total": 6,
	  "week": 14,
	  "year": 2020
	},
	{
	  "date": "2020-03-31",
	  "day": 2,
	  "month": 3,
	  "total": 0,
	  "week": 14,
	  "year": 2020
	}
  ]`
