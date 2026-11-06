# Settings

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

Agent-level settings blocks (moderation, autonomy, voice, feedback, model selection). See [03 The agent](../../03-agent.md).

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

**103 definitions** on this page (54 with a `kind:`). [A–Z index of all definitions](index.md).

<details><summary>Contents</summary>

[AdaptationManagementServiceModelProfile](#adaptationmanagementservicemodelprofile), [AdaptationManagementServiceModelProfileNoKind](#adaptationmanagementservicemodelprofilenokind), [AdaptationManagementServiceModelType](#adaptationmanagementservicemodeltype), [AdaptationManagementServiceModelTypeNoKind](#adaptationmanagementservicemodeltypenokind), [AgentModelType](#agentmodeltype), [AISettings](#aisettings), [AISettingsNoKind](#aisettingsnokind), [AlwaysRequireMCPApproval](#alwaysrequiremcpapproval), [AnalyticsSettings](#analyticssettings), [AnalyticsSettingsNoKind](#analyticssettingsnokind), [AutonomousSettings](#autonomoussettings), [BotComponentReusePolicy](#botcomponentreusepolicy), [BotSettingsBase](#botsettingsbase), [BotSpeechSettings](#botspeechsettings), [BotSpeechSettingsNoKind](#botspeechsettingsnokind), [ClassicVoiceProcessingMode](#classicvoiceprocessingmode), [ComputerUsingAgentAllowlist](#computerusingagentallowlist), [Connection](#connection), [ConnectionParameterSettings](#connectionparametersettings), [ConnectionProperties](#connectionproperties), [ConnectionPropertiesNoKind](#connectionpropertiesnokind), [ConnectionReferenceNameInFlowToConnectionParametersSetMappingNoKind](#connectionreferencenameinflowtoconnectionparameterssetmappingnokind), [ContentModerationLevel](#contentmoderationlevel), [ContentModerationSettings](#contentmoderationsettings), [ConversationHistory](#conversationhistory), [CopilotPolicy](#copilotpolicy), [CopilotPolicyNoKind](#copilotpolicynokind), [CurrentModels](#currentmodels), [CurrentModelsNoKind](#currentmodelsnokind), [CustomModels](#custommodels), [DefaultFeedbackCollection](#defaultfeedbackcollection), [DefaultHoldSettings](#defaultholdsettings), [DefaultHoldSettingsNoKind](#defaultholdsettingsnokind), [DeferredUpdateModels](#deferredupdatemodels), [DigitalListeningMode](#digitallisteningmode), [DigitalNoiseReductionMode](#digitalnoisereductionmode), [DigitalTurnDetectionType](#digitalturndetectiontype), [DigitalVoiceSettings](#digitalvoicesettings), [DtmfMappingKey](#dtmfmappingkey), [DtmfMultipleChoiceOptions](#dtmfmultiplechoiceoptions), [DtmfMultipleChoiceOptionsNoKind](#dtmfmultiplechoiceoptionsnokind), [DtmfTerminationKey](#dtmfterminationkey), [ExperimentalModels](#experimentalmodels), [FineTunedModelsPropertyDetails](#finetunedmodelspropertydetails), [FineTunedModelsPropertyDetailsNoKind](#finetunedmodelspropertydetailsnokind), [GenerationSettings](#generationsettings), [GenerationSettingsNoKind](#generationsettingsnokind), [GptCapabilities](#gptcapabilities), [GPTSettings](#gptsettings), [GPTSettingsNoKind](#gptsettingsnokind), [HistoryType](#historytype), [InboundTranscriptRedactionSettings](#inboundtranscriptredactionsettings), [InboundTranscriptRedactionSettingsNoKind](#inboundtranscriptredactionsettingsnokind), [InterruptionEagernessLevel](#interruptioneagernesslevel), [InterruptionPolicy](#interruptionpolicy), [InterruptionPolicyNoKind](#interruptionpolicynokind), [InterruptionTopicListFilter](#interruptiontopiclistfilter), [IvrBotSettings](#ivrbotsettings), [Language](#language), [LatencySettingsBase](#latencysettingsbase), [LatencySettingsBaseLegacy](#latencysettingsbaselegacy), [ModelApiType](#modelapitype), [ModelManifestResource](#modelmanifestresource), [ModelProvider](#modelprovider), [NetworkIsolationType](#networkisolationtype), [NetworkSettings](#networksettings), [NetworkSettingsNoKind](#networksettingsnokind), [NeverRequireMCPApproval](#neverrequiremcpapproval), [NoFeedbackCollection](#nofeedbackcollection), [NoHistory](#nohistory), [PreviewModels](#previewmodels), [QuestionInputSettings](#questioninputsettings), [RealtimeVoiceAudioLatencySettings](#realtimevoiceaudiolatencysettings), [RealTimeVoiceBotSettings](#realtimevoicebotsettings), [RealtimeVoiceMessageLatencySettings](#realtimevoicemessagelatencysettings), [RealtimeVoiceMessageLatencySettingsNoKind](#realtimevoicemessagelatencysettingsnokind), [RealtimeVoiceProcessingMode](#realtimevoiceprocessingmode), [ReasoningExperimentalModels](#reasoningexperimentalmodels), [ReasoningPreviewModels](#reasoningpreviewmodels), [ReasoningProductionModels](#reasoningproductionmodels), [RoiCalculationSettingsBase](#roicalculationsettingsbase), [SensitiveDataMaskingSettings](#sensitivedatamaskingsettings), [SensitivityLabelStatusCode](#sensitivitylabelstatuscode), [SensitivityLevel](#sensitivitylevel), [SessionBasedRoiCalculationSettings](#sessionbasedroicalculationsettings), [TenantFineTunedModelDetails](#tenantfinetunedmodeldetails), [TenantFineTunedModelDetailsNoKind](#tenantfinetunedmodeldetailsnokind), [ToneType](#tonetype), [ToolBasedRoiCalculationSettings](#toolbasedroicalculationsettings), [ToolRoiCalculationConfiguration](#toolroicalculationconfiguration), [ToolRoiCalculationConfigurationNoKind](#toolroicalculationconfigurationnokind), [TurboExperimentalModels](#turboexperimentalmodels), [TurboPreviewModels](#turbopreviewmodels), [TurboProductionModels](#turboproductionmodels), [TurnDetectionType](#turndetectiontype), [VoiceAudioLatencySettings](#voiceaudiolatencysettings), [VoiceInputSettings](#voiceinputsettings), [VoiceMessageLatencySettings](#voicemessagelatencysettings), [VoiceMessageLatencySettingsNoKind](#voicemessagelatencysettingsnokind), [VoiceProcessingModeBase](#voiceprocessingmodebase), [VoiceTypeBase](#voicetypebase), [VoiceTypeBaseNoKind](#voicetypebasenokind), [WasmBrowserTelemetryLogLevel](#wasmbrowsertelemetryloglevel)

</details>

### AdaptationManagementServiceModelProfile

YAML: `kind: AdaptationManagementServiceModelProfile`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `profileId` | `string` |  |  |  |
| `lastTrainedByUser` | `string` |  |  |  |
| `lastTrainedOnUtc` | `string` |  |  |  |
| `modelTypes` | map of name → [AdaptationManagementServiceModelTypeNoKind](#adaptationmanagementservicemodeltypenokind) |  |  |  |

_No other properties are allowed._

### AdaptationManagementServiceModelProfileNoKind

*Title:* AdaptationManagementServiceModelProfile  
Same shape as [AdaptationManagementServiceModelProfile](#adaptationmanagementservicemodelprofile) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `profileId` | `string` |  |  |  |
| `lastTrainedByUser` | `string` |  |  |  |
| `lastTrainedOnUtc` | `string` |  |  |  |
| `modelTypes` | map of name → [AdaptationManagementServiceModelTypeNoKind](#adaptationmanagementservicemodeltypenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotSynchronizationDetails](components.md#botsynchronizationdetails), [BotSynchronizationDetailsNoKind](components.md#botsynchronizationdetailsnokind)</sub>

### AdaptationManagementServiceModelType

YAML: `kind: AdaptationManagementServiceModelType`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `adaptationDetails` | map of name → [AdaptationManagementServiceTrainingMetadataNoKind](components.md#adaptationmanagementservicetrainingmetadatanokind) |  |  |  |

_No other properties are allowed._

### AdaptationManagementServiceModelTypeNoKind

*Title:* AdaptationManagementServiceModelType  
Same shape as [AdaptationManagementServiceModelType](#adaptationmanagementservicemodeltype) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `adaptationDetails` | map of name → [AdaptationManagementServiceTrainingMetadataNoKind](components.md#adaptationmanagementservicetrainingmetadatanokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AdaptationManagementServiceModelProfile](#adaptationmanagementservicemodelprofile), [AdaptationManagementServiceModelProfileNoKind](#adaptationmanagementservicemodelprofilenokind)</sub>

### AgentModelType


**One of:** [CustomModels](#custommodels), [CurrentModels](#currentmodels), [CurrentModelsNoKind](#currentmodelsnokind), [PreviewModels](#previewmodels), [ExperimentalModels](#experimentalmodels), [TurboProductionModels](#turboproductionmodels), [TurboPreviewModels](#turbopreviewmodels), [TurboExperimentalModels](#turboexperimentalmodels), [ReasoningProductionModels](#reasoningproductionmodels), [ReasoningPreviewModels](#reasoningpreviewmodels), [ReasoningExperimentalModels](#reasoningexperimentalmodels), [DeferredUpdateModels](#deferredupdatemodels)

<sub>Used by: [AISettings](#aisettings), [AISettingsNoKind](#aisettingsnokind), [InvokeComputerUsingAgentTaskAction](task-actions.md#invokecomputerusingagenttaskaction), [PromptGrader](evaluation.md#promptgrader)</sub>

### AISettings

YAML: `kind: AISettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `useModelKnowledge` | `boolean` |  | `false` |  |
| `isFileAnalysisEnabled` | `boolean` |  | `false` |  |
| `isSemanticSearchEnabled` | `boolean` |  | `true` |  |
| `contentModeration` | [ContentModerationLevel](#contentmoderationlevel) |  | `"High"` |  |
| `optInUseLatestModels` | `boolean` |  | `false` |  |
| `model` | [AgentModelType](#agentmodeltype) |  |  |  |

_No other properties are allowed._

### AISettingsNoKind

*Title:* AISettings  
Same shape as [AISettings](#aisettings) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `useModelKnowledge` | `boolean` |  | `false` |  |
| `isFileAnalysisEnabled` | `boolean` |  | `false` |  |
| `isSemanticSearchEnabled` | `boolean` |  | `true` |  |
| `contentModeration` | [ContentModerationLevel](#contentmoderationlevel) |  | `"High"` |  |
| `optInUseLatestModels` | `boolean` |  | `false` |  |
| `model` | [AgentModelType](#agentmodeltype) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotConfiguration](components.md#botconfiguration), [BotConfigurationNoKind](components.md#botconfigurationnokind), [GptComponentMetadata](components.md#gptcomponentmetadata), [GptComponentMetadataNoKind](components.md#gptcomponentmetadatanokind)</sub>

### AlwaysRequireMCPApproval

YAML: `kind: AlwaysRequireMCPApproval`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `tools` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

### AnalyticsSettings

YAML: `kind: AnalyticsSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `includeNodeLevelDetailsInTranscript` | `boolean` |  | `false` |  |
| `roiCalculation` | [RoiCalculationSettingsBase](#roicalculationsettingsbase) |  |  |  |
| `includeTestDataInAnalytics` | `boolean` |  | `false` |  |

_No other properties are allowed._

### AnalyticsSettingsNoKind

*Title:* AnalyticsSettings  
Same shape as [AnalyticsSettings](#analyticssettings) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `includeNodeLevelDetailsInTranscript` | `boolean` |  | `false` |  |
| `roiCalculation` | [RoiCalculationSettingsBase](#roicalculationsettingsbase) |  |  |  |
| `includeTestDataInAnalytics` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [BotConfiguration](components.md#botconfiguration), [BotConfigurationNoKind](components.md#botconfigurationnokind)</sub>

### AutonomousSettings

YAML: `kind: AutonomousSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `enterpriseKnowledgeConnectionReference` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotSettingsBase](#botsettingsbase)</sub>

### BotComponentReusePolicy


**String enum.** Allowed values: `None`, `Private`, `Public`

<sub>Used by: [ContentShareContext](messages.md#contentsharecontext), [ContentShareContextNoKind](components.md#contentsharecontextnokind)</sub>

### BotSettingsBase


**One of:** [AutonomousSettings](#autonomoussettings), [ContentModerationSettings](#contentmoderationsettings), [ConnectionParameterSettings](#connectionparametersettings), [IvrBotSettings](#ivrbotsettings), [DigitalVoiceSettings](#digitalvoicesettings), [RealTimeVoiceBotSettings](#realtimevoicebotsettings), [NoFeedbackCollection](#nofeedbackcollection), [DefaultFeedbackCollection](#defaultfeedbackcollection)

<sub>Used by: [BotSettingsComponent](components.md#botsettingscomponent)</sub>

### BotSpeechSettings

YAML: `kind: BotSpeechSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `interdigitTimeoutInMilliseconds` | `number` |  | `3000` |  |
| `terminationTimeoutInMilliseconds` | `number` |  | `2000` |  |
| `isDtmfSettingsEnabled` | `boolean` |  | `true` |  |
| `dtmfTerminationKey` | `=` expression \| [DtmfTerminationKey](#dtmfterminationkey) |  |  | accepts a `=` Power Fx expression |
| `silenceDetectionTimeoutInMilliseconds` | `number` |  | `0` |  |
| `utteranceEndTimeoutInMilliseconds` | `number` |  | `1500` |  |
| `speechRecognitionTimeoutInMilliseconds` | `number` |  | `12000` |  |
| `latencyMessage` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |
| `latencyMessageDelayInMilliseconds` | `number` |  | `500` |  |
| `minimumPlaybackInMilliseconds` | `number` |  | `5000` |  |
| `isDomainLanguageModelEnabled` | `boolean` |  | `true` |  |
| `domainLanguageModelWeightValue` | `number` |  | `0.2` |  |
| `wordsetsWeightValue` | `number` |  | `0.1` |  |
| `speechRecognitionSensitivity` | `number` |  | `0.5` |  |
| `speechRecognitionSpeedVsAccuracy` | `number` |  | `0.5` |  |
| `autoAddLanguageTagInSSML` | `boolean` |  | `false` |  |
| `prefixPaddingInMilliseconds` | `number` |  | `300` |  |
| `audioThreshold` | `number` |  | `0.5` |  |
| `speechEndTimeoutInMilliseconds` | `number` |  | `500` |  |

_No other properties are allowed._

### BotSpeechSettingsNoKind

*Title:* BotSpeechSettings  
Same shape as [BotSpeechSettings](#botspeechsettings) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `interdigitTimeoutInMilliseconds` | `number` |  | `3000` |  |
| `terminationTimeoutInMilliseconds` | `number` |  | `2000` |  |
| `isDtmfSettingsEnabled` | `boolean` |  | `true` |  |
| `dtmfTerminationKey` | `=` expression \| [DtmfTerminationKey](#dtmfterminationkey) |  |  | accepts a `=` Power Fx expression |
| `silenceDetectionTimeoutInMilliseconds` | `number` |  | `0` |  |
| `utteranceEndTimeoutInMilliseconds` | `number` |  | `1500` |  |
| `speechRecognitionTimeoutInMilliseconds` | `number` |  | `12000` |  |
| `latencyMessage` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |
| `latencyMessageDelayInMilliseconds` | `number` |  | `500` |  |
| `minimumPlaybackInMilliseconds` | `number` |  | `5000` |  |
| `isDomainLanguageModelEnabled` | `boolean` |  | `true` |  |
| `domainLanguageModelWeightValue` | `number` |  | `0.2` |  |
| `wordsetsWeightValue` | `number` |  | `0.1` |  |
| `speechRecognitionSensitivity` | `number` |  | `0.5` |  |
| `speechRecognitionSpeedVsAccuracy` | `number` |  | `0.5` |  |
| `autoAddLanguageTagInSSML` | `boolean` |  | `false` |  |
| `prefixPaddingInMilliseconds` | `number` |  | `300` |  |
| `audioThreshold` | `number` |  | `0.5` |  |
| `speechEndTimeoutInMilliseconds` | `number` |  | `500` |  |

_No other properties are allowed._

<sub>Used by: [BotConfiguration](components.md#botconfiguration), [BotConfigurationNoKind](components.md#botconfigurationnokind)</sub>

### ClassicVoiceProcessingMode

YAML: `kind: ClassicVoiceProcessingMode`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [VoiceProcessingModeBase](#voiceprocessingmodebase)</sub>

### ComputerUsingAgentAllowlist

YAML: `kind: ComputerUsingAgentAllowlist`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `websites` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `desktopApps` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

### Connection


**One of:** [ConnectionProperties](#connectionproperties), [ConnectionPropertiesNoKind](#connectionpropertiesnokind)

<sub>Used by: [CurrentModels](#currentmodels), [CurrentModelsNoKind](#currentmodelsnokind)</sub>

### ConnectionParameterSettings

YAML: `kind: ConnectionParameterSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectorSettings` | map of name → [ConnectionParametersSetNoKind](components.md#connectionparameterssetnokind) |  |  |  |
| `flowSettings` | map of name → [ConnectionReferenceNameInFlowToConnectionParametersSetMappingNoKind](#connectionreferencenameinflowtoconnectionparameterssetmappingnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotSettingsBase](#botsettingsbase)</sub>

### ConnectionProperties

YAML: `kind: ConnectionProperties`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `usageDescription` | `string` |  |  |  |
| `name` | `string` |  |  |  |
| `mode` | [ConnectionMode](components.md#connectionmode) |  | `"Maker"` |  |

_No other properties are allowed._

<sub>Used by: [Connection](#connection)</sub>

### ConnectionPropertiesNoKind

*Title:* ConnectionProperties  
Same shape as [ConnectionProperties](#connectionproperties) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `usageDescription` | `string` |  |  |  |
| `name` | `string` |  |  |  |
| `mode` | [ConnectionMode](components.md#connectionmode) |  | `"Maker"` |  |

_No other properties are allowed._

<sub>Used by: [Connection](#connection), [ConnectorTriggerDefinition](components.md#connectortriggerdefinition), [ContinueConversationConnectorTrigger](triggers.md#continueconversationconnectortrigger), [CreateConversationConnectorTrigger](triggers.md#createconversationconnectortrigger), [FederatedStructuredSearchSource](knowledge.md#federatedstructuredsearchsource), [InvokeComputerUsingAgentTaskAction](task-actions.md#invokecomputerusingagenttaskaction), [InvokeConnectorAction](actions.md#invokeconnectoraction), [InvokeConnectorTaskAction](task-actions.md#invokeconnectortaskaction), [InvokeExternalAgentTaskAction](task-actions.md#invokeexternalagenttaskaction), [WaitForConnectorTrigger](actions.md#waitforconnectortrigger)</sub>

### ConnectionReferenceNameInFlowToConnectionParametersSetMappingNoKind

*Title:* ConnectionReferenceNameInFlowToConnectionParametersSetMapping  
Same shape as [ConnectionReferenceNameInFlowToConnectionParametersSetMapping](components.md#connectionreferencenameinflowtoconnectionparameterssetmapping) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionParameters` | map of name → [ConnectionParametersSetNoKind](components.md#connectionparameterssetnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ConnectionParameterSettings](#connectionparametersettings)</sub>

### ContentModerationLevel


**String enum.** Allowed values: `Minimum`, `Low`, `Medium`, `High`, `Maximum`

<sub>Used by: [AISettings](#aisettings), [AISettingsNoKind](#aisettingsnokind), [AnswerQuestionWithAI](actions.md#answerquestionwithai), [CreateSearchQuery](actions.md#createsearchquery), [SearchAndSummarizeContent](actions.md#searchandsummarizecontent), [SearchKnowledgeSources](actions.md#searchknowledgesources)</sub>

### ContentModerationSettings

YAML: `kind: ContentModerationSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `moderatedCustomResponse` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotSettingsBase](#botsettingsbase)</sub>

### ConversationHistory

YAML: `kind: ConversationHistory`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `numberOfPastUserMessagesToInclude` | [IntExpression](other.md#intexpression) |  | `5` | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [HistoryType](#historytype)</sub>

### CopilotPolicy

YAML: `kind: CopilotPolicy`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `allowGpt` | `boolean` |  | `true` |  |
| `allowBing` | `boolean` |  | `true` |  |
| `allowMakerCredentials` | `boolean` |  | `true` |  |

_No other properties are allowed._

### CopilotPolicyNoKind

*Title:* CopilotPolicy  
Same shape as [CopilotPolicy](#copilotpolicy) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `allowGpt` | `boolean` |  | `true` |  |
| `allowBing` | `boolean` |  | `true` |  |
| `allowMakerCredentials` | `boolean` |  | `true` |  |

_No other properties are allowed._

### CurrentModels

YAML: `kind: CurrentModels`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelNameHint` | `string` |  |  |  |
| `summarization` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |
| `apiType` | [ModelApiType](#modelapitype) |  | `"Chat"` |  |
| `provider` | [ModelProvider](#modelprovider) |  |  |  |
| `connection` | [Connection](#connection) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentModelType](#agentmodeltype)</sub>

### CurrentModelsNoKind

*Title:* CurrentModels  
Same shape as [CurrentModels](#currentmodels) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelNameHint` | `string` |  |  |  |
| `summarization` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |
| `apiType` | [ModelApiType](#modelapitype) |  | `"Chat"` |  |
| `provider` | [ModelProvider](#modelprovider) |  |  |  |
| `connection` | [Connection](#connection) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentModelType](#agentmodeltype)</sub>

### CustomModels

YAML: `kind: CustomModels`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelNameHint` | `string` |  |  |  |
| `summarization` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentModelType](#agentmodeltype)</sub>

### DefaultFeedbackCollection

YAML: `kind: DefaultFeedbackCollection`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `disclaimer` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotSettingsBase](#botsettingsbase)</sub>

### DefaultHoldSettings

YAML: `kind: DefaultHoldSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `holdWords` | list of `string` |  |  |  |
| `holdMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `resumeWords` | list of `string` |  |  |  |
| `holdMusic` | `string` |  | `"string.Empty"` |  |
| `holdTimeoutInMilliseconds` | `number` |  | `15000` |  |
| `holdRepeatCount` | `number` |  | `2` |  |
| `holdTimeoutMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `resumeResponse` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

### DefaultHoldSettingsNoKind

Same shape as [DefaultHoldSettings](#defaultholdsettings) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `holdWords` | list of `string` |  |  |  |
| `holdMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `resumeWords` | list of `string` |  |  |  |
| `holdMusic` | `string` |  | `"string.Empty"` |  |
| `holdTimeoutInMilliseconds` | `number` |  | `15000` |  |
| `holdRepeatCount` | `number` |  | `2` |  |
| `holdTimeoutMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `resumeResponse` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [IvrBotSettings](#ivrbotsettings)</sub>

### DeferredUpdateModels

YAML: `kind: DeferredUpdateModels`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelNameHint` | `string` |  |  |  |
| `summarization` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentModelType](#agentmodeltype)</sub>

### DigitalListeningMode


**String enum.** Allowed values: `AlwaysListening`, `PushToTalk`

<sub>Used by: [DigitalVoiceSettings](#digitalvoicesettings)</sub>

### DigitalNoiseReductionMode


**String enum.** Allowed values: `None`, `DeepNoiseSuppression`

<sub>Used by: [DigitalVoiceSettings](#digitalvoicesettings)</sub>

### DigitalTurnDetectionType


**String enum.** Allowed values: `None`, `SilenceBasedVoiceActivityDetection`, `SemanticVoiceActivityDetection`

<sub>Used by: [DigitalVoiceSettings](#digitalvoicesettings)</sub>

### DigitalVoiceSettings

YAML: `kind: DigitalVoiceSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `voiceFont` | [VoiceTypeBaseNoKind](#voicetypebasenokind) |  |  |  |
| `turnDetection` | `=` expression \| [DigitalTurnDetectionType](#digitalturndetectiontype) |  | `"SemanticVoiceActivityDetection"` | accepts a `=` Power Fx expression |
| `prefixPaddingInMilliseconds` | `number` |  | `300` |  |
| `silenceDurationInMilliseconds` | `number` |  | `500` |  |
| `threshold` | `number` |  | `0.5` |  |
| `removeFillerWords` | `boolean` |  | `true` |  |
| `noiseReduction` | `=` expression \| [DigitalNoiseReductionMode](#digitalnoisereductionmode) |  | `"DeepNoiseSuppression"` | accepts a `=` Power Fx expression |
| `echoCancellation` | `boolean` |  | `true` |  |
| `listeningMode` | `=` expression \| [DigitalListeningMode](#digitallisteningmode) |  | `"AlwaysListening"` | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [BotSettingsBase](#botsettingsbase)</sub>

### DtmfMappingKey


**String enum.** Allowed values: `Num0`, `Num1`, `Num2`, `Num3`, `Num4`, `Num5`, `Num6`, `Num7`, `Num8`, `Num9`, `PoundKey`, `StarKey`

<sub>Used by: [ClosedListItem](entities.md#closedlistitem), [ClosedListItemNoKind](dialogs.md#closedlistitemnokind), [OnDtmfKeyPress](triggers.md#ondtmfkeypress), [SelectedClosedListItem](entities.md#selectedclosedlistitem), [SelectedClosedListItemNoKind](entities.md#selectedclosedlistitemnokind), [SelectedClosedListItemOptionalKind](entities.md#selectedclosedlistitemoptionalkind)</sub>

### DtmfMultipleChoiceOptions

YAML: `kind: DtmfMultipleChoiceOptions`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `generateMapping` | [BoolExpression](other.md#boolexpression) |  | `false` | accepts a `=` Power Fx expression |
| `readOutOptions` | `true` \| `false` \| [BoolExpression](other.md#boolexpression) |  | `true` |  |

_No other properties are allowed._

### DtmfMultipleChoiceOptionsNoKind

*Title:* DtmfMultipleChoiceOptions  
Same shape as [DtmfMultipleChoiceOptions](#dtmfmultiplechoiceoptions) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `generateMapping` | [BoolExpression](other.md#boolexpression) |  | `false` | accepts a `=` Power Fx expression |
| `readOutOptions` | `true` \| `false` \| [BoolExpression](other.md#boolexpression) |  | `true` |  |

_No other properties are allowed._

<sub>Used by: [BooleanPrebuiltEntity](entities.md#booleanprebuiltentity), [ClosedListEntity](dialogs.md#closedlistentity), [ClosedListEntityReference](entities.md#closedlistentityreference), [DynamicClosedListEntity](entities.md#dynamicclosedlistentity)</sub>

### DtmfTerminationKey


**String enum.** Allowed values: `PoundKey`, `StarKey`

<sub>Used by: [BotSpeechSettings](#botspeechsettings), [BotSpeechSettingsNoKind](#botspeechsettingsnokind), [FixedLengthTermination](actions.md#fixedlengthtermination), [FlexibleLengthTermination](actions.md#flexiblelengthtermination), [KeyBasedTermination](actions.md#keybasedtermination)</sub>

### ExperimentalModels

YAML: `kind: ExperimentalModels`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelNameHint` | `string` |  |  |  |
| `summarization` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentModelType](#agentmodeltype)</sub>

### FineTunedModelsPropertyDetails

YAML: `kind: FineTunedModelsPropertyDetails`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `tenantFineTunedModelDetails` | list of [TenantFineTunedModelDetailsNoKind](#tenantfinetunedmodeldetailsnokind) |  |  |  |

_No other properties are allowed._

### FineTunedModelsPropertyDetailsNoKind

Same shape as [FineTunedModelsPropertyDetails](#finetunedmodelspropertydetails) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `tenantFineTunedModelDetails` | list of [TenantFineTunedModelDetailsNoKind](#tenantfinetunedmodeldetailsnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [GptCapabilities](#gptcapabilities), [GptCapabilitiesNoKind](components.md#gptcapabilitiesnokind)</sub>

### GenerationSettings

YAML: `kind: GenerationSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

### GenerationSettingsNoKind

*Title:* GenerationSettings  

**Type:** object

<sub>Used by: [BotConfiguration](components.md#botconfiguration), [BotConfigurationNoKind](components.md#botconfigurationnokind)</sub>

### GptCapabilities

YAML: `kind: GptCapabilities`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `webBrowsing` | `boolean` |  | `false` |  |
| `codeInterpreter` | `boolean` |  | `false` |  |
| `generateImages` | `boolean` |  | `false` |  |
| `fineTunedModels` | [FineTunedModelsPropertyDetailsNoKind](#finetunedmodelspropertydetailsnokind) |  |  |  |
| `searchTeams` | `boolean` |  | `false` |  |
| `searchOneDriveAndSharePoint` | `boolean` |  | `false` |  |
| `searchEmails` | `boolean` |  | `false` |  |
| `searchMeetings` | `boolean` |  | `false` |  |
| `searchPeople` | `boolean` |  | `false` |  |
| `searchPeopleIncludeRelatedContent` | `boolean` |  | `false` |  |
| `customExperience` | `string` |  |  |  |

_No other properties are allowed._

### GPTSettings

YAML: `kind: GPTSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `defaultSchemaName` | `string` |  |  |  |

_No other properties are allowed._

### GPTSettingsNoKind

*Title:* GPTSettings  
Same shape as [GPTSettings](#gptsettings) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `defaultSchemaName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotConfiguration](components.md#botconfiguration), [BotConfigurationNoKind](components.md#botconfigurationnokind)</sub>

### HistoryType


**One of:** [NoHistory](#nohistory), [ConversationHistory](#conversationhistory)

<sub>Used by: [AnswerQuestionWithAI](actions.md#answerquestionwithai), [CreateSearchQuery](actions.md#createsearchquery), [ExternalChatCompletionModelConfiguration](actions.md#externalchatcompletionmodelconfiguration), [ExternalChatCompletionModelConfigurationNoKind](actions.md#externalchatcompletionmodelconfigurationnokind), [GptComponentMetadata](components.md#gptcomponentmetadata), [GptComponentMetadataNoKind](components.md#gptcomponentmetadatanokind), [InvokeConnectedAgentTaskAction](task-actions.md#invokeconnectedagenttaskaction), [InvokeExternalAgentTaskAction](task-actions.md#invokeexternalagenttaskaction), [SearchAndSummarizeContent](actions.md#searchandsummarizecontent), [SearchAndSummarizeWithCustomModel](actions.md#searchandsummarizewithcustommodel), [SearchKnowledgeSources](actions.md#searchknowledgesources)</sub>

### InboundTranscriptRedactionSettings

YAML: `kind: InboundTranscriptRedactionSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `isRedactCreditCardEnabled` | `boolean` |  | `false` |  |
| `isRedactSsnEnabled` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [SensitiveDataMaskingSettings](#sensitivedatamaskingsettings)</sub>

### InboundTranscriptRedactionSettingsNoKind

*Title:* InboundTranscriptRedactionSettings  
Same shape as [InboundTranscriptRedactionSettings](#inboundtranscriptredactionsettings) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `isRedactCreditCardEnabled` | `boolean` |  | `false` |  |
| `isRedactSsnEnabled` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [SensitiveDataMaskingSettings](#sensitivedatamaskingsettings)</sub>

### InterruptionEagernessLevel


**String enum.** Allowed values: `Low`, `Medium`, `High`

<sub>Used by: [IvrBotSettings](#ivrbotsettings)</sub>

### InterruptionPolicy

YAML: `kind: InterruptionPolicy`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `allowInterruption` | `boolean` |  | `true` |  |
| `interruptionTopicListFilter` | [InterruptionTopicListFilter](#interruptiontopiclistfilter) |  |  |  |

_No other properties are allowed._

### InterruptionPolicyNoKind

*Title:* InterruptionPolicy  
Same shape as [InterruptionPolicy](#interruptionpolicy) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `allowInterruption` | `boolean` |  | `true` |  |
| `interruptionTopicListFilter` | [InterruptionTopicListFilter](#interruptiontopiclistfilter) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ActivityProtocolMetadata](components.md#activityprotocolmetadata), [AdaptiveCardPrompt](actions.md#adaptivecardprompt), [AgentToAgentProtocolMetadata](components.md#agenttoagentprotocolmetadata), [InvokeConnectedAgentTaskAction](task-actions.md#invokeconnectedagenttaskaction), [InvokeSkillAction](actions.md#invokeskillaction), [InvokeSkillTaskAction](task-actions.md#invokeskilltaskaction), [Question](actions.md#question)</sub>

### InterruptionTopicListFilter


**One of:** [IncludeSelectedTopics](actions.md#includeselectedtopics)

<sub>Used by: [InterruptionPolicy](#interruptionpolicy), [InterruptionPolicyNoKind](#interruptionpolicynokind)</sub>

### IvrBotSettings

YAML: `kind: IvrBotSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `domainLanguageModelWeightValue` | `number` |  |  |  |
| `holdSettings` | [DefaultHoldSettingsNoKind](#defaultholdsettingsnokind) |  |  |  |
| `latencySettings` | [LatencySettingsBase](#latencysettingsbase) |  |  |  |
| `voiceFont` | [VoiceTypeBaseNoKind](#voicetypebasenokind) |  |  |  |
| `turnDetection` | `=` expression \| [TurnDetectionType](#turndetectiontype) |  | `"SilenceBasedVoiceActivityDetection"` | accepts a `=` Power Fx expression |
| `interruptionEagerness` | `=` expression \| [InterruptionEagernessLevel](#interruptioneagernesslevel) |  | `"Medium"` | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [BotSettingsBase](#botsettingsbase)</sub>

### Language


**One of:** `1025`, `2052`, `1028`, `1029`, `1030`, `1043`, `1033`, `1035`, `1036`, `1031`, `1032`, `1081`, `1057`, `1040`, `1041`, `1042`, `1044`, `1045`, `1046`, `1049`, `1034`, `1054`, `1053`, `1055`, `2057`, `3084`, `21514`, `3081`, `1058`, `1037`, `2070`, `20490`

<sub>Used by: [BotComponentCollectionConfiguration](components.md#botcomponentcollectionconfiguration), [BotComponentCollectionConfigurationNoKind](components.md#botcomponentcollectionconfigurationnokind), [BotEntity](components.md#botentity), [BotEntityNoKind](components.md#botentitynokind), [LegacyOrUnknownComponent](components.md#legacyorunknowncomponent), [TranslationsComponent](components.md#translationscomponent), [TranslationsComponentNoKind](components.md#translationscomponentnokind)</sub>

### LatencySettingsBase


**One of:** [VoiceMessageLatencySettings](#voicemessagelatencysettings), [VoiceMessageLatencySettingsNoKind](#voicemessagelatencysettingsnokind), [VoiceAudioLatencySettings](#voiceaudiolatencysettings)

<sub>Used by: [IvrBotSettings](#ivrbotsettings)</sub>

### LatencySettingsBaseLegacy


**One of:** [RealtimeVoiceMessageLatencySettings](#realtimevoicemessagelatencysettings), [RealtimeVoiceMessageLatencySettingsNoKind](#realtimevoicemessagelatencysettingsnokind), [RealtimeVoiceAudioLatencySettings](#realtimevoiceaudiolatencysettings)

<sub>Used by: [RealTimeVoiceBotSettings](#realtimevoicebotsettings)</sub>

### ModelApiType


**String enum.** Allowed values: `Chat`, `Responses`

<sub>Used by: [CurrentModels](#currentmodels), [CurrentModelsNoKind](#currentmodelsnokind)</sub>

### ModelManifestResource

YAML: `kind: ModelManifestResource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [ManifestResources](components.md#manifestresources)</sub>

### ModelProvider


**String enum.** Allowed values: `OpenAI`, `AzureOpenAI`, `Anthropic`, `GoogleVertexAI`

<sub>Used by: [CurrentModels](#currentmodels), [CurrentModelsNoKind](#currentmodelsnokind)</sub>

### NetworkIsolationType


**String enum.** Allowed values: `PublicNetwork`, `VirtualNetwork`

<sub>Used by: [HttpRequestAction](actions.md#httprequestaction), [NetworkSettings](#networksettings), [NetworkSettingsNoKind](#networksettingsnokind)</sub>

### NetworkSettings

YAML: `kind: NetworkSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `networkIsolation` | [NetworkIsolationType](#networkisolationtype) |  | `"PublicNetwork"` |  |

_No other properties are allowed._

### NetworkSettingsNoKind

*Title:* NetworkSettings  
Same shape as [NetworkSettings](#networksettings) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `networkIsolation` | [NetworkIsolationType](#networkisolationtype) |  | `"PublicNetwork"` |  |

_No other properties are allowed._

<sub>Used by: [BotConfiguration](components.md#botconfiguration), [BotConfigurationNoKind](components.md#botconfigurationnokind)</sub>

### NeverRequireMCPApproval

YAML: `kind: NeverRequireMCPApproval`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `tools` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

### NoFeedbackCollection

YAML: `kind: NoFeedbackCollection`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [BotSettingsBase](#botsettingsbase)</sub>

### NoHistory

YAML: `kind: NoHistory`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [HistoryType](#historytype)</sub>

### PreviewModels

YAML: `kind: PreviewModels`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelNameHint` | `string` |  |  |  |
| `summarization` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentModelType](#agentmodeltype)</sub>

### QuestionInputSettings

YAML: `kind: QuestionInputSettings`

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

### RealtimeVoiceAudioLatencySettings

YAML: `kind: RealtimeVoiceAudioLatencySettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `latencyMessageDelayInMilliseconds` | `number` |  |  |  |
| `minimumPlaybackInMilliseconds` | `number` |  |  |  |
| `latencyAudio` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `fallbackLatencyMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [LatencySettingsBaseLegacy](#latencysettingsbaselegacy)</sub>

### RealTimeVoiceBotSettings

YAML: `kind: RealTimeVoiceBotSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `latencySettings` | [LatencySettingsBaseLegacy](#latencysettingsbaselegacy) |  |  |  |
| `voiceType` | [VoiceTypeBaseNoKind](#voicetypebasenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotSettingsBase](#botsettingsbase)</sub>

### RealtimeVoiceMessageLatencySettings

YAML: `kind: RealtimeVoiceMessageLatencySettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `latencyMessageDelayInMilliseconds` | `number` |  |  |  |
| `minimumPlaybackInMilliseconds` | `number` |  |  |  |
| `latencyMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [LatencySettingsBaseLegacy](#latencysettingsbaselegacy)</sub>

### RealtimeVoiceMessageLatencySettingsNoKind

Same shape as [RealtimeVoiceMessageLatencySettings](#realtimevoicemessagelatencysettings) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `latencyMessageDelayInMilliseconds` | `number` |  |  |  |
| `minimumPlaybackInMilliseconds` | `number` |  |  |  |
| `latencyMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [LatencySettingsBaseLegacy](#latencysettingsbaselegacy)</sub>

### RealtimeVoiceProcessingMode

YAML: `kind: RealtimeVoiceProcessingMode`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [VoiceProcessingModeBase](#voiceprocessingmodebase)</sub>

### ReasoningExperimentalModels

YAML: `kind: ReasoningExperimentalModels`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelNameHint` | `string` |  |  |  |
| `summarization` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentModelType](#agentmodeltype)</sub>

### ReasoningPreviewModels

YAML: `kind: ReasoningPreviewModels`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelNameHint` | `string` |  |  |  |
| `summarization` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentModelType](#agentmodeltype)</sub>

### ReasoningProductionModels

YAML: `kind: ReasoningProductionModels`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelNameHint` | `string` |  |  |  |
| `summarization` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentModelType](#agentmodeltype)</sub>

### RoiCalculationSettingsBase


**One of:** [ToolBasedRoiCalculationSettings](#toolbasedroicalculationsettings), [SessionBasedRoiCalculationSettings](#sessionbasedroicalculationsettings)

<sub>Used by: [AnalyticsSettings](#analyticssettings), [AnalyticsSettingsNoKind](#analyticssettingsnokind)</sub>

### SensitiveDataMaskingSettings


**One of:** [InboundTranscriptRedactionSettings](#inboundtranscriptredactionsettings), [InboundTranscriptRedactionSettingsNoKind](#inboundtranscriptredactionsettingsnokind)

<sub>Used by: [BotConfiguration](components.md#botconfiguration), [BotConfigurationNoKind](components.md#botconfigurationnokind)</sub>

### SensitivityLabelStatusCode


**String enum.** Allowed values: `LabelFound`, `LabelNotFound`, `UnsupportedMediaType`, `NotSupportedDueToEncryption`, `NotSupportedOther`, `NotAuthorized`

<sub>Used by: [GraphDriveItemSyncDetails](knowledge.md#graphdriveitemsyncdetails), [GraphDriveItemSyncDetailsNoKind](knowledge.md#graphdriveitemsyncdetailsnokind)</sub>

### SensitivityLevel


**String enum.** Allowed values: `None`, `Sensitive`

<sub>Used by: [AgePrebuiltEntity](entities.md#ageprebuiltentity), [Any](data-types.md#any), [Blank](data-types.md#blank), [Boolean](data-types.md#boolean), [BooleanPrebuiltEntity](entities.md#booleanprebuiltentity), [CityPrebuiltEntity](entities.md#cityprebuiltentity), [ClosedListEntityReference](entities.md#closedlistentityreference), [ClosedListOptionSet](data-types.md#closedlistoptionset), [Color](data-types.md#color), [ColorPrebuiltEntity](entities.md#colorprebuiltentity), [ContinentPrebuiltEntity](entities.md#continentprebuiltentity), [CountryOrRegionPrebuiltEntity](entities.md#countryorregionprebuiltentity), [Date](data-types.md#date), [DatePrebuiltEntity](entities.md#dateprebuiltentity), [DateTime](data-types.md#datetime), [DateTimeNoTimeZone](data-types.md#datetimenotimezone), [DateTimeNoTimeZonePrebuiltEntity](entities.md#datetimenotimezoneprebuiltentity), [DateTimePrebuiltEntity](entities.md#datetimeprebuiltentity), [DurationPrebuiltEntity](entities.md#durationprebuiltentity), [DynamicClosedListEntity](entities.md#dynamicclosedlistentity) … and 41 more</sub>

### SessionBasedRoiCalculationSettings

YAML: `kind: SessionBasedRoiCalculationSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `currencyCode` | `string` |  |  |  |
| `savedCost` | `number` |  |  |  |
| `savedTimeInSeconds` | `number` |  |  |  |

_No other properties are allowed._

<sub>Used by: [RoiCalculationSettingsBase](#roicalculationsettingsbase)</sub>

### TenantFineTunedModelDetails

YAML: `kind: TenantFineTunedModelDetails`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelId` | `string` |  |  |  |
| `scenarioName` | `string` |  |  |  |

_No other properties are allowed._

### TenantFineTunedModelDetailsNoKind

Same shape as [TenantFineTunedModelDetails](#tenantfinetunedmodeldetails) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelId` | `string` |  |  |  |
| `scenarioName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [FineTunedModelsPropertyDetails](#finetunedmodelspropertydetails), [FineTunedModelsPropertyDetailsNoKind](#finetunedmodelspropertydetailsnokind)</sub>

### ToneType


**String enum.** Allowed values: `Friendly`, `Confident`, `Empathetic`, `Conversational`, `Professional`, `Motivational`

### ToolBasedRoiCalculationSettings

YAML: `kind: ToolBasedRoiCalculationSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `currencyCode` | `string` |  |  |  |
| `tools` | list of [ToolRoiCalculationConfigurationNoKind](#toolroicalculationconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [RoiCalculationSettingsBase](#roicalculationsettingsbase)</sub>

### ToolRoiCalculationConfiguration

YAML: `kind: ToolRoiCalculationConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `schemaName` | `string` | ✓ |  |  |
| `savedCost` | `number` |  |  |  |
| `savedTimeInSeconds` | `number` |  |  |  |

_No other properties are allowed._

### ToolRoiCalculationConfigurationNoKind

*Title:* ToolRoiCalculationConfiguration  
Same shape as [ToolRoiCalculationConfiguration](#toolroicalculationconfiguration) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `schemaName` | `string` | ✓ |  |  |
| `savedCost` | `number` |  |  |  |
| `savedTimeInSeconds` | `number` |  |  |  |

_No other properties are allowed._

<sub>Used by: [ToolBasedRoiCalculationSettings](#toolbasedroicalculationsettings)</sub>

### TurboExperimentalModels

YAML: `kind: TurboExperimentalModels`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelNameHint` | `string` |  |  |  |
| `summarization` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentModelType](#agentmodeltype)</sub>

### TurboPreviewModels

YAML: `kind: TurboPreviewModels`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelNameHint` | `string` |  |  |  |
| `summarization` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentModelType](#agentmodeltype)</sub>

### TurboProductionModels

YAML: `kind: TurboProductionModels`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelNameHint` | `string` |  |  |  |
| `summarization` | [SummarizationExternalModelConfigurationNoKind](knowledge.md#summarizationexternalmodelconfigurationnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AgentModelType](#agentmodeltype)</sub>

### TurnDetectionType


**String enum.** Allowed values: `SilenceBasedVoiceActivityDetection`, `SemanticVoiceActivityDetection`

<sub>Used by: [IvrBotSettings](#ivrbotsettings)</sub>

### VoiceAudioLatencySettings

YAML: `kind: VoiceAudioLatencySettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `latencyAudio` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `fallbackLatencyMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [LatencySettingsBase](#latencysettingsbase)</sub>

### VoiceInputSettings

YAML: `kind: VoiceInputSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `wordsetsWeightValue` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `speechRecognitionSensitivity` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `speechRecognitionSpeedVsAccuracy` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `silenceDetectionTimeoutInMilliseconds` | `0` \| [IntExpression](other.md#intexpression) |  |  |  |
| `repeatCountOnSilence` | `0` \| `1` \| `2` \| [IntExpression](other.md#intexpression) |  | `2` |  |
| `inputTimeoutResponse` | [ActivityTemplateBase](messages.md#activitytemplatebase) |  |  |  |
| `fallbackDialogOnSilence` | [DialogExpression](other.md#dialogexpression) |  |  | accepts a `=` Power Fx expression |
| `defaultValueMissingAction` | `=` expression \| [DefaultValueMissingAction](actions.md#defaultvaluemissingaction) |  | `"HangUp"` | accepts a `=` Power Fx expression |
| `defaultValueOnSilence` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `utteranceEndTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `speechRecognitionTimeoutInMilliseconds` | `0` \| [IntExpression](other.md#intexpression) |  |  |  |

_No other properties are allowed._

### VoiceMessageLatencySettings

YAML: `kind: VoiceMessageLatencySettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `latencyMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [LatencySettingsBase](#latencysettingsbase)</sub>

### VoiceMessageLatencySettingsNoKind

Same shape as [VoiceMessageLatencySettings](#voicemessagelatencysettings) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `latencyMessage` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [LatencySettingsBase](#latencysettingsbase)</sub>

### VoiceProcessingModeBase


**One of:** [ClassicVoiceProcessingMode](#classicvoiceprocessingmode), [RealtimeVoiceProcessingMode](#realtimevoiceprocessingmode)

<sub>Used by: [BotConfiguration](components.md#botconfiguration), [BotConfigurationNoKind](components.md#botconfigurationnokind)</sub>

### VoiceTypeBase

YAML: `kind: VoiceTypeBase`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `voiceName` | `string` |  |  |  |
| `speakingSpeed` | `0.25` \| `0.5` \| `0.75` \| `1` \| `1.25` \| `1.5` \| [NumberExpression](other.md#numberexpression) |  | `1` |  |

_No other properties are allowed._

### VoiceTypeBaseNoKind

*Title:* VoiceTypeBase  
Same shape as [VoiceTypeBase](#voicetypebase) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `voiceName` | `string` |  |  |  |
| `speakingSpeed` | `0.25` \| `0.5` \| `0.75` \| `1` \| `1.25` \| `1.5` \| [NumberExpression](other.md#numberexpression) |  | `1` |  |

_No other properties are allowed._

<sub>Used by: [DigitalVoiceSettings](#digitalvoicesettings), [IvrBotSettings](#ivrbotsettings), [RealTimeVoiceBotSettings](#realtimevoicebotsettings)</sub>

### WasmBrowserTelemetryLogLevel


**String enum.** Allowed values: `Critical`, `Error`, `Warning`, `Information`, `Trace`, `Debug`
