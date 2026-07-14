# Application Report — ERPApp-001

**Generated:** 2026-07-14T07:33:55Z  
**Analysis ID:** GD-2026-001  
**App ID:** app001

## Application Overview

ERPApp-001 is a mission-critical finance ERP with a still-supported core stack, but it remains anchored to outdated AIX infrastructure and a high-friction operating model with no CI/CD or containers. Because the application is scheduled for decommission in 2027, modernization should emphasize stabilization, supportability, and selective cost takeout rather than full transformation.

| Field | Value |
| --- | --- |
| App ID | app001 |
| Application Name | ERPApp-001 |
| Business Unit | Finance |
| Criticality | High |
| Deployment | On-Premise |
| Architecture | 1-Tier monolith |
| Users | 350 |
| CI/CD | No |
| Containers | No |
| APIs | 0 |
| External Interfaces | 5 |
| Database | Oracle 19c |
| Database Size | 1000GB |
| License Profile | Licensed |
| Servers | sv01, sv02 |
| Compute | 4 CPU cores / 16GB RAM |
| Logging | Not stated |
| Monitoring | Not stated |
| Planned Decommission | 2027 |

## Technology Assessment

| Component | Version | Status | Assessment | EOL / Support | Confidence |
| --- | --- | --- | --- | --- | --- |
| 🟡 AIX | 7.2 | OUTDATED | Standard support ended; extended support only. | N/A | 7 |
| ✅ COBOL | 2014 | CURRENT | ISO/IEC current standard and still actively supported. | N/A | 8 |
| ✅ Oracle | 19c | CURRENT | Extended support available through April 2027. | 2027-04-30 | 9 |

Technology flags: EOL components = **No**; Outdated components = **Yes**; Missing version data = **No**.

## Complexity Assessment

**Complexity Score:** 9/10 — Very High  
**Multiplier:** 1.8x

| Factor | Why it matters |
| --- | --- |
| High criticality finance workload | Operational disruption directly affects core finance processes and audit commitments. |
| 1-tier monolith | Tight coupling raises migration and regression risk. |
| No CI/CD | Release automation is absent, so every change carries higher delivery friction. |
| No containers | Portability and repeatable environment promotion are limited. |
| Legacy COBOL on proprietary AIX | Specialized skills and platform lock-in increase delivery risk. |
| Five external interfaces | Interface testing must be coordinated across downstream finance integrations. |
| 1000GB Oracle database | Data movement, cutover, and rollback planning are non-trivial. |
| 2027 decommission deadline | Short runway limits the value of deep transformation investments. |

ERPApp-001 is the most difficult application in scope because it combines a monolithic legacy runtime, proprietary infrastructure, a large licensed database, and minimal engineering enablement. The application is still business critical, but the 2027 retirement target means modernization must focus on risk reduction and continuity rather than broad re-architecture.

### Key risks

- Outdated AIX 7.2 increases security, support, and platform continuity risk ahead of the planned retirement window.
- The 1-tier design plus no CI/CD/no containers makes even small changes operationally expensive and slower to validate.
- A 1000GB licensed Oracle footprint and five interfaces raise cutover complexity for any platform or data move.

### Recommended actions

- Patch and harden the AIX estate immediately to reduce near-term operational risk.
- Use the Linux migration as the main platform optimization if the application must remain in service through 2027.
- Avoid large refactoring unless the decommission timeline slips materially or replacement plans are delayed.

## Scenario Applicability

| Scenario | Status | Priority | Effort |
| --- | --- | --- | --- |
| OS Update / Security Patch | APPLICABLE | High | Medium |
| Switch to Standard Linux OS | APPLICABLE | High | High |
| Switch to ARM CPU | BLOCKED | Low | Very High |
| Application Server Replacement | NOT_APPLICABLE | N/A | N/A |
| Cloud Lift & Shift | APPLICABLE | Medium | High |
| Application Containerization | BLOCKED | Low | Very High |
| Application Refactoring / Decoupling | APPLICABLE | High | Very High |
| Upgrade Legacy Databases | FULFILLED | Low | N/A |
| Switch Database to Open Source | APPLICABLE | Medium | High |
| Update Outdated Components | APPLICABLE | High | Medium |

### OS Update / Security Patch (`os_update_security_patch`)

- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** Medium
- **Reasoning:** AIX 7.2 is outdated, so security patching and hardening provide the fastest risk reduction without changing application behavior.
- **Suggestion:** Plan an OS remediation sprint with patch validation, backup verification, and post-change vulnerability review.

### Switch to Standard Linux OS (`switch_to_standard_linux_os`)

- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** High
- **Reasoning:** Moving off proprietary AIX reduces platform dependency and can improve supportability during the remaining service life.
- **Suggestion:** Use a constrained replatform plan that preserves COBOL and Oracle while replacing the host OS and surrounding operational tooling.

### Switch to ARM CPU (`switch_to_arm_cpu`)

- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** Very High
- **Reasoning:** The COBOL/AIX stack is not a natural candidate for ARM adoption without a much larger replatforming effort.
- **Suggestion:** Treat ARM only as a by-product of a future full-stack replacement, not as a standalone modernization workstream.

### Application Server Replacement (`application_server_replacement`)

- **Status:** NOT_APPLICABLE
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** The application is a 1-tier ERP and does not rely on a separate application server layer that can be replaced independently.
- **Suggestion:** Concentrate investment on OS, hosting, database economics, and retirement planning instead.

### Cloud Lift & Shift (`app_deployment_to_cloud`)

- **Status:** APPLICABLE
- **Priority:** Medium
- **Effort:** High
- **Reasoning:** Cloud lift-and-shift is feasible but the monolith, large database, and low automation increase migration effort.
- **Suggestion:** Use cloud migration only if it materially reduces datacenter risk or aligns with the retirement roadmap.

### Application Containerization (`app_containerization`)

- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** Very High
- **Reasoning:** A 1-tier COBOL monolith without CI/CD or service boundaries would need substantial restructuring before containerization becomes practical.
- **Suggestion:** Do not pursue containerization as a near-term goal; the return window is too short.

### Application Refactoring / Decoupling (`app_refactor_decoupling`)

- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** Very High
- **Reasoning:** Refactoring could improve agility, but the cost and delivery risk are exceptionally high for an application due to retire soon.
- **Suggestion:** Use selective decoupling only where it directly supports decommission preparation or a mandatory integration change.

### Upgrade Legacy Databases (`upgrade_legacy_databases`)

- **Status:** FULFILLED
- **Priority:** Low
- **Effort:** N/A
- **Reasoning:** Oracle 19c is already within support, so no immediate database version remediation is required.
- **Suggestion:** Maintain patch cadence and focus on resilience, capacity, and license optimization.

### Switch Database to Open Source (`switch_db_engine_open_source`)

- **Status:** APPLICABLE
- **Priority:** Medium
- **Effort:** High
- **Reasoning:** There is a material licensing cost angle, but data migration and application compatibility work would be significant.
- **Suggestion:** Only justify this move if the application must stay beyond 2027 or if Oracle license costs materially rise.

### Update Outdated Components (`update_outdated_components`)

- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** Medium
- **Reasoning:** AIX is the clear outdated component and should be remediated to avoid support degradation.
- **Suggestion:** Track outdated infrastructure in the same remediation program as OS security patching to avoid duplicated delivery effort.

## Business Case

| Scenario | Adjusted Cost | Yearly Savings | ROI (yrs) |
| --- | --- | --- | --- |
| OS Update | $1,800 | $500 | 3.6 |
| Switch to Linux OS | $540 | $400 | 1.35 |
| Cloud Lift & Shift | $9,000 | $3,000 | 3 |
| App Refactoring | $450,000 | $150,000 | 3 |
| Switch DB to Open-Source | $45,000 | $15,000 | 3 |
| TOTAL | $506,340 | $168,900 | 3 |

Portfolio planning note: ERPApp-001 is best aligned to **Wave 1 for OS hardening and Linux replatform decision; defer deep transformation to replacement planning.**
