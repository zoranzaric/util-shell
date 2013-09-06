#why

Bash is *NOT* readable. Experienced bash users often came up with even more non-readable script.
To prove this, see [String contains in bash](http://stackoverflow.com/questions/229551/string-contains-in-bash).


This is the top answer there:

```
	string='My string';

	if [[ "$string" == *My* ]]
	then
	  echo "It's there!";
	fi
	
```
You have to google for a while to know more bash detail. (for me, that's a month to know every detail.)

Wouldn't it be nice, if we have:

	strings contains 'My string' 'My' && echo "yes" || "no"
	
## yes, util-shell gives you this.

util-shell contains 4 basic binary that solves the more difficult part of shell.

####`math`: bash does not have floating point arithmetic. Only integers. But now you have a calculater!

```
	math '+' 4.5 1.5
	math '-' 5
	math '-' -5
	math '-' 10 5
	math '*' 1.2 4
	math '/' 3.3 3
	math ceil 3.3
	math floor 3.3
	math sqrt 4
	math sqrt 5
	math log 4
	math pow 10 2
	math '<' 1.1 2.2 && echo yes || echo no
	math '>=' 1.1 2.2 && echo yes || echo no
	math pi
```

####`array`: bash syntax for array is very ugly. If you read it, you probably will for get about it tomorrow.

```
	arr=$(array new 1 2 3 4)
	array len "$arr"
	array get 1 "$arr"
	array get 2 "$arr"
	array set 0 3 "$arr"
	
	another_arr=$(array delete 0 "$arr") // store the changes
	
	array set "abc" 3 "$arr"
	array has key 0 "$arr" && echo yes || echo no
	array has key 4 "$arr" && echo yes || echo no
	array has value 4 "$arr" && echo yes || echo no
	array has value 5 "$arr" && echo yes || echo no
	array list "$arr"
```
	
another array goodies:

```
	ls | array new // this is an array of your current directory
```
	
####`associative array (or map)`: bash syntax for this is even uglier; you can't even pass it to a function in bash.
(array is ordered; map is not.)

pass map to a function(note: array is also passable):

```
	map=$(map new k1:hello "k2:new world")
	a_function "$map"
	map len "$map"
	map get k1 "$map"
	map get k3 "$map"
	map delete k3 "$map"
	map delete k1 "$map"
	map set k4 213123 "$map"
	map list "$map"
	map has key k1 "$map" && echo yes || echo no
	map has key k3 "$map" && echo yes || echo no
	map has value "new world" "$map" && echo yes || echo no
	map has value "not exist" "$map" && echo yes || echo no
	
	a_function() {
		map len "$1"
	}
```
	
####`string`: do you rememeber how to get the string length? ${#string} is it readable?

```
	string len "你好" // string len "$string"
	string contains abc c && echo yes || no
	string count aabbccaabbcc bb
	string has-prefix abcde ab && echo yes || no
	string index abc b
	string repeat abc 3
	string to-lower ABc
	string to-upper Abc
```
	
want to split string and create array?

```
	arr=$(string split a,b,c,d,e , | array new)
```

#let's write readable code now
