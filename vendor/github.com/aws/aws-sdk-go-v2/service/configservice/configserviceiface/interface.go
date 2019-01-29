// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package configserviceiface provides an interface to enable mocking the AWS Config service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package configserviceiface

import (
	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

// ConfigServiceAPI provides an interface to enable mocking the
// configservice.ConfigService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Config.
//    func myFunc(svc configserviceiface.ConfigServiceAPI) bool {
//        // Make svc.BatchGetAggregateResourceConfig request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := configservice.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockConfigServiceClient struct {
//        configserviceiface.ConfigServiceAPI
//    }
//    func (m *mockConfigServiceClient) BatchGetAggregateResourceConfig(input *configservice.BatchGetAggregateResourceConfigInput) (*configservice.BatchGetAggregateResourceConfigOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockConfigServiceClient{}
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
type ConfigServiceAPI interface {
	BatchGetAggregateResourceConfigRequest(*configservice.BatchGetAggregateResourceConfigInput) configservice.BatchGetAggregateResourceConfigRequest

	BatchGetResourceConfigRequest(*configservice.BatchGetResourceConfigInput) configservice.BatchGetResourceConfigRequest

	DeleteAggregationAuthorizationRequest(*configservice.DeleteAggregationAuthorizationInput) configservice.DeleteAggregationAuthorizationRequest

	DeleteConfigRuleRequest(*configservice.DeleteConfigRuleInput) configservice.DeleteConfigRuleRequest

	DeleteConfigurationAggregatorRequest(*configservice.DeleteConfigurationAggregatorInput) configservice.DeleteConfigurationAggregatorRequest

	DeleteConfigurationRecorderRequest(*configservice.DeleteConfigurationRecorderInput) configservice.DeleteConfigurationRecorderRequest

	DeleteDeliveryChannelRequest(*configservice.DeleteDeliveryChannelInput) configservice.DeleteDeliveryChannelRequest

	DeleteEvaluationResultsRequest(*configservice.DeleteEvaluationResultsInput) configservice.DeleteEvaluationResultsRequest

	DeletePendingAggregationRequestRequest(*configservice.DeletePendingAggregationRequestInput) configservice.DeletePendingAggregationRequestRequest

	DeleteRetentionConfigurationRequest(*configservice.DeleteRetentionConfigurationInput) configservice.DeleteRetentionConfigurationRequest

	DeliverConfigSnapshotRequest(*configservice.DeliverConfigSnapshotInput) configservice.DeliverConfigSnapshotRequest

	DescribeAggregateComplianceByConfigRulesRequest(*configservice.DescribeAggregateComplianceByConfigRulesInput) configservice.DescribeAggregateComplianceByConfigRulesRequest

	DescribeAggregationAuthorizationsRequest(*configservice.DescribeAggregationAuthorizationsInput) configservice.DescribeAggregationAuthorizationsRequest

	DescribeComplianceByConfigRuleRequest(*configservice.DescribeComplianceByConfigRuleInput) configservice.DescribeComplianceByConfigRuleRequest

	DescribeComplianceByResourceRequest(*configservice.DescribeComplianceByResourceInput) configservice.DescribeComplianceByResourceRequest

	DescribeConfigRuleEvaluationStatusRequest(*configservice.DescribeConfigRuleEvaluationStatusInput) configservice.DescribeConfigRuleEvaluationStatusRequest

	DescribeConfigRulesRequest(*configservice.DescribeConfigRulesInput) configservice.DescribeConfigRulesRequest

	DescribeConfigurationAggregatorSourcesStatusRequest(*configservice.DescribeConfigurationAggregatorSourcesStatusInput) configservice.DescribeConfigurationAggregatorSourcesStatusRequest

	DescribeConfigurationAggregatorsRequest(*configservice.DescribeConfigurationAggregatorsInput) configservice.DescribeConfigurationAggregatorsRequest

	DescribeConfigurationRecorderStatusRequest(*configservice.DescribeConfigurationRecorderStatusInput) configservice.DescribeConfigurationRecorderStatusRequest

	DescribeConfigurationRecordersRequest(*configservice.DescribeConfigurationRecordersInput) configservice.DescribeConfigurationRecordersRequest

	DescribeDeliveryChannelStatusRequest(*configservice.DescribeDeliveryChannelStatusInput) configservice.DescribeDeliveryChannelStatusRequest

	DescribeDeliveryChannelsRequest(*configservice.DescribeDeliveryChannelsInput) configservice.DescribeDeliveryChannelsRequest

	DescribePendingAggregationRequestsRequest(*configservice.DescribePendingAggregationRequestsInput) configservice.DescribePendingAggregationRequestsRequest

	DescribeRetentionConfigurationsRequest(*configservice.DescribeRetentionConfigurationsInput) configservice.DescribeRetentionConfigurationsRequest

	GetAggregateComplianceDetailsByConfigRuleRequest(*configservice.GetAggregateComplianceDetailsByConfigRuleInput) configservice.GetAggregateComplianceDetailsByConfigRuleRequest

	GetAggregateConfigRuleComplianceSummaryRequest(*configservice.GetAggregateConfigRuleComplianceSummaryInput) configservice.GetAggregateConfigRuleComplianceSummaryRequest

	GetAggregateDiscoveredResourceCountsRequest(*configservice.GetAggregateDiscoveredResourceCountsInput) configservice.GetAggregateDiscoveredResourceCountsRequest

	GetAggregateResourceConfigRequest(*configservice.GetAggregateResourceConfigInput) configservice.GetAggregateResourceConfigRequest

	GetComplianceDetailsByConfigRuleRequest(*configservice.GetComplianceDetailsByConfigRuleInput) configservice.GetComplianceDetailsByConfigRuleRequest

	GetComplianceDetailsByResourceRequest(*configservice.GetComplianceDetailsByResourceInput) configservice.GetComplianceDetailsByResourceRequest

	GetComplianceSummaryByConfigRuleRequest(*configservice.GetComplianceSummaryByConfigRuleInput) configservice.GetComplianceSummaryByConfigRuleRequest

	GetComplianceSummaryByResourceTypeRequest(*configservice.GetComplianceSummaryByResourceTypeInput) configservice.GetComplianceSummaryByResourceTypeRequest

	GetDiscoveredResourceCountsRequest(*configservice.GetDiscoveredResourceCountsInput) configservice.GetDiscoveredResourceCountsRequest

	GetResourceConfigHistoryRequest(*configservice.GetResourceConfigHistoryInput) configservice.GetResourceConfigHistoryRequest

	ListAggregateDiscoveredResourcesRequest(*configservice.ListAggregateDiscoveredResourcesInput) configservice.ListAggregateDiscoveredResourcesRequest

	ListDiscoveredResourcesRequest(*configservice.ListDiscoveredResourcesInput) configservice.ListDiscoveredResourcesRequest

	PutAggregationAuthorizationRequest(*configservice.PutAggregationAuthorizationInput) configservice.PutAggregationAuthorizationRequest

	PutConfigRuleRequest(*configservice.PutConfigRuleInput) configservice.PutConfigRuleRequest

	PutConfigurationAggregatorRequest(*configservice.PutConfigurationAggregatorInput) configservice.PutConfigurationAggregatorRequest

	PutConfigurationRecorderRequest(*configservice.PutConfigurationRecorderInput) configservice.PutConfigurationRecorderRequest

	PutDeliveryChannelRequest(*configservice.PutDeliveryChannelInput) configservice.PutDeliveryChannelRequest

	PutEvaluationsRequest(*configservice.PutEvaluationsInput) configservice.PutEvaluationsRequest

	PutRetentionConfigurationRequest(*configservice.PutRetentionConfigurationInput) configservice.PutRetentionConfigurationRequest

	StartConfigRulesEvaluationRequest(*configservice.StartConfigRulesEvaluationInput) configservice.StartConfigRulesEvaluationRequest

	StartConfigurationRecorderRequest(*configservice.StartConfigurationRecorderInput) configservice.StartConfigurationRecorderRequest

	StopConfigurationRecorderRequest(*configservice.StopConfigurationRecorderInput) configservice.StopConfigurationRecorderRequest
}

var _ ConfigServiceAPI = (*configservice.ConfigService)(nil)
