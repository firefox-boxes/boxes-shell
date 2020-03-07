#!/bin/bash

mkdir dist
go build -o dist/boxes-shell $@ boxes-shell.go shared.go
go build -o dist/boxes-ext-native-shell $@ boxes-ext-native-shell.go shared.go
