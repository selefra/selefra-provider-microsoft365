package provider

import (
	"github.com/selefra/selefra-provider-microsoft365/table_schema_generator"
	"github.com/selefra/selefra-provider-microsoft365/tables"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableMicrosoft365OrganizationContactGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableMicrosoft365MyMailMessageGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableMicrosoft365TeamGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableMicrosoft365MyDriveGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableMicrosoft365MyCalendarGroupGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableMicrosoft365MyCalendarEventGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableMicrosoft365MyContactGenerator{}),
	}
}
