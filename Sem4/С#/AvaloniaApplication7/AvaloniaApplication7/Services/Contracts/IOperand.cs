namespace AvaloniaApplication7.Services.Contracts;

public interface IOperand
{
    public double GetNumerical();
    public void SetNumerical(double value);
    
    public string GetString();
    public void SetString(string value);
    public void Append(string value);
    public void PopEnd();

}