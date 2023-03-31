# Table: microsoft365_my_calendar_event

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| original_end_time_zone | string | X | √ | The end time zone that was set when the event was created. | 
| ical_uid | string | X | √ | A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series. | 
| categories | json | X | √ | The categories associated with the event. | 
| start | json | X | √ | The start date, time, and time zone of the event. By default, the start time is in UTC. | 
| body | json | X | √ | The body of the message associated with the event. It can be in HTML or text format. | 
| id | string | X | √ | Unique identifier for the event. | 
| locations | json | X | √ | The locations where the event is held or attended from. | 
| user_id | string | X | √ | ID or email of the user. | 
| web_link | string | X | √ | The URL to open the event in Outlook on the web. | 
| original_start_time_zone | string | X | √ | The start time zone that was set when the event was created. | 
| allow_new_time_proposals | bool | X | √ | True if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Default is true. | 
| end | json | X | √ | The date, time, and time zone that the event ends. By default, the end time is in UTC. | 
| start_time | timestamp | X | √ | The start date and time of the event. By default, the start time is in UTC. | 
| show_as | string | X | √ | The status to show. Possible values are: free, tentative, busy, oof, workingElsewhere, unknown. | 
| online_meeting_provider | string | X | √ | Represents the online meeting service provider. By default, onlineMeetingProvider is unknown. The possible values are unknown, teamsForBusiness, skypeForBusiness, and skypeForConsumer. | 
| location | json | X | √ | The location of the event. | 
| attendees | json | X | √ | The collection of attendees for the event. | 
| reminder_minutes_before_start | int | X | √ | The number of minutes before the event start time that the reminder alert occurs. | 
| transaction_id | string | X | √ | A custom identifier specified by a client app for the server to avoid redundant POST operations in case of client retries to create the same event. | 
| series_master_id | string | X | √ | The ID for the recurring series master item, if this event is part of a recurring series. | 
| organizer | json | X | √ | The organizer of the event. | 
| recurrence | json | X | √ | The recurrence pattern for the event. | 
| last_modified_date_time | timestamp | X | √ | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. | 
| is_all_day | bool | X | √ | True if the event lasts all day. If true, regardless of whether it's a single-day or multi-day event, start and end time must be set to midnight and be in the same time zone. | 
| end_time | timestamp | X | √ | The end date and time of the event. By default, the end time is in UTC. | 
| change_key | string | X | √ | Identifies the version of the event object. Every time the event is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object. | 
| is_reminder_on | bool | X | √ | True if an alert is set to remind the user of the event. | 
| has_attachments | bool | X | √ | True if the event has attachments. | 
| is_draft | bool | X | √ | True if the user has updated the meeting in Outlook but has not sent the updates to attendees. | 
| hide_attendees | bool | X | √ | If set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false. | 
| subject | string | X | √ | The text of the event's subject line. | 
| title | string | X | √ | Title of the resource. | 
| response_status | json | X | √ | Indicates the type of response sent in response to an event message. | 
| is_cancelled | bool | X | √ | True  if the event has been canceled. | 
| is_organizer | bool | X | √ | True if the calendar owner (specified by the owner property of the calendar) is the organizer of the event (specified by the organizer property of the event). | 
| created_date_time | timestamp | X | √ | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. | 
| sensitivity | string | X | √ | The sensitivity of the event. Possible values are: normal, personal, private, confidential. | 
| response_requested | bool | X | √ | If true, it represents the organizer would like an invitee to send a response to the event. | 
| is_online_meeting | bool | X | √ | True if this event has online meeting information (that is, onlineMeeting points to an onlineMeetingInfo resource), false otherwise. Default is false (onlineMeeting is null). | 
| tenant_id | string | X | √ | The Azure Tenant ID where the resource is located. | 
| online_meeting_url | string | X | √ | A URL for an online meeting. The property is set only when an organizer specifies in Outlook that an event is an online meeting such as Skype. | 
| importance | int | X | √ | The importance of the event. The possible values are: low, normal, high. | 
| online_meeting | json | X | √ | Details for an attendee to join the meeting online. Default is null. | 
| body_preview | string | X | √ | The preview of the message associated with the event in text format. | 


