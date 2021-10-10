set -o errexit
set -o pipefail
set -o xtrace

curl -s http://localhost:8080/notes | jq