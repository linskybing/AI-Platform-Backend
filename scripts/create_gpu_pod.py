from kubernetes import client, config
import sys

def create_gpu_pod(pod_name, namespace, vram_limit, thread_percentage):
    """
    Creates a Pod with specific NVIDIA MPS annotations to control GPU resources.
    """
    # Load kube config (works for both in-cluster and local kubeconfig)
    try:
        config.load_kube_config()
    except:
        config.load_incluster_config()

    api_instance = client.CoreV1Api()

    # Define the Pod Manifest
    pod_manifest = {
        "apiVersion": "v1",
        "kind": "Pod",
        "metadata": {
            "name": pod_name,
            "annotations": {
                # Dynamic Resource Control Annotations
                "mps.nvidia.com/vram": vram_limit,       # e.g., "4G", "1280M"
                "mps.nvidia.com/threads": str(thread_percentage) # e.g., "10"
            }
        },
        "spec": {
            "restartPolicy": "Never",
            "containers": [
                {
                    "name": "gpu-workload",
                    "image": "nvidia/cuda:11.8.0-base-ubuntu22.04",
                    "command": ["/bin/bash", "-c"],
                    "args": [
                        "echo 'Starting Workload...'; "
                        "env | grep CUDA_MPS; "
                        "nvidia-smi; "
                        "sleep 3600"
                    ],
                    "resources": {
                        "limits": {
                            # Request 1 Time-Slice (1/25th of a card)
                            # The actual VRAM/Compute is enforced by the Env Vars injected by Kyverno
                            "nvidia.com/gpu": "1" 
                        }
                    },
                    # Volume mount for MPS pipe is required for the container to talk to the daemon
                    "volumeMounts": [
                        {
                            "name": "mps-pipe",
                            "mountPath": "/run/nvidia/mps"
                        },
                        {
                            "name": "shm",
                            "mountPath": "/dev/shm"
                        }
                    ]
                }
            ],
            "volumes": [
                {
                    "name": "mps-pipe",
                    "hostPath": {
                        "path": "/run/nvidia/mps",
                        "type": "Directory"
                    }
                },
                {
                    "name": "shm",
                    "emptyDir": {
                        "medium": "Memory"
                    }
                }
            ]
        }
    }

    try:
        api_response = api_instance.create_namespaced_pod(
            namespace=namespace,
            body=pod_manifest
        )
        print(f"Pod '{pod_name}' created successfully.")
        print(f"  - VRAM Limit: {vram_limit}")
        print(f"  - Thread Limit: {thread_percentage}%")
    except client.exceptions.ApiException as e:
        print(f"Exception when calling CoreV1Api->create_namespaced_pod: {e}")

if __name__ == "__main__":
    # Example Usage: Create a "Premium" tier pod
    create_gpu_pod(
        pod_name="user-101-premium-workload",
        namespace="default",
        vram_limit="4G",      # 4GB VRAM
        thread_percentage=10  # 10% Compute
    )
