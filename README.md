# SCI: Simplified Compliance Infrastructure  [![Go Reference](https://pkg.go.dev/badge/github.com/revanite-io/sci.svg)](https://pkg.go.dev/github.com/revanite-io/sci)

- [Overview](#overview)
- [The Model](#the-model)
  - [Layer 1: Guidance](#layer-1-guidance)
  - [Layer 2: Controls](#layer-2-controls)
    - [Layer 2 Schema](#layer-2-schema)
  - [Layer 3: Policy](#layer-3-policy)
  - [Layer 4: Evaluation](#layer-4-evaluation)
    - [Layer 4 Schema](#layer-4-schema)
  - [Layer 5: Enforcement](#layer-5-enforcement)
  - [Layer 6: Audit](#layer-6-audit)
- [Usage](#usage)
- [Contributing](#contributing)
- [Projects and tooling using SCI](#projects-and-tooling-using-sci)

## Overview

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

Examples include [CIS Benchmarks](https://www.cisecurity.org/cis-benchmarks-overview), [FINOS Common Cloud Controls](https://github.com/finos/common-cloud-controls/blob/main/README.md), and the [Open Source Project Security (OSPS) Baseline](https://baseline.openssf.org/).

#### Layer 2 Schema

The SCI [Layer 2 Schema](./schemas/layer-2.cue) describes the machine-readable format of Layer 2 controls.

The schema allows controls to be mapped to threats or Layer 1 controls by their unique identifiers. Threats may also be expressed in the schema, with mappings to the technology-specific capabilities which may be vulnerable to the threat.

The SCI go module provides Layer 2 support for ingesting YAML and JSON documents that follow this schema.

The [cue](https://cuelang.org) CLI can be used to [validate YAML data](https://cuelang.org/docs/concept/how-cue-works-with-yaml/#validating-yaml-files-against-a-schema) containing a Layer 2 control catalog.

### Layer 3: Policy

Activities in the Policy layer provide risk-informed governance rules that — while based on best practices and industry standards — are tailored to an organization.

Layer 3 controls are typically developed by an organization to compile into organizational policies. Policies cannot be properly developed without consideration for organization-specific risk appetite and risk-acceptance.

These policy documents may be referenced by other policy documents, or used as a starting point for Layer 4 assessments.

### Layer 4: Evaluation

Activities in the Evaluation layer provide inspection of code, configurations, and deployments. Those elements are part of the _software development lifecycle_ which is not represented in this model.

Evaluation activities may be built based on outputs from layers 2 or 3. While automated assessments are often developed by vendors or industry groups, robust evaluation should be informed by organizational policies in order to custom-tailor the assessment to the needs of the compliance program.

#### Layer 4 Schema

The SCI [Layer 4 Schema](./schemas/layer-4.cue) describes the machine-readable format of Layer 4 evaluation results.

The schema allows evaluations to be mapped to Layer 2 controls by their unique identifiers.

The SCI go module provides Layer 2 support for writing assessments that can write results in this schema. The work to add go module Layer 4 support for reading and writing assessments data is planned in [issue 30](https://github.com/revanite-io/sci/issues/30), but not completed.

### Layer 5: Enforcement

Activities in the Enforcement layer provide prevention or remediation. These enforcement actions should be guided by Layer 3 policies and based on assessment findings from Layer 4 evaluations.

This layer ensures that the organization is complying with policy when evidence of noncompliance is found, such as by blocking the deployment of a resource that does not meet the organization's policies.

### Layer 6: Audit

Activities in the Audit layer provide a review of organizational policy and conformance.

Audits consider information from all of the lower layers. These activities are typically performed by internal or external auditors to ensure that the organization has designed and enforced effective policies based on the organization's requirements.

## Usage

Install the go module with `go get github.com/revanite-io/sci` and consult our [go docs](https://pkg.go.dev/github.com/revanite-io/sci)

Use the schemas directly with [cue](https://cuelang.org/) for validating SCI data payloads against the schemas and more.

## Contributing

We're so glad you asked - see [CONTRIBUTING.md](/CONTRIBUTING.md) and if you have any questions or feedback head over to the OpenSSF Slack in [#wg-orbit](https://openssf.slack.com/archives/C08NJTFAL74)

## Projects and tooling using SCI

The SCI Model is intended to be used in conjunction with tooling that can help automate the compliance process.

Examples of such tools include:

* [revanite-io/pvtr-github-repo](https://github.com/revanite-io/pvtr-github-repo) - this tool automates layer 4 evaluation of the [Open Source Project Security Baseline](https://baseline.openssf.org/) control catalog
* [FINOS Common Cloud Controls](https://www.finos.org/common-cloud-controls-project) - the control catalog is modeled using the layer 2 schema and consumed using the layer2 go package.
