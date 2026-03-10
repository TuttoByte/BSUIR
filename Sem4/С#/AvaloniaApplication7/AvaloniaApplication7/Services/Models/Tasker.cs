using System;
using AvaloniaApplication7.Services.Contracts;

namespace AvaloniaApplication7.Services.Models;

public enum TaskerState
{
    WaitingForFirst,
    WaitingForOperation,
    WaitingForSecond,
    RedyToPerform
}

public class Tasker
{
    
    private TaskerState _state = TaskerState.WaitingForFirst;
    private readonly Calculator _calculator = new Calculator([new Sum(), new Subtract(),new Divide()]);
    
    private double _firstOperand, _secondOperand;
    


    public void SetOperand(double a)
    {
        switch (_state)
        {
            case TaskerState.WaitingForFirst:
                _firstOperand = a;
                _state++;
                break;
            case TaskerState.WaitingForSecond:
                _secondOperand = a;
                _state++;
                break;
        }
        
        
        Console.WriteLine(_state.ToString());
    }

    public void SetOperation(string operation)
    {
        if (_state == TaskerState.WaitingForOperation)
        {
            _calculator.SetCurrentOperation(operation);
            _state++;
        }
        Console.WriteLine(_state.ToString());
    }

    public double Perform()
    {
        if (_state == TaskerState.RedyToPerform)
        {
            _calculator.Calculate(_firstOperand, _secondOperand);
            _state = TaskerState.WaitingForFirst;
            Console.WriteLine(_state.ToString());
            return _calculator.GetResult();
        }

        return 0;
    }
    
}