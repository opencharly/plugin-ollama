// Package ollama is the charly plugin OWNING the `charly ollama` command — the host-side
// management CLI for a DEPLOYED Ollama server (list / ps / pull / rm / show / cp / create /
// push / run / stop / version). The plugin owns the ENTIRE logic: the subcommand grammar AND
// the Ollama HTTP API calls, built on the official upstream client
// (github.com/ollama/ollama/api). There is no core ollama logic and no HostBuild seam.
//
// The command talks to the server over plain HTTP, so it is fully self-contained and works
// IDENTICALLY compiled-in OR out-of-process: compiled-in (charly.yml compiled_plugins) the
// host dispatches Invoke(OpRun) in-process (dispatchInProcCommand); out-of-process the host
// fork/execs cmd/serve → CliMain running the SAME runOllamaCLI.
//
// Endpoint resolution is deliberately lightweight: --server flag > OLLAMA_HOST env >
// http://127.0.0.1:11434. The ollama CANDY (candy/ollama) deploys the server; this plugin
// manages it — the two are independent surfaces (the candy's host shell `alias: ollama`
// execs the upstream CLI inside the container; this command runs on the host).
package ollama

import (
	"embed"
	"fmt"
	"os"

	"github.com/opencharly/sdk"
	pb "github.com/opencharly/spec/proto"
)

//go:embed schema/*.cue
var schemaFS embed.FS

// NewProvider returns the ollama provider.
func NewProvider() pb.ProviderServer { return &provider{} }

// NewMeta advertises command:ollama — the COMPILED-IN registry path resolves it
// (registerCompiledPlugin → resolve(ClassCommand,"ollama") → dispatchInProcCommand →
// Invoke(OpRun)) — plus this plugin's OWN self-contained CUE schema (schema/ollama.cue)
// served over Describe: there is NO schema-less plugin. The command's args are
// pass-through CLI tokens rather than a structured plugin_input, so the schema documents
// the command contract and the endpoint resolution.
func NewMeta() pb.PluginMetaServer {
	return sdk.NewMeta("2026.229.1013",
		[]sdk.ProvidedCapability{{Class: "command", Word: "ollama"}},
		schemaFS)
}

// CliMain is the CLI entrypoint (the out-of-process placement + the shared entry). The
// command is self-contained (plain HTTP to the server), so it runs runOllamaCLI directly —
// no reverse channel, works in either placement.
func CliMain(args []string) int {
	if err := runOllamaCLI(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
