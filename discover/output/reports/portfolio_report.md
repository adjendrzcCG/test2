# Portfolio Modernization Report

> **Generated:** 2026-07-13 | **Reference Date:** 2026-07-13 | **Portfolio:** Enterprise Application Portfolio

## Executive Summary

This portfolio modernization report covers **5 applications** assessed for technology lifecycle status, complexity, modernization opportunities, and business value.

| Metric | Value |
|--------|-------|
| Total Applications | 5 |
| Out-of-Scope (Retired) | 1 |
| In-Scope (Production) | 4 |
| Applications with EOL Components | 3 |
| Applications with Outdated Components | 3 |
| Applications with Unknown Version Data | 1 |
| Total Portfolio Investment Required | $963,940.00 |
| Total Yearly Savings | $395,400.00 |
| Portfolio ROI Payback | ~2.44 years |

The portfolio carries **significant technical debt** concentrated in EOL operating systems (RHEL 7, Windows Server 2012), severely EOL middleware (WebSphere 7.0, Apache Tomcat 6.1), and EOL languages/runtimes (Python 3.9, PostgreSQL 13). Three of four in-scope applications have EOL components requiring immediate attention. ERPApp-001 and HRApp-004 represent the highest modernization risk and investment due to their High business criticality and legacy architecture.

## Portfolio Overview

| App ID | Name | Criticality | Status | Deployment | Containerized | CI/CD | Complexity |
|--------|------|-------------|--------|------------|---------------|-------|------------|
| `app001` | **ERPApp-001** | High | Production | On-Premise | No | No | 9/10 (Very High) |
| `app002` | **CRMApp-002** | Medium | Production | AWS | No | Yes | 7/10 (High) |
| `app003` | **AnalyticsApp-003** | Low | Production | AWS | Yes | Yes | 4/10 (Medium) |
| `app004` | **HRApp-004** | High | Production | AWS+On-premise | Yes | Yes | 7/10 (High) |
| `app005` | **EComApp-005** *(Out of Scope)* | Critical | Retired | AWS | Yes | Yes | N/A (Out of Scope) |

## Technology Health

### EOL & Outdated Component Summary

| App | OS | Language | App Server | Database | EOL? | Outdated? | Missing Version? |
|-----|----|----------|------------|----------|------|-----------|-----------------|
| **ERPApp-001** | AIX 7.2 — 🟡 Outdated | COBOL-2014 — ✅ Current | N/A | Oracle 19c — ✅ Current | ✅ No | 🟡 Yes | ✅ No |
| **CRMApp-002** | RHEL 7 — 🔴 EOL | Java 11 — 🟡 Outdated | Websphere 7.0 — 🔴 EOL | Amazon RDS MySQL — ✅ Current | 🔴 Yes | 🟡 Yes | ✅ No |
| **AnalyticsApp-003** | RHEL 7 — 🔴 EOL | Python 3.9 — 🔴 EOL | Apache Tomcat 6.1 — 🔴 EOL | PostgreSQL 13 — 🔴 EOL | 🔴 Yes | ✅ No | ✅ No |
| **HRApp-004** | Windows Server 2012 — 🔴 EOL | .NET Core — ❓ Unknown | Microsoft IIS 8.0 — 🔴 EOL | SQL Server 2019 — 🟡 Outdated | 🔴 Yes | 🟡 Yes | ❓ Yes |

### Key Technology Risks

- 🔴 **RHEL 7** (EOL June 2024) — affects CRMApp-002 and AnalyticsApp-003
- 🔴 **Windows Server 2012** (EOL October 2023) — affects HRApp-004
- 🔴 **WebSphere 7.0** (EOL September 2015) — affects CRMApp-002; critically overdue
- 🔴 **Apache Tomcat 6.1** (EOL December 2016) — affects AnalyticsApp-003; critically overdue
- 🔴 **Python 3.9** (EOL October 2025) — affects AnalyticsApp-003
- 🔴 **PostgreSQL 13** (EOL November 2025) — affects AnalyticsApp-003
- 🔴 **Microsoft IIS 8.0** (EOL October 2023) — affects HRApp-004
- 🟡 **AIX 7.2** (Outdated) — affects ERPApp-001
- 🟡 **Java 11** (Outdated/approaching EOL) — affects CRMApp-002
- 🟡 **SQL Server 2019** (Outdated, mainstream support ended Jan 2025) — affects HRApp-004
- ❓ **.NET Core** (Version unknown) — affects HRApp-004; version discovery required

## Top Modernization Opportunities

Based on ROI, urgency, and feasibility:

| Rank | Scenario | Application(s) | Status | Adj. Cost | Yearly Savings | ROI |
|------|----------|----------------|--------|-----------|----------------|-----|
| 1 | Application Server Replacement (Tomcat 6.1 → 10.x) | AnalyticsApp-003 | 🔵 APPLICABLE | $8,000 | $12,000 | 0.67 yrs |
| 2 | Upgrade Legacy Databases (PostgreSQL 13 → 16+) | AnalyticsApp-003 | 🔵 APPLICABLE | $8,000 | $10,000 | 0.80 yrs |
| 3 | Switch to Standard Linux OS | ERPApp-001 | 🔵 APPLICABLE | $540 | $400 | 1.35 yrs |
| 4 | Application Server Replacement (WebSphere 7.0) | CRMApp-002 | 🔵 APPLICABLE | $14,000 | $12,000 | 1.17 yrs |
| 5 | OS Security Patch (all in-scope apps) | All 4 apps | 🔵 APPLICABLE | ~$5,400 | ~$2,000 | 2.7 yrs |
| 6 | Full Cloud Migration (HRApp-004) | HRApp-004 | 🟡 PARTIALLY_FULFILLED | $7,000 | $3,000 | 2.33 yrs |
| 7 | Switch DB Engine to Open-Source | ERPApp-001 | 🔵 APPLICABLE | $45,000 | $15,000 | 3.0 yrs |
| 8 | Application Refactoring (ERPApp-001 COBOL) | ERPApp-001 | 🔵 APPLICABLE | $450,000 | $150,000 | 3.0 yrs |

## Financial Summary

| Metric | Value |
|--------|-------|
| Total Portfolio Investment | **$963,940.00** |
| Total Yearly Savings | **$395,400.00** |
| Portfolio ROI Payback | **2.44 years** |

### Per-Application Financial Summary

| Application | Complexity | Investment | Yearly Savings | ROI (years) |
|-------------|------------|------------|----------------|-------------|
| **ERPApp-001** (`app001`) | 1.8x | $506,340.00 | $168,900.00 | 3.0 |
| **CRMApp-002** (`app002`) | 1.4x | $15,400.00 | $12,500.00 | 1.23 |
| **AnalyticsApp-003** (`app003`) | 0.8x | $20,800.00 | $23,500.00 | 0.89 |
| **HRApp-004** (`app004`) | 1.4x | $421,400.00 | $190,500.00 | 2.21 |

> *Investment costs are base migration costs adjusted by per-application complexity multiplier (complexity score ÷ 5).*
> *Only APPLICABLE and PARTIALLY_FULFILLED scenarios are included in cost calculations.*
> *Savings are first-year estimates; multi-year compound savings will increase ROI.*

## Risk Applications

### 🔴 Critical Risk: ERPApp-001
- **Risk:** COBOL monolith on EOL AIX 7.2, no CI/CD, no containerization, Confidential data, decommission target 2027
- **Action:** Immediate COBOL code analysis, AIX upgrade, cloud migration planning, Oracle license cost reduction
- **Impact:** Finance/Accounting system failure would be catastrophic to business operations

### 🔴 Critical Risk: CRMApp-002
- **Risk:** WebSphere 7.0 (EOL since 2015) + RHEL 7 (EOL 2024) combination creates severe unpatched vulnerability exposure
- **Action:** Emergency WebSphere 7.0 replacement, RHEL 7 → RHEL 9 upgrade
- **Impact:** 1,200 Marketing users affected; Customer data exposure risk

### 🟠 High Risk: AnalyticsApp-003
- **Risk:** All four technology components are EOL; however, application is already containerized with CI/CD
- **Action:** Systematic component update plan — OS, Python, Tomcat, PostgreSQL all need upgrading
- **Mitigation:** Low complexity (4/10) and existing containerization make this highly actionable

### 🟠 High Risk: HRApp-004
- **Risk:** EOL Windows Server 2012 + IIS 8.0, unknown .NET Core version, HR/Payroll data sensitivity
- **Action:** OS upgrade to Windows Server 2022, identify .NET Core version, complete cloud migration
- **Impact:** Payroll processing disruption would have immediate employee impact

## Per-Application Reports

Detailed reports for each in-scope application:

- [ERPApp-001 (app001)](apps/app_report_app001.md)
- [CRMApp-002 (app002)](apps/app_report_app002.md)
- [AnalyticsApp-003 (app003)](apps/app_report_app003.md)
- [HRApp-004 (app004)](apps/app_report_app004.md)

---

*app005 (EComApp-005) is excluded from this report as its status is **Retired**.*
