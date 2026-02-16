
namespace Lab2.Services;

public class StreamService <T> where T : unmanaged
{


    [Obsolete("Obsolete")]
    public async Task WriteToStreamAync(Stream stream, IEnumerable<T> data)
    {
        var enumerable = data as T[] ?? data.ToArray();
        var byteArray = new List<Byte>();
        foreach (var e in enumerable)
        {
            var bout = ByteFormatter<T>.ObjectToByteArray(e);
            byteArray.AddRange(bout);
        }

        await Task.Delay(3000);
        
        await stream.WriteAsync(byteArray.ToArray(), 0, byteArray.Count);
        
    }


    [Obsolete("Obsolete")]
    public async Task CopyFromStreamAsync(Stream stream, string fyleName,  IProgress<string> progress)
    {

        using StreamWriter writer = new StreamWriter(fyleName);
        var byteArray = new  Byte[stream.Length];
        await stream.ReadAtLeastAsync(byteArray, 0, false);
        var data = ByteFormatter<T>.ByteArrayToIEnum(byteArray);
        writer.Write(data);

        
    }

    public async Task<int> GetStatisticsAsync(string fileName, Func<T, bool> filter)
    {
        int cоunter = 0; 
        
        using FileStream reader = new FileStream(
            fileName,
            FileMode.Open,
            FileAccess.Read,
            FileShare.Read,
            bufferSize: default,
            useAsync:true
            );
        var data = new Byte[reader.Length];
        await reader.ReadExactlyAsync(data);

        var dataType = ByteConverter.Convert<T>(data);
        foreach (var e in dataType)
        {
            if (filter(e))
            {
                cоunter++;
            }
        }


        return cоunter;
    }
    
}
