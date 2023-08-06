- There are 10 types of data:
  - Scalar types
    - Boolean
    - Int
    - Float
    - String
  - Compound types
    - Array
    - Object
    - Callable
    - Iterable
  - Special types
    - Resource
    - Null
- Type hint for multiple values is separated by `|`, `int|float`.
- In PHP 8, type hint `mixed` indicate any data types.

## null

- A variable can be `null` if it is:
  - undefined
  - explicitly assigned a value of `null`
  - destroyed with `unset()`
- There are 2 ways to check if a variable is `null`:
  - `is_null()`
  - `$var === null`

## Boolean

- During implicit type conversion/type casting, any truthy value will be converted to `true`: `1, 1.0, "a"`
- During implicit type conversion/type casting, any falsy value will be converted to `false`: `0, 0.0, "", null`
- The result of `echo true` value is `"1"`
- The result of `echo false` value is `""`

## Int

- The result of value overflow (value > `PHP_INT_MAX`) is implicit type conversion into `float`
- `is_int()` check if value is int or not

## Float

- The result of value overflow (value > `PHP_FLOAT_MAX`) is `INF`
- Never compare two float values directly because of loss of precision issue

## String

- `.` operator is used to concatenate strings.
- `nl2br()` inserts HTML line breaks before all newlines in a string.
- Accessing string character using curly bracket `$name{0}` is deprecated in PHP 7.4.
- Heredoc treat string as double quotes, which mean you can insert a variable within the string.
- Nowdoc treat string as single quotes, which mean you can't.insert a variable within the string
- String in PHP is mutable, you can directly modify a character by accessing it index, `$name[1] = "b"`. Modifying a character by accessing out of range index will result in the string being padded with spaces to fill the length until the specified index, for example `$name` string have length of 15 (last index of 14), `$name[19] = "s"` will result in 4 spaces followed by char "s" appended into the string.
- You can access character of a string by using index, `$name[0]`.
- You can also access character of a string from behind by using negative index, second character from behind `$name[-2]`.

## Array

- `array_pop` Pop the element off the end of array.
- `array_shift()` Shift an element off the beginning of array, after the operation is done, the array will be reindexed.
- `count()` Counts all elements in an array, or something in an object.
- `isset()` determine if a variable is declared and is different than NULL, can be used to check array element exist or not.
- `print_r()` prints human-readable information about a variable, especially useful for printing array.
- `unset($array[0])` will retains the maximum index.
- Another way to check array element exist or not is `array_key_exists()`, one difference compared to `isset()` is when value of the element is `null`, it will return `true`.
- If a key of an element in array is not a string or int, PHP will try to cast it into int, with exception of `null` will be casted into an empty strings.
- To append element to an array, you can use `$array[] = ` or `array_push()`.
