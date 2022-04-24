/*
package s3 provides functions for transfering data to/from
aws s3, as streams of Bytes
*/
package aws

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

/* UploadToS3 uploads data to an s3 bucket */
func UploadToS3(data *bytes.Buffer, bucket, profile, key string, metadata map[string]string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile))

	if err != nil {
		log.Fatalf("Error uploading to s3: %s", err)
	}

	client := s3.NewFromConfig(cfg)

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:   aws.String(bucket),
		Key:      aws.String(key),
		Body:     data,
		Metadata: metadata,
	})

	if err != nil {
		log.Fatalf("Error while uploading to s3: %s", err)
	}
}
