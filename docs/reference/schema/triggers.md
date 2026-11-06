# Triggers

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

Everything that can appear in a topic's `beginDialog:` (or `triggers:`) slot. See [04 Topics & triggers](../../04-topics-and-triggers.md).

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

**30 definitions** on this page (28 with a `kind:`). [A–Z index of all definitions](index.md).

<details><summary>Contents</summary>

[ContinueConversationConnectorTrigger](#continueconversationconnectortrigger), [CreateConversationConnectorTrigger](#createconversationconnectortrigger), [OnActivity](#onactivity), [OnAnalyzeText](#onanalyzetext), [OnAnsweringMachineDetected](#onansweringmachinedetected), [OnConversationStart](#onconversationstart), [OnCopilotRedirect](#oncopilotredirect), [OnDtmfKeyPress](#ondtmfkeypress), [OnError](#onerror), [OnEscalate](#onescalate), [OnEventActivity](#oneventactivity), [OnGeneratedResponse](#ongeneratedresponse), [OnInactivity](#oninactivity), [OnKnowledgeRequested](#onknowledgerequested), [OnOutgoingMessage](#onoutgoingmessage), [OnPlanComplete](#onplancomplete), [OnRecognizedIntent](#onrecognizedintent), [OnRedirect](#onredirect), [OnSelectIntent](#onselectintent), [OnSelectIntentTriggerBehavior](#onselectintenttriggerbehavior), [OnSignIn](#onsignin), [OnSilence](#onsilence), [OnSystemIntent](#onsystemintent), [OnSystemRedirect](#onsystemredirect), [OnToolSelected](#ontoolselected), [OnUnknownDtmfKey](#onunknowndtmfkey), [OnUnknownIntent](#onunknownintent), [OnUnrecognizedSpeech](#onunrecognizedspeech), [RecurrenceTrigger](#recurrencetrigger), [TriggerBase](#triggerbase)

</details>

### ContinueConversationConnectorTrigger

YAML: `kind: ContinueConversationConnectorTrigger`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `connectionReference` | `string` | ✓ |  |  |
| `connectionProperties` | [ConnectionPropertiesNoKind](settings.md#connectionpropertiesnokind) |  |  |  |
| `connectorTriggerId` | `string` | ✓ |  |  |
| `dynamicInputSchema` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `dynamicOutputSchema` | [DataType](data-types.md#datatype) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### CreateConversationConnectorTrigger

YAML: `kind: CreateConversationConnectorTrigger`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `connectionReference` | `string` | ✓ |  |  |
| `connectionProperties` | [ConnectionPropertiesNoKind](settings.md#connectionpropertiesnokind) |  |  |  |
| `connectorTriggerId` | `string` | ✓ |  |  |
| `dynamicInputSchema` | [RecordNoKind](data-types.md#recordnokind) |  |  |  |
| `dynamicOutputSchema` | [DataType](data-types.md#datatype) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnActivity

YAML: `kind: OnActivity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `type` | [ActivityType](messages.md#activitytype) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnAnalyzeText

YAML: `kind: OnAnalyzeText`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnAnsweringMachineDetected

YAML: `kind: OnAnsweringMachineDetected`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnConversationStart

YAML: `kind: OnConversationStart`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnCopilotRedirect

YAML: `kind: OnCopilotRedirect`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `askForConfirmation` | `boolean` |  |  |  |
| `conversationStarters` | list of `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnDtmfKeyPress

YAML: `kind: OnDtmfKeyPress`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `dtmfKey` | [DtmfMappingKey](settings.md#dtmfmappingkey) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnError

YAML: `kind: OnError`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnEscalate

YAML: `kind: OnEscalate`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `intent` | [IntentBase](dialogs.md#intentbase) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnEventActivity

YAML: `kind: OnEventActivity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `eventName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnGeneratedResponse

YAML: `kind: OnGeneratedResponse`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnInactivity

YAML: `kind: OnInactivity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `durationInSeconds` | `900` \| `1800` \| `2700` \| `3600` \| [NumberExpression](other.md#numberexpression) |  | `1800` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnKnowledgeRequested

YAML: `kind: OnKnowledgeRequested`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnOutgoingMessage

YAML: `kind: OnOutgoingMessage`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnPlanComplete

YAML: `kind: OnPlanComplete`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnRecognizedIntent

YAML: `kind: OnRecognizedIntent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `intent` | [IntentBase](dialogs.md#intentbase) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnRedirect

YAML: `kind: OnRedirect`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnSelectIntent

YAML: `kind: OnSelectIntent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `triggerBehavior` | [OnSelectIntentTriggerBehavior](#onselectintenttriggerbehavior) |  | `"Default"` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnSelectIntentTriggerBehavior


**String enum.** Allowed values: `Default`, `Always`

<sub>Used by: [OnSelectIntent](#onselectintent)</sub>

### OnSignIn

YAML: `kind: OnSignIn`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnSilence

YAML: `kind: OnSilence`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `silenceDetectionTimeoutInMilliseconds` | `0` \| [IntExpression](other.md#intexpression) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnSystemIntent

YAML: `kind: OnSystemIntent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `intent` | [IntentBase](dialogs.md#intentbase) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnSystemRedirect

YAML: `kind: OnSystemRedirect`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnToolSelected

YAML: `kind: OnToolSelected`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |
| `description` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnUnknownDtmfKey

YAML: `kind: OnUnknownDtmfKey`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnUnknownIntent

YAML: `kind: OnUnknownIntent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### OnUnrecognizedSpeech

YAML: `kind: OnUnrecognizedSpeech`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### RecurrenceTrigger

YAML: `kind: RecurrenceTrigger`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `actions` | list of [DialogAction](actions.md#dialogaction) |  |  |  |
| `id` | `string` | ✓ |  |  |
| `condition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `priority` | `number` |  | `0` |  |
| `runOnce` | `boolean` |  | `false` |  |

_No other properties are allowed._

<sub>Used by: [TriggerBase](#triggerbase)</sub>

### TriggerBase


**One of:** [RecurrenceTrigger](#recurrencetrigger), [CreateConversationConnectorTrigger](#createconversationconnectortrigger), [ContinueConversationConnectorTrigger](#continueconversationconnectortrigger), [OnRedirect](#onredirect), [OnPlanComplete](#onplancomplete), [OnKnowledgeRequested](#onknowledgerequested), [OnGeneratedResponse](#ongeneratedresponse), [OnOutgoingMessage](#onoutgoingmessage), [OnCopilotRedirect](#oncopilotredirect), [OnActivity](#onactivity), [OnEventActivity](#oneventactivity), [OnDtmfKeyPress](#ondtmfkeypress), [OnToolSelected](#ontoolselected), [OnRecognizedIntent](#onrecognizedintent), [OnSystemIntent](#onsystemintent), [OnEscalate](#onescalate), [OnInactivity](#oninactivity), [OnAnsweringMachineDetected](#onansweringmachinedetected), [OnSystemRedirect](#onsystemredirect), [OnSelectIntent](#onselectintent), [OnError](#onerror), [OnUnrecognizedSpeech](#onunrecognizedspeech), [OnConversationStart](#onconversationstart), [OnAnalyzeText](#onanalyzetext), [OnSignIn](#onsignin), [OnUnknownIntent](#onunknownintent), [OnUnknownDtmfKey](#onunknowndtmfkey), [OnSilence](#onsilence)

<sub>Used by: [AdaptiveDialog](dialogs.md#adaptivedialog), [AdaptiveDialogNoKind](dialogs.md#adaptivedialognokind), [AgentDialog](dialogs.md#agentdialog)</sub>
