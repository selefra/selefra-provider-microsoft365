# Table: microsoft365_team

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| display_name | string | X | √ | The name of the team. | 
| is_archived | bool | X | √ | True if this team is in read-only mode. | 
| internal_id | string | X | √ | A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API. | 
| web_url | string | X | √ | A hyperlink that will go to the team in the Microsoft Teams client. | 
| template | json | X | √ | The template this team was created from. | 
| id | string | X | √ | The unique id of the team. | 
| classification | string | X | √ | An optional label. Typically describes the data or business sensitivity of the team. Must match one of a pre-configured set in the tenant's directory. | 
| title | string | X | √ | Title of the resource. | 
| description | string | X | √ | A description for the team. | 
| created_date_time | timestamp | X | √ | Date and time when the team was created. | 
| visibility | string | X | √ | The visibility of the group and team. Defaults to Public. | 
| specialization | string | X | √ | Indicates whether the team is intended for a particular use case. Each team specialization has access to unique behaviors and experiences targeted to its use case. | 
| summary | json | X | √ | Specifies the team's summary. | 
| tenant_id | string | X | √ | The Azure Tenant ID where the resource is located. | 


