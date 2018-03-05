package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func NextTick(t time.Time, d time.Duration) time.Time {
	var ticks []time.Time
	start := time.Date(
		t.Year(), t.Month(), t.Day(),
		0, 0, 0, 0,
		time.UTC,
	)
	tomorrow := start.Add(24 * time.Hour)
	for tt := start; tt.Before(tomorrow); tt = tt.Add(d) {
		ticks = append(ticks, tt)
	}
	for _, tt := range ticks {
		if tt.After(t) {
			return tt
		}
	}
	return tomorrow
}

const usage = `usage: seeds [-start time] [duration] [count]

seeds calculates when a Farming seed in Old School RuneScape with the given
tick duration and number of ticks will be grown to completion. If the -start
flag is specified, a planting time can be specified in the form e.g.
Jan 2 15:04 2006; otherwise, seeds will use the current time for planting.`

const layout = "Jan _2 15:04 2006"

func main() {
	startp := flag.String("start", "", "the time when the seed was planted (RFC3339)")
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
	loc := start.Location()
	if *startp != "" {
		t, err := time.ParseInLocation(layout, *startp, loc)
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
	t := NextTick(start, d).Add(growd)
	fmt.Println(t.In(loc).Format(layout))
}
