#!/bin/bash

START=
END=
FILE_NAME="block_time_"$START"_"$END".txt"
touch ./"$FILE_NAME"
for i in `seq $START $END`
do
    ununifid q block "$i" | jq -r '.block.header.time' >> ./"$FILE_NAME"
done
