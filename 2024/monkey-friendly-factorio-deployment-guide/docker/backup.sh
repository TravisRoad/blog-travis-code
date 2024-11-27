#!/usr/bin/env bash
TEMP_DIR=$(mktemp -d)
TS=$(date '+%Y-%m-%d_%H%M%S')
BACKUP_NAME="space-age-${TS}"
TARGET_FILE_NAME="${TEMP_DIR}/${BACKUP_NAME}"

PWD="$(dirname "${BASH_SOURCE[0]}")"

{
  cd "${PWD}" || exit 1
  tar --exclude='*.zip' --exclude='.lock' --exclude='data/temp/**' -cf "${TARGET_FILE_NAME}.tar" ./**
  tar -rf "${TARGET_FILE_NAME}.tar" data/saves/_autosave1.zip data/mods
  gzip -c "${TARGET_FILE_NAME}.tar" >"${TARGET_FILE_NAME}.tar.gz"
}

echo "${TARGET_FILE_NAME}.tar.gz"
