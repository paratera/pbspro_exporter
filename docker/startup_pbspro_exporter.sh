#!/bin/bash

readonly IMGPRJ=taylor840326
readonly IMGNAME=pbspro_exporter
readonly IMGTAG=latest
readonly CNAME=pbspro_exporter

docker rm -f $CNAME

docker run \
	--name $CNAME \
	--network host \
	--restart always \
	-e PBS_ADDR=172.18.7.10 \
	-e EXPORTER_PORT=9107 \
	-d $IMGPRJ/$IMGNAME:$IMGTAG
