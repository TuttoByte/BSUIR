using System;
using System.Windows.Input;
using Avalonia.Controls;
using Avalonia.Interactivity;
using AvaloniaApplication7.Services.Models;

namespace AvaloniaApplication7;

public partial class MainWindow : Window
{
    private Tasker _tasker;
    
    public MainWindow()
    {
        InitializeComponent();
        _tasker = new Tasker();
    }

    
    public void OnNumericButtonClick(object? sender, RoutedEventArgs e)
    {
        var button = sender as Button;
        
        Console.WriteLine(button!.Name);

        _tasker.SetOperand(double.Parse(button.Tag!.ToString()!));
    }

    public void OnOperationButtonClick(object? sender, RoutedEventArgs e)
    {
        var button = sender as Button;
        
        Console.WriteLine(button!.Name);
        _tasker.SetOperation(button.Tag!.ToString()!);
    }

    public void OnResultButtonClick(object? sender, RoutedEventArgs e)
    {
        var button = sender as Button;
        
        Console.WriteLine(button!.Name);
        Console.WriteLine(_tasker.Perform());
    }
}