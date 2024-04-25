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

// Blob dev guide list blobs sample

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func listBlobsFlat(client *azblob.Client, containerName string) {
	// List the blobs in the container
	pager := client.NewListBlobsFlatPager(containerName, &azblob.ListBlobsFlatOptions{
		Include: azblob.ListBlobsInclude{Snapshots: true, Versions: true},
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		for _, blob := range resp.Segment.BlobItems {
			fmt.Println(*blob.Name)
		}
	}
}

func listBlobsHierarchy(client *azblob.Client, containerName string, prefix string) {
	// Reference the container as a client object
	containerClient := client.ServiceClient().NewContainerClient(containerName)

	pager := containerClient.NewListBlobsHierarchyPager("/", &container.ListBlobsHierarchyOptions{
		Include: container.ListBlobsInclude{Metadata: true},
		Prefix:  to.Ptr(prefix),
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		// Check to see if the result is a prefix or a blob
		for _, prefix := range resp.Segment.BlobPrefixes {
			fmt.Println("Prefix:", *prefix.Name)

			// Recursively list blobs in the prefix
			listBlobsHierarchy(client, containerName, *prefix.Name)
		}

		for _, blob := range resp.Segment.BlobItems {
			fmt.Println("Blob:", *blob.Name)
		}
	}
}

func main() {
	// TODO: replace <storage-account-name> with your actual storage account name
	url := "https://<storage-account-name>.blob.core.windows.net/"

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	containerName := "sample-container"

	listBlobsFlat(client, containerName)
	listBlobsHierarchy(client, containerName, "/")
}
