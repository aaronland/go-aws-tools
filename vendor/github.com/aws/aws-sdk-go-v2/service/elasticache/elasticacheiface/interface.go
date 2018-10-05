// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elasticacheiface provides an interface to enable mocking the Amazon ElastiCache service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package elasticacheiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

// ElastiCacheAPI provides an interface to enable mocking the
// elasticache.ElastiCache service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon ElastiCache.
//    func myFunc(svc elasticacheiface.ElastiCacheAPI) bool {
//        // Make svc.AddTagsToResource request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := elasticache.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockElastiCacheClient struct {
//        elasticacheiface.ElastiCacheAPI
//    }
//    func (m *mockElastiCacheClient) AddTagsToResource(input *elasticache.AddTagsToResourceInput) (*elasticache.RemoveTagsFromResourceOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockElastiCacheClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ElastiCacheAPI interface {
	AddTagsToResourceRequest(*elasticache.AddTagsToResourceInput) elasticache.AddTagsToResourceRequest

	AuthorizeCacheSecurityGroupIngressRequest(*elasticache.AuthorizeCacheSecurityGroupIngressInput) elasticache.AuthorizeCacheSecurityGroupIngressRequest

	CopySnapshotRequest(*elasticache.CopySnapshotInput) elasticache.CopySnapshotRequest

	CreateCacheClusterRequest(*elasticache.CreateCacheClusterInput) elasticache.CreateCacheClusterRequest

	CreateCacheParameterGroupRequest(*elasticache.CreateCacheParameterGroupInput) elasticache.CreateCacheParameterGroupRequest

	CreateCacheSecurityGroupRequest(*elasticache.CreateCacheSecurityGroupInput) elasticache.CreateCacheSecurityGroupRequest

	CreateCacheSubnetGroupRequest(*elasticache.CreateCacheSubnetGroupInput) elasticache.CreateCacheSubnetGroupRequest

	CreateReplicationGroupRequest(*elasticache.CreateReplicationGroupInput) elasticache.CreateReplicationGroupRequest

	CreateSnapshotRequest(*elasticache.CreateSnapshotInput) elasticache.CreateSnapshotRequest

	DecreaseReplicaCountRequest(*elasticache.DecreaseReplicaCountInput) elasticache.DecreaseReplicaCountRequest

	DeleteCacheClusterRequest(*elasticache.DeleteCacheClusterInput) elasticache.DeleteCacheClusterRequest

	DeleteCacheParameterGroupRequest(*elasticache.DeleteCacheParameterGroupInput) elasticache.DeleteCacheParameterGroupRequest

	DeleteCacheSecurityGroupRequest(*elasticache.DeleteCacheSecurityGroupInput) elasticache.DeleteCacheSecurityGroupRequest

	DeleteCacheSubnetGroupRequest(*elasticache.DeleteCacheSubnetGroupInput) elasticache.DeleteCacheSubnetGroupRequest

	DeleteReplicationGroupRequest(*elasticache.DeleteReplicationGroupInput) elasticache.DeleteReplicationGroupRequest

	DeleteSnapshotRequest(*elasticache.DeleteSnapshotInput) elasticache.DeleteSnapshotRequest

	DescribeCacheClustersRequest(*elasticache.DescribeCacheClustersInput) elasticache.DescribeCacheClustersRequest

	DescribeCacheEngineVersionsRequest(*elasticache.DescribeCacheEngineVersionsInput) elasticache.DescribeCacheEngineVersionsRequest

	DescribeCacheParameterGroupsRequest(*elasticache.DescribeCacheParameterGroupsInput) elasticache.DescribeCacheParameterGroupsRequest

	DescribeCacheParametersRequest(*elasticache.DescribeCacheParametersInput) elasticache.DescribeCacheParametersRequest

	DescribeCacheSecurityGroupsRequest(*elasticache.DescribeCacheSecurityGroupsInput) elasticache.DescribeCacheSecurityGroupsRequest

	DescribeCacheSubnetGroupsRequest(*elasticache.DescribeCacheSubnetGroupsInput) elasticache.DescribeCacheSubnetGroupsRequest

	DescribeEngineDefaultParametersRequest(*elasticache.DescribeEngineDefaultParametersInput) elasticache.DescribeEngineDefaultParametersRequest

	DescribeEventsRequest(*elasticache.DescribeEventsInput) elasticache.DescribeEventsRequest

	DescribeReplicationGroupsRequest(*elasticache.DescribeReplicationGroupsInput) elasticache.DescribeReplicationGroupsRequest

	DescribeReservedCacheNodesRequest(*elasticache.DescribeReservedCacheNodesInput) elasticache.DescribeReservedCacheNodesRequest

	DescribeReservedCacheNodesOfferingsRequest(*elasticache.DescribeReservedCacheNodesOfferingsInput) elasticache.DescribeReservedCacheNodesOfferingsRequest

	DescribeSnapshotsRequest(*elasticache.DescribeSnapshotsInput) elasticache.DescribeSnapshotsRequest

	IncreaseReplicaCountRequest(*elasticache.IncreaseReplicaCountInput) elasticache.IncreaseReplicaCountRequest

	ListAllowedNodeTypeModificationsRequest(*elasticache.ListAllowedNodeTypeModificationsInput) elasticache.ListAllowedNodeTypeModificationsRequest

	ListTagsForResourceRequest(*elasticache.ListTagsForResourceInput) elasticache.ListTagsForResourceRequest

	ModifyCacheClusterRequest(*elasticache.ModifyCacheClusterInput) elasticache.ModifyCacheClusterRequest

	ModifyCacheParameterGroupRequest(*elasticache.ModifyCacheParameterGroupInput) elasticache.ModifyCacheParameterGroupRequest

	ModifyCacheSubnetGroupRequest(*elasticache.ModifyCacheSubnetGroupInput) elasticache.ModifyCacheSubnetGroupRequest

	ModifyReplicationGroupRequest(*elasticache.ModifyReplicationGroupInput) elasticache.ModifyReplicationGroupRequest

	ModifyReplicationGroupShardConfigurationRequest(*elasticache.ModifyReplicationGroupShardConfigurationInput) elasticache.ModifyReplicationGroupShardConfigurationRequest

	PurchaseReservedCacheNodesOfferingRequest(*elasticache.PurchaseReservedCacheNodesOfferingInput) elasticache.PurchaseReservedCacheNodesOfferingRequest

	RebootCacheClusterRequest(*elasticache.RebootCacheClusterInput) elasticache.RebootCacheClusterRequest

	RemoveTagsFromResourceRequest(*elasticache.RemoveTagsFromResourceInput) elasticache.RemoveTagsFromResourceRequest

	ResetCacheParameterGroupRequest(*elasticache.ResetCacheParameterGroupInput) elasticache.ResetCacheParameterGroupRequest

	RevokeCacheSecurityGroupIngressRequest(*elasticache.RevokeCacheSecurityGroupIngressInput) elasticache.RevokeCacheSecurityGroupIngressRequest

	TestFailoverRequest(*elasticache.TestFailoverInput) elasticache.TestFailoverRequest

	WaitUntilCacheClusterAvailable(*elasticache.DescribeCacheClustersInput) error
	WaitUntilCacheClusterAvailableWithContext(aws.Context, *elasticache.DescribeCacheClustersInput, ...aws.WaiterOption) error

	WaitUntilCacheClusterDeleted(*elasticache.DescribeCacheClustersInput) error
	WaitUntilCacheClusterDeletedWithContext(aws.Context, *elasticache.DescribeCacheClustersInput, ...aws.WaiterOption) error

	WaitUntilReplicationGroupAvailable(*elasticache.DescribeReplicationGroupsInput) error
	WaitUntilReplicationGroupAvailableWithContext(aws.Context, *elasticache.DescribeReplicationGroupsInput, ...aws.WaiterOption) error

	WaitUntilReplicationGroupDeleted(*elasticache.DescribeReplicationGroupsInput) error
	WaitUntilReplicationGroupDeletedWithContext(aws.Context, *elasticache.DescribeReplicationGroupsInput, ...aws.WaiterOption) error
}

var _ ElastiCacheAPI = (*elasticache.ElastiCache)(nil)
