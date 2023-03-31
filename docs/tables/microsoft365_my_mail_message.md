# Table: microsoft365_my_mail_message

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| attachments | json | X | √ | The attachments of the message. | 
| tenant_id | string | X | √ | The Azure Tenant ID where the resource is located. | 
| sent_date_time | timestamp | X | √ | The date and time the message was sent. | 
| internet_message_id | string | X | √ | The message ID in the format specified by RFC2822. | 
| web_link | string | X | √ |  | 
| parent_folder_id | string | X | √ | The unique identifier for the message's parent mailFolder. | 
| bcc_recipients | json | X | √ | The Bcc: recipients for the message. | 
| cc_recipients | json | X | √ | The Cc: recipients for the message. | 
| body_preview | string | X | √ | The first 255 characters of the message body in text format. | 
| received_date_time | timestamp | X | √ | The date and time the message was received. | 
| inference_classification | string | X | √ | The classification of the message for the user, based on inferred relevance or importance, or on an explicit override. The possible values are: focused or other. | 
| from | json | X | √ | The owner of the mailbox from which the message is sent. | 
| categories | json | X | √ | The categories associated with the message. | 
| change_key | string | X | √ | The version of the message. | 
| reply_to | json | X | √ | The email addresses to use when replying. | 
| to_recipients | json | X | √ | The To: recipients for the message. | 
| created_date_time | timestamp | X | √ |  | 
| sender | json | X | √ | The date and time the message was created. | 
| subject | string | X | √ | The subject of the message. | 
| conversation_id | string | X | √ | The ID of the conversation the email belongs to. | 
| is_delivery_receipt_requested | bool | X | √ | Indicates whether a read receipt is requested for the message. | 
| is_read | bool | X | √ | Indicates whether the message has been read. | 
| id | string | X | √ | Unique identifier for the message. | 
| last_modified_date_time | timestamp | X | √ |  | 
| user_id | string | X | √ | ID or email of the user. | 
| is_read_receipt_requested | bool | X | √ | Indicates whether a read receipt is requested for the message. | 
| body | json | X | √ | The body of the message. It can be in HTML or text format. | 
| title | string | X | √ | Title of the resource. | 
| has_attachments | bool | X | √ | Indicates whether the message has attachments. | 
| is_draft | bool | X | √ | Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet. | 
| importance | string | X | √ | The importance of the message. The possible values are: low, normal, and high. | 


