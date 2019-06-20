#!/bin/bash

set -e

go test -covermode=set -coverprofile "cover.out" github.com/Heiko-san/mermaidgen/flowchart github.com/Heiko-san/mermaidgen/gantt github.com/Heiko-san/mermaidgen/sequence
go tool cover -html="cover.out" -o cover.html
