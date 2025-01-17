// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_identity_domains_account_recovery_setting", IdentityDomainsAccountRecoverySettingResource())
	tfresource.RegisterResource("oci_identity_domains_api_key", IdentityDomainsApiKeyResource())
	tfresource.RegisterResource("oci_identity_domains_app", IdentityDomainsAppResource())
	tfresource.RegisterResource("oci_identity_domains_app_role", IdentityDomainsAppRoleResource())
	tfresource.RegisterResource("oci_identity_domains_approval_workflow", IdentityDomainsApprovalWorkflowResource())
	tfresource.RegisterResource("oci_identity_domains_approval_workflow_assignment", IdentityDomainsApprovalWorkflowAssignmentResource())
	tfresource.RegisterResource("oci_identity_domains_approval_workflow_step", IdentityDomainsApprovalWorkflowStepResource())
	tfresource.RegisterResource("oci_identity_domains_auth_token", IdentityDomainsAuthTokenResource())
	tfresource.RegisterResource("oci_identity_domains_authentication_factor_setting", IdentityDomainsAuthenticationFactorSettingResource())
	tfresource.RegisterResource("oci_identity_domains_cloud_gate", IdentityDomainsCloudGateResource())
	tfresource.RegisterResource("oci_identity_domains_cloud_gate_mapping", IdentityDomainsCloudGateMappingResource())
	tfresource.RegisterResource("oci_identity_domains_cloud_gate_server", IdentityDomainsCloudGateServerResource())
	tfresource.RegisterResource("oci_identity_domains_condition", IdentityDomainsConditionResource())
	tfresource.RegisterResource("oci_identity_domains_customer_secret_key", IdentityDomainsCustomerSecretKeyResource())
	tfresource.RegisterResource("oci_identity_domains_dynamic_resource_group", IdentityDomainsDynamicResourceGroupResource())
	tfresource.RegisterResource("oci_identity_domains_grant", IdentityDomainsGrantResource())
	tfresource.RegisterResource("oci_identity_domains_group", IdentityDomainsGroupResource())
	tfresource.RegisterResource("oci_identity_domains_identity_provider", IdentityDomainsIdentityProviderResource())
	tfresource.RegisterResource("oci_identity_domains_identity_setting", IdentityDomainsIdentitySettingResource())
	tfresource.RegisterResource("oci_identity_domains_kmsi_setting", IdentityDomainsKmsiSettingResource())
	tfresource.RegisterResource("oci_identity_domains_my_api_key", IdentityDomainsMyApiKeyResource())
	tfresource.RegisterResource("oci_identity_domains_my_auth_token", IdentityDomainsMyAuthTokenResource())
	tfresource.RegisterResource("oci_identity_domains_my_customer_secret_key", IdentityDomainsMyCustomerSecretKeyResource())
	tfresource.RegisterResource("oci_identity_domains_my_oauth2client_credential", IdentityDomainsMyOAuth2ClientCredentialResource())
	tfresource.RegisterResource("oci_identity_domains_my_request", IdentityDomainsMyRequestResource())
	tfresource.RegisterResource("oci_identity_domains_my_smtp_credential", IdentityDomainsMySmtpCredentialResource())
	tfresource.RegisterResource("oci_identity_domains_my_support_account", IdentityDomainsMySupportAccountResource())
	tfresource.RegisterResource("oci_identity_domains_my_user_db_credential", IdentityDomainsMyUserDbCredentialResource())
	tfresource.RegisterResource("oci_identity_domains_network_perimeter", IdentityDomainsNetworkPerimeterResource())
	tfresource.RegisterResource("oci_identity_domains_notification_setting", IdentityDomainsNotificationSettingResource())
	tfresource.RegisterResource("oci_identity_domains_oauth2client_credential", IdentityDomainsOAuth2ClientCredentialResource())
	tfresource.RegisterResource("oci_identity_domains_oauth_client_certificate", IdentityDomainsOAuthClientCertificateResource())
	tfresource.RegisterResource("oci_identity_domains_oauth_partner_certificate", IdentityDomainsOAuthPartnerCertificateResource())
	tfresource.RegisterResource("oci_identity_domains_password_policy", IdentityDomainsPasswordPolicyResource())
	tfresource.RegisterResource("oci_identity_domains_policy", IdentityDomainsPolicyResource())
	tfresource.RegisterResource("oci_identity_domains_rule", IdentityDomainsRuleResource())
	tfresource.RegisterResource("oci_identity_domains_security_question", IdentityDomainsSecurityQuestionResource())
	tfresource.RegisterResource("oci_identity_domains_security_question_setting", IdentityDomainsSecurityQuestionSettingResource())
	tfresource.RegisterResource("oci_identity_domains_self_registration_profile", IdentityDomainsSelfRegistrationProfileResource())
	tfresource.RegisterResource("oci_identity_domains_setting", IdentityDomainsSettingResource())
	tfresource.RegisterResource("oci_identity_domains_smtp_credential", IdentityDomainsSmtpCredentialResource())
	tfresource.RegisterResource("oci_identity_domains_user", IdentityDomainsUserResource())
	tfresource.RegisterResource("oci_identity_domains_user_db_credential", IdentityDomainsUserDbCredentialResource())
}
