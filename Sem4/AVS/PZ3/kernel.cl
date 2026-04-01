

__kernel void compute_kernel(
    __global const float* x_vals,
    __global float* y_vals,
    __global float* s_vals,
    __global int* n_vals,
    const float epsilon,
    const int k)
{
    int idx = get_global_id(0);
    float x = x_vals[idx];
    float y_x, s_x;
    int n;
    if(k == 1){
        y_x = 2.0f * (cos(x) * cos(x) - 1);
        s_x = 0.0;
        n = -1;
    }
    else{
        y_x = y_vals[idx];
        s_x = s_vals[idx];
        n = n_vals[idx];
    }

    
    float c = 1.0f;

    for (int i = 1; i < 2 * k; i++){
        c *= i;
    }

    float temp = pow(-1.0f, k) * pow(2.0f * x, 2.0f * k) / c;
    s_x += temp;

    if (fabs(s_x - y_x) < epsilon && n == -1) {
        n = k;
    }

    y_vals[idx] = y_x;
    s_vals[idx] = s_x;
    n_vals[idx] = n;
}

