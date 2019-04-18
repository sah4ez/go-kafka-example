#!/bin/bash

docker rm testapp_p

docker run --rm \
	-it \
	--name testapp_p \
	--link kafka1:kafka \
	test app -publisher -kafka=kafka:9092
