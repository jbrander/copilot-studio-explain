# Entities

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

Entities tell a `Question` node (and slot filling) what kind of value to extract. See [06 Variables, entities & Power Fx](../../06-variables-entities-powerfx.md).

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

**51 definitions** on this page (7 with a `kind:`). [A–Z index of all definitions](index.md).

<details><summary>Contents</summary>

[AgePrebuiltEntity](#ageprebuiltentity), [AmbiguousDateTimeResolutionBehavior](#ambiguousdatetimeresolutionbehavior), [BooleanPrebuiltEntity](#booleanprebuiltentity), [CityPrebuiltEntity](#cityprebuiltentity), [ClosedListConversionMetadata](#closedlistconversionmetadata), [ClosedListEntityReference](#closedlistentityreference), [ClosedListItem](#closedlistitem), [ColorPrebuiltEntity](#colorprebuiltentity), [ContinentPrebuiltEntity](#continentprebuiltentity), [CopilotConnectorEntity](#copilotconnectorentity), [CountryOrRegionPrebuiltEntity](#countryorregionprebuiltentity), [DatePrebuiltEntity](#dateprebuiltentity), [DateTimeNoTimeZonePrebuiltEntity](#datetimenotimezoneprebuiltentity), [DateTimePrebuiltEntity](#datetimeprebuiltentity), [DurationPrebuiltEntity](#durationprebuiltentity), [DynamicClosedListEntity](#dynamicclosedlistentity), [DynamicClosedListItem](#dynamicclosedlistitem), [DynamicClosedListItemNoKind](#dynamicclosedlistitemnokind), [EmailPrebuiltEntity](#emailprebuiltentity), [EmbeddedEntity](#embeddedentity), [Entity](#entity), [EntityReference](#entityreference), [EventPrebuiltEntity](#eventprebuiltentity), [ExternalEntityReference](#externalentityreference), [FilePrebuiltEntity](#fileprebuiltentity), [LanguagePrebuiltEntity](#languageprebuiltentity), [MoneyPrebuiltEntity](#moneyprebuiltentity), [MultiEntityChoice](#multientitychoice), [MultiEntityChoiceNoKind](#multientitychoicenokind), [NumberPrebuiltEntity](#numberprebuiltentity), [OneOfMultiEntityReference](#oneofmultientityreference), [OrdinalPrebuiltEntity](#ordinalprebuiltentity), [OrganizationPrebuiltEntity](#organizationprebuiltentity), [PercentagePrebuiltEntity](#percentageprebuiltentity), [PersonNamePrebuiltEntity](#personnameprebuiltentity), [PhoneNumberPrebuiltEntity](#phonenumberprebuiltentity), [PointOfInterestPrebuiltEntity](#pointofinterestprebuiltentity), [RegexEntityReference](#regexentityreference), [RegexIntentPattern](#regexintentpattern), [SelectedClosedListItem](#selectedclosedlistitem), [SelectedClosedListItemNoKind](#selectedclosedlistitemnokind), [SelectedClosedListItemOptionalKind](#selectedclosedlistitemoptionalkind), [SpeechRecognitionFormatType](#speechrecognitionformattype), [SpeedPrebuiltEntity](#speedprebuiltentity), [StatePrebuiltEntity](#stateprebuiltentity), [StreetAddressPrebuiltEntity](#streetaddressprebuiltentity), [StringPrebuiltEntity](#stringprebuiltentity), [TemperaturePrebuiltEntity](#temperatureprebuiltentity), [URLPrebuiltEntity](#urlprebuiltentity), [WeightPrebuiltEntity](#weightprebuiltentity), [ZipCodePrebuiltEntity](#zipcodeprebuiltentity)

</details>

### AgePrebuiltEntity


**One of:** `"AgePrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### AmbiguousDateTimeResolutionBehavior


**String enum.** Allowed values: `PromptForDisambiguation`, `PreferFuture`, `PreferPast`

<sub>Used by: [DatePrebuiltEntity](#dateprebuiltentity), [DateTimeNoTimeZonePrebuiltEntity](#datetimenotimezoneprebuiltentity), [DateTimePrebuiltEntity](#datetimeprebuiltentity)</sub>

### BooleanPrebuiltEntity


**One of:** `"BooleanPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### CityPrebuiltEntity


**One of:** `"CityPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### ClosedListConversionMetadata

YAML: `kind: ClosedListConversionMetadata`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `legacyIdToDisplayNameMapping` | map of name → `string` |  |  |  |

_No other properties are allowed._

### ClosedListEntityReference

YAML: `kind: ClosedListEntityReference`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sensitivityLevel` | [SensitivityLevel](settings.md#sensitivitylevel) |  | `"None"` |  |
| `includeMetadata` | `boolean` |  | `false` |  |
| `allowMultipleValues` | `boolean` |  | `false` |  |
| `entityId` | `string` | ✓ |  |  |
| `selectedItems` | list of [SelectedClosedListItemNoKind](#selectedclosedlistitemnokind) |  |  |  |
| `dtmfMultipleChoiceOptions` | [DtmfMultipleChoiceOptionsNoKind](settings.md#dtmfmultiplechoiceoptionsnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [EntityReference](#entityreference)</sub>

### ClosedListItem

YAML: `kind: ClosedListItem`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `displayName` | `string` | ✓ |  |  |
| `synonyms` | list of `string` |  |  |  |
| `dtmfKey` | [DtmfMappingKey](settings.md#dtmfmappingkey) |  |  |  |

_No other properties are allowed._

### ColorPrebuiltEntity


**One of:** `"ColorPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### ContinentPrebuiltEntity


**One of:** `"ContinentPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### CopilotConnectorEntity

YAML: `kind: CopilotConnectorEntity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `entityName` | `string` | ✓ |  |  |

_No other properties are allowed._

### CountryOrRegionPrebuiltEntity


**One of:** `"CountryOrRegionPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### DatePrebuiltEntity


**One of:** `"DatePrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### DateTimeNoTimeZonePrebuiltEntity


**One of:** `"DateTimeNoTimeZonePrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### DateTimePrebuiltEntity


**One of:** `"DateTimePrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### DurationPrebuiltEntity


**One of:** `"DurationPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### DynamicClosedListEntity


**One of:** object

<sub>Used by: [EntityReference](#entityreference)</sub>

### DynamicClosedListItem

YAML: `kind: DynamicClosedListItem`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

### DynamicClosedListItemNoKind


**Type:** object

### EmailPrebuiltEntity


**One of:** `"EmailPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### EmbeddedEntity


**One of:** object

<sub>Used by: [EntityReference](#entityreference)</sub>

### Entity


**One of:** [RegexEntity](dialogs.md#regexentity), [ClosedListEntity](dialogs.md#closedlistentity), [ExternalEntity](dialogs.md#externalentity)

<sub>Used by: [CustomEntityComponent](components.md#customentitycomponent), [EmbeddedEntity](#embeddedentity)</sub>

### EntityReference


**One of:** [OneOfMultiEntityReference](#oneofmultientityreference), [RegexEntityReference](#regexentityreference), [ClosedListEntityReference](#closedlistentityreference), [DynamicClosedListEntity](#dynamicclosedlistentity), [BooleanPrebuiltEntity](#booleanprebuiltentity), [StringPrebuiltEntity](#stringprebuiltentity), [AgePrebuiltEntity](#ageprebuiltentity), [DatePrebuiltEntity](#dateprebuiltentity), [MoneyPrebuiltEntity](#moneyprebuiltentity), [NumberPrebuiltEntity](#numberprebuiltentity), [OrdinalPrebuiltEntity](#ordinalprebuiltentity), [PercentagePrebuiltEntity](#percentageprebuiltentity), [PhoneNumberPrebuiltEntity](#phonenumberprebuiltentity), [SpeedPrebuiltEntity](#speedprebuiltentity), [TemperaturePrebuiltEntity](#temperatureprebuiltentity), [WeightPrebuiltEntity](#weightprebuiltentity), [ZipCodePrebuiltEntity](#zipcodeprebuiltentity), [CityPrebuiltEntity](#cityprebuiltentity), [ColorPrebuiltEntity](#colorprebuiltentity), [ContinentPrebuiltEntity](#continentprebuiltentity), [CountryOrRegionPrebuiltEntity](#countryorregionprebuiltentity), [DateTimePrebuiltEntity](#datetimeprebuiltentity), [DateTimeNoTimeZonePrebuiltEntity](#datetimenotimezoneprebuiltentity), [DurationPrebuiltEntity](#durationprebuiltentity), [EmailPrebuiltEntity](#emailprebuiltentity), [EventPrebuiltEntity](#eventprebuiltentity), [FilePrebuiltEntity](#fileprebuiltentity), [LanguagePrebuiltEntity](#languageprebuiltentity), [OrganizationPrebuiltEntity](#organizationprebuiltentity), [PersonNamePrebuiltEntity](#personnameprebuiltentity), [PointOfInterestPrebuiltEntity](#pointofinterestprebuiltentity), [StatePrebuiltEntity](#stateprebuiltentity), [StreetAddressPrebuiltEntity](#streetaddressprebuiltentity), [URLPrebuiltEntity](#urlprebuiltentity), [EmbeddedEntity](#embeddedentity), [ExternalEntityReference](#externalentityreference)

<sub>Used by: [AutomaticTaskInput](task-actions.md#automatictaskinput), [MultiEntityChoice](#multientitychoice), [MultiEntityChoiceNoKind](#multientitychoicenokind), [Question](actions.md#question)</sub>

### EventPrebuiltEntity


**One of:** `"EventPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### ExternalEntityReference


**One of:** object

<sub>Used by: [EntityReference](#entityreference)</sub>

### FilePrebuiltEntity


**One of:** `"FilePrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### LanguagePrebuiltEntity


**One of:** `"LanguagePrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### MoneyPrebuiltEntity


**One of:** `"MoneyPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### MultiEntityChoice

YAML: `kind: MultiEntityChoice`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `entity` | [EntityReference](#entityreference) | ✓ |  |  |
| `propertyName` | `string` | ✓ |  |  |

_No other properties are allowed._

### MultiEntityChoiceNoKind

Same shape as [MultiEntityChoice](#multientitychoice) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `entity` | [EntityReference](#entityreference) | ✓ |  |  |
| `propertyName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [OneOfMultiEntityReference](#oneofmultientityreference)</sub>

### NumberPrebuiltEntity


**One of:** `"NumberPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### OneOfMultiEntityReference


**One of:** object

<sub>Used by: [EntityReference](#entityreference)</sub>

### OrdinalPrebuiltEntity


**One of:** `"OrdinalPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### OrganizationPrebuiltEntity


**One of:** `"OrganizationPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### PercentagePrebuiltEntity


**One of:** `"PercentagePrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### PersonNamePrebuiltEntity


**One of:** `"PersonNamePrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### PhoneNumberPrebuiltEntity


**One of:** `"PhoneNumberPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### PointOfInterestPrebuiltEntity


**One of:** `"PointOfInterestPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### RegexEntityReference


**One of:** object

<sub>Used by: [EntityReference](#entityreference)</sub>

### RegexIntentPattern

YAML: `kind: RegexIntentPattern`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `intent` | `string` | ✓ |  |  |
| `pattern` | `string` | ✓ |  |  |

_No other properties are allowed._

### SelectedClosedListItem


**One of:** `string`, object

### SelectedClosedListItemNoKind


**One of:** `string`, object

<sub>Used by: [ClosedListEntityReference](#closedlistentityreference)</sub>

### SelectedClosedListItemOptionalKind


**One of:** `string`, object

### SpeechRecognitionFormatType


**String enum.** Allowed values: `Default`, `NumberAsDigits`

<sub>Used by: [StringPrebuiltEntity](#stringprebuiltentity)</sub>

### SpeedPrebuiltEntity


**One of:** `"SpeedPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### StatePrebuiltEntity


**One of:** `"StatePrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### StreetAddressPrebuiltEntity


**One of:** `"StreetAddressPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### StringPrebuiltEntity


**One of:** `"StringPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### TemperaturePrebuiltEntity


**One of:** `"TemperaturePrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### URLPrebuiltEntity


**One of:** `"URLPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### WeightPrebuiltEntity


**One of:** `"WeightPrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>

### ZipCodePrebuiltEntity


**One of:** `"ZipCodePrebuiltEntity"`, object

<sub>Used by: [EntityReference](#entityreference)</sub>
