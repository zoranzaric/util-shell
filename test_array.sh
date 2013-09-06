# map
map=$(map new k1:hello "k2:new world")
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

# array
arr=$(array new 1 2 3 4)
array len "$arr"
array get 1 "$arr"
array get 2 "$arr"
array set 0 3 "$arr"
array delete 0 "$arr"
array set "abc" 3 "$arr"
array has key 0 "$arr" && echo yes || echo no
array has key 4 "$arr" && echo yes || echo no
array has value 4 "$arr" && echo yes || echo no
array has value 5 "$arr" && echo yes || echo no
array list "$arr"
