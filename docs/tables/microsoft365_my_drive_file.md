# Table: microsoft365_my_drive_file

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| etag | string | X | √ | ETag for the entire item (metadata + content). | 
| ctag | string | X | √ | An eTag for the content of the item. This eTag is not changed if only the metadata is changed. This property is not returned if the item is a folder. | 
| last_modified_by | json | X | √ | Identity of the user, device, and application which last modified the item. | 
| parent_Reference | json | X | √ | Parent information, if the item has a parent. | 
| title | string | X | √ | Title of the resource. | 
| user_id | string | X | √ | ID or email of the user. | 
| tenant_id | string | X | √ | The Azure Tenant ID where the resource is located. | 
| path | string | X | √ | URL that displays the resource in the browser. | 
| web_url | string | X | √ | URL that displays the resource in the browser. | 
| last_modified_date_time | timestamp | X | √ | Date and time the item was last modified. | 
| size | int | X | √ | Size of the item in bytes. | 
| web_dav_url | string | X | √ | WebDAV compatible URL for the item. | 
| id | string | X | √ | The unique identifier of the item within the Drive. | 
| drive_id | string | X | √ | The unique id of the drive. | 
| folder | json | X | √ | Folder metadata, if the item is a folder. | 
| name | string | X | √ | The name of the item (filename and extension). | 
| created_date_time | timestamp | X | √ | Date and time of item creation. | 
| created_by | json | X | √ | Identity of the user, device, and application which created the item. | 
| file | json | X | √ | File metadata, if the item is a file. | 
| description | string | X | √ | Provides a user-visible description of the item. | 


