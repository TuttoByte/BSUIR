namespace CarSeller.Domain.Entities;

public interface IUnitOfWork
{
    IRepository<CarInfo> CarInfoRepository { get; }
    IRepository<Seller> SellerRepository { get; }
    public Task SaveAllAsync();
    public Task DeleteDataBaseAsync();
    public Task CreateDataBaseAsync();
}