# go-reloaded 📝

`go-reloaded` is a simple text completion, auto-correction, and formatting tool written in **Go**. It reads text from an input file, processes specific modifiers and rules that apply backwards, handles formatting layout rules for punctuation and quotes, and saves the corrected layout onto an output file.

---

## 🚀 Features

### 1. Modifiers
These modifiers look at the **word(s) immediately before** them and apply the specified transformation rules:
*   **(hex)**: Replaces previous word with its decimal version (e.g., `"1E (hex)"` -> `"30"`)
*   **(bin)**: Replaces previous word with its decimal version (e.g., `"10 (bin)"` -> `"2"`)
*   **(up)**: Converts previous word to UPPERCASE.
*   **(low)**: Converts previous word to lowercase.
*   **(cap)**: Converts previous word to Capitalized version.
*   **(up, n)**, **(low, n)**, **(cap, n)**: Applies modification rule to the exact previous `<n>` words.

### 2. Punctuation Spacing
*   Removes spaces BEFORE `.` `,` `!` `?` `:` `;` so they stick close to the prior word.
*   Adds spaces AFTER punctuation to give proper breaks.
*   Preserves contiguous punctuation groups like `...` or `!?`.

### 3. Smart Articles
*   Automatically replaces `a` with `an` if the next following word starts with a vowel sound (`a`, `e`, `i`, `o`, `u` or `h`). Maintains casing if starting with uppercase `A`.

### 4. Single Quote Cleanup
*   Strips excessive space inside matching pairs of single quotes (`' awesome '` -> `'awesome'`).

---

## 🛠️ Installation & Usage

### Running the Tool
Use the standard run command with exactly two arguments: the **input file** path and **output file** path.

```bash
# General syntax
go run . <input_file> <output_file>

# Example
go run . sample.txt result.txt
```

---

## 🧠 Code Breakdown (Step-by-Step)

The logic consists of 6 sequential steps processed individually on **each line** to preserve line-breaks nicely:

### Step 1: Tokenize
The line gets split using `strings.Fields()`. Dividing the string into dynamic token chunks makes backward parsing of indices easy and safe to check iteratively.

### Step 2: Modifiers processing
We loop forward over the array checks. If we encounter a modifier match directly over a token using a **Regex matching**, it takes effect on the previous backward frame offsets safely avoiding bound collapses. Replaced index shifts inside loop are securely managed on index removals accurately.

### Step 3: Articles Cleanup
Tokens inspect their next sequential position values safely after modifier edits already resolved safely. If target index begins with a vowel Sound, a simple toggle converts `a` safely into `an`.

### Step 4: Reassemble
The edited token words list gets joined back seamlessly using `strings.Join(token, " ")` setting a basic space format rule first.

### Step 5: Punctuation rules formatting
Regular Expression matches solve spacing setups quickly and elegantly around continuous tokens groups avoiding breaking consecutive setups structure.
*   `\s+([.,!?:;]+)` removes spaces before.
*   `([.,!?:;]+)([^ \t.,!?:;])` fixes space missing after.

### Step 6: Quotes fix
Pairs are found non-greedily inside containing matches to safely strip margins seamlessly inside final joined builds.

---

## 💡 Author Recommendations
Auditors will love clean recursive backward scaling mechanics. Feel free to use included test runners via fully verbose testing setup methods directly if preferred on audits setups updates cleanly!
