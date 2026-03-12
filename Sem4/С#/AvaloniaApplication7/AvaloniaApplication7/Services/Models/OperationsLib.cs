using System;
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

    public string Signature()
    {
        return "+";
    }
}

public class Subtract:IOperation 
{
    public double Proceed(double value1, double value2)
    {
        return value1 - value2;
    }
    
    public string Signature()
    {
        return "-";
    }
}

public class Divide :IOperation
{
    public double Proceed(double value1, double value2)
    {
        return value1 / value2;
    }
    
    public string Signature()
    {
        return "/";
    }
}


public class Mul :IOperation
{
    public double Proceed(double value1, double value2)
    {
        return value1 * value2;
      
    }
    
    public string Signature()
    {
        return "x";
    }
}


public class Sqrt: IOperation{
    
    public double Proceed(double value1, double value2)
    {
        return Math.Sqrt(value1);
      
    }
    
    public string Signature()
    {
        return "sqrt";
    }
}

public class Sqr: IOperation{
    
    public double Proceed(double value1, double value2)
    {
        return value1 * value1;

    }
    
    public string Signature()
    {
        return "sqr";
    }
}


public class Rev: IOperation{
    
    public double Proceed(double value1, double value2)
    {
        return 1 / value1;

    }
    
    public string Signature()
    {
        return "sqr";
    }
}

public class Percent : IOperation
{
    public double Proceed(double value1, double value2)
    {
        return value1 / 100;
    }

    public string Signature()
    {
        return "percent";
    }
}