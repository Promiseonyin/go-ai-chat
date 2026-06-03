# Go AI Chat

A local AI chat interface built with Go, Ollama, and HTMX for private, context-aware conversations with your data.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)

---

## The Problem

In healthcare and enterprise environments, patient and customer data exists in silos across multiple systems, formats, and locations. This fragmentation creates critical challenges:

- **Information Loss**: Relevant context is scattered across EMRs, documents, and communication logs, making it difficult for professionals to access complete information quickly.
- **Context Switching**: Users must manually navigate between systems to gather context, leading to inefficiency and increased error risk.
- **Privacy Concerns**: Cloud-based AI services raise compliance questions around HIPAA, GDPR, and data residency requirements.
- **Integration Bottlenecks**: Consolidating fragmented data typically requires complex ETL processes or expensive middleware solutions.

---

## The Solution

**Go AI Chat** provides a **local, private AI context-builder** that:

- **Consolidates Data**: Ingest structured and unstructured data from your systems into a unified interface without sending it to external servers.
- **Builds Context**: Leverage AI to synthesize fragmented information into coherent, actionable summaries and insights.
- **Maintains Privacy**: Run entirely on-premises with Ollama, ensuring sensitive data never leaves your infrastructure.
- **Enables Real-Time Interaction**: Ask questions, explore relationships, and generate reports through a responsive chat interface.
- **Stays Lightweight**: Built with Go for minimal resource footprint, making it deployable on modest hardware.

Perfect for healthcare providers, legal firms, financial institutions, and any organization handling sensitive, distributed data.

---

## The Tech Stack

### Why Go?

- **Performance**: Compiled binary with fast startup times and minimal memory overhead—ideal for containerized deployments.
- **Concurrency**: Goroutines handle multiple chat sessions effortlessly without thread overhead.
- **Simplicity**: Clean syntax reduces maintenance burden; single executable deployment.
- **Standard Library**: Rich built-in packages for HTTP servers, JSON handling, and file I/O eliminate unnecessary dependencies.

### Why Ollama?

- **Privacy First**: Run open-source LLMs locally without API calls or vendor lock-in.
- **Flexibility**: Swap models (Mistral, Llama, Neural Chat, etc.) without code changes.
- **Resource Efficiency**: Optimized for CPU and GPU inference on modest hardware.
- **Ecosystem**: Active community maintaining quantized, production-ready models.

### Supporting Technologies

- **HTMX**: Lightweight interactivity without heavy JavaScript frameworks—fast, responsive UI.
- **Standard Web Stack**: HTML, CSS, and minimal client-side logic for reliability and accessibility.

---

## Getting Started

### Prerequisites

- **Go** 1.21 or later ([Download](https://golang.org/dl))
- **Ollama** ([Download](https://ollama.ai)) or Docker with Ollama image
- **git**

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/Promiseonyin/go-ai-chat.git
   cd go-ai-chat
