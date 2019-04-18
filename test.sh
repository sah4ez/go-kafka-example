#!/bin/bash

docker rm testapp_s testapp_p

docker run \
	-itd \
	--name testapp_s \
	--link kafka1:kafka \
	test app -subscriber -kfakaAddr=kafka:9092
	
