# ERPApp-001 — Modernization Assessment Report

**Date:** 2026-07-14 | **App ID:** app001 | **Business Unit:** Finance | **Status:** Production

---

## Executive Summary

ERPApp-001 is a core financial ERP system built on COBOL-2014 running on IBM AIX 7.2.  
It handles general ledger, financial transactions, and regulatory reporting for 350 users.  
This application presents the **highest modernization complexity** (score: **8/10**) in the portfolio due to its legacy COBOL monolithic architecture on an EOL proprietary platform.

| Attribute | Value |
|---|---|
| Criticality | High |
| Deployment | On-Premise |
| Architecture | 1-Tier (Monolith) |
| Containerized | No |
| CI/CD | No |
| Users | 350 |
| External Interfaces | 5 |
| Complexity Score | **8 / 10 (Very High)** |

---

## Technology Assessment

| Component | Technology | Status | Notes |
|---|---|---|---|
| Operating System | AIX 7.2 | ⛔ EOL | IBM extended support ended ~April 2025 |
| Language | COBOL-2014 | ❓ NO_KNOWLEDGE | Stable standard; runtime vendor support varies |
| Database | Oracle 19c | ✅ CURRENT | Supported until April 2027 |
| App Server | None | N/A | 1-Tier — no app server |

**Overall Tech Risk: CRITICAL** — 2 components at EOL/unknown, database approaching EOL in 2027.

---

## Complexity Analysis (Score: 8/10 — Very High)

- COBOL codebase requires specialist talent (shrinking pool)
- AIX 7.2 platform is EOL with no security patch path
- 1-Tier tightly couples presentation, logic, and data
- No CI/CD pipeline — manual deployment risk
- 5 external interfaces require adapter rewrites
- Oracle 19c license cost with impending EOL (2027)

---

## Scenario Recommendations

| Scenario | Status | Priority | Effort |
|---|---|---|---|
| OS Update / Security Patch | 🟡 APPLICABLE | Critical | High |
| Switch to Standard Linux OS | 🟡 APPLICABLE | High | High |
| App Deployment to Cloud | 🟡 APPLICABLE | Medium | High |
| App Refactor / Decoupling | 🟡 APPLICABLE | High | Very High |
| Switch to Managed DB | 🟡 APPLICABLE | Medium | Medium |
| Switch DB Engine to PostgreSQL | 🟡 APPLICABLE | Low | Very High |
| App Containerization | 🔴 BLOCKED | — | Blocked by legacy stack |
| Application Server Replacement | ⚪ N/A | — | No app server |

---

## Business Case

| Metric | Value |
|---|---|
| Complexity Multiplier | 1.60× |
| Total Migration Cost | $479,200 |
| Annual Savings | $165,500 |
| 3-Year Savings | $496,500 |
| 3-Year ROI | 3.6% |
| Recommended Priority | After infrastructure remediation |

**Key Recommendation:** Address AIX → Linux OS migration as immediate priority.  
Long-term: plan COBOL modernization via strangler-fig pattern over 3-5 years.
