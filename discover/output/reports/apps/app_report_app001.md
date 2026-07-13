# Application Report — ERPApp-001

> Generated: 2026-07-13 | Application ID: `app001`

## Application Overview

| Field | Value |
|-------|-------|
| **App ID** | `app001` |
| **Name** | ERPApp-001 |
| **Description** | Core ERP system |
| **Solution Type** | Custom made |
| **Business Criticality** | High |
| **Status** | Production |
| **Decommission Date** | 2027 |
| **Deployment Type** | On-Premise |
| **Data Classification** | Confidential |
| **Business Unit** | Finance |
| **Business Capabilities** | Financial Planning, Accounting, Budgeting |
| **User Count** | 350 |
| **Operating System** | AIX 7.2 |
| **Programming Language** | COBOL-2014 |
| **Application Server** | None |
| **Architecture** | 1-Tier |
| **Containerized** | No |
| **CI/CD** | No |
| **Environment Count** | 2 |
| **Servers** | sv01, sv02 |
| **CPU Cores** | 4 |
| **Memory (GB)** | 16 |
| **API Endpoints** | 0 |
| **External Interfaces** | 5 |
| **Database Engine** | Oracle 19c |
| **Database Storage (GB)** | 1000 |
| **Database License Required** | Yes |
| **Logging Solution** | None |
| **Monitoring Tool** | None |

## Technology Assessment

🟡 **Outdated components present**

| Component | Type | Version | Status | EOL Date | Notes |
|-----------|------|---------|--------|----------|-------|
| **AIX 7.2** | os | 7.2 | 🟡 OUTDATED | N/A | AIX 7.2 standard support has ended; only extended support available. IBM has shifted focus to newer AIX and Power platform versions. |
| **COBOL-2014** | programming_language | 2014 | ✅ CURRENT_VERSION | N/A | COBOL-2014 is the current ISO/IEC standard for COBOL. Mainframe COBOL compilers (GnuCOBOL, IBM Enterprise COBOL) continue to support the 2014 standard actively. |
| **Oracle 19c** | database | 19c | ✅ CURRENT_VERSION | 2027-04-30 | Oracle 19c is the long-term release with Premier Support through April 2024 and Extended Support through April 2027. Currently under extended support with full patches available. |

## Complexity Assessment

**Complexity Score: 9/10 — 🔴 Very High**

### Complexity Factors

| Factor | Value |
|--------|-------|
| Business Criticality | High |
| Architecture Tier | 1-Tier |
| Containerization | False |
| Ci Cd | False |
| Server Count | 2 |
| External Interface Count | 5 |
| Eol Components | False |
| Outdated Components | True |
| Technology Age | legacy |

### Reasoning

ERPApp-001 scores Very High complexity (9/10) due to multiple compounding factors. The application is a monolithic 1-Tier COBOL system running on AIX 7.2 — a highly proprietary, aging technology stack with very limited modernization tooling. There is no CI/CD pipeline and no containerization, indicating a manual, fragile deployment process. The application carries High business criticality with Confidential data classification and serves core Finance functions (Financial Planning, Accounting, Budgeting). With 0 API endpoints and 5 external interfaces, integration modernization will require significant re-engineering. The 1000GB Oracle database and decommission target of 2027 add time pressure. The COBOL codebase requires specialized skills that are increasingly rare, further elevating modernization risk and effort.

## Scenario Applicability

| Scenario | Status | Priority | Effort | Summary |
|----------|--------|----------|--------|---------|
| **Operating System Update** | 🔵 APPLICABLE | High | Medium | AIX 7.2 is OUTDATED with standard support ended. Security patches are limited to extended support co... |
| **Switch to Standard Linux Operating System** | 🔵 APPLICABLE | High | High | AIX 7.2 is a proprietary IBM UNIX platform with limited vendor ecosystem support. Migrating to RHEL ... |
| **Switch to ARM-based CPU** | 🔴 BLOCKED | Low | Very High | COBOL-2014 on AIX 7.2 is tightly coupled to IBM Power architecture. COBOL compilers with full ARM64 ... |
| **Application Server Replacement** | ⚪ NOT_APPLICABLE | N/A | N/A | ERPApp-001 does not use an application server (application_server: None). The COBOL application runs... |
| **Application Migration to Cloud Infrastructure (Lift & Shift)** | 🔵 APPLICABLE | Medium | High | ERPApp-001 is currently On-Premise with no cloud footprint. A lift-and-shift migration to cloud (e.g... |
| **Application Containerization** | 🔴 BLOCKED | Low | Very High | COBOL on AIX 7.2 with a 1-Tier monolithic architecture and no CI/CD pipeline is not suitable for con... |
| **Application Refactoring and De-coupling** | 🔵 APPLICABLE | High | Very High | The 1-Tier monolithic COBOL architecture with 5 external interfaces and a 2027 decommission target s... |
| **Upgrade Legacy Databases** | ✅ FULFILLED | Low | N/A | Oracle 19c is the current LTS version with support through April 2027. No database upgrade is needed... |
| **Switch DB Engine to Open-Source Database Solution** | 🔵 APPLICABLE | Medium | High | Oracle 19c requires a commercial license (database_license_required: true). Migrating to PostgreSQL ... |
| **Update Outdated Components** | 🔵 APPLICABLE | High | Medium | AIX 7.2 is OUTDATED. Updating to the latest AIX 7.3 or planning OS migration addresses security gaps... |

### Scenario Details

#### 🔵 Operating System Update
- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** Medium
- **Reasoning:** AIX 7.2 is OUTDATED with standard support ended. Security patches are limited to extended support contracts, increasing vulnerability exposure for a High-criticality Confidential application.
- **Suggestion:** Update AIX to the latest supported version or plan migration to a standard Linux distribution to reduce proprietary platform risk.

#### 🔵 Switch to Standard Linux Operating System
- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** High
- **Reasoning:** AIX 7.2 is a proprietary IBM UNIX platform with limited vendor ecosystem support. Migrating to RHEL 9 or similar standard Linux would reduce licensing costs and broaden tooling availability. However, COBOL portability must be validated.
- **Suggestion:** Plan a phased migration from AIX 7.2 to RHEL 9. Validate COBOL compiler compatibility (GnuCOBOL or IBM COBOL for Linux) before migrating.

#### 🔴 Switch to ARM-based CPU
- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** Very High
- **Reasoning:** COBOL-2014 on AIX 7.2 is tightly coupled to IBM Power architecture. COBOL compilers with full ARM64 support are limited, and the application has no CI/CD or containerization to facilitate cross-architecture builds.
- **Suggestion:** This scenario is blocked until the application is modernized to a portable language/runtime stack and containerized.

#### ⚪ Application Server Replacement
- **Status:** NOT_APPLICABLE
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** ERPApp-001 does not use an application server (application_server: None). The COBOL application runs directly on AIX without a middleware tier.
- **Suggestion:** Not applicable. No application server is present to replace.

#### 🔵 Application Migration to Cloud Infrastructure (Lift & Shift)
- **Status:** APPLICABLE
- **Priority:** Medium
- **Effort:** High
- **Reasoning:** ERPApp-001 is currently On-Premise with no cloud footprint. A lift-and-shift migration to cloud (e.g., IBM Power Virtual Server on IBM Cloud, or AWS Bare Metal) would improve scalability and reduce on-premise hardware costs.
- **Suggestion:** Evaluate IBM Power Virtual Server (IBM Cloud) as an AIX-compatible lift-and-shift target to maintain COBOL/AIX compatibility while moving to cloud.

#### 🔴 Application Containerization
- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** Very High
- **Reasoning:** COBOL on AIX 7.2 with a 1-Tier monolithic architecture and no CI/CD pipeline is not suitable for containerization without major refactoring. There is no container runtime or orchestration infrastructure in place.
- **Suggestion:** Containerization is blocked until the COBOL application is refactored or rewritten to a portable technology stack and a CI/CD pipeline is established.

#### 🔵 Application Refactoring and De-coupling
- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** Very High
- **Reasoning:** The 1-Tier monolithic COBOL architecture with 5 external interfaces and a 2027 decommission target strongly suggests refactoring is necessary. Breaking the monolith into services would enable incremental modernization.
- **Suggestion:** Conduct a COBOL code analysis (using tools like Micro Focus, IBM Rational) to identify service boundaries. Gradually expose business functions as APIs to enable phased modernization.

#### ✅ Upgrade Legacy Databases
- **Status:** FULFILLED
- **Priority:** Low
- **Effort:** N/A
- **Reasoning:** Oracle 19c is the current LTS version with support through April 2027. No database upgrade is needed at this time.
- **Suggestion:** Monitor Oracle 19c end-of-extended-support (April 2027) and plan upgrade to Oracle 21c or 23ai before that date.

#### 🔵 Switch DB Engine to Open-Source Database Solution
- **Status:** APPLICABLE
- **Priority:** Medium
- **Effort:** High
- **Reasoning:** Oracle 19c requires a commercial license (database_license_required: true). Migrating to PostgreSQL or other open-source alternatives would eliminate Oracle licensing costs. However, Oracle-specific SQL and PL/SQL must be migrated.
- **Suggestion:** Evaluate AWS Aurora PostgreSQL or standard PostgreSQL as migration targets. Use tools like AWS Schema Conversion Tool (SCT) or Ora2Pg to assess migration complexity.

#### 🔵 Update Outdated Components
- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** Medium
- **Reasoning:** AIX 7.2 is OUTDATED. Updating to the latest AIX 7.3 or planning OS migration addresses security gaps and extends vendor support coverage for this High-criticality application.
- **Suggestion:** Engage IBM support to update AIX 7.2 to AIX 7.3 if hardware permits, or plan OS migration. Review COBOL compiler and runtime versions for available updates.

## Business Case

**Total Adjusted Investment: $506,340.00**  
**Total Yearly Savings: $168,900.00**  
**Estimated ROI Payback: 3.0 years**

| Scenario | Status | Adj. Cost | Yearly Savings | ROI (years) |
|----------|--------|-----------|----------------|-------------|
| Operating System Update | APPLICABLE | $1,800.00 | $500.00 | 3.6 |
| Switch to Standard Linux Operating System | APPLICABLE | $540.00 | $400.00 | 1.35 |
| Application Migration to Cloud Infrastructure (Lift & Shift) | APPLICABLE | $9,000.00 | $3,000.00 | 3.0 |
| Application Refactoring and De-coupling | APPLICABLE | $450,000.00 | $150,000.00 | 3.0 |
| Switch DB Engine to Open-Source Database Solution | APPLICABLE | $45,000.00 | $15,000.00 | 3.0 |

> *Costs adjusted by complexity multiplier: **1.8x** (complexity score / 5)*
