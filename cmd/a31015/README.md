###### September 4, 2026: d60115 readme v10

# [a31015]

Inserts the GWAS Ancestry data to the **Genomic Variation** asset from the EBI [GWAS Ancestry].

```mermaid
flowchart LR
    db1[(EBI GWAS)] --> gwas-catalog-ancestry.tsv.gz
    env1>ENV_AEE_BASE_PATH] --> app((a31015))
    env2>ENV_ALT_DATA_PATH] --> app((a31015))
    uri[[uri.txt]] --> app((a31015))
    gwas-catalog-ancestry.tsv.gz --> uri[[uri.txt]]
    in1([udk_pubmed_id_u.bin]) --> app((a31015))
    app((a31015)) --> ds{{interlace-ex}}
    ds{{interlace-ex}} --> work{{event}}
    ds{{interlace-ex}} --> data{{entity}}
    work{{event}} --> workEntity{{gene}}
    data{{entity}} --> dataEntity{{gene}}
    workEntity{{gene}} --> workAsset{{w31001}}
    dataEntity{{gene}} --> dataAsset{{genomic_variation}}
    workAsset{{w31001}} --> workRelease{{version}}
    dataAsset{{genomic_variation}} --> dataRelease{{version}}
    workRelease{{version}} --> action{{action}}
    workRelease{{version}} --> artifact{{artifact}}
    dataRelease{{version}} --> prop{{property}}
    action{{action}} --> log[[a31015.log]]
    artifact{{artifact}} --> gene_genomic_variation_tmp.tsv.gz
    prop{{property}} --> d1([udk_bit_flags_u.bin])
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
AppName: a31015
Version: beta_2026_08_31_10
Purpose: inserts the GWAS Ancestry data as bit flags to the Gene entity Genomic Variation asset.

Sources:

  1) GWAS_ANCESTRY
  2) udk_pubmed_id_u.bin

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
        Event time tag (required) (default "20260904124204")
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
# Non-persistent URI character sequences may not have a constant filename.
# Files with non-persistent URL are manually downloaded and declared as absolute file paths.
# Because their file names are unknown, the application looks for the unique URI fragment e.g., #GWAS_ANCESTRY appended at the end.
# https://ftp.ebi.ac.uk/pub/databases/gwas/releases/latest/gwas-catalog-ancestry.tsv
file:///Users/foglio/interlace-ex/data/gwas-catalog-ancestry.tsv.gz#GWAS_ANCESTRY
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
a31015 -log a31015.log -event-time-tag 20260904150101 -uri-inp-file uri.txt
```

### 6. Log Example

```text
Inf: 2026/09/04 16:10:27.124510 application name: a31015
Inf: 2026/09/04 16:10:27.124621 application version: beta_2026_08_31_10
Inf: 2026/09/04 16:10:27.124640 application module: interlace-ex
Inf: 2026/09/04 16:10:27.124670 application aim: inserts the GWAS Ancestry data as bit flags to the Gene entity Genomic Variation asset.
Inf: 2026/09/04 16:10:27.124678 application start time tag: 20260904161027
Inf: 2026/09/04 16:10:27.124684 listing configuration variables ...
Inf: 2026/09/04 16:10:27.124718 APP_REVISION: 10
Inf: 2026/09/04 16:10:27.124724 BASE_CONCEPT: gene
Inf: 2026/09/04 16:10:27.124731 BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 16:10:27.124736 DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 16:10:27.124742 ENV_AEE_BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 16:10:27.124748 ENV_ALT_DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 16:10:27.124754 EVENT_TIME_TAG: 20260904150101
Inf: 2026/09/04 16:10:27.124759 HOME: /Users/foglio
Inf: 2026/09/04 16:10:27.124764 LOG_OUTPUT_FILE: a31015.log
Inf: 2026/09/04 16:10:27.124770 URI_INP_FILE: uri.txt
Inf: 2026/09/04 16:10:27.124775 USER_LOGIN_NAME: foglio
Inf: 2026/09/04 16:10:27.124780 USER_NAME: Mario Foglio
Inf: 2026/09/04 16:10:27.124786 WORKFLOW_ID: w31001
Inf: 2026/09/04 16:10:27.124791 setting up the refiner ...
Inf: 2026/09/04 16:10:27.125083 configured file[GWAS_ANCESTRY]=/Users/foglio/InterlaceEx/data/gwas-catalog-ancestry.tsv.gz
Inf: 2026/09/04 16:10:27.125106 integrating the gene genomic variation GWAS ancestry data ...
Inf: 2026/09/04 16:10:27.249966 total GWAS Ancestry lines read: 278864
Inf: 2026/09/04 16:10:27.249992 total GWAS Ancestry table rows examined: 278862
Inf: 2026/09/04 16:10:27.249997 total GWAS Ancestry tuples created: 795158
Inf: 2026/09/04 16:10:27.250363 number of warnings caught: 0
Inf: 2026/09/04 16:10:27.250371 elapsed time just before closing: 127.44525ms
Inf: 2026/09/04 16:10:27.250374 end of process: complete
```

### 7. Output Example

```text
interlace-ex/entity/gene/genomic_variation
└── 20260904150101
    ├── catalog
    │   ├── suk_genomic_variation_hash_id_m.bin
    │   └── uuk_genomic_variation_text_s.bin
    ├── enumeration
    │   └── uuk_bit_flags_cv_s.bin
    └── property
        ├── sdk_entrez_gene_id_u.bin
        ├── udk_bit_flags_u.bin
        ├── udk_gwas_hash_id_m.bin
        ├── udk_ontology_id_s.bin
        ├── udk_pubmed_id_u.bin
        └── udk_refsnp_id_u.bin

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
    │   ├── 20260904155932_a31012.log
    │   ├── 20260904160145_a31013.log
    │   ├── 20260904160604_a31014.log
    │   └── 20260904161027_a31015.log
    ├── artifact
    └── function
```

[a31015]: ../../internal/app/a31015
[GWAS Ancestry]: https://www.ebi.ac.uk/gwas/docs/ancestry-data
