package yandexDiskApiClient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const baseURL = "https://cloud-api.yandex.net/v1/disk"

type Client struct {
	oAuth   string
	baseURl string
	client  *http.Client
}

func NewClient(oAuth string, timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		oAuth:   oAuth,
		baseURl: baseURL,
		client: &http.Client{
			Timeout:       timeout,
			Transport:     transport,
			CheckRedirect: checkRedirect,
		},
	}, nil
}

func (c *Client) sendRequest(req *http.Request, data interface{}) error {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "OAuth "+c.oAuth)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		errorResponse := ErrorResponse{}
		if err = json.NewDecoder(resp.Body).Decode(&errorResponse); err == nil {
			return fmt.Errorf("%sstatus code: %d\n", errorResponse.Info(), resp.StatusCode)
		}

		return fmt.Errorf("unknown error, status code: %d\n", resp.StatusCode)
	}

	json.NewDecoder(resp.Body).Decode(&data)

	return nil
}

func (c *Client) GetDiskInfo(ctx context.Context) (*Disk, error) {
	req, err := http.NewRequest("GET", c.baseURl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	disk := Disk{}

	if err = c.sendRequest(req, &disk); err != nil {
		return nil, err
	}

	return &disk, nil

}

func (c *Client) GetFiles(ctx context.Context, limit int) (*FilesResourceList, error) {
	req, err := http.NewRequest("GET", c.baseURl+"/resources/files", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	filesResourceList := FilesResourceList{}

	if err = c.sendRequest(req, &filesResourceList); err != nil {
		return nil, err
	}

	return &filesResourceList, nil
}

func (c *Client) Delete(ctx context.Context, path string, permanently bool) (*SuccessResponse, error) {
	if path == "" {
		return nil, errors.New("path can't be empty")
	}

	req, err := http.NewRequest("DELETE", c.baseURl+"/resources", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("path", path)
	q.Set("permanently", strconv.FormatBool(permanently))
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	successResponse := SuccessResponse{}

	if err = c.sendRequest(req, &successResponse); err != nil {
		return nil, err
	}

	return &successResponse, nil
}

func (c *Client) Download(ctx context.Context, path string) (*SuccessResponse, error) {
	if path == "" {
		return nil, errors.New("path can't be empty")
	}

	req, err := http.NewRequest("GET", c.baseURl+"/resources/download", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("path", path)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	successResponse := SuccessResponse{}

	if err = c.sendRequest(req, &successResponse); err != nil {
		return nil, err
	}

	return &successResponse, nil
}

func (c *Client) Upload(ctx context.Context, path string) (*SuccessResponse, error) {
	if path == "" {
		return nil, errors.New("path can't be empty")
	}

	req, err := http.NewRequest("GET", c.baseURl+"/resources/upload", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("path", path)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	successResponse := SuccessResponse{}

	if err = c.sendRequest(req, &successResponse); err != nil {
		return nil, err
	}

	return &successResponse, nil
}

func (c *Client) UploadByURL(ctx context.Context, path string, url string) (*SuccessResponse, error) {
	if path == "" || url == "" {
		return nil, errors.New("path or url can't be empty")
	}

	req, err := http.NewRequest("POST", c.baseURl+"/resources/upload", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("url", url)
	q.Set("path", path)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	successResponse := SuccessResponse{}

	if err = c.sendRequest(req, &successResponse); err != nil {
		return nil, err
	}

	return &successResponse, nil
}

func (c *Client) Publish(ctx context.Context, path string) (*SuccessResponse, error) {
	if path == "" {
		return nil, errors.New("path can't be empty")
	}

	req, err := http.NewRequest("PUT", c.baseURl+"/resources/publish", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("path", path)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	successResponse := SuccessResponse{}

	if err = c.sendRequest(req, &successResponse); err != nil {
		return nil, err
	}

	return &successResponse, nil
}

func (c *Client) Unpublish(ctx context.Context, path string) (*SuccessResponse, error) {
	if path == "" {
		return nil, errors.New("path can't be empty")
	}

	req, err := http.NewRequest("PUT", c.baseURl+"/resources/unpublish", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("path", path)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	successResponse := SuccessResponse{}

	if err = c.sendRequest(req, &successResponse); err != nil {
		return nil, err
	}

	return &successResponse, nil
}

func (c *Client) GetPublicFiles(ctx context.Context, limit int) (*FilesResourceList, error) {
	req, err := http.NewRequest("GET", c.baseURl+"/resources/public", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	filesResourceList := FilesResourceList{}

	if err = c.sendRequest(req, &filesResourceList); err != nil {
		return nil, err
	}

	return &filesResourceList, nil
}

func (c *Client) Move(ctx context.Context, from string, path string) (*SuccessResponse, error) {
	if path == "" || from == "" {
		return nil, errors.New("paths can't be empty")
	}

	req, err := http.NewRequest("POST", c.baseURl+"/resources/move", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("from", from)
	q.Set("path", path)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	successResponse := SuccessResponse{}

	if err = c.sendRequest(req, &successResponse); err != nil {
		return nil, err
	}

	return &successResponse, nil
}

func (c *Client) Copy(ctx context.Context, from string, path string) (*SuccessResponse, error) {
	if path == "" || from == "" {
		return nil, errors.New("paths can't be empty")
	}

	req, err := http.NewRequest("POST", c.baseURl+"/resources/copy", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("from", from)
	q.Set("path", path)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	successResponse := SuccessResponse{}

	if err = c.sendRequest(req, &successResponse); err != nil {
		return nil, err
	}

	return &successResponse, nil
}

func (c *Client) Mkdir(ctx context.Context, path string) (*SuccessResponse, error) {
	if path == "" {
		return nil, errors.New("paths can't be empty")
	}

	req, err := http.NewRequest("PUT", c.baseURl+"/resources", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("path", path)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	successResponse := SuccessResponse{}

	if err = c.sendRequest(req, &successResponse); err != nil {
		return nil, err
	}

	return &successResponse, nil
}

func (c *Client) GetTrash(ctx context.Context, path string, limit int) (*TrashResourceList, error) {
	if path == "" {
		return nil, errors.New("path can't be empty")
	}

	req, err := http.NewRequest("GET", c.baseURl+"/trash/resources", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("path", path)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	trashResourceList := TrashResourceList{}

	if err = c.sendRequest(req, &trashResourceList); err != nil {
		return nil, err
	}

	return &trashResourceList, nil
}

// only full trash path, trash:/ works
func (c *Client) ClearTrash(ctx context.Context, path string) (*SuccessResponse, error) {
	if path == "" {
		return nil, errors.New("path can't be empty")
	}

	req, err := http.NewRequest("DELETE", c.baseURl+"/trash/resources", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("path", path)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	successResponse := SuccessResponse{}

	if err = c.sendRequest(req, &successResponse); err != nil {
		return nil, err
	}

	return &successResponse, nil
}

// only full trash path, trash:/ doesn't work
func (c *Client) RestoreTrash(ctx context.Context, path string) (*SuccessResponse, error) {
	if path == "" {
		return nil, errors.New("path can't be empty")
	}

	req, err := http.NewRequest("PUT", c.baseURl+"/trash/resources/restore", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("path", path)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)

	successResponse := SuccessResponse{}

	if err = c.sendRequest(req, &successResponse); err != nil {
		return nil, err
	}

	return &successResponse, nil
}
