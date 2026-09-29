# AGENTS.md — plugin-ollama

Standalone plugin repo owning the `charly ollama` management CLI
(`command:ollama`). The plugin is a Go module at `candy/plugin-ollama/` (module
path `github.com/opencharly/plugin-ollama/candy/plugin-ollama`); the root
`charly.yml` declares `discover: candy` (so the repo is a project and its candy is
scanned) and carries the embedded `ollama-cli-skill:` skill entity.

Canonical files:

- `candy/plugin-ollama/charly.yml` — the `plugin-ollama:` candy entity (`plugin:`
  block, `plan:` check).
- `candy/plugin-ollama/` — the Go source: `plugin.go`, `provider.go`,
  `command.go` (the CLI tree + HTTP API calls), `schema/ollama.cue`,
  `cmd/serve/main.go`.
- `charly.yml` — the root manifest (`discover: candy`) + the `ollama-cli-skill:`
  skill entity.
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.
- `README.md` — user overview only; never agent guidance.

## Load these skills first (R0)

- `/charly-ollama:ollama-cli` — the `charly ollama` CLI reference (projected
  from this candy's own `ollama-cli-skill:` entity). Load before changing the
  command tree or the endpoint resolution.
- `/charly-internals:plugin` — the plugin authoring reference: the `plugin:`
  block, the `command` provider class, the per-plugin CUE-schema contract.
- `/charly-ollama:ollama` / `/charly-ollama:ollama-layer` — the server the CLI
  manages.
- `/charly-internals:git-workflow` — before any git/PR action.

## Build / validate / test

- `go build ./...` in `candy/plugin-ollama/` — compile the plugin module.
- `go test ./...` in `candy/plugin-ollama/` — the plugin's Go tests (the command
  tree + the schema-serve seam).
- `charly box validate` at the repo root — the structural check (the candy +
  `plugin:` block, CUE schema).
- The merge gate is the **org-wide** `charly/pr-validator` (required check
  `validate / validate`, defined in `opencharly/.github`); this repo has **no**
  per-repo candy gate.
- The live R10 witness is the disposable `check-ollama-pod` bed (in `box/fedora`).

## Modify this repo

- Edit the `plugin-ollama:` candy entity, the Go source, and `schema/ollama.cue`
  **together** — the schema is the served declaration surface.
- Keep the command **self-contained over plain HTTP** — it must work identically
  compiled-in or out-of-process; do not introduce a reverse channel.
- Keep the `ollama-cli-skill:` entity in step with any command-tree change — it
  is the projected source for `/charly-ollama:ollama-cli`.

## Landing

Load `/charly-internals:git-workflow` before any git/PR action; it owns the
landing mechanics. The authoritative rulebook is the umbrella `AGENTS.md` in
`opencharly/opencharly` and `charly/AGENTS.md` in the charly repo.
