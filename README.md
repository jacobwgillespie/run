# 🦊 run

A small utility for running `package.json` scripts and npm modules. `run` is written in Go and executes scripts faster than `npm`, `yarn`, or `pnpm` since it can skip the Node.js startup time. Additionally `run` supports dynamic shell completion that suggests available script names as completions.

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
  Time (mean ± σ):     183.3 ms ±   2.0 ms    [User: 185.2 ms, System: 27.2 ms]
  Range (min … max):   181.0 ms … 188.5 ms    15 runs

Benchmark 2: yarn run echo-example
  Time (mean ± σ):     109.7 ms ±   0.8 ms    [User: 93.2 ms, System: 16.6 ms]
  Range (min … max):   108.9 ms … 113.0 ms    27 runs

Benchmark 3: pnpm run echo-example
  Time (mean ± σ):     223.7 ms ±   2.7 ms    [User: 212.5 ms, System: 18.5 ms]
  Range (min … max):   221.1 ms … 231.2 ms    13 runs

Benchmark 4: ./run echo-example
  Time (mean ± σ):       5.1 ms ±   0.2 ms    [User: 1.9 ms, System: 2.5 ms]
  Range (min … max):     4.8 ms …   5.9 ms    507 runs

Summary
  ./run echo-example ran
   21.40 ± 0.67 times faster than yarn run echo-example
   35.75 ± 1.16 times faster than npm run echo-example
   43.62 ± 1.43 times faster than pnpm run echo-example
```

If you would like to run the benchmark on your local machine, ensure you have [hyperfine](https://github.com/sharkdp/hyperfine) installed, then run `npm run benchmark` (or just `run benchark` if `run` is already installed 🦊).

## License

MIT License, see `LICENSE`.
