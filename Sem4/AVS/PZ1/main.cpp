
#include <iostream>
#include <cmath>
#include <sstream>
#include <fstream>
#include <iomanip>
#include <chrono>
#include "functions/functions.h"


int main(){
    double a,b,h,eps;
    std::cout<<"Input a"<<std::endl;
    std::cin>>a;
    std::cout<<"Input b"<<std::endl;
    std::cin>>b;
    std::cout<<"Input h"<<std::endl;
    std::cin>>h;
    std::cout<<"Input epsilon"<<std::endl;
    std::cin>>eps;


    std::ofstream csv("dataOn1.csv");
    csv << "Time,Iterations" << std::endl;

    std::cout << "\n" << std::setw(10) << "x" << std::setw(15) << "Y(x)" << std::setw(15) << "S(x)" << std::setw(10) << "n" << std::abort;
    std::cout <<std::endl <<std::string(55, '-') << std::endl;


    for (double x = a; x <= b; x += h){
        int n = 0;
        double s_x = 0, temp = 0;

        auto start = std::chrono::high_resolution_clock::now();
        double y_x = 2 * (fpu_sqr(fpu_cos(x)) - 1);
       
       for (long i = 0; i < 10000; i ++){
            temp = (pow(-1, i) * pow(fpu_mul(2, x), i)) / fpu_factorial(2 * i);
            if (std::isnan(temp)){
                temp = 0;
            }
            s_x = fpu_add(s_x, temp);

            

            if (fabs(s_x - y_x) < eps && n == 0){
                n = i;
                break;
            }

            if( i % 10 == 0){
                auto end = std::chrono::high_resolution_clock::now();
                std::chrono::duration<double> elapsed =  end - start;
                start = end;

                csv << elapsed.count() << "," <<i<<std::endl;
            }
       }


       std::cout << std::fixed << std::setprecision(4) << std::setw(10) << x
             << std::setw(15) << y_x << std::setw(15) << s_x << std::setw(10) << n << std::endl;
             n = 0;
    }

    csv.close();
    return 0;
}
