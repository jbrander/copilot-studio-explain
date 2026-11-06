# 06 — Variables, entities & Power Fx

This page covers where values live (topic, global and system variables), the data types the schema knows,
[entities](glossary.md#entity) (the value types a question can recognise), and [Power Fx](glossary.md#power-fx), the
Excel-like formula language used in every `=` expression. It includes the complete list of Power Fx functions that
Copilot Studio supports.

## Variable scopes

| Prefix | Scope | Lifetime | Declared |
|---|---|---|---|
| `Topic.<Name>` | [Topic variable](glossary.md#topic-variable) | The current topic only | Implicitly, on first assignment with `init:` (or as a topic input/output) |
| `Global.<Name>` | [Global variable](glossary.md#global-variable) | The whole conversation | In a file under `variables/` |
| `System.<Name>` | [System variable](glossary.md#system-variable) | Provided by the runtime | Built in; mostly read-only |

```yaml
# reading and writing each scope (three nodes)
- kind: SetVariable
  id: setVariable_sC0pe1
  variable: init:Topic.Greeting                     # create a topic variable
  value: ="Hello from " & System.Bot.Name           # read a system variable

- kind: SetVariable
  id: setVariable_sC0pe2
  variable: Global.LastGreeting                     # write a global (declared in variables/)
  value: =Topic.Greeting

- kind: SendActivity
  id: sendMessage_sC0pe3
  activity: "{Global.LastGreeting}"
```

### The `init:` rule

The **first** assignment to a topic variable must create it with `init:`. Later assignments omit it:

```yaml
# fragment — first assignment vs. later assignment
variable: init:Topic.UserEmail   # first assignment (SetVariable, Question, output binding, …)
variable: Topic.UserEmail        # later assignments
```

- `init:` applies to `Topic.*` only. Globals are declared by their file, and system variables already exist. For
  example, the disambiguation template writes `variable: System.Recognizer.SelectedIntent` without `init:`.
- Any node that writes a variable can do the `init:`: `SetVariable`, `Question.variable`, or
  `output: kind: SingleVariableOutputBinding` with `variable: init:Topic.Profile`.
- Topic inputs and outputs (`inputs:` / `outputType:`) declare their `Topic.<name>` variables, so the arithmetic
  template sets `variable: Topic.result` without `init:`.

## Global variables

A global variable is one file in `variables/`, kind [GlobalVariableComponent](reference/schema/components.md#globalvariablecomponent). On disk the file is **flattened**:
the variable's own fields sit at the root, although the schema nests them under `variable:` (see
[02 Project layout](02-project-layout.md#on-disk-shape-vs-schema-shape)).

```yaml
# Name: User Country
# The user's country, from their Microsoft 365 profile
name: UserCountry
aIVisibility: UseInAIContext
scope: Conversation
description: The user's country, from their Microsoft 365 profile
schemaName: cr26e_myAgent.globalvariable.UserCountry
kind: GlobalVariableComponent
defaultValue: Unknown
```

Save it as `variables/UserCountry.mcs.yml`. Keep the two comment lines: the first is the display name, the second
the description.

| Field | Type/values | Req. | Meaning |
|---|---|---|---|
| `name` | string (PascalCase) | ✓ | The identifier. Topics use `Global.<name>`. |
| `schemaName` | `<agentSchema>.globalvariable.<name>` | ✓ | Take the agent prefix from `schemaName` in `settings.mcs.yml`. |
| `aIVisibility` | `UseInAIContext` \| `Hidden` (schema default) | | [AI visibility](glossary.md#ai-visibility). `UseInAIContext` lets the [orchestrator](glossary.md#orchestrator) (the LLM that plans responses) read the value and reason about it: use it for user preferences and state that should affect routing or answers. `Hidden` is for counters, flags and other bookkeeping. |
| `scope` | `Conversation` | | The vendor skill says it is "always `Conversation`" for global variables. The schema's [VariableScope](reference/schema/actions.md#variablescope) also lists `Dialog`, `User` and `Environment`, without vendor guidance. |
| `description` | string | | What it holds. |
| `defaultValue` | value | | Initial value. The vendor template uses the placeholder `DEFAULT` when there's no specific value. |
| `isExternalInitializationAllowed` | boolean, default `false` | | No vendor guidance. By its name, this is the UI's "receive values from other agents" permission that a [connected agent](glossary.md#connected-agent)'s inputs need (inferred; see [09](09-mcp-and-multi-agent.md)). |
| `isOutputToExternalCallers`, `initializationTimeoutInMilliseconds`, `missingInitializationBehavior` (`UseDefault` \| `ThrowError` \| `Escalate`) | | | No vendor guidance. See [VariableNoKind](reference/schema/components.md#variablenokind). |

The `add-global-variable` skill's own template writes `aIVisibility: HideFromAIContext`. **That value isn't in the
schema**; use `Hidden`.

**Where globals are useful:**

- **Cross-topic state.** One topic sets a value and another reads it (preferred language, last search, a
  "conversation was reset" flag).
- **AI-aware context.** With `UseInAIContext`, the orchestrator sees the value when it plans.
- **In instructions.** Interpolate them into the agent's instructions, e.g. `User: {Global.UserDisplayName}` (see
  [03](03-agent.md#writing-instructions)).
- **Dynamic knowledge sources.** A knowledge source URL can be `=$"{Global.VarName}"`, so each user gets a different
  site or folder. The value must be a clean URL. See [07 Knowledge](07-knowledge.md).
- **Feeding tools.** A topic gathers data into a global, and a tool's input reads it with Power Fx (see
  [04 Topic-action chaining](04-topics-and-triggers.md#topic-action-chaining)).

A common way to fill globals is a one-shot initialisation topic: an `OnActivity` topic with `type: Message` and the
condition `=IsBlank(Global.X)` (see [04](04-topics-and-triggers.md#greeting)).

## System variables

Only variables used in the vendor skills, patterns and templates are listed. The runtime has more, but these are the
ones with evidence. Unless noted, they're read-only.

| Variable | Meaning | Where it's available / notes |
|---|---|---|
| `System.Activity.Text` | The user's current message | Everywhere. Common `userInput:` for search nodes. |
| `System.Activity.ChannelId` | The [channel](glossary.md#channel) (surface) the message came through, e.g. `msteams`, `directlinespeech` | Use `StartsWith(Lower(System.Activity.ChannelId), "msteams")` to catch every [Teams](glossary.md#teams) (Microsoft's chat app) surface. See [13 Channel-aware behavior](13-patterns.md#channel-aware-behavior). |
| `System.Activity.Channel` | The channel as an option-set value, compared with `ChannelId.MsTeams` | Appears only in a commented-out condition in the child-agent template. |
| `System.Activity.Value` | Structured payload of the incoming activity | In an `OnRedirect` topic, inputs from a calling agent arrive here, e.g. `System.Activity.Value.expenseReportFileFullPath`. |
| `System.Bot.Name` | The agent's name | e.g. in greetings |
| `System.Bot.Id`, `System.Bot.SchemaName` | The agent's ID and schema name | Diagnostics |
| `System.Bot.EnvironmentId`, `System.Bot.TenantId` | The [environment](glossary.md#environment) and [tenant](glossary.md#tenant) | Diagnostics |
| `System.Conversation.Id` | The conversation ID | Error messages, telemetry |
| `System.Conversation.InTestMode` | `true` in the authoring *Test* pane | Show error details only to makers |
| `System.User.Email`, `System.User.Id`, `System.User.Language` | The signed-in user | e.g. pass `=System.User.Email` to an [MCP](glossary.md#mcp) (Model Context Protocol) tool as a fixed input |
| `System.Error.Message`, `System.Error.Code` | The current error | `OnError` |
| `System.Error.SubCode` | The error's subcategory, e.g. `OpenAIJailBreak` for content-filter errors | `OnError`. See [13 RAI error handling](13-patterns.md#rai-error-handling). |
| `System.FallbackCount` | Number of consecutive fallbacks | `OnUnknownIntent` |
| `System.SignInReason` | Why sign-in was triggered; compare with `SignInReason.SignInRequired` | `OnSignIn` |
| `System.Recognizer.IntentOptions` | The candidate topics (table of `DisplayName`, `TopicId`, `TriggerId`, `Score`) | `OnSelectIntent` |
| `System.Recognizer.SelectedIntent` | The chosen candidate (`.TopicId`, …). **Writable**: set it to choose the topic | `OnSelectIntent` |
| `System.SearchQuery` | AI-rewritten search query | `OnKnowledgeRequested` |
| `System.KeywordSearchQuery` | Keyword version of the query | `OnKnowledgeRequested` |
| `System.SearchResults` | **Writable** table of custom results: `Content`, `ContentLocation`, `Title` | `OnKnowledgeRequested` |
| `System.Response.FormattedText` | The AI-generated response text | `OnGeneratedResponse` |
| `System.ContinueResponse` | **Writable**: set to `false` to stop the generated response being sent | `OnGeneratedResponse` |

`ClearAllVariables` with `variables: ConversationHistory` clears the conversation history the AI sees (see
[05](05-nodes.md#clearallvariables)).

## Data types

Wherever the schema wants a type, it uses the [DataType](reference/schema/data-types.md#datatype) union. That includes `inputType`/`outputType`
properties, `ParseValue.valueType`, and the response schemas of HTTP and connector calls. Simple types can be written
as a bare name; `Record` and `Table` need `kind:` plus nested `properties:`.

| Type | Write as | Notes |
|---|---|---|
| [String](reference/schema/data-types.md#string) | `String` | Text. The most common type. |
| [Number](reference/schema/data-types.md#number) | `Number` | Numeric (decimal). |
| [Float](reference/schema/data-types.md#float) | `Float` | Floating-point number. No vendor guidance on when to prefer it over `Number`. |
| [Boolean](reference/schema/data-types.md#boolean) | `Boolean` | `true` / `false`. |
| [Date](reference/schema/data-types.md#date), [DateTime](reference/schema/data-types.md#datetime), [DateTimeNoTimeZone](reference/schema/data-types.md#datetimenotimezone), [Time](reference/schema/data-types.md#time) | bare name | Dates and times. |
| [Guid](reference/schema/data-types.md#guid), [Hyperlink](reference/schema/data-types.md#hyperlink), [Color](reference/schema/data-types.md#color), [File](reference/schema/data-types.md#file) | bare name | Special scalars. |
| [Any](reference/schema/data-types.md#any) | `Any` | Untyped. |
| [Blank](reference/schema/data-types.md#blank), [Unspecified](reference/schema/data-types.md#unspecified), [UnresolvedDynamicType](reference/schema/data-types.md#unresolveddynamictype) | bare name | Placeholders the tools may write; not something you'd choose. |
| [Record](reference/schema/data-types.md#record) | `kind: Record` + `properties:` | One object with named fields. |
| [Table](reference/schema/data-types.md#table) | `kind: Table` + `properties:` | A list of records, each with the given fields. |
| [ClosedListOptionSet](reference/schema/data-types.md#closedlistoptionset) | `kind: ClosedListOptionSet` + `schemaName` | The values of a custom closed-list entity. |
| [EmbeddedOptionSet](reference/schema/data-types.md#embeddedoptionset) | `kind: EmbeddedOptionSet` + `dialogSchemaName`, `triggerId`, `actionId` | Option set defined inline in a node. Tool-generated. |
| [SystemOptionSet](reference/schema/data-types.md#systemoptionset) | `kind: SystemOptionSet` + `name` | A built-in enum such as `SignInReason`, `VariablesToClear`, `SkipQuestionMode`. |

### In `inputType` / `outputType`

Each property under `properties:` is keyed by the variable name and has a `type:` (plus `displayName`,
`description`, and optionally `isRequired`):

```yaml
kind: AdaptiveDialog
beginDialog:
  kind: OnRecognizedIntent
  id: main
  intent: {}
  actions:
    - kind: SetVariable
      id: setVariable_tYpe01
      variable: Topic.found
      value: =!IsEmpty(Topic.lines)
outputType:
  properties:
    found:
      displayName: found
      description: Whether any order lines were found
      type: Boolean
    order:
      displayName: order
      type:
        kind: Record
        properties:
          orderId: String
          total: Number
          placedOn: DateTime
    lines:
      displayName: lines
      type:
        kind: Table
        properties:
          sku: String
          qty: Number
```

## Entities

An entity tells a `Question` node (or an `AutomaticTaskInput`) **what kind of value to recognise** in the user's
reply, and stores the recognised value in the question's variable. Recognising a value from something
the user has already said is [slot filling](glossary.md#slot-filling) (skipping questions that are already answered).
The `entity:` property takes one of the [EntityReference](reference/schema/entities.md#entityreference) shapes.

### Prebuilt entities

A [prebuilt entity](glossary.md#prebuilt-entity) is referenced by bare name (`entity: NumberPrebuiltEntity`) or as
`kind:` with options (`includeMetadata`, `allowMultipleValues`, `sensitivityLevel`, …). The int-reference skill
lists the first five as the common ones.

| Group | Entities |
|---|---|
| **Common** | [BooleanPrebuiltEntity](reference/schema/entities.md#booleanprebuiltentity) (yes/no), [NumberPrebuiltEntity](reference/schema/entities.md#numberprebuiltentity), [StringPrebuiltEntity](reference/schema/entities.md#stringprebuiltentity) (free text: accepts anything), [DateTimePrebuiltEntity](reference/schema/entities.md#datetimeprebuiltentity), [EmailPrebuiltEntity](reference/schema/entities.md#emailprebuiltentity) |
| Numbers & measures | [AgePrebuiltEntity](reference/schema/entities.md#ageprebuiltentity), [MoneyPrebuiltEntity](reference/schema/entities.md#moneyprebuiltentity), [OrdinalPrebuiltEntity](reference/schema/entities.md#ordinalprebuiltentity), [PercentagePrebuiltEntity](reference/schema/entities.md#percentageprebuiltentity), [SpeedPrebuiltEntity](reference/schema/entities.md#speedprebuiltentity), [TemperaturePrebuiltEntity](reference/schema/entities.md#temperatureprebuiltentity), [WeightPrebuiltEntity](reference/schema/entities.md#weightprebuiltentity) |
| Dates & times | [DatePrebuiltEntity](reference/schema/entities.md#dateprebuiltentity), [DateTimeNoTimeZonePrebuiltEntity](reference/schema/entities.md#datetimenotimezoneprebuiltentity), [DurationPrebuiltEntity](reference/schema/entities.md#durationprebuiltentity) |
| Places | [CityPrebuiltEntity](reference/schema/entities.md#cityprebuiltentity), [StatePrebuiltEntity](reference/schema/entities.md#stateprebuiltentity), [CountryOrRegionPrebuiltEntity](reference/schema/entities.md#countryorregionprebuiltentity), [ContinentPrebuiltEntity](reference/schema/entities.md#continentprebuiltentity), [ZipCodePrebuiltEntity](reference/schema/entities.md#zipcodeprebuiltentity), [StreetAddressPrebuiltEntity](reference/schema/entities.md#streetaddressprebuiltentity), [PointOfInterestPrebuiltEntity](reference/schema/entities.md#pointofinterestprebuiltentity) |
| People & things | [PersonNamePrebuiltEntity](reference/schema/entities.md#personnameprebuiltentity), [OrganizationPrebuiltEntity](reference/schema/entities.md#organizationprebuiltentity), [PhoneNumberPrebuiltEntity](reference/schema/entities.md#phonenumberprebuiltentity), [URLPrebuiltEntity](reference/schema/entities.md#urlprebuiltentity), [ColorPrebuiltEntity](reference/schema/entities.md#colorprebuiltentity), [LanguagePrebuiltEntity](reference/schema/entities.md#languageprebuiltentity), [EventPrebuiltEntity](reference/schema/entities.md#eventprebuiltentity), [FilePrebuiltEntity](reference/schema/entities.md#fileprebuiltentity) |

Spelling gotcha: the int-reference table writes `EMailPrebuiltEntity`. The schema name is **`EmailPrebuiltEntity`**.

### Closed lists

A closed list is a fixed set of choices, each with synonyms. There are two ways to use one in a `Question`.

**Reference a custom entity** defined in the agent, by its schema name `<agentSchema>.entity.<Name>`
([ClosedListEntityReference](reference/schema/entities.md#closedlistentityreference)). This is the form used in the line-breaks and Teams-hardening patterns:

```yaml
- kind: Question
  id: question_c0nF1m
  variable: init:Topic.Confirm
  alwaysPrompt: true
  prompt: Do you want to continue?
  entity:
    kind: ClosedListEntityReference
    entityId: cr26e_myAgent.entity.YesNo
```

The entity itself is a separate component ([CustomEntityComponent](reference/schema/components.md#customentitycomponent) holding a [ClosedListEntity](reference/schema/dialogs.md#closedlistentity)). The source
repo doesn't show its on-disk file, so create custom entities in the UI and reference them from YAML.

**Define the list inline** with [EmbeddedEntity](reference/schema/entities.md#embeddedentity) (a schema-valid shape; not shown in the vendor repo):

```yaml
- kind: Question
  id: question_dEpt01
  variable: init:Topic.Department
  prompt: Which department is this about?
  entity:
    kind: EmbeddedEntity
    definition:
      kind: ClosedListEntity
      items:
        - id: hr
          displayName: HR
          synonyms:
            - Human Resources
            - People team
        - id: it
          displayName: IT
          synonyms:
            - tech support
            - helpdesk
```

`ClosedListEntity` also has `smartMatchingEnabled` (default `true`) and `samples`.

### Regular expressions

A regex entity matches a pattern, such as an order code. Reference a custom one with
[RegexEntityReference](reference/schema/entities.md#regexentityreference) (`entityId: <agentSchema>.entity.<Name>`), or embed a [RegexEntity](reference/schema/dialogs.md#regexentity) with a `pattern:`.
Neither form appears in the vendor repo; this is the schema shape:

```yaml
- kind: Question
  id: question_oRd3r1
  variable: init:Topic.OrderCode
  prompt: What's your order code (e.g. ABC-1234)?
  entity:
    kind: EmbeddedEntity
    definition:
      kind: RegexEntity
      pattern: "^[A-Z]{3}-\\d{4}$"
```

### Dynamic lists

[DynamicClosedListEntity](reference/schema/entities.md#dynamicclosedlistentity) builds the choices at runtime from a table expression in `items:`. The disambiguation
template uses it to offer the matched topics, plus "None of these", as choices:

```yaml
- kind: Question
  id: question_dYn4m1
  interruptionPolicy:
    allowInterruption: false
  alwaysPrompt: true
  variable: System.Recognizer.SelectedIntent
  prompt: "To clarify, did you mean:"
  entity:
    kind: DynamicClosedListEntity
    items: =Topic.IntentOptions
```

### Other entity references

| Kind | Purpose |
|---|---|
| [OneOfMultiEntityReference](reference/schema/entities.md#oneofmultientityreference) | Accept any one of several entities: `entities:` is a list of `{entity, propertyName}`. No vendor guidance. |
| [ExternalEntityReference](reference/schema/entities.md#externalentityreference) | Reference an entity by `entityId` from outside the agent. No vendor guidance. |

## Power Fx

Power Fx is the low-code formula language of the Microsoft [Power Platform](glossary.md#power-platform) (Microsoft's low-code suite). Copilot Studio uses it for
conditions, computed values and interpolation.

### Expressions, templates and literals

| You write | It is | Used in |
|---|---|---|
| `value: =Topic.Count + 1` | **Expression**: the whole value after `=` is Power Fx | `SetVariable.value`, `condition`, `userInput`, `items`, `dialog`, input bindings, `properties` |
| `activity: "You have {Topic.Count} items"` | **Template**: literal text with `{expression}` holes | `SendActivity`, `Question.prompt`, `SetTextVariable.value`, agent `instructions` |
| `activity: Thanks!` | **Literal** | anywhere |

- `{…}` in a template can hold any expression, not only a variable: `{Text(Today(),DateTimeFormat.LongDate)}`.
- To turn a number or date into text, use either a template (`SetTextVariable` with `"Guests: {Topic.N}"`) or
  `=Text(…)`.
- Variable *targets* (`variable:`, output `binding:` values) are names, written **without** `=`.
- A missing `=` is a common validation error. Without it, `value: Topic.A + 1` is the literal text "Topic.A + 1".

### Quoting rules in YAML

Power Fx and YAML share several characters. Use these rules:

| Situation | Write | Why |
|---|---|---|
| Expression contains `: `, `#`, or starts with `={` | `value: "={ Name: \"x\" }"` (quote the whole value; escape inner `"` as `\"`) | Otherwise YAML reads `: ` as a mapping and `#` as a comment. |
| Long or nested expression | `value: \|-` then the expression on indented lines, starting with `=` | A block scalar needs no escaping. Used for multi-line records and `Switch(…)`. |
| Template starts with `{` | `activity: "{Topic.CleanedText}"` | An unquoted leading `{` starts a YAML flow mapping. |
| Text ends with `:` or contains `: ` | `prompt: "To clarify, did you mean:"` | Same as above. |
| Text literals *inside* Power Fx | `"Unknown"` (double quotes) | In Power Fx, single quotes delimit names, not text. |
| [OData](glossary.md#odata) (a REST query convention) `$` names | See [08 `$`-prefixed inputs](08-tools-and-connectors.md#-prefixed-inputs-odata-parameters) | Tools and inline connector calls use different forms. |

### Operators and built-in enums seen in the sources

- Comparison `=` and `<`; logical not `!` (e.g. `!IsBlank(x)`); text concatenation `&`
  (`"AssignedTo/EMail eq '" & Topic.UserEmail & "'"`).
- Record literal `{ Field: value, … }`. Field access `Topic.Profile.mail`. `ThisRecord` inside `ForAll`.
- String interpolation inside an expression: `=$"{Global.VarName}"`.
- Enums: `DateTimeFormat.UTC`, `DateTimeFormat.LongDate`, `SignInReason.SignInRequired`, `ChannelId.MsTeams`.

### Supported functions

**Only these functions work in Copilot Studio.** It supports a subset of Power Fx, and any function not on this list
causes an error. (The list is from the vendor int-reference skill,
which presents it as complete.)

| Category | Functions |
|---|---|
| **Math** | `Abs`, `Acos`, `Acot`, `Asin`, `Atan`, `Atan2`, `Cos`, `Cot`, `Degrees`, `Exp`, `Int`, `Ln`, `Log`, `Mod`, `Pi`, `Power`, `Radians`, `Rand`, `RandBetween`, `Round`, `RoundDown`, `RoundUp`, `Sin`, `Sqrt`, `Sum`, `Tan`, `Trunc` |
| **Text** | `Char`, `Concat`, `Concatenate`, `EncodeHTML`, `EncodeUrl`, `EndsWith`, `Find`, `Left`, `Len`, `Lower`, `Match`, `MatchAll`, `Mid`, `PlainText`, `Proper`, `Replace`, `Right`, `Search`, `Split`, `StartsWith`, `Substitute`, `Text`, `Trim`, `TrimEnds`, `UniChar`, `Upper`, `Value` |
| **Date/Time** | `Date`, `DateAdd`, `DateDiff`, `DateTime`, `DateTimeValue`, `DateValue`, `Day`, `EDate`, `EOMonth`, `Hour`, `IsToday`, `Minute`, `Month`, `Now`, `Second`, `Time`, `TimeValue`, `TimeZoneOffset`, `Today`, `Weekday`, `WeekNum`, `Year` |
| **Logical** | `And`, `Coalesce`, `If`, `IfError`, `IsBlank`, `IsBlankOrError`, `IsEmpty`, `IsError`, `IsMatch`, `IsNumeric`, `IsType`, `Not`, `Or`, `Switch` |
| **Table** | `AddColumns`, `Column`, `ColumnNames`, `Count`, `CountA`, `CountIf`, `CountRows`, `Distinct`, `DropColumns`, `Filter`, `First`, `FirstN`, `ForAll`, `Index`, `Last`, `LastN`, `LookUp`, `Patch`, `Refresh`, `RenameColumns`, `Sequence`, `ShowColumns`, `Shuffle`, `Sort`, `SortByColumns`, `Summarize`, `Table` |
| **Aggregate** | `Average`, `Max`, `Min`, `StdevP`, `VarP` |
| **Type conversion** | `AsType`, `Boolean`, `Dec2Hex`, `Decimal`, `Float`, `GUID`, `Hex2Dec`, `JSON`, `ParseJSON` |
| **Other** | `Blank`, `ColorFade`, `ColorValue`, `Error`, `Language`, `OptionSetInfo`, `RGBA`, `Trace`, `With` |

### Recipes

All taken from vendor templates, skills or patterns:

| Goal | Expression |
|---|---|
| Add two numbers and return text | `=Text(Topic.number1 + Topic.number2)` |
| Current UTC timestamp as text | `=Text(Now(), DateTimeFormat.UTC)` |
| Today's date, unambiguous (in instructions) | `{Text(Today(),DateTimeFormat.LongDate)}` |
| Is a value set? | `=!IsBlank(Topic.Answer)` |
| Default when blank | `=If(IsBlank(Topic.M365Profile.country), "Unknown", Topic.M365Profile.country)` |
| Limit retries | `=System.FallbackCount < 3` |
| Only in the test pane | `=System.Conversation.InTestMode = true` |
| Any Teams surface | `=StartsWith(Lower(System.Activity.ChannelId), "msteams")` |
| Pick a topic dynamically | `=Switch(Topic.N, 1, "cr26e_myAgent.topic.Lesson1", 2, "cr26e_myAgent.topic.Lesson2", "cr26e_myAgent.topic.Fallback")` |
| Random number | `=RandBetween(0,4)` |
| Build a record (quoted) | `"={ DisplayName: Topic.NoneOfTheseDisplayName, TopicId: \"NoTopic\", TriggerId: \"NoTrigger\", Score: 1.0 }"` |
| Reshape a table | `"=ForAll(Topic.RawResults, { Content: ThisRecord.body, ContentLocation: ThisRecord.url, Title: ThisRecord.title })"` |
| Strip text | `=Substitute(System.Response.FormattedText, "[1]", "")` |
| Build an OData filter | `="AssignedTo/EMail eq '" & Topic.UserEmail & "'"` |

## Sources

- `vendor/skills-for-copilot-studio/skills/int-reference/SKILL.md`: scopes, system variables, prebuilt entities, Power Fx examples and the supported-function list
- `vendor/skills-for-copilot-studio/skills/add-global-variable/SKILL.md`: global-variable file format and uses
- `vendor/skills-for-copilot-studio/templates/variables/global-variable.variable.mcs.yml`: template
- `vendor/skills-for-copilot-studio/skills/new-topic/SKILL.md`: `init:`, `SetTextVariable` for coercion
- `vendor/skills-for-copilot-studio/patterns/orchestrator-variables.md`, `patterns/date-context.md`, `patterns/channel-aware-behavior.md`, `patterns/teams-production-hardening.md`, `patterns/rai-error-handling.md`, `patterns/line-breaks-in-messages.md`, `patterns/dynamic-topic-redirect.md`
- `vendor/skills-for-copilot-studio/skills/add-other-agents/SKILL.md`: `System.Activity.Value`, external-source globals
- `vendor/skills-for-copilot-studio/skills/edit-action/SKILL.md`, `skills/edit-action/sharepoint-actions.md`: `System.User.Email`, OData expression
- `vendor/skills-for-copilot-studio/skills/add-knowledge/SKILL.md`: globals in knowledge URLs
- `vendor/skills-for-copilot-studio/templates/topics/*.topic.mcs.yml`: disambiguation (`DynamicClosedListEntity`), error handler, custom knowledge source, remove citations
- Generated reference: [entities](reference/schema/entities.md), [data types](reference/schema/data-types.md), [components](reference/schema/components.md)
