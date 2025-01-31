# Design

## Monoliths


### N-Tier applications


## Microservices


## Service oriented architecture (SOA)
SOA uses software components called _services_ to create business applications.
Each service provides a business capability and communicates with other services. Microservices are typically seen as an evolution of SOA to address SOA's shortcomings.

### Enterprise service bus (ESB)
Services typically communicate over an ESB like [Apache Camel](https://camel.apache.org/) or Mulesoft's [Anypoint Platform](https://www.mulesoft.com/platform/enterprise-integration).

```
               +-------------------+
               |       ESB         |
               +--------+----------+
                        |
        +---------------+-----------------+
        |               |                 |
+-------+-------+ +-----+-----+   +-------+-------+
| User Service  | | Order     |   | Payment       |
|               | | Service   |   | Service       |
+---------------+ +-----+-----+   +-------+-------+
        |               |                 |
+-------+-------+ +-----+-----+   +-------+-------+
| Product       | | Inventory |   | Notification  |
| Service       | | Service   |   | Service       |
+---------------+ +-----------+   +---------------+
```

### SOA vs microservices
Although similar, there are some differences between SOA and microservices:
- SOA typically has more coarse-grained services that handle substantial business processes. Microservices are typically fine-grained and handling a small set of functions.
- SOA relies on complex middleware solutions for communication, such as enterprise services buses (ESBs) with multiple messaging protocols: SOAP, AMQP, or MSMQ. Microservices use HTTP/REST, gRPC, or message brokers.
- SOA are typically deployed to a few large containers. Microservices are typically deployed through containerization via docker and orchestrated via Kubernetes.
- SOAs typically share data storage, unlike microservices.
- SOAs can slow as more services are tacked on; microservices can (maybe) stay consistent.

## SOLID

### S is for single responsibility principle
**Classes should have a single responsibility, a single purpose.**

### O is for open/closed principle
**Classes should be open to extension but closed to modification.**

### L is for Liskov's substitution principle
**Subclasses should be able to be used anywhere a superclass is used.**

### I is for interface segregation principle
**Classes should not be forced to depend upon interfaces that they do not use.**

### D is for dependency inversion principle
**Depend on abstractions, not concretes.**

## Design patterns

### Model-View-Controller

#### Model-View-ViewModel

## Dependency inversion (DI) and inversion-of-control (IoC) containers

## Command-query responsibility segregation (CQRS) principle

## Repository and Unit-of-Work (UoW) patterns

## Domain driven design (DDD)
DDD is a conceptual approach to development that emphasizes understanding and modeling the core business domains via collaborating with domain experts.

### Domains
*Domains* are a subject area for a project: products, users, customers, finance,
and so on.

## Evaluating risk