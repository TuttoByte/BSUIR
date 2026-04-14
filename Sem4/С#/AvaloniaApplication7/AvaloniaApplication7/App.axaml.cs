using System;
using System.Security.Authentication.ExtendedProtection;
using Avalonia;
using Avalonia.Controls.ApplicationLifetimes;

using Avalonia.Markup.Xaml;
using AvaloniaApplication7.Services.Contracts;
using AvaloniaApplication7.Services.Models;
using Microsoft.Extensions.DependencyInjection;

namespace AvaloniaApplication7;

public partial class App : Application
{
    
    public override void Initialize()
    {
        AvaloniaXamlLoader.Load(this);
    }

    public override void OnFrameworkInitializationCompleted()
    {

        var uri = new Uri("https://api.nbrb.by/exrates/rates");
        
        if (ApplicationLifetime is IClassicDesktopStyleApplicationLifetime desktop)
        {
            desktop.MainWindow = new MainWindow();
            var services = new ServiceCollection();
            services.AddTransient<IDbService, SqLiteService>();
            services.AddHttpClient<IRateService, RateService>();
        }

        base.OnFrameworkInitializationCompleted();
    }
}