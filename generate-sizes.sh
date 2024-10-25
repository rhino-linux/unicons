#!/usr/bin/env bash

sizes=(16 18 22 24 32 42 48 64 84 96 128)

for icon in scalable/apps/*.svg; do
  iconname=$(basename "${icon}" .svg)
  for s in "${sizes[@]}"; do
    size="${s}x${s}/apps"
    mkdir -p "${size}"
    rm -rf "${size}"/*
    iconsize="${size}/${iconname}.png"
    echo "Generating ${iconsize}..."
    if command -v rsvg-convert > /dev/null; then
      rsvg-convert -w "${s}" -h "${s}" "${icon}" -o "${iconsize}"
    elif command -v inkscape > /dev/null; then
      inkscape "${icon}" --export-width="${s}" --export-height="${s}" --export-filename="${iconsize}"
    else
      echo "Need rsvg-convert or inkscape to convert."
      exit 1
    fi
  done
done
echo "PNG generation completed."
