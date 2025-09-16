package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3bucket = "bucketgofrank"
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials("", "", ""),
		},
	)
	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
}

func main() {
	dir, err := os.Open("/home/franklin/repositorios/PosGolang/Aulas/15-Uploads/tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 100)
	errorFileUpload := make(chan string, 10)

	go func() {
		for filename := range errorFileUpload {
			uploadControl <- struct{}{}
			wg.Add(1)
			go UploadFile(filename, uploadControl, errorFileUpload)
		}
	}()

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("Failed to read directory, %v", err)
			continue
		}

		wg.Add(1)
		uploadControl <- struct{}{}
		go UploadFile(files[0].Name(), uploadControl, errorFileUpload)

	}
	wg.Wait()

}

func UploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()

	completeFileName := fmt.Sprintf("/home/franklin/repositorios/PosGolang/Aulas/15-Uploads/tmp/%s", filename)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Failed to open file %q, %v", completeFileName, err)
		<-uploadControl
		errorFileUpload <- completeFileName
		return
	}
	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Failed to upload file %q, %v", filename, err)
		<-uploadControl
		errorFileUpload <- completeFileName
		return
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, s3bucket)
	<-uploadControl
}
