# Connector: Office 365 Outlook (`shared_office365`)

> **Generated** by `tools/gen_reference.py` from `reference/connectors/shared_office365.yml` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

Microsoft Office 365 is a cloud-based service that is designed to help meet your organization's needs for robust security, reliability, and user productivity.

| | |
|---|---|
| Connector id | `/providers/Microsoft.PowerApps/apis/shared_office365` |
| Type | Public |
| Operations | 88 |
| Triggers (used by flows / external triggers, not by tools) | 25 |

Use an operation's **operationId** as `action.operationId` in a `TaskDialog` tool (see [08 Tools & connectors](../../08-tools-and-connectors.md)). Inputs map to `inputs:` entries (`propertyName` = the input *Name* below). Input names that start with `$` are [OData](../../glossary.md#odata) query options and need special quoting — see [`$`-prefixed inputs](../../08-tools-and-connectors.md#-prefixed-inputs-odata-parameters).

## Operations

| operationId | Name | Description |
|---|---|---|
| [`GetEventsCalendarView`](#geteventscalendarview) | Get calendar view of events | Get calendar view of events. |
| [`GetEmails`](#getemails) | Get emails | This operation gets emails from a folder. |
| [`SendEmail`](#sendemail) | Send an email | This operation sends an email message. |
| [`GetOutlookCategoryNames`](#getoutlookcategorynames) | Get Outlook category names | This operation gets Outlook category display names. |
| [`DraftEmail`](#draftemail) | Draft an email message | This operation drafts an email message. |
| [`UpdateDraftEmail`](#updatedraftemail) | Updates an email Draft message | This operation updates an an email Draft message. |
| [`SendDraftEmail`](#senddraftemail) | Send a Draft message | This operation sends a Draft message. |
| [`AssignCategory`](#assigncategory) | Assigns an Outlook category | This operation assigns an Outlook category to an email. |
| [`AssignCategoryBulk`](#assigncategorybulk) | Assign a category to multiple emails | This operation assigns an Outlook category to multiple emails. |
| [`GetEmailsV2`](#getemailsv2) | Get emails (V2) | This operation gets emails from a folder. |
| [`SendEmailV2`](#sendemailv2) | Send an email (V2) | This operation sends an email message. |
| [`GetEmail`](#getemail) | Get email | This operation gets an email by id. |
| [`DeleteEmail`](#deleteemail) | Delete email | This operation deletes an email by id. |
| [`GetEmailV2`](#getemailv2) | Get email (V2) | This operation gets an email by id. |
| [`GetEmailsV3`](#getemailsv3) | Get emails (V3) | This operation gets emails from a folder via graph apis. Please note that filtering related to these fields: To, Cc, To Or Cc, From, Importance, Fetch Only With Attachments, Subject Filter, is performed using first 250 items in a given mail folder. To avoid that limitation you can use 'Search Query' field. |
| [`Move`](#move) | Move email | This operation moves an email to the specified folder within the same mailbox. |
| [`MoveV2`](#movev2) | Move email (V2) | This operation moves an email to the specified folder within the same mailbox. |
| [`Flag`](#flag) | Flag email | This operation flags an email. |
| [`MarkAsRead`](#markasread) | Mark as read | This operation marks an email as having been read. |
| [`ReplyTo`](#replyto) | Reply to email | This operation replies to an email. |
| [`ReplyToV2`](#replytov2) | Reply to email (V2) | This operation replies to an email. |
| [`ReplyToV3`](#replytov3) | Reply to email (V3) | This operation replies to an email. |
| [`GetAttachment`](#getattachment) | Get attachment | This operation gets an email attachment by id. |
| [`SetAutomaticRepliesSetting`](#setautomaticrepliessetting) | Set up automatic replies | Set the automatic replies setting for your mailbox. |
| [`GetMailTips`](#getmailtips) | Get mail tips for a mailbox | Get mail tips for a mailbox such as automatic replies / OOF message or if the mailbox is full. |
| [`SendMailWithOptions`](#sendmailwithoptions) | Send email with options | This operation sends an email with multiple options and waits for the recipient to respond back with one of the options. Please refer to the following link regarding the support of actionable messages in different mail clients: https://docs.microsoft.com/outlook/actionable-messages/#outlook-version-requirements-for-actionable-messages. |
| [`SendApprovalMail`](#sendapprovalmail) | Send approval email | This operation sends an approval email and waits for a response from the recipient. Please refer to the following link regarding the support of actionable messages in different mail clients: https://docs.microsoft.com/outlook/actionable-messages/#outlook-version-requirements-for-actionable-messages. |
| [`SharedMailboxSendEmail`](#sharedmailboxsendemail) | Send an email from a shared mailbox | This operation sends an email from a shared mailbox. Your account should have permission to access the mailbox for this operation to succeed. |
| [`SharedMailboxSendEmailV2`](#sharedmailboxsendemailv2) | Send an email from a shared mailbox (V2) | This operation sends an email from a shared mailbox. Your account should have permission to access the mailbox for this operation to succeed. |
| [`CalendarGetTables`](#calendargettables) | Get calendars | This operation lists available calendars. |
| [`CalendarGetItems`](#calendargetitems) | Get events (V1) | This operation gets events from a calendar. (V1) |
| [`CalendarPostItem`](#calendarpostitem) | Create event (V1) | This operation creates a new event in a calendar. (V1) |
| [`CalendarGetItem`](#calendargetitem) | Get event (V1) | This operation gets a specific event from a calendar. (V1) |
| [`CalendarDeleteItem`](#calendardeleteitem) | Delete event | This operation deletes an event in a calendar. |
| [`CalendarPatchItem`](#calendarpatchitem) | Update event (V1) | This operation updates an event in a calendar. (V1) |
| [`V3CalendarGetItems`](#v3calendargetitems) | Get events (V3) | This operation gets events from a calendar. (V3) |
| [`V3CalendarPostItem`](#v3calendarpostitem) | Create event (V3) | This operation creates a new event in a calendar. |
| [`V4CalendarGetItems`](#v4calendargetitems) | Get events (V4) | This operation gets events from a calendar using Graph API. (V4) |
| [`V4CalendarPostItem`](#v4calendarpostitem) | Create event (V4) | This operation creates a new event in a calendar. |
| [`V2CalendarGetItems`](#v2calendargetitems) | Get events (V2) | This operation gets events from a calendar. (V2) |
| [`V2CalendarPostItem`](#v2calendarpostitem) | Create event (V2) | This operation creates a new event in a calendar. (V2) |
| [`GetEventsCalendarViewV2`](#geteventscalendarviewv2) | Get calendar view of events (V2) | This operation gets all events (including instances of recurrences) in a calendar. Recurrence property is null in this case. |
| [`GetEventsCalendarViewV3`](#geteventscalendarviewv3) | Get calendar view of events (V3) | This operation gets all events (including instances of recurrences) in a calendar using Graph API. Recurrence property is null in this case. |
| [`V2CalendarGetItem`](#v2calendargetitem) | Get event (V2) | This operation gets a specific event from a calendar. (V2) |
| [`V2CalendarPatchItem`](#v2calendarpatchitem) | Update event (V2) | This operation updates an event in a calendar. (V2) |
| [`V3CalendarGetItem`](#v3calendargetitem) | Get event (V3) | This operation gets a specific event from a calendar using Graph API. (V3) |
| [`V3CalendarPatchItem`](#v3calendarpatchitem) | Update event (V3) | This operation updates an event in a calendar. |
| [`V4CalendarPatchItem`](#v4calendarpatchitem) | Update event (V4) | This operation updates an event in a calendar using Graph API. |
| [`ContactGetTables`](#contactgettables) | Get contact folders | This operation lists available contacts folders. |
| [`ContactGetTablesV2`](#contactgettablesv2) | Get contact folders (V2) | This operation lists available contacts folders using Graph API |
| [`ContactGetItems`](#contactgetitems) | Get contacts | This operation gets contacts from a contacts folder. |
| [`ContactPostItem`](#contactpostitem) | Create contact | This operation creates a new contact in a contacts folder. |
| [`ContactGetItem`](#contactgetitem) | Get contact | This operation gets a specific contact from a contacts folder. |
| [`ContactDeleteItem`](#contactdeleteitem) | Delete contact | This operation deletes a contact from a contacts folder. |
| [`ContactPatchItem`](#contactpatchitem) | Update contact | This operation updates a contact in a contacts folder. |
| [`ExportEmail`](#exportemail) | Export email | Export the content of the email. |
| [`ExportEmail_V2`](#exportemail_v2) | Export email (V2) | Export the content of the email in the EML file format. |
| [`Flag_V2`](#flag_v2) | Flag email (V2) | This operation updates an email flag. |
| [`MarkAsRead_V2`](#markasread_v2) | Mark as read or unread (V2) | This operation marks an email as read/unread. |
| [`MarkAsRead_V3`](#markasread_v3) | Mark as read or unread (V3) | This operation marks an email as read/unread. |
| [`DeleteEmail_V2`](#deleteemail_v2) | Delete email (V2) | This operation deletes an email by id. |
| [`GetAttachment_V2`](#getattachment_v2) | Get Attachment (V2) | This operation gets an email attachment by id. |
| [`RespondToEvent`](#respondtoevent) | Respond to an event invite | Respond to an event invite. |
| [`RespondToEvent_V2`](#respondtoevent_v2) | Respond to an event invite (V2) | Respond to an event invite. |
| [`ForwardEmail`](#forwardemail) | Forward an email | Forward an email. |
| [`ForwardEmail_V2`](#forwardemail_v2) | Forward an email (V2) | Forward an email. |
| [`GetRoomLists`](#getroomlists) | Get room lists | Get all the room lists defined in the user's tenant |
| [`GetRoomLists_V2`](#getroomlists_v2) | Get room lists (V2) | Get all the room lists defined in the user's tenant |
| [`GetRooms`](#getrooms) | Get rooms | Get all the meeting rooms defined in the user's tenant |
| [`GetRooms_V2`](#getrooms_v2) | Get rooms (V2) | Get all the meeting rooms defined in the user's tenant |
| [`GetRoomsInRoomList`](#getroomsinroomlist) | Get rooms in room list | Get the meeting rooms in a specific room list |
| [`GetRoomsInRoomList_V2`](#getroomsinroomlist_v2) | Get rooms in room list (V2) | Get the meeting rooms in a specific room list |
| [`FindMeetingTimes`](#findmeetingtimes) | Find meeting times | Find meeting time suggestions based on organizer, attendee availability, and time or location constraints |
| [`FindMeetingTimes_V2`](#findmeetingtimes_v2) | Find meeting times (V2) | Find meeting time suggestions based on organizer, attendee availability, and time or location constraints |
| [`SetAutomaticRepliesSetting_V2`](#setautomaticrepliessetting_v2) | Set up automatic replies (V2) | Set the automatic replies setting for your mailbox. |
| [`GetMailTips_V2`](#getmailtips_v2) | Get mail tips for a mailbox (V2) | Get mail tips for a mailbox such as automatic replies / OOF message or if the mailbox is full. This is not available in GccHigh and Mooncake. |
| [`CalendarGetTables_V2`](#calendargettables_v2) | Get calendars (V2) | This operation lists available calendars. |
| [`CalendarDeleteItem_V2`](#calendardeleteitem_v2) | Delete event (V2) | This operation deletes an event in a calendar. |
| [`ContactGetItem_V2`](#contactgetitem_v2) | Get contact (V2) | This operation gets a specific contact from a contacts folder. |
| [`ContactDeleteItem_V2`](#contactdeleteitem_v2) | Delete contact (V2) | This operation deletes a contact from a contacts folder. |
| [`ContactPatchItem_V2`](#contactpatchitem_v2) | Update contact (V2) | This operation updates a contact in a contacts folder. |
| [`ContactGetItems_V2`](#contactgetitems_v2) | Get contacts (V2) | This operation gets contacts from a contacts folder. |
| [`ContactPostItem_V2`](#contactpostitem_v2) | Create contact (V2) | This operation creates a new contact in a contacts folder. |
| [`UpdateMyContactPhoto`](#updatemycontactphoto) | Update my contact's photo | Updates the photo of the specified contact of the current user. The size of the photo must be less than 4 MB. |
| [`HttpRequest`](#httprequest) | Send an HTTP request | Construct a Microsoft Graph REST API request to invoke. These segments are supported: 1st segement: /me, /users/<userId> 2nd segment: messages, mailFolders, events, calendar, calendars, outlook, inferenceClassification. Learn more: https://docs.microsoft.com/en-us/graph/use-the-api. |
| [`mcp_EmailsManagement`](#mcp_emailsmanagement) | Email Management MCP Server | This MCP server manages email messages from your Office 365 account |
| [`mcp_MeetingManagement`](#mcp_meetingmanagement) | Meeting Management MCP Server | This MCP server manages events, calendars and meetings |
| [`mcp_ContactsManagement`](#mcp_contactsmanagement) | Contact Management MCP Server | This MCP server manages contacts |

### GetEventsCalendarView

**Get calendar view of events** — Get calendar view of events.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `calendarId` | Calendar Id | `String` | ✓ | Select a calendar. *values looked up dynamically* |
| `startDateTimeOffset` | Start Time | `String` | ✓ | Start time (example: '2017-01-01T08:00:00-07:00'). |
| `endDateTimeOffset` | End Time | `String` | ✓ | End time (example: '2017-02-01T08:00:00-07:00'). |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Values` |  | `Table` (28 fields) | Values |

### GetEmails

**Get emails** — This operation gets emails from a folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `folderPath` | Folder | `String` |  | Mail folder to retrieve emails from (default: 'Inbox'). |
| `fetchOnlyUnread` | Fetch Only Unread Messages | `Boolean` |  | Retrieve only unread emails?. |
| `includeAttachments` | Include Attachments | `Boolean` |  | If set to true, attachments content will also be retrieved along with the email. |
| `searchQuery` | Search Query | `String` |  | Search query (like in the Outlook client) to filter emails. |
| `top` | Top | `Number` |  | Number of emails to retrieve (default: 10). |

**Output** — `Table` (17 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Importance` | Importance | `String` | The importance of the message |
| `From` | From | `String` | The mailbox owner and sender of the message |
| `To` | To | `String` | The recipients for the message |
| `Cc` | CC | `String` | The Cc recipients for the message *advanced* |
| `Bcc` | BCC | `String` | The Bcc recipients for the message *advanced* |
| `ReplyTo` | Reply To | `String` | The email addresses to use when replying *advanced* |
| `Subject` | Subject | `String` | The subject of the message |
| `Body` | Body | `String` | The body of the message |
| `BodyPreview` | Body Preview | `String` | The preview of the message *advanced* |
| `HasAttachment` | Has Attachment | `Boolean` | Indicates whether the message has attachments |
| `Id` | Message Id | `String` | The unique identifier of the message *advanced* |
| `InternetMessageId` | Internet Message Id | `String` | The message ID in the format specified by RFC2822 *advanced* |
| `ConversationId` | Conversation Id | `String` | The Id of the conversation the email belongs to *advanced* |
| `DateTimeReceived` | Received Time | `DateTime` | The date and time the message was received *advanced* |
| `IsRead` | Is Read | `Boolean` | Indicates whether the message has been read *advanced* |
| `Attachments` | Attachments | `Table` (8 fields) | The file attachments for the message *advanced* |
| `IsHtml` | Is HTML | `Boolean` | Is Html? *advanced* |

### SendEmail

**Send an email** — This operation sends an email message.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `To` | To | `String` | ✓ | Specify email addresses separated by semicolons like someone@contoso.com |
| `Subject` | Subject | `String` | ✓ | Specify the subject of the mail |
| `Body` | Body | `String` | ✓ | Specify the body of the mail |
| `From` | From (Send as) | `String` |  | Email address to send mail from (requires "Send as" or "Send on behalf of" permission for that mailbox). For more info on granting permissions please refer https://docs.microsoft.com/office365/admin/manage/send-email-as-distribution-list *advanced* |
| `Cc` | CC | `String` |  | Specify email addresses separated by semicolons like someone@contoso.com *advanced* |
| `Bcc` | BCC | `String` |  | Specify email addresses separated by semicolons like someone@contoso.com *advanced* |
| `Attachments` | Attachments | `Table` (2 fields) |  | Attachments *advanced* |
| `Sensitivity` | Sensitivity | `String` |  | Sensitivity *advanced* *values looked up dynamically* |
| `ReplyTo` | Reply To | `String` |  | The email addresses to use when replying *advanced* |
| `Importance` | Importance | `String` |  | Importance *advanced* |
| `IsHtml` | Is HTML | `Boolean` |  | Is Html? *advanced* |

**Output** — `String`

### GetOutlookCategoryNames

**Get Outlook category names** — This operation gets Outlook category display names.

**Inputs**

_None._

**Output** — `Table` (2 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `id` |  | `String` | The ID of the Outlook Category. |
| `displayName` |  | `String` | The display name of the Outlook Category. |

### DraftEmail

**Draft an email message** — This operation drafts an email message.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `To` | To | `String` | ✓ | Specify email addresses separated by semicolons like someone@contoso.com |
| `Subject` | Subject | `String` | ✓ | Specify the subject of the mail |
| `Body` | Body | `String` | ✓ | Specify the body of the mail |
| `messageId` | Message Id | `String` |  | Message Id. *advanced* |
| `draftType` | Draft Type | `String` |  | Draft Type. *advanced* |
| `comment` | Comment | `String` |  | Draft Comment *advanced* |
| `From` | From (Send as) | `String` |  | Email address to send mail from (requires "Send as" or "Send on behalf of" permission for that mailbox). For more info on granting permissions please refer https://docs.microsoft.com/office365/admin/manage/send-email-as-distribution-list *advanced* |
| `Cc` | CC | `String` |  | Specify email addresses separated by semicolons like someone@contoso.com *advanced* |
| `Bcc` | BCC | `String` |  | Specify email addresses separated by semicolons like someone@contoso.com *advanced* |
| `Attachments` | Attachments | `Table` (2 fields) |  | Attachments *advanced* |
| `Sensitivity` | Sensitivity | `String` |  | Sensitivity *advanced* *values looked up dynamically* |
| `ReplyTo` | Reply To | `String` |  | The email addresses to use when replying *advanced* |
| `Importance` | Importance | `String` |  | Importance *advanced* |

**Output** — `Record` (19 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `InternetMessageId` |  | `String` | Internet Message Id |
| `BodyPreview` |  | `String` | Body preview |
| `Id` |  | `String` | Id |
| `ConversationId` |  | `String` | Conversation Id |
| `HasAttachments` |  | `Boolean` | Has attachments |
| `IsRead` |  | `Boolean` | Is read |
| `CreatedDateTime` |  | `DateTime` | Created date and time |
| `ReceivedDateTime` |  | `DateTime` | Received date and time |
| `LastModifiedDateTime` |  | `DateTime` | Last modified date and time |
| `Attachments` |  | `Table` (12 fields) | Attachments |
| `ToRecipients` |  | `Table` (1 fields) | To Recipient |
| `CcRecipients` |  | `Table` (1 fields) | Cc Recipients |
| `BccRecipients` |  | `Table` (1 fields) | Bcc Recipients |
| `ReplyTo` |  | `Table` (1 fields) | The email addresses to use when replying |
| `Subject` |  | `String` | Subject |
| `Body` |  | `Record` (2 fields) |  |
| `From` |  | `Record` (1 fields) |  |
| `Importance` |  | `String` | Importance |
| `InternetMessageHeaders` |  | `Table` (2 fields) | Internet message headers |

### UpdateDraftEmail

**Updates an email Draft message** — This operation updates an an email Draft message.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Message Id. |
| `To` | To | `String` | ✓ | Specify email addresses separated by semicolons like someone@contoso.com |
| `Subject` | Subject | `String` | ✓ | Specify the subject of the mail |
| `Body` | Body | `String` | ✓ | Specify the body of the mail |
| `From` | From (Send as) | `String` |  | Email address to send mail from (requires "Send as" or "Send on behalf of" permission for that mailbox). For more info on granting permissions please refer https://docs.microsoft.com/office365/admin/manage/send-email-as-distribution-list *advanced* |
| `Cc` | CC | `String` |  | Specify email addresses separated by semicolons like someone@contoso.com *advanced* |
| `Bcc` | BCC | `String` |  | Specify email addresses separated by semicolons like someone@contoso.com *advanced* |
| `Attachments` | Attachments | `Table` (2 fields) |  | Attachments *advanced* |
| `Sensitivity` | Sensitivity | `String` |  | Sensitivity *advanced* *values looked up dynamically* |
| `ReplyTo` | Reply To | `String` |  | The email addresses to use when replying *advanced* |
| `Importance` | Importance | `String` |  | Importance *advanced* |

**Output** — `String`

### SendDraftEmail

**Send a Draft message** — This operation sends a Draft message.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Message Id. |

**Output** — `String`

### AssignCategory

**Assigns an Outlook category** — This operation assigns an Outlook category to an email.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Message Id. |
| `category` | Category | `String` | ✓ | Category |

**Output** — `String`

### AssignCategoryBulk

**Assign a category to multiple emails** — This operation assigns an Outlook category to multiple emails.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `categoryName` | Category Name | `String` | ✓ | Category Name |
| `messageIds` | List of Message IDs | `Table` (1 fields) | ✓ | List of Message IDs. |

**Output** — `Record` (2 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `successCount` |  | `Number` | The count of messages processed successfully. |
| `failures` |  | `Table` (2 fields) | The list of failed messages with errors. |

### GetEmailsV2

**Get emails (V2)** — This operation gets emails from a folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `folderPath` | Folder | `String` |  | Mail folder to retrieve emails from (default: 'Inbox'). |
| `to` | To | `String` |  | Recipient email addresses separated by semicolons (If any match, the trigger will run). |
| `cc` | CC | `String` |  | CC recipient email addresses separated by semicolons (If any match, the trigger will run). *advanced* |
| `toOrCc` | To or CC | `String` |  | To or CC recipient email addresses separated by semicolons (If any match, the trigger will run). *advanced* |
| `from` | From | `String` |  | Sender email addresses separated by semicolons (If any match, the trigger will run). |
| `importance` | Importance | `String` |  | Importance of the email (Any, High, Normal, Low). *advanced* |
| `fetchOnlyWithAttachment` | Only with Attachments | `Boolean` |  | If set to true, only emails with an attachment will be retrieved. Emails without any attachments will be skipped. If set to false, all emails will be retrieved. *advanced* |
| `subjectFilter` | Subject Filter | `String` |  | String to look for in the subject line. *advanced* |
| `fetchOnlyUnread` | Fetch Only Unread Messages | `Boolean` |  | Retrieve only unread emails?. |
| `includeAttachments` | Include Attachments | `Boolean` |  | If set to true, attachments content will also be retrieved along with the email. |
| `searchQuery` | Search Query | `String` |  | Search query (like in the Outlook client) to filter emails. |
| `top` | Top | `Number` |  | Number of emails to retrieve (default: 10). |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (17 fields) | A list of the response objects |

### SendEmailV2

**Send an email (V2)** — This operation sends an email message.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `To` | To | `String` | ✓ | Specify email addresses separated by semicolons like someone@contoso.com |
| `Subject` | Subject | `String` | ✓ | Specify the subject of the mail |
| `Body` | Body | `String` | ✓ | Specify the body of the mail |
| `From` | From (Send as) | `String` |  | Email address to send mail from (requires "Send as" or "Send on behalf of" permission for that mailbox). For more info on granting permissions please refer https://docs.microsoft.com/office365/admin/manage/send-email-as-distribution-list *advanced* |
| `Cc` | CC | `String` |  | Specify email addresses separated by semicolons like someone@contoso.com *advanced* |
| `Bcc` | BCC | `String` |  | Specify email addresses separated by semicolons like someone@contoso.com *advanced* |
| `Attachments` | Attachments | `Table` (2 fields) |  | Attachments *advanced* |
| `Sensitivity` | Sensitivity | `String` |  | Sensitivity *values looked up dynamically* |
| `ReplyTo` | Reply To | `String` |  | The email addresses to use when replying *advanced* |
| `Importance` | Importance | `String` |  | Importance *advanced* |

**Output** — `String`

### GetEmail

**Get email** — This operation gets an email by id.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email. |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to retrieve mail from. |
| `includeAttachments` | Include Attachments | `Boolean` |  | If set to true, attachments content will also be retrieved along with the email. |
| `internetMessageId` | Internet Message Id | `String` |  | Internet Message Id. *advanced* |

**Output** — `Record` (17 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `From` | From | `String` | The mailbox owner and sender of the message |
| `To` | To | `String` | The recipients for the message |
| `Cc` | CC | `String` | The Cc recipients for the message *advanced* |
| `Bcc` | BCC | `String` | The Bcc recipients for the message *advanced* |
| `ReplyTo` | Reply To | `String` | The email addresses to use when replying *advanced* |
| `Subject` | Subject | `String` | The subject of the message |
| `Body` | Body | `String` | The body of the message |
| `Importance` | Importance | `Number` | The importance of the message (0 - Low, 1 - Normal, 2 - High) |
| `BodyPreview` | Body Preview | `String` | The preview of the message *advanced* |
| `HasAttachment` | Has Attachment | `Boolean` | Indicates whether the message has attachments |
| `Id` | Message Id | `String` | The unique identifier of the message *advanced* |
| `InternetMessageId` | Internet Message Id | `String` | The message ID in the format specified by RFC2822 *advanced* |
| `ConversationId` | Conversation Id | `String` | The Id of the conversation the email belongs to *advanced* |
| `DateTimeReceived` | Received Time | `DateTime` | The date and time the message was received *advanced* |
| `IsRead` | Is Read | `Boolean` | Indicates whether the message has been read *advanced* |
| `Attachments` | Attachments | `Table` (8 fields) | The file attachments for the message *advanced* |
| `IsHtml` | Is HTML | `Boolean` | Is Html? *advanced* |

### DeleteEmail

**Delete email** — This operation deletes an email by id.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to delete. |

**Output** — `String`

### GetEmailV2

**Get email (V2)** — This operation gets an email by id.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email. |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to retrieve mail from. |
| `includeAttachments` | Include Attachments | `Boolean` |  | If set to true, attachments content will also be retrieved along with the email. |
| `internetMessageId` | Internet Message Id | `String` |  | Internet Message Id. *advanced* |
| `extractSensitivityLabel` | Extract Sensitivity Label | `Boolean` |  | Select if you want to extract Sensitivity label ( false, true). |
| `fetchSensitivityLabelMetadata` | Sensitivity Label Metadata | `Boolean` |  | A boolean whether to fetch sensitivity label Metadata for associated LabelId. |

**Output** — `Record` (18 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `from` | From | `String` | The mailbox owner and sender of the message |
| `toRecipients` | To | `String` | The recipients for the message |
| `ccRecipients` | CC | `String` | The Cc recipients for the message *advanced* |
| `bccRecipients` | BCC | `String` | The Bcc recipients for the message *advanced* |
| `replyTo` | Reply To | `String` | The email addresses to use when replying *advanced* |
| `subject` | Subject | `String` | The subject of the message |
| `body` | Body | `String` | The body of the message |
| `importance` | Importance | `String` | The importance of the message (low, normal, high) |
| `bodyPreview` | Body Preview | `String` | The preview of the message *advanced* |
| `hasAttachments` | Has Attachment | `Boolean` | Indicates whether the message has attachments |
| `id` | Message Id | `String` | The unique identifier of the message *advanced* |
| `internetMessageId` | Internet Message Id | `String` | The message ID in the format specified by RFC2822 *advanced* |
| `conversationId` | Conversation Id | `String` | The Id of the conversation the email belongs to *advanced* |
| `receivedDateTime` | Received Time | `DateTime` | The date and time the message was received *advanced* |
| `isRead` | Is Read | `Boolean` | Indicates whether the message has been read *advanced* |
| `attachments` | Attachments | `Table` (8 fields) | The file attachments for the message *advanced* |
| `isHtml` | Is HTML | `Boolean` | Is Html? *advanced* |
| `sensitivityLabelInfo` |  | `Table` (10 fields) |  |

### GetEmailsV3

**Get emails (V3)** — This operation gets emails from a folder via graph apis. Please note that filtering related to these fields: To, Cc, To Or Cc, From, Importance, Fetch Only With Attachments, Subject Filter, is performed using first 250 items in a given mail folder. To avoid that limitation you can use 'Search Query' field.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `folderPath` | Folder | `String` |  | Mail folder to retrieve emails from (default: 'Inbox'). |
| `to` | To | `String` |  | Recipient email addresses separated by semicolons (If any match, the trigger will run). |
| `cc` | CC | `String` |  | CC recipient email addresses separated by semicolons (If any match, the trigger will run). *advanced* |
| `toOrCc` | To or CC | `String` |  | To or CC recipient email addresses separated by semicolons (If any match, the trigger will run). *advanced* |
| `from` | From | `String` |  | Sender email addresses separated by semicolons (If any match, the trigger will run). |
| `importance` | Importance | `String` |  | Importance of the email (Any, High, Normal, Low). *advanced* |
| `fetchOnlyWithAttachment` | Only with Attachments | `Boolean` |  | If set to true, only emails with an attachment will be retrieved. Emails without any attachments will be skipped. If set to false, all emails will be retrieved. *advanced* |
| `subjectFilter` | Subject Filter | `String` |  | String to look for in the subject line. *advanced* |
| `fetchOnlyUnread` | Fetch Only Unread Messages | `Boolean` |  | Retrieve only unread emails?. |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to retrieve mails from. |
| `includeAttachments` | Include Attachments | `Boolean` |  | If set to true, attachments content will also be retrieved along with the email. |
| `searchQuery` | Search Query | `String` |  | Search query to filter emails. How to use '$search' parameter please refer to: https://docs.microsoft.com/graph/query-parameters#search-parameter. |
| `top` | Top | `Number` |  | Number of emails to retrieve (default: 10, max: 1000). |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (18 fields) | A list of the response objects |

### Move

**Move email** — This operation moves an email to the specified folder within the same mailbox.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to be moved. |
| `folderPath` | Folder | `String` | ✓ | Mail folder to move the email to. |

**Output** — `Record` (17 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Importance` | Importance | `String` | The importance of the message |
| `From` | From | `String` | The mailbox owner and sender of the message |
| `To` | To | `String` | The recipients for the message |
| `Cc` | CC | `String` | The Cc recipients for the message *advanced* |
| `Bcc` | BCC | `String` | The Bcc recipients for the message *advanced* |
| `ReplyTo` | Reply To | `String` | The email addresses to use when replying *advanced* |
| `Subject` | Subject | `String` | The subject of the message |
| `Body` | Body | `String` | The body of the message |
| `BodyPreview` | Body Preview | `String` | The preview of the message *advanced* |
| `HasAttachment` | Has Attachment | `Boolean` | Indicates whether the message has attachments |
| `Id` | Message Id | `String` | The unique identifier of the message *advanced* |
| `InternetMessageId` | Internet Message Id | `String` | The message ID in the format specified by RFC2822 *advanced* |
| `ConversationId` | Conversation Id | `String` | The Id of the conversation the email belongs to *advanced* |
| `DateTimeReceived` | Received Time | `DateTime` | The date and time the message was received *advanced* |
| `IsRead` | Is Read | `Boolean` | Indicates whether the message has been read *advanced* |
| `Attachments` | Attachments | `Table` (8 fields) | The file attachments for the message *advanced* |
| `IsHtml` | Is HTML | `Boolean` | Is Html? *advanced* |

### MoveV2

**Move email (V2)** — This operation moves an email to the specified folder within the same mailbox.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to be moved. |
| `folderPath` | Folder | `String` | ✓ | Mail folder to move the email to. |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to move mail from. |

**Output** — `Record` (18 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `from` | From | `String` | The mailbox owner and sender of the message |
| `toRecipients` | To | `String` | The recipients for the message |
| `ccRecipients` | CC | `String` | The Cc recipients for the message *advanced* |
| `bccRecipients` | BCC | `String` | The Bcc recipients for the message *advanced* |
| `replyTo` | Reply To | `String` | The email addresses to use when replying *advanced* |
| `subject` | Subject | `String` | The subject of the message |
| `body` | Body | `String` | The body of the message |
| `importance` | Importance | `String` | The importance of the message (low, normal, high) |
| `bodyPreview` | Body Preview | `String` | The preview of the message *advanced* |
| `hasAttachments` | Has Attachment | `Boolean` | Indicates whether the message has attachments |
| `id` | Message Id | `String` | The unique identifier of the message *advanced* |
| `internetMessageId` | Internet Message Id | `String` | The message ID in the format specified by RFC2822 *advanced* |
| `conversationId` | Conversation Id | `String` | The Id of the conversation the email belongs to *advanced* |
| `receivedDateTime` | Received Time | `DateTime` | The date and time the message was received *advanced* |
| `isRead` | Is Read | `Boolean` | Indicates whether the message has been read *advanced* |
| `attachments` | Attachments | `Table` (8 fields) | The file attachments for the message *advanced* |
| `isHtml` | Is HTML | `Boolean` | Is Html? *advanced* |
| `sensitivityLabelInfo` |  | `Table` (10 fields) |  |

### Flag

**Flag email** — This operation flags an email.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to be flagged. |

**Output** — `String`

### MarkAsRead

**Mark as read** — This operation marks an email as having been read.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to be marked as read. |

**Output** — `String`

### ReplyTo

**Reply to email** — This operation replies to an email.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to reply to. |
| `comment` | Comment | `String` | ✓ | Reply comment. |
| `replyAll` | Reply All | `Boolean` |  | Reply to all recipients. |

**Output** — `String`

### ReplyToV2

**Reply to email (V2)** — This operation replies to an email.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to reply to. |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to reply from. |
| `To` | To | `String` |  | Example: recipient1@domain.com; recipient2@domain.com *advanced* |
| `Cc` | CC | `String` |  | Example: recipient1@domain.com; recipient2@domain.com *advanced* |
| `Bcc` | BCC | `String` |  | Example: recipient1@domain.com; recipient2@domain.com *advanced* |
| `Subject` | Subject | `String` |  | Email subject (if empty, the original subject used). *advanced* |
| `Body` | Body | `String` |  | Content of the email. |
| `ReplyAll` | Reply All | `Boolean` |  | True to reply to all recipients. (default: False) |
| `IsHtml` | Is HTML | `Boolean` |  | True to send the reply as HTML. (default: True) *advanced* |
| `Importance` | Importance | `String` |  | Pick an importance. (default: Low) *advanced* |
| `Attachments` | Attachments | `Table` (2 fields) |  | Details of attachments to be sent along with the reply. |

**Output** — `String`

### ReplyToV3

**Reply to email (V3)** — This operation replies to an email.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to reply to. |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to reply from. |
| `To` | To | `String` |  | Example: recipient1@domain.com; recipient2@domain.com *advanced* |
| `Cc` | CC | `String` |  | Example: recipient1@domain.com; recipient2@domain.com *advanced* |
| `Bcc` | BCC | `String` |  | Example: recipient1@domain.com; recipient2@domain.com *advanced* |
| `Subject` | Subject | `String` |  | Email subject (if empty, the original subject used). *advanced* |
| `Body` | Body | `String` |  | Content of the email. |
| `ReplyAll` | Reply All | `Boolean` |  | True to reply to all recipients. (default: False) |
| `Importance` | Importance | `String` |  | Pick an importance. (default: Low) *advanced* |
| `Attachments` | Attachments | `Table` (2 fields) |  | Details of attachments to be sent along with the reply. |

**Output** — `String`

### GetAttachment

**Get attachment** — This operation gets an email attachment by id.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email. |
| `attachmentId` | Attachment Id | `String` | ✓ | Id of the attachment to download. |

**Output** — `File`

### SetAutomaticRepliesSetting

**Set up automatic replies** — Set the automatic replies setting for your mailbox.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `Status` | Status | `String` | ✓ | Automatic reply setting status |
| `ExternalAudience` | External Audience | `String` | ✓ | The audience that will see the external reply message |
| `ScheduledStartDateTimeOffset` | Start Time | `String` |  | Scheduled start time (example: '2016-11-01T15:30:00-00:00Z') |
| `ScheduledEndDateTimeOffset` | End Time | `String` |  | Scheduled end time (example: '2016-11-01T15:30:00-00:00Z') |
| `InternalReplyMessage` | Internal Reply Message | `String` |  | Message for people within your organization |
| `ExternalReplyMessage` | External Reply Message | `String` |  | Message for people outside your organization |

**Output** — `String`

### GetMailTips

**Get mail tips for a mailbox** — Get mail tips for a mailbox such as automatic replies / OOF message or if the mailbox is full.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `mailboxAddress` | Original Mailbox Address | `String` | ✓ | Address of the mailbox to get mail tips for. |

**Output** — `Record` (7 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `AutomaticReplies` |  | `Record` (1 fields) | Automatic replies as part of mail tips |
| `DeliveryRestricted` |  | `Boolean` | Is delivery restricted |
| `ExternalMemberCount` |  | `Number` | Number of external members |
| `IsModerated` |  | `Boolean` | Is moderated |
| `MailboxFull` |  | `Boolean` | Is mailbox full |
| `MaxMessageSize` |  | `Number` | Maximum message size |
| `TotalMemberCount` |  | `Number` | Total member count |

### SendMailWithOptions

**Send email with options** — This operation sends an email with multiple options and waits for the recipient to respond back with one of the options. Please refer to the following link regarding the support of actionable messages in different mail clients: https://docs.microsoft.com/outlook/actionable-messages/#outlook-version-requirements-for-actionable-messages.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `Message` | Subscription request for options email | `Record` (12 fields) | ✓ | User Options Email Message. This is the message expected as part of user input |

**Output** — `Record` (4 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `id` |  | `String` | Id of the subscription |
| `resource` |  | `String` | Resource of the subscription request |
| `notificationType` |  | `String` | Notification Type |
| `notificationUrl` |  | `String` | Notification Url |

### SendApprovalMail

**Send approval email** — This operation sends an approval email and waits for a response from the recipient. Please refer to the following link regarding the support of actionable messages in different mail clients: https://docs.microsoft.com/outlook/actionable-messages/#outlook-version-requirements-for-actionable-messages.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `Message` | Subscription request for approval email | `Record` (11 fields) | ✓ | Approval Email Message. This is the message expected as part of user input |

**Output** — `Record` (4 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `id` |  | `String` | Id of the subscription |
| `resource` |  | `String` | Resource of the subscription request |
| `notificationType` |  | `String` | Notification Type |
| `notificationUrl` |  | `String` | Notification Url |

### SharedMailboxSendEmail

**Send an email from a shared mailbox** — This operation sends an email from a shared mailbox. Your account should have permission to access the mailbox for this operation to succeed.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `MailboxAddress` | Original Mailbox Address | `String` | ✓ | Specify email address of a shared mailbox like someone@contoso.com |
| `To` | To | `String` | ✓ | Specify email addresses separated by semicolons like someone@contoso.com |
| `Subject` | Subject | `String` | ✓ | Specify the subject of the mail |
| `Body` | Body | `String` | ✓ | Specify the body of the mail |
| `Cc` | CC | `String` |  | Specify email addresses separated by semicolons like someone@contoso.com *advanced* |
| `Bcc` | BCC | `String` |  | Specify email addresses separated by semicolons like someone@contoso.com *advanced* |
| `Attachments` | Attachments | `Table` (2 fields) |  | Attachments *advanced* |
| `Sensitivity` | Sensitivity | `String` |  | Sensitivity *advanced* *values looked up dynamically* |
| `ReplyTo` | Reply To | `String` |  | The email addresses to use when replying *advanced* |
| `Importance` | Importance | `String` |  | Importance *advanced* |
| `IsHtml` | Is HTML | `Boolean` |  | Is Html? *advanced* |

**Output** — `String`

### SharedMailboxSendEmailV2

**Send an email from a shared mailbox (V2)** — This operation sends an email from a shared mailbox. Your account should have permission to access the mailbox for this operation to succeed.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `MailboxAddress` | Original Mailbox Address | `String` | ✓ | Specify email address of a shared mailbox like someone@contoso.com |
| `To` | To | `String` | ✓ | Specify email addresses separated by semicolons like someone@contoso.com |
| `Subject` | Subject | `String` | ✓ | Specify the subject of the mail |
| `Body` | Body | `String` | ✓ | Specify the body of the mail |
| `Cc` | CC | `String` |  | Specify email addresses separated by semicolons like someone@contoso.com *advanced* |
| `Bcc` | BCC | `String` |  | Specify email addresses separated by semicolons like someone@contoso.com *advanced* |
| `Attachments` | Attachments | `Table` (2 fields) |  | Attachments *advanced* |
| `Sensitivity` | Sensitivity | `String` |  | Sensitivity *advanced* *values looked up dynamically* |
| `ReplyTo` | Reply To | `String` |  | The email addresses to use when replying *advanced* |
| `Importance` | Importance | `String` |  | Importance *advanced* |

**Output** — `String`

### CalendarGetTables

**Get calendars** — This operation lists available calendars.

**Inputs**

_None._

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (3 fields) | List of values |

### CalendarGetItems

**Get events (V1)** — This operation gets events from a calendar. (V1)

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `$orderby` | Order By | `String` |  | An ODATA orderBy query for specifying the order of entries. *advanced* |
| `$top` | Top Count | `Number` |  | Total number of entries to retrieve (default = all). *advanced* |
| `$skip` | Skip Count | `Number` |  | The number of entries to skip (default = 0). *advanced* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (30 fields) | List of values |

### CalendarPostItem

**Create event (V1)** — This operation creates a new event in a calendar. (V1)

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `End` | End time | `DateTime` | ✓ | The end time of the event |
| `Start` | Start time | `DateTime` | ✓ | The start time of the event |
| `Subject` | Subject | `String` | ✓ | Event subject |
| `Id` | Id | `String` |  | The Event's unique identifier |
| `Attendees` | Attendees | `Table` (3 fields) |  | List of attendees for the event |
| `Body` | Item | `Record` (2 fields) |  | Event to create |
| `BodyPreview` | Body preview | `String` |  | The preview of the message associated with the event *advanced* |
| `Categories` | Categories | `Table` (1 fields) |  | The categories associated with the event *advanced* |
| `ChangeKey` | Change key | `String` |  | This property identifies the version of the event object. Every time the event is changed, ChangeKey changes as well. *advanced* |
| `DateTimeCreated` | Created time | `DateTime` |  | The date and time that the event was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` |  | The date and time that the event was last modified *advanced* |
| `EndTimeZone` | End time zone | `String` |  | This property specifies the time zone of the meeting end time. The value must be as defined in Windows (example: 'Pacific Standard Time'). *advanced* |
| `HasAttachments` | Has attachments? | `Boolean` |  | Set to true if the event has attachments *advanced* |
| `ICalUId` | Event Unique ID | `String` |  | A unique identifier that is shared by all instances of an event across different calendars *advanced* |
| `Importance` | Importance | `String` |  | The importance of the event: Low, Normal, or High *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` |  | Set to true if the event lasts all day *advanced* |
| `IsCancelled` | Is cancelled? | `Boolean` |  | Set to true if the event has been canceled *advanced* |
| `IsOrganizer` | Is Organizer | `Boolean` |  | Set to true if the message sender is also the organizer *advanced* |
| `Location` | Item | `Record` (3 fields) |  | Event to create |
| `Organizer` | Item | `Record` (1 fields) |  | Event to create |
| `Recurrence` | Item | `Record` (2 fields) |  | Event to create |
| `Reminder` | Reminder | `Number` |  | Time in minutes before event start to remind *advanced* |
| `ResponseRequested` | Response requested | `Boolean` |  | Set to true if the sender would like a response when the event is accepted or declined *advanced* |
| `ResponseStatus` | Item | `Record` (2 fields) |  | Event to create |
| `SeriesMasterId` | Series master id | `String` |  | Unique identifier for Series Master event type *advanced* |
| `ShowAs` | Show as | `String` |  | Shows as free or busy *advanced* |
| `StartTimeZone` | Start time zone | `String` |  | This property specifies the time zone of the meeting start time. The value must be as defined in Windows (example: 'Pacific Standard Time'). *advanced* |
| `Type` | Type | `String` |  | The event type: Single Instance, Occurrence, Exception, or Series Master *advanced* |
| `WebLink` | Web link | `String` |  | The preview of the message associated with the event *advanced* |
| `Reason` | Reason | `String` |  | The reason property used by O365 sync events protocol, will be 'deleted' if its a deleted event. *advanced* |

**Output** — `Record` (30 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Id` | Id | `String` | The Event's unique identifier |
| `Attendees` | Attendees | `Table` (3 fields) | List of attendees for the event |
| `Body` |  | `Record` (2 fields) |  |
| `BodyPreview` | Body preview | `String` | The preview of the message associated with the event *advanced* |
| `Categories` | Categories | `Table` (1 fields) | The categories associated with the event *advanced* |
| `ChangeKey` | Change key | `String` | This property identifies the version of the event object. Every time the event is changed, ChangeKey changes as well. *advanced* |
| `DateTimeCreated` | Created time | `DateTime` | The date and time that the event was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` | The date and time that the event was last modified *advanced* |
| `End` | End time | `DateTime` | The end time of the event |
| `EndTimeZone` | End time zone | `String` | This property specifies the time zone of the meeting end time. The value must be as defined in Windows (example: 'Pacific Standard Time'). *advanced* |
| `HasAttachments` | Has attachments? | `Boolean` | Set to true if the event has attachments *advanced* |
| `ICalUId` | Event Unique ID | `String` | A unique identifier that is shared by all instances of an event across different calendars *advanced* |
| `Importance` | Importance | `String` | The importance of the event: Low, Normal, or High *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` | Set to true if the event lasts all day *advanced* |
| `IsCancelled` | Is cancelled? | `Boolean` | Set to true if the event has been canceled *advanced* |
| `IsOrganizer` | Is Organizer | `Boolean` | Set to true if the message sender is also the organizer *advanced* |
| `Location` |  | `Record` (3 fields) |  |
| `Organizer` |  | `Record` (1 fields) |  |
| `Recurrence` |  | `Record` (2 fields) |  |
| `Reminder` | Reminder | `Number` | Time in minutes before event start to remind *advanced* |
| `ResponseRequested` | Response requested | `Boolean` | Set to true if the sender would like a response when the event is accepted or declined *advanced* |
| `ResponseStatus` |  | `Record` (2 fields) |  |
| `SeriesMasterId` | Series master id | `String` | Unique identifier for Series Master event type *advanced* |
| `ShowAs` | Show as | `String` | Shows as free or busy *advanced* |
| `Start` | Start time | `DateTime` | The start time of the event |
| `StartTimeZone` | Start time zone | `String` | This property specifies the time zone of the meeting start time. The value must be as defined in Windows (example: 'Pacific Standard Time'). *advanced* |
| `Subject` | Subject | `String` | Event subject |
| `Type` | Type | `String` | The event type: Single Instance, Occurrence, Exception, or Series Master *advanced* |
| `WebLink` | Web link | `String` | The preview of the message associated with the event *advanced* |
| `Reason` | Reason | `String` | The reason property used by O365 sync events protocol, will be 'deleted' if its a deleted event. *advanced* |

### CalendarGetItem

**Get event (V1)** — This operation gets a specific event from a calendar. (V1)

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `id` | Item id | `String` | ✓ | Select an event |

**Output** — `Record` (30 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Id` | Id | `String` | The Event's unique identifier |
| `Attendees` | Attendees | `Table` (3 fields) | List of attendees for the event |
| `Body` |  | `Record` (2 fields) |  |
| `BodyPreview` | Body preview | `String` | The preview of the message associated with the event *advanced* |
| `Categories` | Categories | `Table` (1 fields) | The categories associated with the event *advanced* |
| `ChangeKey` | Change key | `String` | This property identifies the version of the event object. Every time the event is changed, ChangeKey changes as well. *advanced* |
| `DateTimeCreated` | Created time | `DateTime` | The date and time that the event was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` | The date and time that the event was last modified *advanced* |
| `End` | End time | `DateTime` | The end time of the event |
| `EndTimeZone` | End time zone | `String` | This property specifies the time zone of the meeting end time. The value must be as defined in Windows (example: 'Pacific Standard Time'). *advanced* |
| `HasAttachments` | Has attachments? | `Boolean` | Set to true if the event has attachments *advanced* |
| `ICalUId` | Event Unique ID | `String` | A unique identifier that is shared by all instances of an event across different calendars *advanced* |
| `Importance` | Importance | `String` | The importance of the event: Low, Normal, or High *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` | Set to true if the event lasts all day *advanced* |
| `IsCancelled` | Is cancelled? | `Boolean` | Set to true if the event has been canceled *advanced* |
| `IsOrganizer` | Is Organizer | `Boolean` | Set to true if the message sender is also the organizer *advanced* |
| `Location` |  | `Record` (3 fields) |  |
| `Organizer` |  | `Record` (1 fields) |  |
| `Recurrence` |  | `Record` (2 fields) |  |
| `Reminder` | Reminder | `Number` | Time in minutes before event start to remind *advanced* |
| `ResponseRequested` | Response requested | `Boolean` | Set to true if the sender would like a response when the event is accepted or declined *advanced* |
| `ResponseStatus` |  | `Record` (2 fields) |  |
| `SeriesMasterId` | Series master id | `String` | Unique identifier for Series Master event type *advanced* |
| `ShowAs` | Show as | `String` | Shows as free or busy *advanced* |
| `Start` | Start time | `DateTime` | The start time of the event |
| `StartTimeZone` | Start time zone | `String` | This property specifies the time zone of the meeting start time. The value must be as defined in Windows (example: 'Pacific Standard Time'). *advanced* |
| `Subject` | Subject | `String` | Event subject |
| `Type` | Type | `String` | The event type: Single Instance, Occurrence, Exception, or Series Master *advanced* |
| `WebLink` | Web link | `String` | The preview of the message associated with the event *advanced* |
| `Reason` | Reason | `String` | The reason property used by O365 sync events protocol, will be 'deleted' if its a deleted event. *advanced* |

### CalendarDeleteItem

**Delete event** — This operation deletes an event in a calendar.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `id` | Id | `String` | ✓ | Select an event |

**Output** — `String`

### CalendarPatchItem

**Update event (V1)** — This operation updates an event in a calendar. (V1)

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `id` | Id | `String` | ✓ | Select an event |
| `End` | End time | `DateTime` | ✓ | The end time of the event |
| `Start` | Start time | `DateTime` | ✓ | The start time of the event |
| `Subject` | Subject | `String` | ✓ | Event subject |
| `Id` | Id | `String` |  | The Event's unique identifier |
| `Attendees` | Attendees | `Table` (3 fields) |  | List of attendees for the event |
| `Body` | Item | `Record` (2 fields) |  | Event to update |
| `BodyPreview` | Body preview | `String` |  | The preview of the message associated with the event *advanced* |
| `Categories` | Categories | `Table` (1 fields) |  | The categories associated with the event *advanced* |
| `ChangeKey` | Change key | `String` |  | This property identifies the version of the event object. Every time the event is changed, ChangeKey changes as well. *advanced* |
| `DateTimeCreated` | Created time | `DateTime` |  | The date and time that the event was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` |  | The date and time that the event was last modified *advanced* |
| `EndTimeZone` | End time zone | `String` |  | This property specifies the time zone of the meeting end time. The value must be as defined in Windows (example: 'Pacific Standard Time'). *advanced* |
| `HasAttachments` | Has attachments? | `Boolean` |  | Set to true if the event has attachments *advanced* |
| `ICalUId` | Event Unique ID | `String` |  | A unique identifier that is shared by all instances of an event across different calendars *advanced* |
| `Importance` | Importance | `String` |  | The importance of the event: Low, Normal, or High *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` |  | Set to true if the event lasts all day *advanced* |
| `IsCancelled` | Is cancelled? | `Boolean` |  | Set to true if the event has been canceled *advanced* |
| `IsOrganizer` | Is Organizer | `Boolean` |  | Set to true if the message sender is also the organizer *advanced* |
| `Location` | Item | `Record` (3 fields) |  | Event to update |
| `Organizer` | Item | `Record` (1 fields) |  | Event to update |
| `Recurrence` | Item | `Record` (2 fields) |  | Event to update |
| `Reminder` | Reminder | `Number` |  | Time in minutes before event start to remind *advanced* |
| `ResponseRequested` | Response requested | `Boolean` |  | Set to true if the sender would like a response when the event is accepted or declined *advanced* |
| `ResponseStatus` | Item | `Record` (2 fields) |  | Event to update |
| `SeriesMasterId` | Series master id | `String` |  | Unique identifier for Series Master event type *advanced* |
| `ShowAs` | Show as | `String` |  | Shows as free or busy *advanced* |
| `StartTimeZone` | Start time zone | `String` |  | This property specifies the time zone of the meeting start time. The value must be as defined in Windows (example: 'Pacific Standard Time'). *advanced* |
| `Type` | Type | `String` |  | The event type: Single Instance, Occurrence, Exception, or Series Master *advanced* |
| `WebLink` | Web link | `String` |  | The preview of the message associated with the event *advanced* |
| `Reason` | Reason | `String` |  | The reason property used by O365 sync events protocol, will be 'deleted' if its a deleted event. *advanced* |

**Output** — `Record` (30 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Id` | Id | `String` | The Event's unique identifier |
| `Attendees` | Attendees | `Table` (3 fields) | List of attendees for the event |
| `Body` |  | `Record` (2 fields) |  |
| `BodyPreview` | Body preview | `String` | The preview of the message associated with the event *advanced* |
| `Categories` | Categories | `Table` (1 fields) | The categories associated with the event *advanced* |
| `ChangeKey` | Change key | `String` | This property identifies the version of the event object. Every time the event is changed, ChangeKey changes as well. *advanced* |
| `DateTimeCreated` | Created time | `DateTime` | The date and time that the event was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` | The date and time that the event was last modified *advanced* |
| `End` | End time | `DateTime` | The end time of the event |
| `EndTimeZone` | End time zone | `String` | This property specifies the time zone of the meeting end time. The value must be as defined in Windows (example: 'Pacific Standard Time'). *advanced* |
| `HasAttachments` | Has attachments? | `Boolean` | Set to true if the event has attachments *advanced* |
| `ICalUId` | Event Unique ID | `String` | A unique identifier that is shared by all instances of an event across different calendars *advanced* |
| `Importance` | Importance | `String` | The importance of the event: Low, Normal, or High *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` | Set to true if the event lasts all day *advanced* |
| `IsCancelled` | Is cancelled? | `Boolean` | Set to true if the event has been canceled *advanced* |
| `IsOrganizer` | Is Organizer | `Boolean` | Set to true if the message sender is also the organizer *advanced* |
| `Location` |  | `Record` (3 fields) |  |
| `Organizer` |  | `Record` (1 fields) |  |
| `Recurrence` |  | `Record` (2 fields) |  |
| `Reminder` | Reminder | `Number` | Time in minutes before event start to remind *advanced* |
| `ResponseRequested` | Response requested | `Boolean` | Set to true if the sender would like a response when the event is accepted or declined *advanced* |
| `ResponseStatus` |  | `Record` (2 fields) |  |
| `SeriesMasterId` | Series master id | `String` | Unique identifier for Series Master event type *advanced* |
| `ShowAs` | Show as | `String` | Shows as free or busy *advanced* |
| `Start` | Start time | `DateTime` | The start time of the event |
| `StartTimeZone` | Start time zone | `String` | This property specifies the time zone of the meeting start time. The value must be as defined in Windows (example: 'Pacific Standard Time'). *advanced* |
| `Subject` | Subject | `String` | Event subject |
| `Type` | Type | `String` | The event type: Single Instance, Occurrence, Exception, or Series Master *advanced* |
| `WebLink` | Web link | `String` | The preview of the message associated with the event *advanced* |
| `Reason` | Reason | `String` | The reason property used by O365 sync events protocol, will be 'deleted' if its a deleted event. *advanced* |

### V3CalendarGetItems

**Get events (V3)** — This operation gets events from a calendar. (V3)

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `$filter` | Filter Query | `String` |  | An ODATA filter query to restrict the entries returned (e.g. stringColumn eq 'string' OR numberColumn lt 123). *advanced* |
| `$orderby` | Order By | `String` |  | An ODATA orderBy query for specifying the order of entries. *advanced* |
| `$top` | Top Count | `Number` |  | Total number of entries to retrieve (default = all). *advanced* |
| `$skip` | Skip Count | `Number` |  | The number of entries to skip (default = 0). *advanced* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (28 fields) | List of calendar items |

### V3CalendarPostItem

**Create event (V3)** — This operation creates a new event in a calendar.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `Subject` | Subject | `String` | ✓ | Event subject |
| `Start` | Start time | `DateTime` | ✓ | Start time of the event (example: '2016-11-01T14:30:00Z') |
| `End` | End time | `DateTime` | ✓ | End time of the event (example: '2016-11-01T15:30:00Z') |
| `TimeZone` | Time zone | `String` |  | Time zone of the event *advanced* |
| `RequiredAttendees` | Required attendees | `String` |  | Required attendees for the event separated by semicolons *advanced* |
| `OptionalAttendees` | Optional attendees | `String` |  | Optional attendees for the event separated by semicolons *advanced* |
| `ResourceAttendees` | Resource attendees | `String` |  | Resource attendees for the event separated by semicolons *advanced* |
| `Body` | Body | `String` |  | Body of the message associated with the event *advanced* |
| `Location` | Location | `String` |  | Location of the event *advanced* |
| `Importance` | Importance | `String` |  | The importance of the event: Low, Normal, or High *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` |  | Set to true if the event lasts all day *advanced* |
| `Recurrence` | Recurrence | `String` |  | The recurrence pattern for the event *advanced* |
| `RecurrenceEnd` | Recurrence end time | `DateTime` |  | End time of the recurrence *advanced* |
| `NumberOfOccurrences` | Number of occurrences | `Number` |  | How many times to repeat the event *advanced* |
| `Reminder` | Reminder | `Number` |  | Time in minutes before event start to remind *advanced* |
| `ShowAs` | Show as | `String` |  | Status to show during the event *advanced* |
| `ResponseRequested` | Response requested | `Boolean` |  | Set to true if the sender would like a response when the event is accepted or declined *advanced* |

**Output** — `Record` (28 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Importance` | Importance | `String` | The importance of the event: Low, Normal, or High *advanced* |
| `ResponseType` | Response type | `String` | The response type of the event: None, Organizer, TentativelyAccepted, Accepted, Declined or NotResponded *advanced* |
| `Recurrence` | Recurrence | `String` | The recurrence pattern for the event *advanced* |
| `ShowAs` | Show as | `String` | Status to show during the event *advanced* |
| `Subject` | Subject | `String` | Event subject |
| `Start` | Start time | `DateTime` | Start time of the event (example: '2016-11-01T14:30:00Z') |
| `End` | End time | `DateTime` | End time of the event (example: '2016-11-01T15:30:00Z') |
| `ResponseTime` | Response time | `DateTime` | The response time of the event *advanced* |
| `ICalUId` | Event Unique ID | `String` | A unique identifier that is shared by all instances of an event across different calendars *advanced* |
| `Id` | Id | `String` | The event's unique identifier *advanced* |
| `DateTimeCreated` | Created time | `DateTime` | The date and time that the event was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` | The date and time that the event was last modified *advanced* |
| `Organizer` | Organizer | `String` | The organizer of the event *advanced* |
| `TimeZone` | Time zone | `String` | Time zone of the event *advanced* |
| `SeriesMasterId` | Series master id | `String` | Unique identifier for Series Master event type *advanced* |
| `Categories` | Categories | `Table` (1 fields) | The categories associated with the event *advanced* |
| `WebLink` | Web link | `String` | The URL to open the event in Outlook Web App *advanced* |
| `RequiredAttendees` | Required attendees | `String` | Required attendees for the event separated by semicolons *advanced* |
| `OptionalAttendees` | Optional attendees | `String` | Optional attendees for the event separated by semicolons *advanced* |
| `ResourceAttendees` | Resource attendees | `String` | Resource attendees for the event separated by semicolons *advanced* |
| `Body` | Body | `String` | Body of the message associated with the event *advanced* |
| `IsHtml` | Is HTML | `Boolean` | Set to true if the body is Html *advanced* |
| `Location` | Location | `String` | Location of the event *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` | Set to true if the event lasts all day *advanced* |
| `RecurrenceEnd` | Recurrence end time | `DateTime` | End time of the recurrence *advanced* |
| `NumberOfOccurrences` | Number of occurrences | `Number` | How many times to repeat the event *advanced* |
| `Reminder` | Reminder | `Number` | Time in minutes before event start to remind *advanced* |
| `ResponseRequested` | Response requested | `Boolean` | Set to true if the sender would like a response when the event is accepted or declined *advanced* |

### V4CalendarGetItems

**Get events (V4)** — This operation gets events from a calendar using Graph API. (V4)

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `$filter` | Filter Query | `String` |  | An ODATA filter query to restrict the entries returned (e.g. stringColumn eq 'string' OR numberColumn lt 123). *advanced* |
| `$orderby` | Order By | `String` |  | An ODATA orderBy query for specifying the order of entries. *advanced* |
| `$top` | Top Count | `Number` |  | Total number of entries to retrieve (default = all). *advanced* |
| `$skip` | Skip Count | `Number` |  | The number of entries to skip (default = 0). *advanced* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (32 fields) | List of calendar items |

### V4CalendarPostItem

**Create event (V4)** — This operation creates a new event in a calendar.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `subject` | Subject | `String` | ✓ | Event subject |
| `start` | Start time | `DateTime` | ✓ | Start time of the event (example: '2017-08-29T04:00:00') |
| `end` | End time | `DateTime` | ✓ | End time of the event (example: '2017-08-29T05:00:00') |
| `timeZone` | Time zone | `String` | ✓ | Time zone of the event |
| `requiredAttendees` | Required attendees | `String` |  | Required attendees for the event separated by semicolons *advanced* |
| `optionalAttendees` | Optional attendees | `String` |  | Optional attendees for the event separated by semicolons *advanced* |
| `resourceAttendees` | Resource attendees | `String` |  | Resource attendees for the event separated by semicolons *advanced* |
| `body` | Body | `String` |  | Body of the message associated with the event *advanced* |
| `categories` | Categories | `Table` (1 fields) |  | The categories associated with the event *advanced* |
| `location` | Location | `String` |  | Location of the event *advanced* |
| `importance` | Importance | `String` |  | The importance of the event: low, normal, or high *advanced* |
| `isAllDay` | Is all day event? | `Boolean` |  | Set to true if the event lasts all day *advanced* |
| `recurrence` | Recurrence | `String` |  | The recurrence pattern for the event: none, daily, weekly, monthly or yearly *advanced* |
| `selectedDaysOfWeek` | Selected days of week | `Table` (1 fields) |  | Days of week for weekly recurrence *advanced* |
| `recurrenceEnd` | Recurrence end date | `Date` |  | End Date of the recurrence *advanced* |
| `numberOfOccurences` | Number of occurrences | `Number` |  | How many times to repeat the event *advanced* |
| `reminderMinutesBeforeStart` | Reminder | `Number` |  | Time in minutes before event start to remind *advanced* |
| `isReminderOn` | Is reminder on | `Boolean` |  | Set to true if an alert is set to remind the user of the event. *advanced* |
| `showAs` | Show as | `String` |  | Status to show during the event: free, tentative, busy, oof, workingElsewhere or unknown *advanced* |
| `responseRequested` | Response requested | `Boolean` |  | Set to true if the sender would like a response when the event is accepted or declined *advanced* |
| `sensitivity` | Sensitivity | `String` |  | The possible values are: normal, personal, private, confidential *advanced* |

**Output** — `Record` (32 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `subject` | Subject | `String` | Event subject |
| `start` | Start time | `DateTime` | Start time of the event (example: '2017-08-29T04:00:00.0000000') |
| `end` | End time | `DateTime` | End time of the event (example: '2017-08-29T05:00:00.0000000') |
| `startWithTimeZone` | Start time with time zone | `DateTime` | Start time of the event with time zone (example: '2017-08-29T04:00:00.0000000+00:00') |
| `endWithTimeZone` | End time with time zone | `DateTime` | End time of the event with time zone (example: '2017-08-29T05:00:00.0000000+00:00') |
| `body` | Body | `String` | Body of the message associated with the event *advanced* |
| `isHtml` | Is HTML | `Boolean` | Set to true if the body is Html *advanced* |
| `responseType` | Response type | `String` | The response type of the event (none, organizer, tentativelyAccepted, accepted, declined or notResponded) *advanced* |
| `responseTime` | Response time | `DateTime` | The response time of the event *advanced* |
| `id` | Id | `String` | The event's unique identifier *advanced* |
| `createdDateTime` | Created time | `DateTime` | The date and time that the event was created *advanced* |
| `lastModifiedDateTime` | Last modified time | `DateTime` | The date and time that the event was last modified *advanced* |
| `organizer` | Organizer | `String` | The organizer of the event *advanced* |
| `timeZone` | Time zone | `String` | Time zone of the event *advanced* |
| `seriesMasterId` | Series master id | `String` | Unique identifier for Series Master event type *advanced* |
| `iCalUId` | iCalUId | `String` | A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series *advanced* |
| `categories` | Categories | `Table` (1 fields) | The categories associated with the event *advanced* |
| `webLink` | Web link | `String` | The URL to open the event in Outlook Web App *advanced* |
| `requiredAttendees` | Required attendees | `String` | Required attendees for the event separated by semicolons *advanced* |
| `optionalAttendees` | Optional attendees | `String` | Optional attendees for the event separated by semicolons *advanced* |
| `resourceAttendees` | Resource attendees | `String` | Resource attendees for the event separated by semicolons *advanced* |
| `location` | Location | `String` | Location of the event *advanced* |
| `importance` | Importance | `String` | The importance of the event: low, normal, or high *advanced* |
| `isAllDay` | Is all day event? | `Boolean` | Set to true if the event lasts all day *advanced* |
| `recurrence` | Recurrence | `String` | The recurrence pattern for the event: none, daily, weekly, monthly or yearly *advanced* |
| `recurrenceEnd` | Recurrence end date | `Date` | End Date of the recurrence *advanced* |
| `numberOfOccurences` | Number of occurrences | `Number` | How many times to repeat the event *advanced* |
| `reminderMinutesBeforeStart` | Reminder | `Number` | Time in minutes before event start to remind *advanced* |
| `isReminderOn` | Is reminder on | `Boolean` | Set to true if an alert is set to remind the user of the event. *advanced* |
| `showAs` | Show as | `String` | Status to show during the event: free, tentative, busy, oof, workingElsewhere or unknown *advanced* |
| `responseRequested` | Response requested | `Boolean` | Set to true if the sender would like a response when the event is accepted or declined *advanced* |
| `sensitivity` | Sensitivity | `String` | The possible values are: normal, personal, private, confidential *advanced* |

### V2CalendarGetItems

**Get events (V2)** — This operation gets events from a calendar. (V2)

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `$filter` | Filter Query | `String` |  | An ODATA filter query to restrict the entries returned (e.g. stringColumn eq 'string' OR numberColumn lt 123). *advanced* |
| `$orderby` | Order By | `String` |  | An ODATA orderBy query for specifying the order of entries. *advanced* |
| `$top` | Top Count | `Number` |  | Total number of entries to retrieve (default = all). *advanced* |
| `$skip` | Skip Count | `Number` |  | The number of entries to skip (default = 0). *advanced* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (28 fields) | List of values |

### V2CalendarPostItem

**Create event (V2)** — This operation creates a new event in a calendar. (V2)

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `Subject` | Subject | `String` | ✓ | Event subject |
| `Start` | Start time | `DateTime` | ✓ | Start time of the event (example: '2016-11-01T14:30:00Z') |
| `End` | End time | `DateTime` | ✓ | End time of the event (example: '2016-11-01T15:30:00Z') |
| `TimeZone` | Time zone | `String` |  | Time zone of the event *advanced* |
| `RequiredAttendees` | Required attendees | `String` |  | Required attendees for the event separated by semicolons *advanced* |
| `OptionalAttendees` | Optional attendees | `String` |  | Optional attendees for the event separated by semicolons *advanced* |
| `ResourceAttendees` | Resource attendees | `String` |  | Resource attendees for the event separated by semicolons *advanced* |
| `Body` | Body | `String` |  | Body of the message associated with the event *advanced* |
| `IsHtml` | Is HTML | `Boolean` |  | Set to true if the body is Html *advanced* |
| `Location` | Location | `String` |  | Location of the event *advanced* |
| `Importance` | Importance | `String` |  | The importance of the event: Low, Normal, or High *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` |  | Set to true if the event lasts all day *advanced* |
| `Recurrence` | Recurrence | `String` |  | The recurrence pattern for the event *advanced* |
| `RecurrenceEnd` | Recurrence end time | `DateTime` |  | End time of the recurrence *advanced* |
| `NumberOfOccurrences` | Number of occurrences | `Number` |  | How many times to repeat the event *advanced* |
| `Reminder` | Reminder | `Number` |  | Time in minutes before event start to remind *advanced* |
| `ShowAs` | Show as | `String` |  | Status to show during the event *advanced* |
| `ResponseRequested` | Response requested | `Boolean` |  | Set to true if the sender would like a response when the event is accepted or declined *advanced* |

**Output** — `Record` (28 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Importance` | Importance | `String` | The importance of the event: Low, Normal, or High *advanced* |
| `ResponseType` | Response type | `String` | The response type of the event: None, Organizer, TentativelyAccepted, Accepted, Declined or NotResponded *advanced* |
| `Recurrence` | Recurrence | `String` | The recurrence pattern for the event *advanced* |
| `ShowAs` | Show as | `String` | Status to show during the event *advanced* |
| `Subject` | Subject | `String` | Event subject |
| `Start` | Start time | `DateTime` | Start time of the event (example: '2016-11-01T14:30:00Z') |
| `End` | End time | `DateTime` | End time of the event (example: '2016-11-01T15:30:00Z') |
| `ResponseTime` | Response time | `DateTime` | The response time of the event *advanced* |
| `ICalUId` | Event Unique ID | `String` | A unique identifier that is shared by all instances of an event across different calendars *advanced* |
| `Id` | Id | `String` | The event's unique identifier *advanced* |
| `DateTimeCreated` | Created time | `DateTime` | The date and time that the event was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` | The date and time that the event was last modified *advanced* |
| `Organizer` | Organizer | `String` | The organizer of the event *advanced* |
| `TimeZone` | Time zone | `String` | Time zone of the event *advanced* |
| `SeriesMasterId` | Series master id | `String` | Unique identifier for Series Master event type *advanced* |
| `Categories` | Categories | `Table` (1 fields) | The categories associated with the event *advanced* |
| `WebLink` | Web link | `String` | The URL to open the event in Outlook Web App *advanced* |
| `RequiredAttendees` | Required attendees | `String` | Required attendees for the event separated by semicolons *advanced* |
| `OptionalAttendees` | Optional attendees | `String` | Optional attendees for the event separated by semicolons *advanced* |
| `ResourceAttendees` | Resource attendees | `String` | Resource attendees for the event separated by semicolons *advanced* |
| `Body` | Body | `String` | Body of the message associated with the event *advanced* |
| `IsHtml` | Is HTML | `Boolean` | Set to true if the body is Html *advanced* |
| `Location` | Location | `String` | Location of the event *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` | Set to true if the event lasts all day *advanced* |
| `RecurrenceEnd` | Recurrence end time | `DateTime` | End time of the recurrence *advanced* |
| `NumberOfOccurrences` | Number of occurrences | `Number` | How many times to repeat the event *advanced* |
| `Reminder` | Reminder | `Number` | Time in minutes before event start to remind *advanced* |
| `ResponseRequested` | Response requested | `Boolean` | Set to true if the sender would like a response when the event is accepted or declined *advanced* |

### GetEventsCalendarViewV2

**Get calendar view of events (V2)** — This operation gets all events (including instances of recurrences) in a calendar. Recurrence property is null in this case.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `calendarId` | Calendar Id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `startDateTimeOffset` | Start Time | `String` | ✓ | Start time (example: '2017-01-01T08:00:00-07:00') |
| `endDateTimeOffset` | End Time | `String` | ✓ | End time (example: '2017-02-01T08:00:00-07:00') |
| `$filter` | Filter Query | `String` |  | An ODATA filter query to restrict the entries returned (e.g. stringColumn eq 'string' OR numberColumn lt 123). *advanced* |
| `$orderby` | Order By | `String` |  | An ODATA orderBy query for specifying the order of entries. *advanced* |
| `$top` | Top Count | `Number` |  | Total number of entries to retrieve (default = all). *advanced* |
| `$skip` | Skip Count | `Number` |  | The number of entries to skip (default = 0). *advanced* |
| `search` | Search | `String` |  | Search text for matching event body and subject *advanced* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (28 fields) | List of values |

### GetEventsCalendarViewV3

**Get calendar view of events (V3)** — This operation gets all events (including instances of recurrences) in a calendar using Graph API. Recurrence property is null in this case.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `calendarId` | Calendar Id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `startDateTimeUtc` | Start Time | `String` | ✓ | Start time in UTC (example: '2017-01-01T08:00:00') |
| `endDateTimeUtc` | End Time | `String` | ✓ | End time in UTC (example: '2017-02-01T08:00:00') |
| `$filter` | Filter Query | `String` |  | An ODATA filter query to restrict the entries returned (e.g. stringColumn eq 'string' OR numberColumn lt 123). *advanced* |
| `$orderby` | Order By | `String` |  | An ODATA orderBy query for specifying the order of entries. *advanced* |
| `$top` | Top Count | `Number` |  | Total number of entries to retrieve (default = all). *advanced* |
| `$skip` | Skip Count | `Number` |  | The number of entries to skip (default = 0). *advanced* |
| `search` | Search | `String` |  | Search text for matching event body and subject *advanced* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (32 fields) | List of values |

### V2CalendarGetItem

**Get event (V2)** — This operation gets a specific event from a calendar. (V2)

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `id` | Item id | `String` | ✓ | Select an event |

**Output** — `Record` (28 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Importance` | Importance | `String` | The importance of the event: Low, Normal, or High *advanced* |
| `ResponseType` | Response type | `String` | The response type of the event: None, Organizer, TentativelyAccepted, Accepted, Declined or NotResponded *advanced* |
| `Recurrence` | Recurrence | `String` | The recurrence pattern for the event *advanced* |
| `ShowAs` | Show as | `String` | Status to show during the event *advanced* |
| `Subject` | Subject | `String` | Event subject |
| `Start` | Start time | `DateTime` | Start time of the event (example: '2016-11-01T14:30:00Z') |
| `End` | End time | `DateTime` | End time of the event (example: '2016-11-01T15:30:00Z') |
| `ResponseTime` | Response time | `DateTime` | The response time of the event *advanced* |
| `ICalUId` | Event Unique ID | `String` | A unique identifier that is shared by all instances of an event across different calendars *advanced* |
| `Id` | Id | `String` | The event's unique identifier *advanced* |
| `DateTimeCreated` | Created time | `DateTime` | The date and time that the event was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` | The date and time that the event was last modified *advanced* |
| `Organizer` | Organizer | `String` | The organizer of the event *advanced* |
| `TimeZone` | Time zone | `String` | Time zone of the event *advanced* |
| `SeriesMasterId` | Series master id | `String` | Unique identifier for Series Master event type *advanced* |
| `Categories` | Categories | `Table` (1 fields) | The categories associated with the event *advanced* |
| `WebLink` | Web link | `String` | The URL to open the event in Outlook Web App *advanced* |
| `RequiredAttendees` | Required attendees | `String` | Required attendees for the event separated by semicolons *advanced* |
| `OptionalAttendees` | Optional attendees | `String` | Optional attendees for the event separated by semicolons *advanced* |
| `ResourceAttendees` | Resource attendees | `String` | Resource attendees for the event separated by semicolons *advanced* |
| `Body` | Body | `String` | Body of the message associated with the event *advanced* |
| `IsHtml` | Is HTML | `Boolean` | Set to true if the body is Html *advanced* |
| `Location` | Location | `String` | Location of the event *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` | Set to true if the event lasts all day *advanced* |
| `RecurrenceEnd` | Recurrence end time | `DateTime` | End time of the recurrence *advanced* |
| `NumberOfOccurrences` | Number of occurrences | `Number` | How many times to repeat the event *advanced* |
| `Reminder` | Reminder | `Number` | Time in minutes before event start to remind *advanced* |
| `ResponseRequested` | Response requested | `Boolean` | Set to true if the sender would like a response when the event is accepted or declined *advanced* |

### V2CalendarPatchItem

**Update event (V2)** — This operation updates an event in a calendar. (V2)

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `id` | Id | `String` | ✓ | Select an event |
| `Subject` | Subject | `String` | ✓ | Event subject |
| `Start` | Start time | `DateTime` | ✓ | Start time of the event (example: '2016-11-01T14:30:00Z') |
| `End` | End time | `DateTime` | ✓ | End time of the event (example: '2016-11-01T15:30:00Z') |
| `TimeZone` | Time zone | `String` |  | Time zone of the event *advanced* |
| `RequiredAttendees` | Required attendees | `String` |  | Required attendees for the event separated by semicolons *advanced* |
| `OptionalAttendees` | Optional attendees | `String` |  | Optional attendees for the event separated by semicolons *advanced* |
| `ResourceAttendees` | Resource attendees | `String` |  | Resource attendees for the event separated by semicolons *advanced* |
| `Body` | Body | `String` |  | Body of the message associated with the event *advanced* |
| `IsHtml` | Is HTML | `Boolean` |  | Set to true if the body is Html *advanced* |
| `Location` | Location | `String` |  | Location of the event *advanced* |
| `Importance` | Importance | `String` |  | The importance of the event: Low, Normal, or High *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` |  | Set to true if the event lasts all day *advanced* |
| `Recurrence` | Recurrence | `String` |  | The recurrence pattern for the event *advanced* |
| `RecurrenceEnd` | Recurrence end time | `DateTime` |  | End time of the recurrence *advanced* |
| `NumberOfOccurrences` | Number of occurrences | `Number` |  | How many times to repeat the event *advanced* |
| `Reminder` | Reminder | `Number` |  | Time in minutes before event start to remind *advanced* |
| `ShowAs` | Show as | `String` |  | Status to show during the event *advanced* |
| `ResponseRequested` | Response requested | `Boolean` |  | Set to true if the sender would like a response when the event is accepted or declined *advanced* |

**Output** — `Record` (28 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Importance` | Importance | `String` | The importance of the event: Low, Normal, or High *advanced* |
| `ResponseType` | Response type | `String` | The response type of the event: None, Organizer, TentativelyAccepted, Accepted, Declined or NotResponded *advanced* |
| `Recurrence` | Recurrence | `String` | The recurrence pattern for the event *advanced* |
| `ShowAs` | Show as | `String` | Status to show during the event *advanced* |
| `Subject` | Subject | `String` | Event subject |
| `Start` | Start time | `DateTime` | Start time of the event (example: '2016-11-01T14:30:00Z') |
| `End` | End time | `DateTime` | End time of the event (example: '2016-11-01T15:30:00Z') |
| `ResponseTime` | Response time | `DateTime` | The response time of the event *advanced* |
| `ICalUId` | Event Unique ID | `String` | A unique identifier that is shared by all instances of an event across different calendars *advanced* |
| `Id` | Id | `String` | The event's unique identifier *advanced* |
| `DateTimeCreated` | Created time | `DateTime` | The date and time that the event was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` | The date and time that the event was last modified *advanced* |
| `Organizer` | Organizer | `String` | The organizer of the event *advanced* |
| `TimeZone` | Time zone | `String` | Time zone of the event *advanced* |
| `SeriesMasterId` | Series master id | `String` | Unique identifier for Series Master event type *advanced* |
| `Categories` | Categories | `Table` (1 fields) | The categories associated with the event *advanced* |
| `WebLink` | Web link | `String` | The URL to open the event in Outlook Web App *advanced* |
| `RequiredAttendees` | Required attendees | `String` | Required attendees for the event separated by semicolons *advanced* |
| `OptionalAttendees` | Optional attendees | `String` | Optional attendees for the event separated by semicolons *advanced* |
| `ResourceAttendees` | Resource attendees | `String` | Resource attendees for the event separated by semicolons *advanced* |
| `Body` | Body | `String` | Body of the message associated with the event *advanced* |
| `IsHtml` | Is HTML | `Boolean` | Set to true if the body is Html *advanced* |
| `Location` | Location | `String` | Location of the event *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` | Set to true if the event lasts all day *advanced* |
| `RecurrenceEnd` | Recurrence end time | `DateTime` | End time of the recurrence *advanced* |
| `NumberOfOccurrences` | Number of occurrences | `Number` | How many times to repeat the event *advanced* |
| `Reminder` | Reminder | `Number` | Time in minutes before event start to remind *advanced* |
| `ResponseRequested` | Response requested | `Boolean` | Set to true if the sender would like a response when the event is accepted or declined *advanced* |

### V3CalendarGetItem

**Get event (V3)** — This operation gets a specific event from a calendar using Graph API. (V3)

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `id` | Item id | `String` | ✓ | Select an event |

**Output** — `Record` (32 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `subject` | Subject | `String` | Event subject |
| `start` | Start time | `DateTime` | Start time of the event (example: '2017-08-29T04:00:00.0000000') |
| `end` | End time | `DateTime` | End time of the event (example: '2017-08-29T05:00:00.0000000') |
| `startWithTimeZone` | Start time with time zone | `DateTime` | Start time of the event with time zone (example: '2017-08-29T04:00:00.0000000+00:00') |
| `endWithTimeZone` | End time with time zone | `DateTime` | End time of the event with time zone (example: '2017-08-29T05:00:00.0000000+00:00') |
| `body` | Body | `String` | Body of the message associated with the event *advanced* |
| `isHtml` | Is HTML | `Boolean` | Set to true if the body is Html *advanced* |
| `responseType` | Response type | `String` | The response type of the event (none, organizer, tentativelyAccepted, accepted, declined or notResponded) *advanced* |
| `responseTime` | Response time | `DateTime` | The response time of the event *advanced* |
| `id` | Id | `String` | The event's unique identifier *advanced* |
| `createdDateTime` | Created time | `DateTime` | The date and time that the event was created *advanced* |
| `lastModifiedDateTime` | Last modified time | `DateTime` | The date and time that the event was last modified *advanced* |
| `organizer` | Organizer | `String` | The organizer of the event *advanced* |
| `timeZone` | Time zone | `String` | Time zone of the event *advanced* |
| `seriesMasterId` | Series master id | `String` | Unique identifier for Series Master event type *advanced* |
| `iCalUId` | iCalUId | `String` | A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series *advanced* |
| `categories` | Categories | `Table` (1 fields) | The categories associated with the event *advanced* |
| `webLink` | Web link | `String` | The URL to open the event in Outlook Web App *advanced* |
| `requiredAttendees` | Required attendees | `String` | Required attendees for the event separated by semicolons *advanced* |
| `optionalAttendees` | Optional attendees | `String` | Optional attendees for the event separated by semicolons *advanced* |
| `resourceAttendees` | Resource attendees | `String` | Resource attendees for the event separated by semicolons *advanced* |
| `location` | Location | `String` | Location of the event *advanced* |
| `importance` | Importance | `String` | The importance of the event: low, normal, or high *advanced* |
| `isAllDay` | Is all day event? | `Boolean` | Set to true if the event lasts all day *advanced* |
| `recurrence` | Recurrence | `String` | The recurrence pattern for the event: none, daily, weekly, monthly or yearly *advanced* |
| `recurrenceEnd` | Recurrence end date | `Date` | End Date of the recurrence *advanced* |
| `numberOfOccurences` | Number of occurrences | `Number` | How many times to repeat the event *advanced* |
| `reminderMinutesBeforeStart` | Reminder | `Number` | Time in minutes before event start to remind *advanced* |
| `isReminderOn` | Is reminder on | `Boolean` | Set to true if an alert is set to remind the user of the event. *advanced* |
| `showAs` | Show as | `String` | Status to show during the event: free, tentative, busy, oof, workingElsewhere or unknown *advanced* |
| `responseRequested` | Response requested | `Boolean` | Set to true if the sender would like a response when the event is accepted or declined *advanced* |
| `sensitivity` | Sensitivity | `String` | The possible values are: normal, personal, private, confidential *advanced* |

### V3CalendarPatchItem

**Update event (V3)** — This operation updates an event in a calendar.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `id` | Id | `String` | ✓ | Select an event |
| `Subject` | Subject | `String` | ✓ | Event subject |
| `Start` | Start time | `DateTime` | ✓ | Start time of the event (example: '2016-11-01T14:30:00Z') |
| `End` | End time | `DateTime` | ✓ | End time of the event (example: '2016-11-01T15:30:00Z') |
| `TimeZone` | Time zone | `String` |  | Time zone of the event *advanced* |
| `RequiredAttendees` | Required attendees | `String` |  | Required attendees for the event separated by semicolons *advanced* |
| `OptionalAttendees` | Optional attendees | `String` |  | Optional attendees for the event separated by semicolons *advanced* |
| `ResourceAttendees` | Resource attendees | `String` |  | Resource attendees for the event separated by semicolons *advanced* |
| `Body` | Body | `String` |  | Body of the message associated with the event *advanced* |
| `Location` | Location | `String` |  | Location of the event *advanced* |
| `Importance` | Importance | `String` |  | The importance of the event: Low, Normal, or High *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` |  | Set to true if the event lasts all day *advanced* |
| `Recurrence` | Recurrence | `String` |  | The recurrence pattern for the event *advanced* |
| `RecurrenceEnd` | Recurrence end time | `DateTime` |  | End time of the recurrence *advanced* |
| `NumberOfOccurrences` | Number of occurrences | `Number` |  | How many times to repeat the event *advanced* |
| `Reminder` | Reminder | `Number` |  | Time in minutes before event start to remind *advanced* |
| `ShowAs` | Show as | `String` |  | Status to show during the event *advanced* |
| `ResponseRequested` | Response requested | `Boolean` |  | Set to true if the sender would like a response when the event is accepted or declined *advanced* |

**Output** — `Record` (28 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Importance` | Importance | `String` | The importance of the event: Low, Normal, or High *advanced* |
| `ResponseType` | Response type | `String` | The response type of the event: None, Organizer, TentativelyAccepted, Accepted, Declined or NotResponded *advanced* |
| `Recurrence` | Recurrence | `String` | The recurrence pattern for the event *advanced* |
| `ShowAs` | Show as | `String` | Status to show during the event *advanced* |
| `Subject` | Subject | `String` | Event subject |
| `Start` | Start time | `DateTime` | Start time of the event (example: '2016-11-01T14:30:00Z') |
| `End` | End time | `DateTime` | End time of the event (example: '2016-11-01T15:30:00Z') |
| `ResponseTime` | Response time | `DateTime` | The response time of the event *advanced* |
| `ICalUId` | Event Unique ID | `String` | A unique identifier that is shared by all instances of an event across different calendars *advanced* |
| `Id` | Id | `String` | The event's unique identifier *advanced* |
| `DateTimeCreated` | Created time | `DateTime` | The date and time that the event was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` | The date and time that the event was last modified *advanced* |
| `Organizer` | Organizer | `String` | The organizer of the event *advanced* |
| `TimeZone` | Time zone | `String` | Time zone of the event *advanced* |
| `SeriesMasterId` | Series master id | `String` | Unique identifier for Series Master event type *advanced* |
| `Categories` | Categories | `Table` (1 fields) | The categories associated with the event *advanced* |
| `WebLink` | Web link | `String` | The URL to open the event in Outlook Web App *advanced* |
| `RequiredAttendees` | Required attendees | `String` | Required attendees for the event separated by semicolons *advanced* |
| `OptionalAttendees` | Optional attendees | `String` | Optional attendees for the event separated by semicolons *advanced* |
| `ResourceAttendees` | Resource attendees | `String` | Resource attendees for the event separated by semicolons *advanced* |
| `Body` | Body | `String` | Body of the message associated with the event *advanced* |
| `IsHtml` | Is HTML | `Boolean` | Set to true if the body is Html *advanced* |
| `Location` | Location | `String` | Location of the event *advanced* |
| `IsAllDay` | Is all day event? | `Boolean` | Set to true if the event lasts all day *advanced* |
| `RecurrenceEnd` | Recurrence end time | `DateTime` | End time of the recurrence *advanced* |
| `NumberOfOccurrences` | Number of occurrences | `Number` | How many times to repeat the event *advanced* |
| `Reminder` | Reminder | `Number` | Time in minutes before event start to remind *advanced* |
| `ResponseRequested` | Response requested | `Boolean` | Set to true if the sender would like a response when the event is accepted or declined *advanced* |

### V4CalendarPatchItem

**Update event (V4)** — This operation updates an event in a calendar using Graph API.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `id` | Id | `String` | ✓ | Select an event |
| `subject` | Subject | `String` | ✓ | Event subject |
| `start` | Start time | `DateTime` | ✓ | Start time of the event (example: '2017-08-29T04:00:00') |
| `end` | End time | `DateTime` | ✓ | End time of the event (example: '2017-08-29T05:00:00') |
| `timeZone` | Time zone | `String` | ✓ | Time zone of the event |
| `requiredAttendees` | Required attendees | `String` |  | Required attendees for the event separated by semicolons *advanced* |
| `optionalAttendees` | Optional attendees | `String` |  | Optional attendees for the event separated by semicolons *advanced* |
| `resourceAttendees` | Resource attendees | `String` |  | Resource attendees for the event separated by semicolons *advanced* |
| `body` | Body | `String` |  | Body of the message associated with the event *advanced* |
| `categories` | Categories | `Table` (1 fields) |  | The categories associated with the event *advanced* |
| `location` | Location | `String` |  | Location of the event *advanced* |
| `importance` | Importance | `String` |  | The importance of the event: low, normal, or high *advanced* |
| `isAllDay` | Is all day event? | `Boolean` |  | Set to true if the event lasts all day *advanced* |
| `recurrence` | Recurrence | `String` |  | The recurrence pattern for the event: none, daily, weekly, monthly or yearly *advanced* |
| `selectedDaysOfWeek` | Selected days of week | `Table` (1 fields) |  | Days of week for weekly recurrence *advanced* |
| `recurrenceEnd` | Recurrence end date | `Date` |  | End Date of the recurrence *advanced* |
| `numberOfOccurences` | Number of occurrences | `Number` |  | How many times to repeat the event *advanced* |
| `reminderMinutesBeforeStart` | Reminder | `Number` |  | Time in minutes before event start to remind *advanced* |
| `isReminderOn` | Is reminder on | `Boolean` |  | Set to true if an alert is set to remind the user of the event. *advanced* |
| `showAs` | Show as | `String` |  | Status to show during the event: free, tentative, busy, oof, workingElsewhere or unknown *advanced* |
| `responseRequested` | Response requested | `Boolean` |  | Set to true if the sender would like a response when the event is accepted or declined *advanced* |
| `sensitivity` | Sensitivity | `String` |  | The possible values are: normal, personal, private, confidential *advanced* |

**Output** — `Record` (32 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `subject` | Subject | `String` | Event subject |
| `start` | Start time | `DateTime` | Start time of the event (example: '2017-08-29T04:00:00.0000000') |
| `end` | End time | `DateTime` | End time of the event (example: '2017-08-29T05:00:00.0000000') |
| `startWithTimeZone` | Start time with time zone | `DateTime` | Start time of the event with time zone (example: '2017-08-29T04:00:00.0000000+00:00') |
| `endWithTimeZone` | End time with time zone | `DateTime` | End time of the event with time zone (example: '2017-08-29T05:00:00.0000000+00:00') |
| `body` | Body | `String` | Body of the message associated with the event *advanced* |
| `isHtml` | Is HTML | `Boolean` | Set to true if the body is Html *advanced* |
| `responseType` | Response type | `String` | The response type of the event (none, organizer, tentativelyAccepted, accepted, declined or notResponded) *advanced* |
| `responseTime` | Response time | `DateTime` | The response time of the event *advanced* |
| `id` | Id | `String` | The event's unique identifier *advanced* |
| `createdDateTime` | Created time | `DateTime` | The date and time that the event was created *advanced* |
| `lastModifiedDateTime` | Last modified time | `DateTime` | The date and time that the event was last modified *advanced* |
| `organizer` | Organizer | `String` | The organizer of the event *advanced* |
| `timeZone` | Time zone | `String` | Time zone of the event *advanced* |
| `seriesMasterId` | Series master id | `String` | Unique identifier for Series Master event type *advanced* |
| `iCalUId` | iCalUId | `String` | A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series *advanced* |
| `categories` | Categories | `Table` (1 fields) | The categories associated with the event *advanced* |
| `webLink` | Web link | `String` | The URL to open the event in Outlook Web App *advanced* |
| `requiredAttendees` | Required attendees | `String` | Required attendees for the event separated by semicolons *advanced* |
| `optionalAttendees` | Optional attendees | `String` | Optional attendees for the event separated by semicolons *advanced* |
| `resourceAttendees` | Resource attendees | `String` | Resource attendees for the event separated by semicolons *advanced* |
| `location` | Location | `String` | Location of the event *advanced* |
| `importance` | Importance | `String` | The importance of the event: low, normal, or high *advanced* |
| `isAllDay` | Is all day event? | `Boolean` | Set to true if the event lasts all day *advanced* |
| `recurrence` | Recurrence | `String` | The recurrence pattern for the event: none, daily, weekly, monthly or yearly *advanced* |
| `recurrenceEnd` | Recurrence end date | `Date` | End Date of the recurrence *advanced* |
| `numberOfOccurences` | Number of occurrences | `Number` | How many times to repeat the event *advanced* |
| `reminderMinutesBeforeStart` | Reminder | `Number` | Time in minutes before event start to remind *advanced* |
| `isReminderOn` | Is reminder on | `Boolean` | Set to true if an alert is set to remind the user of the event. *advanced* |
| `showAs` | Show as | `String` | Status to show during the event: free, tentative, busy, oof, workingElsewhere or unknown *advanced* |
| `responseRequested` | Response requested | `Boolean` | Set to true if the sender would like a response when the event is accepted or declined *advanced* |
| `sensitivity` | Sensitivity | `String` | The possible values are: normal, personal, private, confidential *advanced* |

### ContactGetTables

**Get contact folders** — This operation lists available contacts folders.

**Inputs**

_None._

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (3 fields) | List of values |

### ContactGetTablesV2

**Get contact folders (V2)** — This operation lists available contacts folders using Graph API

**Inputs**

_None._

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (3 fields) | List of values |

### ContactGetItems

**Get contacts** — This operation gets contacts from a contacts folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Folder id | `String` | ✓ | Unique identifier of the contacts folder to retrieve *values looked up dynamically* |
| `$filter` | Filter Query | `String` |  | An ODATA filter query to restrict the entries returned (e.g. stringColumn eq 'string' OR numberColumn lt 123). *advanced* |
| `$orderby` | Order By | `String` |  | An ODATA orderBy query for specifying the order of entries. *advanced* |
| `$top` | Top Count | `Number` |  | Total number of entries to retrieve (default = all). *advanced* |
| `$skip` | Skip Count | `Number` |  | The number of entries to skip (default = 0). *advanced* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (35 fields) | List of values |

### ContactPostItem

**Create contact** — This operation creates a new contact in a contacts folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Folder id | `String` | ✓ | Select a contacts folder *values looked up dynamically* |
| `GivenName` | Given name | `String` | ✓ | The contact's given name |
| `HomePhones` | Home phones | `Table` (1 fields) | ✓ | The contact's home phone numbers |
| `Id` | Id | `String` |  | The contact's unique identifier. |
| `ParentFolderId` | Parent folder id | `String` |  | The ID of the contact's parent folder *advanced* |
| `Birthday` | Birthday | `DateTime` |  | The contact's birthday *advanced* |
| `FileAs` | File as | `String` |  | The name the contact is filed under *advanced* |
| `DisplayName` | Display Name | `String` |  | The contact's display name |
| `Initials` | Initials | `String` |  | The contact's initials *advanced* |
| `MiddleName` | Middle name | `String` |  | The contact's middle name *advanced* |
| `NickName` | Nickname | `String` |  | The contact's nickname *advanced* |
| `Surname` | Surname | `String` |  | The contact's surname *advanced* |
| `Title` | Title | `String` |  | The contact's title *advanced* |
| `Generation` | Generation | `String` |  | The contact's generation *advanced* |
| `EmailAddresses` | Email addresses | `Table` (2 fields) |  | The contact's email addresses |
| `ImAddresses` | IM addresses | `Table` (1 fields) |  | The contact's instant messaging (IM) addresses *advanced* |
| `JobTitle` | JobTitle | `String` |  | The contact's job title *advanced* |
| `CompanyName` | Company name | `String` |  | The name of the contact's company |
| `Department` | Department | `String` |  | The contact's department *advanced* |
| `OfficeLocation` | Office location | `String` |  | The location of the contact's office *advanced* |
| `Profession` | Profession | `String` |  | The contact's profession *advanced* |
| `BusinessHomePage` | Business home page | `String` |  | The business home page of the contact *advanced* |
| `AssistantName` | Assistant name | `String` |  | The name of the contact's assistant *advanced* |
| `Manager` | Manager | `String` |  | The name of the contact's manager *advanced* |
| `BusinessPhones` | Business phones | `Table` (1 fields) |  | The contact's business phone numbers |
| `MobilePhone1` | Mobile phone | `String` |  | The contact's mobile phone number |
| `HomeAddress` | Item | `Record` (5 fields) |  | Contact to create |
| `BusinessAddress` | Item | `Record` (5 fields) |  | Contact to create |
| `OtherAddress` | Item | `Record` (5 fields) |  | Contact to create |
| `YomiCompanyName` | Yomi company name | `String` |  | The phonetic Japanese company name of the contact *advanced* |
| `YomiGivenName` | Yomi given name | `String` |  | The phonetic Japanese given name (first name) of the contact *advanced* |
| `YomiSurname` | Yomi surname | `String` |  | The phonetic Japanese surname (last name) of the contact *advanced* |
| `Categories` | Categories | `Table` (1 fields) |  | The categories associated with the contact *advanced* |
| `ChangeKey` | Change key | `String` |  | Identifies the version of the event object *advanced* |
| `DateTimeCreated` | Created time | `DateTime` |  | The time the contact was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` |  | The time the contact was modified *advanced* |

**Output** — `Record` (35 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `GivenName` | Given name | `String` | The contact's given name |
| `HomePhones` | Home phones | `Table` (1 fields) | The contact's home phone numbers |
| `Id` | Id | `String` | The contact's unique identifier. |
| `ParentFolderId` | Parent folder id | `String` | The ID of the contact's parent folder *advanced* |
| `Birthday` | Birthday | `DateTime` | The contact's birthday *advanced* |
| `FileAs` | File as | `String` | The name the contact is filed under *advanced* |
| `DisplayName` | Display Name | `String` | The contact's display name |
| `Initials` | Initials | `String` | The contact's initials *advanced* |
| `MiddleName` | Middle name | `String` | The contact's middle name *advanced* |
| `NickName` | Nickname | `String` | The contact's nickname *advanced* |
| `Surname` | Surname | `String` | The contact's surname *advanced* |
| `Title` | Title | `String` | The contact's title *advanced* |
| `Generation` | Generation | `String` | The contact's generation *advanced* |
| `EmailAddresses` | Email addresses | `Table` (2 fields) | The contact's email addresses |
| `ImAddresses` | IM addresses | `Table` (1 fields) | The contact's instant messaging (IM) addresses *advanced* |
| `JobTitle` | JobTitle | `String` | The contact's job title *advanced* |
| `CompanyName` | Company name | `String` | The name of the contact's company |
| `Department` | Department | `String` | The contact's department *advanced* |
| `OfficeLocation` | Office location | `String` | The location of the contact's office *advanced* |
| `Profession` | Profession | `String` | The contact's profession *advanced* |
| `BusinessHomePage` | Business home page | `String` | The business home page of the contact *advanced* |
| `AssistantName` | Assistant name | `String` | The name of the contact's assistant *advanced* |
| `Manager` | Manager | `String` | The name of the contact's manager *advanced* |
| `BusinessPhones` | Business phones | `Table` (1 fields) | The contact's business phone numbers |
| `MobilePhone1` | Mobile phone | `String` | The contact's mobile phone number |
| `HomeAddress` |  | `Record` (5 fields) |  |
| `BusinessAddress` |  | `Record` (5 fields) |  |
| `OtherAddress` |  | `Record` (5 fields) |  |
| `YomiCompanyName` | Yomi company name | `String` | The phonetic Japanese company name of the contact *advanced* |
| `YomiGivenName` | Yomi given name | `String` | The phonetic Japanese given name (first name) of the contact *advanced* |
| `YomiSurname` | Yomi surname | `String` | The phonetic Japanese surname (last name) of the contact *advanced* |
| `Categories` | Categories | `Table` (1 fields) | The categories associated with the contact *advanced* |
| `ChangeKey` | Change key | `String` | Identifies the version of the event object *advanced* |
| `DateTimeCreated` | Created time | `DateTime` | The time the contact was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` | The time the contact was modified *advanced* |

### ContactGetItem

**Get contact** — This operation gets a specific contact from a contacts folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Folder id | `String` | ✓ | Select a contacts folder *values looked up dynamically* |
| `id` | Item id | `String` | ✓ | Unique identifier of a contact to retrieve |

**Output** — `Record` (35 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `GivenName` | Given name | `String` | The contact's given name |
| `HomePhones` | Home phones | `Table` (1 fields) | The contact's home phone numbers |
| `Id` | Id | `String` | The contact's unique identifier. |
| `ParentFolderId` | Parent folder id | `String` | The ID of the contact's parent folder *advanced* |
| `Birthday` | Birthday | `DateTime` | The contact's birthday *advanced* |
| `FileAs` | File as | `String` | The name the contact is filed under *advanced* |
| `DisplayName` | Display Name | `String` | The contact's display name |
| `Initials` | Initials | `String` | The contact's initials *advanced* |
| `MiddleName` | Middle name | `String` | The contact's middle name *advanced* |
| `NickName` | Nickname | `String` | The contact's nickname *advanced* |
| `Surname` | Surname | `String` | The contact's surname *advanced* |
| `Title` | Title | `String` | The contact's title *advanced* |
| `Generation` | Generation | `String` | The contact's generation *advanced* |
| `EmailAddresses` | Email addresses | `Table` (2 fields) | The contact's email addresses |
| `ImAddresses` | IM addresses | `Table` (1 fields) | The contact's instant messaging (IM) addresses *advanced* |
| `JobTitle` | JobTitle | `String` | The contact's job title *advanced* |
| `CompanyName` | Company name | `String` | The name of the contact's company |
| `Department` | Department | `String` | The contact's department *advanced* |
| `OfficeLocation` | Office location | `String` | The location of the contact's office *advanced* |
| `Profession` | Profession | `String` | The contact's profession *advanced* |
| `BusinessHomePage` | Business home page | `String` | The business home page of the contact *advanced* |
| `AssistantName` | Assistant name | `String` | The name of the contact's assistant *advanced* |
| `Manager` | Manager | `String` | The name of the contact's manager *advanced* |
| `BusinessPhones` | Business phones | `Table` (1 fields) | The contact's business phone numbers |
| `MobilePhone1` | Mobile phone | `String` | The contact's mobile phone number |
| `HomeAddress` |  | `Record` (5 fields) |  |
| `BusinessAddress` |  | `Record` (5 fields) |  |
| `OtherAddress` |  | `Record` (5 fields) |  |
| `YomiCompanyName` | Yomi company name | `String` | The phonetic Japanese company name of the contact *advanced* |
| `YomiGivenName` | Yomi given name | `String` | The phonetic Japanese given name (first name) of the contact *advanced* |
| `YomiSurname` | Yomi surname | `String` | The phonetic Japanese surname (last name) of the contact *advanced* |
| `Categories` | Categories | `Table` (1 fields) | The categories associated with the contact *advanced* |
| `ChangeKey` | Change key | `String` | Identifies the version of the event object *advanced* |
| `DateTimeCreated` | Created time | `DateTime` | The time the contact was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` | The time the contact was modified *advanced* |

### ContactDeleteItem

**Delete contact** — This operation deletes a contact from a contacts folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Folder id | `String` | ✓ | Select a contacts folder *values looked up dynamically* |
| `id` | Id | `String` | ✓ | Unique identifier of contact to delete |

**Output** — `String`

### ContactPatchItem

**Update contact** — This operation updates a contact in a contacts folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `table` | Folder id | `String` | ✓ | Select a contacts folder *values looked up dynamically* |
| `id` | Id | `String` | ✓ | Unique identifier of contact to update |
| `GivenName` | Given name | `String` | ✓ | The contact's given name |
| `HomePhones` | Home phones | `Table` (1 fields) | ✓ | The contact's home phone numbers |
| `Id` | Id | `String` |  | The contact's unique identifier. |
| `ParentFolderId` | Parent folder id | `String` |  | The ID of the contact's parent folder *advanced* |
| `Birthday` | Birthday | `DateTime` |  | The contact's birthday *advanced* |
| `FileAs` | File as | `String` |  | The name the contact is filed under *advanced* |
| `DisplayName` | Display Name | `String` |  | The contact's display name |
| `Initials` | Initials | `String` |  | The contact's initials *advanced* |
| `MiddleName` | Middle name | `String` |  | The contact's middle name *advanced* |
| `NickName` | Nickname | `String` |  | The contact's nickname *advanced* |
| `Surname` | Surname | `String` |  | The contact's surname *advanced* |
| `Title` | Title | `String` |  | The contact's title *advanced* |
| `Generation` | Generation | `String` |  | The contact's generation *advanced* |
| `EmailAddresses` | Email addresses | `Table` (2 fields) |  | The contact's email addresses |
| `ImAddresses` | IM addresses | `Table` (1 fields) |  | The contact's instant messaging (IM) addresses *advanced* |
| `JobTitle` | JobTitle | `String` |  | The contact's job title *advanced* |
| `CompanyName` | Company name | `String` |  | The name of the contact's company |
| `Department` | Department | `String` |  | The contact's department *advanced* |
| `OfficeLocation` | Office location | `String` |  | The location of the contact's office *advanced* |
| `Profession` | Profession | `String` |  | The contact's profession *advanced* |
| `BusinessHomePage` | Business home page | `String` |  | The business home page of the contact *advanced* |
| `AssistantName` | Assistant name | `String` |  | The name of the contact's assistant *advanced* |
| `Manager` | Manager | `String` |  | The name of the contact's manager *advanced* |
| `BusinessPhones` | Business phones | `Table` (1 fields) |  | The contact's business phone numbers |
| `MobilePhone1` | Mobile phone | `String` |  | The contact's mobile phone number |
| `HomeAddress` | Item | `Record` (5 fields) |  | Contact item to update |
| `BusinessAddress` | Item | `Record` (5 fields) |  | Contact item to update |
| `OtherAddress` | Item | `Record` (5 fields) |  | Contact item to update |
| `YomiCompanyName` | Yomi company name | `String` |  | The phonetic Japanese company name of the contact *advanced* |
| `YomiGivenName` | Yomi given name | `String` |  | The phonetic Japanese given name (first name) of the contact *advanced* |
| `YomiSurname` | Yomi surname | `String` |  | The phonetic Japanese surname (last name) of the contact *advanced* |
| `Categories` | Categories | `Table` (1 fields) |  | The categories associated with the contact *advanced* |
| `ChangeKey` | Change key | `String` |  | Identifies the version of the event object *advanced* |
| `DateTimeCreated` | Created time | `DateTime` |  | The time the contact was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` |  | The time the contact was modified *advanced* |

**Output** — `Record` (35 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `GivenName` | Given name | `String` | The contact's given name |
| `HomePhones` | Home phones | `Table` (1 fields) | The contact's home phone numbers |
| `Id` | Id | `String` | The contact's unique identifier. |
| `ParentFolderId` | Parent folder id | `String` | The ID of the contact's parent folder *advanced* |
| `Birthday` | Birthday | `DateTime` | The contact's birthday *advanced* |
| `FileAs` | File as | `String` | The name the contact is filed under *advanced* |
| `DisplayName` | Display Name | `String` | The contact's display name |
| `Initials` | Initials | `String` | The contact's initials *advanced* |
| `MiddleName` | Middle name | `String` | The contact's middle name *advanced* |
| `NickName` | Nickname | `String` | The contact's nickname *advanced* |
| `Surname` | Surname | `String` | The contact's surname *advanced* |
| `Title` | Title | `String` | The contact's title *advanced* |
| `Generation` | Generation | `String` | The contact's generation *advanced* |
| `EmailAddresses` | Email addresses | `Table` (2 fields) | The contact's email addresses |
| `ImAddresses` | IM addresses | `Table` (1 fields) | The contact's instant messaging (IM) addresses *advanced* |
| `JobTitle` | JobTitle | `String` | The contact's job title *advanced* |
| `CompanyName` | Company name | `String` | The name of the contact's company |
| `Department` | Department | `String` | The contact's department *advanced* |
| `OfficeLocation` | Office location | `String` | The location of the contact's office *advanced* |
| `Profession` | Profession | `String` | The contact's profession *advanced* |
| `BusinessHomePage` | Business home page | `String` | The business home page of the contact *advanced* |
| `AssistantName` | Assistant name | `String` | The name of the contact's assistant *advanced* |
| `Manager` | Manager | `String` | The name of the contact's manager *advanced* |
| `BusinessPhones` | Business phones | `Table` (1 fields) | The contact's business phone numbers |
| `MobilePhone1` | Mobile phone | `String` | The contact's mobile phone number |
| `HomeAddress` |  | `Record` (5 fields) |  |
| `BusinessAddress` |  | `Record` (5 fields) |  |
| `OtherAddress` |  | `Record` (5 fields) |  |
| `YomiCompanyName` | Yomi company name | `String` | The phonetic Japanese company name of the contact *advanced* |
| `YomiGivenName` | Yomi given name | `String` | The phonetic Japanese given name (first name) of the contact *advanced* |
| `YomiSurname` | Yomi surname | `String` | The phonetic Japanese surname (last name) of the contact *advanced* |
| `Categories` | Categories | `Table` (1 fields) | The categories associated with the contact *advanced* |
| `ChangeKey` | Change key | `String` | Identifies the version of the event object *advanced* |
| `DateTimeCreated` | Created time | `DateTime` | The time the contact was created *advanced* |
| `DateTimeLastModified` | Last modified time | `DateTime` | The time the contact was modified *advanced* |

### ExportEmail

**Export email** — Export the content of the email.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to export. |

**Output** — `File`

### ExportEmail_V2

**Export email (V2)** — Export the content of the email in the EML file format.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to export. |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to export from. |

**Output** — `File`

### Flag_V2

**Flag email (V2)** — This operation updates an email flag.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to be flagged. |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to update mail. |
| `flag` | Flag Status | `Record` (1 fields) |  | Flag status |

**Output** — `String`

### MarkAsRead_V2

**Mark as read or unread (V2)** — This operation marks an email as read/unread.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to be marked. |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to update mail. |
| `isRead` | Mark as | `Boolean` |  | Mark as read/unread. |

**Output** — `String`

### MarkAsRead_V3

**Mark as read or unread (V3)** — This operation marks an email as read/unread.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to be marked. |
| `isRead` | Mark as | `Boolean` | ✓ | Mark as read/unread. |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to update mail. |

**Output** — `String`

### DeleteEmail_V2

**Delete email (V2)** — This operation deletes an email by id.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email to delete. |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to delete mail from. |

**Output** — `String`

### GetAttachment_V2

**Get Attachment (V2)** — This operation gets an email attachment by id.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `messageId` | Message Id | `String` | ✓ | Id of the email. |
| `attachmentId` | Attachement Id | `String` | ✓ | Id of the attachment to download. |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to retrieve attachment from. |
| `extractSensitivityLabel` | Extract Sensitivity Label | `Boolean` |  | Select if you want to extract Sensitivity label ( false, true). |
| `fetchSensitivityLabelMetadata` | Sensitivity Label Metadata | `Boolean` |  | A boolean whether to fetch sensitivity label Metadata for associated LabelId. |

**Output** — `Record` (9 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `id` | Id | `String` | Id of the attachment. |
| `name` | Name | `String` | Name of attachment. |
| `contentType` | Content Type | `String` | Content type of attachment. |
| `size` | Size | `Number` | Size of attachment. |
| `contentBytes` | Content Bytes | `File` | Content of attachment. |
| `isInline` | Is Inline | `Boolean` | Set to true if this is an inline attachment. |
| `lastModifiedDateTime` | Last Modified DateTime | `DateTime` | The date and time when the attachment was last modified. |
| `contentId` | Content Id | `String` | Content Id |
| `sensitivityLabelInfo` |  | `Table` (10 fields) |  |

### RespondToEvent

**Respond to an event invite** — Respond to an event invite.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `event_id` | Event Id | `String` | ✓ | Id of the event to respond to. |
| `response` | Response | `String` | ✓ | Response for the event invite |
| `Comment` | Comment | `String` |  | Comment |
| `SendResponse` | Send response? | `Boolean` |  | Send response to organizer? |

**Output** — `String`

### RespondToEvent_V2

**Respond to an event invite (V2)** — Respond to an event invite.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `event_id` | Event Id | `String` | ✓ | Id of the event to respond to. |
| `response` | Response | `String` | ✓ | Response for the event invite |
| `Comment` | Comment | `String` |  | Comment |
| `SendResponse` | Send response? | `Boolean` |  | Send response to organizer? |

**Output** — `String`

### ForwardEmail

**Forward an email** — Forward an email.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `message_id` | Message Id | `String` | ✓ | Id of the message to forward. |
| `ToRecipients` | To | `String` | ✓ | Semicolon separated list of recipients to forward the message to |
| `Comment` | Comment | `String` |  | Comment |

**Output** — `String`

### ForwardEmail_V2

**Forward an email (V2)** — Forward an email.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `message_id` | Message Id | `String` | ✓ | Id of the message to forward. (You can use 'Message Id' from trigger or 'Get Emails' action output) |
| `ToRecipients` | To | `String` | ✓ | Semicolon separated list of recipients to forward the message to |
| `mailboxAddress` | Original Mailbox Address | `String` |  | Address of the shared mailbox to forward mail from. |
| `Comment` | Comment | `String` |  | Comment |

**Output** — `String`

### GetRoomLists

**Get room lists** — Get all the room lists defined in the user's tenant

**Inputs**

_None._

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | value | `Table` (2 fields) | value |

### GetRoomLists_V2

**Get room lists (V2)** — Get all the room lists defined in the user's tenant

**Inputs**

_None._

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | value | `Table` (2 fields) | value |

### GetRooms

**Get rooms** — Get all the meeting rooms defined in the user's tenant

**Inputs**

_None._

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | value | `Table` (2 fields) | value |

### GetRooms_V2

**Get rooms (V2)** — Get all the meeting rooms defined in the user's tenant

**Inputs**

_None._

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | value | `Table` (2 fields) | value |

### GetRoomsInRoomList

**Get rooms in room list** — Get the meeting rooms in a specific room list

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `room_list` | Room list | `String` | ✓ | Room list to find rooms in *values looked up dynamically* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | value | `Table` (2 fields) | value |

### GetRoomsInRoomList_V2

**Get rooms in room list (V2)** — Get the meeting rooms in a specific room list

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `room_list` | Room list | `String` | ✓ | Room list to find rooms in *values looked up dynamically* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | value | `Table` (2 fields) | value |

### FindMeetingTimes

**Find meeting times** — Find meeting time suggestions based on organizer, attendee availability, and time or location constraints

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `RequiredAttendees` | Required attendees | `String` |  | List of semicolon separated email addresses |
| `OptionalAttendees` | Optional attendees | `String` |  | List of semicolon separated email addresses |
| `ResourceAttendees` | Resource attendees | `String` |  | Resource attendees for the event separated by semicolons |
| `MeetingDuration` | Meeting duration | `Number` |  | Duration of the meeting in minutes |
| `Start` | Start time | `DateTime` |  | Start time for meeting time suggestions |
| `End` | End time | `DateTime` |  | End time for meeting time suggestions |
| `MaxCandidates` | Max Candidates | `Number` |  | The maximum number of meeting suggestions to return in the response *advanced* |
| `MinimumAttendeePercentage` | Minimum Attendee Percentage | `String` |  | The minimum required confidence for a time slot to be returned in the response *advanced* |
| `IsOrganizerOptional` | Is Organizer Optional? | `Boolean` |  | true if the organizer doesn't have to attend. The default is false *advanced* |
| `ActivityDomain` | Activity Domain | `String` |  | Work, Personal, Unrestricted, or Unknown *advanced* |

**Output** — `Record` (2 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `EmptySuggestionsReason` | Empty Suggestions Reason | `String` | Empty Suggestions Reason |
| `MeetingTimeSuggestions` | Meeting Time Suggestions | `Table` (6 fields) | Meeting Time Suggestions |

### FindMeetingTimes_V2

**Find meeting times (V2)** — Find meeting time suggestions based on organizer, attendee availability, and time or location constraints

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `RequiredAttendees` | Required attendees | `String` |  | List of semicolon separated email addresses |
| `OptionalAttendees` | Optional attendees | `String` |  | List of semicolon separated email addresses |
| `ResourceAttendees` | Resource attendees | `String` |  | Resource attendees for the event separated by semicolons |
| `MeetingDuration` | Meeting duration | `Number` |  | Duration of the meeting in minutes |
| `Start` | Start time | `DateTime` |  | Start time for meeting time suggestions |
| `End` | End time | `DateTime` |  | End time for meeting time suggestions |
| `MaxCandidates` | Max Candidates | `Number` |  | The maximum number of meeting suggestions to return in the response *advanced* |
| `MinimumAttendeePercentage` | Minimum Attendee Percentage | `String` |  | The minimum required confidence for a time slot to be returned in the response *advanced* |
| `IsOrganizerOptional` | Is Organizer Optional? | `Boolean` |  | true if the organizer doesn't have to attend. The default is false *advanced* |
| `ActivityDomain` | Activity Domain | `String` |  | Work, Personal, Unrestricted, or Unknown *advanced* |

**Output** — `Record` (2 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `emptySuggestionsReason` | Empty Suggestions Reason | `String` | Empty Suggestions Reason |
| `meetingTimeSuggestions` | Meeting Time Suggestions | `Table` (6 fields) | Meeting Time Suggestions |

### SetAutomaticRepliesSetting_V2

**Set up automatic replies (V2)** — Set the automatic replies setting for your mailbox.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `automaticRepliesSetting` | Settings | `Record` (6 fields) |  | Automatic replies setting model for the connector |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `automaticRepliesSetting` |  | `Record` (6 fields) | Automatic replies setting model for the connector |

### GetMailTips_V2

**Get mail tips for a mailbox (V2)** — Get mail tips for a mailbox such as automatic replies / OOF message or if the mailbox is full. This is not available in GccHigh and Mooncake.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `EmailAddresses` | Email Addresses | `Table` (1 fields) | ✓ | Address of the mailbox to get mail tips for. |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (7 fields) |  |

### CalendarGetTables_V2

**Get calendars (V2)** — This operation lists available calendars.

**Inputs**

_None._

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | value | `Table` (3 fields) | value |

### CalendarDeleteItem_V2

**Delete event (V2)** — This operation deletes an event in a calendar.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `calendar` | Calendar id | `String` | ✓ | Select a calendar *values looked up dynamically* |
| `event` | Id | `String` | ✓ | Select an event |

**Output** — `String`

### ContactGetItem_V2

**Get contact (V2)** — This operation gets a specific contact from a contacts folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `folder` | Folder id | `String` | ✓ | Select a contacts folder *values looked up dynamically* |
| `id` | Item id | `String` | ✓ | Unique identifier of a contact to retrieve |

**Output** — `Record` (35 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `id` | Id | `String` | The contact's unique identifier. |
| `parentFolderId` | Parent folder id | `String` | The ID of the contact's parent folder *advanced* |
| `birthday` | Birthday | `DateTime` | The contact's birthday *advanced* |
| `fileAs` | File as | `String` | The name the contact is filed under *advanced* |
| `displayName` | Display Name | `String` | The contact's display name |
| `givenName` | Given name | `String` | The contact's given name |
| `initials` | Initials | `String` | The contact's initials *advanced* |
| `middleName` | Middle name | `String` | The contact's middle name *advanced* |
| `nickName` | Nickname | `String` | The contact's nickname *advanced* |
| `surname` | Surname | `String` | The contact's surname *advanced* |
| `title` | Title | `String` | The contact's title *advanced* |
| `generation` | Generation | `String` | The contact's generation *advanced* |
| `emailAddresses` | Email addresses | `Table` (2 fields) | The contact's email addresses |
| `imAddresses` | IM addresses | `Table` (1 fields) | The contact's instant messaging (IM) addresses *advanced* |
| `jobTitle` | JobTitle | `String` | The contact's job title *advanced* |
| `companyName` | Company name | `String` | The name of the contact's company |
| `department` | Department | `String` | The contact's department *advanced* |
| `officeLocation` | Office location | `String` | The location of the contact's office *advanced* |
| `profession` | Profession | `String` | The contact's profession *advanced* |
| `businessHomePage` | Business home page | `String` | The business home page of the contact *advanced* |
| `assistantName` | Assistant name | `String` | The name of the contact's assistant *advanced* |
| `manager` | Manager | `String` | The name of the contact's manager *advanced* |
| `homePhones` | Home phones | `Table` (1 fields) | The contact's home phone numbers |
| `businessPhones` | Business phones | `Table` (1 fields) | The contact's business phone numbers |
| `mobilePhone` | Mobile phone | `String` | The contact's mobile phone number |
| `homeAddress` |  | `Record` (5 fields) |  |
| `businessAddress` |  | `Record` (5 fields) |  |
| `otherAddress` |  | `Record` (5 fields) |  |
| `yomiCompanyName` | Yomi company name | `String` | The phonetic Japanese company name of the contact *advanced* |
| `yomiGivenName` | Yomi given name | `String` | The phonetic Japanese given name (first name) of the contact *advanced* |
| `yomiSurname` | Yomi surname | `String` | The phonetic Japanese surname (last name) of the contact *advanced* |
| `categories` | Categories | `Table` (1 fields) | The categories associated with the contact *advanced* |
| `changeKey` | Change key | `String` | Identifies the version of the event object *advanced* |
| `createdDateTime` | Created time | `DateTime` | The time the contact was created *advanced* |
| `lastModifiedDateTime` | Last modified time | `DateTime` | The time the contact was modified *advanced* |

### ContactDeleteItem_V2

**Delete contact (V2)** — This operation deletes a contact from a contacts folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `folder` | Folder id | `String` | ✓ | Select a contacts folder *values looked up dynamically* |
| `id` | Id | `String` | ✓ | Unique identifier of contact to delete |

**Output** — `String`

### ContactPatchItem_V2

**Update contact (V2)** — This operation updates a contact in a contacts folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `folder` | Folder id | `String` | ✓ | Select a contacts folder *values looked up dynamically* |
| `id` | Id | `String` | ✓ | Unique identifier of contact to update |
| `givenName` | Given name | `String` | ✓ | The contact's given name |
| `homePhones` | Home phones | `Table` (1 fields) | ✓ | The contact's home phone numbers |
| `id_1` | Id | `String` |  | The contact's unique identifier. |
| `parentFolderId` | Parent folder id | `String` |  | The ID of the contact's parent folder *advanced* |
| `birthday` | Birthday | `DateTime` |  | The contact's birthday *advanced* |
| `fileAs` | File as | `String` |  | The name the contact is filed under *advanced* |
| `displayName` | Display Name | `String` |  | The contact's display name |
| `initials` | Initials | `String` |  | The contact's initials *advanced* |
| `middleName` | Middle name | `String` |  | The contact's middle name *advanced* |
| `nickName` | Nickname | `String` |  | The contact's nickname *advanced* |
| `surname` | Surname | `String` |  | The contact's surname *advanced* |
| `title` | Title | `String` |  | The contact's title *advanced* |
| `generation` | Generation | `String` |  | The contact's generation *advanced* |
| `emailAddresses` | Email addresses | `Table` (2 fields) |  | The contact's email addresses |
| `imAddresses` | IM addresses | `Table` (1 fields) |  | The contact's instant messaging (IM) addresses *advanced* |
| `jobTitle` | JobTitle | `String` |  | The contact's job title *advanced* |
| `companyName` | Company name | `String` |  | The name of the contact's company |
| `department` | Department | `String` |  | The contact's department *advanced* |
| `officeLocation` | Office location | `String` |  | The location of the contact's office *advanced* |
| `profession` | Profession | `String` |  | The contact's profession *advanced* |
| `businessHomePage` | Business home page | `String` |  | The business home page of the contact *advanced* |
| `assistantName` | Assistant name | `String` |  | The name of the contact's assistant *advanced* |
| `manager` | Manager | `String` |  | The name of the contact's manager *advanced* |
| `businessPhones` | Business phones | `Table` (1 fields) |  | The contact's business phone numbers |
| `mobilePhone` | Mobile phone | `String` |  | The contact's mobile phone number |
| `homeAddress` | Item | `Record` (5 fields) |  | Contact item to update |
| `businessAddress` | Item | `Record` (5 fields) |  | Contact item to update |
| `otherAddress` | Item | `Record` (5 fields) |  | Contact item to update |
| `yomiCompanyName` | Yomi company name | `String` |  | The phonetic Japanese company name of the contact *advanced* |
| `yomiGivenName` | Yomi given name | `String` |  | The phonetic Japanese given name (first name) of the contact *advanced* |
| `yomiSurname` | Yomi surname | `String` |  | The phonetic Japanese surname (last name) of the contact *advanced* |
| `categories` | Categories | `Table` (1 fields) |  | The categories associated with the contact *advanced* |
| `changeKey` | Change key | `String` |  | Identifies the version of the event object *advanced* |
| `createdDateTime` | Created time | `DateTime` |  | The time the contact was created *advanced* |
| `lastModifiedDateTime` | Last modified time | `DateTime` |  | The time the contact was modified *advanced* |

**Output** — `Record` (35 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `id` | Id | `String` | The contact's unique identifier. |
| `parentFolderId` | Parent folder id | `String` | The ID of the contact's parent folder *advanced* |
| `birthday` | Birthday | `DateTime` | The contact's birthday *advanced* |
| `fileAs` | File as | `String` | The name the contact is filed under *advanced* |
| `displayName` | Display Name | `String` | The contact's display name |
| `givenName` | Given name | `String` | The contact's given name |
| `initials` | Initials | `String` | The contact's initials *advanced* |
| `middleName` | Middle name | `String` | The contact's middle name *advanced* |
| `nickName` | Nickname | `String` | The contact's nickname *advanced* |
| `surname` | Surname | `String` | The contact's surname *advanced* |
| `title` | Title | `String` | The contact's title *advanced* |
| `generation` | Generation | `String` | The contact's generation *advanced* |
| `emailAddresses` | Email addresses | `Table` (2 fields) | The contact's email addresses |
| `imAddresses` | IM addresses | `Table` (1 fields) | The contact's instant messaging (IM) addresses *advanced* |
| `jobTitle` | JobTitle | `String` | The contact's job title *advanced* |
| `companyName` | Company name | `String` | The name of the contact's company |
| `department` | Department | `String` | The contact's department *advanced* |
| `officeLocation` | Office location | `String` | The location of the contact's office *advanced* |
| `profession` | Profession | `String` | The contact's profession *advanced* |
| `businessHomePage` | Business home page | `String` | The business home page of the contact *advanced* |
| `assistantName` | Assistant name | `String` | The name of the contact's assistant *advanced* |
| `manager` | Manager | `String` | The name of the contact's manager *advanced* |
| `homePhones` | Home phones | `Table` (1 fields) | The contact's home phone numbers |
| `businessPhones` | Business phones | `Table` (1 fields) | The contact's business phone numbers |
| `mobilePhone` | Mobile phone | `String` | The contact's mobile phone number |
| `homeAddress` |  | `Record` (5 fields) |  |
| `businessAddress` |  | `Record` (5 fields) |  |
| `otherAddress` |  | `Record` (5 fields) |  |
| `yomiCompanyName` | Yomi company name | `String` | The phonetic Japanese company name of the contact *advanced* |
| `yomiGivenName` | Yomi given name | `String` | The phonetic Japanese given name (first name) of the contact *advanced* |
| `yomiSurname` | Yomi surname | `String` | The phonetic Japanese surname (last name) of the contact *advanced* |
| `categories` | Categories | `Table` (1 fields) | The categories associated with the contact *advanced* |
| `changeKey` | Change key | `String` | Identifies the version of the event object *advanced* |
| `createdDateTime` | Created time | `DateTime` | The time the contact was created *advanced* |
| `lastModifiedDateTime` | Last modified time | `DateTime` | The time the contact was modified *advanced* |

### ContactGetItems_V2

**Get contacts (V2)** — This operation gets contacts from a contacts folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `folder` | Folder id | `String` | ✓ | Unique identifier of the contacts folder to retrieve *values looked up dynamically* |
| `$filter` | Filter Query | `String` |  | An ODATA filter query to restrict the entries returned (e.g. stringColumn eq 'string' OR numberColumn lt 123). *advanced* |
| `$orderby` | Order By | `String` |  | An ODATA orderBy query for specifying the order of entries. *advanced* |
| `$top` | Top Count | `Number` |  | Total number of entries to retrieve (default = all). *advanced* |
| `$skip` | Skip Count | `Number` |  | The number of entries to skip (default = 0). *advanced* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (35 fields) | List of values |

### ContactPostItem_V2

**Create contact (V2)** — This operation creates a new contact in a contacts folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `folder` | Folder id | `String` | ✓ | Select a contacts folder *values looked up dynamically* |
| `givenName` | Given name | `String` | ✓ | The contact's given name |
| `homePhones` | Home phones | `Table` (1 fields) | ✓ | The contact's home phone numbers |
| `id` | Id | `String` |  | The contact's unique identifier. |
| `parentFolderId` | Parent folder id | `String` |  | The ID of the contact's parent folder *advanced* |
| `birthday` | Birthday | `DateTime` |  | The contact's birthday *advanced* |
| `fileAs` | File as | `String` |  | The name the contact is filed under *advanced* |
| `displayName` | Display Name | `String` |  | The contact's display name |
| `initials` | Initials | `String` |  | The contact's initials *advanced* |
| `middleName` | Middle name | `String` |  | The contact's middle name *advanced* |
| `nickName` | Nickname | `String` |  | The contact's nickname *advanced* |
| `surname` | Surname | `String` |  | The contact's surname *advanced* |
| `title` | Title | `String` |  | The contact's title *advanced* |
| `generation` | Generation | `String` |  | The contact's generation *advanced* |
| `emailAddresses` | Email addresses | `Table` (2 fields) |  | The contact's email addresses |
| `imAddresses` | IM addresses | `Table` (1 fields) |  | The contact's instant messaging (IM) addresses *advanced* |
| `jobTitle` | JobTitle | `String` |  | The contact's job title *advanced* |
| `companyName` | Company name | `String` |  | The name of the contact's company |
| `department` | Department | `String` |  | The contact's department *advanced* |
| `officeLocation` | Office location | `String` |  | The location of the contact's office *advanced* |
| `profession` | Profession | `String` |  | The contact's profession *advanced* |
| `businessHomePage` | Business home page | `String` |  | The business home page of the contact *advanced* |
| `assistantName` | Assistant name | `String` |  | The name of the contact's assistant *advanced* |
| `manager` | Manager | `String` |  | The name of the contact's manager *advanced* |
| `businessPhones` | Business phones | `Table` (1 fields) |  | The contact's business phone numbers |
| `mobilePhone` | Mobile phone | `String` |  | The contact's mobile phone number |
| `homeAddress` | Item | `Record` (5 fields) |  | Contact to create |
| `businessAddress` | Item | `Record` (5 fields) |  | Contact to create |
| `otherAddress` | Item | `Record` (5 fields) |  | Contact to create |
| `yomiCompanyName` | Yomi company name | `String` |  | The phonetic Japanese company name of the contact *advanced* |
| `yomiGivenName` | Yomi given name | `String` |  | The phonetic Japanese given name (first name) of the contact *advanced* |
| `yomiSurname` | Yomi surname | `String` |  | The phonetic Japanese surname (last name) of the contact *advanced* |
| `categories` | Categories | `Table` (1 fields) |  | The categories associated with the contact *advanced* |
| `changeKey` | Change key | `String` |  | Identifies the version of the event object *advanced* |
| `createdDateTime` | Created time | `DateTime` |  | The time the contact was created *advanced* |
| `lastModifiedDateTime` | Last modified time | `DateTime` |  | The time the contact was modified *advanced* |

**Output** — `Record` (35 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `id` | Id | `String` | The contact's unique identifier. |
| `parentFolderId` | Parent folder id | `String` | The ID of the contact's parent folder *advanced* |
| `birthday` | Birthday | `DateTime` | The contact's birthday *advanced* |
| `fileAs` | File as | `String` | The name the contact is filed under *advanced* |
| `displayName` | Display Name | `String` | The contact's display name |
| `givenName` | Given name | `String` | The contact's given name |
| `initials` | Initials | `String` | The contact's initials *advanced* |
| `middleName` | Middle name | `String` | The contact's middle name *advanced* |
| `nickName` | Nickname | `String` | The contact's nickname *advanced* |
| `surname` | Surname | `String` | The contact's surname *advanced* |
| `title` | Title | `String` | The contact's title *advanced* |
| `generation` | Generation | `String` | The contact's generation *advanced* |
| `emailAddresses` | Email addresses | `Table` (2 fields) | The contact's email addresses |
| `imAddresses` | IM addresses | `Table` (1 fields) | The contact's instant messaging (IM) addresses *advanced* |
| `jobTitle` | JobTitle | `String` | The contact's job title *advanced* |
| `companyName` | Company name | `String` | The name of the contact's company |
| `department` | Department | `String` | The contact's department *advanced* |
| `officeLocation` | Office location | `String` | The location of the contact's office *advanced* |
| `profession` | Profession | `String` | The contact's profession *advanced* |
| `businessHomePage` | Business home page | `String` | The business home page of the contact *advanced* |
| `assistantName` | Assistant name | `String` | The name of the contact's assistant *advanced* |
| `manager` | Manager | `String` | The name of the contact's manager *advanced* |
| `homePhones` | Home phones | `Table` (1 fields) | The contact's home phone numbers |
| `businessPhones` | Business phones | `Table` (1 fields) | The contact's business phone numbers |
| `mobilePhone` | Mobile phone | `String` | The contact's mobile phone number |
| `homeAddress` |  | `Record` (5 fields) |  |
| `businessAddress` |  | `Record` (5 fields) |  |
| `otherAddress` |  | `Record` (5 fields) |  |
| `yomiCompanyName` | Yomi company name | `String` | The phonetic Japanese company name of the contact *advanced* |
| `yomiGivenName` | Yomi given name | `String` | The phonetic Japanese given name (first name) of the contact *advanced* |
| `yomiSurname` | Yomi surname | `String` | The phonetic Japanese surname (last name) of the contact *advanced* |
| `categories` | Categories | `Table` (1 fields) | The categories associated with the contact *advanced* |
| `changeKey` | Change key | `String` | Identifies the version of the event object *advanced* |
| `createdDateTime` | Created time | `DateTime` | The time the contact was created *advanced* |
| `lastModifiedDateTime` | Last modified time | `DateTime` | The time the contact was modified *advanced* |

### UpdateMyContactPhoto

**Update my contact's photo** — Updates the photo of the specified contact of the current user. The size of the photo must be less than 4 MB.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `folder` | Folder id | `String` | ✓ | Select a contacts folder *values looked up dynamically* |
| `id` | Item id | `String` | ✓ | Unique identifier of a contact to update photo |
| `file` | Image content | `File` | ✓ | Image content |

**Output** — `String`

### HttpRequest

**Send an HTTP request** — Construct a Microsoft Graph REST API request to invoke. These segments are supported: 1st segement: /me, /users/<userId> 2nd segment: messages, mailFolders, events, calendar, calendars, outlook, inferenceClassification. Learn more: https://docs.microsoft.com/en-us/graph/use-the-api.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `Uri` | URI | `String` | ✓ | The full or relative URI. Example: https://graph.microsoft.com/{version}/{resource}. |
| `Method` | Method | `String` | ✓ | The HTTP method (default is GET). |
| `ContentType` | Content-Type | `String` |  | The content-type header for the body (default is application/json). |
| `CustomHeader1` | CustomHeader1 | `String` |  | Custom header 1. Specify in format: header-name: header-value *advanced* |
| `CustomHeader2` | CustomHeader2 | `String` |  | Custom header 2. Specify in format: header-name: header-value *advanced* |
| `CustomHeader3` | CustomHeader3 | `String` |  | Custom header 3. Specify in format: header-name: header-value *advanced* |
| `CustomHeader4` | CustomHeader4 | `String` |  | Custom header 4. Specify in format: header-name: header-value *advanced* |
| `CustomHeader5` | CustomHeader5 | `String` |  | Custom header 5. Specify in format: header-name: header-value *advanced* |
| `file` | Body | `File` |  | The request body content. |

**Output** — `Any`

### mcp_EmailsManagement

**Email Management MCP Server** — This MCP server manages email messages from your Office 365 account

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `sessionId` |  | `String` |  |  |
| `jsonrpc` |  | `String` |  |  |
| `id` |  | `String` |  |  |
| `method` |  | `String` |  |  |
| `params` |  | `Any` |  |  |
| `result` |  | `Any` |  |  |
| `error` |  | `Any` |  |  |

**Output** — `Record` (6 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `jsonrpc` |  | `String` |  |
| `id` |  | `String` |  |
| `method` |  | `String` |  |
| `params` |  | `Any` |  |
| `result` |  | `Any` |  |
| `error` |  | `Any` |  |

### mcp_MeetingManagement

**Meeting Management MCP Server** — This MCP server manages events, calendars and meetings

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `sessionId` |  | `String` |  |  |
| `jsonrpc` |  | `String` |  |  |
| `id` |  | `String` |  |  |
| `method` |  | `String` |  |  |
| `params` |  | `Any` |  |  |
| `result` |  | `Any` |  |  |
| `error` |  | `Any` |  |  |

**Output** — `Record` (6 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `jsonrpc` |  | `String` |  |
| `id` |  | `String` |  |
| `method` |  | `String` |  |
| `params` |  | `Any` |  |
| `result` |  | `Any` |  |
| `error` |  | `Any` |  |

### mcp_ContactsManagement

**Contact Management MCP Server** — This MCP server manages contacts

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `sessionId` |  | `String` |  |  |
| `jsonrpc` |  | `String` |  |  |
| `id` |  | `String` |  |  |
| `method` |  | `String` |  |  |
| `params` |  | `Any` |  |  |
| `result` |  | `Any` |  |  |
| `error` |  | `Any` |  |  |

**Output** — `Record` (6 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `jsonrpc` |  | `String` |  |
| `id` |  | `String` |  |
| `method` |  | `String` |  |
| `params` |  | `Any` |  |
| `result` |  | `Any` |  |
| `error` |  | `Any` |  |

## Triggers

Connector triggers start a flow or an external-trigger when something happens in the connected service. See [10 Flows & autonomous triggers](../../10-flows-and-autonomous-triggers.md).

| operationId | Name | Description |
|---|---|---|
| `OnUpcomingEvents` | When an upcoming event is starting soon | This operation triggers a flow when an upcoming calendar event is starting. |
| `OnUpcomingEventsV2` | When an upcoming event is starting soon (V2) | This operation triggers a flow when an upcoming calendar event is starting. |
| `OnUpcomingEventsV3` | When an upcoming event is starting soon (V3) | This operation triggers a flow when an upcoming calendar event is starting. |
| `OnNewEmail` | When a new email arrives | This operation triggers a flow when a new email arrives. It will skip any email that has a total message size greater than the limit put by your Exchange Admin or 50 MB, whichever is less. It may also skip protected emails and emails with invalid body or attachments. |
| `OnNewEmailV2` | When a new email arrives (V2) | This operation triggers a flow when a new email arrives. It will skip any email that has a total message size greater than the limit put by your Exchange Admin or 50 MB, whichever is less. It may also skip protected emails and emails with invalid body or attachments. |
| `OnNewEmailV3` | When a new email arrives (V3) | This operation triggers a flow when a new email arrives. It will skip any email that has a total message size greater than the limit put by your Exchange Admin or 50 MB, whichever is less. It may also skip protected emails and emails with invalid body or attachments. |
| `OnFlaggedEmail` | When an email is flagged | This operation triggers a flow when an email is flagged. |
| `OnFlaggedEmailV2` | When an email is flagged (V2) | This operation triggers a flow when an email is flagged. |
| `OnFlaggedEmailV3` | When an email is flagged (V3) | This operation triggers a flow when an email is flagged. |
| `OnFlaggedEmailV4` | When an email is flagged (V4) | This operation triggers a flow when an email is flagged. |
| `OnNewMentionMeEmail` | When a new email mentioning me arrives | This operation triggers a flow when a new email mentioning me arrives. It will skip any email that has a total message size greater than the limit put by your Exchange Admin or 50 MB, whichever is less. It may also skip protected emails and emails with invalid body or attachments. |
| `OnNewMentionMeEmailV2` | When a new email mentioning me arrives (V2) | This operation triggers a flow when a new email mentioning me arrives. It will skip any email that has a total message size greater than the limit put by your Exchange Admin or 50 MB, whichever is less. It may also skip protected emails and emails with invalid body or attachments. |
| `OnNewMentionMeEmailV3` | When a new email mentioning me arrives (V3) | This operation triggers a flow when a new email mentioning me arrives. It will skip any email that has a total message size greater than the limit put by your Exchange Admin or 50 MB, whichever is less. It may also skip protected emails and emails with invalid body or attachments. |
| `CreateOnNewEmailSubscription` | When a new email arrives (webhook) | This operation triggers a flow when a new email arrives. |
| `SharedMailboxOnNewEmail` | When a new email arrives in a shared mailbox | This operation triggers a flow when a new email arrives in a shared mailbox. Your account should have permission to access the mailbox for this operation to succeed. It will skip any email that has a total message size greater than the limit put by your Exchange Admin or 50 MB, whichever is less. It may also skip protected emails and emails with invalid body or attachments. |
| `SharedMailboxOnNewEmailV2` | When a new email arrives in a shared mailbox (V2) | This operation triggers a flow when a new email arrives in a shared mailbox. Your account should have permission to access the mailbox for this operation to succeed. It will skip any email that has a total message size greater than the limit put by your Exchange Admin or 50 MB, whichever is less. It may also skip protected emails and emails with invalid body or attachments. |
| `CalendarGetOnNewItems` | When a new event is created (V1) | This operation triggers a flow when a new event is created in a calendar. (V1) |
| `CalendarGetOnUpdatedItems` | When an event is modified (V1) | This operation triggers a flow when an event is modified in a calendar. (V1) |
| `CalendarGetOnNewItemsV2` | When a new event is created (V2) | This operation triggers a flow when a new event is created in a calendar. (V2) |
| `CalendarGetOnNewItemsV3` | When a new event is created (V3) | This operation triggers a flow when a new event is created in a calendar. (V3) |
| `CalendarGetOnUpdatedItemsV2` | When an event is modified (V2) | This operation triggers a flow when an event is modified in a calendar. (V2) |
| `CalendarGetOnUpdatedItemsV3` | When an event is modified (V3) | This operation triggers a flow when an event is modified in a calendar. (V3) |
| `CalendarGetOnChangedItems` | When an event is added, updated or deleted | This operation triggers a flow when an event is added, updated or deleted in a calendar. |
| `CalendarGetOnChangedItemsV2` | When an event is added, updated or deleted (V2) | This operation triggers a flow when an event is added, updated or deleted in a calendar. (V2) |
| `CalendarGetOnChangedItemsV3` | When an event is added, updated or deleted (V3) | This operation triggers a flow when an event is added, updated or deleted in a calendar. (V3) This is not available in Mooncake. |
