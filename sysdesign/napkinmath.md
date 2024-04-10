# Napkin math
"Estimation is a vital skill for an engineer. It’s something you can get better at by practicing." - [Roberto Vitillo](https://robertovitillo.com/back-of-the-envelope-estimation-hacks/)

Also known as back of the envelope calculations. These numbers are **heavily** rounded for memorizing. We want to be in the ballpark for creating useful mental models and to develop a gut feel for numbers being discussed to streamline conversations and make sure we've got a top-notch mental model of what's being talked about. If you need accurate numbers, then do your Ti-83 math instead.

## Little's Law
Lot's of things can be queues: message brokers like Google's Pub/Sub or even a web service. Little's law can help us relate the the average of new items arriving into a queue, the average time an item spends in a queue, and the number of items in the queue.
`avg_items_in_queue = avg_rate_of_new_items_coming_into_the_queue * time_items_stay_in_queue` or `q = r * t`. Say we have a service that takes an average of a 100ms to process a request and we are getting 2 million requests per second (rps). The equation via Little's Law is `r * t = 2_000_000 request/S * 0.1 S/request = 200_000 requests total = q`. We can take this a step further and calculate the number of computers we need. Assume that the requests are data heavy, and that we'll need an entire CPU thread for each request. If we have 200K requests per second and we have 8 core machines available, we'll need `200_000 requests / 8 core machines = 25K machines`.

## The rule of 72
The rule of 72 estimates how long it'll take for a value to double if it grows some percent rate. The rule is `time = 72/rate` or `t = 72/r`. For example, if requests are increasing at 10% per day, then `t = 72/r = 72/10 = 7.2 days`.

## Sorting
We need to memorize a basic sorting algorithm, like insertion sort:
```go
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		// This is the magic. j starts at the ith element and walks smaller elements to the front of the array.
		// So, while j-1 > j, swap it so that j-1 is < j and decrement j. Note j > 0 cause we're doing j-- stuff.
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
      arr[j-1], arr[j] = arr[j], arr[j-1] // Swap
		}
	}
}
```

## Stacks and queues
If you need to implement your own go stack or queue, you don't need to go crazy with locking, generics, or efficiency. Here's how to code rudimentary stacks and queues.

```go
// A straightforward stack without generics, locking, etc.
type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}
```

```go
// A straightforward queue without generics, locking, etc.
type Queue []int

func (s *Queue) Enq(x int) {
	*s = append(*s, x)
}

func (s *Queue) Deq() int {
	result := (*s)[0]
	*s = (*s)[1:len(*s)]
	return result
}

func (s *Queue) Peek() int {
	return (*s)[0]
}
```

## Costs
Cloud costs are important. It's good to have a sense of how much cloud computing costs will run. If we're using 400, four core CPUs per month, we want to know what that'll cost.

Name | $/time/quantity | Description
--- | --- | ---
Core | $10/mo/core | Cost of cloud computing cores per month (typically have many)
SSD | ¢10/mo/GB | Cost of SSD storage per month GB
HDD | ¢1/mo/GB | Cost of a HDD storage per month per GB
CDN | ¢1/mo/GB | Cost in AWS or GCP for CDN storage per month per GB
Network | ¢1/mo/GB | Networking utilization costs in AWS or GCP per month per GB

## Uptime in nines
Obviously, zero downtime in systems is ideal, but this ain't realistic. We want to be as close as financially reasonable to 100% up time. We talk about this using the "number of nines". The nines themselves don't really tell you the actual amount of downtime allowed. For example, what's 4 nines (99.99%) of 365 days? Why, it's 0.0365 days obviously! Yeah, no, still not helpful. Memorize the times.

| No. of Nines | %        | Daily downtime | Annual downtime |
|--------------|----------|----------------|-----------------|
| 1 nine       | 90%      | 150min         | 36days          |
| 1.5 nines    | 95%      | 75min          | 18days          |
| 2 nines      | 99%      | 15min          | 4days           |
| 3 nines      | 99.9%    | 2min           | 9hrs            |
| 4 nines      | 99.99%   | 10s            | 1hr             |
| 5 nines      | 99.999%  | 1s             | 5min            |
| 6 nines      | 99.9999% | 100ms          | 30sec           |

Also note that downtime among serial systems is typically additive. Say we've get two services named A and B where A depends on B. If they both have 5 nines of uptime, then they could both be down for 1s per day but there's no guarantee these times overlap. So, the worst case is going to be 2 services * 1s of downtime = 2s of downtime.

The more nines, the more expensive it is. For example, in terms of load balancers (LB), you may need to consider LBs+1, LBs+2, or 2*LBs to get your desire uptime. Now, this formula isn't perfect, but the availability is roughly `1 - ((1 - server_nines_as_decimal)^server_Count)`. So, two servers with 99% availability becomes `1-(0.01*0.01)=0.9999`. Now, again, this doesn't account for chained failures among many replicas.


## Data storage

## Networking

## Computations

## DBs with SQL

## DBs with NoSQL

## RAM disk providers
Like Redis or Memcached.

## Storage devices

## Serialization

## Hashing

## Common HTTP status codes
Code | Name | Description
--- | --- | ---
200 | OK | Indicates the request succeeded.
201 | Created | XXXXXXX
202 | Accepted | XXXXXX
204 | No content | XXXXXXX
307 | Temporary redirect | XXXXXX
308 | Permanent redirect | XXXXXXX
400 | Bad request | XXXXXX
401 | Unauthorized | XXXXXX
403 | Forbidden | XXXXXXX
404 | Not found | XXXXXXX
408 | Timeout | XXXXXXX
409 | Conflict | XXXXXX
429 | Too many requests | XXXXXXX
500 | Internal server error | XXXXXX
501 | Not implemented | XXXXXXX
502 | Bad gateway | XXXXXXXX
503 | Service unavailable | XXXXXXX
504 | Gateway timeout | XXXXXXX
