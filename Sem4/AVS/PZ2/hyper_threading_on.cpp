#include <iostream>
#include <fstream>
#include <cmath>
#include <iomanip>
#include <chrono>
#include <vector>
#include <omp.h>
#include <sched.h>
#include <unistd.h>
#include "functions/functions.h"


struct IterationResult {
    int thread_id;
    double x;
    long long time_us;
    int k;
};


int main(){
    double a, b, h, epsilon;

    std::cout << "Input a: "; 
    std::cin >> a;
    std::cout << "Input b: "; 
    std::cin >> b;
    std::cout << "Input h: "; 
    std::cin >> h;
    std::cout << "Input epsilon: ";
    std::cin >> epsilon;

    // std::ofstream csv("data_smt_on.csv");
    // csv << "Time,Iterations" << std::endl;

    std::cout << "\n" << std::setw(10) << "x" << std::setw(15) << "Y(x)" << std::setw(15) << "S(x)" << std::setw(10) << "n" << std::endl;
    std::cout << std::string(55, '-') << std::endl;

    int total_steps = static_cast<int>((b - a) / h) + 1;
    int num_threads = omp_get_max_threads();
    int logical_cores = (int)sysconf(_SC_NPROCESSORS_ONLN);

    std::vector<std::vector<IterationResult>> results(num_threads);

    #pragma omp parallel num_threads(num_threads)
    {
        int tid = omp_get_thread_num();

        cpu_set_t cpuset;
        CPU_ZERO(&cpuset);
        CPU_SET(tid % logical_cores, &cpuset);
        sched_setaffinity(0, sizeof(cpu_set_t), &cpuset);

        int items_per_thread = total_steps / num_threads;
        int start_idx = tid * items_per_thread;
        int end_idx = (tid == num_threads - 1) ? total_steps : start_idx + items_per_thread;

        results[tid].reserve((end_idx - start_idx) * 10000);

        for (int i = start_idx; i < end_idx; ++i) {
            double x = a + i * h;
            if (x > b) continue;

            double y_x = (fpu_sqr(fpu_cos(x)) - 1);
            double s_x = 0;
            int n = -1;

            for (int k = 1; k <= 10000; k++) {
                auto t0 = std::chrono::high_resolution_clock::now();
                
                volatile double temp = (pow(-1, i) * pow(fpu_mul(2, x), i)) / fpu_factorial(2 * i);
                s_x = fpu_add(s_x, temp);

                auto t1 = std::chrono::high_resolution_clock::now();
                
                long long duration = std::chrono::duration_cast<std::chrono::microseconds>(t1 - t0).count();

                if ((fabs(s_x - y_x) < epsilon) && (n == -1)) {
                    n = k;
                }

                results[tid].push_back({tid, x, duration, k});

            }
            #pragma omp critical
            {
                std::cout << std::fixed << std::setprecision(8) << std::setw(10) << x
                << std::setw(15) << y_x << std::setw(15) << s_x << std::setw(10) << n << std::endl;
            }  
            
        }
    }

    std::cout << "Writing start!\n";

    std::ofstream csv("data_ht_on.csv");
    csv << "Thread,X,Time_us,Iteration\n";
    for (const auto& t_vec : results) {
        for (const auto& r : t_vec) {
            csv << r.thread_id << "," << r.x << "," << r.time_us << "," << r.k << "\n";
        }
    }

    csv.close();
}