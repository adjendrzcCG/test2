# Application Report — CRMApp-002

> Generated: 2026-07-13 | Application ID: `app002`

## Application Overview

| Field | Value |
|-------|-------|
| **App ID** | `app002` |
| **Name** | CRMApp-002 |
| **Description** | CRM system |
| **Solution Type** | 3rd party |
| **Business Criticality** | Medium |
| **Status** | Production |
| **Decommission Date** | not planned |
| **Deployment Type** | AWS |
| **Data Classification** | Internal |
| **Business Unit** | Marketing |
| **Business Capabilities** | Customer Management, Lead Tracking, Sales Analytics |
| **User Count** | 1200 |
| **Operating System** | RHEL 7 |
| **Programming Language** | Java 11 |
| **Application Server** | Websphere 7.0 |
| **Architecture** | unknown |
| **Containerized** | No |
| **CI/CD** | Yes |
| **Environment Count** | 2 |
| **Servers** | sv05, sv07 |
| **CPU Cores** | 2 |
| **Memory (GB)** | 8 |
| **API Endpoints** | 15 |
| **External Interfaces** | 8 |
| **Database Engine** | Amazon RDS MySQL |
| **Database Storage (GB)** | 500 |
| **Database License Required** | No |
| **Logging Solution** | CloudWatch |
| **Monitoring Tool** | Datadog |

## Technology Assessment

🔴 **EOL components present** | 🟡 **Outdated components present**

| Component | Type | Version | Status | EOL Date | Notes |
|-----------|------|---------|--------|----------|-------|
| **RHEL 7** | os | 7 | 🔴 EOL | 2024-06-30 | Red Hat Enterprise Linux 7 reached End of Life on June 30, 2024. No further security patches or bug fixes are available without a paid ELS subscription. |
| **Java 11** | programming_language | 11 | 🟡 OUTDATED | 2026-09-30 | Java 11 LTS is in extended support; Oracle's free public updates ended in September 2023. Java 17 and 21 are current LTS releases. Extended support continues through September 2026. |
| **Websphere 7.0** | application_server | 7.0 | 🔴 EOL | 2015-09-30 | IBM WebSphere Application Server 7.0 reached End of Service on September 30, 2015. Running this version poses significant security and compliance risks. |
| **Amazon RDS MySQL** | database | managed | ✅ CURRENT_VERSION | N/A | Amazon RDS for MySQL is a managed service that receives automatic patching and version updates from AWS. AWS manages security and compatibility. *(managed)* |

## Complexity Assessment

**Complexity Score: 7/10 — 🟠 High**

### Complexity Factors

| Factor | Value |
|--------|-------|
| Business Criticality | Medium |
| Architecture Tier | unknown |
| Containerization | False |
| Ci Cd | True |
| Server Count | 2 |
| External Interface Count | 8 |
| Eol Components | True |
| Outdated Components | True |
| Technology Age | aging |

### Reasoning

CRMApp-002 scores High complexity (7/10). As a 3rd-party CRM solution hosted on AWS, it benefits from cloud infrastructure and an existing CI/CD pipeline. However, it runs on EOL RHEL 7 and uses EOL WebSphere 7.0 (EOL since 2015), which creates significant security and compliance exposure. Java 11 is in extended support and approaching EOL. The application is not containerized, and its architecture is unknown, adding discovery risk. With 1,200 users across Marketing and 8 external interfaces plus 15 API endpoints, integration surface is considerable. The 3rd-party nature means modernization may be constrained by vendor roadmaps. No database license cost and existing monitoring tools (CloudWatch, Datadog) reduce some operational complexity.

## Scenario Applicability

| Scenario | Status | Priority | Effort | Summary |
|----------|--------|----------|--------|---------|
| **Operating System Update** | 🔵 APPLICABLE | Critical | Medium | RHEL 7 is EOL as of June 30, 2024. The OS is no longer receiving security patches, exposing a 1,200-... |
| **Switch to Standard Linux Operating System** | ✅ FULFILLED | N/A | N/A | CRMApp-002 already runs on RHEL, which is a standard enterprise Linux OS. The requirement for a stan... |
| **Switch to ARM-based CPU** | 🔴 BLOCKED | Low | Very High | WebSphere Application Server 7.0 does not support ARM architecture. Without replacing the EOL applic... |
| **Application Server Replacement** | 🔵 APPLICABLE | Critical | High | WebSphere 7.0 is severely EOL (since September 2015) and represents the highest risk component. Repl... |
| **Application Migration to Cloud Infrastructure (Lift & Shift)** | ✅ FULFILLED | N/A | N/A | CRMApp-002 is already deployed on AWS. The cloud deployment requirement is fulfilled. |
| **Application Containerization** | 🔴 BLOCKED | Low | Very High | The 3rd-party CRM cannot be easily containerized without vendor support. WebSphere 7.0 is EOL and no... |
| **Application Refactoring and De-coupling** | 🔴 BLOCKED | Low | N/A | As a 3rd-party application, source code refactoring is not available. Decoupling would need to happe... |
| **Upgrade Legacy Databases** | ✅ FULFILLED | N/A | N/A | Amazon RDS MySQL is a managed service that AWS automatically keeps updated. Database version managem... |
| **Switch DB Engine to Open-Source Database Solution** | ✅ FULFILLED | N/A | N/A | Amazon RDS MySQL is already an open-source database engine (MySQL is open-source). No commercial dat... |
| **Update Outdated Components** | 🔴 BLOCKED | High | High | RHEL 7 and WebSphere 7.0 are EOL, not just outdated. Java 11 is approaching EOL. For a 3rd-party app... |

### Scenario Details

#### 🔵 Operating System Update
- **Status:** APPLICABLE
- **Priority:** Critical
- **Effort:** Medium
- **Reasoning:** RHEL 7 is EOL as of June 30, 2024. The OS is no longer receiving security patches, exposing a 1,200-user CRM application to unpatched vulnerabilities.
- **Suggestion:** Upgrade from RHEL 7 to RHEL 9 (current) using in-place upgrade tools (Leapp) or re-platform to a new RHEL 9 instance. Validate Java 11 and WebSphere compatibility.

#### ✅ Switch to Standard Linux Operating System
- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** CRMApp-002 already runs on RHEL, which is a standard enterprise Linux OS. The requirement for a standard Linux OS is met; the OS version upgrade is tracked under os_update_security_patch.
- **Suggestion:** OS family already fulfilled. Proceed with RHEL 7 → RHEL 9 upgrade.

#### 🔴 Switch to ARM-based CPU
- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** Very High
- **Reasoning:** WebSphere Application Server 7.0 does not support ARM architecture. Without replacing the EOL application server first, ARM migration is not feasible.
- **Suggestion:** Replace WebSphere 7.0 with a modern application server (Tomcat, WildFly, or Liberty) before evaluating ARM migration.

#### 🔵 Application Server Replacement
- **Status:** APPLICABLE
- **Priority:** Critical
- **Effort:** High
- **Reasoning:** WebSphere 7.0 is severely EOL (since September 2015) and represents the highest risk component. Replacement with a modern, supported application server is urgently needed.
- **Suggestion:** Migrate from WebSphere 7.0 to IBM Liberty, WildFly, or Apache Tomcat (latest). For a 3rd-party CRM, check vendor support for alternative application servers before migrating.

#### ✅ Application Migration to Cloud Infrastructure (Lift & Shift)
- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** CRMApp-002 is already deployed on AWS. The cloud deployment requirement is fulfilled.
- **Suggestion:** Cloud deployment is already fulfilled. Focus on modernizing EOL components (RHEL 7, WebSphere 7.0) in the existing AWS environment.

#### 🔴 Application Containerization
- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** Very High
- **Reasoning:** The 3rd-party CRM cannot be easily containerized without vendor support. WebSphere 7.0 is EOL and not designed for container deployment. Architecture is unknown, adding discovery risk.
- **Suggestion:** Containerization is blocked for a 3rd-party application on EOL WebSphere. Evaluate vendor-provided containerized versions or SaaS alternatives.

#### 🔴 Application Refactoring and De-coupling
- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** N/A
- **Reasoning:** As a 3rd-party application, source code refactoring is not available. Decoupling would need to happen via API layer additions, which may conflict with vendor licensing terms.
- **Suggestion:** Refactoring is blocked for 3rd-party software. Evaluate vendor API capabilities or plan replacement with a SaaS CRM (e.g., Salesforce, HubSpot).

#### ✅ Upgrade Legacy Databases
- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** Amazon RDS MySQL is a managed service that AWS automatically keeps updated. Database version management is handled by AWS.
- **Suggestion:** Database lifecycle management is handled by AWS RDS. Ensure the RDS MySQL major version is current (8.0+) and enable auto minor version upgrades.

#### ✅ Switch DB Engine to Open-Source Database Solution
- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** Amazon RDS MySQL is already an open-source database engine (MySQL is open-source). No commercial database license is required (database_license_required: false).
- **Suggestion:** Open-source database requirement is already fulfilled with Amazon RDS MySQL.

#### 🔴 Update Outdated Components
- **Status:** BLOCKED
- **Priority:** High
- **Effort:** High
- **Reasoning:** RHEL 7 and WebSphere 7.0 are EOL, not just outdated. Java 11 is approaching EOL. For a 3rd-party application, component updates are constrained by vendor certification requirements — updating RHEL or Java may break the certified support matrix.
- **Suggestion:** Coordinate with the CRM vendor to obtain a certified upgrade path. Check vendor support matrix for RHEL 9 and Java 17/21 compatibility.

## Business Case

**Total Adjusted Investment: $15,400.00**  
**Total Yearly Savings: $12,500.00**  
**Estimated ROI Payback: 1.23 years**

| Scenario | Status | Adj. Cost | Yearly Savings | ROI (years) |
|----------|--------|-----------|----------------|-------------|
| Operating System Update | APPLICABLE | $1,400.00 | $500.00 | 2.8 |
| Application Server Replacement | APPLICABLE | $14,000.00 | $12,000.00 | 1.17 |

> *Costs adjusted by complexity multiplier: **1.4x** (complexity score / 5)*
