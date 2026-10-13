# AI Agent in Golang with OpenRouter

A lightweight implementation demonstrating how to build an autonomous AI agent using **Golang** and the **OpenRouter API**, featuring tool execution loops and structured prompt handling.

---

## Overview

This project provides a foundational architecture for building AI agents in Go. Unlike simple request-response wrappers, this setup implements an agentic execution flow where the language model can invoke local tools or functions to complete complex, multi-step tasks.

### Key Features
- **Golang Backend:** Fast, concurrent, and statically typed implementation suited for production microservices.
- **OpenRouter Integration:** Flexible access to multi-provider LLM endpoints using standard API interfaces.
- **Tool Calling Loop:** Demonstrates how to parse model outputs, intercept tool execution requests, run local code/functions, and feed the results back into the conversation context.

---

## Project Structure

```text
├── main.go          # Entry point and application configuration
├── go.mod           # Go modules and dependencies
└── README.md        # Project documentation