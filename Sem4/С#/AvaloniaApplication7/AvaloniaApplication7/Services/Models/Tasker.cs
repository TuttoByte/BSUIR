using System;
using System.Linq.Expressions;
using Avalonia.Controls;
using AvaloniaApplication7.Services.classes;
using AvaloniaApplication7.Services.Contracts;

namespace AvaloniaApplication7.Services.Models;

public enum TaskerState
{
    WaitingForFirst,
    WaitingForSecond,
}


public class Tasker
{
    
    
   public delegate void OperandHandler(string operand);
   public delegate void OperationHandler(string operation);


    public event OperandHandler? OperandHandlerNotify;
    public event OperationHandler? OperationHandlerNotify;
        
    private TaskerState _state = TaskerState.WaitingForFirst;
    
    private readonly Calculator _calculator = new Calculator([new Sum(), new Subtract(),new Divide(), new Mul(), new Sqrt(), new Rev(), new Sqr(), new Percent()]);
    private readonly IOperand[] _operands = { new Operand(), new Operand() };
    


    public void SetOperand(string a)
    {
        IOperand currentOperand = _operands[(int)_state];
        currentOperand.Append(a);
        OperandHandlerNotify?.Invoke(currentOperand.GetString());
        
        Console.WriteLine(_state.ToString());
    }

    public void SetOperation(string operation)
    {
        
        if(_state == TaskerState.WaitingForFirst)
        {
            IOperand currentOperand = _operands[(int)_state];
            _calculator.SetCurrentOperation(operation);
            _state = TaskerState.WaitingForSecond;
            OperationHandlerNotify?.Invoke( currentOperand.GetString() + " "  + _calculator.GetOperation(operation).Signature());
            OperandHandlerNotify?.Invoke("");
              
        }
        Console.WriteLine(_state.ToString());
    }
    
    public void SetUnarOperation(string operation)
    {

        IOperand currentOperand = _operands[(int)_state];
        if (_state == TaskerState.WaitingForFirst)
        {
          
            _calculator.SetCurrentOperation(operation);
            _calculator.Calculate(currentOperand.GetNumerical());
            Console.WriteLine(_operands[(int)_state].GetString());
            _resetOperands();
            var result = _calculator.GetResult();
            _operands[0].SetString(result.ToString());
            OperandHandlerNotify?.Invoke(_operands[0].GetString());
        }
        
    }

    
    // TODO move to caclcualtor
    public void SetSignChangeOperator()
    {
        IOperand currentOperand = _operands[(int)_state];
        double num =  currentOperand.GetNumerical();
        currentOperand.SetString((num * -1).ToString());
        OperandHandlerNotify?.Invoke(currentOperand.GetString());
    }
    
    public double Perform()
    {
        
        
        if (_state == TaskerState.WaitingForSecond)
        {
            IOperand currentOperand = _operands[(int)_state];
            
            _calculator.Calculate(_operands[0].GetNumerical(), currentOperand.GetNumerical());
            _state = TaskerState.WaitingForFirst;
            Console.WriteLine(_state.ToString());
            _resetOperands();
            
            var result = _calculator.GetResult();
            _operands[0].SetString(result.ToString());
            OperandHandlerNotify?.Invoke(_operands[0].GetString());
            OperationHandlerNotify?.Invoke("");
            
            
            return result;
        }

        return 0;
    }

    private void _resetOperands()
    {
        foreach (var operand in _operands)
        {
            operand.SetString("");
        }
    }


    public void ClearAll()
    {
        _state = TaskerState.WaitingForFirst;
        _resetOperands();
        OperandHandlerNotify?.Invoke("");
        OperationHandlerNotify?.Invoke("");
    }

    public void ClearEntry()
    {
        if (_state == TaskerState.WaitingForSecond)
        {
            _operands[(int)_state].SetString("");
            OperandHandlerNotify?.Invoke(_operands[(int)_state].GetString());
        }
    }

    public void EraseCurrentOperand()
    {
        
        IOperand currentOperand = _operands[(int)_state];
        currentOperand.PopEnd();
        OperandHandlerNotify?.Invoke(currentOperand.GetString());
    }

    public void PoinAdd()
    {
        IOperand currentOperand = _operands[(int)_state];
        currentOperand.Append(".");
        OperandHandlerNotify?.Invoke(currentOperand.GetString());
    }

    public double GetCurrentOperand()
    {
        return _operands[(int)_state].GetNumerical();
    }

    public void SetCurrentOperand(double value)
    {
        
        _operands[(int)_state].SetString(value.ToString());
        OperandHandlerNotify?.Invoke(_operands[(int)_state].GetString());
    }


}