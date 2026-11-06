# Connector: Azure DevOps (`shared_visualstudioteamservices`)

> **Generated** by `tools/gen_reference.py` from `reference/connectors/shared_visualstudioteamservices.yml` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

Azure DevOps provides services for teams to share code, track work, and ship software - for any language, all in a single package. It's the perfect complement to your IDE.

| | |
|---|---|
| Connector id | `/providers/Microsoft.PowerApps/apis/shared_visualstudioteamservices` |
| Type | Public |
| Operations | 26 |
| Triggers (used by flows / external triggers, not by tools) | 13 |

Use an operation's **operationId** as `action.operationId` in a `TaskDialog` tool (see [08 Tools & connectors](../../08-tools-and-connectors.md)). Inputs map to `inputs:` entries (`propertyName` = the input *Name* below). Input names that start with `$` are [OData](../../glossary.md#odata) query options and need special quoting — see [`$`-prefixed inputs](../../08-tools-and-connectors.md#-prefixed-inputs-odata-parameters).

## Operations

| operationId | Name | Description |
|---|---|---|
| [`SearchWorkItemsAsync`](#searchworkitemsasync) | Search for work items | Provides a set of results for the search text. |
| [`ListAccounts`](#listaccounts) | List Organizations | Retrieves the list of Organizations the user is a member of. |
| [`GetProfile`](#getprofile) | Get a User Profile | Gets a user profile. |
| [`ListIterations`](#listiterations) | List iterations | Retrieves the list of iterations for a project. |
| [`QueueNewBuild`](#queuenewbuild) | Queue a new build | Create a build, based on an existing build definition, and add it to the build queue. |
| [`ListGitRepositories`](#listgitrepositories) | List Git repositories | Retrieves the list of Git repositories in a project. |
| [`ListProjects`](#listprojects) | List projects | Retrieves the list of projects in an Organization. |
| [`ListReleaseDefinitions`](#listreleasedefinitions) | List release definitions | Retrieves the list of release definitions associated with a project. |
| [`CreateRelease`](#createrelease) | Create a new release | Create a release, based on an existing release definition. |
| [`HttpRequest`](#httprequest) | Send an HTTP request to Azure DevOps | Construct an Azure DevOps REST API request to invoke. Please refer to the following link for all available endpoints: https://docs.microsoft.com/rest/api/azure/devops |
| [`ListWorkItemTypes`](#listworkitemtypes) | List work item types | Retrieves the list of Work Item Types accessible in a given Azure DevOps project. |
| [`GetWorkItemDetails`](#getworkitemdetails) | Get work item details | Get the details of a single work item. |
| [`UpdateWorkItem`](#updateworkitem) | Update a work item | Update an existing work item by id. |
| [`GetWorkItemChildren`](#getworkitemchildren) | Get work item children | Get a list of children (max size 1000) of a single work item. |
| [`GetWorkItemCommentsAsync`](#getworkitemcommentsasync) | Get work item comments | Retrieves comments for a work item. |
| [`CreateWorkItemCommentAsync`](#createworkitemcommentasync) | Create work item comment | Adds a comment to a work item. |
| [`DeleteWorkItemCommentAsync`](#deleteworkitemcommentasync) | Delete work item comment | Delete a comment on a work item. |
| [`UpdateWorkItemCommentAsync`](#updateworkitemcommentasync) | Update work item comment | Updates a comment on a work item. |
| [`CreateWorkItem`](#createworkitem) | Create a work item | Create a new work item with the provided attributes. |
| [`GetQueryResults`](#getqueryresults) | Get query results | Retrieves the result of a work item query. |
| [`GetQueryResultsV2`](#getqueryresultsv2) | Get query results | Retrieves the result of a work item query. |
| [`ListRootQueryFolders`](#listrootqueryfolders) | List root level queries | Retrieves the list of root-level work item queries in a project. |
| [`ListQueriesInFolder`](#listqueriesinfolder) | List queries within folder | Retrieves the list of work item queries within a folder for a given project. |
| [`ListWorkItems`](#listworkitems) | List work items | Returns a list of work items. |
| [`ListPipelines`](#listpipelines) | List pipelines | Retrieves the list of pipelines in a project. |
| [`ListPipelineRuns`](#listpipelineruns) | List pipeline runs | Gets top 10000 runs for a particular pipeline |

### SearchWorkItemsAsync

**Search for work items** — Provides a set of results for the search text.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have \"Third-party application access via OAuth\" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `searchText` | Search Text | `String` |  | The search text. |
| `$skip` | Skip | `Number` |  | Number of results to be skipped. Defaults to 0. *advanced* |
| `$top` | Top | `Number` |  | Number of results to be returned. Defaults to 100. *advanced* |
| `filters` | Filters | `Any` |  | Filters to be applied. Set it to null if there are no filters to be applied. *advanced* |
| `includeFacets` | Include Facets | `Boolean` |  | Flag to opt for faceting in the result. Default behavior is false. *advanced* |
| `$orderBy` | Order By | `Table` (2 fields) |  | Options for sorting search results. If set to null, the results will be returned sorted by relevance. If more than one sort option is provided, the results are sorted in the order specified in the OrderBy. *advanced* |

**Output** — `Record` (4 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `count` | Count | `Number` | Total number of matched work items. |
| `results` | Results | `Table` (4 fields) | List of top matched work items. |
| `infoCode` | Info Code | `Number` | Numeric code indicating any additional information. *advanced* |
| `facets` | Facets | `Any` | A dictionary storing an array of Filter object against each facet. *advanced* |

### ListAccounts

**List Organizations** — Retrieves the list of Organizations the user is a member of.

**Inputs**

_None._

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | Value | `Table` (6 fields) | The VSTS Models |

### GetProfile

**Get a User Profile** — Gets a user profile.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `id` | Profile Id | `String` | ✓ | The ID of the profile within the same organization, or 'me' to get the profile of the current authenticated user. |

**Output** — `Record` (7 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `displayName` | displayName | `String` | Display Name of User |
| `id` | id | `String` | The unique identifier of the profile. |
| `publicAlias` | publicAlias | `String` | Public Alias. |
| `emailAddress` | Email Address | `String` | Email Address of the account. |
| `timeStamp` | timeStamp | `DateTime` | The time at which this profile was last changed. *advanced* |
| `revision` | revision | `Number` | The maximum revision number of any attribute. *advanced* |
| `coreRevision` | CoreRevision | `Number` | CoreRevision *advanced* |

### ListIterations

**List iterations** — Retrieves the list of iterations for a project.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `team` | Team Name | `String` | ✓ | The name of the team within the project. *values looked up dynamically* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | Value | `Table` (6 fields) | The VSTS Models |

### QueueNewBuild

**Queue a new build** — Create a build, based on an existing build definition, and add it to the build queue.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `buildDefId` | Build Definition Id | `String` | ✓ | The identifier of the build definition to use. *values looked up dynamically* |
| `sourceBranch` | Source Branch | `String` |  | The source branch of the build. |
| `parameters` | Parameters | `String` |  | A JSON dictionary of optional build parameters. |

**Output** — `Record` (15 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `id` | Build Id | `Number` | The unique identifier of the build. |
| `buildNumber` | Build Number | `String` | The build number. |
| `sourceBranch` | Source Branch | `String` | The source control branch the build used. *advanced* |
| `sourceVersion` | Source version | `String` | The source control version the build used. *advanced* |
| `status` | Status | `String` | The status of the build. *advanced* |
| `priority` | Priority | `String` | The priority of the build. *advanced* |
| `queueTime` | Queue Time | `DateTime` | The time the build was queued. *advanced* |
| `startTime` | Start Time | `DateTime` | The time the build started. *advanced* |
| `finishTime` | Finish Time | `DateTime` | The time the build finished. *advanced* |
| `reason` | Reason | `String` | The reason for the build. *advanced* |
| `result` | Result | `String` | The result of the build. |
| `requestedFor` |  | `Record` (1 fields) | Build Request User |
| `parameters` | Parameters | `String` | Parameters passed to the build, if any. *advanced* |
| `definition` |  | `Record` (2 fields) | Build Result Definition |
| `_links` |  | `Record` (1 fields) | Links |

### ListGitRepositories

**List Git repositories** — Retrieves the list of Git repositories in a project.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | Value | `Table` (4 fields) | The VSTS Models |

### ListProjects

**List projects** — Retrieves the list of projects in an Organization.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | Value | `Table` (3 fields) | The VSTS Models |

### ListReleaseDefinitions

**List release definitions** — Retrieves the list of release definitions associated with a project.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | Value | `Table` (13 fields) | The VSTS Models |

### CreateRelease

**Create a new release** — Create a release, based on an existing release definition.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `releaseDefId` | Release Definition Id | `String` | ✓ | The identifier of the release definition to use. *values looked up dynamically* |
| `Description` | Description | `String` |  | The description of the release start metadata. |
| `IsDraft` | Is Draft | `Boolean` |  | Specifies whether the release is a draft. |
| `Reason` | Reason | `String` |  | The reason for the release start metadata. |
| `Variables` | Release variables | `Table` (2 fields) |  | Release variables. |

**Output** — `Record` (14 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `Comment` | Comment | `String` | The comment of the release. *advanced* |
| `CreatedBy` |  | `Record` (5 fields) | Identity Ref |
| `CreatedOn` | Created On | `DateTime` | The date on which the release was created. |
| `Description` | Description | `String` | The description of the release. |
| `Id` | Id | `Number` | The unique Identifier of the release. *advanced* |
| `KeepForever` | Keep Forever | `Boolean` | Whether to exclude the release from retention policies. *advanced* |
| `LogsContainerUrl` | Logs Container URL | `String` | The logs Container URL of this release. *advanced* |
| `ModifiedBy` |  | `Record` (5 fields) | Identity Ref |
| `ModifiedOn` | Modified On | `DateTime` | Date on which the release was modified. *advanced* |
| `Name` | Name | `String` | The name of the release. *advanced* |
| `Reason` | Reason | `String` | The reason of the release. *advanced* |
| `ReleaseDefinition` |  | `Record` (3 fields) | Release Definition Shallow Reference |
| `Status` | Status | `String` | The status of the release. *advanced* |
| `Url` | URL | `String` | The URL of the release. *advanced* |

### HttpRequest

**Send an HTTP request to Azure DevOps** — Construct an Azure DevOps REST API request to invoke. Please refer to the following link for all available endpoints: https://docs.microsoft.com/rest/api/azure/devops

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `Method` | Method | `String` | ✓ | The HTTP method. |
| `Uri` | Relative URI | `String` | ✓ | The relative URI. Example: {project}/{team}/_apis/wit/templates?api-version=5.0-preview.1. |
| `Headers` | Headers | `Any` |  | The request headers. |
| `Body` | Body | `String` |  | The request content. Can be JSON object or binary data encoded as Base64 string. |
| `IsBase64` | Body is Base64 | `Boolean` |  | Must be true if request content is encoded as Base64 string. |

**Output** — `Any`

### ListWorkItemTypes

**List work item types** — Retrieves the list of Work Item Types accessible in a given Azure DevOps project.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | Value | `Table` (6 fields) | The VSTS Models. |

### GetWorkItemDetails

**Get work item details** — Get the details of a single work item.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `typeName` | Work Item Type | `String` | ✓ | The work item type of the work item. *values looked up dynamically* |
| `id` | Work Item Id | `String` | ✓ | The unique id of the work item to get details for.The work item must be of the type specified in 'Work Item Type'. |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### UpdateWorkItem

**Update a work item** — Update an existing work item by id.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `id` | Work Item Id | `String` | ✓ | The unique id of the work item to update. |
| `project` | Project Name | `String` |  | The name of the project within the specified Organization. *values looked up dynamically* |
| `type` | Work Item Type | `String` |  | The work item type to generate schema for. *values looked up dynamically* |
| `title` | Title | `String` |  | The title of the work item. |
| `description` | Description | `String` |  | The description of the work item. |
| `priority` | Priority | `Number` |  | The priority of the work item. *advanced* |
| `iteration` | Iteration Path | `String` |  | The iteration path of the work item. *advanced* |
| `area` | Area Path | `String` |  | The area path of the work item. *advanced* |
| `linkUrl` | Link URL | `String` |  | The URL of the work item to link with this one. *advanced* |
| `linkType` | Link Type | `String` |  | The type of link to between the work items. *advanced* |
| `linkComment` | Link Comment | `String` |  | The comment for the link. *advanced* |
| `dynamicFields` | Work Item Details | `UnresolvedDynamicType` |  | The details of the work item. |
| `userEnteredFields` | Other Fields | `Any` |  | Other work item fields, including custom fields. 'Key' is field name and 'Value' is field value. *advanced* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### GetWorkItemChildren

**Get work item children** — Get a list of children (max size 1000) of a single work item.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `id` | Work Item Id | `String` | ✓ | The unique id of the work item to get the children for. |
| `workItemType` | Children Work Item Type | `String` |  | The type of the children work item(s) *values looked up dynamically* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### GetWorkItemCommentsAsync

**Get work item comments** — Retrieves comments for a work item.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `workItemId` | Work Item Id | `Number` | ✓ | The unique id of the work item to get comments for. |
| `$top` | Maximum comments to return | `Number` |  | Maximum number of comments to return. *advanced* |
| `order` | Comment sort order | `String` |  | Sort order for comments. Defaults to 'desc'. *advanced* |
| `includeDeleted` | Include deleted comments | `Boolean` |  | Whether to include comments that have been soft-deleted. *advanced* |
| `$expand` | Expand options | `String` |  | Controls which additional comment data should be returned. *advanced* |

**Output** — `Record` (2 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `count` | Count | `Number` | The total count of comments. |
| `comments` | Comments | `Table` (17 fields) | The list of comments. |

### CreateWorkItemCommentAsync

**Create work item comment** — Adds a comment to a work item.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | Project ID or project name *values looked up dynamically* |
| `workItemId` | Work Item Id | `Number` | ✓ | Id of a work item. |
| `text` | Comment Text | `String` | ✓ | The text of the comment. |
| `format` | Comment Format | `String` |  | Format of a work item comment (Markdown or Html). *advanced* |

**Output** — `Record` (17 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `_links` |  | `Any` | Collection of REST reference links. |
| `id` | Comment Id | `Number` | The unique id of the comment. |
| `workItemId` | Work Item Id | `Number` | The work item id the comment belongs to. *advanced* |
| `version` | Version | `Number` | The work item revision associated with this comment. *advanced* |
| `format` | Comment Format | `String` | Specifies how comment text is interpreted. *advanced* |
| `text` | Comment Text | `String` | The text of the comment. |
| `renderedText` | Rendered Text | `String` | The rendered HTML version of the comment text. *advanced* |
| `createdBy` |  | `Record` (5 fields) | Identity Ref |
| `createdDate` | Created Date | `DateTime` | The date the comment was created. *advanced* |
| `createdOnBehalfOf` |  | `Record` (5 fields) | Identity Ref |
| `createdOnBehalfDate` | Created On Behalf Date | `DateTime` | Effective creation date on behalf of another identity. *advanced* |
| `modifiedBy` |  | `Record` (5 fields) | Identity Ref |
| `modifiedDate` | Modified Date | `DateTime` | The date the comment was last modified. *advanced* |
| `isDeleted` | Is Deleted | `Boolean` | Indicates if the comment has been deleted. *advanced* |
| `mentions` | Mentions | `Table` (6 fields) | Mentions included in the comment. *advanced* |
| `reactions` | Reactions | `Table` (6 fields) | Reactions on the comment. *advanced* |
| `url` | URL | `String` | The REST API URL of the comment. *advanced* |

### DeleteWorkItemCommentAsync

**Delete work item comment** — Delete a comment on a work item.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `workItemId` | Work Item Id | `Number` | ✓ | The unique id of the work item. |
| `commentId` | Comment Id | `Number` | ✓ | The id of the comment to delete. |

**Output** — `String`

### UpdateWorkItemCommentAsync

**Update work item comment** — Updates a comment on a work item.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | Project ID or project name *values looked up dynamically* |
| `workItemId` | Work Item Id | `Number` | ✓ | Id of a work item. |
| `commentId` | Comment Id | `Number` | ✓ | Id of the comment to update. |
| `text` | Comment Text | `String` | ✓ | The text of the comment. |
| `format` | Comment Format | `String` |  | Format of a work item comment (Markdown or Html). *advanced* |

**Output** — `Record` (17 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `_links` |  | `Any` | Collection of REST reference links. |
| `id` | Comment Id | `Number` | The unique id of the comment. |
| `workItemId` | Work Item Id | `Number` | The work item id the comment belongs to. *advanced* |
| `version` | Version | `Number` | The work item revision associated with this comment. *advanced* |
| `format` | Comment Format | `String` | Specifies how comment text is interpreted. *advanced* |
| `text` | Comment Text | `String` | The text of the comment. |
| `renderedText` | Rendered Text | `String` | The rendered HTML version of the comment text. *advanced* |
| `createdBy` |  | `Record` (5 fields) | Identity Ref |
| `createdDate` | Created Date | `DateTime` | The date the comment was created. *advanced* |
| `createdOnBehalfOf` |  | `Record` (5 fields) | Identity Ref |
| `createdOnBehalfDate` | Created On Behalf Date | `DateTime` | Effective creation date on behalf of another identity. *advanced* |
| `modifiedBy` |  | `Record` (5 fields) | Identity Ref |
| `modifiedDate` | Modified Date | `DateTime` | The date the comment was last modified. *advanced* |
| `isDeleted` | Is Deleted | `Boolean` | Indicates if the comment has been deleted. *advanced* |
| `mentions` | Mentions | `Table` (6 fields) | Mentions included in the comment. *advanced* |
| `reactions` | Reactions | `Table` (6 fields) | Reactions on the comment. *advanced* |
| `url` | URL | `String` | The REST API URL of the comment. *advanced* |

### CreateWorkItem

**Create a work item** — Create a new work item with the provided attributes.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `type` | Work Item Type | `String` | ✓ | Type of this work item *values looked up dynamically* |
| `title` | Title | `String` | ✓ | The title of the work item. |
| `shouldReturnAllFields` | Return all fields | `Boolean` |  | When setting the value to 'No', only the work item ID will be returned. If the value is 'Yes', all fields of the created Item along with the ID will be returned. The default value is 'Yes' *advanced* |
| `description` | Description | `String` |  | The description of the work item. |
| `priority` | Priority | `Number` |  | The priority of the work item. *advanced* |
| `iteration` | Iteration Path | `String` |  | The iteration path of the work item. *advanced* |
| `area` | Area Path | `String` |  | The area path of the work item. *advanced* |
| `linkUrl` | Link URL | `String` |  | The URL of the work item to link with this one. *advanced* |
| `linkType` | Link Type | `String` |  | The type of the link to between the work items. *advanced* |
| `linkComment` | Link Comment | `String` |  | The comment for the link. *advanced* |
| `dynamicFields` | Work Item Details | `UnresolvedDynamicType` |  | Information to include in the new work item |
| `userEnteredFields` | Other Fields | `Any` |  | Other work item fields, including custom fields. 'Key' is field name and 'Value' is field value. *advanced* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### GetQueryResults

**Get query results** — Retrieves the result of a work item query.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `queryId` | Query Id | `String` | ✓ | The Id of the stored query to run. |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | Value | `Table` (1 fields) | The VSTS Models |

### GetQueryResultsV2

**Get query results** — Retrieves the result of a work item query.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `queryId` | Query Id | `String` | ✓ | The Id of the stored query to run. |
| `workItemsCount` | Number of work items to return | `Number` |  | Max work items count to retrieve (1 - 20000) *advanced* |
| `throwIfQueryChanged` | Fail if query results change | `Boolean` |  | If query results change during flow execution then flow run should fail to avoid inconsistencies. *advanced* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### ListRootQueryFolders

**List root level queries** — Retrieves the list of root-level work item queries in a project.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | Value | `Table` (20 fields) | The VSTS Models |

### ListQueriesInFolder

**List queries within folder** — Retrieves the list of work item queries within a folder for a given project.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `folderPath` | Folder Path | `String` | ✓ | The folder path under which to look for queries. |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | Value | `Table` (20 fields) | The VSTS Models |

### ListWorkItems

**List work items** — Returns a list of work items.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `workItemIds` | Work Item IDs | `String` | ✓ | The comma-separated list of requested work item IDs. |
| `workItemType` | Work Item Type | `String` |  | The work item type to generate schema for. *values looked up dynamically* |

**Output** — `UnresolvedDynamicType` (shape depends on the list/table chosen at design time)

### ListPipelines

**List pipelines** — Retrieves the list of pipelines in a project.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |

**Output** — `Record` (2 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `count` |  | `Number` | count |
| `value` |  | `Table` (6 fields) | Array that has the list of all pipelines |

### ListPipelineRuns

**List pipeline runs** — Gets top 10000 runs for a particular pipeline

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `account` | Organization Name | `String` | ✓ | The name of the Organization. This can be found in the instance URL i.e. {Organization}.visualstudio.com. The Organization should have "Third-party application access via OAuth" enabled to prevent errors in the operation. *values looked up dynamically* |
| `project` | Project Name | `String` | ✓ | The name of the project within the specified Organization. *values looked up dynamically* |
| `pipelineId` | The pipeline ID | `Number` | ✓ | The pipeline ID *advanced* |

**Output** — `Record` (2 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `count` |  | `Number` | count |
| `value` |  | `Table` (9 fields) | Pipeline Run |

## Triggers

Connector triggers start a flow or an external-trigger when something happens in the connected service. See [10 Flows & autonomous triggers](../../10-flows-and-autonomous-triggers.md).

| operationId | Name | Description |
|---|---|---|
| `OnBuildCompleted` | When a build completes | Triggers a flow when a build completes. |
| `OnGitPush` | When code is pushed (Git) | Triggers a flow when code is pushed to a Git repository. |
| `OnGitPullCreated` | When a pull request is created (Git) | Triggers a flow when a pull request which matches provided criteria is created. |
| `OnGitPullClosed` | When a pull request is closed (Git) | Triggers a flow when a pull request which matches provided criteria is closed. |
| `OnTfvcCheckIn` | When code is checked in (TFVC) | Triggers a flow when code is checked in to a Team Foundation Version Control project. |
| `OnWorkItemAssigned` | When a work item is assigned | Triggers a flow when a work item which matches provided criteria is assigned to the given user. |
| `OnWorkItemAssignedV2` | When a work item is assigned | Triggers a flow when a work item which matches provided criteria is assigned to the given user. |
| `OnWorkItemUpdated` | When a work item is updated | Triggers a flow when a work item which matches provided criteria is updated. |
| `OnWorkItemUpdatedV2` | When a work item is updated | Triggers a flow when a work item which matches provided criteria is updated. |
| `OnWorkItemCreated` | When a work item is created | Triggers a flow when a work item which matches provided criteria is created. |
| `OnWorkItemCreatedV2` | When a work item is created | Triggers a flow when a work item which matches provided criteria is created. |
| `OnWorkItemClosed` | When a work item is closed | Triggers a flow when a work item which matches provided criteria is closed. |
| `OnWorkItemClosedV2` | When a work item is closed | Triggers a flow when a work item which matches provided criteria is closed. |
