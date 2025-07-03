# 🧠 Ragcli — AI-Powered Document RAG CLI Tool (powered by docai)

Integrate powerful RAG (Retrieval-Augmented Generation) workflows in your Go applications like a pro.
Built using the modular docai library – the LangChain-equivalent for Go.

## ✨ Features

✅ Add .pdf, .txt, .docx documents

🔍 Ask natural language questions from full or specific document

🧩 Modular architecture (built with docai)

🧠 Uses Ollama for local embeddings + LLM responses

📄 Summarize documents in seconds

💾 Uses SQLite for metadata and persistent vector storage

🧱 Chain-based RAG pipeline (inspired by LangChain, built for Go)

🛠️ Built on Cobra CLI — extendable and scriptable

## 📦 Installation

```bash
go install github.com/SrabanMondal/ragcli@latest
```

Or clone:

```bash
git clone https://github.com/SrabanMondal/ragcli
cd ragcli
go run ./cmd
```

## 📚 Commands

```bash
ragcli add --file path/to/file.pdf     # Add a document
ragcli list                            # List added documents
ragcli remove --doc filename           # Remove a document
ragcli ask --query "your question?"    # Ask from all documents
ragcli ask --query "..." --doc file    # Ask from a specific document
ragcli summarize --doc file            # Summarize a document
```

## 🔌 Under the Hood

ragcli uses docai — a powerful framework written in Go:  

📄 reader: Supports PDF, TXT, DOCX readers  

🔪 chunker: Sentence chunking for long docs  

📌 store: Persistent SQLite store for chunks & vectors  

🧠 embedder: Embedding interface with Ollama-backed models  

🤖 generator: Text generation interface (Ollama-supported)  

🔗 chain: LangChain-style modular pipeline in Go  

🧭 retriever: Cosine similarity-based top-K retriever  

📂 File Extensions Supported  
-.pdf
-.txt
-.docx

## 🛠️ Configuration (Set flags using config command)

```bash
--db ./data/docai.db                     # SQLite path
--embedder-url http://localhost:11434   # Ollama embedding endpoint
--generator-url http://localhost:11434  # Ollama generation endpoint
--embed-model all-minilm                # Ollama embedding model
--generator-model llama3                # Ollama chat model
```

## 💡 Example Workflow

```bash
ragcli add --file notes.pdf
ragcli add --file report.txt
ragcli list
ragcli ask --query "What are key insights?"
ragcli ask --query "Timeline of events?" --doc report
ragcli summarize --doc notes
```

## 💻 Sample Output

```bash
$ ragcli ask --query "What is the summary of this contract?"
✅ Embedded query and retrieved top 5 chunks
🤖 AI Answer:
"This contract outlines mutual responsibilities, termination clauses, and a confidentiality agreement between both parties..."
```

## 🧠 Inspired By

LangChain (Python)  
Ollama  
LlamaIndex  

## 👨‍💻 Contributing

PRs, issues, and ideas are welcome. Let’s bring Go into the AI-first era 🧠💪

## ⚡ Future Additions

🔧 Plug-and-play support for OpenAI, Cohere, and HuggingFace APIs

📈 Local analytics on queries and answer quality

🧠 Smart caching layer for faster repeated answers

🌐 Web UI using Fiber or React