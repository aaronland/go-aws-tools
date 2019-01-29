// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iamiface provides an interface to enable mocking the AWS Identity and Access Management service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iamiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

// IAMAPI provides an interface to enable mocking the
// iam.IAM service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Identity and Access Management.
//    func myFunc(svc iamiface.IAMAPI) bool {
//        // Make svc.AddClientIDToOpenIDConnectProvider request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := iam.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIAMClient struct {
//        iamiface.IAMAPI
//    }
//    func (m *mockIAMClient) AddClientIDToOpenIDConnectProvider(input *iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIAMClient{}
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
type IAMAPI interface {
	AddClientIDToOpenIDConnectProviderRequest(*iam.AddClientIDToOpenIDConnectProviderInput) iam.AddClientIDToOpenIDConnectProviderRequest

	AddRoleToInstanceProfileRequest(*iam.AddRoleToInstanceProfileInput) iam.AddRoleToInstanceProfileRequest

	AddUserToGroupRequest(*iam.AddUserToGroupInput) iam.AddUserToGroupRequest

	AttachGroupPolicyRequest(*iam.AttachGroupPolicyInput) iam.AttachGroupPolicyRequest

	AttachRolePolicyRequest(*iam.AttachRolePolicyInput) iam.AttachRolePolicyRequest

	AttachUserPolicyRequest(*iam.AttachUserPolicyInput) iam.AttachUserPolicyRequest

	ChangePasswordRequest(*iam.ChangePasswordInput) iam.ChangePasswordRequest

	CreateAccessKeyRequest(*iam.CreateAccessKeyInput) iam.CreateAccessKeyRequest

	CreateAccountAliasRequest(*iam.CreateAccountAliasInput) iam.CreateAccountAliasRequest

	CreateGroupRequest(*iam.CreateGroupInput) iam.CreateGroupRequest

	CreateInstanceProfileRequest(*iam.CreateInstanceProfileInput) iam.CreateInstanceProfileRequest

	CreateLoginProfileRequest(*iam.CreateLoginProfileInput) iam.CreateLoginProfileRequest

	CreateOpenIDConnectProviderRequest(*iam.CreateOpenIDConnectProviderInput) iam.CreateOpenIDConnectProviderRequest

	CreatePolicyRequest(*iam.CreatePolicyInput) iam.CreatePolicyRequest

	CreatePolicyVersionRequest(*iam.CreatePolicyVersionInput) iam.CreatePolicyVersionRequest

	CreateRoleRequest(*iam.CreateRoleInput) iam.CreateRoleRequest

	CreateSAMLProviderRequest(*iam.CreateSAMLProviderInput) iam.CreateSAMLProviderRequest

	CreateServiceLinkedRoleRequest(*iam.CreateServiceLinkedRoleInput) iam.CreateServiceLinkedRoleRequest

	CreateServiceSpecificCredentialRequest(*iam.CreateServiceSpecificCredentialInput) iam.CreateServiceSpecificCredentialRequest

	CreateUserRequest(*iam.CreateUserInput) iam.CreateUserRequest

	CreateVirtualMFADeviceRequest(*iam.CreateVirtualMFADeviceInput) iam.CreateVirtualMFADeviceRequest

	DeactivateMFADeviceRequest(*iam.DeactivateMFADeviceInput) iam.DeactivateMFADeviceRequest

	DeleteAccessKeyRequest(*iam.DeleteAccessKeyInput) iam.DeleteAccessKeyRequest

	DeleteAccountAliasRequest(*iam.DeleteAccountAliasInput) iam.DeleteAccountAliasRequest

	DeleteAccountPasswordPolicyRequest(*iam.DeleteAccountPasswordPolicyInput) iam.DeleteAccountPasswordPolicyRequest

	DeleteGroupRequest(*iam.DeleteGroupInput) iam.DeleteGroupRequest

	DeleteGroupPolicyRequest(*iam.DeleteGroupPolicyInput) iam.DeleteGroupPolicyRequest

	DeleteInstanceProfileRequest(*iam.DeleteInstanceProfileInput) iam.DeleteInstanceProfileRequest

	DeleteLoginProfileRequest(*iam.DeleteLoginProfileInput) iam.DeleteLoginProfileRequest

	DeleteOpenIDConnectProviderRequest(*iam.DeleteOpenIDConnectProviderInput) iam.DeleteOpenIDConnectProviderRequest

	DeletePolicyRequest(*iam.DeletePolicyInput) iam.DeletePolicyRequest

	DeletePolicyVersionRequest(*iam.DeletePolicyVersionInput) iam.DeletePolicyVersionRequest

	DeleteRoleRequest(*iam.DeleteRoleInput) iam.DeleteRoleRequest

	DeleteRolePermissionsBoundaryRequest(*iam.DeleteRolePermissionsBoundaryInput) iam.DeleteRolePermissionsBoundaryRequest

	DeleteRolePolicyRequest(*iam.DeleteRolePolicyInput) iam.DeleteRolePolicyRequest

	DeleteSAMLProviderRequest(*iam.DeleteSAMLProviderInput) iam.DeleteSAMLProviderRequest

	DeleteSSHPublicKeyRequest(*iam.DeleteSSHPublicKeyInput) iam.DeleteSSHPublicKeyRequest

	DeleteServerCertificateRequest(*iam.DeleteServerCertificateInput) iam.DeleteServerCertificateRequest

	DeleteServiceLinkedRoleRequest(*iam.DeleteServiceLinkedRoleInput) iam.DeleteServiceLinkedRoleRequest

	DeleteServiceSpecificCredentialRequest(*iam.DeleteServiceSpecificCredentialInput) iam.DeleteServiceSpecificCredentialRequest

	DeleteSigningCertificateRequest(*iam.DeleteSigningCertificateInput) iam.DeleteSigningCertificateRequest

	DeleteUserRequest(*iam.DeleteUserInput) iam.DeleteUserRequest

	DeleteUserPermissionsBoundaryRequest(*iam.DeleteUserPermissionsBoundaryInput) iam.DeleteUserPermissionsBoundaryRequest

	DeleteUserPolicyRequest(*iam.DeleteUserPolicyInput) iam.DeleteUserPolicyRequest

	DeleteVirtualMFADeviceRequest(*iam.DeleteVirtualMFADeviceInput) iam.DeleteVirtualMFADeviceRequest

	DetachGroupPolicyRequest(*iam.DetachGroupPolicyInput) iam.DetachGroupPolicyRequest

	DetachRolePolicyRequest(*iam.DetachRolePolicyInput) iam.DetachRolePolicyRequest

	DetachUserPolicyRequest(*iam.DetachUserPolicyInput) iam.DetachUserPolicyRequest

	EnableMFADeviceRequest(*iam.EnableMFADeviceInput) iam.EnableMFADeviceRequest

	GenerateCredentialReportRequest(*iam.GenerateCredentialReportInput) iam.GenerateCredentialReportRequest

	GenerateServiceLastAccessedDetailsRequest(*iam.GenerateServiceLastAccessedDetailsInput) iam.GenerateServiceLastAccessedDetailsRequest

	GetAccessKeyLastUsedRequest(*iam.GetAccessKeyLastUsedInput) iam.GetAccessKeyLastUsedRequest

	GetAccountAuthorizationDetailsRequest(*iam.GetAccountAuthorizationDetailsInput) iam.GetAccountAuthorizationDetailsRequest

	GetAccountPasswordPolicyRequest(*iam.GetAccountPasswordPolicyInput) iam.GetAccountPasswordPolicyRequest

	GetAccountSummaryRequest(*iam.GetAccountSummaryInput) iam.GetAccountSummaryRequest

	GetContextKeysForCustomPolicyRequest(*iam.GetContextKeysForCustomPolicyInput) iam.GetContextKeysForCustomPolicyRequest

	GetContextKeysForPrincipalPolicyRequest(*iam.GetContextKeysForPrincipalPolicyInput) iam.GetContextKeysForPrincipalPolicyRequest

	GetCredentialReportRequest(*iam.GetCredentialReportInput) iam.GetCredentialReportRequest

	GetGroupRequest(*iam.GetGroupInput) iam.GetGroupRequest

	GetGroupPolicyRequest(*iam.GetGroupPolicyInput) iam.GetGroupPolicyRequest

	GetInstanceProfileRequest(*iam.GetInstanceProfileInput) iam.GetInstanceProfileRequest

	GetLoginProfileRequest(*iam.GetLoginProfileInput) iam.GetLoginProfileRequest

	GetOpenIDConnectProviderRequest(*iam.GetOpenIDConnectProviderInput) iam.GetOpenIDConnectProviderRequest

	GetPolicyRequest(*iam.GetPolicyInput) iam.GetPolicyRequest

	GetPolicyVersionRequest(*iam.GetPolicyVersionInput) iam.GetPolicyVersionRequest

	GetRoleRequest(*iam.GetRoleInput) iam.GetRoleRequest

	GetRolePolicyRequest(*iam.GetRolePolicyInput) iam.GetRolePolicyRequest

	GetSAMLProviderRequest(*iam.GetSAMLProviderInput) iam.GetSAMLProviderRequest

	GetSSHPublicKeyRequest(*iam.GetSSHPublicKeyInput) iam.GetSSHPublicKeyRequest

	GetServerCertificateRequest(*iam.GetServerCertificateInput) iam.GetServerCertificateRequest

	GetServiceLastAccessedDetailsRequest(*iam.GetServiceLastAccessedDetailsInput) iam.GetServiceLastAccessedDetailsRequest

	GetServiceLastAccessedDetailsWithEntitiesRequest(*iam.GetServiceLastAccessedDetailsWithEntitiesInput) iam.GetServiceLastAccessedDetailsWithEntitiesRequest

	GetServiceLinkedRoleDeletionStatusRequest(*iam.GetServiceLinkedRoleDeletionStatusInput) iam.GetServiceLinkedRoleDeletionStatusRequest

	GetUserRequest(*iam.GetUserInput) iam.GetUserRequest

	GetUserPolicyRequest(*iam.GetUserPolicyInput) iam.GetUserPolicyRequest

	ListAccessKeysRequest(*iam.ListAccessKeysInput) iam.ListAccessKeysRequest

	ListAccountAliasesRequest(*iam.ListAccountAliasesInput) iam.ListAccountAliasesRequest

	ListAttachedGroupPoliciesRequest(*iam.ListAttachedGroupPoliciesInput) iam.ListAttachedGroupPoliciesRequest

	ListAttachedRolePoliciesRequest(*iam.ListAttachedRolePoliciesInput) iam.ListAttachedRolePoliciesRequest

	ListAttachedUserPoliciesRequest(*iam.ListAttachedUserPoliciesInput) iam.ListAttachedUserPoliciesRequest

	ListEntitiesForPolicyRequest(*iam.ListEntitiesForPolicyInput) iam.ListEntitiesForPolicyRequest

	ListGroupPoliciesRequest(*iam.ListGroupPoliciesInput) iam.ListGroupPoliciesRequest

	ListGroupsRequest(*iam.ListGroupsInput) iam.ListGroupsRequest

	ListGroupsForUserRequest(*iam.ListGroupsForUserInput) iam.ListGroupsForUserRequest

	ListInstanceProfilesRequest(*iam.ListInstanceProfilesInput) iam.ListInstanceProfilesRequest

	ListInstanceProfilesForRoleRequest(*iam.ListInstanceProfilesForRoleInput) iam.ListInstanceProfilesForRoleRequest

	ListMFADevicesRequest(*iam.ListMFADevicesInput) iam.ListMFADevicesRequest

	ListOpenIDConnectProvidersRequest(*iam.ListOpenIDConnectProvidersInput) iam.ListOpenIDConnectProvidersRequest

	ListPoliciesRequest(*iam.ListPoliciesInput) iam.ListPoliciesRequest

	ListPoliciesGrantingServiceAccessRequest(*iam.ListPoliciesGrantingServiceAccessInput) iam.ListPoliciesGrantingServiceAccessRequest

	ListPolicyVersionsRequest(*iam.ListPolicyVersionsInput) iam.ListPolicyVersionsRequest

	ListRolePoliciesRequest(*iam.ListRolePoliciesInput) iam.ListRolePoliciesRequest

	ListRoleTagsRequest(*iam.ListRoleTagsInput) iam.ListRoleTagsRequest

	ListRolesRequest(*iam.ListRolesInput) iam.ListRolesRequest

	ListSAMLProvidersRequest(*iam.ListSAMLProvidersInput) iam.ListSAMLProvidersRequest

	ListSSHPublicKeysRequest(*iam.ListSSHPublicKeysInput) iam.ListSSHPublicKeysRequest

	ListServerCertificatesRequest(*iam.ListServerCertificatesInput) iam.ListServerCertificatesRequest

	ListServiceSpecificCredentialsRequest(*iam.ListServiceSpecificCredentialsInput) iam.ListServiceSpecificCredentialsRequest

	ListSigningCertificatesRequest(*iam.ListSigningCertificatesInput) iam.ListSigningCertificatesRequest

	ListUserPoliciesRequest(*iam.ListUserPoliciesInput) iam.ListUserPoliciesRequest

	ListUserTagsRequest(*iam.ListUserTagsInput) iam.ListUserTagsRequest

	ListUsersRequest(*iam.ListUsersInput) iam.ListUsersRequest

	ListVirtualMFADevicesRequest(*iam.ListVirtualMFADevicesInput) iam.ListVirtualMFADevicesRequest

	PutGroupPolicyRequest(*iam.PutGroupPolicyInput) iam.PutGroupPolicyRequest

	PutRolePermissionsBoundaryRequest(*iam.PutRolePermissionsBoundaryInput) iam.PutRolePermissionsBoundaryRequest

	PutRolePolicyRequest(*iam.PutRolePolicyInput) iam.PutRolePolicyRequest

	PutUserPermissionsBoundaryRequest(*iam.PutUserPermissionsBoundaryInput) iam.PutUserPermissionsBoundaryRequest

	PutUserPolicyRequest(*iam.PutUserPolicyInput) iam.PutUserPolicyRequest

	RemoveClientIDFromOpenIDConnectProviderRequest(*iam.RemoveClientIDFromOpenIDConnectProviderInput) iam.RemoveClientIDFromOpenIDConnectProviderRequest

	RemoveRoleFromInstanceProfileRequest(*iam.RemoveRoleFromInstanceProfileInput) iam.RemoveRoleFromInstanceProfileRequest

	RemoveUserFromGroupRequest(*iam.RemoveUserFromGroupInput) iam.RemoveUserFromGroupRequest

	ResetServiceSpecificCredentialRequest(*iam.ResetServiceSpecificCredentialInput) iam.ResetServiceSpecificCredentialRequest

	ResyncMFADeviceRequest(*iam.ResyncMFADeviceInput) iam.ResyncMFADeviceRequest

	SetDefaultPolicyVersionRequest(*iam.SetDefaultPolicyVersionInput) iam.SetDefaultPolicyVersionRequest

	SimulateCustomPolicyRequest(*iam.SimulateCustomPolicyInput) iam.SimulateCustomPolicyRequest

	SimulatePrincipalPolicyRequest(*iam.SimulatePrincipalPolicyInput) iam.SimulatePrincipalPolicyRequest

	TagRoleRequest(*iam.TagRoleInput) iam.TagRoleRequest

	TagUserRequest(*iam.TagUserInput) iam.TagUserRequest

	UntagRoleRequest(*iam.UntagRoleInput) iam.UntagRoleRequest

	UntagUserRequest(*iam.UntagUserInput) iam.UntagUserRequest

	UpdateAccessKeyRequest(*iam.UpdateAccessKeyInput) iam.UpdateAccessKeyRequest

	UpdateAccountPasswordPolicyRequest(*iam.UpdateAccountPasswordPolicyInput) iam.UpdateAccountPasswordPolicyRequest

	UpdateAssumeRolePolicyRequest(*iam.UpdateAssumeRolePolicyInput) iam.UpdateAssumeRolePolicyRequest

	UpdateGroupRequest(*iam.UpdateGroupInput) iam.UpdateGroupRequest

	UpdateLoginProfileRequest(*iam.UpdateLoginProfileInput) iam.UpdateLoginProfileRequest

	UpdateOpenIDConnectProviderThumbprintRequest(*iam.UpdateOpenIDConnectProviderThumbprintInput) iam.UpdateOpenIDConnectProviderThumbprintRequest

	UpdateRoleRequest(*iam.UpdateRoleInput) iam.UpdateRoleRequest

	UpdateRoleDescriptionRequest(*iam.UpdateRoleDescriptionInput) iam.UpdateRoleDescriptionRequest

	UpdateSAMLProviderRequest(*iam.UpdateSAMLProviderInput) iam.UpdateSAMLProviderRequest

	UpdateSSHPublicKeyRequest(*iam.UpdateSSHPublicKeyInput) iam.UpdateSSHPublicKeyRequest

	UpdateServerCertificateRequest(*iam.UpdateServerCertificateInput) iam.UpdateServerCertificateRequest

	UpdateServiceSpecificCredentialRequest(*iam.UpdateServiceSpecificCredentialInput) iam.UpdateServiceSpecificCredentialRequest

	UpdateSigningCertificateRequest(*iam.UpdateSigningCertificateInput) iam.UpdateSigningCertificateRequest

	UpdateUserRequest(*iam.UpdateUserInput) iam.UpdateUserRequest

	UploadSSHPublicKeyRequest(*iam.UploadSSHPublicKeyInput) iam.UploadSSHPublicKeyRequest

	UploadServerCertificateRequest(*iam.UploadServerCertificateInput) iam.UploadServerCertificateRequest

	UploadSigningCertificateRequest(*iam.UploadSigningCertificateInput) iam.UploadSigningCertificateRequest

	WaitUntilInstanceProfileExists(*iam.GetInstanceProfileInput) error
	WaitUntilInstanceProfileExistsWithContext(aws.Context, *iam.GetInstanceProfileInput, ...aws.WaiterOption) error

	WaitUntilUserExists(*iam.GetUserInput) error
	WaitUntilUserExistsWithContext(aws.Context, *iam.GetUserInput, ...aws.WaiterOption) error
}

var _ IAMAPI = (*iam.IAM)(nil)
