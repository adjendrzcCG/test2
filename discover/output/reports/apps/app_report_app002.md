# Application Report — CRMApp-002

**Generated:** 2026-07-14T07:33:55Z  
**Analysis ID:** GD-2026-001  
**App ID:** app002

## Application Overview

CRMApp-002 already has cloud hosting and delivery automation, but it remains exposed through EOL RHEL 7 and WebSphere 7.0 plus an aging Java 11 runtime. The best value path is to remediate the middleware layer first, because that removes the highest support risk and unlocks follow-on modernization.

| Field | Value |
| --- | --- |
| App ID | app002 |
| Application Name | CRMApp-002 |
| Business Unit | Marketing |
| Criticality | Medium |
| Deployment | AWS |
| Architecture | Unknown |
| Users | 1200 |
| CI/CD | Yes |
| Containers | No |
| APIs | 15 |
| External Interfaces | 8 |
| Database | Amazon RDS MySQL |
| Database Size | Managed service |
| License Profile | No license cost |
| Servers | sv05, sv07 |
| Compute | 2 CPU cores / 8GB RAM |
| Logging | Cloud-native / monitoring tools |
| Monitoring | CloudWatch, Datadog |
| Planned Decommission | Not stated |

## Technology Assessment

| Component | Version | Status | Assessment | EOL / Support | Confidence |
| --- | --- | --- | --- | --- | --- |
| 🔴 RHEL | 7 | EOL | End of life reached June 30, 2024. | 2024-06-30 | 10 |
| 🟡 Java | 11 | OUTDATED | Extended support only; free Oracle updates ended in 2023. | 2026-09-30 | 8 |
| 🔴 WebSphere | 7.0 | EOL | Vendor support ended in 2015. | 2015-09-30 | 10 |
| ✅ Amazon RDS MySQL | Managed | CURRENT | AWS-managed service keeps database patching current. | Managed by AWS | 9 |

Technology flags: EOL components = **Yes**; Outdated components = **Yes**; Missing version data = **No**.

## Complexity Assessment

**Complexity Score:** 7/10 — High  
**Multiplier:** 1.4x

| Factor | Why it matters |
| --- | --- |
| Third-party product constraints | Vendor packaging limits how aggressively components can be changed. |
| Unknown architecture | Discovery work is needed before estimating migration sequencing confidently. |
| RHEL 7 and WebSphere 7.0 are EOL | Unsupported infrastructure and middleware create immediate support and security pressure. |
| Java 11 is aging | Runtime uplift will likely be needed during or after app server migration. |
| 15 APIs and 8 external interfaces | Regression testing and contract validation are significant. |
| 1200 users | Business-facing downtime and user acceptance windows must be tightly managed. |
| No containers | Modern hosting flexibility is limited until the middleware stack is refreshed. |

CRMApp-002 is easier to modernize than the ERP because it already runs in AWS and has CI/CD, but it is still constrained by severe middleware obsolescence and third-party packaging. The priority is to replace WebSphere and remediate RHEL before broader optimization, because those moves unlock most other scenario paths.

### Key risks

- WebSphere 7.0 and RHEL 7 are both unsupported, creating urgent supportability and security exposure.
- The application has a wide integration footprint, so middleware failure or rushed migration would have broad downstream impact.
- Third-party ownership constrains refactoring and can make component upgrades dependent on vendor-supported paths.

### Recommended actions

- Prioritize application server replacement as the main structural modernization move.
- Patch and remediate the OS in parallel to reduce EOL exposure while the application server program is mobilized.
- Delay containerization and broader component refresh until a supported application server baseline is in place.

## Scenario Applicability

| Scenario | Status | Priority | Effort |
| --- | --- | --- | --- |
| OS Update / Security Patch | APPLICABLE | Critical | Medium |
| Switch to Standard Linux OS | FULFILLED | N/A | N/A |
| Switch to ARM CPU | BLOCKED | Low | Very High |
| Application Server Replacement | APPLICABLE | Critical | High |
| Cloud Lift & Shift | FULFILLED | N/A | N/A |
| Application Containerization | BLOCKED | Low | Very High |
| Application Refactoring / Decoupling | BLOCKED | Low | N/A |
| Upgrade Legacy Databases | FULFILLED | N/A | N/A |
| Switch Database to Open Source | FULFILLED | N/A | N/A |
| Update Outdated Components | BLOCKED | High | High |

### OS Update / Security Patch (`os_update_security_patch`)

- **Status:** APPLICABLE
- **Priority:** Critical
- **Effort:** Medium
- **Reasoning:** RHEL 7 is already EOL, so patch governance and near-term remediation are urgent for continued safe operation.
- **Suggestion:** Establish a short remediation window to reduce exposure while the WebSphere replacement plan is finalized.

### Switch to Standard Linux OS (`switch_to_standard_linux_os`)

- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** The application already runs on a standard Linux distribution, so no platform family change is required.
- **Suggestion:** Focus effort on moving from an EOL release to a supported target during the middleware transition.

### Switch to ARM CPU (`switch_to_arm_cpu`)

- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** Very High
- **Reasoning:** The current vendor stack and WebSphere dependency make ARM migration impractical at this stage.
- **Suggestion:** Revisit only after the product is on a supported middleware and hosting baseline.

### Application Server Replacement (`application_server_replacement`)

- **Status:** APPLICABLE
- **Priority:** Critical
- **Effort:** High
- **Reasoning:** WebSphere 7.0 is long out of support and is the largest structural source of technical debt in the application.
- **Suggestion:** Treat application server replacement as the anchor modernization initiative and align vendor supportability checks early.

### Cloud Lift & Shift (`app_deployment_to_cloud`)

- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** The workload is already hosted in AWS.
- **Suggestion:** Use the existing cloud footprint to simplify migration environments, testing, and rollback options.

### Application Containerization (`app_containerization`)

- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** Very High
- **Reasoning:** Containerization is not attractive while the application still depends on EOL middleware and vendor-controlled packaging.
- **Suggestion:** Assess containers only after the application server is modernized and vendor support boundaries are clear.

### Application Refactoring / Decoupling (`app_refactor_decoupling`)

- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** N/A
- **Reasoning:** As a third-party CRM, the application is not a good candidate for major internal refactoring by the enterprise owner.
- **Suggestion:** Constrain modernization to vendor-supported platform, middleware, and integration patterns.

### Upgrade Legacy Databases (`upgrade_legacy_databases`)

- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** The database runs as a managed RDS MySQL service and does not represent an immediate legacy database issue.
- **Suggestion:** Continue to rely on managed-service lifecycle updates and focus application work higher in the stack.

### Switch Database to Open Source (`switch_db_engine_open_source`)

- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** MySQL is already an open-source engine delivered via managed AWS hosting.
- **Suggestion:** No database engine replacement is required.

### Update Outdated Components (`update_outdated_components`)

- **Status:** BLOCKED
- **Priority:** High
- **Effort:** High
- **Reasoning:** Java uplift and wider component refresh depend on first replacing the unsupported WebSphere foundation.
- **Suggestion:** Sequence runtime and dependency modernization immediately after the application server change.

## Business Case

| Scenario | Adjusted Cost | Yearly Savings | ROI (yrs) |
| --- | --- | --- | --- |
| OS Update | $1,400 | $500 | 2.8 |
| App Server Replacement | $14,000 | $12,000 | 1.17 |
| TOTAL | $15,400 | $12,500 | 1.23 |

Portfolio planning note: CRMApp-002 is best aligned to **Wave 1 for OS remediation, then Wave 2 for the WebSphere replacement and follow-on runtime updates.**
