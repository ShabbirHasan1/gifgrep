#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
tmp_dir="$(mktemp -d)"
bin="$tmp_dir/termcaps-check"

cleanup() { rm -rf "$tmp_dir"; }
trap cleanup EXIT

cd "$repo_root"
go build -o "$bin" ./cmd/termcaps-check

wait_for_file() {
  local f="$1"
  local tries=0
  until [[ -s "$f" ]]; do
    tries=$((tries + 1))
    if [[ $tries -gt 80 ]]; then
      return 1
    fi
    sleep 0.1
  done
}

run_terminal() {
  local app="$1"
  local expect="$2"
  local out="$3"

  /usr/bin/osascript <<'APPLESCRIPT' "$app" "$bin" "$expect" "$out"
on run argv
  set appName to item 1 of argv
  set toolPath to item 2 of argv
  set expectProto to item 3 of argv
  set outPath to item 4 of argv
  set cmd to "/bin/zsh -lc " & quoted form of (toolPath & " --expect " & expectProto & " > " & outPath & " 2>&1; exit")

  if appName is "Terminal" then
    tell application "Terminal"
      activate
      set t to do script cmd
    end tell
  else if appName is "iTerm" then
    tell application "iTerm"
      activate
      set w to (create window with default profile)
      tell current session of w
        write text cmd
      end tell
    end tell
  else
    error "unsupported app: " & appName
  end if
end run
APPLESCRIPT
}

pass() { printf "PASS %s\n" "$1"; }
fail() { printf "FAIL %s\n" "$1"; exit 1; }
skip() { printf "SKIP %s\n" "$1"; }

echo "termcaps e2e (macOS)"

# Current terminal (sanity)
current_out="$tmp_dir/current.json"
if "$bin" --json >"$current_out"; then
  pass "current-terminal"
else
  fail "current-terminal"
fi

# Apple Terminal (expect none)
term_out="$tmp_dir/terminal.json"
if run_terminal "Terminal" "none" "$term_out"; then
  if wait_for_file "$term_out"; then
    pass "Apple Terminal"
  else
    fail "Apple Terminal (no output)"
  fi
else
  skip "Apple Terminal (osascript failed)"
fi

# iTerm2 (expect iterm) — skip if not installed
if /usr/bin/osascript -e 'application "iTerm" exists' >/dev/null 2>&1; then
  iterm_out="$tmp_dir/iterm.json"
  if run_terminal "iTerm" "iterm" "$iterm_out"; then
    if wait_for_file "$iterm_out"; then
      pass "iTerm2"
    else
      fail "iTerm2 (no output)"
    fi
  else
    fail "iTerm2 (osascript failed)"
  fi
else
  skip "iTerm2 (not installed)"
fi

echo "Outputs:"
ls -1 "$tmp_dir"/*.json 2>/dev/null || true

