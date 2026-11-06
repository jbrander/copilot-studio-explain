# Components & agent definitions

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

Top-level file kinds and solution components: the agent itself (`GptComponentMetadata`), global variables, knowledge-source components, external triggers, connection references, and so on. See [02 Project layout](../../02-project-layout.md) and [03 The agent](../../03-agent.md).

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

**168 definitions** on this page (84 with a `kind:`). [A–Z index of all definitions](index.md).

<details><summary>Contents</summary>

[A2AExtensions](#a2aextensions), [A2AExtensionsNoKind](#a2aextensionsnokind), [ActivityProtocolMetadata](#activityprotocolmetadata), [AdaptationManagementServiceTrainingMetadata](#adaptationmanagementservicetrainingmetadata), [AdaptationManagementServiceTrainingMetadataNoKind](#adaptationmanagementservicetrainingmetadatanokind), [AgentDefinition](#agentdefinition), [AgentToAgentProtocolMetadata](#agenttoagentprotocolmetadata), [AIModelDefinition](#aimodeldefinition), [AIModelDefinitionNoKind](#aimodeldefinitionnokind), [AIPluginComponent](#aiplugincomponent), [AIPluginOperation](#aipluginoperation), [AIPluginOperationComponent](#aipluginoperationcomponent), [AIPluginOperationNoKind](#aipluginoperationnokind), [ApplicationDetails](#applicationdetails), [ApplicationDetailsNoKind](#applicationdetailsnokind), [AuditInfo](#auditinfo), [AuditInfoNoKind](#auditinfonokind), [BotAccessControlPolicy](#botaccesscontrolpolicy), [BotAuthenticationConfiguration](#botauthenticationconfiguration), [BotAuthenticationConfigurationNoKind](#botauthenticationconfigurationnokind), [BotAuthenticationMode](#botauthenticationmode), [BotAuthenticationTrigger](#botauthenticationtrigger), [BotComponentBase](#botcomponentbase), [BotComponentCollection](#botcomponentcollection), [BotComponentCollectionConfiguration](#botcomponentcollectionconfiguration), [BotComponentCollectionConfigurationNoKind](#botcomponentcollectionconfigurationnokind), [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotComponentCollectionNoKind](#botcomponentcollectionnokind), [BotComponentStatusReason](#botcomponentstatusreason), [BotConfiguration](#botconfiguration), [BotConfigurationNoKind](#botconfigurationnokind), [BotDefinition](#botdefinition), [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind), [BotRegistrationDetails](#botregistrationdetails), [BotRegistrationDetailsNoKind](#botregistrationdetailsnokind), [BotSettingsComponent](#botsettingscomponent), [BotSynchronizationDetails](#botsynchronizationdetails), [BotSynchronizationDetailsNoKind](#botsynchronizationdetailsnokind), [Category](#category), [CategoryNoKind](#categorynokind), [ChannelDefinition](#channeldefinition), [ChannelDefinitionNoKind](#channeldefinitionnokind), [ChannelId](#channelid), [CloudFlowConnectionType](#cloudflowconnectiontype), [CloudFlowDefinition](#cloudflowdefinition), [CloudFlowDefinitionNoKind](#cloudflowdefinitionnokind), [CloudFlowTriggerType](#cloudflowtriggertype), [ConnectedAgentDefinition](#connectedagentdefinition), [ConnectedAgentDefinitionNoKind](#connectedagentdefinitionnokind), [ConnectionMode](#connectionmode), [ConnectionParametersSet](#connectionparametersset), [ConnectionParametersSetNoKind](#connectionparameterssetnokind), [ConnectionReference](#connectionreference), [ConnectionReferenceNameInFlowToConnectionParametersSetMapping](#connectionreferencenameinflowtoconnectionparameterssetmapping), [ConnectionReferenceNoKind](#connectionreferencenokind), [ConnectionReferencesSourceFile](#connectionreferencessourcefile), [ConnectorDefinition](#connectordefinition), [ConnectorDefinitionComponent](#connectordefinitioncomponent), [ConnectorDefinitionNoKind](#connectordefinitionnokind), [ConnectorDefinitionType](#connectordefinitiontype), [ConnectorDynamicValuesCapability](#connectordynamicvaluescapability), [ConnectorDynamicValuesConfig](#connectordynamicvaluesconfig), [ConnectorOperation](#connectoroperation), [ConnectorOperationNoKind](#connectoroperationnokind), [ConnectorTriggerDefinition](#connectortriggerdefinition), [ConnectorTriggerOperation](#connectortriggeroperation), [ConnectorTriggerOperationNoKind](#connectortriggeroperationnokind), [ContentShareContextNoKind](#contentsharecontextnokind), [ConversationStarter](#conversationstarter), [ConversationStarterNoKind](#conversationstarternokind), [ConversionMetadata](#conversionmetadata), [CopilotChatApplicationDetails](#copilotchatapplicationdetails), [CopilotChatApplicationDetailsNoKind](#copilotchatapplicationdetailsnokind), [CustomEntityComponent](#customentitycomponent), [CustomMetricCategory](#custommetriccategory), [CustomMetricCategoryNoKind](#custommetriccategorynokind), [CustomMetricDefinition](#custommetricdefinition), [CustomMetricDefinitionComponent](#custommetricdefinitioncomponent), [CustomMetricDefinitionNoKind](#custommetricdefinitionnokind), [DataverseTableSearch](#dataversetablesearch), [DataverseTableSearchEntityColumnSynonym](#dataversetablesearchentitycolumnsynonym), [DataverseTableSearchEntityColumnSynonymNoKind](#dataversetablesearchentitycolumnsynonymnokind), [DataverseTableSearchEntityConfiguration](#dataversetablesearchentityconfiguration), [DataverseTableSearchEntityConfigurationNoKind](#dataversetablesearchentityconfigurationnokind), [DataverseTableSearchGlossaryConfiguration](#dataversetablesearchglossaryconfiguration), [DataverseTableSearchGlossaryConfigurationNoKind](#dataversetablesearchglossaryconfigurationnokind), [DefaultResponseMode](#defaultresponsemode), [DiagnosticType](#diagnostictype), [DialogComponent](#dialogcomponent), [DialogComponentNoKind](#dialogcomponentnokind), [DialogComponentReference](#dialogcomponentreference), [DlpComponentType](#dlpcomponenttype), [EnvironmentVariableDefinition](#environmentvariabledefinition), [EnvironmentVariableDefinitionNoKind](#environmentvariabledefinitionnokind), [EnvironmentVariableReference](#environmentvariablereference), [EnvironmentVariableReferenceNoKind](#environmentvariablereferencenokind), [EnvironmentVariableSecretReadStatus](#environmentvariablesecretreadstatus), [EnvironmentVariableType](#environmentvariabletype), [EnvironmentVariableValue](#environmentvariablevalue), [EnvironmentVariableValueNoKind](#environmentvariablevaluenokind), [ExternalTriggerComponent](#externaltriggercomponent), [ExternalTriggerConfiguration](#externaltriggerconfiguration), [ExternalTriggerConfigurationNoKind](#externaltriggerconfigurationnokind), [ExternalTriggerDefinition](#externaltriggerdefinition), [FederatedTableSearch](#federatedtablesearch), [FileAttachmentComponent](#fileattachmentcomponent), [FileAttachmentComponentMetadataNoKind](#fileattachmentcomponentmetadatanokind), [GlobalVariableComponent](#globalvariablecomponent), [GptCapabilitiesNoKind](#gptcapabilitiesnokind), [GptComponent](#gptcomponent), [GptComponentMetadata](#gptcomponentmetadata), [GptComponentMetadataNoKind](#gptcomponentmetadatanokind), [KnowledgeSourceComponent](#knowledgesourcecomponent), [KnowledgeSourceConfiguration](#knowledgesourceconfiguration), [KnowledgeSourceConfigurationNoKind](#knowledgesourceconfigurationnokind), [LegacyOrUnknownComponent](#legacyorunknowncomponent), [ManagedProperties](#managedproperties), [ManagedPropertiesNoKind](#managedpropertiesnokind), [ManifestResources](#manifestresources), [McpToolDefinition](#mcptooldefinition), [McpToolDefinitionNoKind](#mcptooldefinitionnokind), [Microsoft365ApplicationDetails](#microsoft365applicationdetails), [Microsoft365ApplicationDetailsNoKind](#microsoft365applicationdetailsnokind), [ModelContextProtocolMetadata](#modelcontextprotocolmetadata), [PlanStepType](#plansteptype), [PropertyOverride](#propertyoverride), [ProtocolMetadataBase](#protocolmetadatabase), [ProvisioningState](#provisioningstate), [ProvisioningStatus](#provisioningstatus), [PublishResult](#publishresult), [PublishResultDetails](#publishresultdetails), [PublishResultDetailsNoKind](#publishresultdetailsnokind), [RecurrenceFrequency](#recurrencefrequency), [RecurrenceTriggerDefinition](#recurrencetriggerdefinition), [ReferenceItemSourceFile](#referenceitemsourcefile), [ReferenceItemSourceFileNoKind](#referenceitemsourcefilenokind), [ReferencesSourceFile](#referencessourcefile), [ReferenceType](#referencetype), [RuntimeProvider](#runtimeprovider), [SkillActionDefinition](#skillactiondefinition), [SkillActionDefinitionNoKind](#skillactiondefinitionnokind), [SkillComponent](#skillcomponent), [SkillDefinition](#skilldefinition), [SkillDefinitionNoKind](#skilldefinitionnokind), [SolutionComponent](#solutioncomponent), [StateCode](#statecode), [SuccessfulPublishResultDetails](#successfulpublishresultdetails), [SuccessfulPublishResultDetailsNoKind](#successfulpublishresultdetailsnokind), [SynchronizationState](#synchronizationstate), [SynchronizationStateNoKind](#synchronizationstatenokind), [SynchronizedChannel](#synchronizedchannel), [SynchronizedChannelNoKind](#synchronizedchannelnokind), [SynchronizingState](#synchronizingstate), [SystemUser](#systemuser), [TableSearchType](#tablesearchtype), [TeamsApplicationDetails](#teamsapplicationdetails), [TeamsApplicationDetailsNoKind](#teamsapplicationdetailsnokind), [TestCaseComponent](#testcasecomponent), [ToolManifestResource](#toolmanifestresource), [ToolsSelectionMode](#toolsselectionmode), [TranslationsComponent](#translationscomponent), [TranslationsComponentNoKind](#translationscomponentnokind), [UseAllTools](#usealltools), [UserPrivilege](#userprivilege), [UseSpecificTools](#usespecifictools), [VariableNoKind](#variablenokind), [WorkflowExternalTrigger](#workflowexternaltrigger)

</details>

### A2AExtensions

YAML: `kind: A2AExtensions`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |

_No other properties are allowed._

### A2AExtensionsNoKind

*Title:* A2AExtensions  
Same shape as [A2AExtensions](#a2aextensions) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentToAgentProtocolMetadata](#agenttoagentprotocolmetadata)</sub>

### ActivityProtocolMetadata

YAML: `kind: ActivityProtocolMetadata`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `operationId` | `string` | ✓ |  |  |
| `dynamicInputSchema` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `interruptionPolicy` | [InterruptionPolicyNoKind](settings.md#interruptionpolicynokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ProtocolMetadataBase](#protocolmetadatabase)</sub>

### AdaptationManagementServiceTrainingMetadata

YAML: `kind: AdaptationManagementServiceTrainingMetadata`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelId` | `string` |  |  |  |
| `contextTagUrn` | `string` |  |  |  |
| `buildId` | `string` |  |  |  |
| `tagId` | `string` |  |  |  |
| `trainingStatusCode` | `string` |  |  |  |
| `trainingErrorMessage` | `string` |  |  |  |

_No other properties are allowed._

### AdaptationManagementServiceTrainingMetadataNoKind

*Title:* AdaptationManagementServiceTrainingMetadata  
Same shape as [AdaptationManagementServiceTrainingMetadata](#adaptationmanagementservicetrainingmetadata) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelId` | `string` |  |  |  |
| `contextTagUrn` | `string` |  |  |  |
| `buildId` | `string` |  |  |  |
| `tagId` | `string` |  |  |  |
| `trainingStatusCode` | `string` |  |  |  |
| `trainingErrorMessage` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [AdaptationManagementServiceModelType](settings.md#adaptationmanagementservicemodeltype), [AdaptationManagementServiceModelTypeNoKind](settings.md#adaptationmanagementservicemodeltypenokind)</sub>

### AgentDefinition


**One of:** [GptComponentMetadata](#gptcomponentmetadata), [GptComponentMetadataNoKind](#gptcomponentmetadatanokind), [ConnectedAgentDefinition](#connectedagentdefinition)

### AgentToAgentProtocolMetadata

YAML: `kind: AgentToAgentProtocolMetadata`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `operationId` | `string` | ✓ |  |  |
| `dynamicInputSchema` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `interruptionPolicy` | [InterruptionPolicyNoKind](settings.md#interruptionpolicynokind) |  |  |  |
| `extensions` | [A2AExtensionsNoKind](#a2aextensionsnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ProtocolMetadataBase](#protocolmetadatabase)</sub>

### AIModelDefinition

YAML: `kind: AIModelDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `name` | `string` | ✓ |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### AIModelDefinitionNoKind

*Title:* AIModelDefinition  
Same shape as [AIModelDefinition](#aimodeldefinition) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `name` | `string` | ✓ |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotDefinition](#botdefinition)</sub>

### AIPluginComponent

YAML: `kind: AIPluginComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `name` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### AIPluginOperation

YAML: `kind: AIPluginOperation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `exportKey` | `string` | ✓ |  |  |
| `pluginName` | `string` |  |  |  |
| `operationId` | `string` |  |  |  |
| `displayName` | `string` | ✓ |  |  |
| `description` | `string` |  |  |  |
| `iconUri` | `string` |  |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) | ✓ |  |  |
| `outputType` | [DataType](data-types.md#datatype) | ✓ |  |  |

_No other properties are allowed._

### AIPluginOperationComponent

YAML: `kind: AIPluginOperationComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `operationId` | `string` | ✓ |  |  |
| `displayName` | `string` |  |  |  |
| `exportKey` | `string` |  |  |  |
| `parentAIPluginId` | `string` |  |  |  |
| `parentAIPluginName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### AIPluginOperationNoKind

*Title:* AIPluginOperation  
Same shape as [AIPluginOperation](#aipluginoperation) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `exportKey` | `string` | ✓ |  |  |
| `pluginName` | `string` |  |  |  |
| `operationId` | `string` |  |  |  |
| `displayName` | `string` | ✓ |  |  |
| `description` | `string` |  |  |  |
| `iconUri` | `string` |  |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) | ✓ |  |  |
| `outputType` | [DataType](data-types.md#datatype) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotDefinition](#botdefinition)</sub>

### ApplicationDetails

YAML: `kind: ApplicationDetails`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `teams` | [TeamsApplicationDetailsNoKind](#teamsapplicationdetailsnokind) |  |  |  |
| `microsoft365` | [Microsoft365ApplicationDetailsNoKind](#microsoft365applicationdetailsnokind) |  |  |  |
| `copilotChat` | [CopilotChatApplicationDetailsNoKind](#copilotchatapplicationdetailsnokind) |  |  |  |

_No other properties are allowed._

### ApplicationDetailsNoKind

*Title:* ApplicationDetails  
Same shape as [ApplicationDetails](#applicationdetails) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `teams` | [TeamsApplicationDetailsNoKind](#teamsapplicationdetailsnokind) |  |  |  |
| `microsoft365` | [Microsoft365ApplicationDetailsNoKind](#microsoft365applicationdetailsnokind) |  |  |  |
| `copilotChat` | [CopilotChatApplicationDetailsNoKind](#copilotchatapplicationdetailsnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind)</sub>

### AuditInfo

YAML: `kind: AuditInfo`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `createdTimeUtc` | `string` |  |  |  |
| `modifiedTimeUtc` | `string` |  |  |  |
| `createdBy` | `string` |  |  |  |
| `modifiedBy` | `string` |  |  |  |

_No other properties are allowed._

### AuditInfoNoKind

*Title:* AuditInfo  
Same shape as [AuditInfo](#auditinfo) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `createdTimeUtc` | `string` |  |  |  |
| `modifiedTimeUtc` | `string` |  |  |  |
| `createdBy` | `string` |  |  |  |
| `modifiedBy` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [AIModelDefinition](#aimodeldefinition), [AIModelDefinitionNoKind](#aimodeldefinitionnokind), [AIPluginComponent](#aiplugincomponent), [AIPluginOperationComponent](#aipluginoperationcomponent), [BotComponentCollection](#botcomponentcollection), [BotComponentCollectionNoKind](#botcomponentcollectionnokind), [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind), [BotSettingsComponent](#botsettingscomponent), [CloudFlowDefinition](#cloudflowdefinition), [CloudFlowDefinitionNoKind](#cloudflowdefinitionnokind), [ConnectionReference](#connectionreference), [ConnectionReferenceNoKind](#connectionreferencenokind), [ConnectorDefinitionComponent](#connectordefinitioncomponent), [CustomEntityComponent](#customentitycomponent), [CustomMetricDefinitionComponent](#custommetricdefinitioncomponent), [DataverseTableSearch](#dataversetablesearch), [DataverseTableSearchEntityColumnSynonym](#dataversetablesearchentitycolumnsynonym), [DataverseTableSearchEntityColumnSynonymNoKind](#dataversetablesearchentitycolumnsynonymnokind), [DataverseTableSearchEntityConfiguration](#dataversetablesearchentityconfiguration) … and 22 more</sub>

### BotAccessControlPolicy


**String enum.** Allowed values: `Any`, `ChatbotReaders`, `GroupMembership`, `AnyMultiTenant`

<sub>Used by: [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind)</sub>

### BotAuthenticationConfiguration

YAML: `kind: BotAuthenticationConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionName` | `string` |  |  |  |

_No other properties are allowed._

### BotAuthenticationConfigurationNoKind

*Title:* BotAuthenticationConfiguration  
Same shape as [BotAuthenticationConfiguration](#botauthenticationconfiguration) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind)</sub>

### BotAuthenticationMode


**String enum.** Allowed values: `Unspecified`, `None`, `Integrated`, `CustomAzureActiveDirectory`, `GenericOAuth2`

<sub>Used by: [BotComponentCollectionConfiguration](#botcomponentcollectionconfiguration), [BotComponentCollectionConfigurationNoKind](#botcomponentcollectionconfigurationnokind), [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind), [SuccessfulPublishResultDetails](#successfulpublishresultdetails), [SuccessfulPublishResultDetailsNoKind](#successfulpublishresultdetailsnokind)</sub>

### BotAuthenticationTrigger


**String enum.** Allowed values: `AsNeeded`, `Always`

<sub>Used by: [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind)</sub>

### BotComponentBase


**One of:** [DialogComponent](#dialogcomponent), [TranslationsComponent](#translationscomponent), [CustomEntityComponent](#customentitycomponent), [GlobalVariableComponent](#globalvariablecomponent), [SkillComponent](#skillcomponent), [FileAttachmentComponent](#fileattachmentcomponent), [GptComponent](#gptcomponent), [KnowledgeSourceComponent](#knowledgesourcecomponent), [ExternalTriggerComponent](#externaltriggercomponent), [BotSettingsComponent](#botsettingscomponent), [TestCaseComponent](#testcasecomponent), [LegacyOrUnknownComponent](#legacyorunknowncomponent), [CustomMetricDefinitionComponent](#custommetricdefinitioncomponent)

<sub>Used by: [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotDefinition](#botdefinition)</sub>

### BotComponentCollection

YAML: `kind: BotComponentCollection`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `id` | `string` | ✓ |  |  |
| `displayName` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `configuration` | [BotComponentCollectionConfigurationNoKind](#botcomponentcollectionconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### BotComponentCollectionConfiguration

YAML: `kind: BotComponentCollectionConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `primaryLanguage` | [Language](settings.md#language) |  |  |  |
| `secondaryLanguages` | list of `Language` |  |  |  |
| `requiredAuthenticationMode` | [BotAuthenticationMode](#botauthenticationmode) |  |  |  |
| `includeIvrVariables` | `boolean` |  |  |  |
| `scopedToBots` | list of `string` |  |  |  |

_No other properties are allowed._

### BotComponentCollectionConfigurationNoKind

*Title:* BotComponentCollectionConfiguration  
Same shape as [BotComponentCollectionConfiguration](#botcomponentcollectionconfiguration) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `primaryLanguage` | [Language](settings.md#language) |  |  |  |
| `secondaryLanguages` | list of `Language` |  |  |  |
| `requiredAuthenticationMode` | [BotAuthenticationMode](#botauthenticationmode) |  |  |  |
| `includeIvrVariables` | `boolean` |  |  |  |
| `scopedToBots` | list of `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentCollection](#botcomponentcollection), [BotComponentCollectionNoKind](#botcomponentcollectionnokind)</sub>

### BotComponentCollectionDefinition

YAML: `kind: BotComponentCollectionDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `components` | list of [BotComponentBase](#botcomponentbase) |  |  |  |
| `environmentVariables` | list of [EnvironmentVariableDefinitionNoKind](#environmentvariabledefinitionnokind) |  |  |  |
| `flows` | list of [CloudFlowDefinitionNoKind](#cloudflowdefinitionnokind) |  |  |  |
| `dataverseTableSearchs` | list of [DataverseTableSearchBase](knowledge.md#dataversetablesearchbase) |  |  |  |
| `dataverseTableSearchGlossaryConfigurations` | list of [DataverseTableSearchGlossaryConfigurationNoKind](#dataversetablesearchglossaryconfigurationnokind) |  |  |  |
| `dataverseTableSearchEntityConfigurations` | list of [DataverseTableSearchEntityConfigurationNoKind](#dataversetablesearchentityconfigurationnokind) |  |  |  |
| `dataverseTableSearchEntityColumnSynonyms` | list of [DataverseTableSearchEntityColumnSynonymNoKind](#dataversetablesearchentitycolumnsynonymnokind) |  |  |  |
| `connectionReferences` | list of [ConnectionReferenceNoKind](#connectionreferencenokind) |  |  |  |
| `connectorDefinitions` | list of [ConnectorDefinitionNoKind](#connectordefinitionnokind) |  |  |  |
| `aIModelDefinitions` | list of [AIModelDefinitionNoKind](#aimodeldefinitionnokind) |  |  |  |
| `aIPluginOperations` | list of [AIPluginOperationNoKind](#aipluginoperationnokind) |  |  |  |
| `connectedAgentDefinitions` | list of [ConnectedAgentDefinitionNoKind](#connectedagentdefinitionnokind) |  |  |  |
| `componentCollection` | [BotComponentCollectionNoKind](#botcomponentcollectionnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DefinitionBase](other.md#definitionbase)</sub>

### BotComponentCollectionNoKind

*Title:* BotComponentCollection  
Same shape as [BotComponentCollection](#botcomponentcollection) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `id` | `string` | ✓ |  |  |
| `displayName` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `configuration` | [BotComponentCollectionConfigurationNoKind](#botcomponentcollectionconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotDefinition](#botdefinition)</sub>

### BotComponentStatusReason


**String enum.** Allowed values: `Active`, `Inactive`

<sub>Used by: [BotSettingsComponent](#botsettingscomponent), [CustomEntityComponent](#customentitycomponent), [CustomMetricDefinitionComponent](#custommetricdefinitioncomponent), [DialogComponent](#dialogcomponent), [DialogComponentNoKind](#dialogcomponentnokind), [ExternalTriggerComponent](#externaltriggercomponent), [FileAttachmentComponent](#fileattachmentcomponent), [GlobalVariableComponent](#globalvariablecomponent), [GptComponent](#gptcomponent), [KnowledgeSourceComponent](#knowledgesourcecomponent), [LegacyOrUnknownComponent](#legacyorunknowncomponent), [SkillComponent](#skillcomponent), [TestCaseComponent](#testcasecomponent), [TranslationsComponent](#translationscomponent), [TranslationsComponentNoKind](#translationscomponentnokind)</sub>

### BotConfiguration

YAML: `kind: BotConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `categories` | list of [CategoryNoKind](#categorynokind) |  |  |  |
| `channels` | list of [ChannelDefinitionNoKind](#channeldefinitionnokind) |  |  |  |
| `settings` | map of name → [DataValue](other.md#datavalue) |  |  |  |
| `isTelephonyEnabled` | `boolean` |  | `false` |  |
| `sensitiveDataMaskingSettings` | [SensitiveDataMaskingSettings](settings.md#sensitivedatamaskingsettings) |  |  |  |
| `isAgentConnectable` | `boolean` |  | `true` |  |
| `voiceProcessingMode` | [VoiceProcessingModeBase](settings.md#voiceprocessingmodebase) |  |  |  |
| `botSpeechSettings` | [BotSpeechSettingsNoKind](settings.md#botspeechsettingsnokind) |  |  |  |
| `networkSettings` | [NetworkSettingsNoKind](settings.md#networksettingsnokind) |  |  |  |
| `generationSettings` | [GenerationSettingsNoKind](settings.md#generationsettingsnokind) |  |  |  |
| `publishOnCreate` | `boolean` |  | `false` |  |
| `publishOnImport` | `boolean` |  | `true` |  |
| `gPTSettings` | [GPTSettingsNoKind](settings.md#gptsettingsnokind) |  |  |  |
| `isLightweightBot` | `boolean` |  | `false` |  |
| `aISettings` | [AISettingsNoKind](settings.md#aisettingsnokind) |  |  |  |
| `recognizer` | [RecognizerBase](dialogs.md#recognizerbase) |  |  |  |
| `analyticsSettings` | [AnalyticsSettingsNoKind](settings.md#analyticssettingsnokind) |  |  |  |

_No other properties are allowed._

### BotConfigurationNoKind

*Title:* BotConfiguration  
Same shape as [BotConfiguration](#botconfiguration) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `categories` | list of [CategoryNoKind](#categorynokind) |  |  |  |
| `channels` | list of [ChannelDefinitionNoKind](#channeldefinitionnokind) |  |  |  |
| `settings` | map of name → [DataValue](other.md#datavalue) |  |  |  |
| `isTelephonyEnabled` | `boolean` |  | `false` |  |
| `sensitiveDataMaskingSettings` | [SensitiveDataMaskingSettings](settings.md#sensitivedatamaskingsettings) |  |  |  |
| `isAgentConnectable` | `boolean` |  | `true` |  |
| `voiceProcessingMode` | [VoiceProcessingModeBase](settings.md#voiceprocessingmodebase) |  |  |  |
| `botSpeechSettings` | [BotSpeechSettingsNoKind](settings.md#botspeechsettingsnokind) |  |  |  |
| `networkSettings` | [NetworkSettingsNoKind](settings.md#networksettingsnokind) |  |  |  |
| `generationSettings` | [GenerationSettingsNoKind](settings.md#generationsettingsnokind) |  |  |  |
| `publishOnCreate` | `boolean` |  | `false` |  |
| `publishOnImport` | `boolean` |  | `true` |  |
| `gPTSettings` | [GPTSettingsNoKind](settings.md#gptsettingsnokind) |  |  |  |
| `isLightweightBot` | `boolean` |  | `false` |  |
| `aISettings` | [AISettingsNoKind](settings.md#aisettingsnokind) |  |  |  |
| `recognizer` | [RecognizerBase](dialogs.md#recognizerbase) |  |  |  |
| `analyticsSettings` | [AnalyticsSettingsNoKind](settings.md#analyticssettingsnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind)</sub>

### BotDefinition

YAML: `kind: BotDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `components` | list of [BotComponentBase](#botcomponentbase) |  |  |  |
| `environmentVariables` | list of [EnvironmentVariableDefinitionNoKind](#environmentvariabledefinitionnokind) |  |  |  |
| `flows` | list of [CloudFlowDefinitionNoKind](#cloudflowdefinitionnokind) |  |  |  |
| `dataverseTableSearchs` | list of [DataverseTableSearchBase](knowledge.md#dataversetablesearchbase) |  |  |  |
| `dataverseTableSearchGlossaryConfigurations` | list of [DataverseTableSearchGlossaryConfigurationNoKind](#dataversetablesearchglossaryconfigurationnokind) |  |  |  |
| `dataverseTableSearchEntityConfigurations` | list of [DataverseTableSearchEntityConfigurationNoKind](#dataversetablesearchentityconfigurationnokind) |  |  |  |
| `dataverseTableSearchEntityColumnSynonyms` | list of [DataverseTableSearchEntityColumnSynonymNoKind](#dataversetablesearchentitycolumnsynonymnokind) |  |  |  |
| `connectionReferences` | list of [ConnectionReferenceNoKind](#connectionreferencenokind) |  |  |  |
| `connectorDefinitions` | list of [ConnectorDefinitionNoKind](#connectordefinitionnokind) |  |  |  |
| `aIModelDefinitions` | list of [AIModelDefinitionNoKind](#aimodeldefinitionnokind) |  |  |  |
| `aIPluginOperations` | list of [AIPluginOperationNoKind](#aipluginoperationnokind) |  |  |  |
| `connectedAgentDefinitions` | list of [ConnectedAgentDefinitionNoKind](#connectedagentdefinitionnokind) |  |  |  |
| `publishedVersion` | `string` |  |  |  |
| `forceUseVersion` | `boolean` |  |  |  |
| `entity` | [BotEntityNoKind](#botentitynokind) | ✓ |  |  |
| `componentCollections` | list of [BotComponentCollectionNoKind](#botcomponentcollectionnokind) |  |  |  |
| `connectedBots` | list of [BotEntityNoKind](#botentitynokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DefinitionBase](other.md#definitionbase)</sub>

### BotEntity

YAML: `kind: BotEntity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `componentIdUnique` | `string` |  |  |  |
| `cdsBotId` | `string` | ✓ |  |  |
| `accessControlPolicy` | [BotAccessControlPolicy](#botaccesscontrolpolicy) |  | `"Any"` |  |
| `authenticationMode` | [BotAuthenticationMode](#botauthenticationmode) |  | `"Unspecified"` |  |
| `authenticationTrigger` | [BotAuthenticationTrigger](#botauthenticationtrigger) |  | `"AsNeeded"` |  |
| `authorizedSecurityGroupIds` | list of `string` |  |  |  |
| `teamsAppInformation` | [ApplicationDetailsNoKind](#applicationdetailsnokind) |  |  |  |
| `configuration` | [BotConfigurationNoKind](#botconfigurationnokind) |  |  |  |
| `authenticationConfiguration` | [BotAuthenticationConfigurationNoKind](#botauthenticationconfigurationnokind) |  |  |  |
| `publishedOn` | `string` |  |  |  |
| `synchronizationStatus` | [BotSynchronizationDetailsNoKind](#botsynchronizationdetailsnokind) |  |  |  |
| `template` | `string` |  |  |  |
| `iconBase64` | `string` |  |  |  |
| `language` | [Language](settings.md#language) |  |  |  |
| `runtimeProvider` | [RuntimeProvider](#runtimeprovider) |  | `"PowerVirtualAgents"` |  |
| `providerConnectionReferenceId` | `string` |  |  |  |
| `supportedLanguages` | list of `Language` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [ProvisioningStatus](#provisioningstatus) |  |  |  |
| `origin` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### BotEntityNoKind

*Title:* BotEntity  
Same shape as [BotEntity](#botentity) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `componentIdUnique` | `string` |  |  |  |
| `cdsBotId` | `string` | ✓ |  |  |
| `accessControlPolicy` | [BotAccessControlPolicy](#botaccesscontrolpolicy) |  | `"Any"` |  |
| `authenticationMode` | [BotAuthenticationMode](#botauthenticationmode) |  | `"Unspecified"` |  |
| `authenticationTrigger` | [BotAuthenticationTrigger](#botauthenticationtrigger) |  | `"AsNeeded"` |  |
| `authorizedSecurityGroupIds` | list of `string` |  |  |  |
| `teamsAppInformation` | [ApplicationDetailsNoKind](#applicationdetailsnokind) |  |  |  |
| `configuration` | [BotConfigurationNoKind](#botconfigurationnokind) |  |  |  |
| `authenticationConfiguration` | [BotAuthenticationConfigurationNoKind](#botauthenticationconfigurationnokind) |  |  |  |
| `publishedOn` | `string` |  |  |  |
| `synchronizationStatus` | [BotSynchronizationDetailsNoKind](#botsynchronizationdetailsnokind) |  |  |  |
| `template` | `string` |  |  |  |
| `iconBase64` | `string` |  |  |  |
| `language` | [Language](settings.md#language) |  |  |  |
| `runtimeProvider` | [RuntimeProvider](#runtimeprovider) |  | `"PowerVirtualAgents"` |  |
| `providerConnectionReferenceId` | `string` |  |  |  |
| `supportedLanguages` | list of `Language` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [ProvisioningStatus](#provisioningstatus) |  |  |  |
| `origin` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotDefinition](#botdefinition)</sub>

### BotRegistrationDetails

YAML: `kind: BotRegistrationDetails`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `botRegistrationIdConsumptionTime` | `string` |  |  |  |
| `applicationId` | `string` |  |  |  |
| `isAppAvailableInTenant` | `boolean` |  |  |  |

_No other properties are allowed._

### BotRegistrationDetailsNoKind

*Title:* BotRegistrationDetails  
Same shape as [BotRegistrationDetails](#botregistrationdetails) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `botRegistrationIdConsumptionTime` | `string` |  |  |  |
| `applicationId` | `string` |  |  |  |
| `isAppAvailableInTenant` | `boolean` |  |  |  |

_No other properties are allowed._

<sub>Used by: [SynchronizationState](#synchronizationstate), [SynchronizationStateNoKind](#synchronizationstatenokind)</sub>

### BotSettingsComponent

YAML: `kind: BotSettingsComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `settings` | [BotSettingsBase](settings.md#botsettingsbase) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### BotSynchronizationDetails

YAML: `kind: BotSynchronizationDetails`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `contentVersion` | `number` |  |  |  |
| `lastFinishedPublishOperation` | [PublishResultDetailsNoKind](#publishresultdetailsnokind) |  |  |  |
| `lastPublishedDetails` | [SuccessfulPublishResultDetailsNoKind](#successfulpublishresultdetailsnokind) |  |  |  |
| `lastPublishedUserId` | `string` |  |  |  |
| `lastPublishedOnUtc` | `string` |  |  |  |
| `currentSynchronizationState` | [SynchronizationStateNoKind](#synchronizationstatenokind) |  |  |  |
| `lastSynchronizedOnUtc` | `string` |  |  |  |
| `domainLanguageModelURN` | `string` |  |  |  |
| `naturalLanguageModelURN` | `string` |  |  |  |
| `modelProfile` | [AdaptationManagementServiceModelProfileNoKind](settings.md#adaptationmanagementservicemodelprofilenokind) |  |  |  |
| `graphDriveItems` | list of [GraphDriveItemSyncDetailsNoKind](knowledge.md#graphdriveitemsyncdetailsnokind) |  |  |  |

_No other properties are allowed._

### BotSynchronizationDetailsNoKind

*Title:* BotSynchronizationDetails  
Same shape as [BotSynchronizationDetails](#botsynchronizationdetails) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `contentVersion` | `number` |  |  |  |
| `lastFinishedPublishOperation` | [PublishResultDetailsNoKind](#publishresultdetailsnokind) |  |  |  |
| `lastPublishedDetails` | [SuccessfulPublishResultDetailsNoKind](#successfulpublishresultdetailsnokind) |  |  |  |
| `lastPublishedUserId` | `string` |  |  |  |
| `lastPublishedOnUtc` | `string` |  |  |  |
| `currentSynchronizationState` | [SynchronizationStateNoKind](#synchronizationstatenokind) |  |  |  |
| `lastSynchronizedOnUtc` | `string` |  |  |  |
| `domainLanguageModelURN` | `string` |  |  |  |
| `naturalLanguageModelURN` | `string` |  |  |  |
| `modelProfile` | [AdaptationManagementServiceModelProfileNoKind](settings.md#adaptationmanagementservicemodelprofilenokind) |  |  |  |
| `graphDriveItems` | list of [GraphDriveItemSyncDetailsNoKind](knowledge.md#graphdriveitemsyncdetailsnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind)</sub>

### Category

YAML: `kind: Category`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `displayName` | `string` |  |  |  |

_No other properties are allowed._

### CategoryNoKind

*Title:* Category  
Same shape as [Category](#category) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `displayName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotConfiguration](#botconfiguration), [BotConfigurationNoKind](#botconfigurationnokind)</sub>

### ChannelDefinition

YAML: `kind: ChannelDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `channelId` | [ChannelId](#channelid) | ✓ |  |  |
| `channelSpecifier` | `string` |  |  |  |
| `displayName` | `string` |  |  |  |

_No other properties are allowed._

### ChannelDefinitionNoKind

*Title:* ChannelDefinition  
Same shape as [ChannelDefinition](#channeldefinition) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `channelId` | [ChannelId](#channelid) | ✓ |  |  |
| `channelSpecifier` | `string` |  |  |  |
| `displayName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotConfiguration](#botconfiguration), [BotConfigurationNoKind](#botconfigurationnokind)</sub>

### ChannelId


**String enum.** Allowed values: `Directline`, `Facebook`, `MsTeams`, `Telephony`, `Omnichannel`, `Microsoft365Copilot`, `Agents`, `MsTeamsPhone`

<sub>Used by: [ChannelDefinition](#channeldefinition), [ChannelDefinitionNoKind](#channeldefinitionnokind)</sub>

### CloudFlowConnectionType


**String enum.** Allowed values: `EmbeddedOnly`, `InvokerOnly`, `EmbeddedAndInvoker`, `NoConnections`

<sub>Used by: [CloudFlowDefinition](#cloudflowdefinition), [CloudFlowDefinitionNoKind](#cloudflowdefinitionnokind)</sub>

### CloudFlowDefinition

YAML: `kind: CloudFlowDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `isEnabled` | `boolean` |  | `true` |  |
| `workflowId` | `string` | ✓ |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `triggerType` | [CloudFlowTriggerType](#cloudflowtriggertype) |  | `"Standard"` |  |
| `connectionType` | [CloudFlowConnectionType](#cloudflowconnectiontype) |  |  |  |
| `secureInputs` | `boolean` |  | `false` |  |
| `secureOutputs` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### CloudFlowDefinitionNoKind

*Title:* CloudFlowDefinition  
Same shape as [CloudFlowDefinition](#cloudflowdefinition) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `isEnabled` | `boolean` |  | `true` |  |
| `workflowId` | `string` | ✓ |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `triggerType` | [CloudFlowTriggerType](#cloudflowtriggertype) |  | `"Standard"` |  |
| `connectionType` | [CloudFlowConnectionType](#cloudflowconnectiontype) |  |  |  |
| `secureInputs` | `boolean` |  | `false` |  |
| `secureOutputs` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotDefinition](#botdefinition)</sub>

### CloudFlowTriggerType


**String enum.** Allowed values: `Standard`, `Copilot`

<sub>Used by: [CloudFlowDefinition](#cloudflowdefinition), [CloudFlowDefinitionNoKind](#cloudflowdefinitionnokind)</sub>

### ConnectedAgentDefinition

YAML: `kind: ConnectedAgentDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `publishedVersion` | `number` |  |  |  |
| `connectedAgentSchemaName` | `string` | ✓ |  |  |
| `isAgentConnectable` | `boolean` |  | `true` |  |

_No other properties are allowed._

<sub>Used by: [AgentDefinition](#agentdefinition)</sub>

### ConnectedAgentDefinitionNoKind

Same shape as [ConnectedAgentDefinition](#connectedagentdefinition) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `publishedVersion` | `number` |  |  |  |
| `connectedAgentSchemaName` | `string` | ✓ |  |  |
| `isAgentConnectable` | `boolean` |  | `true` |  |

_No other properties are allowed._

<sub>Used by: [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotDefinition](#botdefinition)</sub>

### ConnectionMode


**String enum.** Allowed values: `Maker`, `Invoker`

<sub>Used by: [ConnectionProperties](settings.md#connectionproperties), [ConnectionPropertiesNoKind](settings.md#connectionpropertiesnokind), [RecurrenceTriggerDefinition](#recurrencetriggerdefinition)</sub>

### ConnectionParametersSet

YAML: `kind: ConnectionParametersSet`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionParametersSetName` | `string` |  |  |  |
| `parameters` | map of name → [ValueExpression](other.md#valueexpression) |  |  |  |

_No other properties are allowed._

### ConnectionParametersSetNoKind

*Title:* ConnectionParametersSet  
Same shape as [ConnectionParametersSet](#connectionparametersset) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionParametersSetName` | `string` |  |  |  |
| `parameters` | map of name → [ValueExpression](other.md#valueexpression) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ConnectionParameterSettings](settings.md#connectionparametersettings), [ConnectionReferenceNameInFlowToConnectionParametersSetMapping](#connectionreferencenameinflowtoconnectionparameterssetmapping), [ConnectionReferenceNameInFlowToConnectionParametersSetMappingNoKind](settings.md#connectionreferencenameinflowtoconnectionparameterssetmappingnokind)</sub>

### ConnectionReference

YAML: `kind: ConnectionReference`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `connectionId` | `string` |  |  |  |
| `customConnectorId` | `string` |  |  |  |
| `connectorId` | `string` |  |  |  |
| `connectionReferenceLogicalName` | `string` | ✓ |  |  |
| `displayName` | `string` |  |  |  |
| `sharedConnectionParameters` | `string` |  |  |  |
| `allowSharing` | `boolean` |  |  |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### ConnectionReferenceNameInFlowToConnectionParametersSetMapping

YAML: `kind: ConnectionReferenceNameInFlowToConnectionParametersSetMapping`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionParameters` | map of name → [ConnectionParametersSetNoKind](#connectionparameterssetnokind) |  |  |  |

_No other properties are allowed._

### ConnectionReferenceNoKind

*Title:* ConnectionReference  
Same shape as [ConnectionReference](#connectionreference) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `connectionId` | `string` |  |  |  |
| `customConnectorId` | `string` |  |  |  |
| `connectorId` | `string` |  |  |  |
| `connectionReferenceLogicalName` | `string` | ✓ |  |  |
| `displayName` | `string` |  |  |  |
| `sharedConnectionParameters` | `string` |  |  |  |
| `allowSharing` | `boolean` |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotDefinition](#botdefinition), [ConnectionReferencesSourceFile](#connectionreferencessourcefile)</sub>

### ConnectionReferencesSourceFile

YAML: `kind: ConnectionReferencesSourceFile`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionReferences` | list of [ConnectionReferenceNoKind](#connectionreferencenokind) |  |  |  |

_No other properties are allowed._

### ConnectorDefinition

YAML: `kind: ConnectorDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectorId` | `string` | ✓ |  |  |
| `snapshotId` | `string` |  |  |  |
| `version` | `number` |  |  |  |
| `displayName` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `isCustom` | `boolean` |  |  |  |
| `connectorType` | [ConnectorDefinitionType](#connectordefinitiontype) |  |  |  |
| `customConnectorId` | `string` |  |  |  |
| `iconUri` | `string` |  |  |  |
| `changedTime` | `string` |  |  |  |
| `hasPublicData` | `boolean` |  | `false` |  |
| `isSSOSupported` | `boolean` |  | `false` |  |
| `operations` | list of [ConnectorOperationNoKind](#connectoroperationnokind) |  |  |  |
| `triggers` | list of [ConnectorTriggerOperationNoKind](#connectortriggeroperationnokind) |  |  |  |

_No other properties are allowed._

### ConnectorDefinitionComponent

YAML: `kind: ConnectorDefinitionComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `connectorId` | `string` | ✓ |  |  |
| `connectorGuidId` | `string` | ✓ |  |  |
| `iconBrandColor` | `string` |  |  |  |
| `displayName` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `iconUri` | `string` |  |  |  |
| `openApiDefinition` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### ConnectorDefinitionNoKind

*Title:* ConnectorDefinition  
Same shape as [ConnectorDefinition](#connectordefinition) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectorId` | `string` | ✓ |  |  |
| `snapshotId` | `string` |  |  |  |
| `version` | `number` |  |  |  |
| `displayName` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `isCustom` | `boolean` |  |  |  |
| `connectorType` | [ConnectorDefinitionType](#connectordefinitiontype) |  |  |  |
| `customConnectorId` | `string` |  |  |  |
| `iconUri` | `string` |  |  |  |
| `changedTime` | `string` |  |  |  |
| `hasPublicData` | `boolean` |  | `false` |  |
| `isSSOSupported` | `boolean` |  | `false` |  |
| `operations` | list of [ConnectorOperationNoKind](#connectoroperationnokind) |  |  |  |
| `triggers` | list of [ConnectorTriggerOperationNoKind](#connectortriggeroperationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotDefinition](#botdefinition)</sub>

### ConnectorDefinitionType


**String enum.** Allowed values: `Solution`, `Environment`, `Public`

<sub>Used by: [ConnectorDefinition](#connectordefinition), [ConnectorDefinitionNoKind](#connectordefinitionnokind)</sub>

### ConnectorDynamicValuesCapability


**String enum.** Allowed values: `List`, `FilePicker`

<sub>Used by: [ConnectorDynamicValuesConfig](#connectordynamicvaluesconfig), [ConnectorDynamicValuesConfigNoKind](data-types.md#connectordynamicvaluesconfignokind)</sub>

### ConnectorDynamicValuesConfig

YAML: `kind: ConnectorDynamicValuesConfig`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `capability` | [ConnectorDynamicValuesCapability](#connectordynamicvaluescapability) |  |  |  |

_No other properties are allowed._

### ConnectorOperation

YAML: `kind: ConnectorOperation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `displayName` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `nameForModel` | `string` |  |  |  |
| `descriptionForModel` | `string` |  |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) | ✓ |  |  |
| `outputType` | [DataType](data-types.md#datatype) | ✓ |  |  |
| `operationId` | `string` | ✓ |  |  |

_No other properties are allowed._

### ConnectorOperationNoKind

*Title:* ConnectorOperation  
Same shape as [ConnectorOperation](#connectoroperation) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `displayName` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `nameForModel` | `string` |  |  |  |
| `descriptionForModel` | `string` |  |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) | ✓ |  |  |
| `outputType` | [DataType](data-types.md#datatype) | ✓ |  |  |
| `operationId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [ConnectorDefinition](#connectordefinition), [ConnectorDefinitionNoKind](#connectordefinitionnokind)</sub>

### ConnectorTriggerDefinition

YAML: `kind: ConnectorTriggerDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `callbackMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) | ✓ |  |  |
| `inputs` | list of [TaskInput](task-actions.md#taskinput) |  |  |  |
| `connectionReference` | `string` | ✓ |  |  |
| `connectionProperties` | [ConnectionPropertiesNoKind](settings.md#connectionpropertiesnokind) |  |  |  |
| `triggerId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [ExternalTriggerDefinition](#externaltriggerdefinition)</sub>

### ConnectorTriggerOperation

YAML: `kind: ConnectorTriggerOperation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `displayName` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `nameForModel` | `string` |  |  |  |
| `descriptionForModel` | `string` |  |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) | ✓ |  |  |
| `outputType` | [DataType](data-types.md#datatype) | ✓ |  |  |
| `triggerId` | `string` | ✓ |  |  |

_No other properties are allowed._

### ConnectorTriggerOperationNoKind

*Title:* ConnectorTriggerOperation  
Same shape as [ConnectorTriggerOperation](#connectortriggeroperation) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `displayName` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `nameForModel` | `string` |  |  |  |
| `descriptionForModel` | `string` |  |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) | ✓ |  |  |
| `outputType` | [DataType](data-types.md#datatype) | ✓ |  |  |
| `triggerId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [ConnectorDefinition](#connectordefinition), [ConnectorDefinitionNoKind](#connectordefinitionnokind)</sub>

### ContentShareContextNoKind

*Title:* ContentShareContext  
Same shape as [ContentShareContext](messages.md#contentsharecontext) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `reusePolicy` | [BotComponentReusePolicy](settings.md#botcomponentreusepolicy) |  | `"None"` |  |
| `iconUri` | `string` |  |  |  |
| `helpUri` | `string` |  |  |  |
| `accentColor` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotSettingsComponent](#botsettingscomponent), [CustomEntityComponent](#customentitycomponent), [CustomMetricDefinitionComponent](#custommetricdefinitioncomponent), [DialogComponent](#dialogcomponent), [DialogComponentNoKind](#dialogcomponentnokind), [ExternalTriggerComponent](#externaltriggercomponent), [FileAttachmentComponent](#fileattachmentcomponent), [GlobalVariableComponent](#globalvariablecomponent), [GptComponent](#gptcomponent), [KnowledgeSourceComponent](#knowledgesourcecomponent), [LegacyOrUnknownComponent](#legacyorunknowncomponent), [SkillComponent](#skillcomponent), [TestCaseComponent](#testcasecomponent), [TranslationsComponent](#translationscomponent), [TranslationsComponentNoKind](#translationscomponentnokind)</sub>

### ConversationStarter

YAML: `kind: ConversationStarter`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `title` | `string` | ✓ |  |  |
| `text` | `string` | ✓ |  |  |

_No other properties are allowed._

### ConversationStarterNoKind

*Title:* ConversationStarter  
Same shape as [ConversationStarter](#conversationstarter) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `title` | `string` | ✓ |  |  |
| `text` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [GptComponentMetadata](#gptcomponentmetadata), [GptComponentMetadataNoKind](#gptcomponentmetadatanokind)</sub>

### ConversionMetadata

YAML: `kind: ConversionMetadata`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `variableIdToNameMapping` | map of name → `string` |  |  |  |
| `initializedVariableIdToTypeMapping` | map of name → [DataType](data-types.md#datatype) |  |  |  |

_No other properties are allowed._

### CopilotChatApplicationDetails

YAML: `kind: CopilotChatApplicationDetails`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `isEnabled` | `boolean` |  |  |  |

_No other properties are allowed._

### CopilotChatApplicationDetailsNoKind

*Title:* CopilotChatApplicationDetails  
Same shape as [CopilotChatApplicationDetails](#copilotchatapplicationdetails) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `isEnabled` | `boolean` |  |  |  |

_No other properties are allowed._

<sub>Used by: [ApplicationDetails](#applicationdetails), [ApplicationDetailsNoKind](#applicationdetailsnokind)</sub>

### CustomEntityComponent

YAML: `kind: CustomEntityComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `entity` | [Entity](entities.md#entity) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### CustomMetricCategory

YAML: `kind: CustomMetricCategory`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `label` | `string` |  |  |  |
| `description` | `string` |  |  |  |

_No other properties are allowed._

### CustomMetricCategoryNoKind

Same shape as [CustomMetricCategory](#custommetriccategory) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `label` | `string` |  |  |  |
| `description` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [CustomMetricDefinition](#custommetricdefinition), [CustomMetricDefinitionNoKind](#custommetricdefinitionnokind)</sub>

### CustomMetricDefinition

YAML: `kind: CustomMetricDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `evaluationPrompt` | `string` |  |  |  |
| `categories` | list of [CustomMetricCategoryNoKind](#custommetriccategorynokind) |  |  |  |
| `fallbackCategory` | `string` |  |  |  |

_No other properties are allowed._

### CustomMetricDefinitionComponent

YAML: `kind: CustomMetricDefinitionComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `definition` | [CustomMetricDefinitionNoKind](#custommetricdefinitionnokind) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### CustomMetricDefinitionNoKind

Same shape as [CustomMetricDefinition](#custommetricdefinition) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `evaluationPrompt` | `string` |  |  |  |
| `categories` | list of [CustomMetricCategoryNoKind](#custommetriccategorynokind) |  |  |  |
| `fallbackCategory` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [CustomMetricDefinitionComponent](#custommetricdefinitioncomponent)</sub>

### DataverseTableSearch

YAML: `kind: DataverseTableSearch`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `searchType` | [TableSearchType](#tablesearchtype) |  | `"DataverseSearch"` |  |
| `name` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DataverseTableSearchBase](knowledge.md#dataversetablesearchbase), [SolutionComponent](#solutioncomponent)</sub>

### DataverseTableSearchEntityColumnSynonym

YAML: `kind: DataverseTableSearchEntityColumnSynonym`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `dataverseTableSearchEntityConfiguration` | `string` |  |  |  |
| `synonyms` | `string` | ✓ |  |  |
| `columnLogicalName` | `string` | ✓ |  |  |
| `columnDescription` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### DataverseTableSearchEntityColumnSynonymNoKind

*Title:* DataverseTableSearchEntityColumnSynonym  
Same shape as [DataverseTableSearchEntityColumnSynonym](#dataversetablesearchentitycolumnsynonym) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `dataverseTableSearchEntityConfiguration` | `string` |  |  |  |
| `synonyms` | `string` | ✓ |  |  |
| `columnLogicalName` | `string` | ✓ |  |  |
| `columnDescription` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotDefinition](#botdefinition)</sub>

### DataverseTableSearchEntityConfiguration

YAML: `kind: DataverseTableSearchEntityConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `dataverseTableSearch` | `string` |  |  |  |
| `name` | `string` | ✓ |  |  |
| `entityLogicalName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### DataverseTableSearchEntityConfigurationNoKind

*Title:* DataverseTableSearchEntityConfiguration  
Same shape as [DataverseTableSearchEntityConfiguration](#dataversetablesearchentityconfiguration) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `dataverseTableSearch` | `string` |  |  |  |
| `name` | `string` | ✓ |  |  |
| `entityLogicalName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotDefinition](#botdefinition)</sub>

### DataverseTableSearchGlossaryConfiguration

YAML: `kind: DataverseTableSearchGlossaryConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `dataverseTableSearch` | `string` |  |  |  |
| `term` | `string` | ✓ |  |  |
| `termDescription` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### DataverseTableSearchGlossaryConfigurationNoKind

*Title:* DataverseTableSearchGlossaryConfiguration  
Same shape as [DataverseTableSearchGlossaryConfiguration](#dataversetablesearchglossaryconfiguration) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `dataverseTableSearch` | `string` |  |  |  |
| `term` | `string` | ✓ |  |  |
| `termDescription` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotDefinition](#botdefinition)</sub>

### DefaultResponseMode


**String enum.** Allowed values: `Auto`, `ThinkDeeper`, `QuickResponse`

<sub>Used by: [GptComponentMetadata](#gptcomponentmetadata), [GptComponentMetadataNoKind](#gptcomponentmetadatanokind)</sub>

### DiagnosticType


**String enum.** Allowed values: `BotElementError`, `BotElementWarning`, `All`

### DialogComponent

YAML: `kind: DialogComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `dialog` | [DialogBase](dialogs.md#dialogbase) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### DialogComponentNoKind

*Title:* DialogComponent  
Same shape as [DialogComponent](#dialogcomponent) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `dialog` | [DialogBase](dialogs.md#dialogbase) | ✓ |  |  |

_No other properties are allowed._

### DialogComponentReference

YAML: `kind: DialogComponentReference`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `schemaName` | `string` | ✓ |  |  |

_No other properties are allowed._

### DlpComponentType


**String enum.** Allowed values: `Flow`, `Skill`, `HttpRequest`, `AnonymousAccess`, `Channel`, `Connector`, `GenerativeAnswers`, `GenerativeActions`, `GenerativeBuilder`, `ApplicationInsights`, `Knowledge`

### EnvironmentVariableDefinition

YAML: `kind: EnvironmentVariableDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `id` | `string` |  |  |  |
| `defaultValue` | `string` |  |  |  |
| `valueComponent` | [EnvironmentVariableValueNoKind](#environmentvariablevaluenokind) |  |  |  |
| `type` | [EnvironmentVariableType](#environmentvariabletype) |  | `"String"` |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### EnvironmentVariableDefinitionNoKind

*Title:* EnvironmentVariableDefinition  
Same shape as [EnvironmentVariableDefinition](#environmentvariabledefinition) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `id` | `string` |  |  |  |
| `defaultValue` | `string` |  |  |  |
| `valueComponent` | [EnvironmentVariableValueNoKind](#environmentvariablevaluenokind) |  |  |  |
| `type` | [EnvironmentVariableType](#environmentvariabletype) |  | `"String"` |  |

_No other properties are allowed._

<sub>Used by: [BotComponentCollectionDefinition](#botcomponentcollectiondefinition), [BotDefinition](#botdefinition)</sub>

### EnvironmentVariableReference

YAML: `kind: EnvironmentVariableReference`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `schemaName` | `string` | ✓ |  |  |

_No other properties are allowed._

### EnvironmentVariableReferenceNoKind

*Title:* EnvironmentVariableReference  
Same shape as [EnvironmentVariableReference](#environmentvariablereference) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `schemaName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [FabricAISkillSource](knowledge.md#fabricaiskillsource), [GraphConnectorSearchSource](knowledge.md#graphconnectorsearchsource), [SkillDefinition](#skilldefinition), [SkillDefinitionNoKind](#skilldefinitionnokind)</sub>

### EnvironmentVariableSecretReadStatus


**String enum.** Allowed values: `Success`, `BotForbidden`, `ServiceForbidden`, `NotFound`, `InternalError`

### EnvironmentVariableType


**String enum.** Allowed values: `String`, `Number`, `Boolean`, `Json`, `DataSource`, `Secret`

<sub>Used by: [EnvironmentVariableDefinition](#environmentvariabledefinition), [EnvironmentVariableDefinitionNoKind](#environmentvariabledefinitionnokind)</sub>

### EnvironmentVariableValue

YAML: `kind: EnvironmentVariableValue`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` |  |  |  |
| `definitionId` | `string` |  |  |  |
| `value` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [SolutionComponent](#solutioncomponent)</sub>

### EnvironmentVariableValueNoKind

*Title:* EnvironmentVariableValue  
Same shape as [EnvironmentVariableValue](#environmentvariablevalue) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` |  |  |  |
| `definitionId` | `string` |  |  |  |
| `value` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [EnvironmentVariableDefinition](#environmentvariabledefinition), [EnvironmentVariableDefinitionNoKind](#environmentvariabledefinitionnokind)</sub>

### ExternalTriggerComponent

YAML: `kind: ExternalTriggerComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `configuration` | [ExternalTriggerConfigurationNoKind](#externaltriggerconfigurationnokind) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### ExternalTriggerConfiguration

YAML: `kind: ExternalTriggerConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `externalTriggerSource` | [ExternalTriggerDefinition](#externaltriggerdefinition) |  |  |  |

_No other properties are allowed._

### ExternalTriggerConfigurationNoKind

*Title:* ExternalTriggerConfiguration  
Same shape as [ExternalTriggerConfiguration](#externaltriggerconfiguration) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `externalTriggerSource` | [ExternalTriggerDefinition](#externaltriggerdefinition) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ExternalTriggerComponent](#externaltriggercomponent)</sub>

### ExternalTriggerDefinition


**One of:** [WorkflowExternalTrigger](#workflowexternaltrigger), [RecurrenceTriggerDefinition](#recurrencetriggerdefinition), [ConnectorTriggerDefinition](#connectortriggerdefinition)

<sub>Used by: [ExternalTriggerConfiguration](#externaltriggerconfiguration), [ExternalTriggerConfigurationNoKind](#externaltriggerconfigurationnokind)</sub>

### FederatedTableSearch

YAML: `kind: FederatedTableSearch`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `searchType` | [TableSearchType](#tablesearchtype) |  | `"DataverseSearch"` |  |
| `name` | `string` | ✓ |  |  |
| `connectionReference` | `string` |  |  |  |
| `connectionReferenceLogicalName` | `string` |  |  |  |
| `connectorId` | `string` |  |  |  |
| `connectionId` | `string` |  |  |  |
| `knowledgeConfig` | [BaseSearchConfiguration](knowledge.md#basesearchconfiguration) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DataverseTableSearchBase](knowledge.md#dataversetablesearchbase), [SolutionComponent](#solutioncomponent)</sub>

### FileAttachmentComponent

YAML: `kind: FileAttachmentComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `metadata` | [FileAttachmentComponentMetadataNoKind](#fileattachmentcomponentmetadatanokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### FileAttachmentComponentMetadataNoKind

Same shape as [FileAttachmentComponentMetadata](messages.md#fileattachmentcomponentmetadata) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [FileAttachmentComponent](#fileattachmentcomponent)</sub>

### GlobalVariableComponent

YAML: `kind: GlobalVariableComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `variable` | [VariableNoKind](#variablenokind) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### GptCapabilitiesNoKind

Same shape as [GptCapabilities](settings.md#gptcapabilities) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `webBrowsing` | `boolean` |  | `false` |  |
| `codeInterpreter` | `boolean` |  | `false` |  |
| `generateImages` | `boolean` |  | `false` |  |
| `fineTunedModels` | [FineTunedModelsPropertyDetailsNoKind](settings.md#finetunedmodelspropertydetailsnokind) |  |  |  |
| `searchTeams` | `boolean` |  | `false` |  |
| `searchOneDriveAndSharePoint` | `boolean` |  | `false` |  |
| `searchEmails` | `boolean` |  | `false` |  |
| `searchMeetings` | `boolean` |  | `false` |  |
| `searchPeople` | `boolean` |  | `false` |  |
| `searchPeopleIncludeRelatedContent` | `boolean` |  | `false` |  |
| `customExperience` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [GptComponentMetadata](#gptcomponentmetadata), [GptComponentMetadataNoKind](#gptcomponentmetadatanokind)</sub>

### GptComponent

YAML: `kind: GptComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `metadata` | [GptComponentMetadataNoKind](#gptcomponentmetadatanokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### GptComponentMetadata

YAML: `kind: GptComponentMetadata`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `historyType` | [HistoryType](settings.md#historytype) |  |  |  |
| `instructions` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `responseInstructions` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `tools` | list of [TaskAction](task-actions.md#taskaction) |  |  |  |
| `knowledgeSources` | [SearchKnowledgeSourcesMode](knowledge.md#searchknowledgesourcesmode) |  |  |  |
| `fileSearchDataSource` | [FileSearchDataSourceNoKind](knowledge.md#filesearchdatasourcenokind) |  |  |  |
| `gptCapabilities` | [GptCapabilitiesNoKind](#gptcapabilitiesnokind) |  |  |  |
| `conversationStarters` | list of [ConversationStarterNoKind](#conversationstarternokind) |  |  |  |
| `defaultResponseMode` | [DefaultResponseMode](#defaultresponsemode) |  |  |  |
| `aISettings` | [AISettingsNoKind](settings.md#aisettingsnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentDefinition](#agentdefinition)</sub>

### GptComponentMetadataNoKind

Same shape as [GptComponentMetadata](#gptcomponentmetadata) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `historyType` | [HistoryType](settings.md#historytype) |  |  |  |
| `instructions` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `responseInstructions` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `tools` | list of [TaskAction](task-actions.md#taskaction) |  |  |  |
| `knowledgeSources` | [SearchKnowledgeSourcesMode](knowledge.md#searchknowledgesourcesmode) |  |  |  |
| `fileSearchDataSource` | [FileSearchDataSourceNoKind](knowledge.md#filesearchdatasourcenokind) |  |  |  |
| `gptCapabilities` | [GptCapabilitiesNoKind](#gptcapabilitiesnokind) |  |  |  |
| `conversationStarters` | list of [ConversationStarterNoKind](#conversationstarternokind) |  |  |  |
| `defaultResponseMode` | [DefaultResponseMode](#defaultresponsemode) |  |  |  |
| `aISettings` | [AISettingsNoKind](settings.md#aisettingsnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentDefinition](#agentdefinition), [AgentDialog](dialogs.md#agentdialog), [GptComponent](#gptcomponent)</sub>

### KnowledgeSourceComponent

YAML: `kind: KnowledgeSourceComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `configuration` | [KnowledgeSourceConfigurationNoKind](#knowledgesourceconfigurationnokind) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### KnowledgeSourceConfiguration

YAML: `kind: KnowledgeSourceConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `source` | [KnowledgeSource](knowledge.md#knowledgesource) |  |  |  |

_No other properties are allowed._

### KnowledgeSourceConfigurationNoKind

Same shape as [KnowledgeSourceConfiguration](#knowledgesourceconfiguration) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `source` | [KnowledgeSource](knowledge.md#knowledgesource) |  |  |  |

_No other properties are allowed._

<sub>Used by: [KnowledgeSourceComponent](#knowledgesourcecomponent)</sub>

### LegacyOrUnknownComponent

YAML: `kind: LegacyOrUnknownComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `componentTypeInt` | `number` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `content` | `string` |  |  |  |
| `data` | `string` |  |  |  |
| `language` | [Language](settings.md#language) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### ManagedProperties

YAML: `kind: ManagedProperties`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `isManaged` | `boolean` |  |  |  |
| `isCustomizable` | `boolean` |  | `false` |  |
| `solutionId` | `string` |  |  |  |

_No other properties are allowed._

### ManagedPropertiesNoKind

*Title:* ManagedProperties  
Same shape as [ManagedProperties](#managedproperties) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `isManaged` | `boolean` |  |  |  |
| `isCustomizable` | `boolean` |  | `false` |  |
| `solutionId` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [AIModelDefinition](#aimodeldefinition), [AIModelDefinitionNoKind](#aimodeldefinitionnokind), [AIPluginComponent](#aiplugincomponent), [AIPluginOperationComponent](#aipluginoperationcomponent), [BotComponentCollection](#botcomponentcollection), [BotComponentCollectionNoKind](#botcomponentcollectionnokind), [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind), [BotSettingsComponent](#botsettingscomponent), [CloudFlowDefinition](#cloudflowdefinition), [CloudFlowDefinitionNoKind](#cloudflowdefinitionnokind), [ConnectionReference](#connectionreference), [ConnectionReferenceNoKind](#connectionreferencenokind), [ConnectorDefinitionComponent](#connectordefinitioncomponent), [CustomEntityComponent](#customentitycomponent), [CustomMetricDefinitionComponent](#custommetricdefinitioncomponent), [DataverseTableSearch](#dataversetablesearch), [DataverseTableSearchEntityColumnSynonym](#dataversetablesearchentitycolumnsynonym), [DataverseTableSearchEntityColumnSynonymNoKind](#dataversetablesearchentitycolumnsynonymnokind), [DataverseTableSearchEntityConfiguration](#dataversetablesearchentityconfiguration) … and 20 more</sub>

### ManifestResources


**One of:** [ModelManifestResource](settings.md#modelmanifestresource), [ToolManifestResource](#toolmanifestresource)

### McpToolDefinition

YAML: `kind: McpToolDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `name` | `string` | ✓ |  |  |
| `description` | `string` |  |  |  |
| `inputSchema` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputSchema` | [DataType](data-types.md#datatype) |  |  |  |

_No other properties are allowed._

### McpToolDefinitionNoKind

*Title:* McpToolDefinition  
Same shape as [McpToolDefinition](#mcptooldefinition) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `name` | `string` | ✓ |  |  |
| `description` | `string` |  |  |  |
| `inputSchema` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputSchema` | [DataType](data-types.md#datatype) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ModelContextProtocolMetadata](#modelcontextprotocolmetadata)</sub>

### Microsoft365ApplicationDetails

YAML: `kind: Microsoft365ApplicationDetails`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `titleId` | `string` |  |  |  |
| `shareLink` | `string` |  |  |  |
| `appId` | `string` |  |  |  |
| `lobTitleId` | `string` |  |  |  |
| `lobAppId` | `string` |  |  |  |
| `lastSharingUpdateDateTimeUtc` | `string` |  |  |  |

_No other properties are allowed._

### Microsoft365ApplicationDetailsNoKind

*Title:* Microsoft365ApplicationDetails  
Same shape as [Microsoft365ApplicationDetails](#microsoft365applicationdetails) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `titleId` | `string` |  |  |  |
| `shareLink` | `string` |  |  |  |
| `appId` | `string` |  |  |  |
| `lobTitleId` | `string` |  |  |  |
| `lobAppId` | `string` |  |  |  |
| `lastSharingUpdateDateTimeUtc` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [ApplicationDetails](#applicationdetails), [ApplicationDetailsNoKind](#applicationdetailsnokind)</sub>

### ModelContextProtocolMetadata

YAML: `kind: ModelContextProtocolMetadata`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `operationId` | `string` | ✓ |  |  |
| `dynamicInputSchema` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `isElicitationEnabled` | `boolean` |  |  |  |
| `tools` | [ToolsSelectionMode](#toolsselectionmode) |  |  |  |
| `knownTools` | list of [McpToolDefinitionNoKind](#mcptooldefinitionnokind) |  |  |  |
| `toolsLastRefreshedAt` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [ProtocolMetadataBase](#protocolmetadatabase)</sub>

### PlanStepType


**String enum.** Allowed values: `Action`, `Topic`, `Agent`, `KnowledgeSource`, `BuiltIn`

### PropertyOverride

YAML: `kind: PropertyOverride`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `property` | `string` | ✓ |  |  |
| `sensitivityLevel` | [SensitivityLevel](settings.md#sensitivitylevel) |  | `"None"` |  |

_No other properties are allowed._

### ProtocolMetadataBase


**One of:** [ModelContextProtocolMetadata](#modelcontextprotocolmetadata), [ActivityProtocolMetadata](#activityprotocolmetadata), [AgentToAgentProtocolMetadata](#agenttoagentprotocolmetadata)

<sub>Used by: [InvokeExternalAgentTaskAction](task-actions.md#invokeexternalagenttaskaction)</sub>

### ProvisioningState


**String enum.** Allowed values: `Provisioning`, `Provisioned`, `ProvisionedWithoutRegistration`, `ContentReady`, `ProvisionFailed`

<sub>Used by: [SynchronizationState](#synchronizationstate), [SynchronizationStateNoKind](#synchronizationstatenokind)</sub>

### ProvisioningStatus


**One of:** `1`, `2`, `3`, `4`, `5`

<sub>Used by: [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind)</sub>

### PublishResult


**String enum.** Allowed values: `Succeeded`, `Failed`

<sub>Used by: [PublishResultDetails](#publishresultdetails), [PublishResultDetailsNoKind](#publishresultdetailsnokind)</sub>

### PublishResultDetails

YAML: `kind: PublishResultDetails`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `userId` | `string` |  |  |  |
| `operationStart` | `string` |  |  |  |
| `operationEnd` | `string` |  |  |  |
| `status` | [PublishResult](#publishresult) |  |  |  |

_No other properties are allowed._

### PublishResultDetailsNoKind

*Title:* PublishResultDetails  
Same shape as [PublishResultDetails](#publishresultdetails) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `userId` | `string` |  |  |  |
| `operationStart` | `string` |  |  |  |
| `operationEnd` | `string` |  |  |  |
| `status` | [PublishResult](#publishresult) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotSynchronizationDetails](#botsynchronizationdetails), [BotSynchronizationDetailsNoKind](#botsynchronizationdetailsnokind)</sub>

### RecurrenceFrequency


**String enum.** Allowed values: `Minute`, `Day`, `Hour`, `Month`, `Week`, `Year`

### RecurrenceTriggerDefinition

YAML: `kind: RecurrenceTriggerDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `callbackMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) | ✓ |  |  |
| `inputs` | list of [TaskInput](task-actions.md#taskinput) |  |  |  |
| `mode` | [ConnectionMode](#connectionmode) |  | `"Maker"` |  |

_No other properties are allowed._

<sub>Used by: [ExternalTriggerDefinition](#externaltriggerdefinition)</sub>

### ReferenceItemSourceFile

YAML: `kind: ReferenceItemSourceFile`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `schemaName` | `string` |  |  |  |
| `directory` | `string` |  |  |  |

_No other properties are allowed._

### ReferenceItemSourceFileNoKind

*Title:* ReferenceItemSourceFile  
Same shape as [ReferenceItemSourceFile](#referenceitemsourcefile) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `schemaName` | `string` |  |  |  |
| `directory` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [ReferencesSourceFile](#referencessourcefile)</sub>

### ReferencesSourceFile

YAML: `kind: ReferencesSourceFile`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `componentCollections` | list of [ReferenceItemSourceFileNoKind](#referenceitemsourcefilenokind) |  |  |  |

_No other properties are allowed._

### ReferenceType


**String enum.** Allowed values: `Skill`, `CloudFlow`, `AIModel`, `Dialog`, `Trigger`, `Action`, `EnvironmentVariable`, `Entity`, `ConnectionReference`, `Gpt`, `File`, `Knowledge`

### RuntimeProvider


**String enum.** Allowed values: `PowerVirtualAgents`, `NuanceMixShell`

<sub>Used by: [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind)</sub>

### SkillActionDefinition

YAML: `kind: SkillActionDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `activityType` | `string` |  |  |  |
| `actionId` | `string` | ✓ |  |  |
| `title` | `string` |  |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |

_No other properties are allowed._

### SkillActionDefinitionNoKind

Same shape as [SkillActionDefinition](#skillactiondefinition) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `activityType` | `string` |  |  |  |
| `actionId` | `string` | ✓ |  |  |
| `title` | `string` |  |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [SkillDefinition](#skilldefinition), [SkillDefinitionNoKind](#skilldefinitionnokind)</sub>

### SkillComponent

YAML: `kind: SkillComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `skill` | [SkillDefinitionNoKind](#skilldefinitionnokind) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### SkillDefinition

YAML: `kind: SkillDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `appId` | [EnvironmentVariableReferenceNoKind](#environmentvariablereferencenokind) | ✓ |  |  |
| `appEndpoint` | [EnvironmentVariableReferenceNoKind](#environmentvariablereferencenokind) | ✓ |  |  |
| `manifestVersion` | `string` |  |  |  |
| `privacyPolicy` | `string` |  |  |  |
| `actions` | list of [SkillActionDefinitionNoKind](#skillactiondefinitionnokind) |  |  |  |

_No other properties are allowed._

### SkillDefinitionNoKind

*Title:* SkillDefinition  
Same shape as [SkillDefinition](#skilldefinition) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `appId` | [EnvironmentVariableReferenceNoKind](#environmentvariablereferencenokind) | ✓ |  |  |
| `appEndpoint` | [EnvironmentVariableReferenceNoKind](#environmentvariablereferencenokind) | ✓ |  |  |
| `manifestVersion` | `string` |  |  |  |
| `privacyPolicy` | `string` |  |  |  |
| `actions` | list of [SkillActionDefinitionNoKind](#skillactiondefinitionnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [SkillComponent](#skillcomponent)</sub>

### SolutionComponent


**One of:** [BotEntity](#botentity), [BotComponentCollection](#botcomponentcollection), [CloudFlowDefinition](#cloudflowdefinition), [AIModelDefinition](#aimodeldefinition), [ConnectorDefinitionComponent](#connectordefinitioncomponent), [EnvironmentVariableDefinition](#environmentvariabledefinition), [EnvironmentVariableValue](#environmentvariablevalue), [ConnectionReference](#connectionreference), [DataverseTableSearchEntityColumnSynonym](#dataversetablesearchentitycolumnsynonym), [DataverseTableSearchGlossaryConfiguration](#dataversetablesearchglossaryconfiguration), [DataverseTableSearchEntityConfiguration](#dataversetablesearchentityconfiguration), [DataverseTableSearch](#dataversetablesearch), [FederatedTableSearch](#federatedtablesearch), [DialogComponent](#dialogcomponent), [TranslationsComponent](#translationscomponent), [CustomEntityComponent](#customentitycomponent), [GlobalVariableComponent](#globalvariablecomponent), [SkillComponent](#skillcomponent), [FileAttachmentComponent](#fileattachmentcomponent), [GptComponent](#gptcomponent), [KnowledgeSourceComponent](#knowledgesourcecomponent), [ExternalTriggerComponent](#externaltriggercomponent), [BotSettingsComponent](#botsettingscomponent), [TestCaseComponent](#testcasecomponent), [LegacyOrUnknownComponent](#legacyorunknowncomponent), [CustomMetricDefinitionComponent](#custommetricdefinitioncomponent), [AIPluginOperationComponent](#aipluginoperationcomponent), [AIPluginComponent](#aiplugincomponent)

### StateCode


**String enum.** Allowed values: `Active`, `Inactive`

<sub>Used by: [BotEntity](#botentity), [BotEntityNoKind](#botentitynokind), [BotSettingsComponent](#botsettingscomponent), [CustomEntityComponent](#customentitycomponent), [CustomMetricDefinitionComponent](#custommetricdefinitioncomponent), [DialogComponent](#dialogcomponent), [DialogComponentNoKind](#dialogcomponentnokind), [ExternalTriggerComponent](#externaltriggercomponent), [FileAttachmentComponent](#fileattachmentcomponent), [GlobalVariableComponent](#globalvariablecomponent), [GptComponent](#gptcomponent), [KnowledgeSourceComponent](#knowledgesourcecomponent), [LegacyOrUnknownComponent](#legacyorunknowncomponent), [SkillComponent](#skillcomponent), [TestCaseComponent](#testcasecomponent), [TranslationsComponent](#translationscomponent), [TranslationsComponentNoKind](#translationscomponentnokind)</sub>

### SuccessfulPublishResultDetails

YAML: `kind: SuccessfulPublishResultDetails`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `authenticationMode` | [BotAuthenticationMode](#botauthenticationmode) |  |  |  |

_No other properties are allowed._

### SuccessfulPublishResultDetailsNoKind

*Title:* SuccessfulPublishResultDetails  
Same shape as [SuccessfulPublishResultDetails](#successfulpublishresultdetails) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `authenticationMode` | [BotAuthenticationMode](#botauthenticationmode) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotSynchronizationDetails](#botsynchronizationdetails), [BotSynchronizationDetailsNoKind](#botsynchronizationdetailsnokind)</sub>

### SynchronizationState

YAML: `kind: SynchronizationState`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `botRegistration` | [BotRegistrationDetailsNoKind](#botregistrationdetailsnokind) |  |  |  |
| `provisioningStatus` | [ProvisioningState](#provisioningstate) |  |  |  |
| `state` | [SynchronizingState](#synchronizingstate) |  |  |  |
| `channels` | list of [SynchronizedChannelNoKind](#synchronizedchannelnokind) |  |  |  |

_No other properties are allowed._

### SynchronizationStateNoKind

*Title:* SynchronizationState  
Same shape as [SynchronizationState](#synchronizationstate) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `botRegistration` | [BotRegistrationDetailsNoKind](#botregistrationdetailsnokind) |  |  |  |
| `provisioningStatus` | [ProvisioningState](#provisioningstate) |  |  |  |
| `state` | [SynchronizingState](#synchronizingstate) |  |  |  |
| `channels` | list of [SynchronizedChannelNoKind](#synchronizedchannelnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotSynchronizationDetails](#botsynchronizationdetails), [BotSynchronizationDetailsNoKind](#botsynchronizationdetailsnokind)</sub>

### SynchronizedChannel

YAML: `kind: SynchronizedChannel`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `channelName` | `string` |  |  |  |

_No other properties are allowed._

### SynchronizedChannelNoKind

*Title:* SynchronizedChannel  
Same shape as [SynchronizedChannel](#synchronizedchannel) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `channelName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [SynchronizationState](#synchronizationstate), [SynchronizationStateNoKind](#synchronizationstatenokind)</sub>

### SynchronizingState


**String enum.** Allowed values: `Synchronized`, `Synchronizing`

<sub>Used by: [SynchronizationState](#synchronizationstate), [SynchronizationStateNoKind](#synchronizationstatenokind)</sub>

### SystemUser

YAML: `kind: SystemUser`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `id` | `string` |  |  |  |
| `aadUserId` | `string` |  |  |  |

_No other properties are allowed._

### TableSearchType


**String enum.** Allowed values: `DataverseSearch`, `FederatedSearch`

<sub>Used by: [DataverseTableSearch](#dataversetablesearch), [FederatedTableSearch](#federatedtablesearch)</sub>

### TeamsApplicationDetails

YAML: `kind: TeamsApplicationDetails`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `lobDisplayName` | `string` |  |  |  |
| `version` | `string` |  |  |  |
| `botChannelRegistrationAppId` | `string` |  |  |  |
| `colorIcon` | `string` |  |  |  |
| `outlineIcon` | `string` |  |  |  |
| `accentColor` | `string` |  |  |  |
| `shortDescription` | `string` |  |  |  |
| `longDescription` | `string` |  |  |  |
| `disclaimer` | `string` |  |  |  |
| `developerName` | `string` |  |  |  |
| `teamsCollaborationEnabled` | `boolean` |  |  |  |
| `websiteLink` | `string` |  |  |  |
| `termsLink` | `string` |  |  |  |
| `supportsCalling` | `boolean` |  |  |  |
| `privacyLink` | `string` |  |  |  |
| `microsoftPartnerNetworkId` | `string` |  |  |  |
| `isPowerPlatformStoreDiscoverable` | `boolean` |  |  |  |
| `singleSignOnAadApplicationClientId` | `string` |  |  |  |
| `singleSignOnResourceUri` | `string` |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `isTeamsIvrEnabled` | `boolean` |  |  |  |
| `scopes` | list of `string` |  |  |  |

_No other properties are allowed._

### TeamsApplicationDetailsNoKind

*Title:* TeamsApplicationDetails  
Same shape as [TeamsApplicationDetails](#teamsapplicationdetails) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `lobDisplayName` | `string` |  |  |  |
| `version` | `string` |  |  |  |
| `botChannelRegistrationAppId` | `string` |  |  |  |
| `colorIcon` | `string` |  |  |  |
| `outlineIcon` | `string` |  |  |  |
| `accentColor` | `string` |  |  |  |
| `shortDescription` | `string` |  |  |  |
| `longDescription` | `string` |  |  |  |
| `disclaimer` | `string` |  |  |  |
| `developerName` | `string` |  |  |  |
| `teamsCollaborationEnabled` | `boolean` |  |  |  |
| `websiteLink` | `string` |  |  |  |
| `termsLink` | `string` |  |  |  |
| `supportsCalling` | `boolean` |  |  |  |
| `privacyLink` | `string` |  |  |  |
| `microsoftPartnerNetworkId` | `string` |  |  |  |
| `isPowerPlatformStoreDiscoverable` | `boolean` |  |  |  |
| `singleSignOnAadApplicationClientId` | `string` |  |  |  |
| `singleSignOnResourceUri` | `string` |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `isTeamsIvrEnabled` | `boolean` |  |  |  |
| `scopes` | list of `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [ApplicationDetails](#applicationdetails), [ApplicationDetailsNoKind](#applicationdetailsnokind)</sub>

### TestCaseComponent

YAML: `kind: TestCaseComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `definition` | [TestDefinitionBase](evaluation.md#testdefinitionbase) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### ToolManifestResource

YAML: `kind: ToolManifestResource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `options` | object |  |  |  |

_No other properties are allowed._

<sub>Used by: [ManifestResources](#manifestresources)</sub>

### ToolsSelectionMode


**One of:** [UseAllTools](#usealltools), [UseSpecificTools](#usespecifictools)

<sub>Used by: [ModelContextProtocolMetadata](#modelcontextprotocolmetadata)</sub>

### TranslationsComponent

YAML: `kind: TranslationsComponent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `language` | [Language](settings.md#language) |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `content` | [LocalizableContentContainer](dialogs.md#localizablecontentcontainer) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [BotComponentBase](#botcomponentbase), [SolutionComponent](#solutioncomponent)</sub>

### TranslationsComponentNoKind

*Title:* TranslationsComponent  
Same shape as [TranslationsComponent](#translationscomponent) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `managedProperties` | [ManagedPropertiesNoKind](#managedpropertiesnokind) |  |  |  |
| `auditInfo` | [AuditInfoNoKind](#auditinfonokind) |  |  |  |
| `displayName` | `string` |  |  |  |
| `id` | `string` |  |  |  |
| `parentBotComponentId` | `string` |  |  |  |
| `parentBotId` | `string` |  |  |  |
| `parentBotComponentCollectionId` | `string` |  |  |  |
| `description` | `string` |  |  |  |
| `shareContext` | [ContentShareContextNoKind](#contentsharecontextnokind) |  |  |  |
| `category` | `string` |  |  |  |
| `state` | [StateCode](#statecode) |  | `"Active"` |  |
| `status` | [BotComponentStatusReason](#botcomponentstatusreason) |  |  |  |
| `publisherUniqueName` | `string` |  |  |  |
| `language` | [Language](settings.md#language) |  |  |  |
| `schemaName` | `string` | ✓ |  |  |
| `content` | [LocalizableContentContainer](dialogs.md#localizablecontentcontainer) | ✓ |  |  |

_No other properties are allowed._

### UseAllTools

YAML: `kind: UseAllTools`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [ToolsSelectionMode](#toolsselectionmode)</sub>

### UserPrivilege

YAML: `kind: UserPrivilege`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `version` | `number` |  |  |  |
| `id` | `string` |  |  |  |
| `name` | `string` |  |  |  |

_No other properties are allowed._

### UseSpecificTools

YAML: `kind: UseSpecificTools`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `tools` | list of `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [ToolsSelectionMode](#toolsselectionmode)</sub>

### VariableNoKind

Same shape as [Variable](actions.md#variable) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `name` | `string` | ✓ |  |  |
| `scope` | [VariableScope](actions.md#variablescope) |  | `"Conversation"` |  |
| `defaultValue` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `isExternalInitializationAllowed` | `boolean` |  | `false` |  |
| `isOutputToExternalCallers` | `boolean` |  | `false` |  |
| `initializationTimeoutInMilliseconds` | `number` |  | `0` |  |
| `missingInitializationBehavior` | [InitializationBehavior](actions.md#initializationbehavior) |  | `"UseDefault"` |  |
| `aIVisibility` | [VariableAIVisibility](actions.md#variableaivisibility) |  | `"Hidden"` |  |

_No other properties are allowed._

<sub>Used by: [GlobalVariableComponent](#globalvariablecomponent)</sub>

### WorkflowExternalTrigger

YAML: `kind: WorkflowExternalTrigger`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `flowId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [ExternalTriggerDefinition](#externaltriggerdefinition)</sub>
