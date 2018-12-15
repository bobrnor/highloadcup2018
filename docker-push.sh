#!/usr/bin/env bash

source ./credentials

docker login -u $USERNAME -p $PASSWORD stor.highloadcup.ru
docker push stor.highloadcup.ru/accounts/barracuda_workery
