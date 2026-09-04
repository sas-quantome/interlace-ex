###### September 4, 2026: d60105 readme v13

# [a31005]

Creates the human **Gene Function** asset from the NCBI [Gene DB].

```mermaid
flowchart LR
    db1[(NCBI Gene DB)] --> generifs_basic.gz
    env1>ENV_AEE_BASE_PATH] --> app((a31005))
    env2>ENV_ALT_DATA_PATH] --> app((a31005))
    uri[[uri.txt]] --> app((a31005))
    generifs_basic.gz --> uri[[uri.txt]]
    app((a31005)) --> ds{{interlace-ex}}
    ds{{interlace-ex}} --> work{{event}}
    ds{{interlace-ex}} --> data{{entity}}
    work{{event}} --> workEntity{{gene}}
    data{{entity}} --> dataEntity{{gene}}
    workEntity{{gene}} --> workAsset{{w31001}}
    dataEntity{{gene}} --> dataAsset{{function}}
    workAsset{{w31001}} --> workRelease{{version}}
    dataAsset{{function}} --> dataRelease{{version}}
    workRelease{{version}} --> action{{action}}
    workRelease{{version}} --> artifact{{artifact}}
    dataRelease{{version}} --> prop{{property}}
    dataRelease{{version}} --> cat{{catalog}}
    action{{action}} --> log[[a31005.log]]
    artifact{{artifact}} --> gene_function_tmp.tsv.gz
    prop{{property}} --> d1([sdk_entrez_gene_id_u.bin])
    prop{{property}} --> d2([udk_pubmed_id_u.bin])
    prop{{property}} --> d3([udk_taxonomy_id_u.bin])
    prop{{property}} --> d4([udk_rif_hash_id_m.bin])
    cat{{catalog}} --> c1([suk_function_hash_id_m.bin])
    cat{{catalog}} --> c2([uuk_function_text_s.bin])
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
AppName: a31005
Version: beta_2026_08_31_13
Purpose: creates the Homo sapiens Gene entity Function asset from NCBI Gene RIF.

Sources:

  1) generifs_basic.gz

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
        Event time tag (required) (default "20260904112330")
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
https://ftp.ncbi.nih.gov/gene/GeneRIF/generifs_basic.gz
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
a31005 -log a31005.log -event-time-tag 20260904150101 -uri-inp-file uri.txt
```

### 6. Log Example

```text
Inf: 2026/09/04 15:41:33.994395 application name: a31005
Inf: 2026/09/04 15:41:33.994484 application version: beta_2026_08_31_13
Inf: 2026/09/04 15:41:33.994500 application module: interlace-ex
Inf: 2026/09/04 15:41:33.994511 application aim: creates the Homo sapiens Gene entity Function asset from NCBI Gene RIF.
Inf: 2026/09/04 15:41:33.994520 application start time tag: 20260904154133
Inf: 2026/09/04 15:41:33.994526 listing configuration variables ...
Inf: 2026/09/04 15:41:33.994716 APP_REVISION: 13
Inf: 2026/09/04 15:41:33.994725 BASE_CONCEPT: gene
Inf: 2026/09/04 15:41:33.994731 BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:41:33.994737 DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:41:33.994743 ENV_AEE_BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:41:33.994748 ENV_ALT_DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:41:33.994754 EVENT_TIME_TAG: 20260904150101
Inf: 2026/09/04 15:41:33.994759 HOME: /Users/foglio
Inf: 2026/09/04 15:41:33.994765 LOG_OUTPUT_FILE: a31005.log
Inf: 2026/09/04 15:41:33.994771 URI_INP_FILE: uri.txt
Inf: 2026/09/04 15:41:33.994776 USER_LOGIN_NAME: foglio
Inf: 2026/09/04 15:41:33.994782 USER_NAME: Mario Foglio
Inf: 2026/09/04 15:41:33.994787 WORKFLOW_ID: w31001
Inf: 2026/09/04 15:41:33.994794 setting up the refiner ...
Inf: 2026/09/04 15:41:33.995776 file exists: /Users/foglio/interlace-ex/data/generifs_basic.gz
Inf: 2026/09/04 15:41:33.995788 creating the gene function asset ...
Inf: 2026/09/04 15:41:38.699742 total Entrez Gene RIF lines read: 1695117
Inf: 2026/09/04 15:41:38.699784 total Entrez Gene RIF table rows examined: 1051428
Inf: 2026/09/04 15:41:38.699790 total Entrez Gene RIF tuples created: 1051428
Inf: 2026/09/04 15:41:40.341332 /Users/foglio/interlace-ex/event/gene/w31001/20260904150101/artifact/function_4172511447.tsv.gz size: 41644681
Inf: 2026/09/04 15:41:42.807483 number of warnings caught: 0
Inf: 2026/09/04 15:41:42.807502 elapsed time just before closing: 8.81505425s
Inf: 2026/09/04 15:41:42.807507 end of process: complete
```

### 7. Output Example

```text
interlace-ex/entity/gene/function
└── 20260904150101
    ├── catalog
    │   ├── suk_function_hash_id_m.bin
    │   └── uuk_function_text_s.bin
    ├── enumeration
    └── property
        ├── sdk_entrez_gene_id_u.bin
        ├── udk_pubmed_id_u.bin
        ├── udk_rif_hash_id_m.bin
        └── udk_taxonomy_id_u.bin

interlace-ex/event/gene/w31001
└── 20260904150101
    ├── action
    │   ├── 20260904151330_a31001.log
    │   ├── 20260904151952_a31002.log
    │   ├── 20260904152558_a31011.log
    │   ├── 20260904153405_a31003.log
    │   ├── 20260904153833_a31004.log
    │   └── 20260904154142_a31005.log
    ├── artifact
    └── function
```

[a31005]: ../../internal/app/a31005
[Gene DB]: https://www.ncbi.nlm.nih.gov/gene/
