using System.Linq.Expressions;
using System.Runtime.Serialization;
using CarSeller.Presistence.Data;
using Microsoft.EntityFrameworkCore;

namespace CarSeller.Presistence.Repository;

public class EfRepository <T>:IRepository<T> where T : Entity
{
    
    protected readonly AppDbContext _dbContext;
    protected readonly DbSet<T> _dbSet;

    public EfRepository(AppDbContext dbContext)
    {
        _dbContext = dbContext;
        _dbSet = _dbContext.Set<T>();
    }
    
    public async Task<T?> GetByIdAsync(int id, CancellationToken cancellationToken = default, params Expression<Func<T, object>>[]? includeProperties)
    {
        return await _dbSet.FindAsync([id], cancellationToken);
    }

    public async Task<IReadOnlyList<T>> ListAllAsync(CancellationToken cancellationToken = default)
    {
        return await _dbSet.ToListAsync(cancellationToken);
    }

    public async Task<IReadOnlyList<T>> ListAsync(Expression<Func<T, bool>>? filter, CancellationToken cancellationToken = default,
        params Expression<Func<T, object>>[] includesProperties)
    {
        IQueryable<T>? query = _dbSet.AsQueryable();
        if (includesProperties.Any())
        {
            foreach (Expression<Func<T, object>>? included in includesProperties)
            {
                query = query.Include(included);
            }
        }

        if (filter != null)
        {
            query = query.Where(filter);
        }
        return await query.ToListAsync(cancellationToken: cancellationToken);
    }

    public Task AddAsync(T entity, CancellationToken cancellationToken = default)
    {
        _dbSet.Add(entity);
        _dbContext.Entry(entity).State = EntityState.Added;
        return Task.CompletedTask;
    }

    public Task UpdateAsync(T entity, CancellationToken cancellationToken = default)
    {
        _dbContext.Entry(entity).State = EntityState.Modified;
        return Task.CompletedTask;
    }

    public Task DeleteAsync(T entity, CancellationToken cancellationToken = default)
    {
        _dbSet.Remove(entity);
        _dbContext.Entry(entity).State = EntityState.Deleted;
        return Task.CompletedTask;
    }

    public async Task<T?> FirstOrDefaultAsync(Expression<Func<T, bool>> filter, CancellationToken cancellationToken = default)
    {
        return await _dbSet.FirstOrDefaultAsync(filter, cancellationToken);
    }
}