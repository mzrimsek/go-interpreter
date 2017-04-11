let reduce = fn(arr, initial, f) { 
    let iter = fn(arr, result) { 
        if (len(arr) == 0) { 
            result
        } else { 
            iter(tail(arr) , f(result , first (arr)))
        } 
    };
    iter(arr, initial); 
};

let sum = fn(arr) { 
    reduce(arr, 0, fn(initial, el) { 
        initial + el 
    }); 
};

let array = [1, 2, 3, 4, 5];
let sum = sum(array);
putln("Sum: " + sum);
putln(push(array, 6));

let addNums = fn(x, y) {
    return x + y;
} 

let a = 3;
if(true) {
    let a = a + 2;
    putln(a);
}
putln(a);
putln();

let fizz = fn(num) { num % 3 == 0 };
let buzz = fn(num) { num % 5 == 0 };
let num = 1;
while(num <= 100) {
    if(fizz(num)) {
        put("fizz");
    }
    if (buzz(num)) {
        put("buzz");
    }
    if(!fizz(num) && !buzz(num)) {
        put(num);
    }
    putln();
    ++num;
}

let a = "hello";
putln(substr(a, 0, 3) * 3);
