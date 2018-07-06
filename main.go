package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"net/url"
	"os"
)

func printUsageAndExit() {
	fmt.Fprintf(os.Stderr, "USAGE: %s s3://YOUT_BUCKET/path/to/object\n", os.Args[0])
	os.Exit(1)
}

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

func main() {
	if len(os.Args) < 2 {
		printUsageAndExit()
	}
	u, err := url.Parse(os.Args[1])
	if err != nil || u.Scheme != "s3" {
		printUsageAndExit()
	}

	svc := s3.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))

	ret, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(u.Host),
		Key:    aws.String(u.Path),
	})
	exitIfError(err)

	_, err = io.Copy(os.Stdout, ret.Body)
	exitIfError(err)
}
