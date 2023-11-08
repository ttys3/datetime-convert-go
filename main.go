package main

import (
	"fmt"
	"time"
)

func main() {
	// Define a time to convert
	currentTime := time.Now()
	fmt.Println("Current Time in Local    :", currentTime, currentTime.Unix())
	currentTimestamp := currentTime.Unix()

	// Display the current time in UTC
	fmt.Println("Current Time in UTC      :", currentTime.UTC(), currentTime.UTC().Unix())

	// Define the time zone you want to convert from
	newYorkLoc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error loading source location:", err)
		return
	}

	// Convert the current time to the source time zone
	newYorkTime := currentTime.In(newYorkLoc)
	fmt.Println("Current Time in New York :", newYorkTime, newYorkTime.Unix())

	// Define the time zone you want to convert from
	hkLoc, err := time.LoadLocation("Asia/Hong_Kong")
	if err != nil {
		fmt.Println("Error loading source location:", err)
		return
	}
	// Convert the current time to the source time zone
	hkTime := currentTime.In(hkLoc)
	fmt.Println("Current Time in Hong Kong:", hkTime, hkTime.Unix())

	// Define the time zone you want to convert to
	tokyoLoc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("Error loading destination location:", err)
		return
	}
	// Convert the time to the destination time zone
	tokyoTime := newYorkTime.In(tokyoLoc)
	fmt.Println("Current Time in Tokyo    :", tokyoTime, tokyoTime.Unix())

	fmt.Println("------------------------------------------------------")

	/*
	   Current Time in Local: 2023-11-08 16:58:03.193033945 +0000 UTC m=+0.000020900 1699462683
	   Current Time in UTC: 2023-11-08 16:58:03.193033945 +0000 UTC 1699462683
	   Current Time in New York: 2023-11-08 11:58:03.193033945 -0500 EST 1699462683
	   Current Time in Hong Kong: 2023-11-09 00:58:03.193033945 +0800 HKT 1699462683
	   Current Time in Tokyo: 2023-11-09 01:58:03.193033945 +0900 JST 1699462683
	*/

	revTime := time.Unix(currentTimestamp, 0)
	fmt.Println("revTime in UTC      :", revTime.UTC(), revTime.UTC().Unix())
	fmt.Println("revTime in Local    :", revTime, revTime.Unix())
	fmt.Println("revTime in New York :", revTime.In(newYorkLoc), revTime.In(newYorkLoc).Unix())
	fmt.Println("revTime in Hong Kong:", revTime.In(hkLoc), revTime.In(hkLoc).Unix())
	fmt.Println("revTime in Tokyo    :", revTime.In(tokyoLoc), revTime.In(tokyoLoc).Unix())

	fmt.Println("------------------------------------------------------")

	// unify to UTC timezone datetime
	fmt.Println("UTC:      ", revTime.UTC(), "in UTC:", time.Unix(currentTimestamp, 0).UTC(), time.Unix(currentTimestamp, 0).UTC().Unix())

	_, offsetLocal := currentTime.Zone()
	fmt.Println("Local:    ", revTime.Local(), "in UTC:", time.Unix(currentTimestamp-int64(offsetLocal), 0).Local().UTC(), time.Unix(currentTimestamp-int64(offsetLocal), 0).Local().UTC().Unix())

	_, offsetNewYork := newYorkTime.Zone()
	fmt.Println("New York: ", revTime.In(newYorkLoc), "in UTC:", time.Unix(currentTimestamp-int64(offsetNewYork), 0).UTC(), time.Unix(currentTimestamp-int64(offsetNewYork), 0).UTC().Unix())

	_, offsetHK := hkTime.Zone()
	fmt.Println("Hong Kong:", revTime.In(hkLoc), "in UTC:", time.Unix(currentTimestamp-int64(offsetHK), 0).UTC(), time.Unix(currentTimestamp-int64(offsetHK), 0).UTC().Unix())

	_, offsetTokyo := tokyoTime.Zone()
	fmt.Println("Tokyo:    ", revTime.In(tokyoLoc), "in UTC:", time.Unix(currentTimestamp-int64(offsetTokyo), 0).UTC(), time.Unix(currentTimestamp-int64(offsetTokyo), 0).UTC().Unix())
}
