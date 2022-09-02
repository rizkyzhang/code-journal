## SortableJS - Disable sort on a specific element

You need to add both:

the filter: 'selector' property
the onMove: `function (e) { return e.related.className.indexOf('static') === -1; }` callback listener function
to completely prevent an element from being rearranged at the beginning or end of a list.

## Reference

https://stackoverflow.com/a/71350502
