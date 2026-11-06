# Glossary

Plain-language definitions of the Microsoft, Power Platform and Copilot Studio terms used throughout these docs.
Guides link here the first time they use a term. Entries are alphabetical. Where a term has an exact YAML
spelling (a `kind:` or property name), it is shown in `code`.

Jump to: [A](#a) · [B](#b) · [C](#c) · [D](#d) · [E](#e) · [F](#f) · [G](#g) · [H](#h) · [I](#i) ·
[K](#k) · [L](#l) · [M](#m) · [N](#n) · [O](#o) · [P](#p) · [Q](#q) · [R](#r) · [S](#s) · [T](#t) ·
[V](#v) · [Y](#y)

---

## A

### A2A
**Agent-to-Agent protocol.** An open protocol that lets AI agents built on different platforms call each other
over HTTP. Copilot Studio can call an A2A-compatible agent as a tool, described in YAML as
`InvokeExternalAgentTaskAction` with `AgentToAgentProtocolMetadata`. See also [External agent](#external-agent) and
[09 MCP & multi-agent](09-mcp-and-multi-agent.md).
Learn more: [a2a-protocol.org](https://a2a-protocol.org/)

### Action
An overloaded word in Copilot Studio, with two meanings:
1. **A node in a topic**: one step in a topic's `actions:` list, such as sending a message or asking a question.
   See [Node](#node) and [05 Nodes](05-nodes.md).
2. **A tool**: a file in the `actions/` folder (`kind: TaskDialog`) that wraps a [connector](#connector) operation,
   an [MCP](#mcp) server, a [flow](#cloud-flow) or another agent, so the [orchestrator](#orchestrator) can call it.
   The Copilot Studio UI now calls these **tools**. See [Tool](#tool) and [08 Tools & connectors](08-tools-and-connectors.md).

### Activity
The [Bot Framework](#bot-framework) word for **anything sent in a conversation**: a chat message, a typing
indicator, an event, a card. In YAML, `SendActivity` sends one, `System.Activity.Text` is the text of the user's
latest message, and the `OnActivity` trigger fires when a given activity type arrives.
See [11 Messages & Adaptive Cards](11-messages-and-adaptive-cards.md).

### Adaptive Card
A **JSON description of a small UI** (text blocks, images, fact lists, input fields, buttons). The chat client
renders it natively, so the same card looks right in Teams, web chat and Outlook. Copilot Studio can send a card as
a message (`AdaptiveCardTemplate`) or use one as a form that waits for the user to submit (`AdaptiveCardPrompt`).
The bundled card schema is version 1.6; the skills' examples declare `"version": "1.5"` for broad client support. See [11 Messages & Adaptive Cards](11-messages-and-adaptive-cards.md) and the
[Adaptive Card schema reference](reference/adaptive-card-schema.md).
Learn more: [Adaptive Cards (Microsoft Learn)](https://learn.microsoft.com/en-us/adaptive-cards/) ·
[Ask with an Adaptive Card](https://learn.microsoft.com/en-us/microsoft-copilot-studio/authoring-ask-with-adaptive-card)

### Agent
A conversational AI assistant built in [Copilot Studio](#copilot-studio) (older docs call it a "bot" or "copilot",
which is why the schema still uses `Bot…` names). An agent is a bundle of [instructions](#instructions),
[topics](#topic), [tools](#tool), [knowledge sources](#knowledge-source) and [variables](#global-variable),
stored as a set of `.mcs.yml` files. See [01 Concepts](01-concepts.md) and [03 The agent](03-agent.md).
Learn more: [What is Copilot Studio?](https://learn.microsoft.com/en-us/microsoft-copilot-studio/fundamentals-what-is-copilot-studio)

### Agent flow
A deterministic, step-by-step automation (like a [Power Automate cloud flow](#cloud-flow)) that is built and stored
inside Copilot Studio and can be called by an agent as a tool or from a topic. In YAML it appears through
`InvokeFlowAction` / `InvokeFlowTaskAction` with a `flowId`. See [10 Flows & autonomous triggers](10-flows-and-autonomous-triggers.md).
Learn more: [Agent flows overview](https://learn.microsoft.com/en-us/microsoft-copilot-studio/flows-overview)

### AI Builder
A Power Platform service for prebuilt and custom AI models: document processing, classification, and
**prompts** (reusable LLM instructions with inputs). Topics call them with `InvokeAIBuilderModelAction`. The skills
repo can't author AI Builder prompts in YAML; create them in the UI and reference them.
Learn more: [AI Builder overview](https://learn.microsoft.com/en-us/ai-builder/overview) ·
[Prompts in Copilot Studio](https://learn.microsoft.com/en-us/microsoft-copilot-studio/prompts-overview)

### AI visibility
The `aIVisibility` setting on a [global variable](#global-variable). `UseInAIContext` lets the
[orchestrator](#orchestrator) see the variable's value and reason about it. `Hidden` keeps it out of the model's
context (use this for internal flags). See [06 Variables](06-variables-entities-powerfx.md).

### App registration
An identity for an application, created in [Entra ID](#entra-id) (Azure Portal → *App registrations*). Test tools
use one to sign a user in and call Microsoft APIs on their behalf: the skills repo's `test-auth` skill needs an app
registration with delegated Power Platform API permissions (`CopilotStudio.MakerOperations.Read`/`ReadWrite`,
`CopilotStudio.Copilots.Invoke`), redirect URI `http://localhost` and public client flows enabled, to run
evaluations and chat with integrated-auth agents. See [12 Evaluations & testing](12-evaluations-and-testing.md#test-auth-the-app-registration).
Learn more: [Register an application](https://learn.microsoft.com/en-us/entra/identity-platform/quickstart-register-app)

### Application Insights
Azure's application monitoring service. Copilot Studio can send conversation telemetry to it, including custom
events you emit with the `LogCustomTelemetryEvent` node. See [Telemetry](#telemetry).
Learn more: [Application Insights overview](https://learn.microsoft.com/en-us/azure/azure-monitor/app/app-insights-overview) ·
[Capture telemetry](https://learn.microsoft.com/en-us/microsoft-copilot-studio/advanced-bot-framework-composer-capture-telemetry)

### Authentication mode
How an agent identifies its end users, set by `authenticationMode` in `settings.mcs.yml`: `None` (anonymous),
manual (`GenericOAuth2` or `CustomAzureActiveDirectory`: you configure an identity provider), or `Integrated`
([Entra ID](#entra-id) single sign-on, as in Teams / Microsoft 365). It decides how test tools can chat with the agent: [Direct Line](#direct-line) for no or manual auth,
the Copilot Studio SDK for integrated auth. See [12 Evaluations & testing](12-evaluations-and-testing.md).
Learn more: [Configure user authentication](https://learn.microsoft.com/en-us/microsoft-copilot-studio/configuration-end-user-authentication)

### Autonomous agent
An agent that acts without a user typing to it. An **event trigger** (a new email, a SharePoint item changed, a
schedule) starts it and hands it a payload to act on. In YAML these show up as `ExternalTriggerComponent`s and the
`AutonomousSettings` settings block. Creating the trigger needs the Copilot Studio UI because it also creates
Power Platform resources. See [10 Flows & autonomous triggers](10-flows-and-autonomous-triggers.md).
Learn more: [Event triggers](https://learn.microsoft.com/en-us/microsoft-copilot-studio/authoring-triggers-about)

### Azure AI Search
Azure's managed search index service (formerly *Azure Cognitive Search*). It can be a [knowledge source](#knowledge-source)
(`AzureAISearchSource`) when your documents are already indexed there.
Learn more: [What is Azure AI Search?](https://learn.microsoft.com/en-us/azure/search/search-what-is-azure-search)

### Azure OpenAI
Microsoft's hosted version of OpenAI's models, run inside an Azure subscription. Its **"On Your Data"** feature
lets a model answer from your own Azure AI Search index. The schema can describe such a deployment as a knowledge
source (`AzureOpenAIOnYourDataKnowledgeSource`), but the skills repo gives no guidance on it. See
[07 Knowledge](07-knowledge.md).

## B

### Bing Custom Search
A Bing feature that limits web search to a list of sites you choose. It appears as the knowledge source kind
`BingCustomSearchSource`. For most "search this website" needs, use `PublicSiteSearchSource` instead.
See [07 Knowledge](07-knowledge.md).

### Bot Framework
Microsoft's older SDK and protocol for chatbots. Copilot Studio runs on it, which is where
[activities](#activity), [Direct Line](#direct-line), [hero cards](#hero-card) and the many `Bot…` names in the
schema come from.

### Bot Framework skill
A bot built with the [Bot Framework](#bot-framework) that another bot can call as a sub-component, through a
manifest that lists its actions. Copilot Studio can call one as a tool (`InvokeSkillTaskAction`) or from a topic
(`InvokeSkillAction`). Don't confuse it with the "skills" in Microsoft's `skills-for-copilot-studio` repo, which are
instructions for AI coding assistants. See [08 Tools & connectors](08-tools-and-connectors.md).

## C

### Channel
A place where users talk to the agent: Microsoft [Teams](#teams), Microsoft 365 Copilot, a website (web chat),
[Direct Line](#direct-line) clients, telephony, and so on. Channels render things differently (Markdown, cards,
line breaks), so YAML sometimes branches on `System.Activity.ChannelId`. See the
[channel-aware behaviour pattern](13-patterns.md#channel-aware-behavior).
Learn more: [Publish and deploy to channels](https://learn.microsoft.com/en-us/microsoft-copilot-studio/publication-fundamentals-publish-channels)

### Child agent
A smaller, specialised agent defined **inside** the parent agent's project (`agents/<name>/agent.mcs.yml`,
`kind: AgentDialog`). The parent's [orchestrator](#orchestrator) hands it work through the `OnToolSelected` trigger.
Compare [connected agent](#connected-agent). See [09 MCP & multi-agent](09-mcp-and-multi-agent.md).
Learn more: [Add other agents](https://learn.microsoft.com/en-us/microsoft-copilot-studio/authoring-add-other-agents)

### Citation
A reference marker (such as `[1]`) that a [generative answer](#generative-answers) attaches to show which
knowledge-source document a sentence came from. The [remove-citations template](07-knowledge.md) strips these with
`OnGeneratedResponse`.

### Classic orchestration
The original routing mode. Each user message is matched against every topic's [trigger phrases](#trigger-phrases),
and the best-matching topic runs. It is used when `GenerativeActionsEnabled` is `false`. Compare
[generative orchestration](#generative-orchestration). See [01 Concepts](01-concepts.md#how-a-message-is-routed).

### Cloud flow
A [Power Automate](#power-automate) workflow: a trigger followed by steps (call APIs, transform data, send
approvals) that runs in the cloud. Agents call flows to do deterministic, multi-step work. The schema represents a
flow as `CloudFlowDefinition` and calls it with `InvokeFlowAction` (from a topic) or `InvokeFlowTaskAction`
(as a tool). See [10 Flows & autonomous triggers](10-flows-and-autonomous-triggers.md).
Learn more: [Cloud flows overview](https://learn.microsoft.com/en-us/power-automate/overview-cloud) ·
[Call a flow from an agent](https://learn.microsoft.com/en-us/microsoft-copilot-studio/advanced-flow)

### Component
One unit of agent content stored in [Dataverse](#dataverse): a topic, a tool, a knowledge source, a variable. Each
has a globally unique **schema name** such as `cr26e_myAgent.topic.Greeting`, built from the agent's schema prefix,
the component type and a name. Topics reference each other by schema name (for example `dialog:` in
`BeginDialog`). See [02 Project layout](02-project-layout.md#schema-names-and-ids).

### Computer-using agent
A tool that operates a website or desktop application by looking at the screen and clicking and typing, the way a
person would. Use it for systems that have no API. In YAML it is a `TaskDialog` with
`InvokeComputerUsingAgentTaskAction`, which can carry an allow-list of websites and desktop apps. The skills repo
gives no guidance on it. See [08 Tools & connectors](08-tools-and-connectors.md).

### Connected agent
A separate, independently published Copilot Studio agent that your agent can call as a tool
(`InvokeConnectedAgentTaskAction`). Unlike a [child agent](#child-agent) it has its own project, settings and
lifecycle. See [09 MCP & multi-agent](09-mcp-and-multi-agent.md).

### Connection
A stored, authenticated sign-in to an external service (for example "a particular user's Outlook account"). Connections live
in the [environment](#environment) and are never written in YAML. YAML points at them indirectly through a
[connection reference](#connection-reference).

### Connection reference
A named pointer from your agent to a [connection](#connection), registered in `connectionreferences.mcs.yml`. Tool
YAML sets `connectionReference: <logical name>` instead of embedding credentials, so the same agent can be moved
between environments and re-pointed at different connections. Connection references can only be **created** by
adding the tool in the Copilot Studio UI. See [08 Tools & connectors](08-tools-and-connectors.md).
Learn more: [Connection references](https://learn.microsoft.com/en-us/power-apps/maker/data-platform/create-connection-reference)

### Connector
A Power Platform wrapper around an external API (Outlook, SharePoint, Teams, Azure DevOps, Salesforce, and over a
thousand more) that exposes named **operations**, each with an `operationId`, typed inputs and outputs. Agents call
connector operations as [tools](#tool). The skills repo bundles full operation lists for four connectors; see
[connector references](reference/connectors/index.md).
Learn more: [Connectors list](https://learn.microsoft.com/en-us/connectors/connectors)

### Connector trigger
An event exposed by a [connector](#connector), such as "When a new email arrives" (Office 365 Outlook
`OnNewEmailV3`) or "When an item is created" (SharePoint `GetOnNewItems`). [Cloud flows](#cloud-flow) start on
these events, and an [autonomous agent](#autonomous-agent) can be started by one through a
`ConnectorTriggerDefinition`. The bundled connector catalogues list each connector's triggers. See
[10 Flows & autonomous triggers](10-flows-and-autonomous-triggers.md).

### Content moderation
A sensitivity setting for generative AI output (`ContentModerationSettings`, levels `Minimum`, `Low`, `Medium`, `High`, `Maximum`). Stricter
levels block more borderline content, and blocked answers surface as [RAI](#rai) errors.
See [03 The agent](03-agent.md) and the [RAI error-handling pattern](13-patterns.md#rai-error-handling).

### Conversation starter
A suggested prompt (title plus text) that a [channel](#channel) shows before the user has typed anything. Set in
`agent.mcs.yml` under `conversationStarters:`. See [03 The agent](03-agent.md).

### Copilot connector
See [Graph connector](#graph-connector).

### Copilot Studio
Microsoft's low-code platform for building AI [agents](#agent), part of the [Power Platform](#power-platform). You
design agents in a web UI at copilotstudio.microsoft.com, or, as these docs describe, as YAML files that you pull,
edit and push.
Learn more: [What is Copilot Studio?](https://learn.microsoft.com/en-us/microsoft-copilot-studio/fundamentals-what-is-copilot-studio)

### Copilot Studio Kit
A free Microsoft "Power CAT" add-on, installed into an environment, that runs batch test suites against agents and
stores results in Dataverse. The skills repo's `run-tests-kit` skill drives it.
See [12 Evaluations & testing](12-evaluations-and-testing.md).
Learn more: [Power CAT Copilot Studio Kit (GitHub)](https://github.com/microsoft/Power-CAT-Copilot-Studio-Kit)

### CSAT
**Customer satisfaction** score. The `CSATQuestion` node asks the user to rate the conversation (1–5), and the
results feed Copilot Studio analytics.
Learn more: [Customer satisfaction analytics](https://learn.microsoft.com/en-us/microsoft-copilot-studio/analytics-csat)

### Custom connector
A [connector](#connector) you define yourself, from an OpenAPI description of your own REST API, so agents can call
it like any built-in connector.
Learn more: [Custom connectors](https://learn.microsoft.com/en-us/connectors/custom-connectors/)

## D

### Dataverse
The Power Platform's built-in database. Agents, topics, tools and every other [component](#component) are rows in
Dataverse tables, and the `.mcs.yml` files are a text projection of those rows. Dataverse tables can also be a
[knowledge source](#knowledge-source) (`DataverseStructuredSearchSource`).
Learn more: [What is Dataverse?](https://learn.microsoft.com/en-us/power-apps/maker/data-platform/data-platform-intro)

### Direct Line
A REST/WebSocket API for talking to a published agent from your own code. Test tools use it to send messages and
read replies for agents that don't use integrated auth. See [12 Evaluations & testing](12-evaluations-and-testing.md).
Learn more: [Direct Line API concepts](https://learn.microsoft.com/en-us/azure/bot-service/rest-api/bot-framework-rest-direct-line-3-0-concepts)

### Disambiguation
Asking "did you mean A or B?" when a message matches several topics equally well. Implemented with the
`OnSelectIntent` trigger; see the disambiguation template in [04 Topics & triggers](04-topics-and-triggers.md).

### DLP
**Data loss prevention** policies. Admin rules in Power Platform that decide which [connectors](#connector) may be
used together, or at all, in an [environment](#environment). A tool can fail at runtime because of a DLP policy even
when its YAML is valid.
Learn more: [Data policies](https://learn.microsoft.com/en-us/power-platform/admin/wp-data-loss-prevention)

### DTMF
**Dual-tone multi-frequency**: telephone keypad tones. Voice agents ([IVR](#ivr)) can accept keypad input, so the
schema includes DTMF triggers (`OnDtmfKeyPress`, `OnUnknownDtmfKey`) and settings.

## E

### Entity
A type of value that the agent knows how to recognise and extract from what the user says: a number, a date, an
email address, a choice from a list you define. A `Question` node's `entity:` decides what counts as a valid
answer. Entities can be **prebuilt** (`NumberPrebuiltEntity`, `EmailPrebuiltEntity`, …), **closed lists** (fixed
options with synonyms), or **regex**. See [06 Variables, entities & Power Fx](06-variables-entities-powerfx.md#entities).
Learn more: [Entities and slot filling](https://learn.microsoft.com/en-us/microsoft-copilot-studio/advanced-entities-slot-filling)

### Entra ID
Microsoft's identity service (formerly *Azure Active Directory*). It signs users in, issues tokens, and holds **app
registrations** (the identities that test tools use to call an agent).
Learn more: [What is Microsoft Entra ID?](https://learn.microsoft.com/en-us/entra/fundamentals/whatis) ·
[Register an app](https://learn.microsoft.com/en-us/entra/identity-platform/quickstart-register-app)

### Environment
An isolated Power Platform container (roughly a "workspace" or "tenant partition") with its own
[Dataverse](#dataverse) database, agents, flows, [connections](#connection) and security. Organisations usually have
separate dev, test and prod environments. A cloned agent's `.mcs/conn.json` records which environment it came from.
Learn more: [Environments overview](https://learn.microsoft.com/en-us/power-platform/admin/environments-overview)

### Environment variable
A named configuration value stored in a [solution](#solution), such as a URL or a site address that differs between
dev and prod. Schema kinds: `EnvironmentVariableDefinition`, `EnvironmentVariableValue`.
Learn more: [Environment variables](https://learn.microsoft.com/en-us/power-apps/maker/data-platform/environmentvariables)

### Evaluation
An automated quality check. You run a **test set** (questions plus expected answers or behaviours) against the
agent, and **graders** score each response for relevance, groundedness, exact match, expected tool use, and so on.
See [12 Evaluations & testing](12-evaluations-and-testing.md).
Learn more: [Agent evaluation](https://learn.microsoft.com/en-us/microsoft-copilot-studio/analytics-agent-evaluation-intro)

### External agent
An agent hosted outside Copilot Studio that your agent calls over a protocol such as [A2A](#a2a). MCP servers are
modelled with the same action kind (`InvokeExternalAgentTaskAction`). See [09 MCP & multi-agent](09-mcp-and-multi-agent.md).

## F

### Fabric
Microsoft Fabric, an analytics and data platform. A Fabric **data agent** can be a knowledge source
(`FabricAISkillSource`).
Learn more: [What is Microsoft Fabric?](https://learn.microsoft.com/en-us/fabric/fundamentals/microsoft-fabric-overview)

### Fallback
What happens when no topic matches the user's message: the `OnUnknownIntent` trigger fires (the **Fallback**
[system topic](#system-topic)). It typically tries [generative answers](#generative-answers), then rephrasing, then
escalation. `System.FallbackCount` counts consecutive fallbacks.
Learn more: [Fallback topic](https://learn.microsoft.com/en-us/microsoft-copilot-studio/authoring-system-fallback-topic)

## G

### Generative answers
LLM-written answers [grounded](#grounding) in your [knowledge sources](#knowledge-source). They happen
automatically under [generative orchestration](#generative-orchestration), or explicitly through the
`SearchAndSummarizeContent` node. See [07 Knowledge](07-knowledge.md).
Learn more: [Generative answers](https://learn.microsoft.com/en-us/microsoft-copilot-studio/nlu-boost-conversations)

### Generative orchestration
The AI-driven routing mode, switched on by `GenerativeActionsEnabled: true` in settings. An LLM
[orchestrator](#orchestrator) reads the agent's [instructions](#instructions) and every topic's and tool's
[model description](#model-description). It then plans which topics, tools, knowledge and child agents to use
(possibly several in one turn) and fills their inputs from the conversation. Compare
[classic orchestration](#classic-orchestration). See [01 Concepts](01-concepts.md#how-a-message-is-routed).
Learn more: [Generative orchestration](https://learn.microsoft.com/en-us/microsoft-copilot-studio/advanced-generative-actions)

### Global variable
A variable that keeps its value for the whole conversation and is visible to every topic: `Global.<Name>`. Each one
is declared in its own file under `variables/` (`kind: GlobalVariableComponent`). Compare
[topic variable](#topic-variable) and [system variable](#system-variable). See [06 Variables](06-variables-entities-powerfx.md).
Learn more: [Work with variables](https://learn.microsoft.com/en-us/microsoft-copilot-studio/authoring-variables)

### Graph connector
Also called a **Copilot connector**. It indexes content from a third-party system (ServiceNow, Confluence, a file
share) into Microsoft 365, so that Microsoft 365 search and Copilot can find it. Agents can use a Graph connector as a
knowledge source (`GraphConnectorSearchSource`).
Learn more: [Connectors overview](https://learn.microsoft.com/en-us/microsoftsearch/connectors-overview)

### Grounding
Making an LLM answer from specific retrieved documents (your [knowledge sources](#knowledge-source)) rather than from
its general training data. Grounded answers carry [citations](#citation) and are less prone to hallucination.

## H

### Hero card
A simple [Bot Framework](#bot-framework) card: title, subtitle, text, one image and a row of buttons. It's less
flexible than an [Adaptive Card](#adaptive-card) but widely supported. YAML: `HeroCardTemplate`.
Learn more: [Card types reference](https://learn.microsoft.com/en-us/microsoftteams/platform/task-modules-and-cards/cards/cards-reference)

## I

### Instructions
The agent's system prompt: free-text guidance (in `agent.mcs.yml` → `instructions:`) that tells the
[orchestrator](#orchestrator) its role, tone, rules and when to use which tools. See [03 The agent](03-agent.md).
Learn more: [Write agent instructions](https://learn.microsoft.com/en-us/microsoft-copilot-studio/authoring-instructions)

### Intent
What the user is trying to do ("reset my password"). Topics declare the intent they handle with `intent:` on an
`OnRecognizedIntent` trigger, using [trigger phrases](#trigger-phrases) and/or a [model description](#model-description).

### IVR
**Interactive voice response.** Phone-based agents that speak and listen, and can accept [DTMF](#dtmf) keypad input.
Many schema kinds (`OnSilence`, `OnUnrecognizedSpeech`, voice settings, latency settings) exist only for this
scenario.
Learn more: [Voice-enabled agents](https://learn.microsoft.com/en-us/microsoft-copilot-studio/voice-overview)

## K

### Knowledge source
A body of content the agent can search to [ground](#grounding) its answers: a public website, a SharePoint site,
uploaded files, Dataverse tables, Azure AI Search, a Graph connector, and others. Each source is one file under
`knowledge/` (`kind: KnowledgeSourceConfiguration`). See [07 Knowledge](07-knowledge.md).
Learn more: [Knowledge sources](https://learn.microsoft.com/en-us/microsoft-copilot-studio/knowledge-copilot-studio)

## L

### LCID
**Locale identifier.** A Windows number that names a language and region, for example `1033` for English (US) or
`2057` for English (UK). Copilot Studio stores an agent's primary language as an LCID in `settings.mcs.yml`
(`language: 1033`). The schema's `Language` definition lists the 32 values it accepts. See
[03 The agent](03-agent.md#settingsmcsyml).

### LSP
**Language Server Protocol**: a standard way for editors to get validation, completion and hover help from a
language-specific server. The [Copilot Studio VS Code extension](#vs-code-extension) ships a *LanguageServerHost*
binary that validates `.mcs.yml` files and performs clone/pull/push. The skills repo's scripts drive that same
binary.
Learn more: [Language Server Protocol](https://microsoft.github.io/language-server-protocol/)

## M

### Maker and invoker credentials
Who a [tool](#tool) acts as when it calls a [connector](#connector) (`connectionProperties.mode`). In **`Maker`**
mode it uses the agent author's [connection](#connection), so every user acts as the author. In **`Invoker`** mode it
uses the end user's own credentials, so each user can only see and do what they're allowed to. Invoker is the safer
default for personal data (mail, calendar, files). See [08 Tools & connectors](08-tools-and-connectors.md).

### MCP
**Model Context Protocol**, an open standard for exposing tools and data to LLM applications. An **MCP server**
publishes a list of tools with their parameters, and Copilot Studio can attach a whole MCP server as one tool
(`InvokeExternalAgentTaskAction` + `ModelContextProtocolMetadata`). The server's own tool list decides what the agent
can call. See [09 MCP & multi-agent](09-mcp-and-multi-agent.md).
Learn more: [modelcontextprotocol.io](https://modelcontextprotocol.io/) ·
[MCP in Copilot Studio](https://learn.microsoft.com/en-us/microsoft-copilot-studio/agent-extend-action-mcp)

### Microsoft 365 Copilot
Microsoft's AI assistant built into Office apps and Teams. Copilot Studio agents can be published into it as a
[channel](#channel), where it takes over some of the work your agent would do in other channels (for example
rendering and orchestration).
Learn more: [Microsoft 365 Copilot overview](https://learn.microsoft.com/en-us/copilot/microsoft-365/microsoft-365-copilot-overview)

### Microsoft Graph
The single REST API for Microsoft 365 data: users, mail, calendars, files, Teams. Connectors such as Office 365
Users call it behind the scenes.
Learn more: [Microsoft Graph overview](https://learn.microsoft.com/en-us/graph/overview)

### Model description
`modelDescription` (with its companion `modelDisplayName`): the text that the [generative orchestrator](#generative-orchestration)
reads to decide **when** to use a topic, tool or child agent, and what it does. It's the most important routing
field under generative orchestration, so write it like documentation for a colleague: what it does, when to use
it, and when *not* to. See [04 Topics & triggers](04-topics-and-triggers.md).

## N

### Node
One step in a topic's flow chart, for example "send a message", "ask a question", "branch on a condition" or "call a
tool". In YAML each node is one item in an `actions:` list, with a `kind:` and a unique `id:`. The schema calls
these `DialogAction`s. See [05 Nodes](05-nodes.md).

## O

### OData
**Open Data Protocol**, a REST convention for querying data with URL options such as `$filter`, `$orderby`, `$top`
and `$select`. SharePoint and Dataverse connector operations accept these, and they need special quoting in tool
YAML. See [08 Tools & connectors](08-tools-and-connectors.md#-prefixed-inputs-odata-parameters).
Learn more: [OData overview](https://learn.microsoft.com/en-us/odata/overview)

### Orchestrator
The part of Copilot Studio that decides what happens with each user message: which topic or tool to run, whether to
search knowledge, how to fill inputs, and how to phrase the reply. Its behaviour depends on whether
[generative](#generative-orchestration) or [classic](#classic-orchestration) orchestration is enabled.
See [01 Concepts](01-concepts.md).

## P

### Power Automate
The Power Platform's workflow automation service. Its workflows are called [cloud flows](#cloud-flow).
Learn more: [Cloud flows overview](https://learn.microsoft.com/en-us/power-automate/overview-cloud)

### Power Fx
The Excel-like formula language used across the Power Platform. In Copilot Studio YAML, any value that starts with
`=` is a Power Fx expression, for example `=Topic.Count + 1`, `=IsBlank(Global.UserName)` or
`=Text(Now(), DateTimeFormat.UTC)`. Copilot Studio supports only a **subset** of functions.
See [06 Variables, entities & Power Fx](06-variables-entities-powerfx.md#power-fx).
Learn more: [Power Fx overview](https://learn.microsoft.com/en-us/power-platform/power-fx/overview) ·
[Power Fx in Copilot Studio](https://learn.microsoft.com/en-us/microsoft-copilot-studio/advanced-power-fx)

### Power Platform
Microsoft's family of low-code products (Power Apps, [Power Automate](#power-automate), Power BI, Copilot Studio)
that share [environments](#environment), [Dataverse](#dataverse), [connectors](#connector) and
[solutions](#solution).

### Prebuilt entity
An [entity](#entity) that Copilot Studio already knows how to recognise, such as numbers, dates, emails, money,
cities or ages. Referenced by name, for example `entity: NumberPrebuiltEntity`.

### Publish
Copilot Studio keeps a **draft** and a **published** version of every agent. Pushing YAML or saving in the UI
updates the draft, and only **publishing** makes changes live in the [channels](#channel) and testable through
[Direct Line](#direct-line). See [02 Project layout](02-project-layout.md#lifecycle-clone--edit--validate--push--publish).

## Q

### Quick reply
Clickable suggestion buttons shown under a message (`quickReplies:` on a message, kinds `MessageBack`, `OpenUrl`,
`MakeCall`). Clicking one sends its text as if the user had typed it. Also called *suggested actions*.
See [11 Messages & Adaptive Cards](11-messages-and-adaptive-cards.md).

## R

### RAI
**Responsible AI.** Microsoft's safety filters on generative output and input (harmful content, jailbreak attempts,
copyright). When a filter blocks something, the agent raises an error (for example `ContentFiltered`) that you can
handle in an `OnError` topic. See the [RAI error-handling pattern](13-patterns.md#rai-error-handling).
Learn more: [Responsible AI FAQ for Copilot Studio](https://learn.microsoft.com/en-us/microsoft-copilot-studio/responsible-ai-overview)

### Redirect
Jumping from one topic to another. `BeginDialog` calls a topic and comes back afterwards, `ReplaceDialog` hands over
and doesn't come back, and the `OnSystemRedirect` trigger marks a topic that can *only* be reached this way.
See [05 Nodes](05-nodes.md).

## S

### SharePoint
Microsoft's intranet and document-management platform (sites, document libraries, lists). It's a common
[knowledge source](#knowledge-source) (`SharePointSearchSource`) and a common [connector](#connector)
(`shared_sharepointonline`).
Learn more: [SharePoint introduction](https://learn.microsoft.com/en-us/sharepoint/introduction)

### Slot filling
Automatically filling a question's answer (a "slot") from something the user already said. If the user says
"book 4 seats for Friday", the date and party-size questions are skipped. Driven by [entities](#entity) and, under
generative orchestration, by topic `inputs:`.

### Solution
The Power Platform's packaging unit, used to move [components](#component) (agents, flows, connection references,
[environment variables](#environment-variable)) between environments. Solutions are *unmanaged* (editable, in dev) or
*managed* (locked, in test/prod).
Learn more: [Solution concepts](https://learn.microsoft.com/en-us/power-platform/alm/solution-concepts-alm)

### System topic
A topic that every agent has and that fires on system events rather than user intent: Conversation Start, Fallback,
On Error, Escalate, Multiple Topics Matched, Sign In, End of Conversation, and others. You can edit them but not
delete them. See [04 Topics & triggers](04-topics-and-triggers.md).
Learn more: [System topics](https://learn.microsoft.com/en-us/microsoft-copilot-studio/authoring-system-topics)

### System variable
A read-only, built-in variable that the runtime provides, under `System.*`: for example `System.Activity.Text` (the
user's message), `System.Conversation.Id`, `System.User.Email`, `System.Error.Message`.
See [06 Variables](06-variables-entities-powerfx.md#system-variables).

## T

### Teams
Microsoft Teams, the most common [channel](#channel) for internal agents. It has its own rendering rules (limited Markdown,
its own Adaptive Card support) and production quirks; see the
[Teams hardening pattern](13-patterns.md#teams-production-hardening).
Learn more: [Publish to Teams](https://learn.microsoft.com/en-us/microsoft-copilot-studio/publication-add-bot-to-microsoft-teams)

### Telemetry
Usage and diagnostic data about conversations. Copilot Studio records analytics itself, and you can also send custom
events to [Application Insights](#application-insights) with `LogCustomTelemetryEvent`.

### Tenant
An organisation's Microsoft cloud account (one Entra ID directory), identified by a GUID or a domain like
`contoso.onmicrosoft.com`. Environments live inside a tenant.

### Test pane
The chat panel built into the Copilot Studio web UI that talks to the agent's **draft** (pushed, not necessarily
published) content. Inside it `System.Conversation.InTestMode` is `true`, and `System.Activity.ChannelId` has its own
value, so it is not a substitute for testing in a real [channel](#channel).
See [12 Evaluations & testing](12-evaluations-and-testing.md#level-2-the-test-pane).

### Tool
Anything the [orchestrator](#orchestrator) can call to get something done: a [connector](#connector) operation, an
[MCP](#mcp) server, a [flow](#cloud-flow), an AI Builder prompt, a REST API, or another agent. Tools are files in
`actions/` with `kind: TaskDialog`, whose `action:` block says what is called. Older docs call them "actions" or
"plugins". See [08 Tools & connectors](08-tools-and-connectors.md).
Learn more: [Add tools (plugins)](https://learn.microsoft.com/en-us/microsoft-copilot-studio/advanced-plugin-actions)

### Topic
A reusable conversation script, roughly a flow chart: a [trigger](#trigger) that says when it runs, plus a list of
[nodes](#node) that say what happens. Stored in `topics/*.topic.mcs.yml` as `kind: AdaptiveDialog`. Under
[generative orchestration](#generative-orchestration) a topic can also declare inputs and outputs and be used like a
tool. See [04 Topics & triggers](04-topics-and-triggers.md).
Learn more: [Create and edit topics](https://learn.microsoft.com/en-us/microsoft-copilot-studio/authoring-create-edit-topics)

### Topic variable
A variable that exists only while its topic runs: `Topic.<Name>`. You create it on first assignment by prefixing the
target with `init:` (for example `variable: init:Topic.Answer`). See [06 Variables](06-variables-entities-powerfx.md).

### Trigger
What starts a topic: a recognised intent, the conversation starting, an error, an unknown message, an incoming
event, a timer, and so on. Each topic has one trigger in `beginDialog:`, with a `kind:` such as `OnRecognizedIntent`
or `OnConversationStart`. See [04 Topics & triggers](04-topics-and-triggers.md).
Learn more: [Topic triggers](https://learn.microsoft.com/en-us/microsoft-copilot-studio/authoring-triggers)

### Trigger phrases
Example user sentences (`triggerQueries:`) that a topic should respond to. They're the main routing signal under
[classic orchestration](#classic-orchestration) and only a hint under generative orchestration, where the
[model description](#model-description) matters more.

## V

### VS Code extension
The **Copilot Studio extension for Visual Studio Code**. It clones an agent from an environment into local
`.mcs.yml` files, validates them through its [language server](#lsp), and pushes and pulls changes. The skills repo
needs it installed (it reuses the extension's binary), but VS Code itself doesn't have to be running.
See [02 Project layout](02-project-layout.md).
Learn more: [VS Code extension overview](https://learn.microsoft.com/en-us/microsoft-copilot-studio/visual-studio-code-extension-overview) ·
[GitHub repo](https://github.com/microsoft/vscode-copilotstudio)

## Y

### YAML
A human-friendly data format, indentation-based like Python. Copilot Studio's code view and the
[VS Code extension](#vs-code-extension) represent every component as YAML in files ending `.mcs.yml`.
Watch out for YAML's quoting rules: values containing `: `, `#`, `{` or a leading `=` or `'` usually need quotes.
Learn more: [YAML 1.2 specification](https://yaml.org/spec/1.2.2/)
