// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kinesisanalyticsv2iface provides an interface to enable mocking the Amazon Kinesis Analytics service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kinesisanalyticsv2iface

import (
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
)

// KinesisAnalyticsV2API provides an interface to enable mocking the
// kinesisanalyticsv2.KinesisAnalyticsV2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Kinesis Analytics.
//    func myFunc(svc kinesisanalyticsv2iface.KinesisAnalyticsV2API) bool {
//        // Make svc.AddApplicationCloudWatchLoggingOption request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := kinesisanalyticsv2.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKinesisAnalyticsV2Client struct {
//        kinesisanalyticsv2iface.KinesisAnalyticsV2API
//    }
//    func (m *mockKinesisAnalyticsV2Client) AddApplicationCloudWatchLoggingOption(input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKinesisAnalyticsV2Client{}
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
type KinesisAnalyticsV2API interface {
	AddApplicationCloudWatchLoggingOptionRequest(*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionRequest

	AddApplicationInputRequest(*kinesisanalyticsv2.AddApplicationInputInput) kinesisanalyticsv2.AddApplicationInputRequest

	AddApplicationInputProcessingConfigurationRequest(*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) kinesisanalyticsv2.AddApplicationInputProcessingConfigurationRequest

	AddApplicationOutputRequest(*kinesisanalyticsv2.AddApplicationOutputInput) kinesisanalyticsv2.AddApplicationOutputRequest

	AddApplicationReferenceDataSourceRequest(*kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) kinesisanalyticsv2.AddApplicationReferenceDataSourceRequest

	CreateApplicationRequest(*kinesisanalyticsv2.CreateApplicationInput) kinesisanalyticsv2.CreateApplicationRequest

	CreateApplicationSnapshotRequest(*kinesisanalyticsv2.CreateApplicationSnapshotInput) kinesisanalyticsv2.CreateApplicationSnapshotRequest

	DeleteApplicationRequest(*kinesisanalyticsv2.DeleteApplicationInput) kinesisanalyticsv2.DeleteApplicationRequest

	DeleteApplicationCloudWatchLoggingOptionRequest(*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionRequest

	DeleteApplicationInputProcessingConfigurationRequest(*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationRequest

	DeleteApplicationOutputRequest(*kinesisanalyticsv2.DeleteApplicationOutputInput) kinesisanalyticsv2.DeleteApplicationOutputRequest

	DeleteApplicationReferenceDataSourceRequest(*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) kinesisanalyticsv2.DeleteApplicationReferenceDataSourceRequest

	DeleteApplicationSnapshotRequest(*kinesisanalyticsv2.DeleteApplicationSnapshotInput) kinesisanalyticsv2.DeleteApplicationSnapshotRequest

	DescribeApplicationRequest(*kinesisanalyticsv2.DescribeApplicationInput) kinesisanalyticsv2.DescribeApplicationRequest

	DescribeApplicationSnapshotRequest(*kinesisanalyticsv2.DescribeApplicationSnapshotInput) kinesisanalyticsv2.DescribeApplicationSnapshotRequest

	DiscoverInputSchemaRequest(*kinesisanalyticsv2.DiscoverInputSchemaInput) kinesisanalyticsv2.DiscoverInputSchemaRequest

	ListApplicationSnapshotsRequest(*kinesisanalyticsv2.ListApplicationSnapshotsInput) kinesisanalyticsv2.ListApplicationSnapshotsRequest

	ListApplicationsRequest(*kinesisanalyticsv2.ListApplicationsInput) kinesisanalyticsv2.ListApplicationsRequest

	StartApplicationRequest(*kinesisanalyticsv2.StartApplicationInput) kinesisanalyticsv2.StartApplicationRequest

	StopApplicationRequest(*kinesisanalyticsv2.StopApplicationInput) kinesisanalyticsv2.StopApplicationRequest

	UpdateApplicationRequest(*kinesisanalyticsv2.UpdateApplicationInput) kinesisanalyticsv2.UpdateApplicationRequest
}

var _ KinesisAnalyticsV2API = (*kinesisanalyticsv2.KinesisAnalyticsV2)(nil)
