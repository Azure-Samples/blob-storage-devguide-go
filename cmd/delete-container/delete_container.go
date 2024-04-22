package main

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// Blob dev guide delete container sample

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func deleteContainer(client *azblob.Client, containerName string) {
	// Delete the container
	_, err := client.DeleteContainer(context.TODO(), containerName, nil)
	handleError(err)
}

func restoreDeletedContainer(client *azblob.Client, containerName string) {
	// List containers, included deleted ones
	pager := client.NewListContainersPager(&azblob.ListContainersOptions{
		Include: azblob.ListContainersInclude{Deleted: true},
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		for _, container := range resp.ContainerItems {
			if *container.Name == containerName && *container.Deleted {
				// Restore the deleted container
				_, err := client.ServiceClient().RestoreContainer(context.TODO(), containerName, *container.Version, nil)
				handleError(err)
			}
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

	deleteContainer(client, containerName)
	restoreDeletedContainer(client, containerName)
}
