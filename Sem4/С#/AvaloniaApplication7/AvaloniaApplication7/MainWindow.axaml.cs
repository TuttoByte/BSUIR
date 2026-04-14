using System.Runtime.Serialization;
using Avalonia.Controls;
using Avalonia.Interactivity;
using Avalonia.Media;
using AvaloniaApplication7.Content;

namespace AvaloniaApplication7
{
    public partial class MainWindow : Window
    {
        private bool _sidebarOpen = false;
        private TranslateTransform _menuTransform;
        
        Calculus _calculus;
        MainPage _main;
        CustomProgressBar _progressBar;
        Picker _picker;
        CustomConv _conv;

        public MainWindow()
        {
            InitializeComponent();
            
            
            
            _menuTransform = Sidebar.RenderTransform as TranslateTransform;
            _calculus = new Calculus();
            _main = new MainPage();
            _progressBar = new CustomProgressBar();
            _picker = new Picker();
            _conv = new CustomConv();
            MainContent.Content = _progressBar;
            ButtonCalculus.Classes.Remove("active");
            ButtonMain.Classes.Remove("active");
            ProgressBar.Classes.Add("active");
            Picker.Classes.Remove("active");
            Converter.Classes.Remove("active");

        }

        private void ToggleSidebar_Click(object? sender, RoutedEventArgs e)
        {
            if (_sidebarOpen)
            {
         
                _menuTransform.X = -Sidebar.Width;
                OverlayBackground.IsVisible = false;
                OverlayBackground.Opacity = 0;
            }
            else
            {
                // Открыть меню
                _menuTransform.X = 0;
                OverlayBackground.IsVisible = true;
                OverlayBackground.Opacity = 1;
            }

            _sidebarOpen = !_sidebarOpen;
        }

        private void ShowCalculus(object? sender, RoutedEventArgs e)
        {
            PageTitle.Text = "Calculator";
            ButtonCalculus.Classes.Add("active");
            ButtonMain.Classes.Remove("active");
            ProgressBar.Classes.Remove("active");
            Picker.Classes.Remove("active");
            Converter.Classes.Remove("active");
            MainContent.Content = _calculus;
            ToggleSidebar_Click(sender, e);
       
        }

        private void ShowMain(object? sender, RoutedEventArgs e)
        {
            PageTitle.Text = "Main";
            ButtonCalculus.Classes.Remove("active");
            ButtonMain.Classes.Add("active");
            ProgressBar.Classes.Remove("active");
            Picker.Classes.Remove("active");
            Converter.Classes.Remove("active");
            MainContent.Content = _main;
            ToggleSidebar_Click(sender, e);
        }

        public void ShowProgressBar(object? sender, RoutedEventArgs e)
        {
            PageTitle.Text = "ProgressBar";
            ButtonCalculus.Classes.Remove("active");
            ButtonMain.Classes.Remove("active");
            Picker.Classes.Remove("active");
            ProgressBar.Classes.Add("active");
            Converter.Classes.Remove("active");
            MainContent.Content =  _progressBar;
            ToggleSidebar_Click(sender, e);
            
        }

        public void ShowPicker(object? sender, RoutedEventArgs e)
        {
            PageTitle.Text = "Picker";
            ButtonCalculus.Classes.Remove("active");
            ButtonMain.Classes.Remove("active");
            ProgressBar.Classes.Remove("active");
            Picker.Classes.Add("active");
            Converter.Classes.Remove("active");
            MainContent.Content =  _picker;
            ToggleSidebar_Click(sender, e);
        }
        
        public void ShowConverter(object? sender, RoutedEventArgs e)
        {
            PageTitle.Text = "Converter";
            ButtonCalculus.Classes.Remove("active");
            ButtonMain.Classes.Remove("active");
            ProgressBar.Classes.Remove("active");
            Picker.Classes.Remove("active");
            Converter.Classes.Add("active");
            MainContent.Content =  _conv;
            ToggleSidebar_Click(sender, e);
        }
    }
}