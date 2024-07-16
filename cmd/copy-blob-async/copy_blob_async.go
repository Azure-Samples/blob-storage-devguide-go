package main

import (
	"context"
	"log"

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

// <snippet_copy_from_source_url>
func copyFromSourceAsync(srcBlob *blockblob.Client, destBlob *blockblob.Client) {
	// Lease the source blob during copy to prevent other clients from modifying it
	blobLeaseClient, err := lease.NewBlobClient(srcBlob, nil)
	handleError(err)

	_, err = blobLeaseClient.AcquireLease(context.TODO(), int32(15), nil)
	handleError(err)

	// Retrieve the SAS token for the source blob and append it to the URL
	sas := "<sas-token>"
	url := srcBlob.URL() + "?" + sas

	// Set copy options
	copyOptions := blob.StartCopyFromURLOptions{
		Tier: to.Ptr(blob.AccessTierCool),
	}

	// Copy the blob from the source URL to the destination blob
	_, err = destBlob.StartCopyFromURL(context.TODO(), url, &copyOptions)
	handleError(err)

	// Release the lease on the source blob
	_, err = blobLeaseClient.ReleaseLease(context.TODO(), nil)
	handleError(err)
}

// </snippet_copy_from_source_url>

// <snippet_copy_from_external_source>
func copyFromExternalSource(srcURL string, destBlob *blockblob.Client) {
	// Set copy options
	copyOptions := blockblob.UploadBlobFromURLOptions{
		Tier: to.Ptr(blob.AccessTierCool),
	}

	// Copy the blob from the source URL to the destination blob
	_, err := destBlob.UploadBlobFromURL(context.TODO(), srcURL, &copyOptions)
	handleError(err)
}

// </snippet_copy_from_external_source>

func main() {
	// <snippet_copy_from_source_url_usage>
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
	// </snippet_copy_from_source_url_usage>

	// <snippet_copy_from_external_source_usage>
	//externalURL := "<source-url>"

	//destBlob = destClient.ServiceClient().NewContainerClient("destination-container").NewBlockBlobClient("destination-blob-2")

	//copyFromExternalSource(externalURL, destBlob)
	// </snippet_copy_from_external_source_usage>
}
