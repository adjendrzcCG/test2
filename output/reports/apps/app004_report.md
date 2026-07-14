# HRApp-004 — Modernization Assessment Report

**Date:** 2026-07-14 | **App ID:** app004 | **Business Unit:** HR | **Status:** Production

---

## Executive Summary

HRApp-004 is a custom-built HR management system handling payroll, benefits, and employee records.  
It serves 670 HR users across a hybrid AWS + On-Premise deployment.  
High criticality and sensitive HR/payroll data combined with EOL infrastructure create significant risk.  
Complexity score: **7/10 (High)**.

| Attribute | Value |
|---|---|
| Criticality | High |
| Deployment | AWS + On-Premise (Hybrid) |
| Architecture | 2-Tier |
| Containerized | Yes ✅ |
| CI/CD | Yes |
| Users | 670 |
| External Interfaces | 6 |
| Complexity Score | **7 / 10 (High)** |

---

## Technology Assessment

| Component | Technology | Status | Notes |
|---|---|---|---|
| Operating System | Windows Server 2012 | ⛔ EOL | EOL October 2023 |
| Language | .NET Core | ⚠️ OUTDATED | Version ambiguous; may be EOL |
| App Server | Microsoft IIS 8.0 | ⛔ EOL | EOL October 2023 |
| Database | SQL Server 2019 | ⚠️ OUTDATED | Extended support until Jan 2030 |

**Overall Tech Risk: HIGH** — 2 EOL, 2 outdated components. Payroll data sensitivity amplifies risk.

---

## Complexity Analysis (Score: 7/10 — High)

- Windows Server 2012 + IIS 8.0 both EOL — no security patches for HR/payroll infrastructure
- .NET Core version unconfirmed — may be EOL (3.1) or current (6+)
- Hybrid deployment complicates migration planning
- SQL Server licensing cost in extended support
- 6 external interfaces (payroll providers, benefits systems)
- Custom-built application — full control but full responsibility

---

## Scenario Recommendations

| Scenario | Status | Priority | Effort |
|---|---|---|---|
| OS Update / Security Patch | 🟡 APPLICABLE | Critical | Medium |
| Switch to Standard Linux OS | 🟡 APPLICABLE | Medium | High |
| Switch to ARM CPU | 🟡 APPLICABLE | Medium | Medium |
| Application Server Replacement | 🟡 APPLICABLE | Critical | Medium |
| App Deployment to Cloud | 🟡 PARTIALLY FULFILLED | High | Medium |
| App Containerization | ✅ FULFILLED | — | Already containerized |
| App Refactor / Decoupling | 🟡 APPLICABLE | Medium | High |
| Upgrade Legacy Databases | 🟡 APPLICABLE | Medium | Medium |
| Switch to Managed DB | 🟡 APPLICABLE | Medium | Medium |
| Switch DB Engine to PostgreSQL | 🟡 APPLICABLE | Low | High |
| Managed ARM DB | ⚪ N/A | — | SQL Server not on ARM |
| Serverless DB | ⚪ N/A | — | Not suitable for payroll |

---

## Business Case

| Metric | Value |
|---|---|
| Complexity Multiplier | 1.40× |
| Total Migration Cost | $327,260 |
| Annual Savings | $197,000 |
| 3-Year Savings | $591,000 |
| 3-Year ROI | 80.6% |
| Recommended Priority | **Immediate — OS and IIS EOL critical** |

**Key Recommendation:** Immediately address Windows Server 2012 EOL — HR/payroll data on unpatched OS is a compliance violation.  
Confirm .NET Core version. Migrate to full cloud deployment to eliminate hybrid complexity.
