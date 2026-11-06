# 13 — Patterns

The skills repo ships a library of 15 design patterns: tested solutions to recurring Copilot Studio problems, each
with YAML. This page summarises every pattern (status, problem, solution, key YAML, caveats) and links to the
original file. The YAML here is adapted to these docs' conventions (`cr26e_myAgent` schema prefix, real IDs) and
checked against the schema. Where the original YAML doesn't match the schema, the section says so.

## How patterns are chosen

Each pattern file has a frontmatter `status`:

| Status | Meaning (how the skills present it) |
|---|---|
| `proven` | "Used in production — consider it for…" |
| `recommended` | "A recommended approach that works well" |
| `experimental` | "Not yet fully validated — test thoroughly before adopting" |

In the skills plugin, the **Advisor** agent suggests patterns ("you could consider…", never "you must…") and waits
for the user's decision. The **Author** agent implements a pattern only after it has been chosen, reading only the
relevant pattern file for its YAML. The `int-patterns` skill is the index that both use: for each pattern it lists
the symptoms that should prompt reading it (e.g. "users report seeing raw JSON" → *Prevent tool call leaks*). The
same advice applies to you: patterns are recommendations, not requirements.

| Pattern | Status | Solves |
|---|---|---|
| [Chain-of-thought logging](#chain-of-thought-logging) | recommended | Silent waits during multi-step tool/agent chains |
| [Channel-aware behavior](#channel-aware-behavior) | experimental | Features that work in one channel but break in another |
| [Conversation history variable](#conversation-history-variable) | proven | Passing the conversation transcript to escalation or tools |
| [Date context](#date-context) | proven | The orchestrator misreading "next week", "recent" |
| [Deterministic MCP calls](#deterministic-mcp-calls) | experimental | The orchestrator skipping a required MCP tool |
| [Dynamic topic redirect](#dynamic-topic-redirect) | proven | Deeply nested condition chains for routing |
| [JIT glossary](#jit-glossary) | proven | Internal acronyms ruining knowledge search |
| [JIT user context](#jit-user-context) | proven | Generic answers that ignore the user's country/department |
| [Knowledge hold message](#knowledge-hold-message) | proven | Silence while knowledge search runs |
| [Line breaks in messages](#line-breaks-in-messages) | proven | Multi-line messages collapsing into a wall of text |
| [Orchestrator variables](#orchestrator-variables) | proven | Classifying/extracting data without extra AI calls |
| [Prevent child agent responses](#prevent-child-agent-responses) | proven | Child agents messaging users directly |
| [Prevent tool call leaks](#prevent-tool-call-leaks) | recommended | Raw tool-call JSON shown to users |
| [RAI error handling](#rai-error-handling) | recommended | Generic errors when content filters block a request |
| [Teams production hardening](#teams-production-hardening) | recommended | Stale sessions, reinstalls, resets and diagnostics in Teams |

Common combinations from the index:

- **JIT glossary + JIT user context**: merge into one `OnActivity` initialisation topic (the vendor template
  `templates/topics/conversation-init.topic.mcs.yml` does exactly this).
- **Date context + JIT user context**: both are references in the agent instructions.
- **Orchestrator variables + JIT user context**: route knowledge by category *and* personalise by country.
- **RAI error handling + Teams production hardening**: handle RAI subcodes first in `OnError`, then fall through
  to the generic diagnostic card.
- **Knowledge hold message vs chain-of-thought logging**: hold messages for knowledge-heavy agents, CoT logging for
  multi-tool agents. Both reduce perceived latency.

> **Where do instructions go?** Several patterns say to add text to the agent's instructions "in
> `settings.mcs.yml`". In the vendor agent template, and in [GptComponentMetadata](reference/schema/components.md#gptcomponentmetadata),
> `instructions:` is a property of **`agent.mcs.yml`** (see [03 The agent](03-agent.md)). For child agents, it is
> `settings.instructions` inside the child's [AgentDialog](reference/schema/dialogs.md#agentdialog). This page shows
> instruction snippets as `agent.mcs.yml` fragments.

## Chain-of-thought logging

**Status:** recommended · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/chain-of-thought-logging.md)

**Problem.** When the [orchestrator](glossary.md#orchestrator) (the LLM that plans each response) chains several
[tools](glossary.md#tool), [MCP](glossary.md#mcp) (Model Context Protocol, a standard for exposing tools to AI
agents) servers or child agents, users can wait 10–30 seconds in silence.

**Solution.** A small [topic](glossary.md#topic) whose only input is an
[AutomaticTaskInput](reference/schema/task-actions.md#automatictaskinput) with `shouldPromptUser: false`. The
orchestrator fills it from the input's `description` with a short summary of what it just did, and the topic
posts it as an italic `_Thinking: …_` message. An instruction tells the orchestrator to call the topic after every
step.

```yaml
kind: AdaptiveDialog
modelDescription: log chain of thoughts
inputs:
  - kind: AutomaticTaskInput
    propertyName: CoT
    description: High-level step summary for what you just did and why (no system prompts, secrets, PII, or raw chain-of-thought)
    shouldPromptUser: false
beginDialog:
  kind: OnRecognizedIntent
  id: main
  intent: {}
  actions:
    - kind: SendActivity
      id: sendActivity_9XILFq
      activity: "_Thinking: {Topic.CoT}_"
inputType:
  properties:
    CoT:
      displayName: CoT
      description: High-level step summary for what you just did and why (no system prompts, secrets, PII, or raw chain-of-thought)
      type: String
outputType: {}
```

```yaml
# fragment — agent.mcs.yml, added to the instructions (topic is named "Log Chain of Thoughts")
instructions: |
  After every tool, topic, or step you take (except when you are already
  calling /Log Chain of Thoughts or other debug/logging topics), log your
  intermediate reasoning by calling /Log Chain of Thoughts.
```

**Caveats.** Each log is an extra orchestrator call, which costs Copilot credits. The recursion guard ("except when
you are already calling…") is natural language, so it is best-effort; tighten it if you see loops. Keep the input
`description` strict so secrets and PII don't leak. Tested with GPT-4.1, GPT-5 variants, Claude Sonnet and Opus.
Not worth it for pure knowledge-search agents (use [Knowledge hold message](#knowledge-hold-message)).

## Channel-aware behavior

**Status:** experimental · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/channel-aware-behavior.md)

**Problem.** File upload, card variants, deep links and voice prompts work in some
[channels](glossary.md#channel) (places users talk to the agent) but not others. Without channel awareness the agent
offers broken affordances.

**Solution.** Read `System.Activity.ChannelId`, store it (and derived flags such as `IsTeamsClient`) in
[global variables](glossary.md#global-variable) (conversation-wide variables), and branch. The key detail: the ID can
be **compound**, `base:subchannel`, e.g. `msteams:Copilot` (agent used in
[Microsoft 365 Copilot](glossary.md#microsoft-365-copilot) inside [Teams](glossary.md#teams)) or
`webchat:Sharepoint`. Always `Lower()` it and use `StartsWith` for the base channel.

```yaml
kind: AdaptiveDialog
beginDialog:
  kind: OnRecognizedIntent
  id: main
  intent:
    triggerQueries:
      - dump activity
  actions:
    - kind: SetVariable
      id: setVariable_cH4nP1
      variable: Global.ClientPlatform
      value: =System.Activity.ChannelId
    - kind: SetVariable
      id: setVariable_cH4nP2
      variable: Global.IsTeamsClient
      value: =StartsWith(Lower(System.Activity.ChannelId), "msteams")
    - kind: SendActivity
      id: sendMessage_cH4nP3
      activity: "Channel: {Global.ClientPlatform} | IsTeams: {Global.IsTeamsClient}"
    - kind: ConditionGroup
      id: conditionGroup_cH4nP4
      conditions:
        - id: conditionItem_cH4nP5
          condition: =Global.IsTeamsClient
          actions:
            - kind: SendActivity
              id: sendMessage_cH4nP6
              activity: This action is not available in Teams.
            - kind: EndDialog
              id: endDialog_cH4nP7
      elseActions:
        - kind: SendActivity
          id: sendMessage_cH4nP8
          activity: Continuing in a non-Teams channel.
```

| ChannelId (lowercased) | Surface |
|---|---|
| `msteams` | Teams (standalone agent) |
| `msteams:copilot` | Microsoft 365 Copilot in Teams |
| `m365extensions` | Microsoft 365 Copilot extension surface |
| `webchat` / `webchat:sharepoint` | Embedded web chat / inside [SharePoint](glossary.md#sharepoint) |
| `directline` / `directlinespeech` | [Direct Line](glossary.md#direct-line) clients / voice clients |
| `outlook`, `omnichannel` | Outlook; Dynamics 365 Omnichannel handoff |
| `slack`, `telegram`, `facebook`, `skype`, `sms`, `email`, `telephony` | Other channels |

To let the orchestrator reason about the channel, give `ClientPlatform` `aIVisibility: UseInAIContext` and mention
`{Global.ClientPlatform}` in the instructions.

**Caveats.** The Test pane reports its own channel value, so test in the real channel. Casing isn't guaranteed. If
many topics need the value, detect it once in an `OnActivity` initialisation topic (as in
[JIT user context](#jit-user-context)). **Channel is not a security boundary**: Direct Line callers can spoof it, so
gate sensitive actions on the authenticated user instead.

## Conversation history variable

**Status:** proven · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/conversation-history-variable.md)

**Problem.** Escalation, logging and downstream tools need the conversation so far, and rebuilding it turn by turn
is fragile.

**Solution.** Same trick as chain-of-thought logging: an `AutomaticTaskInput` whose `description` asks the
orchestrator to reconstruct the whole conversation in a given format. The topic then shows, stores or forwards it.

```yaml
kind: AdaptiveDialog
modelDescription: save the conversation history
inputs:
  - kind: AutomaticTaskInput
    propertyName: ConversationHistory
    description: "Entire conversation history in the format \"User: question <br /><br /> Agent : response <br /><br /> User: Question <br /><br /> Agent: response\""
    shouldPromptUser: false
beginDialog:
  kind: OnRecognizedIntent
  id: main
  intent: {}
  actions:
    - kind: SendActivity
      id: sendActivity_gDSbH3
      activity: "Here is your saved conversation history: {Topic.ConversationHistory}"
inputType:
  properties:
    ConversationHistory:
      displayName: ConversationHistory
      description: "Entire conversation history in the format \"User: question <br /><br /> Agent : response <br /><br /> User: Question <br /><br /> Agent: response\""
      type: String
outputType: {}
```

Variations: replace the `SendActivity` with a `SetVariable` to store it; change the `description` to ask for a
summary or the last N turns; call the topic from a `RecognizeIntent` node at escalation time.

**Caveats.** It's a **best-effort reconstruction** from the orchestrator's context window, not a verbatim or
compliance-grade record, and long conversations may be truncated. It can contain PII, so a summary is often safer.

## Date context

**Status:** proven · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/date-context.md)

**Problem.** The orchestrator doesn't know today's date, so "next week", "upcoming" and "recent" are misread, and
time-sensitive knowledge is filtered wrongly.

**Solution.** Put the date into the [instructions](glossary.md#instructions) (the agent's system prompt) with a
[Power Fx](glossary.md#power-fx) template. `DateTimeFormat.LongDate` gives an unambiguous value such as
"Thursday, March 13, 2026".

```yaml
# fragment — agent.mcs.yml instructions excerpt
instructions: |
  ## Current Context
  Date: {Text(Today(),DateTimeFormat.LongDate)}

  When users ask about dates, schedules, or time-sensitive information, use this date
  as your reference point. Interpret "next", "upcoming", "recent", and other relative
  terms based on this current date.
```

**Caveats.** Costs about 5–10 tokens per orchestrator call. It is evaluated once per conversation, not per
message. `Today()` uses the server's time zone; add explicit time-zone handling if precision across zones matters.

## Deterministic MCP calls

**Status:** experimental · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/deterministic-mcp-calls.md)

**Problem.** MCP tools are *generative actions*: the orchestrator decides whether to call them. You can't force one
with `/ToolName` in instructions, and (per the pattern) topics can't call MCP tools directly. A tool that must
always run for an intent may be skipped.

**Solution.** Two workarounds, lightest first:

1. **Name the tool in the instructions** and map intents to it explicitly. Reliability is high but not guaranteed.
2. **Wrap the tool in a [child agent](glossary.md#child-agent)** (a sub-agent the main agent delegates to) whose
   only job is to call it. The parent routes to the child by its description, which is near-deterministic.

```yaml
# fragment — agent.mcs.yml instructions excerpt (Option 1)
instructions: |
  ## Tool Usage Rules
  When the user asks about <specific intent>, you MUST call the
  <MCP Tool Name> tool to retrieve the answer.
  Do not attempt to answer from your own knowledge.
```

```yaml
# Option 2 — agents/PartLookup/agent.mcs.yml
kind: AgentDialog
beginDialog:
  kind: OnToolSelected
  id: main
  description: Specialized agent for part-number lookups. Always calls the Part Catalog tool. Use for requests like "look up part 4471" or "what is the part number for the M8 bolt".
settings:
  instructions: |
    You are a specialized agent that answers part-number questions.

    CRITICAL: You MUST call the Part Catalog tool for EVERY request you receive.
    Never answer from your own knowledge. Always use the tool and return its results.
inputType: {}
outputType: {}
```

**Caveats.** Option 2 adds a hop and maintenance; its description must not clash with other topics or agents.
If the parent should format the final answer, combine with
[Prevent child agent responses](#prevent-child-agent-responses).

Corrections to the pattern file: its Option 2 YAML puts `intent.triggerQueries` under `OnToolSelected`, which the
schema doesn't allow ([OnToolSelected](reference/schema/triggers.md#ontoolselected) takes a required
`description` instead), so the example above drops it. It also names the file `agents/MyMcpChild.agent.mcs.yml`,
while child agents normally live at `agents/<Name>/agent.mcs.yml` ([02](02-project-layout.md)). Finally, the
schema *does* define a topic node [InvokeMcpToolAction](reference/schema/actions.md#invokemcptoolaction)
(`serverSchemaName`, `toolName`), which seems to contradict "topics cannot call MCP tools". The sources never use
it, so treat it as unverified.

## Dynamic topic redirect

**Status:** proven · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/dynamic-topic-redirect.md)

**Problem.** Routing to one of N topics by a variable's value becomes a deep, hard-to-maintain `ConditionGroup`.

**Solution.** [BeginDialog](reference/schema/actions.md#begindialog)'s `dialog` accepts a Power Fx expression that
evaluates to a topic schema name, so one `Switch()` replaces the whole chain:

```yaml
# fragment — two nodes of a topic (see the note below about validators)
- kind: SetVariable
  id: setVariable_7bgfoP
  variable: init:Topic.LessonNumber
  value: =RandBetween(0,4)
- kind: BeginDialog
  id: beginDialog_A4lDAn
  dialog: |-
    =Switch(
        Topic.LessonNumber,
        1, "cr26e_myAgent.topic.Lesson1",
        2, "cr26e_myAgent.topic.Lesson2",
        3, "cr26e_myAgent.topic.Lesson3",
        "cr26e_myAgent.topic.Fallback"
    )
```

**Caveats.** Each branch value must be a **fully qualified** schema name (prefix from `settings.mcs.yml`). Always
give `Switch()` a default as the last argument. For routing on several variables or compound conditions, a
`ConditionGroup` is still clearer. The skills' validator accepts this YAML. A strict JSON-Schema validator rejects
it, because the schema's `DialogExpression` is a `oneOf` of two string shapes that both match a `=` string; that's
why the block above is marked as a fragment.

## JIT glossary

**Status:** proven · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/jit-glossary.md)

**Problem.** The agent doesn't know the organisation's acronyms, so knowledge search for "WFH policy for the GTM
team" goes wrong.

**Solution.** Load a glossary file into `Global.Glossary` the first time the user sends a message ("just in time"),
and reference it once in the instructions so the orchestrator silently expands acronyms before searching:

1. Upload the glossary as a [knowledge source](glossary.md#knowledge-source) (content the agent can search)
   stored in [Dataverse](glossary.md#dataverse) (the [Power Platform](glossary.md#power-platform) database), as a **`.txt` file with CSV
   content**, with `triggerCondition: =false` so it is never searched automatically.
2. An [OnActivity](reference/schema/triggers.md#onactivity) topic with `type: Message` and
   `condition: =IsBlank(Global.Glossary)` retrieves it with
   [SearchAndSummarizeContent](reference/schema/actions.md#searchandsummarizecontent).
3. The instructions include `{Global.Glossary}`.

```yaml
kind: AdaptiveDialog
beginDialog:
  kind: OnActivity
  id: main
  type: Message
  condition: =IsBlank(Global.Glossary)
  actions:
    - kind: SearchAndSummarizeContent
      id: searchContent_gL0ss1
      autoSend: false
      variable: Topic.GlossaryResult
      userInput: ="*"
      responseCaptureType: FullResponse
      applyModelKnowledgeSetting: false
      knowledgeSources:
        kind: SearchSpecificKnowledgeSources
        knowledgeSources:
          - cr26e_myAgent.knowledge.glossary
    - kind: SetVariable
      id: setVariable_gL0ss2
      variable: Global.Glossary
      value: =Topic.GlossaryResult
```

```yaml
# fragment — agent.mcs.yml instructions excerpt
instructions: |
  ## Glossary
  {Global.Glossary}
  The above is the customer glossary (format: ACRONYM,Definition, one per line).
  Silently expand any acronym found in it before interpreting the user's message or searching knowledge sources.
  Do not mention the glossary to the user unless they explicitly ask for a list of acronyms.
```

**Why `OnActivity` and not `OnConversationStart`?** `OnConversationStart` doesn't fire in Microsoft 365 Copilot and
other embedded surfaces. A `Message` activity with an `IsBlank` guard runs exactly once, on the first real message.

**Caveats.** Reference `{Global.Glossary}` **once** only, since every reference injects the whole glossary into
the prompt. SharePoint knowledge sources can't return a whole file (semantic search returns chunks), hence
Dataverse. The knowledge-source name in the topic must match the knowledge file name. The pattern's YAML has
`modelDescription: null` and `userInput: ='*'` (a single-quoted Power Fx string); this version omits the null (see
[14 Gotchas](14-gotchas.md#modeldescription-null-in-a-vendor-template)) and uses double quotes.

## JIT user context

**Status:** proven · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/jit-user-context.md)

**Problem.** The agent can't tailor answers (WFH policy, holidays, local rules) to the user's country or department
without asking.

**Solution.** On the first message, call the Office 365 Users [connector](glossary.md#connector) (a pre-built
wrapper around an external API) operation `UserGet_V2` ("get my profile") with the user's own credentials, copy
fields into global variables with `aIVisibility: UseInAIContext`, and reference them in the instructions.

```yaml
kind: AdaptiveDialog
beginDialog:
  kind: OnActivity
  id: main
  type: Message
  condition: =IsBlank(Global.UserCountry)
  actions:
    - kind: InvokeConnectorAction
      id: invokeConnector_pR0f1L
      connectionReference: shared_office365users
      connectionProperties:
        mode: Invoker
      operationId: UserGet_V2
      output:
        kind: SingleVariableOutputBinding
        variable: init:Topic.M365Profile
    - kind: SetVariable
      id: setVariable_pR0f2C
      variable: Global.UserCountry
      value: =If(IsBlank(Topic.M365Profile.country), "Unknown", Topic.M365Profile.country)
```

```yaml
# variables/UserCountry.mcs.yml (on-disk, flattened shape)
name: UserCountry
aIVisibility: UseInAIContext
scope: Conversation
description: The user's country from their M365 profile, loaded JIT at conversation start.
schemaName: cr26e_myAgent.globalvariable.UserCountry
kind: GlobalVariableComponent
defaultValue: DEFAULT
```

```yaml
# fragment — agent.mcs.yml instructions excerpt
instructions: |
  ## User Context
  The current user's country is {Global.UserCountry}.
  When the user asks a location-dependent question, use this country to search the correct
  knowledge sources. Do not ask the user for their country — it is already known.
  If the country is "Unknown", answer with the general policy and note it may vary by country.
```

**Caveats.** The connection must exist first (set up in the Copilot Studio UI), and the agent needs user
authentication, otherwise the call fails silently. Directory fields such as `country` are often empty, so always
fall back. Without `UseInAIContext` the orchestrator ignores the variable
([AI visibility](glossary.md#ai-visibility)). Combine with the glossary in a single `OnActivity` topic rather than
two. `connectionReference: shared_office365users` is the pattern's placeholder; in a real agent use the logical
name from `connectionreferences.mcs.yml` ([08](08-tools-and-connectors.md)).

## Knowledge hold message

**Status:** proven · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/knowledge-hold-message.md)

**Problem.** Where the model or channel doesn't stream responses, users stare at a silent screen during knowledge
search, then resend or abandon.

**Solution.** An [OnKnowledgeRequested](reference/schema/triggers.md#onknowledgerequested) topic (fires whenever
the orchestrator decides to search knowledge) that immediately sends a short hold message, picked at random from a
Power Fx table. The orchestrator then carries on with the search as usual; you add no search node.

```yaml
kind: AdaptiveDialog
beginDialog:
  kind: OnKnowledgeRequested
  id: main
  actions:
    - kind: SetVariable
      id: setVariable_hM7xQ2
      variable: init:Topic.HoldMessages
      value: "=Table({Value: \"Let me dig into that for you...\"}, {Value: \"One moment while I look that up!\"}, {Value: \"Searching my knowledge base now...\"}, {Value: \"Give me just a sec to find that...\"})"
    - kind: SetVariable
      id: setVariable_rP4kN8
      variable: init:Topic.SelectedMessage
      value: =Index(Topic.HoldMessages, RoundDown(Rand() * CountRows(Topic.HoldMessages), 0) + 1).Value
    - kind: SendActivity
      id: sendActivity_wK9mT3
      activity: "{Topic.SelectedMessage}"
```

(The original has 40 messages. For a fixed message, keep only the `SendActivity`.)

**Caveats.** Every knowledge answer becomes two messages, which some stakeholders find chatty. Keep hold messages
short. `OnKnowledgeRequested` is **YAML-only**: the UI can't create or show it ([14 Gotchas](14-gotchas.md)). For
per-channel behaviour, wrap the `SendActivity` in a condition on `System.Activity.ChannelId`.

## Line breaks in messages

**Status:** proven · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/line-breaks-in-messages.md)

**Problem.** Most channels render plain YAML newlines as spaces, so multi-paragraph messages become walls of text.

**Solution.** Put `<br />` (line break) or `<br /><br />` (paragraph) inside `SendActivity.activity` and
`Question.prompt`, written in a `|-` block for readability.

```yaml
- kind: Question
  id: question_bR3k01
  alwaysPrompt: true
  variable: init:Topic.UserChoice
  prompt: |-
    I need a few details to get started.
    <br /><br />
    Please select one of the options below to continue.
  entity: BooleanPrebuiltEntity
```

**Caveats.** Don't rely on plain newlines. Apply the same style to messages and questions. Test long messages in the
target channel. More in [11 Messages & Adaptive Cards](11-messages-and-adaptive-cards.md#line-breaks-and-formatting).

## Orchestrator variables

**Status:** proven · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/orchestrator-variables.md)
(titled *Orchestrator-Generated Variables*)

**Problem.** Classifying a request (HR / IT / Finance) or extracting a value normally means asking the user or
running a separate AI prompt, which costs credits and latency.

**Solution.** Declare a topic input as `AutomaticTaskInput` with `shouldPromptUser: false` and a `description`
that is the extraction prompt, listing the allowed values. Under
[generative orchestration](glossary.md#generative-orchestration) the orchestrator fills it in the **same** LLM call
it uses to pick the topic, so there's no extra cost or latency. Declare it in **both** `inputs` and
`inputType.properties`.

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
      id: setVariable_cL4ss1
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

A companion `OnKnowledgeRequested` topic then searches only the matching sources, writing to
`System.SearchResults`:

```yaml
kind: AdaptiveDialog
beginDialog:
  kind: OnKnowledgeRequested
  id: main
  actions:
    - kind: ConditionGroup
      id: conditionGroup_r0uT31
      conditions:
        - id: conditionItem_r0uTHR
          condition: =Global.searchCategory = "HR"
          actions:
            - kind: SearchAndSummarizeContent
              id: searchContent_r0uTH2
              variable: System.SearchResults
              userInput: =System.SearchQuery
              knowledgeSources:
                kind: SearchSpecificKnowledgeSources
                knowledgeSources:
                  - cr26e_myAgent.knowledge.hr-policies
        - id: conditionItem_r0uTIT
          condition: =Global.searchCategory = "IT"
          actions:
            - kind: SearchAndSummarizeContent
              id: searchContent_r0uTI2
              variable: System.SearchResults
              userInput: =System.SearchQuery
              knowledgeSources:
                kind: SearchSpecificKnowledgeSources
                knowledgeSources:
                  - cr26e_myAgent.knowledge.it-documentation
      elseActions:
        - kind: SearchAndSummarizeContent
          id: searchContent_r0uTEL
          variable: System.SearchResults
          userInput: =System.SearchQuery
```

**Caveats.** Write the description as a precise instruction with an enum-style list; 3–6 categories work well,
and more than 10 degrades accuracy. Always handle a fallback (`Other`, `elseActions`). The value is set once, when
the topic is invoked. A knowledge source's `triggerCondition` can't use runtime variables like this, so do the
routing inside `OnKnowledgeRequested` ([07 Knowledge](07-knowledge.md)). Not for complex reasoning or long documents;
use an AI prompt for those.

## Prevent child agent responses

**Status:** proven · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/prevent-child-agent-responses.md)

**Problem.** [Child agents](glossary.md#child-agent) and [connected agents](glossary.md#connected-agent) can
message the user directly (through their built-in `SendMessageTool`), bypassing the parent's formatting, filtering
and tone.

**Solution.** Forbid it in the **child's** instructions and have it return its answer in an output variable that the
parent reads. The child agent's *completion setting* does **not** control this; it only decides what the parent
does after the child finishes.

```yaml
# agents/OrderStatus/agent.mcs.yml
kind: AgentDialog
beginDialog:
  kind: OnToolSelected
  id: main
  description: Answers order-status questions such as "where is order 1042" or "has my order shipped".
settings:
  instructions: |
    CRITICAL - DO NOT MESSAGE USERS
    - DO NOT respond directly to the user
    - DO NOT call SendMessageTool or send any messages
    - ONLY populate the output variables with your response
    - Let the parent orchestrator deliver the response to the user
inputType: {}
outputType:
  properties:
    AgentOutput:
      displayName: AgentOutput
      description: The answer for the parent agent to deliver to the user.
      type: String
```

The parent can then read `AgentOutput` after a [BeginDialog](reference/schema/actions.md#begindialog), using the
schema's output-binding shape:

```yaml
- kind: BeginDialog
  id: beginDialog_cH1ld1
  dialog: cr26e_myAgent.agent.OrderStatus
  output:
    binding:
      AgentOutput: Topic.ChildResult
- kind: SendActivity
  id: sendMessage_cH1ld2
  activity: "{Topic.ChildResult}"
```

**Caveats.** It's instruction-based, not enforced by the platform, so test that the child complies. The child must
declare output variables, or it has no way to return data. The pattern file writes the parent call with
`output: { kind: VariableOutputBindings, bindings: [...] }`; **those kinds don't exist in the schema**, so the
version above uses `output.binding` ([ActionOutputBindingNoKind](reference/schema/actions.md#actionoutputbindingnokind))
instead. The `<agent>.agent.<Name>` dialog name is as given in the pattern; see
[09 MCP & multi-agent](09-mcp-and-multi-agent.md) for how child agents are invoked.

## Prevent tool call leaks

**Status:** recommended · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/prevent-tool-call-leaks.md)

**Problem.** The orchestrator sometimes returns raw JSON with internal fields like `explanation_of_tool_call` and
`new_instruction` instead of a clean answer, especially when it calls connector tools on its own initiative.

**Solution.** Two approaches:

1. **Instructions** (simpler, less reliable): "Return only the final user-facing answer. Do not include internal
   reasoning, tool call explanations, or diagnostic JSON."
2. **Intercept the response** (more reliable): an
   [OnGeneratedResponse](reference/schema/triggers.md#ongeneratedresponse) topic (fires just before an AI-generated
   answer is sent) that matches the leak and suppresses sending with `System.ContinueResponse = false`.

```yaml
kind: AdaptiveDialog
beginDialog:
  kind: OnGeneratedResponse
  id: main
  actions:
    - kind: ConditionGroup
      id: conditionGroup_r9Qpg6
      conditions:
        - id: conditionItem_0mxghZ
          condition: =IsMatch(Lower(System.Response.FormattedText), "explanation_of_tool_call")
          actions:
            - kind: SetVariable
              id: setVariable_vaDcla
              variable: System.ContinueResponse
              value: false
inputType: {}
outputType: {}
```

**Caveats.** Approach 1 has worked inconsistently across configurations. Approach 2 only catches the strings you
match, so add conditions for other metadata formats. This is known orchestrator behaviour with no platform fix yet.
The pattern says `OnGeneratedResponse` is YAML-only (the UI can't create it).

## RAI error handling

**Status:** recommended · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/rai-error-handling.md)

**Problem.** When a [RAI](glossary.md#rai) (Responsible AI) content filter blocks a request, users see a generic
error with no explanation or next step.

**Solution.** Override the `OnError` [system topic](glossary.md#system-topic) (a built-in topic every agent has).
All filter blocks arrive with error code `ContentFiltered`, and `System.Error.SubCode` says which category fired.
Branch on it in a **single** `ConditionGroup` (it stops at the first match and has an `elseActions` fallback), log
[telemetry](glossary.md#telemetry), and cancel all dialogs.

| Subcode | Catches |
|---|---|
| `OpenAIViolence` | Violence, weapons, physical harm |
| `OpenAIHate` | Hateful or discriminatory content |
| `OpenAISexual` | Sexually explicit content |
| `OpenAISelfHarm` | Self-injury or suicide content |
| `OpenAIJailBreak` | The **user** trying to override instructions |
| `OpenAIndirectAttack` | A prompt attack inside **grounding data** (documents, knowledge). Note the spelling: one `I` |

```yaml
kind: AdaptiveDialog
startBehavior: UseLatestPublishedContentAndCancelOtherTopics
beginDialog:
  kind: OnError
  id: main
  actions:
    - kind: SetVariable
      id: setVariable_r41Tm1
      variable: init:Topic.CurrentTime
      value: =Text(Now(), DateTimeFormat.UTC)
    - kind: ConditionGroup
      id: conditionGroup_r41Sw1
      conditions:
        - id: conditionItem_r41Sh1
          condition: =System.Error.SubCode = "OpenAISelfHarm"
          actions:
            - kind: SendActivity
              id: sendMessage_r41Sh2
              activity: If you or someone you know is in crisis, please contact emergency services or a crisis helpline.
        - id: conditionItem_r41Jb1
          condition: =System.Error.SubCode = "OpenAIJailBreak"
          actions:
            - kind: SendActivity
              id: sendMessage_r41Jb2
              activity: Your message was flagged as an attempt to override system instructions.
        - id: conditionItem_r41Ia1
          condition: =System.Error.SubCode = "OpenAIndirectAttack"
          actions:
            - kind: SendActivity
              id: sendMessage_r41Ia2
              activity: A prompt injection attack was detected in external data. The request has been blocked.
      elseActions:
        - kind: SendActivity
          id: sendMessage_r41El1
          activity: |-
            An error has occurred.
            Error code: {System.Error.Code}
            Error subcode: {System.Error.SubCode}
            Conversation Id: {System.Conversation.Id}
            Time (UTC): {Topic.CurrentTime}.
    - kind: LogCustomTelemetryEvent
      id: logTelemetry_r41Lg1
      eventName: OnErrorLog
      properties: "={ErrorMessage: System.Error.Message, ErrorCode: System.Error.Code, SubCode: System.Error.SubCode, TimeUTC: Topic.CurrentTime, ConversationId: System.Conversation.Id}"
    - kind: CancelAllDialogs
      id: cancelAllDialogs_r41Cx1
inputType: {}
outputType: {}
```

(The original has a branch for each of the six subcodes, with `[PLACEHOLDER]` texts to replace with approved
wording.)

A secondary approach classifies `System.Activity.Text` with an [AI Builder](glossary.md#ai-builder) prompt
([InvokeAIBuilderModelAction](reference/schema/actions.md#invokeaibuildermodelaction)). It adds about 4–6 seconds
and needs the model to be set up manually in the UI, so use it only when you need finer categories.

**Caveats.** **Azure OpenAI models only.** Anthropic (Claude) and xAI (Grok) models don't produce `ContentFiltered`
or these subcodes, so agents using them need a different approach. The pattern doesn't weaken filtering; it only
changes what the user sees. Keep `elseActions`, because new subcodes may appear.

## Teams production hardening

**Status:** recommended · [pattern file](https://github.com/microsoft/skills-for-copilot-studio/blob/main/patterns/teams-production-hardening.md)

**Problem.** Agents in [Teams](glossary.md#teams) and Microsoft 365 Copilot suffer from stale sessions, no greeting
after an app reinstall, inconsistent context, and unhelpful reset and error experiences.

**Solution.** Eight coordinated sub-patterns, meant to be applied **together and in order**, because later ones rely
on variables set by earlier ones:

| # | Sub-pattern | Mechanism |
|---|---|---|
| 1 | Handle app reinstalls | `OnActivity` with `type: InstallationUpdate` → `BeginDialog` to the Conversation Start topic (reinstalling doesn't fire the normal start) |
| 2 | Clear stale context after inactivity | [OnInactivity](reference/schema/triggers.md#oninactivity) (Teams only): clear variables and `ConversationHistory`, set `Global.InactiveConversation`, cancel dialogs |
| 3 | Notify after an inactivity reset | `OnActivity` (`Message`) when `Global.InactiveConversation` is true: reset it, send a "Session expired" [hero card](glossary.md#hero-card) with **Start over** |
| 4 | Set global context cross-channel | Low-priority (`priority: -2`) `OnActivity` initialising `Global.UserContext` when blank |
| 5 | Rebuild Reset Conversation | `OnSystemRedirect` clearing scoped variables and history, then back to Conversation Start |
| 6 | Rebuild Start Over with diagnostics | Adaptive Card question on a `YesNo` closed-list [entity](glossary.md#entity), with an "Advanced options" diagnostics panel |
| 7 | Rebuild OnError with diagnostics | Error card with the same diagnostics, `LogCustomTelemetryEvent`, `CancelAllDialogs` |
| 8 | Suggested prompts | 3–4 [conversation starters](glossary.md#conversation-starter) at agent level |

Sub-patterns 1, 2 and 5:

```yaml
kind: AdaptiveDialog
startBehavior: UseLatestPublishedContentAndCancelOtherTopics
beginDialog:
  kind: OnActivity
  id: main
  type: InstallationUpdate
  actions:
    - kind: BeginDialog
      id: beginDialog_Bqmh4L
      dialog: cr26e_myAgent.topic.ConversationStart
```

```yaml
kind: AdaptiveDialog
beginDialog:
  kind: OnInactivity
  id: main
  condition: =System.Activity.ChannelId = "msteams"
  durationInSeconds: 43200
  actions:
    - kind: ClearAllVariables
      id: clearAllVariables_mXHosp
      variables: ConversationHistory
    - kind: ClearAllVariables
      id: clearAllVariables_Vsemgr
    - kind: SetVariable
      id: setVariable_6CUITr
      variable: Global.InactiveConversation
      value: true
    - kind: CancelAllDialogs
      id: cancelAllDialogs_webE3j
inputType: {}
outputType: {}
```

```yaml
kind: AdaptiveDialog
startBehavior: UseLatestPublishedContentAndCancelOtherTopics
beginDialog:
  kind: OnSystemRedirect
  id: main
  actions:
    - kind: ClearAllVariables
      id: clearAllVariables_73bTFR
      variables: ConversationScopedVariables
    - kind: ClearAllVariables
      id: clearAllVariables_SLgE7u
      variables: ConversationHistory
    - kind: BeginDialog
      id: beginDialog_U14iCH
      dialog: cr26e_myAgent.topic.ConversationStart
    - kind: CancelAllDialogs
      id: cancelAllDialogs_12Gt21
```

Sub-pattern 3's hero card is shown in [11 Messages & Adaptive Cards](11-messages-and-adaptive-cards.md#card-attachments).
The diagnostics panel in 6 and 7 shows `System.Bot.EnvironmentId`, `System.Bot.TenantId`, `System.Bot.Name`,
`System.Bot.Id`, `System.Bot.SchemaName`, `System.User.Language`, `System.User.Id`, `System.Activity.ChannelId`,
`System.Conversation.Id` and the UTC time, with buttons sending `/debug clearstate`, `/debug clearhistory` and
`/debug conversationid`. The source gives no YAML for 6 and 8.

**Caveats.** Replace the pattern's `cat_B2EAgent.topic.*` references with your own prefix.
`startBehavior: UseLatestPublishedContentAndCancelOtherTopics` matters on 1, 5 and 7, so published content wins over
stale in-flight dialogs. Sub-pattern 2's channel condition uses `=` on `"msteams"`, which misses compound IDs such as
`msteams:Copilot` (compare [Channel-aware behavior](#channel-aware-behavior)). Its 43,200 seconds (12 h) is not one
of the schema's listed `durationInSeconds` values (900/1800/2700/3600) but is accepted as a number expression. The
pattern says suggested prompts live in "`agent.mcs.yml` or `settings.mcs.yml`"; in the vendor template,
`conversationStarters` is in `agent.mcs.yml`. Validation checklist: reinstall the app, open the agent in M365
Copilot, idle past the timeout, trigger Start Over, force an error, and check `OnErrorLog` in telemetry.

## Sources

- `vendor/skills-for-copilot-studio/skills/int-patterns/SKILL.md`: pattern index, status language, combinations, Advisor/Author roles
- `vendor/skills-for-copilot-studio/patterns/*.md`: all 15 patterns
- `vendor/skills-for-copilot-studio/templates/topics/conversation-init.topic.mcs.yml`: combined JIT topic
- `vendor/skills-for-copilot-studio/templates/agents/agent.mcs.yml`, `templates/agents/child-agent.mcs.yml`: where instructions and conversation starters live
- `vendor/skills-for-copilot-studio/agents/copilot-studio-advisor.md`, `agents/copilot-studio-author.md`: how patterns are suggested and applied
- `vendor/skills-for-copilot-studio/CONTRIBUTING.md`: pattern file format (frontmatter, sections)
