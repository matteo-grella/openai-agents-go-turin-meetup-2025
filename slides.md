---
marp: true
theme: default
paginate: true
class: lead
title: "Building Multi-Agent Workflows in Go"
description: "Turin Go Meetup 2025 — Matteo Grella"
---

# 🧠 Building Multi-Agent Workflows in Go  
### Turin Go Meetup 2025  
**Matteo Grella — @matteo-grella**  
[github.com/nlpodyssey/openai-agents-go](https://github.com/nlpodyssey/openai-agents-go)
[github.com/matteo-grella/openai-agents-go-turin-meetup-2025](https://github.com/matteo-grella/openai-agents-go-turin-meetup-2025)

---

## 🎯 Goals of this session

- Understand **LLMs → Agents → Multi-Agent Workflows**
- Explore [`openai-agents-go`](https://github.com/nlpodyssey/openai-agents-go)
- Learn **Agent patterns and orchestration**
- See **Go code examples** and real use cases

---

## 🧩 From LLMs to Agents

**LLM:** a large probabilistic text generator.  
**Agent:** an LLM + memory + tools + goals.

| Concept | Description |
|----------|--------------|
| LLM | Predicts next token from context |
| Agent | Adds purpose, state, and actions |
| Multi-Agent System | Agents collaborating and handing off tasks |

---

## ⚙️ LLM Basics

- GPT-*, Claude, Gemini, etc.  
- APIs → *Responses* / *Chat Completions*  
- Prompt → Context → Structured Output  
- Stateless by default  
- But: **no persistence**, **no tools**, **no reasoning loop**

---

## 🤖 What’s an Agent?

> “An LLM configured with instructions, tools, memory, and goals.”

- **Instructions** → role and behavior  
- **Tools** → call external functions  
- **Memory** → persistent context  
- **Guardrails** → input/output validation  
- **Handoffs** → transfer control between agents

---

## 🧠 Agents in Go (SDK Overview)

**`openai-agents-go`**  
A lightweight, provider-agnostic SDK for multi-agent workflows.  
Go port of [OpenAI Agents Python SDK](https://openai.github.io/openai-agents-python/).

```bash
go get github.com/nlpodyssey/openai-agents-go
````

---

## 🪶 Hello World Example

```go
agent := agents.New("Assistant").
  WithInstructions("You are a helpful assistant").
  WithModel("gpt-4o")

result, _ := agents.Run(context.Background(), agent,
  "If you could join us at this meetup, what snack would you bring?")

fmt.Println(result.FinalOutput)
```

---

## 🔄 The Agent Loop

1. LLM call with message history
2. Response may include:

   * Tool calls
   * Handoffs
   * Final output
3. SDK executes tools or new agent
4. Loop continues until final output

`MaxTurns` prevents infinite loops.

---

## 🧰 Tools Example

```go
type GetWeatherParams struct { City string `json:"city"` }

func getWeather(_ context.Context, p GetWeatherParams) (string, error) {
  return fmt.Sprintf("The weather in %s is sunny.", p.City), nil
}

var tool = agents.NewFunctionTool("GetWeather", "", getWeather)

agent := agents.New("WeatherBot").
  WithInstructions("You are a helpful agent").
  WithModel("gpt-4o").
  WithTools(tool)
```

---

## 🌐 Handoffs Example

```go
math := agents.New("Math").
  WithInstructions("Handle all math and logic problems precisely.").
  WithModel("gpt-5")

story := agents.New("Storyteller").
  WithInstructions("Handle narrative or creative writing tasks.").
  WithModel("gpt-4o")

triage := agents.New("Triage").
  WithInstructions("Route math problems to Math, and stories to Storyteller.").
  WithAgentHandoffs(math, story).
  WithModel("gpt-4.1-mini")
```

---

## 🤝 Multi-Agent Workflows

Why multiple agents?

* **Decomposition:** smaller, specialized tasks
* **Parallelization:** speed up reasoning
* **Transparency:** logs per agent
* **Scalability:** plug in new agents easily

---

## 🧭 Workflow Examples

* ✈️ **Customer service:** triage → language → solution
* 💼 **Financial research:** data → analysis → report
* 🔍 **Research bot:** planner → searcher → writer
* 🎧 **Voice interaction:** text → audio responses

👉 See examples in `/examples`:

* `basic/`, `handoffs/`, `research_bot/`, `workflowrunner/`

---

## 🧩 The `workflowrunner` Package

> Experimental orchestration layer

* Declarative agent graph definition
* **prompt** > **multi-agent workflow**

---

## 🧠 Why Multi-Agent Architectures Matter

* **Modularity:** each agent focuses on one skill
* **Transparency:** observable reasoning paths
* **Hybrid logic:** symbolic + neural
* **Reuse:** same agents across workflows
* **Resilience:** failure isolation and retries

---

## 🛠️ Patterns & Best Practices

* Keep instructions concise and role-specific
* Use handoffs for routing and delegation
* Add structured `OutputType` for reliability
* Monitor tokens & cost (`usage` package)
* Trace interactions (`tracing` package)

---

## 💡 Demo Ideas

* `examples/basic` → Hello world agent
* `examples/handoffs` → Triage with handoffs
* `examples/research_bot` → Planner + Searcher + Writer
* `examples/workflowrunner` → Graph orchestration

Each runs with:

```bash
go run examples/<name>/main.go
```

---

## 🚀 When to Use Agents

✅ LLM needs reasoning loops
✅ Coordination between roles
✅ Structured outputs or tool use
✅ Persistent context

❌ Simple one-shot completions
❌ Static prompt engineering

---

## 🧰 Integrations and Ecosystem

* [Model Context Protocol](https://github.com/modelcontextprotocol/go-sdk) (MCP)
* Custom model providers via LiteLLM
* Compatible with OpenAI Responses & Chat APIs
* Cross-language concept parity with Python SDK

---

## 👏 Credits

**Project Authors:**
[Matteo Grella](https://github.com/matteo-grella)
[Marco Nicola](https://github.com/marco-nicola)
[Avi Tal](https://github.com/avi3tal)

Thanks to:

* [OpenAI](https://openai.com) for the original SDK
* [Anthropic](https://anthropic.com) for MCP
* The [nlpodyssey](https://github.com/nlpodyssey) community!

---

## 🧩 Repo Resources

* 🧠 [Main SDK Repo](https://github.com/nlpodyssey/openai-agents-go)
* 💬 [Meetup Materials](https://github.com/matteo-grella/openai-agents-go-turin-meetup-2025)
* 💻 Examples in `/examples/`
* 📜 OpenAI Agents Python: [openai.github.io/openai-agents-python](https://openai.github.io/openai-agents-python)

---

# 🙌 Thank you!

**Q&A / Discussion**

Let’s build something awesome in Go 💙
[@matteo-grella](https://github.com/matteo-grella)

---
