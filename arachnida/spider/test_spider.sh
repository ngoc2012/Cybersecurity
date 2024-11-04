#!/bin/bash

# Function to run a command and check its output
run_test() {
    local cmd="$1"
    local expected_output="$2"
    local test_name="$3"

    echo "Running test: $test_name"
    output=$(eval $cmd 2>&1)
    if [[ "$output" == *"$expected_output"* ]]; then
        echo "Test passed: $test_name"
    else
        echo "Test failed: $test_name"
        echo "Expected: $expected_output"
        echo "Got: $output"
    fi
    echo
}

# Test missing URL argument
run_test "./spider -r -l 3 -p ./images" "Usage: ./spider [-rlp] URL" "Missing URL argument"

# Test invalid URL
run_test "./spider -r -l 3 -p ./images http://invalidurl" "failed to fetch" "Invalid URL"

# Test missing recursive flag
run_test "./spider -l 3 -p ./images http://example.com" "Usage: ./spider [-rlp] URL" "Missing recursive flag"

# Test missing depth flag with recursive flag
run_test "./spider -r -p ./images http://example.com" "Usage: ./spider [-rlp] URL" "Missing depth flag with recursive flag"

# Test missing path flag
run_test "./spider -r -l 3 http://example.com" "Usage: ./spider [-rlp] URL" "Missing path flag"

# Test valid command
run_test "./spider -r -l 3 -p ./images http://example.com" "" "Valid command"

# Test invalid depth value
run_test "./spider -r -l -1 -p ./images http://example.com" "invalid value" "Invalid depth value"

# Test invalid path value
run_test "./spider -r -l 3 -p /invalid/path http://example.com" "no such file or directory" "Invalid path value"

# Test missing all flags
run_test "./spider http://example.com" "Usage: ./spider [-rlp] URL" "Missing all flags"

# Test invalid flag
run_test "./spider -x http://example.com" "flag provided but not defined" "Invalid flag"

# Test invalid URL format
run_test "./spider -r -l 3 -p ./images example.com" "invalid URL escape" "Invalid URL format"

# Test recursive download with depth 0
run_test "./spider -r -l 0 -p ./images http://example.com" "" "Recursive download with depth 0"

# Test recursive download with depth greater than default
run_test "./spider -r -l 10 -p ./images http://example.com" "" "Recursive download with depth greater than default"

# Test recursive download with default depth
run_test "./spider -r -p ./images http://example.com" "" "Recursive download with default depth"

# Test non-recursive download
run_test "./spider -l 3 -p ./images http://example.com" "" "Non-recursive download"

# Test non-recursive download with default depth
run_test "./spider -p ./images http://example.com" "" "Non-recursive download with default depth"

# Test non-recursive download with default path
run_test "./spider -l 3 http://example.com" "" "Non-recursive download with default path"

# Test non-recursive download with default depth and path
run_test "./spider http://example.com" "" "Non-recursive download with default depth and path"
