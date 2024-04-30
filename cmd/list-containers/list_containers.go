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

// <snippet_list_containers>
func listContainers(client *azblob.Client) {
	// List the containers in the storage account and include metadata
	pager := client.NewListContainersPager(&azblob.ListContainersOptions{
		Include: azblob.ListContainersInclude{Metadata: true},
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		for _, container := range resp.ContainerItems {
			fmt.Println(*container.Name)
			for k, v := range container.Metadata {
				fmt.Printf("%v: %v\n", k, *v)
			}
		}
	}
}

// </snippet_list_containers>

// <snippet_list_containers_prefix>
func listContainersWithPrefix(client *azblob.Client, prefix string) {
	// List the containers in the storage account with a prefix
	pager := client.NewListContainersPager(&azblob.ListContainersOptions{
		Prefix: &prefix,
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		for _, container := range resp.ContainerItems {
			fmt.Println(*container.Name)
		}
	}
}

// </snippet_list_containers_prefix>

// <snippet_list_containers_pages>
func listContainersWithMaxResults(client *azblob.Client, maxResults int32) {
	// List the containers in the storage account with a maximum number of results
	pager := client.NewListContainersPager(&azblob.ListContainersOptions{
		MaxResults: &maxResults,
	})

	i := 0
	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		// Show page number to demonstrate pagination with max results
		i++
		fmt.Printf("Page %d:\n", i)

		for _, container := range resp.ContainerItems {
			fmt.Println(*container.Name)
		}
	}
}

// </snippet_list_containers_pages>

func main() {
	// TODO: replace <storage-account-name> with your actual storage account name
	url := "https://<storage-account-name>.blob.core.windows.net/"

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	listContainers(client)
	listContainersWithPrefix(client, "sample")
	listContainersWithMaxResults(client, 2)
}
