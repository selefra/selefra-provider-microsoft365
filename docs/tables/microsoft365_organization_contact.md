# Table: microsoft365_organization_contact

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| display_name | string | X | √ | The contact's display name. | 
| given_name | string | X | √ | The contact's given name. | 
| id | string | X | √ | The contact's unique identifier. | 
| direct_reports | json | X | √ |  | 
| manager | json | X | √ |  | 
| mail_nickname | string | X | √ |  | 
| on_premises_last_sync_date_time | timestamp | X | √ |  | 
| on_premises_sync_enabled | bool | X | √ |  | 
| on_premises_provisioning_errors | json | X | √ |  | 
| surname | string | X | √ | The contact's surname. | 
| department | string | X | √ | The contact's department. | 
| job_title | string | X | √ | The contact’s job title. | 
| mail | string | X | √ |  | 
| phones | json | X | √ |  | 
| proxy_addresses | json | X | √ |  | 
| transitive_member_of | json | X | √ |  | 
| tenant_id | string | X | √ | The Azure Tenant ID where the resource is located. | 
| company_name | string | X | √ | The name of the contact's company. | 
| addresses | json | X | √ |  | 
| member_of | json | X | √ |  | 
| title | string | X | √ | Title of the resource. | 


