# fizz - pluggable email delivery/notifications microservice written in Go

Email sending is fairly general thing in the development world. It's often repetitive and time consuming to write it again and again.
This microservice is stateless and pluggable that can be hooked up into any microservices.

Written in Go, with optional Postgres persistence layer [WIP], you can deploy it to your Kubernetes cluster

This repo also features a complete terraform suite to provision a blank, fresh EKS cluster that can be used to deploy your services.

### WIP

1. Github Actions CI/CD Pipeline
2. Unit testing and test coverage
3. Helm charts

### Basic Features

1. Send email
2. Mailgun Backend
3. REST API
4. Docker and Kubernetes Deployment
5. Terraform AWS
6. Terraform EKS Deployment

### Extended features

1. Reliability: retry, timeout, fallback, panic recover, rollback
2. webhook events
3. database
4. update db according to webhook events
5. support for multiple backend
6. HTML Email
7. email sending status lookup
8. Observability: Metrics, Logging, Traces
9. Writing Tests
