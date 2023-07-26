# Discussions and Issues

Trivy uses [GitHub Discussion](https://github.com/aquasecurity/trivy/discussions) for bug reports, feature requests, and questions, and [GitHub Issues](https://github.com/aquasecurity/trivy/issues) to track work items and progress.
We do not allow users and contributors to open and assign issues directly, instead we as you to start a discussion. If maintainers decide to accept a new feature or acknowledge a bug, they will convert the discussion into a [GitHub Issue](https://github.com/aquasecurity/trivy/issues) work item, where you can track it's progress and discuss implementation details.

- Feel free to open discussions for any reason. When you open a new discussion, you'll need to select the appropriate discussion category.
- Before starting a new discussion, please search existing issues/discussion to confirm it's not a duplicate. If it is, please add your comment to the existing thread instead of starting a new one.
- Please give your discussion a meaningful title to help others find it in the future.
- The issue should clearly explain the reason for opening, the proposal if you have any, and any relevant technical information.

There are 4 categories:

- 💡 [Ideas](https://github.com/aquasecurity/trivy/discussions/categories/ideas)
    - Share ideas for new features
- 🔎 [False Detection](https://github.com/aquasecurity/trivy/discussions/categories/false-detection)
    - Report false positives/negatives, meaning an problem with the results that Trivy gave you
- 🐛 [Bugs](https://github.com/aquasecurity/trivy/discussions/categories/bugs)
    - Report something that is not working as expected
- 🙏 [Q&A](https://github.com/aquasecurity/trivy/discussions/categories/q-a)
    - Ask the community for help
- 📖 [Documentation](https://github.com/aquasecurity/trivy/discussions/new?category=documentation)
    - Any suggestion or problem you find in Trivy's documentation

!!! note
    If you find any false positives or false negatives, please make sure to report them under the "False Detection" category, not "Bugs".

## False detection
Trivy depends on [multiple data sources](https://aquasecurity.github.io/trivy/latest/docs/vulnerability/detection/data-source/). Sometime these sources contain mistakes, in which case we do not consider it a Trivy issue.  
To check if the issue you have stems from a data source, please follow the steps below:

1. Run Trivy with `-f json` to show which data sources were used.
2. Find the problematic finding, and identify it's data source.
3. Go to the data source and search for the CVE ID. Data source locations are detailed [here for OS packages](https://aquasecurity.github.io/trivy/latest/docs/scanner/vulnerability/os/#data-source-selection) and [here for programming language packages](https://aquasecurity.github.io/trivy/latest/docs/scanner/vulnerability/language/#data-sources).
4. If the problem appears also in the data source, please report it to the data source maintainers. If the problem is in Trivy, please report it to us (as a "false detection").
