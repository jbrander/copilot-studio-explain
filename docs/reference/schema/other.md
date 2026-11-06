# Other definitions

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

Supporting types that don't belong to a single category: expression wrappers, enums, metadata, and helper records.

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

**18 definitions** on this page (1 with a `kind:`). [A–Z index of all definitions](index.md).

<details><summary>Contents</summary>

[AdaptiveCardExpression](#adaptivecardexpression), [BoolExpression](#boolexpression), [BotElement](#botelement), [DataValue](#datavalue), [DefinitionBase](#definitionbase), [DialogExpression](#dialogexpression), [ExpressionErrorCode](#expressionerrorcode), [FileExpression](#fileexpression), [InitializablePropertyPath](#initializablepropertypath), [InitializablePropertyPathNoKind](#initializablepropertypathnokind), [IntExpression](#intexpression), [InvalidReferenceErrorCode](#invalidreferenceerrorcode), [NumberExpression](#numberexpression), [StringExpression](#stringexpression), [StructuredRecordExpression](#structuredrecordexpression), [StructuredRecordExpressionNoKind](#structuredrecordexpressionnokind), [ValidationErrorCode](#validationerrorcode), [ValueExpression](#valueexpression)

</details>

### AdaptiveCardExpression

*Title:* LiteralValue  

**Type:** `string`

<sub>Used by: [AdaptiveCardPrompt](actions.md#adaptivecardprompt), [AdaptiveCardTemplate](messages.md#adaptivecardtemplate)</sub>

### BoolExpression


**One of:** `boolean`, `=` expression

<sub>Used by: [ActivateExternalTrigger](actions.md#activateexternaltrigger), [AdaptiveDialog](dialogs.md#adaptivedialog), [AdaptiveDialogNoKind](dialogs.md#adaptivedialognokind), [AzureAISearchSource](knowledge.md#azureaisearchsource), [AzureCognitiveSearchInternalDataSource](knowledge.md#azurecognitivesearchinternaldatasource), [AzureCognitiveSearchInternalDataSourceNoKind](knowledge.md#azurecognitivesearchinternaldatasourcenokind), [AzureOpenAIOnYourDataKnowledgeSource](knowledge.md#azureopenaionyourdataknowledgesource), [BeginDialog](actions.md#begindialog), [BingCustomSearchSource](knowledge.md#bingcustomsearchsource), [CancelAllDialogs](actions.md#cancelalldialogs), [CancelDialog](actions.md#canceldialog), [ConditionItem](actions.md#conditionitem), [ConditionItemNoKind](actions.md#conditionitemnokind), [ContinueConversationConnectorTrigger](triggers.md#continueconversationconnectortrigger), [CopilotConnectorKnowledgeSourceDemo](knowledge.md#copilotconnectorknowledgesourcedemo), [CreateConversationConnectorTrigger](triggers.md#createconversationconnectortrigger), [CustomKnowledgeSource](knowledge.md#customknowledgesource), [DataverseStructuredSearchSource](knowledge.md#dataversestructuredsearchsource), [DtmfMultipleChoiceOptions](settings.md#dtmfmultiplechoiceoptions), [DtmfMultipleChoiceOptionsNoKind](settings.md#dtmfmultiplechoiceoptionsnokind) … and 49 more</sub>

### BotElement


**One of 457 types** (see the [A–Z index](index.md)).

### DataValue


**One of:** `string`, `number`, `boolean`, list of any, object

<sub>Used by: [BotConfiguration](components.md#botconfiguration), [BotConfigurationNoKind](components.md#botconfigurationnokind), [PropertyInfo](data-types.md#propertyinfo), [PropertyInfoNoKind](data-types.md#propertyinfonokind), [SimpleEvaluationCase](evaluation.md#simpleevaluationcase), [SimpleEvaluationCaseNoKind](evaluation.md#simpleevaluationcasenokind)</sub>

### DefinitionBase


**One of:** [BotDefinition](components.md#botdefinition), [BotComponentCollectionDefinition](components.md#botcomponentcollectiondefinition)

### DialogExpression


**One of:** `string`, `=` expression

<sub>Used by: [ActivateExternalTrigger](actions.md#activateexternaltrigger), [BeginDialog](actions.md#begindialog), [DisableTrigger](actions.md#disabletrigger), [Question](actions.md#question), [ReplaceDialog](actions.md#replacedialog), [VoiceInputSettings](settings.md#voiceinputsettings), [VoiceInputSettingsNoKind](actions.md#voiceinputsettingsnokind)</sub>

### ExpressionErrorCode


**String enum.** Allowed values: `UnexpectedCharacter`, `IdentifierNotRecognized`, `IncompatibleTypes`, `PowerFxError`

### FileExpression


**Type:** `=` expression

### InitializablePropertyPath


**Type:** `string`

### InitializablePropertyPathNoKind


**Type:** `string`

### IntExpression


**One of:** `number`, `=` expression

<sub>Used by: [AdaptiveCardPrompt](actions.md#adaptivecardprompt), [AutomaticHoldSettings](actions.md#automaticholdsettings), [AzureCognitiveSearchInternalDataSource](knowledge.md#azurecognitivesearchinternaldatasource), [AzureCognitiveSearchInternalDataSourceNoKind](knowledge.md#azurecognitivesearchinternaldatasourcenokind), [AzureOpenAIOnYourDataKnowledgeSource](knowledge.md#azureopenaionyourdataknowledgesource), [AzureOpenAIOnYourDataSource](knowledge.md#azureopenaionyourdatasource), [AzureOpenAIOnYourDataSourceNoKind](actions.md#azureopenaionyourdatasourcenokind), [ConversationHistory](settings.md#conversationhistory), [ExternalChatCompletionModelConfiguration](actions.md#externalchatcompletionmodelconfiguration), [ExternalChatCompletionModelConfigurationNoKind](actions.md#externalchatcompletionmodelconfigurationnokind), [FixedLengthTermination](actions.md#fixedlengthtermination), [FlexibleLengthTermination](actions.md#flexiblelengthtermination), [HttpRequestAction](actions.md#httprequestaction), [InvokeConnectorAction](actions.md#invokeconnectoraction), [InvokeConnectorTaskAction](task-actions.md#invokeconnectortaskaction), [InvokeCustomModelAction](actions.md#invokecustommodelaction), [KeyBasedTermination](actions.md#keybasedtermination), [LatencyMessageSettings](messages.md#latencymessagesettings), [LatencyMessageSettingsNoKind](actions.md#latencymessagesettingsnokind), [OnSilence](triggers.md#onsilence) … and 9 more</sub>

### InvalidReferenceErrorCode


**String enum.** Allowed values: `NotFound`, `Disabled`, `InvalidKind`

### NumberExpression


**One of:** `number`, `=` expression

<sub>Used by: [AzureOpenAIOnYourDataKnowledgeSource](knowledge.md#azureopenaionyourdataknowledgesource), [AzureOpenAIOnYourDataSource](knowledge.md#azureopenaionyourdatasource), [AzureOpenAIOnYourDataSourceNoKind](actions.md#azureopenaionyourdatasourcenokind), [ExternalChatCompletionModelConfiguration](actions.md#externalchatcompletionmodelconfiguration), [ExternalChatCompletionModelConfigurationNoKind](actions.md#externalchatcompletionmodelconfigurationnokind), [OnInactivity](triggers.md#oninactivity), [VoiceInputSettings](settings.md#voiceinputsettings), [VoiceInputSettingsNoKind](actions.md#voiceinputsettingsnokind), [VoiceTypeBase](settings.md#voicetypebase), [VoiceTypeBaseNoKind](settings.md#voicetypebasenokind)</sub>

### StringExpression

*Title:* LiteralValue  

**Type:** `string`

<sub>Used by: [Activity](messages.md#activity), [AnswerQuestionWithAI](actions.md#answerquestionwithai), [AutomaticHoldSettings](actions.md#automaticholdsettings), [AzureCognitiveSearchInternalDataSource](knowledge.md#azurecognitivesearchinternaldatasource), [AzureCognitiveSearchInternalDataSourceNoKind](knowledge.md#azurecognitivesearchinternaldatasourcenokind), [AzureOpenAIOnYourDataKnowledgeSource](knowledge.md#azureopenaionyourdataknowledgesource), [AzureOpenAIOnYourDataSource](knowledge.md#azureopenaionyourdatasource), [AzureOpenAIOnYourDataSourceNoKind](actions.md#azureopenaionyourdatasourcenokind), [BingCustomSearchDataSource](knowledge.md#bingcustomsearchdatasource), [BingCustomSearchSource](knowledge.md#bingcustomsearchsource), [BingSearchOptions](knowledge.md#bingsearchoptions), [BingSearchOptionsNoKind](actions.md#bingsearchoptionsnokind), [CancelAllDialogs](actions.md#cancelalldialogs), [CancelDialog](actions.md#canceldialog), [CreateSearchQuery](actions.md#createsearchquery), [DeleteActivity](actions.md#deleteactivity), [EmailSearchSource](knowledge.md#emailsearchsource), [EmitEvent](actions.md#emitevent), [EventActivity](messages.md#eventactivity), [ExternalChatCompletionModelConfiguration](actions.md#externalchatcompletionmodelconfiguration) … and 37 more</sub>

### StructuredRecordExpression

YAML: `kind: StructuredRecordExpression`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `properties` | map of name → [ValueExpression](#valueexpression) |  |  |  |

_No other properties are allowed._

### StructuredRecordExpressionNoKind

*Title:* StructuredRecordExpression  
Same shape as [StructuredRecordExpression](#structuredrecordexpression) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `properties` | map of name → [ValueExpression](#valueexpression) |  |  |  |

_No other properties are allowed._

### ValidationErrorCode


**String enum.** Allowed values: `MissingRequiredProperty`, `PropertyNotSupported`, `ElementNotSupported`, `EmptyCollection`, `EmptyActions`, `SignInRequiredWithManualAuthentication`, `SynonymsNotUnique`, `UnknownEnumValue`, `InvalidPropertyPath`, `DisplayNameMustBeScalar`, `MissingDisplayName`, `SynonymsMustBeArrayOfScalars`, `AuthenticationInputNotEnabled`, `ManualAuthenticationInputNotEnabled`, `DuplicateItem`, `DuplicateVariableAssignment`, `DuplicateActionId`, `DuplicateConditionId`, `OverlappingTaskDialogInputPropertyPath`, `DuplicateTaskDialogInputPropertyPath`, `DuplicateDtmfKey`, `ValueCannotBeNegative`, `ConnectorDynamicTypeError`, `ConnectorNotSupported`, `ConnectorActionNotSupported`, `McpToolNotFound`, `ConnectionNotSet`, `ConnectionDoesNotExist`, `ConnectionNotShared`, `ReservedOptionSetName`, `DuplicateOptionSetName`, `DuplicateMultiEntitySelection`, `DuplicateMultiEntityPropertyName`, `OneOfMultiEntityMinimumSelectedEntities`, `SameDTMFKeyAcrossSelectedMultiEntities`, `OptionSetNameHidesBuiltInName`, `AdaptiveCardJsonParsingError`, `AdaptiveCardMissingActionSubmit`, `AdaptiveCardUseOfReservedKeyword`, `AdaptiveCardActionSubmitIdTypeMismatch`, `AdaptiveCardInputIsRequiredMissingErrorMessage`, `AdaptiveCardInputIsRequiredMissingLabel`, `AdaptiveCardVersionNotSupported`, `AdaptiveCardToggleInputMissingTitle`, `OpenAIDialogActionInvalidOutput`, `BingCustomSearchDataSourceInvalidCustomConfigurationId`, `PublicSiteSearchDataSourceTooManySites`, `PublicSiteSearchDataSourceNotSupported`, `SharePointSearchDataSourceTooManySites`, `TooManyAzureOpenAIOnYourDataKnowledgeSources`, `TooManyPublicSiteSearchSources`, `TooManyBingCustomSearchSources`, `TooManyDataverseStructuredSearchSources`, `TooManyFederatedStructuredSearchSources`, `TooManyFederatedIngestionBasedSearchSources`, `TooManyAzureAISearchSources`, `FederatedStructuredSearchSourceNotConfigured`, `ExternalAgentTaskActionDoesNotAllowRedirect`, `InvalidMakerConnectionReference`, `TooManyFabricAISkillsSearchSources`, `EnvironmentVariableMissingValue`, `TemplateNotFound`, `TemplateLanguageNotSupported`, `SynchronizationSystemError`, `InvalidSipHeader`, `InvalidTransferPhoneNumber`, `InvalidDtmfSequence`, `DtmfSequenceLengthExceeded`, `GptBlocked`, `BingSearchBlocked`, `AdditionalInstructionsLengthGreaterThanMaxLengthAllowed`, `UserConnectionUnsupportedAuthentication`, `UserConnectionRequiresAuthentication`, `MustPromptUserWhenRequired`, `MaxCountExceeded`, `MinCountNotMet`, `MaxLengthExceeded`, `GraphConnectorSearchSourceTooManyConnections`, `GraphConnectionNotSet`, `GptValidationMissingName`, `GptValidationMissingDescription`, `GptValidationMissingInstructions`, `GptValidationOffensiveContent`, `GptValidationJailbreakAttempted`, `GptValidationRAIViolation`, `GptValidationUnknownFailure`, `InvalidGuid`, `MakerConnectionIsUsed`, `PublishWithAIFeaturesDisabled`, `TemperatureShouldBeNonNegative`, `TopPMustBeBetweenZeroAndOne`, `MaxTokensShouldBePositive`, `ChatMlSyntaxErrorEmptyMessageContent`, `ChatMlSyntaxErrorExpectedEndTag`, `ChatMlSyntaxErrorInvalidMessageRole`, `ChatMlSyntaxErrorExpectedStartTag`, `CustomModelMissingInstructions`, `CustomModelMissingDataSourcesVariableUsage`, `CustomModelMissingUserQueryVariableUsage`, `ExternalChatCompletionModelConfigurationInvalidInput`, `AnnotationMissingForEntitySample`, `InvalidAnnotationForEntitySample`, `MultipleAnnotationsNotAllowedForEntitySample`, `InvalidAnnotationForTopicSample`, `TriggerRecurrenceTooSmall`, `TriggerRecurrenceInvalidInput`, `TriggerRecurrenceInMakerModeOnlySupportManualTaskInput`, `InvalidBindingSensitivityLevel`, `InvalidBindingOutput`, `InvalidBindingInput`, `InvalidBindingInvokeAction`, `InvalidBinding`, `IncorrectTypeAssignment`, `IncompatibleTypeComparison`, `AgentDialogContainsActionsWithinTrigger`, `MakerAuthenticationIsBlocked`, `InvalidUrl`, `InvalidGraphConnectorPath`, `InvalidEmailAddress`, `InvalidSipXHeaderName`, `InvalidSipXHeaderValue`, `InvalidAppId`, `InvalidScopeUrl`, `MissingAgentIdentity`

### ValueExpression


**One of:** `string`, `number`, `boolean`, list of any, object

<sub>Used by: [ActionInputBinding](actions.md#actioninputbinding), [ActionInputBindingNoKind](actions.md#actioninputbindingnokind), [Activity](messages.md#activity), [AddItemOperation](actions.md#additemoperation), [AutomaticTaskInput](task-actions.md#automatictaskinput), [BooleanCondition](actions.md#booleancondition), [CancelAllDialogs](actions.md#cancelalldialogs), [CancelDialog](actions.md#canceldialog), [ConnectionParametersSet](components.md#connectionparametersset), [ConnectionParametersSetNoKind](components.md#connectionparameterssetnokind), [EditTable](actions.md#edittable), [EmitEvent](actions.md#emitevent), [EndDialog](actions.md#enddialog), [EventActivity](messages.md#eventactivity), [ExpressionSegment](messages.md#expressionsegment), [ExpressionTransferContext](actions.md#expressiontransfercontext), [Foreach](actions.md#foreach), [JsonRequestContent](actions.md#jsonrequestcontent), [LogCustomTelemetryEvent](actions.md#logcustomtelemetryevent), [ManualTaskInput](task-actions.md#manualtaskinput) … and 19 more</sub>
