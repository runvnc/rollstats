#!/bin/bash
gpp rollstats.go -o rollstats_proc.go
./tealang -c rollstats_proc.go -o rollstats.teal
