using System.Collections.Generic;
using AvaloniaApplication7.Services.Contracts;

namespace AvaloniaApplication7.Services.Models;

public class Calculator : ICalculator
{
    Dictionary<string, IOperation> _operations;
    private double _result, _previousResult;
    private IOperation _currentOperation;

    public Calculator(IOperation[] operations)
    {
        _operations = new Dictionary<string, IOperation>();
        foreach (var op in operations)
        {
            _operations[op.GetType().Name] = op;
        }
    }
    
    public double Calculate(double a, double b)
    {
        _result = GetCurrentOperation().Proceed(a, b);
        _previousResult = _result;
        return _result;
    }

    public double GetResult()
    {
        return _result;
    }

    public double GetPreviousResult()
    {
        return _previousResult;
    }

    public IOperation GetCurrentOperation()
    {
        return _currentOperation;
    }
    
    public void SetCurrentOperation(string name)
    {
        _currentOperation = _operations[name];
    }

}


