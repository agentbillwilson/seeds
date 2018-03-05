seeds
=====

Calculates Farming seed completion times in Old School RuneScape

Building
--------

    go build


Usage
-----

    seeds [-start time] [duration] [count]

seeds calculates when a Farming seed in Old School RuneScape with the given
tick duration and number of ticks will be grown to completion. If the -start
flag is specified, a planting time can be specified in the form e.g.
Jan 2 15:04 2006; otherwise, seeds will use the current time for planting.
