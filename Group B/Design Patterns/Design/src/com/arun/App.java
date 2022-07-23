package com.arun;

public class App {
    public static void main(String[] args) throws Exception {
        Person p = new Person("Arun");
        p.sayHello();

        // Problems from design perspective
        // 1. Tightly coupled system, if I add/remove any field/method in Person class
        // the Main class will break.
        // 2. After making any change in perosn class, it needs to be rebuild thus it
        // also forces our Main class to recompile.

        // Application built this way is really hard to extend and maintaine as it leads to alot of cascading changes.
    }
}
