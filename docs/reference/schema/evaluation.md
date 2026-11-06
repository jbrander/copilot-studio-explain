# Evaluation & testing

> **Generated** by `tools/gen_reference.py` from `reference/bot.schema.yaml-authoring.json` (vendor/skills-for-copilot-studio @ `9c42d3a`). Do not edit by hand; re-run the script.

Test sets, evaluation runs, graders and assertions. See [12 Evaluations & testing](../../12-evaluations-and-testing.md).

**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.

**52 definitions** on this page (35 with a `kind:`). [A–Z index of all definitions](index.md).

<details><summary>Contents</summary>

[AbstentionAssertion](#abstentionassertion), [AgentEvaluationTarget](#agentevaluationtarget), [AnnotatedSample](#annotatedsample), [AnnotatedSampleNoKind](#annotatedsamplenokind), [CodeGrader](#codegrader), [CompareMeaningGrader](#comparemeaninggrader), [CompletenessAssertion](#completenessassertion), [ContainsAllAssertion](#containsallassertion), [ContainsAllGrader](#containsallgrader), [ContainsAnyAssertion](#containsanyassertion), [ContainsAnyGrader](#containsanygrader), [CosineEmbeddingSimilarityAssertion](#cosineembeddingsimilarityassertion), [DataPolicyEvaluationViolationType](#datapolicyevaluationviolationtype), [EvaluationData](#evaluationdata), [EvaluationRunDefinition](#evaluationrundefinition), [EvaluationSet](#evaluationset), [EvaluationTargetBase](#evaluationtargetbase), [ExactMatchAssertion](#exactmatchassertion), [ExactMatchGrader](#exactmatchgrader), [ExpectedExecutionStepBase](#expectedexecutionstepbase), [ExpectedToolCallStep](#expectedtoolcallstep), [ExpectedTopicTriggeredStep](#expectedtopictriggeredstep), [FlowEvaluationTarget](#flowevaluationtarget), [GeneralQualityGrader](#generalqualitygrader), [GraderBase](#graderbase), [GraderOutcome](#graderoutcome), [GroundednessAssertion](#groundednessassertion), [IntentMatchAssertion](#intentmatchassertion), [MultiTurnActivity](#multiturnactivity), [MultiTurnActivityNoKind](#multiturnactivitynokind), [MultiTurnEvaluationCase](#multiturnevaluationcase), [PromptEvaluationTarget](#promptevaluationtarget), [PromptGrader](#promptgrader), [PromptGraderLabel](#promptgraderlabel), [PromptGraderLabelNoKind](#promptgraderlabelnokind), [RelevanceAssertion](#relevanceassertion), [SampleDataType](#sampledatatype), [SampleSegment](#samplesegment), [SampleSegmentNoKind](#samplesegmentnokind), [SimpleEvaluationCase](#simpleevaluationcase), [SimpleEvaluationCaseNoKind](#simpleevaluationcasenokind), [TaskSuccessGrader](#tasksuccessgrader), [TestActivityDefinition](#testactivitydefinition), [TestActivitySourceType](#testactivitysourcetype), [TestAssertionBase](#testassertionbase), [TestCaseDefinition](#testcasedefinition), [TestDefinitionBase](#testdefinitionbase), [TestSetDefinition](#testsetdefinition), [TestTranscriptDefinition](#testtranscriptdefinition), [TestTranscriptDefinitionNoKind](#testtranscriptdefinitionnokind), [TextSimilarityGrader](#textsimilaritygrader), [ToolCallGrader](#toolcallgrader)

</details>

### AbstentionAssertion

YAML: `kind: AbstentionAssertion`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [TestAssertionBase](#testassertionbase)</sub>

### AgentEvaluationTarget

YAML: `kind: AgentEvaluationTarget`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `agentSchemaName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [EvaluationTargetBase](#evaluationtargetbase)</sub>

### AnnotatedSample


**Type:** `string`

### AnnotatedSampleNoKind


**Type:** `string`

<sub>Used by: [ClosedListEntity](dialogs.md#closedlistentity), [DynamicClosedListEntity](entities.md#dynamicclosedlistentity), [Intent](dialogs.md#intent), [IntentNoKind](dialogs.md#intentnokind), [RegexEntity](dialogs.md#regexentity)</sub>

### CodeGrader

YAML: `kind: CodeGrader`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `name` | `string` | ✓ |  |  |
| `code` | `string` | ✓ |  |  |
| `threshold` | `number` |  | `0.5` |  |

_No other properties are allowed._

<sub>Used by: [GraderBase](#graderbase)</sub>

### CompareMeaningGrader

YAML: `kind: CompareMeaningGrader`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `threshold` | `number` |  | `0.6` |  |

_No other properties are allowed._

<sub>Used by: [GraderBase](#graderbase)</sub>

### CompletenessAssertion

YAML: `kind: CompletenessAssertion`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [TestAssertionBase](#testassertionbase)</sub>

### ContainsAllAssertion

YAML: `kind: ContainsAllAssertion`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `expectedResponses` | list of `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [TestAssertionBase](#testassertionbase)</sub>

### ContainsAllGrader

YAML: `kind: ContainsAllGrader`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [GraderBase](#graderbase)</sub>

### ContainsAnyAssertion

YAML: `kind: ContainsAnyAssertion`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `expectedResponses` | list of `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [TestAssertionBase](#testassertionbase)</sub>

### ContainsAnyGrader

YAML: `kind: ContainsAnyGrader`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [GraderBase](#graderbase)</sub>

### CosineEmbeddingSimilarityAssertion

YAML: `kind: CosineEmbeddingSimilarityAssertion`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `expectedResponse` | `string` |  |  |  |
| `scoreThreshold` | `number` |  | `0.7` |  |

_No other properties are allowed._

<sub>Used by: [TestAssertionBase](#testassertionbase)</sub>

### DataPolicyEvaluationViolationType


**String enum.** Allowed values: `NotSpecified`, `BlockedConnector`, `BusinessAndNonBusinessConnector`, `BlockedEndpoint`, `BlockedAction`

### EvaluationData

YAML: `kind: EvaluationData`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `rows` | list of [SimpleEvaluationCaseNoKind](#simpleevaluationcasenokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TestDefinitionBase](#testdefinitionbase)</sub>

### EvaluationRunDefinition

YAML: `kind: EvaluationRunDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `target` | [EvaluationTargetBase](#evaluationtargetbase) | ✓ |  |  |
| `evaluationSetSchemaName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [TestDefinitionBase](#testdefinitionbase)</sub>

### EvaluationSet

YAML: `kind: EvaluationSet`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `graders` | list of [GraderBase](#graderbase) |  |  |  |
| `inputType` | [DataType](data-types.md#datatype) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TestDefinitionBase](#testdefinitionbase)</sub>

### EvaluationTargetBase


**One of:** [AgentEvaluationTarget](#agentevaluationtarget), [PromptEvaluationTarget](#promptevaluationtarget), [FlowEvaluationTarget](#flowevaluationtarget)

<sub>Used by: [EvaluationRunDefinition](#evaluationrundefinition)</sub>

### ExactMatchAssertion

YAML: `kind: ExactMatchAssertion`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `expectedResponse` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [TestAssertionBase](#testassertionbase)</sub>

### ExactMatchGrader

YAML: `kind: ExactMatchGrader`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [GraderBase](#graderbase)</sub>

### ExpectedExecutionStepBase


**One of:** [ExpectedToolCallStep](#expectedtoolcallstep), [ExpectedTopicTriggeredStep](#expectedtopictriggeredstep)

<sub>Used by: [MultiTurnEvaluationCase](#multiturnevaluationcase), [SimpleEvaluationCase](#simpleevaluationcase), [SimpleEvaluationCaseNoKind](#simpleevaluationcasenokind)</sub>

### ExpectedToolCallStep

YAML: `kind: ExpectedToolCallStep`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `schemaName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [ExpectedExecutionStepBase](#expectedexecutionstepbase)</sub>

### ExpectedTopicTriggeredStep

YAML: `kind: ExpectedTopicTriggeredStep`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `schemaName` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [ExpectedExecutionStepBase](#expectedexecutionstepbase)</sub>

### FlowEvaluationTarget

YAML: `kind: FlowEvaluationTarget`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `flowId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [EvaluationTargetBase](#evaluationtargetbase)</sub>

### GeneralQualityGrader

YAML: `kind: GeneralQualityGrader`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [GraderBase](#graderbase)</sub>

### GraderBase


**One of:** [GeneralQualityGrader](#generalqualitygrader), [TaskSuccessGrader](#tasksuccessgrader), [CompareMeaningGrader](#comparemeaninggrader), [TextSimilarityGrader](#textsimilaritygrader), [ExactMatchGrader](#exactmatchgrader), [ContainsAllGrader](#containsallgrader), [ContainsAnyGrader](#containsanygrader), [ToolCallGrader](#toolcallgrader), [PromptGrader](#promptgrader), [CodeGrader](#codegrader)

<sub>Used by: [EvaluationSet](#evaluationset)</sub>

### GraderOutcome


**String enum.** Allowed values: `Pass`, `Fail`

<sub>Used by: [PromptGraderLabel](#promptgraderlabel), [PromptGraderLabelNoKind](#promptgraderlabelnokind)</sub>

### GroundednessAssertion

YAML: `kind: GroundednessAssertion`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [TestAssertionBase](#testassertionbase)</sub>

### IntentMatchAssertion

YAML: `kind: IntentMatchAssertion`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `expectedResponse` | `string` |  |  |  |
| `scoreThreshold` | `number` |  | `60` |  |

_No other properties are allowed._

<sub>Used by: [TestAssertionBase](#testassertionbase)</sub>

### MultiTurnActivity

YAML: `kind: MultiTurnActivity`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `activity` | [ActivityTemplateBase](messages.md#activitytemplatebase) | ✓ |  |  |
| `expectedOutput` | `string` |  |  |  |

_No other properties are allowed._

### MultiTurnActivityNoKind

*Title:* MultiTurnActivity  
Same shape as [MultiTurnActivity](#multiturnactivity) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `activity` | [ActivityTemplateBase](messages.md#activitytemplatebase) | ✓ |  |  |
| `expectedOutput` | `string` |  |  |  |

_No other properties are allowed._

<sub>Used by: [MultiTurnEvaluationCase](#multiturnevaluationcase)</sub>

### MultiTurnEvaluationCase

YAML: `kind: MultiTurnEvaluationCase`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `source` | [TestActivitySourceType](#testactivitysourcetype) |  | `"Manual"` |  |
| `expectedOutput` | `string` |  |  |  |
| `expectedKeywords` | list of `string` |  |  |  |
| `expectedExecutionSteps` | list of [ExpectedExecutionStepBase](#expectedexecutionstepbase) |  |  |  |
| `activities` | list of [MultiTurnActivityNoKind](#multiturnactivitynokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TestDefinitionBase](#testdefinitionbase)</sub>

### PromptEvaluationTarget

YAML: `kind: PromptEvaluationTarget`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `aIModelId` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [EvaluationTargetBase](#evaluationtargetbase)</sub>

### PromptGrader

YAML: `kind: PromptGrader`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `name` | `string` | ✓ |  |  |
| `instructions` | `string` | ✓ |  |  |
| `model` | [AgentModelType](settings.md#agentmodeltype) |  |  |  |
| `labels` | list of [PromptGraderLabelNoKind](#promptgraderlabelnokind) |  |  |  |

_No other properties are allowed._

<sub>Used by: [GraderBase](#graderbase)</sub>

### PromptGraderLabel

YAML: `kind: PromptGraderLabel`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `name` | `string` | ✓ |  |  |
| `description` | `string` |  |  |  |
| `outcome` | [GraderOutcome](#graderoutcome) |  |  |  |

_No other properties are allowed._

### PromptGraderLabelNoKind

*Title:* PromptGraderLabel  
Same shape as [PromptGraderLabel](#promptgraderlabel) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `name` | `string` | ✓ |  |  |
| `description` | `string` |  |  |  |
| `outcome` | [GraderOutcome](#graderoutcome) |  |  |  |

_No other properties are allowed._

<sub>Used by: [PromptGrader](#promptgrader)</sub>

### RelevanceAssertion

YAML: `kind: RelevanceAssertion`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [TestAssertionBase](#testassertionbase)</sub>

### SampleDataType


**String enum.** Allowed values: `Json`, `Yaml`

### SampleSegment

YAML: `kind: SampleSegment`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `text` | `string` | ✓ |  |  |
| `variable` | `string` |  |  |  |

_No other properties are allowed._

### SampleSegmentNoKind

*Title:* SampleSegment  
Same shape as [SampleSegment](#samplesegment) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `text` | `string` | ✓ |  |  |
| `variable` | `string` |  |  |  |

_No other properties are allowed._

### SimpleEvaluationCase

YAML: `kind: SimpleEvaluationCase`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `source` | [TestActivitySourceType](#testactivitysourcetype) |  | `"Manual"` |  |
| `expectedOutput` | `string` |  |  |  |
| `expectedKeywords` | list of `string` |  |  |  |
| `expectedExecutionSteps` | list of [ExpectedExecutionStepBase](#expectedexecutionstepbase) |  |  |  |
| `input` | [DataValue](other.md#datavalue) | ✓ |  |  |

_No other properties are allowed._

### SimpleEvaluationCaseNoKind

*Title:* SimpleEvaluationCase  
Same shape as [SimpleEvaluationCase](#simpleevaluationcase) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `source` | [TestActivitySourceType](#testactivitysourcetype) |  | `"Manual"` |  |
| `expectedOutput` | `string` |  |  |  |
| `expectedKeywords` | list of `string` |  |  |  |
| `expectedExecutionSteps` | list of [ExpectedExecutionStepBase](#expectedexecutionstepbase) |  |  |  |
| `input` | [DataValue](other.md#datavalue) | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [EvaluationData](#evaluationdata)</sub>

### TaskSuccessGrader

YAML: `kind: TaskSuccessGrader`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [GraderBase](#graderbase)</sub>

### TestActivityDefinition


**One of:** [SendUserActivity](messages.md#senduseractivity)

<sub>Used by: [TestTranscriptDefinition](#testtranscriptdefinition), [TestTranscriptDefinitionNoKind](#testtranscriptdefinitionnokind)</sub>

### TestActivitySourceType


**String enum.** Allowed values: `Manual`, `Generated`, `Edited`, `Imported`, `AnalyticsTheme`

<sub>Used by: [MultiTurnEvaluationCase](#multiturnevaluationcase), [SendUserActivity](messages.md#senduseractivity), [SimpleEvaluationCase](#simpleevaluationcase), [SimpleEvaluationCaseNoKind](#simpleevaluationcasenokind)</sub>

### TestAssertionBase


**One of:** [RelevanceAssertion](#relevanceassertion), [AbstentionAssertion](#abstentionassertion), [CompletenessAssertion](#completenessassertion), [GroundednessAssertion](#groundednessassertion), [ExactMatchAssertion](#exactmatchassertion), [IntentMatchAssertion](#intentmatchassertion), [CosineEmbeddingSimilarityAssertion](#cosineembeddingsimilarityassertion), [ContainsAllAssertion](#containsallassertion), [ContainsAnyAssertion](#containsanyassertion)

<sub>Used by: [SendUserActivity](messages.md#senduseractivity)</sub>

### TestCaseDefinition

YAML: `kind: TestCaseDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `transcriptDefinition` | [TestTranscriptDefinitionNoKind](#testtranscriptdefinitionnokind) |  |  |  |
| `index` | `number` |  | `0` |  |

_No other properties are allowed._

<sub>Used by: [TestDefinitionBase](#testdefinitionbase)</sub>

### TestDefinitionBase


**One of:** [TestSetDefinition](#testsetdefinition), [TestCaseDefinition](#testcasedefinition), [EvaluationRunDefinition](#evaluationrundefinition), [EvaluationSet](#evaluationset), [EvaluationData](#evaluationdata), [MultiTurnEvaluationCase](#multiturnevaluationcase)

<sub>Used by: [TestCaseComponent](components.md#testcasecomponent)</sub>

### TestSetDefinition

YAML: `kind: TestSetDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|

_No other properties are allowed._

<sub>Used by: [TestDefinitionBase](#testdefinitionbase)</sub>

### TestTranscriptDefinition

YAML: `kind: TestTranscriptDefinition`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `testActivities` | list of [TestActivityDefinition](#testactivitydefinition) |  |  |  |

_No other properties are allowed._

### TestTranscriptDefinitionNoKind

Same shape as [TestTranscriptDefinition](#testtranscriptdefinition) but written **without** a `kind:` line (used where the type is implied by the parent property).

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `testActivities` | list of [TestActivityDefinition](#testactivitydefinition) |  |  |  |

_No other properties are allowed._

<sub>Used by: [TestCaseDefinition](#testcasedefinition)</sub>

### TextSimilarityGrader

YAML: `kind: TextSimilarityGrader`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `threshold` | `number` |  | `0.7` |  |

_No other properties are allowed._

<sub>Used by: [GraderBase](#graderbase)</sub>

### ToolCallGrader

YAML: `kind: ToolCallGrader`

| Property | Type | Req. | Default | Notes |
|---|---|---|---|---|
| `name` | `string` | ✓ |  |  |

_No other properties are allowed._

<sub>Used by: [GraderBase](#graderbase)</sub>
