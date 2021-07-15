#!/usr/bin/env bash

want=`ls -t migrations 2>&1 | head -1 | egrep -o '[0-9]+'`
c=$(echo $want | xargs) 
if [[ -z $want ]]; then
    echo "Nothing migrate";
    exit
fi

echo "Want: '$want'";

count=3
timeout=3
while [ $count -gt 0 ]
do
    got=`psql "postgres://user:password@localhost:5432/test_db?sslmode=disable" -t -c 'select version from schema_migrations;' 2>&1 | head -1`
    got=$(echo $got | xargs) 
    echo "Got: '$got'"

    if [[ $want == $got ]]; then
        echo "EXEC...";
        exec $1
        break
    fi

    echo "Waiting migrate...";
    sleep $timeout
    count=$(($count-1))
done