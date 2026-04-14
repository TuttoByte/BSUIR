using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text.Json;
using System.Threading.Tasks;
using AvaloniaApplication7.Services.Contracts;
using NbrbAPI.Models;

namespace AvaloniaApplication7.Services.Models;

public class RateService :IRateService
{
    private static HttpClient httpClient = new HttpClient();
    
 
    public async Task<IEnumerable<Rate>?> GetRates(DateTime startDate)
    {
        var parametrs = $"https://api.nbrb.by/exrates/rates?ondate={startDate.Date.ToString("yyyy-MM-dd")}&periodicity=0";
        Console.WriteLine(parametrs);
        var rates = await httpClient.GetFromJsonAsync<IEnumerable<Rate>>( parametrs, new JsonSerializerOptions(JsonSerializerDefaults.Web));
        Console.WriteLine(rates);
        return rates;
    }
}