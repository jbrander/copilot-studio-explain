# 03 — The agent

This page covers the two files that define the agent as a whole: `agent.mcs.yml` (its name,
[instructions](glossary.md#instructions) (the system prompt), conversation starters and model choice) and
`settings.mcs.yml` (its schema name, sign-in mode, orchestration mode, recognizer and language). It also lists the
agent-level settings blocks the schema defines but the source repo doesn't document.

## agent.mcs.yml

`agent.mcs.yml` sits at the root of the agent folder and has `kind: GptComponentMetadata` ([schema](reference/schema/components.md#gptcomponentmetadata)). This is the vendor
template, with the placeholders filled in:

```yaml
mcs.metadata:
  componentName: Contoso Support
kind: GptComponentMetadata
displayName: Contoso Support
instructions: |
  You are a helpful assistant that provides accurate and friendly responses.

  Guidelines:
  - Be concise and clear in your responses
  - Ask clarifying questions when needed
  - Escalate to a human agent when you cannot help
conversationStarters:
  - title: Get Started
    text: How can you help me?

  - title: FAQ
    text: What questions can you answer?

aISettings:
  model:
    modelNameHint: GPT5Chat
```

`displayName` at the root is part of the on-disk flattening described in
[02 Project layout](02-project-layout.md#on-disk-shape-vs-schema-shape): it isn't a `GptComponentMetadata`
property, but the authoring tools expect it there. Keep the `mcs.metadata:` block as you found it.

> **You can't create this file from nothing.** The skills refuse to scaffold `agent.mcs.yml` or `settings.mcs.yml`.
> Create the agent in the Copilot Studio UI, clone it, then edit these files. See
> [02 Project layout](02-project-layout.md#lifecycle-clone--edit--validate--push--publish).

### Properties

| Property | Type/values | Req. | Meaning |
|---|---|---|---|
| `displayName` | string | | The agent's name as users see it (on-disk only; see above). |
| `instructions` | text template | | The system prompt. Supports `{…}` [Power Fx](glossary.md#power-fx) (the agent's formula language) interpolation, e.g. `{Today()}`. See [Writing instructions](#writing-instructions). |
| `responseInstructions` | text template | | Present in the schema; the source repo has no guidance on it. By its name, extra guidance for how responses are phrased (inferred). |
| `conversationStarters` | list of `{title, text}` | | [Conversation starters](glossary.md#conversation-starter) (suggested opening prompts). Both `title` and `text` are required per item. |
| `aISettings` | [AISettingsNoKind](reference/schema/settings.md#aisettingsnokind) | | Model choice and AI switches. See [Model and AI settings](#model-and-ai-settings). |
| `gptCapabilities` | [GptCapabilitiesNoKind](reference/schema/components.md#gptcapabilitiesnokind) | | Built-in capabilities (web browsing, code interpreter, image generation, Microsoft 365 searches). See [Capabilities](#capabilities-gptcapabilities). |
| `knowledgeSources` | [SearchKnowledgeSourcesMode](reference/schema/knowledge.md#searchknowledgesourcesmode) | | Which [knowledge sources](glossary.md#knowledge-source) the agent searches: `kind: SearchAllKnowledgeSources` or `kind: SearchSpecificKnowledgeSources` with a list of knowledge schema names. No vendor guidance at the agent level; see [07 Knowledge](07-knowledge.md). |
| `defaultResponseMode` | `Auto` \| `ThinkDeeper` \| `QuickResponse` | | No guidance in the source repo. |
| `inputType`, `outputType`, `historyType`, `tools`, `fileSearchDataSource` | | | Present in the schema; not used by the vendor templates or skills. See [GptComponentMetadata](reference/schema/components.md#gptcomponentmetadata). |

### Model and AI settings

`aISettings.model` selects the LLM. The `edit-agent` skill's rule is to use the **no-`kind`** form
([CurrentModelsNoKind](reference/schema/settings.md#currentmodelsnokind)) and fill in only `modelNameHint` and, for non-OpenAI models, `provider`:

```yaml
kind: GptComponentMetadata
displayName: Contoso Support
aISettings:
  model:
    modelNameHint: Sonnet46       # model hint; GPT5Chat is the template default
    provider: Anthropic           # required for non-OpenAI models
```

| `aISettings.model` property | Type/values | Meaning |
|---|---|---|
| `modelNameHint` | string | Which model to use. The vendor examples are `GPT5Chat` and `Sonnet46`; the schema doesn't enumerate the valid hints. The skills tell you to run `schema-lookup.bundle.js models` before choosing, but that command only prints the model-related schema definitions, not a list of hints. |
| `provider` | `OpenAI` \| `AzureOpenAI` \| `Anthropic` \| `GoogleVertexAI` | The model's provider. Required for non-OpenAI models. |
| `apiType` | `Chat` (default) \| `Responses` | No guidance in the source repo. |
| `summarization`, `connection` | objects | No guidance in the source repo. See [CurrentModelsNoKind](reference/schema/settings.md#currentmodelsnokind). |

The schema's [AgentModelType](reference/schema/settings.md#agentmodeltype) union also allows a `kind:` that names a model *channel*. The source repo doesn't
explain these, and its rule is to leave `kind` out. They are listed here only so you recognise them in a cloned
file:

| `kind:` | Extra properties beyond `modelNameHint`, `summarization` |
|---|---|
| [CurrentModels](reference/schema/settings.md#currentmodels) | `apiType`, `provider`, `connection` (same as the no-kind form) |
| [CustomModels](reference/schema/settings.md#custommodels), [PreviewModels](reference/schema/settings.md#previewmodels), [ExperimentalModels](reference/schema/settings.md#experimentalmodels) | none |
| [TurboProductionModels](reference/schema/settings.md#turboproductionmodels), [TurboPreviewModels](reference/schema/settings.md#turbopreviewmodels), [TurboExperimentalModels](reference/schema/settings.md#turboexperimentalmodels) | none |
| [ReasoningProductionModels](reference/schema/settings.md#reasoningproductionmodels), [ReasoningPreviewModels](reference/schema/settings.md#reasoningpreviewmodels), [ReasoningExperimentalModels](reference/schema/settings.md#reasoningexperimentalmodels) | none |
| [DeferredUpdateModels](reference/schema/settings.md#deferredupdatemodels) | none |

The other `aISettings` switches come from the schema only. The source repo doesn't describe them, so the meanings
below are inferred from the names:

| `aISettings` property | Default | Meaning (inferred) |
|---|---|---|
| `useModelKnowledge` | `false` | Let the model answer from its general training knowledge, not only from knowledge sources. |
| `isFileAnalysisEnabled` | `false` | Allow analysis of files the user uploads. |
| `isSemanticSearchEnabled` | `true` | Use semantic search over knowledge. |
| `contentModeration` | `High` | [Content moderation](glossary.md#content-moderation) (how strictly generated output is filtered) level: `Minimum`, `Low`, `Medium`, `High`, `Maximum`. |
| `optInUseLatestModels` | `false` | Opt in to newer model versions. |

### Capabilities (`gptCapabilities`)

All booleans default to `false`. None are described in the source repo; the names are self-explanatory.

| Property | Turns on |
|---|---|
| `webBrowsing` | Web search |
| `codeInterpreter` | Code execution |
| `generateImages` | Image generation |
| `searchTeams`, `searchEmails`, `searchMeetings`, `searchPeople`, `searchPeopleIncludeRelatedContent`, `searchOneDriveAndSharePoint` | Searches over the user's [Microsoft 365 Copilot](glossary.md#microsoft-365-copilot) (Microsoft's assistant in Office apps) data |
| `fineTunedModels`, `customExperience` | (objects/strings; no guidance) |

## Writing instructions

`instructions` is the system prompt, and the vendor `instructions-guide.md` calls it "the most powerful lever for
improving knowledge-based responses". Write it as a YAML literal block (`instructions: |`).

Under [generative orchestration](glossary.md#generative-orchestration) (an LLM plans which topics and tools to run),
instructions shape how the [orchestrator](glossary.md#orchestrator) behaves, but routing to a specific topic or tool
is decided mainly by each component's `modelDescription`
(see [04 Topics & triggers](04-topics-and-triggers.md#modeldescription-vs-trigger-phrases)).

Instructions control five things:

| Concern | Example directive (from the guide) |
|---|---|
| **[Grounding](glossary.md#grounding)** (stick to retrieved content) | "Always base your answers on the provided knowledge sources. Do not use general model knowledge to answer questions about [topic]." |
| **Tone & format** | "Keep answers concise — 3 sentences or fewer unless the user asks for more detail. Use bullet points when listing steps." |
| **[Citation](glossary.md#citation) behavior** | "Always cite the source document when providing an answer. Do not fabricate sources or links that are not in the search results." |
| **Fallback behavior** | "If you cannot find a relevant answer in the knowledge sources, do not guess. Instead, offer to connect the user with a human agent." |
| **Scope enforcement** | "You only answer questions about [subject]. If the user asks about something outside this scope, politely redirect them…" |

**Inject the current date.** The orchestrator doesn't know today's date, so relative phrases ("next week",
"recent") get misread. Interpolate a Power Fx expression directly into the instructions. No variable or topic is
needed:

```yaml
# fragment — the instructions property of agent.mcs.yml
instructions: |
  ## Current Context
  Date: {Text(Today(),DateTimeFormat.LongDate)}

  When users ask about dates, schedules, or time-sensitive information, use this date
  as your reference point.

  ## Your Role
  You are a customer support agent for Contoso Ltd.
  - Be professional and empathetic
  - Escalate billing issues to a human agent
```

`DateTimeFormat.LongDate` produces an unambiguous date such as "Thursday, March 13, 2026". It's evaluated once per
conversation, uses the server's time zone, and costs about 5–10 tokens per orchestrator call. Use `Now()` for
time-of-day precision. You can interpolate [global variables](glossary.md#global-variable) the same way, for example
`User: {Global.UserDisplayName} from {Global.UserCountry}`. (Source: `patterns/date-context.md`. That pattern says
the instructions live in `settings.mcs.yml`; they live in `agent.mcs.yml`.)

**What instructions can't do:**

- Override the model's safety guardrails.
- Force a search of one particular knowledge source. Use a `SearchAndSummarizeContent` node restricted with
  `knowledgeSources:` instead (see [07 Knowledge](07-knowledge.md)).
- Guarantee no hallucination. Grounding reduces it but doesn't eliminate it.
- Change chunking or embedding behavior. That depends on the knowledge content.

## Fields you must not change

The `edit-agent` skill marks these as platform-managed. Warn before touching them:

| Field | File | Why |
|---|---|---|
| `schemaName` | `settings.mcs.yml` | The agent's internal [Power Platform](glossary.md#power-platform) (Microsoft's low-code suite) identifier. Changing it breaks the agent. |
| `publishedOn`, `template` | `settings.mcs.yml` (if present) | Managed by the platform. |
| `language` | `settings.mcs.yml` | Changing it "can corrupt the agent". |

## settings.mcs.yml

`settings.mcs.yml` has **no `kind:`**, so it isn't schema-validated as a whole. Its fields have the same names as
the schema's [BotEntity](reference/schema/components.md#botentity) definition (the agent's own row in [Dataverse](glossary.md#dataverse), the Power Platform database), so that page is the best reference for its
shape (inferred from matching property names). This is the eval-fixture file from the source repo:

```yaml
# fragment — settings.mcs.yml (no kind:, not schema-validated)
displayName: Eval Test Agent
schemaName: eval_testAgent
accessControlPolicy: ChatbotReaders
authenticationMode: None
configuration:
  settings:
    GenerativeActionsEnabled: false

  recognizer:
    kind: GenerativeAIRecognizer

language: 1033
```

| Field | Type/values | Meaning |
|---|---|---|
| `displayName` | string | Display name. |
| `schemaName` | string | The agent's unique schema name. Every component's schema name starts with it (`cr26e_myAgent.topic.Escalate`). **Never change it.** |
| `accessControlPolicy` | [BotAccessControlPolicy](reference/schema/components.md#botaccesscontrolpolicy): `Any` (schema default), `ChatbotReaders`, `GroupMembership`, `AnyMultiTenant` | Who may chat with the agent. The source only shows `ChatbotReaders`. The other values aren't explained (by name: anyone, members of a group, anyone in any [tenant](glossary.md#tenant)). |
| `authenticationMode` | [BotAuthenticationMode](reference/schema/components.md#botauthenticationmode): `Unspecified` (schema default), `None`, `Integrated`, `CustomAzureActiveDirectory`, `GenericOAuth2` | The [authentication mode](glossary.md#authentication-mode) (how end users are identified). `None` = anonymous; `Integrated` = [Entra ID](glossary.md#entra-id) (Microsoft's identity service) single sign-on as in [Teams](glossary.md#teams) (Microsoft's chat app); `CustomAzureActiveDirectory` / `GenericOAuth2` = "manual" auth against an identity provider you configure. It also decides which test tools can reach the published agent (see [12 Evaluations & testing](12-evaluations-and-testing.md)). |
| `authenticationTrigger` | [BotAuthenticationTrigger](reference/schema/components.md#botauthenticationtrigger): `AsNeeded` (default), `Always` | When sign-in is requested. |
| `configuration.settings.GenerativeActionsEnabled` | `true` \| `false` | **Orchestration mode.** `true` = generative orchestration; `false` = [classic orchestration](glossary.md#classic-orchestration) (trigger-phrase matching). See [01 Concepts](01-concepts.md#how-a-message-is-routed). |
| `configuration.recognizer` | [RecognizerBase](reference/schema/dialogs.md#recognizerbase) | Intent recognizer. See [below](#recognizers). |
| `configuration.aISettings` | [AISettingsNoKind](reference/schema/settings.md#aisettingsnokind) | The `edit-agent` skill lists `configuration.aISettings.*` ("various booleans") as editable here. It's the same shape as `aISettings` in `agent.mcs.yml`. |
| `language` | [Language](reference/schema/settings.md#language) (a number) | The agent's primary language as a Windows [LCID](glossary.md#lcid) (locale ID). **Don't change it.** |

`configuration.settings` is a free-form map in the schema ([BotConfiguration](reference/schema/components.md#botconfiguration) types it as `object`), so the
schema can't check the keys inside it. `GenerativeActionsEnabled` is the only key the source repo documents. The
`edit-agent` skill also mentions a web-browsing switch under `configuration.settings.*.content.capabilities.webBrowsing`,
without giving its full path. In the schema, web browsing is `gptCapabilities.webBrowsing` in `agent.mcs.yml`.

The `language` values the schema allows are: 1025 Arabic, 2052 Chinese (Simplified), 1028 Chinese (Traditional),
1029 Czech, 1030 Danish, 1043 Dutch, 1033 English (US), 2057 English (UK), 3081 English (Australia), 1035 Finnish,
1036 French, 3084 French (Canada), 1031 German, 1032 Greek, 1037 Hebrew, 1081 Hindi, 1057 Indonesian, 1040 Italian,
1041 Japanese, 1042 Korean, 1044 Norwegian, 1045 Polish, 1046 Portuguese (Brazil), 2070 Portuguese (Portugal),
1049 Russian, 1034 Spanish, 21514 Spanish (US), 20490 Spanish (Puerto Rico), 1053 Swedish, 1054 Thai, 1055 Turkish,
1058 Ukrainian.

Other `configuration` properties in [BotConfiguration](reference/schema/components.md#botconfiguration) (`isTelephonyEnabled`, `isAgentConnectable`,
`botSpeechSettings`, `voiceProcessingMode`, `sensitiveDataMaskingSettings`, `channels`, `publishOnCreate`, …) have no
guidance in the source repo.

### Recognizers

`configuration.recognizer` takes one of the four [RecognizerBase](reference/schema/dialogs.md#recognizerbase) kinds. None has any properties besides `kind`.

| `kind:` | Notes |
|---|---|
| [GenerativeAIRecognizer](reference/schema/dialogs.md#generativeairecognizer) | The only one in the source repo's fixtures. They pair it with `GenerativeActionsEnabled: false`, so the recognizer and the orchestration switch are set separately. |
| [ClassicAIRecognizer](reference/schema/dialogs.md#classicairecognizer) | No guidance in the source repo. |
| [AdvanceNLURecognizer](reference/schema/dialogs.md#advancenlurecognizer) | No guidance in the source repo. |
| [CLIAgentRecognizer](reference/schema/dialogs.md#cliagentrecognizer) | No guidance in the source repo. |

Don't confuse this with a topic's own `recognizer:` property, which takes a different union
([Recognizer](reference/schema/dialogs.md#recognizer): `RegexRecognizer`). See [04 Topics & triggers](04-topics-and-triggers.md#adaptivedialog-root-properties).

## Agent settings components (`BotSettingsBase`)

The schema has a component kind, [BotSettingsComponent](reference/schema/components.md#botsettingscomponent), whose `settings:` holds one block from the
[BotSettingsBase](reference/schema/settings.md#botsettingsbase) union. These blocks cover content moderation, autonomous runs, voice and feedback. **The source
repo has no examples of these blocks, doesn't say which file they live in, and doesn't describe their behavior.**
A cloned agent may contain them. If so, leave them as they are or check them against the generated reference. The
one-line purposes below are inferred from the names.

| `kind:` | Properties | Purpose (inferred) |
|---|---|---|
| [ContentModerationSettings](reference/schema/settings.md#contentmoderationsettings) | `moderatedCustomResponse` (a message) | The reply sent when content moderation blocks a response. The moderation *level* itself is `aISettings.contentModeration` (agent-wide) or `moderationLevel` on generative-answer nodes (`Minimum` … `Maximum`). For handling blocked responses, see the [RAI error-handling pattern](13-patterns.md#rai-error-handling). |
| [AutonomousSettings](reference/schema/settings.md#autonomoussettings) | `enterpriseKnowledgeConnectionReference` | Settings for [autonomous agents](glossary.md#autonomous-agent) (agents started by events or schedules, not by a user). See [10 Flows & autonomous triggers](10-flows-and-autonomous-triggers.md). |
| [ConnectionParameterSettings](reference/schema/settings.md#connectionparametersettings) | `connectorSettings`, `flowSettings` (free-form maps) | Per-connector and per-flow connection parameters. |
| [IvrBotSettings](reference/schema/settings.md#ivrbotsettings) | `voiceFont`, `turnDetection`, `interruptionEagerness`, `holdSettings`, `latencySettings`, `domainLanguageModelWeightValue` | Voice settings for [IVR](glossary.md#ivr) (phone) agents. |
| [DigitalVoiceSettings](reference/schema/settings.md#digitalvoicesettings) | `voiceFont`, `turnDetection`, `silenceDurationInMilliseconds`, `noiseReduction`, `echoCancellation`, `listeningMode`, … | Voice settings for digital (non-phone) voice channels. |
| [RealTimeVoiceBotSettings](reference/schema/settings.md#realtimevoicebotsettings) | `voiceType`, `latencySettings` | Real-time voice settings. |
| [DefaultFeedbackCollection](reference/schema/settings.md#defaultfeedbackcollection) | `disclaimer` (text) | Turns on thumbs-up/down feedback on responses, with an optional disclaimer. |
| [NoFeedbackCollection](reference/schema/settings.md#nofeedbackcollection) | none | Turns feedback collection off. |

## Sources

- `vendor/skills-for-copilot-studio/skills/edit-agent/SKILL.md`: editable fields, model rule, fields never to modify
- `vendor/skills-for-copilot-studio/skills/edit-agent/instructions-guide.md`: writing instructions
- `vendor/skills-for-copilot-studio/patterns/date-context.md`: date injection
- `vendor/skills-for-copilot-studio/templates/agents/agent.mcs.yml`: agent template
- `vendor/skills-for-copilot-studio/evals/fixtures/basic-agent/settings.mcs.yml`, `evals/fixtures/agent-with-mcp-action/settings.mcs.yml`: settings examples
- `vendor/skills-for-copilot-studio/skills/int-reference/SKILL.md`: file types
- `vendor/skills-for-copilot-studio/agents/copilot-studio-author.md`: "don't scaffold agent files" rule
- `vendor/skills-for-copilot-studio/skills/chat-with-agent/SKILL.md`: auth mode decides Direct Line vs. SDK testing
- Generated reference: [components](reference/schema/components.md), [settings](reference/schema/settings.md), [dialogs](reference/schema/dialogs.md)
