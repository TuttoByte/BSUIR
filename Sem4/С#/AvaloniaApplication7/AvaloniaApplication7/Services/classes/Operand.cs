using System.Globalization;
using AvaloniaApplication7.Services.Contracts;

namespace AvaloniaApplication7.Services.classes;

public class Operand: IOperand
{
    
    
    private string? _innerValue = "";
    private double _value = 0 ;
    
    public double GetNumerical()
    {
        if (_innerValue == "")
        {
            _value = 0;
            return _value;
        }
        
        var  ok = double.TryParse(_innerValue, out _value);
        if (!ok)
        {
            throw new System.Exception("Invalid number");
        }
        return _value;
    }

    public void SetNumerical(double value)
    {
        _value = value;
        _innerValue = value.ToString(CultureInfo.InvariantCulture);
    }

    public string GetString()
    {
        return _innerValue;
    }

    public void SetString(string value)
    {
        _innerValue = value;
    }
    
    public void Append(string value)
    {
        if (value == "." && _innerValue.Contains("."))
        {
            return;
        }
        _innerValue += value;
    }

    public void PopEnd()
    {
        if (_innerValue == "")
        {
            return;
        }
        _innerValue = _innerValue.Remove(_innerValue.Length - 1, 1);
    }
   
}