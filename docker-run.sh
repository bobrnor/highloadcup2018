#!/usr/bin/env bash

docker rm barracuda_workery
docker run --name barracuda_workery -p:80:80 -v /Users/daniil.zinenko/Development/pets/hlc2018/testdata/data:/tmp/data -it stor.highloadcup.ru/accounts/barracuda_workery