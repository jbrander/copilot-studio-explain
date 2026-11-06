# 05 — Nodes

A topic's `beginDialog.actions:` list holds its [nodes](glossary.md#node) (the steps of a topic). The schema's
[DialogAction](reference/schema/actions.md#dialogaction) union has **52 node kinds**. This page covers the properties every node shares, the ~20 you'll
use most (with examples and property tables), and gives a one-line summary of the rest. Nodes that have their own
guide (knowledge, connectors, flows, cards) link there rather than repeating it.

> **Nodes vs. actions vs. tools.** The topic property is called `actions:`, and the vendor skills call its items
> *nodes*. The files in the `actions/` folder are something else: [tools](glossary.md#tool)
> (`TaskDialog`), covered in [08 Tools & connectors](08-tools-and-connectors.md).

## All node kinds at a glance

| Group | Kinds |
|---|---|
| [Messaging](#messaging) | `SendActivity`, `UpdateActivity`, `DeleteActivity` |
| [Questions](#questions) | `Question`, `AdaptiveCardPrompt`, `CSATQuestion`, `OAuthInput` |
| [Variables](#variables) | `SetVariable`, `SetTextVariable`, `SetMultipleVariables`, `ParseValue`, `ResetVariable`, `ClearAllVariables`, `EditTable`, `EditTableV2` |
| [Branching & loops](#branching--loops) | `ConditionGroup`, `Foreach`, `BreakLoop`, `ContinueLoop`, `GotoAction`, `ActionScope`, `DisconnectedNodeContainer` |
| [Topic navigation](#topic-navigation) | `BeginDialog`, `ReplaceDialog`, `RepeatDialog`, `EndDialog`, `CancelDialog`, `CancelAllDialogs`, `EndConversation`, `RecognizeIntent`, `DisableTrigger`, `ActivateExternalTrigger` |
| [AI & knowledge](#ai--knowledge) | `SearchAndSummarizeContent`, `AnswerQuestionWithAI`, `SearchKnowledgeSources`, `CreateSearchQuery`, `SearchAndSummarizeWithCustomModel`, `InvokeCustomModelAction`, `InvokeAIBuilderModelAction` |
| [Integration](#integration) | `InvokeConnectorAction`, `InvokeFlowAction`, `InvokeMcpToolAction`, `InvokeSkillAction`, `HttpRequestAction`, `WaitForConnectorTrigger` |
| [Conversation control](#conversation-control) | `TransferConversation`, `TransferConversationV2`, `SignOutUser`, `GetActivityMembers`, `GetConversationMembers`, `EmitEvent` |
| [Telemetry](#telemetry) | `LogCustomTelemetryEvent` |

## Properties every node has

| Property | Type/values | Req. | Meaning |
|---|---|---|---|
| `kind` | one of the 52 kinds | ✓ | The node type. **Verify every `kind:` exists** (see the [A–Z index](reference/schema/index.md)). The vendor skills call invented kinds "the #1 source of hallucination errors". |
| `id` | string | ✓ | Unique within the topic. Convention: `<nodeType>_<6–8 random alphanumerics>`, e.g. `sendMessage_k3Jd9a`. Duplicate IDs fail on push. |
| `displayName` | string | | The label shown on the node in the authoring canvas. |
| `disabled` | boolean, default `false` | | No vendor guidance. By its name, the node is skipped at runtime (inferred). |
| `conversationOutcome` | `ResolvedConfirmed` \| `ResolvedImplied` \| `Escalated` | | No vendor guidance. By its name, marks the conversation's outcome for analytics when this node runs (inferred). |

AI and integration nodes also accept `latencyMessageSettings` (`allowLatencyMessage`, `latencyMessage`,
`latencyMessageDelayInMilliseconds`, `minimumPlaybackInMilliseconds`). The source has no guidance on it; by its
name, it's an interim "please wait" message for slow calls (inferred). See also the knowledge-hold-message pattern
in [13 Patterns](13-patterns.md).

Expressions follow the usual rules: `=` starts a [Power Fx](glossary.md#power-fx) (the formula language) expression,
`{…}` interpolates inside text, anything else is literal. See [06 Power Fx](06-variables-entities-powerfx.md#power-fx).

## Messaging

### SendActivity

[SendActivity](reference/schema/actions.md#sendactivity) sends a message. **When to use it:** for greetings, explanations and mid-flow status. Under
[generative orchestration](glossary.md#generative-orchestration) (an LLM plans the reply), prefer **topic outputs**
for final results so the orchestrator writes the answer (see
[04](04-topics-and-triggers.md#topic-inputs-and-outputs)).

```yaml
# three SendActivity nodes, from simplest to richest
- kind: SendActivity
  id: sendMessage_k3Jd9a
  activity: "Thanks, {Topic.UserName}! Your request is logged."

- kind: SendActivity
  id: sendMessage_vR4r1x
  activity:
    text:
      - I'm not sure I understand. Could you rephrase that?
      - I didn't quite get that. Can you try asking in a different way?

- kind: SendActivity
  id: sendMessage_q8Rp2m
  activity:
    text:
      - An error has occurred. Error code {System.Error.Code}.
    speak:
      - An error has occurred, please try again.
    quickReplies:
      - kind: MessageBack
        title: Start over
        text: Start over
```

| `activity` form / property | Meaning |
|---|---|
| plain string | One message, literal or `{…}` template. Use `\|-` for multi-line text. Channels collapse plain newlines, so use `<br /><br />` for paragraph breaks (see [11](11-messages-and-adaptive-cards.md)). |
| `text:` (list) | **Variations**: one is picked at random each time (`randomizeVariantSelection`, default `true`). |
| `speak:` (list) | What voice channels say instead of `text`. |
| `quickReplies:` | [Quick reply](glossary.md#quick-reply) buttons (`MessageBack`, `OpenUrl`, `MakeCall`, `SendHiddenMessageBack`). |
| `attachments:` | Cards (`HeroCardTemplate`, `AdaptiveCardTemplate`, …). For interactive [Adaptive Cards](glossary.md#adaptive-card) (JSON-described UI cards) the vendor rule is to use `AdaptiveCardPrompt` instead. See [11](11-messages-and-adaptive-cards.md). |

All `activity` properties are listed under [MessageOptionalKind](reference/schema/messages.md#messageoptionalkind).

### UpdateActivity and DeleteActivity

| Kind | Purpose |
|---|---|
| [UpdateActivity](reference/schema/actions.md#updateactivity) | Replace a previously sent message (`activityId`, new `activity`). No vendor guidance. |
| [DeleteActivity](reference/schema/actions.md#deleteactivity) | Delete a previously sent message (`activityId`). No vendor guidance. |

## Questions

### Question

[Question](reference/schema/actions.md#question) asks the user something and stores the recognised answer in a variable. The
[entity](glossary.md#entity) (the type of value to recognise) decides what counts as a valid answer.
**When to use it:** under generative orchestration, only for *conditional* questions ("ask for a phone number only
if the group is larger than 6") or end-of-flow confirmations. Use topic `inputs:` for anything the topic always
needs ([04](04-topics-and-triggers.md#topic-inputs-and-outputs)). Under
[classic orchestration](glossary.md#classic-orchestration), questions are the normal way to collect input.

```yaml
- kind: Question
  id: question_k7xPm2
  variable: init:Topic.UserName
  prompt: What is your name?
  entity: StringPrebuiltEntity
  alwaysPrompt: true
  interruptionPolicy:
    allowInterruption: false

- kind: Question
  id: question_eM4il9
  variable: init:Topic.WorkEmail
  prompt: What's your work email address?
  entity: EmailPrebuiltEntity
  repeatCount: 1
  validation: =EndsWith(Topic.WorkEmail, "@contoso.com")
  invalidPrompt: Please use your @contoso.com address.
```

| Property | Type/values | Req. | Meaning |
|---|---|---|---|
| `variable` | variable name | ✓ | Where the answer goes. Use `init:Topic.X` the first time. |
| `prompt` | message (same forms as `SendActivity.activity`) | ✓ | The question text. |
| `entity` | entity reference | ✓ | What to recognise: a [prebuilt entity](glossary.md#prebuilt-entity) name (`StringPrebuiltEntity`, `NumberPrebuiltEntity`, `BooleanPrebuiltEntity`, `DateTimePrebuiltEntity`, `EmailPrebuiltEntity`, …), a closed list, a regex, or a dynamic list. See [06 Entities](06-variables-entities-powerfx.md#entities). |
| `alwaysPrompt` | boolean or `=` expression, default `false` | | `true` = ask even if the variable already has a value. Vendor templates set it to `true`. |
| `skipQuestionMode` | `SkipOnFirstExecutionIfVariableHasValue` (default) \| `AlwaysSkipIfVariableHasValue` \| `AlwaysAsk` | | When an existing value lets the question be skipped (this is how [slot filling](glossary.md#slot-filling) (auto-filling answers from earlier input) works). No vendor guidance beyond the names. |
| `interruptionPolicy.allowInterruption` | boolean, default `true` | | `false` = the user's reply can't switch to another topic. Vendor templates set `false`. |
| `repeatCount` | `0`–`2` or expression, default `2` | | Re-prompts after an unrecognised answer (inferred from the name). |
| `invalidPrompt`, `unrecognizedPrompt` | message | | Shown when `validation` fails / nothing is recognised (inferred). |
| `validation` | `=` boolean expression | | Extra check on the recognised value (inferred). |
| `defaultValue`, `defaultValueResponse` | value / message | | Value (and message) used after retries run out (inferred). |

`voiceInputSettings`, `holdSettings` and `fallbackDialogOnInvalidEntity` also exist, without vendor guidance.

**Gotcha:** compare text answers with **double-quoted** strings (`=Topic.UserChoice = "Basic"`). The vendor
`question-topic` template uses single quotes, which in Power Fx delimit *names*, not text.

### AdaptiveCardPrompt, CSATQuestion, OAuthInput

| Kind | Purpose | Key properties |
|---|---|---|
| [AdaptiveCardPrompt](reference/schema/actions.md#adaptivecardprompt) | Show an Adaptive Card form and wait for the user to submit it. The vendor rule is to use it for **every** Adaptive Card, even display-only ones. | `card` (JSON in a `\|` block, required), `output.binding` (card input id → `Topic.X`, required), `outputType` (required). Every card needs an `Action.Submit`. Details in [11 Messages & Adaptive Cards](11-messages-and-adaptive-cards.md). |
| [CSATQuestion](reference/schema/actions.md#csatquestion) | Ask the built-in [CSAT](glossary.md#csat) (customer-satisfaction rating) question. | `variable` (optional). Listed in int-reference; no example. |
| [OAuthInput](reference/schema/actions.md#oauthinput) | Show a sign-in card. Used in the `OnSignIn` topic. | `title`, `text`, `resource`. See the [sign-in example](04-topics-and-triggers.md#sign-in). |

```yaml
- kind: OAuthInput
  id: oauthInput_lG1n4x
  title: Login
  text: To continue, please login
```

## Variables

Variable names are `Topic.X`, `Global.X` or `System.X`, written **without** `=` when they're the *target*. Prefix
the first assignment of a [topic variable](glossary.md#topic-variable) (a variable that lives only while its topic
runs) with `init:`. See [06 Variables](06-variables-entities-powerfx.md).

### SetVariable

[SetVariable](reference/schema/actions.md#setvariable) sets a variable to a value or a Power Fx expression. **When to use it:** for almost any computation,
copy or flag.

```yaml
- kind: SetVariable
  id: setVariable_cT1me0
  variable: init:Topic.CurrentTime
  value: =Text(Now(), DateTimeFormat.UTC)

- kind: SetVariable
  id: setVariable_fLag01
  variable: Global.InactiveConversation
  value: true

- kind: SetVariable
  id: setVariable_rEc0rd
  variable: Global.UserContext
  value: |-
    ={
        Country: "USA",
        Language: "English"
    }
```

| Property | Type | Req. | Meaning |
|---|---|---|---|
| `variable` | variable name | ✓ | Target. `init:` on first assignment of a topic variable. |
| `value` | literal or `=` expression | ✓ | New value. Literal `true`/numbers keep their type; `=…` is evaluated. |

### SetTextVariable

[SetTextVariable](reference/schema/actions.md#settextvariable) sets a variable to a **text template** (`{…}` interpolation). **When to use it:** to turn
non-text values into text (`"Guests: {Topic.NumberOfGuests}"`) or to set a literal string. The `add-node` skill notes
it's YAML-only: it doesn't render on the authoring canvas.

```yaml
- kind: SetTextVariable
  id: setTextVariable_nOn3x
  variable: Topic.NoneOfTheseDisplayName
  value: None of these
```

| Property | Type | Req. | Meaning |
|---|---|---|---|
| `variable` | variable name | ✓ | Target. |
| `value` | text template | ✓ | Literal text with optional `{…}` expressions. |

### ClearAllVariables

[ClearAllVariables](reference/schema/actions.md#clearallvariables) resets a whole group of variables. **When to use it:** in reset, inactivity and
"start over" topics (see the [inactivity example](04-topics-and-triggers.md#inactivity-reset-teams)).

```yaml
- kind: ClearAllVariables
  id: clearAllVariables_hIst01
  variables: ConversationHistory
```

| Property | Type/values | Meaning |
|---|---|---|
| `variables` | `ConversationScopedVariables` (default) \| `ConversationHistory` \| `UserScopedVariables` \| `AllGlobalVariables` | Which variables to clear. The vendor reset topic clears both `ConversationScopedVariables` and `ConversationHistory` (two nodes). |

### EditTable

[EditTable](reference/schema/actions.md#edittable) changes a table (list) variable in place. **When to use it:** to add or remove rows, as in the
disambiguation template, which appends a "None of these" row:

```yaml
- kind: EditTable
  id: editTable_aDd001
  changeType: Add
  itemsVariable: Topic.IntentOptions
  value: "={ DisplayName: Topic.NoneOfTheseDisplayName, TopicId: \"NoTopic\", TriggerId: \"NoTrigger\", Score: 1.0 }"
```

| Property | Type/values | Req. | Meaning |
|---|---|---|---|
| `itemsVariable` | variable name | ✓ | The table to change. |
| `changeType` | `Add` (default) \| `Remove` \| `Clear` \| `TakeFirst` \| `TakeLast` | | The operation. |
| `value` | value / `=` expression | | The row to add or remove. |
| `resultVariable` | variable name | | No vendor guidance. By its name, receives the taken/removed item (inferred). |

### Other variable nodes

These have no vendor guidance. The shapes below come from the schema.

| Kind | Purpose (inferred from the schema) |
|---|---|
| [SetMultipleVariables](reference/schema/actions.md#setmultiplevariables) | Several assignments in one node: `assignments:` is a list of `{variable, value}`. |
| [ParseValue](reference/schema/actions.md#parsevalue) | Parse a value (typically JSON text) into a typed variable: `variable`, `value`, `valueType` (a [data type](06-variables-entities-powerfx.md#data-types)), `errorPolicy` (`DefaultBlank` default \| `ThrowError`). |
| [ResetVariable](reference/schema/actions.md#resetvariable) | Clear one variable (`variable`). |
| [EditTableV2](reference/schema/actions.md#edittablev2) | Newer form of `EditTable`: `changeType` is an object (`AddItemOperation` / `RemoveItemOperation` with `value`, `TakeFirstItemOperation`, `TakeLastItemOperation`, `ClearItemsOperation`). |

```yaml
# schema-valid shapes; not from the vendor repo
- kind: SetMultipleVariables
  id: setMultipleVariables_m1x2Q3
  assignments:
    - variable: Topic.Attempts
      value: =0
    - variable: Topic.Status
      value: pending

- kind: ParseValue
  id: parseValue_pJs0n1
  variable: Topic.Order
  value: =Topic.RawJson
  valueType:
    kind: Record
    properties:
      orderId: String
      total: Number
```

## Branching & loops

### ConditionGroup

[ConditionGroup](reference/schema/actions.md#conditiongroup) is if / else-if / else. Conditions are checked in order; the first true one runs its `actions`.
If none is true, `elseActions` runs.

```yaml
- kind: ConditionGroup
  id: conditionGroup_m2Nf7a
  conditions:
    - id: conditionItem_Qw3rTy
      condition: =System.FallbackCount < 3
      actions:
        - kind: SendActivity
          id: sendMessage_rPh01a
          activity: Could you rephrase that?

    - id: conditionItem_Zx8vBn
      condition: =!IsBlank(Global.LastDiscussedCity)
      actions:
        - kind: SendActivity
          id: sendMessage_cIty02
          activity: "Last time we discussed {Global.LastDiscussedCity}."

  elseActions:
    - kind: BeginDialog
      id: beginDialog_eSc003
      dialog: cr26e_myAgent.topic.Escalate
```

| Property | Type | Req. | Meaning |
|---|---|---|---|
| `conditions` | list of `{id, condition, actions}` | ✓ | Branches. **Each item needs its own unique `id`** and a `condition` (`=` expression). |
| `elseActions` | list of nodes | | Runs when no condition is true. |

A chain of conditions that only picks which topic to call can often be one `BeginDialog` with a `Switch()`
expression. See [BeginDialog](#begindialog).

### Loops and structure

None of these appear in the vendor repo. The descriptions come from the schema property names.

| Kind | Purpose (inferred) | Key properties |
|---|---|---|
| [Foreach](reference/schema/actions.md#foreach) | Loop over a table | `items` (`=` expression, required), `value` (loop variable, required), `index`, `actions` |
| [BreakLoop](reference/schema/actions.md#breakloop) | Leave the enclosing `Foreach` | — |
| [ContinueLoop](reference/schema/actions.md#continueloop) | Skip to the next iteration | — |
| [GotoAction](reference/schema/actions.md#gotoaction) | Jump to another node in the same topic | `actionId` (target node's `id`, required) |
| [ActionScope](reference/schema/actions.md#actionscope) | Group nodes into a block | `actions` |
| [DisconnectedNodeContainer](reference/schema/actions.md#disconnectednodecontainer) | Holds nodes not connected to the flow (canvas leftovers) | `actions` |

```yaml
# schema-valid shape; not from the vendor repo
- kind: Foreach
  id: foreach_lP9q2w
  items: =Topic.Results
  value: Topic.Item
  index: Topic.ItemIndex
  actions:
    - kind: SendActivity
      id: sendMessage_lP9q3e
      activity: "{Topic.ItemIndex}: {Topic.Item.Title}"
```

## Topic navigation

In the source, "dialog" means topic (or child agent). A topic is referenced by its schema name,
`<agentSchema>.topic.<TopicName>`. A [redirect](glossary.md#redirect) is a jump from one topic to another.

### BeginDialog

[BeginDialog](reference/schema/actions.md#begindialog) calls another topic and **comes back** when it ends. **When to use it:** escalation, shared
sub-flows, calling a child agent or an `OnSystemRedirect` helper.

```yaml
- kind: BeginDialog
  id: beginDialog_eSc4la
  dialog: cr26e_myAgent.topic.Escalate

- kind: BeginDialog
  id: beginDialog_sRch01
  dialog: cr26e_myAgent.topic.CustomSearch
  input:
    binding:
      query: =System.SearchQuery
  output:
    binding:
      results: Topic.RawResults
```

| Property | Type | Req. | Meaning |
|---|---|---|---|
| `dialog` | schema name, or `=` expression | ✓ | The topic to call. |
| `input` | [binding](#input-and-output-bindings) | | Values for the called topic's inputs. |
| `output` | [binding](#input-and-output-bindings) | | Where to store the called topic's outputs. |
| `activityProcessed` | boolean / expression, default `true` | | No vendor guidance. |

**Dynamic redirect.** `dialog:` accepts a Power Fx expression, so one node can replace a chain of condition
branches. This is from the `dynamic-topic-redirect` pattern:

```yaml
# fragment — dynamic dialog expression (see the note below on strict schema validators)
- kind: BeginDialog
  id: redirect_A4lDAn
  dialog: |-
    =Switch(
        Topic.LessonNumber,
        1, "cr26e_myAgent.topic.Lesson1",
        2, "cr26e_myAgent.topic.Lesson2",
        3, "cr26e_myAgent.topic.Lesson3",
        "cr26e_myAgent.topic.Fallback"
    )
```

The vendor pattern marks this as proven, and `schema-lookup.bundle.js validate` accepts it. A generic
JSON-Schema validator (such as the one these docs use to check their examples) rejects it, though. The schema's
[DialogExpression](reference/schema/other.md#dialogexpression) is a `oneOf` of "expression string" and "literal string", and an `=…` value matches both.

### EndDialog and CancelAllDialogs

| Kind | Purpose | Key properties |
|---|---|---|
| [EndDialog](reference/schema/actions.md#enddialog) | End the **current** topic and return to the caller, if any. | `clearTopicQueue` (default `true`), `value` |
| [CancelAllDialogs](reference/schema/actions.md#cancelalldialogs) | End **every** running topic. Error handlers and reset topics finish with it. | `eventName`, `eventValue`, `activityProcessed` |

```yaml
- kind: EndDialog
  id: endDialog_dOne01
  clearTopicQueue: true

- kind: CancelAllDialogs
  id: cancelAllDialogs_aLl001
```

### Other navigation nodes

| Kind | Purpose |
|---|---|
| [ReplaceDialog](reference/schema/actions.md#replacedialog) | Go to another topic and **don't come back**: the current topic is replaced. The disambiguation template uses it to hand over to Fallback. Same `dialog`/`input` as `BeginDialog`; no `output`. |
| [RepeatDialog](reference/schema/actions.md#repeatdialog) | Restart the current topic (inferred from the name). No vendor guidance. |
| [CancelDialog](reference/schema/actions.md#canceldialog) | Cancel the current topic, optionally raising `eventName`/`eventValue` (inferred). No vendor guidance. |
| [EndConversation](reference/schema/actions.md#endconversation) | End the whole conversation (inferred). No vendor guidance. |
| [RecognizeIntent](reference/schema/actions.md#recognizeintent) | Run intent recognition on `userInput` (inferred). No vendor guidance. |
| [DisableTrigger](reference/schema/actions.md#disabletrigger) | Disable the trigger of the topic named in `dialog` (inferred). No vendor guidance. |
| [ActivateExternalTrigger](reference/schema/actions.md#activateexternaltrigger) | Activate an external (autonomous) trigger: `dialog`, `externalTriggerId` (inferred). No vendor guidance; see [10](10-flows-and-autonomous-triggers.md). |

## Input and output bindings

`BeginDialog`, `ReplaceDialog` and the integration nodes (`InvokeConnectorAction`, `InvokeFlowAction`,
`InvokeAIBuilderModelAction`, …) pass values in and out through **bindings**.

| Shape | YAML | Meaning |
|---|---|---|
| Input map ([ActionInputBinding](reference/schema/actions.md#actioninputbinding), written without `kind`) | `input:`<br>`  binding:`<br>`    <inputName>: <value or =expr>` | One entry per input of the target. |
| Single input ([SingleValueInputBinding](reference/schema/actions.md#singlevalueinputbinding)) | `input:`<br>`  kind: SingleValueInputBinding`<br>`  value: =…` | One value (no vendor example). |
| Output map ([ActionOutputBinding](reference/schema/actions.md#actionoutputbinding), written without `kind`) | `output:`<br>`  binding:`<br>`    <outputName>: Topic.X` | Each output into a variable. **No `=`** on the variable. |
| Whole result ([SingleVariableOutputBinding](reference/schema/actions.md#singlevariableoutputbinding)) | `output:`<br>`  kind: SingleVariableOutputBinding`<br>`  variable: init:Topic.X` | The whole response (a record) into one variable. Read fields with `Topic.X.field`. |

```yaml
# InvokeConnectorAction with a single-variable output, then reading a field
- kind: InvokeConnectorAction
  id: getProfile_m3Pr0f
  connectionReference: shared_office365users
  connectionProperties:
    mode: Invoker
  operationId: UserGet_V2
  output:
    kind: SingleVariableOutputBinding
    variable: init:Topic.UserProfile

- kind: SetVariable
  id: setEmail_d4f5g6
  variable: init:Topic.UserEmail
  value: =Topic.UserProfile.mail

# InvokeAIBuilderModelAction with binding maps (model id is a placeholder GUID)
- kind: InvokeAIBuilderModelAction
  id: invokeAIBuilderModelAction_xxoZmJ
  aIModelId: 00000000-0000-0000-0000-000000000000
  input:
    binding:
      User_20Message: =System.Activity.Text
  output:
    binding:
      predictionOutput: Topic.ContentFilteringReason
```

Two binding forms in the vendor files **don't validate against the schema**:

- The vendor [SharePoint](glossary.md#sharepoint) (Microsoft's document-management service) and int-reference docs put connector parameters directly under `input:`
  (`input: { parameters/$filter: … }`). The schema needs them under `binding:`, as
  `input: binding: parameters/$filter: …`. See
  [08 `$`-prefixed inputs](08-tools-and-connectors.md#-prefixed-inputs-odata-parameters).
- The `prevent-child-agent-responses` pattern uses `output: kind: VariableOutputBindings` with
  `bindings: - kind: VariableOutputBinding`. Neither kind exists in the schema. Use `output: binding:` instead.

## AI & knowledge

All covered in [07 Knowledge](07-knowledge.md). Summary:

| Kind | Purpose |
|---|---|
| [SearchAndSummarizeContent](reference/schema/actions.md#searchandsummarizecontent) | [Generative answers](glossary.md#generative-answers) (an AI answer grounded in knowledge sources). Restrict it to specific [knowledge sources](glossary.md#knowledge-source) with `knowledgeSources:`; set `autoSend: false` plus `variable:` to capture the answer instead of sending it. |
| [AnswerQuestionWithAI](reference/schema/actions.md#answerquestionwithai) | AI answer from conversation history and general model knowledge only (not grounded in documents). |
| [SearchKnowledgeSources](reference/schema/actions.md#searchknowledgesources) | Search knowledge and return raw results (`result:`), with no AI summary. |
| [CreateSearchQuery](reference/schema/actions.md#createsearchquery) | Let the AI rewrite the user's message into a standalone search query (`result:`; read `.SearchQuery`). |
| [SearchAndSummarizeWithCustomModel](reference/schema/actions.md#searchandsummarizewithcustommodel) | Generative answers with a custom model configuration (`modelConfiguration`). No vendor guidance. |
| [InvokeCustomModelAction](reference/schema/actions.md#invokecustommodelaction) | Call a custom/external model (`modelConfiguration`). No vendor guidance. |
| [InvokeAIBuilderModelAction](reference/schema/actions.md#invokeaibuildermodelaction) | Run an [AI Builder](glossary.md#ai-builder) ([Power Platform](glossary.md#power-platform)'s prompt/model service) model by `aIModelId`. The model must be created in the UI; YAML can only reference it. Example above; see the [RAI pattern](13-patterns.md#rai-error-handling). |

The vendor author agent's rule: knowledge sources are searched automatically, so add `SearchAndSummarizeContent`
only when you need explicit control (for example, over the query). Use `AnswerQuestionWithAI` only for general,
ungrounded answers.

## Integration

| Kind | Purpose | Guide |
|---|---|---|
| [InvokeConnectorAction](reference/schema/actions.md#invokeconnectoraction) | Call a [connector](glossary.md#connector) (pre-built API wrapper) operation inline in a topic. Needs `connectionReference` (a [connection reference](glossary.md#connection-reference) logical name), `operationId` and optional `connectionProperties.mode` (`Maker`/`Invoker`; see [maker and invoker credentials](glossary.md#maker-and-invoker-credentials)). | [08](08-tools-and-connectors.md) |
| [InvokeFlowAction](reference/schema/actions.md#invokeflowaction) | Run a [Power Automate](glossary.md#power-automate) [cloud flow](glossary.md#cloud-flow) (`flowId`) with input/output bindings. | [10](10-flows-and-autonomous-triggers.md) |
| [InvokeMcpToolAction](reference/schema/actions.md#invokemcptoolaction) | Call one tool on an [MCP](glossary.md#mcp) server (`serverSchemaName`, `toolName`). | [09](09-mcp-and-multi-agent.md) |
| [InvokeSkillAction](reference/schema/actions.md#invokeskillaction) | Call a [Bot Framework](glossary.md#bot-framework) skill (`skillId`, `actionId`). No vendor guidance. | — |
| [HttpRequestAction](reference/schema/actions.md#httprequestaction) | Raw HTTP call: `method` (`Get` default, `Post`, `Patch`, `Put`, `Delete`), `url`, `headers`, `body`, `response` (variable), `responseSchema`, `errorHandling`, `requestTimeoutInMilliseconds` (default 30000). | [08](08-tools-and-connectors.md) |
| [WaitForConnectorTrigger](reference/schema/actions.md#waitforconnectortrigger) | Pause until a connector trigger fires (`connectionReference`, `operationId`). No vendor guidance. | [10](10-flows-and-autonomous-triggers.md) |

## Conversation control

None of these appear in the vendor repo. The purposes are inferred from the names and properties.

| Kind | Purpose (inferred) | Key properties |
|---|---|---|
| [TransferConversation](reference/schema/actions.md#transferconversation) | Hand the conversation to a human agent or phone number | `messageToAgent`, `context`, `telephoneNumber` |
| [TransferConversationV2](reference/schema/actions.md#transferconversationv2) | Newer transfer node | `transferType`: `kind: TransferToAgent` (`messageToAgent`, `context`) or `kind: TransferToPhoneNumber` (`phoneNumber`, SIP headers, `dtmfSequence`) |
| [SignOutUser](reference/schema/actions.md#signoutuser) | Sign the user out | — |
| [GetActivityMembers](reference/schema/actions.md#getactivitymembers) | Read the members of an activity | `variable`, `activityId` |
| [GetConversationMembers](reference/schema/actions.md#getconversationmembers) | Read the members of the conversation (e.g. a [Teams](glossary.md#teams) group chat) | `variable` |
| [EmitEvent](reference/schema/actions.md#emitevent) | Raise an event | `eventName` (required), `eventValue`, `eventType` (`ActivityIfUnhandled` default \| `Activity` \| `DialogEvent`), `bubbleEvent` |

## Telemetry

### LogCustomTelemetryEvent

[LogCustomTelemetryEvent](reference/schema/actions.md#logcustomtelemetryevent) writes a custom [telemetry](glossary.md#telemetry) event, viewable in
[Application Insights](glossary.md#application-insights) (Azure's monitoring service) if it's connected. **When to
use it:** in error handlers and at key business steps.

```yaml
- kind: LogCustomTelemetryEvent
  id: logTelemetry_eRr006
  eventName: OnErrorLog
  properties: "={ErrorMessage: System.Error.Message, ErrorCode: System.Error.Code, TimeUTC: Text(Now(), DateTimeFormat.UTC), ConversationId: System.Conversation.Id}"
```

| Property | Type | Meaning |
|---|---|---|
| `eventName` | string / expression | The event name to search for. |
| `properties` | `=` record expression | Custom dimensions. **Quote the whole value**, because it starts with `={`. |

## Sources

- `vendor/skills-for-copilot-studio/skills/add-node/SKILL.md`: common nodes, IDs, Question example
- `vendor/skills-for-copilot-studio/skills/int-reference/SKILL.md`: action-type table, Power Fx examples, `parameters/` input form
- `vendor/skills-for-copilot-studio/skills/new-topic/SKILL.md`: `SetTextVariable` for type coercion, inputs-vs-questions
- `vendor/skills-for-copilot-studio/skills/add-global-variable/SKILL.md`: `ConditionGroup` / `SetTextVariable` on globals
- `vendor/skills-for-copilot-studio/skills/add-adaptive-card/SKILL.md`: `AdaptiveCardPrompt` rules
- `vendor/skills-for-copilot-studio/skills/edit-action/sharepoint-actions.md`: `InvokeConnectorAction`, `SingleVariableOutputBinding`
- `vendor/skills-for-copilot-studio/agents/copilot-studio-author.md`: when to use `SearchAndSummarizeContent` / `AnswerQuestionWithAI`
- `vendor/skills-for-copilot-studio/templates/topics/*.topic.mcs.yml`: node examples
- `vendor/skills-for-copilot-studio/patterns/dynamic-topic-redirect.md`, `patterns/teams-production-hardening.md`, `patterns/rai-error-handling.md`, `patterns/line-breaks-in-messages.md`, `patterns/prevent-child-agent-responses.md`
- Generated reference: [actions](reference/schema/actions.md)
