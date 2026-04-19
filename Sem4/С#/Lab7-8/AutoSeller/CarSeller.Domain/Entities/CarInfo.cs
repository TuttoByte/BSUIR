namespace CarSeller.Domain.Entities;

public class CarInfo: Entity
{
    private CarInfo() { }


    public CarInfo(string brand, string model)
    {
        Brand = brand;
        Model = model;
    }
    
    
    public int? SellerId { get; set; }
    
    public double Milleage { get; set; }
    public double Price { get; set; }
    public DateTime MadeYear { get; set; }
    public UInt32 OwnerNumbers {get; set;}
    public string Brand {get; set;}
    public string Model {get; set;}
    
    public void AddToSeller(int sellerId)
    {
        if (sellerId <= 0)
        {
            return;
        }
        SellerId = sellerId;
    }
    
}