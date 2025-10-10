#!/bin/bash

filename="$1"
desired_width="$2"

if [ -z filename ] || [ -z desired_width ]; then
  echo "Usage $0 <input-image> <desired_width>"
  exit 1
fi

function convert_file() {
  local width=$1
  local density=$2
  local format=$3

  local result_file="${basename}@${density}x.${format}"

  local parameters=""
  if [ $format = "web" ]; then
    parameters="-quality 80 -define webp:lossless=false -define webp:method=4"
  elif [ $format = "avif" ]; then
    parameters="-quality 50 -define avif:lossless=false -define avif:method=4"
  fi

  magick $filename \
    -resize ${width}x \
    ${parameters} \
    ${result_file}

  echo "✅ Created ${result_file} (width=${width})"
}

basename="${filename%.*}"
base_format="${filename##*.}"
echo "Converting image '$basename' for serving on the web"

densities=(1 2 3)
formats=(${base_format} webp avif)

for format in "${formats[@]}"; do 
  for density in "${densities[@]}"; do
    width=$((density * desired_width))
    convert_file ${width} ${density} ${format}
  done
done

