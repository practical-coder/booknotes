set -o errexit
set -o pipefail
set -o xtrace

curl -s -XDELETE http://localhost:8080/notes/$1 | jq