namespace AvaloniaApplication7.Services.Models;

public class Memory
{
    private double _memoryBucket = 0;

    public void LoadToMemory(double memory)
    {
        _memoryBucket = memory;
    }

    public void MemoryAdd(double memory)
    {
        _memoryBucket += memory;
    }

    public void MemorySubtract(double memory)
    {
        _memoryBucket -= memory;
    }

    public void MemoryClear()
    {
        _memoryBucket = 0;
    }

    public double MemoryRecall()
    {
        return _memoryBucket;
    }
    
}
