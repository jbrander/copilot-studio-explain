# 14 — Gotchas

The pitfalls from the source repo's skills, patterns and templates, plus a few found while checking them against
the schema, collected in one place. Each entry has a short explanation and a link to the guide that covers the topic
in depth.

**Contents:** [What YAML can't do](#what-yaml-cant-do) · [Lifecycle](#lifecycle) ·
[YAML syntax and values](#yaml-syntax-and-values) · [Triggers and orchestration](#triggers-and-orchestration) ·
[Tools and connectors](#tools-and-connectors) · [Knowledge](#knowledge) ·
[Messages and cards](#messages-and-cards) · [Testing](#testing) ·
[Schema and source quirks](#schema-and-source-quirks)

## What YAML can't do

### You can't create an agent from scratch

There is no supported way to scaffold `agent.mcs.yml` and `settings.mcs.yml` from nothing. Create the agent (even
an empty one) in the Copilot Studio UI, then clone it with the [VS Code extension](glossary.md#vs-code-extension)
or `manage-agent clone`. The skills' Author agent refuses to start if no `agent.mcs.yml` exists.
See [02 Project layout](02-project-layout.md#lifecycle-clone--edit--validate--push--publish).

### Some components must be created in the UI

New [connection references](glossary.md#connection-reference) (the named pointer to an authenticated
[connection](glossary.md#connection)), and therefore the **first** use of a [connector](glossary.md#connector)
(a pre-built wrapper around an external API) or MCP tool, autonomous (event and schedule) triggers, and
[AI Builder](glossary.md#ai-builder) ([Power Platform](glossary.md#power-platform)'s prompt/model service) prompt nodes all create Power Platform
resources beyond the agent's own rows. Create them in the UI, pull, then edit the YAML. You *can* reference
existing ones from new topics. The connector lookup script covers only a few connectors; for others, add the tool
in the portal first. See [08 Tools & connectors](08-tools-and-connectors.md) and
[10 Flows & autonomous triggers](10-flows-and-autonomous-triggers.md).

### YAML-only features can be lost in a UI round trip

Some features work at runtime but the Copilot Studio UI can't show them, and **editing that component in the UI
may silently remove them**:

- custom [Power Fx](glossary.md#power-fx) (the formula language) in a [knowledge source](glossary.md#knowledge-source)'s `triggerCondition` (the UI
  only has an on/off toggle, which writes `=false`);
- `OnKnowledgeRequested` custom-search topics;
- `OnGeneratedResponse` topics (per the *Prevent tool call leaks* pattern).

Warn your team before anyone opens these in the UI. See [07 Knowledge](07-knowledge.md) and
[04 Topics & triggers](04-topics-and-triggers.md).

### Some structures don't render in the canvas

YAML that is valid but uses structures the visual designer doesn't support can push fine yet fail to display in
the authoring canvas. Simplify, or make that part in the UI.
See [02 Project layout](02-project-layout.md#common-pushvalidation-errors).

## Lifecycle

### Pull before push

Push uses optimistic concurrency: each row carries a version stamp. Pushing files that are older than the server
copy fails with `ConcurrencyVersionMismatch`. Always pull first.
See [02 Project layout](02-project-layout.md#lifecycle-clone--edit--validate--push--publish).

### Push isn't publish

A push updates the **draft**, which only the authoring canvas, its Test pane and draft [evaluations](glossary.md#evaluation) see. [Teams](glossary.md#teams), web
chat, [Direct Line](glossary.md#direct-line) point tests and [Copilot Studio Kit](glossary.md#copilot-studio-kit) suites see only
[published](glossary.md#publish) content. And publishing affects **every** user the agent is shared with, so
confirm before publishing an agent with production users.
See [02 Project layout](02-project-layout.md#lifecycle-clone--edit--validate--push--publish) and
[12 Evaluations & testing](12-evaluations-and-testing.md).

### Fields you must never change

In `settings.mcs.yml`/the agent: `schemaName` (breaks the agent), `publishedOn` and `template` (platform-managed),
`language` (can corrupt the agent). In a tool: `action.operationId` / `operationDetails.operationId` and
`action.connectionReference`, which identify what runs and with which credentials. Don't hard-code an agent's name
in scripts either: every clone uses its own, so locate the folder containing `agent.mcs.yml`.
See [03 The agent](03-agent.md#settingsmcsyml) and [08 Tools & connectors](08-tools-and-connectors.md).

## YAML syntax and values

### The `=` prefix decides what a value means

`value: Topic.Count + 1` is the **literal text** "Topic.Count + 1". `value: =Topic.Count + 1` is a formula.
`activity: "Hi {Topic.Name}"` is a template (no `=`). A missing `=` is one of the most common validation errors.
The exception: variable *names*, such as `variable: Topic.X` and `AdaptiveCardPrompt` `output.binding` values,
take no `=`. See [06 Variables, entities & Power Fx](06-variables-entities-powerfx.md#power-fx).

### `init:` on the first assignment

The first `SetVariable` (or question, or output binding) that creates a [topic variable](glossary.md#topic-variable)
writes `variable: init:Topic.Name`; later assignments write `Topic.Name`. [Global variables](glossary.md#global-variable) are declared in
`variables/` instead. See [06 Variables, entities & Power Fx](06-variables-entities-powerfx.md).

### Quote values containing `{`, `: ` or `#`

YAML gives these characters meaning: a value **starting** with `{` is a mapping, `: ` inside a plain value
starts a nested key, and ` #` starts a comment. Quote the **whole** value, including any leading `=`:

```yaml
- kind: SetVariable
  id: setVariable_qU0t31
  variable: init:Topic.Record
  value: "={ Name: \"Ada\", Team: \"Ops\" }"

- kind: SendActivity
  id: sendMessage_qU0t32
  activity: "{Topic.Record.Name} is in team #{Topic.Record.Team}"
```

Unquoted, both lines fail: the first because of the `: ` inside it, the second because YAML reads the leading `{` as
the start of a mapping:

```text
value: ={ Name: "Ada" }          # parse error: "mapping values are not allowed here"
activity: {Topic.Record.Name} is in team   # '{' starts a YAML flow mapping, not text
```

Multi-line text and long expressions are safest in a `|-` block. See [README conventions](README.md#the-three-kinds-of-value).

### Only a subset of Power Fx

Copilot Studio supports a fixed list of Power Fx functions (math, text, date/time, logical, table), given in the
`int-reference` skill. Functions that work in Power Apps may not exist here, and the validator flags them.
See [06 Variables, entities & Power Fx](06-variables-entities-powerfx.md#power-fx).

### Unique node IDs, and no `_REPLACE` left over

Every node and every condition branch needs an `id` that is unique within the topic. Copy-pasted nodes keep their
IDs and fail to push. Vendor templates contain `_REPLACE1`, `_REPLACE2`, … placeholders that must all be replaced.
See [02 Project layout](02-project-layout.md#node-ids).

### Never invent a `kind:`

The `new-topic` skill calls an unverified `kind:` value "the #1 source of hallucination errors". If a kind isn't in
the [schema index](reference/schema/index.md), it doesn't exist. This applies to names in the sources too; see
[schema and source quirks](#schema-and-source-quirks).

## Triggers and orchestration

### `OnOutgoingMessage` doesn't fire

[OnOutgoingMessage](reference/schema/triggers.md#onoutgoingmessage) exists in the schema, but the `int-reference`
skill records it as **non-functional (2026-03-15)**: it doesn't fire at runtime. To act on AI-generated answers,
use `OnGeneratedResponse`. See [04 Topics & triggers](04-topics-and-triggers.md).

### `OnConversationStart` doesn't fire everywhere

In [Microsoft 365 Copilot](glossary.md#microsoft-365-copilot) and other embedded surfaces, `OnConversationStart`
never fires, so initialisation placed there silently doesn't happen. Use an `OnActivity` topic with
`type: Message` and an `IsBlank(Global.X)` condition, so it runs once, on the first real message.
See [13 Patterns](13-patterns.md#jit-glossary).

### A topic output that says "done" stops the chain

Under [generative orchestration](glossary.md#generative-orchestration), the
[orchestrator](glossary.md#orchestrator) reads a topic's outputs to decide what to do next. If a topic that only
*collects* data for a tool outputs "Your complaint has been sent!", the orchestrator believes the job is finished
and never calls the tool that would send it. Output the **data itself** (the complaint text), or drop the topic
and let the tool's inputs gather the data. Relatedly, return final results as topic **outputs** rather than
`SendActivity`, so the orchestrator writes the reply. See [04 Topics & triggers](04-topics-and-triggers.md).

### Child agents talk to users directly

[Child agents](glossary.md#child-agent) can message the user themselves (via `SendMessageTool`), bypassing the
parent. The child agent's *completion setting* doesn't change this. Tell the child, in its instructions, not to
message users and to fill output variables instead; then test, because it's a behavioural instruction, not a hard
block. Also: an [AgentDialog](reference/schema/dialogs.md#agentdialog) must **not** have `beginDialog.actions`;
child agents are driven entirely by `settings.instructions`.
See [13 Patterns](13-patterns.md#prevent-child-agent-responses) and [09 MCP & multi-agent](09-mcp-and-multi-agent.md).

### Tool-call metadata leaks into answers

The orchestrator sometimes sends raw JSON with `explanation_of_tool_call` / `new_instruction` fields to the user.
Instructions alone fix this inconsistently; an `OnGeneratedResponse` topic that sets
`System.ContinueResponse = false` on a match is more reliable. There is no platform fix yet.
See [13 Patterns](13-patterns.md#prevent-tool-call-leaks).

### MCP tools can't be forced

[MCP](glossary.md#mcp) tools are generative actions: you can't force one with `/Tool` in instructions, and the
sources say topics can't call them. Name the tool explicitly in instructions, or wrap it in a single-purpose child
agent. See [13 Patterns](13-patterns.md#deterministic-mcp-calls) and [09 MCP & multi-agent](09-mcp-and-multi-agent.md).

### `ContentFiltered` subcodes only come from Azure OpenAI models

The [RAI](glossary.md#rai) error code `ContentFiltered` and its `OpenAI*` subcodes are produced only by Azure OpenAI
models (GPT-4.1, GPT-5, …). Agents on Anthropic or xAI models never see them, so an `OnError` switch on those
subcodes does nothing there. The subcode is spelled `OpenAIndirectAttack` (one `I`).
See [13 Patterns](13-patterns.md#rai-error-handling).

### Variables the orchestrator can't see

A global variable is invisible to the orchestrator unless it has
`aIVisibility: UseInAIContext` ([AI visibility](glossary.md#ai-visibility)). And every `{Global.X}` reference in
instructions injects the full value into each prompt, so reference large values (such as a glossary) only once.
See [06 Variables, entities & Power Fx](06-variables-entities-powerfx.md) and [13 Patterns](13-patterns.md#jit-glossary).

### `ChannelId` is compound, and the Test pane isn't a channel

`System.Activity.ChannelId` can be `msteams:Copilot` or `webchat:Sharepoint`, so `= "msteams"` misses them: use
`StartsWith(Lower(...), "msteams")`. The Test pane reports its own value. And a [channel](glossary.md#channel) isn't a security boundary.
See [13 Patterns](13-patterns.md#channel-aware-behavior).

## Tools and connectors

### `$filter` quoting differs between tools and topic nodes

[OData](glossary.md#odata) (a query-string standard) parameters such as `$filter`, `$orderby` and `$top` need
different spellings in two places:

```yaml
# fragment — TaskDialog input (actions/*.mcs.yml): property name is the literal string '$filter'
- kind: ManualTaskInput
  propertyName: "'$filter'"
  value: "Status eq 'Active'"
```

```yaml
# fragment — InvokeConnectorAction inside a topic: parameters/ prefix, no inner quotes
- kind: InvokeConnectorAction
  id: invokeConnector_f1Lt3r
  connectionReference: cr26e_myAgent.shared_sharepointonline.a1b2c3
  operationId: GetItems
  input:
    binding:
      parameters/$filter: "Status eq 'Active'"
```

`$filter`, `"$filter"` and `'$filter'` alone all fail in a TaskDialog. The vendor docs write the topic form as
`input:` → `parameters/$filter:` directly, but the schema requires the `binding:` level shown above. The filter text
itself is OData (`eq`, `ne`, `gt`), not Power Fx.
See [08 Tools & connectors](08-tools-and-connectors.md#-prefixed-inputs-odata-parameters).

### `ManualTaskInput` hard-codes strings only

A [ManualTaskInput](reference/schema/task-actions.md#manualtaskinput) supplies a fixed value to a tool input, but
only **strings** can be hard-coded reliably. Review non-string values (IDs, enums) in the UI after pushing. For MCP
tools, don't add `AutomaticTaskInput`s at all: the MCP protocol discovers parameters itself, though
`ManualTaskInput`s are fine for passing context. See [08 Tools & connectors](08-tools-and-connectors.md).

## Knowledge

### Public website sources: depth ≤ 2, and Bing snippets

A [PublicSiteSearchSource](reference/schema/knowledge.md#publicsitesearchsource) URL may go **at most two path levels**
below the domain: `https://microsoft.com/products/surface` works, `…/en-us/microsoft-365/business` is
**ignored**. The source uses Bing search within that scope; it doesn't crawl or summarise whole pages, and the
site must be public. See [07 Knowledge](07-knowledge.md).

### SharePoint knowledge returns chunks, not files

[SharePoint](glossary.md#sharepoint) knowledge sources use semantic search and return relevant chunks, never the
whole file. Anything that needs complete file content (a glossary, a config file) must come from
[Dataverse](glossary.md#dataverse) storage or an [agent flow](glossary.md#agent-flow). Dynamic knowledge URLs held in a global variable must be
clean folder URLs, not `AllItems.aspx?…` links. See [07 Knowledge](07-knowledge.md).

### Search with a rewritten query, not the raw message

The `add-generative-answers` skill says never to pass `=System.Activity.Text` straight to
`SearchAndSummarizeContent`, because a follow-up like "tell me more about that" has no context. Run
[CreateSearchQuery](reference/schema/actions.md#createsearchquery) first and search its `SearchQuery` result. Some
vendor templates and examples still pass `System.Activity.Text` directly. Inside `OnKnowledgeRequested`, use
`System.SearchQuery`, which is already rewritten. See [07 Knowledge](07-knowledge.md).

## Messages and cards

### Plain newlines collapse

Most channels render YAML newlines as spaces. Use `<br />` / `<br /><br />`.
See [11 Messages & Adaptive Cards](11-messages-and-adaptive-cards.md#line-breaks-and-formatting).

### Adaptive Card rules that bite

- `AdaptiveCardPrompt` needs `output`, `outputType` **and** an `Action.Submit`, even for display-only cards.
- The JSON goes under `card: |` (literal block), never `>` or inline.
- JSON cards are static: `{Topic.X}` isn't interpolated. Use a message next to the card, or a Power Fx formula card.
- `style: "Email"` / `"Tel"` don't validate input; use `regex` and `isRequired`.
- Declare `"version": "1.5"` for Teams; the bundled schema is 1.6.

See [11 Messages & Adaptive Cards](11-messages-and-adaptive-cards.md#adaptive-cards).

### `kind: Activity` is not a message

`kind: Activity` is the non-message activity template (default `activityType: Typing`). For a rich message, write
`activity:` as an object with `text:`, `speak:`, `attachments:`, and no `kind` (or `kind: Message`).
See [11 Messages & Adaptive Cards](11-messages-and-adaptive-cards.md#message-object).

## Testing

### Draft vs published, auth and quotas

Point tests and Kit suites need a **published** agent; the Evaluation API can test the draft. Direct Line tokens
are bound to one conversation and expire after about 30 minutes. An evaluation run without a connection ID runs
anonymously and **skips** authenticated tools and knowledge. The API allows one run at a time and 20 runs per
agent per 24 hours. A CSV test-set import ignores any test-method column.
See [12 Evaluations & testing](12-evaluations-and-testing.md).

## Schema and source quirks

### The schema's descriptions are placeholders

Every `title`/`description` in `bot.schema.yaml-authoring.json` is a token such as
`BotSchema:$$OnRecognizedIntent_description$$`, so the schema tells you shapes, not meanings. The generated
[reference](reference/schema/index.md) inherits this; the explanations in these guides come from the skills,
patterns and templates.

### `modelDescription: null` in a vendor template

`templates/topics/conversation-init.topic.mcs.yml` (and the JIT glossary / JIT user context pattern examples) set
`modelDescription: null`. The schema types `modelDescription` as a string, so strict validation rejects it; this
docs project's checker flags it. The topic is triggered by `OnActivity`, not by the orchestrator, so simply omit
`modelDescription`. See [13 Patterns](13-patterns.md#jit-glossary).

### On-disk files are flattened

`agent.mcs.yml` carries `displayName` at its root, and `variables/*.mcs.yml` puts the variable's fields at the root
instead of under `variable:`. Validators that use the raw schema must undo this first. In `agent.mcs.yml`, the
model block has no `kind` (`CurrentModelsNoKind`). See [02 Project layout](02-project-layout.md#on-disk-shape-vs-schema-shape).

### "Instructions in `settings.mcs.yml`"

Several patterns say to put instructions "in `settings.mcs.yml`". The agent's `instructions` belong in
`agent.mcs.yml` ([GptComponentMetadata](reference/schema/components.md#gptcomponentmetadata)), and a child agent's
in `settings.instructions` of its AgentDialog. See [13 Patterns](13-patterns.md#how-patterns-are-chosen).

### Vendor examples that don't match the schema

- *Prevent child agent responses* uses `kind: VariableOutputBindings` / `VariableOutputBinding`, which don't
  exist; use `output.binding`.
- *Deterministic MCP calls* puts `intent.triggerQueries` under `OnToolSelected`, which the schema doesn't allow.
- `InvokeConnectorAction` examples omit the `binding:` level under `input:` (see [above](#filter-quoting-differs-between-tools-and-topic-nodes)).
- `System.SearchResults` rows are `{snippet, title, url}` in the knowledge guide but `Content, ContentLocation, Title`
  in the custom-knowledge-source template and `int-reference`. The latter matches the schema's
  [CustomSearchInputContent](reference/schema/knowledge.md#customsearchinputcontent); prefer it.
- `add-global-variable` offers `aIVisibility: HideFromAIContext`; the schema value is `Hidden`
  ([VariableAIVisibility](reference/schema/actions.md#variableaivisibility)).
- `int-reference` spells `EMailPrebuiltEntity`; the schema kind is
  [EmailPrebuiltEntity](reference/schema/entities.md#emailprebuiltentity).
- `knowledge-guide.md` and `instructions-guide.md` mention a `knowledgeSourceIds` filter; the
  `SearchAndSummarizeContent` property is `knowledgeSources:` (`SearchSpecificKnowledgeSources`). See
  [07 Knowledge](07-knowledge.md).
- `knowledge-guide.md`'s SharePoint-list example uses `connectionProperties: kind: ConnectionReferenceBySchema`; no
  such kind exists, and `connectionProperties` allows only `name`, `mode` and `usageDescription`.
- `knowledge-guide.md` refers to `templates/knowledge/graph-connector.knowledge.mcs.yml`, which isn't in the repo.
- `edit-agent` places web browsing at `configuration.settings.*.content.capabilities.webBrowsing`; the schema has
  `gptCapabilities.webBrowsing` in `agent.mcs.yml`. See [03 The agent](03-agent.md).
- `int-patterns` mentions an `OnInstallationUpdate` trigger; it doesn't exist. Use `OnActivity` with
  `type: InstallationUpdate` ([04](04-topics-and-triggers.md)).
- The conversation-init template and JIT user-context pattern call Office 365 Users operation `UserGet_V2`, which
  isn't in the bundled [Office 365 Users catalogue](reference/connectors/shared_office365users.md) (whose
  current-user operation is `MyProfile_V2`). It may exist in the live connector; verify in the UI.
- The question-topic template compares text with single quotes (`= '<Option 1>'`). In Power Fx single quotes
  delimit *names*; use double quotes for text (`="Option 1"`).

See [13 Patterns](13-patterns.md) for corrected versions.

### Strict JSON-Schema validators reject `=` dialog expressions

In the schema, `BeginDialog.dialog` (`DialogExpression`) is a `oneOf` of "any string" and "string matching `=.*`".
A `=Switch(...)` value matches both, so a strict generic validator rejects it. The skills' own validator and the
Copilot Studio tooling accept it. This affects only the
[dynamic topic redirect pattern](13-patterns.md#dynamic-topic-redirect) style of `dialog:` value. (This docs
project's `tools/check_docs.py` treats the schema's `oneOf` as `anyOf` for the same reason.)

## Sources

- `vendor/skills-for-copilot-studio/skills/int-reference/SKILL.md`: YAML-only features, `OnOutgoingMessage`, `$filter` formats, `ManualTaskInput`, Power Fx function list, `init:`
- `vendor/skills-for-copilot-studio/skills/new-topic/SKILL.md`: topic outputs vs chaining, unverified kinds, `TaskOutput`
- `vendor/skills-for-copilot-studio/agents/copilot-studio-author.md`: what can't be created in YAML
- `vendor/skills-for-copilot-studio/skills/manage-agent/SKILL.md`, `SETUP_GUIDE.md`: pull before push, publish warnings
- `vendor/skills-for-copilot-studio/skills/edit-agent/SKILL.md`, `skills/edit-action/SKILL.md`, `skills/edit-action/sharepoint-actions.md`, `skills/add-action/SKILL.md`: fields never to change, OData quoting, MCP inputs
- `vendor/skills-for-copilot-studio/skills/add-knowledge/SKILL.md`, `skills/add-knowledge/knowledge-guide.md`, `skills/add-generative-answers/SKILL.md`: URL depth, SharePoint limits, `CreateSearchQuery`, `System.SearchResults`
- `vendor/skills-for-copilot-studio/skills/add-other-agents/SKILL.md`: no `beginDialog.actions` on child agents
- `vendor/skills-for-copilot-studio/skills/add-adaptive-card/SKILL.md`: card rules
- `vendor/skills-for-copilot-studio/patterns/*.md`: pattern-specific pitfalls
- `vendor/skills-for-copilot-studio/templates/topics/conversation-init.topic.mcs.yml`, `templates/topics/custom-knowledge-source.topic.mcs.yml`, `templates/topics/search-topic.topic.mcs.yml`: template quirks
- `vendor/skills-for-copilot-studio/skills/run-eval/SKILL.md`, `skills/chat-directline/SKILL.md`, `skills/create-eval-set/SKILL.md`, `scripts/src/eval-api.js`: testing limits
- `vendor/skills-for-copilot-studio/reference/bot.schema.yaml-authoring.json`: placeholder descriptions, `DialogExpression`
