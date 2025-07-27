package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/awserr"
// 	"github.com/aws/aws-sdk-go/aws/credentials"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/s3"
// )

func main() {
	// configureSession()
}

// func configureSession() {

// 	sess, err := session.NewSession(&aws.Config{
// 		Region:      aws.String("us-east-2"),
// 	})

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	svc := s3.New(sess)
// 	input := &s3.ListBucketsInput{}

// 	result, err := svc.ListBuckets(input)
// 	if err != nil {
// 		if aerr, ok := err.(awserr.Error); ok {
// 			switch aerr.Code() {
// 			default:
// 				fmt.Println(aerr.Error())
// 			}
// 		} else {
// 			// Print the error, cast err to awserr.Error to get the Code and
// 			// Message from an error.
// 			fmt.Println(err.Error())
// 		}
// 		return
// 	}

// 	fmt.Println(result)

// }
