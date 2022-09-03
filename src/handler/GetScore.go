package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetRealTimeScore(address string) (string, error) {
	tureTimeScore := "https://api.forta.network/stats/sla/scanner/" + address
	resp, err := http.Get(tureTimeScore)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var dat map[string]interface{}
	//var score map[string]interface{}
	if err := json.Unmarshal(body, &dat); err == nil {
		test := dat["statistics"].(map[string]interface{})["avg"]
		println(test)
		avg, _ := fmt.Printf("%.1f", dat["statistics"].(map[string]interface{})["avg"])
		return strconv.Itoa(avg), nil
	} else {
		return "", err
	}
}

func GetWeekScore(address string) (string, error) {
	lastMonday := getMondayOfWeek(time.Now(), "2006-01-02T15:04:05Z")
	weekScore := "https://api.forta.network/stats/sla/scanner/" + address + "?startTime=" + lastMonday
	resp, err := http.Get(weekScore)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var dat map[string]interface{}
	//var score map[string]interface{}
	if err := json.Unmarshal(body, &dat); err == nil {
		avg, _ := fmt.Printf("%.1f", dat["statistics"].(map[string]interface{})["avg"])
		return strconv.Itoa(avg), nil
	} else {
		return "", err
	}
}

func getMondayOfWeek(t time.Time, fmtStr string) (dayStr string) {
	dayObj := getZeroTime(t)
	if t.Weekday() == time.Monday {
		//修改hour、min、sec = 0后格式化
		dayStr = dayObj.Format(fmtStr)
	} else {
		offset := int(time.Monday - t.Weekday())
		if offset > 0 {
			offset = -6
		}
		dayStr = dayObj.AddDate(0, 0, offset).Format(fmtStr)
	}
	return
}

func getZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
