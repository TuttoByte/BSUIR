
using System.Collections.Immutable;
using System.Runtime.InteropServices;
using System.Text;
using System.Text.Json.Serialization;
using Newtonsoft.Json;


namespace Lab2.Services;

public class StreamService <T>
{

    
    public async Task WriteToStreamAync(Stream stream, IEnumerable<T> data, IProgress<string> progress)
    {
        var byteArray = new List<Byte>();
        foreach (var e in data)
        {
            var byteObj = Encoding.UTF8.GetBytes(JsonConvert.SerializeObject(e));
            var byteObjSize = byteObj.Length;
            byteArray.AddRange(BitConverter.GetBytes(byteObjSize));
            byteArray.AddRange(byteObj);
        }
        
        progress.Report($"Write start in {Thread.CurrentThread.ManagedThreadId}\n");
        await stream.WriteAsync(byteArray.ToArray(), 0, byteArray.Count);
        await Task.Delay(3000);
        progress.Report($"Write ended in {Thread.CurrentThread.ManagedThreadId} \n");
    }
    
    public async Task CopyFromStreamAsync(Stream stream, string fileName,  IProgress<string> progress)
    {

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
            useAsync:true
            );
        if (reader.CanSeek) reader.Seek(0, SeekOrigin.Begin);

        int stepInFile = 0;
        var size = new byte[4];
        var element =new byte[reader.Length];
        while (stepInFile < reader.Length)
        { 
            await reader.ReadExactlyAsync(size, 0, 4);
            var objSize = BitConverter.ToInt32(size, 0);
            reader.Seek(stepInFile + 4, SeekOrigin.Begin);
            await reader.ReadExactlyAsync(element, 0, objSize);
            stepInFile += objSize + 4;
            reader.Seek(stepInFile, SeekOrigin.Begin);
            
            var obj  =  JsonConvert.DeserializeObject<T>(Encoding.UTF8.GetString(element.Take(objSize).ToArray()));
            if (filter(obj))
            {
                cоunter++;
            }
        }
        return cоunter;
    }
    
    
    
    
}
