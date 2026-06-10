# Helm Golang Test

A short exercise: build a small HTTP service that calculates profit and loss for a portfolio of shares from a CSV file. It should take around an hour — we're not looking for a production system, just clean, idiomatic Go.

## Prerequisites

- Go 1.22+

## Getting Started

```
go run .
curl localhost:8080/health
```

The boilerplate in `main.go` gives you a running HTTP server with a health endpoint. Everything else is up to you.

## The Data

`data/positions.csv` contains the portfolio:

```
ticker,quantity,buy_price,current_price
AAPL,10,150.00,175.50
GOOG,5,2800.00,2650.75
...
```

Each row is a position: a number of shares bought at `buy_price`, currently worth `current_price`. The same ticker can appear more than once (multiple lots).

## The Task

Add a `GET /pnl` endpoint that reads the CSV and returns the profit and loss for the portfolio as JSON.

For a single position:

```
pnl = (current_price - buy_price) * quantity
```

The response should include the P&L per ticker (lots for the same ticker combined) and a total for the whole portfolio. Something like:

```json
{
  "positions": [
    { "ticker": "AAPL", "quantity": 14, "pnl": 297.00 },
    { "ticker": "TSLA", "quantity": 17, "pnl": -616.00 }
  ],
  "total_pnl": 1100.80
}
```

Don't worry about matching this shape exactly — design the response how you see fit.

## Concurrency

We'd like to see you use goroutines and channels. The calculation is small enough that concurrency isn't strictly necessary, so this is about demonstrating the pattern, not the performance win.

A suggested approach: a worker pool. Feed positions into a jobs channel, have a fixed number of worker goroutines calculate the P&L for each, and collect the results on a results channel before aggregating. A `sync.WaitGroup` to know when the workers are done is the usual companion. Imagine the portfolio had a million rows — structure it as if that were the case.

## Goals

1. `GET /pnl` returns per-ticker and total P&L as JSON, calculated from the CSV.
2. The calculation fans out across a worker pool using channels.
3. Errors are handled sensibly — a missing or malformed CSV shouldn't crash the server or return a 200.

## Stretch Goals

Only if you have time left — none of these are required:

- Support `GET /pnl?ticker=AAPL` to return the P&L for a single ticker.
- Include a percentage return per ticker alongside the absolute P&L.
- Add a unit test or two around the calculation.
- Make the number of workers configurable.

## What We're Looking For

- Clean, readable, idiomatic Go — sensible names, small functions, standard project layout.
- Correct use of channels and goroutines — no leaks, no races (`go run -race .` is your friend).
- Sensible error handling and HTTP status codes.
- Stdlib only is perfectly fine; don't reach for a framework.
