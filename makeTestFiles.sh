#!/bin/bash

rm -rf testdata/*
for i in {0..10};do
    mkdir testdata/dir${i}
    for j in {0..50};do
        mkfile 50m testdata/dir${i}/test-${j}.txt
    done
done

du -h testdata
