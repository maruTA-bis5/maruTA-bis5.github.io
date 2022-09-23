#!/bin/bash
docker run --rm \
	-u `id -u`:`id -g` \
	-v $PWD:/home/circleci/project \
	--network host \
	cibuilds/hugo:0.102.3 hugo $@
