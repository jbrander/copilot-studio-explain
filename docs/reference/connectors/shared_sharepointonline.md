# Connector: SharePoint (`shared_sharepointonline`)

> **Generated** by `tools/gen_reference.py` from `reference/connectors/shared_sharepointonline.yml` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

SharePoint helps organizations share and collaborate with colleagues, partners, and customers. You can connect to SharePoint Online or to an on-premises SharePoint 2016 or 2019 farm using the On-Premises Data Gateway to manage documents and list items.

| | |
|---|---|
| Connector id | `/providers/Microsoft.PowerApps/apis/shared_sharepointonline` |
| Type | Public |
| Operations | 53 |
| Triggers (used by flows / external triggers, not by tools) | 10 |

Use an operation's **operationId** as `action.operationId` in a `TaskDialog` tool (see [08 Tools & connectors](../../08-tools-and-connectors.md)). Inputs map to `inputs:` entries (`propertyName` = the input *Name* below). Input names that start with `$` are [OData](../../glossary.md#odata) query options and need special quoting — see [`$`-prefixed inputs](../../08-tools-and-connectors.md#-prefixed-inputs-odata-parameters).

## Operations

| operationId | Name | Description |
|---|---|---|
| [`GetAllTables`](#getalltables) | Get all lists and libraries | Get all lists and libraries. |
| [`ApproveHubSiteJoin`](#approvehubsitejoin) | Approve hub site join request | Approve hub site join request. This will return an approval token that can be used to complete the join request using the join hub site action. |
| [`CancelHubSiteJoinApproval`](#cancelhubsitejoinapproval) | Cancel hub site join request | Cancel hub join request. If applicable, you should specify the same Approval Correlation Id as used in the "Set hub site join status to pending" action. |
| [`CreateSharingLink`](#createsharinglink) | Create sharing link for a file or folder | Create sharing link for a file or folder. |
| [`CopyFile`](#copyfile) | Copy file (deprecated) | Copies a file to a SharePoint site. |
| [`CopyFileAsync`](#copyfileasync) | Copy file | Copies a file. Works in a similar way to the "Copy to" command in SharePoint libraries. Returns information about the new file after copy. |
| [`CopyFolderAsync`](#copyfolderasync) | Copy folder | Copies a folder. Works in a similar way to the "Copy to" command in SharePoint libraries. Returns information about the new folder after copy. |
| [`ExtractFolderV2`](#extractfolderv2) | Extract folder | Extracts an archive file into a SharePoint folder (example: .zip). |
| [`CreateFile`](#createfile) | Create file | Uploads a file to a SharePoint site. Make sure to pick an existing library. |
| [`GetFileMetadata`](#getfilemetadata) | Get file metadata | Gets information about the file such as size, etag, created date, etc. Uses a file identifier to pick the file. Use "Get file properties" action to get to the values stored in the columns in the library. |
| [`UpdateFile`](#updatefile) | Update file | Updates the contents of the file specified by the file identifier. |
| [`DeleteFile`](#deletefile) | Delete file | Deletes the file specified by the file identifier. |
| [`GetFileContent`](#getfilecontent) | Get file content | Gets file contents using the file identifier. The contents can be copied somewhere else, or be used as an attachment. |
| [`ListRootFolder`](#listrootfolder) | List root folder | Returns files in the root SharePoint folder. |
| [`ListFolder`](#listfolder) | List folder | Returns files contained in a SharePoint folder. |
| [`GetFileMetadataByPath`](#getfilemetadatabypath) | Get file metadata using path | Gets information about the file such as size, etag, created date, etc. Uses a file path to pick the file. Use "Get file properties" action to get to the values stored in the columns in the library. |
| [`GetFileContentByPath`](#getfilecontentbypath) | Get file content using path | Gets file contents using the file path. |
| [`GetFolderMetadata`](#getfoldermetadata) | Get folder metadata | Gets information about the folder. Uses a file identifier to pick the folder. |
| [`GetFolderMetadataByPath`](#getfoldermetadatabypath) | Get folder metadata using path | Gets information about the folder. Uses a folder path to pick the folder. |
| [`HttpRequest`](#httprequest) | Send an HTTP request to SharePoint | Construct a SharePoint REST API to invoke. Note – This action may execute any SharePoint REST API you have access to. Please proceed with caution. |
| [`JoinHubSite`](#joinhubsite) | Join hub site | Join the requested site to the hub site. An Approval Token is required to complete the join successfully if that hub requires approval. If applicable, you should specify the same Approval Correlation Id as used in the "Set hub site join status to pending" action. |
| [`MoveFileAsync`](#movefileasync) | Move file | Moves a file. Works in a similar way to the "Move to" command in SharePoint libraries. Returns information about the new file after move. |
| [`MoveFolderAsync`](#movefolderasync) | Move folder | Moves a folder. Works in a similar way to the "Move to" command in SharePoint libraries. Returns information about the new folder after move. |
| [`NotifyHubSiteJoinApprovalStarted`](#notifyhubsitejoinapprovalstarted) | Set hub site join status to pending | Set the requested site's hub join request status to pending. The Approval Correlation Id is an optional parameter that helps SharePoint identify a particular hub join request. The requesting site can only have one pending request at a given time. |
| [`GetTables`](#gettables) | Get lists | Gets SharePoint lists from a site. |
| [`CreateNewDocumentSet`](#createnewdocumentset) | Create new document set | Creates a new document set list item. |
| [`CreateNewFolder`](#createnewfolder) | Create new folder | Creates a new folder or folder path. |
| [`SearchForUser`](#searchforuser) | Resolve person | Returns a single matching user value so it can be assigned to a column of type person. If there are no matches, or multiple matches, this action will error out. |
| [`GetFileItems`](#getfileitems) | Get files (properties only) | Gets the properties saved in the columns in the library for all folders and files stored in the library.             You can also filter down to the items that match a condition. An "Apply to each" section is usually used to work with the output from this action.             When using this with the On-Premises Data Gateway, the name of the library to connect to may need to be entered manually. |
| [`GetItems`](#getitems) | Get items | Gets items from a SharePoint list. |
| [`PostItem`](#postitem) | Create item | Creates a new item in a SharePoint list. |
| [`GetItem`](#getitem) | Get item | Gets a single item by its id from a SharePoint list. |
| [`DeleteItem`](#deleteitem) | Delete item | Deletes an item from a SharePoint list. |
| [`PatchItem`](#patchitem) | Update item | Updates an item in a SharePoint list. |
| [`CreateApprovalRequest`](#createapprovalrequest) | Create an approval request for an item or file | Creates an approval request for an item or file. |
| [`GetItemChanges`](#getitemchanges) | Get changes for an item or a file (properties only) | Returns information about columns that have changed within a given time window. Note: The list must have Versioning turned on. |
| [`CheckIfFileIsPublished`](#checkiffileispublished) | Check if the scheduled version of the item is published | Returns the result in the output variable IsFilePublished. |
| [`CheckInFile`](#checkinfile) | Check in file | Check in a checked out file in a document library, which makes the version of the document available to others. |
| [`CheckOutFile`](#checkoutfile) | Check out file | Check out a file in a document library to prevent others from editing the document, and your changes from being visible until the documented is checked in. |
| [`DiscardFileCheckOut`](#discardfilecheckout) | Discard check out | If you check out a file and don’t make changes to it, or you make changes that you don’t want to keep, you can simply discard the checkout, rather than saving the file. If your organization tracks versions, a new version is created each time you check a file back into the library. By discarding the checkout, you can avoid making new versions when you haven’t made any changes to the file. |
| [`GetFileItem`](#getfileitem) | Get file properties | Gets the properties saved in the columns in the library for the item specified by the item id.             You can add a "Get file content" step and use the "File identifier" property returned by this action to get to the contents of the file.             When using this with the On-Premises Data Gateway, the name of the library to connect to may need to be entered manually. |
| [`GrantAccess`](#grantaccess) | Grant access to an item or a folder | Grant access to an item or a folder in SharePoint to specific people. |
| [`PatchFileItem`](#patchfileitem) | Update file properties | Updates the properties stored in columns in a library for the item specified by the item id. Use "Update file" action to update file contents.             When using this with the On-Premises Data Gateway, the name of the library to connect to may need to be entered manually. |
| [`PatchFileItemWithPredictedValues`](#patchfileitemwithpredictedvalues) | Update file properties using AI Builder model results | Updates the values stored in library columns for a file analyzed by the model specified by the ModelId. |
| [`SetApprovalStatus`](#setapprovalstatus) | Set content approval status | Sets the content approval status for an item in a list or library that has content approval turned on. You must provide an ETag for pages and files. You can get the ETag using the Get File Metadata action. This action is only available for SharePoint Online and SharePoint 2019. |
| [`UnshareItem`](#unshareitem) | Stop sharing an item or a file | Delete all links giving access to an item or a file and remove all people with direct access except for owners. |
| [`GetItemAttachments`](#getitemattachments) | Get attachments | Returns the list of attachments for the specified list item. You can add a "Get attachment content" step and use the "File identifier" property returned by this action to get to the contents of the file. |
| [`CreateAttachment`](#createattachment) | Add attachment | Adds a new attachment to the specified list item. |
| [`DeleteAttachment`](#deleteattachment) | Delete attachment | Deletes the specified attachment. |
| [`GetAttachmentContent`](#getattachmentcontent) | Get attachment content | Returns file contents using the file identifier. The contents can be copied somewhere else, or be used as an attachment. |
| [`CreateContentAssemblyDocument`](#createcontentassemblydocument) | Generate document using Microsoft Syntex (preview) | Use this action to create documents based on modern templates from Microsoft Syntex. This preview requires a Syntex license. Pricing is subject to change. For more info see: https://docs.microsoft.com/en-us/microsoft-365/contentunderstanding/content-assembly. |
| [`GetTableViews`](#gettableviews) | Get list views | Gets views from a SharePoint list. |
| [`CreateAgreementsSolutionDocument`](#createagreementssolutiondocument) | Agreements Solution - Generate document within Agreements Solution workspace | Use this action to create documents based on modern templates in a Agreements Solution workspace. This is behind a payment wall currently in planning (either license or PayG). |

### GetAllTables

**Get all lists and libraries** — Get all lists and libraries.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (3 fields) | List of Tables |

### ApproveHubSiteJoin

**Approve hub site join request** — Approve hub site join request. This will return an approval token that can be used to complete the join request using the join hub site action.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Hub Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `joiningSiteId` | Requesting Site Id | `String` | ✓ | Id of the requesting site. |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ApprovalToken` |  | `String` | Approval Token |

### CancelHubSiteJoinApproval

**Cancel hub site join request** — Cancel hub join request. If applicable, you should specify the same Approval Correlation Id as used in the "Set hub site join status to pending" action.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Requesting Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `approvalCorrelationId` | Approval Correlation Id | `String` |  | Approval correlation identifier for this request. *advanced* |

**Output** — `String`

### CreateSharingLink

**Create sharing link for a file or folder** — Create sharing link for a file or folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | Library Name | `String` | ✓ | SharePoint library name. *values looked up dynamically* |
| `id` | Item Id | `Number` | ✓ | Id of the file or folder item. |
| `type` | Link Type | `String` | ✓ | The type of sharing link to create *values looked up dynamically* |
| `scope` | Link Scope | `String` | ✓ | Choose who your sharing link gives access to. "Anyone" option will only work if your administrator has enabled it. *values looked up dynamically* |
| `expirationDateTime` | Link Expiration | `DateTime` |  | The date after which the link will expire in yyyy-MM-dd format. Only applicable for anonymous links. *advanced* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `link` |  | `Record` (1 fields) | Internal structure for sharing links |

### CopyFile

**Copy file (deprecated)** — Copies a file to a SharePoint site.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `source` | Source File Path | `String` | ✓ | Path to the source file. |
| `destination` | Destination File Path | `String` | ✓ | Path to the destination file. |
| `overwrite` | Overwrite Flag | `Boolean` |  | Whether or not to overwrite the destination file if it exists. |

**Output** — `Record` (10 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### CopyFileAsync

**Copy file** — Copies a file. Works in a similar way to the "Copy to" command in SharePoint libraries. Returns information about the new file after copy.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Current Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `sourceFileId` | File to Copy | `String` | ✓ | File Identifier |
| `destinationDataset` | Destination Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |
| `destinationFolderPath` | Destination Folder | `String` | ✓ | Destination Folder |
| `nameConflictBehavior` | If another file is already there | `Number` | ✓ | Pick one of the options available *values looked up dynamically* |

**Output** — `Record` (11 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ItemId` |  | `Number` | The value that can be used to Get or Update file properties in libraries. |
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### CopyFolderAsync

**Copy folder** — Copies a folder. Works in a similar way to the "Copy to" command in SharePoint libraries. Returns information about the new folder after copy.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Current Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `sourceFolderId` | Folder to Copy | `String` | ✓ | File Identifier |
| `destinationDataset` | Destination Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |
| `destinationFolderPath` | Destination Folder | `String` | ✓ | Destination Folder |
| `nameConflictBehavior` | If another folder is already there | `Number` | ✓ | Pick one of the options available *values looked up dynamically* |

**Output** — `Record` (11 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ItemId` |  | `Number` | The value that can be used to Get or Update file properties in libraries. |
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### ExtractFolderV2

**Extract folder** — Extracts an archive file into a SharePoint folder (example: .zip).

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `source` | Source File Path | `String` | ✓ | Path to the source file. |
| `destination` | Destination Folder Path | `String` | ✓ | Path to the destination folder. |
| `overwrite` | Overwrite Flag | `Boolean` |  | Whether or not to overwrite the destination file if it exists. |

**Output** — `Table` (10 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### CreateFile

**Create file** — Uploads a file to a SharePoint site. Make sure to pick an existing library.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `folderPath` | Folder Path | `String` | ✓ | Must start with an existing library. Add folders if needed. |
| `name` | File Name | `String` | ✓ | Name of the file. |
| `file` | File Content | `File` | ✓ | Content of the file. |

**Output** — `Record` (11 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ItemId` |  | `Number` | The value that can be used to Get or Update file properties in libraries. |
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### GetFileMetadata

**Get file metadata** — Gets information about the file such as size, etag, created date, etc. Uses a file identifier to pick the file. Use "Get file properties" action to get to the values stored in the columns in the library.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `id` | File Identifier | `String` | ✓ | Select a file. |

**Output** — `Record` (11 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ItemId` |  | `Number` | The value that can be used to Get or Update file properties in libraries. |
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### UpdateFile

**Update file** — Updates the contents of the file specified by the file identifier.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `id` | File Identifier | `String` | ✓ | Select a file. |
| `file` | File Content | `File` | ✓ | Content of the file. |

**Output** — `Record` (10 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### DeleteFile

**Delete file** — Deletes the file specified by the file identifier.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `id` | File Identifier | `String` | ✓ | Select a file. |

**Output** — `String`

### GetFileContent

**Get file content** — Gets file contents using the file identifier. The contents can be copied somewhere else, or be used as an attachment.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `id` | File Identifier | `String` | ✓ | Select a file. |
| `inferContentType` | Infer Content Type | `Boolean` |  | Infer content-type based on extension. *advanced* |

**Output** — `File`

### ListRootFolder

**List root folder** — Returns files in the root SharePoint folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |

**Output** — `Table` (10 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### ListFolder

**List folder** — Returns files contained in a SharePoint folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `id` | File Identifier | `String` | ✓ | Unique identifier of the folder. |

**Output** — `Table` (10 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### GetFileMetadataByPath

**Get file metadata using path** — Gets information about the file such as size, etag, created date, etc. Uses a file path to pick the file. Use "Get file properties" action to get to the values stored in the columns in the library.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `path` | File Path | `String` | ✓ | Select a file. |

**Output** — `Record` (11 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ItemId` |  | `Number` | The value that can be used to Get or Update file properties in libraries. |
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### GetFileContentByPath

**Get file content using path** — Gets file contents using the file path.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `path` | File Path | `String` | ✓ | Select a file. |
| `inferContentType` | Infer Content Type | `Boolean` |  | Infer content-type based on extension. *advanced* |

**Output** — `File`

### GetFolderMetadata

**Get folder metadata** — Gets information about the folder. Uses a file identifier to pick the folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `id` | File Identifier | `String` | ✓ | Select a folder. |

**Output** — `Record` (11 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ItemId` |  | `Number` | The value that can be used to Get or Update file properties in libraries. |
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### GetFolderMetadataByPath

**Get folder metadata using path** — Gets information about the folder. Uses a folder path to pick the folder.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `path` | Folder Path | `String` | ✓ | Select a folder. |

**Output** — `Record` (11 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ItemId` |  | `Number` | The value that can be used to Get or Update file properties in libraries. |
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### HttpRequest

**Send an HTTP request to SharePoint** — Construct a SharePoint REST API to invoke. Note – This action may execute any SharePoint REST API you have access to. Please proceed with caution.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `method` | Method | `String` | ✓ | Http Method |
| `uri` | Uri | `String` | ✓ | Example: _api/web/lists/getbytitle('Documents') |
| `headers` | Headers | `Any` |  | Enter JSON object of request headers |
| `body` | Body | `String` |  | Enter request content in JSON |

**Output** — `Any`

### JoinHubSite

**Join hub site** — Join the requested site to the hub site. An Approval Token is required to complete the join successfully if that hub requires approval. If applicable, you should specify the same Approval Correlation Id as used in the "Set hub site join status to pending" action.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Requesting Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `hubSiteId` | Hub Site Id | `String` | ✓ | Id of the hub site. |
| `approvalToken` | Approval Token | `String` |  | Approval token for this request. *advanced* |
| `approvalCorrelationId` | Approval Correlation Id | `String` |  | Approval correlation identifier for this request. *advanced* |

**Output** — `String`

### MoveFileAsync

**Move file** — Moves a file. Works in a similar way to the "Move to" command in SharePoint libraries. Returns information about the new file after move.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Current Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `sourceFileId` | File to Move | `String` | ✓ | File Identifier |
| `destinationDataset` | Destination Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |
| `destinationFolderPath` | Destination Folder | `String` | ✓ | Destination Folder |
| `nameConflictBehavior` | If another file is already there | `Number` | ✓ | Pick one of the options available *values looked up dynamically* |

**Output** — `Record` (11 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ItemId` |  | `Number` | The value that can be used to Get or Update file properties in libraries. |
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### MoveFolderAsync

**Move folder** — Moves a folder. Works in a similar way to the "Move to" command in SharePoint libraries. Returns information about the new folder after move.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Current Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `sourceFolderId` | Folder to Move | `String` | ✓ | File Identifier |
| `destinationDataset` | Destination Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |
| `destinationFolderPath` | Destination Folder | `String` | ✓ | Destination Folder |
| `nameConflictBehavior` | If another folder is already there | `Number` | ✓ | Pick one of the options available *values looked up dynamically* |

**Output** — `Record` (11 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ItemId` |  | `Number` | The value that can be used to Get or Update file properties in libraries. |
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### NotifyHubSiteJoinApprovalStarted

**Set hub site join status to pending** — Set the requested site's hub join request status to pending. The Approval Correlation Id is an optional parameter that helps SharePoint identify a particular hub join request. The requesting site can only have one pending request at a given time.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Requesting Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `approvalCorrelationId` | Approval Correlation Id | `String` |  | Approval correlation identifier for this request. *advanced* |

**Output** — `String`

### GetTables

**Get lists** — Gets SharePoint lists from a site.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Table` (3 fields) | List of Tables |

### CreateNewDocumentSet

**Create new document set** — Creates a new document set list item.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | Library | `String` | ✓ | SharePoint library name. *values looked up dynamically* |
| `path` | Document Set Path | `String` | ✓ | Example: folder1/folder2/dsName |
| `contentTypeId` | Content Type Id | `String` | ✓ | Example: 0x0120D520 *values looked up dynamically* |
| `DynamicProperties` | Parameters to this operation passed in the body | `UnresolvedDynamicType` |  | Parameters to this operation passed in the body. |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### CreateNewFolder

**Create new folder** — Creates a new folder or folder path.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | List or Library | `String` | ✓ | SharePoint list or library name. *values looked up dynamically* |
| `path` | Folder Path | `String` | ✓ | Example: folder1/folder2/folder3 |
| `view` | Limit Columns by View | `String` |  | Avoid column threshold issues by only using columns defined in a view. *advanced* *values looked up dynamically* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### SearchForUser

**Resolve person** — Returns a single matching user value so it can be assigned to a column of type person. If there are no matches, or multiple matches, this action will error out.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | List or Library | `String` | ✓ | Pick the list or library that the column is in. *values looked up dynamically* |
| `entityId` | Column | `String` | ✓ | Pick the column you want to assign the value to. *values looked up dynamically* |
| `searchValue` | Email or name | `String` | ✓ | Use the email address, or the full name of the user. |
| `view` | Limit Columns by View | `String` |  | Avoid column threshold issues by only using columns defined in a view. *advanced* *values looked up dynamically* |

**Output** — `Record` (7 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Claims` |  | `String` | user claims |
| `DisplayName` |  | `String` | user title |
| `Email` |  | `String` | user email |
| `Picture` |  | `String` | user picture |
| `Department` |  | `String` | user department |
| `JobTitle` |  | `String` | user job title |
| `@odata.type` |  | `String` |  |

### GetFileItems

**Get files (properties only)** — Gets the properties saved in the columns in the library for all folders and files stored in the library.             You can also filter down to the items that match a condition. An "Apply to each" section is usually used to work with the output from this action.             When using this with the On-Premises Data Gateway, the name of the library to connect to may need to be entered manually.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |
| `table` | Library Name | `String` | ✓ | SharePoint library name *values looked up dynamically* |
| `$filter` | Filter Query | `String` |  | An ODATA filter query to restrict the entries returned (e.g. stringColumn eq 'string' OR numberColumn lt 123). *advanced* |
| `$orderby` | Order By | `String` |  | An ODATA orderBy query for specifying the order of entries. *advanced* |
| `$top` | Top Count | `Number` |  | Total number of entries to retrieve (default = all). *advanced* |
| `folderPath` | Limit Entries to Folder | `String` |  | Select a folder, or leave blank for the whole library |
| `viewScopeOption` | Include Nested Items | `String` |  | Return entries contained in sub-folders (default = true) *values looked up dynamically* |
| `view` | Limit Columns by View | `String` |  | Avoid column threshold issues by only using columns defined in a view *advanced* *values looked up dynamically* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### GetItems

**Get items** — Gets items from a SharePoint list.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |
| `table` | List Name | `String` | ✓ | SharePoint list name *values looked up dynamically* |
| `$filter` | Filter Query | `String` |  | An ODATA filter query to restrict the entries returned (e.g. stringColumn eq 'string' OR numberColumn lt 123). *advanced* |
| `$orderby` | Order By | `String` |  | An ODATA orderBy query for specifying the order of entries. *advanced* |
| `$top` | Top Count | `Number` |  | Total number of entries to retrieve (default = all). *advanced* |
| `folderPath` | Limit Entries to Folder | `String` |  | Select a folder, or leave blank for the whole list |
| `viewScopeOption` | Include Nested Items | `String` |  | Return entries contained in sub-folders (default = true) *values looked up dynamically* |
| `view` | Limit Columns by View | `String` |  | Avoid column threshold issues by only using columns defined in a view *advanced* *values looked up dynamically* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### PostItem

**Create item** — Creates a new item in a SharePoint list.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |
| `table` | List Name | `String` | ✓ | SharePoint list name *values looked up dynamically* |
| `item` | Item | `UnresolvedDynamicType` | ✓ | Item to create |
| `view` | Limit Columns by View | `String` |  | Avoid column threshold issues by only using columns defined in a view *advanced* *values looked up dynamically* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### GetItem

**Get item** — Gets a single item by its id from a SharePoint list.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |
| `table` | List Name | `String` | ✓ | SharePoint list name *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Unique identifier of item to be retrieved |
| `view` | Limit Columns by View | `String` |  | Avoid column threshold issues by only using columns defined in a view *advanced* *values looked up dynamically* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### DeleteItem

**Delete item** — Deletes an item from a SharePoint list.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |
| `table` | List Name | `String` | ✓ | SharePoint list name *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Unique identifier of item to be deleted |

**Output** — `String`

### PatchItem

**Update item** — Updates an item in a SharePoint list.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |
| `table` | List Name | `String` | ✓ | SharePoint list name *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Unique identifier of item to be updated |
| `item` | Item | `UnresolvedDynamicType` | ✓ | Item with changed properties |
| `view` | Limit Columns by View | `String` |  | Avoid column threshold issues by only using columns defined in a view *advanced* *values looked up dynamically* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### CreateApprovalRequest

**Create an approval request for an item or file** — Creates an approval request for an item or file.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | List or Library | `String` | ✓ | SharePoint list or library name. *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Id of the item the approval request is being created for. |
| `approvalType` | Approval Type | `Number` | ✓ | Select an approval type. *values looked up dynamically* |
| `approvalSchema` | Approval Schema | `UnresolvedDynamicType` | ✓ | Schema for the selected approval type. |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ApprovalId` | Approval Request ID | `String` | The ID of the approval request created |

### GetItemChanges

**Get changes for an item or a file (properties only)** — Returns information about columns that have changed within a given time window. Note: The list must have Versioning turned on.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | List or Library Name | `String` | ✓ | SharePoint list or library name. *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Id of the item you want to fetch changes for. |
| `since` | Since | `String` | ✓ | Trigger Window Start token or an item version label (ex: 3.0) or an ISO 8601 date (YYYY-MM-DDThh:mmZ). |
| `until` | Until | `String` |  | Trigger Window End token or an item version label (ex: 3.0) or an ISO 8601 date (YYYY-MM-DDThh:mmZ). If blank, defaults to latest version. |
| `includeDrafts` | Include Minor Versions | `Boolean` |  | Boolean for whether to consider changes from minor (draft) versions. *advanced* |
| `view` | Limit Columns by View | `String` |  | Returns only columns defined in a view. *advanced* *values looked up dynamically* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### CheckIfFileIsPublished

**Check if the scheduled version of the item is published** — Returns the result in the output variable IsFilePublished.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | Library Name | `String` | ✓ | SharePoint library name. *values looked up dynamically* |
| `id` | Item ID | `Number` | ✓ | Unique identifier of the file. |
| `scheduledVersion` | Scheduled Version | `String` | ✓ | Version of file that was scheduled for publish. |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `IsFilePublished` |  | `Boolean` | A boolean value (true, false) to indicate whether the scheduled version of the file has been published |

### CheckInFile

**Check in file** — Check in a checked out file in a document library, which makes the version of the document available to others.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | Library Name | `String` | ✓ | SharePoint library name. *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Id of the list item the file is attached to. |
| `comment` | Comments | `String` | ✓ | Type comments describing what has changed in this version |
| `checkinType` | Check in type | `Number` | ✓ | Select the type of version you would like to check in *values looked up dynamically* |

**Output** — `String`

### CheckOutFile

**Check out file** — Check out a file in a document library to prevent others from editing the document, and your changes from being visible until the documented is checked in.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | Library Name | `String` | ✓ | SharePoint library name. *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Id of the list item the file is attached to. |

**Output** — `String`

### DiscardFileCheckOut

**Discard check out** — If you check out a file and don’t make changes to it, or you make changes that you don’t want to keep, you can simply discard the checkout, rather than saving the file. If your organization tracks versions, a new version is created each time you check a file back into the library. By discarding the checkout, you can avoid making new versions when you haven’t made any changes to the file.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | Library Name | `String` | ✓ | SharePoint library name. *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Id of the list item the file is attached to. |

**Output** — `String`

### GetFileItem

**Get file properties** — Gets the properties saved in the columns in the library for the item specified by the item id.             You can add a "Get file content" step and use the "File identifier" property returned by this action to get to the contents of the file.             When using this with the On-Premises Data Gateway, the name of the library to connect to may need to be entered manually.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |
| `table` | Library Name | `String` | ✓ | SharePoint library name *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Unique identifier of item to be retrieved |
| `view` | Limit Columns by View | `String` |  | Avoid column threshold issues by only using columns defined in a view *advanced* *values looked up dynamically* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### GrantAccess

**Grant access to an item or a folder** — Grant access to an item or a folder in SharePoint to specific people.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | List or Library Name | `String` | ✓ | SharePoint list or library name. *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Unique identifier of the item or folder you want to grant access to. |
| `recipients` | Recipients | `String` | ✓ | A collection of recipients who will receive the sharing invitation |
| `roleValue` | Roles | `String` | ✓ | Specify a role that is to be granted to the recipients *values looked up dynamically* |
| `emailBody` | Message | `String` |  | A plain text formatted message that is included in the sharing invitation |
| `sendEmail` | Notify Recipients | `Boolean` |  | Specify whether recipients should receive an email notification message |

**Output** — `String`

### PatchFileItem

**Update file properties** — Updates the properties stored in columns in a library for the item specified by the item id. Use "Update file" action to update file contents.             When using this with the On-Premises Data Gateway, the name of the library to connect to may need to be entered manually.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename *values looked up dynamically* |
| `table` | Library Name | `String` | ✓ | SharePoint library name *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Unique identifier of item to be updated |
| `item` | Item | `UnresolvedDynamicType` | ✓ | Item with changed properties |
| `view` | Limit Columns by View | `String` |  | Avoid column threshold issues by only using columns defined in a view *advanced* *values looked up dynamically* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### PatchFileItemWithPredictedValues

**Update file properties using AI Builder model results** — Updates the values stored in library columns for a file analyzed by the model specified by the ModelId.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | Library Name | `String` | ✓ | Library Name. *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Unique identifier of item to be updated. |
| `modelId` | ModelId | `String` |  | Enter modelId of the Model which is used for prediction |
| `predictResult` | PredictResult | `String` |  | Enter request content in JSON |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### SetApprovalStatus

**Set content approval status** — Sets the content approval status for an item in a list or library that has content approval turned on. You must provide an ETag for pages and files. You can get the ETag using the Get File Metadata action. This action is only available for SharePoint Online and SharePoint 2019.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | Library Name | `String` | ✓ | SharePoint library name. *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Id of the item you are setting the status of. |
| `approvalAction` | Action | `String` | ✓ | Pick the approval action. |
| `comments` | Comments | `String` |  | Add the comments from the approver. |
| `entityTag` | ETag | `String` |  | Add an ETag (required for files and pages). |

**Output** — `Record` (5 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ETag` |  | `String` | ETag of the item after the approval status was set |
| `ApprovalLink` |  | `String` | A link to the item that needs approval |
| `PublishStartDate` |  | `String` | Date time at which the item will be Published |
| `ContentApprovalStatus` |  | `String` | The content approval status of the list item |
| `ScheduledVersion` |  | `String` | The version of the item that has been scheduled |

### UnshareItem

**Stop sharing an item or a file** — Delete all links giving access to an item or a file and remove all people with direct access except for owners.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | List or Library Name | `String` | ✓ | SharePoint list or library name. *values looked up dynamically* |
| `id` | Id | `Number` | ✓ | Unique identifier of the item or file you want to stop sharing. |

**Output** — `String`

### GetItemAttachments

**Get attachments** — Returns the list of attachments for the specified list item. You can add a "Get attachment content" step and use the "File identifier" property returned by this action to get to the contents of the file.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | List Name | `String` | ✓ | SharePoint list name. *values looked up dynamically* |
| `itemId` | Id | `String` | ✓ | Id of the list item to get attachments from. |

**Output** — `Table` (3 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Id` |  | `String` | File identifier |
| `AbsoluteUri` |  | `String` | Link to attachment |
| `DisplayName` |  | `String` | Name |

### CreateAttachment

**Add attachment** — Adds a new attachment to the specified list item.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | List Name | `String` | ✓ | SharePoint list name. *values looked up dynamically* |
| `itemId` | Id | `Number` | ✓ | Id of the list item to attach the file to. |
| `displayName` | File Name | `String` | ✓ | File name. |
| `file` | File Content | `File` | ✓ | Content of the file. |

**Output** — `Record` (3 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Id` |  | `String` | File identifier |
| `AbsoluteUri` |  | `String` | Link to attachment |
| `DisplayName` |  | `String` | Name |

### DeleteAttachment

**Delete attachment** — Deletes the specified attachment.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | List Name | `String` | ✓ | SharePoint list name. *values looked up dynamically* |
| `itemId` | Id | `Number` | ✓ | Id of the list item the file is attached to. |
| `attachmentId` | File Identifier | `String` | ✓ | File identifier for the attachment. |

**Output** — `String`

### GetAttachmentContent

**Get attachment content** — Returns file contents using the file identifier. The contents can be copied somewhere else, or be used as an attachment.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | List Name | `String` | ✓ | SharePoint list name. *values looked up dynamically* |
| `itemId` | Id | `Number` | ✓ | Id of the list item the file is attached to. |
| `attachmentId` | File Identifier | `String` | ✓ | File identifier for the attachment. |

**Output** — `File`

### CreateContentAssemblyDocument

**Generate document using Microsoft Syntex (preview)** — Use this action to create documents based on modern templates from Microsoft Syntex. This preview requires a Syntex license. Pricing is subject to change. For more info see: https://docs.microsoft.com/en-us/microsoft-365/contentunderstanding/content-assembly.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | Document Library Name | `String` | ✓ | SharePoint document library name. *values looked up dynamically* |
| `template` | Document Template | `String` | ✓ | Document template. *values looked up dynamically* |
| `item` | Placeholders | `UnresolvedDynamicType` | ✓ | Document placeholder values. |
| `folderPath` | Folder Path | `String` |  | Must start with an existing library. |
| `fileName` | File Name | `String` |  | Document file name. |
| `view` | View (no effect) | `String` |  | This parameter does nothing. Please do not use. *advanced* |

**Output** — `Record` (11 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ItemId` |  | `Number` | The value that can be used to Get or Update file properties in libraries. |
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

### GetTableViews

**Get list views** — Gets views from a SharePoint list.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Site Address | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `table` | List Name | `String` | ✓ | SharePoint list name. *values looked up dynamically* |

**Output** — `Table` (3 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Name` |  | `String` | The name of the table. The name is used at runtime. |
| `DisplayName` |  | `String` | The display name of the table. |
| `DynamicProperties` |  | `Any` | Additional table properties provided by the connector to the clients. |

### CreateAgreementsSolutionDocument

**Agreements Solution - Generate document within Agreements Solution workspace** — Use this action to create documents based on modern templates in a Agreements Solution workspace. This is behind a payment wall currently in planning (either license or PayG).

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `dataset` | Agreements Solution Workspace | `String` | ✓ | Example: https://contoso.sharepoint.com/sites/sitename. *values looked up dynamically* |
| `template` | Agreements Solution Template | `String` | ✓ | Agreements Solution template. *values looked up dynamically* |
| `item` | Fields | `UnresolvedDynamicType` | ✓ | Document placeholder values. |
| `documentName` | File Name | `String` |  | Document file name. |
| `table` | Table (no effect) | `String` |  | This parameter does nothing. Please do not use. *advanced* |
| `view` | View (no effect) | `String` |  | This parameter does nothing. Please do not use. *advanced* |

**Output** — `Record` (11 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `ItemId` |  | `Number` | The value that can be used to Get or Update file properties in libraries. |
| `Id` |  | `String` | The unique id of the file or folder. |
| `Name` |  | `String` | The name of the file or folder. |
| `DisplayName` |  | `String` | The display name of the file or folder. |
| `Path` |  | `String` | The path of the file or folder. |
| `LastModified` |  | `DateTime` | The date and time the file or folder was last modified. |
| `Size` |  | `Number` | The size of the file or folder. |
| `MediaType` |  | `String` | The media type of the file or folder. |
| `IsFolder` |  | `Boolean` | A boolean value (true, false) to indicate whether or not the blob is a folder. |
| `ETag` |  | `String` | The etag of the file or folder. |
| `FileLocator` |  | `String` | The filelocator of the file or folder. |

## Triggers

Connector triggers start a flow or an external-trigger when something happens in the connected service. See [10 Flows & autonomous triggers](../../10-flows-and-autonomous-triggers.md).

| operationId | Name | Description |
|---|---|---|
| `GetOnChangedItems` | When an item or a file is modified | Triggers when an item is modified (but not when it is created). |
| `GetOnDeletedFileItems` | When a file is deleted | Triggers when a file is deleted in a library. You can optionally specify a folder to watch as well. When a folder is deleted, the trigger will fire only once for the deleted folder. This can only be used by site collection admins of the site where the list is located. |
| `GetOnDeletedItems` | When an item is deleted | Triggers when an item is deleted in a list. This can only be used by site collection admins of the site where the list is located. |
| `GetOnNewFileItems` | When a file is created (properties only) | Triggers when an item is created in a library. Returns only the properties stored in the library columns.             You can add a "Get file content" step and use the "File identifier" property returned by this action to get to the contents of the file.             When using this with the On-Premises Data Gateway, the name of the library to connect to may need to be entered manually. |
| `GetOnNewItems` | When an item is created | Triggers when an item is created. |
| `GetOnUpdatedFileClassifiedTimes` | When a file is classified by a Microsoft Syntex model | Triggers a flow when Microsoft Syntex changes the classification date of any file in the library. The date changes when a document processing model classifies or extracts information. |
| `GetOnUpdatedFileItems` | When a file is created or modified (properties only) | Triggers when an item is created, or modified in a library. Returns only the properties stored in the library columns.             You can add a "Get file content" step and use the "File identifier" property returned by this action to get to the contents of the file.             When using this with the On-Premises Data Gateway, the name of the library to connect to may need to be entered manually. |
| `GetOnUpdatedItems` | When an item is created or modified | Triggers when an item is created, and also each time it is modified. |
| `OnNewFile` | When a file is created in a folder (deprecated) | Triggers when a file is created in a SharePoint folder. The trigger does not fire if a file is added/updated in a subfolder. If it is required to trigger on subfolders, multiple triggers should be created. |
| `OnUpdatedFile` | When a file is created or modified in a folder (deprecated) | Triggers when a file is created, and also each time it is modified in a SharePoint folder. The trigger does not fire if a file is added/updated in a subfolder. If it is required to trigger on subfolders, multiple triggers should be created. |
