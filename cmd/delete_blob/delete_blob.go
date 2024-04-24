package main

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
)

// Blob dev guide delete blob sample

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func deleteBlob(client *azblob.Client, containerName string, blobName string) {
	// Delete the blob
	_, err := client.DeleteBlob(context.TODO(), containerName, blobName, nil)
	handleError(err)
}

func deleteBlobWithSnapshots(client *azblob.Client, containerName string, blobName string) {
	// Delete the blob and its snapshots
	_, err := client.DeleteBlob(context.TODO(), containerName, blobName, &blob.DeleteOptions{
		DeleteSnapshots: to.Ptr(blob.DeleteSnapshotsOptionTypeInclude),
	})
	handleError(err)
}

func restoreDeletedBlob(client *azblob.Client, containerName string, blobName string) {
	// Reference the blob as a client object
	blobClient := client.ServiceClient().NewContainerClient(containerName).NewBlobClient(blobName)

	// Restore the deleted blob
	_, err := blobClient.Undelete(context.TODO(), &blob.UndeleteOptions{})
	handleError(err)
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

	deleteBlob(client, containerName, blobName)
	deleteBlobWithSnapshots(client, containerName, blobName)
	restoreDeletedBlob(client, containerName, blobName)
}
