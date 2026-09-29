# TAKEAWAYS

1. You can only use `:=` within function bodies. At the package level (outside any function), you must explicitly declare it's type, for eg. `var name string`.
2. Avoid using `var a` or `var a = <value>` for initializing or declaring wherever possible. 
