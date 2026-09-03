###### September 3, 2026: interlace-ex cmd readme v2

# CMD

15 **interlace-go** command-line applications build the **interlace-ex** Gene Entity datastack in less than 5 minutes using a cap of four CPUs max in a MacBook Pro M1.

Reducing storage size by 10X or more yields organization-wide infrastructure cost savings.

```mermaid
flowchart TB
    sz1>57 GB of uncompressed text data] --- inp{{Raw Data}}
    inp{{Raw Data}} --> a01((a31001))
    a01((a31001)) --> a02((a31002))
    a02((a31002)) --> a11((a31011))
    a02((a31002)) --> a14((a31014))
    inp{{Raw Data}} --> a03((a31003))
    inp{{Raw Data}} --> a04((a31004))
    inp{{Raw Data}} --> a05((a31005))
    inp{{Raw Data}} --> a06((a31006))
    inp{{Raw Data}} --> a07((a31007))
    inp{{Raw Data}} --> a08((a31008))
    inp{{Raw Data}} --> a09((a31009))
    inp{{Raw Data}} --> a10((a31010))
    inp{{Raw Data}} --> a12((a31012))
    inp{{Raw Data}} --> a13((a31013))
    a11((a31011)) --> a15((a31015))
    a02((a31002)) --> db[(Gene Entity)]
    a03((a31003)) --> db[(Gene Entity)]
    a04((a31004)) --> db[(Gene Entity)]
    a05((a31005)) --> db[(Gene Entity)]
    a06((a31006)) --> db[(Gene Entity)]
    a07((a31007)) --> db[(Gene Entity)]
    a08((a31008)) --> db[(Gene Entity)]
    a09((a31009)) --> db[(Gene Entity)]
    a10((a31010)) --> db[(Gene Entity)]
    a12((a31012)) --> db[(Gene Entity)]
    a13((a31013)) --> db[(Gene Entity)]
    a14((a31014)) --> db[(Gene Entity)]
    a15((a31015)) --> db[(Gene Entity)]
    db[(Refined Data)] --- sz2>1.4 GB of uncompressed Interlace Go binary files]
 ```

##### Flowchart Legend

- ⬡ **Hexagon:** Diverse omics database flat-files
- ◯ **Circles:** Interlace Go applications
- ⛁ **Cylinder:** Refined binary Gene assets
- 𓉴 **Flags:** Disk space 

Notice the difference in disk space: 

- the input is 57 GB of uncompressed text data, and 
- the output is 1.4 GB of uncompressed and refined binary data assets. 

Part of this reduction comes from excluding non-human genes, yet most input files in the **interlace-ex** are downloaded exclusively for human genes. 

Text is encoded as integers using bit flags or controlled-vocabulary enumerators, and catalogs store text strings once, greatly reducing the flat-file bloat.

At the end of the data refinement process, the complete human gene dataset fits in memory, making information retrieval very fast.

## Contents

| AppID    | Dependencies | Data Assets & Sources                                                          |
|----------|--------------|--------------------------------------------------------------------------------|
| [a31001] |              | `Creates` the **Gene Product** Gob assets from NCBI Gene DB and UniProt.       |
| [a31002] | [a31001]     | `Creates` the **Gene Synonym** Gob assets from the **Gene Product**.           |
| [a31003] |              | `Creates` the **Gene Metadata** Gob assets from NCBI Gene DB.                  |
| [a31004] |              | `Creates` the **Gene PubMed Reference** Gob assets from NCBI Gene DB.          |
| [a31005] |              | `Creates` the **Gene Function** Gob assets from NCBI Gene DB RIF.              |
| [a31006] |              | `Creates` the **Gene Ontology** Gob assets from NCBI Gene DB GO.               |
| [a31007] |              | `Creates` the **Gene Medical Genetics** Gob assets from NCBI Gene DB MedGen.   |
| [a31008] |              | `Creates` the **Gene Ortholog** Gob assets from NCBI Gene DB.                  |
| [a31009] |              | `Creates` the **Gene Molecular Interaction** Gob assets from NCBI Gene DB RIF. |
| [a31010] |              | `Creates` the **Gene Genetic Variation** Gob assets from NCBI ClinVar.         |
| [a31011] | [a31002]     | `Creates` the **Gene Genomic Variation** Gob assets from EBI GWAS Catalog.     |
| [a31012] |              | `Creates` the **Gene Pathway** Gob assets from Reactome.                       |
| [a31013] |              | `Creates` the **Gene Chemical Reaction** Gob assets from Reactome.             |
| [a31014] | [a31002]     | `Creates` the **Gene Protein Interaction** Gob assets from STRING DB.          |
| [a31015] | [a31011]     | `Inserts` the GWAS Ancestry data to the **Gene Genomic Variation** asset.      |

### Economic Impact & Data Gravity

Reducing storage size by 10X or more yields organization-wide infrastructure cost savings and overcomes data gravity.

Data gravity refers to how large data attracts applications, services, and other data—the larger the mass, the stronger its pull.

Coined by Dave McCrory in 2010, the term draws on physics: large data, like massive objects, exerts a costly, complex gravitational pull on workloads.

In genomics and bioinformatics, this is very tangible. Reference genomes, variant databases (gnomAD, ClinVar, dbSNP), and large sequencing cohorts exert enormous gravity — which is why tools like AWS HealthOmics or Terra on GCP co-locate compute with the data rather than streaming it out.

The pipelines direct to the data, not the other way around.

This effect means that data gravity often leads to greater centralization and lock-in.


[a31001]: ./a31001/README.md
[a31002]: ./a31002/README.md
[a31003]: ./a31003/README.md
[a31004]: ./a31004/README.md
[a31005]: ./a31005/README.md
[a31006]: ./a31006/README.md
[a31007]: ./a31007/README.md
[a31008]: ./a31008/README.md
[a31009]: ./a31009/README.md
[a31010]: ./a31010/README.md
[a31011]: ./a31011/README.md
[a31012]: ./a31012/README.md
[a31013]: ./a31013/README.md
[a31014]: ./a31014/README.md
[a31015]: ./a31015/README.md
