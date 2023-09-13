package auth

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type StorageClient struct {
	accountName   string
	accountKey    string
	containerName string
}

func NewStorageClient(accountKey string) *StorageClient {
	return &StorageClient{
		accountName:   "projectzenstorage",
		accountKey:    accountKey,
		containerName: "mails",
	}
}

func (client *StorageClient) createContainerURL() (azblob.ContainerURL, error) {
	credential, err := azblob.NewSharedKeyCredential(client.accountName, client.accountKey)
	if err != nil {
		return azblob.ContainerURL{}, err
	}

	// Construct the service URL using the client's account name
	serviceURL, err := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net", client.accountName))
	if err != nil {
		return azblob.ContainerURL{}, err
	}

	pipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	serviceURLWithPipeline := azblob.NewServiceURL(*serviceURL, pipeline)

	return serviceURLWithPipeline.NewContainerURL(client.containerName), nil
}

func (client *StorageClient) UploadTextToBlob(content []byte) {
	// Create a container client
	containerURL, err := client.createContainerURL()
	if err != nil {
		fmt.Printf("Failed to create blob client: %v\n", err)
	}

	// Create a unique file name using the current date and time
	fileName := fmt.Sprintf("mails_%s.json", getCurrentTimestamp())

	// Upload the content to the Blob Storage
	if err := uploadTextToBlob(containerURL, fileName, content); err != nil {
		fmt.Printf("Error uploading file to Azure Blob Storage: %v\n", err)
	}

	fmt.Println("File uploaded successfully!")
}

func getCurrentTimestamp() string {
	return now().Format("20060102_150405")
}

func now() time.Time {
	return time.Now()
}

func uploadTextToBlob(containerURL azblob.ContainerURL, fileName string, content []byte) error {
	blobURL := containerURL.NewBlockBlobURL(fileName)
	ctx := context.Background()

	// Create a reader for the content
	reader := bytes.NewReader(content)

	// Upload the content to the Blob Storage
	_, err := azblob.UploadStreamToBlockBlob(ctx, reader, blobURL, azblob.UploadStreamToBlockBlobOptions{
		BufferSize: 4 * 1024 * 1024, // 4MB buffer size
	})
	return err
}
