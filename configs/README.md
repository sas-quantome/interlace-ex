###### September 4, 2026: interlace-ex configs readme v1

# Configs

```text
configs
├── README.md
└── input
    ├── uri.txt
    └── url.txt
```

## Contents

### [URI]

- Use this URI file to facilitate applications finding their correct resources.
- There is no need to edit this file unless a URL has changed.
- Before using it verify if the source databases have newer releases or URL updates.

### [URL]

- This file contains URLs in the URI file.
- The Bash scripts require this file to download all the input files in advance.
- The scripts also check for broken URLs.

[URI]: input/uri.txt
[URL]: input/url.txt