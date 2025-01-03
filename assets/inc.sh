#!/bin/bash

# Increase brightness by 10%
current_brightness=0

while IFS= read -r line; do
  if [[ $line =~ (current|brightness) ]]; then
    current_brightness=$(echo "$line" | awk '{print $NF}')
    break
  fi
done < <(sudo ddcutil getvcp 10)

if [[ -z "$current_brightness" ]]; then
  echo "Error: Could not parse brightness value"
  exit 1
fi

new_brightness=$((current_brightness + 10))

# Ensure brightness stays within 0-100 range
if [[ $new_brightness -gt 100 ]]; then
        new_brightness=80
fi

sudo ddcutil setvcp 10 $new_brightness

echo "Brightness increased to $new_brightness%"
