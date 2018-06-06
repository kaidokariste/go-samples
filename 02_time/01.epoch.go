package main

import (
	"fmt"
	"time"
)

func main() {

	//now := time.Now()
	//secs := now.Unix()
	//fmt.Println(now)
	//fmt.Println(secs)

	//Layout should follow the standard Mon Jan 2 15:04:05 -0700 MST 2006 (look time.Parse Doc)
	const longForm= "2006-01-02T15:04:05-07:00"

	tm, err := time.Parse(longForm, "2018-02-02T18:09:34+02:00") // -> returnt Time struct in UTC
	// expected 1517587774 from 2018-02-02 16:09:34.000000 or 2018-02-02T18:09:34-02:00
	if err != nil {
		panic(err)
	}
	unixts := tm.Unix()

	//From unix to local time
	unxToloc := time.Unix(1528212031,0)

	fmt.Println(tm)
	fmt.Println(unixts)
	fmt.Println(unxToloc)

}

