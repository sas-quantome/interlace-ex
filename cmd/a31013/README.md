###### September 4, 2026: d60113 readme v13

# [a31013]

Creates the human **Gene Chemical Reaction** asset from [Reactome].

```mermaid
flowchart LR
    db1[(Reactome)] --> NCBI2Reactome_PE_Reactions.txt
    db1[(Reactome)] --> ReactionPMIDS.txt
    env1>ENV_AEE_BASE_PATH] --> app((a31013))
    env2>ENV_ALT_DATA_PATH] --> app((a31013))
    uri[[uri.txt]] --> app((a31013))
    NCBI2Reactome_PE_Reactions.txt --> uri[[uri.txt]]
    ReactionPMIDS.txt --> uri[[uri.txt]]
    app((a31013)) --> ds{{interlace-ex}}
    ds{{interlace-ex}} --> work{{event}}
    ds{{interlace-ex}} --> data{{entity}}
    work{{event}} --> workEntity{{gene}}
    data{{entity}} --> dataEntity{{gene}}
    workEntity{{gene}} --> workAsset{{w31001}}
    dataEntity{{gene}} --> dataAsset{{chemical_reaction}}
    workAsset{{w31001}} --> workRelease{{version}}
    dataAsset{{chemical_reaction}} --> dataRelease{{version}}
    workRelease{{version}} --> action{{action}}
    workRelease{{version}} --> artifact{{artifact}}
    dataRelease{{version}} --> prop{{property}}
    dataRelease{{version}} --> cat{{catalog}}
    action{{action}} --> log[[a31013.log]]
    artifact{{artifact}} --> gene_chemical_reaction_tmp.tsv.gz
    prop{{property}} --> d1([sdk_entrez_gene_id_u.bin])
    prop{{property}} --> d2([udk_reactome_hsa_id_u.bin])
    prop{{property}} --> d3([udk_pubmed_id_u.bin])
    prop{{property}} --> d4([udk_entity_cv_d.bin])
    prop{{property}} --> d5([udk_reaction_hash_id_m.bin])
    cat{{catalog}} --> c1([suk_chemical_reaction_hash_id_m.bin])
    cat{{catalog}} --> c2([uuk_chemical_reaction_text_s.bin])
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
AppName: a31013
Version: beta_2026_08_31_14
Purpose: creates the Homo sapiens Gene entity Chemical Reaction asset from the Reactome DB.

Sources:

  1) NCBI2Reactome_PE_Reactions.txt
  2) ReactionPMIDS.txt

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
        Event time tag (required) (default "20260904123126")
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
https://reactome.org/download/current/NCBI2Reactome_PE_Reactions.txt
https://reactome.org/download/current/ReactionPMIDS.txt
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
a31013 -log a31013.log -event-time-tag 20260904150101 -uri-inp-file uri.txt
```

### 6. Log Example

```text
Inf: 2026/09/04 16:01:44.209150 application name: a31013
Inf: 2026/09/04 16:01:44.209202 application version: beta_2026_08_31_14
Inf: 2026/09/04 16:01:44.209212 application module: interlace-ex
Inf: 2026/09/04 16:01:44.209219 application aim: creates the Homo sapiens Gene entity Chemical Reaction asset from the Reactome DB.
Inf: 2026/09/04 16:01:44.209226 application start time tag: 20260904160144
Inf: 2026/09/04 16:01:44.209232 listing configuration variables ...
Inf: 2026/09/04 16:01:44.209262 APP_REVISION: 14
Inf: 2026/09/04 16:01:44.209268 BASE_CONCEPT: gene
Inf: 2026/09/04 16:01:44.209274 BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 16:01:44.209280 DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 16:01:44.209286 ENV_AEE_BASE_PATH: /Users/foglio/interlace-ex
Inf: 2026/09/04 16:01:44.209292 ENV_ALT_DATA_PATH: /Users/foglio/interlace-ex/data
Inf: 2026/09/04 16:01:44.209298 EVENT_TIME_TAG: 20260904150101
Inf: 2026/09/04 16:01:44.209303 HOME: /Users/foglio
Inf: 2026/09/04 16:01:44.209309 LOG_OUTPUT_FILE: a31013.log
Inf: 2026/09/04 16:01:44.209314 URI_INP_FILE: uri.txt
Inf: 2026/09/04 16:01:44.209320 USER_LOGIN_NAME: foglio
Inf: 2026/09/04 16:01:44.209325 USER_NAME: Mario Foglio
Inf: 2026/09/04 16:01:44.209330 WORKFLOW_ID: w31001
Inf: 2026/09/04 16:01:44.209335 setting up the refiner ...
Inf: 2026/09/04 16:01:44.210097 file exists: /Users/foglio/interlace-ex/data/NCBI2Reactome_PE_Reactions.txt
Inf: 2026/09/04 16:01:44.210151 file exists: /Users/foglio/interlace-ex/data/ReactionPMIDS.txt
Inf: 2026/09/04 16:01:44.210166 creating the gene chemical reaction asset ...
Inf: 2026/09/04 16:01:44.217368 total Reactome PubMed reference lines read: 50196
Inf: 2026/09/04 16:01:44.217401 total Reactome PubMed reference table rows examined: 50195
Inf: 2026/09/04 16:01:44.217409 total Reactome PubMed reference tuples created: 43581
Inf: 2026/09/04 16:01:44.861931 total Reactome reactions lines read: 814425
Inf: 2026/09/04 16:01:44.861949 total Reactome reactions table rows examined: 814424
Inf: 2026/09/04 16:01:44.861954 total Reactome reactions tuples created: 184295
Inf: 2026/09/04 16:01:44.936003 /Users/foglio/interlace-ex/event/gene/w31001/20260904150101/artifact/chemical_reaction_1076806535.tsv.gz size: 8239416
Inf: 2026/09/04 16:01:45.545595 number of warnings caught: 0
Inf: 2026/09/04 16:01:45.545669 elapsed time just before closing: 1.337878417s
Inf: 2026/09/04 16:01:45.545687 end of process: complete
```

### 7. Output Example

```text
interlace-ex/entity/gene/chemical_reaction
└── 20260904150101
    ├── catalog
    │   ├── suk_chemical_reaction_hash_id_m.bin
    │   └── uuk_chemical_reaction_text_s.bin
    ├── enumeration
    │   └── uuk_reaction_entity_cv_s.bin
    └── property
        ├── sdk_entrez_gene_id_u.bin
        ├── udk_pubmed_id_u.bin
        ├── udk_reaction_entity_cv_d.bin
        ├── udk_reaction_hash_id_m.bin
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
    │   ├── 20260904155932_a31012.log
    │   └── 20260904160145_a31013.log
    ├── artifact
    └── function
```

[a31013]: ../../internal/app/a31013
[Reactome]: https://reactome.org/
