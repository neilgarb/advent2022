#!/bin/bash

(
  s=0
  cat input.txt | while read l; do [ -z $l ] && echo $s && s=0 || s=$((s + $l)); done
  echo $s
) | sort -n | tail -n 1
