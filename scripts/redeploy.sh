#!/bin/bash
set -e

echo "=== Quick Redeploy Script ==="

# 1. Build the images
echo "[1/3] Building Docker images..."
./scripts/build_images.sh

# 2. Restart the pod to pull new image
echo "[2/3] Restarting go-api deployment..."
kubectl rollout restart deployment/go-api

# 3. Wait for rollout
echo "[3/3] Waiting for deployment to be ready..."
kubectl rollout status deployment/go-api --timeout=120s

# 4. Get new pod name
POD_NAME=$(kubectl get pods -l app=go-api -o jsonpath="{.items[0].metadata.name}")
echo ""
echo "✅ Deployment complete!"
echo "Pod name: $POD_NAME"
echo ""
echo "To view logs:"
echo "  kubectl logs -f $POD_NAME"
