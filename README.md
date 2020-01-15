# Revolut Stocks List ![GitHub](https://img.shields.io/github/license/nmapx/revolut-stocks-list?style=flat) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/nmapx/revolut-stocks-list?style=flat) [![Go Report Card](https://goreportcard.com/badge/github.com/nmapx/revolut-stocks-list)](https://goreportcard.com/report/github.com/nmapx/revolut-stocks-list) [![Build Status](https://travis-ci.com/nmapx/revolut-stocks-list.svg?branch=master)](https://travis-ci.com/nmapx/revolut-stocks-list)

Extract Revolut stocks list from the list screenshot(s) ([example screenshot](./example-input.jpg) - use your mobile app for this).

I made this app since there is no official list of all available Revolut stocks.

**You can get a full list at the bottom of this page. It may not be up to date but I will do my best to keep it updated!**

## Technologies

- Go
- Docker with **compose**
- Tesseract OCR
- Make

## Requirements

- Docker >= 19.03 with **compose** >= 1.25
- Make >= 4.1 *(optional)*
> Note: use appropriate commands (found in Makefiles) without Make

## Setup guide

1. Clone the repository and navigate to the project root
2. Copy `.env.dist` file to `.env` and fill missing parameters
> Note: HOST_UID is your system user ID
3. Build docker image `make build`
4. Start container `make up` (development purposes only)
5. Run the script by executing appropriate command inside the container (see below)
6. Compile the app by executing `make -f Makefile.native build` inside the container

## Basic usage

### Docker

```bash
docker exec {CONTAINER_NAME} go run . extract
```

or

```bash
docker-compose -p {PROJECT_NAME} exec app go run . extract
```

### Binary

Binary not supported yet - use Docker environment.

### I'm just looking for the list

Scroll down. The list is as the bottom of this page.

## Advanced usage

### Input files (images/screenshots)

First of all - check [example screenshot](./example-input.jpg). Your screenshot(s) must follow the pattern!

You can specify multiple input files by adding more parameters. By default it's processing 1 file (filepath `./input.jpg`) which can be very long (complete stock list in 1 screenshot (possible on some mobiles)). See example below.

```bash
... extract --input ./this/is/input/file_1.jpg --input ./this/is/input/file_2.jpg --input ./this/is/input/file_3.jpg
```

### Output file (CSV)

Script generates 1 output file (including all input files in it). By default it's `./output.csv` - you can change the filepath. See example below.

```bash
... extract --output ./this/is/output/file.csv
```

### Languages

By default only English is being installed and used by the Tesseract library. In order to use also other languages like Polish or Russian you can add an optional parameters. You need to install appropriate language library (Tesseract) first (eg. [tesseract-ocr-data-pol](https://pkgs.alpinelinux.org/package/edge/community/x86_64/tesseract-ocr-data-pol)). See example below.

```bash
... extract --lang eng --lang pol --lang rus
```

### Whitelist

By default whitelist includes only A-Z characters (upper-case) for tickers and whole alphabet for companies names. It's the best setup but you can change it if you want to experiment a bit. See example below.

```bash
... extract --whitelist-tickers abcDEF123 --whitelist-names qwertyQWERTY
```

## License

[MIT License](./LICENSE)

## The list

As the result you will get a file with all recognized tickers and companies names.

**OCR is still not the most accurate solution so you can find few mistakes there (about 2-3% misidentified tickers).**

[Here](./LIST.md) you can find already extracted list of all available stocks. I will update it from time to time.
