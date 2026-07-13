# Application Report — AnalyticsApp-003

> Generated: 2026-07-13 | Application ID: `app003`

## Application Overview

| Field | Value |
|-------|-------|
| **App ID** | `app003` |
| **Name** | AnalyticsApp-003 |
| **Description** | Analytics platform |
| **Solution Type** | Open Source |
| **Business Criticality** | Low |
| **Status** | Production |
| **Decommission Date** | not planned |
| **Deployment Type** | AWS |
| **Data Classification** | Public |
| **Business Unit** | IT |
| **Business Capabilities** | Data Analytics, Reporting, Business Intelligence |
| **User Count** | 480 |
| **Operating System** | RHEL 7 |
| **Programming Language** | Python 3.9 |
| **Application Server** | Apache Tomcat 6.1 |
| **Architecture** | 3-Tier |
| **Containerized** | Yes |
| **CI/CD** | Yes |
| **Environment Count** | 1 |
| **Servers** | sv03 |
| **CPU Cores** | 2 |
| **Memory (GB)** | 8 |
| **API Endpoints** | 8 |
| **External Interfaces** | 3 |
| **Database Engine** | PostgreSQL 13 |
| **Database Storage (GB)** | 200 |
| **Database License Required** | No |
| **Logging Solution** | ELK |
| **Monitoring Tool** | Prometheus |

## Technology Assessment

🔴 **EOL components present**

| Component | Type | Version | Status | EOL Date | Notes |
|-----------|------|---------|--------|----------|-------|
| **RHEL 7** | os | 7 | 🔴 EOL | 2024-06-30 | Red Hat Enterprise Linux 7 reached End of Life on June 30, 2024. Security patches are no longer provided without ELS subscription. |
| **Python 3.9** | programming_language | 3.9 | 🔴 EOL | 2025-10-05 | Python 3.9 reached End of Life on October 5, 2025. No further security updates are released by the Python Software Foundation. Python 3.12+ is recommended. |
| **Apache Tomcat 6.1** | application_server | 6.1 | 🔴 EOL | 2016-12-31 | Apache Tomcat 6.x reached End of Life in December 2016. This version has known unpatched security vulnerabilities and should be urgently migrated. |
| **PostgreSQL 13** | database | 13 | 🔴 EOL | 2025-11-13 | PostgreSQL 13 reached End of Life on November 13, 2025. No further bug fix or security releases are available. PostgreSQL 16 or 17 is recommended. |

## Complexity Assessment

**Complexity Score: 4/10 — 🟡 Medium**

### Complexity Factors

| Factor | Value |
|--------|-------|
| Business Criticality | Low |
| Architecture Tier | 3-Tier |
| Containerization | True |
| Ci Cd | True |
| Server Count | 1 |
| External Interface Count | 3 |
| Eol Components | True |
| Outdated Components | False |
| Technology Age | mixed |

### Reasoning

AnalyticsApp-003 scores Medium complexity (4/10). The application has several modernization-friendly attributes: it is already containerized, uses a 3-Tier architecture, has CI/CD in place, and runs on a single server. It has Low business criticality and a small number of external interfaces (3), reducing risk during modernization. However, all core components are EOL: RHEL 7 (EOL June 2024), Python 3.9 (EOL October 2025), Apache Tomcat 6.1 (EOL December 2016), and PostgreSQL 13 (EOL November 2025). The breadth of EOL components creates security exposure that must be addressed promptly. The existing containerization and CI/CD pipeline make updates technically straightforward. The open-source nature of the stack enables cost-effective version upgrades.

## Scenario Applicability

| Scenario | Status | Priority | Effort | Summary |
|----------|--------|----------|--------|---------|
| **Operating System Update** | 🔵 APPLICABLE | High | Low | RHEL 7 is EOL (June 30, 2024) and no longer receives security patches. The application is containeri... |
| **Switch to Standard Linux Operating System** | ✅ FULFILLED | N/A | N/A | AnalyticsApp-003 already runs on RHEL, a standard enterprise Linux distribution. The scenario is ful... |
| **Switch to ARM-based CPU** | 🔵 APPLICABLE | Medium | Medium | The application is containerized with CI/CD on AWS, making it a candidate for ARM-based EC2 instance... |
| **Application Server Replacement** | 🔵 APPLICABLE | Critical | Medium | Apache Tomcat 6.1 is severely EOL (December 2016) and has numerous unpatched CVEs. Replacement with ... |
| **Application Migration to Cloud Infrastructure (Lift & Shift)** | ✅ FULFILLED | N/A | N/A | AnalyticsApp-003 is already deployed on AWS. Cloud deployment is fulfilled. |
| **Application Containerization** | ✅ FULFILLED | N/A | N/A | AnalyticsApp-003 is already containerized. Container deployment is fully implemented. |
| **Application Refactoring and De-coupling** | 🔴 BLOCKED | Low | Very High | AnalyticsApp-003 is an Open Source platform. Direct source code refactoring may conflict with the op... |
| **Upgrade Legacy Databases** | 🔵 APPLICABLE | High | Low | PostgreSQL 13 reached EOL on November 13, 2025. Upgrading to PostgreSQL 16 or 17 is needed to restor... |
| **Switch DB Engine to Open-Source Database Solution** | ✅ FULFILLED | N/A | N/A | PostgreSQL is already an open-source database. No commercial license is required (database_license_r... |
| **Update Outdated Components** | 🔵 APPLICABLE | Critical | Medium | All components are EOL (RHEL 7, Python 3.9, Tomcat 6.1, PostgreSQL 13). Systematic component updates... |

### Scenario Details

#### 🔵 Operating System Update
- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** Low
- **Reasoning:** RHEL 7 is EOL (June 30, 2024) and no longer receives security patches. The application is containerized and has CI/CD, making OS updates operationally straightforward.
- **Suggestion:** Upgrade the container base image and host OS from RHEL 7 to RHEL 9 or equivalent (e.g., Rocky Linux 9, AlmaLinux 9). Update CI/CD pipeline to use new base images.

#### ✅ Switch to Standard Linux Operating System
- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** AnalyticsApp-003 already runs on RHEL, a standard enterprise Linux distribution. The scenario is fulfilled; only version upgrade is needed.
- **Suggestion:** Linux OS family is already standard. Proceed with RHEL 7 → RHEL 9 version upgrade.

#### 🔵 Switch to ARM-based CPU
- **Status:** APPLICABLE
- **Priority:** Medium
- **Effort:** Medium
- **Reasoning:** The application is containerized with CI/CD on AWS, making it a candidate for ARM-based EC2 instances (AWS Graviton). Python and PostgreSQL have strong ARM64 support. Apache Tomcat compatibility with ARM must be validated.
- **Suggestion:** After upgrading to a supported Tomcat version, test on AWS Graviton (ARM64) EC2 instances. Update CI/CD to build multi-arch container images.

#### 🔵 Application Server Replacement
- **Status:** APPLICABLE
- **Priority:** Critical
- **Effort:** Medium
- **Reasoning:** Apache Tomcat 6.1 is severely EOL (December 2016) and has numerous unpatched CVEs. Replacement with Apache Tomcat 10.x or 11.x (current) is urgently needed.
- **Suggestion:** Upgrade Apache Tomcat 6.1 to Tomcat 10.1 or 11.0. Note: Tomcat 10+ uses Jakarta EE namespace; migration from javax.* to jakarta.* packages may be required.

#### ✅ Application Migration to Cloud Infrastructure (Lift & Shift)
- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** AnalyticsApp-003 is already deployed on AWS. Cloud deployment is fulfilled.
- **Suggestion:** Cloud deployment already fulfilled. Consider AWS-native analytics services (EMR, Redshift, QuickSight) for future optimization.

#### ✅ Application Containerization
- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** AnalyticsApp-003 is already containerized. Container deployment is fully implemented.
- **Suggestion:** Containerization is fulfilled. Improve container security by updating base images to remove EOL OS layers.

#### 🔴 Application Refactoring and De-coupling
- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** Very High
- **Reasoning:** AnalyticsApp-003 is an Open Source platform. Direct source code refactoring may conflict with the open-source project's architecture and community roadmap.
- **Suggestion:** For open-source platforms, contribute to upstream projects rather than forking. Evaluate whether replacing with a managed analytics service is more efficient than custom refactoring.

#### 🔵 Upgrade Legacy Databases
- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** Low
- **Reasoning:** PostgreSQL 13 reached EOL on November 13, 2025. Upgrading to PostgreSQL 16 or 17 is needed to restore security patch coverage.
- **Suggestion:** Use pg_upgrade or pg_logical replication to upgrade from PostgreSQL 13 to PostgreSQL 16 or 17. Test application compatibility with the new version before production migration.

#### ✅ Switch DB Engine to Open-Source Database Solution
- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** PostgreSQL is already an open-source database. No commercial license is required (database_license_required: false).
- **Suggestion:** Open-source database requirement is fulfilled. Focus on upgrading PostgreSQL version rather than switching engines.

#### 🔵 Update Outdated Components
- **Status:** APPLICABLE
- **Priority:** Critical
- **Effort:** Medium
- **Reasoning:** All components are EOL (RHEL 7, Python 3.9, Tomcat 6.1, PostgreSQL 13). Systematic component updates are required across the full stack. The containerized, CI/CD-enabled setup makes updates highly feasible.
- **Suggestion:** Create a systematic component update plan: (1) Update base OS image to RHEL 9/Rocky 9, (2) Upgrade Python to 3.12+, (3) Upgrade Tomcat to 10.1+, (4) Upgrade PostgreSQL to 16+. Execute through the existing CI/CD pipeline.

## Business Case

**Total Adjusted Investment: $20,800.00**  
**Total Yearly Savings: $23,500.00**  
**Estimated ROI Payback: 0.89 years**

| Scenario | Status | Adj. Cost | Yearly Savings | ROI (years) |
|----------|--------|-----------|----------------|-------------|
| Operating System Update | APPLICABLE | $800.00 | $500.00 | 1.6 |
| Switch to ARM-based CPU | APPLICABLE | $4,000.00 | $1,000.00 | 4.0 |
| Application Server Replacement | APPLICABLE | $8,000.00 | $12,000.00 | 0.67 |
| Upgrade Legacy Databases | APPLICABLE | $8,000.00 | $10,000.00 | 0.8 |

> *Costs adjusted by complexity multiplier: **0.8x** (complexity score / 5)*
