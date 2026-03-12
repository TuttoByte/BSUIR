using Avalonia.Controls;

namespace AvaloniaApplication7.Services.Models;

public class Printer
{
    
    private TextBlock _mainZone, _secondaryZone;
    

    public Printer(TextBlock mainZone, TextBlock secondaryZone)
    {
        _mainZone = mainZone;
        _secondaryZone = secondaryZone;
    }
    
    public void PrintOnMain(string text)
    {
        _mainZone.Text = text;
    }
    
    public void PrintOnSecond(string text)
    {
        _secondaryZone.Text = text;
        
    }
}