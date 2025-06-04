# terraform-batch
**A lightweight CLI tool to automate Terraform commands across multiple directories.**

## ✨ Features

- Run `terraform init` across directories
- Run `terraform validate` across directories
- Run `terraform apply` with auto-approval across directories
- Run `terraform destroy` with auto-approval across directories
- Validates directories contain a `main.tf` file before running commands
- Simple, no-frills CLI ideal for local development and CI/CD pipelines
- No dependencies besides the Terraform CLI

## 📥 Installation
Download pre-built binaries from [Releases](https://github.com/AdhamBasheir/terraform-batch/releases) or build from source:

```bash
git clone https://github.com/AdhamBasheir/terraform-batch.git
cd terraform-batch
CGO_ENABLED=0 go build -o <binary-name> .
```
### ⚠️ Note:
If you don’t add <binary-name> to your system’s PATH, you must run the tool from your Terraform root directory (the directory containing your Terraform configs) by specifying the relative or absolute path to the executable.

## 🛠️ Usage
if in path
```bash
<binary-name> <command> [directory1] [directory2] ...
```
if not in path
```bash
./<binary-name> <command> [directory1] [directory2] ...
```
### commands
| Command   | Description                                       |
| --------- | --------------------------------------------------|
| `init`    | Initialize Terraform configurations               |
| `validate`| Validate Terraform configurations                 |
| `apply`   | Apply Terraform configurations (auto-approve)     |
| `destroy` | Destroy Terraform configurations (auto-approve)   |
| `help`    | Show help information                             |

## 🗂️ Terraform Directory Structure
```
terraform/
├── foo/
│   └── main
│       └── main.tf
├── bar/
│   └── main
│       └── main.tf
├── baz/
│   └── main
│       └── main.tf
```

## 📝 Notes
- The tool requires Terraform CLI installed and available in your PATH.
- Directories must contain a valid main.tf file.
- Intended for developer use and automation scripts; avoids emojis or colors for CI compatibility.
