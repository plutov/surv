## SURV Project 🏄 [![Build Status](https://travis-ci.org/plutov/surv.svg?branch=master)](https://travis-ci.org/plutov/surv) 

Why SURV? I believe projects with a name have a soul.

The SURV project consists of few main components: multiple survey services and a centralized dashboard.

High level diagram:

![surv high level diagram](https://s3.amazonaws.com/pliutau.com/High+Level+Diagram.png)

In this project we use JSON over HTTP for communication as survey services can be external ones and HTTP is the most common format nowadays. For internal services I'd go with protocol buffers as they are type safe and beter serialized.

### Data

Dashboard displays all survey submissions from all survey services.

![surv data diagram](https://s3.amazonaws.com/pliutau.com/Data+Diagram.png)

### Survey services

Two survey services use the same codebase (`./survey`) as they solve the similar problem, but they have different configuration and deployed separately. We can scale only one survey service if necessary and easily redeploy one without affecting other. If we develop 2 completely different survey services it may have sense to have a separate codebase as well so different teams can work on them. 

We separate survey services by domain, for example one service for gaming, another for online store.

Dashboard may access these services via known REST API endpoints, if there are multiple survey services dashboard may connect to `etcd` for service discovery to find an address of specific service.

#### Survey service configuration:

 - Name
 - Address

I manage local configuration in environment variables of Docker Compose, the production config can be stored in Vault / KMS / etc.

#### API

Basically each survey service is a separate REST API service. Now it is a public API, but in real world it's more safe to guard it with OAuth 2.0 or at least Basic Auth.

API is documented using Swagger UI.

### Dashboard service

Dashboard service (`./dashboard`) pulls data from different (but known) surveys. It has a connector mechanism (client) (`./dashboard/pkg/api/connectors.go`) for each type of survey, then it pulls data, aggregate it and stores in Dashboard storage. Here I assume that survey service API is almost always up, in real world we have to think about circuit breakers.

As data fetch / aggregation may take some time Dashboard has a Queue and processes data requests 1 by 1. Data from survey service can be requested by simple POST API request to a Dashboard service (or a button in future interface).

After aggregation data is available for GET requests.

#### API

For now dashboard service consists only of REST API and connectors, it doesn't have GUI / Web interface yet. Now it is a public API, but in real world it's more safe to guard it with OAuth 2.0 or at least Basic Auth.

API is documented using Swagger UI.

In future it may be better to use GraphQL for dashboard so we can decide which data do we need.

#### Queue

I use queue to process all survey data fetch sequentially. In real world it may help to separate Dashboard API and Dashboard Consumers and scale only Consumers for example if necessary.

In real world I'd use RabbitMQ / Beanstalkd or Cloud-based queue, in this project I simply use Go channels.

I implemented very simple consumers, without retries, statuses updates. It should be more robust in real world, for example if API call to survey service failed - we have to delay this message.

#### Dashboard configuration

- List of survey services: address, name, connector name

The local configuration is stored in `./dashboard/config.json` file and contains 2 local services.

### Storage

For now each service has in-memory storage, I decided not to go with real database just not to overengineer. The real world system must have persistent database or cloud-based storage. Also each service configured to connect to any storage instance.

### Local environment

To orchestrate services locally I used Docker Compose as it's commonly used between developers. However for testing / production environments I'd create a Kubernetes configurations. Depends on the cloud choice I'd setup Continuous Delivery to those environments: AWS EKS + AWS Code Pipeline or GCP Kubernetes Engine + Google Cloud Build. No manual builds / pushes!

#### Run it locally

Requirements:

 - Docker
 - Docker Compose

*Tested with Docker 18.06.0-ce-mac70 (26399)*

```
docker-compose up -d --build
```

Services configured locally via env files
Services:

- [localhost:7771](http://localhost:7771) - Dashboard API
- [localhost:7772](http://localhost:7772) - Gaming Survey Service API
- [localhost:7773](http://localhost:7773) - Online Store Survey Service API

API specs:

- [Dashboard Swagger UI](http://petstore.swagger.io/?url=http://localhost:7771/swagger)
- [Gaming Survey Swagger UI](http://petstore.swagger.io/?url=http://localhost:7772/swagger)
- [Store Survey Swagger UI](http://petstore.swagger.io/?url=http://localhost:7773/swagger)

#### How to verify that it works with Swagger UI

- Go to [Gaming Survey Swagger UI](http://petstore.swagger.io/?url=http://localhost:7772/swagger), submit an answer using `POST /answers`.

Example JSON:
```
{
	"survey_id": 1,
	"values": {"age": "21"},
	"user": "John"
}
```

- Go to [Store Survey Swagger UI](http://petstore.swagger.io/?url=http://localhost:7773/swagger), submit an answer using `POST /answers` using similar JSON from gaming survey service.
- Go to [Dashboard Swagger UI](http://petstore.swagger.io/?url=http://localhost:7771/swagger), request data fetch from survey services using `POST /request`.
- After some time you will be able to get dashboard data using `GET /dashboard`.

Voilà!

### Development

I use Go to implement REST APIs as it is efficient language, easy to distribute via single binary file and type safe, which is crutial for APIs.

#### Testing

Unit tests are executed using TravisCI.

Run them manually:
```
go test -v -race ./...
```

### Monitoring

I'd go with StackDriver if project is deployed to Kubernetes Engine, because it has good integration with different metrics.

### TODO List

- Persistent volume storage instead of in-memory one.
- Separate storages for each survey and dashboard.
- API Auth.
- Use RabbitMQ (or similar technology) for queue management instead of go channels.
- As Survey service may have a lot of data, we need a faster way to request only latest data so we don't check duplicates.
- Use GraphQL for Daashboard API, as we can decide what data do we need.
- Write more Unit tests.
- Setup Prometheus monitoring.