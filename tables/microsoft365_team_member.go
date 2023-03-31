package tables

import (
	"context"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/selefra/selefra-provider-microsoft365/microsoft365_client"

	"github.com/microsoftgraph/msgraph-sdk-go/groups/item/members"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/selefra/selefra-provider-microsoft365/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMicrosoft365TeamMemberGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMicrosoft365TeamMemberGenerator{}

func (x *TableMicrosoft365TeamMemberGenerator) GetTableName() string {
	return "microsoft365_team_member"
}

func (x *TableMicrosoft365TeamMemberGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMicrosoft365TeamMemberGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMicrosoft365TeamMemberGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMicrosoft365TeamMemberGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			teamData := task.ParentRawResult.(*Microsoft365TeamInfo)
			teamID := teamData.ID

			client, adapter, err := microsoft365_client.GetGraphClient(ctx, taskClient.(*microsoft365_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			headers := map[string]string{
				"ConsistencyLevel": "eventual",
			}

			includeCount := true
			requestParameters := &members.MembersRequestBuilderGetQueryParameters{
				Count: &includeCount,
			}

			pageSize := int64(999)

			requestParameters.Top = microsoft365_client.Int32(int32(pageSize))

			config := &members.MembersRequestBuilderGetRequestConfiguration{
				Headers:         headers,
				QueryParameters: requestParameters,
			}

			members, err := client.GroupsById(teamID).Members().Get(ctx, config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			pageIterator, err := msgraphcore.NewPageIterator(members, adapter, models.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
				member := pageItem.(models.DirectoryObjectable)
				resultChannel <- &Microsoft365TeamMemberInfo{
					TeamID:   teamID,
					MemberID: *member.GetId(),
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

type Microsoft365TeamMemberInfo struct {
	TeamID   string
	MemberID string
}

func (x *TableMicrosoft365TeamMemberGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMicrosoft365TeamMemberGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("team_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier of the team.").
			Extractor(column_value_extractor.StructSelector("TeamID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("member_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier of the member.").
			Extractor(column_value_extractor.StructSelector("MemberID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The Azure Tenant ID where the resource is located.").Build(),
	}
}

func (x *TableMicrosoft365TeamMemberGenerator) GetSubTables() []*schema.Table {
	return nil
}
