# Actions (topic nodes)

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

Every `kind` that can appear in an `actions:` list inside a topic, plus their supporting types. See [05 Nodes](../../05-nodes.md).

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

**137 definitions** on this page (88 with a `kind:`). [A–Z index of all definitions](index.md).

<details><summary>Contents</summary>

[ActionInputBinding](#actioninputbinding), [ActionInputBindingBase](#actioninputbindingbase), [ActionInputBindingNoKind](#actioninputbindingnokind), [ActionOutputBinding](#actionoutputbinding), [ActionOutputBindingBase](#actionoutputbindingbase), [ActionOutputBindingNoKind](#actionoutputbindingnokind), [ActionScope](#actionscope), [ActivateExternalTrigger](#activateexternaltrigger), [AdaptiveCardPrompt](#adaptivecardprompt), [AddItemOperation](#additemoperation), [AndConditionGroup](#andconditiongroup), [AnswerQuestionWithAI](#answerquestionwithai), [AutomaticHoldSettings](#automaticholdsettings), [AutomaticTransferContext](#automatictransfercontext), [AutomaticTransferVariables](#automatictransfervariables), [AzureOpenAIOnYourDataSourceNoKind](#azureopenaionyourdatasourcenokind), [BeginDialog](#begindialog), [BingSearchOptionsNoKind](#bingsearchoptionsnokind), [BooleanCondition](#booleancondition), [BooleanConditionBase](#booleanconditionbase), [BooleanConditionGroup](#booleanconditiongroup), [BooleanConditionOperator](#booleanconditionoperator), [BreakLoop](#breakloop), [CancelAllDialogs](#cancelalldialogs), [CancelDialog](#canceldialog), [ClearAllVariables](#clearallvariables), [ClearItemsOperation](#clearitemsoperation), [ConditionGroup](#conditiongroup), [ConditionItem](#conditionitem), [ConditionItemNoKind](#conditionitemnokind), [ContinueLoop](#continueloop), [ContinueOnErrorBehavior](#continueonerrorbehavior), [ConversationOutcome](#conversationoutcome), [CreateSearchQuery](#createsearchquery), [CSATQuestion](#csatquestion), [CustomDataSourceNoKind](#customdatasourcenokind), [DefaultValueMissingAction](#defaultvaluemissingaction), [DeleteActivity](#deleteactivity), [DialogAction](#dialogaction), [DisableTrigger](#disabletrigger), [DisconnectedNodeContainer](#disconnectednodecontainer), [DoNotHoldSettings](#donotholdsettings), [DoNotSearchFiles](#donotsearchfiles), [DtmfTerminationOptions](#dtmfterminationoptions), [EditTable](#edittable), [EditTableOperation](#edittableoperation), [EditTableV2](#edittablev2), [EmitEvent](#emitevent), [EndConversation](#endconversation), [EndDialog](#enddialog), [EventType](#eventtype), [ExpressionTransferContext](#expressiontransfercontext), [ExternalChatCompletionModelConfiguration](#externalchatcompletionmodelconfiguration), [ExternalChatCompletionModelConfigurationNoKind](#externalchatcompletionmodelconfigurationnokind), [ExternalModelConfiguration](#externalmodelconfiguration), [FixedLengthTermination](#fixedlengthtermination), [FlexibleLengthTermination](#flexiblelengthtermination), [Foreach](#foreach), [GenerativeModality](#generativemodality), [GenerativeResponseOutputType](#generativeresponseoutputtype), [GetActivityMembers](#getactivitymembers), [GetConversationMembers](#getconversationmembers), [GotoAction](#gotoaction), [HoldSettings](#holdsettings), [HttpErrorHandlingBehavior](#httperrorhandlingbehavior), [HttpMethodType](#httpmethodtype), [HttpRequestAction](#httprequestaction), [HttpRequestContent](#httprequestcontent), [IncludeSelectedTopics](#includeselectedtopics), [InitializationBehavior](#initializationbehavior), [InvocationSource](#invocationsource), [InvokeAIBuilderModelAction](#invokeaibuildermodelaction), [InvokeConnectorAction](#invokeconnectoraction), [InvokeCustomModelAction](#invokecustommodelaction), [InvokeFlowAction](#invokeflowaction), [InvokeMcpToolAction](#invokemcptoolaction), [InvokeSkillAction](#invokeskillaction), [JsonRequestContent](#jsonrequestcontent), [KeyBasedTermination](#keybasedtermination), [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind), [LogCustomTelemetryEvent](#logcustomtelemetryevent), [MultipleTopicsMatchedReason](#multipletopicsmatchedreason), [NoRequestContent](#norequestcontent), [NoRequestContentNoKind](#norequestcontentnokind), [OAuthInput](#oauthinput), [OrConditionGroup](#orconditiongroup), [ParseValue](#parsevalue), [ParseValueErrorPolicy](#parsevalueerrorpolicy), [PartType](#parttype), [PropertyOverrideNoKind](#propertyoverridenokind), [Question](#question), [RaiseErrorBehavior](#raiseerrorbehavior), [RaiseErrorBehaviorNoKind](#raiseerrorbehaviornokind), [RawRequestContent](#rawrequestcontent), [RecognizeIntent](#recognizeintent), [RemoveItemOperation](#removeitemoperation), [RepeatDialog](#repeatdialog), [ReplaceDialog](#replacedialog), [ResetVariable](#resetvariable), [SearchAllFiles](#searchallfiles), [SearchAndSummarizeContent](#searchandsummarizecontent), [SearchAndSummarizeWithCustomModel](#searchandsummarizewithcustommodel), [SearchFilesMode](#searchfilesmode), [SearchKnowledgeSources](#searchknowledgesources), [SearchSpecificFiles](#searchspecificfiles), [SendActivity](#sendactivity), [SetMultipleVariables](#setmultiplevariables), [SetTextVariable](#settextvariable), [SetVariable](#setvariable), [SharePointSearchDataSourceNoKind](#sharepointsearchdatasourcenokind), [SignInReason](#signinreason), [SignOutUser](#signoutuser), [SingleValueInputBinding](#singlevalueinputbinding), [SingleVariableOutputBinding](#singlevariableoutputbinding), [SkipQuestionMode](#skipquestionmode), [TableChangeType](#tablechangetype), [TakeFirstItemOperation](#takefirstitemoperation), [TakeLastItemOperation](#takelastitemoperation), [TaskOutput](#taskoutput), [ToolDialogVariableContext](#tooldialogvariablecontext), [TransferContext](#transfercontext), [TransferConversation](#transferconversation), [TransferConversationV2](#transferconversationv2), [TransferToAgent](#transfertoagent), [TransferToPhoneNumber](#transfertophonenumber), [TransferType](#transfertype), [UpdateActivity](#updateactivity), [UserInputType](#userinputtype), [UserRequestedHoldSettings](#userrequestedholdsettings), [Variable](#variable), [VariableAIVisibility](#variableaivisibility), [VariableAssignment](#variableassignment), [VariableAssignmentNoKind](#variableassignmentnokind), [VariableScope](#variablescope), [VariablesToClear](#variablestoclear), [VoiceInputSettingsNoKind](#voiceinputsettingsnokind), [WaitForConnectorTrigger](#waitforconnectortrigger)

</details>

### ActionInputBinding

YAML: `kind: ActionInputBinding`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `binding` | map of name → [ValueExpression](other.md#valueexpression) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ActionInputBindingBase](#actioninputbindingbase)</sub>

### ActionInputBindingBase


**One of:** [SingleValueInputBinding](#singlevalueinputbinding), [ActionInputBinding](#actioninputbinding), [ActionInputBindingNoKind](#actioninputbindingnokind)

<sub>Used by: [ActivateExternalTrigger](#activateexternaltrigger), [BeginDialog](#begindialog), [InvokeAIBuilderModelAction](#invokeaibuildermodelaction), [InvokeConnectorAction](#invokeconnectoraction), [InvokeFlowAction](#invokeflowaction), [InvokeMcpToolAction](#invokemcptoolaction), [InvokeSkillAction](#invokeskillaction), [RepeatDialog](#repeatdialog), [ReplaceDialog](#replacedialog), [WaitForConnectorTrigger](#waitforconnectortrigger)</sub>

### ActionInputBindingNoKind

*Title:* ActionInputBinding  
Same shape as [ActionInputBinding](#actioninputbinding) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `binding` | map of name → [ValueExpression](other.md#valueexpression) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ActionInputBindingBase](#actioninputbindingbase)</sub>

### ActionOutputBinding

YAML: `kind: ActionOutputBinding`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `outputOverrides` | list of [PropertyOverrideNoKind](#propertyoverridenokind) |  |  |  |
| `binding` | map of name → `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [ActionOutputBindingBase](#actionoutputbindingbase)</sub>

### ActionOutputBindingBase


**One of:** [ActionOutputBinding](#actionoutputbinding), [ActionOutputBindingNoKind](#actionoutputbindingnokind), [SingleVariableOutputBinding](#singlevariableoutputbinding)

<sub>Used by: [BeginDialog](#begindialog), [InvokeAIBuilderModelAction](#invokeaibuildermodelaction), [InvokeConnectorAction](#invokeconnectoraction), [InvokeFlowAction](#invokeflowaction), [InvokeMcpToolAction](#invokemcptoolaction), [InvokeSkillAction](#invokeskillaction), [WaitForConnectorTrigger](#waitforconnectortrigger)</sub>

### ActionOutputBindingNoKind

*Title:* ActionOutputBinding  
Same shape as [ActionOutputBinding](#actionoutputbinding) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `outputOverrides` | list of [PropertyOverrideNoKind](#propertyoverridenokind) |  |  |  |
| `binding` | map of name → `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [ActionOutputBindingBase](#actionoutputbindingbase), [AdaptiveCardPrompt](#adaptivecardprompt)</sub>

### ActionScope

YAML: `kind: ActionScope`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `actions` | list of [DialogAction](#dialogaction) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### ActivateExternalTrigger

YAML: `kind: ActivateExternalTrigger`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `input` | [ActionInputBindingBase](#actioninputbindingbase) |  |  |  |
| `activityProcessed` | [BoolExpression](other.md#boolexpression) |  | `true` | accepts a `=` Power Fx expression |
| `dialog` | [DialogExpression](other.md#dialogexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `externalTriggerId` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### AdaptiveCardPrompt

YAML: `kind: AdaptiveCardPrompt`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `interruptionPolicy` | [InterruptionPolicyNoKind](settings.md#interruptionpolicynokind) |  |  |  |
| `unrecognizedPrompt` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |
| `repeatCount` | `0` \| `1` \| `2` \| [IntExpression](other.md#intexpression) |  | `2` |  |
| `card` | [AdaptiveCardExpression](other.md#adaptivecardexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `output` | [ActionOutputBindingNoKind](#actionoutputbindingnokind) | ✓ |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### AddItemOperation

YAML: `kind: AddItemOperation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `value` | [ValueExpression](other.md#valueexpression) | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [EditTableOperation](#edittableoperation)</sub>

### AndConditionGroup

YAML: `kind: AndConditionGroup`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `conditions` | list of [BooleanConditionBase](#booleanconditionbase) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BooleanConditionBase](#booleanconditionbase), [BooleanConditionGroup](#booleanconditiongroup)</sub>

### AnswerQuestionWithAI

YAML: `kind: AnswerQuestionWithAI`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `latencyMessageSettings` | [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind) |  |  |  |
| `autoSend` | `boolean` |  | `true` |  |
| `variable` | `string` |  |  |  |
| `userInput` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `moderationLevel` | `=` expression \| [ContentModerationLevel](settings.md#contentmoderationlevel) |  | `"High"` | accepts a `=` Power Fx expression |
| `historyType` | [HistoryType](settings.md#historytype) |  |  |  |
| `isFileAnalysisEnabled` | `boolean` |  |  |  |
| `additionalInstructions` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### AutomaticHoldSettings

YAML: `kind: AutomaticHoldSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `resumeWords` | list of `string` |  |  |  |
| `holdMusic` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `holdTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `holdRepeatCount` | `0` \| `1` \| `2` \| `3` \| `4` \| [IntExpression](other.md#intexpression) |  |  |  |
| `holdTimeoutMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `resumeResponse` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [HoldSettings](#holdsettings)</sub>

### AutomaticTransferContext

YAML: `kind: AutomaticTransferContext`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `variables` | [AutomaticTransferVariables](#automatictransfervariables) |  | `"All"` |  |

_No other properties are allowed._

<sub>Used by: [TransferContext](#transfercontext)</sub>

### AutomaticTransferVariables


**String enum.** Allowed values: `Topic`, `Global`, `All`

<sub>Used by: [AutomaticTransferContext](#automatictransfercontext)</sub>

### AzureOpenAIOnYourDataSourceNoKind

Same shape as [AzureOpenAIOnYourDataSource](knowledge.md#azureopenaionyourdatasource) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionReference` | `string` |  |  |  |
| `dataSources` | list of [AzureCognitiveSearchInternalDataSourceNoKind](knowledge.md#azurecognitivesearchinternaldatasourcenokind) | ✓ |  |  |
| `apiVersion` | `"2023-06-01-preview"` \| `"2024-02-15-preview"` \| [StringExpression](other.md#stringexpression) | ✓ |  |  |
| `deployment` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `temperature` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `topP` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `stopSequence` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `maxTokens` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [SearchAndSummarizeContent](#searchandsummarizecontent)</sub>

### BeginDialog

YAML: `kind: BeginDialog`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `input` | [ActionInputBindingBase](#actioninputbindingbase) |  |  |  |
| `activityProcessed` | [BoolExpression](other.md#boolexpression) |  | `true` | accepts a `=` Power Fx expression |
| `dialog` | [DialogExpression](other.md#dialogexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `output` | [ActionOutputBindingBase](#actionoutputbindingbase) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### BingSearchOptionsNoKind

*Title:* BingSearchOptions  
Same shape as [BingSearchOptions](knowledge.md#bingsearchoptions) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `market` | `"es-AR"` \| `"en-AU"` \| `"de-AT"` \| `"nl-BE"` \| `"fr-BE"` \| `"pt-BR"` \| `"en-CA"` \| `"fr-CA"` \| `"es-CL"` \| `"da-DK"` \| `"fi-FI"` \| `"fr-FR"` \| `"de-DE"` \| `"zh-HK"` \| `"en-IN"` \| `"en-ID"` \| `"it-IT"` \| `"ja-JP"` \| `"ko-KR"` \| `"en-MY"` \| `"es-MX"` \| `"nl-NL"` \| `"en-NZ"` \| `"no-NO"` \| `"zh-CN"` \| `"pl-PL"` \| `"en-PH"` \| `"ru-RU"` \| `"en-ZA"` \| `"es-ES"` \| `"sv-SE"` \| `"fr-CH"` \| `"de-CH"` \| `"zh-TW"` \| `"tr-TR"` \| `"en-GB"` \| `"en-US"` \| `"es-US"` \| [StringExpression](other.md#stringexpression) |  |  |  |

_No other properties are allowed._

<sub>Used by: [SearchAndSummarizeContent](#searchandsummarizecontent), [SearchAndSummarizeWithCustomModel](#searchandsummarizewithcustommodel), [SearchKnowledgeSources](#searchknowledgesources)</sub>

### BooleanCondition

YAML: `kind: BooleanCondition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `variable` | `string` | ✓ |  |  |
| `operator` | [BooleanConditionOperator](#booleanconditionoperator) | ✓ |  |  |
| `value` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [BooleanConditionBase](#booleanconditionbase)</sub>

### BooleanConditionBase


**One of:** [AndConditionGroup](#andconditiongroup), [OrConditionGroup](#orconditiongroup), [BooleanCondition](#booleancondition)

<sub>Used by: [AndConditionGroup](#andconditiongroup), [OrConditionGroup](#orconditiongroup)</sub>

### BooleanConditionGroup


**One of:** [AndConditionGroup](#andconditiongroup), [OrConditionGroup](#orconditiongroup)

### BooleanConditionOperator


**String enum.** Allowed values: `Equal`, `NotEqual`, `Greater`, `GreaterEqual`, `Less`, `LessEqual`, `In`, `ExactIn`, `IsBlank`, `IsNotBlank`, `IsEmpty`, `IsNotEmpty`

<sub>Used by: [BooleanCondition](#booleancondition)</sub>

### BreakLoop

YAML: `kind: BreakLoop`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### CancelAllDialogs

YAML: `kind: CancelAllDialogs`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `activityProcessed` | [BoolExpression](other.md#boolexpression) |  | `true` | accepts a `=` Power Fx expression |
| `eventName` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `eventValue` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### CancelDialog

YAML: `kind: CancelDialog`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `activityProcessed` | [BoolExpression](other.md#boolexpression) |  | `true` | accepts a `=` Power Fx expression |
| `eventName` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `eventValue` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### ClearAllVariables

YAML: `kind: ClearAllVariables`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `variables` | `=` expression \| [VariablesToClear](#variablestoclear) |  | `"ConversationScopedVariables"` | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### ClearItemsOperation

YAML: `kind: ClearItemsOperation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [EditTableOperation](#edittableoperation)</sub>

### ConditionGroup

YAML: `kind: ConditionGroup`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `conditions` | list of [ConditionItemNoKind](#conditionitemnokind) | ✓ |  |  |
| `elseActions` | list of [DialogAction](#dialogaction) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### ConditionItem

YAML: `kind: ConditionItem`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `actions` | list of [DialogAction](#dialogaction) |  |  |  |

_No other properties are allowed._

### ConditionItemNoKind

Same shape as [ConditionItem](#conditionitem) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `actions` | list of [DialogAction](#dialogaction) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ConditionGroup](#conditiongroup)</sub>

### ContinueLoop

YAML: `kind: ContinueLoop`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### ContinueOnErrorBehavior

YAML: `kind: ContinueOnErrorBehavior`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `statusCode` | `string` |  |  |  |
| `errorResponseBody` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [HttpErrorHandlingBehavior](#httperrorhandlingbehavior)</sub>

### ConversationOutcome


**String enum.** Allowed values: `ResolvedConfirmed`, `ResolvedImplied`, `Escalated`

<sub>Used by: [ActionScope](#actionscope), [ActivateExternalTrigger](#activateexternaltrigger), [AdaptiveCardPrompt](#adaptivecardprompt), [AnswerQuestionWithAI](#answerquestionwithai), [BeginDialog](#begindialog), [BreakLoop](#breakloop), [CSATQuestion](#csatquestion), [CancelAllDialogs](#cancelalldialogs), [CancelDialog](#canceldialog), [ClearAllVariables](#clearallvariables), [ConditionGroup](#conditiongroup), [ContinueLoop](#continueloop), [CreateSearchQuery](#createsearchquery), [DeleteActivity](#deleteactivity), [DisableTrigger](#disabletrigger), [DisconnectedNodeContainer](#disconnectednodecontainer), [EditTable](#edittable), [EditTableV2](#edittablev2), [EmitEvent](#emitevent), [EndConversation](#endconversation) … and 32 more</sub>

### CreateSearchQuery

YAML: `kind: CreateSearchQuery`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `userInput` | [StringExpression](other.md#stringexpression) |  | `"=System.Activity.Text"` | accepts a `=` Power Fx expression |
| `moderationLevel` | `=` expression \| [ContentModerationLevel](settings.md#contentmoderationlevel) |  | `"High"` | accepts a `=` Power Fx expression |
| `historyType` | [HistoryType](settings.md#historytype) |  |  |  |
| `result` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### CSATQuestion

YAML: `kind: CSATQuestion`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `variable` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### CustomDataSourceNoKind

Same shape as [CustomDataSource](knowledge.md#customdatasource) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `searchResults` | `=` expression |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [SearchAndSummarizeContent](#searchandsummarizecontent), [SearchAndSummarizeWithCustomModel](#searchandsummarizewithcustommodel)</sub>

### DefaultValueMissingAction


**String enum.** Allowed values: `HangUp`, `Escalate`, `SetValue`, `SetEmpty`, `GoToDialog`

<sub>Used by: [VoiceInputSettings](settings.md#voiceinputsettings), [VoiceInputSettingsNoKind](#voiceinputsettingsnokind)</sub>

### DeleteActivity

YAML: `kind: DeleteActivity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `activityId` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### DialogAction


**One of 52 types** (see the [A–Z index](index.md)).

<sub>Used by: [ActionScope](#actionscope), [ConditionGroup](#conditiongroup), [ConditionItem](#conditionitem), [ConditionItemNoKind](#conditionitemnokind), [ContinueConversationConnectorTrigger](triggers.md#continueconversationconnectortrigger), [CreateConversationConnectorTrigger](triggers.md#createconversationconnectortrigger), [DisconnectedNodeContainer](#disconnectednodecontainer), [Foreach](#foreach), [OnActivity](triggers.md#onactivity), [OnAnalyzeText](triggers.md#onanalyzetext), [OnAnsweringMachineDetected](triggers.md#onansweringmachinedetected), [OnConversationStart](triggers.md#onconversationstart), [OnCopilotRedirect](triggers.md#oncopilotredirect), [OnDtmfKeyPress](triggers.md#ondtmfkeypress), [OnError](triggers.md#onerror), [OnEscalate](triggers.md#onescalate), [OnEventActivity](triggers.md#oneventactivity), [OnGeneratedResponse](triggers.md#ongeneratedresponse), [OnInactivity](triggers.md#oninactivity), [OnKnowledgeRequested](triggers.md#onknowledgerequested) … and 14 more</sub>

### DisableTrigger

YAML: `kind: DisableTrigger`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `dialog` | [DialogExpression](other.md#dialogexpression) | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### DisconnectedNodeContainer

YAML: `kind: DisconnectedNodeContainer`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `actions` | list of [DialogAction](#dialogaction) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### DoNotHoldSettings

YAML: `kind: DoNotHoldSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [HoldSettings](#holdsettings)</sub>

### DoNotSearchFiles

YAML: `kind: DoNotSearchFiles`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [SearchFilesMode](#searchfilesmode)</sub>

### DtmfTerminationOptions


**One of:** [FixedLengthTermination](#fixedlengthtermination), [FlexibleLengthTermination](#flexiblelengthtermination), [KeyBasedTermination](#keybasedtermination)

<sub>Used by: [AgePrebuiltEntity](entities.md#ageprebuiltentity), [DatePrebuiltEntity](entities.md#dateprebuiltentity), [MoneyPrebuiltEntity](entities.md#moneyprebuiltentity), [NumberPrebuiltEntity](entities.md#numberprebuiltentity), [OrdinalPrebuiltEntity](entities.md#ordinalprebuiltentity), [PercentagePrebuiltEntity](entities.md#percentageprebuiltentity), [PhoneNumberPrebuiltEntity](entities.md#phonenumberprebuiltentity), [RegexEntityReference](entities.md#regexentityreference), [SpeedPrebuiltEntity](entities.md#speedprebuiltentity), [StringPrebuiltEntity](entities.md#stringprebuiltentity), [TemperaturePrebuiltEntity](entities.md#temperatureprebuiltentity), [WeightPrebuiltEntity](entities.md#weightprebuiltentity), [ZipCodePrebuiltEntity](entities.md#zipcodeprebuiltentity)</sub>

### EditTable

YAML: `kind: EditTable`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `changeType` | [TableChangeType](#tablechangetype) |  | `"Add"` |  |
| `itemsVariable` | `string` | ✓ |  |  |
| `resultVariable` | `string` |  |  |  |
| `value` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### EditTableOperation


**One of:** [AddItemOperation](#additemoperation), [RemoveItemOperation](#removeitemoperation), [TakeFirstItemOperation](#takefirstitemoperation), [TakeLastItemOperation](#takelastitemoperation), [ClearItemsOperation](#clearitemsoperation)

<sub>Used by: [EditTableV2](#edittablev2)</sub>

### EditTableV2

YAML: `kind: EditTableV2`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `itemsVariable` | `string` | ✓ |  |  |
| `changeType` | [EditTableOperation](#edittableoperation) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### EmitEvent

YAML: `kind: EmitEvent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `eventType` | [EventType](#eventtype) |  | `"ActivityIfUnhandled"` |  |
| `eventName` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `eventValue` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `bubbleEvent` | [BoolExpression](other.md#boolexpression) |  | `false` | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### EndConversation

YAML: `kind: EndConversation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### EndDialog

YAML: `kind: EndDialog`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `clearTopicQueue` | `boolean` |  | `true` |  |
| `value` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### EventType


**String enum.** Allowed values: `ActivityIfUnhandled`, `Activity`, `DialogEvent`

<sub>Used by: [EmitEvent](#emitevent)</sub>

### ExpressionTransferContext

YAML: `kind: ExpressionTransferContext`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `value` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [TransferContext](#transfercontext)</sub>

### ExternalChatCompletionModelConfiguration

YAML: `kind: ExternalChatCompletionModelConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionReference` | `string` | ✓ |  |  |
| `temperature` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `topP` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `stopSequence` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `maxTokens` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `additionalProperties` | `=` expression \| object |  |  | accepts a `=` Power Fx expression |
| `deployment` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `historyType` | [HistoryType](settings.md#historytype) |  |  |  |
| `prompt` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `systemPrompt` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ExternalModelConfiguration](#externalmodelconfiguration)</sub>

### ExternalChatCompletionModelConfigurationNoKind

*Title:* ExternalChatCompletionModelConfiguration  
Same shape as [ExternalChatCompletionModelConfiguration](#externalchatcompletionmodelconfiguration) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionReference` | `string` | ✓ |  |  |
| `temperature` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `topP` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `stopSequence` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `maxTokens` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `additionalProperties` | `=` expression \| object |  |  | accepts a `=` Power Fx expression |
| `deployment` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `historyType` | [HistoryType](settings.md#historytype) |  |  |  |
| `prompt` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `systemPrompt` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [SummarizationExternalModelConfiguration](knowledge.md#summarizationexternalmodelconfiguration), [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind)</sub>

### ExternalModelConfiguration


**One of:** [ExternalChatCompletionModelConfiguration](#externalchatcompletionmodelconfiguration)

<sub>Used by: [InvokeCustomModelAction](#invokecustommodelaction)</sub>

### FixedLengthTermination

YAML: `kind: FixedLengthTermination`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `acceptsVoice` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `key` | `=` expression \| [DtmfTerminationKey](settings.md#dtmfterminationkey) |  |  | accepts a `=` Power Fx expression |
| `interdigitTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `terminationTimeoutInMilliseconds` | `0` \| [IntExpression](other.md#intexpression) |  |  |  |
| `length` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DtmfTerminationOptions](#dtmfterminationoptions)</sub>

### FlexibleLengthTermination

YAML: `kind: FlexibleLengthTermination`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `acceptsVoice` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `key` | `=` expression \| [DtmfTerminationKey](settings.md#dtmfterminationkey) |  |  | accepts a `=` Power Fx expression |
| `interdigitTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `terminationTimeoutInMilliseconds` | `0` \| [IntExpression](other.md#intexpression) |  |  |  |
| `minLength` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `maxLength` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DtmfTerminationOptions](#dtmfterminationoptions)</sub>

### Foreach

YAML: `kind: Foreach`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `actions` | list of [DialogAction](#dialogaction) |  |  |  |
| `items` | [ValueExpression](other.md#valueexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `value` | `string` | ✓ |  |  |
| `index` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### GenerativeModality


**String enum.** Allowed values: `Automatic`, `Text`, `Speech`, `TextAndSpeech`

<sub>Used by: [SearchAndSummarizeContent](#searchandsummarizecontent)</sub>

### GenerativeResponseOutputType


**String enum.** Allowed values: `TextOnly`, `FullResponse`

<sub>Used by: [SearchAndSummarizeContent](#searchandsummarizecontent), [SearchAndSummarizeWithCustomModel](#searchandsummarizewithcustommodel)</sub>

### GetActivityMembers

YAML: `kind: GetActivityMembers`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `variable` | `string` |  |  |  |
| `activityId` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### GetConversationMembers

YAML: `kind: GetConversationMembers`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `variable` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### GotoAction

YAML: `kind: GotoAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `actionId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### HoldSettings


**One of:** [DoNotHoldSettings](#donotholdsettings), [AutomaticHoldSettings](#automaticholdsettings), [UserRequestedHoldSettings](#userrequestedholdsettings)

<sub>Used by: [Question](#question)</sub>

### HttpErrorHandlingBehavior


**One of:** [RaiseErrorBehavior](#raiseerrorbehavior), [RaiseErrorBehaviorNoKind](#raiseerrorbehaviornokind), [ContinueOnErrorBehavior](#continueonerrorbehavior)

<sub>Used by: [HttpRequestAction](#httprequestaction), [InvokeConnectorAction](#invokeconnectoraction), [InvokeCustomModelAction](#invokecustommodelaction), [WaitForConnectorTrigger](#waitforconnectortrigger)</sub>

### HttpMethodType


**String enum.** Allowed values: `Get`, `Post`, `Patch`, `Put`, `Delete`

<sub>Used by: [HttpRequestAction](#httprequestaction)</sub>

### HttpRequestAction

YAML: `kind: HttpRequestAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `latencyMessageSettings` | [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind) |  |  |  |
| `method` | `=` expression \| [HttpMethodType](#httpmethodtype) |  | `"Get"` | accepts a `=` Power Fx expression |
| `url` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `headers` | map of name → [StringExpression](other.md#stringexpression) |  |  |  |
| `body` | [HttpRequestContent](#httprequestcontent) |  |  |  |
| `errorHandling` | [HttpErrorHandlingBehavior](#httperrorhandlingbehavior) |  |  |  |
| `requestTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  | `30000` | accepts a `=` Power Fx expression |
| `networkIsolation` | `=` expression \| [NetworkIsolationType](settings.md#networkisolationtype) |  |  | accepts a `=` Power Fx expression |
| `response` | `string` |  |  |  |
| `responseSchema` | [DataType](data-types.md#datatype) |  |  |  |
| `responseHeaders` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### HttpRequestContent


**One of:** [NoRequestContent](#norequestcontent), [NoRequestContentNoKind](#norequestcontentnokind), [JsonRequestContent](#jsonrequestcontent), [RawRequestContent](#rawrequestcontent)

<sub>Used by: [HttpRequestAction](#httprequestaction)</sub>

### IncludeSelectedTopics

YAML: `kind: IncludeSelectedTopics`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `allowedInterruptTopics` | list of `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [InterruptionTopicListFilter](settings.md#interruptiontopiclistfilter)</sub>

### InitializationBehavior


**String enum.** Allowed values: `UseDefault`, `ThrowError`, `Escalate`

<sub>Used by: [Variable](#variable), [VariableNoKind](components.md#variablenokind)</sub>

### InvocationSource


**String enum.** Allowed values: `User`, `Agent`

### InvokeAIBuilderModelAction

YAML: `kind: InvokeAIBuilderModelAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `latencyMessageSettings` | [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind) |  |  |  |
| `input` | [ActionInputBindingBase](#actioninputbindingbase) |  |  |  |
| `output` | [ActionOutputBindingBase](#actionoutputbindingbase) |  |  |  |
| `aIModelId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### InvokeConnectorAction

YAML: `kind: InvokeConnectorAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `latencyMessageSettings` | [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind) |  |  |  |
| `input` | [ActionInputBindingBase](#actioninputbindingbase) |  |  |  |
| `output` | [ActionOutputBindingBase](#actionoutputbindingbase) |  |  |  |
| `connectionReference` | `string` | ✓ |  |  |
| `connectionProperties` | [ConnectionPropertiesNoKind](settings.md#connectionpropertiesnokind) |  |  |  |
| `dynamicInputSchema` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `dynamicOutputSchema` | [DataType](data-types.md#datatype) |  |  |  |
| `errorHandling` | [HttpErrorHandlingBehavior](#httperrorhandlingbehavior) |  |  |  |
| `requestTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  | `30000` | accepts a `=` Power Fx expression |
| `operationId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### InvokeCustomModelAction

YAML: `kind: InvokeCustomModelAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `latencyMessageSettings` | [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind) |  |  |  |
| `autoSend` | `boolean` |  | `true` |  |
| `variable` | `string` |  |  |  |
| `modelConfiguration` | [ExternalModelConfiguration](#externalmodelconfiguration) | ✓ |  |  |
| `requestTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  | `30000` | accepts a `=` Power Fx expression |
| `errorHandling` | [HttpErrorHandlingBehavior](#httperrorhandlingbehavior) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### InvokeFlowAction

YAML: `kind: InvokeFlowAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `latencyMessageSettings` | [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind) |  |  |  |
| `input` | [ActionInputBindingBase](#actioninputbindingbase) |  |  |  |
| `output` | [ActionOutputBindingBase](#actionoutputbindingbase) |  |  |  |
| `flowId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### InvokeMcpToolAction

YAML: `kind: InvokeMcpToolAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `latencyMessageSettings` | [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind) |  |  |  |
| `input` | [ActionInputBindingBase](#actioninputbindingbase) |  |  |  |
| `output` | [ActionOutputBindingBase](#actionoutputbindingbase) |  |  |  |
| `serverSchemaName` | `string` | ✓ |  |  |
| `toolName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### InvokeSkillAction

YAML: `kind: InvokeSkillAction`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `latencyMessageSettings` | [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind) |  |  |  |
| `input` | [ActionInputBindingBase](#actioninputbindingbase) |  |  |  |
| `output` | [ActionOutputBindingBase](#actionoutputbindingbase) |  |  |  |
| `skillId` | `string` | ✓ |  |  |
| `actionId` | `string` | ✓ |  |  |
| `interruptionPolicy` | [InterruptionPolicyNoKind](settings.md#interruptionpolicynokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### JsonRequestContent

YAML: `kind: JsonRequestContent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `content` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [HttpRequestContent](#httprequestcontent)</sub>

### KeyBasedTermination

YAML: `kind: KeyBasedTermination`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `acceptsVoice` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `key` | `=` expression \| [DtmfTerminationKey](settings.md#dtmfterminationkey) |  |  | accepts a `=` Power Fx expression |
| `interdigitTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `terminationTimeoutInMilliseconds` | `0` \| [IntExpression](other.md#intexpression) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DtmfTerminationOptions](#dtmfterminationoptions)</sub>

### LatencyMessageSettingsNoKind

*Title:* LatencyMessageSettings  
Same shape as [LatencyMessageSettings](messages.md#latencymessagesettings) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `allowLatencyMessage` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `latencyMessage` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |
| `latencyMessageDelayInMilliseconds` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `minimumPlaybackInMilliseconds` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [AnswerQuestionWithAI](#answerquestionwithai), [HttpRequestAction](#httprequestaction), [InvokeAIBuilderModelAction](#invokeaibuildermodelaction), [InvokeConnectorAction](#invokeconnectoraction), [InvokeCustomModelAction](#invokecustommodelaction), [InvokeFlowAction](#invokeflowaction), [InvokeMcpToolAction](#invokemcptoolaction), [InvokeSkillAction](#invokeskillaction), [SearchAndSummarizeContent](#searchandsummarizecontent), [SearchAndSummarizeWithCustomModel](#searchandsummarizewithcustommodel), [SearchKnowledgeSources](#searchknowledgesources), [WaitForConnectorTrigger](#waitforconnectortrigger)</sub>

### LogCustomTelemetryEvent

YAML: `kind: LogCustomTelemetryEvent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `eventName` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `properties` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### MultipleTopicsMatchedReason


**String enum.** Allowed values: `Interruption`, `Recognition`

### NoRequestContent

YAML: `kind: NoRequestContent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [HttpRequestContent](#httprequestcontent)</sub>

### NoRequestContentNoKind


**Type:** object

<sub>Used by: [HttpRequestContent](#httprequestcontent)</sub>

### OAuthInput

YAML: `kind: OAuthInput`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `title` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `text` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `resource` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### OrConditionGroup

YAML: `kind: OrConditionGroup`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `conditions` | list of [BooleanConditionBase](#booleanconditionbase) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BooleanConditionBase](#booleanconditionbase), [BooleanConditionGroup](#booleanconditiongroup)</sub>

### ParseValue

YAML: `kind: ParseValue`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `variable` | `string` | ✓ |  |  |
| `valueType` | [DataType](data-types.md#datatype) | ✓ |  |  |
| `value` | [ValueExpression](other.md#valueexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `errorPolicy` | [ParseValueErrorPolicy](#parsevalueerrorpolicy) |  | `"DefaultBlank"` |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### ParseValueErrorPolicy


**String enum.** Allowed values: `ThrowError`, `DefaultBlank`

<sub>Used by: [ParseValue](#parsevalue)</sub>

### PartType


**String enum.** Allowed values: `OneNotePart`

<sub>Used by: [SharepointIds](knowledge.md#sharepointids), [SharepointIdsNoKind](knowledge.md#sharepointidsnokind)</sub>

### PropertyOverrideNoKind

*Title:* PropertyOverride  
Same shape as [PropertyOverride](components.md#propertyoverride) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `property` | `string` | ✓ |  |  |
| `sensitivityLevel` | [SensitivityLevel](settings.md#sensitivitylevel) |  | `"None"` |  |

_No other properties are allowed._

<sub>Used by: [ActionOutputBinding](#actionoutputbinding), [ActionOutputBindingNoKind](#actionoutputbindingnokind), [SingleVariableOutputBinding](#singlevariableoutputbinding)</sub>

### Question

YAML: `kind: Question`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `interruptionPolicy` | [InterruptionPolicyNoKind](settings.md#interruptionpolicynokind) |  |  |  |
| `unrecognizedPrompt` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |
| `repeatCount` | `0` \| `1` \| `2` \| [IntExpression](other.md#intexpression) |  | `2` |  |
| `alwaysPrompt` | `false` \| `true` \| [BoolExpression](other.md#boolexpression) |  | `false` |  |
| `skipQuestionMode` | `=` expression \| [SkipQuestionMode](#skipquestionmode) |  | `"SkipOnFirstExecutionIfVariableHasValue"` | accepts a `=` Power Fx expression |
| `variable` | `string` | ✓ |  |  |
| `prompt` | [ActivityTemplateBase](messages.md#activitytemplatebase) | ✓ |  |  |
| `invalidPrompt` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |
| `defaultValueResponse` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |
| `validation` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `defaultValue` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `entity` | [EntityReference](entities.md#entityreference) | ✓ |  |  |
| `voiceInputSettings` | [VoiceInputSettingsNoKind](#voiceinputsettingsnokind) |  |  |  |
| `holdSettings` | [HoldSettings](#holdsettings) |  |  |  |
| `fallbackDialogOnInvalidEntity` | [DialogExpression](other.md#dialogexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### RaiseErrorBehavior

YAML: `kind: RaiseErrorBehavior`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [HttpErrorHandlingBehavior](#httperrorhandlingbehavior)</sub>

### RaiseErrorBehaviorNoKind


**Type:** object

<sub>Used by: [HttpErrorHandlingBehavior](#httperrorhandlingbehavior)</sub>

### RawRequestContent

YAML: `kind: RawRequestContent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `contentType` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `content` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [HttpRequestContent](#httprequestcontent)</sub>

### RecognizeIntent

YAML: `kind: RecognizeIntent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `userInput` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### RemoveItemOperation

YAML: `kind: RemoveItemOperation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `value` | [ValueExpression](other.md#valueexpression) | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [EditTableOperation](#edittableoperation)</sub>

### RepeatDialog

YAML: `kind: RepeatDialog`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `input` | [ActionInputBindingBase](#actioninputbindingbase) |  |  |  |
| `activityProcessed` | [BoolExpression](other.md#boolexpression) |  | `true` | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### ReplaceDialog

YAML: `kind: ReplaceDialog`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `input` | [ActionInputBindingBase](#actioninputbindingbase) |  |  |  |
| `activityProcessed` | [BoolExpression](other.md#boolexpression) |  | `true` | accepts a `=` Power Fx expression |
| `dialog` | [DialogExpression](other.md#dialogexpression) | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### ResetVariable

YAML: `kind: ResetVariable`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `variable` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### SearchAllFiles

YAML: `kind: SearchAllFiles`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [SearchFilesMode](#searchfilesmode)</sub>

### SearchAndSummarizeContent

YAML: `kind: SearchAndSummarizeContent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `latencyMessageSettings` | [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind) |  |  |  |
| `autoSend` | `boolean` |  | `true` |  |
| `variable` | `string` |  |  |  |
| `userInput` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `moderationLevel` | `=` expression \| [ContentModerationLevel](settings.md#contentmoderationlevel) |  | `"High"` | accepts a `=` Power Fx expression |
| `historyType` | [HistoryType](settings.md#historytype) |  |  |  |
| `isFileAnalysisEnabled` | `boolean` |  |  |  |
| `additionalInstructions` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `webBrowsing` | `boolean` |  | `false` |  |
| `webBrowsingOptions` | [BingSearchOptionsNoKind](#bingsearchoptionsnokind) |  |  |  |
| `searchEmails` | `boolean` |  | `false` |  |
| `fileSearchDataSource` | [FileSearchDataSourceNoKind](knowledge.md#filesearchdatasourcenokind) |  |  |  |
| `knowledgeSources` | [SearchKnowledgeSourcesMode](knowledge.md#searchknowledgesourcesmode) |  |  |  |
| `customDataSource` | [CustomDataSourceNoKind](#customdatasourcenokind) |  |  |  |
| `publicDataSource` | [PublicDataSource](knowledge.md#publicdatasource) |  |  |  |
| `sharePointSearchDataSource` | [SharePointSearchDataSourceNoKind](#sharepointsearchdatasourcenokind) |  |  |  |
| `azureOpenAIOnYourDataSource` | [AzureOpenAIOnYourDataSourceNoKind](#azureopenaionyourdatasourcenokind) |  |  |  |
| `applyModelKnowledgeSetting` | `boolean` |  |  |  |
| `responseCaptureType` | [GenerativeResponseOutputType](#generativeresponseoutputtype) |  | `"TextOnly"` |  |
| `modality` | [GenerativeModality](#generativemodality) |  | `"Automatic"` |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### SearchAndSummarizeWithCustomModel

YAML: `kind: SearchAndSummarizeWithCustomModel`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `latencyMessageSettings` | [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind) |  |  |  |
| `autoSend` | `boolean` |  | `true` |  |
| `variable` | `string` |  |  |  |
| `userInput` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `historyType` | [HistoryType](settings.md#historytype) |  |  |  |
| `webBrowsing` | `boolean` |  | `false` |  |
| `webBrowsingOptions` | [BingSearchOptionsNoKind](#bingsearchoptionsnokind) |  |  |  |
| `searchEmails` | `boolean` |  | `false` |  |
| `fileSearchDataSource` | [FileSearchDataSourceNoKind](knowledge.md#filesearchdatasourcenokind) |  |  |  |
| `knowledgeSources` | [SearchKnowledgeSourcesMode](knowledge.md#searchknowledgesourcesmode) |  |  |  |
| `customDataSource` | [CustomDataSourceNoKind](#customdatasourcenokind) |  |  |  |
| `modelConfiguration` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |
| `responseCaptureType` | [GenerativeResponseOutputType](#generativeresponseoutputtype) |  | `"TextOnly"` |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### SearchFilesMode


**One of:** [SearchAllFiles](#searchallfiles), [SearchSpecificFiles](#searchspecificfiles), [DoNotSearchFiles](#donotsearchfiles)

<sub>Used by: [FileSearchDataSource](knowledge.md#filesearchdatasource), [FileSearchDataSourceNoKind](knowledge.md#filesearchdatasourcenokind)</sub>

### SearchKnowledgeSources

YAML: `kind: SearchKnowledgeSources`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `latencyMessageSettings` | [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind) |  |  |  |
| `webBrowsing` | `boolean` |  | `false` |  |
| `webBrowsingOptions` | [BingSearchOptionsNoKind](#bingsearchoptionsnokind) |  |  |  |
| `searchEmails` | `boolean` |  | `false` |  |
| `fileSearchDataSource` | [FileSearchDataSourceNoKind](knowledge.md#filesearchdatasourcenokind) |  |  |  |
| `knowledgeSources` | [SearchKnowledgeSourcesMode](knowledge.md#searchknowledgesourcesmode) |  |  |  |
| `historyType` | [HistoryType](settings.md#historytype) |  |  |  |
| `result` | `string` |  |  |  |
| `userInput` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `moderationLevel` | `=` expression \| [ContentModerationLevel](settings.md#contentmoderationlevel) |  | `"High"` | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### SearchSpecificFiles

YAML: `kind: SearchSpecificFiles`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `files` | `=` expression \| list of `string` | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [SearchFilesMode](#searchfilesmode)</sub>

### SendActivity

YAML: `kind: SendActivity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `activity` | [ActivityTemplateBase](messages.md#activitytemplatebase) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### SetMultipleVariables

YAML: `kind: SetMultipleVariables`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `assignments` | list of [VariableAssignmentNoKind](#variableassignmentnokind) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### SetTextVariable

YAML: `kind: SetTextVariable`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `variable` | `string` | ✓ |  |  |
| `value` | [TemplateLineNoKind](messages.md#templatelinenokind) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### SetVariable

YAML: `kind: SetVariable`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `variable` | `string` | ✓ |  |  |
| `value` | [ValueExpression](other.md#valueexpression) | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### SharePointSearchDataSourceNoKind

Same shape as [SharePointSearchDataSource](knowledge.md#sharepointsearchdatasource) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sites` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `indexBehavior` | [SharePointIndexBehavior](knowledge.md#sharepointindexbehavior) |  | `"EnabledWithFallback"` |  |
| `additionalSearchTerms` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [SearchAndSummarizeContent](#searchandsummarizecontent)</sub>

### SignInReason


**String enum.** Allowed values: `SignInRequired`, `Initializer`

### SignOutUser

YAML: `kind: SignOutUser`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### SingleValueInputBinding

YAML: `kind: SingleValueInputBinding`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `value` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [ActionInputBindingBase](#actioninputbindingbase)</sub>

### SingleVariableOutputBinding

YAML: `kind: SingleVariableOutputBinding`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `outputOverrides` | list of [PropertyOverrideNoKind](#propertyoverridenokind) |  |  |  |
| `variable` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [ActionOutputBindingBase](#actionoutputbindingbase)</sub>

### SkipQuestionMode


**String enum.** Allowed values: `AlwaysAsk`, `SkipOnFirstExecutionIfVariableHasValue`, `AlwaysSkipIfVariableHasValue`

<sub>Used by: [Question](#question)</sub>

### TableChangeType


**String enum.** Allowed values: `Add`, `TakeLast`, `TakeFirst`, `Remove`, `Clear`

<sub>Used by: [EditTable](#edittable)</sub>

### TakeFirstItemOperation

YAML: `kind: TakeFirstItemOperation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `resultVariable` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [EditTableOperation](#edittableoperation)</sub>

### TakeLastItemOperation

YAML: `kind: TakeLastItemOperation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `resultVariable` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [EditTableOperation](#edittableoperation)</sub>

### TaskOutput

YAML: `kind: TaskOutput`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `propertyName` | `string` | ✓ |  |  |
| `name` | `string` |  |  |  |
| `description` | `string` |  |  |  |

_No other properties are allowed._

### ToolDialogVariableContext


**String enum.** Allowed values: `Inputs`, `Confirmation`, `Outputs`

### TransferContext


**One of:** [AutomaticTransferContext](#automatictransfercontext), [ExpressionTransferContext](#expressiontransfercontext)

<sub>Used by: [TransferConversation](#transferconversation), [TransferToAgent](#transfertoagent)</sub>

### TransferConversation

YAML: `kind: TransferConversation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  | `"Escalated"` |  |
| `displayName` | `string` |  |  |  |
| `messageToAgent` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |
| `context` | [TransferContext](#transfercontext) |  |  |  |
| `telephoneNumber` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### TransferConversationV2

YAML: `kind: TransferConversationV2`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  | `"Escalated"` |  |
| `displayName` | `string` |  |  |  |
| `transferType` | [TransferType](#transfertype) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### TransferToAgent

YAML: `kind: TransferToAgent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `messageToAgent` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |
| `context` | [TransferContext](#transfercontext) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TransferType](#transfertype)</sub>

### TransferToPhoneNumber

YAML: `kind: TransferToPhoneNumber`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `phoneNumber` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `sipUuiHeader` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `sipxHeaders` | map of name → [StringExpression](other.md#stringexpression) |  |  |  |
| `dtmfSequence` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [TransferType](#transfertype)</sub>

### TransferType


**One of:** [TransferToAgent](#transfertoagent), [TransferToPhoneNumber](#transfertophonenumber)

<sub>Used by: [TransferConversationV2](#transferconversationv2)</sub>

### UpdateActivity

YAML: `kind: UpdateActivity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `activity` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |
| `activityId` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>

### UserInputType


**String enum.** Allowed values: `Text`, `Speech`, `DTMF`

### UserRequestedHoldSettings

YAML: `kind: UserRequestedHoldSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `holdWords` | list of `string` |  |  |  |
| `holdMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `resumeWords` | list of `string` |  |  |  |
| `holdMusic` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `holdTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `holdRepeatCount` | `0` \| `1` \| `2` \| `3` \| `4` \| [IntExpression](other.md#intexpression) |  |  |  |
| `holdTimeoutMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `resumeResponse` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [HoldSettings](#holdsettings)</sub>

### Variable

YAML: `kind: Variable`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `name` | `string` | ✓ |  |  |
| `scope` | [VariableScope](#variablescope) |  | `"Conversation"` |  |
| `defaultValue` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `isExternalInitializationAllowed` | `boolean` |  | `false` |  |
| `isOutputToExternalCallers` | `boolean` |  | `false` |  |
| `initializationTimeoutInMilliseconds` | `number` |  | `0` |  |
| `missingInitializationBehavior` | [InitializationBehavior](#initializationbehavior) |  | `"UseDefault"` |  |
| `aIVisibility` | [VariableAIVisibility](#variableaivisibility) |  | `"Hidden"` |  |

_No other properties are allowed._

### VariableAIVisibility


**String enum.** Allowed values: `Hidden`, `UseInAIContext`

<sub>Used by: [Variable](#variable), [VariableNoKind](components.md#variablenokind)</sub>

### VariableAssignment

YAML: `kind: VariableAssignment`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `variable` | `string` | ✓ |  |  |
| `value` | [ValueExpression](other.md#valueexpression) | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

### VariableAssignmentNoKind

*Title:* VariableAssignment  
Same shape as [VariableAssignment](#variableassignment) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `variable` | `string` | ✓ |  |  |
| `value` | [ValueExpression](other.md#valueexpression) | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [SetMultipleVariables](#setmultiplevariables)</sub>

### VariableScope


**String enum.** Allowed values: `Dialog`, `Conversation`, `User`, `Environment`

<sub>Used by: [Variable](#variable), [VariableNoKind](components.md#variablenokind)</sub>

### VariablesToClear


**String enum.** Allowed values: `ConversationScopedVariables`, `ConversationHistory`, `UserScopedVariables`, `AllGlobalVariables`

<sub>Used by: [ClearAllVariables](#clearallvariables)</sub>

### VoiceInputSettingsNoKind

*Title:* VoiceInputSettings  
Same shape as [VoiceInputSettings](settings.md#voiceinputsettings) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `wordsetsWeightValue` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `speechRecognitionSensitivity` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `speechRecognitionSpeedVsAccuracy` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `silenceDetectionTimeoutInMilliseconds` | `0` \| [IntExpression](other.md#intexpression) |  |  |  |
| `repeatCountOnSilence` | `0` \| `1` \| `2` \| [IntExpression](other.md#intexpression) |  | `2` |  |
| `inputTimeoutResponse` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |
| `fallbackDialogOnSilence` | [DialogExpression](other.md#dialogexpression) |  |  | accepts a `=` Power Fx expression |
| `defaultValueMissingAction` | `=` expression \| [DefaultValueMissingAction](#defaultvaluemissingaction) |  | `"HangUp"` | accepts a `=` Power Fx expression |
| `defaultValueOnSilence` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `utteranceEndTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `speechRecognitionTimeoutInMilliseconds` | `0` \| [IntExpression](other.md#intexpression) |  |  |  |

_No other properties are allowed._

<sub>Used by: [Question](#question)</sub>

### WaitForConnectorTrigger

YAML: `kind: WaitForConnectorTrigger`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `disabled` | `boolean` |  | `false` |  |
| `conversationOutcome` | [ConversationOutcome](#conversationoutcome) |  |  |  |
| `displayName` | `string` |  |  |  |
| `latencyMessageSettings` | [LatencyMessageSettingsNoKind](#latencymessagesettingsnokind) |  |  |  |
| `input` | [ActionInputBindingBase](#actioninputbindingbase) |  |  |  |
| `output` | [ActionOutputBindingBase](#actionoutputbindingbase) |  |  |  |
| `connectionReference` | `string` | ✓ |  |  |
| `connectionProperties` | [ConnectionPropertiesNoKind](settings.md#connectionpropertiesnokind) |  |  |  |
| `dynamicInputSchema` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `dynamicOutputSchema` | [DataType](data-types.md#datatype) |  |  |  |
| `errorHandling` | [HttpErrorHandlingBehavior](#httperrorhandlingbehavior) |  |  |  |
| `requestTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  | `30000` | accepts a `=` Power Fx expression |
| `operationId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [DialogAction](#dialogaction)</sub>
