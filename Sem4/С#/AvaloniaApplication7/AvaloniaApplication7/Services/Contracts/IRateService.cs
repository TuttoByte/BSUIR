using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using NbrbAPI.Models;

namespace AvaloniaApplication7.Services.Contracts;

public interface IRateService
{
    Task<IEnumerable<Rate>?> GetRates(DateTime startDate);
}