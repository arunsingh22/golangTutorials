package main

// Use the pattern when your class has a massive conditional statement
// that switches between different variants of the same algorithm.
// The Strategy pattern lets you do away with such a conditional by extracting all
// algorithms into separate classes, all of which implement the same interface.
// The original object delegates execution to one of these objects,
//  instead of implementing all variants of the algorithm.
