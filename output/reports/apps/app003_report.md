# AnalyticsApp-003 — Modernization Assessment Report

**Date:** 2026-07-14 | **App ID:** app003 | **Business Unit:** IT | **Status:** Production

---

## Executive Summary

AnalyticsApp-003 is an open-source analytics platform deployed on AWS and already containerized.  
It serves 480 IT users for business intelligence and reporting.  
Despite having the most EOL components, it is the **most straightforward to modernize** (score: **5/10**) due to its containerized, open-source stack and low criticality.

| Attribute | Value |
|---|---|
| Criticality | Low |
| Deployment | AWS |
| Architecture | 3-Tier |
| Containerized | **Yes** ✅ |
| CI/CD | Yes |
| Users | 480 |
| External Interfaces | 3 |
| Complexity Score | **5 / 10 (Medium)** |

---

## Technology Assessment

| Component | Technology | Status | Notes |
|---|---|---|---|
| Operating System | RHEL 7 | ⛔ EOL | EOL June 2024 |
| Language | Python 3.9 | ⛔ EOL | EOL October 2025 |
| App Server | Apache Tomcat 6.1 | ⛔ EOL | EOL December 2016 |
| Database | PostgreSQL 13 | ⛔ EOL | EOL November 2025 |

**Overall Tech Risk: HIGH** — All 4 components are EOL.

---

## Complexity Analysis (Score: 5/10 — Medium)

- All 4 technology components at EOL
- **BUT:** Already containerized — major modernization accelerator
- Open-source stack — well-documented upgrade paths
- Single environment — simpler migration window
- Low business criticality reduces risk tolerance
- Only 3 external interfaces

---

## Scenario Recommendations

| Scenario | Status | Priority | Effort |
|---|---|---|---|
| OS Update / Security Patch | 🟡 APPLICABLE | Critical | Low |
| Switch to Standard Linux OS | ✅ FULFILLED | — | Already Linux |
| Switch to ARM CPU | 🟡 APPLICABLE | Medium | Low |
| Application Server Replacement | 🟡 APPLICABLE | Critical | Low |
| App Deployment to Cloud | ✅ FULFILLED | — | Already AWS |
| App Containerization | ✅ FULFILLED | — | Already containerized |
| App Refactor / Decoupling | 🟡 PARTIALLY FULFILLED | Low | Medium |
| Upgrade Legacy Databases | 🟡 APPLICABLE | Critical | Low |
| Switch to Managed DB | 🟡 APPLICABLE | Medium | Low |
| Managed ARM DB | 🟡 APPLICABLE | Low | Low |
| Serverless DB Migration | 🟡 APPLICABLE | Medium | Low |
| Switch DB Engine to PostgreSQL | ✅ FULFILLED | — | Already PostgreSQL |

---

## Business Case

| Metric | Value |
|---|---|
| Complexity Multiplier | 1.00× |
| Total Migration Cost | $386,000 |
| Annual Savings | $188,000 |
| 3-Year Savings | $564,000 |
| 3-Year ROI | 46.1% |
| Recommended Priority | **Short-term — quick wins available** |

**Key Recommendation:** Upgrade all EOL components in sequence: OS → Python → Tomcat → PostgreSQL.  
Each upgrade is low-effort given containerized, open-source foundation. Best ROI per effort in portfolio.
