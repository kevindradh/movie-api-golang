package config

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var MinioBucket = "movie-posters"

func ConnectMinio() {
	endpoint := "localhost:9000"
	accessKey := "minioadmin"
	secretKey := "minioadmin"
	useSSL := false

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal("MinIO connect fail:", err)
	}

	MinioClient = client
	log.Println("MinIO connect success!")
}
