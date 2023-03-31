package tables

import (
	"context"
	"github.com/iancoleman/strcase"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"
	"github.com/selefra/selefra-provider-microsoft365/microsoft365_client"
	"github.com/selefra/selefra-provider-microsoft365/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableMicrosoft365MyMailMessageGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMicrosoft365MyMailMessageGenerator{}

func (x *TableMicrosoft365MyMailMessageGenerator) GetTableName() string {
	return "microsoft365_my_mail_message"
}

func (x *TableMicrosoft365MyMailMessageGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMicrosoft365MyMailMessageGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMicrosoft365MyMailMessageGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMicrosoft365MyMailMessageGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, adapter, err := microsoft365_client.GetGraphClient(ctx, taskClient.(*microsoft365_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			userID, _ := microsoft365_client.GetUserID(ctx, taskClient.(*microsoft365_client.Client).Config)

			input := &messages.MessagesRequestBuilderGetQueryParameters{}

			pageSize := int64(9999)

			input.Top = microsoft365_client.Int32(int32(pageSize))

			givenColumns := []string{}
			selectColumns := buildMailMessageRequestFields(ctx, givenColumns)
			input.Select = selectColumns

			options := &messages.MessagesRequestBuilderGetRequestConfiguration{
				QueryParameters: input,
			}

			result, err := client.UsersById(userID.(string)).Messages().Get(ctx, options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateMessageCollectionResponseFromDiscriminatorValue)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
				message := pageItem.(models.Messageable)
				resultChannel <- &Microsoft365MailMessageInfo{message, userID.(string)}

				return true
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type Microsoft365MailMessageInfo struct {
	models.Messageable
	UserID string
}

func buildMailMessageRequestFields(ctx context.Context, queryColumns []string) []string {
	var selectColumns []string

	for _, columnName := range queryColumns {
		if columnName == "title" || columnName == "filter" || columnName == "user_id" || columnName == "_ctx" || columnName == "tenant_id" {
			continue
		}
		selectColumns = append(selectColumns, strcase.ToLowerCamel(columnName))
	}

	return selectColumns
}

func (x *TableMicrosoft365MyMailMessageGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMicrosoft365MyMailMessageGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("attachments").ColumnType(schema.ColumnTypeJSON).Description("The attachments of the message.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The Azure Tenant ID where the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sent_date_time").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the message was sent.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internet_message_id").ColumnType(schema.ColumnTypeString).Description("The message ID in the format specified by RFC2822.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_link").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent_folder_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier for the message's parent mailFolder.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bcc_recipients").ColumnType(schema.ColumnTypeJSON).Description("The Bcc: recipients for the message.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cc_recipients").ColumnType(schema.ColumnTypeJSON).Description("The Cc: recipients for the message.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("body_preview").ColumnType(schema.ColumnTypeString).Description("The first 255 characters of the message body in text format.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("received_date_time").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the message was received.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inference_classification").ColumnType(schema.ColumnTypeString).Description("The classification of the message for the user, based on inferred relevance or importance, or on an explicit override. The possible values are: focused or other.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("from").ColumnType(schema.ColumnTypeJSON).Description("The owner of the mailbox from which the message is sent.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("categories").ColumnType(schema.ColumnTypeJSON).Description("The categories associated with the message.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("change_key").ColumnType(schema.ColumnTypeString).Description("The version of the message.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reply_to").ColumnType(schema.ColumnTypeJSON).Description("The email addresses to use when replying.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("to_recipients").ColumnType(schema.ColumnTypeJSON).Description("The To: recipients for the message.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sender").ColumnType(schema.ColumnTypeJSON).Description("The date and time the message was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subject").ColumnType(schema.ColumnTypeString).Description("The subject of the message.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("conversation_id").ColumnType(schema.ColumnTypeString).Description("The ID of the conversation the email belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_delivery_receipt_requested").ColumnType(schema.ColumnTypeBool).Description("Indicates whether a read receipt is requested for the message.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_read").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the message has been read.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the message.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_date_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("ID or email of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_read_receipt_requested").ColumnType(schema.ColumnTypeBool).Description("Indicates whether a read receipt is requested for the message.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("body").ColumnType(schema.ColumnTypeJSON).Description("The body of the message. It can be in HTML or text format.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_attachments").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the message has attachments.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_draft").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("importance").ColumnType(schema.ColumnTypeString).Description("The importance of the message. The possible values are: low, normal, and high.").Build(),
	}
}

func (x *TableMicrosoft365MyMailMessageGenerator) GetSubTables() []*schema.Table {
	return nil
}
