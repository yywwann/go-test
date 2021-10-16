package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	endpoint := "127.0.0.1:9000"
	accessKeyID := "UBAFOWP15F5G3E5T063Q"
	secretAccessKey := "75KTn55luxsHKD7mFW+y0JqB5DkkLrKpMlKqX7e2"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.NewCore(
		endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "dtalk-test"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the zip file
	key := "2021_03_04_22_33_29.jpg"
	filePath := "/Users/cccccccccchy/Pictures/壁纸now/已备份/2021_03_04_22_33_29.jpg"
	//contentType := "application/zip"
	f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fstats, err := f.Stat()
	if err != nil {
		panic(err)
	}

	res, err := minioClient.PutObject(ctx, bucketName, key, f, fstats.Size(), "", "", minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	// Upload the zip file with FPutObject
	//info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	//if err != nil {
	//	log.Fatalln(err)
	//}

	log.Printf("Successfully uploaded %s of size %d\n", key, res.Size)
}
