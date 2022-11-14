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
	"github.com/gographics/imagick/imagick"
)

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

func UploadFile(bucket *storage.BucketHandle, object string, decodedImage multipart.File) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	wc := bucket.Object(object).NewWriter(ctx)

	imagick.Initialize()
	defer imagick.Terminate()
	// 下地
	mw1 := imagick.NewMagickWand()
	defer mw1.Destroy()
	
	b,err:=io.ReadAll(decodedImage)
	if err!=nil{
		return err
	}

	err = mw1.ReadImageBlob(b)
	if err != nil {
		return err
	}
	err = mw1.SetFormat("webp")
    if err != nil {
        return err
    }
	if _, err =wc.Write(mw1.GetImageBlob()); err != nil {
		fmt.Errorf("io.Copy:%v", err)
		return err
	}
	if err := wc.Close(); err != nil {
		fmt.Errorf("wc.Close:%v", err)
		return err
	}

	fmt.Printf("Blob %v uploaded \n", object)

	return nil
}
