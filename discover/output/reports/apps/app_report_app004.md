# Application Report — HRApp-004

> Generated: 2026-07-13 | Application ID: `app004`

## Application Overview

| Field | Value |
|-------|-------|
| **App ID** | `app004` |
| **Name** | HRApp-004 |
| **Description** | HR management system |
| **Solution Type** | Custom made |
| **Business Criticality** | High |
| **Status** | Production |
| **Decommission Date** | not planned |
| **Deployment Type** | AWS+On-premise |
| **Data Classification** | Internal |
| **Business Unit** | HR |
| **Business Capabilities** | Employee Management, Payroll Processing, Benefits Administration |
| **User Count** | 670 |
| **Operating System** | Windows Server 2012 |
| **Programming Language** | .NET Core |
| **Application Server** | Microsoft IIS 8.0 |
| **Architecture** | 2-Tier |
| **Containerized** | Yes |
| **CI/CD** | Yes |
| **Environment Count** | 2 |
| **Servers** | sv06, sv02 |
| **CPU Cores** | 4 |
| **Memory (GB)** | 16 |
| **API Endpoints** | 12 |
| **External Interfaces** | 6 |
| **Database Engine** | SQL Server 2019 |
| **Database Storage (GB)** | 750 |
| **Database License Required** | Yes |
| **Logging Solution** | Splunk |
| **Monitoring Tool** | CloudWatch |

## Technology Assessment

🔴 **EOL components present** | 🟡 **Outdated components present** | ❓ **Missing version data**

| Component | Type | Version | Status | EOL Date | Notes |
|-----------|------|---------|--------|----------|-------|
| **Windows Server 2012** | os | 2012 | 🔴 EOL | 2023-10-10 | Windows Server 2012 reached End of Extended Support on October 10, 2023. Microsoft no longer provides security updates, exposing the system to unpatched vulnerabilities. |
| **.NET Core** | programming_language | unknown | ❓ NO_KNOWLEDGE | N/A | The .NET Core version is not specified. Without knowing the exact version (e.g., 3.1, 5.0, 6.0, 8.0), lifecycle status cannot be determined. Version discovery is required. |
| **Microsoft IIS 8.0** | application_server | 8.0 | 🔴 EOL | 2023-10-10 | IIS 8.0 is bundled with Windows Server 2012 and reached End of Life on October 10, 2023, together with Windows Server 2012. |
| **SQL Server 2019** | database | 2019 | 🟡 OUTDATED | 2030-01-08 | SQL Server 2019 mainstream support ended January 7, 2025. It remains in Extended Support through January 8, 2030, but is no longer the current version (SQL Server 2022 is latest). |

## Complexity Assessment

**Complexity Score: 7/10 — 🟠 High**

### Complexity Factors

| Factor | Value |
|--------|-------|
| Business Criticality | High |
| Architecture Tier | 2-Tier |
| Containerization | True |
| Ci Cd | True |
| Server Count | 2 |
| External Interface Count | 6 |
| Eol Components | True |
| Outdated Components | True |
| Technology Age | aging |

### Reasoning

HRApp-004 scores High complexity (7/10). The application has modernization-positive traits: containerized, CI/CD present, and already partially deployed on AWS. However, it runs on EOL Windows Server 2012 (EOL October 2023) with EOL IIS 8.0, presenting immediate security risk. The .NET Core version is unspecified, introducing lifecycle uncertainty. As a Custom made, High criticality application handling Payroll Processing and HR data (Internal classification), modernization carries significant business risk. The hybrid AWS+On-premise deployment adds architectural complexity requiring careful migration planning. SQL Server 2019 is in extended support (mainstream ended Jan 2025), and the licensed database engine adds cost and migration considerations. With 6 external interfaces and 12 API endpoints, integration re-testing will be substantial.

## Scenario Applicability

| Scenario | Status | Priority | Effort | Summary |
|----------|--------|----------|--------|---------|
| **Operating System Update** | 🔵 APPLICABLE | Critical | Medium | Windows Server 2012 is EOL (October 10, 2023). The OS is no longer receiving security patches, repre... |
| **Switch to Standard Linux Operating System** | ⚪ NOT_APPLICABLE | N/A | N/A | HRApp-004 uses IIS 8.0 and .NET Core on Windows Server, which are Windows-specific components. While... |
| **Switch to ARM-based CPU** | 🔴 BLOCKED | Low | Very High | Windows Server 2012 does not support ARM architecture. IIS 8.0 is also Windows x86/x64 only. ARM mig... |
| **Application Server Replacement** | 🔵 APPLICABLE | Critical | Medium | Microsoft IIS 8.0 is EOL (tied to Windows Server 2012 EOL, October 2023). Upgrade to IIS 10.0 (Windo... |
| **Application Migration to Cloud Infrastructure (Lift & Shift)** | 🟡 PARTIALLY_FULFILLED | High | High | HRApp-004 has a hybrid AWS+On-premise deployment. Some components are already in AWS, but on-premise... |
| **Application Containerization** | ✅ FULFILLED | N/A | N/A | HRApp-004 is already containerized. Container deployment is implemented. |
| **Application Refactoring and De-coupling** | 🔵 APPLICABLE | Medium | Very High | HRApp-004 is a Custom made 2-Tier application. Refactoring to a microservices or API-first architect... |
| **Upgrade Legacy Databases** | 🔵 APPLICABLE | Medium | Medium | SQL Server 2019 mainstream support ended January 7, 2025. While extended support continues until 203... |
| **Switch DB Engine to Open-Source Database Solution** | 🔵 APPLICABLE | Medium | High | SQL Server 2019 requires a commercial license (database_license_required: true). Migrating to Postgr... |
| **Update Outdated Components** | 🔵 APPLICABLE | High | High | Windows Server 2012 (EOL), IIS 8.0 (EOL), and SQL Server 2019 (outdated) all require updates. The .N... |

### Scenario Details

#### 🔵 Operating System Update
- **Status:** APPLICABLE
- **Priority:** Critical
- **Effort:** Medium
- **Reasoning:** Windows Server 2012 is EOL (October 10, 2023). The OS is no longer receiving security patches, representing a critical compliance and security risk for an HR application handling Payroll data.
- **Suggestion:** Upgrade to Windows Server 2022 (current LTS) or migrate to Windows Server 2019. For containerized components, evaluate migrating to Linux-based containers to reduce Windows licensing costs.

#### ⚪ Switch to Standard Linux Operating System
- **Status:** NOT_APPLICABLE
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** HRApp-004 uses IIS 8.0 and .NET Core on Windows Server, which are Windows-specific components. While .NET Core is cross-platform, IIS is Windows-only. Switching to Linux would require replacing IIS with a Linux-compatible web server.
- **Suggestion:** Not directly applicable without replacing IIS. Consider migrating to Linux after replacing IIS with Nginx or Kestrel as part of a broader modernization initiative.

#### 🔴 Switch to ARM-based CPU
- **Status:** BLOCKED
- **Priority:** Low
- **Effort:** Very High
- **Reasoning:** Windows Server 2012 does not support ARM architecture. IIS 8.0 is also Windows x86/x64 only. ARM migration requires OS modernization and application server replacement first.
- **Suggestion:** ARM migration is blocked until Windows Server is upgraded to a version with ARM support and IIS is replaced with an ARM-compatible web server.

#### 🔵 Application Server Replacement
- **Status:** APPLICABLE
- **Priority:** Critical
- **Effort:** Medium
- **Reasoning:** Microsoft IIS 8.0 is EOL (tied to Windows Server 2012 EOL, October 2023). Upgrade to IIS 10.0 (Windows Server 2019/2022) or replace with a cross-platform alternative (Kestrel, Nginx).
- **Suggestion:** Upgrade IIS to version 10.0 as part of the Windows Server 2022 migration. Alternatively, evaluate Kestrel (built-in .NET web server) or Nginx for cross-platform deployment.

#### 🟡 Application Migration to Cloud Infrastructure (Lift & Shift)
- **Status:** PARTIALLY_FULFILLED
- **Priority:** High
- **Effort:** High
- **Reasoning:** HRApp-004 has a hybrid AWS+On-premise deployment. Some components are already in AWS, but on-premise components remain. Full cloud migration is not yet complete.
- **Suggestion:** Complete the cloud migration by moving remaining on-premise components to AWS. Evaluate AWS ECS or EKS for the containerized components. Ensure data residency requirements for HR/Payroll data are met.

#### ✅ Application Containerization
- **Status:** FULFILLED
- **Priority:** N/A
- **Effort:** N/A
- **Reasoning:** HRApp-004 is already containerized. Container deployment is implemented.
- **Suggestion:** Containerization is fulfilled. Update container base images to resolve Windows Server 2012 EOL risk and improve container security posture.

#### 🔵 Application Refactoring and De-coupling
- **Status:** APPLICABLE
- **Priority:** Medium
- **Effort:** Very High
- **Reasoning:** HRApp-004 is a Custom made 2-Tier application. Refactoring to a microservices or API-first architecture would enable independent scaling of HR modules (Payroll, Benefits, Employee Management) and facilitate future cloud-native migration.
- **Suggestion:** Decompose the 2-Tier .NET Core application into domain-aligned microservices. Expose payroll, benefits, and employee management as independent services with REST/GraphQL APIs.

#### 🔵 Upgrade Legacy Databases
- **Status:** APPLICABLE
- **Priority:** Medium
- **Effort:** Medium
- **Reasoning:** SQL Server 2019 mainstream support ended January 7, 2025. While extended support continues until 2030, upgrading to SQL Server 2022 (current) would restore full feature and security patch coverage.
- **Suggestion:** Upgrade from SQL Server 2019 to SQL Server 2022. Use the Database Experimentation Assistant (DEA) to identify compatibility issues before migration.

#### 🔵 Switch DB Engine to Open-Source Database Solution
- **Status:** APPLICABLE
- **Priority:** Medium
- **Effort:** High
- **Reasoning:** SQL Server 2019 requires a commercial license (database_license_required: true). Migrating to PostgreSQL would eliminate SQL Server licensing costs. .NET Core has strong PostgreSQL support via Npgsql.
- **Suggestion:** Evaluate migration from SQL Server 2019 to PostgreSQL 16+. Use AWS Schema Conversion Tool or Babelfish for PostgreSQL to assess T-SQL compatibility. Test via the existing CI/CD pipeline.

#### 🔵 Update Outdated Components
- **Status:** APPLICABLE
- **Priority:** High
- **Effort:** High
- **Reasoning:** Windows Server 2012 (EOL), IIS 8.0 (EOL), and SQL Server 2019 (outdated) all require updates. The .NET Core version is unknown and must be identified and updated. CI/CD pipeline and existing containerization support systematic updates.
- **Suggestion:** Identify and document the exact .NET Core version in use. Create a component update roadmap: (1) Upgrade to Windows Server 2022 + IIS 10, (2) Identify and update .NET Core to 8.0 LTS, (3) Upgrade SQL Server 2019 to 2022.

## Business Case

**Total Adjusted Investment: $421,400.00**  
**Total Yearly Savings: $190,500.00**  
**Estimated ROI Payback: 2.21 years**

| Scenario | Status | Adj. Cost | Yearly Savings | ROI (years) |
|----------|--------|-----------|----------------|-------------|
| Operating System Update | APPLICABLE | $1,400.00 | $500.00 | 2.8 |
| Application Server Replacement | APPLICABLE | $14,000.00 | $12,000.00 | 1.17 |
| Application Migration to Cloud Infrastructure (Lift & Shift) | PARTIALLY_FULFILLED | $7,000.00 | $3,000.00 | 2.33 |
| Application Refactoring and De-coupling | APPLICABLE | $350,000.00 | $150,000.00 | 2.33 |
| Upgrade Legacy Databases | APPLICABLE | $14,000.00 | $10,000.00 | 1.4 |
| Switch DB Engine to Open-Source Database Solution | APPLICABLE | $35,000.00 | $15,000.00 | 2.33 |

> *Costs adjusted by complexity multiplier: **1.4x** (complexity score / 5)*
