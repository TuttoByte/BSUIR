namespace AvaloniaApplication7.Services.Contracts;

public interface ICalculator
{
    public double Calculate(double value1, double value2);
    public double GetResult();
    public double GetPreviousResult();
    public IOperation GetCurrentOperation();
    public void SetCurrentOperation(string name);
}