
using CarSeller.Presistence.Repository;

class FakeUnitOfWork : IUnitOfWork{
    
    
    private IRepository<CarInfo> _carInfoRepository;
    private IRepository<Seller> _sellerInfoRepository;



    public FakeUnitOfWork()
    {
        _carInfoRepository = new FakeCarInfoRepository();   
        _sellerInfoRepository = new FakeSellerRepository();
    }   
    
    
    
    public IRepository<CarInfo> CarInfoRepository  => _carInfoRepository;
    public IRepository<Seller> SellerRepository => _sellerInfoRepository;
    public Task SaveAllAsync()
    {
        return Task.CompletedTask;
    }

    public Task DeleteDataBaseAsync()
    {
        return Task.CompletedTask;
    }

    public Task CreateDataBaseAsync()
    {
        return Task.CompletedTask;
    }
}
    
    
    
    
    
    