#!/usr/bin/env bash

go build
hyperfine --warmup 5 -N 'npm run echo-example' 'yarn run echo-example' 'pnpm run echo-example' './run echo-example'
