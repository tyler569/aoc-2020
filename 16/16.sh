cat input | grep -A 1000 'nearby tickets' | grep -v nearby | tr ',' '\n' | sort -n | awk '$0 > 971 || $0 < 25 { sum += $0 } END { print sum }'
