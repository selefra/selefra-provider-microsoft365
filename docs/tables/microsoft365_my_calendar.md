# Table: microsoft365_my_calendar

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| calendar_group_id | string | X | √ | The ID of the group. | 
| tenant_id | string | X | √ | The Azure Tenant ID where the resource is located. | 
| user_id | string | X | √ | ID or email of the user. | 
| is_tallying_responses | bool | X | √ | Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses. | 
| allowed_online_meeting_providers | json | X | √ | Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness. | 
| is_default_calendar | bool | X | √ | True if this is the default calendar where new events are created by default, false otherwise. | 
| can_view_private_items | bool | X | √ | True if the user can read calendar items that have been marked private, false otherwise. | 
| can_edit | bool | X | √ | True if the user can write to the calendar, false otherwise. | 
| hex_color | string | X | √ | The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. | 
| change_key | string | X | √ | Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. | 
| default_online_meeting_provider | string | X | √ | The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness. | 
| multi_value_extended_properties | json | X | √ | The collection of multi-value extended properties defined for the calendar. | 
| name | string | X | √ | The calendar name. | 
| owner | json | X | √ | Represents the user who created or added the calendar. | 
| permissions | json | X | √ | Represents the user who created or added the calendar. | 
| title | string | X | √ | Title of the resource. | 
| id | string | X | √ | The calendar's unique identifier. | 
| color | string | X | √ | Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor. | 
| can_share | bool | X | √ | True if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it. | 
| is_removable | bool | X | √ | Indicates whether this user calendar can be deleted from the user mailbox. | 


