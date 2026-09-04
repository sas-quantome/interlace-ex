###### September 4, 2026: d60103 readme v13

# [a31003]

Creates the human **Gene Metadata** asset from the NCBI [Gene DB].

```mermaid
flowchart LR
    db1[(NCBI Gene DB)] --> gene_info.gz
    db1[(NCBI Gene DB)] --> gene_summary.gz
    env1>ENV_AEE_BASE_PATH] --> app((a31003))
    env2>ENV_ALT_DATA_PATH] --> app((a31003))
    uri[[uri.txt]] --> app((a31003))
    gene_info.gz --> uri[[uri.txt]]
    gene_summary.gz --> uri[[uri.txt]]
    app((a31001)) --> ds{{interlace-ex}}
    ds{{interlace-ex}} --> work{{event}}
    ds{{interlace-ex}} --> data{{entity}}
    work{{event}} --> workEntity{{gene}}
    data{{entity}} --> dataEntity{{gene}}
    workEntity{{gene}} --> workAsset{{w31001}}
    dataEntity{{gene}} --> dataAsset{{metadata}}
    workAsset{{w31003}} --> workRelease{{version}}
    dataAsset{{metadata}} --> dataRelease{{version}}
    workRelease{{version}} --> action{{action}}
    workRelease{{version}} --> artifact{{artifact}}
    dataRelease{{version}} --> prop{{property}}
    dataRelease{{version}} --> cat{{catalog}}
    action{{action}} --> log[[a31003.log]]
    artifact{{artifact}} --> gene_metadata_tmp.tsv.gz
    prop{{property}} --> d1([suk_entrez_gene_id_u.bin])
    prop{{property}} --> d2([udk_taxonomy_id_u.bin])
    prop{{property}} --> d3([udk_chromosome_s.bin])
    prop{{property}} --> d4([udk_gene_symbol_s.bin])
    prop{{property}} --> d5([udk_gene_name_s.bin])
    prop{{property}} --> d6([udk_symbol_synonyms_set_s.bin])
    prop{{property}} --> d7([udk_name_synonyms_set_s.bin])
    prop{{property}} --> d8([udk_xdb_id_set_s.bin])
    prop{{property}} --> d9([udk_bit_flags_u.bin])
    prop{{property}} --> d10([udk_metadata_hash_id_m.bin])
    cat{{catalog}} --> c1([suk_metadata_hash_id_m.bin])
    cat{{catalog}} --> c2([uuk_metadata_text_s.bin])
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
AppName: a31003
Version: beta_2026_08_31_13
Purpose: creates the Homo sapiens Gene entity Metadata asset from the NCBI Gene DB.

Sources:

  1) gene_info.gz
  2) gene_summary.gz

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
        Event time tag (required) (default "20260903223133")
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
https://ftp.ncbi.nih.gov/gene/DATA/gene_info.gz
https://ftp.ncbi.nih.gov/gene/DATA/gene_summary.gz
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
a31003 -log a31003.log -event-time-tag 20260904150101 -uri-inp-file uri.txt
```

### 6. Log Example

```text
Inf: 2026/09/04 15:33:27.575657 application name: a31003
Inf: 2026/09/04 15:33:27.575756 application version: beta_2026_08_31_13
Inf: 2026/09/04 15:33:27.575773 application module: interlace-ex
Inf: 2026/09/04 15:33:27.575783 application aim: creates the Homo sapiens Gene entity Metadata asset from the NCBI Gene DB.
Inf: 2026/09/04 15:33:27.575791 application start time tag: 20260904153327
Inf: 2026/09/04 15:33:27.575796 listing configuration variables ...
Inf: 2026/09/04 15:33:27.575827 APP_REVISION: 13
Inf: 2026/09/04 15:33:27.575832 BASE_CONCEPT: gene
Inf: 2026/09/04 15:33:27.575837 BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:33:27.575842 DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:33:27.575847 ENV_AEE_BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:33:27.575852 ENV_ALT_DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:33:27.575857 EVENT_TIME_TAG: 20260904150101
Inf: 2026/09/04 15:33:27.575862 HOME: /Users/foglio
Inf: 2026/09/04 15:33:27.575867 LOG_OUTPUT_FILE: a31003.log
Inf: 2026/09/04 15:33:27.575872 URI_INP_FILE: uri.txt
Inf: 2026/09/04 15:33:27.575877 USER_LOGIN_NAME: foglio
Inf: 2026/09/04 15:33:27.575882 USER_NAME: Mario Foglio
Inf: 2026/09/04 15:33:27.575886 WORKFLOW_ID: w31001
Inf: 2026/09/04 15:33:27.575891 setting up the refiner ...
Inf: 2026/09/04 15:33:27.576682 file exists: /Users/foglio/interlace-ex/data/gene_info.gz
Inf: 2026/09/04 15:33:27.576759 file exists: /Users/foglio/interlace-ex/data/gene_summary.gz
Inf: 2026/09/04 15:33:27.576776 creating the gene metadata asset ...
Inf: 2026/09/04 15:33:28.932473 total Entrez Gene RefSeq summary lines read: 3132497
Inf: 2026/09/04 15:33:28.932524 total Entrez Gene RefSeq summary table rows examined: 151658
Inf: 2026/09/04 15:33:28.932530 total Entrez Gene RefSeq summary tuples created: 151658
Inf: 2026/09/04 15:34:03.388676 total gene info lines read: 71719969
Inf: 2026/09/04 15:34:03.388730 total gene info table rows examined: 193812
Inf: 2026/09/04 15:34:03.388737 total gene info tuples created: 193812
Inf: 2026/09/04 15:34:03.505473 /Users/foglio/interlace-ex/event/gene/w31001/20260904150101/artifact/metadata_1447365567.tsv.gz size: 26781602
Inf: 2026/09/04 15:34:05.734500 number of warnings caught: 0
Inf: 2026/09/04 15:34:05.734519 elapsed time just before closing: 38.161047666s
Inf: 2026/09/04 15:34:05.734527 end of process: complete
```

### 7. Output Example

```text
interlace-ex/entity/gene/metadata
└── 20260904150101
    ├── catalog
    │   ├── suk_metadata_hash_id_m.bin
    │   └── uuk_metadata_text_s.bin
    ├── enumeration
    │   └── uuk_bit_flags_cv_s.bin
    └── property
        ├── suk_entrez_gene_id_u.bin
        ├── udk_bit_flags_u.bin
        ├── udk_chromosome_s.bin
        ├── udk_gene_name_s.bin
        ├── udk_gene_symbol_s.bin
        ├── udk_metadata_hash_id_m.bin
        ├── udk_name_synonyms_set_s.bin
        ├── udk_symbol_synonyms_set_s.bin
        ├── udk_taxonomy_id_u.bin
        └── udk_xdb_id_set_s.bin

interlace-ex/event/gene/w31001
└── 20260904150101
    ├── action
    │   ├── 20260904151330_a31001.log
    │   ├── 20260904151952_a31002.log
    │   ├── 20260904152558_a31011.log
    │   └── 20260904153405_a31003.log
    ├── artifact
    └── function
```

[a31003]: ../../internal/app/a31003
[Gene DB]: https://www.ncbi.nlm.nih.gov/gene/
