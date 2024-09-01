# Go Message Broker

## Description

`go-message-broker` is a Go project that consumes messages from an SQS queue. The project uses Localstack to emulate SQS and assume role for authentication.

## Requirements

- [Go](https://golang.org/dl/) 1.16 or higher
- [Docker](https://www.docker.com/get-started) installed
- AWS CLI configured to use assume role

## Setup

### Localstack

1. Make sure Docker is installed and running.

2. Start Localstack using Docker Compose:

    ```sh
    docker-compose up -d
    ```

3. Create a role in Localstack:

    ```sh
    awslocal iam create-role \
        --role-name localstack-role \
        --assume-role-policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":"arn:aws:iam::000000000000:root"},"Action":"sts:AssumeRole"}]}'
    ```
4. Attach the policy to the IAM role

   ```sh
    awslocal iam attach-role-policy \
    --role-name localstack-role \
    --policy-arn arn:aws:iam::aws:policy/AdministratorAccess
    ```

5. Create an SQS queue in Localstack:

    ```sh
    aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name my-queue --profile localstack
    ```

### Go Project

1. Clone the project repository:

    ```sh
    git clone https://github.com/daluzsi/go-message-broker.git
    cd go-message-broker
    ```

2. Install dependencies:

    ```sh
   go mod tidy
    ```

## Usage

To start the message consumer, run:

```sh
go run main.go
