###### September 4, 2026: d60108 readme v13

# [a31008]

Creates the human **Gene Ortholog** asset from the NCBI [Gene DB] orthologs.

```mermaid
flowchart LR
    db1[(NCBI Gene DB)] --> gene_orthologs.gz
    env1>ENV_AEE_BASE_PATH] --> app((a31008))
    env2>ENV_ALT_DATA_PATH] --> app((a31008))
    uri[[uri.txt]] --> app((a31008))
    gene_orthologs.gz --> uri[[uri.txt]]
    app((a31008)) --> ds{{interlace-ex}}
    ds{{interlace-ex}} --> work{{event}}
    ds{{interlace-ex}} --> data{{entity}}
    work{{event}} --> workEntity{{gene}}
    data{{entity}} --> dataEntity{{gene}}
    workEntity{{gene}} --> workAsset{{w31001}}
    dataEntity{{gene}} --> dataAsset{{ortholog}}
    workAsset{{w31001}} --> workRelease{{version}}
    dataAsset{{ortholog}} --> dataRelease{{version}}
    workRelease{{version}} --> action{{action}}
    workRelease{{version}} --> artifact{{artifact}}
    dataRelease{{version}} --> prop{{property}}
    action{{action}} --> log[[a31008.log]]
    artifact{{artifact}} --> gene_ortholog_tmp.tsv.gz
    prop{{property}} --> d1([sdk_entrez_gene_id_u.bin])
    prop{{property}} --> d2([udk_taxonomy_id_u.bin])
    prop{{property}} --> d3([udk_ortholog_taxonomy_id_u.bin])
    prop{{property}} --> d4([udk_ortholog_entrez_gene_id_u.bin])
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
AppName: a31008
Version: beta_2026_08_31_12
Purpose: creates the Homo sapiens Gene entity Ortholog asset from the NCBI Gene DB.

Sources:

  1) gene_orthologs.gz

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
        Event time tag (required) (default "20260904114315")
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
https://ftp.ncbi.nih.gov/gene/DATA/gene_orthologs.gz
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
a31008 -log a31008.log -event-time-tag 20260904150101 -uri-inp-file uri.txt
```

### 6. Log Example

```text
Inf: 2026/09/04 15:50:29.714732 application name: a31008
Inf: 2026/09/04 15:50:29.714840 application version: beta_2026_08_31_12
Inf: 2026/09/04 15:50:29.714858 application module: interlace-ex
Inf: 2026/09/04 15:50:29.714870 application aim: creates the Homo sapiens Gene entity Ortholog asset from the NCBI Gene DB.
Inf: 2026/09/04 15:50:29.714880 application start time tag: 20260904155029
Inf: 2026/09/04 15:50:29.714885 listing configuration variables ...
Inf: 2026/09/04 15:50:29.714918 APP_REVISION: 12
Inf: 2026/09/04 15:50:29.714924 BASE_CONCEPT: gene
Inf: 2026/09/04 15:50:29.714931 BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:50:29.714936 DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:50:29.714942 ENV_AEE_BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:50:29.714948 ENV_ALT_DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:50:29.714953 EVENT_TIME_TAG: 20260904150101
Inf: 2026/09/04 15:50:29.714958 HOME: /Users/foglio
Inf: 2026/09/04 15:50:29.714964 LOG_OUTPUT_FILE: a31008.log
Inf: 2026/09/04 15:50:29.714969 URI_INP_FILE: uri.txt
Inf: 2026/09/04 15:50:29.714974 USER_LOGIN_NAME: foglio
Inf: 2026/09/04 15:50:29.714979 USER_NAME: Mario Foglio
Inf: 2026/09/04 15:50:29.714984 WORKFLOW_ID: w31001
Inf: 2026/09/04 15:50:29.714990 setting up the refiner ...
Inf: 2026/09/04 15:50:29.715930 file exists: /Users/foglio/interlace-ex/data/gene_orthologs.gz
Inf: 2026/09/04 15:50:29.715943 creating the gene ortholog asset ...
Inf: 2026/09/04 15:50:36.429571 total Entrez Gene Orthologs lines read: 18756043
Inf: 2026/09/04 15:50:36.429620 total Entrez Gene Orthologs table rows examined: 11267602
Inf: 2026/09/04 15:50:36.429629 total Entrez Gene Orthologs tuples created: 11267602
Inf: 2026/09/04 15:50:36.430316 /Users/foglio/interlace-ex/event/gene/w31001/20260904150101/artifact/ortholog_3146324453.tsv.gz size: 304068890
Inf: 2026/09/04 15:50:52.925678 number of warnings caught: 0
Inf: 2026/09/04 15:50:52.930659 elapsed time just before closing: 23.218216833s
Inf: 2026/09/04 15:50:52.930667 end of process: complete
```

### 7. Output Example

```text
interlace-ex/entity/gene/ortholog
└── 20260904150101
    ├── catalog
    ├── enumeration
    └── property
        ├── sdk_entrez_gene_id_u.bin
        ├── udk_ortholog_entrez_gene_id_u.bin
        ├── udk_ortholog_taxonomy_id_u.bin
        └── udk_taxonomy_id_u.bin

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
    │   ├── 20260904154823_a31007.log
    │   └── 20260904155052_a31008.log
    ├── artifact
    └── function
```

[a31008]: ../../internal/app/a31008
[Gene DB]: https://www.ncbi.nlm.nih.gov/gene/
