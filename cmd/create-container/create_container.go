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

// <snippet_create_container>
func createContainer(client *azblob.Client, containerName string) {
	// Create a container
	_, err := client.CreateContainer(context.TODO(), containerName, nil)
	handleError(err)
}

// </snippet_create_container>

// <snippet_create_root_container>
func createRootContainer(client *azblob.Client) {
	// Create root container
	_, err := client.CreateContainer(context.TODO(), "$root", nil)
	handleError(err)
}

// </snippet_create_root_container>

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
