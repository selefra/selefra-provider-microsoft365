package tables

import (
	"context"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/contacts"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/selefra/selefra-provider-microsoft365/microsoft365_client"
	"github.com/selefra/selefra-provider-microsoft365/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableMicrosoft365OrganizationContactGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMicrosoft365OrganizationContactGenerator{}

func (x *TableMicrosoft365OrganizationContactGenerator) GetTableName() string {
	return "microsoft365_organization_contact"
}

func (x *TableMicrosoft365OrganizationContactGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMicrosoft365OrganizationContactGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMicrosoft365OrganizationContactGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMicrosoft365OrganizationContactGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, adapter, err := microsoft365_client.GetGraphClient(ctx, taskClient.(*microsoft365_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			input := &contacts.ContactsRequestBuilderGetQueryParameters{}

			pageSize := int64(999)

			input.Top = microsoft365_client.Int32(int32(pageSize))

			options := &contacts.ContactsRequestBuilderGetRequestConfiguration{
				QueryParameters: input,
			}

			result, err := client.Contacts().Get(ctx, options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateOrgContactCollectionResponseFromDiscriminatorValue)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
				contact := pageItem.(models.OrgContactable)
				resultChannel <- &Microsoft365OrgContactInfo{contact}

				return true
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type Microsoft365OrgContactInfo struct {
	models.OrgContactable
}

func (x *TableMicrosoft365OrganizationContactGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMicrosoft365OrganizationContactGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The contact's display name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("given_name").ColumnType(schema.ColumnTypeString).Description("The contact's given name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The contact's unique identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("direct_reports").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("manager").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mail_nickname").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("on_premises_last_sync_date_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("on_premises_sync_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("on_premises_provisioning_errors").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("surname").ColumnType(schema.ColumnTypeString).Description("The contact's surname.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("department").ColumnType(schema.ColumnTypeString).Description("The contact's department.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("job_title").ColumnType(schema.ColumnTypeString).Description("The contact’s job title.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mail").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phones").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("proxy_addresses").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transitive_member_of").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The Azure Tenant ID where the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("company_name").ColumnType(schema.ColumnTypeString).Description("The name of the contact's company.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("addresses").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("member_of").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
	}
}

func (x *TableMicrosoft365OrganizationContactGenerator) GetSubTables() []*schema.Table {
	return nil
}
