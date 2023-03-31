package tables

import (
	"context"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/selefra/selefra-provider-microsoft365/microsoft365_client"

	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/selefra/selefra-provider-microsoft365/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMicrosoft365MyDriveFileGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMicrosoft365MyDriveFileGenerator{}

func (x *TableMicrosoft365MyDriveFileGenerator) GetTableName() string {
	return "microsoft365_my_drive_file"
}

func (x *TableMicrosoft365MyDriveFileGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMicrosoft365MyDriveFileGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMicrosoft365MyDriveFileGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMicrosoft365MyDriveFileGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			driveData := task.ParentRawResult.(*Microsoft365DriveInfo)

			var driveID string
			if driveData != nil {
				driveID = *driveData.GetId()
			}

			client, adapter, err := microsoft365_client.GetGraphClient(ctx, taskClient.(*microsoft365_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			userID, _ := microsoft365_client.GetUserID(ctx, taskClient.(*microsoft365_client.Client).Config)

			result, err := client.UsersById(userID.(string)).DrivesById(driveID).Root().Children().Get(ctx, nil)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateDriveItemCollectionResponseFromDiscriminatorValue)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
				var resultFiles []Microsoft365DriveItemInfo

				item := pageItem.(models.DriveItemable)

				resultFiles = append(resultFiles, Microsoft365DriveItemInfo{item, driveID, userID.(string)})
				if item.GetFolder() != nil && item.GetFolder().GetChildCount() != nil && *item.GetFolder().GetChildCount() != 0 {
					childData, err := expandDriveFolders(ctx, client, adapter, item, userID.(string), driveID)
					if err != nil {
						return false
					}
					resultFiles = append(resultFiles, childData...)
				}

				for _, i := range resultFiles {
					resultChannel <- i

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

type Microsoft365DriveItemInfo struct {
	models.DriveItemable
	DriveID string
	UserID  string
}

func expandDriveFolders(ctx context.Context, c *msgraphsdkgo.GraphServiceClient, a *msgraphsdkgo.GraphRequestAdapter, data models.DriveItemable, userID string, driveID string) ([]Microsoft365DriveItemInfo, error) {
	var items []Microsoft365DriveItemInfo
	result, err := c.UsersById(userID).DrivesById(driveID).ItemsById(*data.GetId()).Children().Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, a, models.CreateDriveItemCollectionResponseFromDiscriminatorValue)
	if err != nil {

		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		var data []Microsoft365DriveItemInfo

		item := pageItem.(models.DriveItemable)

		data = append(data, Microsoft365DriveItemInfo{item, driveID, userID})
		if item.GetFolder() != nil && item.GetFolder().GetChildCount() != nil && *item.GetFolder().GetChildCount() != 0 {
			childData, err := expandDriveFolders(ctx, c, a, item, userID, driveID)
			if err != nil {
				return false
			}
			data = append(data, childData...)
		}
		items = append(items, data...)

		return true
	})
	if err != nil {

		return nil, err
	}

	return items, nil
}

func (x *TableMicrosoft365MyDriveFileGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMicrosoft365MyDriveFileGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).Description("ETag for the entire item (metadata + content).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ctag").ColumnType(schema.ColumnTypeString).Description("An eTag for the content of the item. This eTag is not changed if only the metadata is changed. This property is not returned if the item is a folder.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_by").ColumnType(schema.ColumnTypeJSON).Description("Identity of the user, device, and application which last modified the item.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent_Reference").ColumnType(schema.ColumnTypeJSON).Description("Parent information, if the item has a parent.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("ID or email of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The Azure Tenant ID where the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("path").ColumnType(schema.ColumnTypeString).Description("URL that displays the resource in the browser.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_url").ColumnType(schema.ColumnTypeString).Description("URL that displays the resource in the browser.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_date_time").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time the item was last modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeInt).Description("Size of the item in bytes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_dav_url").ColumnType(schema.ColumnTypeString).Description("WebDAV compatible URL for the item.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The unique identifier of the item within the Drive.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("drive_id").ColumnType(schema.ColumnTypeString).Description("The unique id of the drive.").
			Extractor(column_value_extractor.StructSelector("DriveID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("folder").ColumnType(schema.ColumnTypeJSON).Description("Folder metadata, if the item is a folder.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the item (filename and extension).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date_time").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time of item creation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeJSON).Description("Identity of the user, device, and application which created the item.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file").ColumnType(schema.ColumnTypeJSON).Description("File metadata, if the item is a file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Provides a user-visible description of the item.").Build(),
	}
}

func (x *TableMicrosoft365MyDriveFileGenerator) GetSubTables() []*schema.Table {
	return nil
}
