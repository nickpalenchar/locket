/*
package s3 provides functions for transfering data to/from
aws s3, as streams of Bytes
*/
package aws

import (
	"context"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Client struct {
	Bucket *string
	Client s3.Client
}

func NewS3Client(profile, bucket string) S3Client {
	client, err := getClient(profile)
	if err != nil {
		log.Fatalf("Cannot create s3 client: %s", err)
	}
	return S3Client{
		Client: *client,
		Bucket: aws.String(bucket),
	}
}

func (s *S3Client) Upload(data io.Reader, key string, metadata map[string]string) error {
	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:   s.Bucket,
		Key:      aws.String(key),
		Body:     data,
		Metadata: metadata,
	})

	return err
}

func (s *S3Client) Download(key string) io.ReadCloser {
	output, err := s.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: s.Bucket,
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatalf("Could not retrieve backup from s3: %s", err)
	}
	return output.Body
}

func (s *S3Client) List() []types.Object {
	objs, err := s.Client.ListObjectsV2(
		context.TODO(),
		&s3.ListObjectsV2Input{Bucket: s.Bucket},
	)

	if err != nil {
		log.Fatalf("Error listing recent backups: %s", err)
	}

	return objs.Contents
}

/*
getClient creates an aws s3 client with default credentials
*/
func getClient(profile string) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile))

	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil

}
