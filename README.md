# aadhaar

[![CI](https://github.com/Rugved2204/aadhaar/actions/workflows/ci.yml/badge.svg)](https://github.com/Rugved2204/aadhaar/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/Rugved2204/aadhaar.svg)](https://pkg.go.dev/github.com/Rugved2204/aadhaar)

Validate Indian Aadhaar numbers in Go using the Verhoeff checksum. Zero dependencies, no I/O.

> **Important:** a passing check means the number is *well-formed* (12 digits,
> first digit 2–9, valid Verhoeff check digit). It does **not** prove the number
> was issued or belongs to anyone. Only UIDAI can verify that.

## Install

```bash
go get github.com/Rugved2204/aadhaar
```

## Usage

```go
import "github.com/Rugved2204/aadhaar"

aadhaar.Validate("XXXX XXXX XXXX") // spaces and hyphens are ignored
aadhaar.Mask(n)                    // "XXXX XXXX 0123"
aadhaar.Normalize("2345-6789-0123")

// Test fixtures only; these are NOT real Aadhaar numbers.
n := aadhaar.GenerateTestNumber(nil)
```

## Privacy

Aadhaar numbers are sensitive personal data (Aadhaar Act, DPDP Act). This
library never logs, stores or transmits input. You are responsible for lawful
handling in your application. Do not commit real Aadhaar numbers to this repo,
including in tests or issues.

## Development

```bash
go test -race -cover ./...
go test -fuzz=FuzzValidate -fuzztime=30s
```

## License

MIT
