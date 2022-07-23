package com.codingwithme;

public class Main {
    public static void main(String[] args) {
        TaxCalculater calculator = getCalculator();
        calculator.calculateTax();

    }

    static TaxCalculater getCalculator() {
        return new TaxCalculator2019();
    }
}
