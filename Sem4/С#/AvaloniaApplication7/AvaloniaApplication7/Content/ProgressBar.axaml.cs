using System;
using System.Globalization;
using System.Runtime.InteropServices.JavaScript;
using System.Threading;
using System.Threading.Tasks;
using Avalonia;
using Avalonia.Controls;
using Avalonia.Interactivity;
using Avalonia.Markup.Xaml;
using Avalonia.Threading;
using AvaloniaApplication7.Services.classes;
using AvaloniaApplication7.Services.Models;

namespace AvaloniaApplication7.Content;



public partial class CustomProgressBar : UserControl
{
    private static CancellationTokenSource cancelTokenSource =  new CancellationTokenSource();
    private CancellationToken token;
    private const int MaxValue = 100;
    public Integral IntegralData =  new(MaxValue) ;


    public double CurrentValue
    {
        get => GetValue(CurrentValueProperty);
        set => SetValue(CurrentValueProperty, value);
    }
    
    public static readonly StyledProperty<double> CurrentValueProperty =
        AvaloniaProperty.Register<CustomProgressBar, double>(nameof(CurrentValue));
    
   
    public CustomProgressBar()
    {
        InitializeComponent();
        DataContext = this;
        Stop.IsEnabled =  false;
    }


    public void OnButtonStartClicked(object sender, RoutedEventArgs e)
    {
        Status.Text = "Вычисление";
        Start.IsEnabled = false;
        Stop.IsEnabled = true;
        cancelTokenSource.Dispose();
        cancelTokenSource =  new CancellationTokenSource();
        token = cancelTokenSource.Token;
        CurrentValue = 0;
        var progress = new Progress<double>(value =>
        {
            CurrentValue = value;
            Console.WriteLine(CurrentValue);
            
        });
        Task.Run(() =>
        {
            var result = IntegralData.Calculate(Functions.Sin, 0, 1, MaxValue, progress, token);

            if (!cancelTokenSource.IsCancellationRequested)
            {
                Dispatcher.UIThread.InvokeAsync(() =>
                {
                    Status.Text = $"Итог вычисления: \n {result.ToString("F3", CultureInfo.InvariantCulture)}";
                });
            }
            

        }, token);
        
    }

    
    
    
    public void OnButtonStopClicked(object sender, RoutedEventArgs e)
    {
        Status.Text = "Задание отменено";
        Start.IsEnabled = true;
        Stop.IsEnabled = false;
        cancelTokenSource.Cancel();
        IntegralData.Reset();
    }
}



