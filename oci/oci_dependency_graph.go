// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

var DependencyGraph map[string][]string

func initDependencyGraph() {
	DependencyGraph = make(map[string][]string)

	DependencyGraph["agreement"] = append(DependencyGraph["agreement"], "MarketplaceAcceptedAgreement")
	DependencyGraph["application"] = append(DependencyGraph["application"], "DataflowInvokeRun")
	DependencyGraph["application"] = append(DependencyGraph["application"], "FunctionsFunction")
	DependencyGraph["asset"] = append(DependencyGraph["asset"], "CoreVolumeBackupPolicyAssignment")
	DependencyGraph["autonomousContainerDatabase"] = append(DependencyGraph["autonomousContainerDatabase"], "DatabaseAutonomousDatabase")
	DependencyGraph["autonomousDataWarehouse"] = append(DependencyGraph["autonomousDataWarehouse"], "DatabaseAutonomousDataWarehouseBackup")
	DependencyGraph["autonomousDataWarehouse"] = append(DependencyGraph["autonomousDataWarehouse"], "DatabaseAutonomousDataWarehouseWallet")
	DependencyGraph["autonomousDatabase"] = append(DependencyGraph["autonomousDatabase"], "DatabaseAutonomousDatabaseBackup")
	DependencyGraph["autonomousDatabase"] = append(DependencyGraph["autonomousDatabase"], "DatabaseAutonomousDatabaseInstanceWalletManagement")
	DependencyGraph["autonomousDatabase"] = append(DependencyGraph["autonomousDatabase"], "DatabaseAutonomousDatabaseWallet")
	DependencyGraph["autonomousExadataInfrastructure"] = append(DependencyGraph["autonomousExadataInfrastructure"], "DatabaseAutonomousContainerDatabase")
	DependencyGraph["backupDestination"] = append(DependencyGraph["backupDestination"], "DatabaseDbHome")
	DependencyGraph["backupPolicy"] = append(DependencyGraph["backupPolicy"], "CoreBootVolume")
	DependencyGraph["backupPolicy"] = append(DependencyGraph["backupPolicy"], "CoreVolume")
	DependencyGraph["backupSubnet"] = append(DependencyGraph["backupSubnet"], "DatabaseDbSystem")
	DependencyGraph["bootVolume"] = append(DependencyGraph["bootVolume"], "CoreBootVolumeAttachment")
	DependencyGraph["bootVolume"] = append(DependencyGraph["bootVolume"], "CoreBootVolumeBackup")
	DependencyGraph["bootVolume"] = append(DependencyGraph["bootVolume"], "CoreInstance")
	DependencyGraph["budget"] = append(DependencyGraph["budget"], "BudgetAlertRule")
	DependencyGraph["catalog"] = append(DependencyGraph["catalog"], "DatacatalogConnection")
	DependencyGraph["catalog"] = append(DependencyGraph["catalog"], "DatacatalogDataAsset")
	DependencyGraph["cluster"] = append(DependencyGraph["cluster"], "ContainerengineClusterKubeConfig")
	DependencyGraph["cluster"] = append(DependencyGraph["cluster"], "ContainerengineNodePool")
	DependencyGraph["cpe"] = append(DependencyGraph["cpe"], "CoreIpSecConnection")
	DependencyGraph["cpeDeviceShape"] = append(DependencyGraph["cpeDeviceShape"], "CoreCpe")
	DependencyGraph["crossConnectGroup"] = append(DependencyGraph["crossConnectGroup"], "CoreCrossConnect")
	DependencyGraph["database"] = append(DependencyGraph["database"], "DatabaseBackup")
	DependencyGraph["database"] = append(DependencyGraph["database"], "DatabaseDataGuardAssociation")
	DependencyGraph["dbHome"] = append(DependencyGraph["dbHome"], "DatabaseDatabase")
	DependencyGraph["dbNode"] = append(DependencyGraph["dbNode"], "DatabaseDbNodeConsoleConnection")
	DependencyGraph["dbSystem"] = append(DependencyGraph["dbSystem"], "DatabaseDbHome")
	DependencyGraph["dedicatedVmHost"] = append(DependencyGraph["dedicatedVmHost"], "CoreInstance")
	DependencyGraph["dhcpOptions"] = append(DependencyGraph["dhcpOptions"], "CoreSubnet")
	DependencyGraph["drg"] = append(DependencyGraph["drg"], "CoreDrgAttachment")
	DependencyGraph["drg"] = append(DependencyGraph["drg"], "CoreIpSecConnection")
	DependencyGraph["drg"] = append(DependencyGraph["drg"], "CoreRemotePeeringConnection")
	DependencyGraph["exadataInfrastructure"] = append(DependencyGraph["exadataInfrastructure"], "DatabaseExadataInfrastructureDownloadConfigFile")
	DependencyGraph["exadataInfrastructure"] = append(DependencyGraph["exadataInfrastructure"], "DatabaseVmCluster")
	DependencyGraph["exadataInfrastructure"] = append(DependencyGraph["exadataInfrastructure"], "DatabaseVmClusterNetwork")
	DependencyGraph["exadataInfrastructure"] = append(DependencyGraph["exadataInfrastructure"], "DatabaseVmClusterNetworkDownloadConfigFile")
	DependencyGraph["exadataInfrastructure"] = append(DependencyGraph["exadataInfrastructure"], "DatabaseVmClusterRecommendedNetwork")
	DependencyGraph["exportSet"] = append(DependencyGraph["exportSet"], "FileStorageExport")
	DependencyGraph["farCrossConnectOrCrossConnectGroup"] = append(DependencyGraph["farCrossConnectOrCrossConnectGroup"], "CoreCrossConnect")
	DependencyGraph["fileSystem"] = append(DependencyGraph["fileSystem"], "FileStorageExport")
	DependencyGraph["fileSystem"] = append(DependencyGraph["fileSystem"], "FileStorageSnapshot")
	DependencyGraph["function"] = append(DependencyGraph["function"], "FunctionsInvokeFunction")
	DependencyGraph["gateway"] = append(DependencyGraph["gateway"], "ApigatewayDeployment")
	DependencyGraph["gateway"] = append(DependencyGraph["gateway"], "CoreVirtualCircuit")
	DependencyGraph["group"] = append(DependencyGraph["group"], "IdentityIdpGroupMapping")
	DependencyGraph["group"] = append(DependencyGraph["group"], "IdentityUserGroupMembership")
	DependencyGraph["healthCheckMonitor"] = append(DependencyGraph["healthCheckMonitor"], "DnsSteeringPolicy")
	DependencyGraph["identityProvider"] = append(DependencyGraph["identityProvider"], "IdentityIdpGroupMapping")
	DependencyGraph["instance"] = append(DependencyGraph["instance"], "CoreBootVolumeAttachment")
	DependencyGraph["instance"] = append(DependencyGraph["instance"], "CoreConsoleHistory")
	DependencyGraph["instance"] = append(DependencyGraph["instance"], "CoreInstanceConsoleConnection")
	DependencyGraph["instance"] = append(DependencyGraph["instance"], "CoreVolumeAttachment")
	DependencyGraph["instanceConfiguration"] = append(DependencyGraph["instanceConfiguration"], "CoreInstancePool")
	DependencyGraph["internetGateway"] = append(DependencyGraph["internetGateway"], "CoreRouteTable")
	DependencyGraph["key"] = append(DependencyGraph["key"], "KmsDecryptedData")
	DependencyGraph["key"] = append(DependencyGraph["key"], "KmsEncryptedData")
	DependencyGraph["key"] = append(DependencyGraph["key"], "KmsGeneratedKey")
	DependencyGraph["key"] = append(DependencyGraph["key"], "KmsKeyVersion")
	DependencyGraph["key"] = append(DependencyGraph["key"], "VaultSecret")
	DependencyGraph["kmsKey"] = append(DependencyGraph["kmsKey"], "ContainerengineCluster")
	DependencyGraph["kmsKey"] = append(DependencyGraph["kmsKey"], "CoreBootVolume")
	DependencyGraph["kmsKey"] = append(DependencyGraph["kmsKey"], "CoreVolume")
	DependencyGraph["kmsKey"] = append(DependencyGraph["kmsKey"], "FileStorageFileSystem")
	DependencyGraph["kmsKey"] = append(DependencyGraph["kmsKey"], "ObjectStorageBucket")
	DependencyGraph["listing"] = append(DependencyGraph["listing"], "CoreAppCatalogSubscription")
	DependencyGraph["listing"] = append(DependencyGraph["listing"], "MarketplaceAcceptedAgreement")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerBackend")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerBackendSet")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerCertificate")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerHostname")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerListener")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerPathRouteSet")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerRuleSet")
	DependencyGraph["maintenanceRun"] = append(DependencyGraph["maintenanceRun"], "DatabaseMaintenanceRun")
	DependencyGraph["managedInstance"] = append(DependencyGraph["managedInstance"], "OsmanagementManagedInstanceManagement")
	DependencyGraph["metricCompartment"] = append(DependencyGraph["metricCompartment"], "MonitoringAlarm")
	DependencyGraph["model"] = append(DependencyGraph["model"], "DatascienceModelProvenance")
	DependencyGraph["nearCrossConnectOrCrossConnectGroup"] = append(DependencyGraph["nearCrossConnectOrCrossConnectGroup"], "CoreCrossConnect")
	DependencyGraph["networkSecurityGroup"] = append(DependencyGraph["networkSecurityGroup"], "CoreNetworkSecurityGroupSecurityRule")
	DependencyGraph["parent"] = append(DependencyGraph["parent"], "OsmanagementSoftwareSource")
	DependencyGraph["policy"] = append(DependencyGraph["policy"], "CoreVolumeBackupPolicyAssignment")
	DependencyGraph["privateIp"] = append(DependencyGraph["privateIp"], "CorePublicIp")
	DependencyGraph["project"] = append(DependencyGraph["project"], "DatascienceModel")
	DependencyGraph["project"] = append(DependencyGraph["project"], "DatascienceNotebookSession")
	DependencyGraph["providerService"] = append(DependencyGraph["providerService"], "CoreVirtualCircuit")
	DependencyGraph["routeTable"] = append(DependencyGraph["routeTable"], "CoreDrgAttachment")
	DependencyGraph["routeTable"] = append(DependencyGraph["routeTable"], "CoreLocalPeeringGateway")
	DependencyGraph["routeTable"] = append(DependencyGraph["routeTable"], "CoreServiceGateway")
	DependencyGraph["routeTable"] = append(DependencyGraph["routeTable"], "CoreSubnet")
	DependencyGraph["softwareSource"] = append(DependencyGraph["softwareSource"], "OsmanagementManagedInstanceManagement")
	DependencyGraph["steeringPolicy"] = append(DependencyGraph["steeringPolicy"], "DnsSteeringPolicyAttachment")
	DependencyGraph["streamPool"] = append(DependencyGraph["streamPool"], "StreamingStream")
	DependencyGraph["subnet"] = append(DependencyGraph["subnet"], "ApigatewayGateway")
	DependencyGraph["subnet"] = append(DependencyGraph["subnet"], "CoreInstance")
	DependencyGraph["subnet"] = append(DependencyGraph["subnet"], "DataSafeDataSafePrivateEndpoint")
	DependencyGraph["subnet"] = append(DependencyGraph["subnet"], "DatabaseAutonomousDatabase")
	DependencyGraph["subnet"] = append(DependencyGraph["subnet"], "DatabaseAutonomousExadataInfrastructure")
	DependencyGraph["subnet"] = append(DependencyGraph["subnet"], "DatacatalogCatalogPrivateEndpoint")
	DependencyGraph["subnet"] = append(DependencyGraph["subnet"], "DataintegrationWorkspace")
	DependencyGraph["subnet"] = append(DependencyGraph["subnet"], "FileStorageMountTarget")
	DependencyGraph["tableNameOr"] = append(DependencyGraph["tableNameOr"], "NosqlIndex")
	DependencyGraph["tagDefinition"] = append(DependencyGraph["tagDefinition"], "IdentityTagDefault")
	DependencyGraph["tagNamespace"] = append(DependencyGraph["tagNamespace"], "IdentityTag")
	DependencyGraph["targetCompartment"] = append(DependencyGraph["targetCompartment"], "BudgetBudget")
	DependencyGraph["tenancy"] = append(DependencyGraph["tenancy"], "IdentityRegionSubscription")
	DependencyGraph["tenancy"] = append(DependencyGraph["tenancy"], "OceOceInstance")
	DependencyGraph["topic"] = append(DependencyGraph["topic"], "OnsSubscription")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentityApiKey")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentityAuthToken")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentityCustomerSecretKey")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentitySmtpCredential")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentitySwiftPassword")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentityUiPassword")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentityUserGroupMembership")
	DependencyGraph["vault"] = append(DependencyGraph["vault"], "VaultSecret")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "ContainerengineCluster")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreDhcpOptions")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreDrgAttachment")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreInternetGateway")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreLocalPeeringGateway")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreNatGateway")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreNetworkSecurityGroup")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreRouteTable")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreSecurityList")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreServiceGateway")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreSubnet")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "DataSafeDataSafePrivateEndpoint")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "DataintegrationWorkspace")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "LoadBalancerLoadBalancer")
	DependencyGraph["vmClusterNetwork"] = append(DependencyGraph["vmClusterNetwork"], "DatabaseVmCluster")
	DependencyGraph["vmClusterNetwork"] = append(DependencyGraph["vmClusterNetwork"], "DatabaseVmClusterNetworkDownloadConfigFile")
	DependencyGraph["vnic"] = append(DependencyGraph["vnic"], "CorePrivateIp")
	DependencyGraph["volume"] = append(DependencyGraph["volume"], "CoreVolumeAttachment")
	DependencyGraph["volume"] = append(DependencyGraph["volume"], "CoreVolumeBackup")
	DependencyGraph["volumeGroup"] = append(DependencyGraph["volumeGroup"], "CoreVolumeGroupBackup")
	DependencyGraph["waasPolicy"] = append(DependencyGraph["waasPolicy"], "WaasPurgeCache")
	DependencyGraph["zone"] = append(DependencyGraph["zone"], "DnsSteeringPolicyAttachment")
	DependencyGraph["zoneNameOr"] = append(DependencyGraph["zoneNameOr"], "DnsRrset")
}
