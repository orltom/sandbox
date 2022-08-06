#!/bin/bash

echo "CREATE TABLE chuck_norris (id SERIAL NOT NULL, joke VARCHAR NOT NULL, PRIMARY KEY (id));"

for i in {1..10}; do
  JOKE=$(http https://api.chucknorris.io/jokes/random | jq -r '.value' | sed "s/'/''/g")
  echo "INSERT INTO chuck_norris(joke) VALUES ('${JOKE}');"
done