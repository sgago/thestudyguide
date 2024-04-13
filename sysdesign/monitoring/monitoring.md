# Monitoring
The main goal of monitoring is to alert teams to failures and overall system health via dashboards. You'll see line charts, gauges, counts, pass/fails, and many other graphical representations of various aspects of the systems. We can monitor pub/sub queue sizes, replica counts, replica restarts, request counts, P99 response times, count server status codes, SLO/SLI values, and much more.

## Metrics
Metrics are raw measurements or samples stored as floating point values. To keep a long story short, to save money and disk space, the data are pre-aggregated (summed, averaged, etc.) into certain buckets (1s, 5s, 10s, 15s, 1 min, etc.) within the time-series databases. See [monitoring](../../sysdesign.md) in the system design document for a deeper explanation of all this. This guide is focusing on how to get started and use Go, Prometheus, and Grafana and less interested in what's under the hood of time-series databases.

Anyway, this example project will store metrics in Prometheus, an open-source time-series database.

Some metric examples include:
```
# Total number of HTTP request
http_requests_total

# Response status of HTTP request
response_status

# Duration of HTTP requests in seconds
http_response_time_seconds
```

Note that metrics typically follow a suffix convention to make them easier to understand and query. Even in our basic example here, you'll see tons of different metrics.

## Labels
Now, if we have many replicas, services, availability zones, URL paths, and so on, it'll be really hard to find metric values we care about without labels. Labels are key-values that we can slap onto our metrics that'll make querying later easier. For example,

```
# Total number of HTTP request
http_requests_total{service="builder"}

# Response status of HTTP request
response_status{path="/"}
response_status{path="/articles"}
```

## Metric types
Prometheus supplies four different types of metrics: counters, gauges, histograms, and summaries.
- *Counters* are a simple, single metric that can only be incremented or reset to zero, and are used to count total number of requests or total number of tasks completed. Most counters are suffixed with `_total` like `http_requests_total` or `jobs_completed_total`. Counters aren't super helpful by themselves; more information can be gather from them when querying with the `rate()` function.
- *Gauges* are also simple, single values but these values can go up and down, and are used for CPU utilization, temperature, RAM usage, and so on. Unlike counters, we can use gauges directly on graphs.
- *Histograms* are used to measure the frequency of values following into predefined buckets, and can be used for request response time, job completion time, and similar. This is useful for say serving requests within 300ms for 95% of all requests. Prometheus uses default buckets of .005, .01, .025, .05, .075, .1, .25, .5, .75, 1, 2.5, 5, 7.5, and 10 though they can be changed.
- *Summaries* are basically the same as histograms, but the values are collected client-side. This yields better accuracy but is more resource intensive. Histograms are typically preferred.

## Running the project
Run `docker compose up -d`. Using default values, you can hit:
- Prometheus at [http://localhost:9090](http://localhost:9090)
- Grafana at [http://localhost:3000/](http://localhost:3000). If you're spinning this up for the first time, the username and password are both `admin`. You'll also need to add Prometheus as a data source in Grafana. Use Connections -> Add data source -> Prometheus. Set the Prometheus server URL to `http://prometheus:9090` and click `Save & Test` at the bottom of the page. You should receive a success message.
- Our boring Go Gin app is at [http://localhost:8080](http://localhost:8080)

## Moving the metrics
So, we need to make the metrics move around some so that it's at least interesting to look at. We can use the load testing tool hey to make that happen. The hey load testing tool repo is [here](https://github.com/rakyll/hey). Basic use is
`hey -z 5m -q 5 -m GET -H "Accept: text/html" http://localhost:8080`
where:
-z = duration of load testing, like 10s or 5m.
-q = Rate limit of load testing, in queries per second (QPS).
-m = HTTP method: GET, POST, PUT, DELETE, or HEAD.
-H = a customer HTTP header like `-H "Accept: text/html. You can chain many -H's together like -H "Accept: text/html" -H "Content-Type: application/xml"`.

## Prometheus query language
Prometheus query language (PromQL) lets users select and aggregate their time series data. It has nothing in common with other query languages like SQL.

Again, to get started, navigate to this directory via CLI and
- Run `docker compose up -d` or `dcupd` to start this project
- Run the load testing tool via `hey -z 5m -q 5 -m GET -H "Accept: text/html" http://localhost:8080` to make the metrics change over time. Increase the `5m` (five minutes) to `10m` or `60m` if you're going to be studying longer.
- 

This project includes 6 self-defined metrics
- cpu_temperature_celsius
- ping_path_total
- ping_path_response_bucket
- index_path_total
- index_path_response_bucket

## References
- Gabriel Tanner has a great blog on getting setup with Go, Prometheus, and Grafana that I followed [here](https://gabrieltanner.org/blog/collecting-prometheus-metrics-in-golang/). Very easy to follow.
- PromQL querying basics [here](https://prometheus.io/docs/prometheus/latest/querying/basics/)
- PromQL examples [here](https://prometheus.io/docs/prometheus/latest/querying/examples/)
- PromQL quick start from medium [here](https://valyala.medium.com/promql-tutorial-for-beginners-9ab455142085)