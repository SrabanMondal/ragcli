package internal

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Ashank007/docai/chunker"
	"github.com/Ashank007/docai/embedder"
	"github.com/Ashank007/docai/generator"
	"github.com/Ashank007/docai/reader"
	"github.com/Ashank007/docai/store"
	//"github.com/Ashank007/docai/types"
)

type RAGProcessor struct {
	MetaStore   store.MetadataStore
	VectorStore store.VectorStore
	Embedder    embedder.Embedder
	Generator   generator.Generator
	Chunker     chunker.Chunker
	Reader      *reader.PDFReader
	Config      *Config
}

func NewRAGProcessor(cfg *Config) *RAGProcessor {
	meta := store.NewSQLiteStore()
	if err := meta.Init(cfg.DBPath); err != nil {
		log.Fatalf("❌ Failed to init DB: %v", err)
	}
	vector, err := store.NewSQLiteVectorStore(meta.DB())
	if err != nil {
		log.Fatal("❌ Vector store init failed:", err)
	}
	return &RAGProcessor{
		MetaStore:   meta,
		VectorStore: vector,
		Embedder:    embedder.NewOllama(cfg.Embed_Model, cfg.EmbedderURL),
		Generator:   generator.NewOllama(cfg.Generator_Model, cfg.GeneratorURL),
		Chunker:     chunker.NewSentenceChunker(cfg.ChunkSize),
		Reader:      reader.NewPDFReader(),
		Config:      cfg,
	}
}

func (p *RAGProcessor) AddFile(path string) error {
	text, err := p.Reader.Extract(path)
	if err != nil {
		return fmt.Errorf("❌ Failed to read PDF: %w", err)
	}

	docName := getDocName(path)

	chunks, err := p.Chunker.Chunk(text)
	if err != nil {
		return fmt.Errorf("❌ Failed to chunk: %w", err)
	}

	for i, chunk := range chunks {
		chunk.Page = 0
		chunk.Position = i
		chunk.Source = docName

		id, err := p.MetaStore.SaveChunk(docName, chunk)
		if err != nil {
			log.Printf("⚠️ Could not save chunk: %v", err)
			continue
		}

		vec, err := p.Embedder.Embed(chunk.Text)
		if err != nil {
			log.Printf("⚠️ Failed to embed chunk: %v", err)
			continue
		}
		p.VectorStore.AddVector(id, vec)
	}

	fmt.Printf("✅ %d chunks indexed from: %s\n", len(chunks), docName)
	return nil
}

func (p *RAGProcessor) AskQuestion(query string) (string, error) {
	queryVec, err := p.Embedder.Embed(query)
	if err != nil {
		return "", fmt.Errorf("❌ Failed to embed query: %w", err)
	}

	ids, err := p.VectorStore.SearchSimilar(queryVec, p.Config.TopK)
	if err != nil {
		return "", fmt.Errorf("❌ Vector search failed: %w", err)
	}
	var contexts []string
	for _, id := range ids {
		chunk, err := p.MetaStore.GetChunkByID(id)
		if err == nil {
			contexts = append(contexts, chunk.Text)
		}
	}

	answer, err := p.Generator.Generate(query, contexts)
	if err != nil {
		return "", fmt.Errorf("❌ Generator failed: %w", err)
	}
	return answer, nil
}

func (p *RAGProcessor) ListFiles() error {
	files, err := p.MetaStore.ListFiles()
	if err != nil {
		return err
	}
	if len(files) == 0 {
		fmt.Println("📭 No documents added.")
		return nil
	}
	for _, f := range files {
		fmt.Println("📄", f.Name)
	}
	return nil
}

func (p *RAGProcessor) RemoveFile(name string) error {
	if err := p.MetaStore.DeleteFile(name); err != nil {
		return err
	}
	p.VectorStore.Reset()
	fmt.Println("🗑️  File removed:", name)
	return nil
}

func getDocName(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return base[:len(base)-len(ext)]
}
