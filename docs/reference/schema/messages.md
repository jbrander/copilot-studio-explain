# Messages, activities & cards

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

What a `SendActivity` / `Question` prompt can contain: text, message variations, cards, quick replies and attachments. See [11 Messages & Adaptive Cards](../../11-messages-and-adaptive-cards.md).

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

**42 definitions** on this page (22 with a `kind:`). [A–Z index of all definitions](index.md).

<details><summary>Contents</summary>

[Activity](#activity), [ActivityAttachment](#activityattachment), [ActivityImportance](#activityimportance), [ActivityTemplateBase](#activitytemplatebase), [ActivityType](#activitytype), [AdaptiveCardTemplate](#adaptivecardtemplate), [AgentMessageContentType](#agentmessagecontenttype), [AgentMessageRole](#agentmessagerole), [AgentMessageSortOrder](#agentmessagesortorder), [AttachmentLayout](#attachmentlayout), [BotFileReference](#botfilereference), [ContentShareContext](#contentsharecontext), [ConversationFileReference](#conversationfilereference), [EventActivity](#eventactivity), [ExpressionSegment](#expressionsegment), [FileAttachmentComponentMetadata](#fileattachmentcomponentmetadata), [FileReference](#filereference), [FileTemplate](#filetemplate), [HeroCardTemplate](#herocardtemplate), [ImageCardTemplate](#imagecardtemplate), [InputHint](#inputhint), [LatencyMessageSettings](#latencymessagesettings), [MakeCall](#makecall), [McpResourceFileReference](#mcpresourcefilereference), [Message](#message), [MessageBack](#messageback), [MessageNoKind](#messagenokind), [MessageOptionalKind](#messageoptionalkind), [OpenUrl](#openurl), [QuickReplyBase](#quickreplybase), [ResponseTemplate](#responsetemplate), [ResponseTemplateNoKind](#responsetemplatenokind), [SendHiddenMessageBack](#sendhiddenmessageback), [SendUserActivity](#senduseractivity), [TemplateLine](#templateline), [TemplateLineNoKind](#templatelinenokind), [TemplateResponseMode](#templateresponsemode), [TemplateSegment](#templatesegment), [TextSegment](#textsegment), [UriFileReference](#urifilereference), [VideoCardTemplate](#videocardtemplate), [WasmInvalidStateMessage](#wasminvalidstatemessage)

</details>

### Activity

YAML: `kind: Activity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `channelData` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `value` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `activityType` | `=` expression \| [ActivityType](#activitytype) |  | `"Typing"` | accepts a `=` Power Fx expression |
| `name` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `inputHint` | [InputHint](#inputhint) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ActivityTemplateBase](#activitytemplatebase)</sub>

### ActivityAttachment


**One of:** [VideoCardTemplate](#videocardtemplate), [ImageCardTemplate](#imagecardtemplate), [AdaptiveCardTemplate](#adaptivecardtemplate), [HeroCardTemplate](#herocardtemplate), [FileTemplate](#filetemplate)

<sub>Used by: [Message](#message), [MessageNoKind](#messagenokind), [MessageOptionalKind](#messageoptionalkind)</sub>

### ActivityImportance


**String enum.** Allowed values: `Low`, `Normal`, `High`

<sub>Used by: [Message](#message), [MessageNoKind](#messagenokind), [MessageOptionalKind](#messageoptionalkind)</sub>

### ActivityTemplateBase


**One of:** [EventActivity](#eventactivity), [MessageOptionalKind](#messageoptionalkind), [Activity](#activity)

<sub>Used by: [AdaptiveCardPrompt](actions.md#adaptivecardprompt), [BotSpeechSettings](settings.md#botspeechsettings), [BotSpeechSettingsNoKind](settings.md#botspeechsettingsnokind), [ContentModerationSettings](settings.md#contentmoderationsettings), [LatencyMessageSettings](#latencymessagesettings), [LatencyMessageSettingsNoKind](actions.md#latencymessagesettingsnokind), [MultiTurnActivity](evaluation.md#multiturnactivity), [MultiTurnActivityNoKind](evaluation.md#multiturnactivitynokind), [Question](actions.md#question), [ResponseTemplate](#responsetemplate), [ResponseTemplateNoKind](#responsetemplatenokind), [SendActivity](actions.md#sendactivity), [SendUserActivity](#senduseractivity), [TransferConversation](actions.md#transferconversation), [TransferToAgent](actions.md#transfertoagent), [UpdateActivity](actions.md#updateactivity), [VoiceInputSettings](settings.md#voiceinputsettings), [VoiceInputSettingsNoKind](actions.md#voiceinputsettingsnokind)</sub>

### ActivityType


**String enum.** Allowed values: `ContactRelationUpdate`, `ConversationUpdate`, `EndOfConversation`, `Event`, `Delay`, `DeleteUserData`, `Handoff`, `InstallationUpdate`, `Invoke`, `InvokeResponse`, `Message`, `MessageDelete`, `MessageReaction`, `MessageUpdate`, `Suggestion`, `Trace`, `Typing`, `Command`, `CommandResult`

<sub>Used by: [Activity](#activity), [OnActivity](triggers.md#onactivity)</sub>

### AdaptiveCardTemplate

YAML: `kind: AdaptiveCardTemplate`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `cardContent` | [AdaptiveCardExpression](other.md#adaptivecardexpression) | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [ActivityAttachment](#activityattachment)</sub>

### AgentMessageContentType


**String enum.** Allowed values: `Text`, `ImageFile`, `ImageUrl`

### AgentMessageRole


**String enum.** Allowed values: `User`, `Agent`

### AgentMessageSortOrder


**String enum.** Allowed values: `NewestFirst`, `OldestFirst`

### AttachmentLayout


**String enum.** Allowed values: `List`, `Carousel`

<sub>Used by: [Message](#message), [MessageNoKind](#messagenokind), [MessageOptionalKind](#messageoptionalkind)</sub>

### BotFileReference

YAML: `kind: BotFileReference`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `schemaName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [FileReference](#filereference)</sub>

### ContentShareContext

YAML: `kind: ContentShareContext`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `reusePolicy` | [BotComponentReusePolicy](settings.md#botcomponentreusepolicy) |  | `"None"` |  |
| `iconUri` | `string` |  |  |  |
| `helpUri` | `string` |  |  |  |
| `accentColor` | `string` |  |  |  |

_No other properties are allowed._

### ConversationFileReference

YAML: `kind: ConversationFileReference`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `value` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [FileReference](#filereference)</sub>

### EventActivity

YAML: `kind: EventActivity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `channelData` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `value` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |
| `name` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [ActivityTemplateBase](#activitytemplatebase)</sub>

### ExpressionSegment

YAML: `kind: ExpressionSegment`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `expression` | [ValueExpression](other.md#valueexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [TemplateSegment](#templatesegment)</sub>

### FileAttachmentComponentMetadata

YAML: `kind: FileAttachmentComponentMetadata`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

### FileReference


**One of:** [BotFileReference](#botfilereference), [ConversationFileReference](#conversationfilereference), [UriFileReference](#urifilereference), [McpResourceFileReference](#mcpresourcefilereference)

### FileTemplate

YAML: `kind: FileTemplate`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `fileContent` | `=` expression | ✓ |  | accepts a `=` Power Fx expression |
| `name` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `contentType` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [ActivityAttachment](#activityattachment)</sub>

### HeroCardTemplate

YAML: `kind: HeroCardTemplate`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `title` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `subtitle` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `image` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `text` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `buttons` | list of [QuickReplyBase](#quickreplybase) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ActivityAttachment](#activityattachment)</sub>

### ImageCardTemplate

YAML: `kind: ImageCardTemplate`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `title` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `image` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [ActivityAttachment](#activityattachment)</sub>

### InputHint


**String enum.** Allowed values: `AcceptingInput`, `IgnoringInput`, `ExpectingInput`

<sub>Used by: [Activity](#activity)</sub>

### LatencyMessageSettings

YAML: `kind: LatencyMessageSettings`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `allowLatencyMessage` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `latencyMessage` | [ActivityTemplateBase](#activitytemplatebase) |  |  |  |
| `latencyMessageDelayInMilliseconds` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `minimumPlaybackInMilliseconds` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

### MakeCall

YAML: `kind: MakeCall`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `title` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `image` | `string` |  |  |  |
| `imageAltText` | `string` |  |  |  |
| `phoneNumber` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [QuickReplyBase](#quickreplybase)</sub>

### McpResourceFileReference

YAML: `kind: McpResourceFileReference`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `resourcePath` | `string` | ✓ |  |  |
| `serverSchemaName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [FileReference](#filereference)</sub>

### Message


**One of:** `string`, object

### MessageBack

YAML: `kind: MessageBack`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `title` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `image` | `string` |  |  |  |
| `imageAltText` | `string` |  |  |  |
| `text` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `displayText` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [QuickReplyBase](#quickreplybase)</sub>

### MessageNoKind


**One of:** `string`, object

### MessageOptionalKind


**One of:** `string`, object

<sub>Used by: [ActivityTemplateBase](#activitytemplatebase)</sub>

### OpenUrl

YAML: `kind: OpenUrl`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `title` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `image` | `string` |  |  |  |
| `imageAltText` | `string` |  |  |  |
| `url` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [QuickReplyBase](#quickreplybase)</sub>

### QuickReplyBase


**One of:** [MessageBack](#messageback), [OpenUrl](#openurl), [MakeCall](#makecall), [SendHiddenMessageBack](#sendhiddenmessageback)

<sub>Used by: [HeroCardTemplate](#herocardtemplate), [Message](#message), [MessageNoKind](#messagenokind), [MessageOptionalKind](#messageoptionalkind), [VideoCardTemplate](#videocardtemplate)</sub>

### ResponseTemplate

YAML: `kind: ResponseTemplate`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `activity` | [ActivityTemplateBase](#activitytemplatebase) |  |  |  |
| `mode` | `=` expression \| [TemplateResponseMode](#templateresponsemode) |  | `"Strict"` | accepts a `=` Power Fx expression |

_No other properties are allowed._

### ResponseTemplateNoKind

Same shape as [ResponseTemplate](#responsetemplate) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `activity` | [ActivityTemplateBase](#activitytemplatebase) |  |  |  |
| `mode` | `=` expression \| [TemplateResponseMode](#templateresponsemode) |  | `"Strict"` | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [AdaptiveDialog](dialogs.md#adaptivedialog), [AdaptiveDialogNoKind](dialogs.md#adaptivedialognokind), [AgentDialog](dialogs.md#agentdialog), [QuestionInputSettings](settings.md#questioninputsettings), [QuestionInputSettingsNoKind](task-actions.md#questioninputsettingsnokind), [TaskDialog](dialogs.md#taskdialog)</sub>

### SendHiddenMessageBack

YAML: `kind: SendHiddenMessageBack`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `title` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `image` | `string` |  |  |  |
| `imageAltText` | `string` |  |  |  |
| `text` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [QuickReplyBase](#quickreplybase)</sub>

### SendUserActivity

YAML: `kind: SendUserActivity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `originType` | [TestActivitySourceType](evaluation.md#testactivitysourcetype) |  | `"Manual"` |  |
| `activity` | [ActivityTemplateBase](#activitytemplatebase) |  |  |  |
| `activityAssertions` | list of [TestAssertionBase](evaluation.md#testassertionbase) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TestActivityDefinition](evaluation.md#testactivitydefinition)</sub>

### TemplateLine


**Type:** `string`

### TemplateLineNoKind


**Type:** `string`

<sub>Used by: [AnswerQuestionWithAI](actions.md#answerquestionwithai), [AutomaticHoldSettings](actions.md#automaticholdsettings), [ConnectorTriggerDefinition](components.md#connectortriggerdefinition), [DefaultFeedbackCollection](settings.md#defaultfeedbackcollection), [DefaultHoldSettings](settings.md#defaultholdsettings), [DefaultHoldSettingsNoKind](settings.md#defaultholdsettingsnokind), [ExternalChatCompletionModelConfiguration](actions.md#externalchatcompletionmodelconfiguration), [ExternalChatCompletionModelConfigurationNoKind](actions.md#externalchatcompletionmodelconfigurationnokind), [FileGroupKnowledgeSource](knowledge.md#filegroupknowledgesource), [GptComponentMetadata](components.md#gptcomponentmetadata), [GptComponentMetadataNoKind](components.md#gptcomponentmetadatanokind), [GraphConnectorSearchSource](knowledge.md#graphconnectorsearchsource), [InvokeComputerUsingAgentTaskAction](task-actions.md#invokecomputerusingagenttaskaction), [Message](#message), [MessageNoKind](#messagenokind), [MessageOptionalKind](#messageoptionalkind), [RealtimeVoiceAudioLatencySettings](settings.md#realtimevoiceaudiolatencysettings), [RealtimeVoiceMessageLatencySettings](settings.md#realtimevoicemessagelatencysettings), [RealtimeVoiceMessageLatencySettingsNoKind](settings.md#realtimevoicemessagelatencysettingsnokind), [RecurrenceTriggerDefinition](components.md#recurrencetriggerdefinition) … and 9 more</sub>

### TemplateResponseMode


**String enum.** Allowed values: `Strict`, `Generated`

<sub>Used by: [ResponseTemplate](#responsetemplate), [ResponseTemplateNoKind](#responsetemplatenokind)</sub>

### TemplateSegment


**One of:** [TextSegment](#textsegment), [ExpressionSegment](#expressionsegment)

### TextSegment

YAML: `kind: TextSegment`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `value` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [TemplateSegment](#templatesegment)</sub>

### UriFileReference

YAML: `kind: UriFileReference`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `uri` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [FileReference](#filereference)</sub>

### VideoCardTemplate

YAML: `kind: VideoCardTemplate`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `title` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `subtitle` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `image` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `media` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `text` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `buttons` | list of [QuickReplyBase](#quickreplybase) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ActivityAttachment](#activityattachment)</sub>

### WasmInvalidStateMessage


**String enum.** Allowed values: `ContentIsNull`, `FormatIsNull`, `DialogActionIsNull`, `SchemaIsNull`, `InvalidLanguageCode`, `InvalidFormat`, `ComponentNotFound`, `DialogNotFound`, `SkillNotFound`, `ActionNotFound`, `FlowNotFound`, `InvokeSkillActionNotFound`, `BotNotLoaded`, `FeatureNotLoaded`, `ErrorGettingDynamicPlanToolDefinition`, `TranslationsAppliedWithCurrentLanguage`
