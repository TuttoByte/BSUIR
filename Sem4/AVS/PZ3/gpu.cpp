#include <CL/cl.h>
#include <iostream>
#include <vector>
#include <fstream>
#include <string>
#include <cmath>
#include <iomanip>
#include <chrono>

int main() {
    // const int N = 1000;
    // double a = 1.0, h = 0.001, epsilon = 1e-7;
    int max_k = 10000;

    float a, b, h, epsilon;

    std::cout << "Input a: "; 
    std::cin >> a;
    std::cout << "Input b: "; 
    std::cin >> b;
    std::cout << "Input h: "; 
    std::cin >> h;
    std::cout << "Input epsilon: ";
    std::cin >> epsilon;

    int N = static_cast<int>((b - a) / h);
    
    std::vector<float> x(N), y(N), s(N);
    std::vector<int> n(N);

    for (int i = 0; i < N; i++)
        x[i] = a + i * h;

    cl_platform_id platform;
    cl_device_id device;
    clGetPlatformIDs(1, &platform, nullptr);
    clGetDeviceIDs(platform, CL_DEVICE_TYPE_GPU, 1, &device, nullptr);

    char platformName[128], deviceName[128];
    clGetPlatformInfo(platform, CL_PLATFORM_NAME, 128, platformName, nullptr);
    clGetDeviceInfo(device, CL_DEVICE_NAME, 128, deviceName, nullptr);
    std::cout << "Platform: " << platformName << "\nDevice: " << deviceName << std::endl;

    cl_context context = clCreateContext(nullptr, 1, &device, nullptr, nullptr, nullptr);
    cl_queue_properties props[] = {0};
    cl_command_queue queue = clCreateCommandQueueWithProperties(context, device, props, nullptr);

    cl_mem d_x = clCreateBuffer(context, CL_MEM_READ_ONLY | CL_MEM_COPY_HOST_PTR,
                                sizeof(float)*N, x.data(), nullptr);
    cl_mem d_y = clCreateBuffer(context, CL_MEM_READ_WRITE, sizeof(float)*N, nullptr, nullptr);
    cl_mem d_s = clCreateBuffer(context, CL_MEM_READ_WRITE, sizeof(float)*N, nullptr, nullptr);
    cl_mem d_n = clCreateBuffer(context, CL_MEM_READ_WRITE, sizeof(int)*N, nullptr, nullptr);

    std::ifstream f("kernel.cl");
    if(!f.is_open()) { std::cerr << "Cannot open kernel.cl\n"; return 1; }
    std::string source((std::istreambuf_iterator<char>(f)), std::istreambuf_iterator<char>());
    const char* kernelSource = source.c_str();

    cl_program program = clCreateProgramWithSource(context, 1, &kernelSource, nullptr, nullptr);

    cl_int build_err = clBuildProgram(program, 1, &device, nullptr, nullptr, nullptr);
    if(build_err != CL_SUCCESS) {
    size_t log_size;
    clGetProgramBuildInfo(program, device, CL_PROGRAM_BUILD_LOG, 0, nullptr, &log_size);
    std::vector<char> build_log(log_size);
    clGetProgramBuildInfo(program, device, CL_PROGRAM_BUILD_LOG, log_size, build_log.data(), nullptr);
    std::cerr << "Build log:\n" << build_log.data() << std::endl;
    return 1;
}
    std::cout << "Kernel compiled OK\n";

    cl_kernel kernel = clCreateKernel(program, "compute_kernel", nullptr);

    size_t globalSize = N;

    std::vector<long long> times_us(max_k); 

    for (int k_iter = 1; k_iter <= max_k; ++k_iter) {
        int max_terms = k_iter;

        clSetKernelArg(kernel, 0, sizeof(cl_mem), &d_x);
        clSetKernelArg(kernel, 1, sizeof(cl_mem), &d_y);
        clSetKernelArg(kernel, 2, sizeof(cl_mem), &d_s);
        clSetKernelArg(kernel, 3, sizeof(cl_mem), &d_n);
        clSetKernelArg(kernel, 4, sizeof(float), &epsilon);
        clSetKernelArg(kernel, 5, sizeof(int), &max_terms);

        clFinish(queue); 

        auto t0 = std::chrono::high_resolution_clock::now();

        clEnqueueNDRangeKernel(queue, kernel, 1, nullptr, &globalSize, nullptr, 0, nullptr, nullptr);
        clFinish(queue);

        auto t1 = std::chrono::high_resolution_clock::now();

        times_us[k_iter-1] = std::chrono::duration_cast<std::chrono::microseconds>(t1 - t0).count();
    }

    std::ofstream csv("data_gpu.csv");
    csv << "Iteration,Time_us\n";
    for (int k = 0; k < max_k; ++k)
        csv << (k+1) << "," << times_us[k] << "\n";
    csv.close();

    clEnqueueReadBuffer(queue, d_y, CL_TRUE, 0, sizeof(float)*N, y.data(), 0, nullptr, nullptr);
    clEnqueueReadBuffer(queue, d_s, CL_TRUE, 0, sizeof(float)*N, s.data(), 0, nullptr, nullptr);
    clEnqueueReadBuffer(queue, d_n, CL_TRUE, 0, sizeof(int)*N, n.data(), 0, nullptr, nullptr);

    std::cout << "\n" << std::setw(10) << "x" << std::setw(15) << "Y(x)" << std::setw(15) << "S(x)" << std::setw(10) << "n" << std::endl;
    std::cout << std::string(55, '-') << std::endl;
    for (int i = 0; i < N; i++) {
        std::cout << std::fixed << std::setprecision(8) << std::setw(10) << x[i]
            << std::setw(15) << y[i] << std::setw(15) << s[i] << std::setw(10) << n[i] << std::endl;
    }

    clReleaseMemObject(d_x);
    clReleaseMemObject(d_y);
    clReleaseMemObject(d_s);
    clReleaseMemObject(d_n);
    clReleaseKernel(kernel);
    clReleaseProgram(program);
    clReleaseCommandQueue(queue);
    clReleaseContext(context);

    return 0;
}