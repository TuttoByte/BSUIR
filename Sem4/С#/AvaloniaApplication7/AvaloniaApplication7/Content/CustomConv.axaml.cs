using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Net.Http;
using System.Runtime.InteropServices.JavaScript;
using System.Threading.Tasks;
using Avalonia;
using Avalonia.Controls;
using Avalonia.Interactivity;
using Avalonia.Markup.Xaml;
using Avalonia.Media;
using Avalonia.Threading;
using AvaloniaApplication7.Services.Contracts;
using AvaloniaApplication7.Services.Models;
using NbrbAPI.Models;

namespace AvaloniaApplication7.Content;

public partial class CustomConv : UserControl
{
    readonly HashSet<String> _requestedRates = new HashSet<String>() {"RUB", "EUR", "USD", "CHF", "CNY", "GBP"};
    private IRateService _rateService;
    
    private int _selectedRate = -1;
    private bool isLeft = true;
    
    public ObservableCollection<Rate> Rates { get; set; } = new();
    
    public CustomConv()
    {
        InitializeComponent();
        DataContext = this;
        _rateService = new RateService();
    }
    
    public void OnLoadButtonClicked(object? sender, RoutedEventArgs e)
    {
        var currentDate = CalendarDatePicker.DisplayDate.Date;
        Console.WriteLine($"Current date: {currentDate}");

        Task.Run(() =>
        {
            var rates =   _rateService.GetRates(currentDate).Result;
            if (rates == null)
            {
                return;
            }
            rates = rates.Select(r => r)
                .Where(r => _requestedRates.Contains(r.Cur_Abbreviation))
                .ToList();
      
            Rates.Clear();
            foreach (var rate in rates)
            {
                Rates.Add(rate);
            }
        });

    }
    
    public void OnSelectionClicked(object? sender, SelectionChangedEventArgs e)
    {
        _selectedRate = ListBox.SelectedIndex;
        Console.WriteLine($"Selected rate: {_selectedRate}");
        setConvertedCurrentcy();
    }


    private void setConvertedCurrentcy()
    {
        if (_selectedRate == -1)
        {
            return;
        }
        
        if (isLeft)
        {
            Currency1.Text = "BYN";
            Currency2.Text = Rates[_selectedRate].Cur_Abbreviation;
        }
        else
        {
            Currency1.Text = Rates[_selectedRate].Cur_Abbreviation;
            Currency2.Text = "BYN";
        }
        
    }

    private void ConvertButtonClicked(object? sender, RoutedEventArgs e)
    {
        if (_selectedRate == -1)
        {
            return;
        }
        if (isLeft)
        {
            convertFromBYN();
        }
        else
        {
            convertToBYN();
        }
    }


    public void convertFromBYN()
    {
        var currentRate = Rates[_selectedRate];

        decimal leftValue;
        string? leftValueString = FirsVal.Text;
        bool result = decimal.TryParse(leftValueString, out leftValue);
        if (!result)
        {
            ConvertError();
            return;
        }


        var converted = leftValue * currentRate.Cur_Scale / currentRate.Cur_OfficialRate;
        SecondVal.Text =  converted.ToString();
    }

    public void convertToBYN()
    {
        
        var currentRate = Rates[_selectedRate];
        decimal leftValue;
        string? leftValueString = FirsVal.Text;
        bool result = decimal.TryParse(leftValueString, out leftValue);
        if (!result)
        {
            ConvertError();
            return;
        }


        var converted = leftValue * currentRate.Cur_OfficialRate/ currentRate.Cur_Scale;
        SecondVal.Text =  converted.ToString();
    }

    public void OnSwapCurrency(object? sender, RoutedEventArgs e)
    {
        isLeft = !isLeft;
        (FirsVal.Text, SecondVal.Text) = (SecondVal.Text, FirsVal.Text);
        (Currency1.Text, Currency2.Text) = (Currency2.Text, Currency1.Text);
    }
    
    
    private void ConvertError(){
        Task.Run(() =>
        {
            Dispatcher.UIThread.InvokeAsync((() =>
            {
                FirsVal.Foreground = Brushes.Red;
                ConvertTo.IsEnabled = false;
            }));
            
            Task.Delay(2000).Wait();
            
            Dispatcher.UIThread.InvokeAsync((() =>
            {
                FirsVal.Foreground = Brush.Parse("#0982e9");
                ConvertTo.IsEnabled = true;
            }));
        });
    }
}