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
  score=$(($score + $scoreB + 1))
  [ "$scoreA" -eq "$scoreB" ] && score=$(($score + 3))
  [ "$scoreB" -eq "$((($scoreA + 1) % 3))" ] && score=$(($score + 6))
done
echo $score
