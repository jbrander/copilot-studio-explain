#!/usr/bin/env python3
"""Verify the hand-written docs in docs/.

Checks:
  1. YAML snippets: every ```yaml block parses. If its top-level mapping (or each item of a top-level list)
     has a `kind:` that names a schema definition, it is validated against that definition.
     Blocks whose first line contains `# fragment` are only parsed, not validated.
     Blocks whose first line contains `# invalid` are expected to FAIL validation (used for "wrong" examples).
  2. Control: every vendor template / eval fixture is validated the same way, to show the validator
     accepts known-good YAML (reported separately; they're not our files).
  3. Links: relative links resolve to files; #anchors resolve to headings (GitHub slug rules).
  4. Kinds: every `kind: X` in a YAML block and every backticked CamelCase word that looks like a kind
     exists in the schema (catches typos / invented kinds).
  5. Glossary: every glossary entry is linked from at least one other page.

  6. (--external) every external http(s) link returns HTTP 200 (uses curl; slow).

Usage:  python3 -I tools/check_docs.py [--verbose] [--external]
Exit code 1 if any check fails.
"""
from __future__ import annotations

import json
import re
import sys
from pathlib import Path

import yaml
from jsonschema import Draft7Validator
from referencing import Registry, Resource

ROOT = Path(__file__).resolve().parent.parent
DOCS = ROOT / "docs"
VENDOR = ROOT / "vendor" / "skills-for-copilot-studio"
SCHEMA = json.loads((VENDOR / "reference" / "bot.schema.yaml-authoring.json").read_text())
DEFS = SCHEMA["definitions"]
VERBOSE = "--verbose" in sys.argv
EXTERNAL = "--external" in sys.argv

# Keys the authoring tools add/allow outside the schema.
STRIP_KEYS = {"mcs.metadata"}



def _oneof_to_anyof(node):
    """The schema uses oneOf for unions whose branches overlap (e.g. DialogExpression = any string | "=.*"),
    so strict oneOf rejects valid values like `dialog: =Switch(...)`. The vendor's own validator accepts them;
    treat oneOf as anyOf here."""
    if isinstance(node, dict):
        return {("anyOf" if k == "oneOf" else k): _oneof_to_anyof(v) for k, v in node.items()}
    if isinstance(node, list):
        return [_oneof_to_anyof(v) for v in node]
    return node


registry = Registry().with_resource("urn:bot", Resource.opaque(_oneof_to_anyof(SCHEMA)))
_validators: dict[str, Draft7Validator] = {}


def validator_for(kind: str) -> Draft7Validator:
    if kind not in _validators:
        _validators[kind] = Draft7Validator({"$ref": f"urn:bot#/definitions/{kind}"}, registry=registry)
    return _validators[kind]


def validate_obj(obj) -> list[str]:
    """Return a list of error strings ('' list if valid or not validatable)."""
    if isinstance(obj, list):
        errs = []
        for i, item in enumerate(obj):
            errs += [f"[{i}] {e}" for e in validate_obj(item)]
        return errs
    if not isinstance(obj, dict):
        return []
    kind = obj.get("kind")
    if not isinstance(kind, str) or kind not in DEFS:
        return []
    obj = {k: v for k, v in obj.items() if k not in STRIP_KEYS}
    obj = file_format_to_schema(kind, obj)
    errors = sorted(validator_for(kind).iter_errors(obj), key=lambda e: list(e.absolute_path))
    out = []
    for e in errors[:5]:
        path = "/".join(str(p) for p in e.absolute_path) or "(root)"
        msg = e.message if len(e.message) < 300 else e.message[:300] + "…"
        out.append(f"{kind} @ {path}: {msg}")
    return out


def file_format_to_schema(kind: str, obj: dict) -> dict:
    """Undo the flattening the authoring tools apply to some component files on disk.

    - agent.mcs.yml (GptComponentMetadata) carries the component's `displayName` at the root.
    - variables/*.mcs.yml (GlobalVariableComponent) put the schema's nested `variable:` fields
      (name, scope, aIVisibility, defaultValue, ...) at the root.
    """
    if kind == "GptComponentMetadata":
        obj = {k: v for k, v in obj.items() if k != "displayName"}
    elif kind == "GlobalVariableComponent" and "variable" not in obj:
        var_props = set(DEFS["VariableNoKind"]["properties"])
        nested = {k: v for k, v in obj.items() if k in var_props}
        obj = {k: v for k, v in obj.items() if k not in var_props}
        obj["variable"] = nested
    return obj


def iter_kinds(obj):
    if isinstance(obj, dict):
        k = obj.get("kind")
        if isinstance(k, str):
            yield k
        for v in obj.values():
            yield from iter_kinds(v)
    elif isinstance(obj, list):
        for v in obj:
            yield from iter_kinds(v)


# --------------------------------------------------------------------------- markdown helpers

FENCE = re.compile(r"^```(\w*)[^\n]*\n(.*?)^```", re.M | re.S)
LINK = re.compile(r"(?<!!)\[[^\]]*\]\(([^)\s]+)\)")
HEADING = re.compile(r"^(#{1,6})\s+(.*?)\s*#*\s*$", re.M)
HTML_ANCHOR = re.compile(r'<a\s+(?:name|id)="([^"]+)"')


def gh_slug(text: str) -> str:
    text = re.sub(r"`([^`]*)`", r"\1", text)          # inline code
    text = re.sub(r"\[([^\]]*)\]\([^)]*\)", r"\1", text)  # links
    s = text.strip().lower()
    s = re.sub(r"[^\w\- ]", "", s)
    return s.replace(" ", "-")


_anchor_cache: dict[Path, set[str]] = {}


def anchors(path: Path) -> set[str]:
    if path not in _anchor_cache:
        text = path.read_text(encoding="utf-8")
        text_nocode = FENCE.sub("", text)
        seen: dict[str, int] = {}
        out = set()
        for _, title in HEADING.findall(text_nocode):
            base = gh_slug(title)
            n = seen.get(base, 0)
            out.add(base if n == 0 else f"{base}-{n}")
            seen[base] = n + 1
        out |= set(HTML_ANCHOR.findall(text))
        _anchor_cache[path] = out
    return _anchor_cache[path]


# --------------------------------------------------------------------------- checks

def main() -> int:
    failures: list[str] = []
    warnings: list[str] = []
    md_files = sorted(DOCS.rglob("*.md"))
    hand_written = [p for p in md_files if "reference" not in p.relative_to(DOCS).parts]

    # 1 + 4. YAML blocks
    n_blocks = n_validated = 0
    for p in hand_written:
        text = p.read_text(encoding="utf-8")
        for m in FENCE.finditer(text):
            lang, body = m.group(1).lower(), m.group(2)
            if lang not in ("yaml", "yml"):
                continue
            n_blocks += 1
            line = text[: m.start()].count("\n") + 1
            where = f"{p.relative_to(ROOT)}:{line}"
            first = body.splitlines()[0] if body.strip() else ""
            # docs use `...` / `# …` ellipses in fragments; make them parseable
            try:
                data = yaml.safe_load(body)
            except yaml.YAMLError as e:
                failures.append(f"YAML parse error {where}: {str(e).splitlines()[0]}")
                continue
            for k in set(iter_kinds(data)):
                if k not in DEFS and not k.startswith("<"):
                    failures.append(f"Unknown kind `{k}` in {where}")
            if "# fragment" in first:
                continue
            errs = validate_obj(data)
            n_validated += 1
            if "# invalid" in first:
                if not errs:
                    failures.append(f"Block marked '# invalid' unexpectedly validates: {where}")
                continue
            for e in errs:
                failures.append(f"Schema error {where}: {e}")

    # 2. control: vendor templates & fixtures
    control_ok = control_bad = 0
    for p in sorted(list(VENDOR.glob("templates/**/*.yml")) + list(VENDOR.glob("evals/fixtures/**/*.yml"))):
        try:
            data = yaml.safe_load(p.read_text(encoding="utf-8").split("\n//")[0])
        except yaml.YAMLError as e:
            warnings.append(f"[control] parse error {p.relative_to(VENDOR)}: {str(e).splitlines()[0]}")
            control_bad += 1
            continue
        errs = validate_obj(data)
        if errs:
            control_bad += 1
            for e in errs:
                warnings.append(f"[control] {p.relative_to(VENDOR)}: {e}")
        else:
            control_ok += 1

    # 3. links
    n_links = 0
    for p in md_files:
        text = FENCE.sub("", p.read_text(encoding="utf-8"))
        text = re.sub(r"`[^`\n]*`", "", text)
        for target in LINK.findall(text):
            if re.match(r"^[a-z]+:", target):
                continue
            n_links += 1
            path_part, _, frag = target.partition("#")
            dest = (p.parent / path_part).resolve() if path_part else p
            where = f"{p.relative_to(ROOT)} → {target}"
            if not dest.exists():
                failures.append(f"Broken link {where}")
                continue
            if frag and dest.suffix == ".md" and frag not in anchors(dest):
                failures.append(f"Missing anchor {where}")

    # 4b. backticked identifiers that look like kinds
    kind_like = re.compile(r"`((?:On|Invoke|Search|Send|Set|Begin|End|Cancel|Edit|Adaptive|Task|Agent|Gpt|Knowledge|"
                           r"Global|Question|Condition|Http|Log|Create|Answer|Replace|Repeat|Transfer|Clear|Reset|"
                           r"Parse|Emit|Foreach|Break|Continue|Goto|Recognize|Activate|Disable|Wait|OAuth|CSAT|"
                           r"Automatic|Manual|Public|SharePoint|Dataverse|Azure|Graph|Bing|Custom|File|Email|Teams|"
                           r"Meeting|Fabric|Federated|Cloud|External|Recurrence|Workflow|Connector|Connection|"
                           r"Model|Closed|Regex|Dynamic|Single|Delete|Update|Get|Sign|Log)[A-Za-z0-9]*)`")
    non_kinds = {"OnOutgoingMessage"}  # documented as broken but exists; keep list for exceptions
    for p in hand_written:
        text = p.read_text(encoding="utf-8")
        for m in kind_like.finditer(text):
            name = m.group(1)
            if name in DEFS or name in non_kinds:
                continue
            if re.fullmatch(r"[A-Z][a-z]+[A-Z]\w*", name) and VERBOSE:
                warnings.append(f"Backticked name not in schema (ok if not a kind): `{name}` in {p.relative_to(ROOT)}")

    # 5. glossary coverage
    gloss = DOCS / "glossary.md"
    if gloss.exists():
        entries = {a for a in anchors(gloss) if a not in ("glossary",)}
        top = {gh_slug(t) for lvl, t in HEADING.findall(FENCE.sub("", gloss.read_text())) if len(lvl) <= 2}
        entries -= top
        linked = set()
        for p in md_files:
            if p == gloss:
                continue
            for target in LINK.findall(p.read_text(encoding="utf-8")):
                if "glossary.md#" in target:
                    linked.add(target.split("#", 1)[1])
        for e in sorted(entries - linked):
            warnings.append(f"Glossary entry never linked from another page: #{e}")

    # 6. external links
    if EXTERNAL:
        import subprocess
        urls: dict[str, list[str]] = {}
        for p in hand_written:
            for target in LINK.findall(FENCE.sub("", p.read_text(encoding="utf-8"))):
                if target.startswith("http"):
                    urls.setdefault(target.split("#")[0], []).append(str(p.relative_to(ROOT)))
        for url, where in sorted(urls.items()):
            code = subprocess.run(["curl", "-s", "-o", "/dev/null", "-L", "--max-time", "20", "-w", "%{http_code}", url],
                                  capture_output=True, text=True).stdout
            if code != "200":
                failures.append(f"External link {url} returned {code} (in {', '.join(sorted(set(where)))})")
        print(f"External URLs checked: {len(urls)}")

    print(f"YAML blocks: {n_blocks} parsed, {n_validated} validated against the schema")
    print(f"Links checked: {n_links}")
    print(f"Control (vendor templates/fixtures): {control_ok} valid, {control_bad} with schema errors")
    for w in warnings:
        print("WARN ", w)
    for f in failures:
        print("FAIL ", f)
    print(f"\n{len(failures)} failure(s), {len(warnings)} warning(s)")
    return 1 if failures else 0


if __name__ == "__main__":
    sys.exit(main())
