#!/bin/bash

# Create S3 bucket
awslocal s3 mb s3://habit-feeds
echo "S3 bucket habit-feeds created"

#create SQS queue
awslocal sqs create-queue --queue-name habit-updates
echo "SQS Created"
