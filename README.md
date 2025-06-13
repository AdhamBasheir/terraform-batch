# terraform-batch
**A lightweight CLI wrapper for running Terraform commands across multiple directories.**

## ✨ Features
- Simple, efficient CLI suited for local development and CI/CD automation
- Execute any Terraform command (plan, apply, destroy, etc.)
- Operate sequentially on one or more directories
- Optional `-ignore-errors` mode to continue execution despite failures
- Clean, colored terminal output
- No dependencies beyond the Terraform CLI

## 📥 Installation
Download pre-built binaries from [Releases](https://github.com/AdhamBasheir/terraform-batch/releases) or build from source:
```bash
go install github.com/AdhamBasheir/terraform-batch@latest
```
or
```bash
git clone https://github.com/AdhamBasheir/terraform-batch.git
cd terraform-batch
CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o terraform-batch .
```

### ⚠️ **Note:**  
> Using `go install` places the binary in your system’s `PATH`, allowing you to run it by simply typing `myapp` anywhere.  
> If you clone the repository and build locally, the binary remains in the build directory, so you must execute it using its full or relative path (e.g., `/path/to/terraform-batch`).


## 🛠️ Usage
if the binary is in the system's PATH
```bash
terraform-batch [flags] <terraform command>
```

## ✅ Examples
Run terraform plan in the current directory:
```bash
terraform-batch plan
```

Run terraform apply in multiple directories:
```bash
terraform-batch -dir=dir1,dir2 apply
```

Ignore errors and continue execution:
```bash
terraform-batch -dir=dir1,dir2,dir3 -ignore-errors=true destroy
```

## ⚙️ Flags
| Flag             | Description                                         |
| ---------------- | --------------------------------------------------- |
| `-dir`           | Comma-separated list of directories                 |
| `-ignore-errors` | Continue even if a directory fails (default: false) |

## 📌 Notes
- Do not include the terraform keyword in your command. Just write the subcommand (e.g., `plan`, `apply`).

## 🚀 Planned Features
- Parallel execution with `-parallel`
- Output logging per directory with `-log`
