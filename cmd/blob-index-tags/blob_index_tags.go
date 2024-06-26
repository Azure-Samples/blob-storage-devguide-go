package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
)

// Blob dev guide delete blob sample

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// <snippet_set_blob_tags>
func setBlobTags(client *azblob.Client, containerName string, blobName string) {
	// Reference the blob as a client object
	blobClient := client.ServiceClient().NewContainerClient(containerName).NewBlobClient(blobName)

	// Get existing tags for the blob if they need to be preserved
	resp, err := blobClient.GetTags(context.TODO(), nil)
	handleError(err)
	tags := make(map[string]string)
	for _, v := range resp.BlobTags.BlobTagSet {
		tags[*v.Key] = *v.Value
	}

	// Add or modify blob tags
	var updated_tags = make(map[string]*string)
	updated_tags["tag1"] = to.Ptr("value1")
	updated_tags["tag2"] = to.Ptr("value2")

	// Combine existing tags with new tags
	for k, v := range updated_tags {
		tags[k] = *v
	}

	// Set blob tags
	_, err = blobClient.SetTags(context.TODO(), tags, nil)
	handleError(err)
}

// </snippet_set_blob_tags>

// <snippet_clear_blob_tags>
func clearBlobTags(client *azblob.Client, containerName string, blobName string) {
	// Reference the blob as a client object
	blobClient := client.ServiceClient().NewContainerClient(containerName).NewBlobClient(blobName)

	// Clear blob tags
	_, err := blobClient.SetTags(context.TODO(), make(map[string]string), nil)
	handleError(err)
}

// </snippet_clear_blob_tags>

// <snippet_get_blob_tags>
func getBlobTags(client *azblob.Client, containerName string, blobName string) {
	// Reference the blob as a client object
	blobClient := client.ServiceClient().NewContainerClient(containerName).NewBlobClient(blobName)

	// Get the blob tags
	resp, err := blobClient.GetTags(context.TODO(), nil)
	handleError(err)

	// Print the blob tags
	for _, v := range resp.BlobTags.BlobTagSet {
		fmt.Printf("Key: %v, Value: %v\n", *v.Key, *v.Value)
	}
}

// </snippet_get_blob_tags>

// <snippet_find_blobs_by_tags>
func findBlobsByTags(client *azblob.Client, containerName string, blobName string) {
	// Reference the container as a client object
	containerClient := client.ServiceClient().NewContainerClient(containerName)

	// Filter blobs by tags
	where := "\"Content\"='image'"
	opts := container.FilterBlobsOptions{MaxResults: to.Ptr(int32(10))}
	resp, err := containerClient.FilterBlobs(context.TODO(), where, &opts)
	handleError(err)

	// Print the blobs found
	for _, blobItem := range resp.FilterBlobSegment.Blobs {
		fmt.Printf("Blob name: %v\n", *blobItem.Name)
	}
}

// </snippet_find_blobs_by_tags>

func main() {
	// TODO: replace <storage-account-name> with your actual storage account name
	url := "https://<storage-account-name>.blob.core.windows.net/"

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	containerName := "sample-container"
	blobName := "sample-blob"

	setBlobTags(client, containerName, blobName)
	getBlobTags(client, containerName, blobName)
	findBlobsByTags(client, containerName, blobName)
	clearBlobTags(client, containerName, blobName)
}
