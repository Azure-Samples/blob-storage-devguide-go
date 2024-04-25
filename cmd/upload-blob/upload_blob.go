package main

import (
	"context"
	"log"
	"os"
	"strings"

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

func uploadBlobStream(client *azblob.Client, containerName string, blobName string) {
	data := "Hello, world!"
	blobContentReader := strings.NewReader(data)

	// Upload the file to the specified container with the specified blob name
	_, err := client.UploadStream(context.TODO(), containerName, blobName, blobContentReader, nil)
	handleError(err)
}

func uploadBlobWithIndexTags(client *azblob.Client, containerName string, blobName string) {
	// Create a buffer with the content of the file to upload
	data := []byte("Hello, world!")

	// Upload the data to a block blob with index tags
	_, err := client.UploadBuffer(context.TODO(), containerName, blobName, data, &azblob.UploadBufferOptions{
		Tags: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	})
	handleError(err)
}

func uploadBlobWithTransferOptions(client *azblob.Client, containerName string, blobName string) {
	// Open the file for reading
	file, err := os.OpenFile("path/to/sample/file", os.O_RDONLY, 0)
	handleError(err)

	defer file.Close()

	// Upload the data to a block blob with transfer options
	_, err = client.UploadFile(context.TODO(), containerName, blobName, file,
		&azblob.UploadFileOptions{
			BlockSize:   int64(4 * 1024 * 1024), // 4 MiB
			Concurrency: uint16(2),
		})
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
	uploadBlobStream(client, containerName, blobName)
	uploadBlobWithIndexTags(client, containerName, blobName)
	uploadBlobWithTransferOptions(client, containerName, blobName)
}
