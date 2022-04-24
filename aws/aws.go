/*
package s3 provides functions for transfering data to/from
aws s3, as streams of Bytes
*/
package aws

import (
	"bytes"
	"context"
	"io"
	"locket/cli"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

/* UploadToS3 uploads data to an s3 bucket */
func UploadToS3(data *bytes.Buffer, bucket, profile, key string, metadata map[string]string) {

	client, err := getClient(profile)

	if err != nil {
		log.Fatalf("Error uploading to s3: %s", err)
	}

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

func DownloadFromS3(bucket, profile, key string) io.ReadCloser {
	client, err := getClient(profile)
	if err != nil {
		log.Fatalf("Error downloading from s3: %s", err)
	}
	output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)})
	if err != nil {
		log.Fatalf("Could not retrieve backup from s3: %s", err)
	}
	cli.Prompt("$$")

	return output.Body
}

/*
ListTopLevelObject returns only objects at the root
of an s3 bucket, so that the list is similar to a
unix command (top level directories and files but
not the contents of said directories)
*/
func ListObjects(bucket, profile string) []types.Object {
	client, err := getClient(profile)

	if err != nil {
		log.Fatalf("Error listing recent backups: %s", err)
	}

	objs, err := client.ListObjectsV2(
		context.TODO(),
		&s3.ListObjectsV2Input{Bucket: aws.String(bucket)},
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
