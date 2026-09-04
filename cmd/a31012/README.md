###### September 4, 2026: d60112 readme v13

# [a31012]

Creates the human **Gene Pathway** asset from [Reactome].

```mermaid
flowchart LR
    db1[(Reactome)] --> NCBI2Reactome_All_Levels.txt
    env1>ENV_AEE_BASE_PATH] --> app((a31012))
    env2>ENV_ALT_DATA_PATH] --> app((a31012))
    uri[[uri.txt]] --> app((a31012))
    NCBI2Reactome_All_Levels.txt --> uri[[uri.txt]]
    app((a31012)) --> ds{{interlace-ex}}
    ds{{interlace-ex}} --> work{{event}}
    ds{{interlace-ex}} --> data{{entity}}
    work{{event}} --> workEntity{{gene}}
    data{{entity}} --> dataEntity{{gene}}
    workEntity{{gene}} --> workAsset{{w31001}}
    dataEntity{{gene}} --> dataAsset{{pathway}}
    workAsset{{w31001}} --> workRelease{{version}}
    dataAsset{{pathway}} --> dataRelease{{version}}
    workRelease{{version}} --> action{{action}}
    workRelease{{version}} --> artifact{{artifact}}
    dataRelease{{version}} --> prop{{property}}
    dataRelease{{version}} --> cat{{catalog}}
    action{{action}} --> log[[a31012.log]]
    artifact{{artifact}} --> gene_pathway_tmp.tsv.gz
    prop{{property}} --> d1([sdk_entrez_gene_id_u.bin])
    prop{{property}} --> d2([udk_reactome_hsa_id_u.bin])
    prop{{property}} --> d3([udk_pathway_hash_id_m.bin])
    cat{{catalog}} --> c1([suk_pathway_hash_id_m.bin])
    cat{{catalog}} --> c2([uuk_pathway_text_s.bin])
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
AppName: a31012
Version: beta_2026_08_31_14
Purpose: creates the Homo sapiens Gene entity Pathway asset from the Reactome DB.

Sources:

  1) NCBI2Reactome_All_Levels.txt

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
        Event time tag (required) (default "20260904122726")
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
https://reactome.org/download/current/NCBI2Reactome_All_Levels.txt
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
a31012 -log a31012.log -event-time-tag 20260904150101 -uri-inp-file uri.txt
```

### 6. Log Example

```text
Inf: 2026/09/04 15:59:31.136087 application name: a31012
Inf: 2026/09/04 15:59:31.136144 application version: beta_2026_08_31_14
Inf: 2026/09/04 15:59:31.136154 application module: interlace-ex
Inf: 2026/09/04 15:59:31.136160 application aim: creates the Homo sapiens Gene entity Pathway asset from the Reactome DB.
Inf: 2026/09/04 15:59:31.136166 application start time tag: 20260904155931
Inf: 2026/09/04 15:59:31.136169 listing configuration variables ...
Inf: 2026/09/04 15:59:31.136303 APP_REVISION: 14
Inf: 2026/09/04 15:59:31.136309 BASE_CONCEPT: gene
Inf: 2026/09/04 15:59:31.136313 BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:59:31.136317 DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:59:31.136321 ENV_AEE_BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:59:31.136325 ENV_ALT_DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:59:31.136329 EVENT_TIME_TAG: 20260904150101
Inf: 2026/09/04 15:59:31.136332 HOME: /Users/foglio
Inf: 2026/09/04 15:59:31.136336 LOG_OUTPUT_FILE: a31012.log
Inf: 2026/09/04 15:59:31.136340 URI_INP_FILE: uri.txt
Inf: 2026/09/04 15:59:31.136343 USER_LOGIN_NAME: foglio
Inf: 2026/09/04 15:59:31.136346 USER_NAME: Mario Foglio
Inf: 2026/09/04 15:59:31.136350 WORKFLOW_ID: w31001
Inf: 2026/09/04 15:59:31.136353 setting up the refiner ...
Inf: 2026/09/04 15:59:31.137069 file exists: /Users/foglio/interlace-ex/data/NCBI2Reactome_All_Levels.txt
Inf: 2026/09/04 15:59:31.137078 creating the gene pathway asset ...
Inf: 2026/09/04 15:59:31.620151 total Reactome pathways lines read: 788195
Inf: 2026/09/04 15:59:31.620167 total Reactome pathways table rows examined: 788194
Inf: 2026/09/04 15:59:31.620171 total Reactome pathways tuples created: 159339
Inf: 2026/09/04 15:59:31.659707 /Users/foglio/interlace-ex/event/gene/w31001/20260904150101/artifact/pathway_312170101.tsv.gz size: 5289740
Inf: 2026/09/04 15:59:32.046920 number of warnings caught: 0
Inf: 2026/09/04 15:59:32.046950 elapsed time just before closing: 912.055792ms
Inf: 2026/09/04 15:59:32.046955 end of process: complete
```

### 7. Output Example

```text
interlace-ex/entity/gene/pathway
└── 20260904150101
    ├── catalog
    │   ├── suk_pathway_hash_id_m.bin
    │   └── uuk_pathway_text_s.bin
    ├── enumeration
    └── property
        ├── sdk_entrez_gene_id_u.bin
        ├── udk_pathway_hash_id_m.bin
        └── udk_reactome_hsa_id_u.bin

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
    │   ├── 20260904155052_a31008.log
    │   ├── 20260904155318_a31009.log
    │   ├── 20260904155712_a31010.log
    │   └── 20260904155932_a31012.log
    ├── artifact
    └── function
```

[a31012]: ../../internal/app/a31012
[Reactome]: https://reactome.org/
