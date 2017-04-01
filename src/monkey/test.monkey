let reduce = fn (arr , initial , f) { let iter = fn (arr , result) { if ( len (arr) == 0 ) { result } else { iter ( tail (arr) , f (result , first (arr))) ; } }; iter (arr , initial) ; };
let sum = fn (arr) { reduce (arr , 0 , fn (initial , el) { initial + el } ) ; };
let array = [1, 2, 3, 4, 5];
sum(array)