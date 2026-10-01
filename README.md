###### October 1, 2026: interlace-ex readme v24

# interlace-ex

Welcome to **interlace-ex**! Learn how to use the **interlace-go** library with the **interlace-ex** code examples to build the Gene data stack. This open-source educational repository guides you through consolidating human gene data, including products, interactants, genetic variants, and diseases, as a hands-on data refinement exercise.

![interlace-ex](assets/interlace-ex.png)

## Contents

1) [Introduction](#1-introduction)
2) [Data Providers](#2-data-providers)
3) [Data Creators and Peer-reviewed Journal References](#3-data-creators-and-peer-reviewed-journal-references)
4) [Clone, Compile, and Install](#4-clone-compile-and-install)
5) [Build the *Gene* Datastack](#5-build-the-gene-datastack)
6) [Query the *Gene* Datastack](#6-query-the-gene-datastack)
7) [Develop Your Own Datastack](#7-develop-your-own-datastack)
8) [License & Legal Notice](#8-license--legal-notice)
9) [Support and Commercial SLAs](#9-support-and-commercial-slas)

### 1. Introduction

Gene data from bioinformatics resources is indispensable to basic research, yet it is massive, nested, and messy. The **interlace-go** [library](https://github.com/sas-quantome/interlace-go#interlace-go) simplifies data refinement for building and maintaining clean-slate datasets as primary storage ready for downstream processing. To demonstrate this in practice, human gene data is a good test case for **interlace-go**.

The software engineering outcome of **interlace-ex** consolidates 57 GB of uncompressed gene and disease information into 1.4 GB of refined **interlace-go** binary data assets, ready to load into memory and power up basic research. 

Consolidation and data refinement take less than five minutes on a modern personal computer with a cap of four threads. This agility makes daily updates possible across [23 IP addresses](configs/input/url.txt). 

The **interlace-ex** code provides many examples of transforming raw data into refined binary assets using simply **interlace-go**.

[15 Go programs](./cmd/README.md) are offered as starting models for creating command-line tools that make robust datasets. The ensemble of these Go programs creates one **Interlace Entity**, named *Gene*, and 14 **Interlace Assets** that serve as a point of reference to evaluate quality and performance:

1) *Gene Product*
2) *Gene Synonym*
3) *Gene Metadata*
4) *Gene PubMed Reference*
5) *Gene Function*
6) *Gene Ontology*
7) *Gene Medical Genetics*
8) *Gene Ortholog*
9) *Gene Molecular Interaction*
10) *Gene Genetic Variation*
11) *Gene Genomic Variation*
12) *Gene Pathway*
13) *Gene Chemical Reaction*
14) *Gene Protein Interaction*

The **interlace-ex** repository tools crystallize knowledge from all known genes in the human genome, along with code examples to help you add data assets to the Gene entity or build your own in other domains.

### 2. Data Providers

The *Gene* datastack is possible thanks to the following bioinformatics resources:

- [Ancestry Category Label Distribution in the GWAS Catalog](https://www.ebi.ac.uk/gwas/docs/ancestry-data)
- [NCBI ClinVar](https://www.ncbi.nlm.nih.gov/clinvar/)
- [EBI GWAS Catalog](https://www.ebi.ac.uk/gwas/)
- [Gender, Sex, and Sexual Orientation Ontology](https://obofoundry.org/ontology/gsso.html)
- [NCBI GeneRIF](https://www.ncbi.nlm.nih.gov/gene/about-generif)
- [Gene Ontology](https://www.geneontology.org/)
- [Human Phenotype Ontology](https://hpo.jax.org/)
- [NCBI MedGen](https://www.ncbi.nlm.nih.gov/medgen/)
- [Mondo Disease Ontology](https://mondo.monarchinitiative.org/)
- [NCBI Gene DB](https://www.ncbi.nlm.nih.gov/gene/)
- [Ontology of Biological Attributes](https://github.com/obophenotype/bio-attribute-ontology)
- [Open Biological and Biomedical Ontology Foundry](https://obofoundry.org/)
- [Reactome](https://reactome.org/)
- [STRING DB](https://string-db.org/)
- [UniProt](https://www.uniprot.org/)

### 3. Data Creators and Peer-reviewed Journal References

The following papers describe in detail the datasets used to create the *Gene* datastack:

#### [A standardized framework for representation of ancestry data in genomics studies, with application to the NHGRI-EBI GWAS Catalog.](https://pubmed.ncbi.nlm.nih.gov/29448949/)

Morales J, Welter D, Bowler EH, Cerezo M, Harris LW, McMahon AC, Hall P, Junkins HA, Milano A, Hastings E, Malangone C, Buniello A, Burdett T, Flicek P, Parkinson H, Cunningham F, Hindorff LA, MacArthur JAL.

Genome Biol. 2018 Feb 15;19(1):21.

doi: 10.1186/s13059-018-1396-2. PMID: 29448949; PMCID: PMC5815218.


#### [ClinVar: public archive of relationships among sequence variation and human phenotype.](https://pubmed.ncbi.nlm.nih.gov/24234437/)

Landrum MJ, Lee JM, Riley GR, Jang W, Rubinstein WS, Church DM, Maglott DR.

Nucleic Acids Res. 2014 Jan; 42 (Database issue): D980-5.

doi: 10.1093/nar/gkt1113. Epub 2013 Nov 14. PMID: 24234437; PMCID: PMC3965032.


#### [The Human Phenotype Ontology in 2024: phenotypes around the world.](https://pubmed.ncbi.nlm.nih.gov/37953324/)

Gargano MA, Matentzoglu N, Coleman B, Addo-Lartey EB, Anagnostopoulos AV, Anderton J, Avillach P, Bagley AM, Bakštein E, Balhoff JP, Baynam G, Bello SM, Berk M, Bertram H, Bishop S, Blau H, Bodenstein DF, Botas P, Boztug K, Čady J, Callahan TJ, Cameron R, Carbon SJ, Castellanos F, Caufield JH, Chan LE, Chute CG, Cruz-Rojo J, Dahan-Oliel N, Davids JR, de Dieuleveult M, de Souza V, de Vries BBA, de Vries E, DePaulo JR, Derfalvi B, Dhombres F, Diaz-Byrd C, Dingemans AJM, Donadille B, Duyzend M, Elfeky R, Essaid S, Fabrizzi C, Fico G, Firth HV, Freudenberg-Hua Y, Fullerton JM, Gabriel DL, Gilmour K, Giordano J, Goes FS, Moses RG, Green I, Griese M, Groza T, Gu W, Guthrie J, Gyori B, Hamosh A, Hanauer M, Hanušová K, He YO, Hegde H, Helbig I, Holasová K, Hoyt CT, Huang S, Hurwitz E, Jacobsen JOB, Jiang X, Joseph L, Keramatian K, King B, Knoflach K, Koolen DA, Kraus ML, Kroll C, Kusters M, Ladewig MS, Lagorce D, Lai MC, Lapunzina P, Laraway B, Lewis-Smith D, Li X, Lucano C, Majd M, Marazita ML, Martinez-Glez V, McHenry TH, McInnis MG, McMurry JA, Mihulová M, Millett CE, Mitchell PB, Moslerová V, Narutomi K, Nematollahi S, Nevado J, Nierenberg AA, Čajbiková NN, Nurnberger JI Jr, Ogishima S, Olson D, Ortiz A, Pachajoa H, Perez de Nanclares G, Peters A, Putman T, Rapp CK, Rath A, Reese J, Rekerle L, Roberts AM, Roy S, Sanders SJ, Schuetz C, Schulte EC, Schulze TG, Schwarz M, Scott K, Seelow D, Seitz B, Shen Y, Similuk MN, Simon ES, Singh B, Smedley D, Smith CL, Smolinsky JT, Sperry S, Stafford E, Stefancsik R, Steinhaus R, Strawbridge R, Sundaramurthi JC, Talapova P, Tenorio Castano JA, Tesner P, Thomas RH, Thurm A, Turnovec M, van Gijn ME, Vasilevsky NA, Vlčková M, Walden A, Wang K, Wapner R, Ware JS, Wiafe AA, Wiafe SA, Wiggins LD, Williams AE, Wu C, Wyrwoll MJ, Xiong H, Yalin N, Yamamoto Y, Yatham LN, Yocum AK, Young AH, Yüksel Z, Zandi PP, Zankl A, Zarante I, Zvolský M, Toro S, Carmody LC, Harris NL, Munoz-Torres MC, Danis D, Mungall CJ, Köhler S, Haendel MA, Robinson PN.

Nucleic Acids Res. 2024 Jan 5;52(D1):D1333-D1346.

doi: 10.1093/nar/gkad1005. PMID: 37953324; PMCID: PMC10767975.


#### [Gene: a gene-centered information resource at NCBI.](https://pubmed.ncbi.nlm.nih.gov/25355515/)

Brown GR, Hem V, Katz KS, Ovetsky M, Wallin C, Ermolaeva O, Tolstoy I, Tatusova T, Pruitt KD, Maglott DR, Murphy TD.

Nucleic Acids Res. 2015 Jan; 43 (Database issue): D36-42.

doi: 10.1093/nar/gku1055. Epub 2014 Oct 29. PMID: 25355515; PMCID: PMC4383897.


#### [The Gene Ontology knowledgebase in 2023.](https://pubmed.ncbi.nlm.nih.gov/36866529/)

Gene Ontology Consortium; Aleksander SA, Balhoff J, Carbon S, Cherry JM, Drabkin HJ, Ebert D, Feuermann M, Gaudet P, Harris NL, Hill DP, Lee R, Mi H, Moxon S, Mungall CJ, Muruganugan A, Mushayahama T, Sternberg PW, Thomas PD, Van Auken K, Ramsey J, Siegele DA, Chisholm RL, Fey P, Aspromonte MC, Nugnes MV, Quaglia F, Tosatto S, Giglio M, Nadendla S, Antonazzo G, Attrill H, Dos Santos G, Marygold S, Strelets V, Tabone CJ, Thurmond J, Zhou P, Ahmed SH, Asanitthong P, Luna Buitrago D, Erdol MN, Gage MC, Ali Kadhum M, Li KYC, Long M, Michalak A, Pesala A, Pritazahra A, Saverimuttu SCC, Su R, Thurlow KE, Lovering RC, Logie C, Oliferenko S, Blake J, Christie K, Corbani L, Dolan ME, Drabkin HJ, Hill DP, Ni L, Sitnikov D, Smith C, Cuzick A, Seager J, Cooper L, Elser J, Jaiswal P, Gupta P, Jaiswal P, Naithani S, Lera-Ramirez M, Rutherford K, Wood V, De Pons JL, Dwinell MR, Hayman GT, Kaldunski ML, Kwitek AE, Laulederkind SJF, Tutaj MA, Vedi M, Wang SJ, D'Eustachio P, Aimo L, Axelsen K, Bridge A, Hyka-Nouspikel N, Morgat A, Aleksander SA, Cherry JM, Engel SR, Karra K, Miyasato SR, Nash RS, Skrzypek MS, Weng S, Wong ED, Bakker E, Berardini TZ, Reiser L, Auchincloss A, Axelsen K, Argoud-Puy G, Blatter MC, Boutet E, Breuza L, Bridge A, Casals-Casas C, Coudert E, Estreicher A, Livia Famiglietti M, Feuermann M, Gos A, Gruaz-Gumowski N, Hulo C, Hyka-Nouspikel N, Jungo F, Le Mercier P, Lieberherr D, Masson P, Morgat A, Pedruzzi I, Pourcel L, Poux S, Rivoire C, Sundaram S, Bateman A, Bowler-Barnett E, Bye-A-Jee H, Denny P, Ignatchenko A, Ishtiaq R, Lock A, Lussi Y, Magrane M, Martin MJ, Orchard S, Raposo P, Speretta E, Tyagi N, Warner K, Zaru R, Diehl AD, Lee R, Chan J, Diamantakis S, Raciti D, Zarowiecki M, Fisher M, James-Zorn C, Ponferrada V, Zorn A, Ramachandran S, Ruzicka L, Westerfield M.

Genetics. 2023 May 4;224(1):iyad031.

doi: 10.1093/genetics/iyad031. PMID: 36866529; PMCID: PMC10158837.


#### [GeneRIF indexing: sentence selection based on machine learning.](https://pubmed.ncbi.nlm.nih.gov/23725347/)

Jimeno-Yepes AJ, Sticco JC, Mork JG, Aronson AR.

BMC Bioinformatics. 2013 May 31;14:171.

doi: 10.1186/1471-2105-14-171. PMID: 23725347; PMCID: PMC3687823.


#### [MedGen: NCBI's Portal to Information on Medical Conditions with a Genetic Component.](https://pubmed.ncbi.nlm.nih.gov/32329672/)

Louden DN.

Med Ref Serv Q. 2020 Apr-Jun; 39(2): 183-191.

doi: 10.1080/02763869.2020.1726152. PMID: 32329672.


#### [Mondo: Integrating Disease Terminology Across Communities.](https://pubmed.ncbi.nlm.nih.gov/41052288/)

Vasilevsky NA, Toro S, Matentzoglu N, Flack JE, Mullen KR, Hegde H, Gehrke S, Whetzel PL, Shwetar Y, Harris NL, Ngu MS, Alyea GL, Kane MS, Roncaglia P, Sid E, Thaxton CL, Wood V, Abraham RS, Achatz MI, Ajuyah P, Amberger JS, Babb L, Baker J, Balhoff JP, Berg JS, Bhalla A, Bofill-De Ros X, Braun IR, Broeren EC, Byer BK, Byrne AB, Callahan TJ, Carmody LC, Chan LE, Clause AR, Cohen JS, DeLuca M, Deuitch NT, Flowers M, Fraser J, Fujiwara T, Gitau V, Goldstein JL, Gration D, Groza T, Gyori BM, Hankey W, Hilton JA, Himmelstein DS, Hong SS, Hoyt CT, Huether R, Hurwitz E, Jacobsen JOB, Kikuchi A, Köhler S, Korn DR, Lagorce D, Laraway BJ, Li JY, Malheiro AJ, McLaughlin J, Meldal BHM, Mohan S, Moxon SAT, Munoz-Torres MC, Nelson TH, Nicholas FW, Ochoa D, Olson D, Oprea TI, Oskotsky TT, Osumi-Sutherland D, Paris K, Parkinson HE, Pendlington ZM, Peng XP, Pizzino A, Plon SE, Powell BC, Ratliff JC, Rehm HL, Remennik L, Riggs ER, Roberts S, Robinson PN, Ross JE, Schaper K, Schilder BM, Schmidt JL, Sharp EW, Similuk MN, Smedley D, Sneddon TP, Sparks R, Stefancsik R, Stupp GS, Sundar S, Takatsuki T, Tammen I, Tshering KC, Unni DR, Valasek E, Vanderver A, Wagner AH, Webb RF, Welter D, Yaya-Stupp D, Zankl A, Zhang XA, McMurry JA, Chute CG, Hamosh A, Mungall CJ, Haendel MA.

Genetics. 2025 Oct 6:iyaf215.

doi: 10.1093/genetics/iyaf215. Epub ahead of print. PMID: 41052288.


#### [The NHGRI-EBI GWAS Catalog: standards for reusability, sustainability, and diversity.](https://pubmed.ncbi.nlm.nih.gov/39530240/)

Cerezo M, Sollis E, Ji Y, Lewis E, Abid A, Bircan KO, Hall P, Hayhurst J, John S, Mosaku A, Ramachandran S, Foreman A, Ibrahim A, McLaughlin J, Pendlington Z, Stefancsik R, Lambert SA, McMahon A, Morales J, Keane T, Inouye M, Parkinson H, Harris LW.

Nucleic Acids Res. 2025 Jan 6; 53(D1): D998-D1005.

doi: 10.1093/nar/gkae1070. PMID: 39530240; PMCID: PMC11701593.


#### [The Ontology of Biological Attributes (OBA)-computational traits for the life sciences.](https://pubmed.ncbi.nlm.nih.gov/37076585/)

Stefancsik R, Balhoff JP, Balk MA, Ball RL, Bello SM, Caron AR, Chesler EJ, de Souza V, Gehrke S, Haendel M, Harris LW, Harris NL, Ibrahim A, Koehler S, Matentzoglu N, McMurry JA, Mungall CJ, Munoz-Torres MC, Putman T, Robinson P, Smedley D, Sollis E, Thessen AE, Vasilevsky N, Walton DO, Osumi-Sutherland D.

Mamm Genome. 2023 Sep;34(3):364-378.

doi: 10.1007/s00335-023-09992-1. Epub 2023 Apr 19. PMID: 37076585; PMCID: PMC10382347.


#### [The Reactome Pathway Knowledgebase 2024.](https://pubmed.ncbi.nlm.nih.gov/37941124/)

Milacic M, Beavers D, Conley P, Gong C, Gillespie M, Griss J, Haw R, Jassal B, Matthews L, May B, Petryszak R, Ragueneau E, Rothfels K, Sevilla C, Shamovsky V, Stephan R, Tiwari K, Varusai T, Weiser J, Wright A, Wu G, Stein L, Hermjakob H, D'Eustachio P.

Nucleic Acids Res. 2024 Jan 5; 52(D1): D672-D678.

doi: 10.1093/nar/gkad1025. PMID: 37941124; PMCID: PMC10767911.


#### [The STRING database in 2023: protein-protein association networks and functional enrichment analyses for any sequenced genome of interest.](https://pubmed.ncbi.nlm.nih.gov/36370105/)

Szklarczyk D, Kirsch R, Koutrouli M, Nastou K, Mehryary F, Hachilif R, Gable AL, Fang T, Doncheva NT, Pyysalo S, Bork P, Jensen LJ, von Mering C.

Nucleic Acids Res. 2023 Jan 6; 51(D1): D638-D646.

doi: 10.1093/nar/gkac1000. PMID: 36370105; PMCID: PMC9825434.


#### [UniProt: the Universal Protein Knowledgebase in 2025.](https://pubmed.ncbi.nlm.nih.gov/39552041/)

UniProt Consortium.

Nucleic Acids Res. 2025 Jan 6; 53(D1): D609-D617.

doi: 10.1093/nar/gkae1010. PMID: 39552041; PMCID: PMC11701636.

### 4. Clone, Compile, and Install

#### 4.1 Download and Install Go

Download and install Go from [https://go.dev/doc/install](https://go.dev/doc/install). It takes less than 5 minutes to complete!

#### 4.2 Declare Your Interlace Codebase Path
Declare a place where to clone the GitHub repositories the standard way somewhere in your computer. Include the `sas-quantome` directory to store multiple Interlace repositories:
```bash
REPO_PATH="${HOME}/github.com/sas-quantome"
```

#### 4.3 Create Your Interlace Codebase Directory
Create the directory:
```bash
mkdir -p "${REPO_PATH}"
```

#### 4.4 Clone the interlace-ex Repository
Clone the repository in the `sas-quantome` folder:
```bash
cd "${REPO_PATH}"
git clone https://github.com/sas-quantome/interlace-ex.git
```

#### 4.5 Check the Makefile
Review the Makefile and then run the following commands:
```bash
cd "${REPO_PATH}/interlace-ex"
make help
```

#### 4.6 Build the Toolchain
If all is well in your installation, you will find the compiled products in the `bin` folder:
```bash
cd "${REPO_PATH}/interlace-ex"
make build
```

#### 4.7 Make Install
Compile the software and install the binaries, you will find them in your `$GOPATH/bin`:
```bash
cd "${REPO_PATH}/interlace-ex"
make install
```

#### 4.8 Cleanup
Remove the `bin` folder:
```bash
cd "${REPO_PATH}/interlace-ex"
make clean
```

### 5. Build the *Gene* Datastack

#### 5.1 Declare Your Interlace Environment Variables
Define a place where to create the **Interlace Basepath** and download files somewhere in your computer, normally as follows:
```bash
ENV_AEE_BASE_PATH="${HOME}/interlace-ex"
ENV_ALT_DATA_PATH="${ENV_AEE_BASE_PATH}/data"
```
**Attention:** `ENV_ALT_DATA_PATH` and `ENV_AEE_BASE_PATH` are visible only in the active shell. It is better to add them in your `.bashrc` or `.zshrc` file.

```bash
export ENV_AEE_BASE_PATH="${HOME}/interlace-ex"
export ENV_ALT_DATA_PATH="${ENV_AEE_BASE_PATH}/data"
```

#### 5.2 Create Your Interlace Base
Create the Interlace directory structure:
```bash
mkdir -p ${ENV_AEE_BASE_PATH}/agent ${ENV_AEE_BASE_PATH}/entity ${ENV_AEE_BASE_PATH}/event
mkdir -p $ENV_ALT_DATA_PATH
```

- **agent:** Agent data.
- **data:** Downloaded datasets.
- **entity:** Refined binary data.
- **event:** Intermediate compressed files, Bash scripts, and logs.

#### 5.3 Review the [URI](./configs/input/uri.txt), [URL](./configs/input/url.txt), and [b31003](./scripts/bash/b31003)

- Take a look to these files before proceeding.
- Adapt the [URI](./configs/input/uri.txt) `file:///` paths to your local file system.

#### 5.4 Download the *Gene* Datasets
Validate URL links, download, rename non-persistent filenames, and move into the `ENV_ALT_DATA_PATH` in less than 8 minutes:
```bash
make download
```

#### 5.5 Create the *Gene* Datasets
Create the *Gene* binary files in about ~7 minutes. Jobs are executed sequentially, not in parallel.
```bash
make gene
```

- Review the log files in the event action directory.
- You should get the same tree but with your actual time tag:

```text
interlace-ex
├── agent
├── data
│   ├── HUMAN_9606_idmapping.dat.gz
│   ├── NCBI2Reactome_All_Levels.txt
│   ├── NCBI2Reactome_PE_Reactions.txt
│   ├── ReactionPMIDS.txt
│   ├── gene2accession.gz
│   ├── gene2ensembl.gz
│   ├── gene2go.gz
│   ├── gene2pubmed.gz
│   ├── gene_info.gz
│   ├── gene_orthologs.gz
│   ├── gene_summary.gz
│   ├── generifs_basic.gz
│   ├── gsso.obo
│   ├── gwas-catalog-ancestry.tsv.gz
│   ├── gwas-catalog-associations-annotated.tsv.gz
│   ├── gwas-efo-trait-mappings.tsv.gz
│   ├── hp.obo
│   ├── interactions.gz
│   ├── mim2gene_medgen
│   ├── mondo.obo
│   ├── oba.obo
│   ├── string-protein-links.tsv.gz
│   └── variant_summary.txt.gz
├── entity
│   └── gene
│       ├── chemical_reaction
│       │   └── 20260906180032
│       │       ├── catalog
│       │       │   ├── suk_chemical_reaction_hash_id_m.bin
│       │       │   └── uuk_chemical_reaction_text_s.bin
│       │       ├── enumeration
│       │       │   └── uuk_reaction_entity_cv_s.bin
│       │       └── property
│       │           ├── sdk_entrez_gene_id_u.bin
│       │           ├── udk_pubmed_id_u.bin
│       │           ├── udk_reaction_entity_cv_d.bin
│       │           ├── udk_reaction_hash_id_m.bin
│       │           └── udk_reactome_hsa_id_u.bin
│       ├── function
│       │   └── 20260906180032
│       │       ├── catalog
│       │       │   ├── suk_function_hash_id_m.bin
│       │       │   └── uuk_function_text_s.bin
│       │       ├── enumeration
│       │       └── property
│       │           ├── sdk_entrez_gene_id_u.bin
│       │           ├── udk_pubmed_id_u.bin
│       │           ├── udk_rif_hash_id_m.bin
│       │           └── udk_taxonomy_id_u.bin
│       ├── genetic_variation
│       │   └── 20260906180032
│       │       ├── catalog
│       │       │   ├── suk_genetic_variation_hash_id_m.bin
│       │       │   └── uuk_genetic_variation_text_s.bin
│       │       ├── enumeration
│       │       │   └── uuk_bit_flags_cv_s.bin
│       │       └── property
│       │           ├── sdk_entrez_gene_id_u.bin
│       │           ├── udk_bit_flags_m.bin
│       │           ├── udk_conical_allele_id_u.bin
│       │           ├── udk_omim_id_u.bin
│       │           ├── udk_refsnp_id_u.bin
│       │           ├── udk_variation_hash_id_m.bin
│       │           └── udk_variation_id_u.bin
│       ├── genomic_variation
│       │   └── 20260906180032
│       │       ├── catalog
│       │       │   ├── suk_genomic_variation_hash_id_m.bin
│       │       │   └── uuk_genomic_variation_text_s.bin
│       │       ├── enumeration
│       │       │   └── uuk_bit_flags_cv_s.bin
│       │       └── property
│       │           ├── sdk_entrez_gene_id_u.bin
│       │           ├── udk_bit_flags_u.bin
│       │           ├── udk_gwas_hash_id_m.bin
│       │           ├── udk_ontology_id_s.bin
│       │           ├── udk_pubmed_id_u.bin
│       │           └── udk_refsnp_id_u.bin
│       ├── medical_genetics
│       │   └── 20260906180032
│       │       ├── catalog
│       │       ├── enumeration
│       │       │   └── uuk_bit_flags_cv_s.bin
│       │       └── property
│       │           ├── sdk_entrez_gene_id_u.bin
│       │           ├── udk_bit_flags_u.bin
│       │           ├── udk_medgen_cui_s.bin
│       │           └── udk_omim_id_u.bin
│       ├── metadata
│       │   └── 20260906180032
│       │       ├── catalog
│       │       │   ├── suk_metadata_hash_id_m.bin
│       │       │   └── uuk_metadata_text_s.bin
│       │       ├── enumeration
│       │       │   └── uuk_bit_flags_cv_s.bin
│       │       └── property
│       │           ├── suk_entrez_gene_id_u.bin
│       │           ├── udk_bit_flags_u.bin
│       │           ├── udk_chromosome_s.bin
│       │           ├── udk_gene_name_s.bin
│       │           ├── udk_gene_symbol_s.bin
│       │           ├── udk_metadata_hash_id_m.bin
│       │           ├── udk_name_synonyms_set_s.bin
│       │           ├── udk_symbol_synonyms_set_s.bin
│       │           ├── udk_taxonomy_id_u.bin
│       │           └── udk_xdb_id_set_s.bin
│       ├── molecular_interaction
│       │   └── 20260906180032
│       │       ├── catalog
│       │       ├── enumeration
│       │       └── property
│       │           ├── sdk_entrez_gene_id_u.bin
│       │           ├── udk_interactant_entrez_gene_id_u.bin
│       │           ├── udk_interactant_taxonomy_id_u.bin
│       │           ├── udk_pubmed_id_set_s.bin
│       │           └── udk_taxonomy_id_u.bin
│       ├── ontology
│       │   └── 20260906180032
│       │       ├── catalog
│       │       │   ├── suk_ontology_hash_id_m.bin
│       │       │   └── uuk_ontology_text_s.bin
│       │       ├── enumeration
│       │       │   └── uuk_bit_flags_cv_s.bin
│       │       └── property
│       │           ├── sdk_entrez_gene_id_u.bin
│       │           ├── udk_bit_flags_m.bin
│       │           ├── udk_go_hash_id_m.bin
│       │           ├── udk_go_id_u.bin
│       │           └── udk_pubmed_id_u.bin
│       ├── ortholog
│       │   └── 20260906180032
│       │       ├── catalog
│       │       ├── enumeration
│       │       └── property
│       │           ├── sdk_entrez_gene_id_u.bin
│       │           ├── udk_ortholog_entrez_gene_id_u.bin
│       │           ├── udk_ortholog_taxonomy_id_u.bin
│       │           └── udk_taxonomy_id_u.bin
│       ├── pathway
│       │   └── 20260906180032
│       │       ├── catalog
│       │       │   ├── suk_pathway_hash_id_m.bin
│       │       │   └── uuk_pathway_text_s.bin
│       │       ├── enumeration
│       │       └── property
│       │           ├── sdk_entrez_gene_id_u.bin
│       │           ├── udk_pathway_hash_id_m.bin
│       │           └── udk_reactome_hsa_id_u.bin
│       ├── product
│       │   └── 20260906180032
│       │       ├── catalog
│       │       ├── enumeration
│       │       │   └── uuk_synonym_category_cv_s.bin
│       │       └── property
│       │           ├── sdk_entrez_gene_id_u.bin
│       │           ├── udk_protein_id_s.bin
│       │           ├── udk_synonym_category_cv_d.bin
│       │           └── udk_synonym_s.bin
│       ├── protein_interaction
│       │   └── 20260906180032
│       │       ├── catalog
│       │       ├── enumeration
│       │       └── property
│       │           ├── sdk_entrez_gene_id_u.bin
│       │           ├── udk_association_ensp_id_i.bin
│       │           ├── udk_association_entrez_gene_id_u.bin
│       │           ├── udk_ensp_id_i.bin
│       │           └── udk_likelihood_score_n.bin
│       ├── pubmed_reference
│       │   └── 20260906180032
│       │       ├── catalog
│       │       ├── enumeration
│       │       └── property
│       │           ├── sdk_entrez_gene_id_u.bin
│       │           └── udk_pubmed_id_u.bin
│       └── synonym
│           └── 20260906180032
│               ├── catalog
│               ├── enumeration
│               │   └── uuk_synonym_category_cv_s.bin
│               └── property
│                   ├── sdk_synonym_s.bin
│                   ├── udk_entrez_gene_id_u.bin
│                   ├── udk_protein_id_s.bin
│                   ├── udk_synonym_category_cv_d.bin
│                   └── udk_synonym_version_s.bin
└── event
    └── gene
        └── w31001
            └── 20260906180032
                ├── action
                │   ├── 20260906180328_a31001.log
                │   ├── 20260906180406_a31002.log
                │   ├── 20260906180445_a31003.log
                │   ├── 20260906180456_a31004.log
                │   ├── 20260906180505_a31005.log
                │   ├── 20260906180540_a31006.log
                │   ├── 20260906180540_a31007.log
                │   ├── 20260906180603_a31008.log
                │   ├── 20260906180610_a31009.log
                │   ├── 20260906180754_a31010.log
                │   ├── 20260906180803_a31011.log
                │   ├── 20260906180804_a31012.log
                │   ├── 20260906180806_a31013.log
                │   ├── 20260906180916_a31014.log
                │   └── 20260906180917_a31015.log
                ├── artifact
                └── function
```

### 6. Query the *Gene* Datastack

Refining and retrieving data have very distinct workflows. Refining data is project-specific, and retrieving is not. Retrieval applications must read any binary file part of the **Interlace Architecture**. For this reason, the **interlace-do** [toolset](https://github.com/sas-quantome/interlace-do) is built separately and uses the **interlace-ex** *Gene* as the data playground.

### 7. Develop Your Own Datastack

Follow these instructions and find the code patterns that will guide you in making an **interlace-go** command line application:

- Find out the differences in the main files [m1], [m2], [m13], and [m14].
- Look for `Start()`, `Commit()`, and `End()` in [i1], [i2], [i13], and [i14].
- Read the **interlace-go** [PKG] page and focus on **Manager**, **Dispenser**, **Refiner**, **Interlacer**, and **Cowriter**.
- Take a good glimpse at the [Interlace Architecture].
- Discover the simplicity of the [Binary File Nomenclature].
- Learn the [File Integrity Check] design.
- See how simple it is to collect data from a [tab-delimited file](./internal/app/a31003/a31003_inp_gene_info.go).

### 8. License & Legal Notice

This project is licensed under the **Apache License 2.0**.

* **Permissive Reuse:** You are free to modify, distribute, and use this software in your own commercial or private applications.
* **No Lock-In:** Your downstream solutions remain entirely your own.
* **As-Is Basis:** Per Sections 7 and 8 of the Apache 2.0 License, this software is provided strictly on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND.

For full details, please review the [LICENSE](LICENSE) and [NOTICE](NOTICE) files in the root of this repository.

### 9. Support and Commercial SLAs

This repository is an **open-source transparency initiative** and does not come with a dedicated maintenance or support SLA.

Bug reports and feature suggestions can be submitted via GitHub Issues.

If your organization requires guaranteed technical support, custom feature development, or integration assistance, please contact your account manager or reach out via [Quantome SAS](https://www.sas-quantome.com).

***
*Copyright © 2026 Quantome SAS. All rights reserved.*

[m1]: ./cmd/a31001/main.go
[m2]: ./cmd/a31002/main.go
[m13]: ./cmd/a31013/main.go
[m14]: ./cmd/a31014/main.go

[i1]: ./internal/app/a31001/a31001.go
[i2]: ./internal/app/a31002/a31002.go
[i13]: ./internal/app/a31013/a31013.go
[i14]: ./internal/app/a31014/a31014.go

[PKG]: https://github.com/sas-quantome/interlace-go/blob/main/pkg/README.md
[Interlace Architecture]: https://github.com/sas-quantome/interlace-go/blob/main/pkg/README.md#9-directory-architecture-entities-and-events
[Binary File Nomenclature]: https://github.com/sas-quantome/interlace-go/blob/main/pkg/README.md#10-binary-file-nomenclature
[File Integrity Check]: https://github.com/sas-quantome/interlace-go/blob/main/pkg/README.md#11-binary-file-integrity-check
