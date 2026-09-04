###### September 4, 2026: d60106 readme v13

# [a31006]

Creates the human **Gene Ontology** asset from the NCBI [Gene DB] GO annotation.

```mermaid
flowchart LR
    db1[(NCBI Gene DB)] --> gene2go.gz
    env1>ENV_AEE_BASE_PATH] --> app((a31006))
    env2>ENV_ALT_DATA_PATH] --> app((a31006))
    uri[[uri.txt]] --> app((a31006))
    gene2go.gz --> uri[[uri.txt]]
    app((a31006)) --> ds{{interlace-ex}}
    ds{{interlace-ex}} --> work{{event}}
    ds{{interlace-ex}} --> data{{entity}}
    work{{event}} --> workEntity{{gene}}
    data{{entity}} --> dataEntity{{gene}}
    workEntity{{gene}} --> workAsset{{w31001}}
    dataEntity{{gene}} --> dataAsset{{ontology}}
    workAsset{{w31001}} --> workRelease{{version}}
    dataAsset{{ontology}} --> dataRelease{{version}}
    workRelease{{version}} --> action{{action}}
    workRelease{{version}} --> artifact{{artifact}}
    dataRelease{{version}} --> prop{{property}}
    dataRelease{{version}} --> cat{{catalog}}
    action{{action}} --> log[[a31006.log]]
    artifact{{artifact}} --> gene_ontology_tmp.tsv.gz
    prop{{property}} --> d1([sdk_entrez_gene_id_u.bin])
    prop{{property}} --> d2([udk_pubmed_id_u.bin])
    prop{{property}} --> d3([udk_go_id_u.bin])
    prop{{property}} --> d4([udk_bit_flags_m.bin])
    prop{{property}} --> d5([udk_go_hash_id_m.bin])
    cat{{catalog}} --> c1([suk_ontology_hash_id_m.bin])
    cat{{catalog}} --> c2([uuk_ontology_text_s.bin])
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
AppName: a31006
Version: beta_2026_08_31_13
Purpose: creates the Homo sapiens Gene entity Gene Ontology asset from the NCBI Gene DB.

Sources:

  1) gene2go.gz

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
        Event time tag (required) (default "20260904113216")
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
https://ftp.ncbi.nih.gov/gene/DATA/gene2go.gz
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
a31006 -log a31006.log -event-time-tag 20260904150101 -uri-inp-file uri.txt
```

### 6. Log Example

```text
Inf: 2026/09/04 15:44:21.245545 application name: a31006
Inf: 2026/09/04 15:44:21.245701 application version: beta_2026_08_31_13
Inf: 2026/09/04 15:44:21.245719 application module: interlace-ex
Inf: 2026/09/04 15:44:21.245729 application aim: creates the Homo sapiens Gene entity Gene Ontology asset from the NCBI Gene DB.
Inf: 2026/09/04 15:44:21.245740 application start time tag: 20260904154421
Inf: 2026/09/04 15:44:21.245747 listing configuration variables ...
Inf: 2026/09/04 15:44:21.245781 APP_REVISION: 13
Inf: 2026/09/04 15:44:21.245787 BASE_CONCEPT: gene
Inf: 2026/09/04 15:44:21.245793 BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:44:21.245799 DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:44:21.245805 ENV_AEE_BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:44:21.245810 ENV_ALT_DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:44:21.245816 EVENT_TIME_TAG: 20260904150101
Inf: 2026/09/04 15:44:21.245821 HOME: /Users/foglio
Inf: 2026/09/04 15:44:21.245827 LOG_OUTPUT_FILE: a31006.log
Inf: 2026/09/04 15:44:21.245832 URI_INP_FILE: uri.txt
Inf: 2026/09/04 15:44:21.245838 USER_LOGIN_NAME: foglio
Inf: 2026/09/04 15:44:21.245843 USER_NAME: Mario Foglio
Inf: 2026/09/04 15:44:21.245848 WORKFLOW_ID: w31001
Inf: 2026/09/04 15:44:21.245878 setting up the refiner ...
Inf: 2026/09/04 15:44:21.246949 file exists: /Users/foglio/interlace-ex/data/gene2go.gz
Inf: 2026/09/04 15:44:21.246963 creating the gene ontology asset ...
Inf: 2026/09/04 15:44:55.326887 total Entrez Gene GO lines read: 124473676
Inf: 2026/09/04 15:44:55.326943 total Entrez Gene GO table rows examined: 446675
Inf: 2026/09/04 15:44:55.326949 total Entrez Gene GO tuples created: 300081
Inf: 2026/09/04 15:44:55.455937 /Users/foglio/interlace-ex/event/gene/w31001/20260904150101/artifact/ontology_1205120959.tsv.gz size: 16577392
Inf: 2026/09/04 15:44:56.141002 number of warnings caught: 0
Inf: 2026/09/04 15:44:56.141024 elapsed time just before closing: 34.897659666s
Inf: 2026/09/04 15:44:56.141030 end of process: complete
```

### 7. Output Example

```text
interlace-ex/entity/gene/ontology
└── 20260904150101
    ├── catalog
    │   ├── suk_ontology_hash_id_m.bin
    │   └── uuk_ontology_text_s.bin
    ├── enumeration
    │   └── uuk_bit_flags_cv_s.bin
    └── property
        ├── sdk_entrez_gene_id_u.bin
        ├── udk_bit_flags_m.bin
        ├── udk_go_hash_id_m.bin
        ├── udk_go_id_u.bin
        └── udk_pubmed_id_u.bin

interlace-ex/event/gene/w31001
└── 20260904150101
    ├── action
    │   ├── 20260904151330_a31001.log
    │   ├── 20260904151952_a31002.log
    │   ├── 20260904152558_a31011.log
    │   ├── 20260904153405_a31003.log
    │   ├── 20260904153833_a31004.log
    │   ├── 20260904154142_a31005.log
    │   └── 20260904154456_a31006.log
    ├── artifact
    └── function
```

[a31006]: ../../internal/app/a31006
[Gene DB]: https://www.ncbi.nlm.nih.gov/gene/
