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
tick duration (e.g. 160m) and number of ticks will be grown to completion.
Planting time can be specified with the -start flag (e.g. Jan 2 15:04 MST 2006);
otherwise, seeds will use the current time for planting.
