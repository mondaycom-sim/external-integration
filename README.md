# monday.com External Integration

Third-party vendor integration service for monday.com. Provides AI-powered
data enrichment using Cohere and OpenAI models via cross-account OIDC trust.

## Architecture

- **Runtime:** AWS Lambda (Go, arm64)
- **AI Providers:** Cohere, OpenAI
- **Auth:** GitHub Actions OIDC -> AWS IAM Role
- **Data Flow:** External vendor API -> monday.com boards
