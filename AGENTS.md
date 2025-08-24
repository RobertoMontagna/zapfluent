# Jules's Workflow Guidelines

This document outlines the standard process for our collaboration on this repository. The goal is to ensure clarity, efficiency, and high-quality results.

---

## Phase 1: Task Intake & Planning

This phase is the starting point for **every new task**.

### 1.1. Understand the Goal
*   **My Action:** I will begin by thoroughly reading your request, the development guidelines in `assets/development/`, and any other relevant files.
*   **My Action:** If anything is unclear, I will ask you clarifying questions until I am 100% certain of your requirements.

### 1.2. Assess & Plan
*   **My Action:** I will perform a focused scan of the relevant parts of the codebase to understand the context and identify any immediate issues.
*   **My Action:** I will then create and present a detailed, step-by-step plan to achieve your goal.

> **Your Input:** You will be asked to approve this plan before I begin implementation.
> *   **(A)pprove:** I will proceed with the approved plan.
> *   **(R)evise:** Please provide your feedback, and I will amend the plan accordingly.

---

## Phase 2: Development & Verification

Once a plan is approved, I will follow this cycle for **each step** in the plan.

### 2.1. Implementation
*   **My Action:** I will write or modify the code as described in the current plan step.

### 2.2. Quality & Verification
*   **My Action:** After implementing the changes, I will run all required quality checks to ensure the code is correct and adheres to our standards. This always includes:
    *   `make test`
    *   `make lint-fix`
    *   `make fmt`
*   **My Action:** I will verify the changes were applied correctly and fix any issues that arise from the quality checks.

### 2.3. Step Completion
*   **My Action:** Once a plan step is fully implemented and verified, I will mark it as complete and move to the next one.

---

## Phase 3: Submission

### 3.1. Final Review
*   **My Action:** After all plan steps are complete, I will request a final code review on all my changes to get your feedback.

### 3.2. Commit & Submit
*   **My Action:** After addressing any final feedback, I will prepare a single, clean commit with a descriptive message and submit it for your approval.

---

## Special Case: Full Codebase Audit

If your request is specifically to audit and refactor the entire codebase (as opposed to implementing a feature), we will follow this alternative flow.

1.  **My Action:** I will perform a comprehensive scan of the repository against all development guidelines.
2.  **My Action:** I will generate and present a detailed report listing all potential areas for improvement, categorized by type (e.g., `1. Testing`, `2. Error Handling`, `3. Naming`).

> **Your Input:** Please choose how to proceed based on the report.
> *   **(T)ackle <Area #>:** Tell me to start working on a specific area (e.g., `T 1` to tackle Testing). I will then begin at **Phase 1 (Task Intake & Planning)** for that specific area.
> *   **(S)top:** End the audit process.

---

## Content Style Guide

*   Always use lists (unordered or ordered) to define a rule.
