package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyRouterInterfaceAttribute invokes the vpc.ModifyRouterInterfaceAttribute API synchronously
// api document: https://help.aliyun.com/api/vpc/modifyrouterinterfaceattribute.html
func (client *Client) ModifyRouterInterfaceAttribute(request *ModifyRouterInterfaceAttributeRequest) (response *ModifyRouterInterfaceAttributeResponse, err error) {
	response = CreateModifyRouterInterfaceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyRouterInterfaceAttributeWithChan invokes the vpc.ModifyRouterInterfaceAttribute API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyrouterinterfaceattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyRouterInterfaceAttributeWithChan(request *ModifyRouterInterfaceAttributeRequest) (<-chan *ModifyRouterInterfaceAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyRouterInterfaceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyRouterInterfaceAttribute(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyRouterInterfaceAttributeWithCallback invokes the vpc.ModifyRouterInterfaceAttribute API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyrouterinterfaceattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyRouterInterfaceAttributeWithCallback(request *ModifyRouterInterfaceAttributeRequest, callback func(response *ModifyRouterInterfaceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyRouterInterfaceAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyRouterInterfaceAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyRouterInterfaceAttributeRequest is the request struct for api ModifyRouterInterfaceAttribute
type ModifyRouterInterfaceAttributeRequest struct {
	*requests.RpcRequest
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RouterInterfaceId        string           `position:"Query" name:"RouterInterfaceId"`
	Name                     string           `position:"Query" name:"Name"`
	Description              string           `position:"Query" name:"Description"`
	OppositeInterfaceId      string           `position:"Query" name:"OppositeInterfaceId"`
	OppositeRouterId         string           `position:"Query" name:"OppositeRouterId"`
	OppositeRouterType       string           `position:"Query" name:"OppositeRouterType"`
	OppositeInterfaceOwnerId requests.Integer `position:"Query" name:"OppositeInterfaceOwnerId"`
	HealthCheckSourceIp      string           `position:"Query" name:"HealthCheckSourceIp"`
	HealthCheckTargetIp      string           `position:"Query" name:"HealthCheckTargetIp"`
	DeleteHealthCheckIp      requests.Boolean `position:"Query" name:"DeleteHealthCheckIp"`
}

// ModifyRouterInterfaceAttributeResponse is the response struct for api ModifyRouterInterfaceAttribute
type ModifyRouterInterfaceAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyRouterInterfaceAttributeRequest creates a request to invoke ModifyRouterInterfaceAttribute API
func CreateModifyRouterInterfaceAttributeRequest() (request *ModifyRouterInterfaceAttributeRequest) {
	request = &ModifyRouterInterfaceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyRouterInterfaceAttribute", "vpc", "openAPI")
	return
}

// CreateModifyRouterInterfaceAttributeResponse creates a response to parse from ModifyRouterInterfaceAttribute response
func CreateModifyRouterInterfaceAttributeResponse() (response *ModifyRouterInterfaceAttributeResponse) {
	response = &ModifyRouterInterfaceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}