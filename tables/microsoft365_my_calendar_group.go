package tables

import (
	"context"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/selefra/selefra-provider-microsoft365/microsoft365_client"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups"
	"github.com/selefra/selefra-provider-microsoft365/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableMicrosoft365MyCalendarGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMicrosoft365MyCalendarGroupGenerator{}

func (x *TableMicrosoft365MyCalendarGroupGenerator) GetTableName() string {
	return "microsoft365_my_calendar_group"
}

func (x *TableMicrosoft365MyCalendarGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMicrosoft365MyCalendarGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMicrosoft365MyCalendarGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMicrosoft365MyCalendarGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, adapter, err := microsoft365_client.GetGraphClient(ctx, taskClient.(*microsoft365_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			userID, _ := microsoft365_client.GetUserID(ctx, taskClient.(*microsoft365_client.Client).Config)

			input := &calendargroups.CalendarGroupsRequestBuilderGetQueryParameters{}

			pageSize := int64(9999)

			input.Top = microsoft365_client.Int32(int32(pageSize))

			options := &calendargroups.CalendarGroupsRequestBuilderGetRequestConfiguration{
				QueryParameters: input,
			}

			result, err := client.UsersById(userID.(string)).CalendarGroups().Get(ctx, options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateCalendarGroupCollectionResponseFromDiscriminatorValue)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
				calendarGroup := pageItem.(models.CalendarGroupable)
				resultChannel <- &Microsoft365CalendarGroupInfo{calendarGroup, userID.(string)}

				return true
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableMicrosoft365MyCalendarGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMicrosoft365MyCalendarGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The group name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The group's unique identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("change_key").ColumnType(schema.ColumnTypeString).Description("Identifies the version of the calendar group. Every time the calendar group is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("class_id").ColumnType(schema.ColumnTypeString).Description("The class identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The Azure Tenant ID where the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("ID or email of the user.").Build(),
	}
}

func (x *TableMicrosoft365MyCalendarGroupGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableMicrosoft365MyCalendarGenerator{}),
	}
}
