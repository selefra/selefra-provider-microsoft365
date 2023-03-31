package tables

import (
	"context"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview"
	"github.com/selefra/selefra-provider-microsoft365/microsoft365_client"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/events"
	"github.com/selefra/selefra-provider-microsoft365/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableMicrosoft365MyCalendarEventGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMicrosoft365MyCalendarEventGenerator{}

func (x *TableMicrosoft365MyCalendarEventGenerator) GetTableName() string {
	return "microsoft365_my_calendar_event"
}

func (x *TableMicrosoft365MyCalendarEventGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMicrosoft365MyCalendarEventGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMicrosoft365MyCalendarEventGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMicrosoft365MyCalendarEventGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, adapter, err := microsoft365_client.GetGraphClient(ctx, taskClient.(*microsoft365_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			userID, _ := microsoft365_client.GetUserID(ctx, taskClient.(*microsoft365_client.Client).Config)

			input := &calendarview.CalendarViewRequestBuilderGetQueryParameters{}

			pageSize := int64(9999)

			input.Top = microsoft365_client.Int32(int32(pageSize))

			var result models.EventCollectionResponseable

			// Filter event using timestamp
			var startTime, endTime string

			currentTime := time.Now().Format(time.RFC3339)
			if startTime != "" && endTime != "" {
				input.StartDateTime = &startTime
				input.EndDateTime = &endTime
			} else if startTime != "" && endTime == "" {
				input.StartDateTime = &startTime
				input.EndDateTime = &currentTime
			} else if startTime == "" && endTime != "" {
				input.StartDateTime = &currentTime
				input.EndDateTime = &endTime
			}

			if startTime != "" || endTime != "" {
				options := &calendarview.CalendarViewRequestBuilderGetRequestConfiguration{
					QueryParameters: input,
				}

				result, err = client.UsersById(userID.(string)).CalendarView().Get(ctx, options)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
			} else {
				input := &events.EventsRequestBuilderGetQueryParameters{}

				pageSize := int64(9999)

				input.Top = microsoft365_client.Int32(int32(pageSize))

				options := &events.EventsRequestBuilderGetRequestConfiguration{
					QueryParameters: input,
				}

				result, err = client.UsersById(userID.(string)).Events().Get(ctx, options)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
			}

			pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateEventCollectionResponseFromDiscriminatorValue)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
				event := pageItem.(models.Eventable)
				resultChannel <- &Microsoft365CalendarEventInfo{event, userID.(string)}

				return true
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type Microsoft365CalendarEventInfo struct {
	models.Eventable
	UserID string
}

func (x *TableMicrosoft365MyCalendarEventGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMicrosoft365MyCalendarEventGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("original_end_time_zone").ColumnType(schema.ColumnTypeString).Description("The end time zone that was set when the event was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ical_uid").ColumnType(schema.ColumnTypeString).Description("A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("categories").ColumnType(schema.ColumnTypeJSON).Description("The categories associated with the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start").ColumnType(schema.ColumnTypeJSON).Description("The start date, time, and time zone of the event. By default, the start time is in UTC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("body").ColumnType(schema.ColumnTypeJSON).Description("The body of the message associated with the event. It can be in HTML or text format.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locations").ColumnType(schema.ColumnTypeJSON).Description("The locations where the event is held or attended from.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("ID or email of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_link").ColumnType(schema.ColumnTypeString).Description("The URL to open the event in Outlook on the web.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("original_start_time_zone").ColumnType(schema.ColumnTypeString).Description("The start time zone that was set when the event was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_new_time_proposals").ColumnType(schema.ColumnTypeBool).Description("True if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Default is true.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end").ColumnType(schema.ColumnTypeJSON).Description("The date, time, and time zone that the event ends. By default, the end time is in UTC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeTimestamp).Description("The start date and time of the event. By default, the start time is in UTC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("show_as").ColumnType(schema.ColumnTypeString).Description("The status to show. Possible values are: free, tentative, busy, oof, workingElsewhere, unknown.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("online_meeting_provider").ColumnType(schema.ColumnTypeString).Description("Represents the online meeting service provider. By default, onlineMeetingProvider is unknown. The possible values are unknown, teamsForBusiness, skypeForBusiness, and skypeForConsumer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeJSON).Description("The location of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attendees").ColumnType(schema.ColumnTypeJSON).Description("The collection of attendees for the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reminder_minutes_before_start").ColumnType(schema.ColumnTypeInt).Description("The number of minutes before the event start time that the reminder alert occurs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transaction_id").ColumnType(schema.ColumnTypeString).Description("A custom identifier specified by a client app for the server to avoid redundant POST operations in case of client retries to create the same event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("series_master_id").ColumnType(schema.ColumnTypeString).Description("The ID for the recurring series master item, if this event is part of a recurring series.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organizer").ColumnType(schema.ColumnTypeJSON).Description("The organizer of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recurrence").ColumnType(schema.ColumnTypeJSON).Description("The recurrence pattern for the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_date_time").ColumnType(schema.ColumnTypeTimestamp).Description("The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_all_day").ColumnType(schema.ColumnTypeBool).Description("True if the event lasts all day. If true, regardless of whether it's a single-day or multi-day event, start and end time must be set to midnight and be in the same time zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeTimestamp).Description("The end date and time of the event. By default, the end time is in UTC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("change_key").ColumnType(schema.ColumnTypeString).Description("Identifies the version of the event object. Every time the event is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_reminder_on").ColumnType(schema.ColumnTypeBool).Description("True if an alert is set to remind the user of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_attachments").ColumnType(schema.ColumnTypeBool).Description("True if the event has attachments.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_draft").ColumnType(schema.ColumnTypeBool).Description("True if the user has updated the meeting in Outlook but has not sent the updates to attendees.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hide_attendees").ColumnType(schema.ColumnTypeBool).Description("If set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subject").ColumnType(schema.ColumnTypeString).Description("The text of the event's subject line.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("response_status").ColumnType(schema.ColumnTypeJSON).Description("Indicates the type of response sent in response to an event message.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_cancelled").ColumnType(schema.ColumnTypeBool).Description("True  if the event has been canceled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_organizer").ColumnType(schema.ColumnTypeBool).Description("True if the calendar owner (specified by the owner property of the calendar) is the organizer of the event (specified by the organizer property of the event).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date_time").ColumnType(schema.ColumnTypeTimestamp).Description("The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sensitivity").ColumnType(schema.ColumnTypeString).Description("The sensitivity of the event. Possible values are: normal, personal, private, confidential.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("response_requested").ColumnType(schema.ColumnTypeBool).Description("If true, it represents the organizer would like an invitee to send a response to the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_online_meeting").ColumnType(schema.ColumnTypeBool).Description("True if this event has online meeting information (that is, onlineMeeting points to an onlineMeetingInfo resource), false otherwise. Default is false (onlineMeeting is null).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The Azure Tenant ID where the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("online_meeting_url").ColumnType(schema.ColumnTypeString).Description("A URL for an online meeting. The property is set only when an organizer specifies in Outlook that an event is an online meeting such as Skype.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("importance").ColumnType(schema.ColumnTypeInt).Description("The importance of the event. The possible values are: low, normal, high.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("online_meeting").ColumnType(schema.ColumnTypeJSON).Description("Details for an attendee to join the meeting online. Default is null.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("body_preview").ColumnType(schema.ColumnTypeString).Description("The preview of the message associated with the event in text format.").Build(),
	}
}

func (x *TableMicrosoft365MyCalendarEventGenerator) GetSubTables() []*schema.Table {
	return nil
}
