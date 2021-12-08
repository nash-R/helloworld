package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {

	u := os.Args[1]

	for {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("get Dump file :", "null")
			}
		}()
		res, err := http.Get(u)
		getYear, getMonth, getDay := time.Now().Year(), time.Now().Month(), time.Now().Day()
		getHour, getMinute, getSecond := time.Now().Hour(), time.Now().Minute(), time.Now().Second()
		getTime := strconv.Itoa(getYear) + "-" + strconv.Itoa(int(getMonth)) + "-" + strconv.Itoa(getDay) + " " + strconv.Itoa(getHour) + ":" + strconv.Itoa(getMinute) + ":" + strconv.Itoa(getSecond)
		if err != nil {
			getDump(err)
		}
		fmt.Println("success:", res.StatusCode, "at", getTime)
		res.Body.Close()
		time.Sleep(time.Second * 1)
	}
}

func getDump(err error) {
	fmt.Println("error:", err)
	// 抓dump
	cmdOut, _ := exec.Command("hostname").CombinedOutput()
	fmt.Println(string(cmdOut))
	// return
}
