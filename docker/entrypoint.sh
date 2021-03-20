#!/bin/bash

readonly BASEDIR=/usr/bin
readonly EXPORTER_PATH=$BASEDIR/pbspro_exporter

#environments
file_env() {
	local var="$1"
	local fileVar="${var}_FILE"
	local def="${2:-}"
	if [ "${!var:-}" ] && [ "${!fileVar:-}" ]; then
		echo >&2 "error: both $var and $fileVar are set (but are exclusive)"
		exit 1
	fi
	local val="$def"
	if [ "${!var:-}" ]; then
		val="${!var}"
	elif [ "${!fileVar:-}" ]; then
		val="$(< "${!fileVar}")"
	fi
	export "$var"="$val"
	unset "$fileVar"
}

file_env 'PBS_ADDR' '127.0.0.1'
file_env 'EXPORT_PORT' '9107'

export LD_LIBRARY_PATH=/usr/lib

$EXPORTER_PATH --collector.pbspro.url="$PBS_ADDR" --web.listen-address=":$EXPORT_PORT"


