# Consensus-Breaking Change in crudchain

## What is Breaking Consensus?
Breaking consensus refers to making a change to the blockchain's application logic that is not backward compatible. This means that nodes running the new version of the application will validate and process transactions differently than nodes running the old version, leading to a situation where blocks produced by one version are rejected by another.

## Description of the Change
I change the logic so that every resource is assigned a fixed ID (e.g., 50) instead of using an auto-incrementing counter. This means that if some nodes use the old logic and others use this new logic, they will derive different state roots when processing the same transactions.

## Why This Breaks Consensus:

1. State Mismatch: If some nodes have not upgraded and still use the auto-increment counter, the same transaction processed on different nodes will yield different IDs for the resource. This inconsistency leads to a state mismatch.

2. Determinism: Consensus requires that all nodes deterministically compute the same state given the same transactions. Changing the logic to use a fixed value violates this rule and causes blocks produced by one set of nodes to be rejected by the other.