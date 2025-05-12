## Deployment options

### Docker Compose

- Complexity: Minimal—just one YAML file.
- Scalability: Single-host only; no autoscaling.
- Control & Features: Basic networking, volumes; no built-in health-checks or self-healing.
- Cost & Ops: Free, zero infra ops—ideal for local dev or small proofs.

### Docker Swarm

- Complexity: Low—native Docker CLI.
- Scalability: Multi-host clustering with simple scaling.
- Control & Features: Rolling updates, service discovery; limited ecosystem.
- Cost & Ops: Light ops; fewer third-party tools than K8s.

### Kubernetes (self-hosted)

- Complexity: High—requires managing control-plane, networking, storage.
- Scalability: Enterprise-grade; auto-scale, rolling upgrades, self-healing.
- Control & Features: Maximum flexibility and extensibility.
- Cost & Ops: Significant operational overhead; steep learning curve.

### Managed Kubernetes (GKE/EKS/AKS)

- Complexity: Medium—the control plane is managed for you.
- Scalability: Same as K8s; easy cluster scaling with cloud APIs.
- Control & Features: Most K8s features, plus cloud integrations (IAM, load-balancers).
- Cost & Ops: Higher hourly charges; less DIY ops but some vendor lock-in.

### Serverless Containers (Cloud Run, Fargate)

- Complexity: Very low—deploy your image, and it “just runs.”
- Scalability: Instant, per-request autoscaling.
- Control & Features: Stateless only; limited to platform’s runtime constraints.
- Cost & Ops: Pay-for-use; no cluster to manage, but potential cold-start latency.
