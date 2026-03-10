using System.Transactions;

namespace AvaloniaApplication7.Services.Contracts;

public interface IOperation
{
    public double Proceed(double value1, double value2);
}