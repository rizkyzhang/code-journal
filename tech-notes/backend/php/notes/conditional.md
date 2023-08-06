## Different between switch and match

- Switch is a statement, while match is an expression.
- A match arm can have multiple expressions (separated by a comma), while switch case may only have a single expression.
- In match, you can only have single-line expressions that return a value when a match arm matches. In switch, however, each case can have multiple statements (i.e. a block of statements) that are executed if the case matches.
- Switch comparison is loose, while match is strict.
- When there's no match found, and a default does not exist, the match expression throws an `UnhandledMatchError`, while switch does nothing.
- switch continues to execute the statements until the end of the switch block (fallthrough) till a terminating statement (such as break, continue or return) is found.

```php
switch ($subject) {
    case 'one':
        echo 'one;';
    case 'two':
        echo 'two;';
    case 'three':
        echo 'three;';
        break;
    default:
        echo 'default;';
}
```

```php
$result = match ($subject) {
    'one' => 'one;',
    'two' => 'two;',
    'three' => 'three;',
    default => 'default;',
};
```

- The null-safe operator supports method while the null-coalescing operator doesn't. The null-coalescing operator supports array while the null-safe operator doesn't.
