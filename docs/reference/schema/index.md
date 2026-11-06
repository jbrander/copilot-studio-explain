# Copilot Studio YAML schema — A–Z index

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

All **744** definitions in the Copilot Studio YAML authoring schema. A file's root must be one of the root kinds the schema allows ([AdaptiveDialog](dialogs.md#adaptivedialog), [TaskDialog](dialogs.md#taskdialog)); other files (agent, knowledge, variables, settings) use component kinds that the authoring tools validate separately.

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

## Pages

- [Components & agent definitions](components.md) — 168 definitions
- [Dialogs (topics, tools, child agents) & recognizers](dialogs.md) — 28 definitions
- [Triggers](triggers.md) — 30 definitions
- [Actions (topic nodes)](actions.md) — 137 definitions
- [Tool (TaskDialog) actions & inputs](task-actions.md) — 22 definitions
- [Knowledge sources](knowledge.md) — 62 definitions
- [Entities](entities.md) — 51 definitions
- [Data types](data-types.md) — 31 definitions
- [Settings](settings.md) — 103 definitions
- [Evaluation & testing](evaluation.md) — 52 definitions
- [Messages, activities & cards](messages.md) — 42 definitions
- [Other definitions](other.md) — 18 definitions

## All definitions

| Definition | Page | `kind:` | Shape |
|---|---|---|---|
| [A2AExtensions](components.md#a2aextensions) | Components & agent definitions | ✓ | object |
| [A2AExtensionsNoKind](components.md#a2aextensionsnokind) | Components & agent definitions |  | object |
| [AbstentionAssertion](evaluation.md#abstentionassertion) | Evaluation & testing | ✓ | object |
| [ActionInputBinding](actions.md#actioninputbinding) | Actions (topic nodes) | ✓ | object |
| [ActionInputBindingBase](actions.md#actioninputbindingbase) | Actions (topic nodes) |  | union |
| [ActionInputBindingNoKind](actions.md#actioninputbindingnokind) | Actions (topic nodes) |  | object |
| [ActionOutputBinding](actions.md#actionoutputbinding) | Actions (topic nodes) | ✓ | object |
| [ActionOutputBindingBase](actions.md#actionoutputbindingbase) | Actions (topic nodes) |  | union |
| [ActionOutputBindingNoKind](actions.md#actionoutputbindingnokind) | Actions (topic nodes) |  | object |
| [ActionScope](actions.md#actionscope) | Actions (topic nodes) | ✓ | object |
| [ActivateExternalTrigger](actions.md#activateexternaltrigger) | Actions (topic nodes) | ✓ | object |
| [Activity](messages.md#activity) | Messages, activities & cards | ✓ | object |
| [ActivityAttachment](messages.md#activityattachment) | Messages, activities & cards |  | union |
| [ActivityImportance](messages.md#activityimportance) | Messages, activities & cards |  | enum |
| [ActivityProtocolMetadata](components.md#activityprotocolmetadata) | Components & agent definitions | ✓ | object |
| [ActivityTemplateBase](messages.md#activitytemplatebase) | Messages, activities & cards |  | union |
| [ActivityType](messages.md#activitytype) | Messages, activities & cards |  | enum |
| [AdaptationManagementServiceModelProfile](settings.md#adaptationmanagementservicemodelprofile) | Settings | ✓ | object |
| [AdaptationManagementServiceModelProfileNoKind](settings.md#adaptationmanagementservicemodelprofilenokind) | Settings |  | object |
| [AdaptationManagementServiceModelType](settings.md#adaptationmanagementservicemodeltype) | Settings | ✓ | object |
| [AdaptationManagementServiceModelTypeNoKind](settings.md#adaptationmanagementservicemodeltypenokind) | Settings |  | object |
| [AdaptationManagementServiceTrainingMetadata](components.md#adaptationmanagementservicetrainingmetadata) | Components & agent definitions | ✓ | object |
| [AdaptationManagementServiceTrainingMetadataNoKind](components.md#adaptationmanagementservicetrainingmetadatanokind) | Components & agent definitions |  | object |
| [AdaptiveCardExpression](other.md#adaptivecardexpression) | Other definitions |  | string |
| [AdaptiveCardPrompt](actions.md#adaptivecardprompt) | Actions (topic nodes) | ✓ | object |
| [AdaptiveCardTemplate](messages.md#adaptivecardtemplate) | Messages, activities & cards | ✓ | object |
| [AdaptiveDialog](dialogs.md#adaptivedialog) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [AdaptiveDialogNoKind](dialogs.md#adaptivedialognokind) | Dialogs (topics, tools, child agents) & recognizers |  | object |
| [AddItemOperation](actions.md#additemoperation) | Actions (topic nodes) | ✓ | object |
| [AdvanceNLURecognizer](dialogs.md#advancenlurecognizer) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [AgentDefinition](components.md#agentdefinition) | Components & agent definitions |  | union |
| [AgentDialog](dialogs.md#agentdialog) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [AgentEvaluationTarget](evaluation.md#agentevaluationtarget) | Evaluation & testing | ✓ | object |
| [AgentMessageContentType](messages.md#agentmessagecontenttype) | Messages, activities & cards |  | enum |
| [AgentMessageRole](messages.md#agentmessagerole) | Messages, activities & cards |  | enum |
| [AgentMessageSortOrder](messages.md#agentmessagesortorder) | Messages, activities & cards |  | enum |
| [AgentModelType](settings.md#agentmodeltype) | Settings |  | union |
| [AgentToAgentProtocolMetadata](components.md#agenttoagentprotocolmetadata) | Components & agent definitions | ✓ | object |
| [AgePrebuiltEntity](entities.md#ageprebuiltentity) | Entities |  | union |
| [AIModelDefinition](components.md#aimodeldefinition) | Components & agent definitions | ✓ | object |
| [AIModelDefinitionNoKind](components.md#aimodeldefinitionnokind) | Components & agent definitions |  | object |
| [AIPluginComponent](components.md#aiplugincomponent) | Components & agent definitions | ✓ | object |
| [AIPluginConnectionReference](task-actions.md#aipluginconnectionreference) | Tool (TaskDialog) actions & inputs |  | union |
| [AIPluginOperation](components.md#aipluginoperation) | Components & agent definitions | ✓ | object |
| [AIPluginOperationComponent](components.md#aipluginoperationcomponent) | Components & agent definitions | ✓ | object |
| [AIPluginOperationNoKind](components.md#aipluginoperationnokind) | Components & agent definitions |  | object |
| [AiSensitivityLevel](data-types.md#aisensitivitylevel) | Data types |  | enum |
| [AISettings](settings.md#aisettings) | Settings | ✓ | object |
| [AISettingsNoKind](settings.md#aisettingsnokind) | Settings |  | object |
| [AlwaysRequireMCPApproval](settings.md#alwaysrequiremcpapproval) | Settings | ✓ | object |
| [AmbiguousDateTimeResolutionBehavior](entities.md#ambiguousdatetimeresolutionbehavior) | Entities |  | enum |
| [AnalyticsSettings](settings.md#analyticssettings) | Settings | ✓ | object |
| [AnalyticsSettingsNoKind](settings.md#analyticssettingsnokind) | Settings |  | object |
| [AndConditionGroup](actions.md#andconditiongroup) | Actions (topic nodes) | ✓ | object |
| [AnnotatedSample](evaluation.md#annotatedsample) | Evaluation & testing |  | string |
| [AnnotatedSampleNoKind](evaluation.md#annotatedsamplenokind) | Evaluation & testing |  | string |
| [AnswerQuestionWithAI](actions.md#answerquestionwithai) | Actions (topic nodes) | ✓ | object |
| [Any](data-types.md#any) | Data types |  | union |
| [ApplicationDetails](components.md#applicationdetails) | Components & agent definitions | ✓ | object |
| [ApplicationDetailsNoKind](components.md#applicationdetailsnokind) | Components & agent definitions |  | object |
| [AttachmentLayout](messages.md#attachmentlayout) | Messages, activities & cards |  | enum |
| [AuditInfo](components.md#auditinfo) | Components & agent definitions | ✓ | object |
| [AuditInfoNoKind](components.md#auditinfonokind) | Components & agent definitions |  | object |
| [AutomaticHoldSettings](actions.md#automaticholdsettings) | Actions (topic nodes) | ✓ | object |
| [AutomaticTaskInput](task-actions.md#automatictaskinput) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [AutomaticTransferContext](actions.md#automatictransfercontext) | Actions (topic nodes) | ✓ | object |
| [AutomaticTransferVariables](actions.md#automatictransfervariables) | Actions (topic nodes) |  | enum |
| [AutonomousSettings](settings.md#autonomoussettings) | Settings | ✓ | object |
| [AzureAISearchSource](knowledge.md#azureaisearchsource) | Knowledge sources | ✓ | object |
| [AzureCognitiveSearchInternalDataSource](knowledge.md#azurecognitivesearchinternaldatasource) | Knowledge sources | ✓ | object |
| [AzureCognitiveSearchInternalDataSourceNoKind](knowledge.md#azurecognitivesearchinternaldatasourcenokind) | Knowledge sources |  | object |
| [AzureOpenAIOnYourDataKnowledgeSource](knowledge.md#azureopenaionyourdataknowledgesource) | Knowledge sources | ✓ | object |
| [AzureOpenAIOnYourDataSource](knowledge.md#azureopenaionyourdatasource) | Knowledge sources | ✓ | object |
| [AzureOpenAIOnYourDataSourceNoKind](actions.md#azureopenaionyourdatasourcenokind) | Actions (topic nodes) |  | object |
| [BaseSearchConfiguration](knowledge.md#basesearchconfiguration) | Knowledge sources |  | union |
| [BeginDialog](actions.md#begindialog) | Actions (topic nodes) | ✓ | object |
| [BingCustomSearchDataSource](knowledge.md#bingcustomsearchdatasource) | Knowledge sources | ✓ | object |
| [BingCustomSearchSource](knowledge.md#bingcustomsearchsource) | Knowledge sources | ✓ | object |
| [BingSearchOptions](knowledge.md#bingsearchoptions) | Knowledge sources | ✓ | object |
| [BingSearchOptionsNoKind](actions.md#bingsearchoptionsnokind) | Actions (topic nodes) |  | object |
| [Blank](data-types.md#blank) | Data types |  | union |
| [Boolean](data-types.md#boolean) | Data types |  | union |
| [BooleanCondition](actions.md#booleancondition) | Actions (topic nodes) | ✓ | object |
| [BooleanConditionBase](actions.md#booleanconditionbase) | Actions (topic nodes) |  | union |
| [BooleanConditionGroup](actions.md#booleanconditiongroup) | Actions (topic nodes) |  | union |
| [BooleanConditionOperator](actions.md#booleanconditionoperator) | Actions (topic nodes) |  | enum |
| [BooleanPrebuiltEntity](entities.md#booleanprebuiltentity) | Entities |  | union |
| [BoolExpression](other.md#boolexpression) | Other definitions |  | union |
| [BotAccessControlPolicy](components.md#botaccesscontrolpolicy) | Components & agent definitions |  | enum |
| [BotAuthenticationConfiguration](components.md#botauthenticationconfiguration) | Components & agent definitions | ✓ | object |
| [BotAuthenticationConfigurationNoKind](components.md#botauthenticationconfigurationnokind) | Components & agent definitions |  | object |
| [BotAuthenticationMode](components.md#botauthenticationmode) | Components & agent definitions |  | enum |
| [BotAuthenticationTrigger](components.md#botauthenticationtrigger) | Components & agent definitions |  | enum |
| [BotComponentBase](components.md#botcomponentbase) | Components & agent definitions |  | union |
| [BotComponentCollection](components.md#botcomponentcollection) | Components & agent definitions | ✓ | object |
| [BotComponentCollectionConfiguration](components.md#botcomponentcollectionconfiguration) | Components & agent definitions | ✓ | object |
| [BotComponentCollectionConfigurationNoKind](components.md#botcomponentcollectionconfigurationnokind) | Components & agent definitions |  | object |
| [BotComponentCollectionDefinition](components.md#botcomponentcollectiondefinition) | Components & agent definitions | ✓ | object |
| [BotComponentCollectionNoKind](components.md#botcomponentcollectionnokind) | Components & agent definitions |  | object |
| [BotComponentReusePolicy](settings.md#botcomponentreusepolicy) | Settings |  | enum |
| [BotComponentStatusReason](components.md#botcomponentstatusreason) | Components & agent definitions |  | enum |
| [BotConfiguration](components.md#botconfiguration) | Components & agent definitions | ✓ | object |
| [BotConfigurationNoKind](components.md#botconfigurationnokind) | Components & agent definitions |  | object |
| [BotDefinition](components.md#botdefinition) | Components & agent definitions | ✓ | object |
| [BotElement](other.md#botelement) | Other definitions |  | union |
| [BotEntity](components.md#botentity) | Components & agent definitions | ✓ | object |
| [BotEntityNoKind](components.md#botentitynokind) | Components & agent definitions |  | object |
| [BotFileReference](messages.md#botfilereference) | Messages, activities & cards | ✓ | object |
| [BotRegistrationDetails](components.md#botregistrationdetails) | Components & agent definitions | ✓ | object |
| [BotRegistrationDetailsNoKind](components.md#botregistrationdetailsnokind) | Components & agent definitions |  | object |
| [BotSettingsBase](settings.md#botsettingsbase) | Settings |  | union |
| [BotSettingsComponent](components.md#botsettingscomponent) | Components & agent definitions | ✓ | object |
| [BotSpeechSettings](settings.md#botspeechsettings) | Settings | ✓ | object |
| [BotSpeechSettingsNoKind](settings.md#botspeechsettingsnokind) | Settings |  | object |
| [BotSynchronizationDetails](components.md#botsynchronizationdetails) | Components & agent definitions | ✓ | object |
| [BotSynchronizationDetailsNoKind](components.md#botsynchronizationdetailsnokind) | Components & agent definitions |  | object |
| [BreakLoop](actions.md#breakloop) | Actions (topic nodes) | ✓ | object |
| [CancelAllDialogs](actions.md#cancelalldialogs) | Actions (topic nodes) | ✓ | object |
| [CancelDialog](actions.md#canceldialog) | Actions (topic nodes) | ✓ | object |
| [Category](components.md#category) | Components & agent definitions | ✓ | object |
| [CategoryNoKind](components.md#categorynokind) | Components & agent definitions |  | object |
| [ChannelDefinition](components.md#channeldefinition) | Components & agent definitions | ✓ | object |
| [ChannelDefinitionNoKind](components.md#channeldefinitionnokind) | Components & agent definitions |  | object |
| [ChannelId](components.md#channelid) | Components & agent definitions |  | enum |
| [Citation](knowledge.md#citation) | Knowledge sources | ✓ | object |
| [CitationNoKind](knowledge.md#citationnokind) | Knowledge sources |  | object |
| [CitationsInModelOutput](knowledge.md#citationsinmodeloutput) | Knowledge sources | ✓ | object |
| [CityPrebuiltEntity](entities.md#cityprebuiltentity) | Entities |  | union |
| [ClassicAIRecognizer](dialogs.md#classicairecognizer) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [ClassicVoiceProcessingMode](settings.md#classicvoiceprocessingmode) | Settings | ✓ | object |
| [ClearAllVariables](actions.md#clearallvariables) | Actions (topic nodes) | ✓ | object |
| [ClearItemsOperation](actions.md#clearitemsoperation) | Actions (topic nodes) | ✓ | object |
| [CLIAgentRecognizer](dialogs.md#cliagentrecognizer) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [ClosedListConversionMetadata](entities.md#closedlistconversionmetadata) | Entities | ✓ | object |
| [ClosedListConversionMetadataNoKind](dialogs.md#closedlistconversionmetadatanokind) | Dialogs (topics, tools, child agents) & recognizers |  | object |
| [ClosedListEntity](dialogs.md#closedlistentity) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [ClosedListEntityReference](entities.md#closedlistentityreference) | Entities | ✓ | object |
| [ClosedListItem](entities.md#closedlistitem) | Entities | ✓ | object |
| [ClosedListItemNoKind](dialogs.md#closedlistitemnokind) | Dialogs (topics, tools, child agents) & recognizers |  | object |
| [ClosedListOptionSet](data-types.md#closedlistoptionset) | Data types | ✓ | object |
| [CloudFlowConnectionType](components.md#cloudflowconnectiontype) | Components & agent definitions |  | enum |
| [CloudFlowDefinition](components.md#cloudflowdefinition) | Components & agent definitions | ✓ | object |
| [CloudFlowDefinitionNoKind](components.md#cloudflowdefinitionnokind) | Components & agent definitions |  | object |
| [CloudFlowTriggerType](components.md#cloudflowtriggertype) | Components & agent definitions |  | enum |
| [CodeGrader](evaluation.md#codegrader) | Evaluation & testing | ✓ | object |
| [Color](data-types.md#color) | Data types |  | union |
| [ColorPrebuiltEntity](entities.md#colorprebuiltentity) | Entities |  | union |
| [CompareMeaningGrader](evaluation.md#comparemeaninggrader) | Evaluation & testing | ✓ | object |
| [CompletenessAssertion](evaluation.md#completenessassertion) | Evaluation & testing | ✓ | object |
| [ComputerUsingAgentAllowlist](settings.md#computerusingagentallowlist) | Settings | ✓ | object |
| [ComputerUsingAgentAllowlistNoKind](task-actions.md#computerusingagentallowlistnokind) | Tool (TaskDialog) actions & inputs |  | object |
| [ConditionGroup](actions.md#conditiongroup) | Actions (topic nodes) | ✓ | object |
| [ConditionItem](actions.md#conditionitem) | Actions (topic nodes) | ✓ | object |
| [ConditionItemNoKind](actions.md#conditionitemnokind) | Actions (topic nodes) |  | object |
| [ConnectedAgentDefinition](components.md#connectedagentdefinition) | Components & agent definitions | ✓ | object |
| [ConnectedAgentDefinitionNoKind](components.md#connectedagentdefinitionnokind) | Components & agent definitions |  | object |
| [Connection](settings.md#connection) | Settings |  | union |
| [ConnectionMode](components.md#connectionmode) | Components & agent definitions |  | enum |
| [ConnectionParameterSettings](settings.md#connectionparametersettings) | Settings | ✓ | object |
| [ConnectionParametersSet](components.md#connectionparametersset) | Components & agent definitions | ✓ | object |
| [ConnectionParametersSetNoKind](components.md#connectionparameterssetnokind) | Components & agent definitions |  | object |
| [ConnectionProperties](settings.md#connectionproperties) | Settings | ✓ | object |
| [ConnectionPropertiesNoKind](settings.md#connectionpropertiesnokind) | Settings |  | object |
| [ConnectionReference](components.md#connectionreference) | Components & agent definitions | ✓ | object |
| [ConnectionReferenceNameInFlowToConnectionParametersSetMapping](components.md#connectionreferencenameinflowtoconnectionparameterssetmapping) | Components & agent definitions | ✓ | object |
| [ConnectionReferenceNameInFlowToConnectionParametersSetMappingNoKind](settings.md#connectionreferencenameinflowtoconnectionparameterssetmappingnokind) | Settings |  | object |
| [ConnectionReferenceNoKind](components.md#connectionreferencenokind) | Components & agent definitions |  | object |
| [ConnectionReferencesSourceFile](components.md#connectionreferencessourcefile) | Components & agent definitions | ✓ | object |
| [ConnectorDefinition](components.md#connectordefinition) | Components & agent definitions | ✓ | object |
| [ConnectorDefinitionComponent](components.md#connectordefinitioncomponent) | Components & agent definitions | ✓ | object |
| [ConnectorDefinitionNoKind](components.md#connectordefinitionnokind) | Components & agent definitions |  | object |
| [ConnectorDefinitionType](components.md#connectordefinitiontype) | Components & agent definitions |  | enum |
| [ConnectorDynamicValuesCapability](components.md#connectordynamicvaluescapability) | Components & agent definitions |  | enum |
| [ConnectorDynamicValuesConfig](components.md#connectordynamicvaluesconfig) | Components & agent definitions | ✓ | object |
| [ConnectorDynamicValuesConfigNoKind](data-types.md#connectordynamicvaluesconfignokind) | Data types |  | object |
| [ConnectorOperation](components.md#connectoroperation) | Components & agent definitions | ✓ | object |
| [ConnectorOperationNoKind](components.md#connectoroperationnokind) | Components & agent definitions |  | object |
| [ConnectorTriggerDefinition](components.md#connectortriggerdefinition) | Components & agent definitions | ✓ | object |
| [ConnectorTriggerOperation](components.md#connectortriggeroperation) | Components & agent definitions | ✓ | object |
| [ConnectorTriggerOperationNoKind](components.md#connectortriggeroperationnokind) | Components & agent definitions |  | object |
| [ContainsAllAssertion](evaluation.md#containsallassertion) | Evaluation & testing | ✓ | object |
| [ContainsAllGrader](evaluation.md#containsallgrader) | Evaluation & testing | ✓ | object |
| [ContainsAnyAssertion](evaluation.md#containsanyassertion) | Evaluation & testing | ✓ | object |
| [ContainsAnyGrader](evaluation.md#containsanygrader) | Evaluation & testing | ✓ | object |
| [ContentModerationLevel](settings.md#contentmoderationlevel) | Settings |  | enum |
| [ContentModerationSettings](settings.md#contentmoderationsettings) | Settings | ✓ | object |
| [ContentShareContext](messages.md#contentsharecontext) | Messages, activities & cards | ✓ | object |
| [ContentShareContextNoKind](components.md#contentsharecontextnokind) | Components & agent definitions |  | object |
| [ContinentPrebuiltEntity](entities.md#continentprebuiltentity) | Entities |  | union |
| [ContinueConversationConnectorTrigger](triggers.md#continueconversationconnectortrigger) | Triggers | ✓ | object |
| [ContinueLoop](actions.md#continueloop) | Actions (topic nodes) | ✓ | object |
| [ContinueOnErrorBehavior](actions.md#continueonerrorbehavior) | Actions (topic nodes) | ✓ | object |
| [ConversationFileReference](messages.md#conversationfilereference) | Messages, activities & cards | ✓ | object |
| [ConversationHistory](settings.md#conversationhistory) | Settings | ✓ | object |
| [ConversationOutcome](actions.md#conversationoutcome) | Actions (topic nodes) |  | enum |
| [ConversationStarter](components.md#conversationstarter) | Components & agent definitions | ✓ | object |
| [ConversationStarterNoKind](components.md#conversationstarternokind) | Components & agent definitions |  | object |
| [ConversionMetadata](components.md#conversionmetadata) | Components & agent definitions | ✓ | object |
| [ConversionMetadataNoKind](dialogs.md#conversionmetadatanokind) | Dialogs (topics, tools, child agents) & recognizers |  | object |
| [CopilotChatApplicationDetails](components.md#copilotchatapplicationdetails) | Components & agent definitions | ✓ | object |
| [CopilotChatApplicationDetailsNoKind](components.md#copilotchatapplicationdetailsnokind) | Components & agent definitions |  | object |
| [CopilotConnectorEntity](entities.md#copilotconnectorentity) | Entities | ✓ | object |
| [CopilotConnectorEntityNoKind](knowledge.md#copilotconnectorentitynokind) | Knowledge sources |  | object |
| [CopilotConnectorKnowledgeSourceDemo](knowledge.md#copilotconnectorknowledgesourcedemo) | Knowledge sources | ✓ | object |
| [CopilotPolicy](settings.md#copilotpolicy) | Settings | ✓ | object |
| [CopilotPolicyNoKind](settings.md#copilotpolicynokind) | Settings |  | object |
| [CosineEmbeddingSimilarityAssertion](evaluation.md#cosineembeddingsimilarityassertion) | Evaluation & testing | ✓ | object |
| [CountryOrRegionPrebuiltEntity](entities.md#countryorregionprebuiltentity) | Entities |  | union |
| [CreateConversationConnectorTrigger](triggers.md#createconversationconnectortrigger) | Triggers | ✓ | object |
| [CreateSearchQuery](actions.md#createsearchquery) | Actions (topic nodes) | ✓ | object |
| [CSATQuestion](actions.md#csatquestion) | Actions (topic nodes) | ✓ | object |
| [CurrentModels](settings.md#currentmodels) | Settings | ✓ | object |
| [CurrentModelsNoKind](settings.md#currentmodelsnokind) | Settings |  | object |
| [CustomDataSource](knowledge.md#customdatasource) | Knowledge sources | ✓ | object |
| [CustomDataSourceNoKind](actions.md#customdatasourcenokind) | Actions (topic nodes) |  | object |
| [CustomEntityComponent](components.md#customentitycomponent) | Components & agent definitions | ✓ | object |
| [CustomKnowledgeSource](knowledge.md#customknowledgesource) | Knowledge sources | ✓ | object |
| [CustomMetricCategory](components.md#custommetriccategory) | Components & agent definitions | ✓ | object |
| [CustomMetricCategoryNoKind](components.md#custommetriccategorynokind) | Components & agent definitions |  | object |
| [CustomMetricDefinition](components.md#custommetricdefinition) | Components & agent definitions | ✓ | object |
| [CustomMetricDefinitionComponent](components.md#custommetricdefinitioncomponent) | Components & agent definitions | ✓ | object |
| [CustomMetricDefinitionNoKind](components.md#custommetricdefinitionnokind) | Components & agent definitions |  | object |
| [CustomModelCitationHandlingType](knowledge.md#custommodelcitationhandlingtype) | Knowledge sources |  | union |
| [CustomModels](settings.md#custommodels) | Settings | ✓ | object |
| [CustomSearchInputContent](knowledge.md#customsearchinputcontent) | Knowledge sources | ✓ | object |
| [CustomSearchInputContentNoKind](knowledge.md#customsearchinputcontentnokind) | Knowledge sources |  | object |
| [DataPolicyEvaluationViolationType](evaluation.md#datapolicyevaluationviolationtype) | Evaluation & testing |  | enum |
| [DataType](data-types.md#datatype) | Data types |  | union |
| [DataValue](other.md#datavalue) | Other definitions |  | union |
| [DataverseStructuredSearchSource](knowledge.md#dataversestructuredsearchsource) | Knowledge sources | ✓ | object |
| [DataverseTableSearch](components.md#dataversetablesearch) | Components & agent definitions | ✓ | object |
| [DataverseTableSearchBase](knowledge.md#dataversetablesearchbase) | Knowledge sources |  | union |
| [DataverseTableSearchEntityColumnSynonym](components.md#dataversetablesearchentitycolumnsynonym) | Components & agent definitions | ✓ | object |
| [DataverseTableSearchEntityColumnSynonymNoKind](components.md#dataversetablesearchentitycolumnsynonymnokind) | Components & agent definitions |  | object |
| [DataverseTableSearchEntityConfiguration](components.md#dataversetablesearchentityconfiguration) | Components & agent definitions | ✓ | object |
| [DataverseTableSearchEntityConfigurationNoKind](components.md#dataversetablesearchentityconfigurationnokind) | Components & agent definitions |  | object |
| [DataverseTableSearchGlossaryConfiguration](components.md#dataversetablesearchglossaryconfiguration) | Components & agent definitions | ✓ | object |
| [DataverseTableSearchGlossaryConfigurationNoKind](components.md#dataversetablesearchglossaryconfigurationnokind) | Components & agent definitions |  | object |
| [Date](data-types.md#date) | Data types |  | union |
| [DatePrebuiltEntity](entities.md#dateprebuiltentity) | Entities |  | union |
| [DateTime](data-types.md#datetime) | Data types |  | union |
| [DateTimeNoTimeZone](data-types.md#datetimenotimezone) | Data types |  | union |
| [DateTimeNoTimeZonePrebuiltEntity](entities.md#datetimenotimezoneprebuiltentity) | Entities |  | union |
| [DateTimePrebuiltEntity](entities.md#datetimeprebuiltentity) | Entities |  | union |
| [DefaultFeedbackCollection](settings.md#defaultfeedbackcollection) | Settings | ✓ | object |
| [DefaultHoldSettings](settings.md#defaultholdsettings) | Settings | ✓ | object |
| [DefaultHoldSettingsNoKind](settings.md#defaultholdsettingsnokind) | Settings |  | object |
| [DefaultResponseMode](components.md#defaultresponsemode) | Components & agent definitions |  | enum |
| [DefaultValueMissingAction](actions.md#defaultvaluemissingaction) | Actions (topic nodes) |  | enum |
| [DeferredUpdateModels](settings.md#deferredupdatemodels) | Settings | ✓ | object |
| [DefinitionBase](other.md#definitionbase) | Other definitions |  | union |
| [DeleteActivity](actions.md#deleteactivity) | Actions (topic nodes) | ✓ | object |
| [DiagnosticType](components.md#diagnostictype) | Components & agent definitions |  | enum |
| [DialogAction](actions.md#dialogaction) | Actions (topic nodes) |  | union |
| [DialogBase](dialogs.md#dialogbase) | Dialogs (topics, tools, child agents) & recognizers |  | union |
| [DialogComponent](components.md#dialogcomponent) | Components & agent definitions | ✓ | object |
| [DialogComponentNoKind](components.md#dialogcomponentnokind) | Components & agent definitions |  | object |
| [DialogComponentReference](components.md#dialogcomponentreference) | Components & agent definitions | ✓ | object |
| [DialogExpression](other.md#dialogexpression) | Other definitions |  | union |
| [DigitalListeningMode](settings.md#digitallisteningmode) | Settings |  | enum |
| [DigitalNoiseReductionMode](settings.md#digitalnoisereductionmode) | Settings |  | enum |
| [DigitalTurnDetectionType](settings.md#digitalturndetectiontype) | Settings |  | enum |
| [DigitalVoiceSettings](settings.md#digitalvoicesettings) | Settings | ✓ | object |
| [DisableToolAugmentation](task-actions.md#disabletoolaugmentation) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [DisableTrigger](actions.md#disabletrigger) | Actions (topic nodes) | ✓ | object |
| [DisconnectedNodeContainer](actions.md#disconnectednodecontainer) | Actions (topic nodes) | ✓ | object |
| [DlpComponentType](components.md#dlpcomponenttype) | Components & agent definitions |  | enum |
| [DoNotHoldSettings](actions.md#donotholdsettings) | Actions (topic nodes) | ✓ | object |
| [DoNotSearchFiles](actions.md#donotsearchfiles) | Actions (topic nodes) | ✓ | object |
| [DtmfMappingKey](settings.md#dtmfmappingkey) | Settings |  | enum |
| [DtmfMultipleChoiceOptions](settings.md#dtmfmultiplechoiceoptions) | Settings | ✓ | object |
| [DtmfMultipleChoiceOptionsNoKind](settings.md#dtmfmultiplechoiceoptionsnokind) | Settings |  | object |
| [DtmfTerminationKey](settings.md#dtmfterminationkey) | Settings |  | enum |
| [DtmfTerminationOptions](actions.md#dtmfterminationoptions) | Actions (topic nodes) |  | union |
| [DurationPrebuiltEntity](entities.md#durationprebuiltentity) | Entities |  | union |
| [DynamicClosedListEntity](entities.md#dynamicclosedlistentity) | Entities |  | union |
| [DynamicClosedListItem](entities.md#dynamicclosedlistitem) | Entities | ✓ | object |
| [DynamicClosedListItemNoKind](entities.md#dynamicclosedlistitemnokind) | Entities |  | object |
| [EditTable](actions.md#edittable) | Actions (topic nodes) | ✓ | object |
| [EditTableOperation](actions.md#edittableoperation) | Actions (topic nodes) |  | union |
| [EditTableV2](actions.md#edittablev2) | Actions (topic nodes) | ✓ | object |
| [EmailPrebuiltEntity](entities.md#emailprebuiltentity) | Entities |  | union |
| [EmailSearchSource](knowledge.md#emailsearchsource) | Knowledge sources | ✓ | object |
| [EmbeddedEntity](entities.md#embeddedentity) | Entities |  | union |
| [EmbeddedOptionSet](data-types.md#embeddedoptionset) | Data types | ✓ | object |
| [EmbeddedOptionSetNoKind](data-types.md#embeddedoptionsetnokind) | Data types |  | object |
| [EmitEvent](actions.md#emitevent) | Actions (topic nodes) | ✓ | object |
| [EnableToolAugmentation](task-actions.md#enabletoolaugmentation) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [EndConversation](actions.md#endconversation) | Actions (topic nodes) | ✓ | object |
| [EndDialog](actions.md#enddialog) | Actions (topic nodes) | ✓ | object |
| [Entity](entities.md#entity) | Entities |  | union |
| [EntityReference](entities.md#entityreference) | Entities |  | union |
| [EnvironmentVariableDefinition](components.md#environmentvariabledefinition) | Components & agent definitions | ✓ | object |
| [EnvironmentVariableDefinitionNoKind](components.md#environmentvariabledefinitionnokind) | Components & agent definitions |  | object |
| [EnvironmentVariableReference](components.md#environmentvariablereference) | Components & agent definitions | ✓ | object |
| [EnvironmentVariableReferenceNoKind](components.md#environmentvariablereferencenokind) | Components & agent definitions |  | object |
| [EnvironmentVariableSecretReadStatus](components.md#environmentvariablesecretreadstatus) | Components & agent definitions |  | enum |
| [EnvironmentVariableType](components.md#environmentvariabletype) | Components & agent definitions |  | enum |
| [EnvironmentVariableValue](components.md#environmentvariablevalue) | Components & agent definitions | ✓ | object |
| [EnvironmentVariableValueNoKind](components.md#environmentvariablevaluenokind) | Components & agent definitions |  | object |
| [EvaluationData](evaluation.md#evaluationdata) | Evaluation & testing | ✓ | object |
| [EvaluationRunDefinition](evaluation.md#evaluationrundefinition) | Evaluation & testing | ✓ | object |
| [EvaluationSet](evaluation.md#evaluationset) | Evaluation & testing | ✓ | object |
| [EvaluationTargetBase](evaluation.md#evaluationtargetbase) | Evaluation & testing |  | union |
| [EventActivity](messages.md#eventactivity) | Messages, activities & cards | ✓ | object |
| [EventPrebuiltEntity](entities.md#eventprebuiltentity) | Entities |  | union |
| [EventType](actions.md#eventtype) | Actions (topic nodes) |  | enum |
| [ExactMatchAssertion](evaluation.md#exactmatchassertion) | Evaluation & testing | ✓ | object |
| [ExactMatchGrader](evaluation.md#exactmatchgrader) | Evaluation & testing | ✓ | object |
| [ExpectedExecutionStepBase](evaluation.md#expectedexecutionstepbase) | Evaluation & testing |  | union |
| [ExpectedToolCallStep](evaluation.md#expectedtoolcallstep) | Evaluation & testing | ✓ | object |
| [ExpectedTopicTriggeredStep](evaluation.md#expectedtopictriggeredstep) | Evaluation & testing | ✓ | object |
| [ExperimentalModels](settings.md#experimentalmodels) | Settings | ✓ | object |
| [ExpressionErrorCode](other.md#expressionerrorcode) | Other definitions |  | enum |
| [ExpressionSegment](messages.md#expressionsegment) | Messages, activities & cards | ✓ | object |
| [ExpressionTransferContext](actions.md#expressiontransfercontext) | Actions (topic nodes) | ✓ | object |
| [ExternalChatCompletionModelConfiguration](actions.md#externalchatcompletionmodelconfiguration) | Actions (topic nodes) | ✓ | object |
| [ExternalChatCompletionModelConfigurationNoKind](actions.md#externalchatcompletionmodelconfigurationnokind) | Actions (topic nodes) |  | object |
| [ExternalEntity](dialogs.md#externalentity) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [ExternalEntityReference](entities.md#externalentityreference) | Entities |  | union |
| [ExternalIntent](dialogs.md#externalintent) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [ExternalModelConfiguration](actions.md#externalmodelconfiguration) | Actions (topic nodes) |  | union |
| [ExternalTriggerComponent](components.md#externaltriggercomponent) | Components & agent definitions | ✓ | object |
| [ExternalTriggerConfiguration](components.md#externaltriggerconfiguration) | Components & agent definitions | ✓ | object |
| [ExternalTriggerConfigurationNoKind](components.md#externaltriggerconfigurationnokind) | Components & agent definitions |  | object |
| [ExternalTriggerDefinition](components.md#externaltriggerdefinition) | Components & agent definitions |  | union |
| [FabricAISkillSource](knowledge.md#fabricaiskillsource) | Knowledge sources | ✓ | object |
| [FederatedStructuredSearchSource](knowledge.md#federatedstructuredsearchsource) | Knowledge sources | ✓ | object |
| [FederatedTableSearch](components.md#federatedtablesearch) | Components & agent definitions | ✓ | object |
| [File](data-types.md#file) | Data types |  | union |
| [FileAttachmentComponent](components.md#fileattachmentcomponent) | Components & agent definitions | ✓ | object |
| [FileAttachmentComponentMetadata](messages.md#fileattachmentcomponentmetadata) | Messages, activities & cards | ✓ | object |
| [FileAttachmentComponentMetadataNoKind](components.md#fileattachmentcomponentmetadatanokind) | Components & agent definitions |  | object |
| [FileExpression](other.md#fileexpression) | Other definitions |  | string |
| [FileGroupKnowledgeSource](knowledge.md#filegroupknowledgesource) | Knowledge sources | ✓ | object |
| [FileIndexingStatus](knowledge.md#fileindexingstatus) | Knowledge sources |  | enum |
| [FilePrebuiltEntity](entities.md#fileprebuiltentity) | Entities |  | union |
| [FileReference](messages.md#filereference) | Messages, activities & cards |  | union |
| [FileSearchDataSource](knowledge.md#filesearchdatasource) | Knowledge sources | ✓ | object |
| [FileSearchDataSourceNoKind](knowledge.md#filesearchdatasourcenokind) | Knowledge sources |  | object |
| [FileTemplate](messages.md#filetemplate) | Messages, activities & cards | ✓ | object |
| [FineTunedModelsPropertyDetails](settings.md#finetunedmodelspropertydetails) | Settings | ✓ | object |
| [FineTunedModelsPropertyDetailsNoKind](settings.md#finetunedmodelspropertydetailsnokind) | Settings |  | object |
| [FixedLengthTermination](actions.md#fixedlengthtermination) | Actions (topic nodes) | ✓ | object |
| [FlexibleLengthTermination](actions.md#flexiblelengthtermination) | Actions (topic nodes) | ✓ | object |
| [Float](data-types.md#float) | Data types |  | union |
| [FlowEvaluationTarget](evaluation.md#flowevaluationtarget) | Evaluation & testing | ✓ | object |
| [Foreach](actions.md#foreach) | Actions (topic nodes) | ✓ | object |
| [GeneralQualityGrader](evaluation.md#generalqualitygrader) | Evaluation & testing | ✓ | object |
| [GenerationSettings](settings.md#generationsettings) | Settings | ✓ | object |
| [GenerationSettingsNoKind](settings.md#generationsettingsnokind) | Settings |  | object |
| [GenerativeAIRecognizer](dialogs.md#generativeairecognizer) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [GenerativeModality](actions.md#generativemodality) | Actions (topic nodes) |  | enum |
| [GenerativeResponseOutputType](actions.md#generativeresponseoutputtype) | Actions (topic nodes) |  | enum |
| [GenericDatasetNameSearchConfiguration](knowledge.md#genericdatasetnamesearchconfiguration) | Knowledge sources | ✓ | object |
| [GetActivityMembers](actions.md#getactivitymembers) | Actions (topic nodes) | ✓ | object |
| [GetConversationMembers](actions.md#getconversationmembers) | Actions (topic nodes) | ✓ | object |
| [GlobalVariableComponent](components.md#globalvariablecomponent) | Components & agent definitions | ✓ | object |
| [GotoAction](actions.md#gotoaction) | Actions (topic nodes) | ✓ | object |
| [GptCapabilities](settings.md#gptcapabilities) | Settings | ✓ | object |
| [GptCapabilitiesNoKind](components.md#gptcapabilitiesnokind) | Components & agent definitions |  | object |
| [GptComponent](components.md#gptcomponent) | Components & agent definitions | ✓ | object |
| [GptComponentMetadata](components.md#gptcomponentmetadata) | Components & agent definitions | ✓ | object |
| [GptComponentMetadataNoKind](components.md#gptcomponentmetadatanokind) | Components & agent definitions |  | object |
| [GPTSettings](settings.md#gptsettings) | Settings | ✓ | object |
| [GPTSettingsNoKind](settings.md#gptsettingsnokind) | Settings |  | object |
| [GraderBase](evaluation.md#graderbase) | Evaluation & testing |  | union |
| [GraderOutcome](evaluation.md#graderoutcome) | Evaluation & testing |  | enum |
| [GraphConnectorItemByExternalId](knowledge.md#graphconnectoritembyexternalid) | Knowledge sources | ✓ | object |
| [GraphConnectorItemByExternalIdNoKind](knowledge.md#graphconnectoritembyexternalidnokind) | Knowledge sources |  | object |
| [GraphConnectorItemByPath](knowledge.md#graphconnectoritembypath) | Knowledge sources | ✓ | object |
| [GraphConnectorItemByPathNoKind](knowledge.md#graphconnectoritembypathnokind) | Knowledge sources |  | object |
| [GraphConnectorItemByUrl](knowledge.md#graphconnectoritembyurl) | Knowledge sources | ✓ | object |
| [GraphConnectorItemByUrlNoKind](knowledge.md#graphconnectoritembyurlnokind) | Knowledge sources |  | object |
| [GraphConnectorItemType](knowledge.md#graphconnectoritemtype) | Knowledge sources |  | enum |
| [GraphConnectorSearchSource](knowledge.md#graphconnectorsearchsource) | Knowledge sources | ✓ | object |
| [GraphDriveItemSyncDetails](knowledge.md#graphdriveitemsyncdetails) | Knowledge sources | ✓ | object |
| [GraphDriveItemSyncDetailsNoKind](knowledge.md#graphdriveitemsyncdetailsnokind) | Knowledge sources |  | object |
| [GroundednessAssertion](evaluation.md#groundednessassertion) | Evaluation & testing | ✓ | object |
| [Guid](data-types.md#guid) | Data types |  | union |
| [HeroCardTemplate](messages.md#herocardtemplate) | Messages, activities & cards | ✓ | object |
| [HistoryType](settings.md#historytype) | Settings |  | union |
| [HoldSettings](actions.md#holdsettings) | Actions (topic nodes) |  | union |
| [HttpErrorHandlingBehavior](actions.md#httperrorhandlingbehavior) | Actions (topic nodes) |  | union |
| [HttpMethodType](actions.md#httpmethodtype) | Actions (topic nodes) |  | enum |
| [HttpRequestAction](actions.md#httprequestaction) | Actions (topic nodes) | ✓ | object |
| [HttpRequestContent](actions.md#httprequestcontent) | Actions (topic nodes) |  | union |
| [Hyperlink](data-types.md#hyperlink) | Data types |  | union |
| [ImageCardTemplate](messages.md#imagecardtemplate) | Messages, activities & cards | ✓ | object |
| [InboundTranscriptRedactionSettings](settings.md#inboundtranscriptredactionsettings) | Settings | ✓ | object |
| [InboundTranscriptRedactionSettingsNoKind](settings.md#inboundtranscriptredactionsettingsnokind) | Settings |  | object |
| [IncludeSelectedTopics](actions.md#includeselectedtopics) | Actions (topic nodes) | ✓ | object |
| [IngestionBasedGraphSearchConfiguration](knowledge.md#ingestionbasedgraphsearchconfiguration) | Knowledge sources | ✓ | object |
| [InitializablePropertyPath](other.md#initializablepropertypath) | Other definitions |  | string |
| [InitializablePropertyPathNoKind](other.md#initializablepropertypathnokind) | Other definitions |  | string |
| [InitializationBehavior](actions.md#initializationbehavior) | Actions (topic nodes) |  | enum |
| [InputHint](messages.md#inputhint) | Messages, activities & cards |  | enum |
| [Intent](dialogs.md#intent) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [IntentBase](dialogs.md#intentbase) | Dialogs (topics, tools, child agents) & recognizers |  | union |
| [IntentMatchAssertion](evaluation.md#intentmatchassertion) | Evaluation & testing | ✓ | object |
| [IntentNoKind](dialogs.md#intentnokind) | Dialogs (topics, tools, child agents) & recognizers |  | object |
| [InterruptionEagernessLevel](settings.md#interruptioneagernesslevel) | Settings |  | enum |
| [InterruptionPolicy](settings.md#interruptionpolicy) | Settings | ✓ | object |
| [InterruptionPolicyNoKind](settings.md#interruptionpolicynokind) | Settings |  | object |
| [InterruptionTopicListFilter](settings.md#interruptiontopiclistfilter) | Settings |  | union |
| [IntExpression](other.md#intexpression) | Other definitions |  | union |
| [InvalidReferenceErrorCode](other.md#invalidreferenceerrorcode) | Other definitions |  | enum |
| [InvocationSource](actions.md#invocationsource) | Actions (topic nodes) |  | enum |
| [InvokeAgentTaskAction](task-actions.md#invokeagenttaskaction) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [InvokeAIBuilderModelAction](actions.md#invokeaibuildermodelaction) | Actions (topic nodes) | ✓ | object |
| [InvokeAIBuilderModelTaskAction](task-actions.md#invokeaibuildermodeltaskaction) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [InvokeAIPluginTaskAction](task-actions.md#invokeaiplugintaskaction) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [InvokeClientTaskAction](task-actions.md#invokeclienttaskaction) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [InvokeComputerUsingAgentTaskAction](task-actions.md#invokecomputerusingagenttaskaction) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [InvokeConnectedAgentTaskAction](task-actions.md#invokeconnectedagenttaskaction) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [InvokeConnectorAction](actions.md#invokeconnectoraction) | Actions (topic nodes) | ✓ | object |
| [InvokeConnectorTaskAction](task-actions.md#invokeconnectortaskaction) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [InvokeCustomModelAction](actions.md#invokecustommodelaction) | Actions (topic nodes) | ✓ | object |
| [InvokeExternalAgentTaskAction](task-actions.md#invokeexternalagenttaskaction) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [InvokeFlowAction](actions.md#invokeflowaction) | Actions (topic nodes) | ✓ | object |
| [InvokeFlowTaskAction](task-actions.md#invokeflowtaskaction) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [InvokeMcpToolAction](actions.md#invokemcptoolaction) | Actions (topic nodes) | ✓ | object |
| [InvokerConnectionReference](task-actions.md#invokerconnectionreference) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [InvokeSkillAction](actions.md#invokeskillaction) | Actions (topic nodes) | ✓ | object |
| [InvokeSkillTaskAction](task-actions.md#invokeskilltaskaction) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [IvrBotSettings](settings.md#ivrbotsettings) | Settings | ✓ | object |
| [JsonRequestContent](actions.md#jsonrequestcontent) | Actions (topic nodes) | ✓ | object |
| [KeyBasedTermination](actions.md#keybasedtermination) | Actions (topic nodes) | ✓ | object |
| [KnowledgeSource](knowledge.md#knowledgesource) | Knowledge sources |  | union |
| [KnowledgeSourceComponent](components.md#knowledgesourcecomponent) | Components & agent definitions | ✓ | object |
| [KnowledgeSourceConfiguration](components.md#knowledgesourceconfiguration) | Components & agent definitions | ✓ | object |
| [KnowledgeSourceConfigurationNoKind](components.md#knowledgesourceconfigurationnokind) | Components & agent definitions |  | object |
| [Language](settings.md#language) | Settings |  | union |
| [LanguagePrebuiltEntity](entities.md#languageprebuiltentity) | Entities |  | union |
| [LatencyMessageSettings](messages.md#latencymessagesettings) | Messages, activities & cards | ✓ | object |
| [LatencyMessageSettingsNoKind](actions.md#latencymessagesettingsnokind) | Actions (topic nodes) |  | object |
| [LatencySettingsBase](settings.md#latencysettingsbase) | Settings |  | union |
| [LatencySettingsBaseLegacy](settings.md#latencysettingsbaselegacy) | Settings |  | union |
| [LegacyOrUnknownComponent](components.md#legacyorunknowncomponent) | Components & agent definitions | ✓ | object |
| [LocalizableContentContainer](dialogs.md#localizablecontentcontainer) | Dialogs (topics, tools, child agents) & recognizers |  | union |
| [LogCustomTelemetryEvent](actions.md#logcustomtelemetryevent) | Actions (topic nodes) | ✓ | object |
| [MakeCall](messages.md#makecall) | Messages, activities & cards | ✓ | object |
| [MakerConnectionReference](task-actions.md#makerconnectionreference) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [ManagedProperties](components.md#managedproperties) | Components & agent definitions | ✓ | object |
| [ManagedPropertiesNoKind](components.md#managedpropertiesnokind) | Components & agent definitions |  | object |
| [ManifestResources](components.md#manifestresources) | Components & agent definitions |  | union |
| [ManualTaskInput](task-actions.md#manualtaskinput) | Tool (TaskDialog) actions & inputs | ✓ | object |
| [McpResourceFileReference](messages.md#mcpresourcefilereference) | Messages, activities & cards | ✓ | object |
| [McpToolDefinition](components.md#mcptooldefinition) | Components & agent definitions | ✓ | object |
| [McpToolDefinitionNoKind](components.md#mcptooldefinitionnokind) | Components & agent definitions |  | object |
| [MeetingSearchSource](knowledge.md#meetingsearchsource) | Knowledge sources | ✓ | object |
| [Message](messages.md#message) | Messages, activities & cards |  | union |
| [MessageBack](messages.md#messageback) | Messages, activities & cards | ✓ | object |
| [MessageNoKind](messages.md#messagenokind) | Messages, activities & cards |  | union |
| [MessageOptionalKind](messages.md#messageoptionalkind) | Messages, activities & cards |  | union |
| [Microsoft365ApplicationDetails](components.md#microsoft365applicationdetails) | Components & agent definitions | ✓ | object |
| [Microsoft365ApplicationDetailsNoKind](components.md#microsoft365applicationdetailsnokind) | Components & agent definitions |  | object |
| [ModelApiType](settings.md#modelapitype) | Settings |  | enum |
| [ModelContextProtocolMetadata](components.md#modelcontextprotocolmetadata) | Components & agent definitions | ✓ | object |
| [ModelManifestResource](settings.md#modelmanifestresource) | Settings | ✓ | object |
| [ModelProvider](settings.md#modelprovider) | Settings |  | enum |
| [MoneyPrebuiltEntity](entities.md#moneyprebuiltentity) | Entities |  | union |
| [MultiEntityChoice](entities.md#multientitychoice) | Entities | ✓ | object |
| [MultiEntityChoiceNoKind](entities.md#multientitychoicenokind) | Entities |  | object |
| [MultipleTopicsMatchedReason](actions.md#multipletopicsmatchedreason) | Actions (topic nodes) |  | enum |
| [MultiTurnActivity](evaluation.md#multiturnactivity) | Evaluation & testing | ✓ | object |
| [MultiTurnActivityNoKind](evaluation.md#multiturnactivitynokind) | Evaluation & testing |  | object |
| [MultiTurnEvaluationCase](evaluation.md#multiturnevaluationcase) | Evaluation & testing | ✓ | object |
| [NetworkIsolationType](settings.md#networkisolationtype) | Settings |  | enum |
| [NetworkSettings](settings.md#networksettings) | Settings | ✓ | object |
| [NetworkSettingsNoKind](settings.md#networksettingsnokind) | Settings |  | object |
| [NeverRequireMCPApproval](settings.md#neverrequiremcpapproval) | Settings | ✓ | object |
| [NoCitationsInModelOutput](knowledge.md#nocitationsinmodeloutput) | Knowledge sources | ✓ | object |
| [NoFeedbackCollection](settings.md#nofeedbackcollection) | Settings | ✓ | object |
| [NoHistory](settings.md#nohistory) | Settings | ✓ | object |
| [NoRequestContent](actions.md#norequestcontent) | Actions (topic nodes) | ✓ | object |
| [NoRequestContentNoKind](actions.md#norequestcontentnokind) | Actions (topic nodes) |  | object |
| [Number](data-types.md#number) | Data types |  | union |
| [NumberExpression](other.md#numberexpression) | Other definitions |  | union |
| [NumberPrebuiltEntity](entities.md#numberprebuiltentity) | Entities |  | union |
| [OAuthInput](actions.md#oauthinput) | Actions (topic nodes) | ✓ | object |
| [OnActivity](triggers.md#onactivity) | Triggers | ✓ | object |
| [OnAnalyzeText](triggers.md#onanalyzetext) | Triggers | ✓ | object |
| [OnAnsweringMachineDetected](triggers.md#onansweringmachinedetected) | Triggers | ✓ | object |
| [OnConversationStart](triggers.md#onconversationstart) | Triggers | ✓ | object |
| [OnCopilotRedirect](triggers.md#oncopilotredirect) | Triggers | ✓ | object |
| [OnDtmfKeyPress](triggers.md#ondtmfkeypress) | Triggers | ✓ | object |
| [OneOfMultiEntityReference](entities.md#oneofmultientityreference) | Entities |  | union |
| [OnError](triggers.md#onerror) | Triggers | ✓ | object |
| [OnEscalate](triggers.md#onescalate) | Triggers | ✓ | object |
| [OnEventActivity](triggers.md#oneventactivity) | Triggers | ✓ | object |
| [OnGeneratedResponse](triggers.md#ongeneratedresponse) | Triggers | ✓ | object |
| [OnInactivity](triggers.md#oninactivity) | Triggers | ✓ | object |
| [OnKnowledgeRequested](triggers.md#onknowledgerequested) | Triggers | ✓ | object |
| [OnOutgoingMessage](triggers.md#onoutgoingmessage) | Triggers | ✓ | object |
| [OnPlanComplete](triggers.md#onplancomplete) | Triggers | ✓ | object |
| [OnRecognizedIntent](triggers.md#onrecognizedintent) | Triggers | ✓ | object |
| [OnRedirect](triggers.md#onredirect) | Triggers | ✓ | object |
| [OnSelectIntent](triggers.md#onselectintent) | Triggers | ✓ | object |
| [OnSelectIntentTriggerBehavior](triggers.md#onselectintenttriggerbehavior) | Triggers |  | enum |
| [OnSignIn](triggers.md#onsignin) | Triggers | ✓ | object |
| [OnSilence](triggers.md#onsilence) | Triggers | ✓ | object |
| [OnSystemIntent](triggers.md#onsystemintent) | Triggers | ✓ | object |
| [OnSystemRedirect](triggers.md#onsystemredirect) | Triggers | ✓ | object |
| [OnToolSelected](triggers.md#ontoolselected) | Triggers | ✓ | object |
| [OnUnknownDtmfKey](triggers.md#onunknowndtmfkey) | Triggers | ✓ | object |
| [OnUnknownIntent](triggers.md#onunknownintent) | Triggers | ✓ | object |
| [OnUnrecognizedSpeech](triggers.md#onunrecognizedspeech) | Triggers | ✓ | object |
| [OpenUrl](messages.md#openurl) | Messages, activities & cards | ✓ | object |
| [OptionSetDataType](data-types.md#optionsetdatatype) | Data types |  | union |
| [OrConditionGroup](actions.md#orconditiongroup) | Actions (topic nodes) | ✓ | object |
| [OrdinalPrebuiltEntity](entities.md#ordinalprebuiltentity) | Entities |  | union |
| [OrganizationPrebuiltEntity](entities.md#organizationprebuiltentity) | Entities |  | union |
| [ParseValue](actions.md#parsevalue) | Actions (topic nodes) | ✓ | object |
| [ParseValueErrorPolicy](actions.md#parsevalueerrorpolicy) | Actions (topic nodes) |  | enum |
| [PartType](actions.md#parttype) | Actions (topic nodes) |  | enum |
| [PercentagePrebuiltEntity](entities.md#percentageprebuiltentity) | Entities |  | union |
| [PersonNamePrebuiltEntity](entities.md#personnameprebuiltentity) | Entities |  | union |
| [PhoneNumberPrebuiltEntity](entities.md#phonenumberprebuiltentity) | Entities |  | union |
| [PlanStepType](components.md#plansteptype) | Components & agent definitions |  | enum |
| [PointOfInterestPrebuiltEntity](entities.md#pointofinterestprebuiltentity) | Entities |  | union |
| [PreviewModels](settings.md#previewmodels) | Settings | ✓ | object |
| [PromptEvaluationTarget](evaluation.md#promptevaluationtarget) | Evaluation & testing | ✓ | object |
| [PromptGrader](evaluation.md#promptgrader) | Evaluation & testing | ✓ | object |
| [PromptGraderLabel](evaluation.md#promptgraderlabel) | Evaluation & testing | ✓ | object |
| [PromptGraderLabelNoKind](evaluation.md#promptgraderlabelnokind) | Evaluation & testing |  | object |
| [PropertyInfo](data-types.md#propertyinfo) | Data types |  | union |
| [PropertyInfoNoKind](data-types.md#propertyinfonokind) | Data types |  | union |
| [PropertyOverride](components.md#propertyoverride) | Components & agent definitions | ✓ | object |
| [PropertyOverrideNoKind](actions.md#propertyoverridenokind) | Actions (topic nodes) |  | object |
| [PropertyVisibility](data-types.md#propertyvisibility) | Data types |  | enum |
| [ProtocolMetadataBase](components.md#protocolmetadatabase) | Components & agent definitions |  | union |
| [ProvisioningState](components.md#provisioningstate) | Components & agent definitions |  | enum |
| [ProvisioningStatus](components.md#provisioningstatus) | Components & agent definitions |  | union |
| [PublicDataSource](knowledge.md#publicdatasource) | Knowledge sources |  | union |
| [PublicSiteSearchDataSource](knowledge.md#publicsitesearchdatasource) | Knowledge sources | ✓ | object |
| [PublicSiteSearchDataSourceNoKind](knowledge.md#publicsitesearchdatasourcenokind) | Knowledge sources |  | object |
| [PublicSiteSearchSource](knowledge.md#publicsitesearchsource) | Knowledge sources | ✓ | object |
| [PublishResult](components.md#publishresult) | Components & agent definitions |  | enum |
| [PublishResultDetails](components.md#publishresultdetails) | Components & agent definitions | ✓ | object |
| [PublishResultDetailsNoKind](components.md#publishresultdetailsnokind) | Components & agent definitions |  | object |
| [Question](actions.md#question) | Actions (topic nodes) | ✓ | object |
| [QuestionInputSettings](settings.md#questioninputsettings) | Settings | ✓ | object |
| [QuestionInputSettingsNoKind](task-actions.md#questioninputsettingsnokind) | Tool (TaskDialog) actions & inputs |  | object |
| [QuickReplyBase](messages.md#quickreplybase) | Messages, activities & cards |  | union |
| [RaiseErrorBehavior](actions.md#raiseerrorbehavior) | Actions (topic nodes) | ✓ | object |
| [RaiseErrorBehaviorNoKind](actions.md#raiseerrorbehaviornokind) | Actions (topic nodes) |  | object |
| [RawRequestContent](actions.md#rawrequestcontent) | Actions (topic nodes) | ✓ | object |
| [RealtimeVoiceAudioLatencySettings](settings.md#realtimevoiceaudiolatencysettings) | Settings | ✓ | object |
| [RealTimeVoiceBotSettings](settings.md#realtimevoicebotsettings) | Settings | ✓ | object |
| [RealtimeVoiceMessageLatencySettings](settings.md#realtimevoicemessagelatencysettings) | Settings | ✓ | object |
| [RealtimeVoiceMessageLatencySettingsNoKind](settings.md#realtimevoicemessagelatencysettingsnokind) | Settings |  | object |
| [RealtimeVoiceProcessingMode](settings.md#realtimevoiceprocessingmode) | Settings | ✓ | object |
| [ReasoningExperimentalModels](settings.md#reasoningexperimentalmodels) | Settings | ✓ | object |
| [ReasoningPreviewModels](settings.md#reasoningpreviewmodels) | Settings | ✓ | object |
| [ReasoningProductionModels](settings.md#reasoningproductionmodels) | Settings | ✓ | object |
| [RecognizeIntent](actions.md#recognizeintent) | Actions (topic nodes) | ✓ | object |
| [Recognizer](dialogs.md#recognizer) | Dialogs (topics, tools, child agents) & recognizers |  | union |
| [RecognizerBase](dialogs.md#recognizerbase) | Dialogs (topics, tools, child agents) & recognizers |  | union |
| [Record](data-types.md#record) | Data types | ✓ | object |
| [RecordNoKind](data-types.md#recordnokind) | Data types |  | object |
| [RecurrenceFrequency](components.md#recurrencefrequency) | Components & agent definitions |  | enum |
| [RecurrenceTrigger](triggers.md#recurrencetrigger) | Triggers | ✓ | object |
| [RecurrenceTriggerDefinition](components.md#recurrencetriggerdefinition) | Components & agent definitions | ✓ | object |
| [ReferenceItemSourceFile](components.md#referenceitemsourcefile) | Components & agent definitions | ✓ | object |
| [ReferenceItemSourceFileNoKind](components.md#referenceitemsourcefilenokind) | Components & agent definitions |  | object |
| [ReferencesSourceFile](components.md#referencessourcefile) | Components & agent definitions | ✓ | object |
| [ReferenceType](components.md#referencetype) | Components & agent definitions |  | enum |
| [RegexEntity](dialogs.md#regexentity) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [RegexEntityReference](entities.md#regexentityreference) | Entities |  | union |
| [RegexIntentPattern](entities.md#regexintentpattern) | Entities | ✓ | object |
| [RegexIntentPatternNoKind](dialogs.md#regexintentpatternnokind) | Dialogs (topics, tools, child agents) & recognizers |  | object |
| [RegexRecognizer](dialogs.md#regexrecognizer) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [RelevanceAssertion](evaluation.md#relevanceassertion) | Evaluation & testing | ✓ | object |
| [RemoveItemOperation](actions.md#removeitemoperation) | Actions (topic nodes) | ✓ | object |
| [RepeatDialog](actions.md#repeatdialog) | Actions (topic nodes) | ✓ | object |
| [ReplaceDialog](actions.md#replacedialog) | Actions (topic nodes) | ✓ | object |
| [ResetVariable](actions.md#resetvariable) | Actions (topic nodes) | ✓ | object |
| [ResponseTemplate](messages.md#responsetemplate) | Messages, activities & cards | ✓ | object |
| [ResponseTemplateNoKind](messages.md#responsetemplatenokind) | Messages, activities & cards |  | object |
| [RoiCalculationSettingsBase](settings.md#roicalculationsettingsbase) | Settings |  | union |
| [RuntimeProvider](components.md#runtimeprovider) | Components & agent definitions |  | enum |
| [SampleDataType](evaluation.md#sampledatatype) | Evaluation & testing |  | enum |
| [SampleSegment](evaluation.md#samplesegment) | Evaluation & testing | ✓ | object |
| [SampleSegmentNoKind](evaluation.md#samplesegmentnokind) | Evaluation & testing |  | object |
| [SearchAllFiles](actions.md#searchallfiles) | Actions (topic nodes) | ✓ | object |
| [SearchAllKnowledgeSources](knowledge.md#searchallknowledgesources) | Knowledge sources | ✓ | object |
| [SearchAndSummarizeContent](actions.md#searchandsummarizecontent) | Actions (topic nodes) | ✓ | object |
| [SearchAndSummarizeWithCustomModel](actions.md#searchandsummarizewithcustommodel) | Actions (topic nodes) | ✓ | object |
| [SearchFilesMode](actions.md#searchfilesmode) | Actions (topic nodes) |  | union |
| [SearchKnowledgeSources](actions.md#searchknowledgesources) | Actions (topic nodes) | ✓ | object |
| [SearchKnowledgeSourcesMode](knowledge.md#searchknowledgesourcesmode) | Knowledge sources |  | union |
| [SearchResultRecord](knowledge.md#searchresultrecord) | Knowledge sources | ✓ | object |
| [SearchSpecificFiles](actions.md#searchspecificfiles) | Actions (topic nodes) | ✓ | object |
| [SearchSpecificKnowledgeSources](knowledge.md#searchspecificknowledgesources) | Knowledge sources | ✓ | object |
| [SearchSummarizationConfigurationState](knowledge.md#searchsummarizationconfigurationstate) | Knowledge sources |  | enum |
| [SelectedClosedListItem](entities.md#selectedclosedlistitem) | Entities |  | union |
| [SelectedClosedListItemNoKind](entities.md#selectedclosedlistitemnokind) | Entities |  | union |
| [SelectedClosedListItemOptionalKind](entities.md#selectedclosedlistitemoptionalkind) | Entities |  | union |
| [SendActivity](actions.md#sendactivity) | Actions (topic nodes) | ✓ | object |
| [SendHiddenMessageBack](messages.md#sendhiddenmessageback) | Messages, activities & cards | ✓ | object |
| [SendUserActivity](messages.md#senduseractivity) | Messages, activities & cards | ✓ | object |
| [SensitiveDataMaskingSettings](settings.md#sensitivedatamaskingsettings) | Settings |  | union |
| [SensitivityLabelStatusCode](settings.md#sensitivitylabelstatuscode) | Settings |  | enum |
| [SensitivityLevel](settings.md#sensitivitylevel) | Settings |  | enum |
| [SessionBasedRoiCalculationSettings](settings.md#sessionbasedroicalculationsettings) | Settings | ✓ | object |
| [SetMultipleVariables](actions.md#setmultiplevariables) | Actions (topic nodes) | ✓ | object |
| [SetTextVariable](actions.md#settextvariable) | Actions (topic nodes) | ✓ | object |
| [SetVariable](actions.md#setvariable) | Actions (topic nodes) | ✓ | object |
| [SharepointIds](knowledge.md#sharepointids) | Knowledge sources | ✓ | object |
| [SharepointIdsNoKind](knowledge.md#sharepointidsnokind) | Knowledge sources |  | object |
| [SharePointIndexBehavior](knowledge.md#sharepointindexbehavior) | Knowledge sources |  | enum |
| [SharePointSearchDataSource](knowledge.md#sharepointsearchdatasource) | Knowledge sources | ✓ | object |
| [SharePointSearchDataSourceNoKind](actions.md#sharepointsearchdatasourcenokind) | Actions (topic nodes) |  | object |
| [SharePointSearchSource](knowledge.md#sharepointsearchsource) | Knowledge sources | ✓ | object |
| [SharepointSourceType](knowledge.md#sharepointsourcetype) | Knowledge sources |  | enum |
| [SignInReason](actions.md#signinreason) | Actions (topic nodes) |  | enum |
| [SignOutUser](actions.md#signoutuser) | Actions (topic nodes) | ✓ | object |
| [SimpleEvaluationCase](evaluation.md#simpleevaluationcase) | Evaluation & testing | ✓ | object |
| [SimpleEvaluationCaseNoKind](evaluation.md#simpleevaluationcasenokind) | Evaluation & testing |  | object |
| [SingleValueInputBinding](actions.md#singlevalueinputbinding) | Actions (topic nodes) | ✓ | object |
| [SingleVariableOutputBinding](actions.md#singlevariableoutputbinding) | Actions (topic nodes) | ✓ | object |
| [SkillActionDefinition](components.md#skillactiondefinition) | Components & agent definitions | ✓ | object |
| [SkillActionDefinitionNoKind](components.md#skillactiondefinitionnokind) | Components & agent definitions |  | object |
| [SkillComponent](components.md#skillcomponent) | Components & agent definitions | ✓ | object |
| [SkillDefinition](components.md#skilldefinition) | Components & agent definitions | ✓ | object |
| [SkillDefinitionNoKind](components.md#skilldefinitionnokind) | Components & agent definitions |  | object |
| [SkipQuestionMode](actions.md#skipquestionmode) | Actions (topic nodes) |  | enum |
| [SolutionComponent](components.md#solutioncomponent) | Components & agent definitions |  | union |
| [SpeechRecognitionFormatType](entities.md#speechrecognitionformattype) | Entities |  | enum |
| [SpeedPrebuiltEntity](entities.md#speedprebuiltentity) | Entities |  | union |
| [SqlFederatedTableSearchConfiguration](knowledge.md#sqlfederatedtablesearchconfiguration) | Knowledge sources | ✓ | object |
| [StateCode](components.md#statecode) | Components & agent definitions |  | enum |
| [StatePrebuiltEntity](entities.md#stateprebuiltentity) | Entities |  | union |
| [StreetAddressPrebuiltEntity](entities.md#streetaddressprebuiltentity) | Entities |  | union |
| [String](data-types.md#string) | Data types |  | union |
| [StringExpression](other.md#stringexpression) | Other definitions |  | string |
| [StringPrebuiltEntity](entities.md#stringprebuiltentity) | Entities |  | union |
| [StructuredRecordExpression](other.md#structuredrecordexpression) | Other definitions | ✓ | object |
| [StructuredRecordExpressionNoKind](other.md#structuredrecordexpressionnokind) | Other definitions |  | object |
| [SuccessfulPublishResultDetails](components.md#successfulpublishresultdetails) | Components & agent definitions | ✓ | object |
| [SuccessfulPublishResultDetailsNoKind](components.md#successfulpublishresultdetailsnokind) | Components & agent definitions |  | object |
| [SummarizationExternalModelConfiguration](knowledge.md#summarizationexternalmodelconfiguration) | Knowledge sources | ✓ | object |
| [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) | Knowledge sources |  | object |
| [SynchronizationState](components.md#synchronizationstate) | Components & agent definitions | ✓ | object |
| [SynchronizationStateNoKind](components.md#synchronizationstatenokind) | Components & agent definitions |  | object |
| [SynchronizedChannel](components.md#synchronizedchannel) | Components & agent definitions | ✓ | object |
| [SynchronizedChannelNoKind](components.md#synchronizedchannelnokind) | Components & agent definitions |  | object |
| [SynchronizingState](components.md#synchronizingstate) | Components & agent definitions |  | enum |
| [SystemOptionSet](data-types.md#systemoptionset) | Data types | ✓ | object |
| [SystemOptionSetName](data-types.md#systemoptionsetname) | Data types |  | enum |
| [SystemUser](components.md#systemuser) | Components & agent definitions | ✓ | object |
| [Table](data-types.md#table) | Data types | ✓ | object |
| [TableChangeType](actions.md#tablechangetype) | Actions (topic nodes) |  | enum |
| [TableSearchType](components.md#tablesearchtype) | Components & agent definitions |  | enum |
| [TakeFirstItemOperation](actions.md#takefirstitemoperation) | Actions (topic nodes) | ✓ | object |
| [TakeLastItemOperation](actions.md#takelastitemoperation) | Actions (topic nodes) | ✓ | object |
| [TaskAction](task-actions.md#taskaction) | Tool (TaskDialog) actions & inputs |  | union |
| [TaskActionOutputMode](dialogs.md#taskactionoutputmode) | Dialogs (topics, tools, child agents) & recognizers |  | enum |
| [TaskDialog](dialogs.md#taskdialog) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [TaskInput](task-actions.md#taskinput) | Tool (TaskDialog) actions & inputs |  | union |
| [TaskOutput](actions.md#taskoutput) | Actions (topic nodes) | ✓ | object |
| [TaskOutputNoKind](dialogs.md#taskoutputnokind) | Dialogs (topics, tools, child agents) & recognizers |  | object |
| [TaskSuccessGrader](evaluation.md#tasksuccessgrader) | Evaluation & testing | ✓ | object |
| [TeamsApplicationDetails](components.md#teamsapplicationdetails) | Components & agent definitions | ✓ | object |
| [TeamsApplicationDetailsNoKind](components.md#teamsapplicationdetailsnokind) | Components & agent definitions |  | object |
| [TeamsMessageSearchSource](knowledge.md#teamsmessagesearchsource) | Knowledge sources | ✓ | object |
| [TemperaturePrebuiltEntity](entities.md#temperatureprebuiltentity) | Entities |  | union |
| [TemplateLine](messages.md#templateline) | Messages, activities & cards |  | string |
| [TemplateLineNoKind](messages.md#templatelinenokind) | Messages, activities & cards |  | string |
| [TemplateResponseMode](messages.md#templateresponsemode) | Messages, activities & cards |  | enum |
| [TemplateSegment](messages.md#templatesegment) | Messages, activities & cards |  | union |
| [TenantFineTunedModelDetails](settings.md#tenantfinetunedmodeldetails) | Settings | ✓ | object |
| [TenantFineTunedModelDetailsNoKind](settings.md#tenantfinetunedmodeldetailsnokind) | Settings |  | object |
| [TestActivityDefinition](evaluation.md#testactivitydefinition) | Evaluation & testing |  | union |
| [TestActivitySourceType](evaluation.md#testactivitysourcetype) | Evaluation & testing |  | enum |
| [TestAssertionBase](evaluation.md#testassertionbase) | Evaluation & testing |  | union |
| [TestCaseComponent](components.md#testcasecomponent) | Components & agent definitions | ✓ | object |
| [TestCaseDefinition](evaluation.md#testcasedefinition) | Evaluation & testing | ✓ | object |
| [TestDefinitionBase](evaluation.md#testdefinitionbase) | Evaluation & testing |  | union |
| [TestSetDefinition](evaluation.md#testsetdefinition) | Evaluation & testing | ✓ | object |
| [TestTranscriptDefinition](evaluation.md#testtranscriptdefinition) | Evaluation & testing | ✓ | object |
| [TestTranscriptDefinitionNoKind](evaluation.md#testtranscriptdefinitionnokind) | Evaluation & testing |  | object |
| [TextSegment](messages.md#textsegment) | Messages, activities & cards | ✓ | object |
| [TextSimilarityGrader](evaluation.md#textsimilaritygrader) | Evaluation & testing | ✓ | object |
| [Time](data-types.md#time) | Data types |  | union |
| [ToneType](settings.md#tonetype) | Settings |  | enum |
| [ToolAugmentation](task-actions.md#toolaugmentation) | Tool (TaskDialog) actions & inputs |  | union |
| [ToolBasedRoiCalculationSettings](settings.md#toolbasedroicalculationsettings) | Settings | ✓ | object |
| [ToolCallGrader](evaluation.md#toolcallgrader) | Evaluation & testing | ✓ | object |
| [ToolDialogVariableContext](actions.md#tooldialogvariablecontext) | Actions (topic nodes) |  | enum |
| [ToolManifestResource](components.md#toolmanifestresource) | Components & agent definitions | ✓ | object |
| [ToolRoiCalculationConfiguration](settings.md#toolroicalculationconfiguration) | Settings | ✓ | object |
| [ToolRoiCalculationConfigurationNoKind](settings.md#toolroicalculationconfigurationnokind) | Settings |  | object |
| [ToolsSelectionMode](components.md#toolsselectionmode) | Components & agent definitions |  | union |
| [TopicStartBehavior](dialogs.md#topicstartbehavior) | Dialogs (topics, tools, child agents) & recognizers |  | enum |
| [TransferContext](actions.md#transfercontext) | Actions (topic nodes) |  | union |
| [TransferConversation](actions.md#transferconversation) | Actions (topic nodes) | ✓ | object |
| [TransferConversationV2](actions.md#transferconversationv2) | Actions (topic nodes) | ✓ | object |
| [TransferToAgent](actions.md#transfertoagent) | Actions (topic nodes) | ✓ | object |
| [TransferToPhoneNumber](actions.md#transfertophonenumber) | Actions (topic nodes) | ✓ | object |
| [TransferType](actions.md#transfertype) | Actions (topic nodes) |  | union |
| [TranslationsComponent](components.md#translationscomponent) | Components & agent definitions | ✓ | object |
| [TranslationsComponentNoKind](components.md#translationscomponentnokind) | Components & agent definitions |  | object |
| [TriggerBase](triggers.md#triggerbase) | Triggers |  | union |
| [TurboExperimentalModels](settings.md#turboexperimentalmodels) | Settings | ✓ | object |
| [TurboPreviewModels](settings.md#turbopreviewmodels) | Settings | ✓ | object |
| [TurboProductionModels](settings.md#turboproductionmodels) | Settings | ✓ | object |
| [TurnDetectionType](settings.md#turndetectiontype) | Settings |  | enum |
| [UnresolvedDynamicType](data-types.md#unresolveddynamictype) | Data types |  | union |
| [Unspecified](data-types.md#unspecified) | Data types |  | union |
| [UpdateActivity](actions.md#updateactivity) | Actions (topic nodes) | ✓ | object |
| [UriFileReference](messages.md#urifilereference) | Messages, activities & cards | ✓ | object |
| [URLPrebuiltEntity](entities.md#urlprebuiltentity) | Entities |  | union |
| [UseAllTools](components.md#usealltools) | Components & agent definitions | ✓ | object |
| [UserInputType](actions.md#userinputtype) | Actions (topic nodes) |  | enum |
| [UserPrivilege](components.md#userprivilege) | Components & agent definitions | ✓ | object |
| [UserRequestedHoldSettings](actions.md#userrequestedholdsettings) | Actions (topic nodes) | ✓ | object |
| [UseSpecificTools](components.md#usespecifictools) | Components & agent definitions | ✓ | object |
| [ValidationErrorCode](other.md#validationerrorcode) | Other definitions |  | enum |
| [ValueExpression](other.md#valueexpression) | Other definitions |  | union |
| [Variable](actions.md#variable) | Actions (topic nodes) | ✓ | object |
| [VariableAIVisibility](actions.md#variableaivisibility) | Actions (topic nodes) |  | enum |
| [VariableAssignment](actions.md#variableassignment) | Actions (topic nodes) | ✓ | object |
| [VariableAssignmentNoKind](actions.md#variableassignmentnokind) | Actions (topic nodes) |  | object |
| [VariableNoKind](components.md#variablenokind) | Components & agent definitions |  | object |
| [VariableScope](actions.md#variablescope) | Actions (topic nodes) |  | enum |
| [VariablesToClear](actions.md#variablestoclear) | Actions (topic nodes) |  | enum |
| [VideoCardTemplate](messages.md#videocardtemplate) | Messages, activities & cards | ✓ | object |
| [VirtualAgents.Recognizer](dialogs.md#virtualagentsrecognizer) | Dialogs (topics, tools, child agents) & recognizers | ✓ | object |
| [VoiceAudioLatencySettings](settings.md#voiceaudiolatencysettings) | Settings | ✓ | object |
| [VoiceInputSettings](settings.md#voiceinputsettings) | Settings | ✓ | object |
| [VoiceInputSettingsNoKind](actions.md#voiceinputsettingsnokind) | Actions (topic nodes) |  | object |
| [VoiceMessageLatencySettings](settings.md#voicemessagelatencysettings) | Settings | ✓ | object |
| [VoiceMessageLatencySettingsNoKind](settings.md#voicemessagelatencysettingsnokind) | Settings |  | object |
| [VoiceProcessingModeBase](settings.md#voiceprocessingmodebase) | Settings |  | union |
| [VoiceTypeBase](settings.md#voicetypebase) | Settings | ✓ | object |
| [VoiceTypeBaseNoKind](settings.md#voicetypebasenokind) | Settings |  | object |
| [WaitForConnectorTrigger](actions.md#waitforconnectortrigger) | Actions (topic nodes) | ✓ | object |
| [WasmBrowserTelemetryLogLevel](settings.md#wasmbrowsertelemetryloglevel) | Settings |  | enum |
| [WasmInvalidStateMessage](messages.md#wasminvalidstatemessage) | Messages, activities & cards |  | enum |
| [WeightPrebuiltEntity](entities.md#weightprebuiltentity) | Entities |  | union |
| [WorkflowExternalTrigger](components.md#workflowexternaltrigger) | Components & agent definitions | ✓ | object |
| [ZipCodePrebuiltEntity](entities.md#zipcodeprebuiltentity) | Entities |  | union |
