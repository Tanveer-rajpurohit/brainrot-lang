# BrainRot Lang - Test Runner v6
# Usage: powershell -ExecutionPolicy Bypass -File scripts/test.ps1

$BINARY       = ".\dist\brainrot.exe"
$BANNER_LINES = 6

$pass = 0
$fail = 0

function Get-CleanOutput {
    param([string[]]$lines)
    if ($lines.Count -le $BANNER_LINES) { return "" }
    # skip banner, trim ALL whitespace including \r from each line
    $result = [System.Collections.Generic.List[string]]::new()
    foreach ($line in $lines[$BANNER_LINES..($lines.Count - 1)]) {
        $result.Add($line.Trim("`r`n `t"))
    }
    # drop trailing blank lines
    while ($result.Count -gt 0 -and $result[$result.Count - 1] -eq "") {
        $result.RemoveAt($result.Count - 1)
    }
    return ($result -join "`n")
}

function Get-CleanExpected {
    param([string]$path)
    $raw = [System.IO.File]::ReadAllBytes($path)
    # decode as UTF8, normalize ALL line endings, trim
    $text = [System.Text.Encoding]::UTF8.GetString($raw)
    $text = $text -replace "`r`n", "`n"
    $text = $text -replace "`r",   "`n"
    return $text.Trim()
}

Write-Host ""
Write-Host "======================================" -ForegroundColor Cyan
Write-Host "   BrainRot Lang Test Runner v6" -ForegroundColor Cyan
Write-Host "   (with output comparison)" -ForegroundColor Cyan
Write-Host "======================================" -ForegroundColor Cyan
Write-Host ""

# ── VALID TESTS ───────────────────────────────────────────────────────────────
$validTests = @(
    "test_variables",
    "test_arithmetic",
    "test_ifelse",
    "test_while",
    "test_for",
    "test_functions",
    "test_recursion",
    "test_arrays",
    "test_operators",
    "test_scope"
)

Write-Host "-- VALID TESTS (exit code + output match) --" -ForegroundColor Yellow
Write-Host ""

foreach ($test in $validTests) {
    $brlPath      = [System.IO.Path]::Combine((Get-Location).Path, "tests", "$test.brt")
    $expectedPath = [System.IO.Path]::Combine((Get-Location).Path, "tests", "expected", "$test.expected")

    $rawLines = & $BINARY run $brlPath 2>&1
    $exitCode = $LASTEXITCODE

    if ($exitCode -ne 0) {
        Write-Host "  [FAIL] $test.brt  (crashed)" -ForegroundColor Red
        if ($rawLines.Count -gt $BANNER_LINES) {
            $rawLines[$BANNER_LINES..($rawLines.Count-1)] | ForEach-Object {
                Write-Host "         $($_.Trim())" -ForegroundColor DarkRed
            }
        }
        $fail++
        continue
    }

    if (-not [System.IO.File]::Exists($expectedPath)) {
        Write-Host "  [SKIP] $test.brt  (no .expected file)" -ForegroundColor Yellow
        continue
    }

    $actual   = (Get-CleanOutput $rawLines).Trim()
    $expected = (Get-CleanExpected $expectedPath).Trim()

    if ($actual -eq $expected) {
        Write-Host "  [PASS] $test.brt" -ForegroundColor Green
        $pass++
    } else {
        Write-Host "  [FAIL] $test.brt  (output mismatch)" -ForegroundColor Red
        $aLines = $actual   -split "`n"
        $eLines = $expected -split "`n"
        $max    = [Math]::Max($aLines.Count, $eLines.Count)
        for ($i = 0; $i -lt $max; $i++) {
            $a = if ($i -lt $aLines.Count) { $aLines[$i] } else { "<missing>" }
            $e = if ($i -lt $eLines.Count) { $eLines[$i] } else { "<missing>" }
            if ($a -ne $e) {
                Write-Host "         Line $($i+1)  expected: '$e'" -ForegroundColor Green
                Write-Host "         Line $($i+1)  actual:   '$a'" -ForegroundColor Red
            }
        }
        $fail++
    }
}

# ── ERROR TESTS ────────────────────────────────────────────────────────────────
Write-Host ""
Write-Host "-- ERROR TESTS (should produce errors) --" -ForegroundColor Yellow
Write-Host ""

$errorTests = @(
    @{ file = "test_err_undefined_var";  keyword = "ghosted" },
    @{ file = "test_err_duplicate_var";  keyword = "already declared" },
    @{ file = "test_err_no_main";        keyword = "no main" },
    @{ file = "test_err_wrong_args";     keyword = "args" },
    @{ file = "test_err_toplevel_code";  keyword = "not allowed" }
)

foreach ($test in $errorTests) {
    $path     = [System.IO.Path]::Combine((Get-Location).Path, "tests", "$($test.file).brt")
    $output   = (& $BINARY run $path 2>&1) -join " "
    $exitCode = $LASTEXITCODE

    if ($exitCode -ne 0 -and ($output -match $test.keyword)) {
        Write-Host "  [PASS] $($test.file).brt - error caught" -ForegroundColor Green
        $pass++
    } elseif ($exitCode -eq 0) {
        Write-Host "  [FAIL] $($test.file).brt - expected error but ran ok!" -ForegroundColor Red
        $fail++
    } else {
        Write-Host "  [FAIL] $($test.file).brt - wrong error (wanted: '$($test.keyword)')" -ForegroundColor Red
        $fail++
    }
}

# ── CLI MODE TESTS ─────────────────────────────────────────────────────────────
Write-Host ""
Write-Host "-- CLI MODE TESTS (tokens / ast) --" -ForegroundColor Yellow
Write-Host ""

$cliTests = @(
    @{ cmd = "tokens"; file = "examples\hello.brt";     label = "tokens - hello" },
    @{ cmd = "ast";    file = "examples\hello.brt";     label = "ast    - hello" },
    @{ cmd = "tokens"; file = "examples\fibonacci.brt"; label = "tokens - fibonacci" },
    @{ cmd = "ast";    file = "examples\fibonacci.brt"; label = "ast    - fibonacci" }
)

foreach ($t in $cliTests) {
    $output   = & $BINARY $t.cmd $t.file 2>&1
    $exitCode = $LASTEXITCODE
    if ($exitCode -eq 0) {
        Write-Host "  [PASS] $($t.label)" -ForegroundColor Green
        $pass++
    } else {
        Write-Host "  [FAIL] $($t.label)" -ForegroundColor Red
        $fail++
    }
}

# ── SUMMARY ────────────────────────────────────────────────────────────────────
$total = $pass + $fail
Write-Host ""
Write-Host "======================================" -ForegroundColor Cyan
Write-Host "  Results: $pass / $total passed" -ForegroundColor White
if ($fail -eq 0) {
    Write-Host "  ALL TESTS PASSED - fr fr no cap" -ForegroundColor Green
} else {
    Write-Host "  $fail TESTS FAILED - skill issue bro" -ForegroundColor Red
}
Write-Host "======================================" -ForegroundColor Cyan
Write-Host ""