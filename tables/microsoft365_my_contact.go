package tables

import (
	"context"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/selefra/selefra-provider-microsoft365/microsoft365_client"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/contacts"
	"github.com/selefra/selefra-provider-microsoft365/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableMicrosoft365MyContactGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMicrosoft365MyContactGenerator{}

func (x *TableMicrosoft365MyContactGenerator) GetTableName() string {
	return "microsoft365_my_contact"
}

func (x *TableMicrosoft365MyContactGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMicrosoft365MyContactGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMicrosoft365MyContactGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMicrosoft365MyContactGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, adapter, err := microsoft365_client.GetGraphClient(ctx, taskClient.(*microsoft365_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			userID, _ := microsoft365_client.GetUserID(ctx, taskClient.(*microsoft365_client.Client).Config)

			input := &contacts.ContactsRequestBuilderGetQueryParameters{}

			pageSize := int64(9999)

			input.Top = microsoft365_client.Int32(int32(pageSize))

			options := &contacts.ContactsRequestBuilderGetRequestConfiguration{
				QueryParameters: input,
			}

			result, err := client.UsersById(userID.(string)).Contacts().Get(ctx, options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateContactCollectionResponseFromDiscriminatorValue)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
				contact := pageItem.(models.Contactable)
				resultChannel <- &Microsoft365ContactInfo{contact, userID.(string)}

				return true
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type Microsoft365ContactInfo struct {
	models.Contactable
	UserID string
}

func (x *TableMicrosoft365MyContactGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMicrosoft365MyContactGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("job_title").ColumnType(schema.ColumnTypeString).Description("The contact’s job title.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("office_location").ColumnType(schema.ColumnTypeString).Description("The location of the contact's office.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("yomi_given_name").ColumnType(schema.ColumnTypeString).Description("The phonetic Japanese given name (first name) of the contact.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("home_address").ColumnType(schema.ColumnTypeString).Description("The contact's home address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("ID or email of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("company_name").ColumnType(schema.ColumnTypeString).Description("The name of the contact's company.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("initials").ColumnType(schema.ColumnTypeString).Description("The contact's initials.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("surname").ColumnType(schema.ColumnTypeString).Description("The contact's surname.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("department").ColumnType(schema.ColumnTypeString).Description("The contact's department.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("personal_notes").ColumnType(schema.ColumnTypeString).Description("The user's notes about the contact.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profession").ColumnType(schema.ColumnTypeString).Description("The contact's profession.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spouse_name").ColumnType(schema.ColumnTypeString).Description("The name of the contact's spouse/partner.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("business_phones").ColumnType(schema.ColumnTypeString).Description("The contact's business phone numbers.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("children").ColumnType(schema.ColumnTypeString).Description("The names of the contact's children.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("home_phones").ColumnType(schema.ColumnTypeString).Description("The contact's home phone numbers.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("middle_name").ColumnType(schema.ColumnTypeString).Description("The contact's middle name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The Azure Tenant ID where the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("business_home_page").ColumnType(schema.ColumnTypeString).Description("The business home page of the contact.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_as").ColumnType(schema.ColumnTypeString).Description("The name the contact is filed under.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The contact's display name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The contact's unique identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assistant_name").ColumnType(schema.ColumnTypeString).Description("The name of the contact's assistant.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nick_name").ColumnType(schema.ColumnTypeString).Description("The contact's nickname.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("yomi_surname").ColumnType(schema.ColumnTypeString).Description("The phonetic Japanese surname (last name) of the contact.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("given_name").ColumnType(schema.ColumnTypeString).Description("The contact's given name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("manager").ColumnType(schema.ColumnTypeString).Description("The name of the contact's manager.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent_folder_id").ColumnType(schema.ColumnTypeString).Description("The ID of the contact's parent folder.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_addresses").ColumnType(schema.ColumnTypeJSON).Description("The contact's email addresses.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("other_address").ColumnType(schema.ColumnTypeString).Description("Other addresses for the contact.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("birthday").ColumnType(schema.ColumnTypeTimestamp).Description("The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generation").ColumnType(schema.ColumnTypeString).Description("The contact's generation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("yomi_company_name").ColumnType(schema.ColumnTypeString).Description("The phonetic Japanese company name of the contact.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("business_address").ColumnType(schema.ColumnTypeString).Description("The contact's business address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("im_addresses").ColumnType(schema.ColumnTypeJSON).Description("The contact's instant messaging (IM) addresses.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mobile_phone").ColumnType(schema.ColumnTypeString).Description("The contact's mobile phone number.").Build(),
	}
}

func (x *TableMicrosoft365MyContactGenerator) GetSubTables() []*schema.Table {
	return nil
}
