package tables

import (
	"context"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"

	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/selefra/selefra-provider-microsoft365/microsoft365_client"
	"github.com/selefra/selefra-provider-microsoft365/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMicrosoft365TeamGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMicrosoft365TeamGenerator{}

func (x *TableMicrosoft365TeamGenerator) GetTableName() string {
	return "microsoft365_team"
}

func (x *TableMicrosoft365TeamGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMicrosoft365TeamGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMicrosoft365TeamGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMicrosoft365TeamGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			client, adapter, err := microsoft365_client.GetGraphClient(ctx, taskClient.(*microsoft365_client.Client).Config)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			input := &groups.GroupsRequestBuilderGetQueryParameters{
				Select: []string{"id"},
			}

			pageSize := int64(999)

			input.Top = microsoft365_client.Int32(int32(pageSize))

			filterStr := "resourceProvisioningOptions/Any(x:x eq 'Team')"
			input.Filter = &filterStr

			options := &groups.GroupsRequestBuilderGetRequestConfiguration{
				QueryParameters: input,
			}

			result, err := client.Groups().Get(ctx, options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateGroupCollectionResponseFromDiscriminatorValue)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
				group := pageItem.(models.Groupable)
				resultChannel <- &Microsoft365TeamInfo{
					ID: *group.GetId(),
				}

				return true
			})
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type Microsoft365TeamInfo struct {
	models.Teamable
	ID string
}

func (x *TableMicrosoft365TeamGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMicrosoft365TeamGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The name of the team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_archived").ColumnType(schema.ColumnTypeBool).Description("True if this team is in read-only mode.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internal_id").ColumnType(schema.ColumnTypeString).Description("A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_url").ColumnType(schema.ColumnTypeString).Description("A hyperlink that will go to the team in the Microsoft Teams client.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template").ColumnType(schema.ColumnTypeJSON).Description("The template this team was created from.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The unique id of the team.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("classification").ColumnType(schema.ColumnTypeString).Description("An optional label. Typically describes the data or business sensitivity of the team. Must match one of a pre-configured set in the tenant's directory.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A description for the team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date_time").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the team was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("visibility").ColumnType(schema.ColumnTypeString).Description("The visibility of the group and team. Defaults to Public.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("specialization").ColumnType(schema.ColumnTypeString).Description("Indicates whether the team is intended for a particular use case. Each team specialization has access to unique behaviors and experiences targeted to its use case.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("summary").ColumnType(schema.ColumnTypeJSON).Description("Specifies the team's summary.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The Azure Tenant ID where the resource is located.").Build(),
	}
}

func (x *TableMicrosoft365TeamGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableMicrosoft365TeamMemberGenerator{}),
	}
}
