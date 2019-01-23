#!/bin/bash
cd db
VALUE=(`ruby parse.rb`)
DBNAME=${VALUE[0]}
USER=${VALUE[1]}
PASS=${VALUE[2]}
cd ..
