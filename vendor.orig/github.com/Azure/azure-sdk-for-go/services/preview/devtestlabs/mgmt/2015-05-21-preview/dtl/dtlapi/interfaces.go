package dtlapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/devtestlabs/mgmt/2015-05-21-preview/dtl"
	"github.com/Azure/go-autorest/autorest"
)

// LabClientAPI contains the set of methods on the LabClient type.
type LabClientAPI interface {
	CreateEnvironment(ctx context.Context, resourceGroupName string, name string, labVirtualMachine dtl.LabVirtualMachine) (result dtl.LabCreateEnvironmentFuture, err error)
	CreateOrUpdateResource(ctx context.Context, resourceGroupName string, name string, lab dtl.Lab) (result dtl.LabCreateOrUpdateResourceFuture, err error)
	DeleteResource(ctx context.Context, resourceGroupName string, name string) (result dtl.LabDeleteResourceFuture, err error)
	GenerateUploadURI(ctx context.Context, resourceGroupName string, name string, generateUploadURIParameter dtl.GenerateUploadURIParameter) (result dtl.GenerateUploadURIResponse, err error)
	GetResource(ctx context.Context, resourceGroupName string, name string) (result dtl.Lab, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationLabPage, err error)
	ListBySubscription(ctx context.Context, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationLabPage, err error)
	ListVhds(ctx context.Context, resourceGroupName string, name string) (result dtl.ResponseWithContinuationLabVhdPage, err error)
	PatchResource(ctx context.Context, resourceGroupName string, name string, lab dtl.Lab) (result dtl.Lab, err error)
}

var _ LabClientAPI = (*dtl.LabClient)(nil)

// ArtifactSourceClientAPI contains the set of methods on the ArtifactSourceClient type.
type ArtifactSourceClientAPI interface {
	CreateOrUpdateResource(ctx context.Context, resourceGroupName string, labName string, name string, artifactSource dtl.ArtifactSource) (result dtl.ArtifactSource, err error)
	DeleteResource(ctx context.Context, resourceGroupName string, labName string, name string) (result autorest.Response, err error)
	GetResource(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.ArtifactSource, err error)
	List(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationArtifactSourcePage, err error)
	PatchResource(ctx context.Context, resourceGroupName string, labName string, name string, artifactSource dtl.ArtifactSource) (result dtl.ArtifactSource, err error)
}

var _ ArtifactSourceClientAPI = (*dtl.ArtifactSourceClient)(nil)

// ArtifactClientAPI contains the set of methods on the ArtifactClient type.
type ArtifactClientAPI interface {
	GenerateArmTemplate(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, generateArmTemplateRequest dtl.GenerateArmTemplateRequest) (result dtl.ArmTemplateInfo, err error)
	GetResource(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string) (result dtl.Artifact, err error)
	List(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationArtifactPage, err error)
}

var _ ArtifactClientAPI = (*dtl.ArtifactClient)(nil)

// CostInsightClientAPI contains the set of methods on the CostInsightClient type.
type CostInsightClientAPI interface {
	GetResource(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.CostInsight, err error)
	List(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationCostInsightPage, err error)
	RefreshData(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.CostInsightRefreshDataFuture, err error)
}

var _ CostInsightClientAPI = (*dtl.CostInsightClient)(nil)

// CostClientAPI contains the set of methods on the CostClient type.
type CostClientAPI interface {
	GetResource(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.Cost, err error)
	List(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationCostPage, err error)
	RefreshData(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.CostRefreshDataFuture, err error)
}

var _ CostClientAPI = (*dtl.CostClient)(nil)

// CustomImageClientAPI contains the set of methods on the CustomImageClient type.
type CustomImageClientAPI interface {
	CreateOrUpdateResource(ctx context.Context, resourceGroupName string, labName string, name string, customImage dtl.CustomImage) (result dtl.CustomImageCreateOrUpdateResourceFuture, err error)
	DeleteResource(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.CustomImageDeleteResourceFuture, err error)
	GetResource(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.CustomImage, err error)
	List(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationCustomImagePage, err error)
}

var _ CustomImageClientAPI = (*dtl.CustomImageClient)(nil)

// FormulaClientAPI contains the set of methods on the FormulaClient type.
type FormulaClientAPI interface {
	CreateOrUpdateResource(ctx context.Context, resourceGroupName string, labName string, name string, formula dtl.Formula) (result dtl.FormulaCreateOrUpdateResourceFuture, err error)
	DeleteResource(ctx context.Context, resourceGroupName string, labName string, name string) (result autorest.Response, err error)
	GetResource(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.Formula, err error)
	List(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationFormulaPage, err error)
}

var _ FormulaClientAPI = (*dtl.FormulaClient)(nil)

// GalleryImageClientAPI contains the set of methods on the GalleryImageClient type.
type GalleryImageClientAPI interface {
	List(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationGalleryImagePage, err error)
}

var _ GalleryImageClientAPI = (*dtl.GalleryImageClient)(nil)

// PolicySetClientAPI contains the set of methods on the PolicySetClient type.
type PolicySetClientAPI interface {
	EvaluatePolicies(ctx context.Context, resourceGroupName string, labName string, name string, evaluatePoliciesRequest dtl.EvaluatePoliciesRequest) (result dtl.EvaluatePoliciesResponse, err error)
}

var _ PolicySetClientAPI = (*dtl.PolicySetClient)(nil)

// PolicyClientAPI contains the set of methods on the PolicyClient type.
type PolicyClientAPI interface {
	CreateOrUpdateResource(ctx context.Context, resourceGroupName string, labName string, policySetName string, name string, policy dtl.Policy) (result dtl.Policy, err error)
	DeleteResource(ctx context.Context, resourceGroupName string, labName string, policySetName string, name string) (result autorest.Response, err error)
	GetResource(ctx context.Context, resourceGroupName string, labName string, policySetName string, name string) (result dtl.Policy, err error)
	List(ctx context.Context, resourceGroupName string, labName string, policySetName string, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationPolicyPage, err error)
	PatchResource(ctx context.Context, resourceGroupName string, labName string, policySetName string, name string, policy dtl.Policy) (result dtl.Policy, err error)
}

var _ PolicyClientAPI = (*dtl.PolicyClient)(nil)

// ScheduleClientAPI contains the set of methods on the ScheduleClient type.
type ScheduleClientAPI interface {
	CreateOrUpdateResource(ctx context.Context, resourceGroupName string, labName string, name string, schedule dtl.Schedule) (result dtl.ScheduleCreateOrUpdateResourceFuture, err error)
	DeleteResource(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.ScheduleDeleteResourceFuture, err error)
	Execute(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.ScheduleExecuteFuture, err error)
	GetResource(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.Schedule, err error)
	List(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationSchedulePage, err error)
	PatchResource(ctx context.Context, resourceGroupName string, labName string, name string, schedule dtl.Schedule) (result dtl.Schedule, err error)
}

var _ ScheduleClientAPI = (*dtl.ScheduleClient)(nil)

// VirtualMachineClientAPI contains the set of methods on the VirtualMachineClient type.
type VirtualMachineClientAPI interface {
	ApplyArtifacts(ctx context.Context, resourceGroupName string, labName string, name string, applyArtifactsRequest dtl.ApplyArtifactsRequest) (result dtl.VirtualMachineApplyArtifactsFuture, err error)
	CreateOrUpdateResource(ctx context.Context, resourceGroupName string, labName string, name string, labVirtualMachine dtl.LabVirtualMachine) (result dtl.VirtualMachineCreateOrUpdateResourceFuture, err error)
	DeleteResource(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualMachineDeleteResourceFuture, err error)
	GetResource(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.LabVirtualMachine, err error)
	List(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationLabVirtualMachinePage, err error)
	PatchResource(ctx context.Context, resourceGroupName string, labName string, name string, labVirtualMachine dtl.LabVirtualMachine) (result dtl.LabVirtualMachine, err error)
	Start(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualMachineStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualMachineStopFuture, err error)
}

var _ VirtualMachineClientAPI = (*dtl.VirtualMachineClient)(nil)

// VirtualNetworkClientAPI contains the set of methods on the VirtualNetworkClient type.
type VirtualNetworkClientAPI interface {
	CreateOrUpdateResource(ctx context.Context, resourceGroupName string, labName string, name string, virtualNetwork dtl.VirtualNetwork) (result dtl.VirtualNetworkCreateOrUpdateResourceFuture, err error)
	DeleteResource(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualNetworkDeleteResourceFuture, err error)
	GetResource(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualNetwork, err error)
	List(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result dtl.ResponseWithContinuationVirtualNetworkPage, err error)
	PatchResource(ctx context.Context, resourceGroupName string, labName string, name string, virtualNetwork dtl.VirtualNetwork) (result dtl.VirtualNetwork, err error)
}

var _ VirtualNetworkClientAPI = (*dtl.VirtualNetworkClient)(nil)
