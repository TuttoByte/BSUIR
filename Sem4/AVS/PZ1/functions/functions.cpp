#include "functions.h"
#include <bitset>
#include <iostream>

double fpu_add(double a, double b){
    
    double result;
    asm volatile(
        "fldl %1\n\t"
        "fldl %2\n\t"
        "faddp\n\t"
        "fstpl %0\n\t"
        : "=m"(result)
        :"m"(a), "m"(b)
        : "st"
    );
    return result;
}


double fpu_mul(double a, double b){
    
    double result;
    asm volatile(
        "fldl %1\n\t"
        "fldl %2\n\t"
        "fmulp\n\t"
        "fstpl %0\n\t"
        : "=m"(result)
        :"m"(a), "m"(b)
        : "st"
    );
    return result;
}



double fpu_div(double a, double b){
    
    double result;
    asm volatile(
        "fldl %1\n\t"
        "fldl %2\n\t"
        "fdivp\n\t"
        "fstpl %0\n\t"
        : "=m"(result)
        :"m"(a), "m"(b)
        : "st"
    );
    return result;
}


double fpu_sub(double a, double b){
    double result;
    asm volatile(
        "fldl %1\n\t"
        "fldl %2\n\t"
        "fsubp\n\t"
        "fstpl %0\n\t"
        : "=m"(result)
        : "m"(a), "m"(b)
        : "st"
    );
    return result;
}

double fpu_cos(double a){
    double result;
    asm volatile(
        "fldl %1\n\t"
        "fsin\n\t"
        "fstpl %0\n\t"
        : "=m"(result)
        : "m"(a)
        : "st"
    );
    return result;
}

double fpu_sqr(double a){
    double result;
    asm volatile(
        "fldl %1\n\t"
        "fldl %2\n\t"
        "fmulp\n\t"
        "fstpl %0\n\t"
        :"=m"(result)
        :"m"(a), "m"(a)
        : "st"
    );

    return result;
}


double fpu_factorial(long a){


    if (a < 0){
        return 0;
    }

    if ( a == 0){
        return 1;
    }

   double result = 1.0;
   long temp = 1.0;


   __asm__ volatile(
    "mov %[a], %%eax\n\t"
    "fldl %[result]\n\t"
    "forl:\n\t"
    "fxch %%st(1)\n\t"
    "fildq %[temp]\n\t"
    "fxch %%st(1)\n\t"
    "fincstp\n\t"
    "fmulp\n\t"
    //"jmp endprog\n\t"
    //"jnz overfl\n\t"
    //"jmp next\n\t"
    //"overfl: \n\t"
    //"fsub %%st(0), %%st(0)\n\t"
    //"jmp endprog\n\t"


    "next: \n\t"
    "incl %[temp]\n\t"
    "cmpl %%eax, %[temp]\n\t"
    "jle forl\n\t"
    //"endprog:"
    "fstpl %[result]\n\t"
    : [result] "=m" (result)    
    : [a] "m" (a), [temp] "m" (temp)
    : "eax", "st"
   );

   return result;

}