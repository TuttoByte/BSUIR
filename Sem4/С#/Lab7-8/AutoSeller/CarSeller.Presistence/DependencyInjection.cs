using CarSeller.Presistence.Data;
using CarSeller.Presistence.Repository;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;

namespace CarSeller.Presistence;

public static class DependencyInjection
{
    public static IServiceCollection AddPersistence(this IServiceCollection services)
    {
        services.AddSingleton<IUnitOfWork, FakeUnitOfWork>();
        return services;
    }

    public static IServiceCollection AddPersistence(this IServiceCollection services, DbContextOptions options)
    {
        services.AddPersistence().AddSingleton<AppDbContext>(new AppDbContext((DbContextOptions<AppDbContext>)options));
        return services;
    }
}