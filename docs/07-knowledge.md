# 07 — Knowledge

This page covers how an agent answers questions from documents and websites: the
[knowledge source](glossary.md#knowledge-source) (a searchable body of content the agent can draw on) files in
`knowledge/`, the 15 kinds of source the schema knows, how to control when each one is searched, and the topic
nodes that run a search explicitly. It ends with custom search topics and [citations](glossary.md#citation)
(the `[1]`-style markers that link an answer to its source).

## How knowledge is used

An answer drawn from knowledge is called a [generative answer](glossary.md#generative-answers) (an LLM-written
answer [grounded](glossary.md#grounding) in retrieved text, meaning it is based on that text rather than on the
model's general knowledge). [Copilot Studio](glossary.md#copilot-studio) (Microsoft's agent-building platform) builds one in five steps:

1. **Split.** At index time, each knowledge source is split into overlapping text chunks.
2. **Embed.** The user's query and all chunks are turned into vectors.
3. **Rank.** Chunks are ranked by how similar they are in meaning to the query.
4. **Pass.** The top-ranked chunks are handed to the language model as context.
5. **Generate.** The model writes a response grounded in those chunks, with citations.

Answer quality depends on how relevant the chunks are, how well the documents are structured, and the agent's
`instructions` (see [03 The agent](03-agent.md)).

### Automatic vs explicit retrieval

| Mode | When it happens | You write |
|---|---|---|
| **Automatic** | Under [generative orchestration](glossary.md#generative-orchestration) (`GenerativeActionsEnabled: true`, where an LLM plans each turn), the [orchestrator](glossary.md#orchestrator) decides for itself when a message needs a knowledge search. | Only the `knowledge/*.knowledge.mcs.yml` files. No topic is needed for simple Q&A. |
| **Explicit** | A [topic](glossary.md#topic) runs a [SearchAndSummarizeContent](reference/schema/actions.md#searchandsummarizecontent) (or related) node. | A topic with a search node (see [Generative-answer nodes](#generative-answer-nodes)). |

For most agents, automatic retrieval is enough. Write an explicit search topic only when you need one of these:

- to search **a subset** of sources rather than all of them;
- to **control the flow** around the answer (follow-up questions, formatting, [Adaptive Cards](glossary.md#adaptive-card), which are small JSON-defined UI cards);
- to **post-process** the answer (extract content, combine it with other data);
- to search with an input **other than the user's last message**.

Under [classic orchestration](glossary.md#classic-orchestration) (topics matched by [trigger phrases](glossary.md#trigger-phrases), the example sentences listed on each topic), there's no
automatic search. Instead, knowledge is usually reached through a [fallback](glossary.md#fallback) topic (the topic that runs when nothing else matches) (`OnUnknownIntent`) that contains a
search node. See [Pattern 3](#pattern-3-fallback-search).

### The UniversalSearchTool and the 25-source limit

Every knowledge search goes through one built-in tool, the **UniversalSearchTool**. This is true whether the
orchestrator triggers it automatically or an `OnKnowledgeRequested` topic hooks into it. The tool:

- searches **all configured sources at once**, whatever their type;
- returns the best-matching results. How many results it includes is controlled by the agent's
  [content moderation](glossary.md#content-moderation) slider (Generative AI settings): a higher setting includes
  more results, a lower one is more selective;
- supports **up to 25 sources**. If the agent has 25 or fewer, all of them are searched on every call. If it has
  more, the orchestrator **picks up to 25** that best match the query, using each source's **name and
  description comments** (see [File format](#file-format)).

The vendor guidance also suggests keeping to **about 10 sources per agent**, because too many sources degrade
relevance ranking, and avoiding overlapping sources that cover the same content.

## File format

Each knowledge source is one file, `knowledge/<descriptive-name>.knowledge.mcs.yml`, with
`kind:` [KnowledgeSourceConfiguration](reference/schema/components.md#knowledgesourceconfiguration). It has a single
property, `source:`, whose `kind:` picks one of the 15 [KnowledgeSource](reference/schema/knowledge.md#knowledgesource)
shapes.

```yaml
# Name: HR leave policies
# HR leave policies and employee entitlements for all countries.
kind: KnowledgeSourceConfiguration
source:
  kind: SharePointSearchSource
  site: https://contoso.sharepoint.com/sites/HR/Shared%20Documents/Policies
```

The two comment lines matter:

| Line | Purpose |
|---|---|
| `# Name: <display name>` | The display name shown in the Copilot Studio UI. |
| `# <plain description>` | What the source covers. Write it as plain text, with no `Description:` prefix. When the agent has more than 25 sources, the orchestrator reads this text to decide whether to include the source. `# HR leave policies and employee entitlements` is better than `# HR docs`. |

Use descriptive, lowercase, hyphenated filenames (`hr-policies.knowledge.mcs.yml`, not `ks1.knowledge.mcs.yml`).

### Referring to a knowledge source

Search nodes refer to a source by its **component schema name**, which is the file name without `.mcs.yml`. The
pattern in [02 Project layout](02-project-layout.md#component-schema-names) is
`<agentSchema>.knowledge.<name>`, for example `cr26e_myAgent.knowledge.hrPolicies`. Some of the vendor skills show
`<agentSchema>.topic.<name>_<suffix>` instead, which is how some sources created in the UI are named. **Copy the name
from the actual file in `knowledge/`; never guess it.** The source must already exist (and be pushed) before a
topic can reference it.

## The 15 knowledge source kinds

Every kind accepts an optional `triggerCondition` (see [Controlling when a source is searched](#controlling-when-a-source-is-searched)).
The skills repo can only **create** two kinds from YAML: public websites and [SharePoint](glossary.md#sharepoint)
(Microsoft's intranet and document-library service). The vendor guidance says the following must be created in the
Copilot Studio UI, because they need other [Power Platform](glossary.md#power-platform) (Microsoft's low-code
platform) configuration: [Dataverse](glossary.md#dataverse) (the Power Platform database) tables, uploaded files,
[Azure AI Search](glossary.md#azure-ai-search) (Azure's managed search-index service) and SQL Server. You can still edit their YAML after cloning.

| Kind | What it searches | Create in YAML? |
|---|---|---|
| [PublicSiteSearchSource](reference/schema/knowledge.md#publicsitesearchsource) | A public website, through Bing search scoped to a URL | ✓ |
| [SharePointSearchSource](reference/schema/knowledge.md#sharepointsearchsource) | A SharePoint folder or site | ✓ |
| [GraphConnectorSearchSource](reference/schema/knowledge.md#graphconnectorsearchsource) | Content indexed into Microsoft 365 by a [Graph connector](glossary.md#graph-connector) (ServiceNow, Jira, Salesforce, …) | Partly: prerequisites need admin setup |
| [FileGroupKnowledgeSource](reference/schema/knowledge.md#filegroupknowledgesource) | Files uploaded to the agent | UI |
| [DataverseStructuredSearchSource](reference/schema/knowledge.md#dataversestructuredsearchsource) | Dataverse tables | UI |
| [FederatedStructuredSearchSource](reference/schema/knowledge.md#federatedstructuredsearchsource) | Structured data reached through a connection (inferred: SQL Server and similar) | UI |
| [AzureAISearchSource](reference/schema/knowledge.md#azureaisearchsource) | An Azure AI Search index | UI |
| [AzureOpenAIOnYourDataKnowledgeSource](reference/schema/knowledge.md#azureopenaionyourdataknowledgesource) | An [Azure OpenAI](glossary.md#azure-openai) (Microsoft-hosted OpenAI models) "On Your Data" deployment over Azure AI Search | No guidance |
| [BingCustomSearchSource](reference/schema/knowledge.md#bingcustomsearchsource) | A [Bing Custom Search](glossary.md#bing-custom-search) (a curated Bing web search) configuration | No guidance |
| [FabricAISkillSource](reference/schema/knowledge.md#fabricaiskillsource) | A [Microsoft Fabric](glossary.md#fabric) (Microsoft's analytics platform) data agent ("AI skill") | No guidance |
| [CustomKnowledgeSource](reference/schema/knowledge.md#customknowledgesource) | Results produced by a [Power Fx](glossary.md#power-fx) (the formula language) expression | No guidance |
| [TeamsMessageSearchSource](reference/schema/knowledge.md#teamsmessagesearchsource) | Messages in [Teams](glossary.md#teams) (Microsoft's chat and meetings app) | No guidance |
| [MeetingSearchSource](reference/schema/knowledge.md#meetingsearchsource) | Teams meetings | No guidance |
| [EmailSearchSource](reference/schema/knowledge.md#emailsearchsource) | A mailbox | No guidance |
| [CopilotConnectorKnowledgeSourceDemo](reference/schema/knowledge.md#copilotconnectorknowledgesourcedemo) | A [Copilot connector](glossary.md#copilot-connector) (another name for a Graph connector); the `Demo` suffix suggests a preview shape | No guidance |

"No guidance" means the vendor repo says nothing about the kind beyond its schema. The descriptions in that column
are inferred from the kind and property names. Check the generated reference for the exact shape.

### PublicSiteSearchSource

Searches a public website. It **uses Bing search to find relevant snippets** within the URL you give. It doesn't
crawl the site, and it can't return or summarise a whole page. Think of `site` as telling Bing "only return results
from under this path".

```yaml
# Name: Contoso API docs
# Public developer documentation for the Contoso REST API: endpoints, authentication and error codes.
kind: KnowledgeSourceConfiguration
source:
  kind: PublicSiteSearchSource
  site: https://docs.contoso.com/en-us/api
```

| Property | Type / values | Req. | Meaning |
|---|---|---|---|
| `site` | string or `=` expression | | The URL that scopes the search. |
| `includeSubPages` | boolean, default `true` | | No vendor guidance. Inferred: whether pages below `site` are included. |
| `triggerCondition` | `=` expression | | When the source is searched. See [below](#controlling-when-a-source-is-searched). |

**URL depth limit: at most 2 path levels beyond the domain.** A deeper URL is **silently ignored**. The vendor
guidance calls this a current platform limitation that may be relaxed later.

| URL | Levels | Result |
|---|---|---|
| `https://microsoft.com/products/surface` | 2 (`products/surface`) | works |
| `https://docs.example.com/en-us/azure` | 2 (`en-us/azure`) | works |
| `https://example.com/docs/api/v2` | 3 | ignored |
| `https://microsoft.com/en-us/microsoft-365/business` | 3 | ignored |

Other rules:

- The site must be public, with no login required.
- Avoid sites whose content is rendered by JavaScript or is otherwise dynamic.
- Subdomains count as separate sources. Add each one you need.
- Sites are re-crawled periodically. If a URL is redirected or removed, indexing breaks silently, so keep an eye on
  your source URLs.

### SharePointSearchSource

Searches documents in a SharePoint site or folder. This is the usual choice for internal company content.

```yaml
# Name: IT runbooks
# IT support runbooks: password resets, VPN setup, laptop provisioning.
kind: KnowledgeSourceConfiguration
source:
  kind: SharePointSearchSource
  site: https://contoso.sharepoint.com/sites/IT/Shared%20Documents/Runbooks
```

| Property | Type / values | Req. | Meaning |
|---|---|---|---|
| `site` | string or `=` expression | | Direct folder path (see below). |
| `triggerCondition` | `=` expression | | When the source is searched. |
| `indexBehavior` | `Disabled` \| `EnabledWithFallback` (default) \| `EnabledWithoutFallback` | | No vendor guidance. Inferred: whether a pre-built index is used, and whether to fall back to live search. |
| `cascadeShare` | boolean, default `true` | | No vendor guidance. |
| `additionalSearchTerms` | string | | No vendor guidance. Inferred: extra terms added to each query. |

**The URL must be a direct folder path**, such as
`https://contoso.sharepoint.com/sites/MySite/Shared%20Documents/MyFolder`. Users often paste other forms:

| What you have | What to do |
|---|---|
| Direct path (`/sites/<Site>/Shared%20Documents/<Folder>`) | Use it as-is. |
| `…/AllItems.aspx?id=%2Fsites%2F…&viewid=…` | URL-decode the `id` query parameter, put the origin (`https://contoso.sharepoint.com`) in front of it, and drop all query parameters. |
| Sharing link (`/:f:/s/…`, `/:x:/s/…`) | You can't convert it, because the token hides the path. Open the folder in a browser and copy the address bar instead. |

Other rules:

- **Encode spaces as `%20`** in a literal `site:` value.
- Use the **deepest folder** that covers the documents you need, rather than the site root.
- Supported document types: PDF, Word (`.docx`), PowerPoint (`.pptx`) and plain text.
- **Permissions are enforced per user.** Search runs with the end user's delegated permissions. A user who can't
  open a document in SharePoint won't get answers from it. If the agent says "I don't have information about that",
  check access first.
- New files in the folder are picked up automatically. Renaming or moving a file breaks existing citations.
- **Only chunks come back, never whole files.** For scenarios that need a file's full content (a glossary CSV, a
  configuration file), use Dataverse or an [agent flow](glossary.md#agent-flow) (a deterministic automation called
  from the agent) with the SharePoint connector's *Get file content* operation. See
  [10 Flows](10-flows-and-autonomous-triggers.md).
- **Document libraries only, not SharePoint Lists.** To ground answers on list rows, use a
  [custom search topic](#custom-knowledge-via-onknowledgerequested-yaml-only).

### GraphConnectorSearchSource

Searches enterprise systems (ServiceNow, Salesforce, Jira, Azure DevOps, custom systems) whose content a Graph
connector has indexed into Microsoft 365.

Two prerequisites **can't be done in YAML**:

1. A Microsoft 365 admin must register and enable the Graph connector (Microsoft 365 admin center → Search &
   Intelligence → Data Sources).
2. The connector's connection ID must be stored as a Power Platform
   [environment variable](glossary.md#environment-variable) (a named configuration value stored in the
   [solution](glossary.md#solution)). The YAML refers to that variable, not to a raw GUID.

```yaml
# Name: ServiceNow Knowledge Base
# ServiceNow tickets and knowledge articles indexed via Microsoft Graph connector.
kind: KnowledgeSourceConfiguration
source:
  kind: GraphConnectorSearchSource
  connectionId:
    schemaName: cr26e_GraphConnectorId_ServiceNow   # the environment variable's schema name
  connectionName: servicenow-connection             # logical name of the Graph connection
  contentSourceDisplayName: ServiceNow              # label shown on citations
  publisherName: ServiceNow
```

| Property | Type | Req. | Meaning |
|---|---|---|---|
| `connectionId.schemaName` | string | ✓ | Schema name of the environment variable that holds the connection ID. Copy it from the solution or the UI. |
| `connectionName` | string | ✓ | The Graph connection's logical name, as registered in the Microsoft 365 admin center. |
| `contentSourceDisplayName` | string | | Label shown on citations. |
| `publisherName` | string | | Shown in the Copilot Studio UI. |
| `triggerCondition` | `=` expression | | Supported. `=false` opts the source out of automatic search. |
| `contentSourceIcon`, `additionalSearchTerms`, `itemsByUrl`, `itemsByExternalId`, `itemsByPath` | | | No vendor guidance. Inferred: the last three restrict the search to specific items or containers. |

> The vendor guide mentions a template `templates/knowledge/graph-connector.knowledge.mcs.yml`, but that file
> doesn't exist in the vendored repo. Use the example above.

### Other kinds, briefly

For these kinds, create the source in the UI, pull it, and then edit the YAML only if you need to. Their required
properties, from the schema:

| Kind | Required properties | Notable optional properties |
|---|---|---|
| FileGroupKnowledgeSource | — | `instructions` |
| AzureAISearchSource | `indexName` | `connectionReference`, `semanticConfigurationName`, `searchFields`, `selectFields`, `nearestNeighbors` (default 3) |
| AzureOpenAIOnYourDataKnowledgeSource | `dataSources`, `apiVersion`, `deployment` | `temperature`, `topP`, `maxTokens` |
| BingCustomSearchSource | — | `customConfigurationId` |
| FabricAISkillSource | `workspaceId.schemaName`, `aISkillId.schemaName` (both are environment-variable references) | — |
| CustomKnowledgeSource | `searchResults` (`=` expression) | — |
| DataverseStructuredSearchSource | — | `skillConfiguration` |
| FederatedStructuredSearchSource | — | `connectionProperties`, `skillConfiguration` |
| TeamsMessageSearchSource | — | `url`, `isRecurringMeeting` |
| MeetingSearchSource | — | `url`, `id`, `isSeries` |
| EmailSearchSource | — | `mailbox`, `folders` |
| CopilotConnectorKnowledgeSourceDemo | `consent`, `connectionReference`, `graphConnectionIdPrefix` | `entityLevelConfiguration` |

## Dynamic URLs with global variables

`site:` accepts a Power Fx expression. That means one knowledge
source can point at different content per user. For example, a per-region SharePoint folder held in a
[global variable](glossary.md#global-variable) (a conversation-wide variable).

```yaml
# Name: Regional HR KB
# HR policies from the SharePoint folder for the user's region.
kind: KnowledgeSourceConfiguration
source:
  kind: SharePointSearchSource
  site: =$"{Global.UserKBURL}"
```

```yaml
# Name: Regional Docs
# Product documentation for the user's region.
kind: KnowledgeSourceConfiguration
source:
  kind: PublicSiteSearchSource
  site: =$"https://docs.example.com/{Global.Region}/api"
```

Checklist:

1. Create the global variable (see [06 Variables](06-variables-entities-powerfx.md)).
2. Set it **early**, before any knowledge search can run. A common place is a conversation-init topic that loads the
   user's profile (see [13 Patterns](13-patterns.md#jit-user-context)). If it's blank when the search runs, the URL
   is blank.
3. The variable must hold a **clean, direct URL**, never an `AllItems.aspx?id=…` link. Unlike a literal `site:`,
   spaces are fine here (`Shared Documents`), because the `$"…"` expression handles encoding at runtime.

## Controlling when a source is searched

Every source kind has an optional `triggerCondition`, a Power Fx boolean. The UniversalSearchTool only includes the
source when the condition is `true`.

```yaml
# Name: Armstrong County Knowledge Base
# Employee information specific to Armstrong county; only searched for users located there.
kind: KnowledgeSourceConfiguration
source:
  kind: SharePointSearchSource
  triggerCondition: =Global.UserCounty = "Armstrong"
  site: https://contoso.sharepoint.com/sites/KnowledgeBase/Shared%20Documents/By_County/Armstrong
```

```yaml
# Name: Confidential procedures
# Restricted procedures, only searched when a topic asks for them explicitly.
kind: KnowledgeSourceConfiguration
source:
  kind: SharePointSearchSource
  triggerCondition: =false
  site: https://contoso.sharepoint.com/sites/Confidential/Shared%20Documents
```

**`triggerCondition: =false`** is the common case. The orchestrator never searches the source automatically. Use it
when:

1. **A topic should control the search.** The source is used only when a search node names it in
   `knowledgeSources:`.
2. **A startup topic decides.** For example, `OnConversationStart` sets `Global.UserDepartment`, and other sources use
   that variable in their `triggerCondition`.
3. **An `OnKnowledgeRequested` topic does all retrieval.** Combined with `=false` on the sources, your topic can route
   every knowledge request through custom logic.

> **YAML-only feature: the UI may remove it.** The Copilot Studio UI shows `triggerCondition` only as an on/off
> toggle ("exclude from search", which writes `=false`). **Arbitrary expressions** such as
> `=Global.UserDepartment = "HR"` work at runtime but can only be set in YAML. They're invisible in the UI, and if
> someone edits that knowledge source in the UI the expression may be silently removed. Note it in your repo and
> re-check after UI edits. See [14 Gotchas](14-gotchas.md).

To keep knowledge out of a particular topic entirely (for example a purely transactional flow), don't put search
nodes in it.

## Generative-answer nodes

Four nodes work with knowledge or the model directly. They go in a topic's `actions:` list like any other
[node](glossary.md#node) (a step in a topic; see [05 Nodes](05-nodes.md)).

| Node | Searches | Output | Use for |
|---|---|---|---|
| [SearchAndSummarizeContent](reference/schema/actions.md#searchandsummarizecontent) | Agent knowledge (all sources, or a list) | AI-summarised answer with citations | Most "generative answers" cases |
| [SearchKnowledgeSources](reference/schema/actions.md#searchknowledgesources) | Agent knowledge | **Raw** search results, not summarised | Verbatim needs: policies, legal or HR text |
| [AnswerQuestionWithAI](reference/schema/actions.md#answerquestionwithai) | Nothing: uses conversation history and the model's general knowledge | AI-generated answer | Explicitly ungrounded answers (rarely needed) |
| [CreateSearchQuery](reference/schema/actions.md#createsearchquery) | — | A rewritten search query | Run it **before** the two search nodes |

### Always run CreateSearchQuery first

The raw last message (`System.Activity.Text`) often lacks context. If the user says "tell me more about that", a
search for that literal text finds nothing useful. `CreateSearchQuery` uses AI and the conversation history to
rewrite the message into a self-contained query. The vendor skills say to **always** put it before
`SearchAndSummarizeContent` or `SearchKnowledgeSources`. Read its result as `Topic.<ResultVar>.SearchQuery`.

```yaml
- kind: CreateSearchQuery
  id: createSearchQuery_Qm4t8R
  userInput: =System.Activity.Text      # the default if omitted
  result: Topic.SearchQuery             # use Topic.SearchQuery.SearchQuery afterwards
```

| Property | Type / values | Req. | Meaning |
|---|---|---|---|
| `userInput` | `=` expression, default `=System.Activity.Text` | | Raw input to rewrite. |
| `result` | variable name | | Where to store the result. The query string is in `.SearchQuery`. |
| `historyType` | `kind: ConversationHistory` (with `numberOfPastUserMessagesToInclude`, default 5) or `kind: NoHistory` | | How much conversation history to use. |
| `moderationLevel` | `Minimum` … `Maximum`, default `High` | | Content moderation level. |

### SearchAndSummarizeContent

The main generative-answers node. It searches knowledge, has the model write a grounded answer, and either sends the
answer or stores it.

```yaml
- kind: SearchAndSummarizeContent
  id: searchContent_Hv7kP2
  autoSend: false
  variable: Topic.Answer
  userInput: =Topic.SearchQuery.SearchQuery
  fileSearchDataSource:
    searchFilesMode:
      kind: DoNotSearchFiles
  knowledgeSources:
    kind: SearchSpecificKnowledgeSources
    knowledgeSources:
      - cr26e_myAgent.knowledge.hrPolicies
  applyModelKnowledgeSetting: false
  responseCaptureType: FullResponse
```

**Property reference** (descriptions from the vendor property reference; rows marked *inferred* have no vendor
guidance):

| Property | Type / values | Req. | Meaning |
|---|---|---|---|
| `id` | string | ✓ | Node ID, unique within the topic. |
| `userInput` | `=` expression | ✓ | What to answer. Usually `=Topic.SearchQuery.SearchQuery` (from `CreateSearchQuery`). Can also be a topic input or a crafted expression such as `="Give me information about " & Topic.Subject`. |
| `autoSend` | boolean, default `true` | | `true` sends the answer to the chat. `false` stores it in `variable` so you can send it yourself, put it in a card, or return it to the orchestrator. |
| `variable` | variable name | | Where the answer is stored. With `responseCaptureType: FullResponse`, read `.Text.Content`, `.Text.MarkdownContent` and `.Text.CitationSources`. Blank means nothing was found. |
| `responseCaptureType` | `TextOnly` (default) \| `FullResponse` | | `FullResponse` captures the structured response: content, markdown and citations. |
| `knowledgeSources` | `kind: SearchSpecificKnowledgeSources` with a `knowledgeSources:` list of schema names, or `kind: SearchAllKnowledgeSources` | | Restricts the search to the listed sources. Omit it to search all of the agent's knowledge. |
| `applyModelKnowledgeSetting` | boolean | | `true` also lets the model use its general knowledge. `false` restricts the answer to the configured sources. |
| `fileSearchDataSource.searchFilesMode` | `kind: DoNotSearchFiles` \| `SearchAllFiles` \| `SearchSpecificFiles` (`files:`) | | Controls whether uploaded files are searched. `DoNotSearchFiles` turns file search off. |
| `additionalInstructions` | string | | *Inferred:* extra instructions for the summarising model. |
| `historyType` | as for `CreateSearchQuery` | | Conversation history to include. |
| `moderationLevel` | `Minimum`…`Maximum`, default `High` | | Content moderation level. |
| `webBrowsing` / `webBrowsingOptions.market` | boolean / Bing market code | | Also search the web through Bing (the vendor documents this for `SearchKnowledgeSources`). |
| `searchEmails` | boolean, default `false` | | *Inferred:* also search the user's email. |
| `publicDataSource`, `sharePointSearchDataSource`, `azureOpenAIOnYourDataSource`, `customDataSource` | objects | | *Inferred:* ad-hoc sources defined inline on the node rather than as knowledge files. `customDataSource.searchResults` takes a `=` expression. |
| `isFileAnalysisEnabled` | boolean | | *Inferred:* allow analysis of files the user uploads. |
| `modality` | `Automatic` (default) \| `Text` \| `Speech` \| `TextAndSpeech` | | *Inferred:* output form for voice channels. |
| `latencyMessageSettings` | `allowLatencyMessage`, `latencyMessage`, `latencyMessageDelayInMilliseconds`, … | | *Inferred:* a "please wait" message if the search is slow. Compare the [hold-message pattern](#hold-messages-during-search). |
| `displayName`, `disabled` | | | Canvas label; skip the node. |

### SearchKnowledgeSources

Returns **raw results without summarisation**. Use it when summarising could lose critical detail (insurance
policies, HR documents, legal or compliance text), or when you want to process the results yourself.

| Property | Meaning |
|---|---|
| `userInput` ✓ | The search query. Same options as `SearchAndSummarizeContent.userInput`. |
| `result` | Variable that receives the raw results. |
| `knowledgeSources` | Same format as above. Omit it to search all sources. |
| `webBrowsing` | `true` also searches the web through Bing. |
| `fileSearchDataSource`, `historyType`, `moderationLevel` | As above. |

### AnswerQuestionWithAI

Answers from **conversation history and general model knowledge only**. It doesn't consult knowledge sources. The
vendor authoring rules say to use it only for deliberately ungrounded, general-knowledge answers. Often you don't
need it at all, because the orchestrator answers general questions on its own.

```yaml
- kind: AnswerQuestionWithAI
  id: answerAI_p8Lm2q
  userInput: =System.Activity.Text
  autoSend: false
  variable: Topic.GeneralAnswer
  additionalInstructions: Answer in two sentences or fewer.
```

Properties: `userInput` (required), `autoSend` (default `true`), `variable`, `historyType`, `moderationLevel`,
`additionalInstructions`, `isFileAnalysisEnabled`.

[SearchAndSummarizeWithCustomModel](reference/schema/actions.md#searchandsummarizewithcustommodel) also exists in the
schema: it is the same idea, with an external model under `modelConfiguration`. The vendor repo has no guidance on
it.

## Patterns

Pick one based on orchestration mode:

| Situation | Pattern |
|---|---|
| `GenerativeActionsEnabled: true` | **2: Orchestrator.** The topic returns the answer as an output and the orchestrator phrases the reply. |
| `GenerativeActionsEnabled: false`, need control | **1: Direct response** (`autoSend: false` and send it yourself) |
| `GenerativeActionsEnabled: false`, simple | **3: Fallback search** |
| Verbatim content required | **4: Precision search** |

### Pattern 1: Direct response

Store the answer, check that one was found, and send it yourself. Use this when you want custom formatting or a card.

```yaml
# three nodes from a topic's actions: list
- kind: CreateSearchQuery
  id: createSearchQuery_d1R3ct
  userInput: =System.Activity.Text
  result: Topic.SearchQuery

- kind: SearchAndSummarizeContent
  id: searchContent_d1R3ct
  autoSend: false
  variable: Topic.Answer
  userInput: =Topic.SearchQuery.SearchQuery
  fileSearchDataSource:
    searchFilesMode:
      kind: DoNotSearchFiles
  knowledgeSources:
    kind: SearchSpecificKnowledgeSources
    knowledgeSources:
      - cr26e_myAgent.knowledge.productDocs
  applyModelKnowledgeSetting: false
  responseCaptureType: FullResponse

- kind: ConditionGroup
  id: conditionGroup_d1R3ct
  conditions:
    - id: conditionItem_d1R3ct
      condition: =!IsBlank(Topic.Answer)
      actions:
        - kind: SendActivity
          id: sendActivity_d1R3ct
          activity: "{Topic.Answer.Text.MarkdownContent}"
```

### Pattern 2: Orchestrator

This is the best fit for generative orchestration. The topic takes a question as an input, which the orchestrator
fills in. The topic searches and returns the text as an **output**, and the orchestrator writes the final reply. You
don't need a `SendActivity` node.

```yaml
kind: AdaptiveDialog
modelDescription: Answers questions about Contoso's travel and expense policy.
inputs:
  - kind: AutomaticTaskInput
    propertyName: QuestionToBeAnswered
    description: The user's question about travel or expense policy
    shouldPromptUser: true

beginDialog:
  kind: OnRecognizedIntent
  id: main
  intent: {}
  actions:
    - kind: CreateSearchQuery
      id: createSearchQuery_Tr4v3l
      userInput: =Topic.QuestionToBeAnswered
      result: Topic.SearchQuery

    - kind: SearchAndSummarizeContent
      id: searchContent_Tr4v3l
      autoSend: false
      variable: Topic.SearchResult
      userInput: =Topic.SearchQuery.SearchQuery
      fileSearchDataSource:
        searchFilesMode:
          kind: DoNotSearchFiles
      knowledgeSources:
        kind: SearchSpecificKnowledgeSources
        knowledgeSources:
          - cr26e_myAgent.knowledge.travelPolicy
      applyModelKnowledgeSetting: false
      responseCaptureType: FullResponse

    - kind: SetVariable
      id: setVariable_Tr4v3l
      variable: Topic.Response
      value: =Topic.SearchResult.Text.Content

inputType:
  properties:
    QuestionToBeAnswered:
      displayName: QuestionToBeAnswered
      description: The user's question about travel or expense policy
      type: String

outputType:
  properties:
    Response:
      displayName: Response
      description: The answer to the provided question
      type: String
```

Inputs and outputs must be declared both in `inputs:` and in `inputType`/`outputType`. See
[04 Topics & triggers](04-topics-and-triggers.md).

### Pattern 3: Fallback search

The default "Search" topic: when no other topic matches, search **all** knowledge. `priority: -1` makes it run
before the standard fallback, so knowledge gets a chance to answer before the agent says "I don't understand". If an
answer was found (and auto-sent), `EndDialog` with `clearTopicQueue: true` ends the topic and clears any queued
topics.

```yaml
# Name: Knowledge Search
kind: AdaptiveDialog
beginDialog:
  kind: OnUnknownIntent
  id: main
  priority: -1
  actions:
    - kind: CreateSearchQuery
      id: createSearchQuery_Fb9xL2
      userInput: =System.Activity.Text
      result: Topic.SearchQuery

    - kind: SearchAndSummarizeContent
      id: searchContent_Fb9xL2
      variable: Topic.Answer
      userInput: =Topic.SearchQuery.SearchQuery

    - kind: ConditionGroup
      id: conditionGroup_Fb9xL2
      conditions:
        - id: conditionItem_Fb9xL2
          condition: =!IsBlank(Topic.Answer)
          actions:
            - kind: EndDialog
              id: endDialog_Fb9xL2
              clearTopicQueue: true
```

### Pattern 4: Precision search

Raw results with no summary. Note that `userInput` here is `=Topic.SearchQuery`, exactly as in the vendor pattern,
whereas the other patterns read `.SearchQuery` from the result. If you get odd results, try
`=Topic.SearchQuery.SearchQuery`.

```yaml
- kind: CreateSearchQuery
  id: createSearchQuery_Pr3c1s
  userInput: =System.Activity.Text
  result: Topic.SearchQuery

- kind: SearchKnowledgeSources
  id: searchKnowledge_Pr3c1s
  userInput: =Topic.SearchQuery
  result: Topic.RawResults
  knowledgeSources:
    kind: SearchSpecificKnowledgeSources
    knowledgeSources:
      - cr26e_myAgent.knowledge.insurancePolicies

- kind: ConditionGroup
  id: conditionGroup_Pr3c1s
  conditions:
    - id: conditionItem_Pr3c1s
      condition: =!IsBlank(Topic.RawResults)
      actions:
        - kind: SendActivity
          id: sendActivity_Pr3c1s
          activity: "{Topic.RawResults}"
```

## Custom knowledge via OnKnowledgeRequested (YAML-only)

A topic whose trigger is [OnKnowledgeRequested](reference/schema/triggers.md#onknowledgerequested) fires **every
time the orchestrator invokes the UniversalSearchTool**. It doesn't replace the search; it hooks into it. This
trigger has **no UI designer**: you can only author it in YAML.

Use it to:

- **rewrite the query** (override `System.SearchQuery`) before it reaches the sources;
- **route** to specific sources based on user context (with a search node using `SearchSpecificKnowledgeSources`),
  or gate searches behind a condition;
- **bring in knowledge the platform can't reach**: SharePoint Lists, external REST APIs, databases that aren't in
  Dataverse. Call them yourself and write the results to `System.SearchResults`. The orchestrator then treats them
  like native results, including grounding and citations.

### System variables available only in this trigger

These [system variables](glossary.md#system-variable) (read-only runtime values) exist only inside an
`OnKnowledgeRequested` topic:

| Variable | Type | Meaning |
|---|---|---|
| `System.SearchQuery` | String | The user's query, rewritten for **semantic/vector** search using conversation history and the agent's instructions. You can read it and overwrite it. |
| `System.KeywordSearchQuery` | String | The query rewritten for **keyword/lexical** search engines (for example Azure AI Search in keyword mode). |
| `System.SearchResults` | Table | Starts empty. **Write** your results here. |

See also [06 System variables](06-variables-entities-powerfx.md#system-variables).

### Result record shape

Each row of `System.SearchResults` needs these fields: **`Content`** (the text), **`ContentLocation`** (a URL, used
for the citation) and **`Title`**. This shape comes from the vendor template and the int-reference system-variable
table, and matches the schema's `CustomSearchInputContent` (`content`, `contentLocation`, `title`).

> **Source conflict.** The vendor knowledge guide's SharePoint List example uses `{snippet, title, url}` instead. It
> doesn't match the template or the schema, so prefer `Content`/`ContentLocation`/`Title`, and test.

### Example: search a SharePoint List

The built-in SharePoint source can't index list rows, so this topic calls the SharePoint
[connector](glossary.md#connector) (a pre-built API wrapper) directly. It uses an inline
[InvokeConnectorAction](reference/schema/actions.md#invokeconnectoraction) node (see
[08 Tools & connectors](08-tools-and-connectors.md#inline-connector-calls-in-topics)) and maps each row into the
result shape.

```yaml
kind: AdaptiveDialog
beginDialog:
  kind: OnKnowledgeRequested
  id: main
  actions:
    # 1. Query the Events list (connection reference copied from an existing SharePoint tool)
    - kind: InvokeConnectorAction
      id: queryEvents_Ev3nt5
      connectionReference: cr26e_myAgent.shared_sharepointonline.4f2a9c
      connectionProperties:
        mode: Invoker
      operationId: GetItems
      input:
        binding:
          parameters/dataset: https://contoso.sharepoint.com/sites/HR
          parameters/table: Events
          parameters/$filter: Title ne null
      output:
        kind: SingleVariableOutputBinding
        variable: init:Topic.EventItems

    # 2. Map list rows to search results
    - kind: SetVariable
      id: setResults_Ev3nt5
      variable: System.SearchResults
      value: "=ForAll(Topic.EventItems.value, { Content: ThisRecord.Description & \" | Date: \" & Text(ThisRecord.Date, \"yyyy-mm-dd\") & \" | Location: \" & ThisRecord.Location, ContentLocation: \"https://contoso.sharepoint.com/sites/HR/Lists/Events/DispForm.aspx?ID=\" & Text(ThisRecord.ID), Title: ThisRecord.Title })"
```

The vendor's `templates/topics/custom-knowledge-source.topic.mcs.yml` shows the same idea with a `BeginDialog` that
calls an existing connector tool and binds `query: =System.SearchQuery`. Either way, you need a connector
[connection reference](glossary.md#connection-reference) (an authenticated link to the service), which you can only
create in the UI (see [08](08-tools-and-connectors.md#connection-references)).

Runtime sequence: the user asks "What company events are coming up?" → the orchestrator detects a knowledge need →
the UniversalSearchTool runs → `OnKnowledgeRequested` fires → the topic fetches rows and fills
`System.SearchResults` → the orchestrator grounds its answer on those rows (plus any other sources), with citations
that link to the list items.

Things to know:

- Your results are **merged** with results from the normal sources, unless every source has
  `triggerCondition: =false`.
- Pass `System.SearchQuery` or `System.KeywordSearchQuery` to the external service so it can filter on the server.
- **It adds latency**, because it adds work to every knowledge search. Put a `condition:` on the trigger so it only
  runs when needed (for example only for HR users).
- Classifying the query first (an orchestrator-filled topic input) and then routing by category inside
  `OnKnowledgeRequested` is a proven pattern. See [13 Patterns](13-patterns.md#orchestrator-variables).

> The vendor guide lists "Use the `knowledgeSourceIds` filter in a `SearchAndSummarizeContent` node". No such
> property exists in the schema. The property is `knowledgeSources`.

### Hold messages during search

Because `OnKnowledgeRequested` fires before retrieval finishes, a `SendActivity` in it reaches the user straight
away. That makes it a good place to send "One moment while I look that up…" on channels without streaming. The full
pattern picks a random message from a `Table()`. See
[13 Patterns — Knowledge hold message](13-patterns.md#knowledge-hold-message).

```yaml
kind: AdaptiveDialog
beginDialog:
  kind: OnKnowledgeRequested
  id: main
  actions:
    - kind: SendActivity
      id: sendActivity_wK9mT3
      activity: One moment while I search for that information...
```

The trade-off is two messages per answer, which some stakeholders find chatty.

## Citations

Generative answers carry citation markers (`[1]`, `[2]`, …) that link sentences to source documents. With
`responseCaptureType: FullResponse` the sources are in `<variable>.Text.CitationSources`. Check them when you test:
citations should point at the documents you expect.

### Removing citation markers

To strip the markers from **every** AI-generated answer, add a topic with the
[OnGeneratedResponse](reference/schema/triggers.md#ongeneratedresponse) trigger. It runs just before a generated
response is sent. In it you can read `System.Response.FormattedText`. Setting `System.ContinueResponse` to `false`
stops the original from being sent, so you can send a cleaned copy instead. This is the vendor's
`remove-citations` template:

```yaml
# Name: Remove Citations
kind: AdaptiveDialog
beginDialog:
  kind: OnGeneratedResponse
  id: main
  actions:
    # Suppress the original AI response
    - kind: SetVariable
      id: setVariable_Rc1t01
      variable: System.ContinueResponse
      value: =false

    # Strip markers [1] to [5]
    - kind: SetVariable
      id: setVariable_Rc1t02
      variable: init:Topic.CleanedText
      value: "=Substitute(Substitute(Substitute(Substitute(Substitute(System.Response.FormattedText, \"[1]\", \"\"), \"[2]\", \"\"), \"[3]\", \"\"), \"[4]\", \"\"), \"[5]\", \"\")"

    # Send the cleaned text
    - kind: SendActivity
      id: sendActivity_Rc1t03
      activity: "{Topic.CleanedText}"
```

Limitations:

- Only `[1]` to `[5]` are removed. Add more `Substitute` layers if answers cite more sources.
- `OnGeneratedResponse` is also **YAML-only**: you can't create it in the UI.
- The same trigger is used to block leaked tool-call JSON. See
  [13 Patterns — Prevent tool call leaks](13-patterns.md#prevent-tool-call-leaks).

## Content-quality checklist

- Documents should have clear headings and titles, which are used for chunking and citations.
- Avoid sources that are mostly tables, images or charts with little text.
- Scanned PDFs need OCR text. Image-only PDFs aren't searchable.
- One subject per document gives better answers than one huge handbook.
- Avoid duplicate content across documents, because it confuses ranking.
- Very short documents (under a page) may lack context. Consider combining them.
- Test each source with several representative questions. One passing test isn't enough.
- Review sources quarterly. When a URL changes, edit the existing file rather than adding a duplicate source.

## Sources

- `vendor/skills-for-copilot-studio/skills/add-knowledge/SKILL.md`: source types, URL normalisation, dynamic URLs, YAML-creatable kinds
- `vendor/skills-for-copilot-studio/skills/add-knowledge/knowledge-guide.md`: retrieval pipeline, UniversalSearchTool, 25-source limit, Graph connectors, `triggerCondition`, `OnKnowledgeRequested`, best practices
- `vendor/skills-for-copilot-studio/skills/add-generative-answers/SKILL.md`, `property-reference.md`, `patterns.md`: generative-answer nodes and patterns
- `vendor/skills-for-copilot-studio/skills/int-reference/SKILL.md`: YAML-only features, `System.SearchResults` schema
- `vendor/skills-for-copilot-studio/agents/copilot-studio-author.md`: when to use knowledge vs `SearchAndSummarizeContent` vs `AnswerQuestionWithAI`
- `vendor/skills-for-copilot-studio/templates/knowledge/public-website.knowledge.mcs.yml`, `templates/knowledge/sharepoint.knowledge.mcs.yml`
- `vendor/skills-for-copilot-studio/templates/topics/search-topic.topic.mcs.yml`, `custom-knowledge-source.topic.mcs.yml`, `remove-citations.topic.mcs.yml`, `conversation-init.topic.mcs.yml`
- `vendor/skills-for-copilot-studio/patterns/knowledge-hold-message.md`, `patterns/orchestrator-variables.md`, `patterns/prevent-tool-call-leaks.md`
- Generated reference: [knowledge.md](reference/schema/knowledge.md), [actions.md](reference/schema/actions.md), [triggers.md](reference/schema/triggers.md)
