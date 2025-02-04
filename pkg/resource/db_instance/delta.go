// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package db_instance

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.AllocatedStorage, b.ko.Spec.AllocatedStorage) {
		delta.Add("Spec.AllocatedStorage", a.ko.Spec.AllocatedStorage, b.ko.Spec.AllocatedStorage)
	} else if a.ko.Spec.AllocatedStorage != nil && b.ko.Spec.AllocatedStorage != nil {
		if *a.ko.Spec.AllocatedStorage != *b.ko.Spec.AllocatedStorage {
			delta.Add("Spec.AllocatedStorage", a.ko.Spec.AllocatedStorage, b.ko.Spec.AllocatedStorage)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AutoMinorVersionUpgrade, b.ko.Spec.AutoMinorVersionUpgrade) {
		delta.Add("Spec.AutoMinorVersionUpgrade", a.ko.Spec.AutoMinorVersionUpgrade, b.ko.Spec.AutoMinorVersionUpgrade)
	} else if a.ko.Spec.AutoMinorVersionUpgrade != nil && b.ko.Spec.AutoMinorVersionUpgrade != nil {
		if *a.ko.Spec.AutoMinorVersionUpgrade != *b.ko.Spec.AutoMinorVersionUpgrade {
			delta.Add("Spec.AutoMinorVersionUpgrade", a.ko.Spec.AutoMinorVersionUpgrade, b.ko.Spec.AutoMinorVersionUpgrade)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AvailabilityZone, b.ko.Spec.AvailabilityZone) {
		delta.Add("Spec.AvailabilityZone", a.ko.Spec.AvailabilityZone, b.ko.Spec.AvailabilityZone)
	} else if a.ko.Spec.AvailabilityZone != nil && b.ko.Spec.AvailabilityZone != nil {
		if *a.ko.Spec.AvailabilityZone != *b.ko.Spec.AvailabilityZone {
			delta.Add("Spec.AvailabilityZone", a.ko.Spec.AvailabilityZone, b.ko.Spec.AvailabilityZone)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.BackupRetentionPeriod, b.ko.Spec.BackupRetentionPeriod) {
		delta.Add("Spec.BackupRetentionPeriod", a.ko.Spec.BackupRetentionPeriod, b.ko.Spec.BackupRetentionPeriod)
	} else if a.ko.Spec.BackupRetentionPeriod != nil && b.ko.Spec.BackupRetentionPeriod != nil {
		if *a.ko.Spec.BackupRetentionPeriod != *b.ko.Spec.BackupRetentionPeriod {
			delta.Add("Spec.BackupRetentionPeriod", a.ko.Spec.BackupRetentionPeriod, b.ko.Spec.BackupRetentionPeriod)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CharacterSetName, b.ko.Spec.CharacterSetName) {
		delta.Add("Spec.CharacterSetName", a.ko.Spec.CharacterSetName, b.ko.Spec.CharacterSetName)
	} else if a.ko.Spec.CharacterSetName != nil && b.ko.Spec.CharacterSetName != nil {
		if *a.ko.Spec.CharacterSetName != *b.ko.Spec.CharacterSetName {
			delta.Add("Spec.CharacterSetName", a.ko.Spec.CharacterSetName, b.ko.Spec.CharacterSetName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CopyTagsToSnapshot, b.ko.Spec.CopyTagsToSnapshot) {
		delta.Add("Spec.CopyTagsToSnapshot", a.ko.Spec.CopyTagsToSnapshot, b.ko.Spec.CopyTagsToSnapshot)
	} else if a.ko.Spec.CopyTagsToSnapshot != nil && b.ko.Spec.CopyTagsToSnapshot != nil {
		if *a.ko.Spec.CopyTagsToSnapshot != *b.ko.Spec.CopyTagsToSnapshot {
			delta.Add("Spec.CopyTagsToSnapshot", a.ko.Spec.CopyTagsToSnapshot, b.ko.Spec.CopyTagsToSnapshot)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBClusterIdentifier, b.ko.Spec.DBClusterIdentifier) {
		delta.Add("Spec.DBClusterIdentifier", a.ko.Spec.DBClusterIdentifier, b.ko.Spec.DBClusterIdentifier)
	} else if a.ko.Spec.DBClusterIdentifier != nil && b.ko.Spec.DBClusterIdentifier != nil {
		if *a.ko.Spec.DBClusterIdentifier != *b.ko.Spec.DBClusterIdentifier {
			delta.Add("Spec.DBClusterIdentifier", a.ko.Spec.DBClusterIdentifier, b.ko.Spec.DBClusterIdentifier)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBInstanceClass, b.ko.Spec.DBInstanceClass) {
		delta.Add("Spec.DBInstanceClass", a.ko.Spec.DBInstanceClass, b.ko.Spec.DBInstanceClass)
	} else if a.ko.Spec.DBInstanceClass != nil && b.ko.Spec.DBInstanceClass != nil {
		if *a.ko.Spec.DBInstanceClass != *b.ko.Spec.DBInstanceClass {
			delta.Add("Spec.DBInstanceClass", a.ko.Spec.DBInstanceClass, b.ko.Spec.DBInstanceClass)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBInstanceIdentifier, b.ko.Spec.DBInstanceIdentifier) {
		delta.Add("Spec.DBInstanceIdentifier", a.ko.Spec.DBInstanceIdentifier, b.ko.Spec.DBInstanceIdentifier)
	} else if a.ko.Spec.DBInstanceIdentifier != nil && b.ko.Spec.DBInstanceIdentifier != nil {
		if *a.ko.Spec.DBInstanceIdentifier != *b.ko.Spec.DBInstanceIdentifier {
			delta.Add("Spec.DBInstanceIdentifier", a.ko.Spec.DBInstanceIdentifier, b.ko.Spec.DBInstanceIdentifier)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBName, b.ko.Spec.DBName) {
		delta.Add("Spec.DBName", a.ko.Spec.DBName, b.ko.Spec.DBName)
	} else if a.ko.Spec.DBName != nil && b.ko.Spec.DBName != nil {
		if *a.ko.Spec.DBName != *b.ko.Spec.DBName {
			delta.Add("Spec.DBName", a.ko.Spec.DBName, b.ko.Spec.DBName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBParameterGroupName, b.ko.Spec.DBParameterGroupName) {
		delta.Add("Spec.DBParameterGroupName", a.ko.Spec.DBParameterGroupName, b.ko.Spec.DBParameterGroupName)
	} else if a.ko.Spec.DBParameterGroupName != nil && b.ko.Spec.DBParameterGroupName != nil {
		if *a.ko.Spec.DBParameterGroupName != *b.ko.Spec.DBParameterGroupName {
			delta.Add("Spec.DBParameterGroupName", a.ko.Spec.DBParameterGroupName, b.ko.Spec.DBParameterGroupName)
		}
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.DBSecurityGroupNames, b.ko.Spec.DBSecurityGroupNames) {
		delta.Add("Spec.DBSecurityGroupNames", a.ko.Spec.DBSecurityGroupNames, b.ko.Spec.DBSecurityGroupNames)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBSubnetGroupName, b.ko.Spec.DBSubnetGroupName) {
		delta.Add("Spec.DBSubnetGroupName", a.ko.Spec.DBSubnetGroupName, b.ko.Spec.DBSubnetGroupName)
	} else if a.ko.Spec.DBSubnetGroupName != nil && b.ko.Spec.DBSubnetGroupName != nil {
		if *a.ko.Spec.DBSubnetGroupName != *b.ko.Spec.DBSubnetGroupName {
			delta.Add("Spec.DBSubnetGroupName", a.ko.Spec.DBSubnetGroupName, b.ko.Spec.DBSubnetGroupName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DeletionProtection, b.ko.Spec.DeletionProtection) {
		delta.Add("Spec.DeletionProtection", a.ko.Spec.DeletionProtection, b.ko.Spec.DeletionProtection)
	} else if a.ko.Spec.DeletionProtection != nil && b.ko.Spec.DeletionProtection != nil {
		if *a.ko.Spec.DeletionProtection != *b.ko.Spec.DeletionProtection {
			delta.Add("Spec.DeletionProtection", a.ko.Spec.DeletionProtection, b.ko.Spec.DeletionProtection)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Domain, b.ko.Spec.Domain) {
		delta.Add("Spec.Domain", a.ko.Spec.Domain, b.ko.Spec.Domain)
	} else if a.ko.Spec.Domain != nil && b.ko.Spec.Domain != nil {
		if *a.ko.Spec.Domain != *b.ko.Spec.Domain {
			delta.Add("Spec.Domain", a.ko.Spec.Domain, b.ko.Spec.Domain)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DomainIAMRoleName, b.ko.Spec.DomainIAMRoleName) {
		delta.Add("Spec.DomainIAMRoleName", a.ko.Spec.DomainIAMRoleName, b.ko.Spec.DomainIAMRoleName)
	} else if a.ko.Spec.DomainIAMRoleName != nil && b.ko.Spec.DomainIAMRoleName != nil {
		if *a.ko.Spec.DomainIAMRoleName != *b.ko.Spec.DomainIAMRoleName {
			delta.Add("Spec.DomainIAMRoleName", a.ko.Spec.DomainIAMRoleName, b.ko.Spec.DomainIAMRoleName)
		}
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.EnableCloudwatchLogsExports, b.ko.Spec.EnableCloudwatchLogsExports) {
		delta.Add("Spec.EnableCloudwatchLogsExports", a.ko.Spec.EnableCloudwatchLogsExports, b.ko.Spec.EnableCloudwatchLogsExports)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EnableCustomerOwnedIP, b.ko.Spec.EnableCustomerOwnedIP) {
		delta.Add("Spec.EnableCustomerOwnedIP", a.ko.Spec.EnableCustomerOwnedIP, b.ko.Spec.EnableCustomerOwnedIP)
	} else if a.ko.Spec.EnableCustomerOwnedIP != nil && b.ko.Spec.EnableCustomerOwnedIP != nil {
		if *a.ko.Spec.EnableCustomerOwnedIP != *b.ko.Spec.EnableCustomerOwnedIP {
			delta.Add("Spec.EnableCustomerOwnedIP", a.ko.Spec.EnableCustomerOwnedIP, b.ko.Spec.EnableCustomerOwnedIP)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EnableIAMDatabaseAuthentication, b.ko.Spec.EnableIAMDatabaseAuthentication) {
		delta.Add("Spec.EnableIAMDatabaseAuthentication", a.ko.Spec.EnableIAMDatabaseAuthentication, b.ko.Spec.EnableIAMDatabaseAuthentication)
	} else if a.ko.Spec.EnableIAMDatabaseAuthentication != nil && b.ko.Spec.EnableIAMDatabaseAuthentication != nil {
		if *a.ko.Spec.EnableIAMDatabaseAuthentication != *b.ko.Spec.EnableIAMDatabaseAuthentication {
			delta.Add("Spec.EnableIAMDatabaseAuthentication", a.ko.Spec.EnableIAMDatabaseAuthentication, b.ko.Spec.EnableIAMDatabaseAuthentication)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EnablePerformanceInsights, b.ko.Spec.EnablePerformanceInsights) {
		delta.Add("Spec.EnablePerformanceInsights", a.ko.Spec.EnablePerformanceInsights, b.ko.Spec.EnablePerformanceInsights)
	} else if a.ko.Spec.EnablePerformanceInsights != nil && b.ko.Spec.EnablePerformanceInsights != nil {
		if *a.ko.Spec.EnablePerformanceInsights != *b.ko.Spec.EnablePerformanceInsights {
			delta.Add("Spec.EnablePerformanceInsights", a.ko.Spec.EnablePerformanceInsights, b.ko.Spec.EnablePerformanceInsights)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Engine, b.ko.Spec.Engine) {
		delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
	} else if a.ko.Spec.Engine != nil && b.ko.Spec.Engine != nil {
		if *a.ko.Spec.Engine != *b.ko.Spec.Engine {
			delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion) {
		delta.Add("Spec.EngineVersion", a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion)
	} else if a.ko.Spec.EngineVersion != nil && b.ko.Spec.EngineVersion != nil {
		if *a.ko.Spec.EngineVersion != *b.ko.Spec.EngineVersion {
			delta.Add("Spec.EngineVersion", a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.IOPS, b.ko.Spec.IOPS) {
		delta.Add("Spec.IOPS", a.ko.Spec.IOPS, b.ko.Spec.IOPS)
	} else if a.ko.Spec.IOPS != nil && b.ko.Spec.IOPS != nil {
		if *a.ko.Spec.IOPS != *b.ko.Spec.IOPS {
			delta.Add("Spec.IOPS", a.ko.Spec.IOPS, b.ko.Spec.IOPS)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID) {
		delta.Add("Spec.KMSKeyID", a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID)
	} else if a.ko.Spec.KMSKeyID != nil && b.ko.Spec.KMSKeyID != nil {
		if *a.ko.Spec.KMSKeyID != *b.ko.Spec.KMSKeyID {
			delta.Add("Spec.KMSKeyID", a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.LicenseModel, b.ko.Spec.LicenseModel) {
		delta.Add("Spec.LicenseModel", a.ko.Spec.LicenseModel, b.ko.Spec.LicenseModel)
	} else if a.ko.Spec.LicenseModel != nil && b.ko.Spec.LicenseModel != nil {
		if *a.ko.Spec.LicenseModel != *b.ko.Spec.LicenseModel {
			delta.Add("Spec.LicenseModel", a.ko.Spec.LicenseModel, b.ko.Spec.LicenseModel)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MasterUserPassword, b.ko.Spec.MasterUserPassword) {
		delta.Add("Spec.MasterUserPassword", a.ko.Spec.MasterUserPassword, b.ko.Spec.MasterUserPassword)
	} else if a.ko.Spec.MasterUserPassword != nil && b.ko.Spec.MasterUserPassword != nil {
		if *a.ko.Spec.MasterUserPassword != *b.ko.Spec.MasterUserPassword {
			delta.Add("Spec.MasterUserPassword", a.ko.Spec.MasterUserPassword, b.ko.Spec.MasterUserPassword)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MasterUsername, b.ko.Spec.MasterUsername) {
		delta.Add("Spec.MasterUsername", a.ko.Spec.MasterUsername, b.ko.Spec.MasterUsername)
	} else if a.ko.Spec.MasterUsername != nil && b.ko.Spec.MasterUsername != nil {
		if *a.ko.Spec.MasterUsername != *b.ko.Spec.MasterUsername {
			delta.Add("Spec.MasterUsername", a.ko.Spec.MasterUsername, b.ko.Spec.MasterUsername)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MaxAllocatedStorage, b.ko.Spec.MaxAllocatedStorage) {
		delta.Add("Spec.MaxAllocatedStorage", a.ko.Spec.MaxAllocatedStorage, b.ko.Spec.MaxAllocatedStorage)
	} else if a.ko.Spec.MaxAllocatedStorage != nil && b.ko.Spec.MaxAllocatedStorage != nil {
		if *a.ko.Spec.MaxAllocatedStorage != *b.ko.Spec.MaxAllocatedStorage {
			delta.Add("Spec.MaxAllocatedStorage", a.ko.Spec.MaxAllocatedStorage, b.ko.Spec.MaxAllocatedStorage)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MonitoringInterval, b.ko.Spec.MonitoringInterval) {
		delta.Add("Spec.MonitoringInterval", a.ko.Spec.MonitoringInterval, b.ko.Spec.MonitoringInterval)
	} else if a.ko.Spec.MonitoringInterval != nil && b.ko.Spec.MonitoringInterval != nil {
		if *a.ko.Spec.MonitoringInterval != *b.ko.Spec.MonitoringInterval {
			delta.Add("Spec.MonitoringInterval", a.ko.Spec.MonitoringInterval, b.ko.Spec.MonitoringInterval)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MonitoringRoleARN, b.ko.Spec.MonitoringRoleARN) {
		delta.Add("Spec.MonitoringRoleARN", a.ko.Spec.MonitoringRoleARN, b.ko.Spec.MonitoringRoleARN)
	} else if a.ko.Spec.MonitoringRoleARN != nil && b.ko.Spec.MonitoringRoleARN != nil {
		if *a.ko.Spec.MonitoringRoleARN != *b.ko.Spec.MonitoringRoleARN {
			delta.Add("Spec.MonitoringRoleARN", a.ko.Spec.MonitoringRoleARN, b.ko.Spec.MonitoringRoleARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MultiAZ, b.ko.Spec.MultiAZ) {
		delta.Add("Spec.MultiAZ", a.ko.Spec.MultiAZ, b.ko.Spec.MultiAZ)
	} else if a.ko.Spec.MultiAZ != nil && b.ko.Spec.MultiAZ != nil {
		if *a.ko.Spec.MultiAZ != *b.ko.Spec.MultiAZ {
			delta.Add("Spec.MultiAZ", a.ko.Spec.MultiAZ, b.ko.Spec.MultiAZ)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.NcharCharacterSetName, b.ko.Spec.NcharCharacterSetName) {
		delta.Add("Spec.NcharCharacterSetName", a.ko.Spec.NcharCharacterSetName, b.ko.Spec.NcharCharacterSetName)
	} else if a.ko.Spec.NcharCharacterSetName != nil && b.ko.Spec.NcharCharacterSetName != nil {
		if *a.ko.Spec.NcharCharacterSetName != *b.ko.Spec.NcharCharacterSetName {
			delta.Add("Spec.NcharCharacterSetName", a.ko.Spec.NcharCharacterSetName, b.ko.Spec.NcharCharacterSetName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.OptionGroupName, b.ko.Spec.OptionGroupName) {
		delta.Add("Spec.OptionGroupName", a.ko.Spec.OptionGroupName, b.ko.Spec.OptionGroupName)
	} else if a.ko.Spec.OptionGroupName != nil && b.ko.Spec.OptionGroupName != nil {
		if *a.ko.Spec.OptionGroupName != *b.ko.Spec.OptionGroupName {
			delta.Add("Spec.OptionGroupName", a.ko.Spec.OptionGroupName, b.ko.Spec.OptionGroupName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PerformanceInsightsKMSKeyID, b.ko.Spec.PerformanceInsightsKMSKeyID) {
		delta.Add("Spec.PerformanceInsightsKMSKeyID", a.ko.Spec.PerformanceInsightsKMSKeyID, b.ko.Spec.PerformanceInsightsKMSKeyID)
	} else if a.ko.Spec.PerformanceInsightsKMSKeyID != nil && b.ko.Spec.PerformanceInsightsKMSKeyID != nil {
		if *a.ko.Spec.PerformanceInsightsKMSKeyID != *b.ko.Spec.PerformanceInsightsKMSKeyID {
			delta.Add("Spec.PerformanceInsightsKMSKeyID", a.ko.Spec.PerformanceInsightsKMSKeyID, b.ko.Spec.PerformanceInsightsKMSKeyID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PerformanceInsightsRetentionPeriod, b.ko.Spec.PerformanceInsightsRetentionPeriod) {
		delta.Add("Spec.PerformanceInsightsRetentionPeriod", a.ko.Spec.PerformanceInsightsRetentionPeriod, b.ko.Spec.PerformanceInsightsRetentionPeriod)
	} else if a.ko.Spec.PerformanceInsightsRetentionPeriod != nil && b.ko.Spec.PerformanceInsightsRetentionPeriod != nil {
		if *a.ko.Spec.PerformanceInsightsRetentionPeriod != *b.ko.Spec.PerformanceInsightsRetentionPeriod {
			delta.Add("Spec.PerformanceInsightsRetentionPeriod", a.ko.Spec.PerformanceInsightsRetentionPeriod, b.ko.Spec.PerformanceInsightsRetentionPeriod)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Port, b.ko.Spec.Port) {
		delta.Add("Spec.Port", a.ko.Spec.Port, b.ko.Spec.Port)
	} else if a.ko.Spec.Port != nil && b.ko.Spec.Port != nil {
		if *a.ko.Spec.Port != *b.ko.Spec.Port {
			delta.Add("Spec.Port", a.ko.Spec.Port, b.ko.Spec.Port)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PreferredBackupWindow, b.ko.Spec.PreferredBackupWindow) {
		delta.Add("Spec.PreferredBackupWindow", a.ko.Spec.PreferredBackupWindow, b.ko.Spec.PreferredBackupWindow)
	} else if a.ko.Spec.PreferredBackupWindow != nil && b.ko.Spec.PreferredBackupWindow != nil {
		if *a.ko.Spec.PreferredBackupWindow != *b.ko.Spec.PreferredBackupWindow {
			delta.Add("Spec.PreferredBackupWindow", a.ko.Spec.PreferredBackupWindow, b.ko.Spec.PreferredBackupWindow)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow) {
		delta.Add("Spec.PreferredMaintenanceWindow", a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow)
	} else if a.ko.Spec.PreferredMaintenanceWindow != nil && b.ko.Spec.PreferredMaintenanceWindow != nil {
		if *a.ko.Spec.PreferredMaintenanceWindow != *b.ko.Spec.PreferredMaintenanceWindow {
			delta.Add("Spec.PreferredMaintenanceWindow", a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.ProcessorFeatures, b.ko.Spec.ProcessorFeatures) {
		delta.Add("Spec.ProcessorFeatures", a.ko.Spec.ProcessorFeatures, b.ko.Spec.ProcessorFeatures)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PromotionTier, b.ko.Spec.PromotionTier) {
		delta.Add("Spec.PromotionTier", a.ko.Spec.PromotionTier, b.ko.Spec.PromotionTier)
	} else if a.ko.Spec.PromotionTier != nil && b.ko.Spec.PromotionTier != nil {
		if *a.ko.Spec.PromotionTier != *b.ko.Spec.PromotionTier {
			delta.Add("Spec.PromotionTier", a.ko.Spec.PromotionTier, b.ko.Spec.PromotionTier)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PubliclyAccessible, b.ko.Spec.PubliclyAccessible) {
		delta.Add("Spec.PubliclyAccessible", a.ko.Spec.PubliclyAccessible, b.ko.Spec.PubliclyAccessible)
	} else if a.ko.Spec.PubliclyAccessible != nil && b.ko.Spec.PubliclyAccessible != nil {
		if *a.ko.Spec.PubliclyAccessible != *b.ko.Spec.PubliclyAccessible {
			delta.Add("Spec.PubliclyAccessible", a.ko.Spec.PubliclyAccessible, b.ko.Spec.PubliclyAccessible)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StorageEncrypted, b.ko.Spec.StorageEncrypted) {
		delta.Add("Spec.StorageEncrypted", a.ko.Spec.StorageEncrypted, b.ko.Spec.StorageEncrypted)
	} else if a.ko.Spec.StorageEncrypted != nil && b.ko.Spec.StorageEncrypted != nil {
		if *a.ko.Spec.StorageEncrypted != *b.ko.Spec.StorageEncrypted {
			delta.Add("Spec.StorageEncrypted", a.ko.Spec.StorageEncrypted, b.ko.Spec.StorageEncrypted)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StorageType, b.ko.Spec.StorageType) {
		delta.Add("Spec.StorageType", a.ko.Spec.StorageType, b.ko.Spec.StorageType)
	} else if a.ko.Spec.StorageType != nil && b.ko.Spec.StorageType != nil {
		if *a.ko.Spec.StorageType != *b.ko.Spec.StorageType {
			delta.Add("Spec.StorageType", a.ko.Spec.StorageType, b.ko.Spec.StorageType)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.Tags, b.ko.Spec.Tags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TDECredentialARN, b.ko.Spec.TDECredentialARN) {
		delta.Add("Spec.TDECredentialARN", a.ko.Spec.TDECredentialARN, b.ko.Spec.TDECredentialARN)
	} else if a.ko.Spec.TDECredentialARN != nil && b.ko.Spec.TDECredentialARN != nil {
		if *a.ko.Spec.TDECredentialARN != *b.ko.Spec.TDECredentialARN {
			delta.Add("Spec.TDECredentialARN", a.ko.Spec.TDECredentialARN, b.ko.Spec.TDECredentialARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TDECredentialPassword, b.ko.Spec.TDECredentialPassword) {
		delta.Add("Spec.TDECredentialPassword", a.ko.Spec.TDECredentialPassword, b.ko.Spec.TDECredentialPassword)
	} else if a.ko.Spec.TDECredentialPassword != nil && b.ko.Spec.TDECredentialPassword != nil {
		if *a.ko.Spec.TDECredentialPassword != *b.ko.Spec.TDECredentialPassword {
			delta.Add("Spec.TDECredentialPassword", a.ko.Spec.TDECredentialPassword, b.ko.Spec.TDECredentialPassword)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Timezone, b.ko.Spec.Timezone) {
		delta.Add("Spec.Timezone", a.ko.Spec.Timezone, b.ko.Spec.Timezone)
	} else if a.ko.Spec.Timezone != nil && b.ko.Spec.Timezone != nil {
		if *a.ko.Spec.Timezone != *b.ko.Spec.Timezone {
			delta.Add("Spec.Timezone", a.ko.Spec.Timezone, b.ko.Spec.Timezone)
		}
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.VPCSecurityGroupIDs, b.ko.Spec.VPCSecurityGroupIDs) {
		delta.Add("Spec.VPCSecurityGroupIDs", a.ko.Spec.VPCSecurityGroupIDs, b.ko.Spec.VPCSecurityGroupIDs)
	}

	return delta
}
