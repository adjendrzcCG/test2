# Portfolio Modernization Assessment
Generated: 2026-07-14 | Analysis ID: GD-2026-07-14-PORTFOLIO

## Executive Summary

- The portfolio contains **5** applications; **4** are in scope and **1** is out of scope.
- **app005 / EComApp-005** is excluded because **STATUS=RETIRED**.
- Across the in-scope estate, **8 technology components are EOL**, **6 are OUTDATED**, and only **1** is assessed as CURRENT_VERSION.
- The average modernization complexity is **5.75/10**. The highest-complexity application is **app001 / ERPApp-001 (8/10)**.
- There are **21 APPLICABLE or PARTIALLY_FULFILLED modernization scenarios** across the 4 in-scope applications.
- Portfolio investment across those scenarios is **€1,055,540**, with **€418,400 annual savings**, **€199,660 3-year net**, and **€1,036,460 5-year net**.
- The aggregate payback period is **2.52 years** (~30.2 months).

Key conclusions:

1. **ERPApp-001** is the largest strategic risk because it combines high criticality, legacy COBOL/AIX/Oracle technology, and no CI/CD or containerization. However, its stated 2027 decommission date means major investment should be gated by a confirmed retirement decision.
2. **CRMApp-002** is primarily a vendor-constrained remediation case. The largest exposure is EOL WebSphere 7.0, but the scenario for application server replacement is explicitly BLOCKED.
3. **AnalyticsApp-003** is the best quick-win candidate. Even though all core assessed components are EOL, the app is already containerized, cloud-hosted, and supported by CI/CD.
4. **HRApp-004** is the strongest strategic modernization target because it combines high criticality with strong modernization readiness (containers + CI/CD) and the best long-term economic profile.

## Portfolio Overview

| App ID | Application | Business Unit | Criticality | Deployment | Users | Complexity | In Scope |
|---|---|---|---|---|---:|---:|---|
| app001 | ERPApp-001 | Finance | High | On-Premise | 350 | 8 | Yes |
| app002 | CRMApp-002 | Marketing | Medium | AWS | 1200 | 6 | Yes |
| app003 | AnalyticsApp-003 | IT | Low | AWS | 480 | 4 | Yes |
| app004 | HRApp-004 | HR | High | AWS + On-premise (hybrid) | 670 | 5 | Yes |
| app005 | EComApp-005 | N/A | N/A | N/A | N/A | N/A | No |

| Portfolio Metric | Value |
|---|---:|
| Total applications provided | 5 |
| In-scope applications | 4 |
| Out-of-scope applications | 1 |
| Average complexity score | 5.75 |
| Applicable / partially fulfilled scenarios | 21 |
| Portfolio total cost with multipliers | €1,055,540 |
| Portfolio annual savings | €418,400 |
| Portfolio 3-year net | €199,660 |
| Portfolio 5-year net | €1,036,460 |
| Portfolio 3-year ROI | 18.92% |
| Portfolio 5-year ROI | 98.19% |

## Out-of-Scope Applications

| App ID | Application | Reason |
|---|---|---|
| app005 | EComApp-005 | STATUS=RETIRED |

## Technology Assessment Summary

| Lifecycle Status | Count |
|---|---:|
| CURRENT_VERSION | 1 |
| OUTDATED | 6 |
| EOL | 8 |
| NO_KNOWLEDGE | 0 |

| App ID | Application | OS | Language | App Server | Database | Notes |
|---|---|---|---|---|---|---|
| app001 | ERPApp-001 | OUTDATED (AIX 7.2) | OUTDATED (COBOL-2014) | N/A (None) | OUTDATED (Oracle 19c) | No app server present. |
| app002 | CRMApp-002 | EOL (RHEL 7) | OUTDATED (Java 11) | EOL (WebSphere 7.0) | CURRENT_VERSION (Amazon RDS MySQL) | - |
| app003 | AnalyticsApp-003 | EOL (RHEL 7) | EOL (Python 3.9) | EOL (Apache Tomcat 6.1) | EOL (PostgreSQL 13) | - |
| app004 | HRApp-004 | EOL (Windows Server 2012) | OUTDATED (.NET Core (version unspecified)) | EOL (Microsoft IIS 8.0) | OUTDATED (SQL Server 2019) | Low confidence on .NET Core version assessment. |

## Complexity Assessment

| App ID | Application | Complexity Score | Multiplier | Rationale |
|---|---|---:|---:|---|
| app001 | ERPApp-001 | 8 | 1.8x | High criticality + COBOL + AIX + Oracle + no CI/CD + no containers + legacy architecture |
| app002 | CRMApp-002 | 6 | 1.5x | 3rd party constraints + WebSphere EOL + RHEL 7 EOL + no containers |
| app003 | AnalyticsApp-003 | 4 | 1.3x | Already containerized + CI/CD + low criticality + simpler stack |
| app004 | HRApp-004 | 5 | 1.4x | Already containerized + CI/CD + .NET Core + but Windows EOL + hybrid deployment + High criticality |

## Scenario Applicability Matrix

| Scenario | app001 | app002 | app003 | app004 |
|---|---|---|---|---|
| OS Update & Security Patch | APPLICABLE | APPLICABLE | APPLICABLE | APPLICABLE |
| Switch to Standard Linux OS | APPLICABLE | FULFILLED | FULFILLED | NOT_APPLICABLE |
| Switch to ARM-based CPU | NOT_APPLICABLE | NOT_APPLICABLE | APPLICABLE | APPLICABLE |
| Application Server Replacement | NOT_APPLICABLE | BLOCKED | APPLICABLE | APPLICABLE |
| Lift & Shift to Cloud | APPLICABLE | FULFILLED | FULFILLED | PARTIALLY_FULFILLED |
| Application Containerization | BLOCKED | NOT_APPLICABLE | FULFILLED | FULFILLED |
| Application Refactoring & Decoupling | APPLICABLE | NOT_APPLICABLE | NOT_APPLICABLE | APPLICABLE |
| Upgrade Legacy Databases | APPLICABLE | FULFILLED | APPLICABLE | APPLICABLE |
| Switch DB Engine to Open Source | APPLICABLE | FULFILLED | FULFILLED | APPLICABLE |
| Update Outdated Components | APPLICABLE | NOT_APPLICABLE | APPLICABLE | APPLICABLE |

## Business Case Analysis

Formula used for every APPLICABLE or PARTIALLY_FULFILLED scenario:

- `cost_with_multiplier = base_cost × complexity_multiplier`
- `annual_savings = base_savings`
- `breakeven_years = cost_with_multiplier / annual_savings`
- `3yr_net = (annual_savings × 3) - cost_with_multiplier`
- `5yr_net = (annual_savings × 5) - cost_with_multiplier`

| App ID | Applicable Scenarios | Cost with Multiplier | Annual Savings | Aggregate Breakeven (yrs) | 3-Year Net | 5-Year Net |
|---|---:|---:|---:|---:|---:|---:|
| app001 | 7 | €551,340 | €186,900 | 2.95 | €9,360 | €383,160 |
| app002 | 1 | €1,500 | €500 | 3.0 | €0 | €1,000 |
| app003 | 5 | €53,300 | €31,500 | 1.69 | €41,200 | €104,200 |
| app004 | 8 | €449,400 | €199,500 | 2.25 | €149,100 | €548,100 |

| Portfolio Total Cost | Portfolio Annual Savings | Portfolio Breakeven (yrs) | Portfolio 3-Year Net | Portfolio 5-Year Net | Portfolio 3-Year ROI | Portfolio 5-Year ROI |
|---:|---:|---:|---:|---:|---:|---:|
| €1,055,540 | €418,400 | 2.52 | €199,660 | €1,036,460 | 18.92% | 98.19% |

Top 10 opportunities by 5-year net value:

| Rank | Application | Scenario | Status | Cost | Annual Savings | 5-Year Net | Breakeven (yrs) |
|---:|---|---|---|---:|---:|---:|---:|
| 1 | HRApp-004 | Application Refactoring & Decoupling | APPLICABLE | €350,000 | €150,000 | €400,000 | 2.33 |
| 2 | ERPApp-001 | Application Refactoring & Decoupling | APPLICABLE | €450,000 | €150,000 | €300,000 | 3.0 |
| 3 | AnalyticsApp-003 | Application Server Replacement | APPLICABLE | €13,000 | €12,000 | €47,000 | 1.08 |
| 4 | HRApp-004 | Application Server Replacement | APPLICABLE | €14,000 | €12,000 | €46,000 | 1.17 |
| 5 | HRApp-004 | Switch DB Engine to Open Source | APPLICABLE | €35,000 | €15,000 | €40,000 | 2.33 |
| 6 | AnalyticsApp-003 | Upgrade Legacy Databases | APPLICABLE | €13,000 | €10,000 | €37,000 | 1.3 |
| 7 | HRApp-004 | Upgrade Legacy Databases | APPLICABLE | €14,000 | €10,000 | €36,000 | 1.4 |
| 8 | ERPApp-001 | Upgrade Legacy Databases | APPLICABLE | €18,000 | €10,000 | €32,000 | 1.8 |
| 9 | ERPApp-001 | Switch DB Engine to Open Source | APPLICABLE | €45,000 | €15,000 | €30,000 | 3.0 |
| 10 | AnalyticsApp-003 | Update Outdated Components | APPLICABLE | €19,500 | €8,000 | €20,500 | 2.44 |

## Modernization Roadmap

### Phase 1: Immediate risk reduction (0-3 months)

Patch exposure, confirm retirement boundaries, and arrest unsupported platform risk.

- Execute OS Update & Security Patch for app001, app002, app003, and app004.
- Confirm whether ERPApp-001 will actually decommission in 2027; if yes, keep spending tactical.
- Launch upgrade planning for EOL stacks in app003 and app004 because all core runtime layers are unsupported or outdated.
- Escalate vendor dependency for CRMApp-002 WebSphere 7.0 remediation because application_server_replacement is BLOCKED.

### Phase 2: Platform remediation (3-9 months)

Remove the most exposed EOL middleware and database risks.

- For app003, replace Apache Tomcat 6.1, upgrade PostgreSQL 13, and update Python 3.9 runtime components.
- For app004, remediate Windows Server 2012 / IIS 8.0 exposure and execute SQL Server upgrade planning.
- For app001, upgrade Oracle-related risk controls and complete target-state decisioning for Linux/cloud migration.
- For app002, complete the OS patch path while negotiating vendor-supported middleware replacement options.

### Phase 3: Structural modernization (9-18 months)

Invest where the portfolio gets durable architectural value.

- Advance HRApp-004 refactoring and decoupling because it already has containers and CI/CD and yields the strongest 5-year economics.
- Use AnalyticsApp-003 as the quick-win modernization pattern because it has low criticality and a comparatively simple operational footprint.
- Complete partially fulfilled cloud modernization for HRApp-004 to simplify the hybrid deployment model.
- Pursue ARM-based optimization only where workloads and runtime compatibility are validated first.

### Phase 4: Strategic disposition (18-24+ months)

Resolve long-horizon platform choices and avoid stranded investment.

- If ERPApp-001 remains in service beyond 2027, fund Application Refactoring & Decoupling and associated database/hosting transformation.
- Where applicable, evaluate open-source database switching only after application compatibility and support obligations are verified.
- Retain EComApp-005 as out of scope unless retirement status changes.

## Application Deep Dives

### app001 - ERPApp-001

**Description:** Core ERP system handling financial transactions, general ledger, and regulatory reporting

| Attribute | Value |
|---|---|
| Solution type | Custom made |
| Criticality | High |
| Status | Production |
| Deployment | On-Premise |
| Business unit | Finance |
| Architecture | 1-Tier |
| Containerized | No |
| CI/CD | No |
| API endpoints | 0 |
| External interfaces | 5 |
| Servers | sv01, sv02 |
| CPU | 4 cores |
| RAM | 16GB |
| Users | 350 |
| Database | Oracle 19c |
| DB storage | 1000 GB |
| DB license | Yes |
| Monitoring | unknown |
| Logging | unknown |
| Decommission date | 2027 |
| Complexity | 8 (1.8x multiplier) |

Technology assessment:

| Layer | Component | Status | Evidence |
|---|---|---|---|
| Operating System | AIX 7.2 | OUTDATED | IBM extended support available but proprietary and non-standard. |
| Language | COBOL-2014 | OUTDATED | Supported but legacy language footprint. |
| Application Server | None | N/A | No standalone application server is present. |
| Database | Oracle 19c | OUTDATED | Premier Support ended April 2024; Extended Support runs to April 2027. |

Applicable / partially fulfilled business case scenarios:

| Scenario | Status | Base Cost | Cost with Multiplier | Annual Savings | Breakeven (yrs) | 3-Year Net | 5-Year Net |
|---|---|---:|---:|---:|---:|---:|---:|
| OS Update & Security Patch | APPLICABLE | €1,000 | €1,800 | €500 | 3.6 | €-300 | €700 |
| Switch to Standard Linux OS | APPLICABLE | €300 | €540 | €400 | 1.35 | €660 | €1,460 |
| Lift & Shift to Cloud | APPLICABLE | €5,000 | €9,000 | €3,000 | 3.0 | €0 | €6,000 |
| Application Refactoring & Decoupling | APPLICABLE | €250,000 | €450,000 | €150,000 | 3.0 | €0 | €300,000 |
| Upgrade Legacy Databases | APPLICABLE | €10,000 | €18,000 | €10,000 | 1.8 | €12,000 | €32,000 |
| Switch DB Engine to Open Source | APPLICABLE | €25,000 | €45,000 | €15,000 | 3.0 | €0 | €30,000 |
| Update Outdated Components | APPLICABLE | €15,000 | €27,000 | €8,000 | 3.38 | €-3,000 | €13,000 |

Scenario statuses (full set):

| Scenario | Status |
|---|---|
| OS Update & Security Patch | APPLICABLE |
| Switch to Standard Linux OS | APPLICABLE |
| Switch to ARM-based CPU | NOT_APPLICABLE |
| Application Server Replacement | NOT_APPLICABLE |
| Lift & Shift to Cloud | APPLICABLE |
| Application Containerization | BLOCKED |
| Application Refactoring & Decoupling | APPLICABLE |
| Upgrade Legacy Databases | APPLICABLE |
| Switch DB Engine to Open Source | APPLICABLE |
| Update Outdated Components | APPLICABLE |

**Recommendation:** Contain risk immediately and use the 2027 decommission date as a decision gate: either retire on time with tactical remediation only, or fund a structured refactor if the platform must survive beyond 2027.

### app002 - CRMApp-002

**Description:** Customer relationship management system for tracking leads, opportunities, and customer interactions

| Attribute | Value |
|---|---|
| Solution type | 3rd party software |
| Criticality | Medium |
| Status | Production |
| Deployment | AWS |
| Business unit | Marketing |
| Architecture | unknown |
| Containerized | No |
| CI/CD | Yes |
| API endpoints | 15 |
| External interfaces | 8 |
| Servers | sv05, sv07 |
| CPU | 2 cores |
| RAM | 8GB |
| Users | 1200 |
| Database | Amazon RDS MySQL |
| DB storage | 500 GB |
| DB license | No |
| Monitoring | Datadog |
| Logging | CloudWatch |
| Decommission date | N/A |
| Complexity | 6 (1.5x multiplier) |

Technology assessment:

| Layer | Component | Status | Evidence |
|---|---|---|---|
| Operating System | RHEL 7 | EOL | End of life was June 30, 2024. |
| Language | Java 11 | OUTDATED | Premier support ended September 2023; some community patches still exist. |
| Application Server | WebSphere 7.0 | EOL | End of life was December 2015. |
| Database | Amazon RDS MySQL | CURRENT_VERSION | Managed service assumed current. |

Applicable / partially fulfilled business case scenarios:

| Scenario | Status | Base Cost | Cost with Multiplier | Annual Savings | Breakeven (yrs) | 3-Year Net | 5-Year Net |
|---|---|---:|---:|---:|---:|---:|---:|
| OS Update & Security Patch | APPLICABLE | €1,000 | €1,500 | €500 | 3.0 | €0 | €1,000 |

Scenario statuses (full set):

| Scenario | Status |
|---|---|
| OS Update & Security Patch | APPLICABLE |
| Switch to Standard Linux OS | FULFILLED |
| Switch to ARM-based CPU | NOT_APPLICABLE |
| Application Server Replacement | BLOCKED |
| Lift & Shift to Cloud | FULFILLED |
| Application Containerization | NOT_APPLICABLE |
| Application Refactoring & Decoupling | NOT_APPLICABLE |
| Upgrade Legacy Databases | FULFILLED |
| Switch DB Engine to Open Source | FULFILLED |
| Update Outdated Components | NOT_APPLICABLE |

**Recommendation:** Treat this as a vendor-constrained stabilization effort: patch the OS exposure immediately and pursue middleware remediation through the product vendor rather than a self-directed redesign.

### app003 - AnalyticsApp-003

**Description:** Analytics platform for generating operational reports and business insights from logistics data

| Attribute | Value |
|---|---|
| Solution type | Open Source |
| Criticality | Low |
| Status | Production |
| Deployment | AWS |
| Business unit | IT |
| Architecture | 3-Tier |
| Containerized | Yes |
| CI/CD | Yes |
| API endpoints | 8 |
| External interfaces | 3 |
| Servers | sv03 |
| CPU | 2 cores |
| RAM | 8GB |
| Users | 480 |
| Database | PostgreSQL 13 |
| DB storage | 200 GB |
| DB license | No |
| Monitoring | Prometheus |
| Logging | ELK |
| Decommission date | N/A |
| Complexity | 4 (1.3x multiplier) |

Technology assessment:

| Layer | Component | Status | Evidence |
|---|---|---|---|
| Operating System | RHEL 7 | EOL | End of life was June 30, 2024. |
| Language | Python 3.9 | EOL | Security support ended October 2025. |
| Application Server | Apache Tomcat 6.1 | EOL | End of life was December 2016. |
| Database | PostgreSQL 13 | EOL | End of life was November 13, 2025. |

Applicable / partially fulfilled business case scenarios:

| Scenario | Status | Base Cost | Cost with Multiplier | Annual Savings | Breakeven (yrs) | 3-Year Net | 5-Year Net |
|---|---|---:|---:|---:|---:|---:|---:|
| OS Update & Security Patch | APPLICABLE | €1,000 | €1,300 | €500 | 2.6 | €200 | €1,200 |
| Switch to ARM-based CPU | APPLICABLE | €5,000 | €6,500 | €1,000 | 6.5 | €-3,500 | €-1,500 |
| Application Server Replacement | APPLICABLE | €10,000 | €13,000 | €12,000 | 1.08 | €23,000 | €47,000 |
| Upgrade Legacy Databases | APPLICABLE | €10,000 | €13,000 | €10,000 | 1.3 | €17,000 | €37,000 |
| Update Outdated Components | APPLICABLE | €15,000 | €19,500 | €8,000 | 2.44 | €4,500 | €20,500 |

Scenario statuses (full set):

| Scenario | Status |
|---|---|
| OS Update & Security Patch | APPLICABLE |
| Switch to Standard Linux OS | FULFILLED |
| Switch to ARM-based CPU | APPLICABLE |
| Application Server Replacement | APPLICABLE |
| Lift & Shift to Cloud | FULFILLED |
| Application Containerization | FULFILLED |
| Application Refactoring & Decoupling | NOT_APPLICABLE |
| Upgrade Legacy Databases | APPLICABLE |
| Switch DB Engine to Open Source | FULFILLED |
| Update Outdated Components | APPLICABLE |

**Recommendation:** Prioritize this application as the quickest modernization win: it is already containerized and automated, so component upgrades should be lower friction and produce rapid risk reduction.

### app004 - HRApp-004

**Description:** Human resources management system handling employee records, benefits, and HR workflows

| Attribute | Value |
|---|---|
| Solution type | Custom made |
| Criticality | High |
| Status | Production |
| Deployment | AWS + On-premise (hybrid) |
| Business unit | HR |
| Architecture | 2-Tier |
| Containerized | Yes |
| CI/CD | Yes |
| API endpoints | 12 |
| External interfaces | 6 |
| Servers | sv06, sv02 |
| CPU | 4 cores |
| RAM | 16GB |
| Users | 670 |
| Database | SQL Server 2019 |
| DB storage | 750 GB |
| DB license | Yes |
| Monitoring | CloudWatch |
| Logging | Splunk |
| Decommission date | N/A |
| Complexity | 5 (1.4x multiplier) |

Technology assessment:

| Layer | Component | Status | Evidence |
|---|---|---|---|
| Operating System | Windows Server 2012 | EOL | Extended support ended October 10, 2023. |
| Language | .NET Core (version unspecified) | OUTDATED | Version unspecified; treat as outdated with low confidence. |
| Application Server | Microsoft IIS 8.0 | EOL | Tied to Windows Server 2012 and therefore effectively end of life. |
| Database | SQL Server 2019 | OUTDATED | Mainstream support ended January 7, 2025; Extended Support runs to January 2030. |

Applicable / partially fulfilled business case scenarios:

| Scenario | Status | Base Cost | Cost with Multiplier | Annual Savings | Breakeven (yrs) | 3-Year Net | 5-Year Net |
|---|---|---:|---:|---:|---:|---:|---:|
| OS Update & Security Patch | APPLICABLE | €1,000 | €1,400 | €500 | 2.8 | €100 | €1,100 |
| Switch to ARM-based CPU | APPLICABLE | €5,000 | €7,000 | €1,000 | 7.0 | €-4,000 | €-2,000 |
| Application Server Replacement | APPLICABLE | €10,000 | €14,000 | €12,000 | 1.17 | €22,000 | €46,000 |
| Lift & Shift to Cloud | PARTIALLY_FULFILLED | €5,000 | €7,000 | €3,000 | 2.33 | €2,000 | €8,000 |
| Application Refactoring & Decoupling | APPLICABLE | €250,000 | €350,000 | €150,000 | 2.33 | €100,000 | €400,000 |
| Upgrade Legacy Databases | APPLICABLE | €10,000 | €14,000 | €10,000 | 1.4 | €16,000 | €36,000 |
| Switch DB Engine to Open Source | APPLICABLE | €25,000 | €35,000 | €15,000 | 2.33 | €10,000 | €40,000 |
| Update Outdated Components | APPLICABLE | €15,000 | €21,000 | €8,000 | 2.62 | €3,000 | €19,000 |

Scenario statuses (full set):

| Scenario | Status |
|---|---|
| OS Update & Security Patch | APPLICABLE |
| Switch to Standard Linux OS | NOT_APPLICABLE |
| Switch to ARM-based CPU | APPLICABLE |
| Application Server Replacement | APPLICABLE |
| Lift & Shift to Cloud | PARTIALLY_FULFILLED |
| Application Containerization | FULFILLED |
| Application Refactoring & Decoupling | APPLICABLE |
| Upgrade Legacy Databases | APPLICABLE |
| Switch DB Engine to Open Source | APPLICABLE |
| Update Outdated Components | APPLICABLE |

**Recommendation:** Use this as a strategic modernization target: the platform already has containers and CI/CD, but its Windows/IIS foundation is EOL and the hybrid footprint increases operational complexity.

## Appendix: Scenario Definitions

| Scenario Key | Scenario Name | Priority | Base Cost | Base Annual Savings | Description |
|---|---|---|---:|---:|---|
| os_update_security_patch | OS Update & Security Patch | High | €1,000 | €500 | Update operating systems to supported patch levels and close security gaps. |
| switch_to_standard_linux_os | Switch to Standard Linux OS | Medium | €300 | €400 | Move from proprietary or non-standard operating systems to standard Linux. |
| switch_to_arm_cpu | Switch to ARM-based CPU | Medium | €5,000 | €1,000 | Migrate workloads to ARM-compatible infrastructure to reduce compute cost. |
| application_server_replacement | Application Server Replacement | Medium | €10,000 | €12,000 | Replace obsolete middleware or application servers with supported alternatives. |
| app_deployment_to_cloud | Lift & Shift to Cloud | High | €5,000 | €3,000 | Move workloads to cloud infrastructure without major application redesign. |
| app_containerization | Application Containerization | High | €100,000 | €100,000 | Package workloads into containers and standardize runtime operations. |
| app_refactor_decoupling | Application Refactoring & Decoupling | High | €250,000 | €150,000 | Refactor tightly coupled applications into more modular architectures. |
| upgrade_legacy_databases | Upgrade Legacy Databases | High | €10,000 | €10,000 | Upgrade database platforms that are close to or beyond support deadlines. |
| switch_db_engine_open_source | Switch DB Engine to Open Source | High | €25,000 | €15,000 | Move commercial or legacy database engines to open-source alternatives where feasible. |
| update_outdated_components | Update Outdated Components | High | €15,000 | €8,000 | Refresh outdated language runtimes, frameworks, and dependent components. |
