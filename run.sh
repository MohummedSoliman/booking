#!/bin/bash
#
# The benefits of put this two commands in the same line with && instead of adding every one
# in separate line is :-
# the second command only run if the first command successed.
go build -o booking cmd/web/*.go && ./booking
