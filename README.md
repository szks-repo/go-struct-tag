# go-struct-tag

go-struct-tag is a utility library designed to simplify the retrieval, editing, and manipulation of Go language struct tags. While Go's standard reflect package makes it easy to get tag values, safely editing existing tags or adding new ones often requires manual string parsing and concatenation, which can be complex and error-prone. This library aims to solve that problem by providing a more intuitive and robust way to manage struct tags.

## 🚀 Overview
Go's struct tag is a powerful mechanism for providing metadata to fields (e.g., json:"field_name,omitempty"). However, the Go standard library (reflect package) supports retrieving tag values via the reflect.StructTag.Get() method but does not provide direct means to programmatically modify tags. Consequently, developers typically have to manually parse the tag string, apply the desired changes, and then reconstruct a new string—a tedious and error-prone process.

go-struct-tag abstracts this process, offering a simple and safe API to treat tags as key-value pairs. This frees developers from cumbersome string operations, allowing them to manage tags more efficiently.

## 📄 License
This project is licensed under the MIT License. See the LICENSE file for details.
