###### September 4, 2026: interlace-ex internal readme v2

# Internal

```text
internal
├── README.md
├── app
│   ├── README.md
│   ├── a31001
│   │   ├── a31001.go
│   │   ├── a31001_enums.go
│   │   ├── a31001_gene_product.go
│   │   ├── a31001_inp_ensembl_id_map.go
│   │   ├── a31001_inp_gene_info.go
│   │   ├── a31001_inp_refseq_id_map.go
│   │   ├── a31001_inp_uniprot_id_map.go
│   │   └── a31001_writers.go
│   ├── a31002
│   │   ├── a31002.go
│   │   └── a31002_gene_synonym.go
│   ├── a31003
│   │   ├── a31003.go
│   │   ├── a31003_bit_flags.go
│   │   ├── a31003_gene_metadata.go
│   │   ├── a31003_inp_gene_info.go
│   │   └── a31003_inp_refseq_sum.go
│   ├── a31004
│   │   ├── a31004.go
│   │   ├── a31004_gene_pubmed_reference.go
│   │   └── a31004_inp_gene_pubmed.go
│   ├── a31005
│   │   ├── a31005.go
│   │   ├── a31005_gene_function.go
│   │   └── a31005_inp_gene_rif.go
│   ├── a31006
│   │   ├── a31006.go
│   │   ├── a31006_bit_flags.go
│   │   ├── a31006_gene_ontology.go
│   │   └── a31006_inp_gene_go.go
│   ├── a31007
│   │   ├── a31007.go
│   │   ├── a31007_bit_flags.go
│   │   ├── a31007_gene_medical_genetics.go
│   │   └── a31007_inp_medgen.go
│   ├── a31008
│   │   ├── a31008.go
│   │   ├── a31008_gene_ortholog.go
│   │   └── a31008_inp_orthologs.go
│   ├── a31009
│   │   ├── a31009.go
│   │   ├── a31009_gene_molecular_interaction.go
│   │   └── a31009_inp_rif_interactions.go
│   ├── a31010
│   │   ├── a31010.go
│   │   ├── a31010_bit_flags.go
│   │   ├── a31010_gene_genetic_variation.go
│   │   └── a31010_inp_variant_sum.go
│   ├── a31011
│   │   ├── a31011.go
│   │   ├── a31011_gene_genomic_variation.go
│   │   ├── a31011_inp_efo_terms.go
│   │   ├── a31011_inp_gwas_cat.go
│   │   └── a31011_inp_obo_terms.go
│   ├── a31012
│   │   ├── a31012.go
│   │   ├── a31012_gene_pathway.go
│   │   └── a31012_inp_reactome_all_levels.go
│   ├── a31013
│   │   ├── a31013.go
│   │   ├── a31013_enums.go
│   │   ├── a31013_gene_chemical_reaction.go
│   │   ├── a31013_inp_reactome_gene.go
│   │   └── a31013_inp_reactome_pubmed.go
│   ├── a31014
│   │   ├── a31014.go
│   │   ├── a31014_gene_protein_interaction.go
│   │   └── a31014_inp_protein_link.go
│   └── a31015
│       ├── a31015.go
│       ├── a31015_bit_flags.go
│       ├── a31015_gene_genomic_variation_ancestry.go
│       └── a31015_gwas_ancestry.go
└── pkg
    ├── README.md
    └── arranger
        └── definer.go
```

## Contents

### [app](./app)

Application non-shared Go code.

### [pkg](./pkg)

Internal package shared Go code.
