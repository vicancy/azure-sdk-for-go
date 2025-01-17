//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewGalleryImagesClient() *GalleryImagesClient {
	subClient, _ := NewGalleryImagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGuestAgentClient() *GuestAgentClient {
	subClient, _ := NewGuestAgentClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGuestAgentsClient() *GuestAgentsClient {
	subClient, _ := NewGuestAgentsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHybridIdentityMetadataClient() *HybridIdentityMetadataClient {
	subClient, _ := NewHybridIdentityMetadataClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLogicalNetworksClient() *LogicalNetworksClient {
	subClient, _ := NewLogicalNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMarketplaceGalleryImagesClient() *MarketplaceGalleryImagesClient {
	subClient, _ := NewMarketplaceGalleryImagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkInterfacesClient() *NetworkInterfacesClient {
	subClient, _ := NewNetworkInterfacesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewStorageContainersClient() *StorageContainersClient {
	subClient, _ := NewStorageContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualHardDisksClient() *VirtualHardDisksClient {
	subClient, _ := NewVirtualHardDisksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualMachineInstancesClient() *VirtualMachineInstancesClient {
	subClient, _ := NewVirtualMachineInstancesClient(c.credential, c.options)
	return subClient
}
