# AGENTS.md | ThitNueaHub Master AI Lab (V83)

## 1. System Identity & Mission
* **Scope:** ThitNueaHub Master AI Lab
* **Version:** V83 Trinity Empire
* **Goal:** High-performance, Zero-Garbage AI Orchestration at Edge (Cloudflare).
* **Core Principle:** "Truth is Speed. Less Garbage, More Logic."

## 2. Infrastructure & Deployment Protocol
* **Deployment Policy:** * STRICTLY PROHIBIT direct `wrangler deploy` in local environment.
    * MANDATORY: Utilize `./scripts/deploy_tnh.sh` or trigger via CI/CD (GitHub Actions).
* **Identity Management:**
    * Refer to `./internal/guardians/` for role-based execution (e.g., `l8_guardian`, `l2_scribe`).
    * All operations must be signed by HMAC as per `Logic Helmet` protocol.

## 3. Review Protocol (The Zero-Garbage Standard)
* **Technical Precision:** Omit conversational fillers. Focus on output quality and latency.
* **Code Hygiene:**
    * **Go Structures:** Validate `./cmd/` against Cloudflare D1/KV/AI bindings.
    * **Cleanup:** Zero tolerance for unused variables, unused imports, or zombie Goroutines.
    * **Configuration:** No hardcoded credentials. All environment-specific values must reside in `./.env` (Root dir).
* **Performance:** Target < 1ms latency for internal dispatching.

## 4. Execution Rules
1. **Task Intake First:** Every instruction must undergo Intake/Queue/Priority validation.
2. **Context Integrity:** Use `wrangler.toml` for resource bindings; never bypass bindings to touch DB directly.
3. **Fail-Safe:** If an operation fails, revert to the last known stable state in `TNH-AI-V83-TRINITY-EMPIRE` branch.

