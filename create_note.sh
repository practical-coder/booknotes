set -o errexit
set -o pipefail
set -o xtrace

curl -s -XPOST http://localhost:8080/notes -d "{\"title\":\"${1}\"}" | jq