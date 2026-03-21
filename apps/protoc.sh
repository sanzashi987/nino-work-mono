export PATH="$PATH:$(go env GOPATH)/bin"

files=$(find proto -iname "*.proto")

for file in $files; do
  protoc --micro_out=. --go_out=. $file
done
