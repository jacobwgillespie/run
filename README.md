# 🦊 run

A small utility for running `package.json` scripts and npm modules. `run` is written in Go and executes scripts faster than `npm`, `yarn`, or `pnpm` since it can skip the Node.js startup time.

### Installation

```
brew install jacobwgillespie/tap/run
```

### Usage

```
run [script] [flags]
```

## Benchmarks

Informally, `run` is about 25-50 times faster than Node package managers at starting scripts:

```
Benchmark 1: npm run echo-example
  Time (mean ± σ):     142.4 ms ±   2.1 ms    [User: 132.6 ms, System: 22.0 ms]
  Range (min … max):   139.7 ms … 150.5 ms    21 runs

Benchmark 2: yarn run echo-example
  Time (mean ± σ):      91.1 ms ±   0.7 ms    [User: 77.5 ms, System: 14.3 ms]
  Range (min … max):    89.5 ms …  92.5 ms    33 runs

Benchmark 3: pnpm run echo-example
  Time (mean ± σ):     204.9 ms ±   1.2 ms    [User: 192.4 ms, System: 16.1 ms]
  Range (min … max):   202.7 ms … 206.9 ms    14 runs

Benchmark 4: ./run echo-example
  Time (mean ± σ):       3.6 ms ±   0.2 ms    [User: 1.4 ms, System: 1.6 ms]
  Range (min … max):     3.5 ms …   5.0 ms    726 runs

Summary
  './run echo-example' ran
   24.98 ± 1.37 times faster than 'yarn run echo-example'
   39.05 ± 2.20 times faster than 'npm run echo-example'
   56.19 ± 3.07 times faster than 'pnpm run echo-example'
```

If you would like to run the benchmark on your local machine, ensure you have [hyperfine](https://github.com/sharkdp/hyperfine) installed, then run `npm run benchmark` (or just `run benchark` if `run` is already installed 🦊).

## License

MIT License, see `LICENSE`.
