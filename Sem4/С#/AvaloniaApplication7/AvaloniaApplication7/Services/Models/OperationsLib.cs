using AvaloniaApplication7.Services.Contracts;

namespace AvaloniaApplication7.Services.Models;

public class OperationsLib
{
    
}

public class Sum :IOperation
{
    public double Proceed(double value1, double value2)
    {
        return value1 + value2;
    }
}

public class Subtract:IOperation 
{
    public double Proceed(double value1, double value2)
    {
        return value1 - value2;
    }
}

public class Divide :IOperation
{
    public double Proceed(double value1, double value2)
    {
        return value1 / value2;
    }
}