# 09 — MCP & multi-agent

This page covers two ways to give an agent capabilities that live somewhere else. The first is
[MCP](glossary.md#mcp) servers (Model Context Protocol: an open standard for exposing a set of tools to an LLM).
The second is other agents: [child agents](glossary.md#child-agent) defined inside your project,
[connected agents](glossary.md#connected-agent) (separately published [Copilot Studio](glossary.md#copilot-studio) agents), and
[external agents](glossary.md#external-agent) reached over [A2A](glossary.md#a2a) (the Agent-to-Agent protocol). For
ordinary [tools](glossary.md#tool) that call a [connector](glossary.md#connector) operation, see
[08 Tools & connectors](08-tools-and-connectors.md).

## Which one do I need?

| You want to… | Use | File | `kind:` |
|---|---|---|---|
| Give the agent every tool an MCP server offers | MCP server tool | `actions/*.mcs.yml` | `TaskDialog` + [InvokeExternalAgentTaskAction](reference/schema/task-actions.md#invokeexternalagenttaskaction) + [ModelContextProtocolMetadata](reference/schema/components.md#modelcontextprotocolmetadata) |
| Call one MCP tool at a fixed point in a topic | MCP tool node (schema-only; see the caveats) | node in a topic | [InvokeMcpToolAction](reference/schema/actions.md#invokemcptoolaction) |
| Split a big agent into specialists you own | Child agent | `agents/<Name>/agent.mcs.yml` | [AgentDialog](reference/schema/dialogs.md#agentdialog) |
| Call another team's published Copilot Studio agent | Connected agent | `actions/*.mcs.yml` | `TaskDialog` + [InvokeConnectedAgentTaskAction](reference/schema/task-actions.md#invokeconnectedagenttaskaction) |
| Call an agent hosted outside Copilot Studio | External agent | `actions/*.mcs.yml` | `TaskDialog` + `InvokeExternalAgentTaskAction` + [AgentToAgentProtocolMetadata](reference/schema/components.md#agenttoagentprotocolmetadata) or [ActivityProtocolMetadata](reference/schema/components.md#activityprotocolmetadata) |

In every case the parent's [orchestrator](glossary.md#orchestrator) (the LLM that plans each turn) decides when to
call the other side, based on its description. There is one exception: the topic node runs where you put it.

## MCP servers as tools

An MCP server is a service that advertises a list of tools (name, description and input schema) and runs them on
request. In Copilot Studio, **one `TaskDialog` represents the whole server**. The orchestrator discovers the server's
tools dynamically and fills in their parameters itself.

```yaml
mcs.metadata:
  componentName: Microsoft Learn Docs MCP Server
  description: Searches Microsoft Learn documentation
kind: TaskDialog
modelDisplayName: Microsoft Learn Docs MCP Server
modelDescription: MCP Server to search Microsoft Learn content
action:
  kind: InvokeExternalAgentTaskAction
  connectionReference: cr26e_evalTest.shared_microsoftlearndocsmcpserver.abc123
  connectionProperties:
    mode: Invoker
  operationDetails:
    kind: ModelContextProtocolMetadata
    operationId: microsoft_docs_search
```

(This example is the vendor's eval fixture `agent-with-mcp-action/actions/SearchDocs.mcs.yml`.)

| Property | Type / values | Req. | Meaning |
|---|---|---|---|
| `modelDisplayName`, `modelDescription` | string | | How the orchestrator sees the server. **Keep `modelDescription` on one line**: multi-line descriptions have been reported to break MCP tool registration after push (vendor issue #91). |
| `action.connectionReference` | string | ✓ | The [connection reference](glossary.md#connection-reference) created by the UI. **Never change it.** |
| `action.connectionProperties.mode` | `Maker` \| `Invoker` | | Whose credentials are used. See [08](08-tools-and-connectors.md#maker-vs-invoker). |
| `action.operationDetails.kind` | `ModelContextProtocolMetadata` | ✓ | Marks the action as MCP. |
| `action.operationDetails.operationId` | string | ✓ | Identifies the MCP operation. **Never change it.** |
| `action.historyType` | `kind: ConversationHistory` \| `kind: NoHistory` | | No vendor guidance. *Inferred:* how much conversation history is sent. |

Further `ModelContextProtocolMetadata` properties, with no vendor guidance (descriptions inferred from names):

| Property | Inferred meaning |
|---|---|
| `tools` | `kind: UseAllTools`, or `kind: UseSpecificTools` with a `tools:` list of names. Limits which of the server's tools are exposed. |
| `knownTools` | Cached tool definitions (`name`, `description`, `inputSchema`, `outputSchema`). |
| `toolsLastRefreshedAt` | When the cached list was last refreshed. |
| `isElicitationEnabled` | Lets the server ask the user follow-up questions (MCP "elicitation"). |
| `dynamicInputSchema` | Resolved input shape. |

### Adding an MCP server

You can't add one from YAML, for the same reason as connector tools: the
[connection](glossary.md#connection) has to be created by a person in the portal. In Copilot Studio, go to your
agent → **Tools** (older UI: **Actions**) → **+ Add** → search for the MCP server connector (for example *Microsoft
Learn Docs MCP Server*) → configure the connection → save. Then **pull**. After that you can edit
`modelDisplayName`, `modelDescription` and `connectionProperties.mode`. See
[08 — add in the UI, then pull](08-tools-and-connectors.md#adding-a-connector-tool-add-in-the-ui-then-pull).

### Inputs on an MCP tool

- **Don't add `AutomaticTaskInput` entries.** The MCP protocol discovers tool parameters dynamically.
- **`ManualTaskInput` entries are fine** for passing context to the server, and here they can be
  [Power Fx](glossary.md#power-fx) (formula) expressions. A common use is passing the user's identity:

```yaml
mcs.metadata:
  componentName: Custom MCP Server
  description: Custom MCP server with user context
kind: TaskDialog
inputs:
  - kind: ManualTaskInput
    propertyName: userid
    value: =System.User.Email
modelDisplayName: Custom MCP Server
modelDescription: Custom MCP server for internal tools
action:
  kind: InvokeExternalAgentTaskAction
  connectionReference: cr26e_evalTest.shared_custommcp.def456
  connectionProperties:
    mode: Maker
  operationDetails:
    kind: ModelContextProtocolMetadata
    operationId: InvokeMCP
```

(From the vendor's eval fixture `CustomMCP.mcs.yml`. The vendor describes the value as reaching the server "e.g., as a
header"; the exact transport isn't documented.)

### MCP approval kinds

The schema defines [AlwaysRequireMCPApproval](reference/schema/settings.md#alwaysrequiremcpapproval) and
[NeverRequireMCPApproval](reference/schema/settings.md#neverrequiremcpapproval). Each has an optional `tools` list
(or `=` expression). By their names, they decide whether a user must approve an MCP tool call before it runs.
However, **no other schema definition references them**, so there's no documented place to put them in an agent
file. The vendor repo doesn't mention them either. Treat them as not yet usable from YAML.

## Calling an MCP tool from a topic

MCP tools are *generative*: the orchestrator decides whether to call them. Sometimes a specific tool must run for a
specific request every time.

### The InvokeMcpToolAction node (schema only)

The schema includes a topic node for a direct call:

```yaml
# fragment — schema shape only; no vendor example exists (see warning)
- kind: InvokeMcpToolAction
  id: invokeMcp_Lr4n9X
  serverSchemaName: <schema name of the MCP server tool>
  toolName: microsoft_docs_search
  input:
    binding:
      query: =Topic.SearchQuery.SearchQuery
  output:
    kind: SingleVariableOutputBinding
    variable: init:Topic.DocsResult
```

| Property | Req. | Meaning (inferred) |
|---|---|---|
| `serverSchemaName` | ✓ | Which MCP server tool to use: the schema name of its `TaskDialog` component. |
| `toolName` | ✓ | Which of the server's tools to call. |
| `input.binding` | | Tool arguments, by parameter name. |
| `output` | | Where the result goes. |

> **Caveat.** The vendor pattern *Deterministic MCP calls* (status: *experimental*) says the opposite: "topics cannot
> call MCP tools directly … there is currently no topic node you can drop in". So the node exists in the schema, but
> the vendor guidance doesn't treat it as usable, and there's no example of the `serverSchemaName` format. If you try
> it, validate with the [language server](glossary.md#lsp) (`manage-agent validate`) and test before relying on it.

### Workarounds from the vendor pattern

See [13 Patterns — Deterministic MCP calls](13-patterns.md#deterministic-mcp-calls). There are two options, in
order of effort:

| | Option 1: name the tool in instructions | Option 2: child-agent wrapper |
|---|---|---|
| How | In the agent's [instructions](glossary.md#instructions), map the intent to the tool by name ("When the user asks about X, you MUST call the *Y* tool"), and list example utterances | A [child agent](#child-agents) whose only job is to call that tool, with a narrow `OnToolSelected` description |
| Reliability | High but not guaranteed | Near-deterministic |
| Cost | Edit one block | A new child agent and another hop |

You can't force an MCP tool with the `/ToolName` reference syntax in instructions, as you can with topics and
other tools. If the parent should control the final wording, have the wrapper child return outputs instead of
replying itself (see [below](#preventing-child-agents-from-messaging-the-user)).

## Child agents

A **child agent** is a specialist sub-agent that lives in the parent's project and is owned by the same team. It has
its own instructions (and optionally its own knowledge). The parent's orchestrator delegates to it like a tool.

File: `agents/<ChildAgentName>/agent.mcs.yml`, `kind:` [AgentDialog](reference/schema/dialogs.md#agentdialog). The
first two comment lines are the display name and the internal name:

```yaml
# Name: Billing Specialist
# BillingSpecialist
kind: AgentDialog

beginDialog:
  kind: OnToolSelected
  id: main
  description: Handles billing inquiries, payment issues, refund requests and subscription changes. Route here when users ask about charges, invoices or account billing, e.g. "why was I charged twice?" or "cancel my subscription".

settings:
  instructions: |
    You are a billing support specialist. Help customers with:
    - Understanding their charges and invoices
    - Processing refund requests
    - Managing subscription plans

    Always verify the customer's account before making changes.
    Escalate to a human agent for refunds over $500.

inputs:
  - kind: AutomaticTaskInput
    propertyName: CustomerQuery
    description: The customer's billing-related question or issue

inputType:
  properties:
    CustomerQuery:
      displayName: Customer Query
      description: The customer's billing-related question or issue
      type: String

outputType: {}
```

| Property | Type | Req. | Meaning |
|---|---|---|---|
| `beginDialog.kind` | `OnToolSelected` | | The [trigger](glossary.md#trigger) that makes this a delegatable agent. |
| `beginDialog.description` | string | ✓ (in `OnToolSelected`) | **The routing signal.** The parent orchestrator reads it to decide when to delegate. Be specific and action-oriented, and include example user messages. |
| `beginDialog.condition` | `=` expression | | Restricts when the child is available, for example only on one [channel](glossary.md#channel). |
| `beginDialog.id` | string | ✓ | Conventionally `main`. |
| `settings.instructions` | string | | The child's system prompt: its scope, tone and rules. |
| `settings.*` | other `GptComponentMetadata` fields | | The schema also allows `responseInstructions`, `knowledgeSources`, `tools`, `aISettings`, … No vendor guidance on these. |
| `inputs` | list of `AutomaticTaskInput` | | Context the parent passes. The orchestrator fills these automatically. |
| `inputType` | record type | | Input declarations. They must match `inputs`. Use `inputType: {}` if there are none. |
| `outputType` | record type | | What the child returns. Use `outputType: {}` if there's nothing structured. |

Rules from the vendor `add-other-agents` skill:

- **No `beginDialog.actions`.** A child agent is driven entirely by `settings.instructions` (generative
  orchestration). If you find yourself wanting `Question`, `SendActivity` or `ConditionGroup` nodes, write a topic in
  the parent instead.
- **Two-phase creation.** (1) Create only `agent.mcs.yml` and **push**. (2) Only then add `knowledge/` files under
  the child, because [knowledge sources](glossary.md#knowledge-source) reference the agent they belong to, so it must already exist in the
  environment.
- `OnToolSelected` accepts only `id`, `description`, `condition`, `priority`, `runOnce` and `actions`. The vendor's
  *Deterministic MCP calls* example adds `intent: triggerQueries: …` to it. The schema rejects that, so leave it out
  and put example phrasings in `description`.

### Preventing child agents from messaging the user

By default a child agent can talk to the user directly (through its built-in `SendMessageTool`), bypassing the
parent's control over formatting, filtering and tone. To stop it, put an explicit rule in the **child's**
instructions and give the child output variables to return its answer through:

```yaml
# fragment — the child agent's settings block
settings:
  instructions: |
    CRITICAL - DO NOT MESSAGE USERS
    - DO NOT respond directly to the user
    - DO NOT call SendMessageTool or send any messages
    - ONLY populate the output variables with your response
    - Let the parent orchestrator deliver the response to the user
```

- **This is enforced by instructions, not by the platform.** Test that the child complies.
- **Define outputs** (for example `outputType.properties.AgentOutput`). Without them, the child has no way to return
  data.
- The child's **completion setting** (continue, end or return after it finishes) is unrelated. It doesn't stop
  direct messages.

The vendor pattern then has a parent topic call the child explicitly and read its output. The pattern's YAML uses
output-binding kinds that **don't exist in the schema**. A schema-valid equivalent uses a plain `binding:` map. Its
schema-name format, `<agentSchema>.agent.<Name>`, comes from the pattern and isn't otherwise documented:

```yaml
- kind: BeginDialog
  id: callChild_Ch1ld1
  dialog: cr26e_myAgent.agent.BillingSpecialist
  input:
    binding:
      CustomerQuery: =System.Activity.Text
  output:
    binding:
      AgentOutput: Topic.ChildResult

- kind: SendActivity
  id: sendResult_Ch1ld2
  activity: "{Topic.ChildResult}"
```

See [13 Patterns — Prevent child agent responses](13-patterns.md#prevent-child-agent-responses).

> The vendor patterns say to put these instructions "in the child agent's `settings.mcs.yml`". For a child agent, the
> instructions actually live in `settings.instructions` inside `agents/<Name>/agent.mcs.yml`, as shown above.

## Connected agents

A **connected agent** is an independently built and published Copilot Studio agent, owned by another team or project,
that your agent calls as a tool. The calling side is a `TaskDialog` in your `actions/`:

```yaml
# Name: Expense Report Processor
# Description: Sends expense reports to the expense processing agent
kind: TaskDialog
modelDisplayName: Expense Report Processor
modelDescription: Use this agent to process expense reports by sending them to the expense processing agent.
inputs:
  - kind: AutomaticTaskInput
    propertyName: expenseReportFileFullPath
    description: Full file path, including SharePoint site URL, of an expense report file.
action:
  kind: InvokeConnectedAgentTaskAction
  inputType:
    properties:
      expenseReportFileFullPath:
        displayName: expenseReportFileFullPath
        isRequired: true
        type: String
  botSchemaName: cr26e_expenseAgent
  historyType:
    kind: ConversationHistory
```

| Property | Type / values | Req. | Meaning |
|---|---|---|---|
| `botSchemaName` | string | ✓ | The connected agent's `schemaName`, from its `settings.mcs.yml` or *Agent details* in the UI. |
| `inputType` / `outputType` | record types | | Data passed to it and returned from it. |
| `historyType` | `kind: ConversationHistory` (`numberOfPastUserMessagesToInclude`, default 5) \| `kind: NoHistory` | | How much of the conversation it sees. |
| `interruptionPolicy` | `allowInterruption` (default `true`), `interruptionTopicListFilter` | | No vendor guidance. *Inferred:* whether the user can switch topics while it's running. |
| `shouldSendStartConversation` | boolean, default `false` | | No vendor guidance. *Inferred:* whether to send it a conversation-start event first. |

**The connected agent must be published and reachable** from the same [environment](glossary.md#environment) or
[tenant](glossary.md#tenant). Its owners must set up the **called side**:

1. An **`OnRedirect` topic** as the entry point when another agent calls it.
2. A **[global variable](glossary.md#global-variable) for each input**, with "Receive values from other agents" (external source) enabled in the
   UI.
3. `inputType: {}` and `outputType: {}` on that topic, to signal that it takes part in multi-agent calls.

```yaml
# OnRedirect topic on the connected agent
kind: AdaptiveDialog
modelDescription: Initializes the agent when called by another agent
beginDialog:
  kind: OnRedirect
  id: main
  actions:
    - kind: SetVariable
      id: setVariable_Xp3nse
      variable: Global.expenseReportFileFullPath
      value: =System.Activity.Value.expenseReportFileFullPath
inputType: {}
outputType: {}
```

[InvokeAgentTaskAction](reference/schema/task-actions.md#invokeagenttaskaction) (a required
`gptComponentSchemaName`, plus `inputType`/`outputType`) is a third agent-calling kind. The vendor repo has no
guidance on it. Judging by the property name, it targets an agent component by schema name.

## External agents (A2A and Activity protocol)

Agents hosted outside Copilot Studio are called with the same `InvokeExternalAgentTaskAction` as MCP. Only the
`operationDetails` kind differs:

| `operationDetails.kind` | Protocol | Extra properties |
|---|---|---|
| [ModelContextProtocolMetadata](reference/schema/components.md#modelcontextprotocolmetadata) | MCP | see [above](#mcp-servers-as-tools) |
| [AgentToAgentProtocolMetadata](reference/schema/components.md#agenttoagentprotocolmetadata) | A2A, an open agent-to-agent HTTP protocol | `interruptionPolicy`, `extensions.inputType`, `extensions.outputType`, `dynamicInputSchema` |
| [ActivityProtocolMetadata](reference/schema/components.md#activityprotocolmetadata) | The [Bot Framework](glossary.md#bot-framework) activity protocol (inferred from the name) | `interruptionPolicy`, `dynamicInputSchema` |

> **Schema-only.** The vendor repo has no guidance or examples for A2A or Activity-protocol agents. The shape below is
> valid against the schema, but the `operationId` value and the connection reference are placeholders. As with MCP,
> expect to add the agent in the UI first (which creates the connection), then pull and edit.

```yaml
# Name: Partner order-status agent
kind: TaskDialog
modelDisplayName: Partner order-status agent
modelDescription: Asks the logistics partner's agent for the shipping status of an order. Use when the user asks where their order is.
action:
  kind: InvokeExternalAgentTaskAction
  connectionReference: cr26e_myAgent.shared_a2aagent.5e7d1b
  historyType:
    kind: ConversationHistory
    numberOfPastUserMessagesToInclude: 3
  operationDetails:
    kind: AgentToAgentProtocolMetadata
    operationId: InvokeAgent
    interruptionPolicy:
      allowInterruption: false
```

## Sources

- `vendor/skills-for-copilot-studio/skills/add-other-agents/SKILL.md`: child and connected agents, called-side setup
- `vendor/skills-for-copilot-studio/templates/agents/child-agent.mcs.yml`
- `vendor/skills-for-copilot-studio/templates/actions/mcp-action.mcs.yml`
- `vendor/skills-for-copilot-studio/evals/fixtures/agent-with-mcp-action/actions/SearchDocs.mcs.yml`, `CustomMCP.mcs.yml`
- `vendor/skills-for-copilot-studio/skills/add-action/SKILL.md`, `skills/edit-action/SKILL.md`: MCP creation and editing rules
- `vendor/skills-for-copilot-studio/patterns/deterministic-mcp-calls.md`, `patterns/prevent-child-agent-responses.md`
- `vendor/skills-for-copilot-studio/skills/int-reference/SKILL.md`: `OnToolSelected`
- Generated reference: [task-actions.md](reference/schema/task-actions.md), [dialogs.md](reference/schema/dialogs.md#agentdialog), [components.md](reference/schema/components.md#protocolmetadatabase), [actions.md](reference/schema/actions.md#invokemcptoolaction), [settings.md](reference/schema/settings.md)
