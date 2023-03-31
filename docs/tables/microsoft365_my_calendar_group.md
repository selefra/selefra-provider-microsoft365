# Table: microsoft365_my_calendar_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | The group name. | 
| id | string | X | √ | The group's unique identifier. | 
| change_key | string | X | √ | Identifies the version of the calendar group. Every time the calendar group is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object. | 
| class_id | string | X | √ | The class identifier. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The Azure Tenant ID where the resource is located. | 
| user_id | string | X | √ | ID or email of the user. | 


