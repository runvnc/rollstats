#!/bin/bash
gpp rollstats.go -o rollstats_proc.go
tl3 -c rollstats_proc.go -o rollstats.teal
