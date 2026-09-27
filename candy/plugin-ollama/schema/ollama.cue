// plugin-ollama's OWN self-contained CUE schema — the SINGLE SOURCE for this plugin's
// declaration surface, served over the Describe channel (there is no schema-less
// plugin). SELF-CONTAINED: it references no base def, so it compiles STANDALONE (the
// property the SDK's serve-side compile and `cue exp gengotypes` both need).
//
// `command:ollama`'s authored input is its pass-through CLI grammar (`ollama
// list|ps|pull|…` plus flags), not a structured plugin_input, so this schema DOCUMENTS
// the command contract and the endpoint/config surface the CLI reads.
#OllamaPlugin: {
	// The command word the plugin serves.
	command: "ollama"

	// What the command does, in one line (the public-docs surface).
	contract: string & !=""

	// Endpoint resolution: --server flag > OLLAMA_HOST env > this default.
	default_endpoint: "http://127.0.0.1:11434"
}
