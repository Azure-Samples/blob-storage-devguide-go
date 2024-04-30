package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
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

// <snippet_set_blob_properties>
func setBlobProperties(client *azblob.Client, containerName string, blobName string) {
	// Reference the blob as a client object
	blobClient := client.ServiceClient().NewContainerClient(containerName).NewBlobClient(blobName)

	// Get the existing blob properties
	resp, err := blobClient.GetProperties(context.TODO(), nil)
	handleError(err)

	// Set the new blob properties and include existing properties
	_, err = blobClient.SetHTTPHeaders(context.TODO(), blob.HTTPHeaders{
		BlobContentType:        to.Ptr("text/plain"),
		BlobContentLanguage:    to.Ptr("en-us"),
		BlobContentEncoding:    resp.ContentEncoding,
		BlobContentDisposition: resp.ContentDisposition,
		BlobCacheControl:       resp.CacheControl,
	}, nil)
	handleError(err)
}

// </snippet_set_blob_properties>

// <snippet_get_blob_properties>
func getBlobProperties(client *azblob.Client, containerName string, blobName string) {
	// Reference the blob as a client object
	blobClient := client.ServiceClient().NewContainerClient(containerName).NewBlobClient(blobName)

	// Get the blob properties
	resp, err := blobClient.GetProperties(context.TODO(), nil)
	handleError(err)

	// Print the blob properties
	fmt.Printf("Content type: %v\n", *resp.ContentType)
	fmt.Printf("Content language: %v\n", *resp.ContentLanguage)
}

// </snippet_get_blob_properties>

// <snippet_set_blob_metadata>
func setBlobMetadata(client *azblob.Client, containerName string, blobName string) {
	// Reference the blob as a client object
	blobClient := client.ServiceClient().NewContainerClient(containerName).NewBlobClient(blobName)

	// Set the blob metadata
	var metadata = make(map[string]*string)
	metadata["key1"] = to.Ptr("value1")
	metadata["key2"] = to.Ptr("value2")

	_, err := blobClient.SetMetadata(context.TODO(), metadata, nil)
	handleError(err)
}

// </snippet_set_blob_metadata>

// <snippet_get_blob_metadata>
func getBlobMetadata(client *azblob.Client, containerName string, blobName string) {
	// Reference the blob as a client object
	blobClient := client.ServiceClient().NewContainerClient(containerName).NewBlobClient(blobName)

	// Get the blob properties, which includes metadata
	resp, err := blobClient.GetProperties(context.TODO(), nil)
	handleError(err)

	// Print the blob metadata
	for k, v := range resp.Metadata {
		fmt.Printf("%v: %v\n", k, *v)
	}
}

// </snippet_get_blob_metadata>

func main() {
	// TODO: replace <storage-account-name> with your actual storage account name
	url := "https://<storage-account-name>.blob.core.windows.net/"

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	containerName := "sample-container"
	blobName := "sample-blob"

	setBlobProperties(client, containerName, blobName)
	getBlobProperties(client, containerName, blobName)
	setBlobMetadata(client, containerName, blobName)
	getBlobMetadata(client, containerName, blobName)
}
