#!/bin/bash


name=$1
partion=$2

docker rm $1 

docker run \
	-it \
	--name $name \
	--link kafka1:kafka \
	test app -subscriber -kafka=kafka:9092 -partition=$partion
	
