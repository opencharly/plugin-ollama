# plugin-ollama

The `charly ollama` management CLI for OpenCharly — a compiled-in command plugin
that manages a DEPLOYED Ollama server from the host.

The plugin owns the entire CLI: the hand-rolled subcommand grammar AND the Ollama
HTTP API calls, built on the official upstream Go client
(`github.com/ollama/ollama/api`, pinned in this candy's own `go.mod`). There is no
core Ollama logic and no host-build seam. It talks to the server over plain HTTP,
so it works identically compiled-in or out-of-process.

## What it provides

| Capability | Surface |
|---|---|
| `command:ollama` | the `charly ollama` CLI: `list` / `ps` / `pull` / `rm` / `show` / `cp` / `create` / `push` / `run` / `stop` / `version` |

Endpoint resolution: `--server <url>` flag > `OLLAMA_HOST` env >
`http://127.0.0.1:11434`. A schemeless value gets `http://` prepended; `0.0.0.0`
(the candy's server-bind value) maps to `127.0.0.1` for client use.

## How to use it

Point the CLI at a running server:

```bash
charly ollama list
charly ollama pull llama3
charly ollama run llama3 "hello"
charly ollama version --server http://127.0.0.1:11434
```

Every leaf takes `--server <url>` anywhere in the args. A non-zero server
response propagates as the command's own non-zero exit. Full Modelfile authoring
stays with the upstream CLI; `charly ollama create` covers the lightweight
`--from` form.

## Three disjoint surfaces

- This plugin — manages the server from the HOST over HTTP.
- `candy/ollama` — deploys the server (supervisord `ollama serve` on 11434).
- The candy's host shell `alias: ollama` (`charly alias install ollama`) — execs
  the upstream CLI INSIDE the container.

## Layout

- `candy/plugin-ollama/` — the plugin module: `plugin.go` (provider + meta),
  `provider.go`, `command.go`, `schema/ollama.cue` (the self-contained
  `#OllamaPlugin`), and `cmd/serve/main.go`.
- `charly.yml` — the root project manifest (`discover: candy`) + the embedded
  `ollama-cli-skill:` skill entity.
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.

## R10 witness

The disposable `check-ollama-pod` bed (in `box/fedora`) deploys the Ollama image
and runs `charly ollama list|version --server ...` host-side against the live
server, plus the candy's baked checks.

## Related

- Owning skill: `/charly-ollama:ollama-cli` — the `charly ollama` CLI reference
  (projected from this candy's own `ollama-cli-skill:` entity).
- `/charly-ollama:ollama` / `/charly-ollama:ollama-layer` — the server box and
  candy.
- `/charly-internals:plugin` — the plugin/provider model.
- [`opencharly/charly`](https://github.com/opencharly/charly) — the charly CLI.
