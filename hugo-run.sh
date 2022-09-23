#!/bin/bash
docker run --rm \
	-u `id -u`:`id -g` \
	-v $PWD:/src \
	-p 1313:1313 \
	klakegg/hugo:0.101.0-ext-alpine $@
