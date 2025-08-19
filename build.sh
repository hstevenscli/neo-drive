#! /bin/bash

cd frontend/ && test -f package.json && npm ci && npm run build; cd ..;

go build -o ./app;
