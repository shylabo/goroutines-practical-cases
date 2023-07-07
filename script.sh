#!/bin/bash

read -p "Enter directory name: " dir_name

mkdir "$dir_name"

cd "$dir_name"

cat <<EOF > main.go
package main

/**
- -

*/

func main() {
  // Add code
}
EOF

go mod init "github.com/shylabo/goroutines-practical-cases/$dir_name"

cd ..

go work use "./$dir_name"
