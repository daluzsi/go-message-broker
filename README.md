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

4. Configure AWS CLI to use Localstack:

    ```sh
    aws configure set aws_access_key_id test --profile localstack
    aws configure set aws_secret_access_key test --profile localstack
    aws configure set region us-east-1 --profile localstack
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

2. Set the environment variables:

    ```sh
    export AWS_REGION=us-east-1
    export AWS_ENDPOINT=http://localhost:4566
    export QUEUE_URL=http://localhost:4566/000000000000/my-queue
    export AWS_ROLE_ARN=arn:aws:iam::000000000000:role/localstack-role
    export AWS_ROLE_SESSION_NAME=session-name
    ```

3. Install dependencies:

    ```sh
    go mod tidy
    ```

## Usage

To start the message consumer, run:

```sh
go run main.go
