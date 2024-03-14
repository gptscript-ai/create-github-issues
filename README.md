# Create GitHub Issues

This is a GPTScript tool that uses the GitHub API to create issues.

Set your GitHub access token to the `GPTSCRIPT_GITHUB_TOKEN` environment variable in order to use this tool.
If the variable is not set, the tool will attempt to make unauthenticated API calls.

## Example

```yaml
tools: github.com/gptscript-ai/create-github-issues

Create an issue called "My issue" in the g-linville/test repo. Assign it to g-linville.
The issue body should state that the issue was created by GPTScript.
```

## License

This tool is available under the Apache License 2.0. See [LICENSE](LICENSE) for more information.
