# go-quilting-api

A POC api for creating a pattern based dynamic merge of documents.

A continuation on the ideas learned in https://github.com/TheWozard/dynamic-distributed-merging. Moving from a completely theoretical realm to one with a bit more application. This aims to work on smells found during the design of the first.
- Embedding traversal information into the provided document gets really sloppy especially when higher priority documents can change that. I found myself distinguishing between behavior for  **Values** and **Object/Lists** at almost every point. This either means the abstraction was poor, or I am trying to define two different problems as one. This leads to the idea of a separate traversal document that defines how the data is traversed. In this case calling it the **Pattern Document** hence the quilting theme.
- If we want to be producing a consistent result, and we know what that is ahead of time, why do we have to traverse every document to know if we found every thing. This also supports the idea of a common json-schema like document that defines this traversal.
- While working on this idea I missed the usage of interfaces to strictly define abstractions and type checking, so this is switching to golang.
