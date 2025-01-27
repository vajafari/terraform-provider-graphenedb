package client

import (
	"context"
	"fmt"
)

type ResourceClient interface {
	CreateResource(requestBody interface{}, responseBody interface{}) error
	RestartResource(requestPathId string, responseBody interface{}) error
	CreateResourceWithPathExt(requestPathId string, requestBody interface{}, responseBody interface{}) error
	ModifyResourceWithPathExt(requestPathId string, requestBody interface{}, responseBody interface{}) error
	FetchResource(requestPathId string, responseBody interface{}) error
	GetResourceInfo(ctx context.Context, requestPath string, responseBody interface{}) error
	ModifyResource(ctx context.Context,  requestPath string, requestBody interface{}, responseBody interface{}) error
	DeleteResource(ctx context.Context,  requestPath string, requestBody interface{}, responseBody interface{}) error
}

type DefaultResourceClient struct {
	*AuthenticatedClient
	ResourceDescription string
	ResourceRootPath    string
}


func (c *DefaultResourceClient) RestartResource(requestPathExt string, responseBody interface{}) error {
	path_w_id := fmt.Sprintf("%s/%s/restart", c.ResourceRootPath, requestPathExt)
	request, err := c.newAuthenticatedPutRequest(path_w_id, nil)
	if err != nil {
		return err
	}

	response, err := c.requestAndCheckStatus(fmt.Sprintf("create %s", c.ResourceDescription), request)
	if err != nil {
		return err
	}

	if response.StatusCode == 200 {
		return nil
	}

	return unmarshalResponseBody(response, responseBody)
}

func (c *DefaultResourceClient) CreateResource(requestBody interface{}, responseBody interface{}) error {
	request, err := c.newAuthenticatedPostRequest(c.ResourceRootPath, requestBody)
	if err != nil {
		return err
	}

	response, err := c.requestAndCheckStatus(fmt.Sprintf("create %s", c.ResourceDescription), request)
	if err != nil {
		return err
	}

	return unmarshalResponseBody(response, responseBody)
}

func (c *DefaultResourceClient) CreateResourceWithPathExt(requestPathExt string, requestBody interface{}, responseBody interface{}) error {
	path_w_id := fmt.Sprintf("%s/%s", c.ResourceRootPath, requestPathExt)
	request, err := c.newAuthenticatedPostRequest(path_w_id, requestBody)
	if err != nil {
		return err
	}

	response, err := c.requestAndCheckStatus(fmt.Sprintf("create %s", c.ResourceDescription), request)
	if err != nil {
		return err
	}

	return unmarshalResponseBody(response, responseBody)
}

func (c *DefaultResourceClient) ModifyResourceWithPathExt(requestPathExt string, requestBody interface{}, responseBody interface{}) error {
	path_w_id := fmt.Sprintf("%s/%s", c.ResourceRootPath, requestPathExt)
	request, err := c.newAuthenticatedPutRequest(path_w_id, requestBody)
	if err != nil {
		return err
	}

	response, err := c.requestAndCheckStatus(fmt.Sprintf("create %s", c.ResourceDescription), request)
	if err != nil {
		return err
	}

	// 204 No Content success
	if response.StatusCode == 204 {
		return nil
	}

	return unmarshalResponseBody(response, responseBody)
}

func (c *DefaultResourceClient) FetchResource(requestPathId string, responseBody interface{}) error {
	path_w_id := fmt.Sprintf("%s/%s", c.ResourceRootPath, requestPathId)
	request, err := c.newAuthenticatedGetRequest(path_w_id)
	if err != nil {
		return err
	}

	response, err := c.requestAndCheckStatus(fmt.Sprintf("read %s", c.ResourceDescription), request)
	if err != nil {
		return err
	}

	return unmarshalResponseBody(response, responseBody)
}

func (c *DefaultResourceClient) GetResourceInfo(ctx context.Context, requestPath string, responseBody interface{}) error {
	
	request, err := c.newAuthenticatedGetRequest(requestPath)
	if err != nil {
		return err
	}
	response, err := c.requestAndCheckStatusWithContext(ctx, fmt.Sprintf("Get resource info %s", requestPath), request)
	if err != nil {
		return err
	}
	
	return unmarshalResponseBody(response, responseBody)
}

func (c *DefaultResourceClient) ModifyResource(ctx context.Context,  requestPath string, requestBody interface{}, responseBody interface{}) error {
	request, err := c.newAuthenticatedPutRequest(requestPath, requestBody)
	if err != nil {
		return err
	}

	response, err := c.requestAndCheckStatus(fmt.Sprintf("Upgrade resource info %s", requestPath), request)
	if err != nil {
		return err
	}

	// 204 No Content success
	if response.StatusCode == 204 {
		return nil
	}

	return unmarshalResponseBody(response, responseBody)
}

func (c *DefaultResourceClient) DeleteResource(ctx context.Context,  requestPath string, requestBody interface{}, responseBody interface{}) error {
	request, err := c.newAuthenticatedDeleteRequest(requestPath, requestBody)
	if err != nil {
		return err
	}

	response, err := c.requestAndCheckStatus(fmt.Sprintf("Upgrade resource info %s", requestPath), request)
	if err != nil {
		return err
	}

	// 204 No Content success
	if response.StatusCode == 204 {
		return nil
	}

	return unmarshalResponseBody(response, responseBody)
}
