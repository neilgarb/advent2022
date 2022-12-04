#!/bin/bash

t=0
(
  s=0
  cat input.txt | while read l; do [ -z $l ] && echo $s && s=0 || s=$((s + $l)); done
  echo $s
) | sort -n | tail -n 3 | (
  while read l; do t=$(($t + $l)); done
  echo $t
)
