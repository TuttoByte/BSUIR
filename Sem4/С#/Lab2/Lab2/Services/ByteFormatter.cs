using System.Runtime.InteropServices;
using System.Runtime.Serialization.Formatters.Binary;

namespace Lab2.Services;

public static class ByteFormatter<T>
{

    [Obsolete("Obsolete")]
    public static byte[] ObjectToByteArray(object obj)
    {
        BinaryFormatter bf = new BinaryFormatter();
        using(var ms = new MemoryStream())
        {
            bf.Serialize(ms, obj);
            return ms.ToArray();
        }
    }

        
    [Obsolete("Obsolete")]
    public static T[] ByteArrayToIEnum (byte[] byteArray)
    {
        BinaryFormatter bf = new BinaryFormatter();
        using (var ms = new MemoryStream(byteArray))
        {
            object obj = bf.Deserialize(ms);
            return  (T[])obj;
            
        }

    } 
    
}


public static class ByteConverter
{
    public static T[] Convert<T>(byte[] input)
        where T: unmanaged
        => MemoryMarshal.Cast<byte, T>(input).ToArray();
}