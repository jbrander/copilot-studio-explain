# Dialogs (topics, tools, child agents) & recognizers

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

`AdaptiveDialog` (a topic), `TaskDialog` (a tool/action), `AgentDialog` (a child agent) and the recognizers that route user messages. See [04 Topics & triggers](../../04-topics-and-triggers.md).

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

**28 definitions** on this page (14 with a `kind:`). [A–Z index of all definitions](index.md).

<details><summary>Contents</summary>

[AdaptiveDialog](#adaptivedialog), [AdaptiveDialogNoKind](#adaptivedialognokind), [AdvanceNLURecognizer](#advancenlurecognizer), [AgentDialog](#agentdialog), [ClassicAIRecognizer](#classicairecognizer), [CLIAgentRecognizer](#cliagentrecognizer), [ClosedListConversionMetadataNoKind](#closedlistconversionmetadatanokind), [ClosedListEntity](#closedlistentity), [ClosedListItemNoKind](#closedlistitemnokind), [ConversionMetadataNoKind](#conversionmetadatanokind), [DialogBase](#dialogbase), [ExternalEntity](#externalentity), [ExternalIntent](#externalintent), [GenerativeAIRecognizer](#generativeairecognizer), [Intent](#intent), [IntentBase](#intentbase), [IntentNoKind](#intentnokind), [LocalizableContentContainer](#localizablecontentcontainer), [Recognizer](#recognizer), [RecognizerBase](#recognizerbase), [RegexEntity](#regexentity), [RegexIntentPatternNoKind](#regexintentpatternnokind), [RegexRecognizer](#regexrecognizer), [TaskActionOutputMode](#taskactionoutputmode), [TaskDialog](#taskdialog), [TaskOutputNoKind](#taskoutputnokind), [TopicStartBehavior](#topicstartbehavior), [VirtualAgents.Recognizer](#virtualagentsrecognizer)

</details>

### AdaptiveDialog

YAML: `kind: AdaptiveDialog`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `confirmation` | [ResponseTemplateNoKind](messages.md#responsetemplatenokind) |  |  |  |
| `response` | [ResponseTemplateNoKind](messages.md#responsetemplatenokind) |  |  |  |
| `inputs` | list of [TaskInput](task-actions.md#taskinput) |  |  |  |
| `modelDisplayName` | `string` |  |  |  |
| `modelDescription` | `string` |  |  |  |
| `recognizer` | [Recognizer](#recognizer) |  |  |  |
| `triggers` | list of [TriggerBase](triggers.md#triggerbase) |  |  |  |
| `endAutomatically` | [BoolExpression](other.md#boolexpression) |  | `true` | accepts a `=` Power Fx expression |
| `startBehavior` | [TopicStartBehavior](#topicstartbehavior) |  | `"Default"` |  |
| `beginDialog` | [TriggerBase](triggers.md#triggerbase) |  |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `conversionMetadata` | [ConversionMetadataNoKind](#conversionmetadatanokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogBase](#dialogbase), [LocalizableContentContainer](#localizablecontentcontainer)</sub>

### AdaptiveDialogNoKind

Same shape as [AdaptiveDialog](#adaptivedialog) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `confirmation` | [ResponseTemplateNoKind](messages.md#responsetemplatenokind) |  |  |  |
| `response` | [ResponseTemplateNoKind](messages.md#responsetemplatenokind) |  |  |  |
| `inputs` | list of [TaskInput](task-actions.md#taskinput) |  |  |  |
| `modelDisplayName` | `string` |  |  |  |
| `modelDescription` | `string` |  |  |  |
| `recognizer` | [Recognizer](#recognizer) |  |  |  |
| `triggers` | list of [TriggerBase](triggers.md#triggerbase) |  |  |  |
| `endAutomatically` | [BoolExpression](other.md#boolexpression) |  | `true` | accepts a `=` Power Fx expression |
| `startBehavior` | [TopicStartBehavior](#topicstartbehavior) |  | `"Default"` |  |
| `beginDialog` | [TriggerBase](triggers.md#triggerbase) |  |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `conversionMetadata` | [ConversionMetadataNoKind](#conversionmetadatanokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogBase](#dialogbase)</sub>

### AdvanceNLURecognizer

YAML: `kind: AdvanceNLURecognizer`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [RecognizerBase](#recognizerbase)</sub>

### AgentDialog

YAML: `kind: AgentDialog`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `confirmation` | [ResponseTemplateNoKind](messages.md#responsetemplatenokind) |  |  |  |
| `response` | [ResponseTemplateNoKind](messages.md#responsetemplatenokind) |  |  |  |
| `inputs` | list of [TaskInput](task-actions.md#taskinput) |  |  |  |
| `beginDialog` | [TriggerBase](triggers.md#triggerbase) |  |  |  |
| `settings` | [GptComponentMetadataNoKind](components.md#gptcomponentmetadatanokind) |  |  |  |
| `inputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `outputType` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [DialogBase](#dialogbase), [LocalizableContentContainer](#localizablecontentcontainer)</sub>

### ClassicAIRecognizer

YAML: `kind: ClassicAIRecognizer`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [RecognizerBase](#recognizerbase)</sub>

### CLIAgentRecognizer

YAML: `kind: CLIAgentRecognizer`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [RecognizerBase](#recognizerbase)</sub>

### ClosedListConversionMetadataNoKind

*Title:* ClosedListConversionMetadata  
Same shape as [ClosedListConversionMetadata](entities.md#closedlistconversionmetadata) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `legacyIdToDisplayNameMapping` | map of name → `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [ClosedListEntity](#closedlistentity)</sub>

### ClosedListEntity

YAML: `kind: ClosedListEntity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `samples` | list of [AnnotatedSampleNoKind](evaluation.md#annotatedsamplenokind) |  |  |  |
| `optionSetName` | `string` |  |  |  |
| `smartMatchingEnabled` | `boolean` |  | `true` |  |
| `items` | list of [ClosedListItemNoKind](#closedlistitemnokind) | ✓ |  |  |
| `conversionMetadata` | [ClosedListConversionMetadataNoKind](#closedlistconversionmetadatanokind) |  |  |  |
| `dtmfMultipleChoiceOptions` | [DtmfMultipleChoiceOptionsNoKind](settings.md#dtmfmultiplechoiceoptionsnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [Entity](entities.md#entity), [LocalizableContentContainer](#localizablecontentcontainer)</sub>

### ClosedListItemNoKind

Same shape as [ClosedListItem](entities.md#closedlistitem) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `displayName` | `string` | ✓ |  |  |
| `synonyms` | list of `string` |  |  |  |
| `dtmfKey` | [DtmfMappingKey](settings.md#dtmfmappingkey) |  |  |  |

_No other properties are allowed._

<sub>Used by: [ClosedListEntity](#closedlistentity)</sub>

### ConversionMetadataNoKind

*Title:* ConversionMetadata  
Same shape as [ConversionMetadata](components.md#conversionmetadata) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `variableIdToNameMapping` | map of name → `string` |  |  |  |
| `initializedVariableIdToTypeMapping` | map of name → [DataType](data-types.md#datatype) |  |  |  |

_No other properties are allowed._

<sub>Used by: [AdaptiveDialog](#adaptivedialog), [AdaptiveDialogNoKind](#adaptivedialognokind)</sub>

### DialogBase


**One of:** [TaskDialog](#taskdialog), [AdaptiveDialog](#adaptivedialog), [AdaptiveDialogNoKind](#adaptivedialognokind), [AgentDialog](#agentdialog)

<sub>Used by: [DialogComponent](components.md#dialogcomponent), [DialogComponentNoKind](components.md#dialogcomponentnokind)</sub>

### ExternalEntity

YAML: `kind: ExternalEntity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `valueType` | [DataType](data-types.md#datatype) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [Entity](entities.md#entity), [LocalizableContentContainer](#localizablecontentcontainer)</sub>

### ExternalIntent

YAML: `kind: ExternalIntent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `displayName` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `availableInChannels` | list of `string` |  |  |  |
| `includeInOnSelectIntent` | `boolean` |  | `true` |  |
| `externalIntentId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [IntentBase](#intentbase)</sub>

### GenerativeAIRecognizer

YAML: `kind: GenerativeAIRecognizer`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [RecognizerBase](#recognizerbase)</sub>

### Intent

YAML: `kind: Intent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `displayName` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `availableInChannels` | list of `string` |  |  |  |
| `includeInOnSelectIntent` | `boolean` |  | `true` |  |
| `triggerQueries` | list of [AnnotatedSampleNoKind](evaluation.md#annotatedsamplenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [IntentBase](#intentbase)</sub>

### IntentBase


**One of:** [Intent](#intent), [IntentNoKind](#intentnokind), [ExternalIntent](#externalintent)

<sub>Used by: [OnEscalate](triggers.md#onescalate), [OnRecognizedIntent](triggers.md#onrecognizedintent), [OnSystemIntent](triggers.md#onsystemintent)</sub>

### IntentNoKind

*Title:* Intent  
Same shape as [Intent](#intent) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `displayName` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `availableInChannels` | list of `string` |  |  |  |
| `includeInOnSelectIntent` | `boolean` |  | `true` |  |
| `triggerQueries` | list of [AnnotatedSampleNoKind](evaluation.md#annotatedsamplenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [IntentBase](#intentbase)</sub>

### LocalizableContentContainer


**One of:** [RegexEntity](#regexentity), [ClosedListEntity](#closedlistentity), [ExternalEntity](#externalentity), [TaskDialog](#taskdialog), [AdaptiveDialog](#adaptivedialog), [AgentDialog](#agentdialog)

<sub>Used by: [TranslationsComponent](components.md#translationscomponent), [TranslationsComponentNoKind](components.md#translationscomponentnokind)</sub>

### Recognizer


**One of:** [RegexRecognizer](#regexrecognizer), [VirtualAgents.Recognizer](#virtualagentsrecognizer)

<sub>Used by: [AdaptiveDialog](#adaptivedialog), [AdaptiveDialogNoKind](#adaptivedialognokind)</sub>

### RecognizerBase


**One of:** [AdvanceNLURecognizer](#advancenlurecognizer), [ClassicAIRecognizer](#classicairecognizer), [GenerativeAIRecognizer](#generativeairecognizer), [CLIAgentRecognizer](#cliagentrecognizer)

<sub>Used by: [BotConfiguration](components.md#botconfiguration), [BotConfigurationNoKind](components.md#botconfigurationnokind)</sub>

### RegexEntity

YAML: `kind: RegexEntity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `samples` | list of [AnnotatedSampleNoKind](evaluation.md#annotatedsamplenokind) |  |  |  |
| `pattern` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [Entity](entities.md#entity), [LocalizableContentContainer](#localizablecontentcontainer)</sub>

### RegexIntentPatternNoKind

*Title:* RegexIntentPattern  
Same shape as [RegexIntentPattern](entities.md#regexintentpattern) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `intent` | `string` | ✓ |  |  |
| `pattern` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [RegexRecognizer](#regexrecognizer)</sub>

### RegexRecognizer

YAML: `kind: RegexRecognizer`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `intents` | list of [RegexIntentPatternNoKind](#regexintentpatternnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [Recognizer](#recognizer)</sub>

### TaskActionOutputMode


**String enum.** Allowed values: `All`, `Specific`

<sub>Used by: [TaskDialog](#taskdialog)</sub>

### TaskDialog

YAML: `kind: TaskDialog`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `confirmation` | [ResponseTemplateNoKind](messages.md#responsetemplatenokind) |  |  |  |
| `response` | [ResponseTemplateNoKind](messages.md#responsetemplatenokind) |  |  |  |
| `inputs` | list of [TaskInput](task-actions.md#taskinput) |  |  |  |
| `modelDisplayName` | `string` |  |  |  |
| `modelDescription` | `string` |  |  |  |
| `outputs` | list of [TaskOutputNoKind](#taskoutputnokind) |  |  |  |
| `action` | [TaskAction](task-actions.md#taskaction) |  |  |  |
| `outputMode` | [TaskActionOutputMode](#taskactionoutputmode) |  | `"All"` |  |
| `triggerCondition` | `true` \| `false` \| [BoolExpression](other.md#boolexpression) |  | `true` |  |

_No other properties are allowed._

<sub>Used by: [DialogBase](#dialogbase), [LocalizableContentContainer](#localizablecontentcontainer)</sub>

### TaskOutputNoKind

*Title:* TaskOutput  
Same shape as [TaskOutput](actions.md#taskoutput) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `propertyName` | `string` | ✓ |  |  |
| `name` | `string` |  |  |  |
| `description` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [TaskDialog](#taskdialog)</sub>

### TopicStartBehavior


**String enum.** Allowed values: `Default`, `CancelOtherTopics`, `UseLatestPublishedContentAndCancelOtherTopics`

<sub>Used by: [AdaptiveDialog](#adaptivedialog), [AdaptiveDialogNoKind](#adaptivedialognokind)</sub>

### VirtualAgents.Recognizer

YAML: `kind: VirtualAgents.Recognizer`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [Recognizer](#recognizer)</sub>
