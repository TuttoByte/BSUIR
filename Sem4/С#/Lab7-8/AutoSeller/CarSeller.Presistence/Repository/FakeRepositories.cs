using System.Linq.Expressions;
using CarSeller.Presistence.Data;

namespace CarSeller.Presistence.Repository;

public class FakeCarInfoRepository: IRepository<CarInfo>
{
    List<CarInfo> _carInfos;

    public FakeCarInfoRepository()
    {

        _carInfos = new List<CarInfo>();
        int k = 1;
        for (int i = 1; i <= 2; i++)
        {
            for (int j = 0; j < 10; j++)
            {
                var carInfo = new CarInfo("Audi",$"SF7{k++}");
                carInfo.Id = i;
                _carInfos.Add(carInfo);
            }
        }
    }


    public Task<CarInfo?> GetByIdAsync(int id, CancellationToken cancellationToken = default, params Expression<Func<CarInfo, object>>[]? includeProperties)
    {
        throw new NotImplementedException();
    }

    public async Task<IReadOnlyList<CarInfo>> ListAllAsync(CancellationToken cancellationToken = default)
    {
        return await Task.Run((() => _carInfos));
    }

    public async Task<IReadOnlyList<CarInfo>> ListAsync(Expression<Func<CarInfo, bool>>? filter, CancellationToken cancellationToken = default,
        params Expression<Func<CarInfo, object>>[] includesProperties)
    {
        var data = _carInfos.AsQueryable();
        return data.Where(filter).ToList();
    }

    public Task AddAsync(CarInfo entity, CancellationToken cancellationToken = default)
    {
        throw new NotImplementedException();
    }

    public Task UpdateAsync(CarInfo entity, CancellationToken cancellationToken = default)
    {
        throw new NotImplementedException();
    }

    public Task DeleteAsync(CarInfo entity, CancellationToken cancellationToken = default)
    {
        throw new NotImplementedException();
    }

    public Task<CarInfo?> FirstOrDefaultAsync(Expression<Func<CarInfo, bool>> filter, CancellationToken cancellationToken = default)
    {
        throw new NotImplementedException();
    }
}

public class FakeSellerRepository : IRepository<Seller>
{
    private List<Seller> _sellersInfos = new List<Seller>();

    public FakeSellerRepository ()
    {
        _sellersInfos = new List<Seller>();

        var seller = new Seller("Jordan", "jordanvellini@gmail.com", "+23 344 56 78 123");
        seller.Id = 1;
        _sellersInfos.Add(seller);
        
        var seller1 = new Seller("Mike", "mikevellini@gmail.com", "+23 344 56 78 123");
        seller.Id = 2;
        _sellersInfos.Add(seller1);
    }

    public Task<Seller?> GetByIdAsync(int id, CancellationToken cancellationToken = default, params Expression<Func<Seller, object>>[]? includeProperties)
    {
        throw new NotImplementedException();
    }

    public Task<IReadOnlyList<Seller>> ListAllAsync(CancellationToken cancellationToken = default)
    {
        throw new NotImplementedException();
    }

    public Task<IReadOnlyList<Seller>> ListAsync(Expression<Func<Seller, bool>>? filter, CancellationToken cancellationToken = default,
        params Expression<Func<Seller, object>>[] includesProperties)
    {
        throw new NotImplementedException();
    }

    public Task AddAsync(Seller entity, CancellationToken cancellationToken = default)
    {
        throw new NotImplementedException();
    }

    public Task UpdateAsync(Seller entity, CancellationToken cancellationToken = default)
    {
        throw new NotImplementedException();
    }

    public Task DeleteAsync(Seller entity, CancellationToken cancellationToken = default)
    {
        throw new NotImplementedException();
    }

    public Task<Seller?> FirstOrDefaultAsync(Expression<Func<Seller, bool>> filter, CancellationToken cancellationToken = default)
    {
        throw new NotImplementedException();
    }
}

