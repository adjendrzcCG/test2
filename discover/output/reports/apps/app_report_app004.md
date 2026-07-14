# Application Report — HRApp-004

**Generated:** 2026-07-14T07:33:55Z  
**Analysis ID:** GD-2026-001  
**App ID:** app004

## Application Overview

HRApp-004 has a usable modernization foundation, but its hosting stack is already out of support and key lifecycle data for .NET is missing. The right approach is to stabilize the platform first, close the information gap, and then use that foundation to finish the cloud journey and reduce licensing overhead.

| Field | Value |
| --- | --- |
| App ID | app004 |
| Application Name | HRApp-004 |
| Business Unit | HR |
| Criticality | High |
| Deployment | AWS + On-Premise hybrid |
| Architecture | 2-Tier |
| Users | 670 |
| CI/CD | Yes |
| Containers | Yes |
| APIs | 12 |
| External Interfaces | 6 |
| Database | SQL Server 2019 |
| Database Size | 750GB |
| License Profile | Licensed |
| Servers | sv06, sv02 |
| Compute | 4 CPU cores / 16GB RAM |
| Logging | Splunk |
| Monitoring | CloudWatch |
| Planned Decommission | Not stated |

## Technology Assessment

| Component | Version | Status | Assessment | EOL / Support | Confidence |
| --- | --- | --- | --- | --- | --- |
| 🔴 Windows Server | 2012 | EOL | End of life reached October 10, 2023. | 2023-10-10 | 10 |
| ❓ .NET Core | Unknown | NO_KNOWLEDGE | Version not supplied, so lifecycle status cannot be confirmed. | Unknown | 3 |
| 🔴 Microsoft IIS | 8.0 | EOL | Lifecycle tied to Windows Server 2012 and therefore unsupported. | 2023-10-10 | 10 |
| 🟡 SQL Server | 2019 | OUTDATED | Mainstream support ended in 2025; extended support continues. | 2030-01-07 (extended) | 9 |

Technology flags: EOL components = **Yes**; Outdated components = **Yes**; Missing version data = **Yes**.

## Complexity Assessment

**Complexity Score:** 7/10 — High  
**Multiplier:** 1.4x

| Factor | Why it matters |
| --- | --- |
| High criticality HR and payroll-sensitive data | Security and downtime concerns are elevated because the system processes sensitive employee information. |
| Hybrid AWS + on-prem deployment | Operational ownership is split across environments and complicates cutover sequencing. |
| Windows Server 2012 and IIS 8.0 are EOL | Unsupported platform layers create urgent remediation pressure. |
| Unknown .NET Core version | Missing lifecycle data raises planning risk and can hide larger upgrade effort. |
| Containers and CI/CD already present | Delivery mechanics are stronger than the underlying platform health. |
| Six external interfaces and 12 APIs | Interface validation remains significant for every platform change. |
| Licensed SQL Server footprint | There is a medium-term cost optimization angle beyond support remediation. |

HRApp-004 mixes urgent infrastructure obsolescence with a partially modern engineering posture. It is easier to modernize than the ERP because CI/CD and containers exist, but the hybrid estate, missing .NET version data, and sensitivity of HR data make planning discipline essential before any major refactor or database change.

### Key risks

- Windows Server 2012 and IIS 8.0 are unsupported in a high-sensitivity HR environment.
- The unknown .NET Core version introduces planning uncertainty and could hide additional runtime remediation effort.
- Hybrid deployment increases operational complexity and slows standardization of controls, observability, and release procedures.

### Recommended actions

- Close the .NET version data gap immediately so the rest of the remediation plan can be validated.
- Replace or uplift the unsupported Windows/IIS hosting baseline as the first major engineering action.
- Use the existing CI/CD and container base to complete cloud migration and stage later database/licensing improvements.

## Scenario Applicability

| Scenario | Status | Priority | Effort |
| --- | --- | --- | --- |
| OS Update / Security Patch | APPLICABLE | Critical | Medium |
| Switch to Standard Linux OS | NOT_APPLICABLE | N/A | N/A |
| Switch to ARM CPU | BLOCKED | Low | Very High |
| Application Server Replacement | APPLICABLE | Critical | Medium |
| Cloud Lift & Shift | PARTIALLY_FULFILLED | High | High |
| Application Containerization | FULFILLED | N/A | N/A |
| Application Refactoring / Decoupling | APPLICABLE | Medium | Very High |
| Upgrade Legacy Databases | APPLICABLE | Medium | Medium |
| Switch Database to Open Source | APPLICABLE | Medium | High |
| Update Outdated Components | APPLICABLE | High | High |

### OS Update / Security Patch (`os_update_security_patch`)

- **Status:** APPLICABLE
- **Priority:** Critical
- **Effort:** Medium
- **Reasoning:** The current Windows Server baseline is already EOL, so immediate remediation is required to reduce infrastructure risk.
- **Suggestion:** Run a time-boxed OS remediation program with validation of HR data handling, access controls, and recovery procedures.

### Switch to Standard Linux OS (`switch_to_standard_linux_os`)

- **Status:** NOT_APPLICABLE
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** The application is aligned to a Windows/.NET/IIS hosting model, so a Linux switch is not the direct modernization path.
- **Suggestion:** Focus on supported Windows or platform-service targets instead of forcing an OS family rewrite.

### Switch to ARM CPU (`switch_to_arm_cpu`)

- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** Very High
- **Reasoning:** ARM does not align well with the current Windows/IIS estate and would introduce unnecessary complexity right now.
- **Suggestion:** Reassess only if the application later moves to a different supported hosting model.

### Application Server Replacement (`application_server_replacement`)

- **Status:** APPLICABLE
- **Priority:** Critical
- **Effort:** Medium
- **Reasoning:** IIS 8.0 is unsupported, making application server uplift a priority alongside the OS remediation.
- **Suggestion:** Treat web/application hosting modernization as a paired workstream with the operating system update.

### Cloud Lift & Shift (`app_deployment_to_cloud`)

- **Status:** PARTIALLY_FULFILLED
- **Priority:** High
- **Effort:** High
- **Reasoning:** The application already has a cloud footprint, but the remaining hybrid topology keeps operational complexity high.
- **Suggestion:** Use platform remediation to standardize deployment and retire residual on-prem dependencies where feasible.

### Application Containerization (`app_containerization`)

- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** Containers are already in place, which reduces delivery friction for later modernization phases.
- **Suggestion:** Exploit the container baseline to rehearse runtime and hosting upgrades in lower-risk environments.

### Application Refactoring / Decoupling (`app_refactor_decoupling`)

- **Status:** APPLICABLE
- **Priority:** Medium
- **Effort:** Very High
- **Reasoning:** Deeper refactoring could simplify the hybrid architecture, but it should follow platform stabilization and .NET discovery.
- **Suggestion:** Limit initial refactoring to seams that support cloud completion, supportability, or security segregation.

### Upgrade Legacy Databases (`upgrade_legacy_databases`)

- **Status:** APPLICABLE
- **Priority:** Medium
- **Effort:** Medium
- **Reasoning:** SQL Server 2019 is not EOL, but it is past mainstream support and should be aligned with broader platform modernization.
- **Suggestion:** Sequence the database upgrade after the hosting tier is stabilized to avoid stacking major changes in one release.

### Switch Database to Open Source (`switch_db_engine_open_source`)

- **Status:** APPLICABLE
- **Priority:** Medium
- **Effort:** High
- **Reasoning:** There is a credible license-cost reduction path, but application compatibility and data migration need careful evaluation.
- **Suggestion:** Run a fit-gap assessment before committing to an engine switch; do not combine it with first-wave infrastructure remediation.

### Update Outdated Components (`update_outdated_components`)

- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** High
- **Reasoning:** SQL Server and possibly the .NET runtime need attention, but the exact scope depends on confirming the current .NET version.
- **Suggestion:** Treat component modernization as a second-stage backlog gated by version discovery and hosting remediation.

## Business Case

| Scenario | Adjusted Cost | Yearly Savings | ROI (yrs) |
| --- | --- | --- | --- |
| OS Update | $1,400 | $500 | 2.8 |
| App Server Replacement | $14,000 | $12,000 | 1.17 |
| Cloud Lift & Shift (partial) | $7,000 | $3,000 | 2.33 |
| App Refactoring | $350,000 | $150,000 | 2.33 |
| Upgrade Legacy Databases | $14,000 | $10,000 | 1.4 |
| Switch DB to Open-Source | $35,000 | $15,000 | 2.33 |
| TOTAL | $421,400 | $190,500 | 2.21 |

Portfolio planning note: HRApp-004 is best aligned to **Wave 1 for discovery and platform hardening, Wave 2 for hosting/cloud completion, Wave 3 for refactoring and database economics.**
