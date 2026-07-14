# Portfolio Modernization Report

## Executive Summary

Four in-scope applications require modernization attention, with three carrying EOL components and three carrying outdated components. The portfolio can return its investment in about 2.44 years, but value depends on sequencing quick-win risk reduction ahead of the two large transformation candidates. The strongest near-term moves are OS remediation everywhere, app server replacement where middleware is unsupported, and tightly governed selective transformation for ERP and HR.

| Metric | Value |
| --- | --- |
| Generated timestamp | 2026-07-14T07:33:55Z |
| Analysis ID | GD-2026-001 |
| Applications in file | 5 |
| In scope | 4 |
| Out of scope | app005 — EComApp-005 (Retired) |
| Apps with EOL components | 3 |
| Apps with outdated components | 3 |
| Apps with missing version data | 1 |
| Total portfolio investment | $963,940 |
| Total yearly savings | $395,400 |
| ROI payback | 2.44 years |

## Portfolio Overview

| App | Name | BU | Criticality | Deploy | Containers | CI/CD | Complexity |
| --- | --- | --- | --- | --- | --- | --- | --- |
| app001 | ERPApp-001 | Finance | High | On-Premise | No | No | 9/10 Very High |
| app002 | CRMApp-002 | Marketing | Medium | AWS | No | Yes | 7/10 High |
| app003 | AnalyticsApp-003 | IT | Low | AWS | Yes | Yes | 4/10 Medium |
| app004 | HRApp-004 | HR | High | AWS+On-Prem | Yes | Yes | 7/10 High |
| app005 | EComApp-005 | Operations | Critical | AWS | Yes | Yes | N/A (Retired) |

## Technology Health

| App | Application | EOL Components | Outdated Components | Missing Version Data | Assessment Detail |
| --- | --- | --- | --- | --- | --- |
| app001 | ERPApp-001 | No | Yes | No | AIX 7.2 (OUTDATED), COBOL 2014 (CURRENT), Oracle 19c (CURRENT) |
| app002 | CRMApp-002 | Yes | Yes | No | RHEL 7 (EOL), Java 11 (OUTDATED), WebSphere 7.0 (EOL), Amazon RDS MySQL Managed (CURRENT) |
| app003 | AnalyticsApp-003 | Yes | No | No | RHEL 7 (EOL), Python 3.9 (EOL), Apache Tomcat 6.1 (EOL), PostgreSQL 13 (EOL) |
| app004 | HRApp-004 | Yes | Yes | Yes | Windows Server 2012 (EOL), .NET Core Unknown (NO_KNOWLEDGE), Microsoft IIS 8.0 (EOL), SQL Server 2019 (OUTDATED) |

### Key portfolio risks

- Unsupported operating systems and middleware are concentrated in CRM, Analytics, and HR, creating urgent supportability and security pressure.
- ERP and HR dominate total investment, so poor sequencing could consume budget before faster ROI work is realized.
- app004 has missing version data for .NET Core, which must be resolved early to protect scope, cost, and security assumptions.
- Several strategic options are blocked by current architecture, especially ARM moves and containerization for the most legacy workloads.

## Top Modernization Opportunities

| Rank | Scenario | Application | Adj. Cost | Yearly Savings | ROI (yrs) |
| --- | --- | --- | --- | --- | --- |
| 1 | App Server Replacement | AnalyticsApp-003 | $8,000 | $12,000/yr | 0.67 |
| 2 | Upgrade Legacy Databases | AnalyticsApp-003 | $8,000 | $10,000/yr | 0.8 |
| 3 | App Server Replacement | CRMApp-002 | $14,000 | $12,000/yr | 1.17 |
| 4 | App Server Replacement | HRApp-004 | $14,000 | $12,000/yr | 1.17 |
| 5 | Switch to Linux OS | ERPApp-001 | $540 | $400/yr | 1.35 |
| 6 | Upgrade Legacy Databases | HRApp-004 | $14,000 | $10,000/yr | 1.4 |

## Financial Summary

| Portfolio Metric | Value |
| --- | --- |
| Total Investment | $963,940 |
| Total Yearly Savings | $395,400 |
| Portfolio Payback | 2.44 years |

| App | Application | Investment | Yearly Savings | ROI (yrs) |
| --- | --- | --- | --- | --- |
| app001 | ERPApp-001 | $506,340 | $168,900 | 3 |
| app002 | CRMApp-002 | $15,400 | $12,500 | 1.23 |
| app003 | AnalyticsApp-003 | $20,800 | $23,500 | 0.89 |
| app004 | HRApp-004 | $421,400 | $190,500 | 2.21 |

## Risk Applications

### CRMApp-002 — Critical risk

CRMApp-002 already has cloud hosting and delivery automation, but it remains exposed through EOL RHEL 7 and WebSphere 7.0 plus an aging Java 11 runtime. The best value path is to remediate the middleware layer first, because that removes the highest support risk and unlocks follow-on modernization.

**Key risks**

- WebSphere 7.0 and RHEL 7 are both unsupported, creating urgent supportability and security exposure.
- The application has a wide integration footprint, so middleware failure or rushed migration would have broad downstream impact.
- Third-party ownership constrains refactoring and can make component upgrades dependent on vendor-supported paths.

**Recommended sequencing**

- Wave 1 for OS remediation, then Wave 2 for the WebSphere replacement and follow-on runtime updates.

### HRApp-004 — Critical risk

HRApp-004 has a usable modernization foundation, but its hosting stack is already out of support and key lifecycle data for .NET is missing. The right approach is to stabilize the platform first, close the information gap, and then use that foundation to finish the cloud journey and reduce licensing overhead.

**Key risks**

- Windows Server 2012 and IIS 8.0 are unsupported in a high-sensitivity HR environment.
- The unknown .NET Core version introduces planning uncertainty and could hide additional runtime remediation effort.
- Hybrid deployment increases operational complexity and slows standardization of controls, observability, and release procedures.

**Recommended sequencing**

- Wave 1 for discovery and platform hardening, Wave 2 for hosting/cloud completion, Wave 3 for refactoring and database economics.

### ERPApp-001 — High risk

ERPApp-001 is a mission-critical finance ERP with a still-supported core stack, but it remains anchored to outdated AIX infrastructure and a high-friction operating model with no CI/CD or containers. Because the application is scheduled for decommission in 2027, modernization should emphasize stabilization, supportability, and selective cost takeout rather than full transformation.

**Key risks**

- Outdated AIX 7.2 increases security, support, and platform continuity risk ahead of the planned retirement window.
- The 1-tier design plus no CI/CD/no containers makes even small changes operationally expensive and slower to validate.
- A 1000GB licensed Oracle footprint and five interfaces raise cutover complexity for any platform or data move.

**Recommended sequencing**

- Wave 1 for OS hardening and Linux replatform decision; defer deep transformation to replacement planning.

### AnalyticsApp-003 — High risk

AnalyticsApp-003 is low criticality and already operates with modern engineering practices, but nearly every runtime layer is out of support. It is the best candidate for a bundled remediation sprint because the technical debt is urgent while delivery complexity remains manageable.

**Key risks**

- The application has full-stack EOL exposure across OS, runtime, application server, and database.
- If the low criticality label suppresses urgency, the environment may remain vulnerable even though upgrades are comparatively easy.
- Independent remediation projects could duplicate effort; the stack should be updated as a coordinated package.

**Recommended sequencing**

- Wave 1 / early Wave 2 quick-win modernization bundle covering OS, Tomcat, Python, and PostgreSQL.

## Modernization Roadmap

### Wave 1: Quick Wins (0-6 months)

Reduce immediate support and security exposure while improving planning certainty.

- Execute OS/security remediation across app001-app004, prioritized by EOL severity.
- Complete .NET Core version discovery and hosting baseline assessment for app004.
- Launch app003 as the proving ground for coordinated stack refresh because it is low complexity and already automated.

### Wave 2: Strategic Modernization (6-18 months)

Remove unsupported middleware, finish cloud/platform alignment, and capture high-confidence ROI.

- Replace WebSphere for app002, Tomcat for app003, and IIS hosting for app004.
- Complete partial cloud migration for app004 and align operational tooling around the target platform.
- Bundle app003 runtime and database upgrades to eliminate the largest EOL concentration in the portfolio.

### Wave 3: Transformation (18-36 months)

Pursue selective structural change only where the business case survives timeline and support assumptions.

- Revisit ERP refactoring or database engine change only if the 2027 retirement horizon moves.
- Use app004 platform stabilization outcomes to decide on refactoring scope and database cost optimization.
- Evaluate ARM for app003 after the supported baseline is restored and performance benchmarks are available.

## Per-Application Reports

- [ERPApp-001](apps/app_report_app001.md)
- [CRMApp-002](apps/app_report_app002.md)
- [AnalyticsApp-003](apps/app_report_app003.md)
- [HRApp-004](apps/app_report_app004.md)
