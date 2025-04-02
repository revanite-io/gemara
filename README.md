# SCI: Simplified Compliance Infrastructure

The _Simplified Compliance Infrastructure_ provides a model to describe the categories of compliance activities, how they interact, and the schemas to enable automated interoperability between them.

In order to better facilitate cross-functional communication, the SCI Model seeks to outline the categorical layers of activities related to automated governance.

We will begin by establishing the overall model, and then the following sections will contain detailed breakdowns of each categorical layer, with examples.

This document assumes that the reader is trained in governance, risk, compliance, or cybersecurity— and therefore understands or can find definitions of concepts or terms that are infrequently used herein.

For the purpose of this document, "organization" may refer to a business or an organizational unit within it.

## The Model

Each layer in the model builds upon the lower layer, though in higher-level use cases you may find examples where multiple lower layers are brought into a higher level together. Examples and clarifications can be found in the respective sections below.

| Layer | Name | Description |
|-------|------|-------------|
| 1 | Guidance | High-level guidance on cybersecurity measures |
| 2 | Controls | Technology-specific, threat-informed security controls |
| 3 | Policy | Risk-informed guidance tailored to an organization |
| 4 | Evaluation | Inspection of code, configurations, and deployments |
| 5 | Enforcement | Prevention or remediation based on assessment findings |
| 6 | Audit | Review of organizational policy and conformance |

### Layer 1: Guidance

The Guidance layer is the lowest level of the SCI Model. Activities in this layer provide high-level rules pertaining to cybersecurity measures. Guidance is typically developed by industry groups, government agencies, or international standards bodies. They are intended to be used as a starting point for organizations to develop their own cybersecurity programs.

Guidance frameworks or standards occasionally express their rules using the term "controls" — these should be understood as Layer 1 Controls in the event that the term appears to conflict with Layer 2.

These guidance documents are high-level, abstract controls that may be referenced in the development of other Layer 1 or Layer 2 assets.

Examples include the NIST Cybersecurity Framework, ISO 27001, PCI DSS, HIPPA, GDPR, and CRA.

### Layer 2: Controls

Activities in the Control layer produce technology-specific, threat-informed security controls. Controls are the specific guardrails that organizations put in place to protect their information systems. They are typically informed by the best practices and industry standards which are produced in Layer 1.

Layer 2 controls are typically developed by an organization for its own purposes, or for general use by industry groups, government agencies, or international standards bodies.

Assets in this category may be refined into more specific Layer 2 controls, or combined with organizational risk considerations to form Layer 3 policies.

The recommended process for developing Layer 2 controls is to first assess the technology's capabilities, then identify threats to those capabilities, and finally develop controls to mitigate those threats.

Examples include CIS Benchmarks, FINOS Common Cloud Controls, and the OSPS Baseline.

#### Layer 2 Schema

The SCI [Layer 2 Schema](./schemas/layer-2.cue) provides methods for expressing Layer 2 controls in a machine-readable format.

The schema allows controls to be mapped to threats or Layer 1 controls by their unique identifiers. Threats may also be expressed in the schema, with mappings to the technology-specific capabilities which may be vulnerable to the threat.

The SCI go module provides Layer 2 support for ingesting documents that follow this schema.

### Layer 3: Policy

Activities in the Policy layer provide risk-informed governance rules tailored to an organization. They are typically informed by risks or based on best practices and industry standards.

Layer 3 controls are typically developed by an organization for its own purposes, to compile into organizational policies. Policies cannot be properly developed without consideration for organization-specific risk appetite and risk-acceptance.

These policy documents may be referenced by other policy documents, or used as a starting point for Layer 4 assessments.

### Layer 4: Evaluation

Activities in the Evaluation layer provide inspection of code, configurations, and deployments. Those elements are part of the _software development lifecycle_ which is not represented in this model.

Evaluation activities may be built based on outputs from layers 2 or 3. While automated assessments are often developed by vendors or industry groups, proper evaluation should also be able to ingest organizational policies to custom-tailor the assessment to the needs of the compliance program.

#### Layer 4 Schema

The SCI [Layer 4 Schema](./schemas/layer-4.cue) provides methods for expressing Layer 4 evaluation results in a machine readable format.

The schema allows evaluations to be mapped to Layer 2 controls by their unique identifiers.

The SCI go module provides Layer 2 support for writing assessments that can write results in this schema.

### Layer 5: Enforcement

Activities in the Enforcement layer provide prevention or remediation based on assessment findings.

Enforcement actions should be dictated by Layer 3 policies, and act on information gathered by Layer 4 evaluations.

This layer ensures that the organization is complying with policy when evidence of noncompliance is found, such as by blocking the deployment of a resource that does not meet the organization's policies.

### Layer 6: Audit

Activities in the Audit layer provide a review of organizational policy and conformance.

Audits consider information from all of the lower layers. These activities are typically performed by internal or external auditors to ensure that the organization has designed and enforced effective policies based on the organization's requirements.

## Tooling

The SCI Model is intended to be used in conjunction with tooling that can help automate the compliance process.

This section will be updated as SCI-compatible tooling becomes available for production use cases.
