package tables

import (
	"context"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/drives"
	"github.com/selefra/selefra-provider-microsoft365/microsoft365_client"
	"github.com/selefra/selefra-provider-microsoft365/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableMicrosoft365MyDriveGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMicrosoft365MyDriveGenerator{}

func (x *TableMicrosoft365MyDriveGenerator) GetTableName() string {
	return "microsoft365_my_drive"
}

func (x *TableMicrosoft365MyDriveGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMicrosoft365MyDriveGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMicrosoft365MyDriveGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMicrosoft365MyDriveGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, adapter, err := microsoft365_client.GetGraphClient(ctx, taskClient.(*microsoft365_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			userID, _ := microsoft365_client.GetUserID(ctx, taskClient.(*microsoft365_client.Client).Config)

			input := &drives.DrivesRequestBuilderGetQueryParameters{}

			pageSize := int64(9999)
			input.Top = microsoft365_client.Int32(int32(pageSize))

			options := &drives.DrivesRequestBuilderGetRequestConfiguration{
				QueryParameters: input,
			}

			result, err := client.UsersById(userID.(string)).Drives().Get(ctx, options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateDriveCollectionResponseFromDiscriminatorValue)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
				drive := pageItem.(models.Driveable)
				resultChannel <- &Microsoft365DriveInfo{drive, userID.(string)}

				return true
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type Microsoft365DriveInfo struct {
	models.Driveable
	UserID string
}

func (x *TableMicrosoft365MyDriveGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMicrosoft365MyDriveGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_date_time").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time the item was last modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_by").ColumnType(schema.ColumnTypeJSON).Description("Identity of the user, device, and application which last modified the item.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Provide a user-visible description of the drive.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeJSON).Description("Identity of the user, device, or application which created the item.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent_reference").ColumnType(schema.ColumnTypeJSON).Description("Parent information, if the drive has a parent.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the item.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The Azure Tenant ID where the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("ID or email of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The unique identifier of the drive.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("drive_type").ColumnType(schema.ColumnTypeString).Description("Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_url").ColumnType(schema.ColumnTypeString).Description("URL that displays the resource in the browser.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date_time").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time of item creation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).Description("Specifies the eTag.").Build(),
	}
}

func (x *TableMicrosoft365MyDriveGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableMicrosoft365MyDriveFileGenerator{}),
	}
}
