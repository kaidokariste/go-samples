package main

import (
	"fmt"
	"time"
)

func localOffset() (zonename string, offs int) {
	var (
		offset int
		name   string
	)
	name, offset = time.Now().Local().Zone()
	return name, offset
}

func main() {
	loc, _ := time.LoadLocation("Europe/Tallinn")
	utc := time.Now().UTC().Format("15:04")
	est := time.Now().In(loc)
	unixts := time.Now().In(loc).Unix()
	fmt.Println(localOffset())
	fmt.Println(loc, utc, est, unixts)
}
