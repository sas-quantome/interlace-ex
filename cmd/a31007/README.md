###### September 4, 2026: d60107 readme v14

# [a31007]

Creates the human **Gene Medical Genetics** asset from the NCBI [Gene DB] MedGen ID map.

```mermaid
flowchart LR
    db1[(NCBI Gene DB)] --> mim2gene_medgen
    env1>ENV_AEE_BASE_PATH] --> app((a31007))
    env2>ENV_ALT_DATA_PATH] --> app((a31007))
    uri[[uri.txt]] --> app((a31007))
    mim2gene_medgen --> uri[[uri.txt]]
    app((a31007)) --> ds{{interlace-ex}}
    ds{{interlace-ex}} --> work{{event}}
    ds{{interlace-ex}} --> data{{entity}}
    work{{event}} --> workEntity{{gene}}
    data{{entity}} --> dataEntity{{gene}}
    workEntity{{gene}} --> workAsset{{w31001}}
    dataEntity{{gene}} --> dataAsset{{medical_genetics}}
    workAsset{{w31001}} --> workRelease{{version}}
    dataAsset{{medical_genetics}} --> dataRelease{{version}}
    workRelease{{version}} --> action{{action}}
    workRelease{{version}} --> artifact{{artifact}}
    dataRelease{{version}} --> prop{{medical_genetics}}
    action{{action}} --> log[[a31007.log]]
    artifact{{artifact}} --> gene_medical_genetics_tmp.tsv.gz
    prop{{medical_genetics}} --> d1([sdk_entrez_gene_id_i.bin])
    prop{{medical_genetics}} --> d2([udk_omim_id_i.bin])
    prop{{medical_genetics}} --> d3([udk_medgen_cui_s.bin])
    prop{{medical_genetics}} --> d4([udk_bit_flags_u.bin])
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
AppName: a31007
Version: beta_2026_08_31_14
Purpose: creates the Homo sapiens Gene entity Medical Genetics asset from the NCBI GenMed DB.

Sources:

  1) mim2gene_medgen

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
        Event time tag (required) (default "20260904113752")
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
https://ftp.ncbi.nih.gov/gene/DATA/mim2gene_medgen
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
a31007 -log a31007.log -event-time-tag 20260904150101 -uri-inp-file uri.txt
```

### 6. Log Example

```text
Inf: 2026/09/04 15:48:23.883081 application name: a31007
Inf: 2026/09/04 15:48:23.883184 application version: beta_2026_08_31_14
Inf: 2026/09/04 15:48:23.883205 application module: interlace-ex
Inf: 2026/09/04 15:48:23.883213 application aim: creates the Homo sapiens Gene entity Medical Genetics asset from the NCBI GenMed DB.
Inf: 2026/09/04 15:48:23.883220 application start time tag: 20260904154823
Inf: 2026/09/04 15:48:23.883227 listing configuration variables ...
Inf: 2026/09/04 15:48:23.883259 APP_REVISION: 14
Inf: 2026/09/04 15:48:23.883265 BASE_CONCEPT: gene
Inf: 2026/09/04 15:48:23.883272 BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:48:23.883277 DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:48:23.883283 ENV_AEE_BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:48:23.883289 ENV_ALT_DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:48:23.883294 EVENT_TIME_TAG: 20260904150101
Inf: 2026/09/04 15:48:23.883300 HOME: /Users/foglio
Inf: 2026/09/04 15:48:23.883305 LOG_OUTPUT_FILE: a31007.log
Inf: 2026/09/04 15:48:23.883310 URI_INP_FILE: uri.txt
Inf: 2026/09/04 15:48:23.883316 USER_LOGIN_NAME: foglio
Inf: 2026/09/04 15:48:23.883321 USER_NAME: Mario Foglio
Inf: 2026/09/04 15:48:23.883327 WORKFLOW_ID: w31001
Inf: 2026/09/04 15:48:23.883332 setting up the refiner ...
Inf: 2026/09/04 15:48:23.884151 file exists: /Users/foglio/interlace-ex/data/mim2gene_medgen
Inf: 2026/09/04 15:48:23.884164 creating the gene medical genetics asset ...
Inf: 2026/09/04 15:48:23.897199 total Entrez Gene OMIM lines read: 28751
Inf: 2026/09/04 15:48:23.897221 total Entrez Gene OMIM table rows examined: 28749
Inf: 2026/09/04 15:48:23.897226 total Entrez Gene OMIM tuples created: 25426
Inf: 2026/09/04 15:48:23.897367 /Users/foglio/interlace-ex/event/gene/w31001/20260904150101/artifact/medical_genetics_3005166081.tsv.gz size: 462761
Inf: 2026/09/04 15:48:23.981511 number of warnings caught: 0
Inf: 2026/09/04 15:48:23.981532 elapsed time just before closing: 100.71975ms
Inf: 2026/09/04 15:48:23.981537 end of process: complete
```

### 7. Output Example

```text
interlace-ex/entity/gene/medical_genetics
└── 20260904150101
    ├── catalog
    ├── enumeration
    │   └── uuk_bit_flags_cv_s.bin
    └── property
        ├── sdk_entrez_gene_id_u.bin
        ├── udk_bit_flags_u.bin
        ├── udk_medgen_cui_s.bin
        └── udk_omim_id_u.bin

interlace-ex/event/gene/w31001
└── 20260904150101
    ├── action
    │   ├── 20260904151330_a31001.log
    │   ├── 20260904151952_a31002.log
    │   ├── 20260904152558_a31011.log
    │   ├── 20260904153405_a31003.log
    │   ├── 20260904153833_a31004.log
    │   ├── 20260904154142_a31005.log
    │   ├── 20260904154456_a31006.log
    │   └── 20260904154823_a31007.log
    ├── artifact
    └── function
```

[a31007]: ../../internal/app/a31007
[Gene DB]: https://www.ncbi.nlm.nih.gov/gene/
[UniProt]: https://www.uniprot.org/
