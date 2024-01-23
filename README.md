<div align="center">
  <br />
  <br />
  <br />
  <h3><a href="https://patex.io">Patex</a> is a low-cost and lightning-fast Ethereum L2 blockchain.</h3>
  <br />
</div>

## What is Patex?

Patex is a low-cost and lightning-fast Ethereum L2 blockchain, **but it's also so much more than that**.

Patex is the technical foundation for [the Patex Collective](https://app.patex.io/announcement), a band of communities, companies, and citizens united by a mutually beneficial pact to adhere to the axiom of **impact=profit** — the principle that positive impact to the collective should be rewarded with profit to the individual.
We're trying to solve some of the most critical coordination failures facing the crypto ecosystem today.
**We're particularly focused on creating a sustainable funding stream for the public goods and infrastructure upon which the ecosystem so heavily relies but has so far been unable to adequately reward.**
We'd love for you to check out [The Patex Vision](https://www.patex.io/vision) to understand more about why we do what we do.

## Documentation

If you want to build on top of Patex, take a look at the extensive documentation on the [Patex Community Hub](http://community.patex.io/).
If you want to build Patex, check out the [Protocol Specs](./specs/).

## Community

General discussion happens most frequently on the [Patex discord](https://discord.gg/patex).
Governance discussion can also be found on the [Patex Governance Forum](https://gov.patex.io/).

## Contributing

Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of our contribution process.
Use the [Developer Quick Start](./CONTRIBUTING.md#development-quick-start) to get your development environment set up to start working on the Patex Monorepo.
Then check out our list of [good first issues](https://github.com/patex-ecosystem/patex-network/contribute) to find something fun to work on!

## Security Policy and Vulnerability Reporting

Please refer to our canonical [Security Policy](https://github.com/ethereum-patex/.github/blob/master/SECURITY.md) document for detailed information about how to report vulnerabilities in this codebase.
Bounty hunters are encouraged to check out [our Immunefi bug bounty program](https://immunefi.com/bounty/patex/).



## Directory Structure

<pre>
~~ Production ~~
├── <a href="./packages">packages</a>
│   ├── <a href="https://github.com/patex-ecosystem/contracts">contracts</a>: L1 and L2 smart contracts for Patex
│   ├── <a href="https://github.com/patex-ecosystem/hardhat-deploy-config">hardhat-deploy-config</a>: simple plugin that adds support for global deploy configuration values
│   ├── <a href="https://github.com/patex-ecosystem/core-utils">core-utils</a>: Low-level utilities that make building Patex easier
│   └── <a href="https://github.com/patex-ecosystem/sdk">sdk</a>: provides a set of tools for interacting with Patex
├── <a href="./integration-tests">integration-tests</a>: Various integration tests for the Patex network
├── <a href="./proxyd">proxyd</a>: Configurable RPC request router and proxy
├── <a href="./technical-documents">technical-documents</a>: audits and post-mortem documents

~~ BEDROCK ~~
├── <a href="./packages">packages</a>
│   └── <a href="https://github.com/patex-ecosystem/contracts-bedrock">contracts-bedrock</a>: Bedrock smart contracts. To be merged with ./packages/contracts.
├── <a href="./pt-bindings">pt-bindings</a>: Go bindings for Bedrock smart contracts.
├── <a href="./pt-batcher">pt-batcher</a>: L2-Batch Submitter, submits bundles of batches to L1
├── <a href="./pt-node">pt-node</a>: rollup consensus-layer client.
├── <a href="./pt-proposer">pt-proposer</a>: L2-Output Submitter, submits proposals to L1
├── <a href="./ops-bedrock">ops-bedrock</a>: Bedrock devnet work
└── <a href="./specs">specs</a>: Specs of the rollup starting at the Bedrock upgrade
</pre>

## Branching Model

### Active Branches

| Branch                                                                    | Status      |
|---------------------------------------------------------------------------|-------------|
| [main](https://github.com/patex-ecosystem/patex-network/tree/main/)       | Main branch |
### Overview

We generally follow [this Git branching model](https://nvie.com/posts/a-successful-git-branching-model/).
Please read the linked post if you're planning to make frequent PRs into this repository (e.g., people working at/with Patex).

### Production branch

Our production branch is `main`.
The `main` branch contains the code for our latest "stable" releases.


**Changes to contracts within `packages/contracts/contracts` are usually NOT considered backwards compatible and SHOULD be made against a release candidate branch**.
Some exceptions to this rule exist for cases in which we absolutely must deploy some new contract after a release candidate branch has already been fully deployed.
If you're changing or adding a contract and you're unsure about which branch to make a PR into, default to using the latest release candidate branch.
See below for info about release candidate branches.

### Release candidate branches

Branches marked `release/X.X.X` are **release candidate branches**.
Changes that are not backwards compatible and all changes to contracts within `packages/contracts/contracts` MUST be directed towards a release candidate branch.
Release candidates are merged into `develop` and then into `master` once they've been fully deployed.
We may sometimes have more than one active `release/X.X.X` branch if we're in the middle of a deployment.
See table in the **Active Branches** section above to find the right branch to target.

## Releases

### Changesets

We use [changesets](https://github.com/changesets/changesets) to mark packages for new releases.


To add a changeset, run the command `yarn changeset` in the root of this monorepo.
You will be presented with a small prompt to select the packages to be released, the scope of the release (major, minor, or patch), and the reason for the release.
Comments within changeset files will be automatically included in the changelog of the package.


## License


All other files within this repository are licensed under the [MIT License](https://github.com/patex-ecosystem/patex-network/blob/master/LICENSE) unless stated otherwise.
