<h1 name="title">gh-xk6</h1>

**Maintain k6 extensions hosted on GitHub**

A gh extension that helps maintain k6 extensions hosted on GitHub.

## Installation

1. Install the `gh` CLI - see the [installation](https://github.com/cli/cli#installation)

   _Installation requires a minimum version (2.0.0) of the the GitHub CLI that supports extensions._

2. Install this extension:

   ```sh
   gh extension install grafana/gh-xk6
   ```

# Usage

<!-- #region cli -->
## gh xk6

**Maintain k6 extensions hosted on GitHub**

A gh extension that helps maintain k6 extensions hosted on GitHub.


### Commands

* [gh xk6 catalog](#gh-xk6-catalog)	 - Maintain k6 extension catalog

---
## gh xk6 catalog

**Maintain k6 extension catalog**

Maintain k6 extension catalog based on GitHub search

### SEE ALSO

* [gh xk6](#gh-xk6)	 - Maintain k6 extensions hosted on GitHub
### Commands

* [gh xk6 catalog create](#gh-xk6-catalog-create)	 - Create new extension catalog
* [gh xk6 catalog import](#gh-xk6-catalog-import)	 - Import k6-docs extension registry
* [gh xk6 catalog update](#gh-xk6-catalog-update)	 - Update versions in extension catalog

---
## gh xk6 catalog create

Create new extension catalog

### Synopsis

Create new extension catalog with only the mandatory k6 entry

```
gh xk6 catalog create [flags]
```

### Flags

```
  -f, --file string   Extension catalog filename (default "k6catalog.json")
      --force         Force overwriting of the existing file
  -h, --help          Help for createcommand
```

### SEE ALSO

* [gh xk6 catalog](#gh-xk6-catalog)	 - Maintain k6 extension catalog

---
## gh xk6 catalog import

Import k6-docs extension registry

### Synopsis

Import k6-docs extension registry into k6 extension catalog

```
gh xk6 catalog import [flags]
```

### Flags

```
  -f, --file string    Extension catalog filename (default "k6catalog.json")
  -q, --filter query   JMESPath query for filtering registry entries
      --force          Force overwriting of the existing file
  -h, --help           Help for importcommand
  -p, --preset name    Select a preset JMESPath filter by name (official|cloud)
```

### SEE ALSO

* [gh xk6 catalog](#gh-xk6-catalog)	 - Maintain k6 extension catalog

---
## gh xk6 catalog update

Update versions in extension catalog

### Synopsis

Update versions in extension catalog using GitHub search

```
gh xk6 catalog update [flags]
```

### Flags

```
  -f, --file string   Extension catalog filename (default "k6catalog.json")
  -h, --help          Help for updatecommand
```

### SEE ALSO

* [gh xk6 catalog](#gh-xk6-catalog)	 - Maintain k6 extension catalog

<!-- #endregion cli -->
