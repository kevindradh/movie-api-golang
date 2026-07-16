package config

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var MinioBucket = "movie-posters"

func ConnectMinio() {
	client, err := minio.New(Env.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(Env.MinioAccessKey, Env.MinioSecretKey, ""),
		Secure: Env.MinioUseSSL,
	})
	if err != nil {
		log.Fatal("MinIO connect fail:", err)
	}

	MinioClient = client
	log.Println("MinIO connect success!")
}
