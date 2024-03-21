# Go Serverless Terraform Demo

This is a demo application showcasing the usage of Go with the Fiber framework and Terraform to deploy a Lambda function and an API Gateway on AWS.

## Prerequisites

Before getting started, make sure you have the following installed on your machine:

- Go: [Installation Guide](https://golang.org/doc/install)
- Terraform: [Installation Guide](https://learn.hashicorp.com/tutorials/terraform/install-cli)

## Getting Started

To run this demo application locally, follow these steps:

1. Clone the repository:

```shell
git clone https://github.com/your-username/golang-serverless-terraform-demo.git
```

2. Change into the project directory:

```shell
cd golang-serverless-terraform-demo
```

3. Install the project dependencies:

```shell
go mod download
```


## Deployment

To deploy the application on AWS using Terraform, follow these steps:

```shell
make terraform-deploy
```