# Knowledge sources

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

The `source:` of a `KnowledgeSourceConfiguration` and related search settings. See [07 Knowledge](../../07-knowledge.md).

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

**62 definitions** on this page (39 with a `kind:`). [A–Z index of all definitions](index.md).

<details><summary>Contents</summary>

[AzureAISearchSource](#azureaisearchsource), [AzureCognitiveSearchInternalDataSource](#azurecognitivesearchinternaldatasource), [AzureCognitiveSearchInternalDataSourceNoKind](#azurecognitivesearchinternaldatasourcenokind), [AzureOpenAIOnYourDataKnowledgeSource](#azureopenaionyourdataknowledgesource), [AzureOpenAIOnYourDataSource](#azureopenaionyourdatasource), [BaseSearchConfiguration](#basesearchconfiguration), [BingCustomSearchDataSource](#bingcustomsearchdatasource), [BingCustomSearchSource](#bingcustomsearchsource), [BingSearchOptions](#bingsearchoptions), [Citation](#citation), [CitationNoKind](#citationnokind), [CitationsInModelOutput](#citationsinmodeloutput), [CopilotConnectorEntityNoKind](#copilotconnectorentitynokind), [CopilotConnectorKnowledgeSourceDemo](#copilotconnectorknowledgesourcedemo), [CustomDataSource](#customdatasource), [CustomKnowledgeSource](#customknowledgesource), [CustomModelCitationHandlingType](#custommodelcitationhandlingtype), [CustomSearchInputContent](#customsearchinputcontent), [CustomSearchInputContentNoKind](#customsearchinputcontentnokind), [DataverseStructuredSearchSource](#dataversestructuredsearchsource), [DataverseTableSearchBase](#dataversetablesearchbase), [EmailSearchSource](#emailsearchsource), [FabricAISkillSource](#fabricaiskillsource), [FederatedStructuredSearchSource](#federatedstructuredsearchsource), [FileGroupKnowledgeSource](#filegroupknowledgesource), [FileIndexingStatus](#fileindexingstatus), [FileSearchDataSource](#filesearchdatasource), [FileSearchDataSourceNoKind](#filesearchdatasourcenokind), [GenericDatasetNameSearchConfiguration](#genericdatasetnamesearchconfiguration), [GraphConnectorItemByExternalId](#graphconnectoritembyexternalid), [GraphConnectorItemByExternalIdNoKind](#graphconnectoritembyexternalidnokind), [GraphConnectorItemByPath](#graphconnectoritembypath), [GraphConnectorItemByPathNoKind](#graphconnectoritembypathnokind), [GraphConnectorItemByUrl](#graphconnectoritembyurl), [GraphConnectorItemByUrlNoKind](#graphconnectoritembyurlnokind), [GraphConnectorItemType](#graphconnectoritemtype), [GraphConnectorSearchSource](#graphconnectorsearchsource), [GraphDriveItemSyncDetails](#graphdriveitemsyncdetails), [GraphDriveItemSyncDetailsNoKind](#graphdriveitemsyncdetailsnokind), [IngestionBasedGraphSearchConfiguration](#ingestionbasedgraphsearchconfiguration), [KnowledgeSource](#knowledgesource), [MeetingSearchSource](#meetingsearchsource), [NoCitationsInModelOutput](#nocitationsinmodeloutput), [PublicDataSource](#publicdatasource), [PublicSiteSearchDataSource](#publicsitesearchdatasource), [PublicSiteSearchDataSourceNoKind](#publicsitesearchdatasourcenokind), [PublicSiteSearchSource](#publicsitesearchsource), [SearchAllKnowledgeSources](#searchallknowledgesources), [SearchKnowledgeSourcesMode](#searchknowledgesourcesmode), [SearchResultRecord](#searchresultrecord), [SearchSpecificKnowledgeSources](#searchspecificknowledgesources), [SearchSummarizationConfigurationState](#searchsummarizationconfigurationstate), [SharepointIds](#sharepointids), [SharepointIdsNoKind](#sharepointidsnokind), [SharePointIndexBehavior](#sharepointindexbehavior), [SharePointSearchDataSource](#sharepointsearchdatasource), [SharePointSearchSource](#sharepointsearchsource), [SharepointSourceType](#sharepointsourcetype), [SqlFederatedTableSearchConfiguration](#sqlfederatedtablesearchconfiguration), [SummarizationExternalModelConfiguration](#summarizationexternalmodelconfiguration), [SummarizationExternalModelConfigurationNoKind](#summarizationexternalmodelconfigurationnokind), [TeamsMessageSearchSource](#teamsmessagesearchsource)

</details>

### AzureAISearchSource

YAML: `kind: AzureAISearchSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `connectionReference` | `string` |  |  |  |
| `indexName` | `string` | ✓ |  |  |
| `nearestNeighbors` | `number` |  | `3` |  |
| `selectFields` | list of `string` |  |  |  |
| `searchFields` | list of `string` |  |  |  |
| `semanticConfigurationName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### AzureCognitiveSearchInternalDataSource

YAML: `kind: AzureCognitiveSearchInternalDataSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `indexName` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `titleField` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `urlField` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `filePathField` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `contentFields` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `vectorFields` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `embeddingDeploymentName` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `semanticConfiguration` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `queryType` | `"simple"` \| `"semantic"` \| `"vector"` \| `"vector_simple_hybrid"` \| `"vector_semantic_hybrid"` \| [StringExpression](other.md#stringexpression) |  |  |  |
| `roleInformation` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `limitResponsesToDataContent` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `strictness` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `topNDocuments` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

### AzureCognitiveSearchInternalDataSourceNoKind

*Title:* AzureCognitiveSearchInternalDataSource  
Same shape as [AzureCognitiveSearchInternalDataSource](#azurecognitivesearchinternaldatasource) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `indexName` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `titleField` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `urlField` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `filePathField` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `contentFields` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `vectorFields` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `embeddingDeploymentName` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `semanticConfiguration` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `queryType` | `"simple"` \| `"semantic"` \| `"vector"` \| `"vector_simple_hybrid"` \| `"vector_semantic_hybrid"` \| [StringExpression](other.md#stringexpression) |  |  |  |
| `roleInformation` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `limitResponsesToDataContent` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `strictness` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |
| `topNDocuments` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [AzureOpenAIOnYourDataKnowledgeSource](#azureopenaionyourdataknowledgesource), [AzureOpenAIOnYourDataSource](#azureopenaionyourdatasource), [AzureOpenAIOnYourDataSourceNoKind](actions.md#azureopenaionyourdatasourcenokind)</sub>

### AzureOpenAIOnYourDataKnowledgeSource

YAML: `kind: AzureOpenAIOnYourDataKnowledgeSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `connectionReference` | `string` |  |  |  |
| `dataSources` | list of [AzureCognitiveSearchInternalDataSourceNoKind](#azurecognitivesearchinternaldatasourcenokind) | ✓ |  |  |
| `apiVersion` | `"2023-06-01-preview"` \| `"2024-02-15-preview"` \| [StringExpression](other.md#stringexpression) | ✓ |  |  |
| `deployment` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `temperature` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `topP` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `stopSequence` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `maxTokens` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### AzureOpenAIOnYourDataSource

YAML: `kind: AzureOpenAIOnYourDataSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `connectionReference` | `string` |  |  |  |
| `dataSources` | list of [AzureCognitiveSearchInternalDataSourceNoKind](#azurecognitivesearchinternaldatasourcenokind) | ✓ |  |  |
| `apiVersion` | `"2023-06-01-preview"` \| `"2024-02-15-preview"` \| [StringExpression](other.md#stringexpression) | ✓ |  |  |
| `deployment` | [StringExpression](other.md#stringexpression) | ✓ |  | accepts a `=` Power Fx expression |
| `temperature` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `topP` | [NumberExpression](other.md#numberexpression) |  |  | accepts a `=` Power Fx expression |
| `stopSequence` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `maxTokens` | [IntExpression](other.md#intexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

### BaseSearchConfiguration


**One of:** [SqlFederatedTableSearchConfiguration](#sqlfederatedtablesearchconfiguration), [IngestionBasedGraphSearchConfiguration](#ingestionbasedgraphsearchconfiguration), [GenericDatasetNameSearchConfiguration](#genericdatasetnamesearchconfiguration)

<sub>Used by: [FederatedTableSearch](components.md#federatedtablesearch)</sub>

### BingCustomSearchDataSource

YAML: `kind: BingCustomSearchDataSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `customConfigurationId` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [PublicDataSource](#publicdatasource)</sub>

### BingCustomSearchSource

YAML: `kind: BingCustomSearchSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `customConfigurationId` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### BingSearchOptions

YAML: `kind: BingSearchOptions`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `market` | `"es-AR"` \| `"en-AU"` \| `"de-AT"` \| `"nl-BE"` \| `"fr-BE"` \| `"pt-BR"` \| `"en-CA"` \| `"fr-CA"` \| `"es-CL"` \| `"da-DK"` \| `"fi-FI"` \| `"fr-FR"` \| `"de-DE"` \| `"zh-HK"` \| `"en-IN"` \| `"en-ID"` \| `"it-IT"` \| `"ja-JP"` \| `"ko-KR"` \| `"en-MY"` \| `"es-MX"` \| `"nl-NL"` \| `"en-NZ"` \| `"no-NO"` \| `"zh-CN"` \| `"pl-PL"` \| `"en-PH"` \| `"ru-RU"` \| `"en-ZA"` \| `"es-ES"` \| `"sv-SE"` \| `"fr-CH"` \| `"de-CH"` \| `"zh-TW"` \| `"tr-TR"` \| `"en-GB"` \| `"en-US"` \| `"es-US"` \| [StringExpression](other.md#stringexpression) |  |  |  |

_No other properties are allowed._

### Citation

YAML: `kind: Citation`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `name` | `string` | ✓ |  |  |
| `text` | `string` | ✓ |  |  |
| `url` | `string` |  |  |  |
| `iconUrl` | `string` |  |  |  |

_No other properties are allowed._

### CitationNoKind

Same shape as [Citation](#citation) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` | ✓ |  |  |
| `name` | `string` | ✓ |  |  |
| `text` | `string` | ✓ |  |  |
| `url` | `string` |  |  |  |
| `iconUrl` | `string` |  |  |  |

_No other properties are allowed._

### CitationsInModelOutput

YAML: `kind: CitationsInModelOutput`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [CustomModelCitationHandlingType](#custommodelcitationhandlingtype)</sub>

### CopilotConnectorEntityNoKind

Same shape as [CopilotConnectorEntity](entities.md#copilotconnectorentity) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `entityName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [CopilotConnectorKnowledgeSourceDemo](#copilotconnectorknowledgesourcedemo)</sub>

### CopilotConnectorKnowledgeSourceDemo

YAML: `kind: CopilotConnectorKnowledgeSourceDemo`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `consent` | `boolean` | ✓ |  |  |
| `connectionReference` | `string` | ✓ |  |  |
| `graphConnectionIdPrefix` | `string` | ✓ |  |  |
| `entityLevelConfiguration` | list of [CopilotConnectorEntityNoKind](#copilotconnectorentitynokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### CustomDataSource

YAML: `kind: CustomDataSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `searchResults` | `=` expression |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

### CustomKnowledgeSource

YAML: `kind: CustomKnowledgeSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `searchResults` | `=` expression | ✓ |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### CustomModelCitationHandlingType


**One of:** [NoCitationsInModelOutput](#nocitationsinmodeloutput), [CitationsInModelOutput](#citationsinmodeloutput)

<sub>Used by: [SummarizationExternalModelConfiguration](#summarizationexternalmodelconfiguration), [SummarizationExternalModelConfigurationNoKind](#summarizationexternalmodelconfigurationnokind)</sub>

### CustomSearchInputContent

YAML: `kind: CustomSearchInputContent`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `title` | `string` |  |  |  |
| `contentLocation` | `string` |  |  |  |
| `content` | `string` | ✓ |  |  |

_No other properties are allowed._

### CustomSearchInputContentNoKind

Same shape as [CustomSearchInputContent](#customsearchinputcontent) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `title` | `string` |  |  |  |
| `contentLocation` | `string` |  |  |  |
| `content` | `string` | ✓ |  |  |

_No other properties are allowed._

### DataverseStructuredSearchSource

YAML: `kind: DataverseStructuredSearchSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `skillConfiguration` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### DataverseTableSearchBase


**One of:** [DataverseTableSearch](components.md#dataversetablesearch), [FederatedTableSearch](components.md#federatedtablesearch)

<sub>Used by: [BotComponentCollectionDefinition](components.md#botcomponentcollectiondefinition), [BotDefinition](components.md#botdefinition)</sub>

### EmailSearchSource

YAML: `kind: EmailSearchSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `mailbox` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `folders` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### FabricAISkillSource

YAML: `kind: FabricAISkillSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `workspaceId` | [EnvironmentVariableReferenceNoKind](components.md#environmentvariablereferencenokind) | ✓ |  |  |
| `aISkillId` | [EnvironmentVariableReferenceNoKind](components.md#environmentvariablereferencenokind) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### FederatedStructuredSearchSource

YAML: `kind: FederatedStructuredSearchSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `connectionProperties` | [ConnectionPropertiesNoKind](settings.md#connectionpropertiesnokind) |  |  |  |
| `skillConfiguration` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### FileGroupKnowledgeSource

YAML: `kind: FileGroupKnowledgeSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `instructions` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### FileIndexingStatus


**String enum.** Allowed values: `NotStarted`, `InProgress`, `Completed`

<sub>Used by: [GraphDriveItemSyncDetails](#graphdriveitemsyncdetails), [GraphDriveItemSyncDetailsNoKind](#graphdriveitemsyncdetailsnokind)</sub>

### FileSearchDataSource

YAML: `kind: FileSearchDataSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `searchFilesMode` | [SearchFilesMode](actions.md#searchfilesmode) |  |  |  |

_No other properties are allowed._

### FileSearchDataSourceNoKind

Same shape as [FileSearchDataSource](#filesearchdatasource) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `searchFilesMode` | [SearchFilesMode](actions.md#searchfilesmode) |  |  |  |

_No other properties are allowed._

<sub>Used by: [GptComponentMetadata](components.md#gptcomponentmetadata), [GptComponentMetadataNoKind](components.md#gptcomponentmetadatanokind), [SearchAndSummarizeContent](actions.md#searchandsummarizecontent), [SearchAndSummarizeWithCustomModel](actions.md#searchandsummarizewithcustommodel), [SearchKnowledgeSources](actions.md#searchknowledgesources)</sub>

### GenericDatasetNameSearchConfiguration

YAML: `kind: GenericDatasetNameSearchConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `datasetName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [BaseSearchConfiguration](#basesearchconfiguration)</sub>

### GraphConnectorItemByExternalId

YAML: `kind: GraphConnectorItemByExternalId`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `graphConnectorItemType` | [GraphConnectorItemType](#graphconnectoritemtype) |  |  |  |
| `externalId` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

### GraphConnectorItemByExternalIdNoKind

Same shape as [GraphConnectorItemByExternalId](#graphconnectoritembyexternalid) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `graphConnectorItemType` | [GraphConnectorItemType](#graphconnectoritemtype) |  |  |  |
| `externalId` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [GraphConnectorSearchSource](#graphconnectorsearchsource)</sub>

### GraphConnectorItemByPath

YAML: `kind: GraphConnectorItemByPath`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `graphConnectorItemType` | [GraphConnectorItemType](#graphconnectoritemtype) |  |  |  |
| `path` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

### GraphConnectorItemByPathNoKind

Same shape as [GraphConnectorItemByPath](#graphconnectoritembypath) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `graphConnectorItemType` | [GraphConnectorItemType](#graphconnectoritemtype) |  |  |  |
| `path` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [GraphConnectorSearchSource](#graphconnectorsearchsource)</sub>

### GraphConnectorItemByUrl

YAML: `kind: GraphConnectorItemByUrl`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `graphConnectorItemType` | [GraphConnectorItemType](#graphconnectoritemtype) |  |  |  |
| `url` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

### GraphConnectorItemByUrlNoKind

Same shape as [GraphConnectorItemByUrl](#graphconnectoritembyurl) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `id` | `string` |  |  |  |
| `graphConnectorItemType` | [GraphConnectorItemType](#graphconnectoritemtype) |  |  |  |
| `url` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [GraphConnectorSearchSource](#graphconnectorsearchsource)</sub>

### GraphConnectorItemType


**String enum.** Allowed values: `Item`, `Container`

<sub>Used by: [GraphConnectorItemByExternalId](#graphconnectoritembyexternalid), [GraphConnectorItemByExternalIdNoKind](#graphconnectoritembyexternalidnokind), [GraphConnectorItemByPath](#graphconnectoritembypath), [GraphConnectorItemByPathNoKind](#graphconnectoritembypathnokind), [GraphConnectorItemByUrl](#graphconnectoritembyurl), [GraphConnectorItemByUrlNoKind](#graphconnectoritembyurlnokind)</sub>

### GraphConnectorSearchSource

YAML: `kind: GraphConnectorSearchSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `connectionId` | [EnvironmentVariableReferenceNoKind](components.md#environmentvariablereferencenokind) | ✓ |  |  |
| `connectionName` | `string` | ✓ |  |  |
| `contentSourceDisplayName` | `string` |  |  |  |
| `contentSourceIcon` | `string` |  |  |  |
| `publisherName` | `string` |  |  |  |
| `additionalSearchTerms` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |
| `itemsByUrl` | list of [GraphConnectorItemByUrlNoKind](#graphconnectoritembyurlnokind) |  |  |  |
| `itemsByExternalId` | list of [GraphConnectorItemByExternalIdNoKind](#graphconnectoritembyexternalidnokind) |  |  |  |
| `itemsByPath` | list of [GraphConnectorItemByPathNoKind](#graphconnectoritembypathnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### GraphDriveItemSyncDetails

YAML: `kind: GraphDriveItemSyncDetails`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `displayName` | `string` | ✓ |  |  |
| `webUrl` | `string` | ✓ |  |  |
| `driveId` | `string` | ✓ |  |  |
| `itemId` | `string` | ✓ |  |  |
| `sharepointIds` | [SharepointIdsNoKind](#sharepointidsnokind) | ✓ |  |  |
| `isDeleted` | `boolean` | ✓ |  |  |
| `indexingStatus` | [FileIndexingStatus](#fileindexingstatus) |  |  |  |
| `sensitivityLabelId` | `string` |  |  |  |
| `sensitivityLabelStatusCode` | [SensitivityLabelStatusCode](settings.md#sensitivitylabelstatuscode) |  |  |  |
| `sharepointSourceType` | [SharepointSourceType](#sharepointsourcetype) |  |  |  |

_No other properties are allowed._

### GraphDriveItemSyncDetailsNoKind

*Title:* GraphDriveItemSyncDetails  
Same shape as [GraphDriveItemSyncDetails](#graphdriveitemsyncdetails) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `displayName` | `string` | ✓ |  |  |
| `webUrl` | `string` | ✓ |  |  |
| `driveId` | `string` | ✓ |  |  |
| `itemId` | `string` | ✓ |  |  |
| `sharepointIds` | [SharepointIdsNoKind](#sharepointidsnokind) | ✓ |  |  |
| `isDeleted` | `boolean` | ✓ |  |  |
| `indexingStatus` | [FileIndexingStatus](#fileindexingstatus) |  |  |  |
| `sensitivityLabelId` | `string` |  |  |  |
| `sensitivityLabelStatusCode` | [SensitivityLabelStatusCode](settings.md#sensitivitylabelstatuscode) |  |  |  |
| `sharepointSourceType` | [SharepointSourceType](#sharepointsourcetype) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BotSynchronizationDetails](components.md#botsynchronizationdetails), [BotSynchronizationDetailsNoKind](components.md#botsynchronizationdetailsnokind), [IngestionBasedGraphSearchConfiguration](#ingestionbasedgraphsearchconfiguration)</sub>

### IngestionBasedGraphSearchConfiguration

YAML: `kind: IngestionBasedGraphSearchConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `driveItems` | list of [GraphDriveItemSyncDetailsNoKind](#graphdriveitemsyncdetailsnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [BaseSearchConfiguration](#basesearchconfiguration)</sub>

### KnowledgeSource


**One of:** [FileGroupKnowledgeSource](#filegroupknowledgesource), [PublicSiteSearchSource](#publicsitesearchsource), [BingCustomSearchSource](#bingcustomsearchsource), [CustomKnowledgeSource](#customknowledgesource), [AzureOpenAIOnYourDataKnowledgeSource](#azureopenaionyourdataknowledgesource), [AzureAISearchSource](#azureaisearchsource), [GraphConnectorSearchSource](#graphconnectorsearchsource), [FabricAISkillSource](#fabricaiskillsource), [SharePointSearchSource](#sharepointsearchsource), [TeamsMessageSearchSource](#teamsmessagesearchsource), [MeetingSearchSource](#meetingsearchsource), [EmailSearchSource](#emailsearchsource), [DataverseStructuredSearchSource](#dataversestructuredsearchsource), [FederatedStructuredSearchSource](#federatedstructuredsearchsource), [CopilotConnectorKnowledgeSourceDemo](#copilotconnectorknowledgesourcedemo)

<sub>Used by: [KnowledgeSourceConfiguration](components.md#knowledgesourceconfiguration), [KnowledgeSourceConfigurationNoKind](components.md#knowledgesourceconfigurationnokind)</sub>

### MeetingSearchSource

YAML: `kind: MeetingSearchSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `url` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `id` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `isSeries` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### NoCitationsInModelOutput

YAML: `kind: NoCitationsInModelOutput`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [CustomModelCitationHandlingType](#custommodelcitationhandlingtype)</sub>

### PublicDataSource


**One of:** [PublicSiteSearchDataSource](#publicsitesearchdatasource), [PublicSiteSearchDataSourceNoKind](#publicsitesearchdatasourcenokind), [BingCustomSearchDataSource](#bingcustomsearchdatasource)

<sub>Used by: [SearchAndSummarizeContent](actions.md#searchandsummarizecontent)</sub>

### PublicSiteSearchDataSource

YAML: `kind: PublicSiteSearchDataSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sites` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `pages` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [PublicDataSource](#publicdatasource)</sub>

### PublicSiteSearchDataSourceNoKind

Same shape as [PublicSiteSearchDataSource](#publicsitesearchdatasource) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sites` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `pages` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [PublicDataSource](#publicdatasource)</sub>

### PublicSiteSearchSource

YAML: `kind: PublicSiteSearchSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `site` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `includeSubPages` | `boolean` |  | `true` |  |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### SearchAllKnowledgeSources

YAML: `kind: SearchAllKnowledgeSources`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [SearchKnowledgeSourcesMode](#searchknowledgesourcesmode)</sub>

### SearchKnowledgeSourcesMode


**One of:** [SearchAllKnowledgeSources](#searchallknowledgesources), [SearchSpecificKnowledgeSources](#searchspecificknowledgesources)

<sub>Used by: [GptComponentMetadata](components.md#gptcomponentmetadata), [GptComponentMetadataNoKind](components.md#gptcomponentmetadatanokind), [SearchAndSummarizeContent](actions.md#searchandsummarizecontent), [SearchAndSummarizeWithCustomModel](actions.md#searchandsummarizewithcustommodel), [SearchKnowledgeSources](actions.md#searchknowledgesources)</sub>

### SearchResultRecord

YAML: `kind: SearchResultRecord`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `title` | `string` | ✓ |  |  |
| `contentLocation` | `string` | ✓ |  |  |
| `content` | `string` | ✓ |  |  |
| `knowledgeType` | `string` | ✓ |  |  |
| `knowledgeName` | `string` | ✓ |  |  |

_No other properties are allowed._

### SearchSpecificKnowledgeSources

YAML: `kind: SearchSpecificKnowledgeSources`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `knowledgeSources` | list of `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [SearchKnowledgeSourcesMode](#searchknowledgesourcesmode)</sub>

### SearchSummarizationConfigurationState


**String enum.** Allowed values: `NotConfigured`, `Configured`, `ConfiguredWithLegacyInfo`, `Customized`

### SharepointIds

YAML: `kind: SharepointIds`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `webId` | `string` | ✓ |  |  |
| `listId` | `string` | ✓ |  |  |
| `siteId` | `string` | ✓ |  |  |
| `siteUrl` | `string` | ✓ |  |  |
| `uniqueId` | `string` |  |  |  |
| `partId` | `string` |  |  |  |
| `partType` | [PartType](actions.md#parttype) |  |  |  |

_No other properties are allowed._

### SharepointIdsNoKind

*Title:* SharepointIds  
Same shape as [SharepointIds](#sharepointids) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `webId` | `string` | ✓ |  |  |
| `listId` | `string` | ✓ |  |  |
| `siteId` | `string` | ✓ |  |  |
| `siteUrl` | `string` | ✓ |  |  |
| `uniqueId` | `string` |  |  |  |
| `partId` | `string` |  |  |  |
| `partType` | [PartType](actions.md#parttype) |  |  |  |

_No other properties are allowed._

<sub>Used by: [GraphDriveItemSyncDetails](#graphdriveitemsyncdetails), [GraphDriveItemSyncDetailsNoKind](#graphdriveitemsyncdetailsnokind)</sub>

### SharePointIndexBehavior


**String enum.** Allowed values: `Disabled`, `EnabledWithFallback`, `EnabledWithoutFallback`

<sub>Used by: [SharePointSearchDataSource](#sharepointsearchdatasource), [SharePointSearchDataSourceNoKind](actions.md#sharepointsearchdatasourcenokind), [SharePointSearchSource](#sharepointsearchsource)</sub>

### SharePointSearchDataSource

YAML: `kind: SharePointSearchDataSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sites` | `=` expression \| list of `string` |  |  | accepts a `=` Power Fx expression |
| `indexBehavior` | [SharePointIndexBehavior](#sharepointindexbehavior) |  | `"EnabledWithFallback"` |  |
| `additionalSearchTerms` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

### SharePointSearchSource

YAML: `kind: SharePointSearchSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `cascadeShare` | `boolean` |  | `true` |  |
| `site` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `indexBehavior` | [SharePointIndexBehavior](#sharepointindexbehavior) |  | `"EnabledWithFallback"` |  |
| `additionalSearchTerms` | [TemplateLineNoKind](messages.md#templatelinenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>

### SharepointSourceType


**String enum.** Allowed values: `Site`, `Folder`, `File`, `List`

<sub>Used by: [GraphDriveItemSyncDetails](#graphdriveitemsyncdetails), [GraphDriveItemSyncDetailsNoKind](#graphdriveitemsyncdetailsnokind)</sub>

### SqlFederatedTableSearchConfiguration

YAML: `kind: SqlFederatedTableSearchConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `sqlServerName` | `string` |  |  |  |
| `sqlDbName` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [BaseSearchConfiguration](#basesearchconfiguration)</sub>

### SummarizationExternalModelConfiguration

YAML: `kind: SummarizationExternalModelConfiguration`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelConfiguration` | [ExternalChatCompletionModelConfigurationNoKind](actions.md#externalchatcompletionmodelconfigurationnokind) | ✓ |  |  |
| `requestTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  | `120000` | accepts a `=` Power Fx expression |
| `citationHandling` | [CustomModelCitationHandlingType](#custommodelcitationhandlingtype) |  |  |  |

_No other properties are allowed._

### SummarizationExternalModelConfigurationNoKind

*Title:* SummarizationExternalModelConfiguration  
Same shape as [SummarizationExternalModelConfiguration](#summarizationexternalmodelconfiguration) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `modelConfiguration` | [ExternalChatCompletionModelConfigurationNoKind](actions.md#externalchatcompletionmodelconfigurationnokind) | ✓ |  |  |
| `requestTimeoutInMilliseconds` | [IntExpression](other.md#intexpression) |  | `120000` | accepts a `=` Power Fx expression |
| `citationHandling` | [CustomModelCitationHandlingType](#custommodelcitationhandlingtype) |  |  |  |

_No other properties are allowed._

<sub>Used by: [CurrentModels](settings.md#currentmodels), [CurrentModelsNoKind](settings.md#currentmodelsnokind), [CustomModels](settings.md#custommodels), [DeferredUpdateModels](settings.md#deferredupdatemodels), [ExperimentalModels](settings.md#experimentalmodels), [PreviewModels](settings.md#previewmodels), [ReasoningExperimentalModels](settings.md#reasoningexperimentalmodels), [ReasoningPreviewModels](settings.md#reasoningpreviewmodels), [ReasoningProductionModels](settings.md#reasoningproductionmodels), [SearchAndSummarizeWithCustomModel](actions.md#searchandsummarizewithcustommodel), [TurboExperimentalModels](settings.md#turboexperimentalmodels), [TurboPreviewModels](settings.md#turbopreviewmodels), [TurboProductionModels](settings.md#turboproductionmodels)</sub>

### TeamsMessageSearchSource

YAML: `kind: TeamsMessageSearchSource`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `triggerCondition` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |
| `url` | [StringExpression](other.md#stringexpression) |  |  | accepts a `=` Power Fx expression |
| `isRecurringMeeting` | [BoolExpression](other.md#boolexpression) |  |  | accepts a `=` Power Fx expression |

_No other properties are allowed._

<sub>Used by: [KnowledgeSource](#knowledgesource)</sub>
