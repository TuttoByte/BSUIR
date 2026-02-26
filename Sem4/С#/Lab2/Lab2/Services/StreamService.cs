
using System.Collections.Immutable;
using System.Runtime.InteropServices;
using System.Text;
using System.Text.Json.Serialization;
using Newtonsoft.Json;
using System.Text.Json;
using JsonSerializer = System.Text.Json.JsonSerializer;

namespace Lab2.Services;

public class StreamService <T>
{

    
    public async Task WriteToStreamAync(Stream stream, IEnumerable<T> data, IProgress<string> progress)
    {
        progress.Report($"Enterred to write in {Thread.CurrentThread.ManagedThreadId}");
        progress.Report($"Write start in {Thread.CurrentThread.ManagedThreadId}");
        await JsonSerializer.SerializeAsync(stream, data);
        await stream.FlushAsync();
        await Task.Delay(3000);
        progress.Report($"Write ended in {Thread.CurrentThread.ManagedThreadId}");
    }
    
    public async Task CopyFromStreamAsync(Stream stream, string fileName,  IProgress<string> progress)
    {
        progress.Report($"Enterred to copy in {Thread.CurrentThread.ManagedThreadId}");
        using FileStream writer = new FileStream(
            fileName,
            FileMode.Create,
            FileAccess.Write,
            FileShare.None,
            4096,
            true
            );
     
        progress.Report($"Copy start in {Thread.CurrentThread.ManagedThreadId}\n");
        if (stream.CanSeek) stream.Seek(0, SeekOrigin.Begin);
        await stream.CopyToAsync(writer);
        
        progress.Report($"Copy ended in {Thread.CurrentThread.ManagedThreadId}\n");
       
        
    }

    public async Task<int> GetStatisticsAsync(string fileName, Func<T, bool> filter)
    {
        int cоunter = 0;

        using FileStream reader = new FileStream(
            fileName,
            FileMode.Open,
            FileAccess.ReadWrite,
            FileShare.Read,
            bufferSize: 4096,
            useAsync: true
        );
        
        var data = await JsonSerializer.DeserializeAsync<List<T>>(reader);
        
        foreach (var x1 in data)
        {
         if (filter(x1))
         {
             cоunter++;
         }   
        }
        return cоunter;
    }
}
