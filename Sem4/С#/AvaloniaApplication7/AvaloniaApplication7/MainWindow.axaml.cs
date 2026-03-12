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

        public MainWindow()
        {
            InitializeComponent();
            
            _menuTransform = Sidebar.RenderTransform as TranslateTransform;
            _calculus = new Calculus();
            _main = new MainPage();
            MainContent.Content = _main;
            ButtonCalculus.Classes.Remove("active");
            ButtonMain.Classes.Add("active");
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
            MainContent.Content = _calculus;
            ToggleSidebar_Click(sender, e);
       
        }

        private void ShowMain(object? sender, RoutedEventArgs e)
        {
            PageTitle.Text = "Main";
            ButtonCalculus.Classes.Remove("active");
            ButtonMain.Classes.Add("active");
            MainContent.Content = _main;
            ToggleSidebar_Click(sender, e);
        }
    }
}