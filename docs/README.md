# Copilot Studio YAML authoring reference

These docs explain how to define a **Microsoft Copilot Studio agent**, including its topics, tools, connectors,
knowledge, flows, variables, cards and evaluations, as [YAML](glossary.md#yaml) files you can version, review and edit in any
editor. They assume **no prior knowledge** of Copilot Studio or the Microsoft
[Power Platform](glossary.md#power-platform). Each Microsoft-specific term gets a one-line explanation where it
first appears and links to the [glossary](glossary.md) for more.

> **Source of truth.** Everything here comes from Microsoft's
> [`skills-for-copilot-studio`](https://github.com/microsoft/skills-for-copilot-studio) repository, vendored at
> `vendor/skills-for-copilot-studio`: its JSON schemas, templates, skills and patterns. Each guide ends with a
> **Sources** list pointing to the files it draws on. Microsoft describes that repo as experimental, and the YAML
> schema "may change without notice", so validate before you push (see [02 Project layout](02-project-layout.md#validation)).

## What is Copilot Studio, in one paragraph?

[Copilot Studio](glossary.md#copilot-studio) is Microsoft's low-code platform for building AI
[agents](glossary.md#agent) (chat assistants). An agent has **instructions** (a system prompt), **topics** (scripted
conversation flows), **tools** (calls to external APIs, workflows or other agents), **knowledge sources** (documents
and sites it can search to ground its answers) and **variables**. An LLM [orchestrator](glossary.md#orchestrator)
decides which of these to use for each user message. Agents are stored in
[Dataverse](glossary.md#dataverse), the Power Platform database. The Copilot Studio
[VS Code extension](glossary.md#vs-code-extension) can **clone** an agent into a folder of `.mcs.yml` files, and
**push** your edits back.

## Reading order

| # | Guide | Read it to learn… |
|---|---|---|
| 01 | [Concepts](01-concepts.md) | The mental model: what the pieces are and how a message is routed |
| 02 | [Project layout & lifecycle](02-project-layout.md) | Which file holds what; schema names and IDs; clone → edit → validate → push → publish |
| 03 | [The agent](03-agent.md) | `agent.mcs.yml` (instructions, starters, model) and `settings.mcs.yml` |
| 04 | [Topics & triggers](04-topics-and-triggers.md) | Writing a topic (`AdaptiveDialog`) and choosing its trigger |
| 05 | [Nodes](05-nodes.md) | Every step type a topic can contain: messages, questions, branches, loops, calls |
| 06 | [Variables, entities & Power Fx](06-variables-entities-powerfx.md) | Topic/global/system variables, data types, entities, the formula language |
| 07 | [Knowledge](07-knowledge.md) | Knowledge sources and generative answers |
| 08 | [Tools & connectors](08-tools-and-connectors.md) | `TaskDialog` tools, connector operations, connection references, HTTP calls |
| 09 | [MCP & multi-agent](09-mcp-and-multi-agent.md) | MCP servers, child agents, connected and external agents |
| 10 | [Flows & autonomous triggers](10-flows-and-autonomous-triggers.md) | Calling Power Automate / agent flows; event- and schedule-driven agents |
| 11 | [Messages & Adaptive Cards](11-messages-and-adaptive-cards.md) | Message formatting, variations, quick replies, cards and card forms |
| 12 | [Evaluations & testing](12-evaluations-and-testing.md) | Test sets, graders, and chatting with a published agent |
| 13 | [Patterns](13-patterns.md) | 15 proven design patterns with their YAML |
| 14 | [Gotchas](14-gotchas.md) | The pitfalls in one place |
| — | [Glossary](glossary.md) | Every Microsoft / Copilot Studio term used here |

### Generated reference

These pages are generated from the vendored schemas by `tools/gen_reference.py`. Use them to look up **every
property** of any `kind:`.

- [Schema A–Z index](reference/schema/index.md): all 744 definitions in the YAML authoring schema, organised into
  pages for [triggers](reference/schema/triggers.md), [actions/nodes](reference/schema/actions.md),
  [tool actions](reference/schema/task-actions.md), [knowledge](reference/schema/knowledge.md),
  [entities](reference/schema/entities.md), [data types](reference/schema/data-types.md),
  [messages](reference/schema/messages.md), [settings](reference/schema/settings.md),
  [components](reference/schema/components.md), [dialogs](reference/schema/dialogs.md),
  [evaluation](reference/schema/evaluation.md) and [other](reference/schema/other.md).
- [Adaptive Card schema](reference/adaptive-card-schema.md): every card element, input and action.
- [Connector operation catalogues](reference/connectors/index.md) for Office 365 Outlook, Office 365 Users,
  SharePoint and Azure DevOps.

## Conventions used in these docs

### The three kinds of value

Copilot Studio YAML values come in three flavours. Mixing them up is the most common authoring mistake.

| Looks like | Meaning | Example |
|---|---|---|
| plain text | A literal string | `activity: Hello there` |
| starts with `=` | A [Power Fx](glossary.md#power-fx) **expression**, evaluated at runtime | `value: =Topic.Count + 1` |
| contains `{…}` (no leading `=`) | A **template**: literal text with expressions interpolated inside braces | `activity: "Hi, I am {System.Bot.Name}!"` |

Expressions that contain YAML-special characters (`: `, `{`, `#`, quotes) must be quoted as a whole:
`value: "={ Name: \"x\" }"`. See [06 Variables, entities & Power Fx](06-variables-entities-powerfx.md#power-fx).

### Reading the property tables

Guides show the commonly used properties of each kind. The generated reference shows all of them. In both:

- **Req.** ✓ means the schema rejects the object without it.
- *Type* `=` expression, or "accepts a `=` Power Fx expression", means you can write either a literal or a formula.
- A link in the *Type* column leads to the definition of a nested object.
- `…NoKind` definitions are the same object written *without* its `kind:` line, because the parent property already
  implies the type.

### Example conventions

- IDs such as `sendMessage_a1B2c3` are random, but must be unique within a topic. See
  [02 Project layout](02-project-layout.md#schema-names-and-ids).
- `<angle brackets>` mark placeholders you must replace.
- YAML blocks are machine-checked against the schema by `tools/check_docs.py`. A block whose first line is
  `# fragment` shows a piece of a larger file (for example one node) and is only checked for valid YAML syntax.

## Exploring the schema from the command line

`cs-explain` (in [`cli/`](../cli)) works like `kubectl explain`. Give it a kind and an optional dotted field
path, and it prints each field's type, whether it is required, its default, its enum values, its union members
and a description. The three vendored schemas are embedded in the binary, so it needs no network access and no
Node.js.

```bash
make cli                                    # builds bin/cs-explain (needs Go 1.26+)
bin/cs-explain SendActivity                 # a topic node and its fields
bin/cs-explain AdaptiveDialog.beginDialog   # a union: lists every trigger kind
bin/cs-explain AdaptiveDialog.beginDialog.OnRecognizedIntent.actions.Question.entity
bin/cs-explain GptComponentMetadata --recursive --max-depth 2
bin/cs-explain card:Input.ChoiceSet         # Adaptive Card elements
bin/cs-explain connector:shared_office365users.UserProfile_V2.inputs
bin/cs-explain list                         # the 7 kinds a .mcs.yml file can hold, like `kubectl api-resources`
bin/cs-explain list --all --kinds-only      # every kind; --filter, --kinds-only and --undocumented imply --all
bin/cs-explain ConversationOutcome -o json  # the raw schema fragment, for scripts
source <(bin/cs-explain completion bash)    # tab-complete kinds and field paths
```

When a field holds a union, such as a topic's `actions:`, the next path segment picks the member kind. Type
`<boolean | =expr>` means the field takes a literal or a [Power Fx](glossary.md#power-fx) formula. The schema's
own descriptions are placeholders, so the text comes from a curated overlay,
`cli/internal/describe/descriptions.yaml`, which is written from these guides. `bin/cs-explain coverage` reports
how much of the schema the overlay covers. It also lists any overlay keys that no longer match the schema.

## Maintaining these docs

```bash
python3 -I tools/gen_reference.py   # regenerate docs/reference/** after updating the submodule
python3 -I tools/check_docs.py      # validate YAML examples against the schema, check links and anchors
make schemas test                   # re-embed the updated schemas in cs-explain and run its tests
```

Both scripts need Python 3.9+ with `pyyaml`. The checker also needs `jsonschema` (4.18+).
