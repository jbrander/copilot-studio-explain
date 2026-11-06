# Tool (TaskDialog) actions & inputs

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

The `action:` block of a `TaskDialog` tool, and the `inputs:` it declares. See [08 Tools & connectors](../../08-tools-and-connectors.md) and [09 MCP & multi-agent](../../09-mcp-and-multi-agent.md).

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

**22 definitions** on this page (16 with a `kind:`). [A–Z index of all definitions](index.md).

<details><summary>Contents</summary>

[AIPluginConnectionReference](#aipluginconnectionreference), [AutomaticTaskInput](#automatictaskinput), [ComputerUsingAgentAllowlistNoKind](#computerusingagentallowlistnokind), [DisableToolAugmentation](#disabletoolaugmentation), [EnableToolAugmentation](#enabletoolaugmentation), [InvokeAgentTaskAction](#invokeagenttaskaction), [InvokeAIBuilderModelTaskAction](#invokeaibuildermodeltaskaction), [InvokeAIPluginTaskAction](#invokeaiplugintaskaction), [InvokeClientTaskAction](#invokeclienttaskaction), [InvokeComputerUsingAgentTaskAction](#invokecomputerusingagenttaskaction), [InvokeConnectedAgentTaskAction](#invokeconnectedagenttaskaction), [InvokeConnectorTaskAction](#invokeconnectortaskaction), [InvokeExternalAgentTaskAction](#invokeexternalagenttaskaction), [InvokeFlowTaskAction](#invokeflowtaskaction), [InvokerConnectionReference](#invokerconnectionreference), [InvokeSkillTaskAction](#invokeskilltaskaction), [MakerConnectionReference](#makerconnectionreference), [ManualTaskInput](#manualtaskinput), [QuestionInputSettingsNoKind](#questioninputsettingsnokind), [TaskAction](#taskaction), [TaskInput](#taskinput), [ToolAugmentation](#toolaugmentation)

</details>

### AIPluginConnectionReference


**One of:** [InvokerConnectionReference](#invokerconnectionreference), [MakerConnectionReference](#makerconnectionreference)

<sub>Used by: [InvokeAIPluginTaskAction](#invokeaiplugintaskaction)</sub>

### AutomaticTaskInput

YAML: `kind: AutomaticTaskInput`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `propertyName` | `string` | ✓ |  |  |
| `name` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `entity` | [EntityReference](entities.md#entityreference) |  |  |  |
| `shouldPromptUser` | `boolean` |  | `true` |  |
| `defaultValue` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `inputSettings` | [QuestionInputSettingsNoKind](#questioninputsettingsnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TaskInput](#taskinput)</sub>

### ComputerUsingAgentAllowlistNoKind

*Title:* ComputerUsingAgentAllowlist  
Same shape as [ComputerUsingAgentAllowlist](settings.md#computerusingagentallowlist) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `websites` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `desktopApps` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [InvokeComputerUsingAgentTaskAction](#invokecomputerusingagenttaskaction)</sub>

### DisableToolAugmentation

YAML: `kind: DisableToolAugmentation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [ToolAugmentation](#toolaugmentation)</sub>

### EnableToolAugmentation

YAML: `kind: EnableToolAugmentation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [ToolAugmentation](#toolaugmentation)</sub>

### InvokeAgentTaskAction

YAML: `kind: InvokeAgentTaskAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `gptComponentSchemaName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [TaskAction](#taskaction)</sub>

### InvokeAIBuilderModelTaskAction

YAML: `kind: InvokeAIBuilderModelTaskAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `aIModelId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [TaskAction](#taskaction)</sub>

### InvokeAIPluginTaskAction

YAML: `kind: InvokeAIPluginTaskAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `entityKey` | `string` | ✓ |  |  |
| `primaryConnectionReference` | [AIPluginConnectionReference](#aipluginconnectionreference) | ✓ |  |  |
| `secondaryConnectionReferences` | map of name → [AIPluginConnectionReference](#aipluginconnectionreference) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TaskAction](#taskaction)</sub>

### InvokeClientTaskAction

YAML: `kind: InvokeClientTaskAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `name` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `clientActionInputSchema` | [RecordNoKind](data-types.md#recordnokind) | ✓ |  |  |
| `clientActionResponseSchema` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TaskAction](#taskaction)</sub>

### InvokeComputerUsingAgentTaskAction

YAML: `kind: InvokeComputerUsingAgentTaskAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionReference` | `string` | ✓ |  |  |
| `connectionProperties` | [ConnectionPropertiesNoKind](settings.md#connectionpropertiesnokind) |  |  |  |
| `isElicitationEnabled` | `boolean` |  |  |  |
| `operationId` | `string` | ✓ |  |  |
| `instructions` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `model` | [AgentModelType](settings.md#agentmodeltype) |  |  |  |
| `initializeContext` | `=` expression \| object |  |  | accepts a `=` Power Fx expression |
| `allowlist` | [ComputerUsingAgentAllowlistNoKind](#computerusingagentallowlistnokind) |  |  |  |
| `credentialIds` | list of `string` |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TaskAction](#taskaction)</sub>

### InvokeConnectedAgentTaskAction

YAML: `kind: InvokeConnectedAgentTaskAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `botSchemaName` | `string` | ✓ |  |  |
| `historyType` | [HistoryType](settings.md#historytype) |  |  |  |
| `interruptionPolicy` | [InterruptionPolicyNoKind](settings.md#interruptionpolicynokind) |  |  |  |
| `shouldSendStartConversation` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TaskAction](#taskaction)</sub>

### InvokeConnectorTaskAction

YAML: `kind: InvokeConnectorTaskAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionReference` | `string` | ✓ |  |  |
| `connectionProperties` | [ConnectionPropertiesNoKind](settings.md#connectionpropertiesnokind) |  |  |  |
| `operationId` | `string` | ✓ |  |  |
| `dynamicInputSchema` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `dynamicOutputSchema` | [DataType](data-types.md#datatype) |  |  |  |
| `aIPluginOperationExportKey` | `string` |  |  |  |
| `requestTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  | `30000` | accepts a `=` Power Fx expression |
| `toolAugmentation` | [ToolAugmentation](#toolaugmentation) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TaskAction](#taskaction)</sub>

### InvokeExternalAgentTaskAction

YAML: `kind: InvokeExternalAgentTaskAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionReference` | `string` | ✓ |  |  |
| `connectionProperties` | [ConnectionPropertiesNoKind](settings.md#connectionpropertiesnokind) |  |  |  |
| `operationDetails` | [ProtocolMetadataBase](components.md#protocolmetadatabase) | ✓ |  |  |
| `historyType` | [HistoryType](settings.md#historytype) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TaskAction](#taskaction)</sub>

### InvokeFlowTaskAction

YAML: `kind: InvokeFlowTaskAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `flowId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [TaskAction](#taskaction)</sub>

### InvokerConnectionReference

YAML: `kind: InvokerConnectionReference`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [AIPluginConnectionReference](#aipluginconnectionreference)</sub>

### InvokeSkillTaskAction

YAML: `kind: InvokeSkillTaskAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `skillId` | `string` | ✓ |  |  |
| `actionId` | `string` | ✓ |  |  |
| `interruptionPolicy` | [InterruptionPolicyNoKind](settings.md#interruptionpolicynokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TaskAction](#taskaction)</sub>

### MakerConnectionReference

YAML: `kind: MakerConnectionReference`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionReference` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [AIPluginConnectionReference](#aipluginconnectionreference)</sub>

### ManualTaskInput

YAML: `kind: ManualTaskInput`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `propertyName` | `string` | ✓ |  |  |
| `value` | [ValueExpression](other.md#valueexpression) | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [TaskInput](#taskinput)</sub>

### QuestionInputSettingsNoKind

*Title:* QuestionInputSettings  
Same shape as [QuestionInputSettings](settings.md#questioninputsettings) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `prompt` | [ResponseTemplateNoKind](messages.md#responsetemplatenokind) |  |  |  |
| `repeatCount` | `0` \| `1` \| `2` \| [IntExpression](other.md#intexpression) |  | `2` |  |
| `validation` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `defaultValue` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `invalidPrompt` | [ResponseTemplateNoKind](messages.md#responsetemplatenokind) |  |  |  |
| `defaultValueResponse` | [ResponseTemplateNoKind](messages.md#responsetemplatenokind) |  |  |  |
| `unrecognizedPrompt` | [ResponseTemplateNoKind](messages.md#responsetemplatenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AutomaticTaskInput](#automatictaskinput)</sub>

### TaskAction


**One of:** [InvokeAIPluginTaskAction](#invokeaiplugintaskaction), [InvokeAgentTaskAction](#invokeagenttaskaction), [InvokeConnectedAgentTaskAction](#invokeconnectedagenttaskaction), [InvokeClientTaskAction](#invokeclienttaskaction), [InvokeConnectorTaskAction](#invokeconnectortaskaction), [InvokeExternalAgentTaskAction](#invokeexternalagenttaskaction), [InvokeComputerUsingAgentTaskAction](#invokecomputerusingagenttaskaction), [InvokeAIBuilderModelTaskAction](#invokeaibuildermodeltaskaction), [InvokeFlowTaskAction](#invokeflowtaskaction), [InvokeSkillTaskAction](#invokeskilltaskaction)

<sub>Used by: [GptComponentMetadata](components.md#gptcomponentmetadata), [GptComponentMetadataNoKind](components.md#gptcomponentmetadatanokind), [TaskDialog](dialogs.md#taskdialog)</sub>

### TaskInput


**One of:** [AutomaticTaskInput](#automatictaskinput), [ManualTaskInput](#manualtaskinput)

<sub>Used by: [AdaptiveDialog](dialogs.md#adaptivedialog), [AdaptiveDialogNoKind](dialogs.md#adaptivedialognokind), [AgentDialog](dialogs.md#agentdialog), [ConnectorTriggerDefinition](components.md#connectortriggerdefinition), [RecurrenceTriggerDefinition](components.md#recurrencetriggerdefinition), [TaskDialog](dialogs.md#taskdialog)</sub>

### ToolAugmentation


**One of:** [DisableToolAugmentation](#disabletoolaugmentation), [EnableToolAugmentation](#enabletoolaugmentation)

<sub>Used by: [InvokeConnectorTaskAction](#invokeconnectortaskaction)</sub>
