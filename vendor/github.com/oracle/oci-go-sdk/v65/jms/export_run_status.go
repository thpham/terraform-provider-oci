// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// ExportRunStatusEnum Enum with underlying type: string
type ExportRunStatusEnum string

// Set of constants representing the allowable values for ExportRunStatusEnum
const (
	ExportRunStatusScheduled  ExportRunStatusEnum = "SCHEDULED"
	ExportRunStatusPending    ExportRunStatusEnum = "PENDING"
	ExportRunStatusInProgress ExportRunStatusEnum = "IN_PROGRESS"
	ExportRunStatusFailed     ExportRunStatusEnum = "FAILED"
	ExportRunStatusRetrying   ExportRunStatusEnum = "RETRYING"
	ExportRunStatusSucceeded  ExportRunStatusEnum = "SUCCEEDED"
)

var mappingExportRunStatusEnum = map[string]ExportRunStatusEnum{
	"SCHEDULED":   ExportRunStatusScheduled,
	"PENDING":     ExportRunStatusPending,
	"IN_PROGRESS": ExportRunStatusInProgress,
	"FAILED":      ExportRunStatusFailed,
	"RETRYING":    ExportRunStatusRetrying,
	"SUCCEEDED":   ExportRunStatusSucceeded,
}

var mappingExportRunStatusEnumLowerCase = map[string]ExportRunStatusEnum{
	"scheduled":   ExportRunStatusScheduled,
	"pending":     ExportRunStatusPending,
	"in_progress": ExportRunStatusInProgress,
	"failed":      ExportRunStatusFailed,
	"retrying":    ExportRunStatusRetrying,
	"succeeded":   ExportRunStatusSucceeded,
}

// GetExportRunStatusEnumValues Enumerates the set of values for ExportRunStatusEnum
func GetExportRunStatusEnumValues() []ExportRunStatusEnum {
	values := make([]ExportRunStatusEnum, 0)
	for _, v := range mappingExportRunStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetExportRunStatusEnumStringValues Enumerates the set of values in String for ExportRunStatusEnum
func GetExportRunStatusEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"PENDING",
		"IN_PROGRESS",
		"FAILED",
		"RETRYING",
		"SUCCEEDED",
	}
}

// GetMappingExportRunStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportRunStatusEnum(val string) (ExportRunStatusEnum, bool) {
	enum, ok := mappingExportRunStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
