package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// Blob dev guide client auth sample

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getServiceClientTokenCredential(accountURL string) *azblob.Client {
	// Create a new service client with token credential
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(accountURL, credential, nil)
	handleError(err)

	return client
}

func getServiceClientSharedKey(accountName string, accountKey string) *azblob.Client {
	// Create a new service client with shared key credential
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	handleError(err)

	accountURL := fmt.Sprintf("https://%s.blob.core.windows.net", accountName)

	client, err := azblob.NewClientWithSharedKeyCredential(accountURL, credential, nil)
	handleError(err)

	return client
}

func getServiceClientConnectionString(connectionString string) *azblob.Client {
	// Create a new service client with connection string
	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	handleError(err)

	return client
}

func getServiceClientSAS(accountURL string, sasToken string) *azblob.Client {
	// Create a new service client with an existing SAS token

	// Append the SAS to the account URL with a "?" delimiter
	accountURLWithSAS := fmt.Sprintf("%s?%s", accountURL, sasToken)

	client, err := azblob.NewClientWithNoCredential(accountURLWithSAS, nil)
	handleError(err)

	return client
}

func main() {
	// TODO: Replace placeholders with your actual values for testing purposes
	accountName := "<storage-account-name>"
	accountKey := "<storage-account-key>"
	connectionString := "<connection-string>"
	sasToken := "<sas-token>"
	accountURL := fmt.Sprintf("https://%s.blob.core.windows.net", accountName)

	clientToken := getServiceClientTokenCredential(accountURL)

	clientSharedKey := getServiceClientSharedKey(accountName, accountKey)

	clientConnectionString := getServiceClientConnectionString(connectionString)

	clientSAS := getServiceClientSAS(accountURL, sasToken)

	// Test an operation with the clients
	clientToken.CreateContainer(context.TODO(), "sample-container1", nil)
	clientSharedKey.CreateContainer(context.TODO(), "sample-container2", nil)
	clientConnectionString.CreateContainer(context.TODO(), "sample-container3", nil)
	clientSAS.CreateContainer(context.TODO(), "sample-container4", nil)
}
