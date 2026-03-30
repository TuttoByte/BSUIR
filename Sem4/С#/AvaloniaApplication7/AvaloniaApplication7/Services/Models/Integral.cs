using System;
using System.ComponentModel;
using System.Threading;
using Avalonia.Threading;

namespace AvaloniaApplication7.Services.Models;

public class Integral
{
    
    private const int TimeStop = 100000;
    public double MaxValue  { get; set; }
    public double CurrentValue { get; set; }
    
    public Integral(int maxValue)
    {
        CurrentValue = 0;
        MaxValue = maxValue;
    }
    public double Calculate(Func<double, double> func, double a, double b,  double n, IProgress<double>? indicator, CancellationToken token)
    {
        double res = 0;
        double h = (b - a) / n;
        for (double i = 0; i < n; i+=0.001) {

            if (token.IsCancellationRequested)
            {
                return res;
            }
            
            res += func(a + h/2 +i *h);
           
                CurrentValue += 0.001;
                indicator?.Report(CurrentValue);
                Stop();
        }
        return res;
    }


    public void Reset()
    {
        CurrentValue = 0;
    }
    private void Stop()
    {
        for (int i = 0; i < TimeStop; i++)
        {
            int nil = 2 * 9;
        }
    }
    
}