package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// NextTick returns the tick of duration d that will begin soonest after time t.
func NextTick(t time.Time, d time.Duration) time.Time {
	t = t.UTC()
	start := time.Date(
		t.Year(), t.Month(), t.Day(),
		0, 0, 0, 0,
		time.UTC,
	)
	tomorrow := start.Add(24 * time.Hour)
	for tt := start; tt.Before(tomorrow); tt = tt.Add(d) {
		if tt.After(t) {
			return tt
		}
	}
	return tomorrow
}

const layout = "Jan _2 15:04 MST 2006"

const usage = `usage: seeds [-start time] [duration] [count]

seeds calculates when a Farming seed in Old School RuneScape with the given
tick duration (e.g. 160m) and number of ticks will be grown to completion.
Planting time can be specified with the -start flag (e.g. Jan 2 15:04 MST 2006);
otherwise, seeds will use the current time for planting.`

func main() {
	startp := flag.String("start", "", "the time when the seed was planted")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
	}
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		os.Exit(1)
	}
	start := time.Now()
	if *startp != "" {
		t, err := time.Parse(layout, *startp)
		if err != nil {
			log.Fatal(err)
		}
		start = t
	}
	d, err := time.ParseDuration(args[0])
	if err != nil {
		log.Fatal(err)
	}
	n, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	growd := d * time.Duration(n)
	loc := start.Location()
	complete := NextTick(start, d).Add(growd).In(loc).Format(layout)
	fmt.Println(complete)
}
