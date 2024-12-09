#!/usr/bin/env bash

read -p "Enter year: " year
read -p "Enter day:  " day

printf -v paddedDay "%02d" $day
echo $paddedDay

dir="$year/day-$paddedDay"

mkdir -p $dir

cp -n 2020/template.go $dir/main.go
touch $dir/data
touch $dir/test.data