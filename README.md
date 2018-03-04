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
flag is specified, a planting time can be specified (RFC3339); otherwise, seeds
will use the current time for planting.
