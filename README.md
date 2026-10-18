# Go Network Tools

Small TCP networking tools in Go that I built while working through Chapter 2 of *Black Hat Go* (Steele, Patten and Kottmann, No Starch Press). Each step makes the port scanner a little better, from a sequential scan to a bounded worker pool.

| Tool | What it does |
| --- | --- |
| [`port-scanner-sequential`](cmd/port-scanner-sequential) | Connects to ports 1–1024 one by one with `net.Dial`. Simple, but slow. |
| [`port-scanner-unbounded`](cmd/port-scanner-unbounded) | Starts one goroutine per port with a `sync.WaitGroup`. Fast, but opens 1024 connections at once, which can give unreliable results. |
| [`worker-pool-demo`](cmd/worker-pool-demo) | Shows the worker-pool pattern on its own: 100 workers reading port numbers from a buffered channel. |
| [`port-scanner-worker-pool`](cmd/port-scanner-worker-pool) | The final scanner: a fixed pool of 100 workers, results collected over a channel, open ports sorted and printed. |
| [`tcp-echo-server`](cmd/tcp-echo-server) | A concurrent TCP server on port 20080 that echoes back whatever a client sends, one goroutine per connection. |
| [`tcp-port-forwarder`](cmd/tcp-port-forwarder) | Listens on `:8080` and forwards every connection to a remote host using two `io.Copy` streams, like a minimal TCP proxy. |
| [`io-reader-writer-demo`](cmd/io-reader-writer-demo) | Custom types that implement `io.Reader` and `io.Writer`, the interfaces the proxy is built on. |

## Run

```bash
go run ./cmd/port-scanner-worker-pool
go run ./cmd/tcp-echo-server        # then: nc localhost 20080
go run ./cmd/tcp-port-forwarder     # then: curl -H "Host: scanme.nmap.org" localhost:8080
```

## What I learned

- How a TCP connect scan works: a completed handshake means the port is open; a refused or timed-out connection means it is closed or filtered.
- Why unbounded concurrency breaks scanners, and how a worker pool with channels keeps it under control.
- How to build servers and proxies on the `net.Conn`, `io.Reader` and `io.Writer` interfaces.

## Use responsibly

The scanners target `scanme.nmap.org`, a host the Nmap project provides for testing. Only scan systems you own or are allowed to test.
