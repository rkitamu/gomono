#!/bin/bash
set -euo pipefail

# Make sure ~/.local/bin is in PATH
export PATH="$HOME/.local/bin:$PATH"
echo 'export PATH="$HOME/.local/bin:$PATH"' >> $HOME/.bashrc

echo "[setup.sh] Installing gopls..."
go install golang.org/x/tools/gopls@latest

echo "[setup.sh] Installing dlv (Delve debugger)..."
go install github.com/go-delve/delve/cmd/dlv@latest

echo "[setup.sh] ✅ Done!"
