package firebaseOperation

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
)

type Image_file struct {
}

func UseDefaultBacket() *storage.BucketHandle {
	config := &firebase.Config{
		StorageBucket: os.Getenv("BUCKET_NAME") + ".appspot.com",
	}
	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	return bucket
}

// 'bucket' is an object defined in the cloud.google.com/go/storage package.
// See https://godoc.org/cloud.google.com/go/storage#BucketHandle
// for more details.storage.go

func UploadFile(bucket *storage.BucketHandle, object string,  decodedImage multipart.File) error {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	wc := bucket.Object(object).NewWriter(ctx)

	if _, err = io.Copy(wc, decodedImage); err != nil {
		fmt.Errorf("io.Copy:%v", err)
	}
	if err := wc.Close(); err != nil {
		fmt.Errorf("wc.Close:%v", err)
	}

	fmt.Printf("Blob %v uploaded \n", object)

	return nil
}
