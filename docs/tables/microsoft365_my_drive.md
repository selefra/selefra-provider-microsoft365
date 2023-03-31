# Table: microsoft365_my_drive

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_modified_date_time | timestamp | X | √ | Date and time the item was last modified. | 
| last_modified_by | json | X | √ | Identity of the user, device, and application which last modified the item. | 
| description | string | X | √ | Provide a user-visible description of the drive. | 
| created_by | json | X | √ | Identity of the user, device, or application which created the item. | 
| parent_reference | json | X | √ | Parent information, if the drive has a parent. | 
| name | string | X | √ | The name of the item. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The Azure Tenant ID where the resource is located. | 
| user_id | string | X | √ | ID or email of the user. | 
| id | string | X | √ | The unique identifier of the drive. | 
| drive_type | string | X | √ | Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary. | 
| web_url | string | X | √ | URL that displays the resource in the browser. | 
| created_date_time | timestamp | X | √ | Date and time of item creation. | 
| etag | string | X | √ | Specifies the eTag. | 


