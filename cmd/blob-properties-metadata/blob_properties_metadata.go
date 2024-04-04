package main

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
)

// Blob dev guide properties/metadata sample

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func setBlobProperties(client azblob.Client, containerName string, blobName string) {
	// Reference the blob as a client object
	blobClient := client.ServiceClient().NewContainerClient(containerName).NewBlobClient(blobName)

	// Set the blob properties
	contentType := "text/plain"
	_, err := blobClient.SetHTTPHeaders(context.TODO(), blob.HTTPHeaders{
		BlobContentType: &contentType,
	}, nil)
	handleError(err)
}

func getBlobProperties(client azblob.Client, containerName string, blobName string) {
	// Reference the blob as a client object
	blobClient := client.ServiceClient().NewContainerClient(containerName).NewBlobClient(blobName)

	// Get the blob properties
	resp, err := blobClient.GetProperties(context.TODO(), nil)
	handleError(err)

	// Print the blob properties
	log.Printf("Blob content type: %s", *resp.ContentType)
}

func setBlobMetadata(client azblob.Client, containerName string, blobName string) {

}

func getBlobMetadata(client azblob.Client, containerName string, blobName string) {

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

	setBlobProperties(*client, containerName, blobName)
	getBlobProperties(*client, containerName, blobName)
	setBlobMetadata(*client, containerName, blobName)
	getBlobMetadata(*client, containerName, blobName)
}
