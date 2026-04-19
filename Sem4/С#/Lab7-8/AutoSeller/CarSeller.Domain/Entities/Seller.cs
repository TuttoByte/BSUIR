namespace CarSeller.Domain.Entities;

public class Seller: Entity
{
    private List<CarInfo> _carInfos = new();

    private Seller()
    {
    }

    public Seller(string name, string email, string phoneNumber)
    {
        Name = name;
        Email = email;
        PhoneNumber = phoneNumber;
    }
    
    public string Name { get; set; }
    public string Email { get; set; }
    public string PhoneNumber { get; set; }

    public IReadOnlyList<CarInfo> CarInfos
    {
        get => _carInfos.AsReadOnly();
    }
    
    
}