# types

* bool - `true` or `false`
* number (float64) - `0.123`
* string (unicode) - `"hello"`
* map/hash/table/dictionary - `{1, 2, 3, "named": "value"}`
* function - `fn () end`
* nil - `nil`

# variables

```
a = 5

fn change_a() 
	a = 6
end

change_a()

print(a) # 6
```

Always overwrite outer scope

# branching

```
case (bool expression | nothing) then
	(body)
case (bool expression | nothing) then
	(body)
end
```

# loop

```
loop
	(body | break | continue)
end
```

# operators

All
* `==`
* `!=`

Numbers and strings
* `>`
* `<`
* `>=`
* `<=`

Numbers
* `+`
* `-`
* `*`
* `/`

Strings
* `..`

# function

```
fn (name | nothing) ((arg, arg, arg)) then
	(body)
end
```

# errors

???

# built-in functions

* print(...args: any)
* number(arg: any): int
* string(arg: any): string
* bool(arg: any): bool

# type functions

* string
	* json(): table
	* upper(): string
	* lower(): string
	* contains(str: string): bool
	* split(str: string): table
	* has_prefix(str: string): bool
	* has_suffix(str: string): bool
	* repeat(times: number): string

* table
	* each(func: function)
	* map(func: function): table
	* filter(func: function): table
	* reduce(func: function, init: any): any
	* sort(func: function): table
	* has(key: any): bool


# standard library

* http
	* get(url: string)
	* download(url: string, path: string)

* log
	* log(str: string)
	* error(str: string)

* fs
	* path(...args: any): string
	* exist(path: string): bool
	* mkdir(path: string)
	* read(path: string): string
	* write(path: string, data: string)
	* append(path: string, data: string)
	* create(path: string)
	* home(): string
	* config(): string
	* cache(): string
	* temp_file(): string
	* temp_dir(): string
	* dir(path: string): table | nil
	* cwd(): string
	* glob(path: string): table

* json
	* parse(string): obj
	* string(obj: any): string

* os
	* args()
	* exit(code: number)
	* exec(cmd: string): string
	* cmd(cmd: string): table

* time
	* sleep(seconds: number)

* math
	* ...