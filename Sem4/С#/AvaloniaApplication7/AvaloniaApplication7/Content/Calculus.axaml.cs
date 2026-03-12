using System;
using Avalonia;
using Avalonia.Controls;

using Avalonia.Interactivity;
using Avalonia.Markup.Xaml;
using AvaloniaApplication7.Services.Models;

namespace AvaloniaApplication7.Content;

public partial class Calculus : UserControl
{
    
    private Tasker _tasker;
    private Printer _printer;
    private Memory _memory;
    
    public Calculus()
    {
        InitializeComponent();
        _tasker = new Tasker();
        _memory = new Memory();
        _printer = new Printer(MainInput, SecondaryInput);
        _tasker.OperandHandlerNotify += _printer.PrintOnMain;
        _tasker.OperationHandlerNotify += _printer.PrintOnSecond;
    }
    
    public void OnNumericButtonClick(object? sender, RoutedEventArgs e)
    {
        var button = sender as Button;
        
        Console.WriteLine(button!.Name);

        _tasker.SetOperand(button.Tag!.ToString()!);
    }

    public void OnOperationButtonClick(object? sender, RoutedEventArgs e)
    {
        var button = sender as Button;
        
        Console.WriteLine(button!.Name);
        _tasker.SetOperation(button.Tag!.ToString()!);
    }

    public void OnUnaryOperatorButtonClick(object? sender, RoutedEventArgs e)
    {
        var button = sender as Button;
        
        Console.WriteLine(button!.Name);
        _tasker.SetUnarOperation(button.Tag!.ToString()!);
    }

    public void OnResultButtonClick(object? sender, RoutedEventArgs e)
    {
        var button = sender as Button;
        
        Console.WriteLine(button!.Name);
        Console.WriteLine(_tasker.Perform());
    }

    public void OnClearButtonClick(object? sender, RoutedEventArgs e)
    {
        _tasker.ClearAll();
    }
    public void OnClearEntryButtonClick(object? sender, RoutedEventArgs e)
    {
     
        _tasker.ClearEntry();
    }
    
    public void OnEraseButtonClick(object? sender, RoutedEventArgs e)
    {
        _tasker.EraseCurrentOperand();
    }
    
    
    public void OnChangeSignButtonClick(object? sender, RoutedEventArgs e)
    {
        _tasker.SetSignChangeOperator();
    }

    public void OnPointButtonClick(object? sender, RoutedEventArgs e)
    {
        _tasker.PoinAdd();
    }



    public void OnMemorySetButtonClick(object? sender, RoutedEventArgs e)
    {
        double current = _tasker.GetCurrentOperand();
        _memory.LoadToMemory(current);
    }

    public void OnMemoryAddButtonClick(object? sender, RoutedEventArgs e)
    {
        double current = _tasker.GetCurrentOperand();
        _memory.MemoryAdd(current);
    }
    public void OnMemorySubButtonClick(object? sender, RoutedEventArgs e)
    {
        double current = _tasker.GetCurrentOperand();
        _memory.MemorySubtract(current);
    }
    
    public void OnMemoryRecallButtonClick(object? sender, RoutedEventArgs e)
    {
        _tasker.SetCurrentOperand(_memory.MemoryRecall());
    }
    
    
    public void OnMemoryClearButtonClick(object? sender, RoutedEventArgs e)
    {
        _memory.MemoryClear();
    }
    
}