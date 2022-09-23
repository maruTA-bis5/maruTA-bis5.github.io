#!/bin/bash
docker run --rm \
	-u `id -u`:`id -g` \
	-v $PWD:/home/circleci/project \
	-p 1313:1313 \
	cibuilds/hugo:0.102.3 hugo $@
