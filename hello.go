package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"strings"
)

// Downloads an item from an S3 Bucket in the region configured in the shared config
// or AWS_REGION environment variable.
//
// Usage:
//    go run s3_download.go
func main() {
	bucket := "bucketexamplelalalalalalal"
	key := "TestFile.txt"

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("ASIAQENA3UJ46NRAAE5S", "B+GGbl5ZagBJq1x221HEpRHFa6IoV9gMfolLNaS3", "FwoGZXIvYXdzEM3//////////wEaDGxz/cr6C2fxsowxoCLPAdVPuKxttdqQXIX7j0fzFx55CXcGCemBqcGfTAREERZvhe9Dv8InusRxz4IErwxYrEhL54+EuGH1rIh+JgP6Kb8gP1wU7mczUIl61fK70smTQcVWPv9nMO/c+7URk9dQBRHjOMqPQa7c2NzdLM6Ypw1B+GNcBJs1zlvgF13PAGTWpjEK8dmfhOvNKzC9sXhsfmGpI+pVhVgtqAMhp1cxiuyx+zo5v2VRx/KgsOWAscUDezfn5aLIr8npvkX04IVTiCpltoPBqY9KzeotOp2Qbiizxo76BTIthqLzOK6clc6WjFk4VlRjdBL/UwmPV8Y/x01JDnaJM+YWqV9xc1OAUq7QqyV0"),
	})

	// Create S3 service client
	svc := s3.New(sess)

	_, err = svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: &bucket,
	})
	if err != nil {
		log.Println("Failed to create bucket", err)
		return
	}

	if err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{Bucket: &bucket}); err != nil {
		log.Printf("Failed to wait for bucket to exist %s, %s\n", bucket, err)
		return
	}

	_, err = svc.PutObject(&s3.PutObjectInput{
		Body:   strings.NewReader("Hello World!"),
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		log.Printf("Failed to upload data to %s/%s, %s\n", bucket, key, err)
		return
	}

	log.Printf("Successfully created bucket %s and uploaded data with key %s\n", bucket, key)
}
