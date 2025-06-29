package internal

type Config struct {
	DBPath        string
	EmbedderURL   string
	GeneratorURL  string
	ChunkSize     int
	TopK          int
	Embed_Model   string
	Generator_Model string
}

func DefaultConfig() *Config {
	return &Config{
		DBPath:       "rag.db",
		EmbedderURL:  "http://localhost:11434/api/embeddings",
		GeneratorURL: "http://localhost:11434/api/generate",
		ChunkSize:    200, // words per chunk
		TopK:         3,   // top k chunks for retrieval
		Embed_Model: "nomic-embed-text",
		Generator_Model: "deepseek-r1:1.5b",
	}
}
