# 04 — Topics & triggers

A [topic](glossary.md#topic) (a scripted conversation flow) is one `topics/*.topic.mcs.yml` file with
`kind: AdaptiveDialog` ([schema](reference/schema/dialogs.md#adaptivedialog)). This page covers the topic's root properties, how the
[orchestrator](glossary.md#orchestrator) (the LLM that routes messages) chooses it, topic inputs and outputs, and
all 28 [trigger](glossary.md#trigger) kinds (the event that starts a topic). It ends with complete worked examples.
The steps *inside* a topic are covered in [05 Nodes](05-nodes.md).

## Anatomy of a topic file

```yaml
# Name: Store Hours
kind: AdaptiveDialog
modelDescription: Use this topic when the user asks when a store is open, including weekend and holiday hours.
beginDialog:
  kind: OnRecognizedIntent
  id: main
  intent:
    displayName: Store Hours
    triggerQueries:
      - when are you open
      - opening hours
      - what time do you close
      - are you open on Sunday
      - holiday hours
  actions:
    - kind: SendActivity
      id: sendMessage_St0rH1
      activity: We're open 9am–5pm, Monday to Saturday.
```

- **`# Name:`** is a comment holding the topic's display name in the UI. Keep it as the first line.
- **`beginDialog:`** is the trigger: a mapping whose `kind:` is one of the 28 [trigger kinds](#trigger-kinds).
  Conventionally its `id` is `main`.
- **`beginDialog.actions:`** is the list of [nodes](glossary.md#node) (steps) to run, in order.
- Save the file as `topics/<TopicName>.topic.mcs.yml`. Its schema name is then
  `<agentSchema>.topic.<TopicName>` (for example `cr26e_myAgent.topic.StoreHours`). Other topics use that name to
  jump here (see [02 Project layout](02-project-layout.md#component-schema-names)).

## AdaptiveDialog root properties

| Property | Type/values | Req. | Meaning |
|---|---|---|---|
| `kind` | `AdaptiveDialog` | ✓ | |
| `beginDialog` | one of the [TriggerBase](reference/schema/triggers.md#triggerbase) kinds | | The trigger and its nodes. Every real topic has one. |
| `modelDescription` | string | | What the topic does and when to use it. This is the **routing signal** under [generative orchestration](glossary.md#generative-orchestration) (an LLM plans which topics and tools to run). See [below](#modeldescription-vs-trigger-phrases). |
| `modelDisplayName` | string | | The name the orchestrator sees for the topic. No vendor guidance; topics normally rely on the `# Name:` / intent `displayName`. |
| `inputs` | list of `AutomaticTaskInput` / `ManualTaskInput` | | Topic inputs the orchestrator fills in. See [Topic inputs and outputs](#topic-inputs-and-outputs). |
| `inputType` | [RecordNoKind](reference/schema/data-types.md#recordnokind) (`properties:` map) | | Type declarations for the inputs. Must match `inputs`. |
| `outputType` | [RecordNoKind](reference/schema/data-types.md#recordnokind) (`properties:` map) | | Declares the topic's output variables. |
| `startBehavior` | `Default` (default) \| `CancelOtherTopics` \| `UseLatestPublishedContentAndCancelOtherTopics` | | What happens to topics already running when this one starts. See [startBehavior](#startbehavior). |
| `endAutomatically` | boolean or `=` expression; default `true` | | No vendor guidance. By its name, whether the topic ends when its last node finishes (inferred). |
| `recognizer` | [Recognizer](reference/schema/dialogs.md#recognizer) | | A topic-level recognizer (`RegexRecognizer` with regex `intents`). No vendor guidance. |
| `triggers` | list of [TriggerBase](reference/schema/triggers.md#triggerbase) | | Extra triggers. No vendor guidance; every vendor topic uses a single `beginDialog`. |
| `confirmation`, `response` | [ResponseTemplateNoKind](reference/schema/messages.md#responsetemplatenokind) | | No vendor guidance. |
| `conversionMetadata` | object | | Written by the tools; leave it alone. |

`inputType: {}` and `outputType: {}` (empty records) appear on many vendor topics. The `add-other-agents` skill
says that on an `OnRedirect` topic they "signal it participates in multi-agent scenarios".

### startBehavior

The source repo uses `startBehavior: UseLatestPublishedContentAndCancelOtherTopics` on the error handler, on the
`OnSystemRedirect` reset topic, and on topics that handle app reinstalls. The `teams-production-hardening` pattern
explains that this makes "published content win over stale in-flight dialogs". In practice, the topic cancels
whatever else was running and uses the newest published version of the agent. The source has no example of
`CancelOtherTopics` on its own.

## modelDescription vs trigger phrases

A topic triggered by `OnRecognizedIntent` (the user expressed a particular [intent](glossary.md#intent), i.e. goal) can be reached in two ways. Which one matters depends on the orchestration
mode (`GenerativeActionsEnabled` in [settings.mcs.yml](03-agent.md#settingsmcsyml)):

| | `modelDescription` | [Trigger phrases](glossary.md#trigger-phrases) (`intent.triggerQueries`) |
|---|---|---|
| Where | Root of the `AdaptiveDialog` | `beginDialog.intent.triggerQueries` |
| Used by | **Generative orchestration** (primary routing signal) | **[Classic orchestration](glossary.md#classic-orchestration)** (pattern-matched against the user's message); secondary hints under generative |
| Form | One or two natural-language sentences | 5–10 short example utterances |

**Writing a good `modelDescription`** (from the `edit-triggers` skill; see also
[model description](glossary.md#model-description)):

- Write a clear, specific sentence saying what the topic does **and when to use it**.
- Focus on the **user's intent**: "Use this topic when the user asks for information about our perfumes", not "This
  topic sends a message".
- Name the **key scenarios**: "This topic handles billing inquiries, refund requests, and payment issues".
- Avoid vague text such as "This topic handles requests".
- Always add one to `OnRecognizedIntent` topics when `GenerativeActionsEnabled: true`, and update it when the
  topic's scope changes.
- System-triggered topics (`OnError`, `OnUnknownIntent`, `OnConversationStart`, …) don't need one.

**Writing good trigger phrases:**

- Use **5–10 phrases** per topic.
- Cover synonyms and variations ("help", "assist", "support", "I need help").
- Mix sentence structures (questions, statements, keywords) and formal/informal tone.
- Keep them short and natural, the way users actually type or speak.
- **Avoid overlap** between topics. Close matches trigger [disambiguation](glossary.md#disambiguation) (the "did
  you mean…?" prompt).

Only `OnRecognizedIntent` topics have editable trigger phrases. You can't change a topic's trigger *kind* by editing
the phrases; to change the kind, recreate the topic.

The [Intent](reference/schema/dialogs.md#intent) object under `intent:` also accepts `displayName`, `includeInOnSelectIntent` (default `true`: whether
the topic is offered in disambiguation) and `availableInChannels`. The vendor files only use `displayName` and
`triggerQueries`. A topic with only inputs and a description can use `intent: {}`, as in the
[arithmetic example](#inputs-and-outputs-arithmetic).

## Topic inputs and outputs

Under generative orchestration, a topic can behave like a function. The orchestrator fills its **inputs** from the
conversation, asking only for what's missing, and reads its **outputs** afterwards to compose the reply. The
`new-topic` and `add-node` skills give these rules:

- **Prefer topic inputs over `Question` nodes** for information the topic always needs. Use `Question` nodes only
  for *conditional* questions ("ask for a phone number only if the group is larger than 6") or end-of-flow
  confirmations.
- **Prefer topic outputs over `SendActivity`** for final results. The orchestrator writes the user-facing message
  from the outputs. A `SendActivity` is for precise mid-flow messages.
- Put `inputs`, `inputType` and `outputType` at the **AdaptiveDialog root**, *not* inside `beginDialog`.
- Declare every input in **both** `inputs:` and `inputType.properties:`, with matching names.
- Set outputs with `SetVariable` on `Topic.<outputName>`. **Don't** use `TaskOutput`; it's only valid in tool
  (`TaskDialog`) files.

### `AutomaticTaskInput`

| Property | Type/values | Req. | Meaning |
|---|---|---|---|
| `kind` | `AutomaticTaskInput` | ✓ | The orchestrator fills the value. (`ManualTaskInput` with a fixed `value:` also exists, but it's meant for tools; see [08](08-tools-and-connectors.md).) |
| `propertyName` | string | ✓ | The input's name. Inside the topic it's `Topic.<propertyName>`. |
| `description` | string | | What to extract. The orchestrator uses this as its prompt. |
| `entity` | [entity](06-variables-entities-powerfx.md#entities) reference | | The expected value type, e.g. `NumberPrebuiltEntity`. |
| `shouldPromptUser` | boolean, default `true` | | `false` = never ask; the orchestrator infers the value silently. |
| `defaultValue` | value | | Used when no value is found. |
| `inputSettings` | [QuestionInputSettingsNoKind](reference/schema/task-actions.md#questioninputsettingsnokind) | | Prompt, validation and retry settings, like a `Question` node. No vendor guidance. |

**Orchestrator-generated variables.** With `shouldPromptUser: false` and a precise `description`, an input becomes
a free classifier or extractor. The orchestrator fills it as part of the call it already makes to pick the topic, so
there's no extra cost or latency. List the exact allowed values in the description, keep to 3–6 categories, and
handle a fallback value. The value is filled once, when the topic is invoked:

```yaml
kind: AdaptiveDialog
inputs:
  - kind: AutomaticTaskInput
    propertyName: searchCategory
    description: |-
      Classify the user's query into one of these categories:
      HR, IT, Finance, Other
    shouldPromptUser: false

beginDialog:
  kind: OnRecognizedIntent
  id: main
  intent: {}
  actions:
    - kind: SetVariable
      id: setVariable_cAt3gY
      variable: Global.searchCategory
      value: =Topic.searchCategory

inputType:
  properties:
    searchCategory:
      displayName: searchCategory
      description: |-
        Classify the user's query into one of these categories:
        HR, IT, Finance, Other
      type: String

outputType: {}
```

### `inputType` / `outputType` properties

Each entry under `properties:` is keyed by the variable name and takes `displayName`, `description` and `type`.
`type` is a [data type](06-variables-entities-powerfx.md#data-types) such as `String`, `Number`, `Boolean`,
`DateTime`, or a `Record`/`Table` with nested `properties:`. Other optional keys from [PropertyInfoNoKind](reference/schema/data-types.md#propertyinfonokind)
include `isRequired` (used on tool inputs in the source).

### Topic-action chaining

When a topic sits alongside [tools](glossary.md#tool) or other topics, its **outputs decide whether the orchestrator
keeps going**. The `new-topic` skill distinguishes two cases:

1. **The topic is self-contained** (for example, it confirms a reservation or computes a result). Output a
   confirmation or status. The orchestrator treats the task as done, which is correct.
2. **The topic gathers data for a tool to use** (for example, it collects a complaint that a "send to [Teams](glossary.md#teams)" (Microsoft's chat app) tool
   should post). Output the **data itself**, not a confirmation. If you output "Your complaint has been sent!", the
   orchestrator thinks the job is finished and never calls the tool.

Ask: *does this topic finish the task, or prepare data for something else?* If it prepares data, output the data.
For data-gathering, there are three alternatives:

- **Drop the topic.** Tailor the tool's own input descriptions so the tool collects the data. This is the simplest
  option for one-shot collection.
- **Use a [global variable](glossary.md#global-variable)** (a conversation-wide variable). The topic fills it,
  possibly over several questions, and the tool's inputs read it with [Power Fx](glossary.md#power-fx) (the formula language). Use this for multi-turn gathering or
  when you need logic in between.
- **Expose the data as a topic output** so the orchestrator passes it to the next tool or topic.

## Trigger kinds

All 28 kinds in the [TriggerBase](reference/schema/triggers.md#triggerbase) union. Every vendor topic sets exactly one, under `beginDialog:`.

### Properties every trigger has

| Property | Type | Req. | Meaning |
|---|---|---|---|
| `id` | string | ✓ | Conventionally `main`. |
| `actions` | list of nodes | | The topic body. See [05 Nodes](05-nodes.md). |
| `condition` | `=` expression (or boolean) | | Extra guard: the trigger only fires when this is true. For example, `=IsBlank(Global.UserCountry)` makes an `OnActivity` topic fire once; `=System.Activity.ChannelId = "msteams"` limits it to Teams. |
| `priority` | number, default `0` | | Ordering between topics that could fire on the same event. The vendor fallback and search topics use `-1`; a "low-priority" initialisation topic uses `-2`. |
| `runOnce` | boolean, default `false` | | No vendor guidance. By its name, the trigger fires at most once per conversation (inferred). |

Kind-specific properties are listed with each kind below. `triggerBehavior` exists **only** on `OnSelectIntent`.

### Intent triggers

| `kind:` | Fires when | Extra properties | Notes |
|---|---|---|---|
| [OnRecognizedIntent](reference/schema/triggers.md#onrecognizedintent) | The user's message matches the topic (by `modelDescription` under generative orchestration, or trigger phrases under classic) | `intent:` (`displayName`, `triggerQueries`, …) | The most common trigger. |
| [OnSystemIntent](reference/schema/triggers.md#onsystemintent) | — | `intent:` | No guidance in the source repo. |
| [OnSelectIntent](reference/schema/triggers.md#onselectintent) | Several topics matched closely ("Multiple Topics Matched" [system topic](glossary.md#system-topic)) | `triggerBehavior`: `Default` \| `Always` | Disambiguation. The vendor template sets `triggerBehavior: Always`. See [example](#disambiguation). |
| [OnUnknownIntent](reference/schema/triggers.md#onunknownintent) | No topic matched (the [fallback](glossary.md#fallback)) | — | Commonly `priority: -1`. See [example](#fallback). |
| [OnEscalate](reference/schema/triggers.md#onescalate) | The user asks for a human | `intent:` | The int-reference table lists it; no template. |

### Conversation lifecycle triggers

| `kind:` | Fires when | Extra properties | Notes |
|---|---|---|---|
| [OnConversationStart](reference/schema/triggers.md#onconversationstart) | A conversation begins | — | Greeting topic. **Doesn't fire in [Microsoft 365 Copilot](glossary.md#microsoft-365-copilot) (Microsoft's assistant in Office apps) or other channel-embedded surfaces**, so use `OnActivity` for initialisation (see below). |
| [OnActivity](reference/schema/triggers.md#onactivity) | An incoming [activity](glossary.md#activity) (any message, event or update sent in a conversation) of the given `type` arrives | `type`: an [ActivityType](reference/schema/messages.md#activitytype) (`Message`, `Event`, `InstallationUpdate`, `ConversationUpdate`, `Invoke`, `Typing`, …) | With `type: Message` and an `IsBlank(Global.X)` condition, the topic fires once on the first real user message. This is the vendor "just-in-time" initialisation pattern. With `type: InstallationUpdate`, it handles Teams app reinstalls. |
| [OnInactivity](reference/schema/triggers.md#oninactivity) | The conversation has been idle | `durationInSeconds`: `900`, `1800` (default), `2700`, `3600`, or an expression | The Teams-hardening pattern clears variables and history here, and also sets `durationInSeconds: 43200`, a value outside the schema's listed choices (it validates in the vendor pattern). |
| [OnPlanComplete](reference/schema/triggers.md#onplancomplete) | — | — | No guidance in the source repo. By its name, after the generative orchestrator finishes a plan (inferred). |

### System triggers

| `kind:` | Fires when | Extra properties | Notes |
|---|---|---|---|
| [OnError](reference/schema/triggers.md#onerror) | A runtime error occurs | — | Read `System.Error.Message` / `.Code` / `.SubCode`. See [example](#error-handler) and the [RAI error-handling pattern](13-patterns.md#rai-error-handling). |
| [OnSignIn](reference/schema/triggers.md#onsignin) | Sign-in is required | — | Read `System.SignInReason`. See [example](#sign-in). |
| [OnSystemRedirect](reference/schema/triggers.md#onsystemredirect) | **Only** when another topic jumps here (`BeginDialog` / `ReplaceDialog`) | — | For helper topics with no user-facing trigger, such as a "reset conversation" topic. |
| [OnRedirect](reference/schema/triggers.md#onredirect) | This agent is called by another agent as a [connected agent](glossary.md#connected-agent) | — | Entry point on the *called* side. Declare `inputType: {}` / `outputType: {}` and read inputs from `System.Activity.Value`. See [09 MCP & multi-agent](09-mcp-and-multi-agent.md). |
| [OnCopilotRedirect](reference/schema/triggers.md#oncopilotredirect) | — | `askForConfirmation`, `conversationStarters` | No guidance in the source repo. |

### Event and interception triggers

| `kind:` | Fires when | Extra properties | Notes |
|---|---|---|---|
| [OnEventActivity](reference/schema/triggers.md#oneventactivity) | An event activity with the given name arrives | `eventName` | No guidance in the source repo beyond the property name. |
| [OnAnalyzeText](reference/schema/triggers.md#onanalyzetext) | — | — | No guidance in the source repo. |
| [OnToolSelected](reference/schema/triggers.md#ontoolselected) | The parent orchestrator routes to this [child agent](glossary.md#child-agent) (a sub-agent the main agent delegates to) | `description` (**required**): when to route here | Used as the `beginDialog` of `AgentDialog` files (`agents/<name>/agent.mcs.yml`), not of topics. See [09](09-mcp-and-multi-agent.md). |
| [OnKnowledgeRequested](reference/schema/triggers.md#onknowledgerequested) | The agent is about to search knowledge | — | **YAML-only** (not shown in the UI). Lets you supply custom search results: read `System.SearchQuery` / `System.KeywordSearchQuery` and set `System.SearchResults`. See [07 Knowledge](07-knowledge.md). |
| [OnGeneratedResponse](reference/schema/triggers.md#ongeneratedresponse) | An AI-generated answer is about to be sent | — | Read `System.Response.FormattedText`; set `System.ContinueResponse` to `false` to suppress the original and send your own. See [example](#intercepting-generated-answers). |
| [OnOutgoingMessage](reference/schema/triggers.md#onoutgoingmessage) | — | — | **Non-functional (as of 2026-03-15).** It exists in the schema but doesn't fire at runtime. Don't use it. |

### Voice ([IVR](glossary.md#ivr)) triggers

For phone agents. The source repo has no guidance or examples for any of these. The descriptions are inferred from
the names.

| `kind:` | Fires when (inferred) | Extra properties |
|---|---|---|
| [OnDtmfKeyPress](reference/schema/triggers.md#ondtmfkeypress) | The caller presses a given [DTMF](glossary.md#dtmf) (telephone keypad) key | `dtmfKey` (**required**): `Num0`…`Num9`, `PoundKey`, `StarKey` |
| [OnUnknownDtmfKey](reference/schema/triggers.md#onunknowndtmfkey) | The caller presses an unmapped key | — |
| [OnSilence](reference/schema/triggers.md#onsilence) | The caller says nothing for a while | `silenceDetectionTimeoutInMilliseconds` |
| [OnUnrecognizedSpeech](reference/schema/triggers.md#onunrecognizedspeech) | Speech can't be recognised | — |
| [OnAnsweringMachineDetected](reference/schema/triggers.md#onansweringmachinedetected) | An outbound call reaches voicemail | — |

### Scheduled and connector triggers

These start an [autonomous agent](glossary.md#autonomous-agent) (one that runs on events or schedules without a
user). The author agent refuses to create autonomous triggers in YAML, because they need [Power Platform](glossary.md#power-platform) (Microsoft's low-code suite) setup beyond
YAML. Create them in the UI, then edit them. See [10 Flows & autonomous triggers](10-flows-and-autonomous-triggers.md).

| `kind:` | Fires when | Extra properties |
|---|---|---|
| [RecurrenceTrigger](reference/schema/triggers.md#recurrencetrigger) | On a schedule | — (no schedule fields on the trigger itself) |
| [CreateConversationConnectorTrigger](reference/schema/triggers.md#createconversationconnectortrigger) | A [connector](glossary.md#connector) (a pre-built wrapper around an external API) event starts a new conversation | `connectionReference`, `connectorTriggerId` (both required), `connectionProperties`, `dynamicInputSchema`, `dynamicOutputSchema` |
| [ContinueConversationConnectorTrigger](reference/schema/triggers.md#continueconversationconnectortrigger) | A connector event continues an existing conversation | same as above |

## Worked examples

All adapted from the vendor templates, with placeholder IDs replaced. Each is a complete, valid topic file.

### Greeting

`OnConversationStart` with message variations. One `text:` entry is picked at random each time:

```yaml
# Name: Greeting
kind: AdaptiveDialog
beginDialog:
  kind: OnConversationStart
  id: main
  actions:
    - kind: SendActivity
      id: sendMessage_gR7kQ2
      activity:
        text:
          - Hello! I'm {System.Bot.Name}. How can I help you today?
          - Hi there! Welcome! What can I assist you with?
          - Greetings! I'm here to help. What would you like to know?
```

Because `OnConversationStart` doesn't fire in Microsoft 365 Copilot, put per-conversation setup in an `OnActivity`
topic instead:

```yaml
# Name: Conversation Init
kind: AdaptiveDialog
beginDialog:
  kind: OnActivity
  id: main
  type: Message
  condition: =IsBlank(Global.UserCountry)
  actions:
    - kind: InvokeConnectorAction
      id: getProfile_m3Pr0f
      connectionReference: shared_office365users
      connectionProperties:
        mode: Invoker
      operationId: UserGet_V2
      output:
        kind: SingleVariableOutputBinding
        variable: init:Topic.M365Profile

    - kind: SetVariable
      id: setCountry_Ct9yX1
      variable: Global.UserCountry
      value: =If(IsBlank(Topic.M365Profile.country), "Unknown", Topic.M365Profile.country)
```

### Question and branching

Ask, then branch on the answer. `alwaysPrompt: true` asks even if the variable already has a value.
`allowInterruption: false` stops the user's reply from triggering another topic:

```yaml
# Name: Plan Info
kind: AdaptiveDialog
modelDescription: Use this topic when the user wants to compare or learn about our subscription plans.
beginDialog:
  kind: OnRecognizedIntent
  id: main
  intent:
    displayName: Plan Info
    triggerQueries:
      - tell me about your plans
      - which plan should I choose
      - pricing tiers
      - what plans do you offer

  actions:
    - kind: SendActivity
      id: sendMessage_pL4n01
      activity:
        text:
          - I can tell you about our Basic and Premium plans.

    - kind: Question
      id: question_pL4n02
      variable: init:Topic.UserChoice
      prompt: Which plan are you interested in, Basic or Premium?
      entity: StringPrebuiltEntity
      alwaysPrompt: true
      interruptionPolicy:
        allowInterruption: false

    - kind: ConditionGroup
      id: conditionGroup_pL4n03
      conditions:
        - id: conditionItem_pL4n04
          condition: =Topic.UserChoice = "Basic"
          actions:
            - kind: SendActivity
              id: sendMessage_pL4n05
              activity: Basic includes 10 projects and email support.

        - id: conditionItem_pL4n06
          condition: =Topic.UserChoice = "Premium"
          actions:
            - kind: SendActivity
              id: sendMessage_pL4n07
              activity: Premium includes unlimited projects and phone support.

      elseActions:
        - kind: SendActivity
          id: sendMessage_pL4n08
          activity: Sorry, I only know about the Basic and Premium plans.
```

(The vendor template compares with single quotes, `= '<Option 1>'`. In Power Fx, text literals use double quotes;
single quotes delimit *names*. See [06 Power Fx](06-variables-entities-powerfx.md#power-fx).) A closed-list
[entity](glossary.md#entity) is usually better than free text for a fixed set of choices; see
[06 Entities](06-variables-entities-powerfx.md#entities).

### Fallback

`OnUnknownIntent` asks the user to rephrase twice, then escalates. `System.FallbackCount` counts consecutive
fallbacks:

```yaml
# Name: Fallback
kind: AdaptiveDialog
beginDialog:
  kind: OnUnknownIntent
  id: main
  priority: -1
  actions:
    - kind: ConditionGroup
      id: conditionGroup_fB1k01
      conditions:
        - id: conditionItem_fB1k02
          condition: =System.FallbackCount < 3
          actions:
            - kind: SendActivity
              id: sendMessage_fB1k03
              activity:
                text:
                  - I'm not sure I understand. Could you rephrase that?
                  - I didn't quite get that. Can you try asking in a different way?

      elseActions:
        - kind: SendActivity
          id: sendMessage_fB1k04
          activity:
            text:
              - I'm having trouble understanding your request. Let me connect you with someone who can help.

        - kind: BeginDialog
          id: beginDialog_fB1k05
          dialog: cr26e_myAgent.topic.Escalate
```

The vendor `search-topic` template is an alternative fallback. It runs `CreateSearchQuery` and then
`SearchAndSummarizeContent` over all knowledge. See [07 Knowledge](07-knowledge.md).

### Error handler

`OnError` shows details in the test pane, shows a short message (with a spoken variant) to real users, logs to
[telemetry](glossary.md#telemetry) (custom events sent to [Application Insights](glossary.md#application-insights), Azure's monitoring service) and cancels everything:

```yaml
# Name: On Error
kind: AdaptiveDialog
startBehavior: UseLatestPublishedContentAndCancelOtherTopics
beginDialog:
  kind: OnError
  id: main
  actions:
    - kind: SetVariable
      id: setVariable_eRr001
      variable: init:Topic.CurrentTime
      value: =Text(Now(), DateTimeFormat.UTC)

    - kind: ConditionGroup
      id: conditionGroup_eRr002
      conditions:
        - id: conditionItem_eRr003
          condition: =System.Conversation.InTestMode = true
          actions:
            - kind: SendActivity
              id: sendMessage_eRr004
              activity: |-
                Error Message: {System.Error.Message}
                Error Code: {System.Error.Code}
                Conversation Id: {System.Conversation.Id}
                Time (UTC): {Topic.CurrentTime}

      elseActions:
        - kind: SendActivity
          id: sendMessage_eRr005
          activity:
            text:
              - |-
                An error has occurred.
                Error code: {System.Error.Code}
                Conversation Id: {System.Conversation.Id}
                Time (UTC): {Topic.CurrentTime}.
            speak:
              - An error has occurred, please try again.

    - kind: LogCustomTelemetryEvent
      id: logTelemetry_eRr006
      eventName: OnErrorLog
      properties: "={ErrorMessage: System.Error.Message, ErrorCode: System.Error.Code, TimeUTC: Topic.CurrentTime, ConversationId: System.Conversation.Id}"

    - kind: CancelAllDialogs
      id: cancelAllDialogs_eRr007
```

### Disambiguation

`OnSelectIntent` runs when several topics match. This version adds a "None of these" option to the candidate list,
asks the user to pick one with a dynamic list [entity](06-variables-entities-powerfx.md#entities), and goes to the
fallback if they choose "None of these". Assigning the choice to `System.Recognizer.SelectedIntent` makes the
runtime run the chosen topic:

```yaml
# Name: Multiple Topics Matched
kind: AdaptiveDialog
beginDialog:
  kind: OnSelectIntent
  id: main
  triggerBehavior: Always
  actions:
    - kind: SetVariable
      id: setVariable_dSm001
      variable: init:Topic.IntentOptions
      value: =System.Recognizer.IntentOptions

    - kind: SetTextVariable
      id: setTextVariable_dSm002
      variable: Topic.NoneOfTheseDisplayName
      value: None of these

    - kind: EditTable
      id: editTable_dSm003
      changeType: Add
      itemsVariable: Topic.IntentOptions
      value: "={ DisplayName: Topic.NoneOfTheseDisplayName, TopicId: \"NoTopic\", TriggerId: \"NoTrigger\", Score: 1.0 }"

    - kind: Question
      id: question_dSm004
      interruptionPolicy:
        allowInterruption: false
      alwaysPrompt: true
      variable: System.Recognizer.SelectedIntent
      prompt: "To clarify, did you mean:"
      entity:
        kind: DynamicClosedListEntity
        items: =Topic.IntentOptions

    - kind: ConditionGroup
      id: conditionGroup_dSm005
      conditions:
        - id: conditionItem_dSm006
          condition: =System.Recognizer.SelectedIntent.TopicId = "NoTopic"
          actions:
            - kind: ReplaceDialog
              id: replaceDialog_dSm007
              dialog: cr26e_myAgent.topic.Fallback
```

### Sign-in

`OnSignIn` explains why sign-in is needed (when the reason is `SignInRequired`), then shows the sign-in card with
`OAuthInput`:

```yaml
# Name: Sign In
kind: AdaptiveDialog
beginDialog:
  kind: OnSignIn
  id: main
  actions:
    - kind: ConditionGroup
      id: conditionGroup_sIgn01
      conditions:
        - id: conditionItem_sIgn02
          condition: =System.SignInReason = SignInReason.SignInRequired
          actions:
            - kind: SendActivity
              id: sendMessage_sIgn03
              activity: Hello! To be able to help you, I'll need you to sign in.

    - kind: OAuthInput
      id: oauthInput_sIgn04
      title: Login
      text: To continue, please login
```

### Inputs and outputs (arithmetic)

A topic the orchestrator calls like a function. It has two number inputs (prompted for if missing) and one text
output, and no `SendActivity`: the orchestrator writes the reply from `result`. The trigger has `intent: {}`, so
routing relies entirely on `modelDescription`.

```yaml
# Name: Arithmetic Sum
kind: AdaptiveDialog
inputs:
  - kind: AutomaticTaskInput
    propertyName: number1
    description: This is the first number for the sum
    entity: NumberPrebuiltEntity
    shouldPromptUser: true

  - kind: AutomaticTaskInput
    propertyName: number2
    description: This is the second number for the sum
    entity: NumberPrebuiltEntity
    shouldPromptUser: true

modelDescription: This topic adds the two numbers given as input and returns the sum.
beginDialog:
  kind: OnRecognizedIntent
  id: main
  intent: {}
  actions:
    - kind: SetVariable
      id: setVariable_aRt001
      variable: Topic.result
      value: =Text(Topic.number1 + Topic.number2)

inputType:
  properties:
    number1:
      displayName: number1
      description: This is the first number for the sum
      type: Number

    number2:
      displayName: number2
      description: This is the second number for the sum
      type: Number

outputType:
  properties:
    result:
      displayName: result
      description: This is the result of the sum, converted to text
      type: String
```

(The vendor template's `modelDescription` says "performs the sub". The intent is "sum".)

### Intercepting generated answers

`OnGeneratedResponse` suppresses the AI answer, strips citation markers `[1]`…`[5]` and sends the cleaned text:

```yaml
# Name: Remove Citations
kind: AdaptiveDialog
beginDialog:
  kind: OnGeneratedResponse
  id: main
  actions:
    - kind: SetVariable
      id: setVariable_rCt001
      variable: System.ContinueResponse
      value: =false

    - kind: SetVariable
      id: setVariable_rCt002
      variable: init:Topic.CleanedText
      value: "=Substitute(Substitute(Substitute(Substitute(Substitute(System.Response.FormattedText, \"[1]\", \"\"), \"[2]\", \"\"), \"[3]\", \"\"), \"[4]\", \"\"), \"[5]\", \"\")"

    - kind: SendActivity
      id: sendActivity_rCt003
      activity: "{Topic.CleanedText}"
```

### Inactivity reset (Teams)

From the Teams-hardening pattern. After 12 hours idle in Teams, it clears history and variables, sets a flag that
another `OnActivity` topic reads to say "your session expired", and cancels running topics:

```yaml
kind: AdaptiveDialog
beginDialog:
  kind: OnInactivity
  id: main
  condition: =System.Activity.ChannelId = "msteams"
  durationInSeconds: 43200
  actions:
    - kind: ClearAllVariables
      id: clearAllVariables_iNa001
      variables: ConversationHistory

    - kind: ClearAllVariables
      id: clearAllVariables_iNa002

    - kind: SetVariable
      id: setVariable_iNa003
      variable: Global.InactiveConversation
      value: true

    - kind: CancelAllDialogs
      id: cancelAllDialogs_iNa004
inputType: {}
outputType: {}
```

## Listing an agent's topics

The `list-topics` skill summarises an agent by reading each `topics/*.topic.mcs.yml` and pulling out: the name (from
`# Name:` or the file name), `beginDialog.kind`, `triggerQueries` (for `OnRecognizedIntent`), the number of
`actions`, and every `BeginDialog`/`ReplaceDialog` target. Topics with `OnError`, `OnSelectIntent` and `OnSignIn`
are system topics; change them with care.

## Sources

- `vendor/skills-for-copilot-studio/skills/new-topic/SKILL.md`: topic structure, inputs/outputs, topic-action chaining
- `vendor/skills-for-copilot-studio/skills/edit-triggers/SKILL.md`: `modelDescription` and trigger-phrase guidance
- `vendor/skills-for-copilot-studio/skills/list-topics/SKILL.md`: topic summary fields
- `vendor/skills-for-copilot-studio/skills/int-reference/SKILL.md`: trigger table, `OnOutgoingMessage` status, YAML-only features
- `vendor/skills-for-copilot-studio/skills/add-node/SKILL.md`: inputs-vs-questions rule
- `vendor/skills-for-copilot-studio/skills/add-other-agents/SKILL.md`: `OnToolSelected`, `OnRedirect`
- `vendor/skills-for-copilot-studio/agents/copilot-studio-author.md`: autonomous-trigger limitation
- `vendor/skills-for-copilot-studio/patterns/orchestrator-variables.md`: orchestrator-filled inputs
- `vendor/skills-for-copilot-studio/patterns/jit-user-context.md`, `patterns/jit-glossary.md`: `OnActivity` instead of `OnConversationStart`
- `vendor/skills-for-copilot-studio/patterns/teams-production-hardening.md`: `startBehavior`, `OnInactivity`, `OnActivity` types, `priority`
- `vendor/skills-for-copilot-studio/templates/topics/*.topic.mcs.yml`: greeting, question, fallback, error handler, disambiguation, sign-in, arithmetic, remove-citations, conversation-init, search
- Generated reference: [triggers](reference/schema/triggers.md), [dialogs](reference/schema/dialogs.md)
