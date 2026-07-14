# CRMApp-002 — Modernization Assessment Report

**Date:** 2026-07-14 | **App ID:** app002 | **Business Unit:** Marketing | **Status:** Production

---

## Executive Summary

CRMApp-002 is a 3rd-party CRM system deployed on AWS serving 1,200 marketing users.  
Despite cloud deployment, the application stack includes critically EOL components —  
notably WebSphere 7.0 (EOL 2015) — creating significant security exposure.  
Complexity score: **6/10 (Medium-High)**.

| Attribute | Value |
|---|---|
| Criticality | Medium |
| Deployment | AWS |
| Architecture | Unknown |
| Containerized | No |
| CI/CD | Yes |
| Users | 1,200 |
| External Interfaces | 8 |
| Complexity Score | **6 / 10 (Medium-High)** |

---

## Technology Assessment

| Component | Technology | Status | Notes |
|---|---|---|---|
| Operating System | RHEL 7 | ⛔ EOL | EOL June 2024 |
| Language | Java 11 | ⚠️ OUTDATED | Premier support ended Sep 2023 |
| App Server | Websphere 7.0 | ⛔ EOL | EOL September 2015 — critical |
| Database | Amazon RDS MySQL | ✅ CURRENT | Managed service |

**Overall Tech Risk: HIGH** — 3 components at EOL/outdated.

---

## Complexity Analysis (Score: 6/10 — Medium-High)

- WebSphere 7.0 EOL since 2015 — over 10 years without vendor support
- RHEL 7 EOL creates OS-level security gap
- 3rd-party software limits customization and migration paths
- 8 external interfaces require careful re-integration
- Java 11 needs upgrade to Java 21 LTS

---

## Scenario Recommendations

| Scenario | Status | Priority | Effort |
|---|---|---|---|
| OS Update / Security Patch | 🟡 APPLICABLE | Critical | Medium |
| Switch to Standard Linux OS | ✅ FULFILLED | — | Already Linux |
| Switch to ARM CPU | 🟡 APPLICABLE | Medium | Medium |
| Application Server Replacement | 🟡 APPLICABLE | Critical | High |
| App Deployment to Cloud | ✅ FULFILLED | — | Already on AWS |
| App Containerization | 🟡 APPLICABLE | Medium | Medium |
| Switch to Managed DB | ✅ FULFILLED | — | Already RDS |
| Managed ARM DB | 🟡 APPLICABLE | Low | Low |
| Serverless DB Migration | 🟡 APPLICABLE | Low | Medium |

---

## Business Case

| Metric | Value |
|---|---|
| Complexity Multiplier | 1.20× |
| Total Migration Cost | $143,640 |
| Annual Savings | $117,500 |
| 3-Year Savings | $352,500 |
| 3-Year ROI | 145.4% |
| Recommended Priority | **Immediate — WebSphere replacement** |

**Key Recommendation:** Replace WebSphere 7.0 immediately (critical security risk).  
Then upgrade RHEL 7 → RHEL 9 and Java 11 → Java 21. Containerize for better portability.
