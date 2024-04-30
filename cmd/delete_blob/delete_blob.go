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

// <snippet_delete_blob>
func deleteBlob(client *azblob.Client, containerName string, blobName string) {
	// Delete the blob
	_, err := client.DeleteBlob(context.TODO(), containerName, blobName, nil)
	handleError(err)
}

// </snippet_delete_blob>

// <snippet_delete_blob_snapshots>
func deleteBlobWithSnapshots(client *azblob.Client, containerName string, blobName string) {
	// Delete the blob and its snapshots
	_, err := client.DeleteBlob(context.TODO(), containerName, blobName, &blob.DeleteOptions{
		DeleteSnapshots: to.Ptr(blob.DeleteSnapshotsOptionTypeInclude),
	})
	handleError(err)
}

// </snippet_delete_blob_snapshots>

// <snippet_restore_blob>
func restoreDeletedBlob(client *azblob.Client, containerName string, blobName string) {
	// Reference the blob as a client object
	blobClient := client.ServiceClient().NewContainerClient(containerName).NewBlobClient(blobName)

	// Restore the deleted blob
	_, err := blobClient.Undelete(context.TODO(), &blob.UndeleteOptions{})
	handleError(err)
}

// </snippet_restore_blob>

// <snippet_restore_blob_version>
func restoreDeletedBlobVersion(client *azblob.Client, containerName string, blobName string, versionID string) {
	// Reference the blob as a client object
	baseBlobClient := client.ServiceClient().NewContainerClient(containerName).NewBlobClient(blobName)

	blobVersionClient, err := baseBlobClient.WithVersionID(versionID)
	handleError(err)

	// Restore the blob version by copying it to the base blob
	_, err = baseBlobClient.StartCopyFromURL(context.TODO(), blobVersionClient.URL(), nil)
	handleError(err)
}

// </snippet_restore_blob_version>

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
	restoreDeletedBlobVersion(client, containerName, blobName, "version-id")
}
