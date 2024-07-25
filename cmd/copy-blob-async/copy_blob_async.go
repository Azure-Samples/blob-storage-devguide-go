package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/lease"
)

// Blob dev guide upload sample

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// <snippet_copy_from_source_async>
func copyFromSourceAsync(srcBlob *blockblob.Client, destBlob *blockblob.Client) {
	// Lease the source blob during copy to prevent other clients from modifying it
	blobLeaseClient, err := lease.NewBlobClient(srcBlob, nil)
	handleError(err)

	_, err = blobLeaseClient.AcquireLease(context.TODO(), int32(60), nil)
	handleError(err)

	// Retrieve the SAS token for the source blob and append it to the URL
	sas := "<sas-token>"
	url := srcBlob.URL() + "?" + sas

	// Set copy options
	copyOptions := blob.StartCopyFromURLOptions{
		Tier: to.Ptr(blob.AccessTierCool),
	}

	// Copy the blob from the source URL to the destination blob
	startCopy, err := destBlob.StartCopyFromURL(context.TODO(), url, &copyOptions)
	handleError(err)

	// If startCopy.CopyStatus returns a status of "pending", the operation has started asynchronously
	// You can optionally add logic to poll the copy status and wait for the operation to complete
	// Example:
	copyStatus := *startCopy.CopyStatus
	for copyStatus == blob.CopyStatusTypePending {
		time.Sleep(time.Second * 2)

		properties, err := destBlob.GetProperties(context.TODO(), nil)
		handleError(err)

		copyStatus = *properties.CopyStatus
	}

	// Release the lease on the source blob
	_, err = blobLeaseClient.ReleaseLease(context.TODO(), nil)
	handleError(err)
}

// </snippet_copy_from_source_async>

// <snippet_copy_from_external_source_async>
func copyFromExternalSourceAsync(srcURL string, destBlob *blockblob.Client) {
	// Set copy options
	copyOptions := blob.StartCopyFromURLOptions{
		Tier: to.Ptr(blob.AccessTierCool),
	}

	// Copy the blob from the source URL to the destination blob
	startCopy, err := destBlob.StartCopyFromURL(context.TODO(), srcURL, &copyOptions)
	handleError(err)

	// If startCopy.CopyStatus returns a status of "pending", the operation has started asynchronously
	// You can optionally add logic to poll the copy status and wait for the operation to complete
	// Example:
	copyStatus := *startCopy.CopyStatus
	for copyStatus == blob.CopyStatusTypePending {
		time.Sleep(time.Second * 2)

		properties, err := destBlob.GetProperties(context.TODO(), nil)
		handleError(err)

		copyStatus = *properties.CopyStatus
	}
}

// </snippet_copy_from_external_source_async>

// <snippet_check_copy_status>
func checkCopyStatus(destBlob *blockblob.Client) {
	// Retrieve the properties from the destination blob
	properties, err := destBlob.GetProperties(context.TODO(), nil)
	handleError(err)

	copyID := *properties.CopyID
	copyStatus := *properties.CopyStatus

	fmt.Printf("Copy operation %s is %s\n", copyID, copyStatus)
}

// </snippet_check_copy_status>

// <snippet_abort_copy>
func abortCopy(destBlob *blockblob.Client) {
	// Retrieve the copy ID from the destination blob
	properties, err := destBlob.GetProperties(context.TODO(), nil)
	handleError(err)

	copyID := *properties.CopyID
	copyStatus := *properties.CopyStatus

	// Abort the copy operation if it's still pending
	if copyStatus == blob.CopyStatusTypePending {
		_, err := destBlob.AbortCopyFromURL(context.TODO(), copyID, nil)
		handleError(err)

		fmt.Printf("Copy operation %s aborted\n", copyID)
	}
}

// </snippet_abort_copy>

func main() {
	// <snippet_copy_from_source_async_usage>
	// TODO: replace <storage-account-name> placeholders with actual storage account names
	srcURL := "https://<src-storage-account-name>.blob.core.windows.net/"
	destURL := "https://<dest-storage-account-name>.blob.core.windows.net/"

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	srcClient, err := azblob.NewClient(srcURL, credential, nil)
	handleError(err)
	destClient, err := azblob.NewClient(destURL, credential, nil)
	handleError(err)

	srcBlob := srcClient.ServiceClient().NewContainerClient("source-container").NewBlockBlobClient("source-blob")
	destBlob := destClient.ServiceClient().NewContainerClient("destination-container").NewBlockBlobClient("destination-blob-1")

	copyFromSourceAsync(srcBlob, destBlob)
	// </snippet_copy_from_source_async_usage>

	// <snippet_copy_from_external_source_async_usage>
	externalURL := "<source-url>"

	destBlob = destClient.ServiceClient().NewContainerClient("destination-container").NewBlockBlobClient("destination-blob-2")

	copyFromExternalSourceAsync(externalURL, destBlob)
	// </snippet_copy_from_external_source_async_usage>
}
