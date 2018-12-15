#!/usr/bin/env bash

docker rm barracuda_workery
docker run --name barracuda_workery -v /Users/daniil.zinenko/Development/pets/go/src/github.com/bobrnor/highloadcup2018/testdata/data:/tmp/data -it stor.highloadcup.ru/accounts/barracuda_workery