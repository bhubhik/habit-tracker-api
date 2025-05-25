package worker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

const (
	queueUrl   = "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/habit-updates"
	bucketName = "habit-feeds"
	region     = "us-east-1"
	localUrl   = "http://localhost:4566"
)

func Start() {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	sqsClient := sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		o.BaseEndpoint = aws.String(localUrl)
	})
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = aws.String(localUrl)
	})

	fmt.Println("Worker started, listening for messages...")

	for {
		out, err := sqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueUrl),
			MaxNumberOfMessages: 1,
			WaitTimeSeconds:     5,
		})
		if err != nil {
			log.Printf("Received error: %v\n", err)
			continue
		}

		for _, msg := range out.Messages {
			fmt.Printf("Received: %s\n", *msg.Body)

			key := *msg.Body

			obj, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String(key),
			})
			if err != nil {
				log.Printf("error gettting object from s3: %v", err)
			} else {
				fmt.Printf("Fetched file: %s\n", key)
				//TODO: Process the content and save to DB
				obj.Body.Close()
			}

			//Delete message after processing
			_, _ = sqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueUrl),
				ReceiptHandle: msg.ReceiptHandle,
			})
		}
		time.Sleep(1 * time.Second)
	}
}
