# 11 — Messages & Adaptive Cards

This page covers everything an agent can *say*: plain and templated text, message variations, speech, quick-reply
buttons, card attachments, and [Adaptive Cards](glossary.md#adaptive-card) (JSON-described mini UIs that the chat
client draws natively), both as display-only messages and as forms that collect input.

## Where message content appears

A message in a [topic](glossary.md#topic) (a scripted conversation flow) is sent by a
[SendActivity](reference/schema/actions.md#sendactivity) [node](glossary.md#node) (one step in a topic). Its
`activity:` property has the type [ActivityTemplateBase](reference/schema/messages.md#activitytemplatebase), and the
same type is reused wherever the agent says something:

| Where | Property | Example use |
|---|---|---|
| [SendActivity](reference/schema/actions.md#sendactivity) | `activity` | Any message |
| [Question](reference/schema/actions.md#question) | `prompt`, `invalidPrompt`, `unrecognizedPrompt`, `defaultValueResponse` | Asking the user something (see [05 Nodes](05-nodes.md)) |
| [AdaptiveCardPrompt](reference/schema/actions.md#adaptivecardprompt) | `unrecognizedPrompt` | Re-prompt when the user types instead of submitting the card |
| [UpdateActivity](reference/schema/actions.md#updateactivity) | `activity` | Replace a message that was already sent (by `activityId`) |
| Generative-answer and hold settings | `latencyMessage` and others | "Please wait" style messages |

[ActivityTemplateBase](reference/schema/messages.md#activitytemplatebase) is a union of three shapes:
[MessageOptionalKind](reference/schema/messages.md#messageoptionalkind) (a normal chat message, written as a string
or as an object), [Activity](reference/schema/messages.md#activity) (a non-message
[activity](glossary.md#activity) (anything sent in a conversation) such as a typing indicator) and
[EventActivity](reference/schema/messages.md#eventactivity) (a named event). Almost everything you write is the first.

## Message shapes

### Plain string

The simplest message is a literal string:

```yaml
- kind: SendActivity
  id: sendMessage_pl41nS
  activity: Thanks, that's everything I need.
```

### Template string

Text with `{…}` placeholders is a **template**. Each placeholder is a [Power Fx](glossary.md#power-fx) (the
Excel-like formula language Copilot Studio uses) expression evaluated at runtime. Quote the whole value when it
starts with `{` or contains `: `, and use `|-` for multi-line text:

```yaml
- kind: SendActivity
  id: sendMessage_t3mpL8
  activity: "Hi, I'm {System.Bot.Name}. You have {CountRows(Topic.OpenTickets)} open tickets."

- kind: SendActivity
  id: sendMessage_mUlt1L
  activity: |-
    Error code: {System.Error.Code}
    Conversation Id: {System.Conversation.Id}
```

Don't put a leading `=` on a template: `activity: =...` would make the whole value an expression instead. See
[06 Variables, entities & Power Fx](06-variables-entities-powerfx.md#power-fx) for the expression language and the
`System.*` variables.

### Message object

For anything beyond one line of text, write `activity:` as an object. Its `kind: Message` line is optional
(the schema name is [MessageOptionalKind](reference/schema/messages.md#messageoptionalkind)), and the vendor
templates omit it.

```yaml
- kind: SendActivity
  id: sendMessage_0bj3ct
  activity:
    text:
      - An error has occurred. Error code {System.Error.Code}.
    speak:
      - An error has occurred, please try again.
```

| Property | Type/values | Meaning |
|---|---|---|
| `text` | list of strings (templates allowed) | The message text. Several entries = **message variations**; see below |
| `speak` | list of strings | What a voice channel says aloud instead of `text` |
| `randomizeVariantSelection` | boolean, default `true` | Whether a variation is picked at random *(inferred from the name; no guidance in the source repo)* |
| `attachments` | list of [ActivityAttachment](reference/schema/messages.md#activityattachment) | Cards or files shown with the message |
| `attachmentLayout` | `List` \| `Carousel` | How several attachments are arranged |
| `quickReplies` | list of [QuickReplyBase](reference/schema/messages.md#quickreplybase) | Suggestion buttons under the message |
| `allowBargeIn` | boolean, default `true` | Voice: presumably whether the caller can interrupt the speech *(inferred)* |
| `summary`, `expiration`, `importance` (`Low`/`Normal`/`High`), `channelData`, `value`, `citationEntities` | — | Advanced; no guidance in the source repo, see the [generated reference](reference/schema/messages.md#messageoptionalkind) |

None of the properties is required. The generated reference shows `Message`/`MessageOptionalKind` only as
"string or object", so the property list above was taken from the schema itself.

> **Naming trap.** `kind: Activity` is **not** "a message with extra options". It is the
> [Activity](reference/schema/messages.md#activity) template for non-message activities, whose
> `activityType` defaults to `Typing`. Use the message object above (no `kind:`, or `kind: Message`) for chat
> messages.

### Message variations

When `text:` has more than one entry, Copilot Studio sends **one** of them, chosen at random, which keeps
repetitive messages (greetings, acknowledgements) from sounding robotic. The Copilot Studio UI calls these
*message variations*.

```yaml
- kind: SendActivity
  id: sendMessage_v4r1aT
  activity:
    text:
      - Got it, one moment.
      - Sure, give me a second.
      - On it!
```

For random choice from a larger bank, or random text inside other logic, use Power Fx instead (see the
[knowledge hold message pattern](13-patterns.md#knowledge-hold-message)).

### Speech

`speak:` holds what voice channels (telephony, [IVR](glossary.md#ivr) (phone menu systems), speech clients) read
aloud. The vendor `OnError` template uses it to keep diagnostic details out of the spoken version: the screen shows
the error code and conversation ID, while the voice says only "An error has occurred, please try again." Text
channels ignore `speak:`.

## Line breaks and formatting

Most [channels](glossary.md#channel) (places users chat with the agent: [Teams](glossary.md#teams), web chat, …) collapse plain YAML
newlines into a single space, so a multi-line `|-` block still renders as one paragraph. The proven fix is HTML
break tags:

```yaml
- kind: SendActivity
  id: sendMessage_br3akS
  activity: |-
    Welcome to the HR Support Bot!
    <br /><br />
    I can help you with leave requests, benefits and payroll.
    <br /><br />
    Just type your question to get started.
```

| Syntax | Effect |
|---|---|
| `<br />` | Single line break |
| `<br /><br />` | Paragraph break |
| plain newline | Usually rendered as a space |

Use the same technique in `Question.prompt`. Test long formatted messages in the real target channel, because tall
chat bubbles look different in each client. See the [line breaks pattern](13-patterns.md#line-breaks-in-messages).

## Quick replies

[Quick replies](glossary.md#quick-reply) (clickable suggestion buttons) go under `quickReplies:`. Each entry is one
of four kinds:

| Kind | Required | What clicking it does |
|---|---|---|
| [MessageBack](reference/schema/messages.md#messageback) | `text` | Sends `text` to the agent as if the user had typed it. `displayText` can show something different in the transcript |
| [SendHiddenMessageBack](reference/schema/messages.md#sendhiddenmessageback) | `text` | Sends `text` to the agent without showing it in the transcript *(inferred from the name)* |
| [OpenUrl](reference/schema/messages.md#openurl) | `url` | Opens a web page |
| [MakeCall](reference/schema/messages.md#makecall) | `phoneNumber` | Starts a phone call |

All four also accept `title` (button label, may be a template or `=` expression), `image` and `imageAltText`.

```yaml
- kind: SendActivity
  id: sendMessage_qu1ckR
  activity:
    text:
      - What would you like to do next?
    quickReplies:
      - kind: MessageBack
        title: Check order status
        text: Where is my order?
      - kind: OpenUrl
        title: Open the help centre
        url: https://contoso.example.com/help
      - kind: MakeCall
        title: Call support
        phoneNumber: "+15555550100"
```

## Card attachments

`attachments:` holds one or more cards or files ([ActivityAttachment](reference/schema/messages.md#activityattachment)).
With several attachments, `attachmentLayout: Carousel` shows them side by side and `List` stacks them.

| Kind | Properties | Use for |
|---|---|---|
| [HeroCardTemplate](reference/schema/messages.md#herocardtemplate) | `title`, `subtitle`, `text`, `image`, `buttons` (quick-reply kinds) | A simple [hero card](glossary.md#hero-card) (title, image, buttons). Widely supported |
| [ImageCardTemplate](reference/schema/messages.md#imagecardtemplate) | `title`, `image` | One image with a caption |
| [VideoCardTemplate](reference/schema/messages.md#videocardtemplate) | `title`, `subtitle`, `text`, `image`, `media`, `buttons` | A playable video (`media` is the video URL) |
| [AdaptiveCardTemplate](reference/schema/messages.md#adaptivecardtemplate) | `cardContent` ✓ | A display-only Adaptive Card; see [below](#adaptive-cards) |
| [FileTemplate](reference/schema/messages.md#filetemplate) | `fileContent` ✓ (expression), `name`, `contentType` | Sends a file held in a variable *(no guidance in the source repo)* |

Every text property of these templates accepts a template string or a `=` expression. The
[Teams production hardening pattern](13-patterns.md#teams-production-hardening) sends this hero card when a session
has expired:

```yaml
- kind: SendActivity
  id: sendActivity_pgGjvA
  activity:
    attachments:
      - kind: HeroCardTemplate
        title: Session expired
        subtitle: New conversation started
        text: Your previous session ended due to inactivity. Your query is now treated as new. Restart anytime.
        buttons:
          - kind: MessageBack
            title: Start over
            text: Start over
```

A carousel of two image cards:

```yaml
- kind: SendActivity
  id: sendMessage_c4rOus
  activity:
    text:
      - Here are this week's offers.
    attachmentLayout: Carousel
    attachments:
      - kind: ImageCardTemplate
        title: Laptops
        image: https://contoso.example.com/img/laptops.png
      - kind: ImageCardTemplate
        title: Monitors
        image: https://contoso.example.com/img/monitors.png
```

## Channel differences

The same message can render differently per channel: Markdown support, card support and Adaptive Card version
limits all vary, and the Copilot Studio *Test* pane is not a real channel. To branch on the channel, read
`System.Activity.ChannelId`, which can be a **compound** value such as `msteams:Copilot` (an agent used inside
[Microsoft 365 Copilot](glossary.md#microsoft-365-copilot) in Teams). Compare with
`StartsWith(Lower(...), "msteams")`, never `=`:

```yaml
- kind: ConditionGroup
  id: conditionGroup_chN3l1
  conditions:
    - id: conditionItem_t3amS1
      condition: =StartsWith(Lower(System.Activity.ChannelId), "msteams")
      actions:
        - kind: SendActivity
          id: sendMessage_t3amS2
          activity: You're in Teams, so I'll show the Teams version of this card.
  elseActions:
    - kind: SendActivity
      id: sendMessage_w3bCh1
      activity: Here's the web version.
```

See the [channel-aware behavior pattern](13-patterns.md#channel-aware-behavior) for the table of known channel IDs.

## Adaptive Cards

### What an Adaptive Card is

An Adaptive Card is a **JSON document describing a small piece of UI**: text blocks, images, fact lists, columns,
input fields and buttons. You don't write HTML; the chat client (Teams, web chat, Outlook, …)
reads the JSON and draws the card with its own native controls and styling. The same JSON therefore looks at home
in each client. Cards can be purely informational, or they can be **forms**: when the user presses a submit button,
the values of all input fields are sent back to the agent.

```json
{
  "type": "AdaptiveCard",
  "$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
  "version": "1.5",
  "body": [
    { "type": "TextBlock", "text": "Order 1042", "weight": "Bolder", "size": "Large", "wrap": true },
    { "type": "FactSet", "facts": [ { "title": "Status", "value": "Shipped" } ] }
  ],
  "actions": [ { "type": "Action.Submit", "title": "OK" } ]
}
```

### Two ways to use a card

| | [AdaptiveCardTemplate](reference/schema/messages.md#adaptivecardtemplate) in a message | [AdaptiveCardPrompt](reference/schema/actions.md#adaptivecardprompt) node |
|---|---|---|
| Where | An entry in `SendActivity.activity.attachments` (or a question prompt) | A node of its own in `actions:` |
| Card property | `cardContent` | `card` |
| Waits for the user? | No: the topic continues immediately | Yes: the topic pauses until the card is submitted |
| Collects input? | No | Yes: `output.binding` maps card inputs to variables |
| Needs `Action.Submit`? | No | Yes (validator error `AdaptiveCardMissingActionSubmit` otherwise) |

The two sources disagree on display-only cards:

- The skills repo's `add-adaptive-card` skill says to use `AdaptiveCardPrompt` for **all** cards, including
  display-only ones (with a dummy "OK" button bound to an unused variable), and says not to use `SendActivity` with
  `attachments` for Adaptive Cards.
- Microsoft's [Ask with an Adaptive Card](https://learn.microsoft.com/en-us/microsoft-copilot-studio/authoring-ask-with-adaptive-card)
  page says the Adaptive Card node is meant for *interactive* cards, and that informational cards belong in a
  Message or Question node, which is the `AdaptiveCardTemplate` route.

Both shapes are schema-valid. Use `AdaptiveCardPrompt` when you want the user to acknowledge or answer before the
topic moves on, and an attachment when the card is just output.

A display-only card attached to a message. `cardContent` is a string, so embed the JSON with a `|` literal block:

```yaml
- kind: SendActivity
  id: sendMessage_cArd01
  activity:
    attachments:
      - kind: AdaptiveCardTemplate
        cardContent: |
          {
            "type": "AdaptiveCard",
            "$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
            "version": "1.5",
            "body": [
              { "type": "TextBlock", "text": "Order 1042", "weight": "Bolder", "size": "Large", "wrap": true },
              { "type": "FactSet", "facts": [ { "title": "Status", "value": "Shipped" } ] }
            ]
          }
```

### AdaptiveCardPrompt

[AdaptiveCardPrompt](reference/schema/actions.md#adaptivecardprompt) shows a card and waits for the user to submit
it. It is the node the Copilot Studio UI calls *Ask with Adaptive Card*.

| Property | Type/values | Req. | Meaning |
|---|---|---|---|
| `id` | string | ✓ | Unique node ID |
| `card` | string (card JSON) or `=` Power Fx record | ✓ | The card. Write JSON under `card: \|` |
| `output` | `binding:` map | ✓ | Card input `id` (or `Action.Submit` `data` key) → variable name, e.g. `userName: Topic.UserName` |
| `outputType` | `properties:` map | ✓ | The type of each bound field (`String`, `Boolean`, `Number`, …) |
| `repeatCount` | `0`, `1`, `2` (default) | | How many times the card is re-sent if the user types text instead of submitting |
| `unrecognizedPrompt` | message | | Message sent with each re-prompt |
| `interruptionPolicy` | object | | Whether the user can switch topics while the card is waiting |

Rules from the `add-adaptive-card` skill:

- Embed the JSON as a **literal block scalar**, `card: |`. Don't use `card: >` (it folds newlines) or inline JSON.
- Include `"$schema"` and `"version"` in the card JSON.
- `output`, `outputType` **and** an `Action.Submit` button are required on every card, even display-only ones.
  Omitting them gives [VS Code extension](glossary.md#vs-code-extension) (the Copilot Studio editor plugin) errors `MissingRequiredProperty` / `AdaptiveCardMissingActionSubmit`.
- `output.binding` values are variable names **without** `=` (`Topic.UserName`), and `binding:` has no `kind:`.
  The bound variables are created by the binding; you don't need `init:`.
- Every input needs a `label` when it is required or is an `Input.ChoiceSet`, otherwise you get
  `AdaptiveCardInputIsRequiredMissingLabel`.
- Card input `id`s must match the `output.binding` keys exactly.

### From submit to variables

When the user presses an `Action.Submit` button, the client sends back one object made of every input's `id` and
value, **merged with** the button's optional `data` object (the
[Action.Submit](reference/adaptive-card-schema.md#actionsubmit) schema describes `data` as hidden properties).
`output.binding` picks keys out of that object. `data` therefore lets you tell buttons apart, or send a value with
no input field at all:

```yaml
- kind: AdaptiveCardPrompt
  id: adaptiveCardPrompt_aCk9Lm
  card: |
    {
      "type": "AdaptiveCard",
      "$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
      "version": "1.5",
      "body": [
        { "type": "TextBlock", "text": "Your request has been logged.", "wrap": true }
      ],
      "actions": [
        { "type": "Action.Submit", "title": "OK", "data": { "acknowledged": true } }
      ]
    }
  output:
    binding:
      acknowledged: Topic.Acknowledged
  outputType:
    properties:
      acknowledged:
        type: Boolean
```

Microsoft Learn also recommends putting an identifier such as `"actionSubmitId": "booking_confirm"` in each
button's `data` when an agent shows several cards in a row. Users can still press buttons on older cards, and the ID
lets you detect which card a submission came from.

### Power Fx and dynamic content in cards

A card written as **JSON** is static: `{Topic.X}` placeholders and Power Fx are *not* evaluated inside it (the
`add-adaptive-card` skill is explicit about this). You have two options:

**Option 1: put dynamic text in a message next to the card.** This is what the skill recommends:

```yaml
- kind: SendActivity
  id: sendMessage_Rz4Wq1
  activity: "Hello {Topic.UserName}, please complete the form below."
```

**Option 2: write the whole card as a Power Fx record.** Both `card` and `cardContent` are typed in the schema as
expressions, and Microsoft Learn documents a *Formula* mode for the card node. In Formula mode the card is a Power Fx
record: property names are unquoted (or in single quotes, as with `'$schema'`), and values can reference variables.
The skills repo has no example of this, so the YAML below is assembled from the schema and the Learn page. Once the
UI converts a card to Formula mode it can't convert it back to JSON.

```yaml
- kind: AdaptiveCardPrompt
  id: adaptiveCardPrompt_fX7rQ2
  card: |-
    ={
      '$schema': "http://adaptivecards.io/schemas/adaptive-card.json",
      type: "AdaptiveCard",
      version: "1.5",
      body: [
        { type: "TextBlock", text: "Booking for " & Topic.CustomerName, weight: "Bolder", wrap: true },
        { type: "Input.Text", id: "myName", label: "Your name", isRequired: true, errorMessage: "Please enter your name." }
      ],
      actions: [ { type: "Action.Submit", title: "Submit", data: { actionSubmitId: "booking_submit" } } ]
    }
  output:
    binding:
      myName: Topic.MyName
  outputType:
    properties:
      myName:
        type: String
```

### Input validation

`style: "Email"` and `style: "Tel"` only change the on-screen keyboard; they **don't** validate. Use `regex` (blocks
submit when the pattern fails) and `isRequired: true` (blocks submit when empty), with an `errorMessage`:

```json
{
  "type": "Input.Text",
  "id": "email",
  "label": "Email address",
  "isRequired": true,
  "regex": "^[a-zA-Z0-9._%+\\-]+@[a-zA-Z0-9.\\-]+\\.[a-zA-Z]{2,}$",
  "errorMessage": "Enter a valid email address"
}
```

Inside a YAML literal block (`card: |`) the JSON is copied verbatim, so the backslashes above stay exactly as
written.

### Which version?

The card schema bundled with the skills repo is **1.6**, but the skills' examples all declare
`"version": "1.5"` "for maximum Teams + Copilot Chat compatibility". Microsoft Learn gives the reason: Copilot Studio
supports 1.6 and earlier, but Teams and the live-chat widget are limited to 1.5, and web chat supports 1.6 without
`Action.Execute`. Declare 1.5 unless you know every target channel supports 1.6. Each element and property in the
[Adaptive Card schema reference](reference/adaptive-card-schema.md) lists its minimum version (for example
[Table](reference/adaptive-card-schema.md#table) needs 1.5).

### The card model in brief

| Part | What it is | Common types |
|---|---|---|
| [AdaptiveCard](reference/adaptive-card-schema.md#adaptivecard) | The root object: `type`, `$schema`, `version`, `body`, `actions` | — |
| **Elements** (in `body`) | Display and layout | [TextBlock](reference/adaptive-card-schema.md#textblock), [RichTextBlock](reference/adaptive-card-schema.md#richtextblock), [Image](reference/adaptive-card-schema.md#image), [FactSet](reference/adaptive-card-schema.md#factset), [Container](reference/adaptive-card-schema.md#container), [ColumnSet](reference/adaptive-card-schema.md#columnset), [Table](reference/adaptive-card-schema.md#table), [ActionSet](reference/adaptive-card-schema.md#actionset) |
| **Inputs** (in `body`) | Form fields; each needs a unique `id` | [Input.Text](reference/adaptive-card-schema.md#inputtext), [Input.ChoiceSet](reference/adaptive-card-schema.md#inputchoiceset), [Input.Toggle](reference/adaptive-card-schema.md#inputtoggle), [Input.Date](reference/adaptive-card-schema.md#inputdate), [Input.Time](reference/adaptive-card-schema.md#inputtime), [Input.Number](reference/adaptive-card-schema.md#inputnumber) |
| **Actions** (in `actions` or an `ActionSet`) | Buttons | [Action.Submit](reference/adaptive-card-schema.md#actionsubmit), [Action.OpenUrl](reference/adaptive-card-schema.md#actionopenurl), [Action.ShowCard](reference/adaptive-card-schema.md#actionshowcard) (expand a nested card), [Action.ToggleVisibility](reference/adaptive-card-schema.md#actiontogglevisibility), [Action.Execute](reference/adaptive-card-schema.md#actionexecute) (not in web chat) |

`Input.ChoiceSet` renders as a dropdown with `"style": "compact"`, as radio buttons with `"expanded"`, and as a
searchable dropdown with `"filtered"` (Teams only).

### Card templates

These are adapted from the skills repo's `card-templates.md`, with real IDs filled in. Each block is a list of
nodes from a topic's `actions:`.

#### Form card

Collects several fields in one submission.

```yaml
- kind: AdaptiveCardPrompt
  id: adaptiveCardPrompt_f0rmA1
  card: |
    {
      "type": "AdaptiveCard",
      "$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
      "version": "1.5",
      "body": [
        { "type": "TextBlock", "text": "Request access", "weight": "Bolder", "size": "Large", "wrap": true },
        { "type": "TextBlock", "text": "Please complete all fields and submit.", "isSubtle": true, "wrap": true, "spacing": "None" },
        {
          "type": "Input.Text", "id": "userName", "label": "Your name",
          "placeholder": "Enter your name", "isRequired": true, "errorMessage": "Name is required"
        },
        {
          "type": "Input.Text", "id": "userEmail", "label": "Email address",
          "placeholder": "Enter your email address", "isRequired": true,
          "regex": "^[a-zA-Z0-9._%+\\-]+@[a-zA-Z0-9.\\-]+\\.[a-zA-Z]{2,}$",
          "errorMessage": "Enter a valid email address"
        },
        {
          "type": "Input.ChoiceSet", "id": "category", "label": "Category", "style": "compact",
          "isRequired": true, "errorMessage": "Please select a category",
          "choices": [
            { "title": "Hardware", "value": "hardware" },
            { "title": "Software", "value": "software" },
            { "title": "Other", "value": "other" }
          ]
        },
        { "type": "Input.Text", "id": "comments", "label": "Additional comments", "placeholder": "Optional...", "isMultiline": true }
      ],
      "actions": [ { "type": "Action.Submit", "title": "Submit" } ]
    }
  output:
    binding:
      userName: Topic.UserName
      userEmail: Topic.UserEmail
      category: Topic.Category
      comments: Topic.Comments
  outputType:
    properties:
      userName:
        type: String
      userEmail:
        type: String
      category:
        type: String
      comments:
        type: String
```

#### Confirmation card with branching

Asks yes/no with radio buttons, then branches on the bound variable.

```yaml
- kind: AdaptiveCardPrompt
  id: adaptiveCardPrompt_c0nf1R
  card: |
    {
      "type": "AdaptiveCard",
      "$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
      "version": "1.5",
      "body": [
        { "type": "TextBlock", "text": "Confirm action", "weight": "Bolder", "size": "Large", "wrap": true },
        {
          "type": "TextBlock", "text": "Are you sure you want to proceed? This action cannot be undone.",
          "wrap": true, "spacing": "Medium", "color": "Warning"
        },
        {
          "type": "Input.ChoiceSet", "id": "confirmed", "label": "Your choice", "style": "expanded",
          "isRequired": true, "errorMessage": "Please select an option",
          "choices": [
            { "title": "Yes, confirm", "value": "yes" },
            { "title": "No, cancel", "value": "no" }
          ]
        }
      ],
      "actions": [ { "type": "Action.Submit", "title": "Submit" } ]
    }
  output:
    binding:
      confirmed: Topic.UserConfirmed
  outputType:
    properties:
      confirmed:
        type: String

- kind: ConditionGroup
  id: conditionGroup_c0nf2B
  conditions:
    - id: conditionItem_c0nf3Y
      condition: =Topic.UserConfirmed = "yes"
      actions:
        - kind: SendActivity
          id: sendMessage_c0nf4Y
          activity: Action confirmed. Proceeding...
    - id: conditionItem_c0nf5N
      condition: =Topic.UserConfirmed = "no"
      actions:
        - kind: SendActivity
          id: sendMessage_c0nf6N
          activity: Action cancelled. No changes were made.
```

#### Info card (as a prompt)

Displays facts and waits for **OK**. Following the skill's rule, it binds a dummy acknowledgement variable that is
never read.

```yaml
- kind: AdaptiveCardPrompt
  id: adaptiveCardPrompt_1nf0Ck
  card: |
    {
      "type": "AdaptiveCard",
      "$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
      "version": "1.5",
      "body": [
        { "type": "TextBlock", "text": "Your laptop request", "weight": "Bolder", "size": "Large", "wrap": true },
        { "type": "TextBlock", "text": "Here is a summary of what we have on file.", "isSubtle": true, "wrap": true, "spacing": "None" },
        {
          "type": "FactSet",
          "facts": [
            { "title": "Model", "value": "Surface Laptop 7" },
            { "title": "Status", "value": "Approved" },
            { "title": "Delivery", "value": "5 working days" }
          ]
        }
      ],
      "actions": [ { "type": "Action.Submit", "title": "OK" } ]
    }
  output:
    binding:
      acknowledged: Topic.CardAcknowledged
  outputType:
    properties:
      acknowledged:
        type: String
```

## Other message-related kinds

| Kind | Purpose |
|---|---|
| [Activity](reference/schema/messages.md#activity) | Sends a non-message activity; `activityType` defaults to `Typing` (a typing indicator). Also `name`, `value`, `channelData`, `inputHint` |
| [EventActivity](reference/schema/messages.md#eventactivity) | Sends a named event (`name`, `value`) to the client, e.g. for a custom web-chat page to react to *(inferred; no guidance in the source repo)* |
| [UpdateActivity](reference/schema/actions.md#updateactivity) | Replaces an already-sent activity identified by `activityId` *(inferred)* |
| [ResponseTemplate](reference/schema/messages.md#responsetemplate) | A message plus a `mode` of `Strict` or `Generated`; used for a dialog's `response`/`confirmation` *(no guidance in the source repo)* |
| [LatencyMessageSettings](reference/schema/messages.md#latencymessagesettings) | A "still working" message shown after a delay on long-running nodes |

## Sources

- `vendor/skills-for-copilot-studio/skills/add-adaptive-card/SKILL.md`: `AdaptiveCardPrompt` structure and rules, output binding, validation, static card JSON
- `vendor/skills-for-copilot-studio/skills/add-adaptive-card/card-templates.md`: form, info, confirmation and minimal card templates; ChoiceSet styles; version table
- `vendor/skills-for-copilot-studio/patterns/line-breaks-in-messages.md`: `<br />` line breaks
- `vendor/skills-for-copilot-studio/patterns/channel-aware-behavior.md`: compound channel IDs
- `vendor/skills-for-copilot-studio/patterns/teams-production-hardening.md`: hero card with `MessageBack` button
- `vendor/skills-for-copilot-studio/templates/topics/error-handler.topic.mcs.yml`: `text:` + `speak:` message object
- `vendor/skills-for-copilot-studio/reference/bot.schema.yaml-authoring.json`: `MessageOptionalKind`, card templates, quick replies, `AdaptiveCardPrompt` (via [messages](reference/schema/messages.md) and [actions](reference/schema/actions.md))
- `vendor/skills-for-copilot-studio/reference/adaptive-card.schema.json`: card elements (via the [Adaptive Card schema reference](reference/adaptive-card-schema.md))
- Microsoft Learn, [Ask with an Adaptive Card](https://learn.microsoft.com/en-us/microsoft-copilot-studio/authoring-ask-with-adaptive-card): Formula mode, version limits per channel, `actionSubmitId`
