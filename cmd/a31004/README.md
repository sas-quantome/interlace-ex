###### September 4, 2026: d60104 readme v12

# [a31004]

Creates the human **Gene PubMed Reference** asset from the NCBI [Gene DB].

```mermaid
flowchart LR
    db1[(NCBI Gene DB)] --> gene2pubmed.gz
    env1>ENV_AEE_BASE_PATH] --> app((a31004))
    env2>ENV_ALT_DATA_PATH] --> app((a31004))
    uri[[uri.txt]] --> app((a31004))
    gene2pubmed.gz --> uri[[uri.txt]]
    app((a31004)) --> ds{{interlace-ex}}
    ds{{interlace-ex}} --> work{{event}}
    ds{{interlace-ex}} --> data{{entity}}
    work{{event}} --> workEntity{{gene}}
    data{{entity}} --> dataEntity{{gene}}
    workEntity{{gene}} --> workAsset{{w31001}}
    dataEntity{{gene}} --> dataAsset{{pubmed_reference}}
    workAsset{{w31001}} --> workRelease{{version}}
    dataAsset{{pubmed_reference}} --> dataRelease{{version}}
    workRelease{{version}} --> action{{action}}
    workRelease{{version}} --> artifact{{artifact}}
    dataRelease{{version}} --> prop{{property}}
    action{{action}} --> log[[a31004.log]]
    artifact{{artifact}} --> gene_pubmed_reference_tmp.tsv.gz
    prop{{property}} --> d1([sdk_entrez_gene_id_u.bin])
    prop{{property}} --> d2([udk_pubmed_id_u.bin])
```

##### Flowchart Legend

- ⬡ **Hexagons:** Interlace folder hierarchy
- ⛁ **Cylinders:** External data providers
- ◯ **Circles:** Interlace Go applications
- 𓉴 **Flags:** Environment variables
- 𓋰 **Ovals:** Interlace Go binary assets
- 𓈙 **Rectangles:** | Raw data files | or || configuration/log files ||

## Contents

1) [Help Page](#1-help-page)
2) [URI Input File](#2-uri-input-file)
3) [Requirements](#3-requirements)
4) [Environment Variables](#4-environment-variables)
5) [Command Line Example](#5-command-line-example)
6) [Log Example](#6-log-example)
7) [Output Example](#7-output-example)

### 1. Help Page

The help page is available with the `-help` option or when no arguments are provided.

```text
Project: interlace-ex
AppName: a31004
Version: beta_2026_08_31_12
Purpose: creates the Homo sapiens Gene entity PubMed Reference asset from the NCBI Gene DB.

Sources:

  1) gene2pubmed.gz

Options:

  -base-concept string
        Base concept ontology term to name the entity folder (default "gene")
  -base-path string
        Agent, event, and entity base path (default "/Users/foglio/interlace-ex")
  -config string
        Config file (optional)
  -data-path string
        Data path where to read and download files (default "/Users/foglio/interlace-ex/data")
  -event-time-tag string
        Event time tag (required) (default "20260904111423")
  -help
        Displays this help page
  -log string
        Logfile path to write stderr & stdout regardless of the -silent flag (default no log file)
  -loud
        Writes additional information if available
  -no-commit
        Disables committing the work
  -no-delete
        Disables deleting the artifact files
  -no-download
        Disables downloading data when necessary
  -silent
        Suppresses all standard output messages
  -uri-inp-file string
        URI list of URL, mailto, telnet, absolute-file-path#key (unique key as URI fragment), etc.
  -workflow-id string
        Workflow ID (default "w31001")
```

### 2. URI Input File

```text
https://ftp.ncbi.nih.gov/gene/DATA/gene2pubmed.gz
```

### 3. Requirements
- Base path from the `-base-path` option or the environmental-variable `ENV_AEE_BASE_PATH`
- Data path from the `-data-path` option or the environmental-variable `ENV_ALT_DATA_PATH`
- URI list containing URLs and full-path file sources.

### 4. Environment Variables

```shell
ENV_AEE_BASE_PATH="${HOME}/interlace-ex"
ENV_ALT_DATA_PATH="${ENV_AEE_BASE_PATH}/data"
```

### 5. Command Line Example

```shell
a31004 -log a31004.log -event-time-tag 20260904150101 -uri-inp-file uri.txt
```

### 6. Log Example

```text
Inf: 2026/09/04 15:38:22.527220 application name: a31004
Inf: 2026/09/04 15:38:22.527362 application version: beta_2026_08_31_12
Inf: 2026/09/04 15:38:22.527382 application module: interlace-ex
Inf: 2026/09/04 15:38:22.527390 application aim: creates the Homo sapiens Gene entity PubMed Reference asset from the NCBI Gene DB.
Inf: 2026/09/04 15:38:22.527398 application start time tag: 20260904153822
Inf: 2026/09/04 15:38:22.527404 listing configuration variables ...
Inf: 2026/09/04 15:38:22.527438 APP_REVISION: 12
Inf: 2026/09/04 15:38:22.527445 BASE_CONCEPT: gene
Inf: 2026/09/04 15:38:22.527450 BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:38:22.527456 DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:38:22.527461 ENV_AEE_BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:38:22.527467 ENV_ALT_DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:38:22.527473 EVENT_TIME_TAG: 20260904150101
Inf: 2026/09/04 15:38:22.527478 HOME: /Users/foglio
Inf: 2026/09/04 15:38:22.527483 LOG_OUTPUT_FILE: a31004.log
Inf: 2026/09/04 15:38:22.527488 URI_INP_FILE: uri.txt
Inf: 2026/09/04 15:38:22.527494 USER_LOGIN_NAME: foglio
Inf: 2026/09/04 15:38:22.527499 USER_NAME: Mario Foglio
Inf: 2026/09/04 15:38:22.527504 WORKFLOW_ID: w31001
Inf: 2026/09/04 15:38:22.527509 setting up the refiner ...
Inf: 2026/09/04 15:38:22.528448 file exists: /Users/foglio/interlace-ex/data/gene2pubmed.gz
Inf: 2026/09/04 15:38:22.528461 creating the gene PubMed reference asset ...
Inf: 2026/09/04 15:38:30.128970 total Entrez Gene PubMed lines read: 82395691
Inf: 2026/09/04 15:38:30.129017 total Entrez Gene PubMed table rows examined: 3326697
Inf: 2026/09/04 15:38:30.129023 total Entrez Gene PubMed tuples created: 3326697
Inf: 2026/09/04 15:38:30.129543 /Users/foglio/interlace-ex/event/gene/w31001/20260904150101/artifact/pubmed_reference_4057533600.tsv.gz size: 48179564
Inf: 2026/09/04 15:38:33.389504 number of warnings caught: 0
Inf: 2026/09/04 15:38:33.389538 elapsed time just before closing: 10.864471667s
Inf: 2026/09/04 15:38:33.389543 end of process: complete
```

### 7. Output Example

```text
interlace-ex/entity/gene/pubmed_reference
└── 20260904150101
    ├── catalog
    ├── enumeration
    └── property
        ├── sdk_entrez_gene_id_u.bin
        └── udk_pubmed_id_u.bin

interlace-ex/event/gene/w31001
└── 20260904150101
    ├── action
    │   ├── 20260904151330_a31001.log
    │   ├── 20260904151952_a31002.log
    │   ├── 20260904152558_a31011.log
    │   ├── 20260904153405_a31003.log
    │   └── 20260904153833_a31004.log
    ├── artifact
    └── function
```

[a31004]: ../../internal/app/a31004
[Gene DB]: https://www.ncbi.nlm.nih.gov/gene/
