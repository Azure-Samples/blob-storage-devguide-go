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

// <snippet_list_blobs_flat>
func listBlobsFlat(client *azblob.Client, containerName string) {
	// List the blobs in the container
	pager := client.NewListBlobsFlatPager(containerName, &azblob.ListBlobsFlatOptions{
		Include: azblob.ListBlobsInclude{Snapshots: true, Versions: true},
	})

	fmt.Println("List blobs flat:")
	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		for _, blob := range resp.Segment.BlobItems {
			fmt.Println(*blob.Name)
		}
	}
}

// </snippet_list_blobs_flat>

// <snippet_list_blobs_flat_options>
func listBlobsFlatOptions(client *azblob.Client, containerName string, prefix string) {
	// List the blobs in the container with a prefix
	pager := client.NewListBlobsFlatPager(containerName, &azblob.ListBlobsFlatOptions{
		Prefix: to.Ptr(prefix),
	})

	fmt.Println("List blobs with prefix:")
	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		for _, blob := range resp.Segment.BlobItems {
			fmt.Println(*blob.Name)
		}
	}
}

// </snippet_list_blobs_flat_options>

// <snippet_list_blobs_hierarchical>
func listBlobsHierarchy(client *azblob.Client, containerName string, prefix string) {
	// Reference the container as a client object
	containerClient := client.ServiceClient().NewContainerClient(containerName)

	pager := containerClient.NewListBlobsHierarchyPager("/", &container.ListBlobsHierarchyOptions{
		Prefix:     to.Ptr(prefix),
		MaxResults: to.Ptr(int32(1)), // MaxResults set to 1 for demonstration purposes
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		if resp.Segment.BlobPrefixes != nil {
			for _, prefix := range resp.Segment.BlobPrefixes {
				fmt.Println("Virtual directory prefix:", *prefix.Name)

				// Recursively list blobs in the prefix
				listBlobsHierarchy(client, containerName, *prefix.Name)
			}
		}

		for _, blob := range resp.Segment.BlobItems {
			fmt.Println("Blob:", *blob.Name)
		}
	}
}

// </snippet_list_blobs_hierarchical>

func main() {
	// TODO: replace <storage-account-name> with your actual storage account name
	url := "https://<storage-account-name>.blob.core.windows.net/"

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	containerName := "sample-container"

	listBlobsFlat(client, containerName)
	listBlobsFlatOptions(client, containerName, "sample")
	listBlobsHierarchy(client, containerName, "")
}
