#!/bin/bash

cd ../views && templ generate
cd  ../cmd && go run main.go --server start
