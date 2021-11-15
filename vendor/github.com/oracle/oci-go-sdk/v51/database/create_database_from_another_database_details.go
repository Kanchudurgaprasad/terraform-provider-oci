// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v51/common"
)

// CreateDatabaseFromAnotherDatabaseDetails The representation of CreateDatabaseFromAnotherDatabaseDetails
type CreateDatabaseFromAnotherDatabaseDetails struct {

	// The database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The password to open the TDE wallet.
	BackupTDEPassword *string `mandatory:"false" json:"backupTDEPassword"`

	// The `DB_UNIQUE_NAME` of the Oracle Database being backed up.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The display name of the database to be created from the backup. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	DbName *string `mandatory:"false" json:"dbName"`

	// The point in time of the original database from which the new database is created. If not specifed, the latest backup is used to create the database.
	TimeStampForPointInTimeRecovery *common.SDKTime `mandatory:"false" json:"timeStampForPointInTimeRecovery"`
}

func (m CreateDatabaseFromAnotherDatabaseDetails) String() string {
	return common.PointerString(m)
}
