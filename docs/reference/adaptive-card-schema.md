# Adaptive Card schema reference

> **Generated** by `tools/gen_reference.py` from `reference/adaptive-card.schema.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

An [Adaptive Card](../glossary.md#adaptive-card) is a platform-neutral JSON description of a small UI (text, images, inputs, buttons) that the chat client draws natively. In Copilot Studio YAML you embed the card JSON inside `AdaptiveCardTemplate` (in a message) or `AdaptiveCardPrompt` (to collect input). This page lists every element in the schema bundled with the skills repo (`http://adaptivecards.io/schemas/adaptive-card.json`). For how to use cards in topics see [11 Messages & Adaptive Cards](../11-messages-and-adaptive-cards.md).

Every element is a JSON object with a `"type"` property naming the element (e.g. `"type": "TextBlock"`). *Version* is the minimum Adaptive Card schema version that supports the element or property; this bundled schema is version **1.6**, while the skills' examples declare `"version": "1.5"` for broad client support (see the [card guide](../11-messages-and-adaptive-cards.md)).

## The card

### AdaptiveCard

An Adaptive Card, containing a free-form body of card elements, and an optional set of actions.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `AdaptiveCard` |  |  |  | Must be `AdaptiveCard` |
| `version` | `string` |  |  |  | Schema version that this card requires. If a client is **lower** than this version, the `fallbackText` will be rendered. NOTE: Version is not required for cards within an `Action.ShowCard`. However, it *is* required for the top-level card. |
| `refresh` | [Refresh](#refresh) |  |  | 1.4 | Defines how the card can be refreshed by making a request to the target Bot. |
| `authentication` | [Authentication](#authentication) |  |  | 1.4 | Defines authentication information to enable on-behalf-of single sign on or just-in-time OAuth. |
| `body` | list of [ImplementationsOf.Element](#implementationsofelement) |  |  |  | The card elements to show in the primary card region. |
| `actions` | list of [ImplementationsOf.Action](#implementationsofaction) |  |  |  | The Actions to show in the card's action bar. |
| `selectAction` | [ImplementationsOf.ISelectAction](#implementationsofiselectaction) |  |  | 1.1 | An Action that will be invoked when the card is tapped or selected. `Action.ShowCard` is not supported. |
| `fallbackText` | `string` |  |  |  | Text shown when the client doesn't support the version specified (may contain markdown). |
| `backgroundImage` | [BackgroundImage](#backgroundimage) \| `string` |  |  | 1.2 | Specifies the background image of the card. |
| `metadata` | [Metadata](#metadata) |  |  | 1.6 | Defines various metadata properties typically not used for rendering the card |
| `minHeight` | `string` |  |  | 1.2 | Specifies the minimum height of the card. |
| `rtl` | `boolean` |  |  | 1.5 | When `true` content in this Adaptive Card should be presented right to left. When 'false' content in this Adaptive Card should be presented left to right. If unset, the default platform behavior will apply. |
| `speak` | `string` |  |  |  | Specifies what should be spoken for this entire card. This is simple text or SSML fragment. |
| `lang` | `string` |  |  |  | The 2-letter ISO-639-1 language used in the card. Used to localize any date/time functions. |
| `verticalContentAlignment` | [VerticalContentAlignment](#verticalcontentalignment) |  |  | 1.1 | Defines how the content should be aligned vertically within the container. Only relevant for fixed-height cards, or cards with a `minHeight` specified. |
| `$schema` | `string` |  |  |  | The Adaptive Card schema. |

## Card elements (layout & display)

### ActionSet

Displays a set of actions.  
*Version:* 1.2  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `ActionSet` |  |  |  | Must be `ActionSet` |
| `actions` | list of [ImplementationsOf.Action](#implementationsofaction) | ✓ |  |  | The array of `Action` elements to show. |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Element](#extendableelement))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Element](#extendableelement))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Element](#extendableelement))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Element](#extendableelement))* |
| `id` | `string` |  |  |  | A unique identifier associated with the item. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### CaptionSource

Defines a source for captions  
*Version:* 1.6  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `CaptionSource` |  |  |  | Must be `CaptionSource` |
| `mimeType` | `string` | ✓ |  |  | Mime type of associated caption file (e.g. `"vtt"`). For rendering in JavaScript, only `"vtt"` is supported, for rendering in UWP, `"vtt"` and `"srt"` are supported. |
| `url` | `string` | ✓ |  |  | URL to captions. |
| `label` | `string` | ✓ |  |  | Label of this caption to show to the user. |

### Column

Defines a container that is part of a ColumnSet.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Column` |  |  |  | Must be `Column` |
| `items` | list of [ImplementationsOf.Element](#implementationsofelement) |  |  |  | The card elements to render inside the `Column`. |
| `backgroundImage` | [BackgroundImage](#backgroundimage) \| `string` |  |  | 1.2 | Specifies the background image. Acceptable formats are PNG, JPEG, and GIF |
| `bleed` | `boolean` |  |  | 1.2 | Determines whether the column should bleed through its parent's padding. |
| `fallback` | [Column](#column) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown item is encountered or the requires of this or any children can't be met. |
| `minHeight` | `string` |  |  | 1.2 | Specifies the minimum height of the column in pixels, like `"80px"`. |
| `rtl` | `boolean` |  |  | 1.5 | When `true` content in this column should be presented right to left. When 'false' content in this column should be presented left to right. When unset layout direction will inherit from parent container or column. If unset in all ancestors, the default platform behavior will apply. |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line between this column and the previous column. |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this column and the preceding column. |
| `selectAction` | [ImplementationsOf.ISelectAction](#implementationsofiselectaction) |  |  | 1.1 | An Action that will be invoked when the `Column` is tapped or selected. `Action.ShowCard` is not supported. |
| `style` | [ContainerStyle](#containerstyle) |  |  |  | Style hint for `Column`. |
| `verticalContentAlignment` | [VerticalContentAlignment](#verticalcontentalignment) |  |  | 1.1 | Defines how the content should be aligned vertically within the column. When not specified, the value of verticalContentAlignment is inherited from the parent container. If no parent container has verticalContentAlignment set, it defaults to Top. |
| `width` | `string` \| `number` |  |  |  | `"auto"`, `"stretch"`, a number representing relative width of the column in the column group, or in version 1.1 and higher, a specific pixel width, like `"50px"`. |
| `id` | `string` |  |  |  | A unique identifier associated with the item. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### ColumnSet

ColumnSet divides a region into Columns, allowing elements to sit side-by-side.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `ColumnSet` |  |  |  | Must be `ColumnSet` |
| `columns` | list of [Column](#column) |  |  |  | The array of `Columns` to divide the region into. |
| `selectAction` | [ImplementationsOf.ISelectAction](#implementationsofiselectaction) |  |  | 1.1 | An Action that will be invoked when the `ColumnSet` is tapped or selected. `Action.ShowCard` is not supported. |
| `style` | [ContainerStyle](#containerstyle) |  |  | 1.2 | Style hint for `ColumnSet`. |
| `bleed` | `boolean` |  |  | 1.2 | Determines whether the element should bleed through its parent's padding. |
| `minHeight` | `string` |  |  | 1.2 | Specifies the minimum height of the column set in pixels, like `"80px"`. |
| `horizontalAlignment` | [HorizontalAlignment](#horizontalalignment) |  |  |  | Controls the horizontal alignment of the ColumnSet. When not specified, the value of horizontalAlignment is inherited from the parent container. If no parent container has horizontalAlignment set, it defaults to Left. |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Element](#extendableelement))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Element](#extendableelement))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Element](#extendableelement))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Element](#extendableelement))* |
| `id` | `string` |  |  |  | A unique identifier associated with the item. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### Container

Containers group items together.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Container` |  |  |  | Must be `Container` |
| `items` | list of [ImplementationsOf.Element](#implementationsofelement) | ✓ |  |  | The card elements to render inside the `Container`. |
| `selectAction` | [ImplementationsOf.ISelectAction](#implementationsofiselectaction) |  |  | 1.1 | An Action that will be invoked when the `Container` is tapped or selected. `Action.ShowCard` is not supported. |
| `style` | [ContainerStyle](#containerstyle) |  |  |  | Style hint for `Container`. |
| `verticalContentAlignment` | [VerticalContentAlignment](#verticalcontentalignment) |  |  | 1.1 | Defines how the content should be aligned vertically within the container. When not specified, the value of verticalContentAlignment is inherited from the parent container. If no parent container has verticalContentAlignment set, it defaults to Top. |
| `bleed` | `boolean` |  |  | 1.2 | Determines whether the element should bleed through its parent's padding. |
| `backgroundImage` | [BackgroundImage](#backgroundimage) \| `string` |  |  | 1.2 | Specifies the background image. Acceptable formats are PNG, JPEG, and GIF |
| `minHeight` | `string` |  |  | 1.2 | Specifies the minimum height of the container in pixels, like `"80px"`. |
| `rtl?` | `boolean` |  |  | 1.5 | When `true` content in this container should be presented right to left. When 'false' content in this container should be presented left to right. When unset layout direction will inherit from parent container or column. If unset in all ancestors, the default platform behavior will apply. |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Element](#extendableelement))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Element](#extendableelement))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Element](#extendableelement))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Element](#extendableelement))* |
| `id` | `string` |  |  |  | A unique identifier associated with the item. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### Fact

Describes a Fact in a FactSet as a key/value pair.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Fact` |  |  |  | Must be `Fact` |
| `title` | `string` | ✓ |  |  | The title of the fact. |
| `value` | `string` | ✓ |  |  | The value of the fact. |

### FactSet

The FactSet element displays a series of facts (i.e. name/value pairs) in a tabular form.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `FactSet` |  |  |  | Must be `FactSet` |
| `facts` | list of [Fact](#fact) | ✓ |  |  | The array of `Fact`'s. |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Element](#extendableelement))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Element](#extendableelement))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Element](#extendableelement))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Element](#extendableelement))* |
| `id` | `string` |  |  |  | A unique identifier associated with the item. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### Image

Displays an image. Acceptable formats are PNG, JPEG, and GIF  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Image` |  |  |  | Must be `Image` |
| `url` | `string` | ✓ |  |  | The URL to the image. Supports data URI in version 1.2+ |
| `altText` | `string` |  |  |  | Alternate text describing the image. |
| `backgroundColor` | `string` |  |  | 1.1 | Applies a background to a transparent image. This property will respect the image style. |
| `height` | `string` \| [BlockElementHeight](#blockelementheight) |  | `"auto"` | 1.1 | The desired height of the image. If specified as a pixel value, ending in 'px', E.g., 50px, the image will distort to fit that exact height. This overrides the `size` property. |
| `horizontalAlignment` | [HorizontalAlignment](#horizontalalignment) |  |  |  | Controls how this element is horizontally positioned within its parent. When not specified, the value of horizontalAlignment is inherited from the parent container. If no parent container has horizontalAlignment set, it defaults to Left. |
| `selectAction` | [ImplementationsOf.ISelectAction](#implementationsofiselectaction) |  |  | 1.1 | An Action that will be invoked when the `Image` is tapped or selected. `Action.ShowCard` is not supported. |
| `size` | [ImageSize](#imagesize) |  |  |  | Controls the approximate size of the image. The physical dimensions will vary per host. |
| `style` | [ImageStyle](#imagestyle) |  |  |  | Controls how this `Image` is displayed. |
| `width` | `string` |  |  | 1.1 | The desired on-screen width of the image, ending in 'px'. E.g., 50px. This overrides the `size` property. |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. |
| `id` | `string` |  |  |  | A unique identifier associated with the item. |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. |

### ImageSet

The ImageSet displays a collection of Images similar to a gallery. Acceptable formats are PNG, JPEG, and GIF  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `ImageSet` |  |  |  | Must be `ImageSet` |
| `images` | list of [Image](#image) | ✓ |  |  | The array of `Image` elements to show. |
| `imageSize` | [ImageSize](#imagesize) |  | `"medium"` |  | Controls the approximate size of each image. The physical dimensions will vary per host. Auto and stretch are not supported for ImageSet. The size will default to medium if those values are set. |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Element](#extendableelement))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Element](#extendableelement))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Element](#extendableelement))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Element](#extendableelement))* |
| `id` | `string` |  |  |  | A unique identifier associated with the item. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### TextRun

Defines a single run of formatted text. A TextRun with no properties set can be represented in the json as string containing the text as a shorthand for the json object. These two representations are equivalent.  
*Version:* 1.2  

**One of:** `string`, `object` (case-insensitive)

### Media

Displays a media player for audio or video content.  
*Version:* 1.1  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Media` |  |  |  | Must be `Media` |
| `sources` | list of [MediaSource](#mediasource) | ✓ |  |  | Array of media sources to attempt to play. |
| `poster` | `string` |  |  |  | URL of an image to display before playing. Supports data URI in version 1.2+. If poster is omitted, the Media element will either use a default poster (controlled by the host application) or will attempt to automatically pull the poster from the target video service when the source URL points to a video from a Web provider such as YouTube. |
| `altText` | `string` |  |  |  | Alternate text describing the audio or video. |
| `captionSources` | list of [CaptionSource](#captionsource) |  |  | 1.6 | Array of captions sources for the media element to provide. |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Element](#extendableelement))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Element](#extendableelement))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Element](#extendableelement))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Element](#extendableelement))* |
| `id` | `string` |  |  |  | A unique identifier associated with the item. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### MediaSource

Defines a source for a Media element  
*Version:* 1.1  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `MediaSource` |  |  |  | Must be `MediaSource` |
| `mimeType` | `string` |  |  |  | Mime type of associated media (e.g. `"video/mp4"`). For YouTube and other Web video URLs, `mimeType` can be omitted. |
| `url` | `string` | ✓ |  |  | URL to media. Supports data URI in version 1.2+ |

### RichTextBlock

Defines an array of inlines, allowing for inline text formatting.  
*Version:* 1.2  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `RichTextBlock` |  |  |  | Must be `RichTextBlock` |
| `inlines` | list of [ImplementationsOf.Inline](#implementationsofinline) | ✓ |  |  | The array of inlines. |
| `horizontalAlignment` | [HorizontalAlignment](#horizontalalignment) |  |  |  | Controls the horizontal text alignment. When not specified, the value of horizontalAlignment is inherited from the parent container. If no parent container has horizontalAlignment set, it defaults to Left. |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Element](#extendableelement))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Element](#extendableelement))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Element](#extendableelement))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Element](#extendableelement))* |
| `id` | `string` |  |  |  | A unique identifier associated with the item. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### Table

Provides a way to display data in a tabular form.  
*Version:* 1.5  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Table` |  |  |  | Must be `Table` |
| `columns` | list of [TableColumnDefinition](#tablecolumndefinition) |  |  |  | Defines the number of columns in the table, their sizes, and more. |
| `rows` | list of [TableRow](#tablerow) |  |  |  | Defines the rows of the table. |
| `firstRowAsHeader` | `boolean` |  | `true` |  | Specifies whether the first row of the table should be treated as a header row, and be announced as such by accessibility software. |
| `showGridLines` | `boolean` |  | `true` |  | Specifies whether grid lines should be displayed. |
| `gridStyle` | [ContainerStyle](#containerstyle) |  | `"default"` |  | Defines the style of the grid. This property currently only controls the grid's color. |
| `horizontalCellContentAlignment` | [HorizontalAlignment](#horizontalalignment) |  |  |  | Controls how the content of all cells is horizontally aligned by default. When not specified, horizontal alignment is defined on a per-cell basis. |
| `verticalCellContentAlignment` | [VerticalAlignment](#verticalalignment) |  |  |  | Controls how the content of all cells is vertically aligned by default. When not specified, vertical alignment is defined on a per-cell basis. |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Element](#extendableelement))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Element](#extendableelement))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Element](#extendableelement))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Element](#extendableelement))* |
| `id` | `string` |  |  |  | A unique identifier associated with the item. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### TableCell

Represents a cell within a row of a Table element.  
*Version:* 1.5  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `TableCell` |  |  |  | Must be `TableCell` |
| `items` | list of [ImplementationsOf.Element](#implementationsofelement) | ✓ |  |  | The card elements to render inside the `TableCell`. |
| `selectAction` | [ImplementationsOf.ISelectAction](#implementationsofiselectaction) |  |  | 1.1 | An Action that will be invoked when the `TableCell` is tapped or selected. `Action.ShowCard` is not supported. |
| `style` | [ContainerStyle](#containerstyle) |  |  |  | Style hint for `TableCell`. |
| `verticalContentAlignment` | [VerticalContentAlignment](#verticalcontentalignment) |  |  | 1.1 | Defines how the content should be aligned vertically within the container. When not specified, the value of verticalContentAlignment is inherited from the parent container. If no parent container has verticalContentAlignment set, it defaults to Top. |
| `bleed` | `boolean` |  |  | 1.2 | Determines whether the element should bleed through its parent's padding. |
| `backgroundImage` | [BackgroundImage](#backgroundimage) \| `string` |  |  | 1.2 | Specifies the background image. Acceptable formats are PNG, JPEG, and GIF |
| `minHeight` | `string` |  |  | 1.2 | Specifies the minimum height of the container in pixels, like `"80px"`. |
| `rtl?` | `boolean` |  |  | 1.5 | When `true` content in this container should be presented right to left. When 'false' content in this container should be presented left to right. When unset layout direction will inherit from parent container or column. If unset in all ancestors, the default platform behavior will apply. |

### TableColumnDefinition

Defines the characteristics of a column in a Table element.  
*Version:* 1.5  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `TableColumnDefinition` |  |  |  | Must be `TableColumnDefinition` |
| `width` | `string` \| `number` |  | `1` |  | Specifies the width of the column. If expressed as a number, width represents the weight a the column relative to the other columns in the table. If expressed as a string, width must by in the format "<number>px" (for instance, "50px") and represents an explicit number of pixels. |
| `horizontalCellContentAlignment` | [HorizontalAlignment](#horizontalalignment) |  |  |  | Controls how the content of all cells in the column is horizontally aligned by default. When specified, this value overrides the setting at the table level. When not specified, horizontal alignment is defined at the table, row or cell level. |
| `verticalCellContentAlignment` | [VerticalAlignment](#verticalalignment) |  |  |  | Controls how the content of all cells in the column is vertically aligned by default. When specified, this value overrides the setting at the table level. When not specified, vertical alignment is defined at the table, row or cell level. |

### TableRow

Represents a row of cells within a Table element.  
*Version:* 1.5  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `TableRow` |  |  |  | Must be `TableRow` |
| `cells` | list of [TableCell](#tablecell) |  |  |  | The cells in this row. If a row contains more cells than there are columns defined on the Table element, the extra cells are ignored. |
| `style` | [ContainerStyle](#containerstyle) |  |  |  | Defines the style of the entire row. |
| `horizontalCellContentAlignment` | [HorizontalAlignment](#horizontalalignment) |  |  |  | Controls how the content of all cells in the row is horizontally aligned by default. When specified, this value overrides both the setting at the table and columns level. When not specified, horizontal alignment is defined at the table, column or cell level. |
| `verticalCellContentAlignment` | [VerticalAlignment](#verticalalignment) |  |  |  | Controls how the content of all cells in the column is vertically aligned by default. When specified, this value overrides the setting at the table and column level. When not specified, vertical alignment is defined either at the table, column or cell level. |

### TextBlock

Displays text, allowing control over font sizes, weight, and color.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `TextBlock` |  |  |  | Must be `TextBlock` |
| `text` | `string` | ✓ |  |  | Text to display. A subset of markdown is supported (https://aka.ms/ACTextFeatures) |
| `color` | [Colors](#colors) |  |  |  | Controls the color of `TextBlock` elements. |
| `fontType` | [FontType](#fonttype) |  |  | 1.2 | Type of font to use for rendering |
| `horizontalAlignment` | [HorizontalAlignment](#horizontalalignment) |  |  |  | Controls the horizontal text alignment. When not specified, the value of horizontalAlignment is inherited from the parent container. If no parent container has horizontalAlignment set, it defaults to Left. |
| `isSubtle` | `boolean` |  | `false` |  | If `true`, displays text slightly toned down to appear less prominent. |
| `maxLines` | `number` |  |  |  | Specifies the maximum number of lines to display. |
| `size` | [FontSize](#fontsize) |  |  |  | Controls size of text. |
| `weight` | [FontWeight](#fontweight) |  |  |  | Controls the weight of `TextBlock` elements. |
| `wrap` | `boolean` |  | `false` |  | If `true`, allow text to wrap. Otherwise, text is clipped. |
| `style` | [TextBlockStyle](#textblockstyle) |  | `"default"` |  | The style of this TextBlock for accessibility purposes. |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Element](#extendableelement))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Element](#extendableelement))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Element](#extendableelement))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Element](#extendableelement))* |
| `id` | `string` |  |  |  | A unique identifier associated with the item. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

## Inputs

### Data.Query

The data populated in the event payload for fetching dynamic choices, sent to the card-author to help identify the dataset from which choices might be fetched to be displayed in the dropdown. It might contain auxillary data to limit the maximum number of choices that can be sent and to support pagination.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Data.Query` |  |  |  | Must be `Data.Query` |
| `dataset` | `string` | ✓ |  | 1.6 | The dataset to be queried to get the choices. |
| `count` | `number` |  |  | 1.6 | The maximum number of choices that should be returned by the query. It can be ignored if the card-author wants to send a different number. |
| `skip` | `number` |  |  | 1.6 | The number of choices to be skipped in the list of choices returned by the query. It can be ignored if the card-author does not want pagination. |

### Input.Choice

Describes a choice for use in a ChoiceSet.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Input.Choice` |  |  |  | Must be `Input.Choice` |
| `title` | `string` | ✓ |  |  | Text to display. |
| `value` | `string` | ✓ |  |  | The raw value for the choice. **NOTE:** do not use a `,` in the value, since a `ChoiceSet` with `isMultiSelect` set to `true` returns a comma-delimited string of choice values. |

### Input.ChoiceSet

Allows a user to input a Choice.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Input.ChoiceSet` |  |  |  | Must be `Input.ChoiceSet` |
| `choices` | list of [Input.Choice](#inputchoice) |  |  |  | `Choice` options. |
| `choices.data` | [Data.Query](#dataquery) |  |  | 1.6 | Allows dynamic fetching of choices from the bot to be displayed as suggestions in the dropdown when the user types in the input field. |
| `isMultiSelect` | `boolean` |  | `false` |  | Allow multiple choices to be selected. |
| `style` | [ChoiceInputStyle](#choiceinputstyle) |  |  |  |  |
| `value` | `string` |  |  |  | The initial choice (or set of choices) that should be selected. For multi-select, specify a comma-separated string of values. |
| `placeholder` | `string` |  |  |  | Description of the input desired. Only visible when no selection has been made, the `style` is `compact` and `isMultiSelect` is `false` |
| `wrap` | `boolean` |  |  | 1.2 | If `true`, allow text to wrap. Otherwise, text is clipped. |
| `id` | `string` | ✓ |  |  | Unique identifier for the value. Used to identify collected input when the Submit action is performed. *(from [Extendable.Input](#extendableinput))* |
| `errorMessage` | `string` |  |  | 1.3 | Error message to display when entered input is invalid *(from [Extendable.Input](#extendableinput))* |
| `isRequired` | `boolean` |  |  | 1.3 | Whether or not this input is required *(from [Extendable.Input](#extendableinput))* |
| `label` | `string` |  |  | 1.3 | Label for this input *(from [Extendable.Input](#extendableinput))* |
| `labelPosition` | [InputLabelPosition](#inputlabelposition) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the position of the label. It can take 'inline' and 'above' values. By default, the label is placed 'above' when label position is not specified. *(from [Extendable.Input](#extendableinput))* |
| `labelWidth` | `string` \| `number` |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the width of the label in percent like 40 or a specific pixel width like '40px' when label is placed inline with the input. labelWidth would be ignored when the label is displayed above the input. *(from [Extendable.Input](#extendableinput))* |
| `inputStyle` | [InputStyle](#inputstyle) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Style hint for input fields. Allows input fields to appear as read-only but when user clicks/focuses on the field, it allows them to update those fields. *(from [Extendable.Input](#extendableinput))* |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Input](#extendableinput))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Input](#extendableinput))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Input](#extendableinput))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Input](#extendableinput))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.Input](#extendableinput))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Input](#extendableinput))* |

### Input.Date

Lets a user choose a date.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Input.Date` |  |  |  | Must be `Input.Date` |
| `max` | `string` |  |  |  | Hint of maximum value expressed in YYYY-MM-DD(may be ignored by some clients). |
| `min` | `string` |  |  |  | Hint of minimum value expressed in YYYY-MM-DD(may be ignored by some clients). |
| `placeholder` | `string` |  |  |  | Description of the input desired. Displayed when no selection has been made. |
| `value` | `string` |  |  |  | The initial value for this field expressed in YYYY-MM-DD. |
| `id` | `string` | ✓ |  |  | Unique identifier for the value. Used to identify collected input when the Submit action is performed. *(from [Extendable.Input](#extendableinput))* |
| `errorMessage` | `string` |  |  | 1.3 | Error message to display when entered input is invalid *(from [Extendable.Input](#extendableinput))* |
| `isRequired` | `boolean` |  |  | 1.3 | Whether or not this input is required *(from [Extendable.Input](#extendableinput))* |
| `label` | `string` |  |  | 1.3 | Label for this input *(from [Extendable.Input](#extendableinput))* |
| `labelPosition` | [InputLabelPosition](#inputlabelposition) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the position of the label. It can take 'inline' and 'above' values. By default, the label is placed 'above' when label position is not specified. *(from [Extendable.Input](#extendableinput))* |
| `labelWidth` | `string` \| `number` |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the width of the label in percent like 40 or a specific pixel width like '40px' when label is placed inline with the input. labelWidth would be ignored when the label is displayed above the input. *(from [Extendable.Input](#extendableinput))* |
| `inputStyle` | [InputStyle](#inputstyle) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Style hint for input fields. Allows input fields to appear as read-only but when user clicks/focuses on the field, it allows them to update those fields. *(from [Extendable.Input](#extendableinput))* |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Input](#extendableinput))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Input](#extendableinput))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Input](#extendableinput))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Input](#extendableinput))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.Input](#extendableinput))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Input](#extendableinput))* |

### Input.Number

Allows a user to enter a number.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Input.Number` |  |  |  | Must be `Input.Number` |
| `max` | `number` |  |  |  | Hint of maximum value (may be ignored by some clients). |
| `min` | `number` |  |  |  | Hint of minimum value (may be ignored by some clients). |
| `placeholder` | `string` |  |  |  | Description of the input desired. Displayed when no selection has been made. |
| `value` | `number` |  |  |  | Initial value for this field. |
| `id` | `string` | ✓ |  |  | Unique identifier for the value. Used to identify collected input when the Submit action is performed. *(from [Extendable.Input](#extendableinput))* |
| `errorMessage` | `string` |  |  | 1.3 | Error message to display when entered input is invalid *(from [Extendable.Input](#extendableinput))* |
| `isRequired` | `boolean` |  |  | 1.3 | Whether or not this input is required *(from [Extendable.Input](#extendableinput))* |
| `label` | `string` |  |  | 1.3 | Label for this input *(from [Extendable.Input](#extendableinput))* |
| `labelPosition` | [InputLabelPosition](#inputlabelposition) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the position of the label. It can take 'inline' and 'above' values. By default, the label is placed 'above' when label position is not specified. *(from [Extendable.Input](#extendableinput))* |
| `labelWidth` | `string` \| `number` |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the width of the label in percent like 40 or a specific pixel width like '40px' when label is placed inline with the input. labelWidth would be ignored when the label is displayed above the input. *(from [Extendable.Input](#extendableinput))* |
| `inputStyle` | [InputStyle](#inputstyle) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Style hint for input fields. Allows input fields to appear as read-only but when user clicks/focuses on the field, it allows them to update those fields. *(from [Extendable.Input](#extendableinput))* |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Input](#extendableinput))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Input](#extendableinput))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Input](#extendableinput))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Input](#extendableinput))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.Input](#extendableinput))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Input](#extendableinput))* |

### Input.Text

Lets a user enter text.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Input.Text` |  |  |  | Must be `Input.Text` |
| `isMultiline` | `boolean` |  | `false` |  | If `true`, allow multiple lines of input. |
| `maxLength` | `number` |  |  |  | Hint of maximum length characters to collect (may be ignored by some clients). |
| `placeholder` | `string` |  |  |  | Description of the input desired. Displayed when no text has been input. |
| `regex` | `string` |  |  | 1.3 | Regular expression indicating the required format of this text input. |
| `style` | [TextInputStyle](#textinputstyle) |  |  |  | Style hint for text input. |
| `inlineAction` | [ImplementationsOf.ISelectAction](#implementationsofiselectaction) |  |  | 1.2 | The inline action for the input. Typically displayed to the right of the input. It is strongly recommended to provide an icon on the action (which will be displayed instead of the title of the action). |
| `value` | `string` |  |  |  | The initial value for this field. |
| `id` | `string` | ✓ |  |  | Unique identifier for the value. Used to identify collected input when the Submit action is performed. *(from [Extendable.Input](#extendableinput))* |
| `errorMessage` | `string` |  |  | 1.3 | Error message to display when entered input is invalid *(from [Extendable.Input](#extendableinput))* |
| `isRequired` | `boolean` |  |  | 1.3 | Whether or not this input is required *(from [Extendable.Input](#extendableinput))* |
| `label` | `string` |  |  | 1.3 | Label for this input *(from [Extendable.Input](#extendableinput))* |
| `labelPosition` | [InputLabelPosition](#inputlabelposition) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the position of the label. It can take 'inline' and 'above' values. By default, the label is placed 'above' when label position is not specified. *(from [Extendable.Input](#extendableinput))* |
| `labelWidth` | `string` \| `number` |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the width of the label in percent like 40 or a specific pixel width like '40px' when label is placed inline with the input. labelWidth would be ignored when the label is displayed above the input. *(from [Extendable.Input](#extendableinput))* |
| `inputStyle` | [InputStyle](#inputstyle) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Style hint for input fields. Allows input fields to appear as read-only but when user clicks/focuses on the field, it allows them to update those fields. *(from [Extendable.Input](#extendableinput))* |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Input](#extendableinput))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Input](#extendableinput))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Input](#extendableinput))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Input](#extendableinput))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.Input](#extendableinput))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Input](#extendableinput))* |

### Input.Time

Lets a user select a time.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Input.Time` |  |  |  | Must be `Input.Time` |
| `max` | `string` |  |  |  | Hint of maximum value expressed in HH:MM (may be ignored by some clients). |
| `min` | `string` |  |  |  | Hint of minimum value expressed in HH:MM (may be ignored by some clients). |
| `placeholder` | `string` |  |  |  | Description of the input desired. Displayed when no time has been selected. |
| `value` | `string` |  |  |  | The initial value for this field expressed in HH:MM. |
| `id` | `string` | ✓ |  |  | Unique identifier for the value. Used to identify collected input when the Submit action is performed. *(from [Extendable.Input](#extendableinput))* |
| `errorMessage` | `string` |  |  | 1.3 | Error message to display when entered input is invalid *(from [Extendable.Input](#extendableinput))* |
| `isRequired` | `boolean` |  |  | 1.3 | Whether or not this input is required *(from [Extendable.Input](#extendableinput))* |
| `label` | `string` |  |  | 1.3 | Label for this input *(from [Extendable.Input](#extendableinput))* |
| `labelPosition` | [InputLabelPosition](#inputlabelposition) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the position of the label. It can take 'inline' and 'above' values. By default, the label is placed 'above' when label position is not specified. *(from [Extendable.Input](#extendableinput))* |
| `labelWidth` | `string` \| `number` |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the width of the label in percent like 40 or a specific pixel width like '40px' when label is placed inline with the input. labelWidth would be ignored when the label is displayed above the input. *(from [Extendable.Input](#extendableinput))* |
| `inputStyle` | [InputStyle](#inputstyle) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Style hint for input fields. Allows input fields to appear as read-only but when user clicks/focuses on the field, it allows them to update those fields. *(from [Extendable.Input](#extendableinput))* |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Input](#extendableinput))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Input](#extendableinput))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Input](#extendableinput))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Input](#extendableinput))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.Input](#extendableinput))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Input](#extendableinput))* |

### Input.Toggle

Lets a user choose between two options.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Input.Toggle` |  |  |  | Must be `Input.Toggle` |
| `title` | `string` | ✓ |  |  | Title for the toggle |
| `value` | `string` |  | `"false"` |  | The initial selected value. If you want the toggle to be initially on, set this to the value of `valueOn`'s value. |
| `valueOff` | `string` |  | `"false"` |  | The value when toggle is off |
| `valueOn` | `string` |  | `"true"` |  | The value when toggle is on |
| `wrap` | `boolean` |  |  | 1.2 | If `true`, allow text to wrap. Otherwise, text is clipped. |
| `id` | `string` | ✓ |  |  | Unique identifier for the value. Used to identify collected input when the Submit action is performed. *(from [Extendable.Input](#extendableinput))* |
| `errorMessage` | `string` |  |  | 1.3 | Error message to display when entered input is invalid *(from [Extendable.Input](#extendableinput))* |
| `isRequired` | `boolean` |  |  | 1.3 | Whether or not this input is required *(from [Extendable.Input](#extendableinput))* |
| `label` | `string` |  |  | 1.3 | Label for this input *(from [Extendable.Input](#extendableinput))* |
| `labelPosition` | [InputLabelPosition](#inputlabelposition) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the position of the label. It can take 'inline' and 'above' values. By default, the label is placed 'above' when label position is not specified. *(from [Extendable.Input](#extendableinput))* |
| `labelWidth` | `string` \| `number` |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the width of the label in percent like 40 or a specific pixel width like '40px' when label is placed inline with the input. labelWidth would be ignored when the label is displayed above the input. *(from [Extendable.Input](#extendableinput))* |
| `inputStyle` | [InputStyle](#inputstyle) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Style hint for input fields. Allows input fields to appear as read-only but when user clicks/focuses on the field, it allows them to update those fields. *(from [Extendable.Input](#extendableinput))* |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Input](#extendableinput))* |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. *(from [Extendable.Input](#extendableinput))* |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. *(from [Extendable.Input](#extendableinput))* |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. *(from [Extendable.Input](#extendableinput))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.Input](#extendableinput))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Input](#extendableinput))* |

## Actions (buttons)

### Action.Execute

Gathers input fields, merges with optional data field, and sends an event to the client. Clients process the event by sending an Invoke activity of type adaptiveCard/action to the target Bot. The inputs that are gathered are those on the current card, and in the case of a show card those on any parent cards. See [Universal Action Model](https://docs.microsoft.com/en-us/adaptive-cards/authoring-cards/universal-action-model) documentation for more details.  
*Version:* 1.4  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Action.Execute` |  |  |  | Must be `Action.Execute` |
| `verb` | `string` |  |  |  | The card author-defined verb associated with this action. |
| `data` | `string` \| `object` |  |  |  | Initial data that input fields will be combined with. These are essentially ‘hidden’ properties. |
| `associatedInputs` | [AssociatedInputs](#associatedinputs) |  | `"auto"` |  | Controls which inputs are associated with the action. |
| `title` | `string` |  |  |  | Label for button or link that represents this action. *(from [Extendable.Action](#extendableaction))* |
| `iconUrl` | `string` |  |  | 1.1 | Optional icon to be shown on the action in conjunction with the title. Supports data URI in version 1.2+ *(from [Extendable.Action](#extendableaction))* |
| `id` | `string` |  |  |  | A unique identifier associated with this Action. *(from [Extendable.Action](#extendableaction))* |
| `style` | [ActionStyle](#actionstyle) |  |  | 1.2 | Controls the style of an Action, which influences how the action is displayed, spoken, etc. *(from [Extendable.Action](#extendableaction))* |
| `fallback` | [ImplementationsOf.Action](#implementationsofaction) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Action](#extendableaction))* |
| `tooltip` | `string` |  |  | 1.5 | Defines text that should be displayed to the end user as they hover the mouse over the action, and read when using narration software. *(from [Extendable.Action](#extendableaction))* |
| `isEnabled` | `boolean` |  | `true` | 1.5 | Determines whether the action should be enabled. *(from [Extendable.Action](#extendableaction))* |
| `mode` | [ActionMode](#actionmode) |  | `"primary"` | 1.5 | Determines whether the action should be displayed as a button or in the overflow menu. *(from [Extendable.Action](#extendableaction))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### Action.OpenUrl

When invoked, show the given url either by launching it in an external web browser or showing within an embedded web browser.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Action.OpenUrl` |  |  |  | Must be `Action.OpenUrl` |
| `url` | `string` | ✓ |  |  | The URL to open. |
| `title` | `string` |  |  |  | Label for button or link that represents this action. *(from [Extendable.Action](#extendableaction))* |
| `iconUrl` | `string` |  |  | 1.1 | Optional icon to be shown on the action in conjunction with the title. Supports data URI in version 1.2+ *(from [Extendable.Action](#extendableaction))* |
| `id` | `string` |  |  |  | A unique identifier associated with this Action. *(from [Extendable.Action](#extendableaction))* |
| `style` | [ActionStyle](#actionstyle) |  |  | 1.2 | Controls the style of an Action, which influences how the action is displayed, spoken, etc. *(from [Extendable.Action](#extendableaction))* |
| `fallback` | [ImplementationsOf.Action](#implementationsofaction) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Action](#extendableaction))* |
| `tooltip` | `string` |  |  | 1.5 | Defines text that should be displayed to the end user as they hover the mouse over the action, and read when using narration software. *(from [Extendable.Action](#extendableaction))* |
| `isEnabled` | `boolean` |  | `true` | 1.5 | Determines whether the action should be enabled. *(from [Extendable.Action](#extendableaction))* |
| `mode` | [ActionMode](#actionmode) |  | `"primary"` | 1.5 | Determines whether the action should be displayed as a button or in the overflow menu. *(from [Extendable.Action](#extendableaction))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### Action.ShowCard

Defines an AdaptiveCard which is shown to the user when the button or link is clicked.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Action.ShowCard` |  |  |  | Must be `Action.ShowCard` |
| `card` | [AdaptiveCard](#adaptivecard) |  |  |  | The Adaptive Card to show. Inputs in ShowCards will not be submitted if the submit button is located on a parent card. See https://docs.microsoft.com/en-us/adaptive-cards/authoring-cards/input-validation for more details. |
| `title` | `string` |  |  |  | Label for button or link that represents this action. *(from [Extendable.Action](#extendableaction))* |
| `iconUrl` | `string` |  |  | 1.1 | Optional icon to be shown on the action in conjunction with the title. Supports data URI in version 1.2+ *(from [Extendable.Action](#extendableaction))* |
| `id` | `string` |  |  |  | A unique identifier associated with this Action. *(from [Extendable.Action](#extendableaction))* |
| `style` | [ActionStyle](#actionstyle) |  |  | 1.2 | Controls the style of an Action, which influences how the action is displayed, spoken, etc. *(from [Extendable.Action](#extendableaction))* |
| `fallback` | [ImplementationsOf.Action](#implementationsofaction) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Action](#extendableaction))* |
| `tooltip` | `string` |  |  | 1.5 | Defines text that should be displayed to the end user as they hover the mouse over the action, and read when using narration software. *(from [Extendable.Action](#extendableaction))* |
| `isEnabled` | `boolean` |  | `true` | 1.5 | Determines whether the action should be enabled. *(from [Extendable.Action](#extendableaction))* |
| `mode` | [ActionMode](#actionmode) |  | `"primary"` | 1.5 | Determines whether the action should be displayed as a button or in the overflow menu. *(from [Extendable.Action](#extendableaction))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### Action.Submit

Gathers input fields, merges with optional data field, and sends an event to the client. It is up to the client to determine how this data is processed. For example: With BotFramework bots, the client would send an activity through the messaging medium to the bot. The inputs that are gathered are those on the current card, and in the case of a show card those on any parent cards. See https://docs.microsoft.com/en-us/adaptive-cards/authoring-cards/input-validation for more details.  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Action.Submit` |  |  |  | Must be `Action.Submit` |
| `data` | `string` \| `object` |  |  |  | Initial data that input fields will be combined with. These are essentially ‘hidden’ properties. |
| `associatedInputs` | [AssociatedInputs](#associatedinputs) |  | `"auto"` | 1.3 | Controls which inputs are associated with the submit action. |
| `title` | `string` |  |  |  | Label for button or link that represents this action. *(from [Extendable.Action](#extendableaction))* |
| `iconUrl` | `string` |  |  | 1.1 | Optional icon to be shown on the action in conjunction with the title. Supports data URI in version 1.2+ *(from [Extendable.Action](#extendableaction))* |
| `id` | `string` |  |  |  | A unique identifier associated with this Action. *(from [Extendable.Action](#extendableaction))* |
| `style` | [ActionStyle](#actionstyle) |  |  | 1.2 | Controls the style of an Action, which influences how the action is displayed, spoken, etc. *(from [Extendable.Action](#extendableaction))* |
| `fallback` | [ImplementationsOf.Action](#implementationsofaction) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Action](#extendableaction))* |
| `tooltip` | `string` |  |  | 1.5 | Defines text that should be displayed to the end user as they hover the mouse over the action, and read when using narration software. *(from [Extendable.Action](#extendableaction))* |
| `isEnabled` | `boolean` |  | `true` | 1.5 | Determines whether the action should be enabled. *(from [Extendable.Action](#extendableaction))* |
| `mode` | [ActionMode](#actionmode) |  | `"primary"` | 1.5 | Determines whether the action should be displayed as a button or in the overflow menu. *(from [Extendable.Action](#extendableaction))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### Action.ToggleVisibility

An action that toggles the visibility of associated card elements.  
*Version:* 1.2  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Action.ToggleVisibility` |  |  |  | Must be `Action.ToggleVisibility` |
| `targetElements` | list of [TargetElement](#targetelement) | ✓ |  |  | The array of TargetElements. It is not recommended to include Input elements with validation under Action.Toggle due to confusion that can arise from invalid inputs that are not currently visible. See https://docs.microsoft.com/en-us/adaptive-cards/authoring-cards/input-validation for more information. |
| `title` | `string` |  |  |  | Label for button or link that represents this action. *(from [Extendable.Action](#extendableaction))* |
| `iconUrl` | `string` |  |  | 1.1 | Optional icon to be shown on the action in conjunction with the title. Supports data URI in version 1.2+ *(from [Extendable.Action](#extendableaction))* |
| `id` | `string` |  |  |  | A unique identifier associated with this Action. *(from [Extendable.Action](#extendableaction))* |
| `style` | [ActionStyle](#actionstyle) |  |  | 1.2 | Controls the style of an Action, which influences how the action is displayed, spoken, etc. *(from [Extendable.Action](#extendableaction))* |
| `fallback` | [ImplementationsOf.Action](#implementationsofaction) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. *(from [Extendable.Action](#extendableaction))* |
| `tooltip` | `string` |  |  | 1.5 | Defines text that should be displayed to the end user as they hover the mouse over the action, and read when using narration software. *(from [Extendable.Action](#extendableaction))* |
| `isEnabled` | `boolean` |  | `true` | 1.5 | Determines whether the action should be enabled. *(from [Extendable.Action](#extendableaction))* |
| `mode` | [ActionMode](#actionmode) |  | `"primary"` | 1.5 | Determines whether the action should be displayed as a button or in the overflow menu. *(from [Extendable.Action](#extendableaction))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### TargetElement

Represents an entry for Action.ToggleVisibility's targetElements property  

**One of:** `string`, `object` (case-insensitive)

## Enumerations & style values

### ActionMode

Determines whether an action is displayed with a button or is moved to the overflow menu.  
*Version:* 1.5  

**One of:** `primary` \| `secondary` (case-insensitive)

### ActionStyle

Controls the style of an Action, which influences how the action is displayed, spoken, etc.  
*Version:* 1.2  

**One of:** `default` \| `positive` \| `destructive` (case-insensitive)

### AssociatedInputs


**One of:** `Auto` \| `None` (case-insensitive)

### BlockElementHeight


**One of:** `auto` \| `stretch` (case-insensitive)

### ChoiceInputStyle

Style hint for `Input.ChoiceSet`.  

**One of:** `compact` \| `expanded` \| `filtered` (case-insensitive)

### Colors


**One of:** `default` \| `dark` \| `light` \| `accent` \| `good` \| `warning` \| `attention` (case-insensitive)

### ContainerStyle


**One of:** `default` \| `emphasis` \| `good` \| `attention` \| `warning` \| `accent` (case-insensitive)

### FallbackOption


**One of:** `drop` (case-insensitive)

### FontSize


**One of:** `default` \| `small` \| `medium` \| `large` \| `extraLarge` (case-insensitive)

### FontType


**One of:** `default` \| `monospace` (case-insensitive)

### FontWeight


**One of:** `default` \| `lighter` \| `bolder` (case-insensitive)

### HorizontalAlignment

Controls how content is horizontally positioned within its container.  

**One of:** `left` \| `center` \| `right` (case-insensitive)

### ImageFillMode


**One of:** `cover` \| `repeatHorizontally` \| `repeatVertically` \| `repeat` (case-insensitive)

### ImageSize

Controls the approximate size of the image. The physical dimensions will vary per host. Every option preserves aspect ratio.  

**One of:** `auto` \| `stretch` \| `small` \| `medium` \| `large` (case-insensitive)

### ImageStyle

Controls how this `Image` is displayed.  

**One of:** `default` \| `person` (case-insensitive)

### InputLabelPosition

Specifies position of the label with respect to input.  

**One of:** `inline` \| `above` (case-insensitive)

### InputStyle

`revealOnHover`: Show input fields in read only view unless user takes any action. `default`: Show the default editable state for input controls  

**One of:** `revealOnHover` \| `default` (case-insensitive)

### Spacing

Specifies how much spacing. Hosts pick the exact pixel amounts for each of these.  

**One of:** `default` \| `none` \| `small` \| `medium` \| `large` \| `extraLarge` \| `padding` (case-insensitive)

### TextBlockStyle

Controls how a TextBlock behaves.  
*Version:* 1.5  

**One of:** `default` \| `heading` (case-insensitive)

### TextInputStyle

Style hint for text input.  

**One of:** `text` \| `tel` \| `url` \| `email` \| `password` (case-insensitive)

### VerticalAlignment


**One of:** `top` \| `center` \| `bottom` (case-insensitive)

### VerticalContentAlignment


**One of:** `top` \| `center` \| `bottom` (case-insensitive)

## Card-level extras

### AuthCardButton

Defines a button as displayed when prompting a user to authenticate. This maps to the cardAction type defined by the Bot Framework (https://docs.microsoft.com/dotnet/api/microsoft.bot.schema.cardaction).  
*Version:* 1.4  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `string` | ✓ |  |  | The type of the button. |
| `title` | `string` |  |  |  | The caption of the button. |
| `image` | `string` |  |  |  | A URL to an image to display alongside the button's caption. |
| `value` | `string` | ✓ |  |  | The value associated with the button. The meaning of value depends on the button's type. |

### Authentication

Defines authentication information associated with a card. This maps to the OAuthCard type defined by the Bot Framework (https://docs.microsoft.com/dotnet/api/microsoft.bot.schema.oauthcard)  
*Version:* 1.4  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Authentication` |  |  |  | Must be `Authentication` |
| `text` | `string` |  |  |  | Text that can be displayed to the end user when prompting them to authenticate. |
| `connectionName` | `string` |  |  |  | The identifier for registered OAuth connection setting information. |
| `tokenExchangeResource` | [TokenExchangeResource](#tokenexchangeresource) |  |  |  | Provides information required to enable on-behalf-of single sign-on user authentication. |
| `buttons` | list of [AuthCardButton](#authcardbutton) |  |  |  | Buttons that should be displayed to the user when prompting for authentication. The array MUST contain one button of type "signin". Other button types are not currently supported. |

### BackgroundImage

Specifies a background image. Acceptable formats are PNG, JPEG, and GIF  
*Version:* 1.2  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `BackgroundImage` |  |  |  | Must be `BackgroundImage` |
| `url` | `string` | ✓ |  |  | The URL (or data url) of the image. Acceptable formats are PNG, JPEG, and GIF |
| `fillMode` | [ImageFillMode](#imagefillmode) |  |  |  | Describes how the image should fill the area. |
| `horizontalAlignment` | [HorizontalAlignment](#horizontalalignment) |  |  |  | Describes how the image should be aligned if it must be cropped or if using repeat fill mode. |
| `verticalAlignment` | [VerticalAlignment](#verticalalignment) |  |  |  | Describes how the image should be aligned if it must be cropped or if using repeat fill mode. |

### Metadata

Defines various metadata properties  
*Version:* 1.6  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Metadata` |  |  |  | Must be `Metadata` |
| `webUrl` | `string` |  |  |  | URL that uniquely identifies the card and serves as a browser fallback that can be used by some hosts. |

### Refresh

Defines how a card can be refreshed by making a request to the target Bot.  
*Version:* 1.4  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `Refresh` |  |  |  | Must be `Refresh` |
| `action` | [Action.Execute](#actionexecute) |  |  |  | The action to be executed to refresh the card. Clients can run this refresh action automatically or can provide an affordance for users to trigger it manually. |
| `expires` | `string` |  |  | 1.6 | A timestamp that informs a Host when the card content has expired, and that it should trigger a refresh as appropriate. The format is ISO-8601 Instant format. E.g., 2022-01-01T12:00:00Z |
| `userIds` | list of `string` |  |  |  | A list of user Ids informing the client for which users should the refresh action should be run automatically. Some clients will not run the refresh action automatically unless this property is specified. Some clients may ignore this property and always run the refresh action automatically. |

### TokenExchangeResource

Defines information required to enable on-behalf-of single sign-on user authentication. Maps to the TokenExchangeResource type defined by the Bot Framework (https://docs.microsoft.com/dotnet/api/microsoft.bot.schema.tokenexchangeresource)  
*Version:* 1.4  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `type` | `TokenExchangeResource` |  |  |  | Must be `TokenExchangeResource` |
| `id` | `string` | ✓ |  |  | The unique identified of this token exchange instance. |
| `uri` | `string` | ✓ |  |  | An application ID or resource identifier with which to exchange a token on behalf of. This property is identity provider- and application-specific. |
| `providerId` | `string` | ✓ |  |  | An identifier for the identity provider with which to attempt a token exchange. |

## Abstract groups ("any element", "any action", …)

### ImplementationsOf.Action


**One of:** [Action.Execute](#actionexecute), [Action.OpenUrl](#actionopenurl), [Action.ShowCard](#actionshowcard), [Action.Submit](#actionsubmit), [Action.ToggleVisibility](#actiontogglevisibility) (case-insensitive)

### ImplementationsOf.Item


**One of:** [Action.Execute](#actionexecute), [Action.OpenUrl](#actionopenurl), [Action.ShowCard](#actionshowcard), [Action.Submit](#actionsubmit), [Action.ToggleVisibility](#actiontogglevisibility), [ActionSet](#actionset), [Column](#column), [ColumnSet](#columnset), [Container](#container), [FactSet](#factset), [Image](#image), [ImageSet](#imageset), [Input.ChoiceSet](#inputchoiceset), [Input.Date](#inputdate), [Input.Number](#inputnumber), [Input.Text](#inputtext), [Input.Time](#inputtime), [Input.Toggle](#inputtoggle), [Media](#media), [RichTextBlock](#richtextblock), [Table](#table), [TextBlock](#textblock) (case-insensitive)

### ImplementationsOf.ISelectAction


**One of:** [Action.Execute](#actionexecute), [Action.OpenUrl](#actionopenurl), [Action.Submit](#actionsubmit), [Action.ToggleVisibility](#actiontogglevisibility) (case-insensitive)

### ImplementationsOf.Element


**One of:** [ActionSet](#actionset), [ColumnSet](#columnset), [Container](#container), [FactSet](#factset), [Image](#image), [ImageSet](#imageset), [Input.ChoiceSet](#inputchoiceset), [Input.Date](#inputdate), [Input.Number](#inputnumber), [Input.Text](#inputtext), [Input.Time](#inputtime), [Input.Toggle](#inputtoggle), [Media](#media), [RichTextBlock](#richtextblock), [Table](#table), [TextBlock](#textblock) (case-insensitive)

### ImplementationsOf.ToggleableItem


**One of:** [ActionSet](#actionset), [Column](#column), [ColumnSet](#columnset), [Container](#container), [FactSet](#factset), [Image](#image), [ImageSet](#imageset), [Input.ChoiceSet](#inputchoiceset), [Input.Date](#inputdate), [Input.Number](#inputnumber), [Input.Text](#inputtext), [Input.Time](#inputtime), [Input.Toggle](#inputtoggle), [Media](#media), [RichTextBlock](#richtextblock), [Table](#table), [TextBlock](#textblock) (case-insensitive)

### ImplementationsOf.Inline


**One of:** [TextRun](#textrun) (case-insensitive)

### ImplementationsOf.Input


**One of:** [Input.ChoiceSet](#inputchoiceset), [Input.Date](#inputdate), [Input.Number](#inputnumber), [Input.Text](#inputtext), [Input.Time](#inputtime), [Input.Toggle](#inputtoggle) (case-insensitive)

### Extendable.Action


| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `title` | `string` |  |  |  | Label for button or link that represents this action. |
| `iconUrl` | `string` |  |  | 1.1 | Optional icon to be shown on the action in conjunction with the title. Supports data URI in version 1.2+ |
| `id` | `string` |  |  |  | A unique identifier associated with this Action. |
| `style` | [ActionStyle](#actionstyle) |  |  | 1.2 | Controls the style of an Action, which influences how the action is displayed, spoken, etc. |
| `fallback` | [ImplementationsOf.Action](#implementationsofaction) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. |
| `tooltip` | `string` |  |  | 1.5 | Defines text that should be displayed to the end user as they hover the mouse over the action, and read when using narration software. |
| `isEnabled` | `boolean` |  | `true` | 1.5 | Determines whether the action should be enabled. |
| `mode` | [ActionMode](#actionmode) |  | `"primary"` | 1.5 | Determines whether the action should be displayed as a button or in the overflow menu. |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### Extendable.Element


| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. |
| `id` | `string` |  |  |  | A unique identifier associated with the item. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. *(from [Extendable.ToggleableItem](#extendabletoggleableitem))* |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |

### Extendable.Input

Base input class  

| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `id` | `string` | ✓ |  |  | Unique identifier for the value. Used to identify collected input when the Submit action is performed. |
| `errorMessage` | `string` |  |  | 1.3 | Error message to display when entered input is invalid |
| `isRequired` | `boolean` |  |  | 1.3 | Whether or not this input is required |
| `label` | `string` |  |  | 1.3 | Label for this input |
| `labelPosition` | [InputLabelPosition](#inputlabelposition) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the position of the label. It can take 'inline' and 'above' values. By default, the label is placed 'above' when label position is not specified. |
| `labelWidth` | `string` \| `number` |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Determines the width of the label in percent like 40 or a specific pixel width like '40px' when label is placed inline with the input. labelWidth would be ignored when the label is displayed above the input. |
| `inputStyle` | [InputStyle](#inputstyle) |  |  | 1.6 | [SUPPORTED ONLY IN JAVASCRIPT SDK] Style hint for input fields. Allows input fields to appear as read-only but when user clicks/focuses on the field, it allows them to update those fields. |
| `fallback` | [ImplementationsOf.Element](#implementationsofelement) \| [FallbackOption](#fallbackoption) |  |  | 1.2 | Describes what to do when an unknown element is encountered or the requires of this or any children can't be met. |
| `height` | [BlockElementHeight](#blockelementheight) |  |  | 1.1 | Specifies the height of the element. |
| `separator` | `boolean` |  |  |  | When `true`, draw a separating line at the top of the element. |
| `spacing` | [Spacing](#spacing) |  |  |  | Controls the amount of spacing between this element and the preceding element. |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. |

### Extendable.Item


| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. |

### Extendable.ToggleableItem


| Property | Type | Req. | Default | Version | Description |
|---|---|---|---|---|---|
| `id` | `string` |  |  |  | A unique identifier associated with the item. |
| `isVisible` | `boolean` |  | `true` | 1.2 | If `false`, this item will be removed from the visual tree. |
| `requires` | `object` |  |  | 1.2 | A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered. *(from [Extendable.Item](#extendableitem))* |
