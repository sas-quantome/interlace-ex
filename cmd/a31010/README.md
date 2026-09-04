###### September 4, 2026: d60110 readme v15

# [a31010]

Creates the human **Gene Genetic Variation** asset from the NCBI [ClinVar].

```mermaid
flowchart LR
    db1[(NCBI ClinVar)] --> variant_summary.txt.gz
    env1>ENV_AEE_BASE_PATH] --> app((a31010))
    env2>ENV_ALT_DATA_PATH] --> app((a31010))
    uri[[uri.txt]] --> app((a31010))
    variant_summary.txt.gz --> uri[[uri.txt]]
    app((a31010)) --> ds{{interlace-ex}}
    ds{{interlace-ex}} --> work{{event}}
    ds{{interlace-ex}} --> data{{entity}}
    work{{event}} --> workEntity{{gene}}
    data{{entity}} --> dataEntity{{gene}}
    workEntity{{gene}} --> workAsset{{w31001}}
    dataEntity{{gene}} --> dataAsset{{genetic_variation}}
    workAsset{{w31001}} --> workRelease{{version}}
    dataAsset{{genetic_variation}} --> dataRelease{{version}}
    workRelease{{version}} --> action{{action}}
    workRelease{{version}} --> artifact{{artifact}}
    dataRelease{{version}} --> prop{{property}}
    dataRelease{{version}} --> cat{{catalog}}
    action{{action}} --> log[[a31010.log]]
    artifact{{artifact}} --> gene_genetic_variation_tmp.tsv.gz
    prop{{property}} --> d1([sdk_entrez_gene_id_u.bin])
    prop{{property}} --> d2([udk_variation_id_u.bin])
    prop{{property}} --> d3([udk_refsnp_id_u.bin])
    prop{{property}} --> d4([udk_omim_id_u.bin])
    prop{{property}} --> d5([udk_bit_flags_m.bin])
    prop{{property}} --> d6([udk_variation_hash_id_m.bin])
    prop{{property}} --> d7([udk_conical_allele_id_u.bin])
    cat{{catalog}} --> c1([suk_genetic_variation_hash_id_m.bin])
    cat{{catalog}} --> c2([uuk_genetic_variation_text_s.bin])
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
AppName: a31010
Version: beta_2026_08_31_14
Purpose: creates the Homo sapiens Gene entity Genetic Variation asset from the NCBI ClinVar DB.

Sources:

  1) variant_summary.txt.gz

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
        Event time tag (required) (default "20260904120734")
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
https://ftp.ncbi.nlm.nih.gov/pub/clinvar/tab_delimited/variant_summary.txt.gz
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
a31010 -log a31010.log -event-time-tag 20260904150101 -uri-inp-file uri.txt
```

### 6. Log Example

```text
Inf: 2026/09/04 15:55:30.117072 application name: a31010
Inf: 2026/09/04 15:55:30.117137 application version: beta_2026_08_31_14
Inf: 2026/09/04 15:55:30.117147 application module: interlace-ex
Inf: 2026/09/04 15:55:30.117152 application aim: creates the Homo sapiens Gene entity Genetic Variation asset from the NCBI ClinVar DB.
Inf: 2026/09/04 15:55:30.117173 application start time tag: 20260904155530
Inf: 2026/09/04 15:55:30.117178 listing configuration variables ...
Inf: 2026/09/04 15:55:30.117183 APP_REVISION: 14
Inf: 2026/09/04 15:55:30.117187 BASE_CONCEPT: gene
Inf: 2026/09/04 15:55:30.117191 BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:55:30.117194 DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:55:30.117198 ENV_AEE_BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 15:55:30.117201 ENV_ALT_DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 15:55:30.117205 EVENT_TIME_TAG: 20260904150101
Inf: 2026/09/04 15:55:30.117208 HOME: /Users/foglio
Inf: 2026/09/04 15:55:30.117212 LOG_OUTPUT_FILE: a31010.log
Inf: 2026/09/04 15:55:30.117215 URI_INP_FILE: uri.txt
Inf: 2026/09/04 15:55:30.117219 USER_LOGIN_NAME: foglio
Inf: 2026/09/04 15:55:30.117222 USER_NAME: Mario Foglio
Inf: 2026/09/04 15:55:30.117226 WORKFLOW_ID: w31001
Inf: 2026/09/04 15:55:30.117229 setting up the refiner ...
Inf: 2026/09/04 15:55:30.117719 file exists: /Users/foglio/interlace-ex/data/variant_summary.txt.gz
Inf: 2026/09/04 15:55:30.117732 creating the gene genetic variation asset ...
Inf: 2026/09/04 15:56:42.091289 total ClinVar summary records without GeneID: 65501
Inf: 2026/09/04 15:56:42.091312 total ClinVar summary lines read: 9048964
Inf: 2026/09/04 15:56:42.091316 total ClinVar summary table rows examined: 9048962
Inf: 2026/09/04 15:56:42.091320 total ClinVar summary tuples created: 11942689
Inf: 2026/09/04 15:56:44.134545 /Users/foglio/interlace-ex/event/gene/w31001/20260904150101/artifact/genetic_variation_224914543.tsv.gz size: 801196217
Inf: 2026/09/04 15:57:12.119752 number of warnings caught: 0
Inf: 2026/09/04 15:57:12.123684 elapsed time just before closing: 1m42.008198916s
Inf: 2026/09/04 15:57:12.123691 end of process: complete
```

### 7. Output Example

```text
interlace-ex/entity/gene/genetic_variation
└── 20260904150101
    ├── catalog
    │   ├── suk_genetic_variation_hash_id_m.bin
    │   └── uuk_genetic_variation_text_s.bin
    ├── enumeration
    │   └── uuk_bit_flags_cv_s.bin
    └── property
        ├── sdk_entrez_gene_id_u.bin
        ├── udk_bit_flags_m.bin
        ├── udk_conical_allele_id_u.bin
        ├── udk_omim_id_u.bin
        ├── udk_refsnp_id_u.bin
        ├── udk_variation_hash_id_m.bin
        └── udk_variation_id_u.bin

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
    │   └── 20260904155712_a31010.log
    ├── artifact
    └── function
```

[a31010]: ../../internal/app/a31010
[ClinVar]: https://www.ncbi.nlm.nih.gov/clinvar/
