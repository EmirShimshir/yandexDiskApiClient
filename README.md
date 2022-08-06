# Yandex disk REST API client written in Go

## Installation

```
go get -u github.com/EmirShimshir/yandexDiskApiClient
```

## Description
This Yandex disk REST API client (SDK) will help you simplify the integration of your application with Yandex Disk.

```NewClient(oAuth string, timeout time.Duration) (*Client, error)``` - constructor for new client

```(c *Client) GetDiskInfo(ctx context.Context) (*Disk, error)``` - give you info about your disk

```(c *Client) GetFiles(ctx context.Context, limit int) (*FilesResourceList, error)``` - give you info about your files

```(c *Client) Delete(ctx context.Context, path string, permanently bool) (*SuccessResponse, error)``` - delete file from disk

```(c *Client) Download(ctx context.Context, path string) (*SuccessResponse, error)``` - give you download url for your file in disk

```(c *Client) Upload(ctx context.Context, path string) (*SuccessResponse, error)``` - give you upload url for your file, path in disk

```(c *Client) UploadByURL(ctx context.Context, path string, url string)``` - upload your file by url to your disk

```(c *Client) GetPublicFiles(ctx context.Context, limit int) (*FilesResourceList, error)``` - give you info about your public files

```(c *Client) Publish(ctx context.Context, path string) (*SuccessResponse, error)``` - publish file on your disk

```(c *Client) Unpublish(ctx context.Context, path string) (*SuccessResponse, error)``` - unpublish file on your disk

```(c *Client) Move(ctx context.Context, from string, path string) (*SuccessResponse, error)``` - move your file or folder in disk

```(c *Client) Copy(ctx context.Context, from string, path string) (*SuccessResponse, error)``` - copy your file or folder in disk

```(c *Client) Mkdir(ctx context.Context, path string) (*SuccessResponse, error)``` - make folder in your disk

```(c *Client) GetTrash(ctx context.Context, path string, limit int) (*TrashResourceList, error)``` - give you info about your trash

```(c *Client) ClearTrash(ctx context.Context, path string) (*SuccessResponse, error)``` - clean your trash files, only full trash file path, __"trash:/"__ works

```(c *Client) RestoreTrash(ctx context.Context, path string) (*SuccessResponse, error)``` - restore your trash files, // only full trash path, __"trash:/"__ doesn't work
## Example

```go
package main

import (
	"context"
	"fmt"
	"github.com/EmirShimshir/yandexDiskApiClient"
	"log"
	"time"
)

const oAuth = "YOUR OAUTH HERE"

func main() {
	client, err := yandexDiskApiClient.NewClient(oAuth, 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	disk, err := client.GetDiskInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(disk.Info())
}

```

```
[Sat Aug  6 13:03:50 2022] GET https://cloud-api.yandex.net/v1/disk
UserName: Emir2701
UsedSpace: 1760710963
UsedSpace: 10737418240

```
