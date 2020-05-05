# Template Go Standard

![Go](https://github.com/FrancescoIlario/go-std-tmpl/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/FrancescoIlario/go-std-tmpl)](https://goreportcard.com/report/github.com/FrancescoIlario/go-std-tmpl)
![Docker Image CI](https://github.com/FrancescoIlario/go-std-tmpl/workflows/Docker%20Image%20CI/badge.svg)

Template project based on the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

It's enriched with the following supports:

- Go Modules
- Visual Studio Code's Remote-Containers 
- Github Actions for Go, Docker and Helm on AKS
- Dockerfile for a Release and Distroless build
- Helm and K8S manifests

# Setup

Substitute in every file `FrancescoIlario/go-std-tmpl` with your Github path.
You can use the following line of bash code substituting the fields `<YOUR_GH_ACCOUNT>` and `<YOUR_GH_REPO>`:

```bash
grep -iRl FrancescoIlario/go-std-tmpl . \
    | xargs sed -i 's/FrancescoIlario\/go-std-tmpl/<YOUR_GH_ACCOUNT>\/<YOUR_GH_REPO>/g' 
```

## Visual Studio Code's Remote-Containers

To use this extension you have to install `VSCode` and its extension `Remote-Containers` from its marketplace.

Open the project, press F1, and look for the voice `Remote-Containers: Reopen in Container...`
