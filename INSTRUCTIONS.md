# go-reloaded - Project Instructions

## 🎯 Objectives
Use previous knowledge to build a simple text completion, editing, and auto-correction tool of formatting strings layout directly via Standard packages only.

---

## 📖 Introduction
The tool receives two file arguments:
```bash
go run . <input_file> <output_file>
```

### 📋 Modification Rules

#### 1. Numeric Conversions
*   **(hex)**: Replace previous word with decimal integer version (assuming it was hexadecimal).
*   **(bin)**: Replace previous word with decimal integer version (assuming it was binary).

#### 2. Case Conversion
*   **(up)**: Converts previous word into uppercase.
*   **(low)**: Converts previous word into lowercase.
*   **(cap)**: Converts previous word into Capitalized version.
*   **(up, <number>)**, **(low, <number>)**, **(cap, <number>)**: Applies to the previous `<number>` words backwards.

---

## 🛠️ Formatting Actions

### 1. Punctuation Spacing
*   Applied marks: `.`, `,`, `!`, `?`, `:`, and `;`.
*   Rule Structure: Close to previous word immediately, with spaces apart from the next sequential frames layout securely.
*   Contiguous Groups like `...` or `!?` stay groups without spacing inner intervals split setups.

### 2. Single Quotes Cleanup
*   Stops spaces cushioning inside matching groups buffers matching pairs tightly around containing contents setups tightly.

### 3. Vowel Sound Articles Support
*   Turns `a` into `an` if the following next word starts with any vowel sound chunk (`a`, `e`, `i`, `o`, `u`) or `h`.

---

## 📦 Packages
Only standard library package sets are allowed. No third-party modules.
