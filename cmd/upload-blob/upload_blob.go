package main

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// Blob dev guide upload sample

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func uploadBlobFile(client *azblob.Client, containerName string, blobName string) {
	// Open the file for reading
	file, err := os.OpenFile("path/to/sample/file", os.O_RDONLY, 0)
	handleError(err)

	defer file.Close()

	// Upload the file to the specified container with the specified blob name
	_, err = client.UploadFile(context.TODO(), containerName, blobName, file, nil)
	handleError(err)
}

func uploadBlobBuffer(client *azblob.Client, containerName string, blobName string) {
	// Create a buffer with the content of the file to upload
	data := []byte("Hello, world!")

	// Upload the data to a block blob
	_, err := client.UploadBuffer(context.TODO(), containerName, blobName, data, nil)
	handleError(err)
}

func main() {
	// TODO: replace <storage-account-name> with your actual storage account name
	url := "https://<storage-account-name>.blob.core.windows.net/"

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	containerName := "sample-container"
	blobName := "sample-blob"

	uploadBlobFile(client, containerName, blobName)
	uploadBlobBuffer(client, containerName, blobName)
}
