#!/bin/sh

echo "Aguardando o Localstack iniciar..."
sleep 15  # Espera simples para garantir que o Localstack esteja iniciado

# Criar recursos
echo "Criando recursos no Localstack..."
awslocal iam create-role \
    --role-name localstack-role \
    --assume-role-policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":"arn:aws:iam::000000000000:root"},"Action":"sts:AssumeRole"}]}'

awslocal iam attach-role-policy \
    --role-name localstack-role \
    --policy-arn arn:aws:iam::aws:policy/AdministratorAccess

awslocal --endpoint-url=http://localhost:4566 sqs create-queue --queue-name my-queue-payment
awslocal --endpoint-url=http://localhost:4566 sqs create-queue --queue-name my-queue-invoice
awslocal --endpoint-url=http://localhost:4566 sqs create-queue --queue-name my-queue-deposit
awslocal --endpoint-url=http://localhost:4566 sqs create-queue --queue-name my-queue-withdrawal

echo "Recursos criados."