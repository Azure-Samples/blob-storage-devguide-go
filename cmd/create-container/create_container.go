package main

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// Blob dev guide create container sample

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func createContainer(client *azblob.Client, containerName string) {
	// Create a container
	_, err := client.CreateContainer(context.TODO(), containerName, nil)
	handleError(err)
}

func createRootContainer(client *azblob.Client) {
	// Create root container
	_, err := client.CreateContainer(context.TODO(), "$root", nil)
	handleError(err)
}

// Replace the existing main function with the corrected code
func main() {
	// TODO: replace <storage-account-name> with your actual storage account name
	url := "https://<storage-account-name>.blob.core.windows.net/"

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	containerName := "sample-container"

	createContainer(client, containerName)
	createRootContainer(client)
}
