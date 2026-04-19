using CarSeller.Presistence.Data;

namespace CarSeller.Presistence.Repository;

public class EfUnitOfWork: IUnitOfWork
{
    private readonly Lazy<IRepository<CarInfo>> _carInfoRepository;
    private readonly Lazy<IRepository<Seller>> _sellerInfoRepository;



    public EfUnitOfWork(AppDbContext dbContext)
    {
        _dbContext = dbContext;
        _carInfoRepository = new Lazy<IRepository<CarInfo>>(() =>
            new EfRepository<CarInfo>(dbContext));
        _sellerInfoRepository = new Lazy<IRepository<Seller>>(() =>
            new EfRepository<Seller>(dbContext));
    }   
    
    private readonly AppDbContext _dbContext;
    public IRepository<CarInfo> CarInfoRepository => _carInfoRepository.Value;
    public IRepository<Seller> SellerRepository => _sellerInfoRepository.Value;
    public async Task SaveAllAsync()
    {
        await _dbContext.SaveChangesAsync();
    }

    public async Task DeleteDataBaseAsync()
    {
        await _dbContext.Database.EnsureDeletedAsync();
    }

    public async Task CreateDataBaseAsync()
    {
        await _dbContext.Database.EnsureCreatedAsync();
    }
}