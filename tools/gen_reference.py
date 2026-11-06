#!/usr/bin/env python3
"""Generate the docs/reference/** pages from the vendored skills-for-copilot-studio repo.

Inputs (read-only):
  vendor/skills-for-copilot-studio/reference/bot.schema.yaml-authoring.json
  vendor/skills-for-copilot-studio/reference/adaptive-card.schema.json
  vendor/skills-for-copilot-studio/reference/connectors/*.yml

Outputs (overwritten on every run):
  docs/reference/schema/*.md
  docs/reference/adaptive-card-schema.md
  docs/reference/connectors/*.md

Usage:  python3 -I tools/gen_reference.py
Requires: Python 3.9+, PyYAML.
"""
from __future__ import annotations

import json
import re
import subprocess
from collections import defaultdict
from pathlib import Path

import yaml

ROOT = Path(__file__).resolve().parent.parent
VENDOR = ROOT / "vendor" / "skills-for-copilot-studio"
BOT_SCHEMA = VENDOR / "reference" / "bot.schema.yaml-authoring.json"
CARD_SCHEMA = VENDOR / "reference" / "adaptive-card.schema.json"
CONNECTORS = VENDOR / "reference" / "connectors"
OUT = ROOT / "docs" / "reference"
SCHEMA_OUT = OUT / "schema"
CONNECTOR_OUT = OUT / "connectors"


# --------------------------------------------------------------------------- helpers

def vendor_commit() -> str:
    try:
        return subprocess.run(
            ["git", "-C", str(VENDOR), "rev-parse", "--short", "HEAD"],
            capture_output=True, text=True, check=True,
        ).stdout.strip()
    except Exception:  # pragma: no cover - best effort only
        return "unknown"


def slug(text: str) -> str:
    """GitHub-style heading anchor."""
    s = text.strip().lower()
    s = re.sub(r"[^\w\- ]", "", s)
    return s.replace(" ", "-")


def is_placeholder(text: str | None) -> bool:
    return not text or "BotSchema:$$" in text


def cell(text: str) -> str:
    """Make text safe for a markdown table cell."""
    return str(text).replace("|", "\\|").replace("\n", " ").strip()


def header(title: str, source: str, intro: str = "") -> str:
    lines = [
        f"# {title}",
        "",
        f"> **Generated** by `tools/gen_reference.py` from `{source}` "
        f"(vendor/skills-for-copilot-studio @ `{COMMIT}`). Do not edit by hand; re-run the script.",
        "",
    ]
    if intro:
        lines += [intro, ""]
    return "\n".join(lines)


COMMIT = vendor_commit()


# --------------------------------------------------------------------------- bot schema

CATEGORIES = {
    # key: (title, intro, seed unions, extra names)
    "components": (
        "Components & agent definitions",
        "Top-level file kinds and solution components: the agent itself (`GptComponentMetadata`), "
        "global variables, knowledge-source components, external triggers, connection references, and so on. "
        "See [02 Project layout](../../02-project-layout.md) and [03 The agent](../../03-agent.md).",
        ["SolutionComponent", "BotComponentBase", "AgentDefinition", "ExternalTriggerDefinition", "ProtocolMetadataBase"],
        ["KnowledgeSourceConfiguration", "BotDefinition", "BotEntity"],
    ),
    "dialogs": (
        "Dialogs (topics, tools, child agents) & recognizers",
        "`AdaptiveDialog` (a topic), `TaskDialog` (a tool/action), `AgentDialog` (a child agent) and the "
        "recognizers that route user messages. See [04 Topics & triggers](../../04-topics-and-triggers.md).",
        ["DialogBase", "RecognizerBase", "Recognizer", "IntentBase", "LocalizableContentContainer"],
        [],
    ),
    "triggers": (
        "Triggers",
        "Everything that can appear in a topic's `beginDialog:` (or `triggers:`) slot. "
        "See [04 Topics & triggers](../../04-topics-and-triggers.md).",
        ["TriggerBase"],
        [],
    ),
    "actions": (
        "Actions (topic nodes)",
        "Every `kind` that can appear in an `actions:` list inside a topic, plus their supporting types. "
        "See [05 Nodes](../../05-nodes.md).",
        ["DialogAction", "EditTableOperation", "BooleanConditionBase", "HttpRequestContent", "HttpErrorHandlingBehavior",
         "ActionInputBindingBase", "ActionOutputBindingBase", "DtmfTerminationOptions", "HoldSettings", "SearchFilesMode"],
        ["ConditionItem"],
    ),
    "task-actions": (
        "Tool (TaskDialog) actions & inputs",
        "The `action:` block of a `TaskDialog` tool, and the `inputs:` it declares. "
        "See [08 Tools & connectors](../../08-tools-and-connectors.md) and [09 MCP & multi-agent](../../09-mcp-and-multi-agent.md).",
        ["TaskAction", "TaskInput"],
        [],
    ),
    "knowledge": (
        "Knowledge sources",
        "The `source:` of a `KnowledgeSourceConfiguration` and related search settings. "
        "See [07 Knowledge](../../07-knowledge.md).",
        ["KnowledgeSource", "PublicDataSource", "BaseSearchConfiguration"],
        [],
    ),
    "entities": (
        "Entities",
        "Entities tell a `Question` node (and slot filling) what kind of value to extract. "
        "See [06 Variables, entities & Power Fx](../../06-variables-entities-powerfx.md).",
        ["EntityReference", "Entity"],
        [],
    ),
    "data-types": (
        "Data types",
        "Types used for variables, inputs, outputs and connector schemas. "
        "See [06 Variables, entities & Power Fx](../../06-variables-entities-powerfx.md).",
        ["DataType", "PropertyInfo", "PropertyInfoNoKind", "OptionSetDataType"],
        [],
    ),
    "settings": (
        "Settings",
        "Agent-level settings blocks (moderation, autonomy, voice, feedback, model selection). "
        "See [03 The agent](../../03-agent.md).",
        ["BotSettingsBase", "AgentModelType", "LatencySettingsBase", "LatencySettingsBaseLegacy"],
        [],
    ),
    "evaluation": (
        "Evaluation & testing",
        "Test sets, evaluation runs, graders and assertions. "
        "See [12 Evaluations & testing](../../12-evaluations-and-testing.md).",
        ["GraderBase", "TestAssertionBase", "TestDefinitionBase", "EvaluationTargetBase"],
        [],
    ),
    "messages": (
        "Messages, activities & cards",
        "What a `SendActivity` / `Question` prompt can contain: text, message variations, cards, "
        "quick replies and attachments. See [11 Messages & Adaptive Cards](../../11-messages-and-adaptive-cards.md).",
        ["ActivityTemplateBase", "ActivityAttachment", "QuickReplyBase", "FileReference"],
        ["ResponseTemplate", "ResponseTemplateNoKind", "TemplateLine", "TextSegment", "ExpressionSegment"],
    ),
    "other": (
        "Other definitions",
        "Supporting types that don't belong to a single category: expression wrappers, enums, metadata, "
        "and helper records.",
        [],
        [],
    ),
}
CATEGORY_ORDER = list(CATEGORIES)

# Fallback: name-substring rules for definitions not reached through a union or a single referrer.
# Ordered; first match wins.
KEYWORD_RULES: list[tuple[str, list[str]]] = [
    ("other", ["Expression", "ErrorCode", "PropertyPath", "BotElement", "DataValue", "DefinitionBase"]),
    ("evaluation", ["Grader", "Evaluation", "Test", "Sample", "ExpectedExecution", "ExpectedTool",
                    "ExpectedTopic", "MultiTurn", "Annotated"]),
    ("knowledge", ["Search", "DataSource", "Citation", "Knowledge", "Graph", "Sharepoint", "SharePoint",
                   "FileIndexing", "Summarization", "CustomModelCitation"]),
    ("messages", ["Activity", "Message", "Card", "Attachment", "TemplateSegment", "ContentShare"]),
    ("entities", ["ClosedList", "Entity", "RegexIntent"]),
    ("settings", ["Settings", "Moderation", "Speech", "Voice", "Model", "GPT", "Gpt", "Network", "Redaction",
                  "Masking", "Policy", "Roi", "Analytics", "Dtmf", "Interruption", "Language", "Tone",
                  "History", "Telemetry", "Wasm", "Sensitiv", "Approval", "Allowlist", "Capabilities"]),
    ("components", ["Bot", "Component", "Connector", "Connection", "Channel", "Application", "Registration",
                    "Audit", "Managed", "EnvironmentVariable", "Publish", "Synchroniz", "Manifest", "Skill",
                    "Mcp", "MCP", "AIPlugin", "Copilot", "Reference", "Provisioning", "A2A", "Agent",
                    "ExternalTrigger", "Conversion", "Adaptation", "Category", "CustomMetric", "UserPrivilege",
                    "SystemUser", "Dlp", "DataPolicy", "Recurrence", "ConversationStarter", "PropertyOverride",
                    "Diagnostic", "Plan"]),
    ("actions", ["Variable", "Question", "Hold", "SignInReason", "Condition", "UserInput", "TaskOutput",
                 "Initializ", "InvocationSource", "DefaultValueMissing", "IncludeSelectedTopics",
                 "MultipleTopicsMatched", "MultiEntityChoice", "PartType"]),
]


def ref_name(ref: str) -> str:
    return ref.split("/")[-1]


def union_members(defn: dict) -> list[str]:
    out = []
    for key in ("oneOf", "anyOf"):
        for alt in defn.get(key, []):
            if "$ref" in alt:
                out.append(ref_name(alt["$ref"]))
    return out


def iter_refs(node):
    """Yield every $ref name reachable inside a definition (not following refs)."""
    if isinstance(node, dict):
        for k, v in node.items():
            if k == "$ref" and isinstance(v, str):
                yield ref_name(v)
            else:
                yield from iter_refs(v)
    elif isinstance(node, list):
        for v in node:
            yield from iter_refs(v)


def is_enum(defn: dict) -> bool:
    alts = defn.get("oneOf", [])
    return defn.get("type") == "string" and alts and all("const" in a for a in alts)


def categorize(defs: dict) -> dict[str, str]:
    cat: dict[str, str] = {}

    def assign(name, c):
        if name in defs and name not in cat:
            cat[name] = c

    # 1. seed unions and their direct members (in category order, first wins)
    for c in CATEGORY_ORDER:
        _, _, seeds, extras = CATEGORIES[c]
        for s in seeds:
            assign(s, c)
        for e in extras:
            assign(e, c)
    for c in CATEGORY_ORDER:
        _, _, seeds, _ = CATEGORIES[c]
        for s in seeds:
            for m in union_members(defs.get(s, {})):
                # BotElement-style mega unions are not seeds, so this is safe
                assign(m, c)

    # 2. NoKind / Optional variants follow their base
    for name in defs:
        if name in cat:
            continue
        for suffix in ("NoKind", "OptionalKind"):
            if name.endswith(suffix) and name[: -len(suffix)] in cat:
                cat[name] = cat[name[: -len(suffix)]]

    # 3. propagate: an uncategorised def referenced only from one category joins it
    referrers: dict[str, set[str]] = defaultdict(set)
    for name, defn in defs.items():
        if name == "BotElement":
            continue
        for r in iter_refs(defn):
            if r != name:
                referrers[r].add(name)
    changed = True
    while changed:
        changed = False
        for name in defs:
            if name in cat:
                continue
            cats = {cat.get(r) for r in referrers.get(name, ())}
            if len(cats) == 1:
                only = cats.pop()
                if only is not None:
                    cat[name] = only
                    changed = True

    # 4. keyword fallback, then NoKind variants follow their base again
    for name in defs:
        if name in cat:
            continue
        base = name
        for suffix in ("NoKind", "OptionalKind"):
            if name.endswith(suffix):
                base = name[: -len(suffix)]
        if base in cat:
            cat[name] = cat[base]
            continue
        for c, words in KEYWORD_RULES:
            if any(w in base for w in words):
                cat[name] = c
                break

    for name in defs:
        cat.setdefault(name, "other")
    return cat


class BotSchemaRenderer:
    def __init__(self, defs: dict):
        self.defs = defs
        self.cat = categorize(defs)
        self.referrers: dict[str, set[str]] = defaultdict(set)
        for name, defn in defs.items():
            if name == "BotElement":
                continue
            for r in iter_refs(defn):
                if r != name:
                    self.referrers[r].add(name)

    # -- links ------------------------------------------------------------
    def link(self, name: str, from_cat: str | None) -> str:
        if name not in self.defs:
            return f"`{name}`"
        page = self.cat[name]
        target = f"#{slug(name)}" if page == from_cat else f"{page}.md#{slug(name)}"
        return f"[{name}]({target})"

    # -- types ------------------------------------------------------------
    def type_str(self, s: dict, here: str) -> str:
        if not isinstance(s, dict) or not s:
            return "any"
        if "$ref" in s:
            return self.link(ref_name(s["$ref"]), here)
        if "const" in s:
            return f"`{json.dumps(s['const'])}`"
        for key in ("oneOf", "anyOf"):
            if key in s:
                parts = []
                for alt in s[key]:
                    t = self.type_str(alt, here)
                    if t not in parts:
                        parts.append(t)
                return " \\| ".join(parts)
        if "allOf" in s:
            return " & ".join(self.type_str(a, here) for a in s["allOf"])
        t = s.get("type")
        if t == "array":
            return f"list of {self.type_str(s.get('items', {}), here)}"
        if t == "object":
            ap = s.get("additionalProperties")
            if isinstance(ap, dict):
                return f"map of name → {self.type_str(ap, here)}"
            return "object"
        if t == "string" and s.get("pattern") == "=.*":
            return "`=` expression"
        if t == "string" and s.get("pattern"):
            return f"string (pattern `{s['pattern']}`)"
        if isinstance(t, list):
            return " \\| ".join(f"`{x}`" for x in t)
        if t:
            return f"`{t}`"
        return "any"

    # -- one definition ---------------------------------------------------
    def render_def(self, name: str) -> str:
        d = self.defs[name]
        here = self.cat[name]
        out = [f"### {name}", ""]

        title, desc = d.get("title"), d.get("description")
        if title and not is_placeholder(title) and title != name:
            out.append(f"*Title:* {cell(title)}  ")
        if desc and not is_placeholder(desc):
            out.append(cell(desc) + "  ")

        props = d.get("properties", {})
        kind_const = props.get("kind", {}).get("const")
        if kind_const:
            out.append(f"YAML: `kind: {kind_const}`")
        elif props and name.endswith("NoKind"):
            out.append(f"Same shape as {self.link(name[:-6], here)} but written **without** a `kind:` line "
                       "(used where the type is implied by the parent property).")

        if is_enum(d):
            vals = ", ".join(f"`{a['const']}`" for a in d["oneOf"])
            out += ["", f"**String enum.** Allowed values: {vals}"]
        elif "oneOf" in d or "anyOf" in d:
            members = d.get("oneOf") or d.get("anyOf")
            parts = [self.type_str(a, here) for a in members]
            if len(parts) > 40:
                out += ["", f"**One of {len(parts)} types** (see the [A–Z index](index.md))."]
            else:
                out += ["", "**One of:** " + ", ".join(parts)]
        elif d.get("type") and not props:
            out += ["", f"**Type:** {self.type_str(d, here)}"]

        if props:
            required = set(d.get("required", []))
            out += ["", "| Property | Type | Req. | Default | Notes |", "|---|---|---|---|---|"]
            for pname, ps in props.items():
                if pname == "kind":
                    continue
                notes = []
                if ps.get("$role") == "expression":
                    notes.append("accepts a `=` Power Fx expression")
                ptitle = ps.get("title")
                if ptitle and not is_placeholder(ptitle) and ptitle.lower() != pname.lower():
                    notes.append(f"title: {ptitle}")
                pdesc = ps.get("description")
                if pdesc and not is_placeholder(pdesc):
                    notes.append(pdesc)
                default = ps.get("default")
                out.append(
                    f"| `{pname}` | {self.type_str(ps, here)} | {'✓' if pname in required else ''} | "
                    f"{'' if default is None else '`' + cell(json.dumps(default)) + '`'} | {cell('; '.join(notes))} |"
                )
            if d.get("additionalProperties") is False:
                out += ["", "_No other properties are allowed._"]

        refs = sorted(self.referrers.get(name, ()))
        if refs:
            shown = ", ".join(self.link(r, here) for r in refs[:20])
            more = f" … and {len(refs) - 20} more" if len(refs) > 20 else ""
            out += ["", f"<sub>Used by: {shown}{more}</sub>"]
        out.append("")
        return "\n".join(out)

    # -- pages ------------------------------------------------------------
    def render_pages(self) -> dict[str, str]:
        pages: dict[str, str] = {}
        by_cat: dict[str, list[str]] = defaultdict(list)
        for name in self.defs:
            by_cat[self.cat[name]].append(name)

        for c in CATEGORY_ORDER:
            title, intro, _, _ = CATEGORIES[c]
            names = sorted(by_cat[c], key=str.lower)
            kinds = [n for n in names if "kind" in self.defs[n].get("properties", {})]
            toc = ", ".join(self.link(n, c) for n in names)
            body = [
                header(title, "reference/bot.schema.yaml-authoring.json", intro),
                READING_GUIDE,
                f"**{len(names)} definitions** on this page ({len(kinds)} with a `kind:`). "
                "[A–Z index of all definitions](index.md).",
                "",
                "<details><summary>Contents</summary>",
                "",
                toc,
                "",
                "</details>",
                "",
            ]
            for n in names:
                body.append(self.render_def(n))
            pages[f"{c}.md"] = "\n".join(body).rstrip() + "\n"

        # index
        rows = ["| Definition | Page | `kind:` | Shape |", "|---|---|---|---|"]
        for n in sorted(self.defs, key=str.lower):
            d = self.defs[n]
            props = d.get("properties", {})
            shape = ("enum" if is_enum(d) else
                     "object" if props else
                     "union" if ("oneOf" in d or "anyOf" in d) else
                     d.get("type", ""))
            k = "✓" if "kind" in props else ""
            rows.append(f"| {self.link(n, None)} | {CATEGORIES[self.cat[n]][0]} | {k} | {shape} |")
        cat_list = "\n".join(
            f"- [{CATEGORIES[c][0]}]({c}.md) — {sum(1 for n in self.defs if self.cat[n] == c)} definitions"
            for c in CATEGORY_ORDER
        )
        pages["index.md"] = "\n".join([
            header("Copilot Studio YAML schema — A–Z index", "reference/bot.schema.yaml-authoring.json",
                   f"All **{len(self.defs)}** definitions in the Copilot Studio YAML authoring schema. "
                   "A file's root must be one of the root kinds the schema allows "
                   f"({', '.join(str(self.link(ref_name(a['$ref']), None)) for a in ROOT_ONEOF)}); "
                   "other files (agent, knowledge, variables, settings) use component kinds "
                   "that the authoring tools validate separately."),
            READING_GUIDE,
            "## Pages",
            "",
            cat_list,
            "",
            "## All definitions",
            "",
            *rows,
            "",
        ])
        return pages


READING_GUIDE = """\
**How to read these tables.** *Type* links to another definition when the value is itself a structured object.
`=` expression means the value is a [Power Fx](../../glossary.md#power-fx) formula that starts with `=`.
Properties marked **Req.** are required by the schema. A union ("one of") means you pick
exactly one of the listed shapes, usually by setting `kind:`. Descriptions in the source schema are
placeholder tokens, so the hand-written guides in [`docs/`](../../README.md) explain *what* each kind does.
"""

ROOT_ONEOF: list = []


# --------------------------------------------------------------------------- adaptive card schema

def render_adaptive_card(schema: dict) -> str:
    defs = schema["definitions"]

    def link(name):
        return f"[{name}](#{slug(name)})" if name in defs else f"`{name}`"

    def type_str(s):
        if not isinstance(s, dict) or not s:
            return "(inherited)"
        if "$ref" in s:
            return link(ref_name(s["$ref"]))
        if "enum" in s:
            return " \\| ".join(f"`{v}`" for v in s["enum"])
        if "const" in s:
            return f"`{s['const']}`"
        for key in ("anyOf", "oneOf"):
            if key in s:
                parts = [type_str(a) for a in s[key] if a.get("type") != "null"]
                return " \\| ".join(dict.fromkeys(parts))
        if "allOf" in s:
            return " & ".join(type_str(a) for a in s["allOf"])
        t = s.get("type")
        if t == "array":
            return f"list of {type_str(s.get('items', {}))}"
        if isinstance(t, list):
            return " \\| ".join(f"`{x}`" for x in t if x != "null")
        if t:
            return f"`{t}`"
        return "any"

    def collect_props(name, seen=None):
        """Own properties plus those inherited through allOf, own first."""
        seen = seen or set()
        if name in seen or name not in defs:
            return {}, set()
        seen.add(name)
        d = defs[name]
        props, req = {}, set(d.get("required", []))
        for p, ps in d.get("properties", {}).items():
            if ps:  # empty {} means "declared in a base"
                props[p] = (ps, name)
        for base in d.get("allOf", []):
            if "$ref" in base:
                bprops, breq = collect_props(ref_name(base["$ref"]), seen)
                for p, v in bprops.items():
                    props.setdefault(p, v)
                req |= breq
        return props, req

    groups = [
        ("The card", ["AdaptiveCard"]),
        ("Card elements (layout & display)",
         [n for n in defs if n in ("ActionSet", "Column", "ColumnSet", "Container", "Fact", "FactSet", "Image",
                                   "ImageSet", "Media", "MediaSource", "CaptionSource", "RichTextBlock", "TextRun",
                                   "Table", "TableCell", "TableColumnDefinition", "TableRow", "TextBlock")]),
        ("Inputs", [n for n in defs if n.startswith("Input.") or n == "Data.Query"]),
        ("Actions (buttons)", [n for n in defs if n.startswith("Action.")] + ["TargetElement"]),
        ("Enumerations & style values", [n for n in defs if "anyOf" in defs[n] and not n.startswith("ImplementationsOf.")
                                         and n not in ("TextRun", "TargetElement")]),
        ("Card-level extras", ["AuthCardButton", "Authentication", "BackgroundImage", "Metadata", "Refresh",
                               "TokenExchangeResource"]),
        ("Abstract groups (\"any element\", \"any action\", …)",
         [n for n in defs if n.startswith("ImplementationsOf.") or n.startswith("Extendable.")]),
    ]
    placed = {n for _, ns in groups for n in ns}
    rest = [n for n in defs if n not in placed]
    if rest:
        groups.append(("Other", rest))

    out = [header(
        "Adaptive Card schema reference",
        "reference/adaptive-card.schema.json",
        "An [Adaptive Card](../glossary.md#adaptive-card) is a platform-neutral JSON description of a small UI "
        "(text, images, inputs, buttons) that the chat client draws natively. In Copilot Studio YAML you embed "
        "the card JSON inside `AdaptiveCardTemplate` (in a message) or `AdaptiveCardPrompt` (to collect input). "
        "This page lists every element in the schema bundled with the skills repo "
        f"(`{schema.get('id', '')}`). For how to use cards in topics see "
        "[11 Messages & Adaptive Cards](../11-messages-and-adaptive-cards.md).",
    ),
        "Every element is a JSON object with a `\"type\"` property naming the element (e.g. `\"type\": \"TextBlock\"`). "
        "*Version* is the minimum Adaptive Card schema version that supports the element or property; "
        "this bundled schema is version **1.6**, while the skills' examples declare `\"version\": \"1.5\"` "
        "for broad client support (see the [card guide](../11-messages-and-adaptive-cards.md)).",
        "",
    ]
    for gtitle, names in groups:
        out += [f"## {gtitle}", ""]
        for n in names:
            d = defs[n]
            out += [f"### {n}", ""]
            if d.get("description"):
                out.append(cell(d["description"]) + "  ")
            if d.get("version"):
                out.append(f"*Version:* {d['version']}  ")
            if "anyOf" in d and not d.get("properties"):
                # skip the case-insensitive regex alternates; they just re-state the enum values
                alts = [type_str(a) for a in d["anyOf"] if "pattern" not in a]
                out += ["", "**One of:** " + ", ".join(alts) + " (case-insensitive)"]
            props, req = collect_props(n)
            if props:
                out += ["", "| Property | Type | Req. | Default | Version | Description |", "|---|---|---|---|---|---|"]
                for p, (ps, owner) in props.items():
                    inh = "" if owner == n else f" *(from {link(owner)})*"
                    default = ps.get("default")
                    out.append(
                        f"| `{p}` | {type_str(ps)} | {'✓' if p in req else ''} | "
                        f"{'' if default is None else '`' + cell(json.dumps(default)) + '`'} | "
                        f"{ps.get('version', '')} | {cell(ps.get('description', ''))}{inh} |"
                    )
            out.append("")
    return "\n".join(out).rstrip() + "\n"


# --------------------------------------------------------------------------- connectors

def conn_type(t) -> str:
    if isinstance(t, str):
        return f"`{t}`"
    if isinstance(t, dict):
        if "type" in t and isinstance(t["type"], (str, dict)):
            return conn_type(t["type"])
        kind = t.get("kind", "Record")
        n = len(t.get("properties", {}) or {})
        return f"`{kind}`" + (f" ({n} fields)" if n else "")
    return "`?`"


def render_param_table(props: dict, inputs: bool) -> list[str]:
    if not props:
        return ["_None._", ""]
    if inputs:
        rows = ["| Name | Display name | Type | Req. | Description |", "|---|---|---|---|---|"]
    else:
        rows = ["| Name | Display name | Type | Description |", "|---|---|---|---|"]
    items = sorted(props.items(), key=lambda kv: (kv[1] or {}).get("order", 999) if isinstance(kv[1], dict) else 999)
    for name, p in items:
        p = p if isinstance(p, dict) else {"type": p}
        extra = []
        if p.get("visibility") == "Advanced":
            extra.append("*advanced*")
        if p.get("dynamicValuesConfig"):
            extra.append("*values looked up dynamically*")
        desc = cell(" ".join([p.get("description", "")] + extra))
        if inputs:
            rows.append(f"| `{name}` | {cell(p.get('displayName', ''))} | {conn_type(p)} | "
                        f"{'✓' if p.get('isRequired') else ''} | {desc} |")
        else:
            rows.append(f"| `{name}` | {cell(p.get('displayName', ''))} | {conn_type(p)} | {desc} |")
    return rows + [""]


def render_connector(path: Path, data: dict) -> str:
    api = path.stem
    ops = data.get("operations", []) or []
    trig = data.get("triggers", []) or []
    out = [header(
        f"Connector: {data.get('displayName', api)} (`{api}`)",
        f"reference/connectors/{path.name}",
        cell(data.get("description", "")),
    ),
        "| | |", "|---|---|",
        f"| Connector id | `{data.get('connectorId', '')}` |",
        f"| Type | {data.get('connectorType', '')} |",
        f"| Operations | {len(ops)} |",
        f"| Triggers (used by flows / external triggers, not by tools) | {len(trig)} |",
        "",
        "Use an operation's **operationId** as `action.operationId` in a `TaskDialog` tool "
        "(see [08 Tools & connectors](../../08-tools-and-connectors.md)). Inputs map to `inputs:` entries "
        "(`propertyName` = the input *Name* below). Input names that start with `$` are "
        "[OData](../../glossary.md#odata) query options and need special quoting — see "
        "[`$`-prefixed inputs](../../08-tools-and-connectors.md#-prefixed-inputs-odata-parameters).",
        "",
        "## Operations",
        "",
        "| operationId | Name | Description |",
        "|---|---|---|",
    ]
    for o in ops:
        out.append(f"| [`{o['operationId']}`](#{slug(o['operationId'])}) | {cell(o.get('displayName', ''))} | "
                   f"{cell(o.get('description', ''))} |")
    out.append("")
    for o in ops:
        out += [f"### {o['operationId']}", "", f"**{cell(o.get('displayName', ''))}** — {cell(o.get('description', ''))}", ""]
        it = o.get("inputType") or {}
        out += ["**Inputs**", ""] + render_param_table((it.get("properties") if isinstance(it, dict) else None) or {}, True)
        ot = o.get("outputType")
        if isinstance(ot, dict) and ot.get("properties"):
            out += [f"**Output** — {conn_type(ot)}:", ""] + render_param_table(ot["properties"], False)
        else:
            note = " (shape depends on the list/table chosen at design time)" if ot == "UnresolvedDynamicType" else ""
            out += [f"**Output** — {conn_type(ot) if ot else '_none_'}{note}", ""]
    if trig:
        out += ["## Triggers", "",
                "Connector triggers start a flow or an external-trigger when something happens in the "
                "connected service. See [10 Flows & autonomous triggers](../../10-flows-and-autonomous-triggers.md).",
                "", "| triggerId | Name | Description |", "|---|---|---|"]
        for t in trig:
            out.append(f"| `{t.get('triggerId') or t.get('operationId', '')}` | {cell(t.get('displayName', ''))} | {cell(t.get('description', ''))} |")
        out.append("")
    return "\n".join(out).rstrip() + "\n"


# --------------------------------------------------------------------------- main

def write(path: Path, text: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(text, encoding="utf-8")


def main() -> None:
    bot = json.loads(BOT_SCHEMA.read_text(encoding="utf-8"))
    ROOT_ONEOF[:] = bot.get("oneOf", [])
    renderer = BotSchemaRenderer(bot["definitions"])
    for fname, text in renderer.render_pages().items():
        write(SCHEMA_OUT / fname, text)

    card = json.loads(CARD_SCHEMA.read_text(encoding="utf-8"))
    write(OUT / "adaptive-card-schema.md", render_adaptive_card(card))

    rows = ["| Connector | API name | Operations | Description |", "|---|---|---|---|"]
    for path in sorted(CONNECTORS.glob("*.yml")):
        data = yaml.safe_load(path.read_text(encoding="utf-8"))
        write(CONNECTOR_OUT / f"{path.stem}.md", render_connector(path, data))
        rows.append(f"| [{data.get('displayName', path.stem)}]({path.stem}.md) | `{path.stem}` | "
                    f"{len(data.get('operations', []) or [])} | {cell(data.get('description', ''))[:200]} |")
    write(CONNECTOR_OUT / "index.md", "\n".join([
        header("Bundled connector references", "reference/connectors/*.yml",
               "The skills repo includes full operation catalogues for these four "
               "[connectors](../../glossary.md#connector). Any other connector works the same way; "
               "you just need its operation ids and parameters from somewhere else (for example the "
               "Copilot Studio UI or the connector's documentation). "
               "See [08 Tools & connectors](../../08-tools-and-connectors.md)."),
        *rows, "",
    ]))

    counts = defaultdict(int)
    for n in bot["definitions"]:
        counts[renderer.cat[n]] += 1
    print("schema pages:", dict(counts))


if __name__ == "__main__":
    main()
