# Reverse Proxies
In short, reverse proxies are go-betweens, like a facade or abstraction, that sit between a client and the backend. They are an abstraction, a facade, that sits in front of websites.

They can provide
- Load balancing
- Security, like better DDoS protection
- Caching
- Handle TLS and certificates
- Security and access control
- Rate limiting
- Error handling
- URL rewriting
- Analytics
- Request translation, say from HTTP to gRPC

Acting as a gateway to a ton of disparate microservices, a reverse proxy as an API gateway provides these core services:
- *Routing* requests to the various services and allowing the internal microservices to change, hopefully without disrupting the public-facing side.
- *Composing* data or requests among many disparate microservices. For example, orders and customer data could require two requests internally, but the public-facing side would only need to make one. It relieves clients from having to know about all the different microservices and how to query them.
- *Translating* requests from one IPC mechanism to another. Say, converting a public HTTP request to an internal gRPC request or vice versa.
- *Cross-cutting concerns* such as authentication and authorization, logging, rate limiting, etc.

Now, everything is tradeoffs, and we do make tradeoffs when introducing a reverse proxy like traefik. We need to make sure that the reverse proxy scales well and does not become a single point of failure.

In terms of microservices and distributed systems, we want to minimize the number of endpoints that clients need to be aware of. We can introduce an API gateway, via reverse proxy, that hides the microservice architecture from clients. We can route requests via the reverse proxy, change internal endpoints without breaking clients, stitch together data from multiple services, and handle data inconsistencies.

An API gateway is especially beneficial for public APIs backed by microservices. The gateway can also provide rate-limiting, caching, and authentication/authorization. Again, we need to be mindful of tradeoffs with a gateway: scaling, coupling, and microservice API changes. In terms of implementing an API gateway, you can
- Create your own
- Start with a reverse proxy like [nginx](https://www.nginx.com/) (pronounced engine-X) or [traefik](https://traefik.io/traefik/) (pronounced traffic)
- Use a managed service like Google's [Apigee](https://cloud.google.com/apigee?hl=en)

In this example, we'll learn and explore traefik. Typically, nginx is faster but traefik's convenient, automagic service discovery shines through and can help expedite learning.
![](https://doc.traefik.io/traefik/assets/img/quickstart-diagram.png)

## 

## 

## References
