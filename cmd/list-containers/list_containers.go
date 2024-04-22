package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// Blob dev guide list containers sample

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func listContainers(client *azblob.Client) {
	// List the containers in the storage account
	pager := client.NewListContainersPager(&azblob.ListContainersOptions{})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		for _, container := range resp.ContainerItems {
			fmt.Println(*container.Name)
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

	listContainers(client)
}
