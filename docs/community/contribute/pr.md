Thank you for taking interest in contributing to Trivy!

When creating a Pull Request, please follow the following guidelines:

1. Every Pull Request should have an associated GitHub Issue (unless you are fixing a trivial documentation issue).
1. Please add the associated Issue link in the PR description.
1. Please include before and after examples of your change in the PR description.
1. There's no need to add or tag reviewers.
1. If a reviewer commented on your code or asked for changes, please remember to respond with comment. Do not mark discussion as resolved. It's up to reviewer to mark it resolved (in case if suggested fix addresses problem properly). PRs with unresolved issues should not be merged (even if the comment is unclear or requires no action from your side).
1. Please make sure to add tests if relevant.
1. For feature PRs, please make sure to update the documentation accordingly.

## Understand where your pull request belongs

Trivy is composed of several repositories that work together:

- [Trivy](https://github.com/aquasecurity/trivy) is the command line tool, most of the code, documentation, and release artifacts.
- [trivy-db](https://github.com/aquasecurity/trivy-db) is Trivy's vulnerability database, that is pulled by Trivy. The data source for the database is controlled by code in[vuln-list-update](https://github.com/aquasecurity/vuln-list-update).
- [go-dep-parser](https://github.com/aquasecurity/go-dep-parser) is a library for parsing package management files (e.g package.json).
- [defsec](https://github.com/aquasecurity/defsec) is Trivy's misconfiguration scanning engine (inclusing IaC and cloud scanning).
- [trivy-k8s](https://github.com/aquasecurity/trivy-k8s) has the code for scanning Kubernetes clusters.

## Title format

We follow [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) format.

```
<type>[optional scope]: <description>
```

Allowed `<type>` values:

- **feat** for a new feature for the user, not a new feature for build script. Such commit will trigger a release bumping a MINOR version.
- **fix** for a bug fix for the user, not a fix to a build script. Such commit will trigger a release bumping a PATCH version.
- **docs** for changes to the documentation.
- **chore** for changes outside the core functionality of the project, for example tools and processes.
- **refactor** for refactoring production code, e.g. renaming a variable.
- **test** for adding and fixing tests.
- **revert** for revert to a previous commit

`<scope>` is optional. Allowed values:

- Any of the Trivy scanners (for example `vuln`, `misconf`)
- Any vulnerability scanner supported OS or Language (for example `alpine`, `ruby`)
- Any misconfiguration scanner supported formats (for example `terraform`, `kubernetes`)
- Any of the Trivy targets (for example `image`, `fs`)

`<subject>` is a short description of the change. 

- Use [Imperative mood](https://cbea.ms/git-commit/#imperative) in the subject line. For example, use `add` instead of `added` or `adds`.
- Describe [what is the purpose of the change](https://cbea.ms/git-commit/#why-not-how), not how it is implemented. For example, use `check user input to avoid crash` instead of `check variable value before use`.

### Example titles

```
feat(alma): add support for AlmaLinux
```

```
fix(oracle): handle advisories with ksplice versions
```

```
docs(misconf): add comparison with Conftest and TFsec
```

```
chore(deps): bump go.uber.org/zap from 1.19.1 to 1.20.0
```
