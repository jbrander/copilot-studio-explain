# Connector: Office 365 Users (`shared_office365users`)

> **Generated** by `tools/gen_reference.py` from `reference/connectors/shared_office365users.yml` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

Office 365 Users Connection provider lets you access user profiles in your organization using your Office 365 account. You can get your profile, a user's profile, a user's manager, or direct reports.

| | |
|---|---|
| Connector id | `/providers/Microsoft.PowerApps/apis/shared_office365users` |
| Type | Public |
| Operations | 6 |
| Triggers (used by flows / external triggers, not by tools) | 0 |

Use an operation's **operationId** as `action.operationId` in a `TaskDialog` tool (see [08 Tools & connectors](../../08-tools-and-connectors.md)). Inputs map to `inputs:` entries (`propertyName` = the input *Name* below). Input names that start with `$` are [OData](../../glossary.md#odata) query options and need special quoting — see [`$`-prefixed inputs](../../08-tools-and-connectors.md#-prefixed-inputs-odata-parameters).

## Operations

| operationId | Name | Description |
|---|---|---|
| [`MyProfile_V2`](#myprofile_v2) | Get my profile (V2) | Retrieves the profile of the current user from Microsoft Graph. Learn more: https://developer.microsoft.com/en-us/graph/docs/api-reference/v1.0/resources/user#properties |
| [`UserProfile_V2`](#userprofile_v2) | Get user profile (V2) | Retrieves a specific user's profile from Microsoft Graph by user principal name or user ID. |
| [`UserPhoto`](#userphoto) | Get user photo | Retrieves the photo of the specified user or the current user. |
| [`SearchUser_V2`](#searchuser_v2) | Search for users (V2) | Searches for users in the organization by name, email, or other attributes using Microsoft Graph. |
| [`DirectReports_V2`](#directreports_v2) | Get direct reports (V2) | Retrieves the direct reports of the specified user from Microsoft Graph. |
| [`Manager_V2`](#manager_v2) | Get manager (V2) | Retrieves the manager of the specified user from Microsoft Graph. |

### MyProfile_V2

**Get my profile (V2)** — Retrieves the profile of the current user from Microsoft Graph. Learn more: https://developer.microsoft.com/en-us/graph/docs/api-reference/v1.0/resources/user#properties

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `$select` | Select fields | `String` |  | Comma-separated list of user properties to retrieve (e.g., "id, displayName, mail, department") |

**Output** — `Record` (31 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `aboutMe` | About me | `String` | A freeform text entry field for the user to describe themselves |
| `accountEnabled` | Account enabled | `Boolean` | Whether the account is enabled |
| `birthday` | Birthday | `DateTime` | The birthday of the user |
| `businessPhones` | Business phones | `String` | The telephone numbers for the user |
| `city` | City | `String` | The city in which the user is located |
| `companyName` | Company name | `String` | The company name which the user is associated |
| `country` | Country | `String` | The country/region in which the user is located |
| `department` | Department | `String` | The name for the department in which the user works |
| `displayName` | Display name | `String` | The name displayed in the address book for the user |
| `givenName` | Given name | `String` | The given name (first name) of the user |
| `hireDate` | Hire date | `DateTime` | The hire date of the user |
| `id` | Id | `String` | The unique identifier for the user |
| `interests` | Interests | `String` | A list of the user's interests |
| `jobTitle` | Job title | `String` | The user's job title |
| `mail` | Mail | `String` | The SMTP address for the user |
| `mailNickname` | Mail nickname | `String` | The mail alias for the user |
| `mobilePhone` | Mobile phone | `String` | The primary cellular telephone number for the user |
| `mySite` | My site | `String` | The URL for the user's personal site |
| `officeLocation` | Office location | `String` | The office location in the user's place of business |
| `pastProjects` | Past projects | `String` | A list of the user's past projects |
| `postalCode` | Postal code | `String` | The postal code for the user's postal address |
| `preferredLanguage` | Preferred language | `String` | The preferred language for the user (e.g., en-US) |
| `preferredName` | Preferred name | `String` | The preferred name for the user |
| `responsibilities` | Responsibilities | `String` | A list of the user's responsibilities |
| `schools` | Schools | `String` | A list of the user's schools attended |
| `skills` | Skills | `String` | A list of the user's skills |
| `state` | State | `String` | The state or province in the user's address |
| `streetAddress` | Street address | `String` | The street address of the user's place of business |
| `surname` | Surname | `String` | The user's surname (family name or last name) |
| `userPrincipalName` | User principal name | `String` | The user principal name (UPN) of the user (e.g., user@contoso.com) |
| `userType` | User type | `String` | The user type (Member or Guest) |

### UserProfile_V2

**Get user profile (V2)** — Retrieves a specific user's profile from Microsoft Graph by user principal name or user ID.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `id` | User (UPN) | `String` | ✓ | User principal name or user ID (e.g., user@contoso.com or a GUID) |
| `$select` | Select fields | `String` |  | Comma-separated list of user properties to retrieve |

**Output** — `Record` (31 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `aboutMe` | About me | `String` | A freeform text entry field for the user to describe themselves |
| `accountEnabled` | Account enabled | `Boolean` | Whether the account is enabled |
| `birthday` | Birthday | `DateTime` | The birthday of the user |
| `businessPhones` | Business phones | `String` | The telephone numbers for the user |
| `city` | City | `String` | The city in which the user is located |
| `companyName` | Company name | `String` | The company name which the user is associated |
| `country` | Country | `String` | The country/region in which the user is located |
| `department` | Department | `String` | The name for the department in which the user works |
| `displayName` | Display name | `String` | The name displayed in the address book for the user |
| `givenName` | Given name | `String` | The given name (first name) of the user |
| `hireDate` | Hire date | `DateTime` | The hire date of the user |
| `id` | Id | `String` | The unique identifier for the user |
| `interests` | Interests | `String` | A list of the user's interests |
| `jobTitle` | Job title | `String` | The user's job title |
| `mail` | Mail | `String` | The SMTP address for the user |
| `mailNickname` | Mail nickname | `String` | The mail alias for the user |
| `mobilePhone` | Mobile phone | `String` | The primary cellular telephone number for the user |
| `mySite` | My site | `String` | The URL for the user's personal site |
| `officeLocation` | Office location | `String` | The office location in the user's place of business |
| `pastProjects` | Past projects | `String` | A list of the user's past projects |
| `postalCode` | Postal code | `String` | The postal code for the user's postal address |
| `preferredLanguage` | Preferred language | `String` | The preferred language for the user (e.g., en-US) |
| `preferredName` | Preferred name | `String` | The preferred name for the user |
| `responsibilities` | Responsibilities | `String` | A list of the user's responsibilities |
| `schools` | Schools | `String` | A list of the user's schools attended |
| `skills` | Skills | `String` | A list of the user's skills |
| `state` | State | `String` | The state or province in the user's address |
| `streetAddress` | Street address | `String` | The street address of the user's place of business |
| `surname` | Surname | `String` | The user's surname (family name or last name) |
| `userPrincipalName` | User principal name | `String` | The user principal name (UPN) of the user (e.g., user@contoso.com) |
| `userType` | User type | `String` | The user type (Member or Guest) |

### UserPhoto

**Get user photo** — Retrieves the photo of the specified user or the current user.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `id` | User (UPN) | `String` |  | User principal name or user ID. Leave blank for current user. |

**Output** — `Record` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` | Photo | `String` | The user's photo as a base64-encoded string |

### SearchUser_V2

**Search for users (V2)** — Searches for users in the organization by name, email, or other attributes using Microsoft Graph.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `searchTerm` | Search term | `String` | ✓ | Search text to find users (searches displayName and mail) |
| `top` | Top | `Number` |  | Maximum number of results to return |
| `$select` | Select fields | `String` |  | Comma-separated list of user properties to retrieve |

**Output** — `Table` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Record` (4 fields) | List of matching users |

### DirectReports_V2

**Get direct reports (V2)** — Retrieves the direct reports of the specified user from Microsoft Graph.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `id` | User (UPN) | `String` | ✓ | User principal name or user ID |

**Output** — `Table` (1 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `value` |  | `Record` (5 fields) | List of direct reports |

### Manager_V2

**Get manager (V2)** — Retrieves the manager of the specified user from Microsoft Graph.

**Inputs**

| Name | Display name | Type | Req. | Description |
|---|---|---|---|---|
| `id` | User (UPN) | `String` | ✓ | User principal name or user ID |

**Output** — `Record` (6 fields):

| Name | Display name | Type | Description |
|---|---|---|---|
| `displayName` | Display name | `String` | The manager's display name |
| `mail` | Mail | `String` | The manager's email address |
| `id` | Id | `String` | The manager's unique identifier |
| `jobTitle` | Job title | `String` | The manager's job title |
| `department` | Department | `String` | The manager's department |
| `userPrincipalName` | User principal name | `String` | The manager's UPN |
