package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// Blob dev guide download sample

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// <snippet_download_blob_stream>
func downloadBlobToStream(client *azblob.Client, containerName string, blobName string) {
	// Download the blob
	get, err := client.DownloadStream(context.TODO(), containerName, blobName, nil)
	handleError(err)

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(context.TODO(), &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	handleError(err)

	err = retryReader.Close()
	handleError(err)

	// Print the contents of the blob we created
	fmt.Println("Blob contents:")
	fmt.Println(downloadedData.String())
}

// </snippet_download_blob_stream>

// <snippet_download_blob_file>
func downloadBlobToFile(client *azblob.Client, containerName string, blobName string) {
	// Create or open a local file where we can download the blob
	file, err := os.Create("path/to/sample/file")
	handleError(err)

	// Download the blob to the local file
	_, err = client.DownloadFile(context.TODO(), containerName, blobName, file, nil)
	handleError(err)
}

// </snippet_download_blob_file>

// <snippet_download_blob_transfer_options>
func downloadBlobTransferOptions(client *azblob.Client, containerName string, blobName string) {
	// Create or open a local file where we can download the blob
	file, err := os.Create("path/to/sample/file")
	handleError(err)

	// Download the blob to the local file
	_, err = client.DownloadFile(context.TODO(), containerName, blobName, file,
		&azblob.DownloadFileOptions{
			BlockSize:   int64(4 * 1024 * 1024), // 4 MiB
			Concurrency: uint16(2),
		})
	handleError(err)
}

// </snippet_download_blob_transfer_options>

func main() {

	// TODO: replace <storage-account-name> with your actual storage account name
	url := "https://<storage-account-name>.blob.core.windows.net/"

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	containerName := "sample-container"
	blobName := "sample-blob.txt"

	downloadBlobToStream(client, containerName, blobName)
	downloadBlobToFile(client, containerName, blobName)
	downloadBlobTransferOptions(client, containerName, blobName)
}
