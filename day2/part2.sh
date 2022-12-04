#!/bin/bash

score=0
index_of() {
  tmp=${2%%$1*}
  echo $(echo "${2%%$1*}" | echo ${#tmp})
}
cat input.txt | while read l; do
  playA="ABC"
  playB="XYZ"
  a=$(echo $l | awk '{print $1}')
  b=$(echo $l | awk '{print $2}')
  scoreA=$(index_of $a $playA)
  scoreB=$(index_of $b $playB)
  [ "$scoreB" -eq "0" ] && score=$(($score + ($scoreA - 1 + 3) % 3 + 1))
  [ "$scoreB" -eq "1" ] && score=$(($score + 3 + $scoreA + 1))
  [ "$scoreB" -eq "2" ] && score=$(($score + 6 + ($scoreA + 1) % 3 + 1))
done
echo $score
