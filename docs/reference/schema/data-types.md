# Data types

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

Types used for variables, inputs, outputs and connector schemas. See [06 Variables, entities & Power Fx](../../06-variables-entities-powerfx.md).

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

**31 definitions** on this page (5 with a `kind:`). [A–Z index of all definitions](index.md).

<details><summary>Contents</summary>

[AiSensitivityLevel](#aisensitivitylevel), [Any](#any), [Blank](#blank), [Boolean](#boolean), [ClosedListOptionSet](#closedlistoptionset), [Color](#color), [ConnectorDynamicValuesConfigNoKind](#connectordynamicvaluesconfignokind), [DataType](#datatype), [Date](#date), [DateTime](#datetime), [DateTimeNoTimeZone](#datetimenotimezone), [EmbeddedOptionSet](#embeddedoptionset), [EmbeddedOptionSetNoKind](#embeddedoptionsetnokind), [File](#file), [Float](#float), [Guid](#guid), [Hyperlink](#hyperlink), [Number](#number), [OptionSetDataType](#optionsetdatatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind), [PropertyVisibility](#propertyvisibility), [Record](#record), [RecordNoKind](#recordnokind), [String](#string), [SystemOptionSet](#systemoptionset), [SystemOptionSetName](#systemoptionsetname), [Table](#table), [Time](#time), [UnresolvedDynamicType](#unresolveddynamictype), [Unspecified](#unspecified)

</details>

### AiSensitivityLevel


**String enum.** Allowed values: `None`, `Low`, `High`

<sub>Used by: [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### Any


**One of:** `"Any"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### Blank


**One of:** `"Blank"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### Boolean


**One of:** `"Boolean"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### ClosedListOptionSet

YAML: `kind: ClosedListOptionSet`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sensitivityLevel` | [SensitivityLevel](settings.md#sensitivitylevel) |  |  |  |
| `schemaName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DataType](#datatype), [OptionSetDataType](#optionsetdatatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### Color


**One of:** `"Color"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### ConnectorDynamicValuesConfigNoKind

*Title:* ConnectorDynamicValuesConfig  
Same shape as [ConnectorDynamicValuesConfig](components.md#connectordynamicvaluesconfig) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `capability` | [ConnectorDynamicValuesCapability](components.md#connectordynamicvaluescapability) |  |  |  |

_No other properties are allowed._

<sub>Used by: [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### DataType


**One of:** [ClosedListOptionSet](#closedlistoptionset), [EmbeddedOptionSet](#embeddedoptionset), [SystemOptionSet](#systemoptionset), [Unspecified](#unspecified), [Blank](#blank), [Boolean](#boolean), [Number](#number), [Float](#float), [String](#string), [File](#file), [Hyperlink](#hyperlink), [Color](#color), [Guid](#guid), [Time](#time), [Date](#date), [DateTime](#datetime), [DateTimeNoTimeZone](#datetimenotimezone), [Any](#any), [UnresolvedDynamicType](#unresolveddynamictype), [Table](#table), [Record](#record)

<sub>Used by: [AIPluginOperation](components.md#aipluginoperation), [AIPluginOperationNoKind](components.md#aipluginoperationnokind), [ConnectorOperation](components.md#connectoroperation), [ConnectorOperationNoKind](components.md#connectoroperationnokind), [ConnectorTriggerOperation](components.md#connectortriggeroperation), [ConnectorTriggerOperationNoKind](components.md#connectortriggeroperationnokind), [ContinueConversationConnectorTrigger](triggers.md#continueconversationconnectortrigger), [ConversionMetadata](components.md#conversionmetadata), [ConversionMetadataNoKind](dialogs.md#conversionmetadatanokind), [CreateConversationConnectorTrigger](triggers.md#createconversationconnectortrigger), [EvaluationSet](evaluation.md#evaluationset), [ExternalEntity](dialogs.md#externalentity), [HttpRequestAction](actions.md#httprequestaction), [InvokeConnectorAction](actions.md#invokeconnectoraction), [InvokeConnectorTaskAction](task-actions.md#invokeconnectortaskaction), [McpToolDefinition](components.md#mcptooldefinition), [McpToolDefinitionNoKind](components.md#mcptooldefinitionnokind), [ParseValue](actions.md#parsevalue), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind) … and 1 more</sub>

### Date


**One of:** `"Date"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### DateTime


**One of:** `"DateTime"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### DateTimeNoTimeZone


**One of:** `"DateTimeNoTimeZone"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### EmbeddedOptionSet

YAML: `kind: EmbeddedOptionSet`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sensitivityLevel` | [SensitivityLevel](settings.md#sensitivitylevel) |  |  |  |
| `dialogSchemaName` | `string` | ✓ |  |  |
| `triggerId` | `string` | ✓ |  |  |
| `actionId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DataType](#datatype), [OptionSetDataType](#optionsetdatatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### EmbeddedOptionSetNoKind

*Title:* EmbeddedOptionSet  
Same shape as [EmbeddedOptionSet](#embeddedoptionset) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sensitivityLevel` | [SensitivityLevel](settings.md#sensitivitylevel) |  |  |  |
| `dialogSchemaName` | `string` | ✓ |  |  |
| `triggerId` | `string` | ✓ |  |  |
| `actionId` | `string` | ✓ |  |  |

_No other properties are allowed._

### File


**One of:** `"File"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### Float


**One of:** `"Float"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### Guid


**One of:** `"Guid"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### Hyperlink


**One of:** `"Hyperlink"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### Number


**One of:** `"Number"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### OptionSetDataType


**One of:** [ClosedListOptionSet](#closedlistoptionset), [EmbeddedOptionSet](#embeddedoptionset), [SystemOptionSet](#systemoptionset)

### PropertyInfo


**One of:** [ClosedListOptionSet](#closedlistoptionset), [EmbeddedOptionSet](#embeddedoptionset), [SystemOptionSet](#systemoptionset), [Unspecified](#unspecified), [Blank](#blank), [Boolean](#boolean), [Number](#number), [Float](#float), [String](#string), [File](#file), [Hyperlink](#hyperlink), [Color](#color), [Guid](#guid), [Time](#time), [Date](#date), [DateTime](#datetime), [DateTimeNoTimeZone](#datetimenotimezone), [Any](#any), [UnresolvedDynamicType](#unresolveddynamictype), [Table](#table), [Record](#record), object

### PropertyInfoNoKind


**One of:** [ClosedListOptionSet](#closedlistoptionset), [EmbeddedOptionSet](#embeddedoptionset), [SystemOptionSet](#systemoptionset), [Unspecified](#unspecified), [Blank](#blank), [Boolean](#boolean), [Number](#number), [Float](#float), [String](#string), [File](#file), [Hyperlink](#hyperlink), [Color](#color), [Guid](#guid), [Time](#time), [Date](#date), [DateTime](#datetime), [DateTimeNoTimeZone](#datetimenotimezone), [Any](#any), [UnresolvedDynamicType](#unresolveddynamictype), [Table](#table), [Record](#record), object

<sub>Used by: [Record](#record), [RecordNoKind](#recordnokind), [Table](#table)</sub>

### PropertyVisibility


**String enum.** Allowed values: `Important`, `Advanced`, `Internal`

<sub>Used by: [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### Record

YAML: `kind: Record`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sensitivityLevel` | [SensitivityLevel](settings.md#sensitivitylevel) |  |  |  |
| `properties` | map of name → [PropertyInfoNoKind](#propertyinfonokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### RecordNoKind

*Title:* Record  
Same shape as [Record](#record) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sensitivityLevel` | [SensitivityLevel](settings.md#sensitivitylevel) |  |  |  |
| `properties` | map of name → [PropertyInfoNoKind](#propertyinfonokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [A2AExtensions](components.md#a2aextensions), [A2AExtensionsNoKind](components.md#a2aextensionsnokind), [AIModelDefinition](components.md#aimodeldefinition), [AIModelDefinitionNoKind](components.md#aimodeldefinitionnokind), [AIPluginOperation](components.md#aipluginoperation), [AIPluginOperationNoKind](components.md#aipluginoperationnokind), [ActivityProtocolMetadata](components.md#activityprotocolmetadata), [AdaptiveCardPrompt](actions.md#adaptivecardprompt), [AdaptiveDialog](dialogs.md#adaptivedialog), [AdaptiveDialogNoKind](dialogs.md#adaptivedialognokind), [AgentDialog](dialogs.md#agentdialog), [AgentToAgentProtocolMetadata](components.md#agenttoagentprotocolmetadata), [CloudFlowDefinition](components.md#cloudflowdefinition), [CloudFlowDefinitionNoKind](components.md#cloudflowdefinitionnokind), [ConnectedAgentDefinition](components.md#connectedagentdefinition), [ConnectedAgentDefinitionNoKind](components.md#connectedagentdefinitionnokind), [ConnectorOperation](components.md#connectoroperation), [ConnectorOperationNoKind](components.md#connectoroperationnokind), [ConnectorTriggerOperation](components.md#connectortriggeroperation), [ConnectorTriggerOperationNoKind](components.md#connectortriggeroperationnokind) … and 16 more</sub>

### String


**One of:** `"String"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### SystemOptionSet

YAML: `kind: SystemOptionSet`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sensitivityLevel` | [SensitivityLevel](settings.md#sensitivitylevel) |  |  |  |
| `name` | [SystemOptionSetName](#systemoptionsetname) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DataType](#datatype), [OptionSetDataType](#optionsetdatatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### SystemOptionSetName


**String enum.** Allowed values: `RecurrenceFrequency`, `SkipQuestionMode`, `VariablesToClear`, `SignInReason`, `MultipleTopicsMatchedReason`, `HttpMethod`, `NetworkIsolationType`, `ContentModerationLevel`, `ToneType`, `TemplateResponseMode`, `AgentMessageRole`, `AgentMessageContentType`, `AgentMessageSortOrder`, `Locale`, `DTMFTerminationKey`, `InterruptionEagernessLevel`, `TurnDetectionType`, `DigitalTurnDetectionType`, `DigitalNoiseReductionMode`, `DigitalListeningMode`, `DefaultValueMissingAction`, `ActivityType`, `ChannelId`, `PlanStepType`, `InvocationSource`, `UserInputType`

<sub>Used by: [SystemOptionSet](#systemoptionset)</sub>

### Table

YAML: `kind: Table`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sensitivityLevel` | [SensitivityLevel](settings.md#sensitivitylevel) |  |  |  |
| `properties` | map of name → [PropertyInfoNoKind](#propertyinfonokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### Time


**One of:** `"Time"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### UnresolvedDynamicType


**One of:** `"UnresolvedDynamicType"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>

### Unspecified


**One of:** `"Unspecified"`, object

<sub>Used by: [DataType](#datatype), [PropertyInfo](#propertyinfo), [PropertyInfoNoKind](#propertyinfonokind)</sub>
