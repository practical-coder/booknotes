set -o errexit
set -o pipefail
set -o xtrace

curl -s -XPUT http://localhost:8080/notes/$1 -d '{"description":"Kubernetes, GRPC, Cloud Native Go", "tags":["Go","GRPC","Kubernetes"]}' | jq