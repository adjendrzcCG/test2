# Application Report — AnalyticsApp-003

**Generated:** 2026-07-14T07:33:55Z  
**Analysis ID:** GD-2026-001  
**App ID:** app003

## Application Overview

AnalyticsApp-003 is low criticality and already operates with modern engineering practices, but nearly every runtime layer is out of support. It is the best candidate for a bundled remediation sprint because the technical debt is urgent while delivery complexity remains manageable.

| Field | Value |
| --- | --- |
| App ID | app003 |
| Application Name | AnalyticsApp-003 |
| Business Unit | IT |
| Criticality | Low |
| Deployment | AWS |
| Architecture | 3-Tier |
| Users | 480 |
| CI/CD | Yes |
| Containers | Yes |
| APIs | 8 |
| External Interfaces | 3 |
| Database | PostgreSQL 13 |
| Database Size | 200GB |
| License Profile | No license cost |
| Servers | sv03 |
| Compute | 2 CPU cores / 8GB RAM |
| Logging | ELK |
| Monitoring | Prometheus |
| Planned Decommission | Not stated |

## Technology Assessment

| Component | Version | Status | Assessment | EOL / Support | Confidence |
| --- | --- | --- | --- | --- | --- |
| 🔴 RHEL | 7 | EOL | Operating system end of life reached June 30, 2024. | 2024-06-30 | 10 |
| 🔴 Python | 3.9 | EOL | Community support ended October 5, 2025. | 2025-10-05 | 10 |
| 🔴 Apache Tomcat | 6.1 | EOL | Version line has been unsupported since 2016. | 2016-12-31 | 10 |
| 🔴 PostgreSQL | 13 | EOL | Major version support ended November 13, 2025. | 2025-11-13 | 10 |

Technology flags: EOL components = **Yes**; Outdated components = **No**; Missing version data = **No**.

## Complexity Assessment

**Complexity Score:** 4/10 — Medium  
**Multiplier:** 0.8x

| Factor | Why it matters |
| --- | --- |
| Low criticality | Modernization can be delivered with lower business outage sensitivity. |
| 3-tier architecture | Logical separation makes component-by-component change more manageable. |
| CI/CD already present | Automation reduces regression and release effort. |
| Containers already present | Platform portability is already established. |
| Single server footprint | Infrastructure scope is relatively small. |
| Three external interfaces | Integration coordination is modest compared with the rest of the portfolio. |
| All four major components are EOL | Technical urgency is high even though implementation effort is moderate. |
| Open-source stack | Upgrade paths are usually more direct and licensing lock-in is limited. |

AnalyticsApp-003 has the broadest EOL stack exposure in the portfolio, but it is also the easiest environment in which to execute remediation because it is already cloud-hosted, containerized, and automated. That combination makes it an excellent quick-win candidate: the technical issues are severe, yet the delivery path is comparatively straightforward.

### Key risks

- The application has full-stack EOL exposure across OS, runtime, application server, and database.
- If the low criticality label suppresses urgency, the environment may remain vulnerable even though upgrades are comparatively easy.
- Independent remediation projects could duplicate effort; the stack should be updated as a coordinated package.

### Recommended actions

- Use the application as a quick-win modernization wave to remove the largest EOL concentration in the portfolio.
- Bundle app server, runtime, and database upgrade work to take advantage of existing containers and CI/CD.
- Evaluate ARM only after the core EOL stack has been remediated and performance baselines are stable.

## Scenario Applicability

| Scenario | Status | Priority | Effort |
| --- | --- | --- | --- |
| OS Update / Security Patch | APPLICABLE | High | Low |
| Switch to Standard Linux OS | FULFILLED | N/A | N/A |
| Switch to ARM CPU | APPLICABLE | Medium | Medium |
| Application Server Replacement | APPLICABLE | Critical | Medium |
| Cloud Lift & Shift | FULFILLED | N/A | N/A |
| Application Containerization | FULFILLED | N/A | N/A |
| Application Refactoring / Decoupling | BLOCKED | Low | Very High |
| Upgrade Legacy Databases | APPLICABLE | High | Low |
| Switch Database to Open Source | FULFILLED | N/A | N/A |
| Update Outdated Components | APPLICABLE | Critical | Medium |

### OS Update / Security Patch (`os_update_security_patch`)

- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** Low
- **Reasoning:** RHEL 7 is EOL and can be remediated with a relatively low-effort update path in an already automated environment.
- **Suggestion:** Make OS remediation the opening step in a broader stack-refresh sprint.

### Switch to Standard Linux OS (`switch_to_standard_linux_os`)

- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** The application already runs on a standard Linux platform family.
- **Suggestion:** Upgrade to a supported release rather than changing OS family.

### Switch to ARM CPU (`switch_to_arm_cpu`)

- **Status:** APPLICABLE
- **Priority:** Medium
- **Effort:** Medium
- **Reasoning:** Containers and cloud hosting make ARM testing feasible once the supported stack baseline is restored.
- **Suggestion:** Assess ARM after the core remediation to avoid conflating compatibility and lifecycle risk.

### Application Server Replacement (`application_server_replacement`)

- **Status:** APPLICABLE
- **Priority:** Critical
- **Effort:** Medium
- **Reasoning:** Tomcat 6.1 is long unsupported and should be upgraded as part of the central stack refresh.
- **Suggestion:** Bundle this with runtime testing and container image refresh to reduce repeated validation effort.

### Cloud Lift & Shift (`app_deployment_to_cloud`)

- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** The application is already deployed in AWS.
- **Suggestion:** Leverage cloud elasticity and snapshots for upgrade rehearsals.

### Application Containerization (`app_containerization`)

- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** Containerization is already implemented.
- **Suggestion:** Use the existing container platform to standardize rebuilds and promotion for upgraded dependencies.

### Application Refactoring / Decoupling (`app_refactor_decoupling`)

- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** Very High
- **Reasoning:** Deep refactoring is not justified for a low-criticality analytics platform when lifecycle remediation can address the main risk.
- **Suggestion:** Keep the focus on component upgrades rather than architectural redesign.

### Upgrade Legacy Databases (`upgrade_legacy_databases`)

- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** Low
- **Reasoning:** PostgreSQL 13 is EOL and the open-source upgrade path is usually straightforward.
- **Suggestion:** Coordinate the database major-version upgrade with application regression testing in the same release train.

### Switch Database to Open Source (`switch_db_engine_open_source`)

- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** The workload already uses PostgreSQL, so the open-source database objective is met.
- **Suggestion:** No engine swap is needed; only lifecycle upgrades remain.

### Update Outdated Components (`update_outdated_components`)

- **Status:** APPLICABLE
- **Priority:** Critical
- **Effort:** Medium
- **Reasoning:** Python 3.9 and other stack elements require uplift to regain a supported baseline.
- **Suggestion:** Define a single dependency modernization backlog covering container images, runtime versions, and automated regression checks.

## Business Case

| Scenario | Adjusted Cost | Yearly Savings | ROI (yrs) |
| --- | --- | --- | --- |
| OS Update | $800 | $500 | 1.6 |
| Switch to ARM CPU | $4,000 | $1,000 | 4 |
| App Server Replacement | $8,000 | $12,000 | 0.67 |
| Upgrade Legacy Databases | $8,000 | $10,000 | 0.8 |
| TOTAL | $20,800 | $23,500 | 0.89 |

Portfolio planning note: AnalyticsApp-003 is best aligned to **Wave 1 / early Wave 2 quick-win modernization bundle covering OS, Tomcat, Python, and PostgreSQL.**
