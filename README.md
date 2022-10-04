# txetnoc - An inverse Go context
This package implements a context that is Done when the first of its
children Done channel is closed. When that happens it also propagates
the Err from the child.

## Why
Because I wanted to experiment to see how hard would it be to
implement something like this.
