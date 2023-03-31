package tables

import (
	"context"
	"github.com/selefra/selefra-provider-microsoft365/microsoft365_client"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars"
	"github.com/selefra/selefra-provider-microsoft365/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMicrosoft365MyCalendarGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMicrosoft365MyCalendarGenerator{}

func (x *TableMicrosoft365MyCalendarGenerator) GetTableName() string {
	return "microsoft365_my_calendar"
}

func (x *TableMicrosoft365MyCalendarGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMicrosoft365MyCalendarGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMicrosoft365MyCalendarGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMicrosoft365MyCalendarGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			groupData := task.ParentRawResult.(*Microsoft365CalendarGroupInfo)

			client, adapter, err := microsoft365_client.GetGraphClient(ctx, taskClient.(*microsoft365_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			userID, _ := microsoft365_client.GetUserID(ctx, taskClient.(*microsoft365_client.Client).Config)

			input := &calendars.CalendarsRequestBuilderGetQueryParameters{}

			pageSize := int64(9999)

			input.Top = microsoft365_client.Int32(int32(pageSize))

			options := &calendars.CalendarsRequestBuilderGetRequestConfiguration{
				QueryParameters: input,
			}

			result, err := client.UsersById(userID.(string)).CalendarGroupsById(*groupData.GetId()).Calendars().Get(ctx, options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateCalendarCollectionResponseFromDiscriminatorValue)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
				calendar := pageItem.(models.Calendarable)
				resultChannel <- &Microsoft365CalendarInfo{calendar, *groupData.GetId(), userID.(string)}
				return true
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

type Microsoft365CalendarGroupInfo struct {
	models.CalendarGroupable
	UserID string
}
type Microsoft365CalendarInfo struct {
	models.Calendarable
	CalendarGroupID string
	UserID          string
}

func (x *TableMicrosoft365MyCalendarGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMicrosoft365MyCalendarGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("calendar_group_id").ColumnType(schema.ColumnTypeString).Description("The ID of the group.").
			Extractor(column_value_extractor.StructSelector("CalendarGroupID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The Azure Tenant ID where the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("ID or email of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_tallying_responses").ColumnType(schema.ColumnTypeBool).Description("Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_online_meeting_providers").ColumnType(schema.ColumnTypeJSON).Description("Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_default_calendar").ColumnType(schema.ColumnTypeBool).Description("True if this is the default calendar where new events are created by default, false otherwise.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("can_view_private_items").ColumnType(schema.ColumnTypeBool).Description("True if the user can read calendar items that have been marked private, false otherwise.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("can_edit").ColumnType(schema.ColumnTypeBool).Description("True if the user can write to the calendar, false otherwise.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hex_color").ColumnType(schema.ColumnTypeString).Description("The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("change_key").ColumnType(schema.ColumnTypeString).Description("Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_online_meeting_provider").ColumnType(schema.ColumnTypeString).Description("The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("multi_value_extended_properties").ColumnType(schema.ColumnTypeJSON).Description("The collection of multi-value extended properties defined for the calendar.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The calendar name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeJSON).Description("Represents the user who created or added the calendar.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions").ColumnType(schema.ColumnTypeJSON).Description("Represents the user who created or added the calendar.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The calendar's unique identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("color").ColumnType(schema.ColumnTypeString).Description("Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("can_share").ColumnType(schema.ColumnTypeBool).Description("True if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_removable").ColumnType(schema.ColumnTypeBool).Description("Indicates whether this user calendar can be deleted from the user mailbox.").Build(),
	}
}

func (x *TableMicrosoft365MyCalendarGenerator) GetSubTables() []*schema.Table {
	return nil
}
